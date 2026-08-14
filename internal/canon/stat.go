// Package canon define o vocabulario canonico de stats e o vetor somavel.
//
// O conjunto e fechado: um rotulo fora dele e erro, nunca zero. Somar sem
// normalizar produz totais errados em silencio, que e o pior modo de falha
// possivel num sistema que se propoe exato.
package canon

import (
	"fmt"
	"sort"
	"strings"
)

// Stat e um identificador do vocabulario canonico.
type Stat string

// CONVENCAO DE UNIDADES
//
// A unidade e propriedade do stat, e nao do valor.
//
//   - Stats planos (armadura, vida, dano de ataque, mana, letalidade...) usam o
//     valor absoluto do jogo.
//   - Stats percentuais usam PONTOS PERCENTUAIS: 45% vira 45, nao 0.45.
//
// Onde a fonte publica a MESMA grandeza nas duas formas, sao dois stats
// distintos e eles nunca somam no mesmo campo. No 16.16 isso acontece com
// penetracao magica (4 itens planos, 4 percentuais) e com velocidade de
// movimento (15 planos, 20 percentuais). Colapsar os dois num stat so somaria
// 45 de deslocamento com 4% de deslocamento — erro grande, silencioso, e
// exatamente o que este pacote existe para impedir.
const (
	AbilityHaste        Stat = "ability_haste"
	AbilityPower        Stat = "ability_power"
	Armor               Stat = "armor"
	ArmorPenetrationPct Stat = "armor_penetration_pct"
	AttackDamage        Stat = "attack_damage"
	AttackSpeedPct      Stat = "attack_speed_pct"
	BaseHealthRegenPct  Stat = "base_health_regen_pct"
	BaseManaRegenPct    Stat = "base_mana_regen_pct"
	CriticalChancePct   Stat = "critical_chance_pct"
	CriticalDamagePct   Stat = "critical_damage_pct"
	GoldPer10           Stat = "gold_per_10"
	HealShieldPowerPct  Stat = "heal_shield_power_pct"
	Health              Stat = "health"
	Lethality           Stat = "lethality"
	LifeStealPct        Stat = "life_steal_pct"
	MagicPenetration    Stat = "magic_penetration"
	MagicPenetrationPct Stat = "magic_penetration_pct"
	MagicResist         Stat = "magic_resist"
	Mana                Stat = "mana"
	MoveSpeed           Stat = "move_speed"
	MoveSpeedPct        Stat = "move_speed_pct"
	OmnivampPct         Stat = "omnivamp_pct"
	TenacityPct         Stat = "tenacity_pct"
)

// All lista os stats canonicos em ordem estavel, para saidas reproduziveis.
var All = []Stat{
	AbilityHaste, AbilityPower, Armor, ArmorPenetrationPct, AttackDamage,
	AttackSpeedPct, BaseHealthRegenPct, BaseManaRegenPct, CriticalChancePct,
	CriticalDamagePct, GoldPer10, HealShieldPowerPct, Health, Lethality,
	LifeStealPct, MagicPenetration, MagicPenetrationPct, MagicResist, Mana,
	MoveSpeed, MoveSpeedPct, OmnivampPct, TenacityPct,
}

// Percentual informa se o stat e medido em pontos percentuais.
func (s Stat) Percentual() bool {
	return strings.HasSuffix(string(s), "_pct")
}

// rotulo identifica um stat pelo texto da fonte E pela unidade em que ele veio.
//
// A unidade faz parte da chave, e nao e conferida depois, porque a fonte usa o
// mesmo rotulo para grandezas diferentes: "Magic Penetration" plano e "Magic
// Penetration" percentual sao linhas distintas do mesmo bloco.
type rotulo struct {
	texto      string
	percentual bool
}

