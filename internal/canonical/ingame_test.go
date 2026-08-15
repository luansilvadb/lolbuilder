package canonical

import (
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/gamedata"
)

// O campeao de teste tem crescimento redondo, para que a conta seja conferivel.
var (
	statsDeTeste = gamedata.Stats{
		HP:           gamedata.Scaling{Base: 690, PerLevel: 100},
		Armor:        gamedata.Scaling{Base: 38, PerLevel: 4},
		MagicResist:  gamedata.Scaling{Base: 32, PerLevel: 2},
		AttackDamage: gamedata.Scaling{Base: 69, PerLevel: 5},
	}
)

func datasetIngame() *Dataset {
	s := statsDeTeste
	return &Dataset{
		Patch: "16.99",
		Champions: []Champion{{
			ID: 86, Nome: "Garen", NomeCanonico: "Garen", Alias: "Garen", Stats: &s,
		}},
	}
}

func amostra(nivel int, vida, armadura, mr, ad float64, itens ...string) AmostraIngame {
	return AmostraIngame{
		Nivel: nivel, Campeao: "Garen", Itens: itens,
		MaxHealth: vida, Armor: armadura, MagicResist: mr, AttackDamage: ad,
	}
}

// depoisDeSubir monta a leitura esperada no nivel de destino, aplicando o
// crescimento que o dataset afirma.
//
// A conta sai daqui e nao de numeros digitados no teste: com o fator nao
// linear, escrever os valores a mao convida a erro de aritmetica que passaria
// por defeito do codigo.
func depoisDeSubir(base AmostraIngame, ate int, itens ...string) AmostraIngame {
	de := base.Nivel
	return amostra(ate,
		base.MaxHealth+statsDeTeste.HP.CrescimentoEntre(de, ate),
		base.Armor+statsDeTeste.Armor.CrescimentoEntre(de, ate),
		base.MagicResist+statsDeTeste.MagicResist.CrescimentoEntre(de, ate),
		base.AttackDamage+statsDeTeste.AttackDamage.CrescimentoEntre(de, ate),
		itens...)
}

// TestCrescimentoBateComOPrevisto: o valor absoluto abaixo esta deslocado de
// proposito por um bonus enorme e constante — 500 de vida, 60 de armadura. O
// crescimento cancela tudo isso, que e o motivo de a comparacao ser dele.
func TestCrescimentoBateComOPrevisto(t *testing.T) {
	inicio := amostra(3, 690+200+500, 38+8+60, 32+4, 69+10)
	rel, err := CompararIngame(datasetIngame(),
		[]AmostraIngame{inicio, depoisDeSubir(inicio, 6)}, 0.02)
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
// explicacao inocente. Foi um erro exatamente dessa forma que o oraculo pegou
// numa partida real: o crescimento por nivel do LoL nao e linear, e o modelo
// linear previa 10.25 de resistencia magica onde o jogo somou 8.0975.
func TestCrescimentoAbaixoDoPrevistoDiverge(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	fim.MagicResist = inicio.MagicResist // a resistencia magica nao cresceu nada

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("crescimento zerado num eixo que deveria crescer passou")
	}
	for _, c := range rel.Comparacoes {
		if c.Eixo != "resistencia magica" {
			continue
		}
		if c.Veredito != VereditoDiverge {
			t.Fatalf("veredito = %q, esperado DIVERGE", c.Veredito)
		}
		if !strings.Contains(c.Nota, "bonus soma") {
			t.Errorf("a nota nao explica o motivo: %q", c.Nota)
		}
		return
	}
	t.Fatal("a divergencia nao foi atribuida ao eixo certo")
}

// TestVidaAcimaDoPrevistoEInconclusiva: o fragmento Escalamento de Vida cresce
// COM o nivel e cabem dois, entao ate 20 por nivel a mais e explicavel. E o
// unico eixo assim no LoL moderno, e nao existia no modo Jade.
func TestVidaAcimaDoPrevistoEInconclusiva(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	fim.MaxHealth += 60 // dois fragmentos, 10 por nivel cada, tres niveis

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02)
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
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	fim.MaxHealth += 900

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("crescimento de vida muito acima da margem passou")
	}
}

// TestExcessoSemItemNovoTambemDiverge: crescer acima do previsto so e explicavel
// por algo que some. Com os itens iguais e sem runa que escale naquele eixo, nao
// ha o que some — e tratar todo excesso como inconclusivo deixaria a verificacao
// cega para metade dos erros possiveis.
func TestExcessoSemItemNovoTambemDiverge(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	fim.Armor += 50 // armadura nao tem runa que escale por nivel

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("excesso de armadura sem item novo passou")
	}
}

// TestItemMudadoTornaInconclusivo: item comprado entre as leituras quebra o
// cancelamento do bonus fixo, e sem isso uma bota nova viraria divergencia.
func TestItemMudadoTornaInconclusivo(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6, "Cota de Malha")
	fim.Armor += 45

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02)
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

// TestItemNovoNaoMascaraCrescimentoAbaixo: item soma, nunca subtrai — entao nem
// com item novo o crescimento abaixo do previsto tem explicacao. Foi assim que o
// crescimento nao linear apareceu, com cinco itens equipados no caminho.
func TestItemNovoNaoMascaraCrescimentoAbaixo(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6, "Cota de Malha")
	fim.Armor += 45
	fim.MagicResist -= 3 // cresceu menos que o previsto

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("crescimento abaixo do previsto foi mascarado pela mudanca de itens")
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
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	fim.Campeao = "Darius"
	if _, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02); err == nil {
		t.Fatal("amostras de campeoes diferentes foram comparadas")
	}
}

func TestCampeaoForaDoDataset(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	inicio.Campeao, fim.Campeao = "Inexistente", "Inexistente"
	if _, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02); err == nil {
		t.Fatal("campeao fora do dataset foi comparado")
	}
}

// TestOrdemDasAmostrasNaoImporta: elas sao gravadas conforme a partida avanca, e
// o arquivo pode ser editado a mao.
func TestOrdemDasAmostrasNaoImporta(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fora := []AmostraIngame{depoisDeSubir(inicio, 6), inicio}

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
