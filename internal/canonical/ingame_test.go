package canonical

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
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
		[]AmostraIngame{inicio, depoisDeSubir(inicio, 6)}, 0.02, nil)
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

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil)
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

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil)
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

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil)
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

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil)
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

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil)
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

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("crescimento abaixo do previsto foi mascarado pela mudanca de itens")
	}
}

// naSessao devolve a amostra marcada como de outra partida.
func naSessao(a AmostraIngame, s int) AmostraIngame {
	a.Sessao = s
	return a
}

// TestSessoesDiferentesNaoComparam e a razao de o campo Sessao existir.
//
// O arquivo de amostras acumula entre execucoes. Sem separar por partida, as
// leituras se ordenavam so por nivel e o par comparado podia ter um lado de cada
// partida — com runas e fragmentos diferentes, e mesmo assim veredito confiante.
func TestSessoesDiferentesNaoComparam(t *testing.T) {
	p1 := amostra(3, 890, 46, 36, 79)
	// Segunda partida: mesmo campeao e mesmos niveis, mas 200 de vida a menos
	// porque a pagina de runas era outra. Comparar entre as duas daria -200.
	p2 := naSessao(amostra(6, 690, 38, 32, 69), 1)

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{p1, p2}, 0.02, nil)
	if err != nil {
		t.Fatal(err)
	}
	if rel.Diverge() {
		t.Fatal("amostras de partidas diferentes foram comparadas entre si")
	}
	if len(rel.Comparacoes) != 0 {
		t.Fatalf("comparacoes = %d, esperado nenhuma: nao ha par dentro de uma mesma partida",
			len(rel.Comparacoes))
	}
	if rel.Sessoes != 2 {
		t.Fatalf("sessoes = %d, esperado 2", rel.Sessoes)
	}
	if len(rel.SessoesSemPar) != 2 {
		t.Fatalf("sessoes sem par = %v, esperado as duas avisadas", rel.SessoesSemPar)
	}
}

// TestCadaSessaoCompararSeparadamente: separar partidas nao pode custar as
// comparacoes validas de dentro de cada uma.
func TestCadaSessaoCompararSeparadamente(t *testing.T) {
	a1 := amostra(3, 890, 46, 36, 79)
	a2 := depoisDeSubir(a1, 6)
	b1 := naSessao(amostra(4, 1200, 60, 40, 90), 1)
	b2 := naSessao(depoisDeSubir(b1, 9), 1)

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{a1, b2, a2, b1}, 0.02, nil)
	if err != nil {
		t.Fatal(err)
	}
	if rel.Diverge() {
		for _, c := range rel.Comparacoes {
			t.Logf("%+v", c)
		}
		t.Fatal("crescimento correto dentro de cada partida foi acusado de divergir")
	}
	if len(rel.Comparacoes) != 8 {
		t.Fatalf("comparacoes = %d, esperado 4 eixos em cada uma das 2 partidas",
			len(rel.Comparacoes))
	}
	if len(rel.SessoesSemPar) != 0 {
		t.Fatalf("sessoes sem par = %v, esperado nenhuma", rel.SessoesSemPar)
	}
}

// TestItemNovoNaoContaminaOProximoPar: o sinalizador de itens era do relatorio
// inteiro, entao uma compra no par 1→6 desculpava um excesso no par 6→7 — que
// nao tinha nada a ver com ela. Metade dos pares de uma partida com compras
// ficava cega.
func TestItemNovoNaoContaminaOProximoPar(t *testing.T) {
	a1 := amostra(3, 890, 46, 36, 79)
	a2 := depoisDeSubir(a1, 6, "Cota de Malha") // compra aqui
	a2.Armor += 45
	a3 := depoisDeSubir(a2, 9, "Cota de Malha") // nenhuma compra daqui pra frente
	a3.Armor += 50                              // excesso sem explicacao

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{a1, a2, a3}, 0.02, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("o excesso no par 6→9 foi desculpado pela compra ocorrida no par 3→6")
	}
	for _, c := range rel.Comparacoes {
		if c.Eixo == "armadura" && c.DeNivel == 3 && c.Veredito != VereditoContaminado {
			t.Errorf("o par 3→6 tinha item novo e deveria ser inconclusivo, veio %q", c.Veredito)
		}
	}
}

func TestPrecisaDeDuasAmostras(t *testing.T) {
	_, err := CompararIngame(datasetIngame(), []AmostraIngame{amostra(3, 890, 46, 36, 79)}, 0.02, nil)
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
	if _, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil); err == nil {
		t.Fatal("amostras de campeoes diferentes foram comparadas")
	}
}

func TestCampeaoForaDoDataset(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	inicio.Campeao, fim.Campeao = "Inexistente", "Inexistente"
	if _, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil); err == nil {
		t.Fatal("campeao fora do dataset foi comparado")
	}
}

