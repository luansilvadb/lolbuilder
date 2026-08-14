// Package sync orquestra a captura: descobre o patch, baixa as fontes,
// decodifica, filtra, valida e so entao grava.
//
// A ordem importa. Nada e escrito em disco antes de todas as validacoes de
// sanidade passarem — um snapshot vazio gravado por cima de um patch novo seria
// indistinguivel de "a Riot removeu o conteudo", e o erro so apareceria semanas
// depois, no output.
//
// Diferente do Classicorone, o sync nao fala com o cliente de jogo. No modo
// Jade o cliente era a unica fonte que dizia quais runas estavam a venda; no
// LoL moderno nao ha loja de runas, entao a captura e inteiramente offline do
// cliente e nunca bloqueia por ele estar fechado.
package sync

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/cdragon"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/filter"
	"github.com/luansilvadb/lolbuilder/internal/locale"
	"github.com/luansilvadb/lolbuilder/internal/mapdata"
	"github.com/luansilvadb/lolbuilder/internal/model"
	"github.com/luansilvadb/lolbuilder/internal/snapshot"
)

// catalogFiles sao os arquivos de lista do plugin, baixados no locale canonico
// e tambem no de exibicao.
//
// Os dois locales tem schema identico e diferem so no texto. O canonico e a
// espinha estrutural — e nele que a decodificacao estrita roda; o de exibicao
// entra para que o dataset publique nome e efeito em portugues sem perder o
// nome canonico em ingles, que e como o resto do mundo chama a mesma coisa.
var catalogFiles = []string{
	"items.json",
	"perks.json",
	"perkstyles.json",
	"champion-summary.json",
	"summoner-spells.json",
}

// Syncer executa uma captura.
type Syncer struct {
	cfg   *config.Config
	cli   *cdragon.Client
	store *snapshot.Store
	log   io.Writer
}

// New monta o executor a partir da configuracao.
func New(cfg *config.Config, log io.Writer) *Syncer {
	return &Syncer{
		cfg:   cfg,
		cli:   cdragon.New(cfg.HTTPTimeout()),
		store: snapshot.NewStore(cfg.SnapshotsDir),
		log:   log,
	}
}

func (s *Syncer) logf(format string, a ...any) {
	fmt.Fprintf(s.log, format+"\n", a...)
}

// fetched guarda os bytes de uma fonte junto do seu ETag.
type fetched struct {
	name string
	data []byte
	etag string
	from string // "rede" ou "cache"
}

// Result resume o desfecho de uma captura.
type Result struct {
	Patch     string
	Patchline string
	Skipped   bool
	Counts    map[string]int
}

