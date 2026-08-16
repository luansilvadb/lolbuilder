package sync

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/snapshot"
)

// As fontes abaixo sao minimas de proposito: o que este teste exercita e a
// ORQUESTRACAO — ordem das etapas, GET condicional, validacao antes de escrever,
// e a garantia de que nada e gravado quando algo falha. O conteudo de cada
// arquivo ja e coberto pelos testes do canonical.

const (
	fixItems = `[{"id":1001,"name":"Boots","description":"","active":false,"inStore":true,
		"from":[],"to":[],"categories":["Boots"],"maxStacks":1,"requiredChampion":"",
		"requiredAlly":"","requiredBuffCurrencyName":"","requiredBuffCurrencyCost":0,
		"specialRecipe":0,"isEnchantment":false,"price":300,"priceTotal":300,
		"displayInItemSets":true,"iconPath":"a.png"},
		{"id":3006,"name":"Berserker's Greaves","description":"","active":false,"inStore":true,
		"from":[1001],"to":[],"categories":["Boots"],"maxStacks":1,"requiredChampion":"",
		"requiredAlly":"","requiredBuffCurrencyName":"","requiredBuffCurrencyCost":0,
		"specialRecipe":0,"isEnchantment":false,"price":800,"priceTotal":1100,
		"displayInItemSets":true,"iconPath":"b.png"}]`

	fixPerks = `[{"id":8005,"name":"Press the Attack","majorChangePatchVersion":"14.10",
		"tooltip":"","shortDesc":"","longDesc":"","recommendationDescriptor":"",
		"iconPath":"p.png","endOfGameStatDescs":[],"recommendationDescriptorAttributes":{}}]`

	fixStyles = `{"schemaVersion":2,"styles":[{"id":8000,"name":"Precision","tooltip":"",
		"iconPath":"i.png","assetMap":{},"isAdvanced":false,"allowedSubStyles":[],
		"subStyleBonus":[],"slots":[{"type":"kKeyStone","slotLabel":"","perks":[8005]}],
		"defaultPageName":"p","defaultSubStyle":0,"defaultPerks":[],
		"defaultPerksWhenSplashed":[],"defaultStatModsPerSubStyle":[]}]}`

	fixSummary = `[{"id":-1,"name":"None","description":"","alias":"","contentId":"",
		"squarePortraitPath":"","roles":[]},
		{"id":86,"name":"Garen","description":"d","alias":"Garen","contentId":"c",
		 "squarePortraitPath":"s.png","roles":["fighter"]}]`

	fixSpells = `[{"id":4,"name":"Flash","description":"d","summonerLevel":7,
		"cooldown":300,"gameModes":["CLASSIC"],"iconPath":"f.png"}]`

	fixMapa = `{"Maps/Shipping/Map11/Modes/CLASSIC":{"__type":"GameModeMapData",
		"mModeName":"CLASSIC","itemLists":["{a}"]},
		"{a}":{"__type":"GameModeItemList","mItems":["Items/1001","Items/3006"]}}`

	fixChampion = `{"id":86,"contentId":"c","name":"Garen","alias":"Garen","title":"t",
		"shortBio":"b","isVisibleInClient":true,
		"tacticalInfo":{"style":10,"difficulty":1,"damageType":"kPhysical","attackType":"melee"},
		"playstyleInfo":{"damage":3,"durability":3,"crowdControl":1,"mobility":1,"utility":1},
		"championTagInfo":{"championTagPrimary":"Juggernaut","championTagSecondary":""},
		"squarePortraitPath":"s.png","stingerSfxPath":"","chooseVoPath":"","banVoPath":"",
		"roles":["fighter"],"recommendedItemDefaults":[],"skins":[],
		"passive":{"name":"P","abilityIconPath":"","abilityVideoPath":"",
		 "abilityVideoImagePath":"","description":"d"},
		"spells":[]}`

	fixDump = `{"Characters/Garen/CharacterRecords/Root":{"__type":"CharacterRecord",
		"mCharacterName":"Garen","spells":[],"spellNames":[],"mCharacterPassiveSpell":""}}`

	// Dois itens no mesmo grupo, para que ele restrinja de fato — grupo com um
	// membro so seria descartado por nao poder ser violado.
	fixItensBin = `{
		"Items/ItemGroups/Default":{"__type":"ItemGroup","mItemGroupID":"Default","mMaxGroupOwnable":1},
		"Items/ItemGroups/Boots":{"__type":"ItemGroup","mItemGroupID":"Boots","mMaxGroupOwnable":1},
		"Items/1001":{"__type":"ItemData","itemID":1001,
		 "mItemGroups":["Items/ItemGroups/Default","Items/ItemGroups/Boots"]},
		"Items/3006":{"__type":"ItemData","itemID":3006,
		 "mItemGroups":["Items/ItemGroups/Default","Items/ItemGroups/Boots"]}}`
)

