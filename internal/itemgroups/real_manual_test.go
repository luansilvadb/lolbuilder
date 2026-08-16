package itemgroups

import (
	"encoding/json"
	"os"
	"testing"
)

// TestDumpReal roda so quando LOLBUILDER_ITEMS_BIN aponta para o dump baixado.
func TestDumpReal(t *testing.T) {
	p := os.Getenv("LOLBUILDER_ITEMS_BIN")
	if p == "" {
		t.Skip("sem LOLBUILDER_ITEMS_BIN")
	}
	raw, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	gs, err := Ler(raw)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("grupos que limitam: %d", len(gs))

	// Recorte pelo catalogo real do modo, e nao pela faixa numerica: os ids
	// 22xxxx sao variantes de Arena e caem dentro da faixa.
	ids := idsDoModo(t)
	gs = Restringir(gs, ids)
	t.Logf("apos restringir ao modo: %d", len(gs))
	for _, g := range gs {
		t.Logf("  %-26s max=%d  %d itens  %v", g.ID, g.Maximo, len(g.Itens), g.Itens)
	}
}

func idsDoModo(t *testing.T) []int32 {
	raw, err := os.ReadFile(os.Getenv("LOLBUILDER_ITEMS_JSON"))
	if err != nil {
		t.Fatal(err)
	}
	var its []struct {
		ID      int32 `json:"id"`
		InStore bool  `json:"inStore"`
	}
	if err := json.Unmarshal(raw, &its); err != nil {
		t.Fatal(err)
	}
	var out []int32
	for _, i := range its {
		if i.ID >= 1000 && i.ID <= 699999 {
			out = append(out, i.ID)
		}
	}
	t.Logf("catalogo do modo: %d itens", len(out))
	return out
}
