# Runas — patch 16.16

A página tem 5 estilos. O primário dá 1 pedra fundamental mais 3 runas
menores, uma por linha; o secundário dá 2 runas menores de linhas diferentes;
e 3 fragmentos de stat completam a página.

A coluna `escopo` diz se a runa entra no cálculo de `05-computed.md`:
`sum` soma atributo fixo, `sum_per_level` soma um valor que cresce com o
nível, e `out_of_scope` depende de estado de partida. **A maioria é
`out_of_scope`, e isso não é lacuna do conjunto**: no sistema moderno a pedra
fundamental define comportamento, e não atributo. O motivo de cada uma está
na coluna `motivo`.

## Precisão (Precision)

| id | runa | nome canônico | tipo | escopo | stats | efeito | motivo |
|---|---|---|---|---|---|---|---|
| 8005 | Pressione o Ataque | Press the Attack | pedra fundamental | out_of_scope | — | Atingir um Campeão inimigo 3 vezes consecutivas causa dano adicional e amplifica seu dano até sair de combate contra Campeões. | depende de estado de combate: Atingir um Campeão inimigo 3 vezes consecutivas causa dano adicional e amplifica seu dano até sair de combate |
| 8008 | Ritmo Fatal | Lethal Tempo | pedra fundamental | out_of_scope | — | Atacar um Campeão inimigo concede a você Velocidade de Ataque, até 6 acúmulos. Com o máximo de acúmulos, causa Dano Adaptativo adicional ao contato. | depende de estado de combate: Atacar um Campeão inimigo concede a você Velocidade de Ataque, até 6 acúmulos. Com o máximo de acúmulos, causa |
| 8010 | Conquistador | Conqueror | pedra fundamental | out_of_scope | — | Ganha acúmulos de Força Adaptativa ao atacar Campeões inimigos. Após alcançar 12 acúmulos, cura uma parte do dano causado a Campeões. | depende de estado de combate: Ganha acúmulos de Força Adaptativa ao atacar Campeões inimigos. Após alcançar 12 acúmulos, cura uma parte do d |
| 8021 | Agilidade nos Pés | Fleet Footwork | pedra fundamental | out_of_scope | — | Atacar e se movimentar geram acúmulos de Energia. Com 100 acúmulos, seu próximo ataque cura você e concede um aumento de VdM. | depende de estado de combate: Atacar e se movimentar geram acúmulos de Energia. Com 100 acúmulos, seu próximo ataque cura você e concede um |
| 8009 | Presença de Espírito | Presence of Mind | menor | out_of_scope | — | Restaura uma pequena quantidade de Mana ou Energia ao causar dano a um Campeão inimigo. Eliminações restauram Mana ou Energia. | depende de estado de combate: Restaura uma pequena quantidade de Mana ou Energia ao causar dano a um Campeão inimigo. Eliminações restauram |
| 8014 | Golpe de Misericórdia | Coup de Grace | menor | out_of_scope | — | Causa mais dano a Campeões inimigos com pouca Vida. | depende de estado de combate: Causa mais dano a Campeões inimigos com pouca Vida. |
| 8017 | Dilacerar | Cut Down | menor | out_of_scope | — | Causa mais dano a Campeões inimigos com muita Vida. | depende de estado de combate: Causa mais dano a Campeões inimigos com muita Vida. |
| 8299 | Até a Morte | Last Stand | menor | out_of_scope | — | Cause mais dano a Campeões enquanto estiver com pouca Vida. | depende de estado de combate: Cause mais dano a Campeões enquanto estiver com pouca Vida. |
| 9101 | Absorção Vital | Absorb Life | menor | out_of_scope | — | Abater um alvo cura você. | depende de estado de combate: Abater um alvo cura você. |
| 9103 | Lenda: Linhagem | Legend: Bloodline | menor | out_of_scope | — | Abates de inimigos concedem Roubo de Vida permanente até um limite. Quando o limite é atingido, aumentam sua Vida máxima. Fraco no começo mas fica mais forte do que as outras runas Lenda no fim de jogo. | depende de abates acumulados na partida: Abates de inimigos concedem Roubo de Vida permanente até um limite. Quando o limite é atingido, aumentam sua V |
| 9104 | Lenda: Espontaneidade | Legend: Alacrity | menor | out_of_scope | — | Abates em inimigos concedem Velocidade de Ataque permanente. | depende de abates acumulados na partida: Abates em inimigos concedem Velocidade de Ataque permanente. |
| 9105 | Lenda: Aceleração | Legend: Haste | menor | out_of_scope | — | Eliminar inimigos concede Aceleração permanente de habilidades básicas. | depende de abates acumulados na partida: Eliminar inimigos concede Aceleração permanente de habilidades básicas. |
| 9111 | Triunfo | Triumph | menor | out_of_scope | — | Abates restauram 5% da sua Vida perdida e concedem 20 de ouro adicional. | depende de abates acumulados na partida: Abates restauram 5% da sua Vida perdida e concedem 20 de ouro adicional. |