// servidorDeFontes responde a todas as URLs que o sync busca.
func servidorDeFontes(t *testing.T, falhaEm string) *httptest.Server {
	t.Helper()
	corpo := map[string]string{
		"items.json":               fixItems,
		"perks.json":               fixPerks,
		"perkstyles.json":          fixStyles,
		"champion-summary.json":    fixSummary,
		"summoner-spells.json":     fixSpells,
		"champions/86.json":        fixChampion,
		"map11/map11.bin.json":     fixMapa,
		"garen/garen.bin.json":     fixDump,
		"game/items.cdtb.bin.json": fixItensBin,
		"content-metadata.json":    `{"version":"16.99.1+branch.releases-16-99.content.release"}`,
		"json/status.live.txt":     "2026-08-15T00:00:00Z done",
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if falhaEm != "" && strings.Contains(r.URL.Path, falhaEm) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for sufixo, body := range corpo {
			if strings.HasSuffix(r.URL.Path, sufixo) {
				w.Header().Set("ETag", `"etag-`+sufixo+`"`)
				w.Write([]byte(body))
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}))
}

func cfgDeTeste(t *testing.T, base string) *config.Config {
	t.Helper()
	return &config.Config{
		BaseURL: base, Patchline: "latest",
		PluginPath: "plugins/p", GameDataPath: "game/data/characters",
		MapDataPath:     "game/data/maps/shipping",
		ItemDataPath:    "game/items.cdtb.bin.json",
		LocaleCanonical: "default", LocaleDisplay: "pt_br",
		Mode: config.Mode{
			Name: "classic", GameMode: "CLASSIC", MapID: 11,
			ItemIDMin: 1000, ItemIDMax: 699999,
			ChampionIDMin: 1, ChampionIDMax: 999,
		},
		Minimums: config.Minimums{
			ItemCatalog: 1, ItemShop: 1, Runes: 1,
			PerkStyles: 1, Champions: 1, SummonerSpells: 1, ItemGroups: 1,
		},
		SnapshotsDir:       t.TempDir(),
		HTTPTimeoutSeconds: 5,
	}
}

// TestRunGravaOSnapshotCompleto exercita a orquestracao inteira.
func TestRunGravaOSnapshotCompleto(t *testing.T) {
	srv := servidorDeFontes(t, "")
	cfg := cfgDeTeste(t, srv.URL)

	res, err := New(cfg, io.Discard).Run()
	if err != nil {
		t.Fatal(err)
	}
	if res.Patch != "16.99" || res.Skipped {
		t.Fatalf("resultado = %+v", res)
	}

	dir := filepath.Join(cfg.SnapshotsDir, "16.99")
	esperados := []string{
		"capture.json", "items.json", "perks.json", "perkstyles.json",
		"champion-summary.json", "summoner-spells.json", "map11.bin.json",
		"items.bin.json",
		"champions/86.json", "pt_br/items.json", "pt_br/champions/86.json",
		"characters/garen/garen.bin.json",
	}
	for _, f := range esperados {
		if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(f))); err != nil {
			t.Errorf("%s nao foi gravado", f)
		}
	}

	cap, err := snapshot.NewStore(cfg.SnapshotsDir).LoadCapture("16.99")
	if err != nil || cap == nil {
		t.Fatalf("capture.json = %v (%v)", cap, err)
	}
	if cap.Counts["champions"] != 1 || cap.Counts["item_shop"] != 2 {
		t.Errorf("contagens = %+v", cap.Counts)
	}
	// O grupo Botas tem os dois itens da loja, entao restringe de fato.
	if cap.Counts["item_groups"] != 1 {
		t.Errorf("grupos de exclusividade nao contados: %+v", cap.Counts)
	}
	if cap.Patchline != "latest" {
		t.Errorf("procedencia perdida: %q", cap.Patchline)
	}
}

// TestRunEIdempotente: patch ja capturado, nada acontece. O snapshot gravado
// nunca e reescrito.
func TestRunEIdempotente(t *testing.T) {
	srv := servidorDeFontes(t, "")
	cfg := cfgDeTeste(t, srv.URL)

	if _, err := New(cfg, io.Discard).Run(); err != nil {
		t.Fatal(err)
	}
	res, err := New(cfg, io.Discard).Run()
	if err != nil {
		t.Fatal(err)
	}
	if !res.Skipped {
		t.Fatal("a segunda captura do mesmo patch nao foi pulada")
	}
}

