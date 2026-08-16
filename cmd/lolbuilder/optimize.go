package main

import (
	"fmt"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/optimize"
	"github.com/luansilvadb/lolbuilder/internal/snapshot"
)

// carregarDataset monta o modelo canonico de um patch, offline.
func carregarDataset(configPath, patch string) (*config.Config, *canonical.Dataset, string, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, nil, "", err
	}
	store := snapshot.NewStore(cfg.SnapshotsDir)
	if patch == "" {
		if patch, err = store.LatestPatch(); err != nil {
			return nil, nil, "", err
		}
		if patch == "" {
			return nil, nil, "", fmt.Errorf("nenhum snapshot capturado — rode 'lolbuilder sync' primeiro")
		}
	}
	ds, err := canonical.NewBuilder(cfg, store.PatchDir(patch), "curation").Build(patch)
	return cfg, ds, patch, err
}

func runRunes(configPath, patch, objetivo, adaptativa string, nivel int, formato string) error {
	fmtSaida, err := parseFormatoSaida(formato)
	if err != nil {
		return err
	}
	_, ds, patch, err := carregarDataset(configPath, patch)
	if err != nil {
		return err
	}
	if objetivo == "" {
		return emitirPaginas(fmtSaida, patch, ds.Computed.PaginasDeRuna, nil)
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
	return emitirPaginas(fmtSaida, patch, []optimize.Pagina{p}, &obj)
}

func runBuilds(configPath, patch, objetivo, adaptativa string, ouro int, formato string) error {
	fmtSaida, err := parseFormatoSaida(formato)
	if err != nil {
		return err
	}
	_, ds, patch, err := carregarDataset(configPath, patch)
	if err != nil {
		return err
	}
	if objetivo == "" {
		return emitirBuilds(fmtSaida, patch, ds.Computed.BuildsDeItem, nil)
	}

	obj, err := optimize.ParseObjetivo(objetivo, optimize.Resolucao(adaptativa))
	if err != nil {
		return err
	}
	b, err := optimize.MelhorBuild(canonical.CandidatosDeItem(ds), canonical.GruposParaMochila(ds), obj, 6, int32(ouro))
	if err != nil {
		return err
	}
	return emitirBuilds(fmtSaida, patch, []optimize.Build{b}, &obj)
}
