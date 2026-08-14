package canon

import (
	"regexp"
	"strconv"
	"strings"
)

// O bloco de atributos do item vem dentro do HTML da descricao, e nao num campo
// estruturado — o CommunityDragon nao publica stat de item de outra forma.
//
// A boa noticia, medida no 16.16: a gramatica e uniforme. As 475 linhas de stat
// dos 210 itens compraveis casam com a MESMA forma, sem excecao:
//
//	<stats><attention> 75</attention> Attack Damage<br>
//	       <attention> 25%</attention> Critical Strike Chance</stats>
//
// E, diferente do modo Jade, o bloco ja inclui o deslocamento das botas — a
// limitacao mais citada do dataset original nao se reproduz aqui.

var (
	blocoRe = regexp.MustCompile(`(?s)<stats>(.*?)</stats>`)
	linhaRe = regexp.MustCompile(`(?s)^<attention>\s*([0-9]+(?:\.[0-9]+)?)\s*(%?)\s*</attention>\s*(.+)$`)
	tagRe   = regexp.MustCompile(`<[^>]*>`)
)

// LinhaNaoLida e uma linha do bloco que o parser nao soube ler.
//
// Existe para virar numero de cobertura e sair no relatorio do build. Linha nao
// lida nunca vira zero: o stat correspondente simplesmente nao e publicado, e o
// item aparece com um atributo a menos em vez de com um atributo errado.
type LinhaNaoLida struct {
	Item   string
	Linha  string
	Motivo string
}

// LeituraDeItem e o resultado de ler o bloco de atributos de um item.
type LeituraDeItem struct {
	Stats Vector

	// TemBloco distingue "item sem bloco de atributos" de "bloco vazio".
	//
	// No 16.16 ha exatamente um item compravel sem bloco: World Atlas, cuja
	// descricao vem VAZIA da fonte. Isso e lacuna da fonte, nao falha de
	// leitura, e por isso nao entra no denominador da cobertura.
	TemBloco bool

	// Linhas e quantas linhas o bloco tinha; Lidas, quantas viraram stat.
	Linhas int
	Lidas  int

	NaoLidas []LinhaNaoLida
}

// Completa informa se toda linha do bloco virou stat.
func (l LeituraDeItem) Completa() bool {
	return l.TemBloco && l.Lidas == l.Linhas
}

// LerStatsDeItem extrai o vetor de stats da descricao de um item.
//
// Nunca devolve erro: uma linha ilegivel e dado ausente, e dado ausente e
// reportado, nao fatal. O que vigia esta fonte e a taxa de cobertura que o
// build imprime, com minimo em coverage_minimums.
func LerStatsDeItem(nomeItem, descricao string) LeituraDeItem {
	out := LeituraDeItem{Stats: Vector{}}

	m := blocoRe.FindStringSubmatch(descricao)
	if m == nil {
		return out
	}
	out.TemBloco = true

	for _, bruta := range strings.Split(m[1], "<br>") {
		linha := strings.TrimSpace(bruta)
		if linha == "" {
			continue
		}
		out.Linhas++

		partes := linhaRe.FindStringSubmatch(linha)
		if partes == nil {
			out.NaoLidas = append(out.NaoLidas, LinhaNaoLida{
				Item: nomeItem, Linha: linha,
				Motivo: "fora da forma <attention>VALOR</attention> ROTULO",
			})
			continue
		}

		valor, err := strconv.ParseFloat(partes[1], 64)
		if err != nil {
			out.NaoLidas = append(out.NaoLidas, LinhaNaoLida{
				Item: nomeItem, Linha: linha, Motivo: "valor nao numerico",
			})
			continue
		}

		percentual := partes[2] == "%"
		texto := limparRotulo(partes[3])

		stat, ok := LookupStat(texto, percentual)
		if !ok {
			unidade := "plano"
			if percentual {
				unidade = "percentual"
			}
			out.NaoLidas = append(out.NaoLidas, LinhaNaoLida{
				Item: nomeItem, Linha: linha,
				Motivo: "rotulo " + strconv.Quote(texto) + " (" + unidade +
					") nao esta no vocabulario canonico",
			})
			continue
		}

		out.Stats.Add(stat, valor)
		out.Lidas++
	}
	return out
}

// limparRotulo tira marcacao residual e espaco do texto do rotulo.
func limparRotulo(s string) string {
	return strings.TrimSpace(tagRe.ReplaceAllString(s, ""))
}

// CoberturaDeItens resume a leitura do catalogo inteiro.
type CoberturaDeItens struct {
	// ComBloco e o denominador: itens que a fonte descreve com bloco de
	// atributos. Item sem bloco fica fora, porque nao ha o que ler.
	ComBloco int
	// SemBloco conta os itens que a fonte nao descreve. Aparecem no relatorio
	// por nome, para que "a fonte nao publica" nunca se confunda com "o parser
	// nao leu".
	SemBloco     int
	SemBlocoNome []string

	LidosPorInteiro int
	Linhas          int
	LinhasLidas     int
	NaoLidas        []LinhaNaoLida
}

// Taxa devolve a fracao de itens com bloco lidos por inteiro, em pontos
// percentuais.
func (c CoberturaDeItens) Taxa() float64 {
	if c.ComBloco == 0 {
		return 0
	}
	return 100 * float64(c.LidosPorInteiro) / float64(c.ComBloco)
}

// Acumular incorpora a leitura de um item a cobertura.
func (c *CoberturaDeItens) Acumular(nomeItem string, l LeituraDeItem) {
	if !l.TemBloco {
		c.SemBloco++
		c.SemBlocoNome = append(c.SemBlocoNome, nomeItem)
		return
	}
	c.ComBloco++
	c.Linhas += l.Linhas
	c.LinhasLidas += l.Lidas
	c.NaoLidas = append(c.NaoLidas, l.NaoLidas...)
	if l.Completa() {
		c.LidosPorInteiro++
	}
}