## Dominação (Domination)

| id | runa | nome canônico | tipo | escopo | stats | efeito | motivo |
|---|---|---|---|---|---|---|---|
| 8112 | Eletrocutar | Electrocute | pedra fundamental | out_of_scope | — | Acertar um Campeão com 3 ataques ou habilidades separadas em 3s causa Dano Adaptativo adicional. | depende de estado de combate: Acertar um Campeão com 3 ataques ou habilidades separadas em 3s causa Dano Adaptativo adicional. |
| 8128 | Colheita Sombria | Dark Harvest | pedra fundamental | out_of_scope | — | Causar dano a um Campeão que está com pouca Vida inflige Dano Adaptativo e colhe a alma de sua vítima. | depende de estado de combate: Causar dano a um Campeão que está com pouca Vida inflige Dano Adaptativo e colhe a alma de sua vítima. |
| 9923 | Chuva de Lâminas | Hail of Blades | pedra fundamental | out_of_scope | — | Recebe uma grande quantidade de Velocidade de Ataque e Dano Verdadeiro adicional pelos primeiros 3 ataques contra Campeões inimigos. | depende de estado de combate: Recebe uma grande quantidade de Velocidade de Ataque e Dano Verdadeiro adicional pelos primeiros 3 ataques con |
| 8105 | Caça Incansável | Relentless Hunter | menor | out_of_scope | — | Abates únicos concedem VdM permanente fora de combate. | depende de abates acumulados na partida: Abates únicos concedem VdM permanente fora de combate. |
| 8106 | Caça Suprema | Ultimate Hunter | menor | out_of_scope | — | Abates únicos concedem Redução de Tempo de Recarga permanente para a ultimate. | depende de abates acumulados na partida: Abates únicos concedem Redução de Tempo de Recarga permanente para a ultimate. |
| 8126 | Golpe Desleal | Cheap Shot | menor | out_of_scope | — | Causa Dano Verdadeiro adicional a Campeões inimigos com mobilidade ou ações debilitadas. | depende de estado de combate: Causa Dano Verdadeiro adicional a Campeões inimigos com mobilidade ou ações debilitadas. |
| 8135 | Caçador de Tesouros | Treasure Hunter | menor | out_of_scope | — | Eliminações únicas concedem ouro adicional na primeira vez que são coletadas. | depende de abates acumulados na partida: Eliminações únicas concedem ouro adicional na primeira vez que são coletadas. |
| 8137 | Sexto Sentido | Sixth Sense | menor | out_of_scope | — | Com um longo Tempo de Recarga, detecta automaticamente uma sentinela oculta próxima, rastreando-a para a equipe. Nível 11: também revela a sentinela por 10s. | depende de estado de combate: Com um longo Tempo de Recarga, detecta automaticamente uma sentinela oculta próxima, rastreando-a para a equip |
| 8139 | Gosto de Sangue | Taste of Blood | menor | out_of_scope | — | Cura ao causar dano a um Campeão inimigo. | depende de estado de combate: Cura ao causar dano a um Campeão inimigo. |
| 8140 | Lembranças Aterrorizantes | Grisly Mementos | menor | out_of_scope | — | Colete 1 Lembranças ao eliminar Campeões, até um máximo de 18. Recebe 6 de Aceleração de Amuleto para cada Lembrança coletada. | depende de abates acumulados na partida: Colete 1 Lembranças ao eliminar Campeões, até um máximo de 18. Recebe 6 de Aceleração de Amuleto para cada Lem |
| 8141 | Sentinela Profunda | Deep Ward | menor | out_of_scope | — | Suas sentinelas na selva inimiga são Profundas, recebendo Vida e Duração. Nível 9: As sentinelas no rio também são Profundas. | depende de estado de combate: Suas sentinelas na selva inimiga são Profundas, recebendo Vida e Duração. Nível 9: As sentinelas no rio também |
| 8143 | Impacto Repentino | Sudden Impact | menor | out_of_scope | — | Depois de usar um avanço, salto, teleporte ou ao sair da furtividade, seus ataques básicos e habilidades de dano causam Dano Verdadeiro adicional a Campeões inimigos. | depende de estado de combate: Depois de usar um avanço, salto, teleporte ou ao sair da furtividade, seus ataques básicos e habilidades de da |

