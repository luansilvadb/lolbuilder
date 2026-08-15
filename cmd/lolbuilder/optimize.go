package main

import (
	"fmt"
	"os"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/optimize"
	"github.com/luansilvadb/lolbuilder/internal/snapshot"
)

// carregarDataset monta o modelo canonico de um patch, offline.
func carregarDataset(configPath, patch string) (*config.Config, *canonical.Dataset, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, nil, err
	}
	store := snapshot.NewStore(cfg.SnapshotsDir)
	if patch == "" {
		if patch, err = store.LatestPatch(); err != nil {
			return nil, nil, err
		}
		if patch == "" {
			return nil, nil, fmt.Errorf("nenhum snapshot capturado — rode 'lolbuilder sync' primeiro")
		}
	}
	ds, err := canonical.NewBuilder(cfg, store.PatchDir(patch), "curation").Build(patch)
	return cfg, ds, err
}

func runRunes(configPath, patch, objetivo, adaptativa string, nivel int) error {
	_, ds, err := carregarDataset(configPath, patch)
	if err != nil {
		return err
	}
	if objetivo == "" {
		imprimirPaginas(ds.Computed.PaginasDeRuna)
		return nil
	}

	obj, err := optimize.ParseObjetivo(objetivo, optimize.Resolucao(adaptativa))
	if err != nil {
		return err
	}
	cat, err := canonical.CatalogoDeRunas(ds, "curation")
	if err != nil {
		return err
	}
	p, err := optimize.MelhorPagina(cat, obj, nivel)
	if err != nil {
		return err
	}
	imprimirPaginas([]optimize.Pagina{p})
	return nil
}

func runBuilds(configPath, patch, objetivo, adaptativa string, ouro int) error {
	_, ds, err := carregarDataset(configPath, patch)
	if err != nil {
		return err
	}
	if objetivo == "" {
		imprimirBuilds(ds.Computed.BuildsDeItem)
		return nil
	}

	obj, err := optimize.ParseObjetivo(objetivo, optimize.Resolucao(adaptativa))
	if err != nil {
		return err
	}
	b, err := optimize.MelhorBuild(canonical.CandidatosDeItem(ds), obj, 6, int32(ouro))
	if err != nil {
		return err
	}
	imprimirBuilds([]optimize.Build{b})
	return nil
}

func imprimirPaginas(ps []optimize.Pagina) {
	for _, p := range ps {
		fmt.Fprintf(os.Stdout, "\n%s + %s  (nivel %d, forca adaptativa como %s)\n",
			p.NomePrimario, p.NomeSecundario, p.Nivel, p.Resolucao)
		fmt.Printf("  keystone     %s\n", p.Keystone.Nome)
		for _, e := range p.Menores {
			fmt.Printf("  primaria     %s\n", e.Nome)
		}
		for _, e := range p.Secundarias {
			fmt.Printf("  secundaria   %s\n", e.Nome)
		}
		for _, e := range p.Fragmentos {
			fmt.Printf("  fragmento    %s\n", e.Nome)
		}
		fmt.Printf("  total        %v\n", p.Total)
	}
}

func imprimirBuilds(bs []optimize.Build) {
	for _, b := range bs {
		fmt.Fprintf(os.Stdout, "\n%s\n", b.Rotulo)
		fmt.Printf("  %d de ouro dos %d disponiveis, forca adaptativa como %s\n",
			b.Gasto, b.Orcamento, b.Resolucao)
		for _, it := range b.Itens {
			fmt.Printf("  %6d  %-32s %5d\n", it.ID, it.Nome, it.Custo)
		}
		fmt.Printf("  total        %v\n", b.Total)
	}
}