// Run executa a captura completa.
func (s *Syncer) Run() (*Result, error) {
	// 1. O CDragon terminou de processar o patchline?
	st, err := s.cli.FetchStatus(s.cfg.StatusURL())
	if err != nil {
		return nil, fmt.Errorf("consultando status: %w", err)
	}
	if !st.Done() {
		return nil, fmt.Errorf("CDragon ainda processando (estado %q, carimbo %s) — tente mais tarde", st.State, st.Timestamp)
	}
	s.logf("status do CDragon: %s (%s)", st.State, st.Timestamp)

	// 2. Qual patch esta sendo servido?
	patch, patchFull, err := s.cli.FetchPatch(s.cfg.MetadataURL())
	if err != nil {
		return nil, fmt.Errorf("descobrindo patch: %w", err)
	}
	s.logf("patchline %s serve o patch %s (%s)", s.cfg.Patchline, patch, patchFull)

	// Captura retroativa se autoverifica: pedir o patchline 15.20 e receber
	// outro patch significa que o CDragon nao serve aquele patchline, e seguir
	// gravaria a captura sob um nome que nao corresponde ao conteudo.
	if err := s.checkPatchline(patch); err != nil {
		return nil, err
	}

	// 3. Ja capturado? Nao sobrescreve.
	if s.store.Exists(patch) {
		s.logf("snapshot de %s ja existe — nada a fazer", patch)
		return &Result{Patch: patch, Patchline: s.cfg.Patchline, Skipped: true}, nil
	}

	// 4. Captura anterior fornece os ETags para o GET condicional.
	prevPatch, err := s.store.LatestPatch()
	if err != nil {
		return nil, fmt.Errorf("procurando captura anterior: %w", err)
	}
	var prev *snapshot.Capture
	if prevPatch != "" {
		if prev, err = s.store.LoadCapture(prevPatch); err != nil {
			return nil, err
		}
		s.logf("captura anterior: %s", prevPatch)
	} else {
		s.logf("nenhuma captura anterior — esta e a captura inicial")
	}

	all := make([]fetched, 0, 8+4*200)

	// 5. Catalogos, nos dois locales.
	byName := map[string][]byte{}
	for _, f := range catalogFiles {
		got, err := s.fetchOne(f, s.cfg.DataURL(f), prev, prevPatch)
		if err != nil {
			return nil, err
		}
		byName[f] = got.data
		all = append(all, *got)
	}
	byDisplay := map[string][]byte{}
	for _, f := range catalogFiles {
		name := s.displayPath(f)
		got, err := s.fetchOne(name, s.cfg.DisplayURL(f), prev, prevPatch)
		if err != nil {
			return nil, fmt.Errorf("locale de exibicao %q: %w — "+
				"ajuste locale_display em config.json se a Riot deixou de servi-lo",
				s.cfg.LocaleDisplay, err)
		}
		byDisplay[f] = got.data
		all = append(all, *got)
	}

	// 6. Decodificar (estrito) e filtrar.
	items, err := model.DecodeStrict[[]model.Item]("items.json", byName["items.json"])
	if err != nil {
		return nil, err
	}
	runes, err := model.DecodeStrict[[]model.Rune]("perks.json", byName["perks.json"])
	if err != nil {
		return nil, err
	}
	styles, err := model.DecodeStrict[model.PerkStyles]("perkstyles.json", byName["perkstyles.json"])
	if err != nil {
		return nil, err
	}
	summaries, err := model.DecodeStrict[[]model.ChampionSummary]("champion-summary.json", byName["champion-summary.json"])
	if err != nil {
		return nil, err
	}
	spells, err := model.DecodeStrict[[]model.SummonerSpell]("summoner-spells.json", byName["summoner-spells.json"])
	if err != nil {
		return nil, err
	}

	// 6b. O locale de exibicao passa pela mesma decodificacao estrita e e
	// conferido contra o canonico. Sem isso, um pt_br truncado entraria no
	// snapshot em silencio e so apareceria no fim da linha, como nome em branco
	// no arquivo publicado.
	if err := s.checkDisplayCatalogs(byDisplay, items, runes, styles, spells, summaries); err != nil {
		return nil, err
	}

	// 7. Dump do mapa — a unica fonte que diz o que a loja do modo referencia.
	mapName := mapFileName(s.cfg.Mode.MapID)
	gotMap, err := s.fetchOne(mapName, s.cfg.MapDataURL(), prev, prevPatch)
	if err != nil {
		return nil, fmt.Errorf("dump do mapa %d: %w — "+
			"veja map_data_path e mode.map_id em config.json", s.cfg.Mode.MapID, err)
	}
	all = append(all, *gotMap)

	shopIDs, err := mapdata.ShopItemIDs(gotMap.data, s.cfg.Mode.MapID, s.cfg.Mode.GameMode)
	if err != nil {
		return nil, err
	}

	modeItems := filter.ItemsInRange(items, s.cfg.Mode)
	modeChamps := filter.ChampionsInRange(summaries, s.cfg.Mode)
	modeSpells := filter.SummonerSpellsForMode(spells, s.cfg.Mode)

	// A loja referencia itens do mesmo arquivo. Referencia orfa significa que a
	// loja veio de um mapa cujo catalogo nao e este, e isso publica item
	// fantasma se passar calado.
	if orfaos := filter.UnknownShopIDs(shopIDs, filter.ItemsByID(items)); len(orfaos) > 0 {
		return nil, fmt.Errorf(
			"a loja do modo referencia %d id(s) que nao existem em items.json: %v — "+
				"veja map_data_path e mode.map_id em config.json", len(orfaos), orfaos)
	}
	// Fora da faixa nao e defeito: as listas do mapa misturam o catalogo com
	// buffs de torre e marcadores internos. Fica no log para que um aumento
	// apareca em vez de sumir; a triagem e do modelo canonico.
	if fora := filter.ShopIDsOutOfRange(shopIDs, s.cfg.Mode); len(fora) > 0 {
		s.logf("  nota: %d id(s) da loja caem fora da faixa do modo e serao triados no build: %v",
			len(fora), fora)
	}

	// 8. Detalhe de cada campeao do modo, nos dois locales.
	sort.Slice(modeChamps, func(i, j int) bool { return modeChamps[i].ID < modeChamps[j].ID })
	for _, c := range modeChamps {
		name := fmt.Sprintf("champions/%d.json", c.ID)
		got, err := s.fetchOne(name, s.cfg.DataURL(name), prev, prevPatch)
		if err != nil {
			return nil, err
		}
		champ, err := model.DecodeStrict[model.Champion](name, got.data)
		if err != nil {
			return nil, err
		}
		all = append(all, *got)

		display := s.displayPath(name)
		gotDisplay, err := s.fetchOne(display, s.cfg.DisplayURL(name), prev, prevPatch)
		if err != nil {
			return nil, err
		}
		champDisplay, err := model.DecodeStrict[model.Champion](display, gotDisplay.data)
		if err != nil {
			return nil, err
		}
		if err := checkChampionDisplay(display, champ, champDisplay); err != nil {
			return nil, err
		}
		all = append(all, *gotDisplay)
	}

	// 9. Dump de dados do jogo de cada campeao — a unica fonte com estatistica
	// base e formula de habilidade. Nao passa por DecodeStrict: o conjunto de
	// chaves varia por campeao e inclui identificadores opacos, entao a
	// vigilancia dessa fonte e a cobertura de extracao, feita no build.
	//
	// Tambem nao tem locale: o dump nao e traduzido.
	for _, c := range modeChamps {
		name := "characters/" + config.GameDataFile(c.Alias)
		url := s.cfg.GameDataURL(config.GameDataFile(c.Alias))
		got, err := s.fetchOne(name, url, prev, prevPatch)
		if err != nil {
			return nil, fmt.Errorf("dump de %s (%s): %w — o alias mudou ou "+
				"o caminho do dump saiu do ar (veja game_data_path em config.json)",
				c.Name, url, err)
		}
		all = append(all, *got)
	}

	counted := s.countEntities(decoded{
		modeItems: len(modeItems),
		shopIDs:   len(shopIDs),
		runes:     len(runes),
		styles:    len(styles.Styles),
		champions: len(modeChamps),
		spells:    len(modeSpells),
	})
	counts := countMap(counted)

	// 10. Validar ANTES de escrever qualquer coisa.
	if err := s.validate(counted); err != nil {
		return nil, err
	}
	s.logf("contagens: %s", formatCounts(counts))

	// 11. Escrever.
	w, err := s.store.BeginWrite(patch)
	if err != nil {
		return nil, err
	}
	defer w.Abort()

	for _, f := range all {
		if err := w.Add(f.name, f.data, f.etag); err != nil {
			return nil, err
		}
	}
	if err := w.Commit(patchFull, s.cfg.DataURL(""), s.cfg.Patchline, counts); err != nil {
		return nil, err
	}
	s.logf("snapshot gravado em %s", s.store.PatchDir(patch))

	return &Result{Patch: patch, Patchline: s.cfg.Patchline, Counts: counts}, nil
}

