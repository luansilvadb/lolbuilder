package optimize

import (
	"fmt"
	"math"
	"sort"

	"github.com/luansilvadb/lolbuilder/internal/canon"
)

// A pergunta aqui e "qual o maximo de <stat> que cabe em 6 slots e N de ouro".
//
// NAO e "qual a melhor build". A diferenca nao e retorica: efeito de item nao
// entra no objetivo linear, entao o resultado ignora tudo que faz um item ser
// bom alem dos atributos dele. Uma build otima por stat linear nao e uma build
// boa, e o rotulo publicado tem de dizer isso — ver o comentario de Build.
//
// A solucao e exata: mochila 0/1 por programacao dinamica sobre slots e sobre o
// estado dos grupos de exclusividade, guardando em cada estado a fronteira de
// Pareto em (custo, valor). Verificada contra busca exaustiva num teste, que
// respeita os mesmos grupos — senao a verificacao seria circular.

// ItemCandidato e um item compravel com o que ele concede.
type ItemCandidato struct {
	ID    int32        `json:"id"`
	Nome  string       `json:"nome"`
	Custo int32        `json:"custo"`
	Stats canon.Vector `json:"stats,omitempty"`

	// Grupos sao os limites de posse a que o item esta sujeito. Botas era o
	// unico que este pacote conhecia, e a lista abaixo o substitui: o jogo
	// declara vinte e seis, e ignorar os outros vinte e cinco fazia a mochila
	// devolver o otimo exato de um conjunto viavel que nao existe.
	Grupos []string `json:"grupos,omitempty"`
}

// Grupo e um limite de posse que a mochila precisa respeitar.
type Grupo struct {
	ID     string `json:"id"`
	Maximo int    `json:"maximo"`
}

// maxEstadosDeGrupo limita a dimensao de grupos da grade.
//
// So entra o grupo com mais membros UTEIS que o limite dele sob o objetivo
// corrente, entao na pratica o numero e baixo — 8 no pior objetivo do 16.16.
// Passando daqui a resposta certa e recusar, nao aproximar: um otimo aproximado
// publicado como exato e o defeito que este mecanismo veio corrigir.
const maxEstadosDeGrupo = 4096

// maxPontosNaFronteira limita o tamanho total da fronteira de Pareto.
//
// Existe pelo mesmo motivo do limite de estados: se a instancia crescer a ponto
// de nao caber, a resposta certa e recusar. Medido no 16.16, o pico real fica
// tres ordens de grandeza abaixo deste teto.
const maxPontosNaFronteira = 2_000_000

// ponto e uma combinacao alcancavel: quanto custa, quanto vale e quais itens a
// compoem.
//
// A combinacao vai junto, e nao um ponteiro de volta. Ponteiro de volta parece a
// escolha obvia e ja esteve errado aqui: ele aponta para um estado que um item
// posterior sobrescreve, e a reconstrucao caminha por uma solucao que nao e a
// dele — foi assim que Presságio de Randuin e Couraça do Defunto sairam duas
// vezes na mesma build.
type ponto struct {
	custo int32
	valor float64
	itens []int // indices em uteis
}

// podar remove os pontos dominados e devolve a fronteira.
//
// Um ponto e dominado quando existe outro que custa o mesmo ou menos E vale o
// mesmo ou mais: nenhuma solucao otima passa por ele, entao descarta-lo nao
// perde resposta. E o que mantem a fronteira pequena.
func podar(ps []ponto) []ponto {
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].custo != ps[j].custo {
			return ps[i].custo < ps[j].custo
		}
		return ps[i].valor > ps[j].valor
	})

	out := ps[:0]
	melhor := math.Inf(-1)
	for _, p := range ps {
		if p.valor <= melhor {
			continue
		}
		melhor = p.valor
		out = append(out, p)
	}
	return out
}

// Build e o resultado: a combinacao de maior valor dentro das restricoes.
type Build struct {
	// Rotulo diz o que este numero e. Nunca "build otima".
	Rotulo string `json:"rotulo"`

	Itens []ItemCandidato `json:"itens"`
	Total canon.Vector    `json:"total"`
	Valor float64         `json:"valor"`
	Gasto int32           `json:"gasto"`

	Orcamento int32     `json:"orcamento"`
	Slots     int       `json:"slots"`
	Resolucao Resolucao `json:"resolucao_adaptativa"`
}

