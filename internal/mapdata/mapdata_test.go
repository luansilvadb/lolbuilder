package mapdata

import (
	"strings"
	"testing"
)

// dumpFake reproduz a forma real do dump: chaves opacas entre chaves, um
// registro de modo por gameMode e listas referenciadas por essas chaves.
const dumpFake = `{
  "Maps/Shipping/Map11/Modes/CLASSIC": {
    "__type": "GameModeMapData",
    "mModeName": "CLASSIC",
    "itemLists": ["{aaaa}", "{bbbb}", "{cccc}"]
  },
  "Maps/Shipping/Map11/Modes/ARAM": {
    "__type": "GameModeMapData",
    "mModeName": "ARAM",
    "itemLists": ["{dddd}"]
  },
  "{aaaa}": {"__type": "GameModeItemList", "mItems": ["Items/1001", "Items/3031"]},
  "{bbbb}": {"__type": "GameModeItemList", "mItems": ["Items/3006", "Items/1001"]},
  "{cccc}": {"__type": "GameModeItemList", "mItems": ["Items/2003"]},
  "{dddd}": {"__type": "GameModeItemList", "mItems": ["Items/9999"]},
  "OutroLixo": {"__type": "VfxSystemDefinitionData"}
}`

// TestShopItemIDsJuntaTodasAsListas e a razao de existir deste pacote: as
// listas sao grupos semanticos e ficar so com a maior deixaria de fora as
// botas avancadas, os itens de Doran e os trinkets.
func TestShopItemIDsJuntaTodasAsListas(t *testing.T) {
	got, err := ShopItemIDs([]byte(dumpFake), 11, "CLASSIC")
	if err != nil {
		t.Fatal(err)
	}
	want := []int32{1001, 2003, 3006, 3031}
	if len(got) != len(want) {
		t.Fatalf("devolveu %v, esperado %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("devolveu %v, esperado %v (ordenado e sem repeticao)", got, want)
		}
	}
}

func TestShopItemIDsIsolaOModo(t *testing.T) {
	got, err := ShopItemIDs([]byte(dumpFake), 11, "ARAM")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0] != 9999 {
		t.Fatalf("o recorte por modo vazou: %v", got)
	}
}

func TestShopItemIDsAceitaGameModeEmMinusculas(t *testing.T) {
	if _, err := ShopItemIDs([]byte(dumpFake), 11, "classic"); err != nil {
		t.Fatalf("gameMode em minusculas deveria resolver: %v", err)
	}
}

func TestShopItemIDsFalhaSemOModo(t *testing.T) {
	_, err := ShopItemIDs([]byte(dumpFake), 11, "URF")
	if err == nil {
		t.Fatal("modo ausente passou em silencio")
	}
	if !strings.Contains(err.Error(), "mode.game_mode") {
		t.Fatalf("o erro nao aponta onde corrigir: %v", err)
	}
}

func TestShopItemIDsFalhaComReferenciaQuebrada(t *testing.T) {
	quebrado := strings.Replace(dumpFake, `"{cccc}"]`, `"{naoexiste}"]`, 1)
	if _, err := ShopItemIDs([]byte(quebrado), 11, "CLASSIC"); err == nil {
		t.Fatal("referencia quebrada passou — a loja sairia menor sem ninguem notar")
	}
}

func TestShopItemIDsFalhaComTipoErrado(t *testing.T) {
	trocado := strings.Replace(dumpFake,
		`"{aaaa}": {"__type": "GameModeItemList"`,
		`"{aaaa}": {"__type": "OutraCoisa"`, 1)
	if _, err := ShopItemIDs([]byte(trocado), 11, "CLASSIC"); err == nil {
		t.Fatal("lista com tipo errado passou")
	}
}

func TestShopItemIDsFalhaComReferenciaMalFormada(t *testing.T) {
	ruim := strings.Replace(dumpFake, `"Items/2003"`, `"Perks/2003"`, 1)
	if _, err := ShopItemIDs([]byte(ruim), 11, "CLASSIC"); err == nil {
		t.Fatal("referencia fora do formato Items/<id> passou")
	}
}

func TestModePath(t *testing.T) {
	if got := ModePath(11, "classic"); got != "Maps/Shipping/Map11/Modes/CLASSIC" {
		t.Fatalf("ModePath = %q", got)
	}
}
