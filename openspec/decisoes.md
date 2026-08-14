# Decisões

Este projeto é um porte do **Classicorone** — dataset do modo Jade (League of
Legends Classic) — para o **Summoner's Rift moderno**. Mesma stack, mesma regra
de negócio. O que está aqui são as decisões que **divergem** do original, ou que
o original não precisou tomar, com o motivo de cada uma.

Sem este registro, daqui a alguns patches alguém vai "consertar" uma dessas
divergências achando que é defeito.

## Tomadas antes de olhar para a fonte

| # | Decisão | Motivo |
|---:|---|---|
| 1 | **Escopo: Summoner's Rift moderno, só ele.** ARAM, Arena e Swiftplay ficam fora. | Preserva a regra que o `00-MANIFEST` do original protege: o conjunto descreve um modo e não deve ser completado com o que existe fora dele. ARAM/Arena têm modificadores por mapa que exigiriam um segundo eixo de modelagem inteiro. |
| 2 | **Runes Reforged portado 1:1** no vocabulário de escopo curado (`somável` / `por nível` / `multiplicador` / `fora`), com enumeração exaustiva **exata**. | Tradução fiel da regra de negócio. Preserva a propriedade mais valiosa do projeto: dizer explicitamente o que não calcula, em vez de inventar número para efeito condicional. |
| 3 | **O cliente de jogo sai do `sync`.** Sobrevive só no comando `ingame` (Live Client Data), opcional e nunca bloqueante. | No modo Jade o cliente era a única fonte que dizia quais runas estavam à venda. No moderno não existe loja de runas — o acoplamento perdeu a razão, e mantê-lo seria carregar o custo sem o motivo. |
| 4 | **Mesma estrutura de arquivos, `token_budget_max` recalibrado por medição.** Particionamento de `06` por faixa fixa fica escrito como contingência. | "Completude do dado prevalece sobre economia de tamanho." Reduzir escopo para caber violaria a regra de nunca omitir valor que a fonte publica. Particionar por faixa fixa, e nunca por "encher até caber", mantém os nomes estáveis entre patches. |
| 5 | **Curadoria de runa obrigatória por id**, inclusive `out_of_scope` com `reason`. Runa sem entrada aborta o build; runa alterada vira stale e aborta até revisão. | Sob curadoria parcial, uma runa nova que soma stat entra como zero silencioso — o modo de falha que o projeto inteiro existe para tornar impossível. |
| 6 | **Decodificação estrita em tudo + escape declarado** para subárvores que o projeto não consome. | Estrito puro quebraria o build a cada patch de cosmético, e build que quebra por bobagem treina o operador a contornar o alerta. O escape é o tipo `model.Ignored`, com o motivo em comentário no ponto de uso. |
| 7 | **Bootstrap de cobertura em duas etapas.** Enquanto `provisional` for verdadeiro, `coverage_minimums` e `token_budget_max` podem ser zero e o **export se recusa a publicar**. | Um mínimo só mede degradação futura se nascer de um valor observado que alguém aprovou. Herdar os números do Classicorone seria calibrar com dado de outro jogo; defini-los por expectativa arriscaria congelar um defeito. |
| 8 | **Stats de item cruzados de duas fontes**, com divergência publicada como incerteza. | *Revogada pelo Spike 0 — ver abaixo.* |
| 9 | **Dados de partida não entram.** "Calculado, não observado" passa a ser escolha declarada, não limitação da fonte. | Dado observado quebra quatro propriedades de uma vez: `sync` determinístico, snapshot imutável, build offline e zero dependências. E winrate é justamente o que o consumidor já tem por outros caminhos, enquanto fórmula resolvida do dump é o que ele não tem. |
| 10 | **`latest` por padrão, captura retroativa sob demanda** via `-patchline`. Procedência no `capture.json`. | O CDragon versiona patchlines no jogo base, então "cada patch sem sync é um buraco permanente" deixa de valer. Semear automaticamente arriscaria alimentar a calibração dos mínimos com um patch degradado. |
| 11 | **Sete arquivos**: `00-MANIFEST`, `01-items`, `02-runes`, `03-summoner-spells`, `04-champions`, `05-computed`, `06-champion-stats`. | Feitiço de invocador decide build tanto quanto runa, e o original os capturava, validava e descartava sem publicar. Numeração com buraco seria dívida cosmética permanente num conjunto que se sobe à mão. |
| 12 | **Moldura e efeitos em pt_br; nome canônico em inglês adicional** em itens e runas. | O usuário pergunta em português, mas todo o corpus externo do modelo está em inglês. A ponte entre os dois nomes evita a tradução de volta, que é onde o modelo erra. Duplicar textos de efeito seria caro e sem retorno. |
| 13 | **Otimizador exato de build de item** (mochila de 6 slots sob orçamento), rotulado **"máximo de `<stat>` por ouro, ignorando efeitos"** — nunca "build ótima". | Preenche o `05-computed`, que no moderno ficaria quase vazio só com runas. O rótulo não é cosmético: build ótima por stat linear **não é build boa**, e um modelo lendo "ótima" repassaria bobagem com a autoridade do dataset. |
| 14 | **`lolbuilder`**, módulo `github.com/luansilvadb/lolbuilder`, Go 1.26, apenas biblioteca padrão. | — |
| 15 | **Regime de testes completo + este registro.** | As falhas aqui são silenciosas: um deslocamento de rank publica todo dano errado sem quebrar nada, e um `out_of_scope` errado some com uma runa do otimizador sem aviso. |
| 16 | **Spike 0 → M1 captura → M2 catálogo → M3 dados do jogo → M4 curadoria e otimização → M5 export → M6 fidelidade**, com revisão a cada marco. | O risco aqui não é de arquitetura — o original já provou o desenho. É de fonte, e o spike ataca isso mais barato que uma fatia vertical. |
| 17 | **Força adaptativa é stat de primeira classe com resolução declarada** (`-adaptive ad\|ap`); o pré-computado publica as duas leituras. | Ela é determinística *dada uma declaração*, o que a distingue de efeito condicional de verdade. Tratá-la como `out_of_scope` esvaziaria o otimizador de runas, já que os fragmentos são quase todo o conteúdo somável da página. |

