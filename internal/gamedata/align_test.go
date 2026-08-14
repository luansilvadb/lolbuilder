package gamedata

import (
	"strings"
	"testing"
)

// As series de 7 posicoes reservam o indice 0 para "habilidade nao aprendida",
// entao o rank 1 esta no indice 1. As de 6 comecam no indice 0.
func serieDe7(v ...float64) []float64 { return append([]float64{0}, v...) }

func TestAlinhamentoConcordanteEstaOK(t *testing.T) {
	c := NewAlignChecker(98)
	for i := 0; i < 10; i++ {
		c.Compare("cooldown", "champ", serieDe7(8, 8, 8, 8, 8, 8), []float64{8, 8, 8, 8, 8}, 5)
	}
	rel := c.Reports()
	if len(rel) != 1 {
		t.Fatalf("relatorios = %d", len(rel))
	}
	r := rel[0]
	if !r.OK() {
		t.Fatalf("series identicas nao passaram: %v", r.Err())
	}
	if r.Agreement() != 100 || r.Best != 1 {
		t.Fatalf("concordancia %.1f%%, deslocamento %d", r.Agreement(), r.Best)
	}
}

// TestDeslocamentoTrocadoEDetectado e a razao de existir do pacote: sem ele, um
// erro de indexacao publicaria TODO valor por rank deslocado, em silencio e sem
// sintoma no dado.
func TestDeslocamentoTrocadoEDetectado(t *testing.T) {
	c := NewAlignChecker(98)
	// O dump tem 7 posicoes, entao o esperado e deslocamento 1. Aqui os valores
	// batem com deslocamento 0, o que significa que a convencao mudou.
	for i := 0; i < 10; i++ {
		c.Compare("cooldown", "champ", []float64{8, 7, 6, 5, 4, 4, 4}, []float64{8, 7, 6, 5, 4}, 5)
	}
	r := c.Reports()[0]
	if r.OK() {
		t.Fatal("troca de convencao de indexacao passou")
	}
	if r.Best != 0 || r.Expected != 1 {
		t.Fatalf("melhor deslocamento %d, esperado %d", r.Best, r.Expected)
	}
	if !strings.Contains(r.Err().Error(), "deslocaria todo o valor por rank") {
		t.Fatalf("o erro nao diz o tamanho do estrago: %v", r.Err())
	}
}

func TestConcordanciaAbaixoDoMinimoFalha(t *testing.T) {
	c := NewAlignChecker(98)
	for i := 0; i < 9; i++ {
		c.Compare("mana", "bom", []float64{50, 60, 70, 80, 90, 90}, []float64{50, 60, 70, 80, 90}, 5)
	}
	c.Compare("mana", "ruim", []float64{1, 2, 3, 4, 5, 5}, []float64{50, 60, 70, 80, 90}, 5)

	r := c.Reports()[0]
	if r.OK() {
		t.Fatalf("90%% de concordancia passou num minimo de 98%%: %+v", r)
	}
	if len(r.Divergent) != 1 || r.Divergent[0] != "ruim" {
		t.Fatalf("a divergente nao foi nomeada: %v", r.Divergent)
	}
}

// TestMinimoZeroDesligaSoOMinimo: sem minimo declarado a verificacao de
// deslocamento continua valendo, porque ela nao depende de calibracao.
func TestMinimoZeroDesligaSoOMinimo(t *testing.T) {
	c := NewAlignChecker(0)
	for i := 0; i < 9; i++ {
		c.Compare("mana", "bom", []float64{50, 60, 70, 80, 90, 90}, []float64{50, 60, 70, 80, 90}, 5)
	}
	c.Compare("mana", "ruim", []float64{1, 2, 3, 4, 5, 5}, []float64{50, 60, 70, 80, 90}, 5)
	if !c.Reports()[0].OK() {
		t.Fatal("minimo zero deveria tolerar a divergencia pontual")
	}
}

// TestDivergentesSaoContadasUmaAUma protege o defeito que a lacuna acumulada
// causava: depois da primeira falha, toda habilidade que batia entrava na lista.
func TestDivergentesSaoContadasUmaAUma(t *testing.T) {
	c := NewAlignChecker(98)
	c.Compare("mana", "ruim1", []float64{1, 1, 1, 1, 1, 1}, []float64{9, 9, 9, 9, 9}, 5)
	c.Compare("mana", "bom", []float64{9, 9, 9, 9, 9, 9}, []float64{9, 9, 9, 9, 9}, 5)
	c.Compare("mana", "ruim2", []float64{2, 2, 2, 2, 2, 2}, []float64{9, 9, 9, 9, 9}, 5)

	r := c.Reports()[0]
	if len(r.Divergent) != 2 {
		t.Fatalf("divergentes = %v, esperado ruim1 e ruim2", r.Divergent)
	}
	if r.Divergent[0] != "ruim1" || r.Divergent[1] != "ruim2" {
		t.Fatalf("a habilidade que bateu entrou na lista: %v", r.Divergent)
	}
}

// TestListaCortadaContaOExcedente: uma lista cortada em silencio e pior no
// momento em que mais importa — num evento real de quebra ela estaria cheia, e
// quem lesse concluiria que so aquelas divergiram.
func TestListaCortadaContaOExcedente(t *testing.T) {
	c := NewAlignChecker(98)
	for i := 0; i < tetoDaLista+7; i++ {
		c.Compare("mana", "ruim", []float64{1, 1, 1, 1, 1, 1}, []float64{9, 9, 9, 9, 9}, 5)
	}
	r := c.Reports()[0]
	if len(r.Divergent) != tetoDaLista {
		t.Fatalf("lista com %d entradas, teto e %d", len(r.Divergent), tetoDaLista)
	}
	if r.Omitidas != 7 {
		t.Fatalf("omitidas = %d, esperado 7", r.Omitidas)
	}
}

func TestSerieVaziaNaoEntraNaConta(t *testing.T) {
	c := NewAlignChecker(98)
	c.Compare("mana", "sem dump", nil, []float64{1, 2, 3}, 3)
	c.Compare("mana", "sem plugin", []float64{1, 2, 3}, nil, 3)
	if len(c.Reports()) != 0 {
		t.Fatal("serie ausente de um dos lados entrou na comparacao")
	}
}

func TestRelatorioVazioEstaOK(t *testing.T) {
	r := AlignReport{Series: "mana", Hits: map[int]int{}}
	if !r.OK() || r.Err() != nil {
		t.Fatal("relatorio sem comparacao alguma deveria passar")
	}
	if r.Agreement() != 0 {
		t.Fatalf("Agreement() = %v sem comparacoes", r.Agreement())
	}
}
