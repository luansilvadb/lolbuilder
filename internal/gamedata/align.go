package gamedata

import "fmt"

// AlignReport e o resultado de cruzar as series de rank do dump com as do
// arquivo do plugin.
//
// As duas fontes descrevem a mesma grandeza por caminhos independentes. Sem
// este cruzamento, um erro de indexacao publicaria TODO valor por rank
// deslocado — em silencio, e sem nenhum sintoma no dado.
type AlignReport struct {
	Series   string // cooldown ou mana
	Expected int    // deslocamento esperado para o tamanho da serie
	Best     int    // deslocamento com mais concordancia
	Hits     map[int]int
	Compared int
	// Divergent lista habilidades em que as duas fontes discordam mesmo sob o
	// deslocamento vencedor.
	Divergent []string
	// Omitidas conta quantas divergentes ficaram fora da lista por causa do
	// teto. Uma lista cortada em silencio e pior no momento em que mais
	// importa: num evento real de quebra ela estaria cheia, e quem lesse
	// concluiria que so aquelas divergiram.
	Omitidas int
	// Minimo e a concordancia minima exigida, em pontos percentuais. Zero
	// desliga a verificacao.
	Minimo int
}

// Agreement e a taxa de concordancia sob o deslocamento vencedor.
func (r AlignReport) Agreement() float64 {
	if r.Compared == 0 {
		return 0
	}
	return float64(r.Hits[r.Best]) * 100 / float64(r.Compared)
}

// OK informa se a convencao de indexacao continua sendo a esperada e se a
// concordancia esta acima do minimo.
func (r AlignReport) OK() bool {
	if r.Compared == 0 {
		return true
	}
	if r.Best != r.Expected {
		return false
	}
	return r.Minimo == 0 || r.Agreement() >= float64(r.Minimo)
}

func (r AlignReport) Err() error {
	if r.OK() {
		return nil
	}
	if r.Best != r.Expected {
		return fmt.Errorf(
			"alinhamento de rank da serie %s mudou: o deslocamento de melhor concordancia "+
				"e %d (%d de %d), mas o esperado para series de %d posicoes e %d — "+
				"publicar agora deslocaria todo o valor por rank",
			r.Series, r.Best, r.Hits[r.Best], r.Compared, seriesLen(r.Expected), r.Expected)
	}
	return fmt.Errorf(
		"concordancia da serie %s caiu para %.1f%% (%d de %d), minimo %d%% — "+
			"as duas fontes descrevem a mesma grandeza e deveriam concordar. "+
			"Ou a fonte mudou de forma, ou o projeto voltou a ler a serie herdada "+
			"em vez da redefinida. Divergentes: %v",
		r.Series, r.Agreement(), r.Hits[r.Best], r.Compared, r.Minimo, r.Divergent)
}

func seriesLen(offset int) int {
	if offset == 1 {
		return RankedLen
	}
	return PluginAlignedLen
}

// tetoDaLista limita quantas divergentes sao nomeadas. O excedente e contado em
// Omitidas, e nunca descartado em silencio.
const tetoDaLista = 40

// AlignChecker acumula a comparacao ao longo de todos os campeoes.
type AlignChecker struct {
	reports map[string]*AlignReport
	minimo  int
}

// NewAlignChecker monta o verificador com a concordancia minima exigida.
// Zero desliga a verificacao de minimo, mantendo apenas a de deslocamento.
func NewAlignChecker(minimo int) *AlignChecker {
	return &AlignChecker{reports: map[string]*AlignReport{}, minimo: minimo}
}

// Compare cruza uma serie do dump com a correspondente do plugin.
//
// A comparacao e por deslocamento, e nao por igualdade de cada valor. As duas
// fontes podem discordar pontualmente em algumas habilidades sem que a
// indexacao tenha mudado; o que interessa detectar e a troca de convencao, que
// derruba a concordancia global de uma vez.
func (c *AlignChecker) Compare(series, who string, dump, plugin []float64, ranks int) {
	if len(dump) == 0 || len(plugin) == 0 {
		return
	}
	rep, ok := c.reports[series]
	if !ok {
		rep = &AlignReport{
			Series:   series,
			Expected: RankOffset(len(dump)),
			Hits:     map[int]int{},
		}
		c.reports[series] = rep
	}

	// A pergunta e "esta habilidade bateu?", que so o contador antes e depois
	// responde. Comparar a lacuna acumulada contra Compared-1 responde outra
	// coisa: depois da primeira falha a lacuna fica em 1 e toda habilidade que
	// bate entra na lista, enquanto a segunda falha, que abre a lacuna para 2,
	// fica de fora.
	hitsAntes := rep.Hits[rep.Expected]

	for _, off := range []int{0, 1} {
		match := true
		for rank := 1; rank <= ranks; rank++ {
			i, j := rank-1+off, rank-1
			if i >= len(dump) || j >= len(plugin) {
				match = false
				break
			}
			if d := dump[i] - plugin[j]; d > 0.01 || d < -0.01 {
				match = false
				break
			}
		}
		if match {
			rep.Hits[off]++
		}
	}
	rep.Compared++
	if rep.Hits[rep.Expected] == hitsAntes {
		// O contador nao andou: nao bateu sob o deslocamento esperado.
		if len(rep.Divergent) < tetoDaLista {
			rep.Divergent = append(rep.Divergent, who)
		} else {
			rep.Omitidas++
		}
	}
}

// Reports fecha a apuracao, escolhendo o deslocamento vencedor de cada serie.
func (c *AlignChecker) Reports() []AlignReport {
	out := make([]AlignReport, 0, len(c.reports))
	for _, r := range c.reports {
		best, bestHits := r.Expected, -1
		for _, off := range []int{0, 1} {
			if r.Hits[off] > bestHits {
				best, bestHits = off, r.Hits[off]
			}
		}
		r.Best = best
		r.Minimo = c.minimo
		out = append(out, *r)
	}
	return out
}
