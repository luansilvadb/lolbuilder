package canonical

import "testing"

// TestLimparPreservaONumero e a regra que importa: a marcacao sai, o conteudo
// que ela envolve fica. Uma limpeza que comesse o numero junto com a tag
// publicaria habilidade sem dano.
func TestLimparPreservaONumero(t *testing.T) {
	casos := map[string]string{
		"Causa <magicDamage>60</magicDamage> de dano":  "Causa 60 de dano",
		"<passive>Passiva</passive><br>Anda depressa.": "Passiva Anda depressa.",
		"a<br>b<br/>c<BR />d":                          "a b c d",
		"  <b>espaco</b>   sobrando  ":                 "espaco sobrando",
		"<scaleAD>+50% AD</scaleAD> adicional":         "+50% AD adicional",
		"":                                             "",
		"sem marcacao nenhuma":                         "sem marcacao nenhuma",
	}
	for in, want := range casos {
		if got := Limpar(in); got != want {
			t.Errorf("Limpar(%q) = %q, esperado %q", in, got, want)
		}
	}
}

func TestLimparNormalizaEspacoInquebravel(t *testing.T) {
	if got := Limpar("120 de dano"); got != "120 de dano" {
		t.Fatalf("Limpar = %q", got)
	}
}

// TestTextoDeEfeitoTiraOBlocoDeAtributos: o bloco ja virou vetor de stats, e
// repeti-lo aqui daria duas fontes para o mesmo numero.
func TestTextoDeEfeitoTiraOBlocoDeAtributos(t *testing.T) {
	desc := `<mainText><stats><attention> 75</attention> Attack Damage</stats>` +
		`<br><br><passive>Precisao Infinita</passive><br>Criticos causam mais dano.</mainText>`

	got := TextoDeEfeito(desc)
	want := "Precisao Infinita Criticos causam mais dano."
	if got != want {
		t.Fatalf("TextoDeEfeito = %q, esperado %q", got, want)
	}
}

func TestTextoDeEfeitoSemBloco(t *testing.T) {
	if got := TextoDeEfeito("<mainText>So texto.</mainText>"); got != "So texto." {
		t.Fatalf("TextoDeEfeito = %q", got)
	}
	if got := TextoDeEfeito(""); got != "" {
		t.Fatalf("TextoDeEfeito de vazio = %q", got)
	}
}

// TestTextoDeEfeitoSoTiraOBloco: item cujo texto inteiro e o bloco fica com
// efeito vazio, e nao com sobra de marcacao.
func TestTextoDeEfeitoDeItemSoStats(t *testing.T) {
	desc := `<mainText><stats><attention> 40</attention> Ability Power</stats><br><br></mainText>`
	if got := TextoDeEfeito(desc); got != "" {
		t.Fatalf("TextoDeEfeito = %q, esperado vazio", got)
	}
}
