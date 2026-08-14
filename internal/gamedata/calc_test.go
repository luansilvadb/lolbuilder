package gamedata

import (
	"encoding/json"
	"strings"
	"testing"
)

// statsDeTeste cobre os enums que os testes usam, com os mesmos nomes da
// curadoria real.
var statsDeTeste = StatNames{
	stat:    map[int]string{0: "ability_power", 2: "attack_damage", 12: "max_health"},
	formula: formulaNames,
}

func spellCom(calcs map[string]string, valores ...DataValue) *Spell {
	raw := map[string]json.RawMessage{}
	for k, v := range calcs {
		raw[k] = json.RawMessage(v)
	}
	return &Spell{Calculations: raw, DataValues: valores}
}

func avaliar(t *testing.T, sp *Spell, nome string, rank int) (Expr, error) {
	t.Helper()
	return Evaluate(nome, Context{Rank: rank, Level: 18, Spell: sp, Stats: statsDeTeste})
}

func TestSomaDeParcelasFixaEEscala(t *testing.T) {
	sp := spellCom(map[string]string{
		"TotalDamage": `{"__type":"GameCalculation","mFormulaParts":[
			{"__type":"NamedDataValueCalculationPart","mDataValue":"BaseDamage"},
			{"__type":"StatByCoefficientCalculationPart","mStat":2,"mCoefficient":1.5}]}`,
	}, DataValue{Name: "BaseDamage", Values: []float64{0, 30, 60, 90, 120, 150, 150}})

	e, err := avaliar(t, sp, "TotalDamage", 1)
	if err != nil {
		t.Fatal(err)
	}
	e = e.Normalize()
	if e.Flat != 30 {
		t.Errorf("parcela fixa no rank 1 = %v, esperado 30", e.Flat)
	}
	if len(e.Terms) != 1 || e.Terms[0].Stat != "attack_damage" || e.Terms[0].Coefficient != 1.5 {
		t.Errorf("escala errada: %+v", e.Terms)
	}

	e5, err := avaliar(t, sp, "TotalDamage", 5)
	if err != nil {
		t.Fatal(err)
	}
	if e5.Flat != 150 {
		t.Errorf("parcela fixa no rank 5 = %v, esperado 150", e5.Flat)
	}
}

// TestStatOmitidoEPoderDeHabilidade: a fonte omite mStat quando ele vale zero, e
// zero e um valor real. Ler o enum em int faria "ausente" e "poder de
// habilidade" colapsarem, e o defeito ficaria invisivel enquanto a suposicao
// estivesse certa.
func TestStatOmitidoEPoderDeHabilidade(t *testing.T) {
	sp := spellCom(map[string]string{
		"D": `{"__type":"StatByCoefficientCalculationPart","mCoefficient":0.6}`,
	})
	e, err := avaliar(t, sp, "D", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(e.Terms) != 1 || e.Terms[0].Stat != "ability_power" {
		t.Fatalf("mStat omitido resolveu para %+v", e.Terms)
	}
}

// TestProdutoNaoLinearNaoResolve: dois fatores que escalam com estatistica
// produzem um termo quadratico, e a Expr do projeto e linear por construcao.
// Resolver isso daria um numero errado com cara de certo.
func TestProdutoNaoLinearNaoResolve(t *testing.T) {
	sp := spellCom(map[string]string{
		"D": `{"__type":"ProductOfSubPartsCalculationPart",
			"mPart1":{"__type":"StatByCoefficientCalculationPart","mStat":2,"mCoefficient":1},
			"mPart2":{"__type":"StatByCoefficientCalculationPart","mStat":0,"mCoefficient":1}}`,
	})
	_, err := avaliar(t, sp, "D", 1)
	if err == nil {
		t.Fatal("produto nao linear resolveu")
	}
	if !strings.Contains(err.Error(), "nao e linear") {
		t.Fatalf("o erro nao explica o motivo: %v", err)
	}
}

// TestClampSubParts e a parcela que o avaliador do modo Jade nao tratava.
func TestClampSubParts(t *testing.T) {
	casos := []struct {
		nome string
		json string
		want float64
	}{
		{"dentro dos limites", `{"__type":"ClampSubPartsCalculationPart","mFloor":1,"mCeiling":2,
			"mSubparts":[{"__type":"NumberCalculationPart","mNumber":1.5}]}`, 1.5},
		{"preso no piso", `{"__type":"ClampSubPartsCalculationPart","mFloor":1,"mCeiling":2,
			"mSubparts":[{"__type":"NumberCalculationPart","mNumber":0.2}]}`, 1},
		{"preso no teto", `{"__type":"ClampSubPartsCalculationPart","mFloor":1,"mCeiling":2,
			"mSubparts":[{"__type":"NumberCalculationPart","mNumber":9}]}`, 2},
		{"teto nulo so aplica o piso", `{"__type":"ClampSubPartsCalculationPart","mFloor":0.6,"mCeiling":null,
			"mSubparts":[{"__type":"NumberCalculationPart","mNumber":9}]}`, 9},
		{"soma antes de prender", `{"__type":"ClampSubPartsCalculationPart","mFloor":0,"mCeiling":10,
			"mSubparts":[{"__type":"NumberCalculationPart","mNumber":4},
			             {"__type":"NumberCalculationPart","mNumber":3}]}`, 7},
	}
	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			sp := spellCom(map[string]string{"D": c.json})
			e, err := avaliar(t, sp, "D", 1)
			if err != nil {
				t.Fatal(err)
			}
			if e.Flat != c.want {
				t.Fatalf("resultado = %v, esperado %v", e.Flat, c.want)
			}
		})
	}
}

