package main

import (
	"fmt"
	"os"

	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/sync"
)

func runSync(configPath, patchline string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}

	// O patchline da linha de comando sobrepoe o do arquivo. E assim, e nao
	// editando o config, para que uma captura retroativa nao deixe a
	// configuracao apontando para um patch antigo depois de terminar.
	if patchline != "" {
		cfg.Patchline = patchline
	}

	res, err := sync.New(cfg, os.Stdout).Run()
	if err != nil {
		return err
	}
	if res.Skipped {
		return nil
	}

	fmt.Printf("\ncaptura de %s concluida\n", res.Patch)
	if cfg.Provisional {
		fmt.Println(
			"\naviso: config.json esta marcado como provisional — coverage_minimums e\n" +
				"token_budget_max ainda nao foram medidos, e o export vai se recusar a\n" +
				"publicar ate que sejam.")
	}
	return nil
}
