package gamedata

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Slot e a posicao de uma habilidade na barra.
type Slot int

// A ordem e a mesma do array CharacterRecord.spells.
const (
	SlotQ Slot = iota
	SlotW
	SlotE
	SlotR
	SlotPassive
)

func (s Slot) String() string {
	switch s {
	case SlotQ:
		return "Q"
	case SlotW:
		return "W"
	case SlotE:
		return "E"
	case SlotR:
		return "R"
	case SlotPassive:
		return "P"
	}
	return "?"
}

// IsUltimate informa se o slot e o da habilidade suprema, que tem 3 ranks em
// vez de 5.
func (s Slot) IsUltimate() bool { return s == SlotR }

// DataValue e uma serie de valores indexada por rank.
type DataValue struct {
	Name   string    `json:"name"`
	Values []float64 `json:"values"`
}

// Spell e o mSpell de um SpellObject, reduzido ao que sabemos ler.
//
// Diferente do CharacterRecord, aqui nao ha DecodeStrict: o conjunto de campos
// de um mSpell varia muito entre habilidades, e a vigilancia desta parte da
// fonte e a cobertura de extracao, nao o alarme de campo novo.
type Spell struct {
	Slot   Slot
	Name   string // nome do objeto, ex: GarenQ
	Script string

	DataValues   []DataValue                `json:"DataValues"`
	Calculations map[string]json.RawMessage `json:"mSpellCalculations"`
	CastRange    []float64                  `json:"castRange"`

	// CastRangeValues e a serie redefinida do alcance. NAO e publicada: a
	// evidencia nao decide qual das duas o jogo usa, e a divergencia entre elas
	// e declarada como incerteza em vez de resolvida no chute.
	CastRangeValues serieRedefinida `json:"castRangeValues"`

	// A fonte traz DUAS series para recarga e para custo: uma herdada e outra
	// redefinida. O que o jogo usa e a redefinida.
	//
	// Medido no 16.16 contra o arquivo do plugin, que e fonte independente: a
	// redefinida concorda em 680 de 680 recargas e 568 de 568 custos — 100%. A
	// herdada fica em 98.5% e 98.8%, divergindo em 10 e 7 habilidades.
	//
	// O nome engana: cooldownTime e mana sao os nomes classicos do formato BIN,
	// e Cooldown e manaValues parecem estrutura auxiliar. A escolha nao veio da
	// forma do campo, veio da concordancia medida.
	Cooldown   serieRedefinida `json:"Cooldown"`
	ManaValues serieRedefinida `json:"manaValues"`

	// Herdadas. Ficam mapeadas para nao sumirem do alarme de campo novo, e
	// porque a comparacao entre as duas e o que detecta a fonte mudar de forma.
	ManaHerdada     []float64 `json:"mana"`
	CooldownHerdada []float64 `json:"cooldownTime"`
}

// serieRedefinida e o formato em que a fonte publica o valor redefinido.
type serieRedefinida struct {
	Values []float64 `json:"values"`
}

// Mana devolve o custo por rank que o jogo usa.
//
// Sem a serie redefinida a habilidade sai sem custo, e nao com a herdada no
// lugar: publicar a herdada como se fosse a boa e exatamente o defeito que esta
// escolha evita.
func (s *Spell) Mana() []float64 { return s.ManaValues.Values }

// CooldownTime devolve a recarga por rank que o jogo usa.
func (s *Spell) CooldownTime() []float64 { return s.Cooldown.Values }

// DataValue busca uma serie por nome. As formulas referenciam os valores por
// nome, nao por posicao.
func (s *Spell) DataValue(name string) ([]float64, bool) {
	for _, dv := range s.DataValues {
		if strings.EqualFold(dv.Name, name) {
			if len(dv.Values) == 0 {
				return []float64{0, 0, 0, 0, 0, 0, 0}, true
			}
			return dv.Values, true
		}
	}
	return nil, false
}