// TestClampComEscalaNaoResolve: prender uma expressao simbolica exigiria saber o
// valor da estatistica, que e justamente o que o dataset nao fixa. Publicar sem
// o limite daria um numero que o jogo nunca produz.
func TestClampComEscalaNaoResolve(t *testing.T) {
	sp := spellCom(map[string]string{
		"D": `{"__type":"ClampSubPartsCalculationPart","mFloor":1,"mCeiling":2,
			"mSubparts":[{"__type":"StatByCoefficientCalculationPart","mStat":2,"mCoefficient":1}]}`,
	})
	_, err := avaliar(t, sp, "D", 1)
	if err == nil {
		t.Fatal("clamp sobre expressao simbolica resolveu")
	}
	if !strings.Contains(err.Error(), "limite") {
		t.Fatalf("o erro nao explica o motivo: %v", err)
	}
}

func TestGameCalculationModified(t *testing.T) {
	sp := spellCom(map[string]string{
		"Base":    `{"__type":"GameCalculation","mFormulaParts":[{"__type":"NumberCalculationPart","mNumber":100}]}`,
		"Derived": `{"__type":"GameCalculationModified","mModifiedGameCalculation":"Base","mMultiplier":{"__type":"NumberCalculationPart","mNumber":0.25}}`,
	})
	e, err := avaliar(t, sp, "Derived", 1)
	if err != nil {
		t.Fatal(err)
	}
	if e.Flat != 25 {
		t.Fatalf("Base x 0.25 = %v, esperado 25", e.Flat)
	}

	d, ok := sp.DerivacaoDe("Derived")
	if !ok || d.De != "Base" || d.Multiplicador != 0.25 {
		t.Fatalf("derivacao nao foi reconhecida: %+v (%v)", d, ok)
	}
}

// TestCicloEntreFormulasNaoTrava: a profundidade precisa atravessar o salto
// entre formulas nomeadas, ou Foo->Bar->Foo estoura a pilha.
func TestCicloEntreFormulasNaoTrava(t *testing.T) {
	sp := spellCom(map[string]string{
		"Foo": `{"__type":"GameCalculationModified","mModifiedGameCalculation":"Bar","mMultiplier":{"__type":"NumberCalculationPart","mNumber":1}}`,
		"Bar": `{"__type":"GameCalculationModified","mModifiedGameCalculation":"Foo","mMultiplier":{"__type":"NumberCalculationPart","mNumber":1}}`,
	})
	_, err := avaliar(t, sp, "Foo", 1)
	if err == nil {
		t.Fatal("ciclo entre formulas resolveu")
	}
	if !strings.Contains(err.Error(), "recursiva") {
		t.Fatalf("o erro nao identifica o ciclo: %v", err)
	}
}

