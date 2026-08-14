package canon

import (
	"strings"
	"testing"
)

func TestLerStatsDeItemFormaCanonica(t *testing.T) {
	// Infinity Edge, como a fonte publica.
	desc := `<mainText><stats><attention> 75</attention> Attack Damage<br>` +
		`<attention> 25%</attention> Critical Strike Chance<br>` +
		`<attention> 30%</attention> Critical Strike Damage</stats><br><br></mainText>`

	l := LerStatsDeItem("Infinity Edge", desc)
	if !l.Completa() {
		t.Fatalf("leitura incompleta: %d de %d linhas, nao lidas %+v", l.Lidas, l.Linhas, l.NaoLidas)
	}
	want := Vector{AttackDamage: 75, CriticalChancePct: 25, CriticalDamagePct: 30}
	comparar(t, l.Stats, want)
}

// TestBotasTrazemDeslocamento e o teste que fixa a diferenca mais importante em
// relacao ao dataset do modo Jade: la o deslocamento das botas ficava fora do
// bloco de atributos e nao entrava no total. Aqui entra.
func TestBotasTrazemDeslocamento(t *testing.T) {
	desc := `<mainText><stats><attention> 30%</attention> Attack Speed<br>` +
		`<attention> 45</attention> Move Speed</stats><br><br></mainText>`

	l := LerStatsDeItem("Berserker's Greaves", desc)
	if !l.Completa() {
		t.Fatalf("leitura incompleta: %+v", l.NaoLidas)
	}
	comparar(t, l.Stats, Vector{AttackSpeedPct: 30, MoveSpeed: 45})
}

// TestUnidadeSeparaGrandezas: a fonte usa o mesmo rotulo para deslocamento
// plano e percentual, e para penetracao magica plana e percentual. Somar os
// dois no mesmo campo somaria 45 de deslocamento com 4% de deslocamento.
func TestUnidadeSeparaGrandezas(t *testing.T) {
	desc := `<stats><attention> 45</attention> Move Speed<br>` +
		`<attention> 4%</attention> Move Speed<br>` +
		`<attention> 15</attention> Magic Penetration<br>` +
		`<attention> 30%</attention> Magic Penetration</stats>`

	l := LerStatsDeItem("Item Sintetico", desc)
	if !l.Completa() {
		t.Fatalf("leitura incompleta: %+v", l.NaoLidas)
	}
	comparar(t, l.Stats, Vector{
		MoveSpeed: 45, MoveSpeedPct: 4,
		MagicPenetration: 15, MagicPenetrationPct: 30,
	})
}

// TestOrnnBonusEFormaValida: os itens aprimorados pelo Ornn usam outra tag para
// envolver o mesmo valor. Nenhum deles e compravel hoje, e por isso a forma nao
// aparece no subconjunto publicado — reconhece-la agora e mais barato que
// descobri-la no patch em que passar a importar.
func TestOrnnBonusEFormaValida(t *testing.T) {
	desc := `<stats><attention> 60</attention> Attack Damage<br>` +
		`<ornnBonus> 20</ornnBonus> Ability Haste</stats>`

	l := LerStatsDeItem("Item do Ornn", desc)
	if !l.Completa() {
		t.Fatalf("forma <ornnBonus> nao foi lida: %+v", l.NaoLidas)
	}
	comparar(t, l.Stats, Vector{AttackDamage: 60, AbilityHaste: 20})
}

func TestValorDecimal(t *testing.T) {
	// Doran's Blade publica 2.5% de omnivamp.
	l := LerStatsDeItem("Doran's Blade", `<stats><attention> 2.5%</attention> Omnivamp</stats>`)
	if !l.Completa() {
		t.Fatalf("decimal nao foi lido: %+v", l.NaoLidas)
	}
	if got := l.Stats[OmnivampPct]; got != 2.5 {
		t.Fatalf("omnivamp = %v, esperado 2.5", got)
	}
}

// TestSemBlocoNaoEFalhaDeLeitura: World Atlas vem com descricao vazia da fonte.
// Isso e lacuna da fonte, e nao pode contar como parser que falhou.
func TestSemBlocoNaoEFalhaDeLeitura(t *testing.T) {
	l := LerStatsDeItem("World Atlas", "")
	if l.TemBloco {
		t.Fatal("descricao vazia foi tratada como bloco")
	}
	if l.Completa() {
		t.Fatal("item sem bloco nao pode contar como lido por inteiro")
	}
	if len(l.NaoLidas) != 0 {
		t.Fatalf("item sem bloco gerou linha nao lida: %+v", l.NaoLidas)
	}

	var cob CoberturaDeItens
	cob.Acumular("World Atlas", l)
	if cob.ComBloco != 0 || cob.SemBloco != 1 {
		t.Fatalf("item sem bloco entrou no denominador: %+v", cob)
	}
}

