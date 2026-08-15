package optimize

import (
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
)

func item(id int32, custo int32, stat canon.Stat, v float64, botas bool) ItemCandidato {
	return ItemCandidato{
		ID: id, Nome: "item", Custo: custo,
		Stats: canon.Vector{stat: v}, Botas: botas,
	}
}

// exaustivaBuild enumera TODOS os subconjuntos validos e devolve o maior valor.
//
// Lenta de proposito. Existe para conferir a programacao dinamica: um erro ali
// publicaria um maximo que nao e maximo, e nada no dado denunciaria.
func exaustivaBuild(itens []ItemCandidato, obj Objetivo, slots int, orcamento int32) float64 {
	melhor := 0.0
	n := len(itens)
	var rec func(i, usados int, gasto int32, botas bool, v canon.Vector)
	rec = func(i, usados int, gasto int32, botas bool, v canon.Vector) {
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
			if it.Botas && botas {
				continue
			}
			nv := canon.Vector{}
			nv.Merge(v)
			nv.Merge(it.Stats)
			rec(j+1, usados+1, gasto+it.Custo, botas || it.Botas, nv)
		}
	}
	rec(0, 0, 0, false, canon.Vector{})
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
		b, err := MelhorBuild(itens, obj, c.slots, c.orcamento)
		if err != nil {
			t.Fatal(err)
		}
		want := exaustivaBuild(itens, obj, c.slots, c.orcamento)
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
		b, err := MelhorBuild(itens, obj, 6, orcamento)
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
	b, err := MelhorBuild(itens, obj, 6, 10000)
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, it := range b.Itens {
		if it.Botas {
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
	b, _ := MelhorBuild(itens, obj, 6, 6000)
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
	b, _ := MelhorBuild(itens, obj, 6, 6000)
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
	b, _ := MelhorBuild(itens, obj, 6, 1000)
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
	b, _ := MelhorBuild([]ItemCandidato{item(1, 100, canon.Armor, 5, false)}, obj, 6, 1000)

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
	if _, err := MelhorBuild(nil, obj, 0, 1000); err == nil {
		t.Error("zero slots foi aceito")
	}
	if _, err := MelhorBuild(nil, obj, 6, -1); err == nil {
		t.Error("orcamento negativo foi aceito")
	}
}