func TestParcelasQueDependemDaPartida(t *testing.T) {
	casos := map[string]string{
		"cargas de efeito":  `{"__type":"BuffCounterByCoefficientCalculationPart","mCoefficient":1}`,
		"effectAmounts":     `{"__type":"EffectValueCalculationPart","mEffectIndex":3}`,
		"tipo desconhecido": `{"__type":"ParcelaQueNaoExiste"}`,
	}
	for nome, j := range casos {
		t.Run(nome, func(t *testing.T) {
			sp := spellCom(map[string]string{"D": j})
			if _, err := avaliar(t, sp, "D", 1); err == nil {
				t.Fatal("parcela sem valor honesto resolveu")
			}
		})
	}
}

func TestInterpolacaoPorNivel(t *testing.T) {
	sp := spellCom(map[string]string{
		"D": `{"__type":"ByCharLevelInterpolationCalculationPart","mStartValue":10,"mEndValue":44}`,
	})
	e, err := Evaluate("D", Context{Rank: 1, Level: 1, Spell: sp, Stats: statsDeTeste})
	if err != nil || e.Flat != 10 {
		t.Fatalf("nivel 1 = %v (%v), esperado 10", e.Flat, err)
	}
	e, err = Evaluate("D", Context{Rank: 1, Level: 18, Spell: sp, Stats: statsDeTeste})
	if err != nil || e.Flat != 44 {
		t.Fatalf("nivel 18 = %v (%v), esperado 44", e.Flat, err)
	}
}

func TestBreakpointsPorNivel(t *testing.T) {
	j := `{"__type":"ByCharLevelBreakpointsCalculationPart","mLevel1Value":10,
		"mBreakpoints":[{"mLevel":6,"mAdditionalBonusAtThisLevel":5},
		                {"mLevel":11,"mAdditionalBonusAtThisLevel":5}]}`
	sp := spellCom(map[string]string{"D": j})
	for nivel, want := range map[int]float64{1: 10, 6: 15, 11: 20, 18: 20} {
		e, err := Evaluate("D", Context{Rank: 1, Level: nivel, Spell: sp, Stats: statsDeTeste})
		if err != nil || e.Flat != want {
			t.Errorf("nivel %d = %v (%v), esperado %v", nivel, e.Flat, err, want)
		}
	}
}

// TestStatForaDaTabelaNaoResolve: enum sem curadoria deixa a parcela em aberto
// e entra na cobertura. Publicar com o stat errado mentiria.
func TestStatForaDaTabelaNaoResolve(t *testing.T) {
	sp := spellCom(map[string]string{
		"D": `{"__type":"StatByCoefficientCalculationPart","mStat":99,"mCoefficient":1}`,
	})
	_, err := avaliar(t, sp, "D", 1)
	if err == nil {
		t.Fatal("enum fora da tabela curada resolveu")
	}
	if !strings.Contains(err.Error(), "statenum.json") {
		t.Fatalf("o erro nao diz onde curar: %v", err)
	}
}

func TestNormalizeJuntaESortea(t *testing.T) {
	e := Expr{Flat: 1, Terms: []Term{
		{Stat: "attack_damage", Formula: "total", Coefficient: 0.5},
		{Stat: "ability_power", Formula: "total", Coefficient: 0.2},
		{Stat: "attack_damage", Formula: "total", Coefficient: 0.5},
	}}.Normalize()

	if len(e.Terms) != 2 {
		t.Fatalf("termos iguais nao foram juntados: %+v", e.Terms)
	}
	if e.Terms[0].Stat != "ability_power" {
		t.Errorf("ordem instavel: %+v", e.Terms)
	}
	if e.Terms[1].Coefficient != 1.0 {
		t.Errorf("coeficientes nao somaram: %+v", e.Terms[1])
	}
}

func TestNormalizeDescartaCoeficienteZero(t *testing.T) {
	e := Expr{Terms: []Term{
		{Stat: "attack_damage", Formula: "total", Coefficient: 0.5},
		{Stat: "attack_damage", Formula: "total", Coefficient: -0.5},
	}}.Normalize()
	if len(e.Terms) != 0 {
		t.Fatalf("termo que se anula sobreviveu: %+v", e.Terms)
	}
}
