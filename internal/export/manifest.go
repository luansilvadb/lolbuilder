package export

import (
	"fmt"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
)

// renderManifest monta o 00-MANIFEST.md.
//
// Ele NAO e instrucao: e conhecimento, e vai junto dos outros arquivos. E ele
// que impede o modelo de completar lacuna com o que sabe de outra fonte, porque
// declara o que o conjunto cobre e, principalmente, o que ele nao cobre.
func renderManifest(in Input, arquivos []File) string {
	ds := in.Dataset
	var b strings.Builder

	fmt.Fprintf(&b, "# League of Legends — Summoner's Rift\n\n")
	fmt.Fprintf(&b, "**Patch %s** · capturado em %s · mapa Summoner's Rift (11), modo CLASSIC\n\n",
		ds.Patch, in.CapturadoEm)
	b.WriteString("Este conjunto descreve **apenas o Summoner's Rift**. ARAM, Arena e os modos\n")
	b.WriteString("rotativos têm catálogo de itens e ajustes próprios que não estão aqui.\n\n")

	b.WriteString("## Como ler\n\n")
	b.WriteString("| Arquivo | Conteúdo |\n|---|---|\n")
	b.WriteString("| `01-items.md` | itens compráveis, com custo, componentes e efeito |\n")
	b.WriteString("| `02-runes.md` | runas por estilo, com o escopo de cada uma no cálculo |\n")
	b.WriteString("| `03-summoner-spells.md` | feitiços de invocador válidos no modo |\n")
	b.WriteString("| `04-champions.md` | campeões, papéis e o que cada habilidade faz |\n")
	b.WriteString("| `05-computed.md` | máximos exatos de atributo, pré-calculados |\n")
	b.WriteString("| `06-champion-stats.md` | estatística base e efeito das habilidades por rank |\n\n")

	b.WriteString("## Contagens\n\n| Entidade | Quantidade |\n|---|---:|\n")
	fmt.Fprintf(&b, "| Itens no catálogo do modo | %d |\n", len(ds.Items))
	fmt.Fprintf(&b, "| Itens compráveis | %d |\n", len(ds.Purchasable()))
	fmt.Fprintf(&b, "| Runas jogáveis | %d |\n", contarJogaveis(ds))
	fmt.Fprintf(&b, "| Estilos de runa | %d |\n", len(ds.RuneStyles))
	fmt.Fprintf(&b, "| Campeões | %d |\n", len(ds.Champions))
	fmt.Fprintf(&b, "| Feitiços de invocador | %d |\n\n", len(ds.SummonerSpells))

	b.WriteString("## Proveniência\n\n| Origem | O que vem daí |\n|---|---|\n")
	b.WriteString("| `cdragon`, plugin | itens, runas, campeões, feitiços — dado publicado pela Riot |\n")
	b.WriteString("| `cdragon`, dump do jogo | estatística base e fórmula de habilidade |\n")
	b.WriteString("| `curated` | semântica das runas e tradução dos enums de stat |\n")
	b.WriteString("| `derived` | os máximos de `05-computed.md` |\n\n")

	b.WriteString("## Convenções\n\n")
	b.WriteString("- **Stats percentuais estão em pontos percentuais**: `45` significa 45%, não 0,45.\n")
	b.WriteString("- **Valores por rank são acumulados**: rank 3 de `+2/4/6 armadura` são 6, não 2.\n")
	b.WriteString("- **O efeito de habilidade é uma expressão**: `50 + 0.8 ability_power` significa\n")
	b.WriteString("  50 de dano fixo mais 80% do poder de habilidade.\n")
	b.WriteString("- **Nome em português e nome canônico em inglês** aparecem lado a lado em itens\n")
	b.WriteString("  e runas, porque a discussão externa usa o segundo.\n")
	b.WriteString("- Um item é **comprável** se aparece em `01-items.md`. O catálogo do modo tem\n")
	b.WriteString("  entradas que a loja não vende.\n\n")

	b.WriteString("## Limites conhecidos\n\n")
	b.WriteString("- **Este conjunto publica valores, mas não simula combate.** Dano com build\n")
	b.WriteString("  montada, tempo até matar, vida efetiva e mitigação por armadura não são\n")
	b.WriteString("  calculados aqui. Os insumos dessa conta estão em `06-champion-stats.md`.\n")
	escreverCobertura(&b, ds)
	b.WriteString("- **Efeitos de item não estão modelados.** O texto está aqui para leitura, mas\n")
	b.WriteString("  passiva e ativa dependem de estado de combate.\n")
	b.WriteString("- **As pedras fundamentais não entram em cálculo nenhum.** 58 das runas\n")
	b.WriteString("  jogáveis dependem de estado de partida, e a coluna `escopo` de `02-runes.md`\n")
	b.WriteString("  diz o motivo de cada uma. A página de `05-computed.md` maximiza um atributo,\n")
	b.WriteString("  e **não** a força do campeão.\n")
	b.WriteString("- **A build de `05-computed.md` não é uma build boa.** Ela é o máximo de um\n")
	b.WriteString("  atributo por ouro, ignorando efeitos — útil como piso de comparação, inútil\n")
	b.WriteString("  como recomendação de jogo.\n")
	b.WriteString("- **Não há dados de partidas.** Taxa de vitória, de escolha e de banimento não\n")
	b.WriteString("  estão aqui. Não é limitação da fonte: é escolha. Tudo neste conjunto é\n")
	b.WriteString("  calculado a partir do que a Riot publica, e não observado em partidas.\n\n")

	b.WriteString("## Tamanho\n\n| Arquivo | Tokens estimados |\n|---|---:|\n")
	total := 0
	for _, f := range arquivos {
		fmt.Fprintf(&b, "| `%s` | %d |\n", f.Name, f.Tokens)
		total += f.Tokens
	}
	fmt.Fprintf(&b, "| **total sem este arquivo** | **%d** |\n", total)
	return b.String()
}

// escreverCobertura declara o que a extracao nao alcancou, com numero.
func escreverCobertura(b *strings.Builder, ds *canonical.Dataset) {
	c := ds.Coverage.Campeoes
	fmt.Fprintf(b, "- **A extração das habilidades é parcial.** %d de %d habilidades tiveram a\n",
		c.Habilidades.Resolvidas, c.Habilidades.Total)
	fmt.Fprintf(b, "  fórmula resolvida por inteiro, e %d de %d passivas. As que ficaram sem\n",
		c.Passivas.Resolvidas, c.Passivas.Total)
	b.WriteString("  número estão listadas uma a uma no fim de `06-champion-stats.md`.\n")
	b.WriteString("  **nunca publicamos zero no lugar de uma lacuna** — habilidade que não\n")
	b.WriteString("  aparece naquela lista tem número na tabela.\n")

	if n := c.AlcancesDescartados; n > 0 {
		fmt.Fprintf(b, "- **%d habilidades saem sem alcance.** A fonte publica valor de sentinela\n", n)
		b.WriteString("  para elas — 25000 e coisas maiores, que são limite interno de míssil e não\n")
		b.WriteString("  alcance de uso. Publicar esse número seria pior que não publicar nenhum.\n")
	}
	if n := len(c.LacunasDeStat); n > 0 {
		fmt.Fprintf(b, "- **%d estatísticas opcionais ausentes**, listadas no build. Aparecem como\n", n)
		b.WriteString("  ausentes, e nunca como zero.\n")
	}
}

func contarJogaveis(ds *canonical.Dataset) int {
	n := 0
	for _, r := range ds.Runes {
		if r.TipoSlot != "" {
			n++
		}
	}
	return n
}
