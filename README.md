# lolbuilder

Dataset do **League of Legends no Summoner's Rift** extraído do CommunityDragon,
entregue no formato que um Project do [claude.ai](https://claude.ai) consome.

Porte do **Classicorone** — que faz o mesmo para o modo Jade — com a mesma
stack e a mesma regra de negócio. As divergências entre os dois estão em
`openspec/decisoes.md`, cada uma com o motivo.

Go, apenas biblioteca padrão. Nenhuma dependência.

## Gerar o `_data/`

```sh
go run ./cmd/lolbuilder sync     # baixa o patch atual → snapshots/<patch>/
go run ./cmd/lolbuilder export   # gera _data/ + changelogs/<patch>.md
```

Dois comandos. Depois suba os arquivos de `_data/` no Project, substituindo os
anteriores — os nomes são fixos entre patches, não sobra órfão.

O `sync` é idempotente: patch já capturado, ele não faz nada. O `export` lê de
`snapshots/`, nunca da rede.

## Comandos

| Comando | O que faz |
|---|---|
| `sync` | Baixa os catálogos do CommunityDragon nos dois locales, o dump do mapa e o dump de dados do jogo de cada campeão, e grava um snapshot imutável. Aborta sem escrever se qualquer contagem vier abaixo do mínimo. Não precisa do cliente aberto. |
| `build` | Monta o modelo canônico e imprime a cobertura de extração. Offline. Grava `build/<patch>/canonical.json`. |
| `export` | Gera os arquivos do Project e o changelog. Offline. Recusa publicar se a cobertura cair abaixo dos mínimos ou o conjunto passar do teto de tokens. |
| `runes` | Página de runas de valor máximo para um objetivo. |
| `builds` | Máximo de um atributo por ouro em 6 slots. **Não** é build ótima. |
| `ingame` | Confronta o dataset com o que o jogo está calculando numa partida. |

```sh
go run ./cmd/lolbuilder -objective armor runes
go run ./cmd/lolbuilder -objective ability_power -adaptive ap runes
go run ./cmd/lolbuilder -objective attack_damage -gold 10000 builds
```

Sem `-objective`, os dois imprimem a tabela pré-calculada — a mesma que vai para
`05-computed.md`.

## Fontes

Três, todas do CommunityDragon.

- `plugins/rcp-be-lol-game-data/.../v1/` — itens, runas, campeões, feitiços.
  Baixado em dois locales: o canônico é a espinha estrutural, e o `pt_br` traz o
  texto publicado. Decodificação estrita: campo novo vira erro.
- `game/data/characters/<alias>/` — o dump de dados do jogo, um arquivo por
  campeão. É a única fonte com estatística base e fórmula de habilidade; o
  plugin publica esses campos zerados. Não passa por decodificação estrita,
  porque as chaves variam por campeão e incluem identificadores opacos. O que
  vigia essa fonte é a cobertura de extração que o `build` imprime.
- `game/data/maps/shipping/map11/` — o dump do mapa. É a **única** fonte que diz
  quais itens a loja do modo vende: `items.json` não tem campo de mapa, e
  `inStore` é verdadeiro também para os itens de ARAM e Arena.

O `build` também cruza a indexação de rank entre o dump e o arquivo do plugin.
As séries de recarga e custo estão populadas nos dois lados, então servem para
conferir que o deslocamento não mudou — sem isso, um erro de índice publicaria
todo dano deslocado um rank, em silêncio.

## O que os cálculos garantem

Os dois otimizadores são **exatos**, não heurísticos, e verificados contra busca
exaustiva em teste.

**Página de runas:** com objetivo linear, cada slot é independente dado o par de
estilos, então o máximo da soma é a soma dos máximos — 20 pares em vez de 1,5
milhão de páginas. Slot em que nenhuma opção contribui sai como `indiferente`,
e não com o nome da runa que venceu o desempate.

**Build de item:** mochila 0/1 por programação dinâmica sobre slots e ouro, com
a regra de botas únicas. **Não é uma build boa** — o cálculo ignora passiva e
ativa de item, e o rótulo publicado diz isso.

## O que não está modelado

Estatística base e fórmula de habilidade **estão** no conjunto, em
`06-champion-stats.md`. O que segue fora é o modelo de combate.

- **Dano com build montada, tempo até matar e vida efetiva** — dependem de
  decisões que a fonte não fornece. O conjunto publica os insumos; a conta é de
  quem consome.
- **Parte das fórmulas de habilidade** — cerca de um quinto das habilidades e um
  terço das passivas não resolvem por inteiro. O `build` reporta as taxas sobre
  o total, e o que ficou sem número está listado no fim de
  `06-champion-stats.md`, nunca publicado como zero.
- **As pedras fundamentais.** 58 das 69 runas jogáveis dependem de estado de
  partida, e a coluna `escopo` de `02-runes.md` diz o motivo de cada uma.
- **Efeitos de item** — passiva e ativa dependem de estado de combate.
- **Dados de partidas.** Não há taxa de vitória, de escolha ou de banimento. Não
  é limitação da fonte: é escolha. Tudo aqui é calculado, não observado.

## Conferir contra o jogo

```sh
go run ./cmd/lolbuilder ingame
```

Nenhuma fonte publica o que o jogo faz em tempo de execução: o dump declara base
e crescimento, e se a fórmula que aplicamos sobre eles estiver errada, nada no
dado denuncia. Foi assim que a resistência mágica ficou anos publicada como fixa
no dataset do modo Jade.

O comando acumula uma amostra por execução. Com duas ou mais em níveis
diferentes, ele compara o **crescimento** — e não o valor absoluto, que é
contaminado por item, fragmento de stat e runa. O crescimento cancela todo bônus
fixo e deixa só a fórmula do campeão.

Sem partida em andamento não é erro: ele compara as amostras já gravadas.

## Layout

```
  config.json          parametros do modo, minimos e tetos, com o historico
  curation/            classificado a mao: runas e enums de stat
  snapshots/<patch>/   bytes originais do CDragon, imutaveis
  build/<patch>/       modelo canonico em JSON, regeravel
  _data/               o que vai para o Project
  changelogs/          o que mudou em cada patch — para voce ler
  openspec/            as decisoes do projeto e o motivo de cada uma
  docs/                o relatorio do spike inicial
```

## Rode o sync a cada patch

O CommunityDragon versiona os patchlines do jogo base, então uma captura perdida
pode ser refeita depois com `-patchline 15.20 sync`. Isso muda a urgência, não a
disciplina: o changelog só existe entre dois snapshots, e sem o anterior não há
com o que comparar.

## Atribuição

Ferramenta não oficial, sem vínculo com a Riot Games. Dados do
[CommunityDragon](https://communitydragon.org). League of Legends © Riot Games, Inc.
