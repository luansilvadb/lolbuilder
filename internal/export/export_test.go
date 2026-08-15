package export

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/canonical"
	"github.com/luansilvadb/lolbuilder/internal/gamedata"
)

func datasetDeTeste() *canonical.Dataset {
	return &canonical.Dataset{
		Patch: "16.99",
		Items: []canonical.Item{
			{ID: 3006, Nome: "Grevas do Berserker", NomeCanonico: "Berserkers Greaves",
				Custo: 1100, Combina: 800, Componentes: []int32{1001, 1042},
				Stats:  canon.Vector{canon.AttackSpeedPct: 30, canon.MoveSpeed: 45},
				Efeito: "Anda depressa.", Compravel: true, Botas: true},
			{ID: 1515, Nome: "Blindagem", NomeCanonico: "Turret Plating", Compravel: false},
		},
		Runes: []canonical.Rune{
			{ID: 5008, Nome: "Forca Adaptativa", NomeCanonico: "Adaptive Force",
				TipoSlot: canonical.SlotStatMod, Escopo: "sum",
				StatsDaRuna: canon.Vector{canon.AdaptiveForce: 9}},
			{ID: 8005, Nome: "Pressione o Ataque", NomeCanonico: "Press the Attack",
				TipoSlot: canonical.SlotKeyStone, Estilo: 8000, Escopo: "out_of_scope",
				MotivoDoEscopo: "depende de estado de combate"},
			{ID: 9999, Nome: "Removida", NomeCanonico: "Removed"},
		},
		RuneStyles: []canonical.RuneStyle{{ID: 8000, Nome: "Precisao", NomeCanonico: "Precision"}},
		SummonerSpells: []canonical.SummonerSpell{
			{ID: 4, Nome: "Flash", NomeCanonico: "Flash", Recarga: 300, NivelMinimo: 7,
				Descricao: "Teleporta"},
		},
		Champions: []canonical.Champion{{
			ID: 86, Nome: "Garen", NomeCanonico: "Garen", Titulo: "o Poder de Demacia",
			Papeis: []string{"fighter"}, CorpoACorpo: true,
			Stats: &gamedata.Stats{
				HP:           gamedata.Scaling{Base: 690, PerLevel: 98},
				AttackDamage: gamedata.Scaling{Base: 69, PerLevel: 4.199999809265137},
				Melee:        true, MoveSpeed: 340, AttackRange: 175,
			},
			Passiva: canonical.Habilidade{Slot: "p", Nome: "Perseveranca", Descricao: "Regenera."},
			Habilidades: []canonical.Habilidade{{
				Slot: "q", Nome: "Acerto Decisivo",
				Descricao: "Causa @TotalDamage@ e silencia por @SilenceDuration@s, com @Perdido@ de sobra.",
				Recarga:   []float64{8, 8, 8, 8, 8},
				Efeitos: []canonical.Efeito{{
					Nome: "TotalDamage",
					PorRank: []canonical.Expressao{
						{Fixo: 30, Escalas: []canonical.Escala{{Stat: "attack_damage", Parcela: "total", Coeficiente: 1.5}}},
						{Fixo: 60}, {Fixo: 90}, {Fixo: 120}, {Fixo: 150},
					},
				}},
				TodasAsSeries: []canonical.SerieNomeada{
					{Nome: "SilenceDuration", PorRank: []float64{1.5, 1.5, 1.5, 1.5, 1.5}},
				},
			}},
		}},
	}
}

