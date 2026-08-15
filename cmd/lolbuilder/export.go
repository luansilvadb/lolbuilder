package main

import (
	"fmt"
	"os"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/export"
	"github.com/luansilvadb/lolbuilder/internal/snapshot"
)

func runExport(configPath, patch, out string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}
	store := snapshot.NewStore(cfg.SnapshotsDir)

	if patch == "" {
		if patch, err = store.LatestPatch(); err != nil {
			return err
		}
		if patch == "" {
			return fmt.Errorf("nenhum snapshot capturado — rode 'lolbuilder sync' primeiro")
		}
	}
	if !store.Exists(patch) {
		return fmt.Errorf("snapshot do patch %s nao existe em %s", patch, cfg.SnapshotsDir)
	}

	// A TRAVA. Ver a decisao 7 em openspec/decisoes.md.
	//
	// Um dataset cuja qualidade nunca passou por porteira nenhuma e o pior risco
	// do projeto: o consumidor final e um modelo de linguagem, que nao tem como
	// saber que falta dado.
	if cfg.Provisional {
		return fmt.Errorf(
			"config.json esta marcado como provisional: coverage_minimums e "+
				"token_budget_max ainda nao foram medidos, e publicar sem eles entregaria "+
				"um conjunto cuja qualidade nunca passou por porteira nenhuma.\n\n"+
				"Rode 'lolbuilder build' para medir a cobertura, grave os minimos com folga "+
				"declarada em %s, e troque provisional para false", configPath)
	}

	fmt.Printf("montando o patch %s\n", patch)
	ds, err := canonical.NewBuilder(cfg, store.PatchDir(patch), "curation").Build(patch)
	if err != nil {
		return err
	}
	if err := conferirCobertura(cfg, ds); err != nil {
		return err
	}

	anterior := carregarAnterior(cfg, store, patch)
	cap, err := store.LoadCapture(patch)
	if err != nil {
		return err
	}
	capturadoEm := ""
	if cap != nil {
		capturadoEm = cap.CapturedAt
	}

	res, err := export.Generate(export.Input{
		Dataset: ds, Anterior: anterior, CapturadoEm: capturadoEm,
	})
	if err != nil {
		return err
	}

	fmt.Println()
	for _, f := range res.Files {
		fmt.Printf("  %-24s %7d tokens\n", f.Name, f.Tokens)
	}
	fmt.Printf("  %-24s %7d tokens (teto %d)\n", "TOTAL", res.Total, cfg.TokenBudgetMax)

	if res.Total > cfg.TokenBudgetMax {
		return fmt.Errorf(
			"o conjunto estimou %d tokens e o teto e %d — nada foi escrito.\n\n"+
				"O teto NAO deve ser ajustado para caber no resultado corrente: um teto que "+
				"persegue o tamanho deixa de medir alguma coisa. Investigue o crescimento "+
				"antes de mexer em token_budget_max",
			res.Total, cfg.TokenBudgetMax)
	}

	if err := res.Write(out); err != nil {
		return err
	}
	if err := res.WriteChangelog("changelogs"); err != nil {
		return err
	}
	fmt.Printf("\n%d arquivos em %s/, changelog em changelogs/%s\n",
		len(res.Files), out, res.Changelog.Name)

	if estranhos := res.CheckOutputDir(out); len(estranhos) > 0 {
		fmt.Fprintf(os.Stderr,
			"\naviso: %d arquivo(s) em %s/ que o export nao criou: %v\n"+
				"Nada foi removido. Se forem para o Project junto com os corretos, o modelo "+
				"vai le-los como conhecimento.\n", len(estranhos), out, estranhos)
	}
	return nil
}

// carregarAnterior monta o patch anterior para o changelog. Falhar aqui nao
// derruba o export: sem o anterior, o changelog so diz que e a captura inicial.
func carregarAnterior(cfg *config.Config, store *snapshot.Store, patch string) *canonical.Dataset {
	anterior, err := store.PatchAnteriorA(patch)
	if err != nil || anterior == "" {
		return nil
	}
	ds, err := canonical.NewBuilder(cfg, store.PatchDir(anterior), "curation").Build(anterior)
	if err != nil {
		fmt.Fprintf(os.Stderr, "aviso: o patch anterior (%s) nao montou, o changelog sai sem comparacao: %v\n",
			anterior, err)
		return nil
	}
	return ds
}

// conferirCobertura aborta se a extracao caiu abaixo do minimo declarado.
func conferirCobertura(cfg *config.Config, ds *canonical.Dataset) error {
	c := ds.Coverage.Campeoes
	m := cfg.CoverageMinimums

	checks := []struct {
		nome  string
		valor float64
		min   int
		teto  bool
	}{
		{"champion_stats", ratio(c.CampeoesComStats, c.CampeoesTotal), m.ChampionStats, false},
		{"abilities", ratio(c.Habilidades.Resolvidas, c.Habilidades.Total), m.Abilities, false},
		{"passives", ratio(c.Passivas.Resolvidas, c.Passivas.Total), m.Passives, false},
		{"abilities_without_formula_max", ratio(c.Habilidades.SemFormula, c.Habilidades.Total),
			m.AbilitiesWithoutFormulaMax, true},
		{"passives_without_formula_max", ratio(c.Passivas.SemFormula, c.Passivas.Total),
			m.PassivesWithoutFormulaMax, true},
	}
	for _, ch := range checks {
		if ch.teto && ch.valor > float64(ch.min) {
			return fmt.Errorf(
				"%s ficou em %.1f%%, acima do teto de %d%% — a fracao de entidades sem "+
					"formula alguma cresceu, e a taxa de resolucao sozinha nao detectaria isso",
				ch.nome, ch.valor, ch.min)
		}
		if !ch.teto && ch.valor < float64(ch.min) {
			return fmt.Errorf(
				"%s caiu para %.1f%%, abaixo do minimo de %d%% — nada foi escrito",
				ch.nome, ch.valor, ch.min)
		}
	}
	for _, a := range ds.Coverage.Alinhamento {
		if err := a.Err(); err != nil {
			return err
		}
	}
	return nil
}