// MelhorBuild resolve a mochila.
//
// slots e o numero de espacos de inventario e orcamento o ouro disponivel. O
// resultado nao precisa gastar tudo nem encher todos os slots: um item a mais
// so entra se somar valor.
func MelhorBuild(itens []ItemCandidato, grupos []Grupo, obj Objetivo, slots int, orcamento int32) (Build, error) {
	if slots <= 0 {
		return Build{}, fmt.Errorf("numero de slots invalido: %d", slots)
	}
	if orcamento < 0 {
		return Build{}, fmt.Errorf("orcamento negativo: %d", orcamento)
	}

	// Descarta o que nunca ajuda: item sem valor sob este objetivo so ocuparia
	// slot, e item caro demais nunca cabe. Isso e poda de dominancia, e nao
	// heuristica — nenhuma solucao otima usa um deles.
	uteis := make([]ItemCandidato, 0, len(itens))
	for _, it := range itens {
		if it.Custo <= orcamento && obj.Valor(it.Stats) > 0 {
			uteis = append(uteis, it)
		}
	}
	sort.Slice(uteis, func(i, j int) bool { return uteis[i].ID < uteis[j].ID })

	restr, err := montarRestricoes(uteis, grupos)
	if err != nil {
		return Build{}, err
	}

	// A grade e por (slots usados, estado dos grupos), e cada celula guarda a
	// FRONTEIRA de Pareto em (custo, valor) — nao uma casa por moeda de ouro.
	//
	// A versao anterior indexava o ouro diretamente, o que dava 20001 casas. Com
	// uma dimensao so de botas isso cabia; com os 26 grupos reais do jogo sao ate
	// 256 estados, e 7 x 20001 x 256 celulas nao cabem na memoria. A fronteira
	// resolve porque o que importa nao e cada valor de ouro possivel, e sim os
	// poucos pontos em que gastar mais passa a render mais — e com 6 slots eles
	// sao poucos.
	//
	// Continua exato: um ponto so e descartado quando existe outro que custa o
	// mesmo ou menos E vale o mesmo ou mais.
	frente := make([][][]ponto, slots+1)
	for s := range frente {
		frente[s] = make([][]ponto, restr.estados)
	}
	frente[0][0] = []ponto{{}}

	// Slots em ordem DECRESCENTE: a transicao so escreve em s+1, entao cada item
	// entra no maximo uma vez.
	total := 1
	for idx, it := range uteis {
		valor := obj.Valor(it.Stats)
		custo := it.Custo

		for s := slots - 1; s >= 0; s-- {
			for e := 0; e < restr.estados; e++ {
				atual := frente[s][e]
				if len(atual) == 0 {
					continue
				}
				// O item so entra se couber em TODOS os grupos dele: um item
				// pode estar em mais de um, e o Terminus esta em dois.
				ne, ok := restr.somar(e, idx)
				if !ok {
					continue
				}

				novos := make([]ponto, 0, len(atual))
				for _, p := range atual {
					if p.custo+custo > orcamento {
						continue
					}
					comb := make([]int, len(p.itens), len(p.itens)+1)
					copy(comb, p.itens)
					novos = append(novos, ponto{
						custo: p.custo + custo,
						valor: p.valor + valor,
						itens: append(comb, idx),
					})
				}
				if len(novos) == 0 {
					continue
				}
				total -= len(frente[s+1][ne])
				frente[s+1][ne] = podar(append(frente[s+1][ne], novos...))
				total += len(frente[s+1][ne])

				if total > maxPontosNaFronteira {
					return Build{}, fmt.Errorf(
						"a fronteira de Pareto passou de %d pontos — a mochila recusa em vez "+
							"de aproximar, porque um otimo aproximado publicado como exato e pior "+
							"que nenhum", maxPontosNaFronteira)
				}
			}
		}
	}

	// O melhor entre todos os estados alcancaveis.
	var melhor ponto
	for s := 0; s <= slots; s++ {
		for e := 0; e < restr.estados; e++ {
			for _, p := range frente[s][e] {
				if p.valor > melhor.valor {
					melhor = p
				}
			}
		}
	}
	gasto := melhor.custo

	b := Build{
		Rotulo:    rotuloDe(obj),
		Total:     canon.Vector{},
		Valor:     melhor.valor,
		Gasto:     gasto,
		Orcamento: orcamento,
		Slots:     slots,
		Resolucao: obj.Resolucao,
	}
	for _, idx := range melhor.itens {
		it := uteis[idx]
		b.Itens = append(b.Itens, it)
		b.Total.Merge(it.Stats)
	}
	sort.Slice(b.Itens, func(i, j int) bool { return b.Itens[i].ID < b.Itens[j].ID })
	return b, nil
}