## Feitiçaria (Sorcery)

| id | runa | nome canônico | tipo | escopo | stats | efeito | motivo |
|---|---|---|---|---|---|---|---|
| 8214 | Invocar Aery | Summon Aery | pedra fundamental | out_of_scope | — | Seus ataques e habilidades enviam Aery até um alvo, causando dano a inimigos ou protegendo aliados com um escudo. | depende de estado de combate: Seus ataques e habilidades enviam Aery até um alvo, causando dano a inimigos ou protegendo aliados com um escu |
| 8229 | Cometa Arcano | Arcane Comet | pedra fundamental | out_of_scope | — | Ao causar dano a um Campeão com uma habilidade, um cometa é lançado onde ele estiver. | depende de estado de combate: Ao causar dano a um Campeão com uma habilidade, um cometa é lançado onde ele estiver. |
| 8230 | Avanço da Tempestade | Stormraider's Surge | pedra fundamental | out_of_scope | — | Causar dano equivalente a 25% da Vida máxima de um Campeão concede um impulso de VdM e Resistência a Lentidão. | depende de estado de combate: Causar dano equivalente a 25% da Vida máxima de um Campeão concede um impulso de VdM e Resistência a Lentidão. |
| 8992 | Toque Ígneo | Deathfire Touch | pedra fundamental | out_of_scope | — | Causar dano a um Campeão com uma habilidade o queima ao longo do tempo. | depende de estado de combate: Causar dano a um Campeão com uma habilidade o queima ao longo do tempo. |
| 8210 | Transcendência | Transcendence | menor | sum | 10 acel | Recebe efeitos adicionais ao atingir os seguintes níveis: Nível 5: +5 de Aceleração de Habilidade Nível 8: +5 de Aceleração de Habilidade Nível 11: ao eliminar um Campeão inimigo, reduz o Tempo de Recarga restante das habilidades básicas em 20%. | O nivel 11 concede reducao de recarga da ultimate ao abater, que depende de estado de partida e fica fora. |
| 8224 | Arcanista do Axioma | Axiom Arcanist | menor | out_of_scope | — | Sua ultimate causa mais dano e recebe mais cura e Escudo. Eliminar um Campeão inimigo reduz o Tempo de Recarga atual da sua ultimate. | depende de estado de combate: Sua ultimate causa mais dano e recebe mais cura e Escudo. Eliminar um Campeão inimigo reduz o Tempo de Recarga |
| 8226 | Faixa de Fluxo de Mana | Manaflow Band | menor | out_of_scope | — | Atingir um Campeão inimigo com uma habilidade aumenta permanentemente seu Mana máximo em 25, até o total de 250 de Mana. Após atingir 250 de Mana adicional, 1% da seu Mana perdido é restaurado a cada 5s | depende de estado de combate: Atingir um Campeão inimigo com uma habilidade aumenta permanentemente seu Mana máximo em 25, até o total de 25 |
| 8232 | Caminhar Sobre as Águas | Waterwalking | menor | out_of_scope | — | Concede VdM e PdH ou DdA Adaptativos no rio. | depende de estado de combate: Concede VdM e PdH ou DdA Adaptativos no rio. |
| 8233 | Foco Absoluto | Absolute Focus | menor | out_of_scope | — | Acima de 70% de Vida, recebe Dano Adaptativo adicional. | depende de estado de combate: Acima de 70% de Vida, recebe Dano Adaptativo adicional. |
| 8234 | Celeridade | Celerity | menor | sum | 1 %mov | Todos os efeitos de Velocidade de Movimento adicional são 7% mais eficazes em você, além de conceder 1% de Velocidade de Movimento. | So o 1% incondicional entra. O multiplicador de 7% sobre efeitos de movimentacao e multiplicativo sobre valor que vem de outro lugar, e quebra a separabilidade do objetivo linear. |
| 8236 | Tempestade Crescente | Gathering Storm | menor | out_of_scope | — | Recebe quantidades crescentes de DdA ou PdH Adaptativo ao longo da partida. | depende de abates acumulados na partida: Recebe quantidades crescentes de DdA ou PdH Adaptativo ao longo da partida. |
| 8237 | Chamuscar | Scorch | menor | out_of_scope | — | Sua primeira habilidade de dano a atingir o alvo a cada 10s incinerará Campeões. | depende de estado de combate: Sua primeira habilidade de dano a atingir o alvo a cada 10s incinerará Campeões. |
| 8275 | Manto de Nimbus | Nimbus Cloak | menor | out_of_scope | — | Após conjurar um Feitiço de Invocador, concede um breve aumento de Velocidade de Movimento que permite atravessar unidades. | depende de estado de combate: Após conjurar um Feitiço de Invocador, concede um breve aumento de Velocidade de Movimento que permite atraves |

