package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/lcu"
	"github.com/luansilvadb/lolbuilder/internal/snapshot"
)

// runIngame confronta o dataset com o que o jogo esta calculando na partida.
//
// Nao escreve no dataset. Uma divergencia pode significar extracao errada ou
// leitura contaminada por runa e item, e decidir entre as duas exige olhar para
// o contexto da partida. Ajustar o dataset automaticamente transformaria a
// pagina de runas do jogador em fonte de verdade.
func runIngame(configPath, patch, arquivo string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}

	amostras, err := carregarAmostras(arquivo)
	if err != nil {
		return err
	}

	// Sem partida NAO e erro: e o estado normal fora do jogo. Com amostras ja
	// gravadas, da para comparar com o cliente fechado.
	cli := lcu.NewLive(cfg.HTTPTimeout())
	if cli.EmPartida() {
		am, err := cli.Amostrar()
		if err != nil {
			return err
		}
		modo, _ := cli.Modo()
		fmt.Printf("partida em andamento — %s nivel %d, modo %s, %.0fs de jogo\n",
			am.Campeao, am.Nivel, modo, am.TempoDeJogo)
		if len(am.Itens) > 0 {
			fmt.Printf("  %d item(ns) equipado(s): %v\n", len(am.Itens), am.Itens)
		}
		if modo != "" && modo != cfg.Mode.GameMode {
			fmt.Printf("\naviso: a partida e do modo %s e o dataset descreve %s.\n"+
				"A comparacao segue, mas divergencia aqui pode ser so diferenca de modo.\n",
				modo, cfg.Mode.GameMode)
		}

		amostras = append(amostras, converter(am))
		if err := gravarAmostras(arquivo, amostras); err != nil {
			return err
		}
		fmt.Printf("  amostra gravada em %s (%d no total)\n", arquivo, len(amostras))
	} else {
		fmt.Println("nenhuma partida em andamento — comparando com as amostras ja gravadas")
	}

	if len(amostras) < 2 {
		fmt.Println("\ncom menos de duas amostras nao da para conferir a formula.")
		fmt.Println("suba de nivel e rode de novo: o CRESCIMENTO entre niveis e o que prova")
		fmt.Println("a extracao, porque ele cancela todo bonus fixo de item e de runa.")
		return nil
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
	ds, err := canonical.NewBuilder(cfg, store.PatchDir(patch), "curation").Build(patch)
	if err != nil {
		return err
	}

	rel, err := canonical.CompararIngame(ds, amostras, cfg.IngameTolerance)
	if err != nil {
		return err
	}
	imprimirRelatorioIngame(rel, cfg.IngameTolerance)

	if rel.Diverge() {
		return errors.New("o dataset diverge do jogo — veja as linhas marcadas acima")
	}
	return nil
}

func converter(a *lcu.Amostra) canonical.AmostraIngame {
	return canonical.AmostraIngame{
		Nivel: a.Nivel, Campeao: a.Campeao, Itens: a.Itens,
		MaxHealth:       a.Stats.MaxHealth,
		Armor:           a.Stats.Armor,
		MagicResist:     a.Stats.MagicResist,
		AttackDamage:    a.Stats.AttackDamage,
		AttackSpeed:     a.Stats.AttackSpeed,
		HealthRegenRate: a.Stats.HealthRegenRate,
	}
}

func carregarAmostras(caminho string) ([]canonical.AmostraIngame, error) {
	raw, err := os.ReadFile(caminho)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("lendo %s: %w", caminho, err)
	}
	var out []canonical.AmostraIngame
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("lendo %s: %w", caminho, err)
	}
	return out, nil
}

func gravarAmostras(caminho string, amostras []canonical.AmostraIngame) error {
	sort.Slice(amostras, func(i, j int) bool { return amostras[i].Nivel < amostras[j].Nivel })
	raw, err := json.MarshalIndent(amostras, "", " ")
	if err != nil {
		return err
	}
	if dir := filepath.Dir(caminho); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	return os.WriteFile(caminho, append(raw, '\n'), 0o644)
}

func imprimirRelatorioIngame(rel *canonical.RelatorioIngame, tol float64) {
	fmt.Printf("\n%s — niveis %v\n", rel.Campeao, rel.Niveis)
	fmt.Println("\nA comparacao e do CRESCIMENTO entre niveis, e nao do valor absoluto.")
	fmt.Println("Valor absoluto e contaminado por item, fragmento de stat e runa;")
	fmt.Println("crescimento cancela todo bonus fixo e deixa so a formula do campeao.")

	if len(rel.IntervalosCegos) > 0 {
		fmt.Printf("\naviso: o(s) intervalo(s) %v nao testam a curvatura da formula.\n",
			rel.IntervalosCegos)
		fmt.Println("Neles o crescimento real coincide com o linear ingenuo, entao passariam")
		fmt.Println("mesmo com a formula errada. Amostre um par como 1 ao 6 para testar de fato.")
	}
	if rel.ItensMudaram {
		fmt.Println("\naviso: a lista de itens mudou entre leituras, e isso quebra o")
		fmt.Println("cancelamento. As linhas afetadas saem como inconclusivas.")
	}

	fmt.Printf("\n%-22s %9s %10s %10s %10s   %s\n",
		"estatistica", "niveis", "jogo", "previsto", "diferenca", "veredito")
	for _, c := range rel.Comparacoes {
		marca := " "
		if c.Veredito == canonical.VereditoDiverge {
			marca = "!"
		}
		fmt.Printf("%s %-20s %4d→%-4d %10.4f %10.4f %10.4f   %s\n",
			marca, c.Eixo, c.DeNivel, c.AteNivel, c.Jogo, c.Previsto, c.Diff, c.Veredito)
		if c.Nota != "" {
			fmt.Printf("  %-20s %s\n", "", c.Nota)
		}
	}
	fmt.Printf("\ntolerancia de %.2f: o jogo reporta em float32 e o dataset calcula em\n", tol)
	fmt.Println("float64, entao diferenca na terceira casa e representacao, nao defeito.")
}