// restricoes codifica os grupos ATIVOS num estado inteiro de base mista.
//
// Ativo e o grupo que tem mais membros UTEIS que o limite dele. Os outros nao
// podem ser violados nem que se queira, e carregar dimensao para eles so
// multiplicaria a grade a troco de nada: dos 26 grupos do 16.16, quase todos
// somem depois da poda de dominancia, porque um objetivo linear costuma dar
// valor a poucos itens de cada grupo.
type restricoes struct {
	estados int
	// porItem[i] sao os pares (indice do grupo ativo, peso posicional) do item i.
	porItem [][]par
	// limite[g] e quantos cabem no grupo ativo g, e passo[g] a casa dele.
	limite []int
	passo  []int
}

type par struct{ grupo, passo int }

func montarRestricoes(uteis []ItemCandidato, grupos []Grupo) (*restricoes, error) {
	maxDe := make(map[string]int, len(grupos))
	for _, g := range grupos {
		maxDe[g.ID] = g.Maximo
	}

	uteisNo := map[string]int{}
	for _, it := range uteis {
		for _, g := range it.Grupos {
			if _, ok := maxDe[g]; ok {
				uteisNo[g]++
			}
		}
	}

	var ativos []string
	for g, n := range uteisNo {
		if n > maxDe[g] {
			ativos = append(ativos, g)
		}
	}
	sort.Strings(ativos) // ordem estavel: a grade tem de ser reproduzivel

	r := &restricoes{estados: 1, porItem: make([][]par, len(uteis))}
	indice := map[string]int{}
	for _, g := range ativos {
		radix := maxDe[g] + 1
		if r.estados > maxEstadosDeGrupo/radix {
			return nil, fmt.Errorf(
				"os grupos de exclusividade ativos exigem mais de %d estados na grade "+
					"(%d grupos: %v) — a mochila recusa em vez de aproximar, porque um "+
					"otimo aproximado publicado como exato e pior que nenhum",
				maxEstadosDeGrupo, len(ativos), ativos)
		}
		indice[g] = len(r.limite)
		r.passo = append(r.passo, r.estados)
		r.limite = append(r.limite, maxDe[g])
		r.estados *= radix
	}

	for i, it := range uteis {
		for _, g := range it.Grupos {
			if j, ok := indice[g]; ok {
				r.porItem[i] = append(r.porItem[i], par{grupo: j, passo: r.passo[j]})
			}
		}
	}
	return r, nil
}

// somar tenta acrescentar o item ao estado e diz se ainda cabe.
func (r *restricoes) somar(estado, item int) (int, bool) {
	for _, p := range r.porItem[item] {
		usado := (estado / p.passo) % (r.limite[p.grupo] + 1)
		if usado >= r.limite[p.grupo] {
			return 0, false
		}
		estado += p.passo
	}
	return estado, true
}

// rotuloDe monta o rotulo do resultado.
//
// O texto e deliberadamente longo. "Build otima" seria mais curto e mentiria: o
// calculo ignora passiva e ativa de item, que e o que faz metade dos itens
// valerem o que valem. Um modelo de linguagem lendo "otima" repassaria a
// recomendacao com a autoridade do dataset.
func rotuloDe(obj Objetivo) string {
	return fmt.Sprintf("maximo de %s por ouro, ignorando efeitos de item", obj)
}