## Inspiração (Inspiration)

| id | runa | nome canônico | tipo | escopo | stats | efeito | motivo |
|---|---|---|---|---|---|---|---|
| 8351 | Aprimoramento Glacial | Glacial Augment | pedra fundamental | out_of_scope | — | Imobilizar um Campeão inimigo invocará 3 raios glaciais que causam Lentidão a inimigos próximos e reduzem o dano deles contra seus aliados. | depende de estado de combate: Imobilizar um Campeão inimigo invocará 3 raios glaciais que causam Lentidão a inimigos próximos e reduzem o da |
| 8360 | Livro de Feitiços Deslacrado | Unsealed Spellbook | pedra fundamental | out_of_scope | — | Troque Feitiços de Invocador enquanto estiver fora de combate. Trocar para Feitiços de Invocador únicos aumentará a frequência com que você poderá fazer trocas futuras. | efeito sem atributo correspondente no vocabulario canonico: Troque Feitiços de Invocador enquanto estiver fora de combate. Trocar para Feitiços de Invocador únicos aument |
| 8369 | Primeiro Ataque | First Strike | pedra fundamental | out_of_scope | — | Ao iniciar um combate contra um Campeão, você causa 7% de dano adicional por 3s e recebe ouro com base no dano causado. | depende de estado de combate: Ao iniciar um combate contra um Campeão, você causa 7% de dano adicional por 3s e recebe ouro com base no dano |
| 8304 | Calçados Mágicos | Magical Footwear | menor | out_of_scope | — | Concede botas gratuitamente aos 12 mins, mas não é possível comprar botas antes disso. Cada abate bem-sucedido faz suas botas virem 45s mais cedo. | depende do tempo decorrido de partida: Concede botas gratuitamente aos 12 mins, mas não é possível comprar botas antes disso. Cada abate bem-sucedido |
| 8306 | Flashtração Hextec | Hextech Flashtraption | menor | out_of_scope | — | Enquanto o Flash estiver em Tempo de Recarga, ele é substituído pelo Flash Hextec. Flash Hextec: canalize e depois se teletransporte para um novo local. | depende de estado de combate: Enquanto o Flash estiver em Tempo de Recarga, ele é substituído pelo Flash Hextec. Flash Hextec: canalize e de |
| 8313 | Tônico Triplo | Triple Tonic | menor | out_of_scope | — | Ao alcançar o nível 3, recebe um Elixir da Avareza. Ao alcançar o nível 6, recebe um Elixir da Força. Ao alcançar o nível 9, recebe um Elixir da Habilidade. | depende do tempo decorrido de partida: Ao alcançar o nível 3, recebe um Elixir da Avareza. Ao alcançar o nível 6, recebe um Elixir da Força. Ao alcan |
| 8316 | Quebra-galho | Jack Of All Trades | menor | out_of_scope | — | A cada atributo diferente recebido de itens, recebe um acúmulo de Quebra-galho. Cada acúmulo concede a você 1 de Aceleração de Habilidade. Recebe Força Adaptativa adicional com 5 e 10 acúmulos. | depende de estado de combate: A cada atributo diferente recebido de itens, recebe um acúmulo de Quebra-galho. Cada acúmulo concede a você 1 |
| 8321 | Reembolso | Cash Back | menor | out_of_scope | — | Recebe parte do ouro de volta ao comprar itens Lendários. | depende do tempo decorrido de partida: Recebe parte do ouro de volta ao comprar itens Lendários. |
| 8345 | Entrega de Biscoitos | Biscuit Delivery | menor | out_of_scope | — | Ganha um Biscoito grátis a cada 2min, até o minuto 6. Consumir ou vender Biscoitos aumenta permanentemente sua Vida máxima e restaura Vida. | depende do tempo decorrido de partida: Ganha um Biscoito grátis a cada 2min, até o minuto 6. Consumir ou vender Biscoitos aumenta permanentemente sua |
| 8347 | Perspicácia Cósmica | Cosmic Insight | menor | sum | 10 acel item, 18 acel feitico | +18 de Aceleração de Feitiço de Invocador +10 de Aceleração de Item | — |
| 8352 | Tônico de Distorção no Tempo | Time Warp Tonic | menor | out_of_scope | — | Poções concedem um pouco de restauração imediatamente. | depende do tempo decorrido de partida: Poções concedem um pouco de restauração imediatamente. |
| 8410 | Velocidade de Aproximação | Approach Velocity | menor | out_of_scope | — | VdM adicional em direção a Campeões inimigos próximos com movimento debilitado, aumentada em Campeões inimigos cujo movimento você debilitou. | depende de estado de combate: VdM adicional em direção a Campeões inimigos próximos com movimento debilitado, aumentada em Campeões inimigos |

