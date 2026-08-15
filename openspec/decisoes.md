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

**"Comprável" ficou para o M2.** A união inclui buffs de torre, placeholders e
itens removidos, e nenhum campo isolado da fonte separa isso. O `filter` reporta
o que a fonte diz; a semântica é do modelo canônico, onde mora o
`Purchasable()`. Resolvido pela decisão 18, abaixo — que usa `inStore`, mas
**em conjunção** com a referência do modo, e não sozinho.

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

**O fragmento de stat ocupa mais de uma linha.** `Adaptive Force` está nas
linhas 4 e 5, `Health Scaling` nas 5 e 6. Os outros cinco fragmentos ocupam uma
só, e keystones e menores são sempre únicos. O modelo guarda o conjunto de
linhas, não a primeira: um otimizador que lesse só a primeira nunca colocaria
Força Adaptativa no slot flexível, e publicaria página subótima com cara de
exata — que é o pior erro possível num sistema que promete exatidão.

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

## Revisões durante a revisão do M2

**Medir vocabulário só sobre o publicado esconde forma nova.** A cobertura de
item mede os compráveis, que é o que o consumidor lê — e com isso marcava 100%
enquanto duas formas passavam despercebidas no resto do catálogo: o rótulo
`Adaptive Force`, em 12 itens, e a marcação `<ornnBonus>`, que os itens
aprimorados pelo Ornn usam no lugar de `<attention>`. Nenhuma está em item
comprável hoje, então o alarme só tocaria no patch em que uma delas chegasse à
loja, e aí a correção seria urgência em vez de manutenção.

Duas respostas. As duas formas passam a ser reconhecidas — força adaptativa
entra no vocabulário, o que a decisão 17 já exigia de qualquer jeito. E o build
passa a imprimir um **segundo número**: linhas não lidas em itens do catálogo
que não estão na loja. Ele não entra na taxa, porque não afeta o publicado; ele
existe para avisar antes.

**Hwei tinha 12 habilidades descartadas.** O `spellbookOverride` traz três grupos
de quatro sub-habilidades, com dano e recarga próprios. O campo estava mapeado
em `raw.go` justamente por ser dado de habilidade de verdade, e não estava sendo
publicado — o dataset diria que Hwei tem quatro habilidades.

**As habilidades são pareadas por posição, mas o `spellKey` é conferido.** Sem a
conferência, uma reordenação num dos locales publicaria a descrição do W sob o
nome do Q, em silêncio e em todos os campeões de uma vez. A guarda do `sync`
confere a contagem, não a ordem.

## Achados do M3

**A série redefinida é a que o jogo usa — agora medido, não assumido.** Contra o
arquivo do plugin, que é fonte independente: `Cooldown` concorda em **680 de 680**
recargas e `manaValues` em **568 de 568** custos, ambos 100%. As herdadas
(`cooldownTime`, `mana`) ficam em 98,5% e 98,8%, divergindo em 10 e 7
habilidades. Mesma resposta do modo Jade, por evidência própria.

**Três campeões publicam mais de um `CharacterRecord`.** Braum e Milio têm
`Root`, `URF` e `SLIME`; Cassiopeia tem `Root` e `SLIME`. O registro do
Summoner's Rift é o `Root`. Selecionar por `__type`, como o original fazia,
pegaria qualquer um deles — e num patch qualquer o dataset passaria a publicar
estatística de URF como se fosse do Rift, sem sintoma nenhum.

**`ClampSubPartsCalculationPart` soma as parcelas e prende o resultado entre piso
e teto.** Só resolve quando a soma não escala com estatística: prender uma
expressão simbólica exigiria saber o valor da estatística, que é justamente o que
o dataset não fixa, e publicar sem o limite daria um número que o jogo nunca
produz.

**A curadoria de enums é deliberadamente incompleta.** O moderno usa 18 valores
distintos de `mStat` contra 8 no modo Jade. Doze estão curados com evidência em
`curation/statenum.json`; seis ficaram de fora por falta de evidência suficiente
— 20 ocorrências de cerca de 1590. Enum fora da tabela deixa a parcela sem
resolver e entra na cobertura; nunca vira palpite publicado.

## Revisões durante a revisão do M3

