package optimize

import (
	"fmt"
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
)

// item monta um candidato. `botas` vira pertinencia ao grupo Botas, que agora e
// so mais um grupo de exclusividade entre os vinte e seis que o jogo declara.
func item(id int32, custo int32, stat canon.Stat, v float64, botas bool) ItemCandidato {
	it := ItemCandidato{ID: id, Nome: "item", Custo: custo, Stats: canon.Vector{stat: v}}
	if botas {
		it.Grupos = []string{"Boots"}
	}
	return it
}

// grupoBotas e a restricao que o pacote conhecia sozinho antes dos grupos.
var grupoBotas = []Grupo{{ID: "Boots", Maximo: 1}}

func temGrupo(it ItemCandidato, g string) bool {
	for _, x := range it.Grupos {
		if x == g {
			return true
		}
	}
	return false
}

// exaustivaBuild enumera TODOS os subconjuntos validos e devolve o maior valor.
//
// Lenta de proposito. Existe para conferir a programacao dinamica: um erro ali
// publicaria um maximo que nao e maximo, e nada no dado denunciaria.
func exaustivaBuild(itens []ItemCandidato, grupos []Grupo, obj Objetivo, slots int, orcamento int32) float64 {
	maxDe := map[string]int{}
	for _, g := range grupos {
		maxDe[g.ID] = g.Maximo
	}

	melhor := 0.0
	n := len(itens)
	var rec func(i, usados int, gasto int32, uso map[string]int, v canon.Vector)
	rec = func(i, usados int, gasto int32, uso map[string]int, v canon.Vector) {
		if val := obj.Valor(v); val > melhor {
			melhor = val
		}
		if i >= n || usados >= slots {
			return
		}
		for j := i; j < n; j++ {
			it := itens[j]
			if gasto+it.Custo > orcamento {
				continue
			}
			cabe := true
			for _, g := range it.Grupos {
				if m, ok := maxDe[g]; ok && uso[g] >= m {
					cabe = false
					break
				}
			}
			if !cabe {
				continue
			}
			novoUso := map[string]int{}
			for k, x := range uso {
				novoUso[k] = x
			}
			for _, g := range it.Grupos {
				novoUso[g]++
			}
			nv := canon.Vector{}
			nv.Merge(v)
			nv.Merge(it.Stats)
			rec(j+1, usados+1, gasto+it.Custo, novoUso, nv)
		}
	}
	rec(0, 0, 0, map[string]int{}, canon.Vector{})
	return melhor
}

// TestMelhorBuildBateComBuscaExaustiva sustenta a promessa de exatidao.
func TestMelhorBuildBateComBuscaExaustiva(t *testing.T) {
	itens := []ItemCandidato{
		item(1, 300, canon.Armor, 15, false),
		item(2, 800, canon.Armor, 45, false),
		item(3, 1100, canon.Armor, 60, false),
		item(4, 500, canon.Health, 200, false),
		item(5, 900, canon.Armor, 40, true),  // botas
		item(6, 1000, canon.Armor, 50, true), // botas
		item(7, 2000, canon.Armor, 100, false),
		item(8, 250, canon.Armor, 10, false),
	}
	obj, err := ParseObjetivo("armor", ResolucaoAD)
	if err != nil {
		t.Fatal(err)
	}

	casos := []struct {
		slots     int
		orcamento int32
	}{
		{1, 1000}, {2, 1500}, {3, 3000}, {6, 5000}, {6, 300}, {4, 0},
	}
	for _, c := range casos {
		b, err := MelhorBuild(itens, grupoBotas, obj, c.slots, c.orcamento)
		if err != nil {
			t.Fatal(err)
		}
		want := exaustivaBuild(itens, grupoBotas, obj, c.slots, c.orcamento)
		if b.Valor != want {
			t.Errorf("slots=%d orcamento=%d: DP deu %v, exaustiva achou %v",
				c.slots, c.orcamento, b.Valor, want)
		}
		// A combinacao devolvida tem de ser valida, e nao so o valor.
		vistos := map[int32]bool{}
		soma := 0.0
		for _, it := range b.Itens {
			if vistos[it.ID] {
				t.Errorf("slots=%d orcamento=%d: item %d repetido", c.slots, c.orcamento, it.ID)
			}
			vistos[it.ID] = true
			soma += obj.Valor(it.Stats)
		}
		if len(b.Itens) > c.slots {
			t.Errorf("slots=%d: a build usou %d itens", c.slots, len(b.Itens))
		}
		if soma != b.Valor {
			t.Errorf("slots=%d orcamento=%d: os itens somam %v mas o valor e %v",
				c.slots, c.orcamento, soma, b.Valor)
		}
	}
}

