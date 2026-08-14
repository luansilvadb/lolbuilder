package canonical

import (
	"regexp"
	"strings"
)

// A descricao de item e de habilidade vem como HTML de tooltip do cliente. O
// dataset publica texto corrido, entao a marcacao sai — mas o CONTEUDO nunca
// sai junto: uma tag que envolve numero (<magicDamage>60</magicDamage>) tem o
// numero preservado.

var (
	statsBlocoRe = regexp.MustCompile(`(?s)<stats>.*?</stats>`)
	quebraRe     = regexp.MustCompile(`(?i)<br\s*/?>`)
	tagRe        = regexp.MustCompile(`<[^>]*>`)
	espacoRe     = regexp.MustCompile(`\s+`)
)

// TextoDeEfeito extrai o texto do efeito de um item, sem o bloco de atributos.
//
// O bloco de atributos sai porque ele ja foi lido para o vetor de stats;
// repeti-lo no texto duplicaria a informacao no arquivo publicado e daria ao
// consumidor duas fontes para o mesmo numero.
func TextoDeEfeito(descricao string) string {
	return Limpar(statsBlocoRe.ReplaceAllString(descricao, ""))
}

// Limpar converte HTML de tooltip em texto corrido.
func Limpar(s string) string {
	s = quebraRe.ReplaceAllString(s, " ")
	s = tagRe.ReplaceAllString(s, "")
	// Espaco inquebravel vira espaco comum: a fonte usa os dois de forma
	// intercambiavel, e manter o inquebravel faria a mesma frase sair diferente
	// dependendo do item.
	s = strings.ReplaceAll(s, " ", " ")
	s = espacoRe.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}