**O nível de referência era mudo.** 407 parcelas dependem do nível do campeão,
e o dataset publicava só o valor no 18 sem dizer isso. Quem lê "115" toma por
constante o que no nível 1 vale bem menos. Agora o efeito declara
`nivel_de_referencia` — mas **só quando ele muda o resultado**, senão seria ruído
na maioria, que não depende do nível. A detecção é por medição: avalia no 1 e no
18 e compara. Varrer a árvore exigiria enumerar os tipos de parcela que dependem
do nível, e essa lista sairia do ar no dia em que a Riot criasse mais um. São 288
efeitos que declaram.

**As 12 sub-habilidades do Hwei tinham texto e nenhum número — e eram invisíveis
na cobertura.** Elas não estão no array de slots do `CharacterRecord`, então
`Spells()` não as alcançava. São `SpellObject` próprios, resolvíveis pelo
`ObjectName` (`qq` → `HweiQQ`). Agora saem com recarga, custo e fórmula, e têm
eixo próprio na cobertura: 9 de 12 resolvidas.

**O alcance sentinela não é publicado.** 178 habilidades traziam valores como
25000 no `castRange` — limite interno de míssil, não alcance que o jogador
enxerga. O limiar de 10000 se justifica por lacuna observada: a banda legítima
termina em 7500, nos ultimates globais, e nada existe entre 7500 e 10000.
Verifiquei também se `castRangeValues` traria o valor real: **zero casos**. A
alternativa não existe, então a escolha era entre publicar um número errado e
não publicar nenhum.

**Senna e Thresh não eram lacuna — eram zero, e eu tinha errado.** Medido: os
campos `*PerLevel` **nunca** aparecem presentes valendo zero; a fonte apenas os
omite, e a omissão é como ela diz "não cresce". Senna não ganha dano por nível e
Thresh não ganha armadura por nível, os dois por ganharem isso de almas.
Reportar essas omissões como lacuna publicava uma dúvida que não existe. Os
campos de crescimento saíram da lista de obrigatórios; os campos **base**
continuam nela, e `baseStaticHPRegen` segue como lacuna declarada em 8 campeões,
porque ali a omissão não é interpretável.

## Achados do M4

**A decisão 2 se confirmou por inteiro: só 11 das 69 runas somam atributo.** Os
7 fragmentos de stat mais 4 menores — Cosmic Insight, Revitalize, Celerity e
Transcendence. As outras 58 dependem de estado de partida. Isso não é defeito da
curadoria: é a natureza do sistema moderno, onde a keystone define comportamento
e não atributo. É também exatamente por isso que a decisão 13 existe.

**Duas runas são parcialmente somáveis, e a parte que fica de fora está
registrada.** Revitalize dá 5% de cura e escudo incondicional mais 10%
condicional; Celerity dá 1% de deslocamento fixo mais um multiplicador de 7%
sobre efeitos de movimentação. Nos dois casos o `note` diz o que não entrou.

**Transcendence exigiu um mecanismo novo: degraus por nível.** Ela dá aceleração
de habilidade no nível 5 e de novo no 8 — nem valor fixo nem crescimento linear.
Sem os degraus ela cairia em `out_of_scope`, e o dataset perderia um atributo
real e incondicional só por não saber expressá-lo.

**Duas grandezas entraram no vocabulário sem que item algum as publique.**
Aceleração de feitiço de invocador e de item só existem em Cosmic Insight.
Deixá-las de fora obrigaria a classificar como `out_of_scope` uma runa que dá
atributo fixo — o zero silencioso que a curadoria existe para impedir.

**O atalho do otimizador de runas é o ótimo, e há teste que prova.** Com objetivo
linear, cada slot é independente dado o par de estilos, então o máximo da soma é
a soma dos máximos — 20 pares em vez de 1,5 milhão de páginas. O teste compara
contra busca exaustiva sobre um catálogo de forma real.

**O dado real revelou um defeito que os testes não pegaram.** A mochila de itens
dava o valor certo e reconstruía a combinação errada: `Presságio de Randuin` e
`Couraça do Defunto` saíam **duas vezes** na mesma build de armadura. O ponteiro
de volta apontava para o estado da camada em que o item entrou, e essa célula
podia ter sido sobrescrita por um item posterior. A célula passou a carregar a
combinação inteira — cabe, porque ela tem no máximo 6 itens — e a classe de erro
sumiu. Há teste de regressão que reproduz.

## Revisões durante a revisão do M4