// checkDisplayCatalogs decodifica os catalogos do locale de exibicao com o
// mesmo rigor do canonico e confere que os dois descrevem as mesmas entidades.
//
// A decodificacao estrita aqui nao e redundante com a do canonico: os dois
// arquivos sao servidos por caminhos diferentes e podem divergir de forma —
// foi assim que o dataset original descobriu que uma fonte tinha mudado.
func (s *Syncer) checkDisplayCatalogs(
	byDisplay map[string][]byte,
	items []model.Item,
	runes []model.Rune,
	styles model.PerkStyles,
	spells []model.SummonerSpell,
	summaries []model.ChampionSummary,
) error {
	itemsD, err := model.DecodeStrict[[]model.Item](s.displayPath("items.json"), byDisplay["items.json"])
	if err != nil {
		return err
	}
	runesD, err := model.DecodeStrict[[]model.Rune](s.displayPath("perks.json"), byDisplay["perks.json"])
	if err != nil {
		return err
	}
	stylesD, err := model.DecodeStrict[model.PerkStyles](s.displayPath("perkstyles.json"), byDisplay["perkstyles.json"])
	if err != nil {
		return err
	}
	spellsD, err := model.DecodeStrict[[]model.SummonerSpell](s.displayPath("summoner-spells.json"), byDisplay["summoner-spells.json"])
	if err != nil {
		return err
	}
	summariesD, err := model.DecodeStrict[[]model.ChampionSummary](s.displayPath("champion-summary.json"), byDisplay["champion-summary.json"])
	if err != nil {
		return err
	}

	checks := []struct {
		kind  string
		file  string
		canon locale.Names
		disp  locale.Names
	}{
		{"itens", s.displayPath("items.json"), itemNames(items), itemNames(itemsD)},
		{"runas", s.displayPath("perks.json"), runeNames(runes), runeNames(runesD)},
		{"estilos de runa", s.displayPath("perkstyles.json"), styleNames(styles), styleNames(stylesD)},
		{"feiticos", s.displayPath("summoner-spells.json"), spellNames(spells), spellNames(spellsD)},
		{"campeoes", s.displayPath("champion-summary.json"), summaryNames(summaries), summaryNames(summariesD)},
	}
	for _, c := range checks {
		if err := locale.Check(c.kind, c.file, c.canon, c.disp); err != nil {
			return err
		}
	}
	return nil
}

