package optimize

import (
	"fmt"
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
// A solucao e exata: mochila 0/1 por programacao dinamica sobre slots e ouro,
// com uma dimensao a mais para a regra de botas unicas. Verificada contra busca
// exaustiva num teste.

// ItemCandidato e um item compravel com o que ele concede.
type ItemCandidato struct {
	ID    int32        `json:"id"`
	Nome  string       `json:"nome"`
	Custo int32        `json:"custo"`
	Stats canon.Vector `json:"stats,omitempty"`
	Botas bool         `json:"botas,omitempty"`
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
func MelhorBuild(itens []ItemCandidato, obj Objetivo, slots int, orcamento int32) (Build, error) {
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

	// A celula carrega a COMBINACAO inteira, e nao um ponteiro de volta.
	//
	// Ponteiro de volta parece a escolha obvia e esta errado aqui: ele aponta
	// para o estado da camada em que o item entrou, e essa celula pode ser
	// sobrescrita por um item posterior. Reconstruir pela grade final entao
	// caminha por um estado que ja e de outra solucao, e o item se repete — foi
	// exatamente o que aconteceu, com Presságio de Randuin e Couraça do Defunto
	// saindo duas vezes na mesma build.
	//
	// Guardar a combinacao custa pouco porque ela tem no maximo `slots` itens,
	// e elimina a classe inteira de erro.
	const semBotas, comBotas = 0, 1
	type celula struct {
		valor float64
		usado bool
		itens []int // indices em uteis
	}
	dim := int(orcamento) + 1
	grade := make([][][2]celula, slots+1)
	for s := range grade {
		grade[s] = make([][2]celula, dim)
	}
	grade[0][0][semBotas] = celula{usado: true}

	// Mochila 0/1 em grade unica. O slot percorrido em ordem DECRESCENTE e o
	// que garante que cada item entra no maximo uma vez: a transicao so escreve
	// em s+1, e s+1 ja passou quando se le em s.
	for idx, it := range uteis {
		valor := obj.Valor(it.Stats)
		custo := int(it.Custo)

		for s := slots - 1; s >= 0; s-- {
			for g := 0; g+custo < dim; g++ {
				for b := semBotas; b <= comBotas; b++ {
					c := grade[s][g][b]
					if !c.usado {
						continue
					}
					// Botas sao unicas: o jogo nao deixa carregar duas.
					nb := b
					if it.Botas {
						if b == comBotas {
							continue
						}
						nb = comBotas
					}
					ns, ng := s+1, g+custo
					novoValor := c.valor + valor
					if alvo := grade[ns][ng][nb]; alvo.usado && novoValor <= alvo.valor {
						continue
					}
					combinacao := make([]int, len(c.itens), len(c.itens)+1)
					copy(combinacao, c.itens)
					grade[ns][ng][nb] = celula{
						valor: novoValor, usado: true, itens: append(combinacao, idx),
					}
				}
			}
		}
	}

	// O melhor entre todos os estados alcancaveis.
	melhor := celula{usado: true}
	gasto := 0
	for s := 0; s <= slots; s++ {
		for g := 0; g < dim; g++ {
			for b := semBotas; b <= comBotas; b++ {
				if c := grade[s][g][b]; c.usado && c.valor > melhor.valor {
					melhor, gasto = c, g
				}
			}
		}
	}

	b := Build{
		Rotulo:    rotuloDe(obj),
		Total:     canon.Vector{},
		Valor:     melhor.valor,
		Gasto:     int32(gasto),
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

// rotuloDe monta o rotulo do resultado.
//
// O texto e deliberadamente longo. "Build otima" seria mais curto e mentiria: o
// calculo ignora passiva e ativa de item, que e o que faz metade dos itens
// valerem o que valem. Um modelo de linguagem lendo "otima" repassaria a
// recomendacao com a autoridade do dataset.
func rotuloDe(obj Objetivo) string {
	return fmt.Sprintf("maximo de %s por ouro, ignorando efeitos de item", obj)
}
