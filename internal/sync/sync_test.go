package sync

import (
	"io"
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/config"
)

func syncer(min config.Minimums, patchline string) *Syncer {
	cfg := &config.Config{
		Patchline: patchline,
		Mode:      config.Mode{Name: "classic", GameMode: "CLASSIC", MapID: 11},
		Minimums:  min,
	}
	return New(cfg, io.Discard)
}

var minimosOK = config.Minimums{
	ItemCatalog: 500, ItemShop: 200, Runes: 80,
	PerkStyles: 5, Champions: 150, SummonerSpells: 7,
}

// medidoNo1616 sao as contagens reais do patch 16.16. Servem de referencia
// viva: se os minimos subirem acima delas, este teste avisa antes do sync.
var medidoNo1616 = decoded{
	modeItems: 705, shopIDs: 266, runes: 103,
	styles: 5, champions: 173, spells: 9,
}

func TestValidateAceitaOMedidoNoPatchDeReferencia(t *testing.T) {
	s := syncer(minimosOK, "latest")
	if err := s.validate(s.countEntities(medidoNo1616)); err != nil {
		t.Fatalf("as contagens medidas no 16.16 nao passam nos proprios minimos: %v", err)
	}
}

func TestValidateAbortaAbaixoDoMinimo(t *testing.T) {
	s := syncer(minimosOK, "latest")
	d := medidoNo1616
	d.runes = 12

	err := s.validate(s.countEntities(d))
	if err == nil {
		t.Fatal("catalogo de runas esvaziado passou na validacao")
	}
	if !strings.Contains(err.Error(), "runes") || !strings.Contains(err.Error(), "perks.json") {
		t.Fatalf("o erro nao nomeia a entidade e a fonte: %v", err)
	}
	if !strings.Contains(err.Error(), "Nada foi escrito") {
		t.Fatalf("o erro nao diz que a captura anterior segue intacta: %v", err)
	}
}

// TestValidateAbortaSemMinimoDeclarado: entidade sem minimo passa por qualquer
// contagem, inclusive zero — que e pior do que nao ter a entidade.
func TestValidateAbortaSemMinimoDeclarado(t *testing.T) {
	min := minimosOK
	min.PerkStyles = 0
	s := syncer(min, "latest")

	err := s.validate(s.countEntities(medidoNo1616))
	if err == nil {
		t.Fatal("minimo nao declarado passou")
	}
	if !strings.Contains(err.Error(), "perk_styles") {
		t.Fatalf("o erro nao nomeia a entidade sem minimo: %v", err)
	}
}

// TestCountEntitiesLigaEntidadeAFonte e o teste que impede o defeito de trocar
// a fonte de uma entidade sem trocar o minimo que a protege.
func TestCountEntitiesLigaEntidadeAFonte(t *testing.T) {
	s := syncer(minimosOK, "latest")
	want := map[string]string{
		"item_catalog":    "items.json",
		"item_shop":       "map11.bin.json",
		"runes":           "perks.json",
		"perk_styles":     "perkstyles.json",
		"champions":       "champion-summary.json",
		"summoner_spells": "summoner-spells.json",
	}
	got := s.countEntities(medidoNo1616)
	if len(got) != len(want) {
		t.Fatalf("countEntities devolveu %d entidades, esperado %d", len(got), len(want))
	}
	for _, c := range got {
		w, ok := want[c.key]
		if !ok {
			t.Errorf("entidade inesperada: %q", c.key)
			continue
		}
		if c.source != w {
			t.Errorf("%s vem de %q, esperado %q", c.key, c.source, w)
		}
	}
}

func TestCountMapEFormatCounts(t *testing.T) {
	s := syncer(minimosOK, "latest")
	m := countMap(s.countEntities(medidoNo1616))
	if m["champions"] != 173 || m["item_shop"] != 266 {
		t.Fatalf("countMap perdeu contagem: %+v", m)
	}
	got := formatCounts(m)
	if !strings.HasPrefix(got, "champions=173") {
		t.Fatalf("formatCounts nao esta ordenado: %q", got)
	}
}

// TestCheckPatchline: pedir um patchline versionado e receber outro patch
// significa que o CDragon nao o tem, e gravar poria o conteudo sob um nome
// errado.
func TestCheckPatchline(t *testing.T) {
	casos := []struct {
		patchline string
		patch     string
		querErro  bool
	}{
		{"latest", "16.16", false},
		{"", "16.16", false},
		{"15.20", "15.20", false},
		{"15.20", "16.16", true},
		{"15.2", "15.20", true},
	}
	for _, c := range casos {
		err := syncer(minimosOK, c.patchline).checkPatchline(c.patch)
		if c.querErro && err == nil {
			t.Errorf("patchline %q servindo %q deveria falhar", c.patchline, c.patch)
		}
		if !c.querErro && err != nil {
			t.Errorf("patchline %q servindo %q falhou: %v", c.patchline, c.patch, err)
		}
	}
}

func TestDisplayPathEMapFileName(t *testing.T) {
	cfg := &config.Config{LocaleDisplay: "pt_br"}
	s := New(cfg, io.Discard)
	if got := s.displayPath("champions/1.json"); got != "pt_br/champions/1.json" {
		t.Fatalf("displayPath = %q", got)
	}
	if got := mapFileName(11); got != "map11.bin.json" {
		t.Fatalf("mapFileName = %q", got)
	}
}