// TestNenhumItemSeRepeteNaBuild e o teste do defeito que o dado real revelou.
//
// A DP dava o valor certo, e a RECONSTRUCAO caminhava pela grade final seguindo
// um ponteiro de volta que apontava para o estado de outra camada — ja
// sobrescrito por um item posterior. Presságio de Randuin e Couraça do Defunto
// saiam duas vezes na mesma build de armadura.
//
// O caso abaixo reproduz: muitos itens de custo parecido competindo pelas
// mesmas celulas, que e o que faz a sobrescrita acontecer.
func TestNenhumItemSeRepeteNaBuild(t *testing.T) {
	var itens []ItemCandidato
	for i := int32(1); i <= 12; i++ {
		itens = append(itens, item(i, 2000+i*100, canon.Armor, float64(50+i*3), false))
	}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)

	for _, orcamento := range []int32{5000, 12000, 20000, 30000} {
		b, err := MelhorBuild(itens, grupoBotas, obj, 6, orcamento)
		if err != nil {
			t.Fatal(err)
		}
		vistos := map[int32]bool{}
		for _, it := range b.Itens {
			if vistos[it.ID] {
				t.Fatalf("orcamento %d: o item %d saiu duas vezes na build %+v",
					orcamento, it.ID, b.Itens)
			}
			vistos[it.ID] = true
		}
		// A combinacao reconstruida tem de explicar o valor e o gasto.
		soma, custo := 0.0, int32(0)
		for _, it := range b.Itens {
			soma += obj.Valor(it.Stats)
			custo += it.Custo
		}
		if soma != b.Valor {
			t.Errorf("orcamento %d: os itens somam %v mas o valor publicado e %v",
				orcamento, soma, b.Valor)
		}
		if custo != b.Gasto {
			t.Errorf("orcamento %d: os itens custam %d mas o gasto publicado e %d",
				orcamento, custo, b.Gasto)
		}
	}
}

// TestBotasSaoUnicas: o jogo nao deixa carregar duas. Sem a restricao, o
// resultado seria uma build impossivel com cara de otima.
func TestBotasSaoUnicas(t *testing.T) {
	itens := []ItemCandidato{
		item(5, 900, canon.Armor, 40, true),
		item(6, 1000, canon.Armor, 50, true),
		item(7, 1100, canon.Armor, 30, false),
	}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	b, err := MelhorBuild(itens, grupoBotas, obj, 6, 10000)
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, it := range b.Itens {
		if temGrupo(it, "Boots") {
			n++
		}
	}
	if n > 1 {
		t.Fatalf("a build levou %d pares de botas: %+v", n, b.Itens)
	}
	// 50 das botas mais 30 do outro item.
	if b.Valor != 80 {
		t.Fatalf("valor = %v, esperado 80", b.Valor)
	}
}

// TestItemNaoSeRepete: mochila 0/1. Repetir o melhor item seria a solucao
// trivial e errada.
func TestItemNaoSeRepete(t *testing.T) {
	itens := []ItemCandidato{item(1, 100, canon.Armor, 50, false)}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	b, _ := MelhorBuild(itens, grupoBotas, obj, 6, 6000)
	if len(b.Itens) != 1 || b.Valor != 50 {
		t.Fatalf("o item se repetiu: %d itens, valor %v", len(b.Itens), b.Valor)
	}
}

