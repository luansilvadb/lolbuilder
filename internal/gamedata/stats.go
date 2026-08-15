package gamedata

import "fmt"

// Scaling e uma estatistica que cresce por nivel.
type Scaling struct {
	Base     float64 `json:"base"`
	PerLevel float64 `json:"por_nivel"`
}

// At devolve o valor no nivel pedido (1..18).
//
// O crescimento NAO e linear. O jogo aplica um fator que sai de 0.7025 no nivel
// 2 e chega a 1 no 18:
//
//	valor(n) = base + crescimento * (n-1) * (0.7025 + 0.0175*(n-1))
//
// Isto foi MEDIDO contra uma partida, e nao suposto. Rammus tem 2.05 de
// resistencia magica por nivel; do nivel 1 ao 6 o jogo somou 8.0975, que e
// 2.05 * 5 * 0.79 — e nao 2.05 * 5 = 10.25. A armadura confirmou na mesma
// leitura: com 4.5 por nivel, o crescimento base foi 17.775, deixando exatos
// 70 para os dois itens de armadura equipados.
//
// O erro ficou invisivel por muito tempo porque no nivel 18 o fator vale
// exatamente 1: 0.7025 + 0.0175*17 = 1. Ou seja, o total no nivel maximo
// coincide com o crescimento linear, e so os niveis intermediarios saiam
// errados. Nenhuma fonte publica esse fator; so o oraculo em partida o revela.
func (s Scaling) At(level int) float64 {
	return s.Base + s.PerLevel*fatorDeCrescimento(level)
}

// fatorDeCrescimento e quantos "niveis efetivos" de crescimento ja se
// acumularam ao chegar no nivel pedido.
func fatorDeCrescimento(level int) float64 {
	if level <= 1 {
		return 0
	}
	n := float64(level - 1)
	return n * (0.7025 + 0.0175*n)
}

// CrescimentoEntre devolve quanto a estatistica sobe de um nivel a outro.
//
// E a grandeza que a comparacao com a partida usa, porque ela cancela todo
// bonus fixo de item e de runa.
func (s Scaling) CrescimentoEntre(de, ate int) float64 {
	return s.PerLevel * (fatorDeCrescimento(ate) - fatorDeCrescimento(de))
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

// requiredStats sao as estatisticas BASE, presentes nos 173 campeoes do 16.16.
// Se alguma sumir num patch, e mudanca de formato, nao caracteristica de um
// campeao.
//
// Os campos de CRESCIMENTO por nivel ficam de fora de proposito. Medido no
// 16.16: eles NUNCA aparecem presentes valendo zero — a fonte simplesmente os
// omite, e a omissao e como ela diz "nao cresce". Senna nao ganha dano por
// nivel e Thresh nao ganha armadura por nivel, os dois por ganharem isso de
// almas. Reportar essas omissoes como lacuna publicaria uma duvida que nao
// existe.
var requiredStats = []requiredStat{
	{"vida base", func(r *Record) float64 { return r.BaseHP.BaseValue }},
	{"dano base", func(r *Record) float64 { return r.BaseDamage.BaseValue }},
	{"armadura base", func(r *Record) float64 { return r.BaseArmor.BaseValue }},
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
