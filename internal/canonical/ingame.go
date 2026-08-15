package canonical

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// A comparacao com a partida existe porque nenhuma fonte publica o que o jogo
// faz em tempo de execucao. O dump declara base e crescimento; se a formula que
// aplicamos sobre eles estiver errada, nada no dado denuncia.
//
// O que se compara e o CRESCIMENTO entre dois niveis, e nao o valor absoluto.
//
// Valor absoluto e contaminado por tudo: item comprado, fragmento de stat, runa
// que da atributo. Crescimento entre dois niveis cancela todo bonus FIXO, entao
// resta so o que o campeao ganha por subir de nivel — que e exatamente a
// formula que o dataset afirma.
//
// Uma excecao importa no LoL moderno, e nao existia no modo Jade: o fragmento
// Escalamento de Vida cresce COM o nivel, +10 por nivel, e o jogador pode levar
// dois. Entao o crescimento de vida pode vir ate 20 por nivel acima do previsto
// sem que nada esteja errado. Os outros eixos ficam limpos.

// AmostraIngame e uma leitura da partida, ja no vocabulario do dataset.
type AmostraIngame struct {
	Nivel   int      `json:"nivel"`
	Campeao string   `json:"campeao"`
	Itens   []string `json:"itens,omitempty"`

	MaxHealth       float64 `json:"max_health"`
	Armor           float64 `json:"armor"`
	MagicResist     float64 `json:"magic_resist"`
	AttackDamage    float64 `json:"attack_damage"`
	AttackSpeed     float64 `json:"attack_speed"`
	HealthRegenRate float64 `json:"health_regen_rate"`
}

// Veredito e a conclusao de uma comparacao.
type Veredito string

const (
	// VereditoBate: o jogo e o dataset concordam dentro da tolerancia.
	VereditoBate Veredito = "bate"
	// VereditoDiverge: discordam, e a diferenca nao tem explicacao conhecida.
	VereditoDiverge Veredito = "DIVERGE"
	// VereditoContaminado: a diferenca cabe dentro de um bonus conhecido, entao
	// a leitura nao conclui nada — nem a favor nem contra.
	VereditoContaminado Veredito = "inconclusivo"
)

// ComparacaoIngame e uma linha do relatorio.
type ComparacaoIngame struct {
	Eixo     string   `json:"eixo"`
	DeNivel  int      `json:"de_nivel"`
	AteNivel int      `json:"ate_nivel"`
	Jogo     float64  `json:"jogo"`
	Previsto float64  `json:"previsto"`
	Diff     float64  `json:"diferenca"`
	Veredito Veredito `json:"veredito"`
	Nota     string   `json:"nota,omitempty"`
}

// RelatorioIngame resume o confronto.
type RelatorioIngame struct {
	Campeao     string             `json:"campeao"`
	Niveis      []int              `json:"niveis"`
	Comparacoes []ComparacaoIngame `json:"comparacoes"`
	// ItensMudaram avisa que a lista de itens nao e a mesma nas duas amostras,
	// o que quebra o cancelamento do bonus fixo.
	ItensMudaram bool `json:"itens_mudaram"`
}

// Diverge informa se alguma comparacao acusou divergencia real.
func (r *RelatorioIngame) Diverge() bool {
	for _, c := range r.Comparacoes {
		if c.Veredito == VereditoDiverge {
			return true
		}
	}
	return false
}

// escalamentoDeVidaPorNivel e quanto o fragmento Escalamento de Vida cresce por
// nivel, e quantos o jogador pode levar.
//
// Sao os unicos numeros do jogo moderno que contaminam o CRESCIMENTO. Ficam
// aqui, e nao na curadoria, porque a curadoria descreve o que a runa concede e
// isto e a margem que a comparacao precisa tolerar.
const (
	escalamentoDeVidaPorNivel = 10.0
	maximoDeFragmentosDeVida  = 2
)