func gerar(t *testing.T) *Result {
	t.Helper()
	res, err := Generate(Input{Dataset: datasetDeTeste(), CapturadoEm: "2026-08-14T00:00:00Z"})
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func corpo(t *testing.T, res *Result, nome string) string {
	t.Helper()
	for _, f := range res.Files {
		if f.Name == nome {
			return f.Body
		}
	}
	t.Fatalf("%s nao foi gerado", nome)
	return ""
}

// TestNomesDosArquivosSaoFixos: o upload ao Project e manual, e nome que muda
// entre patches deixa orfao la dentro.
func TestNomesDosArquivosSaoFixos(t *testing.T) {
	res := gerar(t)
	want := []string{
		"00-MANIFEST.md", "01-items.md", "02-runes.md", "03-summoner-spells.md",
		"04-champions.md", "05-computed.md", "06-champion-stats.md",
	}
	if len(res.Files) != len(want) {
		t.Fatalf("gerou %d arquivos, esperado %d", len(res.Files), len(want))
	}
	for i, n := range want {
		if res.Files[i].Name != n {
			t.Errorf("arquivo %d = %q, esperado %q", i, res.Files[i].Name, n)
		}
	}
	if res.Total <= 0 {
		t.Error("o total de tokens ficou em zero")
	}
}

// TestSoOsCompraveisSaoPublicados: item fora da loja nao chega ao consumidor.
func TestSoOsCompraveisSaoPublicados(t *testing.T) {
	itens := corpo(t, gerar(t), "01-items.md")
	if !strings.Contains(itens, "Grevas do Berserker") {
		t.Error("item compravel nao foi publicado")
	}
	if strings.Contains(itens, "Turret Plating") {
		t.Error("item fora da loja foi publicado")
	}
	// O nome canonico sai ao lado, porque o corpus externo usa o ingles.
	if !strings.Contains(itens, "Berserkers Greaves") {
		t.Error("o nome canonico nao foi publicado")
	}
}

// TestRunaQueOJogoNaoOferecFicaDeFora: as 34 entradas removidas e os titulos de
// pagina nao sao runas jogaveis.
func TestRunaQueOJogoNaoOferecFicaDeFora(t *testing.T) {
	runas := corpo(t, gerar(t), "02-runes.md")
	if strings.Contains(runas, "Removed") {
		t.Error("runa que o jogo nao oferece foi publicada")
	}
	if !strings.Contains(runas, "out_of_scope") ||
		!strings.Contains(runas, "depende de estado de combate") {
		t.Error("o escopo e o motivo nao chegaram ao arquivo")
	}
}

// TestMarcadorResolvido e o que transforma "causa @TotalDamage@" em numero.
func TestMarcadorResolvido(t *testing.T) {
	camp := corpo(t, gerar(t), "04-champions.md")
	if !strings.Contains(camp, "30/60/90/120/150 + 1.5 attack_damage") {
		t.Error("o marcador de efeito nao foi resolvido")
	}
	if !strings.Contains(camp, "1.5/1.5/1.5/1.5/1.5s") {
		t.Error("o marcador de serie nomeada nao foi resolvido")
	}
	// Marcador sem valor vira lacuna visivel, e nao buraco na frase.
	if !strings.Contains(camp, "(?)") {
		t.Error("marcador sem valor sumiu em vez de virar lacuna declarada")
	}
	if strings.Contains(camp, "@Perdido@") {
		t.Error("marcador cru foi publicado")
	}
}

// TestPrecisaoDeFloat32: o jogo calcula em float32 e o dump serializa a
// representacao float64 disso. Publicar 4.199999809265137 mentiria sobre a
// precisao do dado e gastaria cinco tokens onde um basta.
func TestPrecisaoDeFloat32(t *testing.T) {
	stats := corpo(t, gerar(t), "06-champion-stats.md")
	if !strings.Contains(stats, "4.2") {
		t.Error("o valor nao saiu na precisao da fonte")
	}
	if strings.Contains(stats, "4.199999") {
		t.Error("o ruido de representacao float32 foi publicado")
	}
}

// TestManifestoDeclaraOsLimites: e ele que impede o modelo de completar lacuna
// com o que sabe de outra fonte.
func TestManifestoDeclaraOsLimites(t *testing.T) {
	m := corpo(t, gerar(t), "00-MANIFEST.md")
	frases := []string{
		"apenas o Summoner's Rift",
		"não simula combate",
		"nunca publicamos zero no lugar de uma lacuna",
		"não é uma build boa",
		"Não há dados de partidas",
	}
	for _, f := range frases {
		if !strings.Contains(m, f) {
			t.Errorf("o manifesto nao declara: %q", f)
		}
	}
}

func TestChangelogInicial(t *testing.T) {
	res := gerar(t)
	if res.Changelog == nil || res.Changelog.Name != "16.99.md" {
		t.Fatalf("changelog = %+v", res.Changelog)
	}
	if !strings.Contains(res.Changelog.Body, "captura inicial") {
		t.Error("sem patch anterior, o changelog deveria dizer que e a captura inicial")
	}
}

func TestChangelogComparaPatches(t *testing.T) {
	antes := datasetDeTeste()
	antes.Patch = "16.98"
	antes.Items[0].Custo = 1000
	antes.Items[0].Stats = canon.Vector{canon.AttackSpeedPct: 25, canon.MoveSpeed: 45}

	res, err := Generate(Input{Dataset: datasetDeTeste(), Anterior: antes})
	if err != nil {
		t.Fatal(err)
	}
	body := res.Changelog.Body
	if !strings.Contains(body, "custo 1000 → 1100") {
		t.Errorf("a mudanca de custo nao foi detectada:\n%s", body)
	}
	if !strings.Contains(body, "%AS 25 → 30") {
		t.Errorf("a mudanca de atributo nao foi detectada:\n%s", body)
	}
}

// TestCheckOutputDirAvisaSemApagar: o destino e um diretorio que o usuario
// tambem manipula a mao, e apagar arquivo que o export nao criou seria
// inaceitavel.
func TestCheckOutputDirAvisaSemApagar(t *testing.T) {
	dir := t.TempDir()
	res := gerar(t)
	if err := res.Write(dir); err != nil {
		t.Fatal(err)
	}
	avulso := filepath.Join(dir, "rascunho.md")
	if err := os.WriteFile(avulso, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	estranhos := res.CheckOutputDir(dir)
	if len(estranhos) != 1 || estranhos[0] != "rascunho.md" {
		t.Fatalf("estranhos = %v", estranhos)
	}
	if _, err := os.Stat(avulso); err != nil {
		t.Error("o arquivo avulso foi removido")
	}
	// PROJECT-INSTRUCTIONS.md pertence ao conjunto e nao pode ser acusado.
	if err := os.WriteFile(filepath.Join(dir, "PROJECT-INSTRUCTIONS.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := res.CheckOutputDir(dir); len(got) != 1 {
		t.Errorf("PROJECT-INSTRUCTIONS.md foi acusado: %v", got)
	}
}

func TestEstimateTokens(t *testing.T) {
	if got := estimateTokens("12345678"); got != 2 {
		t.Fatalf("estimateTokens = %d, esperado 2", got)
	}
}
