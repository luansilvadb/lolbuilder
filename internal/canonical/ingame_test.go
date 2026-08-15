package canonical

import (
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/gamedata"
)

// datasetIngame monta um campeao com crescimento conhecido, para que a
// comparacao possa ser conferida de cabeca.
func datasetIngame() *Dataset {
	return &Dataset{
		Patch: "16.99",
		Champions: []Champion{{
			ID: 86, Nome: "Garen", NomeCanonico: "Garen", Alias: "Garen",
			Stats: &gamedata.Stats{
				HP:           gamedata.Scaling{Base: 690, PerLevel: 100},
				Armor:        gamedata.Scaling{Base: 38, PerLevel: 4},
				MagicResist:  gamedata.Scaling{Base: 32, PerLevel: 2},
				AttackDamage: gamedata.Scaling{Base: 69, PerLevel: 5},
			},
		}},
	}
}

func amostra(nivel int, vida, armadura, mr, ad float64, itens ...string) AmostraIngame {
	return AmostraIngame{
		Nivel: nivel, Campeao: "Garen", Itens: itens,
		MaxHealth: vida, Armor: armadura, MagicResist: mr, AttackDamage: ad,
	}
}

// TestCrescimentoBateComOPrevisto: o valor absoluto abaixo esta deslocado de
// propositos por um bonus enorme e constante — 500 de vida, 60 de armadura. O
// crescimento cancela tudo isso, que e o motivo de a comparacao ser dele.
func TestCrescimentoBateComOPrevisto(t *testing.T) {
	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{
		amostra(3, 690+200+500, 38+8+60, 32+4, 69+10),
		amostra(6, 690+500+500, 38+20+60, 32+10, 69+25),
	}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if rel.Diverge() {
		for _, c := range rel.Comparacoes {
			t.Logf("%+v", c)
		}
		t.Fatal("crescimento correto foi acusado de divergir")
	}
	if len(rel.Comparacoes) != 4 {
		t.Fatalf("comparacoes = %d, esperado 4 eixos", len(rel.Comparacoes))
	}
}

// TestCrescimentoAbaixoDoPrevistoDiverge e a razao de existir do comando.
//
// Bonus SOMA, nunca subtrai — entao crescer abaixo do previsto nao tem
// explicacao inocente. Foi um erro dessa forma que deixou a resistencia magica
// publicada como fixa no dataset do modo Jade.
func TestCrescimentoAbaixoDoPrevistoDiverge(t *testing.T) {
	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{
		amostra(3, 890, 46, 36, 79),
		// A resistencia magica nao cresceu nada entre o 3 e o 6.
		amostra(6, 1190, 58, 36, 94),
	}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("crescimento zerado num eixo que deveria crescer passou")
	}
	achou := false
	for _, c := range rel.Comparacoes {
		if c.Eixo == "resistencia magica" && c.Veredito == VereditoDiverge {
			achou = true
			if c.Previsto != 6 || c.Jogo != 0 {
				t.Errorf("numeros errados: jogo=%v previsto=%v", c.Jogo, c.Previsto)
			}
		}
	}
	if !achou {
		t.Fatal("a divergencia nao foi atribuida ao eixo certo")
	}
}

// TestVidaAcimaDoPrevistoEInconclusiva: o fragmento Escalamento de Vida cresce
// COM o nivel e cabem dois, entao ate 20 por nivel a mais e explicavel. E o
// unico eixo assim no LoL moderno, e nao existia no modo Jade.
func TestVidaAcimaDoPrevistoEInconclusiva(t *testing.T) {
	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{
		amostra(3, 890, 46, 36, 79),
		// 3 niveis: 300 previstos, mais 60 dos dois fragmentos.
		amostra(6, 890+360, 58, 42, 94),
	}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if rel.Diverge() {
		t.Fatal("crescimento de vida dentro da margem do fragmento foi acusado de divergir")
	}
	for _, c := range rel.Comparacoes {
		if c.Eixo != "vida" {
			continue
		}
		if c.Veredito != VereditoContaminado {
			t.Fatalf("veredito da vida = %q, esperado inconclusivo", c.Veredito)
		}
		if !strings.Contains(c.Nota, "Escalamento de Vida") {
			t.Errorf("a nota nao explica a margem: %q", c.Nota)
		}
	}
}

// TestVidaMuitoAcimaDaMargemDiverge: a margem tem limite, senao o eixo de vida
// nunca acusaria nada.
func TestVidaMuitoAcimaDaMargemDiverge(t *testing.T) {
	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{
		amostra(3, 890, 46, 36, 79),
		amostra(6, 890+900, 58, 42, 94),
	}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("crescimento de vida muito acima da margem passou")
	}
}

// TestItemMudadoTornaInconclusivo: item comprado entre as leituras quebra o
// cancelamento do bonus fixo, e sem isso uma bota nova viraria divergencia.
func TestItemMudadoTornaInconclusivo(t *testing.T) {
	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{
		amostra(3, 890, 46, 36, 79),
		amostra(6, 1190, 58+45, 42, 94, "Cota de Malha"),
	}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.ItensMudaram {
		t.Fatal("a mudanca de itens nao foi detectada")
	}
	if rel.Diverge() {
		t.Fatal("com item novo, a leitura deveria ser inconclusiva e nao divergente")
	}
}

func TestPrecisaDeDuasAmostras(t *testing.T) {
	_, err := CompararIngame(datasetIngame(), []AmostraIngame{amostra(3, 890, 46, 36, 79)}, 0.02)
	if err == nil {
		t.Fatal("uma amostra so foi aceita")
	}
	if !strings.Contains(err.Error(), "crescimento") {
		t.Fatalf("o erro nao explica por que duas: %v", err)
	}
}

func TestCampeoesDiferentesNaoComparam(t *testing.T) {
	a := amostra(3, 890, 46, 36, 79)
	b := amostra(6, 1190, 58, 42, 94)
	b.Campeao = "Darius"
	if _, err := CompararIngame(datasetIngame(), []AmostraIngame{a, b}, 0.02); err == nil {
		t.Fatal("amostras de campeoes diferentes foram comparadas")
	}
}

func TestCampeaoForaDoDataset(t *testing.T) {
	a := amostra(3, 890, 46, 36, 79)
	b := amostra(6, 1190, 58, 42, 94)
	a.Campeao, b.Campeao = "Inexistente", "Inexistente"
	if _, err := CompararIngame(datasetIngame(), []AmostraIngame{a, b}, 0.02); err == nil {
		t.Fatal("campeao fora do dataset foi comparado")
	}
}

// TestOrdemDasAmostrasNaoImporta: elas sao gravadas conforme a partida avanca, e
// o arquivo pode ser editado a mao.
func TestOrdemDasAmostrasNaoImporta(t *testing.T) {
	fora := []AmostraIngame{
		amostra(6, 1190, 58, 42, 94),
		amostra(3, 890, 46, 36, 79),
	}
	rel, err := CompararIngame(datasetIngame(), fora, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if rel.Diverge() {
		t.Fatal("amostras fora de ordem foram acusadas de divergir")
	}
	if rel.Niveis[0] != 3 || rel.Niveis[1] != 6 {
		t.Fatalf("niveis = %v, esperado ordenados", rel.Niveis)
	}
}
