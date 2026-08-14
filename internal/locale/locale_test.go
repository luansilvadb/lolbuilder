package locale

import (
	"strings"
	"testing"
)

var canon = Names{1001: "Boots", 3031: "Infinity Edge", 3006: "Berserker's Greaves"}

func TestCheckAceitaLocalesEquivalentes(t *testing.T) {
	disp := Names{1001: "Botas", 3031: "Gume do Infinito", 3006: "Grevas do Frenesi"}
	if err := Check("itens", "pt_br/items.json", canon, disp); err != nil {
		t.Fatalf("locales equivalentes foram rejeitados: %v", err)
	}
}

// TestCheckAceitaNomeIdentico: varios campeoes tem o mesmo nome nos dois
// idiomas, e isso nao e sinal de traducao ausente.
func TestCheckAceitaNomeIdentico(t *testing.T) {
	disp := Names{1001: "Boots", 3031: "Infinity Edge", 3006: "Berserker's Greaves"}
	if err := Check("itens", "pt_br/items.json", canon, disp); err != nil {
		t.Fatalf("nome igual nos dois locales foi tratado como defeito: %v", err)
	}
}

// TestCheckPegaLocaleVazio e o caso que motivou o pacote: sem esta guarda, um
// pt_br vazio entra no snapshot e so aparece como nome em branco no arquivo
// publicado, marcos depois.
func TestCheckPegaLocaleVazio(t *testing.T) {
	err := Check("itens", "pt_br/items.json", canon, Names{})
	if err == nil {
		t.Fatal("locale de exibicao vazio passou")
	}
	if !strings.Contains(err.Error(), "pt_br/items.json") {
		t.Fatalf("o erro nao diz qual arquivo esta errado: %v", err)
	}
}

// TestCheckPegaLocaleTruncado: um minimo solto toleraria metade das entidades.
// A igualdade de conjuntos nao tolera.
func TestCheckPegaLocaleTruncado(t *testing.T) {
	disp := Names{1001: "Botas"}
	err := Check("itens", "pt_br/items.json", canon, disp)
	if err == nil {
		t.Fatal("locale truncado passou")
	}
	if !strings.Contains(err.Error(), "2 de 3") {
		t.Fatalf("o erro nao diz o tamanho do estrago: %v", err)
	}
}

func TestCheckPegaEntidadeExtra(t *testing.T) {
	disp := Names{1001: "Botas", 3031: "Gume do Infinito", 3006: "Grevas", 9999: "Fantasma"}
	err := Check("itens", "pt_br/items.json", canon, disp)
	if err == nil {
		t.Fatal("entidade que so existe no locale de exibicao passou")
	}
	if !strings.Contains(err.Error(), "9999") {
		t.Fatalf("o erro nao nomeia a entidade extra: %v", err)
	}
}

func TestCheckPegaNomeVazio(t *testing.T) {
	disp := Names{1001: "Botas", 3031: "   ", 3006: ""}
	err := Check("itens", "pt_br/items.json", canon, disp)
	if err == nil {
		t.Fatal("nome vazio passou")
	}
	if !strings.Contains(err.Error(), "3006") || !strings.Contains(err.Error(), "3031") {
		t.Fatalf("o erro nao nomeia as entidades sem nome: %v", err)
	}
}

// TestCheckAceitaVazioNosDoisLados fixa a distincao que so apareceu ao rodar a
// guarda contra o dado real: no 16.16 o item 2008 e o feitico 5 nao tem nome em
// locale nenhum. Isso e fato da fonte, e a regra do projeto para entidade sem
// nome e publica-la sem nome — nunca inventar um, e nunca abortar por causa
// disso.
func TestCheckAceitaVazioNosDoisLados(t *testing.T) {
	c := Names{1001: "Boots", 2008: ""}
	d := Names{1001: "Botas", 2008: ""}
	if err := Check("itens", "pt_br/items.json", c, d); err != nil {
		t.Fatalf("entidade sem nome nos dois locales foi tratada como defeito: %v", err)
	}
}

func TestListarTruncaEDizQuantoSobrou(t *testing.T) {
	got := listar([]int64{9, 8, 7, 6, 5, 4, 3})
	if !strings.HasPrefix(got, "3, 4, 5, 6, 7") {
		t.Fatalf("listar nao ordenou: %q", got)
	}
	if !strings.Contains(got, "e mais 2") {
		t.Fatalf("listar nao disse quantos sobraram: %q", got)
	}
}