## Determinação (Resolve)

| id | runa | nome canônico | tipo | escopo | stats | efeito | motivo |
|---|---|---|---|---|---|---|---|
| 8437 | Aperto dos Mortos-Vivos | Grasp of the Undying | pedra fundamental | out_of_scope | — | A cada 4s, seu próximo ataque contra um Campeão causa Dano Mágico adicional, cura você e aumenta permanentemente sua Vida. | depende de estado de combate: A cada 4s, seu próximo ataque contra um Campeão causa Dano Mágico adicional, cura você e aumenta permanentemen |
| 8439 | Pós-choque | Aftershock | pedra fundamental | out_of_scope | — | Após imobilizar um Campeão inimigo, ganha defesas e depois causa um ataque de Dano Mágico explosivo ao seu redor. | depende de estado de combate: Após imobilizar um Campeão inimigo, ganha defesas e depois causa um ataque de Dano Mágico explosivo ao seu red |
| 8465 | Guardião | Guardian | pedra fundamental | out_of_scope | — | Protege aliados nos quais você conjurou habilidades e aliados muito próximos a você. Se você ou um Campeão aliado protegido sofrer dano com base no nível, ambos recebem um escudo. | depende de estado de combate: Protege aliados nos quais você conjurou habilidades e aliados muito próximos a você. Se você ou um Campeão ali |
| 8242 | Inabalável | Unflinching | menor | out_of_scope | — | Recebe Armadura e Resistência Mágica ao sofrer Controle de Grupo. | depende de estado de combate: Recebe Armadura e Resistência Mágica ao sofrer Controle de Grupo. |
| 8401 | Golpe de Escudo | Shield Bash | menor | out_of_scope | — | Sempre que receber um escudo, seu próximo ataque básico contra um Campeão causará Dano Adaptativo adicional. | depende de estado de combate: Sempre que receber um escudo, seu próximo ataque básico contra um Campeão causará Dano Adaptativo adicional. |
| 8429 | Condicionamento | Conditioning | menor | out_of_scope | — | Depois de 12min, recebe +8 de Armadura, +8 de Resistência Mágica e aumenta sua Armadura e Resistência Mágica em 3%. | depende do tempo decorrido de partida: Depois de 12min, recebe +8 de Armadura, +8 de Resistência Mágica e aumenta sua Armadura e Resistência Mágica e |
| 8444 | Ventos Revigorantes | Second Wind | menor | out_of_scope | — | Após receber dano de um Campeão inimigo, um pouco da Vida perdida é restaurada ao longo do tempo. | depende de estado de combate: Após receber dano de um Campeão inimigo, um pouco da Vida perdida é restaurada ao longo do tempo. |
| 8446 | Demolir | Demolish | menor | out_of_scope | — | Seu terceiro ataque contra torres causa dano adicional. | depende de estado de combate: Seu terceiro ataque contra torres causa dano adicional. |
| 8451 | Crescimento Excessivo | Overgrowth | menor | out_of_scope | — | Recebe Vida máxima permanente quando tropas inimigas ou monstros são abatidos perto de você. | depende de abates acumulados na partida: Recebe Vida máxima permanente quando tropas inimigas ou monstros são abatidos perto de você. |
| 8453 | Revitalizar | Revitalize | menor | sum | 5 %cura e escudo | Recebe 5% de Cura e Resistência do Escudo. Curas e Escudos conjurados ou recebidos são 10% mais fortes em alvos com menos de 40% de Vida. | So os 5% incondicionais entram. O bonus de 10% em alvos abaixo de 40% de vida depende de estado de combate e fica fora. |
| 8463 | Fonte da Vida | Font of Life | menor | out_of_scope | — | Debilitar o movimento de um Campeão inimigo cura Campeões aliados próximos. | depende de estado de combate: Debilitar o movimento de um Campeão inimigo cura Campeões aliados próximos. |
| 8473 | Osso Revestido | Bone Plating | menor | out_of_scope | — | Após sofrer dano de um Campeão inimigo, os próximos 3 Ataques ou Habilidades que você sofrer desse inimigo causarão 30 - 60 (com base no nível) a menos de dano. Duração: 1.5s Tempo de Recarga: 55s | depende de estado de combate: Após sofrer dano de um Campeão inimigo, os próximos 3 Ataques ou Habilidades que você sofrer desse inimigo cau |