// ConsumedDataValues devolve os nomes das series que as formulas desta
// habilidade referenciam, em minusculas.
//
// Serve para nao publicar duas vezes o mesmo numero: uma habilidade que resolve
// TotalDamage a partir de BaseDamage, ADRatio e APRatio apresentaria o mesmo
// dano quatro vezes se as tres series saissem ao lado do efeito resolvido, como
// se fossem parcelas somaveis.
//
// A coleta e sobre o campo mDataValue, e nao por busca de texto: o nome de uma
// serie ("Cost", "Duration") aparece em muitos outros lugares do dump, e casar
// por texto suprimiria series legitimas.
func (s *Spell) ConsumedDataValues() map[string]bool {
	out := map[string]bool{}
	for _, raw := range s.Calculations {
		var v any
		if err := json.Unmarshal(raw, &v); err != nil {
			continue
		}
		coletaDataValues(v, out)
	}
	return out
}

func coletaDataValues(v any, out map[string]bool) {
	switch t := v.(type) {
	case map[string]any:
		for k, val := range t {
			if k == "mDataValue" {
				if s, ok := val.(string); ok && s != "" {
					out[strings.ToLower(s)] = true
				}
				continue
			}
			coletaDataValues(val, out)
		}
	case []any:
		for _, val := range t {
			coletaDataValues(val, out)
		}
	}
}

// Derivacao e a relacao que a fonte declara entre um calculo e outro.
type Derivacao struct {
	// De e o calculo de origem, como a fonte o nomeia. Pode ser opaco.
	De string
	// Multiplicador e 1 quando a fonte nao declara nenhum.
	Multiplicador float64
}

// DerivacaoDe devolve a relacao declarada para um calculo, quando existir.
//
// Serve para rotular um calculo cujo nome nao e recuperavel: publicar
// "TotalDPS x 0.25" nao inventa nome nenhum, transcreve o que a fonte declara.
// Sem a relacao, a habilidade publica dois numeros para o mesmo dano e somar os
// dois e a leitura natural.
func (s *Spell) DerivacaoDe(nome string) (Derivacao, bool) {
	raw, ok := s.Calculations[nome]
	if !ok {
		return Derivacao{}, false
	}

	var c struct {
		Type       string `json:"__type"`
		Modified   string `json:"mModifiedGameCalculation"`
		Multiplier struct {
			Number *float64 `json:"mNumber"`
		} `json:"mMultiplier"`
	}
	if err := json.Unmarshal(raw, &c); err != nil {
		return Derivacao{}, false
	}
	if c.Type != "GameCalculationModified" || c.Modified == "" {
		return Derivacao{}, false
	}

	// Multiplicador ausente vale 1: a fonte omite mNumber quando o calculo e o
	// outro sem escala, e tratar a omissao como zero apagaria o valor.
	d := Derivacao{De: c.Modified, Multiplicador: 1}
	if c.Multiplier.Number != nil {
		d.Multiplicador = *c.Multiplier.Number
	}
	return d, true
}

// spellObject e o envelope que carrega o mSpell.
type spellObject struct {
	ObjectName string          `json:"ObjectName"`
	ScriptName string          `json:"mScriptName"`
	Spell      json.RawMessage `json:"mSpell"`
	Type       string          `json:"__type"`
}

// Spells resolve as quatro habilidades de slot mais a passiva.
//
// Os slots saem do array ordenado CharacterRecord.spells, e nunca do nome da
// habilidade: Lux chama as dela de LightBinding e PrismaticWave, Jarvan IV de
// DragonStrike e GoldenAegis. Adivinhar pelo sufixo Q/W/E/R deixaria esses
// campeoes sem nenhuma habilidade, e em silencio.
func (d *Dump) Spells() ([]Spell, error) {
	if n := len(d.Record.Spells); n != 4 {
		return nil, fmt.Errorf("%s: %d habilidades em CharacterRecord.spells, esperado 4",
			d.Alias, n)
	}

	out := make([]Spell, 0, 5)
	for i, path := range d.Record.Spells {
		sp, err := d.spellAt(path)
		if err != nil {
			return nil, fmt.Errorf("%s slot %s: %w", d.Alias, Slot(i), err)
		}
		sp.Slot = Slot(i)
		out = append(out, *sp)
	}

	// A passiva nao esta no array de slots. O CharacterRecord aponta para ela
	// por um identificador que costuma ser opaco — e que e, ele proprio, uma
	// chave de topo do dump.
	if ref := d.Record.PassiveSpell; ref != "" {
		sp, err := d.spellAt(ref)
		if err != nil {
			return nil, fmt.Errorf("%s passiva: %w", d.Alias, err)
		}
		sp.Slot = SlotPassive
		out = append(out, *sp)
	}
	return out, nil
}

