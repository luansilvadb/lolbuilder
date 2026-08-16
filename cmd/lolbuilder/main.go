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
		patch      = flag.String("patch", "", "patch a montar; vazio usa o snapshot mais recente")
		objetivo   = flag.String("objective", "", "stat a maximizar; vazio imprime a tabela pre-calculada")
		adaptativa = flag.String("adaptive", "ad", "como resolver forca adaptativa: ad ou ap")
		nivel      = flag.Int("level", 18, "nivel do campeao, para as runas que escalam")
		ouro       = flag.Int("gold", 20000, "orcamento em ouro, para a build de itens")
		saida      = flag.String("out", "_data", "diretorio de destino do export")
		amostras   = flag.String("samples", "ingame-samples.json", "arquivo de amostras do comando ingame")
		formato    = flag.String("format", "text", "formato da saida de runes/builds: text ou json")
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
	case "build":
		return runBuild(*configPath, *patch)
	case "export":
		return runExport(*configPath, *patch, *saida)
	case "ingame":
		return runIngame(*configPath, *patch, *amostras)
	case "runes":
		return runRunes(*configPath, *patch, *objetivo, *adaptativa, *nivel, *formato)
	case "builds":
		return runBuilds(*configPath, *patch, *objetivo, *adaptativa, *ouro, *formato)
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
  build   Monta o modelo canonico a partir de um snapshot e imprime a cobertura
          de leitura. Offline. Grava build/<patch>/canonical.json.
  export  Gera os arquivos do Project e o changelog do patch. Offline. Recusa
          publicar enquanto config.json estiver marcado como provisional.
  ingame  Confronta o dataset com o que o jogo esta calculando numa partida em
          andamento. Acumula amostras entre execucoes; com duas ou mais em
          niveis diferentes, confere a formula pelo crescimento.
  runes   Pagina de runas de valor maximo para um objetivo.
  builds  Maximo de um stat por ouro em 6 slots. NAO e build otima: o calculo
          ignora passiva e ativa de item.

flags:
`)
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, `
exemplos:
  lolbuilder sync                      # captura o patch corrente
  lolbuilder -patchline 15.20 sync     # captura retroativa de um patch antigo
  lolbuilder build                     # monta o snapshot mais recente
  lolbuilder -patch 16.16 build        # monta um snapshot especifico
  lolbuilder -objective armor runes    # pagina de runas com mais armadura
  lolbuilder -objective ability_power -adaptive ap runes
  lolbuilder -objective armor -gold 10000 builds
  lolbuilder -objective armor -format json runes
  lolbuilder ingame                    # com uma partida aberta, grava e compara
`)
}