// CompararIngame confronta o dataset com as amostras da partida.
//
// Precisa de pelo menos duas amostras em niveis diferentes: com uma so da para
// conferir valor absoluto, que nao conclui nada, e nao a formula.
func CompararIngame(ds *Dataset, amostras []AmostraIngame, tolerancia float64) (*RelatorioIngame, error) {
	if len(amostras) < 2 {
		return nil, fmt.Errorf(
			"sao necessarias ao menos 2 amostras em niveis diferentes: o valor absoluto "+
				"e contaminado por item e runa, e so o crescimento entre niveis prova a formula "+
				"(ha %d)", len(amostras))
	}

	campeao := amostras[0].Campeao
	for _, a := range amostras {
		if a.Campeao != campeao {
			return nil, fmt.Errorf(
				"as amostras sao de campeoes diferentes (%s e %s) — o crescimento so "+
					"cancela o bonus fixo dentro da mesma partida", campeao, a.Campeao)
		}
	}

	c := acharCampeao(ds, campeao)
	if c == nil || c.Stats == nil {
		return nil, fmt.Errorf("campeao %q nao esta no dataset, ou saiu sem estatisticas", campeao)
	}

	ord := append([]AmostraIngame(nil), amostras...)
	sort.Slice(ord, func(i, j int) bool { return ord[i].Nivel < ord[j].Nivel })

	rel := &RelatorioIngame{Campeao: campeao}
	for _, a := range ord {
		rel.Niveis = append(rel.Niveis, a.Nivel)
	}

	// Compara cada par consecutivo. Pares e nao extremos: um erro que so aparece
	// numa faixa de nivel some quando se olha so o comeco e o fim.
	for i := 1; i < len(ord); i++ {
		antes, agora := ord[i-1], ord[i]
		if agora.Nivel == antes.Nivel {
			continue
		}
		if !mesmosItens(antes.Itens, agora.Itens) {
			rel.ItensMudaram = true
		}
		niveis := float64(agora.Nivel - antes.Nivel)

		eixos := []struct {
			nome     string
			jogo     float64
			previsto float64
			margem   float64
			nota     string
		}{
			{"vida", agora.MaxHealth - antes.MaxHealth,
				c.Stats.HP.CrescimentoEntre(antes.Nivel, agora.Nivel),
				escalamentoDeVidaPorNivel * maximoDeFragmentosDeVida,
				"o fragmento Escalamento de Vida cresce com o nivel, e cabem dois"},
			{"armadura", agora.Armor - antes.Armor,
				c.Stats.Armor.CrescimentoEntre(antes.Nivel, agora.Nivel), 0, ""},
			{"resistencia magica", agora.MagicResist - antes.MagicResist,
				c.Stats.MagicResist.CrescimentoEntre(antes.Nivel, agora.Nivel), 0, ""},
			{"dano de ataque", agora.AttackDamage - antes.AttackDamage,
				c.Stats.AttackDamage.CrescimentoEntre(antes.Nivel, agora.Nivel), 0, ""},
		}

		for _, e := range eixos {
			previsto := e.previsto
			diff := e.jogo - previsto
			cmp := ComparacaoIngame{
				Eixo: e.nome, DeNivel: antes.Nivel, AteNivel: agora.Nivel,
				Jogo: e.jogo, Previsto: previsto, Diff: diff,
			}

			// Os dois sentidos acusam, por motivos diferentes.
			//
			// ABAIXO do previsto nunca tem explicacao inocente, porque bonus
			// soma e nunca subtrai — vale mesmo com a lista de itens mudada, e
			// foi assim que o crescimento nao linear apareceu: a resistencia
			// magica do Rammus subiu 8.0975 onde o modelo linear previa 10.25,
			// com cinco itens equipados no caminho.
			//
			// ACIMA do previsto so e explicavel por algo que some. Com os itens
			// iguais e sem runa que escale naquele eixo, nao ha o que some, e
			// tambem e divergencia. Tratar todo excesso como inconclusivo
			// deixaria a verificacao cega para metade dos erros possiveis.
			switch {
			case math.Abs(diff) <= tolerancia:
				cmp.Veredito = VereditoBate
			case diff < -tolerancia:
				cmp.Veredito = VereditoDiverge
				cmp.Nota = "crescer abaixo do previsto nao tem explicacao: bonus soma, nunca subtrai"
			case rel.ItensMudaram:
				cmp.Veredito = VereditoContaminado
				cmp.Nota = "a lista de itens mudou entre as duas leituras"
			case diff <= e.margem*niveis+tolerancia:
				cmp.Veredito = VereditoContaminado
				cmp.Nota = e.nota
			default:
				cmp.Veredito = VereditoDiverge
				cmp.Nota = "crescimento acima do previsto sem item novo nem runa que escale nesse eixo"
			}
			rel.Comparacoes = append(rel.Comparacoes, cmp)
		}
	}
	return rel, nil
}

// acharCampeao casa o nome que a partida reporta com o do dataset.
//
// A partida usa o nome de exibicao no idioma do cliente, e o dataset guarda o
// nome traduzido e o canonico. Tenta os dois.
func acharCampeao(ds *Dataset, nome string) *Champion {
	for i := range ds.Champions {
		c := &ds.Champions[i]
		if strings.EqualFold(c.Nome, nome) || strings.EqualFold(c.NomeCanonico, nome) ||
			strings.EqualFold(c.Alias, nome) {
			return c
		}
	}
	return nil
}

func mesmosItens(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	x := append([]string(nil), a...)
	y := append([]string(nil), b...)
	sort.Strings(x)
	sort.Strings(y)
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