## Revisões após o Spike 0

**Decisão 8 revogada.** Não existe fonte estruturada de stats de item: `game/data/items/`
não existe, e a string `ItemData` não ocorre nenhuma vez nos 10 MB do dump do mapa.
Fica a leitura do texto do plugin — mas o motivo que justificava o cruzamento
desapareceu junto: o bloco `<stats>` moderno **já inclui a velocidade de movimento
das botas**, que era a limitação mais citada do dataset original. Sem cruzamento,
sem incerteza de item.

**Decisão 5 fica mais barata.** `perkstyles.json` publica a estrutura da página
(slots, keystones, estilos secundários permitidos, fragmentos), então não há
constante de estrutura para curar à mão como havia com as maestrias. E
`perks.json` traz `majorChangePatchVersion` por runa — gatilho de revisão nativo,
melhor que comparar `longDesc`, que muda por reescrita de texto sem mudar número.

**Decisão 1 ganha uma fonte.** `items.json` moderno **não tem campo `maps`**. O
recorte "comprável no Summoner's Rift" sai de `map11.bin.json`, em
`Maps/Shipping/Map11/Modes/CLASSIC.itemLists`. Isso acrescenta um arquivo de
9,7 MB por captura.

## Revisões durante o M1

**A união das listas de loja, não a maior delas.** As 11 listas que o modo
declara são grupos semânticos — catálogo principal, trinkets, itens de Doran,
companheiros de selva, itens de suporte, botas de nível 3. Ficar só com a maior
(193 itens) deixaria de fora **todas as botas avançadas e todos os itens
iniciais**. A união dá 266.

**"Comprável" ainda não está decidido, e é problema do M2.** A união inclui
buffs de torre, placeholders e itens removidos. `inStore` não serve de
discriminador: ele é falso em itens compráveis de verdade (Bússola Rúnica, Pedra
de Visão Agitadora) e verdadeiro em itens de outros modos. O `filter` reporta o
que a fonte diz; a semântica é do modelo canônico, onde mora o `Purchasable()`.