// checkChampionDisplay confere que o detalhe traduzido descreve o mesmo
// campeao que o canonico.
//
// A contagem de habilidades entra na conferencia porque e o sintoma de um
// arquivo truncado que a igualdade de id nao pega: o campeao continua sendo o
// mesmo, e metade do texto some.
func checkChampionDisplay(file string, canon, disp model.Champion) error {
	if disp.ID != canon.ID {
		return fmt.Errorf("%s descreve o campeao %d, esperado %d", file, disp.ID, canon.ID)
	}
	if strings.TrimSpace(disp.Name) == "" {
		return fmt.Errorf("%s nao tem nome — traducao ausente publica campo em branco", file)
	}
	if len(disp.Spells) != len(canon.Spells) {
		return fmt.Errorf("%s tem %d habilidade(s), contra %d no locale canonico — arquivo truncado",
			file, len(disp.Spells), len(canon.Spells))
	}
	return nil
}

func itemNames(items []model.Item) locale.Names {
	out := make(locale.Names, len(items))
	for _, it := range items {
		out[int64(it.ID)] = it.Name
	}
	return out
}

func runeNames(runes []model.Rune) locale.Names {
	out := make(locale.Names, len(runes))
	for _, r := range runes {
		out[int64(r.ID)] = r.Name
	}
	return out
}

func styleNames(s model.PerkStyles) locale.Names {
	out := make(locale.Names, len(s.Styles))
	for _, st := range s.Styles {
		out[int64(st.ID)] = st.Name
	}
	return out
}

// spellNames descarta a sentinela sem id, que a fonte repete tres vezes com o
// maximo de um uint32 e sem modo algum.
func spellNames(spells []model.SummonerSpell) locale.Names {
	out := make(locale.Names, len(spells))
	for _, sp := range spells {
		if sp.ID > math.MaxInt32 {
			continue
		}
		out[sp.ID] = sp.Name
	}
	return out
}

func summaryNames(cs []model.ChampionSummary) locale.Names {
	out := make(locale.Names, len(cs))
	for _, c := range cs {
		out[int64(c.ID)] = c.Name
	}
	return out
}

// checkPatchline confere que um patchline versionado serve o patch que promete.
func (s *Syncer) checkPatchline(patch string) error {
	if s.cfg.Patchline == "latest" || s.cfg.Patchline == "" {
		return nil
	}
	if strings.HasPrefix(patch+".", s.cfg.Patchline+".") || patch == s.cfg.Patchline {
		return nil
	}
	return fmt.Errorf(
		"patchline %q serve o patch %s — o CDragon nao tem esse patchline versionado, "+
			"e gravar seguiria sob um nome que nao corresponde ao conteudo",
		s.cfg.Patchline, patch)
}

// displayPath e onde o arquivo do locale de exibicao mora dentro do snapshot.
func (s *Syncer) displayPath(file string) string {
	return s.cfg.LocaleDisplay + "/" + file
}

