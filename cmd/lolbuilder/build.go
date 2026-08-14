package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/canonical"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/snapshot"
)

func runBuild(configPath, patch string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}
	store := snapshot.NewStore(cfg.SnapshotsDir)

	if patch == "" {
		latest, err := store.LatestPatch()
		if err != nil {
			return err
		}
		if latest == "" {
			return fmt.Errorf("nenhum snapshot capturado — rode 'lolbuilder sync' primeiro")
		}
		patch = latest
	}
	if !store.Exists(patch) {
		return fmt.Errorf("snapshot do patch %s nao existe em %s", patch, cfg.SnapshotsDir)
	}
	fmt.Printf("montando modelo canonico do patch %s\n", patch)

	ds, err := canonical.NewBuilder(cfg, store.PatchDir(patch)).Build(patch)
	if err != nil {
		return err
	}

	compraveis := ds.Purchasable()
	fmt.Printf("\n  itens          %3d no catalogo do modo, %3d compraveis\n", len(ds.Items), len(compraveis))
	fmt.Printf("  runas          %3d (%d fragmentos, %d keystones)\n",
		len(ds.Runes), contar(ds, canonical.Rune.Fragmento), contar(ds, canonical.Rune.Keystone))
	fmt.Printf("  estilos        %3d\n", len(ds.RuneStyles))
	fmt.Printf("  campeoes       %3d\n", len(ds.Champions))
	fmt.Printf("  summoner       %3d\n", len(ds.SummonerSpells))

	relatarItens(ds)

	if cfg.Provisional {
		fmt.Println("\naviso: config.json esta marcado como provisional — os minimos de\n" +
			"cobertura ainda nao foram medidos, entao nada aqui esta sendo\n" +
			"comparado contra piso algum.")
	}

	destino := filepath.Join("build", patch)
	if err := os.MkdirAll(destino, 0o755); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(ds, "", "  ")
	if err != nil {
		return err
	}
	arquivo := filepath.Join(destino, "canonical.json")
	if err := os.WriteFile(arquivo, append(raw, '\n'), 0o644); err != nil {
		return err
	}
	fmt.Printf("\nmodelo gravado em %s (%d bytes)\n", arquivo, len(raw)+1)
	return nil
}

func contar(ds *canonical.Dataset, pred func(canonical.Rune) bool) int {
	n := 0
	for _, r := range ds.Runes {
		if pred(r) {
			n++
		}
	}
	return n
}

// relatarItens imprime a cobertura da leitura do bloco de atributos.
//
// A taxa e sobre os itens compraveis, que sao os que o dataset publica. Item
// fora da loja nao entra no arquivo, entao uma linha ilegivel nele nao degrada
// nada que o consumidor va ler.
func relatarItens(ds *canonical.Dataset) {
	c := ds.Coverage.Itens
	fmt.Printf("\n  cobertura do parser de stats de item:\n")
	fmt.Printf("    %d/%d itens compraveis com bloco lidos por inteiro (%.1f%%)\n",
		c.LidosPorInteiro, c.ComBloco, c.Taxa())
	fmt.Printf("    %d/%d linhas de stat lidas\n", c.LinhasLidas, c.Linhas)

	if c.SemBloco > 0 {
		fmt.Printf("    %d item(ns) compravel(is) que a fonte nao descreve com bloco de atributos:\n", c.SemBloco)
		for _, nome := range c.SemBlocoNome {
			fmt.Printf("      %s\n", nome)
		}
	}
	if n := len(c.NaoLidas); n > 0 {
		fmt.Printf("    %d linha(s) nao reconhecida(s):\n", n)
		imprimirLinhas(c.NaoLidas)
	}

	// Segundo numero: o catalogo do modo inteiro, e nao so o publicado.
	//
	// Ele nao entra na taxa — item fora da loja nao chega ao consumidor. Sai
	// aqui porque e o aviso ANTECIPADO de forma nova na fonte: forca adaptativa
	// e <ornnBonus> viveram no catalogo com a taxa marcando 100%, e so seriam
	// descobertos no patch em que um item desses entrasse na loja.
	if n := len(ds.Coverage.VocabularioForaDaLoja); n > 0 {
		fmt.Printf("\n  aviso de vocabulario: %d linha(s) nao reconhecida(s) em itens do\n"+
			"  catalogo do modo que NAO estao na loja. Nao afetam o que e publicado,\n"+
			"  mas sinalizam forma que a fonte usa e o parser ainda nao le:\n", n)
		imprimirLinhas(ds.Coverage.VocabularioForaDaLoja)
	}
}

func imprimirLinhas(linhas []canon.LinhaNaoLida) {
	for i, u := range linhas {
		if i >= 10 {
			fmt.Printf("      ... e mais %d\n", len(linhas)-10)
			break
		}
		fmt.Printf("      %-28s %q\n        %s\n", u.Item, u.Linha, u.Motivo)
	}
}
