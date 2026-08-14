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

	ds, err := canonical.NewBuilder(cfg, store.PatchDir(patch), "curation").Build(patch)
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
	relatarCampeoes(ds)

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

// relatarCampeoes imprime a cobertura de extracao do dump de dados do jogo.
//
// As taxas sao sobre o TOTAL de entidades, e nao sobre o subconjunto que tem
// formula. E por isso que elas nao chegam a 100% mesmo sem nada ter falhado: as
// entidades que a fonte nao descreve por numero permanecem no denominador, para
// que um aumento delas apareca aqui em vez de sumir.
func relatarCampeoes(ds *canonical.Dataset) {
	c := ds.Coverage.Campeoes
	fmt.Printf("\n  cobertura do dump de dados do jogo:\n")
	fmt.Printf("    estatisticas   %3d/%-3d campeoes completos (%.1f%%)\n",
		c.CampeoesComStats, c.CampeoesTotal, ratio(c.CampeoesComStats, c.CampeoesTotal))
	linhaDeEntidade("habilidades", c.Habilidades)
	linhaDeEntidade("passivas", c.Passivas)

	for _, a := range ds.Coverage.Alinhamento {
		fmt.Printf("    alinhamento de rank em %-9s deslocamento %d, %.1f%% de concordancia com o plugin (%d habilidades)\n",
			a.Series, a.Best, a.Agreement(), a.Compared)
		if n := len(a.Divergent); n > 0 {
			fmt.Printf("      %d divergente(s): %v\n", n+a.Omitidas, a.Divergent)
		}
	}

	if n := len(c.LacunasDeStat); n > 0 {
		fmt.Printf("    %d lacuna(s) em estatistica opcional (publicadas como ausentes, nunca como zero):\n", n)
		for i, g := range c.LacunasDeStat {
			if i >= 10 {
				fmt.Printf("      ... e mais %d\n", n-10)
				break
			}
			fmt.Printf("      %s\n", g)
		}
	}
}

func linhaDeEntidade(nome string, e canonical.CoberturaDeEntidade) {
	fmt.Printf("    %-14s %3d/%-3d resolvidas (%.1f%%), %d parciais, %d nao resolvidas, %d sem formula (%.1f%%)\n",
		nome, e.Resolvidas, e.Total, ratio(e.Resolvidas, e.Total),
		e.Parciais, e.NaoResolvidas, e.SemFormula, ratio(e.SemFormula, e.Total))
}

func ratio(parte, total int) float64 {
	if total == 0 {
		return 0
	}
	return 100 * float64(parte) / float64(total)
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
