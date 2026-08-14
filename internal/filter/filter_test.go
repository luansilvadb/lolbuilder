package filter

import (
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/model"
)

var modo = config.Mode{
	Name: "classic", GameMode: "CLASSIC", MapID: 11,
	ItemIDMin: 1000, ItemIDMax: 699999,
	ChampionIDMin: 1, ChampionIDMax: 999,
}

// TestItemsInRangeSeparaOsDoisJogos protege o recorte que impede o dataset de
// misturar o LoL moderno com o modo Jade, que a Riot publica no mesmo arquivo.
func TestItemsInRangeSeparaOsDoisJogos(t *testing.T) {
	items := []model.Item{
		{ID: 999, Name: "abaixo da faixa"},
		{ID: 1001, Name: "Boots"},
		{ID: 3031, Name: "Infinity Edge"},
		{ID: 699999, Name: "limite superior"},
		{ID: 771001, Name: "Boots of Speed (Jade)"},
		{ID: 994403, Name: "Jade fora do teto antigo"},
	}
	got := ItemsInRange(items, modo)
	if len(got) != 3 {
		t.Fatalf("devolveu %d itens, esperado 3: %+v", len(got), got)
	}
	for _, it := range got {
		if it.ID < modo.ItemIDMin || it.ID > modo.ItemIDMax {
			t.Errorf("item %d escapou da faixa", it.ID)
		}
	}
}

// TestChampionsInRangeDescartaSentinela: champion-summary.json publica uma
// entrada de id -1 chamada "None" junto dos campeoes de verdade.
func TestChampionsInRangeDescartaSentinela(t *testing.T) {
	cs := []model.ChampionSummary{
		{ID: -1, Name: "None"},
		{ID: 1, Name: "Annie"},
		{ID: 910, Name: "Hwei"},
		{ID: 60001, Name: "Annie (Jade)"},
	}
	got := ChampionsInRange(cs, modo)
	if len(got) != 2 {
		t.Fatalf("devolveu %d campeoes, esperado 2: %+v", len(got), got)
	}
}

func TestSummonerSpellsForMode(t *testing.T) {
	spells := []model.SummonerSpell{
		{ID: 4, Name: "Flash", GameModes: []string{"ARAM", "CLASSIC", "URF"}},
		{ID: 32, Name: "Mark", GameModes: []string{"ARAM"}},
		{ID: 11, Name: "Smite", GameModes: []string{"classic"}},
		{ID: 99, Name: "Sem modo"},
	}
	got := SummonerSpellsForMode(spells, modo)
	if len(got) != 2 {
		t.Fatalf("devolveu %d feiticos, esperado 2: %+v", len(got), got)
	}
	if got[0].Name != "Flash" || got[1].Name != "Smite" {
		t.Fatalf("recorte errado: %+v", got)
	}
}

// TestUnknownShopIDs: referencia orfa significa que a loja veio de um mapa cujo
// catalogo nao e este, e isso publica item fantasma se passar calado.
func TestUnknownShopIDs(t *testing.T) {
	catalog := ItemsByID([]model.Item{{ID: 1001}, {ID: 3031}})

	if got := UnknownShopIDs([]int32{1001, 3031}, catalog); len(got) != 0 {
		t.Fatalf("acusou orfao onde nao ha: %v", got)
	}
	got := UnknownShopIDs([]int32{1001, 424242}, catalog)
	if len(got) != 1 || got[0] != 424242 {
		t.Fatalf("nao acusou o orfao: %v", got)
	}
}

// TestUnknownShopIDsUsaOCatalogoCompleto fixa a distincao que custou uma
// captura abortada: a loja do Summoner's Rift referencia 771500, que existe em
// items.json mas mora na faixa do modo Jade. Isso e fora da faixa, nao orfao —
// conferir contra o catalogo ja recortado transformava um fato da fonte em erro
// fatal.
func TestUnknownShopIDsUsaOCatalogoCompleto(t *testing.T) {
	completo := ItemsByID([]model.Item{{ID: 1001}, {ID: 771500}})
	if got := UnknownShopIDs([]int32{1001, 771500}, completo); len(got) != 0 {
		t.Fatalf("id fora da faixa foi tratado como orfao: %v", got)
	}
}

func TestShopIDsOutOfRange(t *testing.T) {
	got := ShopIDsOutOfRange([]int32{1001, 3031, 771500}, modo)
	if len(got) != 1 || got[0] != 771500 {
		t.Fatalf("ShopIDsOutOfRange = %v, esperado [771500]", got)
	}
	if got := ShopIDsOutOfRange([]int32{1001, 3031}, modo); len(got) != 0 {
		t.Fatalf("acusou fora da faixa onde nao ha: %v", got)
	}
}
