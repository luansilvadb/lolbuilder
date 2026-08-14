package gamedata

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// StatNames traduz os enums mStat e mStatFormula do dump para os nomes que o
// resto do projeto usa.
//
// A Riot nao publica esse mapeamento. Ele foi determinado cruzando a formula
// resolvida com o nome dos coeficientes e com a descricao que o plugin publica,
// e vive em curation/statenum.json com a evidencia de cada entrada.
//
// Enum ausente da tabela NAO vira zero nem palpite: a parcela fica sem resolver
// e entra na cobertura. Uma habilidade fora da tabela apenas nao informa;
// publicada com o stat errado, ela mente.
type StatNames struct {
	stat    map[int]string
	formula map[int]string
}

// statEnumFile e a forma de curation/statenum.json.
type statEnumFile struct {
	Stat map[string]struct {
		Name     string `json:"name"`
		Evidence string `json:"evidence"`
	} `json:"stat"`
	Formula map[string]string `json:"formula"`
}

// LoadStatNames le a tabela curada.
func LoadStatNames(path string) (StatNames, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return StatNames{}, fmt.Errorf("lendo %s: %w", path, err)
	}
	var f statEnumFile
	if err := json.Unmarshal(raw, &f); err != nil {
		return StatNames{}, fmt.Errorf("parseando %s: %w", path, err)
	}

	out := StatNames{stat: map[int]string{}, formula: map[int]string{}}
	for k, v := range f.Stat {
		n, err := strconv.Atoi(k)
		if err != nil {
			return StatNames{}, fmt.Errorf("%s: chave de stat nao numerica %q", path, k)
		}
		if v.Name == "" {
			return StatNames{}, fmt.Errorf("%s: stat %d sem nome", path, n)
		}
		if v.Evidence == "" {
			return StatNames{}, fmt.Errorf(
				"%s: stat %d sem evidencia — a tabela e curada, e entrada sem "+
					"justificativa e palpite que o proximo leitor nao tem como refazer", path, n)
		}
		out.stat[n] = v.Name
	}
	for k, v := range f.Formula {
		n, err := strconv.Atoi(k)
		if err != nil {
			return StatNames{}, fmt.Errorf("%s: chave de formula nao numerica %q", path, k)
		}
		out.formula[n] = v
	}
	if len(out.stat) == 0 {
		return StatNames{}, fmt.Errorf("%s: tabela de stats vazia", path)
	}
	return out, nil
}

// Name traduz o enum de stat. O ponteiro nulo significa campo omitido, que a
// fonte usa para o valor zero.
func (s StatNames) Name(stat *int) (string, error) {
	v := 0
	if stat != nil {
		v = *stat
	}
	name, ok := s.stat[v]
	if !ok {
		return "", fmt.Errorf("mStat %d fora da tabela curada (veja curation/statenum.json): "+
			"determine o significado cruzando a formula com a descricao do plugin antes de publicar", v)
	}
	return name, nil
}

// Known lista os valores mapeados, em ordem, para as mensagens de diagnostico.
func (s StatNames) Known() []int {
	out := make([]int, 0, len(s.stat))
	for k := range s.stat {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}

// formulaNames traduz mStatFormula. Fica em variavel de pacote e nao na tabela
// curada porque sao so tres valores e eles vem da estrutura do jogo, nao de
// julgamento nosso.
var formulaNames = map[int]string{0: "total", 1: "base", 2: "bonus"}

// formulaName traduz o enum de formula. Ausente e total.
func formulaName(f *int) string {
	if f == nil {
		return "total"
	}
	if n, ok := formulaNames[*f]; ok {
		return n
	}
	return fmt.Sprintf("formula_%d", *f)
}
