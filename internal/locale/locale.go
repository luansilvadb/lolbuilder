// Package locale confere que o locale de exibicao descreve as mesmas entidades
// que o canonico.
//
// Os dois arquivos sao o mesmo conteudo em idiomas diferentes, entao a
// conferencia certa nao e uma contagem minima solta: e a igualdade dos
// conjuntos de id. Um minimo proprio para o locale de exibicao toleraria a
// fonte servir metade das entidades, que e justamente o que precisa ser pego.
//
// Sem isso, um pt_br truncado ou vazio entra no snapshot em silencio e so
// aparece no fim da linha, como nome em branco no arquivo publicado.
package locale

import (
	"fmt"
	"sort"
	"strings"
)

// Names indexa o nome de cada entidade pelo id.
type Names map[int64]string

// maxListados limita quantos ids uma mensagem de erro enumera. Cinco bastam
// para identificar o padrao; a contagem total diz o tamanho do estrago.
const maxListados = 5

// Check compara os dois locales de uma entidade.
//
// kind e o nome da entidade na mensagem de erro, e file o arquivo do locale de
// exibicao — quem le o erro precisa saber qual dos dois lados corrigir.
func Check(kind, file string, canon, disp Names) error {
	if len(disp) == 0 {
		return fmt.Errorf("%s: %s nao tem entidade alguma, contra %d no locale canonico",
			kind, file, len(canon))
	}

	var faltando, sobrando, semNome []int64
	for id := range canon {
		if _, ok := disp[id]; !ok {
			faltando = append(faltando, id)
		}
	}
	for id, nome := range disp {
		nomeCanon, ok := canon[id]
		if !ok {
			sobrando = append(sobrando, id)
			continue
		}
		// So e lacuna de traducao se o canonico nomeia e o de exibicao nao.
		// Vazio nos dois lados e fato da fonte: ha entidades que a Riot nunca
		// nomeia — no 16.16 sao o item 2008, cuja chave de localizacao nao
		// resolve, e o feitico 5, que nao vale em modo nenhum. Tratar isso como
		// defeito abortaria toda captura, e a regra do projeto para entidade
		// sem nome na fonte e publica-la sem nome, nunca inventar um.
		if strings.TrimSpace(nome) == "" && strings.TrimSpace(nomeCanon) != "" {
			semNome = append(semNome, id)
		}
	}

	if len(faltando) > 0 {
		return fmt.Errorf("%s: %d de %d entidade(s) do locale canonico nao existem em %s: %s",
			kind, len(faltando), len(canon), file, listar(faltando))
	}
	if len(sobrando) > 0 {
		return fmt.Errorf("%s: %s tem %d entidade(s) que o locale canonico nao tem: %s",
			kind, file, len(sobrando), listar(sobrando))
	}
	if len(semNome) > 0 {
		return fmt.Errorf("%s: %d entidade(s) sem nome em %s: %s — "+
			"traducao ausente publica campo em branco no dataset",
			kind, len(semNome), file, listar(semNome))
	}
	return nil
}

func listar(ids []int64) string {
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	parts := make([]string, 0, maxListados+1)
	for i, id := range ids {
		if i == maxListados {
			parts = append(parts, fmt.Sprintf("... e mais %d", len(ids)-maxListados))
			break
		}
		parts = append(parts, fmt.Sprint(id))
	}
	return strings.Join(parts, ", ")
}