// TestRotuloDesconhecidoNaoViraZero e a regra central do projeto: a linha e
// reportada, o stat nao e publicado, e nenhum numero e inventado.
func TestRotuloDesconhecidoNaoViraZero(t *testing.T) {
	desc := `<stats><attention> 40</attention> Attack Damage<br>` +
		`<attention> 12</attention> Grandeza Que Nao Existe</stats>`

	l := LerStatsDeItem("Item Novo", desc)
	if l.Completa() {
		t.Fatal("rotulo desconhecido passou como lido")
	}
	if l.Lidas != 1 || l.Linhas != 2 {
		t.Fatalf("contagem errada: %d de %d", l.Lidas, l.Linhas)
	}
	if _, ok := l.Stats["grandeza_que_nao_existe"]; ok {
		t.Fatal("rotulo desconhecido virou stat")
	}
	if len(l.NaoLidas) != 1 {
		t.Fatalf("a linha nao foi reportada: %+v", l.NaoLidas)
	}
	if !strings.Contains(l.NaoLidas[0].Motivo, "vocabulario") {
		t.Fatalf("o motivo nao explica o que houve: %q", l.NaoLidas[0].Motivo)
	}
	// O stat que deu certo continua publicado: uma linha ruim nao derruba o item.
	if l.Stats[AttackDamage] != 40 {
		t.Fatalf("a linha boa foi perdida junto: %+v", l.Stats)
	}
}

// TestUnidadeErradaNaoCoage: "Attack Speed" existe so como percentual. Vindo
// plano, e mudanca de forma na fonte e precisa aparecer, nao ser convertida.
func TestUnidadeErradaNaoCoage(t *testing.T) {
	l := LerStatsDeItem("Item", `<stats><attention> 30</attention> Attack Speed</stats>`)
	if l.Completa() {
		t.Fatal("unidade divergente foi aceita em silencio")
	}
	if !strings.Contains(l.NaoLidas[0].Motivo, "plano") {
		t.Fatalf("o motivo nao diz qual unidade veio: %q", l.NaoLidas[0].Motivo)
	}
}

func TestLinhaForaDaForma(t *testing.T) {
	l := LerStatsDeItem("Item", `<stats>texto solto sem marcacao</stats>`)
	if l.Linhas != 1 || l.Lidas != 0 || len(l.NaoLidas) != 1 {
		t.Fatalf("linha malformada nao foi reportada: %+v", l)
	}
}

func TestCoberturaDeItens(t *testing.T) {
	var cob CoberturaDeItens
	cob.Acumular("Bom", LerStatsDeItem("Bom", `<stats><attention> 40</attention> Attack Damage</stats>`))
	cob.Acumular("Parcial", LerStatsDeItem("Parcial",
		`<stats><attention> 40</attention> Attack Damage<br><attention> 1</attention> Xis</stats>`))
	cob.Acumular("Sem", LerStatsDeItem("Sem", ""))

	if cob.ComBloco != 2 || cob.SemBloco != 1 || cob.LidosPorInteiro != 1 {
		t.Fatalf("cobertura errada: %+v", cob)
	}
	if got := cob.Taxa(); got != 50 {
		t.Fatalf("taxa = %v, esperado 50 (1 de 2 com bloco)", got)
	}
	if cob.Linhas != 3 || cob.LinhasLidas != 2 {
		t.Fatalf("contagem de linhas errada: %+v", cob)
	}
}

func TestTaxaSemItensNaoDivideporZero(t *testing.T) {
	var cob CoberturaDeItens
	if got := cob.Taxa(); got != 0 {
		t.Fatalf("Taxa() = %v em cobertura vazia", got)
	}
}

func comparar(t *testing.T, got, want Vector) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("vetor tem %d stats, esperado %d: %+v", len(got), len(want), got)
	}
	for s, v := range want {
		if got[s] != v {
			t.Errorf("%s = %v, esperado %v", s, got[s], v)
		}
	}
}