// rotulos mapeia cada rotulo observado na fonte para o stat canonico.
//
// Os 23 do patch 16.16, medidos sobre as 475 linhas de stat dos 210 itens
// compraveis. O mapa e a definicao do que o projeto sabe ler: rotulo ausente
// aqui vira lacuna declarada, nunca zero.
var rotulos = map[rotulo]Stat{
	{"ability haste", false}:         AbilityHaste,
	{"ability power", false}:         AbilityPower,
	{"armor", false}:                 Armor,
	{"armor penetration", true}:      ArmorPenetrationPct,
	{"attack damage", false}:         AttackDamage,
	{"attack speed", true}:           AttackSpeedPct,
	{"base health regen", true}:      BaseHealthRegenPct,
	{"base mana regen", true}:        BaseManaRegenPct,
	{"critical strike chance", true}: CriticalChancePct,
	{"critical strike damage", true}: CriticalDamagePct,
	{"gold per 10 seconds", false}:   GoldPer10,
	{"heal and shield power", true}:  HealShieldPowerPct,
	{"health", false}:                Health,
	{"lethality", false}:             Lethality,
	{"life steal", true}:             LifeStealPct,
	{"magic penetration", false}:     MagicPenetration,
	{"magic penetration", true}:      MagicPenetrationPct,
	{"magic resist", false}:          MagicResist,
	{"mana", false}:                  Mana,
	{"move speed", false}:            MoveSpeed,
	{"move speed", true}:             MoveSpeedPct,
	{"omnivamp", true}:               OmnivampPct,
	{"tenacity", true}:               TenacityPct,
}

// LookupStat resolve um rotulo da fonte, junto da unidade em que ele veio.
func LookupStat(texto string, percentual bool) (Stat, bool) {
	s, ok := rotulos[rotulo{strings.ToLower(strings.TrimSpace(texto)), percentual}]
	return s, ok
}

// Vector e um conjunto de stats somaveis.
//
// E um mapa, e nao um struct com um campo por stat, porque stat ausente e stat
// zero precisam ser distinguiveis: um item que nao da armadura e diferente de
// um item que da zero de armadura, e so o primeiro existe.
type Vector map[Stat]float64

// Add soma um valor ao stat, criando a entrada se necessario.
func (v Vector) Add(s Stat, valor float64) {
	v[s] += valor
}

// Merge soma outro vetor neste.
func (v Vector) Merge(o Vector) {
	for s, valor := range o {
		v[s] += valor
	}
}

// Stats devolve os stats presentes em ordem canonica, para saida reproduzivel.
//
// A ordem vem de All, e nao de sort sobre as chaves, porque All e a ordem em
// que o dataset publica — mudar uma implicaria mudar a outra em silencio.
func (v Vector) Stats() []Stat {
	presentes := make(map[Stat]bool, len(v))
	for s := range v {
		presentes[s] = true
	}
	out := make([]Stat, 0, len(v))
	for _, s := range All {
		if presentes[s] {
			out = append(out, s)
			delete(presentes, s)
		}
	}
	// Um stat presente no vetor e ausente de All e defeito de programacao, nao
	// de dado: o vocabulario e fechado. Sai no fim, ordenado, para que apareca
	// em vez de desaparecer.
	if len(presentes) > 0 {
		resto := make([]Stat, 0, len(presentes))
		for s := range presentes {
			resto = append(resto, s)
		}
		sort.Slice(resto, func(i, j int) bool { return resto[i] < resto[j] })
		out = append(out, resto...)
	}
	return out
}

// Validate confere que todo stat do vetor pertence ao vocabulario.
func (v Vector) Validate() error {
	conhecido := make(map[Stat]bool, len(All))
	for _, s := range All {
		conhecido[s] = true
	}
	var fora []string
	for s := range v {
		if !conhecido[s] {
			fora = append(fora, string(s))
		}
	}
	if len(fora) == 0 {
		return nil
	}
	sort.Strings(fora)
	return fmt.Errorf("stat(s) fora do vocabulario canonico: %s", strings.Join(fora, ", "))
}