// TestSlotSoEUsadoSeSomar: o resultado nao precisa encher os seis slots.
func TestSlotSoEUsadoSeSomar(t *testing.T) {
	itens := []ItemCandidato{
		item(1, 100, canon.Armor, 50, false),
		item(2, 100, canon.Health, 500, false), // nao pontua sob "armor"
	}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	b, _ := MelhorBuild(itens, grupoBotas, obj, 6, 6000)
	if len(b.Itens) != 1 || b.Itens[0].ID != 1 {
		t.Fatalf("item sem valor entrou na build: %+v", b.Itens)
	}
}

func TestOrcamentoERespeitado(t *testing.T) {
	itens := []ItemCandidato{
		item(1, 3000, canon.Armor, 100, false),
		item(2, 900, canon.Armor, 20, false),
	}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	b, _ := MelhorBuild(itens, grupoBotas, obj, 6, 1000)
	if b.Gasto > 1000 {
		t.Fatalf("gastou %d com orcamento de 1000", b.Gasto)
	}
	if len(b.Itens) != 1 || b.Itens[0].ID != 2 {
		t.Fatalf("build fora do orcamento: %+v", b.Itens)
	}
}

// TestRotuloNaoDizOtima: build otima por stat linear NAO e build boa, e um
// modelo lendo "otima" repassaria a recomendacao com a autoridade do dataset.
func TestRotuloNaoDizOtima(t *testing.T) {
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	b, _ := MelhorBuild([]ItemCandidato{item(1, 100, canon.Armor, 5, false)}, nil, obj, 6, 1000)

	if !strings.Contains(b.Rotulo, "ignorando efeitos") {
		t.Fatalf("o rotulo nao avisa da limitacao: %q", b.Rotulo)
	}
	if strings.Contains(strings.ToLower(b.Rotulo), "otima") {
		t.Fatalf("o rotulo chama o resultado de otimo: %q", b.Rotulo)
	}
	if b.Resolucao != ResolucaoAD || b.Orcamento != 1000 || b.Slots != 6 {
		t.Errorf("o resultado nao carrega as premissas: %+v", b)
	}
}

func TestArgumentosInvalidos(t *testing.T) {
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	if _, err := MelhorBuild(nil, nil, obj, 0, 1000); err == nil {
		t.Error("zero slots foi aceito")
	}
	if _, err := MelhorBuild(nil, nil, obj, 6, -1); err == nil {
		t.Error("orcamento negativo foi aceito")
	}
}