// mapFileName e o nome do dump do mapa dentro do snapshot — achatado, sem o
// diretorio do jogo, porque so ha um mapa por captura.
func mapFileName(mapID int32) string {
	return fmt.Sprintf("map%d.bin.json", mapID)
}

// fetchOne busca um arquivo com GET condicional. Em 304, reaproveita os bytes
// da captura anterior.
//
// name e o caminho dentro do snapshot; url e de onde buscar. Os dois sao
// separados porque o dump de dados do jogo e o do mapa vivem em prefixos
// diferentes do plugin, mas precisam do mesmo tratamento de ETag e 304.
func (s *Syncer) fetchOne(name, url string, prev *snapshot.Capture, prevPatch string) (*fetched, error) {
	var etag string
	if prev != nil {
		if meta, ok := prev.Files[name]; ok {
			etag = meta.ETag
		}
	}

	res, err := s.cli.Fetch(url, etag)
	if err != nil {
		return nil, err
	}
	if res.NotModified {
		data, err := s.store.ReadFile(prevPatch, name)
		if err != nil {
			return nil, fmt.Errorf("304 em %s mas copia anterior ilegivel: %w", name, err)
		}
		s.logf("  %-44s 304 (reaproveitado de %s)", name, prevPatch)
		return &fetched{name: name, data: data, etag: etag, from: "cache"}, nil
	}
	s.logf("  %-44s %d bytes", name, len(res.Body))
	return &fetched{name: name, data: res.Body, etag: res.ETag, from: "rede"}, nil
}

// decoded sao as contagens ja filtradas de cada fonte, antes de virarem
// entidades. Existe para que countEntities seja testavel sem rede.
type decoded struct {
	modeItems int // items.json, recortado pela faixa de id
	shopIDs   int // map<N>.bin.json, listas do modo
	runes     int // perks.json
	styles    int // perkstyles.json
	champions int // champion-summary.json, recortado pela faixa de id
	spells    int // summoner-spells.json, recortado por gameModes
}

// entityCount liga uma entidade contada ao arquivo que a alimenta e ao minimo
// que a protege.
//
// A fonte e um campo, e nao um comentario, porque o defeito que esta estrutura
// existe para impedir e trocar a fonte de uma entidade sem trocar o minimo que
// a protege. Com a fonte no dado, trocar de fonte quebra um teste.
type entityCount struct {
	key    string
	source string
	n      int
	min    int
}

// countEntities monta as contagens a partir das fontes decodificadas.
func (s *Syncer) countEntities(d decoded) []entityCount {
	m := s.cfg.Minimums
	return []entityCount{
		{"item_catalog", "items.json", d.modeItems, m.ItemCatalog},
		{"item_shop", mapFileName(s.cfg.Mode.MapID), d.shopIDs, m.ItemShop},
		{"runes", "perks.json", d.runes, m.Runes},
		{"perk_styles", "perkstyles.json", d.styles, m.PerkStyles},
		{"champions", "champion-summary.json", d.champions, m.Champions},
		{"summoner_spells", "summoner-spells.json", d.spells, m.SummonerSpells},
	}
}

// countMap reduz as contagens ao formato gravado no capture.json.
func countMap(cs []entityCount) map[string]int {
	out := make(map[string]int, len(cs))
	for _, c := range cs {
		out[c.key] = c.n
	}
	return out
}

// validate aborta se qualquer entidade ficar abaixo do minimo configurado.
func (s *Syncer) validate(cs []entityCount) error {
	sorted := append([]entityCount(nil), cs...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].key < sorted[j].key })

	for _, c := range sorted {
		if c.min <= 0 {
			return fmt.Errorf(
				"minimo de %s nao declarado — uma entidade sem minimo passa por "+
					"qualquer contagem, inclusive zero. Declare em minimums no config.json",
				c.key)
		}
		if c.n < c.min {
			return fmt.Errorf(
				"validacao de sanidade falhou: %s tem %d, minimo %d, na fonte %s — "+
					"os filtros do modo provavelmente mudaram (veja mode.* em config.json). "+
					"Nada foi escrito; a captura anterior segue intacta",
				c.key, c.n, c.min, c.source)
		}
	}
	return nil
}

func formatCounts(c map[string]int) string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := ""
	for i, k := range keys {
		if i > 0 {
			out += "  "
		}
		out += fmt.Sprintf("%s=%d", k, c[k])
	}
	return out
}