**Órfão é ausência do catálogo completo, não da faixa do modo.** A loja do
Summoner's Rift referencia `771500` ("Penetrating Bullets"), que existe em
`items.json` mas mora na faixa de id do modo Jade. Conferir contra o catálogo já
recortado transformava um fato da fonte em erro fatal e abortava a captura. Id
fora da faixa vira nota no log; ausência do catálogo inteiro continua abortando.

**`SummonerSpell.ID` é `int64`.** A fonte usa 4294967295 como sentinela de "sem
id", em três entradas chamadas "Primal Smite" com `gameModes` vazio.

## Decisão 18, tomada no M2

**Comprável = referenciado pelas listas de loja do modo **E** `inStore=true`.**
São 210 itens no 16.16, de um catálogo de 705.

Os dois filtros fazem trabalho real, e nenhum sozinho serve. **333** itens têm
`inStore` verdadeiro e o modo não referencia — são de ARAM e Arena
(`Guardian's Amulet`, `Quest: Support`). **55** o modo referencia e a fonte não
põe na loja — buffs de torre, `Structure Bounty`, `Gangplank Placeholder`.

Isso reproduz o precedente do original, que dizia explicitamente que `InStore`
**não** era a fonte de verdade porque havia itens com `InStore` verdadeiro fora
da loja do modo. Aqui são exatamente esses 333.

Candidatos descartados, com o motivo:

- **`priceTotal > 0`** exclui os trinkets (`Stealth Ward`, `Farsight Alteration`,
  `Oracle Lens`), que são gratuitos e são decisão de jogo real.
- **`active`** não significa "existe no jogo". `Infinity Edge`, `Long Sword` e
  `B. F. Sword` têm `active=false`. O campo marca item com efeito **ativo**,
  usável — potions, elixires, `Stirring Wardstone`.

O critério exclui 26 itens que o modo referencia com preço maior que zero. A
maioria é transformação automática, que não se compra (`Seraph's Embrace`,
`Muramana`, `Runic Compass`, `Bounty of Worlds`) ou item que saiu do Rift
(`Prowler's Claw`, `Night Harvester`, `Chemtech Putrifier`, `Opportunity`). Se
algum deles for comprável hoje, a correção é uma exceção curada com motivo em
`curation/items.json`, não afrouxar a regra.

## Achados do M2, para o M4

**`subStyleBonus` não aponta para runas.** Os ids que ele referencia são
**títulos cosméticos de página** — "The Incontestable Spellslinger" para
Precision + Sorcery, "The Brazen Perfect", "The Eternal Champion". O otimizador
não pode tratá-los como runas que concedem stat.

**Só 69 das 103 entradas de `perks.json` são runas jogáveis.** A conta fecha:
17 keystones + 45 menores (5 estilos × 3 linhas × 3 opções) + 7 fragmentos = 69.
As outras 34 são runas removidas do jogo (`Predator`, `Kleptomancy`,
`Zombie Ward`, `Eyeball Collection`, `Ingenious Hunter`, `Celestial Body`,
`Iron Skin`, `Mirror Shell`, `Chrysalis`), um registro chamado `Template`, e os
títulos de página acima. A curadoria da decisão 5 se aplica às 69 — exigir
entrada para as 34 seria curar entidade que o jogo não oferece.

**As descrições de habilidade trazem marcadores não resolvidos.** O
`dynamicDescription` do plugin publica `@TotalDamage@`, `@Cost@`, `@Cooldown@`
literalmente. Resolvê-los contra o dump de dados do jogo é trabalho do M3 — e é
justamente o que dá valor ao `04-champions.md`, já que sem isso o texto descreve
a habilidade sem dizer quanto ela causa.

## Aberto

- Valor de `coverage_minimums` — medido no M3.
- Valor de `token_budget_max` — medido no M5.
- Qual série de recarga e custo o jogo base usa, a redefinida (`Cooldown`,
  `manaValues`) ou a herdada (`cooldownTime`, `mana`). O original descobriu por
  medição em partida que o modo Jade usa a redefinida; não há razão para assumir
  que a resposta é a mesma aqui. Decide-se no M3, confirma-se no M6.
- Semântica de `ClampSubPartsCalculationPart`, a única parcela de cálculo dos
  173 campeões que o avaliador do original não trata. 6 ocorrências.
- Nada pendente do M2.