// TestRunNaoEscreveNadaQuandoUmaFonteFalha e a garantia central do sync: nada
// vai a disco antes de todas as validacoes passarem, e um aborto no meio deixa
// a captura anterior intacta.
func TestRunNaoEscreveNadaQuandoUmaFonteFalha(t *testing.T) {
	for _, alvo := range []string{"perkstyles.json", "map11", "garen", "champions/86"} {
		t.Run(alvo, func(t *testing.T) {
			srv := servidorDeFontes(t, alvo)
			cfg := cfgDeTeste(t, srv.URL)

			if _, err := New(cfg, io.Discard).Run(); err == nil {
				t.Fatal("a falha na fonte passou em silencio")
			}
			entries, err := os.ReadDir(cfg.SnapshotsDir)
			if err != nil {
				t.Fatal(err)
			}
			for _, e := range entries {
				t.Errorf("a captura abortada deixou %s em disco", e.Name())
			}
		})
	}
}

// TestRunAbortaAbaixoDoMinimo: o snapshot nao pode ser gravado com contagem
// abaixo do piso, porque um snapshot vazio e indistinguivel de "a Riot removeu
// o conteudo" e o erro so apareceria semanas depois.
func TestRunAbortaAbaixoDoMinimo(t *testing.T) {
	srv := servidorDeFontes(t, "")
	cfg := cfgDeTeste(t, srv.URL)
	cfg.Minimums.Champions = 500

	_, err := New(cfg, io.Discard).Run()
	if err == nil {
		t.Fatal("contagem abaixo do minimo passou")
	}
	if !strings.Contains(err.Error(), "Nada foi escrito") {
		t.Fatalf("o erro nao diz que a captura anterior segue intacta: %v", err)
	}
	entries, _ := os.ReadDir(cfg.SnapshotsDir)
	if len(entries) != 0 {
		t.Errorf("a captura abortada deixou %d entrada(s) em disco", len(entries))
	}
}

// TestRunAbortaComPatchlineQueNaoConfere: pedir 15.20 e receber outro patch
// significa que o CDragon nao o tem, e gravar poria o conteudo sob um nome
// errado.
func TestRunAbortaComPatchlineQueNaoConfere(t *testing.T) {
	srv := servidorDeFontes(t, "")
	cfg := cfgDeTeste(t, srv.URL)
	cfg.Patchline = "15.20"

	if _, err := New(cfg, io.Discard).Run(); err == nil {
		t.Fatal("patchline que serve outro patch passou")
	}
}

// TestRunAbortaComCDragonProcessando: baixar durante o processamento pode render
// arquivos incompletos, que e pior que nao baixar.
func TestRunAbortaComCDragonProcessando(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("2026-08-15T00:00:00Z running"))
	}))
	defer srv.Close()

	if _, err := New(cfgDeTeste(t, srv.URL), io.Discard).Run(); err == nil {
		t.Fatal("captura durante o processamento do CDragon passou")
	}
}

// TestRunConfereOLocaleDeExibicao: um pt_br truncado entraria no snapshot em
// silencio e so apareceria no export, como nome em branco.
func TestRunConfereOLocaleDeExibicao(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// O locale de exibicao devolve um catalogo de itens VAZIO.
		if strings.Contains(r.URL.Path, "/pt_br/") && strings.HasSuffix(r.URL.Path, "items.json") {
			w.Write([]byte(`[]`))
			return
		}
		servidorDeFontesHandler(w, r)
	}))
	defer srv.Close()

	_, err := New(cfgDeTeste(t, srv.URL), io.Discard).Run()
	if err == nil {
		t.Fatal("locale de exibicao vazio passou")
	}
	if !strings.Contains(err.Error(), "pt_br/items.json") {
		t.Fatalf("o erro nao aponta o arquivo: %v", err)
	}
}

// servidorDeFontesHandler e a mesma resposta do servidor completo, reaproveitada
// pelo teste que sobrescreve so um caminho.
func servidorDeFontesHandler(w http.ResponseWriter, r *http.Request) {
	corpo := map[string]string{
		"items.json":               fixItems,
		"perks.json":               fixPerks,
		"perkstyles.json":          fixStyles,
		"champion-summary.json":    fixSummary,
		"summoner-spells.json":     fixSpells,
		"champions/86.json":        fixChampion,
		"map11/map11.bin.json":     fixMapa,
		"garen/garen.bin.json":     fixDump,
		"game/items.cdtb.bin.json": fixItensBin,
		"content-metadata.json":    `{"version":"16.99.1+x"}`,
		"json/status.live.txt":     "2026-08-15T00:00:00Z done",
	}
	for sufixo, body := range corpo {
		if strings.HasSuffix(r.URL.Path, sufixo) {
			w.Write([]byte(body))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
