# League of Legends — Summoner's Rift

**Patch 16.16** · capturado em 2026-08-16T04:16:06Z · mapa Summoner's Rift (11), modo CLASSIC

Este conjunto descreve **apenas o Summoner's Rift**. ARAM, Arena e os modos
rotativos têm catálogo de itens e ajustes próprios que não estão aqui.

## Como ler

| Arquivo | Conteúdo |
|---|---|
| `01-items.md` | itens compráveis, com custo, componentes e efeito |
| `02-runes.md` | runas por estilo, com o escopo de cada uma no cálculo |
| `03-summoner-spells.md` | feitiços de invocador válidos no modo |
| `04-champions.md` | campeões, papéis e o que cada habilidade faz |
| `05-computed.md` | máximos exatos de atributo, pré-calculados |
| `06-champion-stats.md` | estatística base e efeito das habilidades por rank |

## Contagens

| Entidade | Quantidade |
|---|---:|
| Itens no catálogo do modo | 705 |
| Itens compráveis | 210 |
| Runas jogáveis | 69 |
| Estilos de runa | 5 |
| Campeões | 173 |
| Feitiços de invocador | 9 |
| Grupos de exclusividade de item | 20 |

## Proveniência

| Origem | O que vem daí |
|---|---|
| `cdragon`, plugin | itens, runas, campeões, feitiços — dado publicado pela Riot |
| `cdragon`, dump do jogo | estatística base, fórmula de habilidade e grupos de exclusividade |
| `curated` | semântica das runas e tradução dos enums de stat |
| `derived` | os máximos de `05-computed.md` |

## Convenções

- **Stats percentuais estão em pontos percentuais**: `45` significa 45%, não 0,45.
- **Valores por rank são acumulados**: rank 3 de `+2/4/6 armadura` são 6, não 2.
- **O efeito de habilidade é uma expressão**: `50 + 0.8 ability_power` significa
  50 de dano fixo mais 80% do poder de habilidade.
- **Nome em português e nome canônico em inglês** aparecem lado a lado em itens
  e runas, porque a discussão externa usa o segundo.
- Um item é **comprável** se aparece em `01-items.md`. O catálogo do modo tem
  entradas que a loja não vende.
- **Itens do mesmo grupo de exclusividade não se acumulam.** Comprar um trava
  os outros na loja. Antes de propor qualquer build, confira a coluna `grupos`
  em `01-items.md`: uma combinação com dois itens do mesmo grupo é impossível
  de executar, por melhor que os números pareçam.

## Limites conhecidos

- **Este conjunto publica valores, mas não simula combate.** Dano com build
  montada, tempo até matar, vida efetiva e mitigação por armadura não são
  calculados aqui. Os insumos dessa conta estão em `06-champion-stats.md`.
- **A extração das habilidades é parcial.** 567 de 692 habilidades tiveram a
  fórmula resolvida por inteiro, e 115 de 173 passivas. As que ficaram sem
  número estão listadas uma a uma no fim de `06-champion-stats.md`.
  **nunca publicamos zero no lugar de uma lacuna** — habilidade que não
  aparece naquela lista tem número na tabela.
- **178 habilidades saem sem alcance.** A fonte publica valor de sentinela
  para elas — 25000 e coisas maiores, que são limite interno de míssil e não
  alcance de uso. Publicar esse número seria pior que não publicar nenhum.
- **8 estatísticas opcionais ausentes**, listadas no build. Aparecem como
  ausentes, e nunca como zero.
- **Efeitos de item não estão modelados.** O texto está aqui para leitura, mas
  passiva e ativa dependem de estado de combate.
- **As pedras fundamentais não entram em cálculo nenhum.** 58 das runas
  jogáveis dependem de estado de partida, e a coluna `escopo` de `02-runes.md`
  diz o motivo de cada uma. A página de `05-computed.md` maximiza um atributo,
  e **não** a força do campeão.
- **20 grupos de exclusividade restringem a loja**, e as combinações de
  `05-computed.md` já os respeitam. Eles vêm do dump do jogo: o catálogo do
  cliente não publica essa regra em campo nenhum, e sem ela o cálculo devolvia
  o ótimo exato de um conjunto que a loja não vende.
- **A build de `05-computed.md` não é uma build boa.** Ela é o máximo de um
  atributo por ouro, ignorando efeitos — útil como piso de comparação, inútil
  como recomendação de jogo.
- **Não há dados de partidas.** Taxa de vitória, de escolha e de banimento não
  estão aqui. Não é limitação da fonte: é escolha. Tudo neste conjunto é
  calculado a partir do que a Riot publica, e não observado em partidas.

## Tamanho

| Arquivo | Tokens estimados |
|---|---:|
| `01-items.md` | 12781 |
| `02-runes.md` | 5620 |
| `03-summoner-spells.md` | 349 |
| `04-champions.md` | 79250 |
| `05-computed.md` | 4189 |
| `06-champion-stats.md` | 92492 |
| **total sem este arquivo** | **194681** |
