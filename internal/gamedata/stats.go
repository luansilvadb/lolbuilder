package gamedata

import "fmt"

// Scaling e uma estatistica que cresce por nivel.
type Scaling struct {
	Base     float64 `json:"base"`
	PerLevel float64 `json:"por_nivel"`
}

// At devolve o valor no nivel pedido (1..18).
//
// O crescimento e linear a partir do nivel 1: no nivel 1 vale a base, e cada
// nivel acima soma o incremento uma vez.
func (s Scaling) At(level int) float64 {
	if level <= 1 {
		return s.Base
	}
	return s.Base + s.PerLevel*float64(level-1)
}

// Stats sao as estatisticas base de um campeao.
type Stats struct {
	HP           Scaling `json:"vida"`
	AttackDamage Scaling `json:"dano_de_ataque"`
	Armor        Scaling `json:"armadura"`
	AttackSpeed  Scaling `json:"velocidade_de_ataque"`
	MagicResist  Scaling `json:"resistencia_magica"`

	// Fixas: nao crescem por nivel.
	MoveSpeed        float64 `json:"velocidade_de_movimento"`
	AttackRange      float64 `json:"alcance_de_ataque"`
	AttackSpeedRatio float64 `json:"razao_de_velocidade_de_ataque"`

	// Opcionais: a fonte omite em alguns campeoes. Ficam como ponteiro para que
	// ausente e zero nunca se confundam — publicar 0 de regeneracao afirmaria
	// que o campeao nao regenera.
	HPRegen              *Scaling `json:"regeneracao_de_vida,omitempty"`
	CritDamageMultiplier *float64 `json:"multiplicador_de_dano_critico,omitempty"`

	Melee bool `json:"corpo_a_corpo"`
}

// requiredStat descreve uma estatistica sem a qual o campeao nao serve para
// nada no dataset.
type requiredStat struct {
	name string
	get  func(*Record) float64
}

// requiredStats sao as presentes nos 173 campeoes do 16.16. Se alguma sumir num
// patch, e mudanca de formato, nao caracteristica de um campeao.
var requiredStats = []requiredStat{
	{"vida base", func(r *Record) float64 { return r.BaseHP.BaseValue }},
	{"vida por nivel", func(r *Record) float64 { return r.HPPerLevel.BaseValue }},
	{"dano base", func(r *Record) float64 { return r.BaseDamage.BaseValue }},
	{"dano por nivel", func(r *Record) float64 { return r.DamagePerLevel.BaseValue }},
	{"armadura base", func(r *Record) float64 { return r.BaseArmor.BaseValue }},
	{"armadura por nivel", func(r *Record) float64 { return r.ArmorPerLevel.BaseValue }},
	{"resistencia magica", func(r *Record) float64 { return r.BaseMR.BaseValue }},
	{"velocidade de ataque", func(r *Record) float64 { return r.BaseAttackSpeed.BaseValue }},
	{"alcance de ataque", func(r *Record) float64 { return r.AttackRange.BaseValue }},
	{"velocidade de movimento", func(r *Record) float64 { return r.BaseMoveSpeed.BaseValue }},
}

// Stats monta as estatisticas base e devolve as lacunas encontradas.
//
// Lacuna nao vira zero. Um campeao sem regeneracao base publicada e um campeao
// cuja regeneracao nao sabemos, e afirmar "0 de regeneracao" seria pior que nao
// afirmar nada — o consumidor somaria esse zero.
func (d *Dump) Stats() (Stats, []string) {
	r := &d.Record
	var gaps []string

	for _, req := range requiredStats {
		if req.get(r) == 0 {
			gaps = append(gaps, fmt.Sprintf("%s: %s ausente", r.ShortName(), req.name))
		}
	}

	s := Stats{
		HP:               Scaling{r.BaseHP.BaseValue, r.HPPerLevel.BaseValue},
		AttackDamage:     Scaling{r.BaseDamage.BaseValue, r.DamagePerLevel.BaseValue},
		Armor:            Scaling{r.BaseArmor.BaseValue, r.ArmorPerLevel.BaseValue},
		AttackSpeed:      Scaling{r.BaseAttackSpeed.BaseValue, r.AttackSpeedPerLvl.BaseValue},
		MagicResist:      Scaling{r.BaseMR.BaseValue, r.MrPerLevel.BaseValue},
		MoveSpeed:        r.BaseMoveSpeed.BaseValue,
		AttackRange:      r.AttackRange.BaseValue,
		AttackSpeedRatio: r.AttackSpeedRatio.BaseValue,
		Melee:            r.Melee(),
	}

	if r.BaseHPRegen != nil {
		s.HPRegen = &Scaling{r.BaseHPRegen.BaseValue, r.HPRegenPerLevel.BaseValue}
	} else {
		gaps = append(gaps, fmt.Sprintf("%s: regeneracao de vida base ausente", r.ShortName()))
	}
	if r.CritDamageMultiplier != nil {
		s.CritDamageMultiplier = r.CritDamageMultiplier
	} else {
		gaps = append(gaps, fmt.Sprintf("%s: multiplicador de dano critico ausente", r.ShortName()))
	}

	return s, gaps
}

// Complete informa se todas as estatisticas obrigatorias foram extraidas. E o
// que o eixo de cobertura champion_stats mede: a ausencia de uma opcional e
// lacuna a reportar, nao campeao perdido.
func (s Stats) Complete() bool {
	return s.HP.Base > 0 && s.AttackDamage.Base > 0 && s.Armor.Base > 0 &&
		s.MagicResist.Base > 0 && s.AttackSpeed.Base > 0 && s.AttackRange > 0 &&
		s.MoveSpeed > 0
}
