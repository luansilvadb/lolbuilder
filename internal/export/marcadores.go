package export

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
)

// O texto de habilidade do cliente vem com marcadores no lugar dos numeros:
// "causa @TotalDamage@ de Dano Fisico". O cliente os preenche em tempo de
// execucao a partir das mesmas series que o dump publica.
//
// Publicar o marcador cru seria descrever a habilidade sem dizer quanto ela
// faz, que e justamente o que o consumidor precisa. Resolver aqui e o que
// transforma "causa @TotalDamage@" em "causa 30/60/90/120/150 + 1.5 AD".
//
// Marcador que NAO resolve vira "(?)", e nao some nem fica como @Nome@.
//
// Apagar quebra a frase: "Garen fica com @ResistsForTooltip@ de Armadura" vira
// "Garen fica com de Armadura", que le pior que a versao com o marcador. E
// deixar @Nome@ nao informa nada e ainda parece defeito de geracao.
//
// "(?)" mantem a frase de pe e diz a verdade: aqui vai um numero que a fonte
// nao publica. E a mesma regra que rege o resto do conjunto — lacuna declarada,
// nunca zero.

// marcadorSemValor e o que fica no lugar de um numero que a fonte nao publica.
const marcadorSemValor = "(?)"

var marcadorRe = regexp.MustCompile(`@([A-Za-z0-9_.]+)(?:\*([0-9.]+))?@`)

// resolvedor traduz os marcadores de uma habilidade.
type resolvedor struct {
	valores map[string]string
	// NaoResolvidos conta os marcadores que sumiram por nao ter valor.
	NaoResolvidos int
}

// novoResolvedor indexa tudo que a habilidade sabe responder.
//
// A busca e por nome em minusculas porque a fonte referencia com caixa
// diferente da declaracao — escreve @AoeDamagePercent@ onde declarou
// AOEDamagePercent.
func novoResolvedor(h canonical.Habilidade) *resolvedor {
	r := &resolvedor{valores: map[string]string{}}

	// Series nomeadas primeiro; efeito resolvido depois, para que o efeito
	// vença quando os dois existirem — ele e a forma completa, com escala.
	// TodasAsSeries e usada aqui, e nao SeriesNomeadas: o texto referencia
	// series que alguma formula ja consome, e com so as nao consumidas o
	// marcador some e a frase sai quebrada — "Garen fica com de Armadura".
	for _, s := range h.TodasAsSeries {
		r.valores[strings.ToLower(s.Nome)] = serie(s.PorRank)
	}
	for _, e := range h.Efeitos {
		if len(e.PorRank) == 0 {
			continue
		}
		r.valores[strings.ToLower(e.Nome)] = expressoes(e.PorRank)
	}
	if len(h.Recarga) > 0 {
		r.valores["cooldown"] = serie(h.Recarga)
	}
	if len(h.Custo) > 0 {
		r.valores["cost"] = serie(h.Custo)
	}
	return r
}

// Resolver substitui os marcadores do texto.
func (r *resolvedor) Resolver(texto string) string {
	out := marcadorRe.ReplaceAllStringFunc(texto, func(m string) string {
		partes := marcadorRe.FindStringSubmatch(m)
		valor, ok := r.valores[strings.ToLower(partes[1])]
		if !ok {
			r.NaoResolvidos++
			return marcadorSemValor
		}
		if partes[2] == "" {
			return valor
		}
		// A forma @Nome*100@ converte fracao em pontos percentuais. Multiplicar
		// so faz sentido sobre serie de numeros; sobre expressao com escala,
		// o texto sai como esta.
		fator, err := strconv.ParseFloat(partes[2], 64)
		if err != nil {
			return valor
		}
		return multiplicarSerie(valor, fator)
	})
	// A substituicao deixa espaco duplo e espaco antes de pontuacao.
	out = strings.ReplaceAll(out, " ,", ",")
	out = strings.ReplaceAll(out, " .", ".")
	return strings.Join(strings.Fields(out), " ")
}

// multiplicarSerie aplica o fator a cada valor de uma serie "a/b/c".
//
// Devolve o texto original se ele nao for uma serie de numeros puros: uma
// expressao com escala nao pode ser multiplicada sem distribuir o fator pelos
// coeficientes, e fazer isso errado publicaria numero falso.
func multiplicarSerie(s string, fator float64) string {
	partes := strings.Split(s, "/")
	out := make([]string, 0, len(partes))
	for _, p := range partes {
		v, err := strconv.ParseFloat(p, 64)
		if err != nil {
			return s
		}
		out = append(out, num(v*fator))
	}
	return strings.Join(out, "/")
}
