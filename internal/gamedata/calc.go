package gamedata

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

// Term e uma parcela que escala com uma estatistica do campeao.
type Term struct {
	Stat        string // nome canonico, ex: attack_damage
	Formula     string // total, base ou bonus
	Coefficient float64
}

// Expr e uma formula resolvida: uma parcela fixa mais os coeficientes de
// escala.
//
// Nao e um numero. O dataset publica quanto a habilidade causa em funcao das
// estatisticas, e quem monta a build faz a conta — simular combate esta fora do
// escopo.
type Expr struct {
	Flat  float64
	Terms []Term
}

// add soma duas expressoes.
func (e Expr) add(o Expr) Expr {
	e.Flat += o.Flat
	e.Terms = append(e.Terms, o.Terms...)
	return e
}

// scale multiplica a expressao por uma constante.
func (e Expr) scale(k float64) Expr {
	e.Flat *= k
	out := make([]Term, len(e.Terms))
	for i, t := range e.Terms {
		t.Coefficient *= k
		out[i] = t
	}
	e.Terms = out
	return e
}

// constant informa se a expressao nao depende de estatistica alguma.
func (e Expr) constant() bool { return len(e.Terms) == 0 }

// Constant expoe a mesma pergunta para fora do pacote.
func (e Expr) Constant() bool { return e.constant() }

// Normalize junta termos do mesmo stat e ordena, para que a saida do export
// seja identica entre execucoes.
func (e Expr) Normalize() Expr {
	if len(e.Terms) < 2 {
		return e
	}
	merged := map[Term]float64{}
	for _, t := range e.Terms {
		key := Term{Stat: t.Stat, Formula: t.Formula}
		merged[key] += t.Coefficient
	}
	out := make([]Term, 0, len(merged))
	for k, c := range merged {
		if c == 0 {
			continue
		}
		k.Coefficient = c
		out = append(out, k)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Stat != out[j].Stat {
			return out[i].Stat < out[j].Stat
		}
		return out[i].Formula < out[j].Formula
	})
	e.Terms = out
	return e
}

// Zero informa se a expressao nao afirma numero nenhum.
//
// Serve para distinguir "resolveu e deu zero" de "resolveu de verdade": uma
// formula que resolve para zero absoluto quase sempre significa que a fonte
// publica a serie zerada, e publicar esse zero mentiria.
func (e Expr) Zero() bool {
	return e.Flat == 0 && len(e.Terms) == 0
}

// ErrUnresolved marca uma formula que o interpretador nao sabe avaliar.
//
// E um erro e nao um zero de proposito: uma habilidade publicada com dano zero
// mente, enquanto uma habilidade ausente da tabela apenas nao informa. So o
// segundo e recuperavel pelo leitor.
type ErrUnresolved struct {
	Part   string
	Reason string
}

func (e *ErrUnresolved) Error() string {
	return fmt.Sprintf("parcela %s nao resolvida: %s", e.Part, e.Reason)
}

// Context reune tudo que uma formula pode consultar.
type Context struct {
	Rank  int
	Level int
	Spell *Spell
	Stats StatNames
}

