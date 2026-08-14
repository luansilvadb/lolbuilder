# Spike 0 — viabilidade das fontes do LoL moderno

Medido em 2026-08-14 contra o patch **16.16** (`content-metadata.json` do patchline
`latest`). Nenhum código de produção foi escrito. Este documento registra o que a
fonte oferece e o que isso muda nas decisões tomadas antes de olhar para ela.

## Resumo

O porte é mais barato do que o desenho supunha em quase tudo, e mais caro em um
ponto só. Duas decisões prévias mudam: a fonte estruturada de stats de item **não
existe**, e o filtro de escopo do modo **não está onde se supunha**.

## Catálogos do plugin

Todos em `plugins/rcp-be-lol-game-data/global/{default,pt_br}/v1/`, HTTP 200 nos
dois locales.

| arquivo | entradas | campos | bytes (default / pt_br) |
|---|---:|---:|---|
| `items.json` | 868 | 19, uniformes | 683 KB / 719 KB |
| `perks.json` | 103 | 10, uniformes | 100 KB / 110 KB |
| `perkstyles.json` | 5 estilos | — | 18 KB / 18 KB |
| `champion-summary.json` | 237 | 7, uniformes | 54 KB / 54 KB |
| `summoner-spells.json` | 39 | 7, uniformes | 13 KB / 14 KB |

**Decodificação estrita é barata aqui.** Os campos são uniformes em 100% das
entradas nos quatro arquivos de lista — não há schema variável a acomodar. A lista
de campos ignorados (decisão 6) provavelmente nasce vazia nestes arquivos.

**pt_br é localização real**, não fallback para inglês: os arquivos são
consistentemente maiores. A decisão 12 (bilíngue seletivo) é viável.

### Faixas de ID separam os dois jogos

`champion-summary.json` traz 237 entradas: **173 modernas** (id 1–999), **63 Jade**
(id 60000–60999) e 1 sentinela (`id: -1`, alias `None`, sem `roles`). O recorte por
faixa de ID é limpo e é o mesmo mecanismo que o `config.json` do Classic já usa.

O mesmo vale para itens: `items.json` mistura os dois jogos — a tag `jadeUnique`
aparece 356 vezes nas descrições.

## O filtro de escopo não está no `items.json`

**`items.json` moderno não tem campo `maps`.** O recorte "comprável no Summoner's
Rift" não sai desta fonte. `inStore` é verdadeiro em 696 dos 868, e inclui ARAM,
Arena e o resto.

O filtro está no dump do mapa, em `game/data/maps/shipping/map11/map11.bin.json`
(10,1 MB). Ele contém 20 registros `GameModeItemList`, e o modo resolve assim:

```
Maps/Shipping/Map11/Modes/CLASSIC.itemLists[0] -> {413f2f94} -> 193 itens
```

193 itens, contra 696 `inStore`. O filtro não é cosmético.

Notas: `map453` existe no mesmo diretório e é o mapa do Jade, o que confirma o
`map_id` do Classic. O mesmo arquivo traz `MapCharacterList` e `GameModeConstants`,
que podem servir para outras validações.

`summoner-spells.json`, ao contrário dos itens, **traz o filtro embutido**: o campo
`gameModes` recorta `CLASSIC` em 9 feitiços — Cleanse, Exhaust, Flash, Ghost, Heal,
Smite, Teleport, Ignite, Barrier.

## Stats de item: não há fonte estruturada

Buscado e não encontrado: `game/data/items/`, `game/items/` e `game/global/` não
contêm definição de item. A string `ItemData` **não ocorre nenhuma vez** nos 10 MB
do `map11.bin.json`. Os únicos tipos com "Item" no nome são de loja e telemetria
(`GameModeItemList`, `ItemShopGameModeData`, `HudItemShopQuickBuyData`).

Conclusão: **o CDragon não publica stat de item de forma estruturada.** A decisão 8
cai no fallback declarado — leitura do bloco de texto do plugin.

Mas o motivo que justificava a fonte dupla desapareceu. O bloco `<stats>` moderno é
regular e **já inclui o que faltava no Classic**:

```
Berserker's Greaves: <attention> 30%</attention> Attack Speed<br>
                     <attention> 45</attention> Move Speed
Sorcerer's Shoes:    <attention> 12</attention> Magic Penetration<br>
                     <attention> 45</attention> Move Speed
```

A velocidade de movimento das botas está **dentro** do bloco de stats, não escondida
no texto do efeito. A limitação mais citada do dataset do Classic não se reproduz
aqui. 811 dos 868 itens têm bloco `<stats>`, com gramática uniforme
(`<attention> VALOR</attention> Nome do Stat<br>`).

## Runas: a estrutura da página é publicada