// TestOrdemDasAmostrasNaoImporta: elas sao gravadas conforme a partida avanca, e
// o arquivo pode ser editado a mao.
func TestOrdemDasAmostrasNaoImporta(t *testing.T) {
	inicio := amostra(3, 890, 46, 36, 79)
	fora := []AmostraIngame{depoisDeSubir(inicio, 6), inicio}

	rel, err := CompararIngame(datasetIngame(), fora, 0.02, nil)
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

// TestConversaoDePassivaEntraNoPrevisto: a API reporta o valor TOTAL do eixo, ja
// com o que a passiva concedeu. Sem somar a conversao ao previsto, o oraculo
// acusa divergencia onde o dataset e o jogo estao os dois certos.
func TestConversaoDePassivaEntraNoPrevisto(t *testing.T) {
	conv := &canon.CuradoriaDeConversoes{Conversoes: []canon.Conversao{{
		Campeao: "Garen", Eixo: canon.AttackDamage,
		Origem:    map[canon.Stat]float64{canon.Armor: 0.15, canon.MagicResist: 0.15},
		Evidencia: "teste",
	}}}

	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	// O jogo entrega, alem do crescimento base, 15% do que armadura e
	// resistencia magica cresceram no caminho.
	fim.AttackDamage += 0.15 * (fim.Armor - inicio.Armor)
	fim.AttackDamage += 0.15 * (fim.MagicResist - inicio.MagicResist)

	comConv, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, conv)
	if err != nil {
		t.Fatal(err)
	}
	if comConv.Diverge() {
		for _, c := range comConv.Comparacoes {
			t.Logf("%+v", c)
		}
		t.Fatal("a conversao da passiva nao foi somada ao previsto")
	}

	// E sem a curadoria o mesmo dado tem de acusar: se passasse dos dois jeitos,
	// a tabela nao estaria fazendo diferenca nenhuma.
	semConv, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !semConv.Diverge() {
		t.Fatal("sem a curadoria o excesso da passiva deveria acusar")
	}
}

// TestConversaoNaoAfrouxaAVerificacao: somar a passiva ao previsto so vale se o
// eixo continuar acusando erro. Com o coeficiente errado na curadoria, o
// resultado tem de divergir — e assim a tabela vira teste dela mesma.
func TestConversaoNaoAfrouxaAVerificacao(t *testing.T) {
	errada := &canon.CuradoriaDeConversoes{Conversoes: []canon.Conversao{{
		Campeao: "Garen", Eixo: canon.AttackDamage,
		Origem:    map[canon.Stat]float64{canon.Armor: 0.40},
		Evidencia: "coeficiente de proposito errado",
	}}}

	inicio := amostra(3, 890, 46, 36, 79)
	fim := depoisDeSubir(inicio, 6)
	fim.AttackDamage += 0.15 * (fim.Armor - inicio.Armor)

	rel, err := CompararIngame(datasetIngame(), []AmostraIngame{inicio, fim}, 0.02, errada)
	if err != nil {
		t.Fatal(err)
	}
	if !rel.Diverge() {
		t.Fatal("coeficiente errado na curadoria passou despercebido")
	}
}

// TestCuradoriaDeConversoesExigeEvidencia: sem medicao que a sustente, a entrada
// e palpite — e palpite aqui faz o oraculo confirmar dataset errado.
func TestCuradoriaDeConversoesExigeEvidencia(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "conversoes.json")
	os.WriteFile(p, []byte(`{"conversoes":[{"campeao":"Garen","eixo":"attack_damage",
		"origem":{"armor":0.15}}]}`), 0o644)

	if _, err := canon.LoadConversoes(p); err == nil {
		t.Fatal("conversao sem evidencia foi aceita")
	} else if !strings.Contains(err.Error(), "evidencia") {
		t.Fatalf("o erro nao aponta a evidencia: %v", err)
	}
}

// TestCuradoriaAusenteNaoEErro: a tabela e opcional, e exigi-la quebraria quem
// so quer rodar o build.
func TestCuradoriaAusenteNaoEErro(t *testing.T) {
	c, err := canon.LoadConversoes(filepath.Join(t.TempDir(), "nao-existe.json"))
	if err != nil {
		t.Fatalf("arquivo ausente virou erro: %v", err)
	}
	if len(c.Conversoes) != 0 {
		t.Fatal("curadoria ausente veio com entradas")
	}
}

// TestCuradoriaRealDoRammusFechaNaMedicao trava a evidencia que a curadoria
// declara contra o arquivo em disco.
//
// Sem isto, alguem poderia editar o coeficiente e nenhum teste notaria — e o
// oraculo passaria a confirmar o dataset com uma conta inventada. Os numeros
// abaixo sao os medidos na Ferramenta de Treino, sem item algum.
func TestCuradoriaRealDoRammusFechaNaMedicao(t *testing.T) {
	c, err := canon.LoadConversoes(filepath.Join("..", "..", "curation", "conversoes.json"))
	if err != nil {
		t.Fatal(err)
	}
	cvs := c.Do([]string{"Rammus"}, canon.AttackDamage)
	if len(cvs) != 1 {
		t.Fatalf("conversoes do Rammus para dano de ataque = %d, esperado 1", len(cvs))
	}

	// Crescimento medido em partida, do nivel 1 ao 6 e do 6 ao 7.
	casos := []struct{ armadura, mr, excesso float64 }{
		{17.7750, 8.0975, 3.8809},
		{4.0275, 1.8348, 0.8793},
	}
	for _, k := range casos {
		previsto := cvs[0].Origem[canon.Armor]*k.armadura + cvs[0].Origem[canon.MagicResist]*k.mr
		if d := previsto - k.excesso; d > 0.0001 || d < -0.0001 {
			t.Errorf("armadura %.4f e mr %.4f: a curadoria preve %.6f, medido %.4f (resto %.6f)",
				k.armadura, k.mr, previsto, k.excesso, d)
		}
	}
}