// part e a forma generica de uma parcela, com todos os campos que os tipos
// observados usam. Decodificar tudo de uma vez evita uma struct por tipo.
type part struct {
	Type string `json:"__type"`

	DataValue   string   `json:"mDataValue"`
	Number      *float64 `json:"mNumber"`
	Coefficient *float64 `json:"mCoefficient"`

	// Enums em ponteiro, nunca em int. A fonte omite o campo quando ele vale
	// zero, e zero e um valor real — para mStat e o poder de habilidade. Em int,
	// "ausente" e "poder de habilidade" colapsam no mesmo valor e o defeito fica
	// invisivel enquanto a suposicao estiver certa.
	Stat        *int `json:"mStat"`
	StatFormula *int `json:"mStatFormula"`

	Level1Value *float64  `json:"mLevel1Value"`
	Breakpoints []brkpt   `json:"mBreakpoints"`
	StartValue  *float64  `json:"mStartValue"`
	EndValue    *float64  `json:"mEndValue"`
	Values      []float64 `json:"values"`

	Ceiling *float64 `json:"mCeiling"`
	Floor   *float64 `json:"mFloor"`

	Subparts []json.RawMessage `json:"mSubparts"`
	Subpart  json.RawMessage   `json:"mSubpart"`
	Part1    json.RawMessage   `json:"mPart1"`
	Part2    json.RawMessage   `json:"mPart2"`

	Multiplier    json.RawMessage   `json:"mMultiplier"`
	ModifiedCalc  string            `json:"mModifiedGameCalculation"`
	MultiplierMod json.RawMessage   `json:"mMultiplierMod"`
	FormulaParts  []json.RawMessage `json:"mFormulaParts"`
	TooltipOnly   *int              `json:"mSimpleTooltipCalculationDisplay"`
	BuffName      string            `json:"mBuffName"`
	EffectIndex   *int              `json:"mEffectIndex"`

	SpellCalcKey    string `json:"mSpellCalculationKey"`
	ConditionalCalc string `json:"mConditionalGameCalculation"`
	OpaqueDataValue string `json:"{137cf12a}"`
}

type brkpt struct {
	Level int     `json:"mLevel"`
	Bonus float64 `json:"mAdditionalBonusAtThisLevel"`
	Type  string  `json:"__type"`
}

// Evaluate resolve uma formula nomeada da habilidade.
func Evaluate(name string, ctx Context) (Expr, error) {
	return evalNamed(name, ctx, 0)
}

// evalNamed e a Evaluate com profundidade, para que um salto de uma formula
// para outra some ao contador em vez de zera-lo. Um ciclo entre formulas
// nomeadas so e cortado se a profundidade sobreviver ao salto.
func evalNamed(name string, ctx Context, depth int) (Expr, error) {
	if depth > maxDepth {
		return Expr{}, &ErrUnresolved{Part: name, Reason: "formula recursiva demais"}
	}
	raw, ok := ctx.Spell.Calculations[name]
	if !ok {
		return Expr{}, &ErrUnresolved{Part: name, Reason: "formula ausente da habilidade"}
	}
	return evalPart(raw, ctx, depth)
}

// maxDepth corta recursao patologica. GameCalculationModified pode referenciar
// outra formula, e um ciclo travaria o build. O corte so funciona porque a
// profundidade atravessa o salto entre formulas nomeadas: reinicia-la a cada
// salto deixaria o ciclo Foo->Bar->Foo estourar a pilha.
const maxDepth = 12