No Classic, as regras da página (30 slots, 30 pontos) eram constantes curadas à mão
em `curation/rules.json`, vigiadas por `-watch-rules`. No moderno, `perkstyles.json`
**publica a estrutura**:

- 5 estilos, com `allowedSubStyles` explícito;
- 1 slot `kKeyStone` (3 opções em Resolve/Domination/Inspiration, 4 em
  Precision/Sorcery);
- 3 slots `kMixedRegularSplashable`, 3 opções cada;
- 3 slots `kStatMod`, 3 opções cada — os fragmentos;
- `subStyleBonus` e `defaultStatModsPerSubStyle` por combinação.

O espaço de busca da página fica na ordem de 10⁶, **enumerável por exaustão exata**
sem esforço. A garantia "exatas, não heurísticas" sobrevive sem DP.

`perks.json` ainda traz **`majorChangePatchVersion`** por runa — um sinal nativo de
alteração, mais confiável que o diff de `longDesc` que a decisão 5 previa.

## O dump de campeão porta quase intacto

173 arquivos baixados, **12 MB no total**, todos com `CharacterRecord` presente.
4.648 `SpellObject`, média de 26,9 por campeão, máximo 84 (Yasuo). 3.522 blocos
`GameCalculation`.

Os nomes de campo do `CharacterRecord` moderno são **idênticos** aos que o
`internal/gamedata` do Classic já mapeia — `baseHPModifiable`, `hpPerLevelModifiable`,
`baseDamageModifiable`, `mrPerLevel` — com a mesma forma
`{"baseValue": N, "__type": "ModifiableFloat"}`. Garen confere: 690 de vida base,
+98 por nível, 69 de dano, 32 de resistência mágica +1,55 por nível.

O `mSpell` moderno tem os mesmos campos que a struct `Spell` do Classic mapeia:
`DataValues`, `mSpellCalculations`, `castRange`, `castRangeValues`, `Cooldown`,
`manaValues`, e as herdadas `cooldownTime` e `mana`.

### Parcelas de cálculo: 18 de 19 já tratadas

Varredura completa dos 173 campeões:

| parcela | ocorrências | tratada pelo Classic |
|---|---:|---|
| `GameCalculation` | 3522 | sim |
| `NamedDataValueCalculationPart` | 2069 | sim |
| `StatByCoefficientCalculationPart` | 812 | sim |
| `StatByNamedDataValueCalculationPart` | 712 | sim |
| `NumberCalculationPart` | 533 | sim |
| `GameCalculationModified` | 280 | sim |
| `ByCharLevelBreakpointsCalculationPart` | 228 | sim |
| `ProductOfSubPartsCalculationPart` | 199 | sim |
| `SumOfSubPartsCalculationPart` | 188 | sim |
| `ByCharLevelInterpolationCalculationPart` | 167 | sim |
| `StatBySubPartCalculationPart` | 63 | sim |
| `EffectValueCalculationPart` | 36 | sim |
| `BuffCounterByNamedDataValueCalculationPart` | 34 | sim |
| `BuffCounterByCoefficientCalculationPart` | 23 | sim |
| `CooldownMultiplierCalculationPart` | 18 | sim |
| `GameCalculationConditional` | 14 | sim |
| `ByCharLevelFormulaCalculationPart` | 12 | sim |
| `AbilityResourceByCoefficientCalculationPart` | 10 | sim |
| **`ClampSubPartsCalculationPart`** | **6** | **não** |

Uma única parcela nova, em 6 ocorrências. O avaliador do Classic cobre o resto.

## Escala

| grandeza | Classic 16.16 | moderno 16.16 | fator |
|---|---:|---:|---:|
| campeões | 63 | 173 | 2,7× |
| itens no recorte do modo | 150 | 193 | 1,3× |
| runas / maestrias | 53 + 56 | 103 | ~0,9× |
| feitiços publicados | 0 (16 capturados) | 9 | — |
| dump de campeão | — | 12 MB | — |

Projeção do conjunto exportado: o `06-champion-stats` do Classic dominava os 53.070
tokens medidos, e cresce ~2,7×. Estimativa de **130k a 160k tokens** no total — a
ser medida de verdade no M5, como a decisão 4 prevê.

## Aberto para medição posterior

- **Qual série de recarga e custo o moderno usa.** O Classic descobriu por medição
  em partida que o modo Jade usa a série redefinida (`Cooldown`, `manaValues`) e não
  a herdada (`cooldownTime`, `mana`). No jogo base as duas também coexistem, e não
  há razão para assumir que a resposta é a mesma. Decide-se no M3 pelo alinhamento
  de rank, e confirma-se no M6 pelo `ingame`.
- **Cobertura de extração de habilidade e passiva.** Só medível no M3; é o insumo do
  bootstrap de `coverage_minimums` (decisão 7).
- **Semântica de `ClampSubPartsCalculationPart`.** 6 ocorrências, a resolver no M3.
