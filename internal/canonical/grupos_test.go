package canonical

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/config"
)

// datasetReal monta o dataset do snapshot em disco, se houver um.
//
// Pula quando nao ha snapshot: o repositorio tem um, mas quem clona sem os
// snapshots ainda consegue rodar o resto da suite.
func datasetReal(t *testing.T) *Dataset {
	t.Helper()
	cfg, err := config.Load(filepath.Join("..", "..", "config.json"))
	if err != nil {
		t.Skipf("sem config: %v", err)
	}
	dir := filepath.Join("..", "..", cfg.SnapshotsDir, "16.16")
	if _, err := os.Stat(dir); err != nil {
		t.Skip("sem snapshot de 16.16")
	}
	ds, err := NewBuilder(cfg, dir, filepath.Join("..", "..", "curation")).Build("16.16")
	if err != nil {
		t.Fatal(err)
	}
	return ds
}

// TestBuildsPublicadasSaoCompraveis e o teste do defeito que motivou os grupos.
//
// Antes deles, 5 das 24 combinacoes publicadas em 05-computed.md nao podiam ser
// compradas: a de maxima penetracao de armadura juntava QUATRO itens do grupo
// LastWhisper, e nada no conjunto dizia que eles se excluem. Um otimo que nao
// existe e pior que nenhum otimo, porque quem le nao tem como desconfiar.
func TestBuildsPublicadasSaoCompraveis(t *testing.T) {
	ds := datasetReal(t)
	if len(ds.GruposDeItem) == 0 {
		t.Fatal("o dataset saiu sem grupo de exclusividade algum")
	}

	limite := map[string]int{}
	for _, g := range ds.GruposDeItem {
		limite[g.ID] = g.Maximo
	}
	grupoDe := map[int32][]string{}
	for _, it := range ds.Items {
		grupoDe[it.ID] = it.Grupos
	}

	if len(ds.Computed.BuildsDeItem) == 0 {
		t.Fatal("nenhuma build calculada")
	}
	for _, bl := range ds.Computed.BuildsDeItem {
		uso := map[string]int{}
		for _, it := range bl.Itens {
			for _, g := range grupoDe[it.ID] {
				uso[g]++
			}
		}
		for g, n := range uso {
			if n > limite[g] {
				t.Errorf("%s: leva %d itens do grupo %s (cabem %d) — build impossivel de comprar",
					bl.Rotulo, n, g, limite[g])
			}
		}
	}
}

// TestGrupoConhecidoSobreviveuAExtracao trava um caso concreto e verificavel a
// mao: com Lembrete Mortal equipado, a loja recusa o Terminus.
//
// Sem uma ancora assim, uma extracao que degradasse para "zero grupos" passaria
// no teste acima por vacuidade — nenhuma build violaria restricao nenhuma.
func TestGrupoConhecidoSobreviveuAExtracao(t *testing.T) {
	ds := datasetReal(t)

	const lembreteMortal, terminus = 3033, 3302
	grupoDe := map[int32][]string{}
	for _, it := range ds.Items {
		grupoDe[it.ID] = it.Grupos
	}

	comum := ""
	for _, a := range grupoDe[lembreteMortal] {
		for _, b := range grupoDe[terminus] {
			if a == b {
				comum = a
			}
		}
	}
	if comum == "" {
		t.Fatalf("Lembrete Mortal (%v) e Terminus (%v) sairam sem grupo em comum — "+
			"a extracao dos grupos degradou", grupoDe[lembreteMortal], grupoDe[terminus])
	}
	for _, g := range ds.GruposDeItem {
		if g.ID == comum && g.Maximo != 1 {
			t.Errorf("o grupo %s aceita %d itens, esperado 1", comum, g.Maximo)
		}
	}
}