// TestGruposDeExclusividadeBatemComBuscaExaustiva e a garantia de exatidao sob a
// restricao nova. A busca exaustiva respeita os mesmos grupos, senao a
// verificacao seria circular.
//
// O caso reproduz o que o dado real revelou: varios itens do mesmo grupo com
// valor alto sob o objetivo (foi assim que a build de penetracao de armadura
// juntou QUATRO itens de LastWhisper), e um item pertencendo a DOIS grupos ao
// mesmo tempo, como o Terminus, que esta em LastWhisper e em VoidPen.
func TestGruposDeExclusividadeBatemComBuscaExaustiva(t *testing.T) {
	g := func(it ItemCandidato, gs ...string) ItemCandidato {
		it.Grupos = gs
		return it
	}
	itens := []ItemCandidato{
		g(item(1, 3000, canon.Armor, 100, false), "LastWhisper"),
		g(item(2, 3000, canon.Armor, 95, false), "LastWhisper"),
		g(item(3, 3300, canon.Armor, 90, false), "LastWhisper"),
		g(item(4, 3000, canon.Armor, 85, false), "LastWhisper", "VoidPen"),
		g(item(5, 2500, canon.Armor, 80, false), "VoidPen"),
		g(item(6, 2500, canon.Armor, 75, false), "VoidPen"),
		g(item(7, 1000, canon.Armor, 40, false), "Tear"),
		g(item(8, 1000, canon.Armor, 38, false), "Tear"),
		item(9, 900, canon.Armor, 35, true),  // botas
		item(10, 950, canon.Armor, 33, true), // botas
		item(11, 2000, canon.Armor, 60, false),
		item(12, 1800, canon.Armor, 55, false),
	}
	grupos := []Grupo{
		{ID: "LastWhisper", Maximo: 1},
		{ID: "VoidPen", Maximo: 1},
		{ID: "Tear", Maximo: 1},
		{ID: "Boots", Maximo: 1},
	}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)

	for _, c := range []struct {
		slots     int
		orcamento int32
	}{{6, 20000}, {6, 8000}, {3, 20000}, {2, 6000}, {6, 3000}} {
		b, err := MelhorBuild(itens, grupos, obj, c.slots, c.orcamento)
		if err != nil {
			t.Fatal(err)
		}
		if want := exaustivaBuild(itens, grupos, obj, c.slots, c.orcamento); b.Valor != want {
			t.Errorf("slots=%d ouro=%d: DP deu %v, exaustiva achou %v",
				c.slots, c.orcamento, b.Valor, want)
		}
		// E a combinacao devolvida tem de ser compravel de verdade.
		uso := map[string]int{}
		for _, it := range b.Itens {
			for _, gr := range it.Grupos {
				uso[gr]++
			}
		}
		for _, gr := range grupos {
			if uso[gr.ID] > gr.Maximo {
				t.Errorf("slots=%d ouro=%d: a build leva %d itens do grupo %s (max %d): %+v",
					c.slots, c.orcamento, uso[gr.ID], gr.ID, gr.Maximo, b.Itens)
			}
		}
	}
}

// TestGrupoSemMembrosUteisNaoEntraNaGrade: um grupo que a poda de dominancia
// esvaziou nao pode custar dimensao. Sem isso os 26 grupos do 16.16 dariam 2^26
// estados e a mochila nao rodaria.
func TestGrupoSemMembrosUteisNaoEntraNaGrade(t *testing.T) {
	var itens []ItemCandidato
	var grupos []Grupo
	for i := int32(1); i <= 30; i++ {
		// Todos pontuam em vida, e o objetivo e armadura: a poda tira todos.
		itens = append(itens, g30(i))
		grupos = append(grupos, Grupo{ID: fmt.Sprintf("G%d", i), Maximo: 1})
	}
	itens = append(itens, item(999, 100, canon.Armor, 50, false))

	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	b, err := MelhorBuild(itens, grupos, obj, 6, 6000)
	if err != nil {
		t.Fatalf("a mochila recusou por causa de grupos que nao restringem nada: %v", err)
	}
	if b.Valor != 50 {
		t.Fatalf("valor = %v, esperado 50", b.Valor)
	}
}

func g30(i int32) ItemCandidato {
	it := item(i, 100, canon.Health, 100, false)
	it.Grupos = []string{fmt.Sprintf("G%d", i)}
	return it
}

// TestMuitosGruposAtivosRecusa: a mochila prefere recusar a aproximar. Um otimo
// aproximado publicado como exato e o defeito que os grupos vieram corrigir.
func TestMuitosGruposAtivosRecusa(t *testing.T) {
	var itens []ItemCandidato
	var grupos []Grupo
	for i := int32(0); i < 20; i++ {
		nome := fmt.Sprintf("G%d", i)
		grupos = append(grupos, Grupo{ID: nome, Maximo: 1})
		for j := int32(0); j < 2; j++ { // dois membros uteis: o grupo fica ativo
			it := item(i*10+j, 100, canon.Armor, float64(10+j), false)
			it.Grupos = []string{nome}
			itens = append(itens, it)
		}
	}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	_, err := MelhorBuild(itens, grupos, obj, 6, 20000)
	if err == nil {
		t.Fatal("a mochila aceitou um numero de grupos que nao cabe na grade")
	}
	if !strings.Contains(err.Error(), "recusa em vez de aproximar") {
		t.Fatalf("o erro nao explica a recusa: %v", err)
	}
}