## Fragmentos de stat

Valem em qualquer estilo. São 3 linhas, e uma runa pode aparecer em mais de
uma: Força Adaptativa está na linha ofensiva e na flexível.

| id | runa | nome canônico | tipo | escopo | stats | efeito | motivo |
|---|---|---|---|---|---|---|---|
| 5001 | Escalamento de Vida | Health Scaling | fragmento | sum_per_level | 180 vida | +10-180 de Vida (com base no nível) | A fonte declara +10-180 de Vida com base no nivel: 10 no nivel 1 e 180 no 18, o que da 10 por nivel acima do primeiro. |
| 5005 | Velocidade de Ataque | Attack Speed | fragmento | sum | 10 %AS | +10% de Velocidade de Ataque | — |
| 5007 | Aceleração de Habilidade | Ability Haste | fragmento | sum | 8 acel | +8 de Aceleração de Habilidade | — |
| 5008 | Força Adaptativa | Adaptive Force | fragmento | sum | 9 adapt | +9 de Força Adaptativa | — |
| 5010 | Velocidade de Movimento | Move Speed | fragmento | sum | 2.5 %mov | +2.5% de Velocidade de Movimento | — |
| 5011 | Vida | Health | fragmento | sum | 65 vida | +65 de Vida | — |
| 5013 | Tenacidade e Resistência a Lentidão | Tenacity and Slow Resist | fragmento | sum | 15 %tenacidade | +15% de Tenacidade e Resistência a Lentidão | A fonte declara tenacidade E resistencia a lentidao no mesmo numero. So a tenacidade tem stat no vocabulario; a resistencia a lentidao vai junto no jogo, mas nao e somada aqui. |