// SpellByObjectName acha uma habilidade pelo nome do objeto, ignorando caixa.
//
// Serve para as habilidades que NAO estao no array de slots: o livro de feiticos
// do Hwei publica doze sub-habilidades como SpellObject proprio, cada uma com
// dano e recarga, e o CharacterRecord nao aponta para nenhuma delas.
func (d *Dump) SpellByObjectName(nome string) (*Spell, bool) {
	for ref, raw := range d.objects {
		var probe struct {
			ObjectName string `json:"ObjectName"`
			Type       string `json:"__type"`
		}
		if err := json.Unmarshal(raw, &probe); err != nil {
			continue
		}
		if probe.Type != "SpellObject" || !strings.EqualFold(probe.ObjectName, nome) {
			continue
		}
		sp, err := d.spellAt(ref)
		if err != nil {
			return nil, false
		}
		return sp, true
	}
	return nil, false
}

// spellAt resolve uma referencia de habilidade, seja ela caminho ou hash.
func (d *Dump) spellAt(ref string) (*Spell, error) {
	raw, ok := d.Object(ref)
	if !ok {
		return nil, fmt.Errorf("referencia %q nao existe no dump", ref)
	}

	var obj spellObject
	if err := json.Unmarshal(raw, &obj); err != nil {
		return nil, fmt.Errorf("referencia %q: %w", ref, err)
	}
	if obj.Type != "SpellObject" {
		return nil, fmt.Errorf("referencia %q aponta para %s, nao SpellObject", ref, obj.Type)
	}

	var sp Spell
	if len(obj.Spell) > 0 {
		if err := json.Unmarshal(obj.Spell, &sp); err != nil {
			return nil, fmt.Errorf("mSpell de %q: %w", ref, err)
		}
	}
	sp.Name = obj.ObjectName
	sp.Script = obj.ScriptName
	if sp.Name == "" {
		sp.Name = ref
	}
	return &sp, nil
}

// Ranks e a quantidade de ranks que uma habilidade realmente tem.
//
// Os arrays da fonte sao maiores que isso: sobra preenchimento que repete o
// ultimo valor ou traz zero. Publicar o preenchimento inventaria um rank 6 que
// nao existe no jogo.
func (s *Spell) Ranks() int {
	switch s.Slot {
	case SlotPassive:
		return 1
	case SlotR:
		return 3
	default:
		return 5
	}
}

// Tamanhos das series do dump. Nao sao convencao nossa: sao o que a fonte
// serializa, e cada tamanho vem com sua propria indexacao.
const (
	// RankedLen e o tamanho das series que reservam o indice 0 para o estado
	// "habilidade nao aprendida": DataValues e cooldownTime. Nelas o rank 1
	// esta no indice 1.
	RankedLen = 7
	// PluginAlignedLen e o tamanho das series que usam a mesma indexacao do
	// arquivo do plugin, com o rank 1 no indice 0. E o caso de mana.
	PluginAlignedLen = 6
)

// RankOffset devolve o deslocamento de uma serie a partir do seu tamanho.
//
// As duas series do mesmo mSpell nao concordam entre si: cooldownTime tem 7
// posicoes e comeca no indice 1, mana tem 6 e comeca no indice 0. Presumir um
// deslocamento unico deslocaria metade dos numeros em um rank, para sempre e
// sem sintoma visivel — por isso o build confere os dois contra o arquivo do
// plugin antes de publicar qualquer coisa.
func RankOffset(length int) int {
	if length >= RankedLen {
		return 1
	}
	return 0
}

// AtRank le uma serie do dump no rank pedido (1..N), aplicando o deslocamento
// proprio do tamanho da serie.
func AtRank(values []float64, rank int) (float64, bool) {
	i := rank - 1 + RankOffset(len(values))
	if i < 0 || i >= len(values) {
		return 0, false
	}
	return values[i], true
}
