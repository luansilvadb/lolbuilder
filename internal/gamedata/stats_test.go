package gamedata

import "testing"

// TestCrescimentoNaoELinear guarda o achado mais importante do projeto, com os
// numeros MEDIDOS numa partida real de Rammus na Ferramenta de Treino.
//
// O crescimento do LoL aplica um fator que sai de 0.7025 no nivel 2 e chega a 1
// no 18. Nenhuma fonte publica esse fator; so o oraculo em partida o revela — e
// o modelo linear que o projeto herdou do dataset do modo Jade previa
// 10.25 onde o jogo somou 8.0975.
func TestCrescimentoNaoELinear(t *testing.T) {
	// Rammus no patch 16.16.
	mr := Scaling{Base: 32, PerLevel: 2.05}
	armadura := Scaling{Base: 35, PerLevel: 4.5}

	casos := []struct {
		nome    string
		s       Scaling
		de, ate int
		medido  float64
	}{
		// Medido: do nivel 1 ao 6 o jogo somou 8.0975 de resistencia magica.
		// O modelo linear previa 2.05 * 5 = 10.25.
		{"resistencia magica 1->6", mr, 1, 6, 8.0975},
		// Medido na mesma leitura: a armadura subiu 87.775, dos quais exatos 70
		// vieram de Cota de Malha (45) e Colete Espinhoso (25).
		{"armadura 1->6", armadura, 1, 6, 17.775},
		// Segundo intervalo, leitura independente.
		{"armadura 6->7", armadura, 6, 7, 4.0275},
		{"resistencia magica 6->7", mr, 6, 7, 1.83475},
	}
	for _, c := range casos {
		got := c.s.CrescimentoEntre(c.de, c.ate)
		if diff := got - c.medido; diff > 0.0001 || diff < -0.0001 {
			t.Errorf("%s = %.4f, medido no jogo %.4f", c.nome, got, c.medido)
		}
	}
}

// TestNoNivel18OFatorVale1 explica por que o erro ficou invisivel: o total no
// nivel maximo coincide com o crescimento linear, e o dataset publica justamente
// a coluna "no 18".
func TestNoNivel18OFatorVale1(t *testing.T) {
	s := Scaling{Base: 100, PerLevel: 10}
	if got := s.At(18); got != 100+10*17 {
		t.Fatalf("At(18) = %v, esperado coincidir com o linear (%v)", got, 100+10*17)
	}
	if got := fatorDeCrescimento(18); got != 17 {
		t.Fatalf("fatorDeCrescimento(18) = %v, esperado 17", got)
	}
	// Nos niveis intermediarios os dois modelos discordam, e e ai que o erro
	// vivia.
	if got := s.At(6); got == 100+10*5 {
		t.Fatal("At(6) ainda coincide com o linear")
	}
}

func TestAtNoNivel1EABase(t *testing.T) {
	s := Scaling{Base: 690, PerLevel: 98}
	if got := s.At(1); got != 690 {
		t.Fatalf("At(1) = %v", got)
	}
	if got := s.At(0); got != 690 {
		t.Fatalf("At(0) = %v; nivel abaixo de 1 deveria devolver a base", got)
	}
	if got := s.CrescimentoEntre(3, 3); got != 0 {
		t.Fatalf("crescimento de um nivel para ele mesmo = %v", got)
	}
}

// TestIntervalosCegos: nem todo par de niveis testa a curvatura. Do 7 ao 12 o
// fator acumula exatamente 5, o mesmo que 5 niveis lineares, e do 1 ao 18
// acumula 17 — uma amostragem nesses pares passaria com nota maxima mesmo com a
// formula errada.
func TestIntervalosCegos(t *testing.T) {
	cegos := [][2]int{{7, 12}, {1, 18}}
	for _, c := range cegos {
		if IntervaloTestaACurvatura(c[0], c[1]) {
			t.Errorf("o intervalo %d->%d foi dado como diagnostico, mas e cego", c[0], c[1])
		}
	}
	// Estes sao os que revelaram o defeito numa partida real.
	uteis := [][2]int{{1, 6}, {6, 7}, {1, 12}, {6, 18}}
	for _, c := range uteis {
		if !IntervaloTestaACurvatura(c[0], c[1]) {
			t.Errorf("o intervalo %d->%d foi dado como cego, mas distingue os modelos", c[0], c[1])
		}
	}
}