func evalPart(raw json.RawMessage, ctx Context, depth int) (Expr, error) {
	if depth > maxDepth {
		return Expr{}, &ErrUnresolved{Part: "?", Reason: "formula recursiva demais"}
	}
	var p part
	if err := json.Unmarshal(raw, &p); err != nil {
		return Expr{}, &ErrUnresolved{Part: "?", Reason: err.Error()}
	}
	if p.DataValue == "" && p.OpaqueDataValue != "" {
		p.DataValue = p.OpaqueDataValue
	}

	switch p.Type {

	case "GameCalculation":
		var out Expr
		for _, sub := range p.FormulaParts {
			e, err := evalPart(sub, ctx, depth+1)
			if err != nil {
				return Expr{}, err
			}
			out = out.add(e)
		}
		return out, nil

	case "GameCalculationModified":
		base, err := evalNamed(p.ModifiedCalc, ctx, depth+1)
		if err != nil {
			return Expr{}, err
		}
		mult, err := evalPart(p.Multiplier, ctx, depth+1)
		if err != nil {
			return Expr{}, err
		}
		switch {
		case mult.constant():
			return base.scale(mult.Flat), nil
		case base.constant():
			return mult.scale(base.Flat), nil
		}
		return Expr{}, &ErrUnresolved{Part: p.Type,
			Reason: "os dois fatores escalam com estatistica, o produto nao e linear"}

	case "NamedDataValueCalculationPart", "{2b25a73a}":
		v, ok := ctx.namedValue(p.DataValue)
		if !ok {
			return Expr{}, &ErrUnresolved{Part: p.Type,
				Reason: "valor " + p.DataValue + " ausente do rank " + itoa(ctx.Rank)}
		}
		return Expr{Flat: v}, nil

	case "NumberCalculationPart":
		v := 0.0
		if p.Number != nil {
			v = *p.Number
		}
		return Expr{Flat: v}, nil

	case "StatByCoefficientCalculationPart":
		if p.Coefficient == nil {
			return Expr{}, &ErrUnresolved{Part: p.Type, Reason: "sem mCoefficient"}
		}
		return ctx.statTerm(p.Stat, p.StatFormula, *p.Coefficient, p.Type)

	case "StatByNamedDataValueCalculationPart":
		v, ok := ctx.namedValue(p.DataValue)
		if !ok {
			return Expr{}, &ErrUnresolved{Part: p.Type,
				Reason: "valor " + p.DataValue + " ausente do rank " + itoa(ctx.Rank)}
		}
		return ctx.statTerm(p.Stat, p.StatFormula, v, p.Type)

	case "SumOfSubPartsCalculationPart", "{8a96ea3c}":
		var out Expr
		for _, sub := range p.Subparts {
			e, err := evalPart(sub, ctx, depth+1)
			if err != nil {
				return Expr{}, err
			}
			out = out.add(e)
		}
		return out, nil

	case "ProductOfSubPartsCalculationPart":
		a, err := evalPart(p.Part1, ctx, depth+1)
		if err != nil {
			return Expr{}, err
		}
		b, err := evalPart(p.Part2, ctx, depth+1)
		if err != nil {
			return Expr{}, err
		}
		switch {
		case a.constant():
			return b.scale(a.Flat), nil
		case b.constant():
			return a.scale(b.Flat), nil
		}
		return Expr{}, &ErrUnresolved{Part: p.Type,
			Reason: "os dois fatores escalam com estatistica, o produto nao e linear"}

	case "StatBySubPartCalculationPart":
		sub, err := evalPart(p.Subpart, ctx, depth+1)
		if err != nil {
			return Expr{}, err
		}
		if !sub.constant() {
			return Expr{}, &ErrUnresolved{Part: p.Type,
				Reason: "coeficiente escala com estatistica, o produto nao e linear"}
		}
		return ctx.statTerm(p.Stat, p.StatFormula, sub.Flat, p.Type)

	case "ClampSubPartsCalculationPart":
		// Soma as parcelas e prende o resultado entre piso e teto.
		//
		// So resolve quando a soma nao escala com estatistica. Prender uma
		// expressao simbolica exigiria saber o valor da estatistica, que e
		// justamente o que o dataset NAO fixa — publicar o resultado sem o
		// limite aplicado daria um numero que o jogo nunca produz.
		var soma Expr
		for _, sub := range p.Subparts {
			e, err := evalPart(sub, ctx, depth+1)
			if err != nil {
				return Expr{}, err
			}
			soma = soma.add(e)
		}
		if !soma.constant() {
			return Expr{}, &ErrUnresolved{Part: p.Type,
				Reason: "a soma escala com estatistica, e o limite so pode ser aplicado sobre um numero"}
		}
		v := soma.Flat
		if p.Floor != nil {
			v = math.Max(v, *p.Floor)
		}
		if p.Ceiling != nil {
			v = math.Min(v, *p.Ceiling)
		}
		return Expr{Flat: v}, nil

	case "ByCharLevelBreakpointsCalculationPart":
		if p.Level1Value == nil {
			return Expr{}, &ErrUnresolved{Part: p.Type, Reason: "sem mLevel1Value"}
		}
		v := *p.Level1Value
		for _, b := range p.Breakpoints {
			if ctx.Level >= b.Level {
				v += b.Bonus
			}
		}
		return Expr{Flat: v}, nil

	case "ByCharLevelInterpolationCalculationPart":
		if p.StartValue == nil && p.EndValue == nil {
			return Expr{}, &ErrUnresolved{Part: p.Type, Reason: "sem extremos de interpolacao"}
		}
		start, end := 0.0, 0.0
		if p.StartValue != nil {
			start = *p.StartValue
		}
		if p.EndValue != nil {
			end = *p.EndValue
		}
		// Interpola linearmente entre o nivel 1 e o 18.
		t := float64(ctx.Level-1) / 17.0
		return Expr{Flat: start + (end-start)*t}, nil

	case "ByCharLevelFormulaCalculationPart":
		if ctx.Level < 0 || ctx.Level >= len(p.Values) {
			return Expr{}, &ErrUnresolved{Part: p.Type,
				Reason: "nivel " + itoa(ctx.Level) + " fora da tabela"}
		}
		return Expr{Flat: p.Values[ctx.Level]}, nil

	case "AbilityResourceByCoefficientCalculationPart":
		if p.Coefficient == nil {
			return Expr{}, &ErrUnresolved{Part: p.Type, Reason: "sem mCoefficient"}
		}
		return Expr{Terms: []Term{{Stat: "mana", Formula: "total", Coefficient: *p.Coefficient}}}, nil

	case "{f3cbe7b2}", "SubCalculationPart", "NamedCalculationPart":
		if p.SpellCalcKey == "" {
			return Expr{}, &ErrUnresolved{Part: p.Type, Reason: "sem mSpellCalculationKey"}
		}
		return evalNamed(p.SpellCalcKey, ctx, depth+1)

	case "GameCalculationConditional":
		if p.ConditionalCalc == "" {
			return Expr{}, &ErrUnresolved{Part: p.Type, Reason: "sem mConditionalGameCalculation"}
		}
		return evalNamed(p.ConditionalCalc, ctx, depth+1)

	case "CooldownMultiplierCalculationPart":
		// Fora de combate, a reducao de recarga estatica base e zero
		// (multiplicador 1.0).
		return Expr{Flat: 1.0}, nil

	// Os que dependem do estado da partida. Nao ha valor honesto a publicar
	// para eles fora de combate.
	case "BuffCounterByNamedDataValueCalculationPart", "BuffCounterByCoefficientCalculationPart":
		return Expr{}, &ErrUnresolved{Part: p.Type,
			Reason: "depende de quantas cargas do efeito estao ativas em partida"}

	case "EffectValueCalculationPart":
		return Expr{}, &ErrUnresolved{Part: p.Type,
			Reason: "aponta para effectAmounts, que a fonte publica zerado"}
	}

	return Expr{}, &ErrUnresolved{Part: p.Type, Reason: "tipo de parcela desconhecido"}
}

// namedValue le uma serie da habilidade no rank do contexto.
func (c Context) namedValue(name string) (float64, bool) {
	values, ok := c.Spell.DataValue(name)
	if !ok {
		return 0, false
	}
	return AtRank(values, c.Rank)
}

// statTerm monta o termo de escala, traduzindo os enums da fonte.
func (c Context) statTerm(stat, formula *int, coef float64, partType string) (Expr, error) {
	name, err := c.Stats.Name(stat)
	if err != nil {
		return Expr{}, &ErrUnresolved{Part: partType, Reason: err.Error()}
	}
	return Expr{Terms: []Term{{
		Stat:        name,
		Formula:     formulaName(formula),
		Coefficient: coef,
	}}}, nil
}

func itoa(n int) string { return fmt.Sprintf("%d", n) }