**A etiqueta de botas da fonte não é confiável.** `Gunmetal Greaves` evolui de
Grevas do Berserker e a fonte a categoriza como `AttackSpeed`, `LifeSteal` e
`NonbootsMovement` — **sem** `Boots`. Confiar na categoria direta deixava o
otimizador livre para montar uma build com dois pares de calçado, impossível no
jogo e com cara de ótima. A regra passou a ser transitiva pela árvore de
componentes: aprimoramento de botas é botas. Ganha exatamente 1 item, sem falso
positivo, e sobrevive à fonte esquecer a etiqueta em mais um item no próximo
patch.

**A curadoria não chegava ao modelo publicado.** `kind`, `reason` e `note` eram
carregados, usados pelo otimizador e descartados. Sem eles o arquivo de runas
listaria as 69 sem explicar por que 58 não aparecem no pré-cálculo — o mesmo que
não declarar limite nenhum. Agora cada runa carrega escopo, motivo e ressalva.

**Objetivo em força adaptativa é rejeitado com explicação.** Ela sempre resolve
para dano de ataque ou poder de habilidade antes de valer alguma coisa, então
`Valor()` a convertia e o objetivo pontuava zero — em silêncio. O erro agora diz
o que pedir no lugar.

## Achados do M5

**A projeção do Spike 0 errou.** Estimei 130k–160k tokens; a primeira medição
real deu **217.403**. Duas correções de qualidade cortaram 23.531 sem tirar
informação nenhuma, e o conjunto fechou em **194.504**, com teto em 240.000.

**A fonte publica ruído de float32 como se fosse precisão.** O jogo calcula em
float32 e o dump serializa a representação float64 disso, então 4,2 chega como
`4.199999809265137`. Publicar os 17 dígitos era duas coisas ruins ao mesmo
tempo: mentir sobre a precisão do dado, e gastar cinco tokens onde um basta —
num arquivo com dezenas de milhares desses números. Formatar na precisão do
float32 cortou **22.253 tokens** sozinho.

**Os marcadores do texto de habilidade agora resolvem.** Eram 3.247 ocorrências
de `@TotalDamage@` e afins — o cliente as preenche em tempo de execução a partir
das mesmas séries que o dump publica. Resolvê-las transforma "causa
@TotalDamage@ de Dano Físico" em "causa 30/60/90/120/150 + 1.5 attack_damage de
Dano Físico". Era a pendência que o M2 registrou.

**Marcador sem valor vira `(?)`, e não some.** Apagar quebrava a frase — "Garen
fica com de Armadura" lê pior que a versão com o marcador. `(?)` mantém a frase
de pé e declara a lacuna, que é a mesma regra do resto do conjunto.

**O resolvedor precisa de TODAS as séries, inclusive as consumidas.** O modelo
filtra as séries que alguma fórmula já consome, para não publicar o mesmo número
duas vezes ao lado do efeito resolvido. Mas o texto referencia essas séries, e
sem elas 63 marcadores a mais viravam `(?)`. O campo com todas existe no modelo
e **não** é serializado.

## Revisões durante a revisão do M5

**60% das escolhas publicadas nas páginas eram desempates arbitrários.** Eram
**49 de 81**. As nove páginas saíam com "pedra fundamental: Pressione o Ataque"
— não por mérito, mas porque nenhuma keystone soma atributo, todas empatam em
zero e o desempate é pelo menor id. É a mesma classe de erro que publicar zero
no lugar de uma lacuna: apresentar um artefato do algoritmo como se fosse
resposta. E contradizia o aviso do próprio arquivo, que diz duas linhas acima
que a keystone não entra no cálculo.

Slot em que nenhuma opção contribui passa a sair como `indiferente`. Isso torna
a tabela mais informativa, e não menos: o leitor vê que a página de armadura
determina 1 escolha e deixa 8 livres, que é a verdade útil.

**A tabela não mostrava de onde vinha o total.** Ela listava keystone e
fragmentos; a página de cura e escudo tem os dois indiferentes, e os 5% vinham
de uma runa secundária que a tabela não exibia — um número sem origem visível.
As colunas viraram "runas que contribuem", com o que cada uma dá, e "slots
livres".

**A coluna `motivo` de `02-runes.md` misturava duas semânticas.** "Fora do
cálculo porque" explica a exclusão de uma runa `out_of_scope`; "ressalva"
registra a parte do efeito que ficou de fora numa runa que soma **em parte**.
Sob rótulo único, o leitor não tinha como saber qual dos dois estava lendo.

**O `export` regrava o `canonical.json`.** Ele monta o modelo de qualquer jeito,
e sem isso os dois artefatos descrevem o mesmo patch e podem discordar — o do
disco fica do último `build`.

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
