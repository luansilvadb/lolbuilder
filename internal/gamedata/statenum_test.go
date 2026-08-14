package gamedata

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func escreverCuradoria(t *testing.T, conteudo string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "statenum.json")
	if err := os.WriteFile(p, []byte(conteudo), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

const curadoriaValida = `{
  "curated_from_patch": "16.16",
  "stat": {
    "0": {"name": "ability_power", "evidence": "APRatio em 166 coeficientes"},
    "2": {"name": "attack_damage", "evidence": "ADRatio em 55 coeficientes"}
  },
  "formula": {"0": "total", "1": "base", "2": "bonus"}
}`

func TestLoadStatNames(t *testing.T) {
	s, err := LoadStatNames(escreverCuradoria(t, curadoriaValida))
	if err != nil {
		t.Fatal(err)
	}
	dois := 2
	if n, err := s.Name(&dois); err != nil || n != "attack_damage" {
		t.Fatalf("mStat 2 = %q (%v)", n, err)
	}
	// Ponteiro nulo e o campo omitido, que a fonte usa para o valor zero.
	if n, err := s.Name(nil); err != nil || n != "ability_power" {
		t.Fatalf("mStat omitido = %q (%v)", n, err)
	}
	if got := s.Known(); len(got) != 2 || got[0] != 0 || got[1] != 2 {
		t.Fatalf("Known() = %v", got)
	}
}

// TestEnumForaDaTabelaNaoResolve: enum sem curadoria deixa a parcela em aberto e
// entra na cobertura. Publicar com o stat errado mentiria, e o consumidor nao
// teria como perceber.
func TestEnumForaDaTabelaNaoResolve(t *testing.T) {
	s, err := LoadStatNames(escreverCuradoria(t, curadoriaValida))
	if err != nil {
		t.Fatal(err)
	}
	noventa := 90
	_, err = s.Name(&noventa)
	if err == nil {
		t.Fatal("enum fora da tabela resolveu")
	}
	if !strings.Contains(err.Error(), "statenum.json") {
		t.Fatalf("o erro nao diz onde curar: %v", err)
	}
}

// TestEntradaSemEvidenciaERejeitada: a tabela e curada, e entrada sem
// justificativa e palpite que o proximo leitor nao tem como refazer.
func TestEntradaSemEvidenciaERejeitada(t *testing.T) {
	semEvidencia := strings.Replace(curadoriaValida,
		`{"name": "attack_damage", "evidence": "ADRatio em 55 coeficientes"}`,
		`{"name": "attack_damage"}`, 1)
	_, err := LoadStatNames(escreverCuradoria(t, semEvidencia))
	if err == nil {
		t.Fatal("entrada sem evidencia foi aceita")
	}
	if !strings.Contains(err.Error(), "evidencia") {
		t.Fatalf("o erro nao explica o motivo: %v", err)
	}
}

func TestCuradoriaInvalida(t *testing.T) {
	casos := map[string]string{
		"json quebrado":      `nao e json`,
		"tabela vazia":       `{"stat": {}, "formula": {}}`,
		"chave nao numerica": strings.Replace(curadoriaValida, `"2": {"name": "attack_damage"`, `"xis": {"name": "attack_damage"`, 1),
		"nome vazio":         strings.Replace(curadoriaValida, `"name": "attack_damage"`, `"name": ""`, 1),
	}
	for nome, conteudo := range casos {
		t.Run(nome, func(t *testing.T) {
			if _, err := LoadStatNames(escreverCuradoria(t, conteudo)); err == nil {
				t.Fatal("curadoria invalida foi aceita")
			}
		})
	}
}

func TestArquivoAusente(t *testing.T) {
	if _, err := LoadStatNames(filepath.Join(t.TempDir(), "naoexiste.json")); err == nil {
		t.Fatal("arquivo ausente foi aceito")
	}
}

func TestFormulaName(t *testing.T) {
	zero, base, bonus := 0, 1, 2
	casos := map[*int]string{nil: "total", &zero: "total", &base: "base", &bonus: "bonus"}
	for in, want := range casos {
		if got := formulaName(in); got != want {
			t.Errorf("formulaName(%v) = %q, esperado %q", in, got, want)
		}
	}
	// Valor desconhecido sai identificado, e nao colapsado em "total": colapsar
	// afirmaria que o coeficiente multiplica o stat inteiro quando ninguem sabe.
	novo := 7
	if got := formulaName(&novo); got != "formula_7" {
		t.Errorf("formulaName(7) = %q", got)
	}
}
