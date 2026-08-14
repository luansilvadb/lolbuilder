package canon

import (
	"strings"
	"testing"
)

// TestAllCobreOVocabulario: Stats() ordena por All, entao um stat fora de All
// sairia no fim, fora de ordem, e o dataset publicaria em ordem instavel.
func TestAllCobreOVocabulario(t *testing.T) {
	emAll := make(map[Stat]bool, len(All))
	for _, s := range All {
		if emAll[s] {
			t.Fatalf("%s aparece duas vezes em All", s)
		}
		emAll[s] = true
	}
	for _, s := range rotulos {
		if !emAll[s] {
			t.Errorf("%s e resolvido por um rotulo mas nao esta em All", s)
		}
	}
	if len(All) != 23 {
		t.Errorf("All tem %d stats; o 16.16 mediu 23 rotulos distintos", len(All))
	}
}

func TestPercentual(t *testing.T) {
	if !AttackSpeedPct.Percentual() || !MoveSpeedPct.Percentual() {
		t.Error("stat percentual nao foi reconhecido")
	}
	if MoveSpeed.Percentual() || Health.Percentual() || MagicPenetration.Percentual() {
		t.Error("stat plano foi reconhecido como percentual")
	}
}

// TestLookupStatDistingueUnidade e a invariante que impede somar 45 de
// deslocamento com 4% de deslocamento.
func TestLookupStatDistingueUnidade(t *testing.T) {
	plano, ok := LookupStat("Move Speed", false)
	if !ok || plano != MoveSpeed {
		t.Fatalf("Move Speed plano resolveu para %q (%v)", plano, ok)
	}
	pct, ok := LookupStat("Move Speed", true)
	if !ok || pct != MoveSpeedPct {
		t.Fatalf("Move Speed percentual resolveu para %q (%v)", pct, ok)
	}
	if plano == pct {
		t.Fatal("as duas unidades colapsaram no mesmo stat")
	}
}

func TestLookupStatNormalizaCaixaEEspaco(t *testing.T) {
	got, ok := LookupStat("  ABILITY HasTe  ", false)
	if !ok || got != AbilityHaste {
		t.Fatalf("LookupStat nao normalizou: %q (%v)", got, ok)
	}
}

func TestLookupStatRejeitaDesconhecido(t *testing.T) {
	if _, ok := LookupStat("Grandeza Inventada", false); ok {
		t.Fatal("rotulo fora do vocabulario resolveu")
	}
	// Rotulo certo, unidade errada, tambem nao resolve.
	if _, ok := LookupStat("Attack Speed", false); ok {
		t.Fatal("attack speed plano resolveu; a fonte so publica percentual")
	}
}

func TestVectorAddEMerge(t *testing.T) {
	v := Vector{}
	v.Add(Health, 100)
	v.Add(Health, 50)
	v.Merge(Vector{Armor: 30, Health: 10})

	if v[Health] != 160 || v[Armor] != 30 {
		t.Fatalf("soma errada: %+v", v)
	}
}

// TestVectorStatsSegueOrdemDeAll: a ordem da saida vem de All e nao de sort
// sobre as chaves, porque All e a ordem em que o dataset publica.
func TestVectorStatsSegueOrdemDeAll(t *testing.T) {
	v := Vector{MoveSpeed: 45, AbilityHaste: 15, Health: 200}
	got := v.Stats()
	want := []Stat{AbilityHaste, Health, MoveSpeed}
	if len(got) != len(want) {
		t.Fatalf("Stats() = %v, esperado %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Stats() = %v, esperado %v", got, want)
		}
	}
}

func TestVectorStatsExpoeStatForaDeAll(t *testing.T) {
	v := Vector{Health: 1, "inventado": 2}
	got := v.Stats()
	if len(got) != 2 || got[1] != "inventado" {
		t.Fatalf("stat fora de All sumiu da saida: %v", got)
	}
}

func TestVectorValidate(t *testing.T) {
	if err := (Vector{Health: 1, Armor: 2}).Validate(); err != nil {
		t.Fatalf("vetor valido rejeitado: %v", err)
	}
	err := (Vector{Health: 1, "nao_existe": 2}).Validate()
	if err == nil {
		t.Fatal("stat fora do vocabulario passou")
	}
	if !strings.Contains(err.Error(), "nao_existe") {
		t.Fatalf("o erro nao nomeia o stat: %v", err)
	}
}
