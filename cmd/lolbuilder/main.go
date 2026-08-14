// Comando lolbuilder: captura, monta e exporta o dataset do Summoner's Rift
// moderno.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		configPath = flag.String("config", "config.json", "arquivo de configuracao")
		patchline  = flag.String("patchline", "", "patchline do CDragon a capturar; vazio usa o do config (latest)")
	)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		usage()
		return fmt.Errorf("nenhum comando informado")
	}

	switch cmd := args[0]; cmd {
	case "sync":
		return runSync(*configPath, *patchline)
	default:
		usage()
		return fmt.Errorf("comando desconhecido: %q", cmd)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `lolbuilder — dataset do League of Legends (Summoner's Rift moderno)

uso:
  lolbuilder [flags] <comando>

comandos:
  sync    Baixa as fontes do CommunityDragon e grava um snapshot imutavel do
          patch. Aborta sem escrever se qualquer contagem vier abaixo do minimo.

flags:
`)
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, `
exemplos:
  lolbuilder sync                      # captura o patch corrente
  lolbuilder -patchline 15.20 sync     # captura retroativa de um patch antigo
`)
}
