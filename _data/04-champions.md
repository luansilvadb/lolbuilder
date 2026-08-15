# Campeões — patch 16.16

173 campeões. As estatísticas base e o efeito das habilidades por rank
estão em `06-champion-stats.md`; aqui está o que cada habilidade faz.

> O texto vem do cliente, que o publica com marcadores no lugar dos números.
> Eles foram **resolvidos** contra as séries do dump: onde a fonte escreve
> `@TotalDamage@`, aqui está o valor por rank. Marcador sem valor na fonte foi
> removido em vez de publicado cru — um `@Nome@` visível não informa nada e
> ainda parece defeito de geração.

## Annie, a Criança Sombria

`Annie` · id 1 · a distância · mage, support

- **P · Piromania** — Após conjurar 4 Habilidades, a próxima Habilidade ofensiva de Annie atordoará o alvo. Annie começa a partida e ressurge com Piromania disponível.
- **Q · Desintegrar** — Annie lança uma bola de fogo, causando 80/125/170/215/260 + 0.8 ability_power de Dano Mágico. Se o alvo morrer, ela recupera o custo de Mana e reduz o Tempo de Recarga da Habilidade em 50%.(?)
- **W · Incinerar** — Annie lança uma onda de fogo, causando 70/110/150/190/230 + 0.8 ability_power de Dano Mágico.(?)
- **E · Escudo Fundido** — Annie concede a um Campeão aliado 60/95/130/165/200 + 0.4 ability_power de Escudo por 3/3/3/3/3s e 0.5/0.5/0.5/0.5/0.5 de Velocidade de Movimento que decai por 1.5/1.5/1.5/1.5/1.5s. Enquanto o Escudo estiver ativo, inimigos que atingirem o aliado protegido com Ataques ou Habilidades sofrerão 25/35/45/55/65 + 0.4 ability_power de Dano Mágico uma vez por Escudo. Tibbers sempre recebe os efeitos de Escudo Fundido quando é invocado.(?)
- **R · Invocar: Tibbers** — Passivo: Annie recebe 10/15/20% de Penetração Mágica. Annie invoca seu urso Tibbers, causando 150/275/400 + 0.75 ability_power de Dano Mágico. Pelos próximos 45/45/45s, Tibbers incinera inimigos próximos com 8/12/16 + 0.04 ability_power de Dano Mágico por segundo. Tibbers ficará enfurecido ao ser invocado, se Annie Atordoar um Campeão inimigo e se Annie morrer. Ao ficar enfurecido, Tibbers recebe 275% de Velocidade de Ataque e 100% de Velocidade de Movimento, que decaem ao longo de 3s. Reconjuração: comanda Tibbers manualmente.(?)

## Olaf, o Berserker

`Olaf` · id 2 · corpo a corpo · fighter, tank

- **P · Fúria Berserker** — Olaf recebe Velocidade de Ataque e Roubo de Vida com base na sua Vida perdida.
- **Q · Ressaca** — Olaf arremessa um machado, causando 70/120/170/220/270 + 1 attack_damage (bonus) de Dano Físico e aplicando 30/35/40/45/50% de Lentidão por até 3/3/3/3/3s (com base na distância percorrida). Campeões atingidos perdem 20/20/20/20/20% de Armadura por 4/4/4/4/4s. Se Olaf apanhar o machado, o Tempo de Recarga desta habilidade é reduzido para 2.5/2.5/2.5/2.5/2.5s, ou completamente restituído se 2.5/2.5/2.5/2.5/2.5s tiverem decorrido.(?)
- **W · Duro na Queda** — Olaf recebe 40/50/60/70/80% de Velocidade de Ataque por 5/5/5/5/5s, e 10/40/70/100/130 mais 17.5/17.5/17.5/17.5/17.5% de Escudo com base na Vida perdida (até um máximo de 10/40/70/100/130 + 0.122499995 max_health de Escudo com menos de 30/30/30/30/30% de Vida) por 2.5/2.5/2.5/2.5/2.5s.(?)
- **E · Balanço Temerário** — Olaf brande ferozmente seus machados, causando 70/115/160/205/250 + 0.5 attack_damage de Dano Verdadeiro. Se o inimigo morrer, o custo será restituído. Ataques reduzem o Tempo de Recarga em 1s, aumentado para 2s ao atacar monstros.(?)
- **R · Ragnarok** — Passivo: Olaf recebe 10/15/20 de Armadura e 10/15/20 de Resistência Mágica. Ativo: Olaf purifica todos os efeitos Imobilizadores e Debilitantes que o estiverem afetando, tornando-se imune a eles durante 3/3/3s. Enquanto a habilidade estiver ativa, Olaf recebe 10/20/30 + 0.25 attack_damage de Dano de Ataque. Atingir um Campeão com um Ataque ou Balanço Temerário aumenta a duração em 2.5/2.5/2.5s. Além disso, Olaf recebe 20/45/70% de Velocidade de Movimento ao mover-se em direção a Campeões inimigos por 1/1/1s.(?)

## Galio, o Colosso

`Galio` · id 3 · corpo a corpo · tank, mage

- **P · Esmagada Colossal** — A cada poucos segundos, o próximo ataque básico de Galio causa Dano Mágico adicional em área.
- **Q · Ventos de Guerra** — Galio dispara duas rajadas de vento que causam 70/105/140/175/210 + 0.7 ability_power de Dano Mágico cada. Quando elas se encontram, combinam-se em um tornado que causa 8/8/8/8/8 + 0.04 ability_power% da Vida máxima como Dano Mágico ao longo de 2/2/2/2/2s.(?)
- **W · Escudo de Durand** — Passivo: Galio recebe 0/0/0/0/0 + 0.075 max_health de Escudo Mágico após ficar sem sofrer dano por 8/8/8/8/8s. Início do carregamento: Galio reduz o Dano Mágico sofrido em 0.25/0.3/0.35/0.4/0.45 + 0.0004 ability_power + 0.0008 magic_resist (bonus) + 0.0001 max_health (bonus), o Dano Físico sofrido em 0.125/0.15/0.175/0.2/0.225 + 0.0002 ability_power + 0.0004 magic_resist (bonus) + 0.00005 max_health (bonus) e aplica 15/15/15/15/15% de Lentidão a si mesmo. Liberar: Galio Provoca Campeões inimigos por 0.5/0.5/0.5/0.5/0.5s - 1.5/1.5/1.5/1.5/1.5s, causa 20/30/40/50/60 + 0.3 ability_power - 60/90/120/150/180 + 0.90000004 ability_power de Dano Mágico e redefine a redução de dano por 2/2/2/2/2s. O dano, o alcance e a duração da Provocação aumentam com o tempo de carregamento.(?)
- **E · Soco Justiceiro** — Galio avança e desfere um golpe poderoso, Arremessando ao ar por 0.75/0.75/0.75/0.75/0.75s e causando 100/135/170/205/240 + 1 ability_power de Dano Mágico ao primeiro Campeão atingido. Outros inimigos em seu trajeto sofrem 80/108/136/164/192 + 0.8 ability_power de Dano Mágico. O avanço de Galio é encerrado ao atingir um terreno.(?)
- **R · Entrada Heroica** — Galio marca a posição atual de um aliado como ponto de aterrissagem e concede o Escudo do efeito passivo de Escudo de Durand por 5/5/5s a todos os Campeões aliados na área. Depois, ele voa para a zona de aterrissagem. Ao aterrissar, Arremessa ao ar os inimigos por 0.75/0.75/0.75s e causa 150/250/350 + 0.7 ability_power + 1 magic_resist (bonus) de Dano Mágico a eles.(?)

## Twisted Fate, o Mestre das Cartas

`Twisted Fate` · id 4 · a distância · mage, marksman

- **P · Dados Viciados** — Ao abater uma unidade, Twisted Fate rola seu dado da ''sorte'' e recebe de 1 a 6 de ouro adicional.
- **Q · Curingas** — Twisted Fate arremessa três cartas que causam 60/105/150/195/240 + 0.85 ability_power + 0.5 attack_damage (bonus) de Dano Mágico cada.(?)
- **W · Escolha uma Carta** — Twisted Fate mistura seu baralho, podendo Reconjurar para selecionar uma das três cartas e aprimorar seu próximo Ataque. A Carta Azul causa (?) de Dano Mágico e restaura (?) de Mana.A Carta Vermelha causa (?) a inimigos próximos e (?)% de Lentidão por 2,5s.A Carta Dourada causa (?) e Atordoa por (?)s.(?)
- **E · Baralho Marcado** — Passiva: Twisted Fate recebe 15/25/35/45/55% de Velocidade de Ataque e todo 4º Ataque causa 65/90/115/140/165 + 0.4 ability_power + 0.2 attack_damage (bonus) de Dano Mágico adicional.(?)
- **R · Destino** — Twisted Fate concentra-se em suas cartas, concedendo Visão Mágica de todos os Campeões inimigos no mapa por (?)s e permitindo que Reconjure. Reconjuração: Twisted Fate teleporta-se para até (?) unidades de distância.(?)

## Xin Zhao, o Senescal de Demacia

`Xin Zhao` · id 5 · corpo a corpo · fighter, tank

- **P · Determinação** — Cada terceiro ataque básico causa dano adicional e cura Xin Zhao.
- **Q · Golpe de Três Garras** — Os próximos 3 Ataques de Xin Zhao causam 15/30/45/60/75 + 0.4 attack_damage (bonus) de Dano Físico adicional e reduzem os Tempos de Recarga das outras habilidades em 1s. O terceiro Ataque também Arremessa ao ar por 0.75/0.75/0.75/0.75/0.75s. (?)
- **W · Vento Vira Relâmpago** — Xin Zhao desfere um golpe cortante, causando 30/40/50/60/70 + 0.3 attack_damage de Dano Físico. Depois, ele desfere um golpe perfurante, causando 50/85/120/155/190 + 0.65 ability_power + 0.9 attack_damage de Dano Físico. Inimigos atingidos pelo golpe perfurante sofrem @Effect6Amount*-100@% de Lentidão por 1.5/1.5/1.5/1.5/1.5 + 0.005 ability_powers. Campeões e monstros grandes atingidos pelo golpe perfurante são marcados como Desafiados por 3/3/3/3/3s e revelados, a não ser que estejam em estado furtivo.(?)
- **E · Investida Audaciosa** — Xin Zhao avança em um inimigo, causando a inimigos próximos 50/75/100/125/150 + 1.2 ability_power de Dano Mágico e 30/30/30/30/30% de Lentidão por 0.5/0.5/0.5/0.5/0.5s. Xin Zhao recebe (?)% (%i:scaleAS%%i:scaleAP%) de Velocidade de Ataque por 5/5/5/5/5s. Esta Habilidade recebe alcance aumentado contra inimigos Desafiados. (?)
- **R · Guarda Crescente** — Passivo: o último Campeão atingido por um Ataque ou pela Investida Audaciosa de Xin Zhao fica Desafiado por 3/3/3s. Ativo: Xin Zhao desfere um golpe giratório ao redor de si, causando 75/175/275 + 1.1 ability_power + 1 attack_damage (bonus) mais 15/15/15% da Vida atual como Dano Físico e Empurrando todos os inimigos não Desafiados para trás. Depois, Xin Zhao se torna imune ao dano causado por inimigos que estiverem fora do alcance do golpe por 4/4/4s. (?)

## Urgot, o Encouraçado

`Urgot` · id 6 · a distância · fighter, tank

- **P · Chamas Ecoantes** — Os ataques básicos de Urgot e Expurgar disparam periodicamente explosões de chamas de suas pernas, causando Dano Físico.
- **Q · Carga Corrosiva** — Urgot detona uma carga explosiva, causando 25/70/115/160/205 + 0.7 attack_damage de Dano Físico e 45/50/55/60/65% de Lentidão por 1.25/1.25/1.25/1.25/1.25s.(?)
- **W · Expurgar** — Passivo: as outras Habilidades de Urgot marcam o último Campeão atingido por 5s. Ativo: Urgot dispara seu canhão automático no inimigo mais próximo, priorizando os marcados. Ele os Ataca 3/3/3/3/3 vezes por segundo, causando 12/12/12/12/12 + 0.2 attack_damage de Dano Físico por disparo. Urgot pode se mover enquanto dispara e tem 40/40/40/40/40% de resistência a Lentidão, mas perde 125/125/125/125/125 de Velocidade de Movimento. No ranque máximo, a Habilidade dura por um período indeterminado e pode ser Alternada entre ativada e desativada.(?)
- **E · Desdém** — Urgot avança, recebendo 55/75/95/115/135 + 1.35 attack_damage (bonus) + 0.135 max_health (bonus) de Escudo por 4/4/4/4/4s. O primeiro Campeão atingido é Atordoado por 1.5/1.5/1.5/1.5/1.5s e jogado para trás de Urgot. Todos os inimigos com quem Urgot colidir sofrerão 90/120/150/180/210 + 1 attack_damage (bonus) de Dano Físico.(?)
- **R · Pior que a Morte** — Urgot dispara uma broca química que perfura o primeiro Campeão inimigo atingido, causando 100/225/350 + 0.5 attack_damage (bonus) de Dano Físico e Lentidão de 1% para cada 1% de Vida perdida por 4/4/4s, até um máximo de 75/75/75%. Se a vítima perfurada ficar com menos de 25/25/25% de Vida, Urgot pode Reconjurar a Habilidade, Suprimindo o alvo e puxando-o em sua direção. Ao chegar na posição de Urgot, a vítima será abatida e os inimigos próximos sofrerão Temor por 1.5/1.5/1.5s.(?)

## LeBlanc, a Farsante

`LeBlanc` · id 7 · a distância · assassin, mage

- **P · Imagem-Espelho** — Quando a Vida de LeBlanc está abaixo de 40%, ela fica invisível por 1s e cria uma Imagem-Espelho que não causa dano e dura até 8s.
- **Q · Sigilo de Malícia** — LeBlanc projeta um sigilo em um inimigo, causando a ele 65/90/115/140/165 + 0.4 ability_power de Dano Mágico e marcando-o por 3.5/3.5/3.5/3.5/3.5s. Causar dano ao inimigo marcado com qualquer Habilidade detona o sigilo, causando 65/90/115/140/165 + 0.4 ability_power de Dano Mágico. Se qualquer parte do combo abater o alvo, LeBlanc restitui 100/100/100/100/100% do custo de Mana e 30/30/30/30/30% do Tempo de Recarga durante o restante desta habilidade. O sigilo inicial causa 10/10/10/10/10 de dano adicional a tropas. (?)
- **W · Distorção** — LeBlanc avança e causa 75/115/155/195/235 + 0.8 ability_power de Dano Mágico aos inimigos próximos. Por 4/4/4/4/4s depois de avançar, LeBlanc pode Reconjurar. Reconjuração: LeBlanc retorna ao local inicial.(?)
- **E · Correntes Etéreas** — LeBlanc lança uma corrente que se prende ao primeiro inimigo atingido, causando 50/70/90/110/130 + 0.4 ability_power de Dano Mágico e concedendo Visão Mágica. Caso permaneça preso por 1.5/1.5/1.5/1.5/1.5s, o inimigo é Enraizado por 1.5/1.5/1.5/1.5/1.5s e sofre 80/120/160/200/240 + 0.85 ability_power de Dano Mágico adicional.(?)
- **R · Mímica** — LeBlanc copia sua Habilidade mais recente, usando-a novamente. A Habilidade copiada causa dano aumentado. Sigilo de Malícia com Mímica causa 70/150/230 + 0.4 ability_power de Dano Mágico quando é aplicado e 140/300/460 + 0.8 ability_power de Dano Mágico quando é consumido. Distorção com Mímica causa 150/315/480 + 0.8 ability_power de Dano Mágico. Correntes Etéreas com Mímica causa 70/150/230 + 0.4 ability_power de Dano Mágico quando se prende ao inimigo e 140/300/460 + 0.85 ability_power de Dano Mágico quando o Enraíza.(?)

## Vladimir, o Sanguinário Escarlate

`Vladimir` · id 8 · a distância · mage, fighter

- **P · Pacto Vermelho** — Cada 30 pontos de Vida adicional concedem 1 de Poder de Habilidade a Vladimir, e cada 1 ponto de Poder de Habilidade concede a ele 1,6 de Vida adicional (não cumulativos entre si).
- **Q · Transfusão** — Vladimir drena a energia vital do alvo, causando 80/100/120/140/160 + 0.6 ability_power de Dano Mágico e restaurando 20/25/30/35/40 + 0.35 ability_power de Vida. Depois de usar esta Habilidade duas vezes, Vladimir recebe 40/40/40/40/40% de Velocidade de Movimento por 0,5s e fortalece o próximo uso da Habilidade por 2.5/2.5/2.5/2.5/2.5s. Quando fortalecida, causa 148/185/222/259/296 + 1.11 ability_power de Dano Mágico e restaura um adicional de 200/200/200/200/200 mais 5/5/5/5/5 + 0.04 ability_power da Vida perdida.(?)
- **W · Poça de Sangue** — Vladimir mergulha numa poça de sangue por 2s, recebendo 37.5/37.5/37.5/37.5/37.5% de Velocidade de Movimento que decai ao longo de 1/1/1/1/1s, tornando-se Inalvejável e recebendo efeito Fantasma enquanto causa @MoveSpeedMod*-100@% de Lentidão aos inimigos dentro da poça. Vladimir causa 80/135/190/245/300 + 0.15 max_health (bonus) de Dano Mágico e restaura 24/40.5/57.000004/73.5/90 + 0.045 max_health (bonus) de Vida por inimigo ao longo da duração.(?)
- **E · Maré de Sangue** — Início do carregamento: Vladimir carrega um reservatório de sangue, gastando até 0/0/0/0/0 + 0.08 max_health de Vida. Com o carregamento máximo, Vladimir sofre 20% de Lentidão. Liberar: Vladimir desfere uma onda de projéteis sanguíneos nos inimigos ao redor, causando 30/45/60/75/90 + 0.35 ability_power + 0.015 max_health - 60/90/120/150/180 + 0.8 ability_power + 0.06 max_health de Dano Mágico com base no tempo de carregamento. Caso a Habilidade seja carregada por pelo menos 1s, ela também causa 40/45/50/55/60% de Lentidão aos alvos por 0,5s.(?)
- **R · Hemopraga** — Vladimir cria uma praga virulenta, fazendo com que suas vítimas sofram 10/10/10% a mais de dano de todas as origens por 4/4/4s. Quando esta habilidade expira, Vladimir causa 150/250/350 + 0.7 ability_power de Dano Mágico a todos os alvos infectados. Vladimir restaura 150/250/350 + 0.7 ability_power de Vida caso atinja um Campeão e mais 60/100/140 + 0.28 ability_power de Vida adicional para cada Campeão além do primeiro.(?)

## Fiddlesticks, o Terror Ancestral

`Fiddlesticks` · id 9 · a distância · mage, support

- **P · Um Espantalho Inofensivo** — O amuleto de Fiddlesticks é substituído por efígies do espantalho.
- **Q · Aterrorizar** — Passivo: se causar dano a um inimigo com uma Habilidade enquanto não estiver sendo visto e estiver fora de combate ou imitando uma Efígie, Fiddlesticks causa Temor a ele por 1.2/1.4/1.6/1.8/2s. Ativo: causa Temor a um inimigo por 1.2/1.4/1.6/1.8/2s e 0.04/0.045/0.05/0.055/0.06 + 0.0003 ability_power da Vida atual como Dano Mágico. Se o alvo tiver sofrido Temor causado por Fiddlesticks recentemente, a Habilidade causa 0.08/0.09/0.1/0.11/0.12 + 0.0006 ability_power da Vida atual do alvo como Dano Mágico. (?)
- **W · Colheita Farta** — Fiddlesticks canaliza e drena as almas de inimigos próximos, causando 60/90/120/150/180 + 0.45 ability_power de Dano Mágico por segundo ao longo de 2s, mais 12/14.5/17/19.5/22% da Vida perdida como Dano Mágico ao final da canalização. Fiddlesticks restaura 25/32.5/40/47.5/55% do dano causado como Vida. Se a canalização terminar sem ser interrompida, o Tempo de Recarga restante é reduzido em 60%. (?)
- **E · Ceifar** — Fiddlesticks libera magia sombria, causando 70/105/140/175/210 + 0.5 ability_power de Dano Mágico e @SlowAmount*-100@% de Lentidão aos inimigos por 1.25/1.25/1.25/1.25/1.25s. Inimigos no centro também são Silenciados ao longo da duração.(?)
- **R · Tempestade de Corvos** — Fiddlesticks canaliza por 1.5/1.5/1.5s e depois se teleporta, liberando uma revoada de corvos e causando 750/1250/1750 + 2.5 ability_power de Dano Mágico ao longo de 5/5/5s.(?)

## Kayle, a Justa

`Kayle` · id 10 · corpo a corpo · mage, marksman

- **P · Ascensão Divina** — Os ataques de Kayle são fortalecidos pelos céus conforme ela aumenta de nível de Campeão e de habilidades. Suas asas ficam flamejantes conforme ela progressivamente ganha Velocidade de Ataque, Velocidade de Movimento, Alcance de Ataque e ondas de fogo em seus ataques.
- **Q · Explosão Radiante** — Kayle dispara uma espada celestial que para no primeiro inimigo atingido, causando 60/90/120/150/180 + 0.5 ability_power + 0.6 attack_damage (bonus) de Dano Mágico e 25/30/35/40/45% de Lentidão por 2/2/2/2/2s. Além disso, ela remove 15/15/15/15/15% de Armadura e de Resistência Mágica do alvo e de inimigos atrás dele por 4/4/4/4/4s.(?)
- **W · Bênção Celestial** — Kayle envolve-se em luz e confere o mesmo efeito a um Campeão aliado, restaurando 55/80/105/130/155 + 0.25 ability_power de Vida e concedendo 0.24/0.28/0.32/0.36/0.4 + 0.0008 ability_power de Velocidade de Movimento por 2/2/2/2/2s.(?)
- **E · Lâmina de Fogo Estelar** — Passivo: os Ataques causam 15/20/25/30/35 + 0.2 ability_power + 0.1 attack_damage (bonus) de Dano Mágico adicional. Ativo: o próximo Ataque de Kayle é realizado à distância e causa 8/8.5/9/9.5/10 + 0.015 ability_power da Vida perdida como Dano Mágico adicional. Esse Ataque é aprimorado quando Kayle alcança o nível @Spell.KaylePassive:LevelForPassiveRank2@, fazendo com que ele exploda ao atingir o alvo e cause dano a inimigos próximos.(?)
- **R · Sentença Divina** — Kayle torna um Campeão aliado Invulnerável por 2.5/2.5/2.5s e purifica a área ao redor dele, causando 200/300/400 + 0.7 ability_power + 1 attack_damage (bonus) de Dano Mágico a inimigos próximos.(?)

## Master Yi, o Espadachim Wuju

`Master Yi` · id 11 · corpo a corpo · fighter, assassin

- **P · Ataque duplo** — A cada alguns ataques básicos consecutivos, Master Yi ataca duas vezes.
- **Q · Ataque Alpha** — Master Yi torna-se Inalvejável, teleporta-se e ataca rapidamente os inimigos próximos ao alvo, causando 20/40/60/80/100 + 0.7 attack_damage de Dano Físico a todos os inimigos atingidos após 4/4/4/4/4 golpes. A Habilidade pode atingir o mesmo inimigo diversas vezes se não houver outros alvos, causando 25/25/25/25/25% de dano nos acertos subsequentes (20/40/60/80/100 + 0.7 attack_damage), até um máximo de 35/70/105/140/175 + 1.225 attack_damage de Dano Físico por alvo.(?)
- **W · Meditar** — Master Yi canaliza, recuperando 120/200/280/360/440 + 1 ability_power de Vida ao longo de 4/4/4/4/4s. A cura aumenta em até 100/100/100/100/100% com base na Vida perdida de Master Yi. Enquanto canaliza e por 0.5/0.5/0.5/0.5/0.5s depois, ele sofre 0.7/0.7/0.7/0.7/0.70000005 a menos de dano, reduzido para 45/47.5/50/52.5/55% depois dos primeiros 0.5/0.5/0.5/0.5/0.5s.(?)
- **E · Estilo Wuju** — Os Ataques de Master Yi causam 20/25/30/35/40 + 0.35 attack_damage (bonus) de Dano Verdadeiro adicional por 5/5/5/5/5s.(?)
- **R · Highlander** — Passivo: eliminar Campeões reduz o Tempo de Recarga restante das Habilidades básicas de Master Yi em 70/70/70%. Ativo: Master Yi entra em transe, recebendo 35/45/55% de Velocidade de Movimento, 25/45/65% de Velocidade de Ataque e imunidade a Lentidão por 7/7/7s. Eliminar um Campeão estende a duração da Habilidade em 7/7/7s.(?)

## Alistar, o Minotauro

`Alistar` · id 12 · corpo a corpo · tank, support

- **P · Urro Triunfante** — Alistar carrega seu urro ao atordoar, deslocar a posição de Campeões inimigos ou quando inimigos próximos morrem. Quando estiver com o máximo de cargas, ele cura a si mesmo e aos Campeões aliados próximos.
- **Q · Pulverizar** — Alistar golpeia o chão, Arremessando ao ar os inimigos por 1/1/1/1/1s e causando 60/100/140/180/220 + 0.8 ability_power de Dano Mágico.(?)
- **W · Cabeçada** — Alistar avança até um inimigo, Empurrando-o e causando 55/110/165/220/275 + 1 ability_power de Dano Mágico.(?)
- **E · Atropelar** — Alistar começa a pisotear o chão, recebendo efeito Fantasma e causando 80/110/140/170/200 + 0.7 ability_power de Dano Mágico a inimigos próximos ao longo de 5/5/5/5/5s. Cada pulso que causar dano a um Campeão concederá um acúmulo. Com 5/5/5/5/5 acúmulos, o próximo Ataque de Alistar contra Campeões Atordoa por 1/1/1/1/1s e causa 20/20/20/20/20 de Dano Mágico adicional.(?)
- **R · Vontade Indestrutível** — Alistar purifica imediatamente todos os efeitos Debilitantes e sofre 55/65/75% de dano reduzido por 7/7/7s.(?)

## Ryze, o Mago Rúnico

`Ryze` · id 13 · a distância · mage

- **P · Maestria Arcana** — As habilidades de Ryze causam dano adicional com base em seu Mana adicional. Ele recebe um percentual de aumento de Mana máximo com base em seu Poder de Habilidade.
- **Q · Sobrecarregar** — Passivo: Prisão de Runa e Fluxo de Feitiço redefinem o Tempo de Recarga desta Habilidade e carrega uma runa por 4/4/4/4/4s (máximo de 2/2/2/2/2 runas). Ativo: Ryze libera uma explosão, causando 75/95/115/135/155 + 0.55 ability_power + 0.02 mana de Dano Mágico ao primeiro inimigo atingido. Se o alvo estiver afetado por Fluxo, ele é consumido, fazendo com que a Habilidade cause @Spell.RyzeR:OverloadDamageBonus@% de dano aumentado e ricocheteie para inimigos próximos afetados por Fluxo. Ryze também descarrega todas as runas, recebendo 28/32/36/40/44% de Velocidade de Movimento por 2/2/2/2/2s se 2/2/2/2/2 runas estavam carregadas. (?)
- **W · Prisão de Runa** — Ryze causa 60/90/120/150/180 + 0.6 ability_power + 0.03 mana de Dano Mágico e 50/50/50/50/50% de Lentidão por 1.5/1.5/1.5/1.5/1.5s. Se o alvo estiver afetado por Fluxo de Feitiço, ele é consumido e a Habilidade Enraíza em vez de causar Lentidão.(?)
- **E · Fluxo de Feitiço** — Ryze dispara um orbe, causando 60/90/120/150/180 + 0.5 ability_power + 0.02 mana de Dano Mágico e aplicando Fluxo por 4/4/4/4/4s ao alvo e aos inimigos próximos. Inimigos já afetados por Fluxo farão com que ele se espalhe ainda mais.(?)
- **R · Portal de Reinos** — Passivo: o dano adicional de Sobrecarregar contra alvos afetados por Fluxo é aumentado para 50/75/100%. Ativo: Ryze abre um portal para outro local. Após 2/2/2s, todos os aliados próximos ao portal são teleportados para o local escolhido.(?)

## Sion, o Colosso Morto-Vivo

`Sion` · id 14 · corpo a corpo · tank, fighter

- **P · In Gloria Mori** — Após ser abatido, Sion será reanimado temporariamente, mas sua Vida decairá rapidamente. Seus ataques ficarão muito rápidos, curando e causando dano adicional com base na Vida máxima do alvo.
- **Q · Golpe Demolidor** — Início do carregamento: Sion carrega um golpe poderoso por até 2s. Lançamento: Sion bate no chão com seu machado, causando uma breve Lentidão e 30/45/60/75/90 + 0.4 attack_damage - 90/155/220/285/350 + 1.2 attack_damage de Dano Físico com base no tempo de carregamento. Se Sion carregar por pelo menos 1s, os inimigos são Arremessados ao ar e Atordoados por 1.25/1.25/1.25/1.25/1.25 - 2,25s com base no tempo de carregamento.(?)
- **W · Fornalha da Alma** — Passivo: Sion recebe 4/4/4/4/4 de Vida máxima ao abater uma unidade ou 15/15/15/15/15 ao eliminar Campeões, tropas grandes e monstros grandes. Ativo: Sion recebe 60/75/90/105/120 + 0.4 ability_power + 0.08 max_health de Escudo por 6s. Se o Escudo ainda estiver ativo depois de 3/3/3/3/3s, Sion poderá Reconjurar para detoná-lo e causar 40/65/90/115/140 + 0.4 ability_power mais 14/14/14/14/14% da Vida máxima como Dano Mágico.(?)
- **E · Urro do Assassino** — Sion dispara uma onda de choque, causando 65/100/135/170/205 + 0.55 ability_power de Dano Mágico, 40/45/50/55/60% de Lentidão por 2.5/2.5/2.5/2.5/2.5s e removendo 25/25/25/25/25% de Armadura por 4/4/4/4/4s. Não Campeões próximos atingidos são Empurrados. Inimigos atingidos pela unidade Empurrada sofrem o mesmo dano e efeitos.(?)
- **R · Investida Incontrolável** — Sion avança sem poder ser interrompido por 8s, virando em direção ao cursor do mouse. Ele para ao Reconjurar ou colidir com um Campeão inimigo ou terreno. No fim do carregamento, Sion causa de 150/300/450 + 0.6 attack_damage (bonus) a 400/800/1200 + 1.2 attack_damage (bonus) de Dano Físico com base na distância percorrida. Inimigos próximos a ele são Atordoados de 0.75/0.75/0.75s a 1.75/1.75/1.75s com base na distância percorrida. Inimigos numa área maior sofrem 40/45/50% de Lentidão por 3s.(?)

## Sivir, a Mestra da Batalha

`Sivir` · id 15 · a distância · marksman

- **P · Pés Ligeiros** — Sivir ganha um pequeno impulso de Velocidade de Movimento quando ela ataca um campeão inimigo.
- **Q · Lâmina Bumerangue** — Sivir arremessa sua arma como um bumerangue, causando 60/85/110/135/160 + 0.6 ability_power + 0.7 attack_damage (bonus) a todos os inimigos que atravessar. Atingir não Campeões reduz o dano nos alvos subsequentes, até um mínimo de 40/40/40/40/40%.(?)
- **W · Ricochete** — Pelos próximos 4/4/4/4/4s, Sivir recebe 20/25/30/35/40% de Velocidade de Ataque e seus ataques básicos são fortalecidos e rebatem em mais inimigos próximos, causando 0/0/0/0/0 + 0.4 attack_damage de Dano Físico por ricochete até um máximo de 8/8/8/8/8 ricochetes. Se um ataque causar Acerto Crítico, o ricochete também causará.(?)
- **E · Escudo de Feitiço** — Sivir cria uma barreira mágica por 1.5/1.5/1.5/1.5/1.5s que bloqueia a próxima Habilidade inimiga. Caso a Habilidade seja bloqueada, Sivir restaura 0/0/0/0/0 + 0.5 ability_power + 0.6 attack_damage de Vida e ativa Pés Ligeiros.(?)
- **R · Na Caçada** — Sivir convoca seus aliados próximos, concedendo a eles 20/25/30% de Velocidade de Movimento por 8/10/12s. Os ataques de Sivir contra Campeões durante Na Caçada reduzem o Tempo de Recarga das habilidades básicas em 0.5/0.5/0.5s. Eliminações de inimigos atingidos recentemente redefinem a duração da caçada.(?)

## Soraka, a Filha das Estrelas

`Soraka` · id 16 · a distância · support, mage

- **P · Salvação** — Soraka corre mais rapidamente em direção a aliados com pouca Vida.
- **Q · Chamado Estelar** — Soraka conjura uma estrela, causando 85/120/155/190/225 + 0.35 ability_power de Dano Mágico e 30/30/30/30/30% de Lentidão aos inimigos por 1.5/1.5/1.5/1.5/1.5s. Atingir um Campeão inimigo concede Rejuvenescimento a Soraka, que restaura 60/75/90/105/120 + 0.3 ability_power de Vida ao longo de 2.5/2.5/2.5/2.5/2.5s e concede 20/22.5/25/27.5/30% de Velocidade de Movimento que decai ao longo da mesma duração.(?)
- **W · Infusão Astral** — Soraka restaura 90/110/130/150/170 + 0.5 ability_power de Vida de outro Campeão aliado. Se ela estiver sob o efeito de Rejuvenescimento, o custo de Vida será reduzido em 80/85/90/95/100% e o alvo receberá Rejuvenescimento por @Spell.SorakaQ:HoTDuration(?)SpellModifierDescriptionAppend@
- **E · Equinócio** — Soraka cria um campo de estrelas que causa 70/95/120/145/170 + 0.4 ability_power de Dano Mágico a Campeões. O campo perdura por 1.5/1.5/1.5/1.5/1.5s, Silenciando os inimigos que estiverem dentro dele. Quando a área desaparece, os Campeões que permanecem no local são Enraizados por 1/1.25/1.5/1.75/2s e sofrem 70/95/120/145/170 + 0.4 ability_power de Dano Mágico.(?)
- **R · Desejo** — Soraka invoca poderes divinos, restaurando 150/250/350 + 0.5 ability_power de Vida de todos os Campeões aliados, independentemente da distância. A cura aumenta para 225/375/525 + 0.75 ability_power em alvos com menos de 40% de Vida.(?)

## Teemo, o Explorador Veloz

`Teemo` · id 17 · a distância · marksman, mage

- **P · Técnicas de Guerrilha** — Se Teemo permanece imóvel e não faz nada por um curto período, fica Invisível por tempo indefinido. Caso esteja em um arbusto, Teemo pode manter sua Invisibilidade enquanto se move. Ao sair da Invisibilidade, Teemo ativa Elemento-Surpresa, aumentando a própria Velocidade de Ataque por alguns segundos.
- **Q · Dardo Ofuscante** — Teemo dispara um dardo, Cegando o alvo por 2/2.25/2.5/2.75/3s e causando 80/125/170/215/260 + 0.7 ability_power de Dano Mágico.(?)
- **W · Mover Depressa** — Passivo: Teemo recebe 12/16/20/24/28% de Velocidade de Movimento quando não tiver sofrido dano de um Campeão ou torre nos últimos 5/5/5/5/5s. Ativo: Teemo acelera, recebendo 24/32/40/48/56% de Velocidade de Movimento que não é perdida ao sofrer dano por 3/3/3/3/3s.(?)
- **E · Tiro Tóxico** — Passivo: os Ataques de Teemo aplicam veneno %i:OnHit% ao contato, causando 9/23/37/51/65 + 0.3 ability_power + 0.05 attack_damage (bonus) de Dano Mágico adicional mais 24/48/72/96/120 + 0.4 ability_power + 0.1 attack_damage (bonus) de Dano Mágico ao longo de 4/4/4/4/4s.(?)
- **R · Armadilha Venenosa** — Teemo arremessa uma armadilha de cogumelo que detona ao ser pisada. As armadilhas causam 30/40/50% de Lentidão e 200/325/450 + 0.5 ability_power de Dano Mágico ao longo de 4/4/4s. Inimigos são revelados pela mesma duração. As armadilhas duram 5/5/5min e ficam em estado furtivo. Um cogumelo arremessado sobre outro ricocheteia antes de se fixar em uma posição. A Habilidade tem 3/4/5 cargas (recarrega a cada (?)s). (?)

## Tristana, a Artilheira Yordle

`Tristana` · id 18 · a distância · marksman, assassin

- **P · Tiro Certeiro** — Aumenta o Alcance de Ataque de Tristana conforme ela sobe de nível.
- **Q · Tiro Rápido** — Tristana aumenta a frequência de disparos, recebendo 60/75/90/105/120% de Velocidade de Ataque por 7/7/7/7/7s.(?)
- **W · Salto-foguete** — Tristana se lança ao ar, causando 70/105/140/175/210 + 0.5 ability_power + 1 attack_damage (bonus) de Dano Mágico e @SlowMod*-100@% de Lentidão por 2/2/2/2/2s ao aterrissar. Eliminações de Campeões e detonações de Carga Explosiva com acúmulos máximos em Campeões redefinem o Tempo de Recarga desta Habilidade.(?)
- **E · Carga Explosiva** — Passivo: os Ataques de Tristana que abaterem inimigos causarão 45/60/75/90/105 + 0.25 ability_power de Dano Mágico aos inimigos ao redor. Ativo: Tristana implanta uma bomba em um inimigo ou em uma torre, causando 60/85/110/135/160 + 0.5 ability_power + 0.8 attack_damage (bonus) de Dano Físico aos inimigos ao redor após 4/4/4/4/4s. O dano aumenta em 25/25/25/25/25% a cada vez que Tristana acerta um Ataque ou Habilidade (máximo de 4 acúmulos). Com 4/4/4/4/4 acúmulos, a bomba explode imediatamente (máximo de 120/170/220/270/320 + 1 ability_power + 1.6 attack_damage (bonus) de Dano Físico).(?)
- **R · Tiro Destruidor** — Tristana dispara uma enorme bola de canhão, causando 225/275/325 + 1 ability_power + 0.7 attack_damage (bonus) de Dano Mágico ao alvo, além de Empurrar e Atordoar o alvo e os inimigos ao redor dele por 0.4/0.55/0.7s.(?)

## Warwick, a Ira Desimpedida de Zaun

`Warwick` · id 19 · corpo a corpo · fighter, tank

- **P · Fome Eterna** — Os ataques básicos de Warwick causam Dano Mágico adicional. Se ele estiver com menos de 50% de Vida, cura a si mesmo com a mesma quantia. Se ele estiver com menos de 25% de Vida, a cura é triplicada.
- **Q · Presas da Fera** — Pressionar: Warwick avança no alvo e morde, causando 0/0/0/0/0 + 1 ability_power + 1.2 attack_damage e mais 6/7/8/9/10% da Vida máxima como Dano Mágico, e curando-se em 25/37.5/50/62.5/75% do dano causado. Segurar: Warwick avança e prende as presas no alvo, saltando para trás dele. Enquanto está com as presas no alvo, Warwick segue todos os movimentos dele. Depois de soltar, ele causa o mesmo dano e cura.(?)
- **W · Caçada Sangrenta** — Passivo: Warwick pode sentir a presença de Campeões com menos de 50% de Vida, recebendo 35/42.5/50/57.5/65% de Velocidade de Movimento ao ir em direção a eles. Habilidades e Ataques contra inimigos com menos de 50% de Vida concedem 70/80/90/100/110% de Velocidade de Ataque. Esses efeitos são aumentados em 200% contra inimigos com menos de 25% de Vida. Ativo: Warwick consegue sentir brevemente a presença de todos os inimigos e recebe o efeito passivo desta Habilidade contra o Campeão mais próximo por 8s, independentemente da Vida. Se nenhum Campeão for encontrado, o Tempo de Recarga da Habilidade é reduzido em 30%.(?)
- **E · Uivo Primitivo** — Warwick recebe 35/40/45/50/55% de redução de dano por 2.75/2.75/2.75/2.75/2.75s. Quando o efeito acaba, Warwick uiva, causando Temor aos inimigos próximos por 1/1/1/1/1s. É possível Reconjurar para finalizar a Habilidade antecipadamente.(?)
- **R · Coerção Infinita** — Warwick dá um longo salto que escala com Velocidade de Movimento, Suprimindo o primeiro Campeão atingido enquanto canaliza por 1.5/1.5/1.5s. Ele Ataca esse Campeão três vezes ao longo da duração, causando 175/350/525 + 1.67 attack_damage (bonus) de Dano Mágico. Warwick cura-se em 100% de todo o dano causado durante a canalização.(?)

## Nunu e Willump, o Garoto e seu Yeti

`Nunu & Willump` · id 20 · corpo a corpo · tank, mage

- **P · Chamado de Freljord** — Nunu aumenta as Velocidades de Ataque e de Movimento de Willump e de um aliado próximo, além de fazer com que os ataques básicos de Willump causem dano a inimigos ao redor do alvo.
- **Q · Consumir** — Nunu pede a Willump que morda o inimigo, causando 400/600/800/1000/1200 de Dano Verdadeiro e restaurando 65/95/125/155/185 + 0.9 ability_power + 0.1 max_health (bonus) de Vida ao ser usado contra uma tropa ou monstro da selva. Contra um Campeão, causa 60/100/140/180/220 + 0.65 ability_power + 0.05 max_health (bonus) de Dano Mágico e restaura 39/57.000004/75/93/111.00001 + 0.54 ability_power + 0.060000002 max_health (bonus) de Vida. A cura é aumentada em 50/50/50/50/50% quando Nunu e Willump estão com menos de 50/50/50/50/50% de Vida.(?)
- **W · A Maior Bola de Neve de Todas!** — Nunu e Willump criam uma bola de neve que aumenta em velocidade e tamanho conforme eles a rolam. Eles tornam-se mais lentos enquanto preparam a bola de neve, mas a velocidade aumenta ao longo da Habilidade. Ela causa 59.940002/74.925/89.91/104.895004/119.880005 + 0.4995 ability_power - 180/225/270/315/360 + 1.5 ability_power de Dano Mágico e Arremessa ao ar por 0.5/0.5/0.5/0.5/0.5s - 1.25/1.25/1.25/1.25/1.25s ao colidir com um Campeão, monstro grande ou terreno. Esses valores escalam de acordo com a distância percorrida. Nunu e Willump podem Reconjurar para permitir que a bola de neve seja lançada antes do tempo.(?)
- **E · Rajada de Bolas de Neve** — Nunu arremessa três bolas de neve, causando 15/22.5/30/37.5/45 + 0.12 ability_power de Dano Mágico por bola de neve e @SlowAmount*-100@% de Lentidão aos inimigos atingidos três vezes por 1/1/1/1/1s. Nunu pode Reconjurar a Habilidade mais duas vezes. Após 3/3/3/3/3s, Nunu Enraíza todos os inimigos próximos que sofreram Lentidão pelas bolas de neve por 1.5/1.5/1.5/1.5/1.5s e causa 20/30/40/50/60 + 0.8 ability_power de Dano Mágico adicional.(?)
- **R · Zero Absoluto** — Nunu e Willump canalizam uma poderosa nevasca por até 3/3/3s. Inimigos dentro da área sofrem @SlowStartAmount*-100@% de Lentidão, que aumenta para @MaxSlowAmount*-100@% ao longo da duração. Nunu e Willump também recebem 65/75/85 + 1.5 ability_power + 0.3 max_health (bonus) de Escudo pela duração, antes de decair depois de 3/3/3s. Quando a nevasca termina, ela detona, causando até 625/925/1275 + 3 ability_power de Dano Mágico com base no tempo de canalização. Nunu e Willump podem Reconjurar para encerrar a nevasca antes do tempo.(?)

## Miss Fortune, a Caçadora de Recompensas

`Miss Fortune` · id 21 · a distância · marksman, mage

- **P · Batida do Amor** — Miss Fortune causa Dano Físico adicional sempre que usar um ataque básico em um novo alvo.
- **Q · Dois por Um** — Miss Fortune dispara um tiro ricocheteante, causando 20/45/70/95/120 + 0.35 ability_power + 1 attack_damage de Dano Físico ao inimigo e a outro alvo atrás dele. O segundo tiro pode causar Acerto Crítico equivalente a (?) de Dano Físico e sempre causará Acerto Crítico se o primeiro tiro abater o alvo.(?)
- **W · Desfilando** — Passivo: depois de 4/4/4/4/4s sem sofrer dano, Miss Fortune recebe 30/35/40/45/50 de Velocidade de Movimento. Depois de mais 3/3/3/3/3s, esse valor aumenta para 60/70/80/90/100. Ativo: recebe o valor total de Velocidade de Movimento adicional do efeito passivo e 40/55/70/85/100% de Velocidade de Ataque por 4/4/4/4/4s. Batida do Amor reduz o Tempo de Recarga desta Habilidade em 2/2/2/2/2s.(?)
- **E · Chuva de Disparos** — Miss Fortune dispara uma saraivada de balas, revelando uma área, causando 0.4/0.4/0.4/0.4/0.4 + 0.0006 ability_power de Lentidão e 35/50/65/80/95 + 0.6 ability_power de Dano Mágico por segundo ao longo de 2/2/2/2/2s (total de 70/100/130/160/190 + 1.2 ability_power de Dano Mágico).(?)
- **R · Metendo Bala** — Miss Fortune canaliza uma saraivada de balas, disparando 14/16/18 ondas ao longo de 3/3/3s e causando 20/30/40 + 0.25 ability_power + 0.6 attack_damage de Dano Físico por onda (total de 280/480/720 + 3.5 ability_power + 8.400001 attack_damage de Dano Físico). Cada onda pode causar Acerto Crítico equivalente a %i:scaleCrit% (?) de Dano Físico.(?)

## Ashe, a Arqueira do Gelo

`Ashe` · id 22 · a distância · marksman, support

- **P · Tiro Congelado** — Os ataques de Ashe reduzem a velocidade do alvo e fazem com que ela cause mais dano a ele. Os Acertos Críticos de Ashe não causam dano adicional, mas reduzem ainda mais a velocidade do alvo.
- **Q · Concentração** — Passivo: os Ataques de Ashe concedem um acúmulo por 4/4/4/4/4s. Ela pode ativar esta Habilidade com 4/4/4/4/4 acúmulos. Ativo: Ashe recebe 20/30/40/50/60% de Velocidade de Ataque e seus Ataques causam 0/0/0/0/0 + 1.1 attack_damage de dano por 6/6/6/6/6s.(?)
- **W · Rajada** — Ashe dispara uma rajada de 7/8/9/10/11 flechas. Cada flecha causa 60/95/130/165/200 + 1 attack_damage (bonus) de Dano Físico. Inimigos podem ser atingidos por várias flechas, mas só sofrem dano da primeira.(?)
- **E · Olhar do Falcão** — Ashe envia um falcão que concede visão de qualquer lugar do mapa por 5s. Ele também revela a área ao seu redor durante o trajeto. Esta Habilidade tem 2 cargas (90/80/70/60/50 na segunda recarga).(?)
- **R · Flecha de Cristal Encantada** — Ashe dispara uma grande flecha de cristal que Atordoa o primeiro Campeão atingido e causa 200/400/600 + 1.2 ability_power de Dano Mágico. A duração do Atordoamento aumenta com a distância percorrida, chegando a 3.5/3.5/3.5s. Inimigos ao redor sofrem Lentidão de Tiro Congelado.(?)

## Tryndamere, o Rei Bárbaro

`Tryndamere` · id 23 · corpo a corpo · fighter, assassin

- **P · Fúria da Batalha** — Tryndamere recebe Fúria para cada ataque, acerto crítico e golpe fatal que realiza. A Fúria passivamente aumenta a sua Chance de Acerto Crítico e pode ser consumida com a habilidade Sanguinário.
- **Q · Sanguinário** — Passivo: Tryndamere tem sede de sangue, recebendo até 20/35/50/65/80 de Dano de Ataque com base na própria Vida perdida. Ativo: Tryndamere consome sua Fúria, restaurando 30/40/50/60/70 + 0.3 ability_power mais 0.5/0.95/1.4/1.85/2.3 + 0.012 ability_power de Vida por Fúria (máx.: (?)). (?)
- **W · Grito Zombador** — Tryndamere profere insultos, reduzindo o Dano de Ataque de Campeões inimigos próximos em @ADReduction*-1@ por 4/4/4/4/4s e causando @SlowPotency*-100@% de Lentidão enquanto fogem dele por 3.25/3.25/3.25/3.25/3.25s.(?)
- **E · Corte Giratório** — Tryndamere gira através dos inimigos, causando 80/120/160/200/240 + 0.8 ability_power + 1 attack_damage (bonus) de Dano Físico e gerando 2/2/2/2/2 de Fúria por inimigo atingido, que aumenta para 5/5/5/5/5 de Fúria contra Campeões inimigos. O Tempo de Recarga desta Habilidade é reduzido em 0.75/0.75/0.75/0.75/0.75s sempre que Tryndamere causa um Acerto Crítico e em 1.5/1.5/1.5/1.5/1.5s quando o Acerto Crítico é em um Campeão.(?)
- **R · Fúria Sem Fim** — Tryndamere fica completamente imune à morte por 5/5/5s, recusando-se a ficar com menos de 30/50/70 de Vida e recebendo 50/75/100 de Fúria imediatamente.(?)

## Jax, o Grão-Mestre das Armas

`Jax` · id 24 · corpo a corpo · fighter

- **P · Investida Implacável** — Os ataques básicos consecutivos de Jax continuamente aumentam sua Velocidade de Ataque.
- **Q · Salto Atacante** — Jax salta até uma unidade ou sentinela, causando 65/105/145/185/225 + 1 attack_damage (bonus) de Dano Físico caso seja um inimigo.(?)
- **W · Energizar** — Jax acumula energia na arma, fazendo com que o próximo Ataque ou Salto Atacante cause 50/85/120/155/190 + 0.6 ability_power de Dano Mágico adicional.(?)
- **E · Contra-Ataque** — Jax adota uma postura defensiva por até 2/2/2/2/2s, desviando de Ataques recebidos e sofrendo 25/25/25/25/25% a menos de dano de Habilidades em área de ação. Depois de 2/2/2/2/2s ou ao Reconjurar a Habilidade, Jax causa 40/70/100/130/160 + 0.7 ability_power (bonus) + 4/4/4/4/4% da Vida máxima como Dano Mágico e Atordoa inimigos próximos por 1/1/1/1/1s. O dano é aumentado em 20/20/20/20/20% por Ataque evitado, até um máximo de 80/140/200/260/320 + 1.4 ability_power (bonus) + 4/4/4/4/4% da Vida máxima.(?)
- **R · Grão-Mestre de Armas** — Passivo: todo terceiro Ataque dentro de 2.5/2.5/2.5s causa 75/130/185 + 0.6 ability_power de Dano Mágico adicional. Ativo: Jax golpeia com a lanterna, causando 100/175/250 + 1 ability_power (bonus) de Dano Mágico aos inimigos próximos. Se acertar um Campeão, ele receberá 45/60/75 + 0.4 attack_damage (bonus) de Armadura e 27.000002/36/45 + 0.24000001 attack_damage (bonus) de Resistência Mágica, além de 20/25/30 + 0.1 attack_damage (bonus) de Armadura e 12/15.000001/18 + 0.060000002 attack_damage (bonus) de Resistência Mágica por Campeão adicional atingido por 8/8/8s. Nesse período, todo segundo Ataque, e não o terceiro, causará Dano Mágico adicional.(?)

## Morgana, a Caída

`Morgana` · id 25 · a distância · support, mage

- **P · Sifão da Alma** — Morgana drena o espírito dos inimigos, curando-se conforme causa dano a Campeões, tropas grandes e monstros médios e grandes da selva.
- **Q · Ligação das Trevas** — Morgana lança uma rajada de fogo estelar que Enraíza o primeiro inimigo atingido por 2/2.25/2.5/2.75/3s e causa 80/135/190/245/300 + 0.9 ability_power de Dano Mágico.(?)
- **W · Sombra Atormentada** — Morgana incendeia o chão por 5/5/5/5/5s, causando 18/31/44/57/70 + 0.2 ability_power de Dano Mágico por segundo, e esse valor aumenta até 36/62/88/114/140 + 0.4 ability_power com base na Vida perdida do alvo. O Tempo de Recarga da Habilidade é reduzido em 5/5/5/5/5% toda vez que Morgana é curada por Sifão da Alma.(?)
- **E · Escudo Negro** — Morgana concede 100/155/210/265/320 + 0.7 ability_power de Escudo Mágico a um Campeão aliado por 5/5/5/5/5s. O escudo protege contra efeitos Debilitantes e Imobilizadores até ser quebrado.(?)
- **R · Grilhões da Alma** — Morgana se acorrenta a Campeões inimigos próximos, causando 200/275/350 + 0.8 ability_power de Dano Mágico e 20/20/20% de Lentidão. Após 3/3/3s, inimigos que não conseguirem quebrar as correntes sofrerão 200/275/350 + 0.8 ability_power de Dano Mágico adicional e serão Atordoados por 1.5/1.75/2s. Durante a conjuração da Habilidade, Morgana recebe 20/40/60% de Velocidade de Movimento.(?)

## Zilean, o Guardião do Tempo

`Zilean` · id 26 · a distância · support, mage

- **P · Tempo Engarrafado** — Zilean armazena tempo como Experiência, podendo concedê-la aos aliados. Quando tiver Experiência o suficiente para completar o nível de um aliado, Zilean pode clicar com o botão direito nele para conceder os pontos. Zilean recebe a mesma quantidade de Experiência que concede.
- **Q · Bomba-relógio** — Zilean arremessa uma bomba temporal de acionamento tardio que se prende à primeira unidade que entrar na pequena área ao redor dela. A bomba detona após 3/3/3/3/3s, causando 75/115/165/230/300 + 0.9 ability_power de Dano Mágico. Posicionar uma segunda bomba em uma unidade já afetada por uma detona a primeira imediatamente, Atordoando inimigos no alcance da explosão por 1.1/1.2/1.3/1.4/1.5s.(?)
- **W · Retroceder** — Zilean acelera o tempo, reduzindo os Tempos de Recarga de suas outras Habilidades básicas em 10/10/10/10/10s.(?)
- **E · Distorção no Tempo** — Zilean causa 40/55/70/85/99% de Lentidão a um Campeão inimigo ou concede 40/55/70/85/99% de Velocidade de Movimento a um aliado por 2.5/2.5/2.5/2.5/2.5s.(?)
- **R · Alteração Temporal** — Zilean concede uma runa temporal protetora a um Campeão aliado por 5/5/5s. Quando o alvo morre, a runa volta no tempo, colocando-o em Estase por 3/3/3s, ressuscitando-o e restaurando 600/850/1100 + 2 ability_power de Vida.(?)

## Singed, o Químico Louco

`Singed` · id 27 · corpo a corpo · tank, mage

- **P · Corrente de Ar Nociva** — Ao passar por Campeões próximos, Singed ganha Velocidade de Movimento.
- **Q · Rastro de Veneno** — Alternar: Singed deixa um rastro venenoso que causa 20/30/40/50/60 + 0.425 ability_power de Dano Mágico por segundo.(?)
- **W · Mega Adesivo** — Singed arremessa um líquido pegajoso, Prendendo ao chão os inimigos na área e causando a eles 50/55/60/65/70% de Lentidão por 3/3/3/3/3s.(?)
- **E · Lançar** — Singed arremessa um inimigo para trás, causando 50/60/70/80/90 + 0.55 ability_power mais 6/6.5/7/7.5/8% da Vida máxima como Dano Mágico. Se Singed arremessar o alvo para seu Mega Adesivo, o alvo será Enraizado por 1/1.25/1.5/1.75/2s.(?)
- **R · Poção da Insanidade** — Singed bebe uma mistura de compostos químicos potentes que concede a ele 25/55/85 de Poder de Habilidade, Armadura, Resistência Mágica, Velocidade de Movimento, Regeneração de Vida e Regeneração de Mana por 25/25/25s. Durante o efeito, o Rastro de Veneno de Singed também aplica 40/40/40% de Feridas Dolorosas por 1/1/1s.(?)

## Evelynn, o Abraço da Agonia

`Evelynn` · id 28 · corpo a corpo · assassin, mage

- **P · Sombra Demoníaca** — Quando está fora de combate, Evelynn entra em Sombra Demoníaca. A Sombra Demoníaca cura Evelynn quando ela está com pouca Vida e concede Camuflagem após o nível 6.
- **Q · Espinho de Ódio** — Evelynn ataca com seu chicote, causando 25/30/35/40/45 + 0.25 ability_power de Dano Mágico ao primeiro inimigo atingido e fazendo com que seus próximos três Ataques ou Habilidades contra ele causem 15/25/35/45/55 + 0.25 ability_power de Dano Mágico adicional. Evelynn pode reconjurar esta Habilidade até 3/3/3/3/3x. Reconjuração: Evelynn dispara espinhos no inimigo mais próximo, causando 25/30/35/40/45 + 0.25 ability_power de Dano Mágico a todos os inimigos atingidos.(?)
- **W · Fascinação** — Evelynn marca um Campeão ou monstro por 5s. Caso Evelynn atinja o alvo com um Ataque ou Habilidade, ela expurga a marca, recupera seu custo e causa 45/45/45/45/45% de Lentidão ao alvo por 0.75/0.75/0.75/0.75/0.75s. Se a marca durar pelo menos 2,5s, expurgá-la causará efeitos adicionais:Contra Campeões: Encanta-os por 1.25/1.5/1.75/2/2.25s e remove 35/37.5/40/42.5/45% de Resistência Mágica por 4/4/4/4/4s.Contra monstros: Encanta-os por 3/3.25/3.5/3.75/4s e causa 250/300/350/400/450 + 0.6 ability_power de Dano Mágico.(?)
- **E · Chicotada** — Evelynn chicoteia o inimigo, causando 60/90/120/150/180 mais 3/3/3/3/3 + 0.015 ability_power da Vida máxima como Dano Mágico. Evelynn recebe 30/35/40/45/50% de Velocidade de Movimento por 2/2/2/2/2s. Entrar na forma Sombra Demoníaca reinicia o Tempo de Recarga da Habilidade e a fortalece. Com a Habilidade fortalecida, Evelynn avança em direção ao alvo e causa 80/120/160/200/240 mais 4/4/4/4/4 + 0.025 ability_power da Vida máxima como Dano Mágico ao alvo e a todos que ela atravessar.(?)
- **R · Última Carícia** — Evelynn libera sua energia demoníaca, causando muito dano, tornando-se Inalvejável e teleportando-se para trás. Ela causa 125/250/375 + 0.75 ability_power de Dano Mágico, aumentado para 300/600/900.00006 + 1.8000001 ability_power contra inimigos abaixo de 30% de Vida. Ao conjurar, coloca Sombra Demoníaca em um Tempo de Recarga de 1,25s.(?)

## Twitch, o Semeador da Peste

`Twitch` · id 29 · a distância · marksman, assassin

- **P · Veneno Mortal** — Os ataques básicos de Twitch infectam o alvo %i:OnHit% ao contato, causando Dano Verdadeiro a cada segundo.
- **Q · Emboscada** — Twitch entra em Camuflagem e recebe 10/10/10/10/10% de Velocidade de Movimento por 10/11/12/13/14s. A Velocidade de Movimento aumenta para 30/30/30/30/30% quando está perto de Campeões inimigos que não podem vê-lo. Após sair da Camuflagem, Twitch recebe 40/45/50/55/60% de Velocidade de Ataque por 6/6/6/6/6s. Quando um Campeão inimigo é abatido com Veneno, o Tempo de Recarga da Habilidade é redefinido.(?)
- **W · Tonel de Veneno** — Twitch arremessa um tonel que adiciona um acúmulo de Veneno Mortal a todos os inimigos atingidos, além de deixar uma nuvem tóxica que persiste por 3/3/3/3/3s. Inimigos que permanecerem na nuvem sofrem 30/35/40/45/50 + 0.06 ability_power% de Lentidão e recebem um acúmulo adicional de Veneno Mortal a cada segundo.(?)
- **E · Contaminar** — Causa 20/30/40/50/60 de Dano Físico a todos os inimigos próximos afetados por Veneno Mortal, além de 15/20/25/30/35 + 0.35 attack_damage (bonus) de Dano Físico e 0/0/0/0/0 + 0.35 ability_power de Dano Mágico adicionais por cada acúmulo de Veneno Mortal. Dano máximo: 110/150/190/230/270 + 2.1 attack_damage (bonus) de Dano Físico e 0/0/0/0/0 + 2.1 ability_power de Dano Mágico.(?)
- **R · Passando Fogo** — Twitch liberta sua balestra, recebendo 300/300/300 de Alcance de Ataque e 30/45/60 de Dano de Ataque, tornando seus Ataques perfurantes por 6/6/6s. Os dardos disparados atingem todos os inimigos pelos quais passarem, mas causam 10/10/10% a menos de dano a alvos subsequentes, chegando a um mínimo de 60/60/60%. (?)

## Karthus, a Voz Mortal

`Karthus` · id 30 · a distância · mage

- **P · Desafio da Morte** — Ao morrer, Karthus entra em forma de espírito e pode continuar conjurando habilidades.
- **Q · Devastar** — Karthus cria uma explosão de magia, causando 40/59/78/97/116 + 0.35 ability_power de Dano Mágico. Se a explosão atingir somente um inimigo, causará 80/118/156/194/232 + 0.7 ability_power de Dano Mágico.(?)
- **W · Barreira da Dor** — Karthus cria uma barreira que dura 5/5/5/5/5s. Inimigos que a atravessam perdem 25/25/25/25/25% de Resistência Mágica por 5/5/5/5/5s e sofrem 40/50/60/70/80% de Lentidão que decai ao longo da duração.(?)
- **E · Perverter** — Passivo: quando Karthus abate uma unidade, ele restaura 10/20/30/40/50 de Mana. Alternar: Karthus cria uma aura necrótica, causando 30/50/70/90/110 + 0.2 ability_power de Dano Mágico por segundo a inimigos próximos.(?)
- **R · Réquiem** — Karthus canaliza por 3s e causa 200/350/500 + 0.7 ability_power de Dano Mágico aos Campeões inimigos, independentemente da distância.(?)

## Cho'Gath, o Terror do Vazio

`Cho'Gath` · id 31 · corpo a corpo · tank, mage

- **P · Carnívoro** — Sempre que Cho'Gath abate uma unidade, ele recupera Vida e Mana. Os valores restaurados aumentam com o nível de Cho'Gath.
- **Q · Ruptura** — Cho'Gath rompe o chão e Arremessa ao ar os inimigos por (?)s, causando a eles 80/135/190/245/300 + 1 ability_power de Dano Mágico e (?)% de Lentidão por (?)s.(?)
- **W · Grito Selvagem** — Cho'Gath ruge, Silenciando os inimigos por (?)s e causando (?) de Dano Mágico.(?)
- **E · Espinhos Vorpais** — Os próximos 3 Ataques de Cho'Gath lançam espinhos que causam 20/40/60/80/100 + 0.3 ability_power mais (?) da Vida máxima do alvo como Dano Mágico e 30/35/40/45/50% de Lentidão, decaindo ao longo de 1.5/1.5/1.5/1.5/1.5s.(?)
- **R · Banquete** — Cho'Gath alimenta-se ferozmente de um inimigo, causando 300/475/650 + 0.5 ability_power + 0.1 max_health (bonus) de Dano Verdadeiro a Campeões ou 1200/1200/1200 + 0.5 ability_power + 0.1 max_health (bonus) a tropas e monstros da selva. Caso a habilidade abata o alvo, Cho'Gath recebe um acúmulo, aumentando de tamanho e ganhando 80/120/160 de Vida máxima. O máximo de acúmulos recebidos por abate de tropas e monstros não épicos da selva é de 6/6/6. (?)

## Amumu, a Múmia Triste

`Amumu` · id 32 · corpo a corpo · tank, support

- **P · Toque Amaldiçoado** — Os ataques básicos de Amumu Amaldiçoam seus inimigos, fazendo com que recebam Dano Verdadeiro adicional de qualquer Dano Mágico causado a eles.
- **Q · Lançar Bandagens** — Amumu arremessa uma bandagem, puxando a si mesmo até o primeiro inimigo atingido, Atordoando-o por 1/1/1/1/1s e causando 70/95/120/145/170 + 0.85 ability_power de Dano Mágico. Essa habilidade tem 2 cargas.(?)
- **W · Desespero** — Alternar: Amumu começa a chorar, causando 10/10/10/10/10 mais 1/1.25/1.5/1.75/2 + 0.005 ability_power% da Vida máxima como Dano Mágico a inimigos próximos por segundo e redefinindo Maldição.(?)
- **E · Chilique** — Passivo: Amumu sofre Dano Físico reduzido em 5/7/9/11/13 + 0.03 armor (bonus) + 0.03 magic_resist (bonus). Além disso, quando ele é atingido por um Ataque, o Tempo de Recarga desta habilidade é reduzido em 0.75/0.75/0.75/0.75/0.75s. Ativo: Amumu dá um chilique, causando 65/95/125/155/185 + 0.5 ability_power de Dano Mágico a inimigos próximos.(?)
- **R · A Maldição da Múmia Triste** — Amumu espalha suas bandagens, Atordoando por 1.5/1.5/1.5s, causando 200/300/400 + 0.8 ability_power de Dano Mágico e aplicando Maldição.(?)

## Rammus, o Tatu Blindado

`Rammus` · id 33 · corpo a corpo · tank

- **P · Casco Espetado** — Rammus recebe escalamento de Dano de Ataque adicional com sua Armadura e sua Resistência Mágica.
- **Q · Bola do Poder** — Rammus se encolhe em uma bola, recebendo 0.391/0.391/0.391/0.391/0.391 de Velocidade de Movimento e chegando a 2.346/2.346/2.346/2.346/2.346 de Velocidade de Movimento ao longo de 6/6/6/6/6s. Rammus para depois de colidir com um inimigo, causando 80/120/160/200/240 + 1 ability_power de Dano Mágico, Empurrando e causando 40/50/60/70/80% de Lentidão aos inimigos próximos por 1/1/1/1/1s. Reconjuração: Rammus finaliza a Habilidade antecipadamente.(?)
- **W · Bola Curva Defensiva** — Rammus assume uma posição defensiva por 7/7/7/7/7s, recebendo 35.1/44/53.649998/64.049995/75.200005 + 0.3 armor de Armadura e 26/34.375/43.5/53.375/64 + 0.3 magic_resist de Resistência Mágica, além de causar 15/15/15/15/15 + 0.1 armor + 0.1 magic_resist de Dano Mágico aos inimigos que o atacarem. Reconjuração: Rammus finaliza a Habilidade antecipadamente.(?)
- **E · Provocação Enlouquecedora** — Rammus Provoca um Campeão inimigo ou monstro por 1.2/1.4/1.6/1.8/2s. Monstros sofrem 80/100/120/140/160 + 0.7 ability_power de Dano Mágico.(?)
- **R · Colisão Ascendente** — Rammus salta e se joga com força em uma área, causando 150/250/350 + 0.6 ability_power de Dano Mágico e 30/40/50% de Lentidão por 1.5/1.5/1.5s. Se conjurada durante Bola do Poder, inimigos no centro sofrem @spell.PowerBall:PowerBallDamage@ de Dano Mágico adicional e são Arremessados ao ar por 0.75/0.75/0.75s. Depois, Rammus cria 3/3/3 tremores na área ao longo de 3.5/3.5/3.5s, repetindo a Lentidão. O alcance da Habilidade aumenta de acordo com a Velocidade de Movimento do Rammus.(?)

## Anivia, a Criofênix

`Anivia` · id 34 · a distância · mage

- **P · Renascimento** — Ao receber dano letal, Anivia volta à forma de ovo e renasce com a Vida completa.
- **Q · Lampejo Gelado** — Anivia dispara um bloco de gelo maciço, causando 50/70/90/110/130 + 0.25 ability_power de Dano Mágico, Congelando e causando @Spell.GlacialStorm:SlowAmount@ de Lentidão aos inimigos atingidos por 3/3/3/3/3s. Quando chega ao alcance máximo, o gelo explode, Atordoando os inimigos por 1.1/1.2/1.3/1.4/1.5s e causando a eles 60/95/130/165/200 + 0.45 ability_power de Dano Mágico. Ela pode Reconjurar a Habilidade durante o trajeto do gelo para detoná-lo antecipadamente.(?)
- **W · Cristalizar** — Anivia invoca um muro de gelo com 400/500/600/700/800 unidades de largura. O muro dura 5/5/5/5/5s antes de derreter.(?)
- **E · Congelamento** — Anivia atinge um inimigo com um vento congelante, causando 55/80/105/130/155 + 0.55 ability_power de Dano Mágico. Contra inimigos Congelados, ela causa 110/160/210/260/310 + 1.1 ability_power de Dano Mágico.(?)
- **R · Tempestade Glacial** — Alternar: Anivia invoca uma chuva de gelo e granizo que causa 20/30/40% de Lentidão aos inimigos e 30/45/60 + 0.125 ability_power de Dano Mágico por segundo. A tempestade aumenta de tamanho ao longo de 1.5/1.5/1.5s. No tamanho máximo, ela Congela os inimigos, causando 30/45/60% de Lentidão e 90/135/180 + 0.375 ability_power de Dano Mágico por segundo.(?)

## Shaco, o Bufão Demoníaco

`Shaco` · id 35 · corpo a corpo · assassin

- **P · Apunhalar** — Os ataques básicos de Shaco e Veneno de Dois Gumes causam dano adicional ao atingirem o inimigo por trás.
- **Q · Enganar** — Shaco se teleporta e fica Invisível por 2.5/2.75/3/3.25/3.5s. Conjurar Caixinha-Surpresa ou Alucinações não anula a Invisibilidade. O próximo Ataque de Shaco enquanto estiver Invisível causará 25/35/45/55/65 + 0.6 attack_damage (bonus) de Dano Físico adicional. Se atingir por trás, o Ataque será um Acerto Crítico que causará 0.39999998/0.39999998/0.39999998/0.39999998/0.39999998 + 0.6 critical_damage de dano.(?)
- **W · Caixinha-Surpresa** — Shaco cria uma armadilha que fica invisível depois de 2/2/2/2/2s e dura 40/40/40/40/40 + 0.1 ability_powers. Ela é ativada quando um inimigo se aproxima ou quando ela é revelada, Aterrorizando Campeões inimigos próximos por 0.5/0.75/1/1.25/1.5s ou tropas e monstros da selva por 2/2/2/2/2s. Uma vez ativada, a armadilha dispara em todos os inimigos próximos por 5 segundos, causando 10/15/20/25/30 + 0.12 ability_power de Dano Mágico, ou 25/40/55/70/85 + 0.18 ability_power de dano se focada em um único alvo. Os ataques da Caixinha-Surpresa causam 20/35/50/65/80 a mais de dano a monstros.(?)
- **E · Veneno de Dois Gumes** — Passivo: enquanto a Habilidade estiver fora do Tempo de Recarga, os Ataques de Shaco causam @SlowAmount*-100@% de Lentidão por 2/2/2/2/2s. Ativo: Shaco arremessa uma faca, causando 70/95/120/145/170 + 0.6 ability_power + 0.8 attack_damage (bonus) de Dano Mágico e @SlowAmount*-100@% de Lentidão por 3/3/3/3/3s. Se o alvo tiver menos que 30/30/30/30/30% de Vida, a faca causa 105/142.5/180/217.5/255 + 0.90000004 ability_power + 1.2 attack_damage (bonus) de dano.(?)
- **R · Alucinações** — Shaco desaparece por um instante e, então, reaparece com um clone que dura 18/18/18s e detona ao morrer, causando 150/225/300 + 0.7 ability_power de Dano Mágico e fazendo surgir três Caixinhas-Surpresa menores que são ativadas imediatamente. O clone causa 60/60/60% do dano de Shaco e sofre 50/50/50% de dano aumentado. As Caixinhas-Surpresa menores causam 10/20/30 + 0.1 ability_power de Dano Mágico, ou 25/50/75 + 0.15 ability_power de Dano Mágico se Atacarem somente um inimigo, além de infligirem Temor por 1/1/1s. (?)

## Dr. Mundo, o Louco de Zaun

`Dr. Mundo` · id 36 · corpo a corpo · tank, fighter

- **P · Vai Para Onde Quer** — Dr. Mundo resiste ao próximo efeito Imobilizador que o atingir, perdendo Vida e derrubando um recipiente químico nas proximidades. Dr. Mundo pode pegar o recipiente ao passar por cima dele, restaurando Vida e reduzindo o Tempo de Recarga da Habilidade. Dr. Mundo também tem Regeneração de Vida significativamente mais alta.
- **Q · Serra Infectada** — Dr. Mundo arremessa sua serra, causando 20/22.5/25/27.5/30% da Vida atual como Dano Mágico e 40/40/40/40/40% de Lentidão por 2/2/2/2/2s ao primeiro inimigo que atingir. Se a serra atingir um Campeão ou monstro, Dr. Mundo restaurará 50/60/70/80/90 de Vida. Se atingir um não Campeão ou não monstro, Dr. Mundo restaurará 25/30/35/40/45 de Vida.(?)
- **W · Choquinho Cardíaco** — Dr. Mundo ativa um desfibrilador, causando 20/35/50/65/80 de Dano Mágico por segundo a inimigos próximos por até 3/3/3/3/3s. Além disso, Mundo armazena 0.95/0.95/0.95/0.95/0.95 do dano sofrido nos primeiros 0.75/0.75/0.75/0.75/0.75s e 25/25/25/25/25% pelo restante da duração como Vida cinza e pode Reconjurar. Reconjuração: detona o desfibrilador, causando 20/35/50/65/80 + 0.07 max_health (bonus) de Dano Mágico a inimigos próximos. Se a Habilidade atingir ao menos um Campeão, Dr. Mundo restaura 100/100/100/100/100% de Vida cinza, caso contrário, ele restaura 50/50/50/50/50% da Vida cinza.(?)
- **E · Traumatismo** — Passivo: Dr. Mundo recebe 0/0/0/0/0 + 2 max_health de Dano de Ataque. Ativo: Dr. Mundo golpeia violentamente com a maleta "médica", fazendo com que o próximo Ataque cause 5/15/25/35/45 + 0.05 max_health (bonus) de Dano Físico adicional, aumentado em até 0.39999998/0.39999998/0.39999998/0.39999998/0.39999998 com base na Vida perdida dele. Se o inimigo for abatido, Mundo o arremessa, causando 5/15/25/35/45 + 0.05 max_health (bonus) de Dano Físico aos inimigos atravessados.(?)
- **R · Dosagem Máxima** — Dr. Mundo enche-se de produtos químicos, recebendo 15/20/25% da própria Vida perdida como Vida máxima, 15/25/35% de Velocidade de Movimento e regenerando 20/40/60% da Vida máxima durante 10/10/10s. No ranque 3, ambos os efeitos de cura são aumentados em 5/5/5% a cada Campeão inimigo próximo.(?)

## Sona, a Mestra das Cordas

`Sona` · id 37 · a distância · support, mage

- **P · Power Chord** — Accelerando: Sona recebe Aceleração permanentemente para as Habilidades básicas (exceto a ultimate), caso use-as corretamente, até um limite. Depois disso, o uso bem-sucedido reduz o Tempo de Recarga restante da ultimate. Power Chord: depois de algumas Habilidades conjuradas, o próximo Ataque de Sona causará Dano Mágico adicional e um efeito extra com base na última Habilidade básica ativada por ela.
- **Q · Hino do Valor** — Sona causa 50/85/120/155/190 + 0.4 ability_power de Dano Mágico aos dois inimigos mais próximos, priorizando Campeões. Então, ela começa uma nova Melodia. Ganha um acúmulo de Accelerando para cada Campeão a quem causar dano com esta Habilidade. Melodia: Sona recebe uma aura por 3/3/3/3/3s que concede 10/15/20/25/30 + 0.1 ability_power de Dano Mágico %i:OnHit% adicional ao próximo Ataque de Campeões aliados dentro de 5/5/5/5/5s. Power Chord – Staccato: dano adicional de Power Chord equivalente a 30/30/30/30/30 + 0.3 ability_power de Dano Mágico total.(?)
- **W · Ária da Perseverança** — Ativo: Sona restaura 30/45/60/75/90 + 0.3 ability_power de Vida para si e para um Campeão aliado próximo, priorizando o mais ferido. Então, ela começa uma nova Melodia. Melodia: por 3/3/3/3/3s, Sona recebe uma aura que concede 25/45/65/85/105 + 0.25 ability_power de Escudo aos Campeões aliados por 1.5/1.5/1.5/1.5/1.5s. Você recebe um acúmulo de Accelerando sempre que curar outro aliado ferido e sempre que proteger outro aliado de pelo menos 25/45/65/85/105 de dano com esse Escudo. Power Chord – Diminuendo: Power Chord também reduz 0.25/0.25/0.25/0.25/0.25 + 0.0004 ability_power do Dano Físico e Mágico causado pelo alvo por 3/3/3/3/3s.(?)
- **E · Canção da Celeridade** — Ativo: ela começa uma nova Melodia e concede 0.2/0.2/0.2/0.2/0.2 + 0.0002 ability_power de Velocidade de Movimento a si mesma por 3/3/3/3/3s, que estende-se em até 7/7/7/7/7s caso não sofra dano. Melodia: por 3/3/3/3/3s, Sona recebe uma aura que concede 0.1/0.12/0.14/0.16/0.18 + 0.0002 ability_power de Velocidade de Movimento aos Campeões aliados por 3/3/3/3/3s. Power Chord – Tempo: Power Chord também causa 0.5/0.5/0.5/0.5/0.5 + 0.0004 ability_power de Lentidão ao alvo por 2/2/2/2/2s.(?)
- **R · Crescendo** — Sona toca um acorde irresistível, Atordoando os Campeões inimigos por 1.5/1.5/1.5s e causando 150/250/350 + 0.5 ability_power de Dano Mágico.(?)

## Kassadin, o Andarilho do Vazio

`Kassadin` · id 38 · corpo a corpo · assassin, mage

- **P · Pedra do Vazio** — Kassadin sofre Dano Mágico reduzido e ignora colisão com unidades.
- **Q · Esfera Nula** — Kassadin dispara um orbe de energia do Vazio, causando 65/95/125/155/185 + 0.7 ability_power de Dano Mágico e interrompendo canalizações. Além disso, ele recebe 80/110/140/170/200 + 0.3 ability_power de Escudo contra Dano Mágico por 1,5s.(?)
- **W · Lâmina Ínfera** — Passivo: os Ataques de Kassadin causam 25/25/25/25/25 + 0.1 ability_power de Dano Mágico adicional. Ativo: Kassadin carrega sua lâmina, fazendo com que seu próximo Ataque cause 50/75/100/125/150 + 0.8 ability_power de Dano Mágico e restaure 4/4.5/5/5.5/6% do Mana perdido. Aumenta para 20/22.5/25/27.5/30% contra Campeões.(?)
- **E · Força de Pulso** — Passivo: o Tempo de Recarga de Força de Pulso é reduzido em 0.75/0.75/0.75/0.75/0.75s sempre que uma habilidade é usada perto de Kassadin. Ativo: Kassadin libera um pulso do Vazio, causando 70/100/130/160/190 + 0.7 ability_power de Dano Mágico e 50/55/60/65/70% de Lentidão por 1/1/1/1/1s.(?)
- **R · Caminhar na Fenda** — Kassadin se teleporta para um local próximo, causando 70/90/110 + 0.5 ability_power + 0.02 mana de Dano Mágico. Cada uso subsequente da Habilidade dentro dos próximos 15/15/15s dobra o custo de Mana e causa 35/45/55 + 0.07 ability_power + 0.01 mana de Dano Mágico adicional. Os aumentos de custo e dano podem acumular até 4/4/4 vezes.(?)

## Irelia, a Dançarina das Lâminas

`Irelia` · id 39 · corpo a corpo · fighter, assassin

- **P · Fervor Ioniano** — Quando Irelia acerta inimigos com suas habilidades, ela ganha acúmulo de Velocidade de Ataque adicional. Ao atingir o máximo de acúmulos, Irelia também ganha dano adicional ao contato.
- **Q · Surto da Lâmina** — Irelia avança até um inimigo, causando 5/25/45/65/85 + 0.7 attack_damage de Dano Físico e restaurando 0/0/0/0/0 + 0.09 attack_damage de Vida. Se o inimigo morrer ou estiver Instável, o Tempo de Recarga é redefinido. Causa 55/75/95/115/135 + 0.7 attack_damage de dano a tropas.(?)
- **W · Dança Desafiadora** — Início do carregamento: Irelia assume uma postura defensiva por 1.5/1.5/1.5/1.5/1.5s, tornando-se incapaz de realizar ações e reduzindo o Dano Físico sofrido em 70/70/70/70/70 + 0.07 ability_power% e o Dano Mágico sofrido em 35/35/35/35/35 + 0.035 ability_power%. Liberar: Irelia brande suas lâminas, causando 10/20/30/40/50 + 0.5 ability_power + 0.4 attack_damage de Dano Físico, aumentando até 30/60/90/120/150 + 1.5 ability_power + 1.2 attack_damage com base no tempo de carregamento.(?)
- **E · Dueto Impecável** — Irelia arremessa uma lâmina no chão e pode Reconjurar em até 3.5/3.5/3.5/3.5/3.5s. Reconjuração: Irelia arremessa uma segunda lâmina. Depois, as duas lâminas se encontram, Atordoando por 0.75/0.75/0.75/0.75/0.75s e causando 70/110/150/190/230 + 1 ability_power de Dano Mágico. Campeões e monstros grandes da selva ficam Instáveis por 5/5/5/5/5s.(?)
- **R · Lâmina da Vanguarda** — Irelia lança uma barragem de lâminas, causando 125/200/275 + 0.7 ability_power de Dano Mágico e deixando Campeões e monstros grandes da selva Instáveis por 5/5/5s. A barragem explode, criando uma prisão em volta do primeiro Campeão atingido por 2.5/2.5/2.5s. A prisão causa 125/200/275 + 0.7 ability_power de Dano Mágico e 90/90/90% de Lentidão por 1.5/1.5/1.5s.(?)

## Janna, a Fúria da Tormenta

`Janna` · id 40 · a distância · support, mage

- **P · Brisa de Impulso** — Os aliados de Janna recebem Velocidade de Movimento ao se moverem em direção a ela. Janna causa uma parte da Velocidade de Movimento adicional como Dano Mágico adicional ao contato com Zéfiro.
- **Q · Ventania Uivante** — Janna invoca um tornado que carrega por 3/3/3/3/3s e, depois, segue um determinado trajeto. Ele causa 55/90/125/160/195 + 0.5 ability_power - 85/135/185/235/285 + 0.8 ability_power de Dano Mágico e Arremessa ao ar por 0.5/0.5/0.5/0.5/0.5s - 1.25/1.25/1.25/1.25/1.25s. A distância, o dano e a duração do Arremesso ao ar aumentam com base no carregamento do tornado. Janna pode Reconjurar para enviar o tornado antecipadamente.(?)
- **W · Zéfiro** — Passivo: Janna recebe 0.06/0.07/0.08/0.09/0.1 + 0.0002 ability_power de Velocidade de Movimento e efeito Fantasma. Ativo: o elemental de Janna ataca um inimigo, causando 20/24/28/32/36 + 0.06 ability_power de Lentidão por 2/2/2/2/2s e 55/85/115/145/175 + 0.5 ability_power + @spell.TailwindSelf:BonusDamage@ de Dano Mágico.(?)
- **E · Olho da Tempestade** — Janna concede a um Campeão aliado ou uma torre aliada 80/120/160/200/240 + 0.55 ability_power de Escudo por 4/4/4/4/4s. Enquanto estiver com o Escudo, o alvo recebe 10/15/20/25/30 + 0.1 ability_power de Dano de Ataque. Janna restitui 20/20/20/20/20% do Tempo de Recarga sempre que debilitar a movimentação de um Campeão inimigo com uma Habilidade.(?)
- **R · Monção** — Janna invoca uma monção mágica, Empurrando inimigos próximos e restaurando 300/450/600 + 1.5 ability_power de Vida de aliados próximos ao longo de 3/3/3s. Mover-se ou usar uma Habilidade encerra a monção antecipadamente.(?)

## Gangplank, o Terror dos Doze Mares

`Gangplank` · id 41 · corpo a corpo · fighter

- **P · Julgamento de Fogo** — Periodicamente, o ataque corpo a corpo de Gangplank deixará seu inimigo em chamas.
- **Q · Negociarrr** — {{Spell_GangplankQWrapper_Tooltip_1/1/1/1/1}}(?)
- **W · Remover Escorbuto** — Gangplank consome uma grande quantidade de frutas cítricas, removendo todos os efeitos Debilitantes e se curando em 45/70/95/120/145 + 0.9 ability_power mais 13/13/13/13/13% da Vida perdida.(?)
- **E · Barril de Pólvora** — Gangplank posiciona um Barril de Pólvora que pode ser atacado por ele e por Campeões inimigos por 25/25/25/25/25s. Quando um inimigo o destrói, o barril é desativado. Quando Gangplank o destrói, ele explode, causando 40/50/60/70/80% de Lentidão por 2/2/2/2/2s e o dano do Ataque, ignorando 40/40/40/40/40% de Armadura. Campeões sofrem 75/95/115/135/155 de Dano Físico adicional. A Vida do barril decai a cada 0.5/0.5/0.5/0.5/0.5s. As explosões dos barris detonam outros barris com as zonas de explosão sobrepostas, mas não causam dano ao mesmo alvo mais de uma vez. Explosões ativadas por Negociarrr concedem ouro adicional pelos alvos abatidos.(?)
- **R · Barragem de Canhão** — Gangplank sinaliza ao seu navio para disparar 12/12/12 salvas de balas de canhão em qualquer lugar do mapa ao longo de 8/8/8s. Cada salva causa 30/30/30% de Lentidão por 0.5/0.5/0.5s e 40/70/100 + 0.1 ability_power de Dano Mágico. Dano máximo: 480/840/1200 + 1.2 ability_power Esta habilidade pode ser aprimorada na loja via Negociarrr. Disparar à Vontade: dispara 6 salvas adicionais de balas de canhão. Filha da Morte: dispara uma Bala de Canhão Gigante que causa 120/210/300 + 0.3 ability_power de Dano Verdadeiro e 75/75/75% de Lentidão por 1/1/1s. Levantar o Moral: aliados dentro da Barragem de Canhão recebem 40/40/40% de Velocidade de Movimento por 2/2/2s.(?)

## Corki, o Bombardeiro Ousado

`Corki` · id 42 · a distância · marksman, mage

- **P · Munição Hextec** — Uma porcentagem do Dano de Ataque Básico de Corki é causada como Dano Verdadeiro adicional.
- **Q · Bomba de Fósforo** — Corki lança uma bomba, causando 60/105/150/195/240 + 1 ability_power + 1.25 attack_damage (bonus) de Dano Mágico. A área e os Campeões atingidos são revelados por 6/6/6/6/6s.(?)
- **W · Valquíria** — Corki sobrevoa e incendeia um trajeto, fazendo com que queime por 2.5/2.5/2.5/2.5/2.5s. Inimigos na área sofrem até 150/225/300/375/450 + 1.5 ability_power + 2 attack_damage (bonus) de Dano Mágico ao longo da duração.(?)
- **E · Metralhadora** — Corki atira à frente com uma metralhadora, causando 80/130/180/230/280 + 2.4 attack_damage (bonus) de Dano Físico ao longo de 4/4/4/4/4s e fragmentando até @ShredMax*-1@ de Armadura e Resistência Mágica.(?)
- **R · Barragem de Mísseis** — Corki dispara um míssil que explode ao atingir o primeiro inimigo, causando 90/170/250 + 0.85 attack_damage (bonus) de Dano Físico aos inimigos ao redor. Cada terceiro míssil causa 180/340/500 + 1.7 attack_damage (bonus) de Dano Físico. A Habilidade tem até 4/4/4 cargas. Ataques básicos contra Campeões reduzem o tempo entre as cargas em 1/1/1 + 2 critical_chances ao contato.(?)

## Karma, a Iluminada

`Karma` · id 43 · a distância · mage, support

- **P · Ímpeto Ardente** — As habilidades de dano de Karma reduzem o Tempo de Recarga de Mantra.
- **Q · Chama Interior** — Karma dispara uma explosão de energia, causando 60/110/160/210/260 + 0.7 ability_power de Dano Mágico ao primeiro alvo atingido e aos inimigos ao redor, além de @SlowAmount*-100@% de Lentidão por 1.5/1.5/1.5/1.5/1.5s.(?)
- **W · Decisão Absorta** — Karma se vincula a um Campeão inimigo ou monstro da selva, causando 40/65/90/115/140 + 0.45 ability_power de Dano Mágico e revelando o alvo por 2/2/2/2/2s. Se o vínculo não for quebrado, o alvo sofre 40/65/90/115/140 + 0.45 ability_power de Dano Mágico novamente e é Enraizado por 1.6/1.7/1.8/1.9/2s.(?)
- **E · Inspiração** — Karma concede a um Campeão aliado 80/130/180/230/280 + 0.6 ability_power de Escudo por 2.5/2.5/2.5/2.5/2.5s e 40/40/40/40/40% de Velocidade de Movimento por 2/2/2/2/2s.(?)
- **R · Mantra** — Karma fortalece sua próxima Habilidade conjurada dentro de 8s. Chama Interior: causa 40/100/160 + 0.3 ability_power de Dano Mágico adicional e deixa para trás um círculo de fogo, causando Lentidão e 40/130/220 + 0.5 ability_power de Dano Mágico adicional aos inimigos.Decisão Absorta: Karma restaura 17/17/17 + 0.01 ability_power da Vida perdida no começo e ao final do vínculo, Enraizando por mais 0.5/0.75/1s.Inspiração: Karma concede 45/85/125 + 0.45 ability_power a mais de Escudo ao alvo, bem como Escudo aos aliados próximos ao alvo, concedendo 45/85/125 + 0.45 ability_power de Escudo e 15/15/15% de Velocidade de Movimento.(?)

## Taric, o Escudo de Valoran

`Taric` · id 44 · corpo a corpo · support, tank

- **P · Bravata** — Conjurações fortalecem os próximos 2 ataques básicos de Taric, causando Dano Mágico adicional, reduzindo Tempos de Recarga e atacando em rápida sucessão.
- **Q · Fulgor Estelar** — Passivo: recebe um acúmulo (máximo de 1/2/3/4/5) a cada 15/15/15/15/15s e quando acertar um Ataque com Bravata. Ativo: consome todos os acúmulos para restaurar 25/25/25/25/25 + 0.15 ability_power + 0.01 max_health de Vida por acúmulo a todos os Campeões aliados próximos (25/25/25/25/25 + 0.15 ability_power + 0.01 max_health com 1/2/3/4/5 acúmulos).(?)
- **W · Bastião** — Passivo: Taric recebe 0/0/0/0/0 + 0.06 armor de Armadura e forma uma ligação entre ele e o aliado escolhido. Enquanto permanecerem próximos, o aliado vinculado recebe 0/0/0/0/0 + 0.06 armor de Armadura e as Habilidades de Taric também são conjuradas pelo aliado vinculado. Ativo: Taric se liga a um Campeão aliado, concedendo Escudo equivalente a 7/8/9/10/11% da Vida máxima por 2.5/2.5/2.5/2.5/2.5s.(?)
- **E · Deslumbrar** — Taric projeta um feixe de luz estelar que explode depois de 1/1/1/1/1s, causando 90/130/170/210/250 + 0.5 ability_power + 0.5 armor (bonus) de Dano Mágico e Atordoando inimigos por 1.5/1.5/1.5/1.5/1.5s.(?)
- **R · Resplendor Cósmico** — Taric conjura a proteção vinda dos céus. Depois de 2.5/2.5/2.5s, Campeões aliados próximos tornam-se Invulneráveis por 2.5/2.5/2.5s.(?)

## Veigar, o Pequeno Mestre do Mal

`Veigar` · id 45 · a distância · mage

- **P · Poder Maligno Fenomenal** — Veigar é o maior ser maligno a chegar ao coração de Runeterra - e vai ficar cada vez maior! Atingir um Campeão inimigo com uma habilidade ou executar um golpe final concede a Veigar um aumento permanente de Poder de Habilidade.
- **Q · Golpe Maligno** — Veigar solta um disparo de energia das trevas que causa 80/120/160/200/240 + 0.5 ability_power de Dano Mágico aos dois primeiros inimigos atingidos. Caso Veigar abata um inimigo com esta Habilidade, ele recebe @Spell.VeigarPassive:dQKillStacks@ acúmulo(s) de Poder Maligno Fenomenal. Tropas grandes e monstros grandes concedem @Spell.VeigarPassive:dQKillStacksLarge@ acúmulos.(?)
- **W · Matéria Escura** — Veigar invoca matéria escura do céu, causando 85/140/195/250/305 + 0.7 ability_power de Dano Mágico. A cada @Spell.VeigarPassive:PStacksPerDarkMatterCDR@ acúmulos de Poder Maligno Fenomenal, o Tempo de Recarga desta Habilidade é reduzido em @Spell.VeigarPassive:DarkMatterCDRIncrement*100@%.(?)
- **E · Horizonte de Eventos** — Veigar distorce os limites do espaço, criando uma prisão que Atordoa por 1.5/1.75/2/2.25/2.5s os inimigos que a atravessam. A prisão dura 3s.(?)
- **R · Explosão Primordial** — Veigar dispara uma magia poderosa em um Campeão inimigo e causa 175/250/325 + 0.65 ability_power - 350/500/650 + 1.3 ability_power de Dano Mágico, aumentando de acordo com a Vida perdida do alvo. O dano chega ao valor máximo contra inimigos com menos de 33% de Vida.(?)

## Trundle, o Rei dos Trolls

`Trundle` · id 48 · corpo a corpo · fighter, tank

- **P · Tributo do Rei** — Quando uma unidade inimiga morre perto de Trundle, ele é curado em um percentual da Vida máxima desta unidade.
- **Q · Mordida** — O próximo Ataque de Trundle causa 10/30/50/70/90 + 1.15 attack_damage de Dano Físico e 75/75/75/75/75% de Lentidão. Depois, Trundle recebe 20/25/30/35/40 de Dano de Ataque por 5/5/5/5/5s e o inimigo perde @SappedAD*-1@ de Dano de Ataque pela mesma duração.(?)
- **W · Domínio Congelado** — Trundle congela uma área por 8/8/8/8/8s. Enquanto está dentro dela, ele recebe 20/28/36/44/52% de Velocidade de Movimento, 30/45/60/75/90% de Velocidade de Ataque e 25/25/25/25/25% a mais de cura.(?)
- **E · Pilar de Gelo** — Trundle cria um pilar de gelo por 6/6/6/6/6s, Empurrando os inimigos brevemente por cima dele e causando 34/38/42/46/50% de Lentidão aos inimigos próximos.(?)
- **R · Subjugar** — Trundle drena a força vital de um Campeão inimigo, causando 0.2/0.25/0.3 + 0.0002 ability_power da Vida máxima como Dano Mágico e roubando 40/40/40% de Armadura e Resistência Mágica ao longo de 5/5/5s.(?)

## Swain, o Grande General Noxiano

`Swain` · id 50 · a distância · mage, support

- **P · Bando Voraz** — Os corvos de Swain coletam Fragmentos de Alma que o curam e aumentam permanentemente a Vida máxima dele.
- **Q · Mão da Morte** — Swain libera 5 raios sombrios, causando 60/90/120/150/180 + 0.45 ability_power de Dano Mágico mais 15/22.5/30/37.5/45 + 0.1125 ability_power de Dano Mágico por cada raio depois do primeiro (máximo de 120/180/240/300/360 + 0.9 ability_power de Dano Mágico).(?)
- **W · Visão do Império** — Swain abre um olho demoníaco que revela um local por 1,5s e causa 70/105/140/175/210 + 0.6 ability_power de Dano Mágico e @Slow*-100@% de Lentidão por 1.5/1.5/1.5/1.5/1.5s. Campeões atingidos concedem um Fragmento de Alma a Swain, além de serem revelados por 6/6/6/6/6s.(?)
- **E · Nuncamova** — Swain lança uma onda de poder demoníaco que retorna, detonando no primeiro inimigo atingido, causando 90/130/170/210/250 + 0.7 ability_power de Dano Mágico e Enraizando inimigos na área por 1.5/1.5/1.5/1.5/1.5s. Enraizar um Campeão permite que Swain reconjure a Habilidade, puxando todos os Campeões Enraizados por Nuncamova para si e recebendo um Fragmento de Alma para cada um deles.(?)
- **R · Ascensão Demoníaca** — Swain libera o demônio, causando 15/25/35 + 0.04 ability_power de Dano Mágico e drenando 15/30/45 + 0.05 ability_power + 0.015 max_health (bonus) de Vida de inimigos próximos por segundo. Sua Energia Demoníaca decai ao longo do tempo, mas pode ser recarregada totalmente ao eliminar Campeões inimigos. Após 2/2/2s e a cada 8/8/8s depois disso, Swain pode conjurar Chama Demoníaca durante a transformação, causando 150/250/350 + 0.4 ability_power de Dano Mágico e 50/50/50% de Lentidão aos inimigos, que decai ao longo de 1.5/1.5/1.5s.(?)

## Caitlyn, a Xerife de Piltover

`Caitlyn` · id 51 · a distância · marksman

- **P · Bem na Mira** — A cada poucos ataques básicos, ou contra um alvo preso em uma armadilha ou rede, Caitlyn fará um disparo Bem na Mira causando dano adicional que escala com sua Chance de Acerto Crítico. Em alvos presos por armadilhas ou pela rede, o alcance de Bem na Mira de Caitlyn é dobrado.
- **Q · Pacificadora de Piltover** — Caitlyn prepara seu rifle para disparar um projétil perfurante que causa 50/90/130/170/210 + 1.25 attack_damage de Dano Físico. Depois de atingir o primeiro alvo, o projétil se abre e fica mais largo, causando 30.000002/54.000004/78/102.00001/126.00001 + 0.75 attack_damage de Dano Físico. Inimigos revelados por Armadilha Mecânica Yordle sempre sofrem dano total.(?)
- **W · Armadilha Mecânica Yordle** — Caitlyn posiciona uma armadilha que Enraíza o primeiro Campeão que pisar nela por 1.5/1.5/1.5/1.5/1.5s, concedendo Visão Mágica dele por 3s. Elas duram 30/35/40/45/50s e até 3/3/4/4/5 Armadilhas podem estar ativas ao mesmo tempo. A Habilidade tem 3/3/4/4/5 cargas (recarrega a cada (?)s). Alvos enraizados por esta Habilidade sofrem mais 35/80/125/170/215 + 0.3 attack_damage (bonus) de Dano Físico de Bem na Mira.(?)
- **E · Rede Calibre 90** — Caitlyn dispara uma rede que a empurra para trás. A rede causa 50/50/50/50/50% de Lentidão por 1/1/1/1/1s e 80/130/180/230/280 + 0.8 ability_power de Dano Mágico ao primeiro alvo atingido.(?)
- **R · Ás na Manga** — Caitlyn canaliza brevemente para preparar o disparo perfeito e depois atira, causando 300/475/650 + 1 attack_damage (bonus) de Dano Físico, mas outros Campeões inimigos podem bloquear o disparo. A Habilidade concede Visão Mágica do alvo durante a canalização. O dano escala com a Chance de Acerto Crítico e o Dano de Acerto Crítico de Caitlyn.(?)

## Blitzcrank, o Grande Golem a Vapor

`Blitzcrank` · id 53 · corpo a corpo · tank, support

- **P · Barreira de Mana** — Blitzcrank recebe um Escudo com base em seu Mana ao ficar com a Vida baixa.
- **Q · Puxão Biônico** — Blitzcrank dispara a mão direita, Puxando o primeiro inimigo atingido em sua direção e causando 110/160/210/260/310 + 1.2 ability_power de Dano Mágico.(?)
- **W · Turbo** — Blitzcrank sobrecarrega-se, recebendo 60/65/70/75/80% de Velocidade de Movimento que decai e 30/40/50/60/70% de Velocidade de Ataque por 5/5/5/5/5s. Depois disso, Blitzcrank sofre 30/30/30/30/30% de Lentidão por 1.5/1.5/1.5/1.5/1.5s.(?)
- **E · Punho do Poder** — Blitzcrank carrega o punho, fazendo com que seu próximo Ataque Arremesse ao ar por 1/1/1/1/1s e cause 0/0/0/0/0 + 0.25 ability_power + 2 attack_damage de Dano Físico.(?)
- **R · Campo Estático** — Passivo: enquanto esta habilidade estiver disponível, os punhos de Blitzcrank serão carregados com eletricidade, marcando quem for Atacado. Após 1s, os marcados sofrem um choque elétrico que causa 50/100/150 + 0.3 ability_power + 0.02 mana de Dano Mágico. Ativo: Blitzcrank se sobrecarrega, causando 275/400/525 + 1 ability_power de Dano Mágico e Silenciando Campeões inimigos próximos por 0.5/0.5/0.5s. Os Escudos inimigos também são destruídos.(?)

## Malphite, o Fragmento do Monolito

`Malphite` · id 54 · corpo a corpo · tank, mage

- **P · Escudo de Granito** — Malphite está protegido por uma camada de rocha que absorve dano equivalente a até 10% de sua Vida máxima. Se Malphite não for atingido por alguns segundos, o efeito é recarregado.
- **Q · Fragmento Sísmico** — Malphite arremessa um fragmento de terra em um inimigo, causando 70/120/170/220/270 + 0.6 ability_power de Dano Mágico e 20/25/30/35/40% de Lentidão por 3/3/3/3/3s. Malphite também rouba o valor da Lentidão causada, recebendo Velocidade de Movimento por 3/3/3/3/3s.(?)
- **W · Trovoada** — Passivo: Malphite recebe 10/15/20/25/30% de Armadura (%i:scaleArmor%(?)). Esse efeito é aumentado para 30/45/60/75/90% (%i:scaleArmor%(?)) enquanto Escudo de Granito estiver ativo. Ativo: o próximo Ataque de Malphite causa 30/40/50/60/70 + 0.2 ability_power + 0.15 armor de Dano Físico adicional e cria uma onda de choque que causa 15/25/35/45/55 + 0.3 ability_power + 0.15 armor de Dano Físico %i:OnHit% ao contato na direção alvejada. Os Ataques continuam criando ondas de choque %i:OnHit% ao contato pelos próximos 5/5/5/5/5s.(?)
- **E · Estrondar Terreno** — Malphite golpeia o chão, causando 60/95/130/165/200 + 0.6 ability_power + 0.4 armor de Dano Mágico e reduzindo a Velocidade de Ataque dos inimigos atingidos em 30/35/40/45/50% por 3/3/3/3/3s.(?)
- **R · Força Incontrolável** — Malphite avança com a força de um deslizamento inabalável. No fim do avanço, os inimigos atingidos são Arremessados ao ar por 1.5/1.5/1.5s e sofrem 200/300/400 + 0.9 ability_power de Dano Mágico.(?)

## Katarina, a Lâmina Sinistra

`Katarina` · id 55 · corpo a corpo · assassin, mage

- **P · Voracidade** — Os Tempos de Recarga de Katarina são reduzidos drasticamente sempre que um Campeão que ela havia causado dano recentemente morrer. Se apanhar uma Adaga, Katarina a usa para cortar todos os inimigos próximos, causando Dano Mágico.
- **Q · Lâmina Saltitante** — Katarina atira uma Adaga, causando 80/115/150/185/220 + 0.4 ability_power de Dano Mágico ao alvo e 2/2/2/2/2 a inimigos próximos. Depois, a Adaga ricocheteia e cai no chão atrás do alvo principal.(?)
- **W · Preparação** — Katarina joga uma Adaga para cima e recebe 50/60/70/80/90% de Velocidade de Movimento que decai ao longo de 1.25/1.25/1.25/1.25/1.25s.(?)
- **E · Shunpo** — Katarina se teleporta em um piscar de olhos até o aliado, inimigo ou Adaga que tiver como alvo. Caso seja um inimigo, Katarina causará a ele 20/30/40/50/60 + 0.25 ability_power + 0.4 attack_damage de Dano Mágico. Caso contrário, ela atingirá o inimigo mais próximo que estiver ao seu alcance. Apanhar uma Adaga reduzirá o Tempo de Recarga da Habilidade em 11.5199995/10.559999/9.599999/8.639999/7.68s (0.96/0.96/0.96/0.96/0.96). Katarina pode se teleportar para qualquer local próximo ao alvo.(?)
- **R · Lótus da Morte** — Katarina se transforma em um turbilhão de lâminas, atacando rapidamente os três Campeões inimigos mais próximos. Cada faca causa 25/37.5/50 + 0.19 ability_power de Dano Mágico, (?) de Dano Físico e 40/40/40% de Feridas Dolorosas por 3/3/3s. Total de dano causado a cada inimigo ao longo de 2.5/2.5/2.5s: 375/562.5/750 + 2.85 ability_power de Dano Mágico e (?) de Dano Físico.(?)

## Nocturne, o Eterno Pesadelo

`Nocturne` · id 56 · corpo a corpo · fighter, assassin

- **P · Lâminas Sombrias** — A cada poucos segundos, o próximo ataque de Nocturne atinge os inimigos ao seu redor para causar Dano Físico adicional e curá-lo. Os ataques básicos de Nocturne reduzem esse Tempo de Recarga.
- **Q · Portador do Anoitecer** — Nocturne arremessa uma lâmina de sombras, causando 65/105/145/185/225 + 0.85 attack_damage (bonus) de Dano Físico e deixando um rastro de crepúsculo por 5/5/5/5/5s. Campeões inimigos atingidos também deixam um rastro. Enquanto está no rastro, Nocturne recebe efeito Fantasma, 20/25/30/35/40% de Velocidade de Movimento e 15/25/35/45/55 de Dano de Ataque.(?)
- **W · Proteção das Trevas** — Passivo: Nocturne recebe 30/35/40/45/50% de Velocidade de Ataque. Ativo: Nocturne cria uma barreira sombria por 1,5s, bloqueando a próxima Habilidade inimiga. Se uma Habilidade for bloqueada, o efeito passivo é aumentado para 60/70/80/90/100% de Velocidade de Ataque por 5/5/5/5/5s.(?)
- **E · Horror Indescritível** — Passivo: Nocturne recebe 90/90/90/90/90% de Velocidade de Movimento ao se mover em direção a inimigos que estão sob o efeito de Temor. Ativo: Nocturne cria um vínculo de pesadelo com o alvo, causando 80/125/170/215/260 + 1 ability_power de Dano Mágico ao longo de 2/2/2/2/2s. Se o vínculo não for quebrado, o alvo sofre Temor por 1.25/1.5/1.75/2/2.25s.(?)
- **R · Paranoia** — Nocturne escurece o mapa, reduzindo o raio de visão de todos os Campeões inimigos e impedindo que vejam os aliados por 6/6/6s. Reconjurar a Habilidade ao longo da duração faz com que Nocturne se arremesse em um Campeão inimigo, causando 150/275/400 + 1.2 attack_damage (bonus) de Dano Físico.(?)

## Maokai, o Ente Sinistro

`Maokai` · id 57 · corpo a corpo · tank, support

- **P · Seiva Mágica** — O ataque básico de Maokai também concede cura a ele e causa dano adicional com um Tempo de Recarga moderado. Sempre que ele conjurar uma habilidade ou for atingido por uma habilidade inimiga, o Tempo de Recarga da cura é reduzido.
- **Q · Esmagamento Espinhoso** — Maokai soca o punho no chão, causando 75/120/165/210/255 + 0.5 ability_power mais 2/2.5/3/3.5/4% da Vida Máxima como Dano Mágico e breve Lentidão. Inimigos próximos também são Empurrados.(?)
- **W · Avanço Retorcido** — Maokai se transforma em uma massa ambulante de raízes, tornando-se Inalvejável e avançando até um inimigo. Ao chegar, ele Enraíza o inimigo por 1/1.1/1.2/1.3/1.4s e causa 60/85/110/135/160 + 0.4 ability_power de Dano Mágico. (?)
- **E · Atirar Mudas** — Maokai atira uma muda que fica de guarda por 30/30/30/30/30s. Mudas perseguem inimigos próximos e detonam ao chegar, causando 50/75/100/125/150 + 0.25 ability_power + 0.05 max_health (bonus) de Dano Mágico e 45/45/45/45/45% de Lentidão aos inimigos ao redor por 2/2/2/2/2s. Se a muda atingir um Campeão inimigo ou um monstro épico, o Tempo de Recarga de Seiva Mágica é reduzido em mais 4s. Se forem plantadas em arbustos, as Mudas duram 30/30/30/30/30 + 0.015 max_health (bonus)s e geram uma explosão maior, causando 100/150/200/250/300 + 0.5 ability_power + 0.1 max_health (bonus) de Dano Mágico ao longo de 2/2/2/2/2s e 0.45/0.45/0.45/0.45/0.45 + 0.0001 ability_power + 0.0001 max_health (bonus) de Lentidão aos inimigos.(?)
- **R · Garras da Natureza** — Maokai invoca uma onda gigantesca de espinheiros, causando 150/225/300 + 0.75 ability_power de Dano Mágico e Enraizando inimigos por 0.75/0.75/0.75s - 2.25/2.25/2.25s, dependendo da distância percorrida. Atingir um Campeão concede a Maokai 40/50/60% de Velocidade de Movimento, que decai ao longo de 2/2/2s. (?)

## Renekton, o Carniceiro das Areias

`Renekton` · id 58 · corpo a corpo · fighter, tank

- **P · Domínio da Ira** — Os ataques de Renekton geram Fúria, que aumenta quando ele estiver com a Vida baixa. Essa Fúria pode fortalecer suas habilidades com efeitos adicionais.
- **Q · Abater os Indefesos** — Renekton brande sua lâmina, causando 60/90/120/150/180 + 1 attack_damage (bonus) de Dano Físico e restaurando 2/3/4/5/6 + 0.02 attack_damage (bonus) de Vida para cada não Campeão e 12/20/28/36/44 + 0.17 attack_damage (bonus) por Campeão atingido. Ele também gera 2.5/2.5/2.5/2.5/2.5 de Fúria para cada não Campeão e 10/10/10/10/10 de Fúria por Campeão atingido. Efeito de Fúria: o dano aumenta para 90/135/180/225/270 + 1.4 attack_damage (bonus) de Dano Físico e a cura aumenta para 6/9/12/15/18 + 0.06 attack_damage (bonus) de Vida contra não Campeões e 36/60/84/108/132 + 0.51 attack_damage (bonus) de Vida contra Campeões. Não gera Fúria.(?)
- **W · Predador Desumano** — O próximo Ataque de Renekton acerta duas vezes, Atordoando o inimigo por 0.75/0.75/0.75/0.75/0.75s e causando 10/40/70/100/130 + 1.5 attack_damage de Dano Físico. Atingir um Campeão gera 10/10/10/10/10 de Fúria adicional. Efeito de Fúria: Renekton Ataca 3 vezes, destruindo Escudos antes de causar 15/60/105/150/195 + 2.25 attack_damage de Dano Físico e Atordoar por 1.5/1.5/1.5/1.5/1.5s. Não gera Fúria.(?)
- **E · Fatiar e Picar** — Renekton avança, causando 40/70/100/130/160 + 0.9 attack_damage (bonus) de Dano Físico. Ele gera 2/2/2/2/2 de Fúria para cada não Campeão e 10/10/10/10/10 de Fúria por Campeão atingido. Atingir ao menos um inimigo permite que Renekton Reconjure a Habilidade uma vez por 4/4/4/4/4s. Efeito de Fúria: o avanço Reconjurado causa 70/115/160/205/250 + 1.35 attack_damage (bonus) de Dano Físico e remove 25/27.5/30/32.5/35% da Armadura por 4/4/4/4/4s. Não gera Fúria.(?)
- **R · Dominus** — Renekton envolve-se com energias sombrias por 15/15/15s, recebendo 300/500/700 de Vida máxima e 20/20/20 de Fúria. Enquanto estiver ativo, ele causa 60/150/240 + 0.1 ability_power + 0.1 attack_damage (bonus) de Dano Mágico e recebe 5/5/5 de Fúria por segundo.(?)

## Jarvan IV, o Exemplo de Demacia

`Jarvan IV` · id 59 · corpo a corpo · fighter, tank

- **P · Cadência Marcial** — O primeiro ataque básico de Jarvan a um inimigo causa Dano Físico adicional com base na Vida atual dele. Este efeito não ocorrerá no mesmo inimigo novamente por alguns segundos.
- **Q · Ataque do Dragão** — Jarvan expande sua lança, causando 90/130/170/210/250 + 1.45 attack_damage (bonus) de Dano Físico e removendo 10/14/18/22/26% de Armadura por (?)s. Se a lança se conectar com Estandarte Demaciano, Jarvan se impulsiona até ele, Arremessando ao ar os inimigos no trajeto por 0,75s.(?)
- **W · Égide de Ouro** — Jarvan invoca uma égide, causando 15/20/25/30/35% de Lentidão aos inimigos próximos por (?)s e concedendo a si mesmo (?) de Escudo, aumentado em 0/0/0/0/0 + 0.013 max_health para cada Campeão inimigo atingido.(?)
- **E · Estandarte Demaciano** — Passivo: Jarvan recebe 20/22.5/25/27.5/30% de Velocidade de Ataque. Ativo: Jarvan arremessa um estandarte no chão, causando (?) de Dano Mágico e concedendo aos aliados próximos 20/22.5/25/27.5/30% de Velocidade de Ataque por (?)s.(?)
- **R · Cataclisma** — Jarvan salta heroicamente até um Campeão inimigo, causando 200/325/450 + 1.8 attack_damage (bonus) de Dano Físico a ele e aos inimigos próximos, criando uma arena de terreno intransponível ao redor deles por 3.5/3.5/3.5s. Jarvan pode Reconjurar para destruir o terreno criado.(?)

## Elise, a Aranha Rainha

`Elise` · id 60 · a distância · assassin, mage

- **P · Aranha Rainha** — Forma Humana: Quando as habilidades de Elise atingem um inimigo, ela recebe uma Cria dormente. Forma de Aranha: Ataques básicos causam Dano Mágico adicional e restauram a Vida de Elise.
- **Q · Neurotoxina / Mordida Venenosa** — Forma Humana: Elise injeta uma neurotoxina no alvo, causando 40/70/100/130/160 mais 4/4/4/4/4 + 0.03 ability_power da Vida atual dela como Dano Mágico.(?)
- **W · Cria Volátil / Frenesi Aracnídeo** — Forma Humana: Elise invoca uma aranha explosiva que se move até um local, explodindo ao se aproximar de um inimigo ou depois de 3s. A aranha causa @spell.EliseHumanW:TotalDamage@ de Dano Mágico.(?)
- **E · Casulo / Rapel** — Forma Humana: Elise arremessa um casulo, Atordoando e revelando o primeiro inimigo atingido por 1.6/1.8/2/2.2/2.4s.(?)
- **R · Forma de Aranha** — Forma Humana: Elise se transforma em uma aranha ameaçadora, tornando-se atacante corpo a corpo, recebendo acesso às Habilidades da Forma de Aranha e invocando todas as Crias dormentes.(?)

## Orianna, a Donzela Mecânica

`Orianna` · id 61 · a distância · mage, support

- **P · Dando Corda** — Os Ataques de Orianna causam Dano Mágico adicional. O dano aumenta a cada vez que ela Ataca o mesmo alvo.
- **Q · Comando: Atacar** — Orianna comanda que sua Esfera se mova para uma área, causando 60/90/120/150/180 + 0.55 ability_power de Dano Mágico aos inimigos próximos e aos que ela atravessar em seu trajeto. Causa 30/30/30/30/30% a menos de dano a todos os inimigos depois do primeiro.(?)
- **W · Comando: Dissonância** — Orianna ordena que sua Esfera lance um pulso elétrico, causando 70/110/150/190/230 + 0.8 ability_power de Dano Mágico a inimigos próximos. O pulso deixa para trás um campo de energia por 3/3/3/3/3s, causando 20/25/30/35/40% de Lentidão a inimigos e concedendo 20/25/30/35/40% de Velocidade de Movimento adicional a aliados, que decai ao longo de 2/2/2/2/2s.(?)
- **E · Comando: Proteger** — Passivo: a Esfera adiciona 6/12/18/24/30 de Armadura e 6/12/18/24/30 de Resistência Mágica ao Campeão aliado a que estiver conectada. Ativo: Orianna comanda que sua Esfera se prenda a um Campeão aliado, concedendo 55/90/125/160/195 + 0.45 ability_power de Escudo por 2.5/2.5/2.5/2.5/2.5s. Os inimigos atingidos pela Esfera sofrem 60/90/120/150/180 + 0.3 ability_power de Dano Mágico.(?)
- **R · Comando: Onda de Choque** — Orianna comanda que sua Esfera desfira uma onda de choque, causando 225/350/475 + 1.1 ability_power de Dano Mágico a inimigos próximos e Arremessando-os ao ar na direção da Esfera.(?)

## Wukong, o Macaco Rei

`Wukong` · id 62 · corpo a corpo · fighter, tank

- **P · Pele de Pedra** — Wukong recebe Armadura e Regeneração de Vida máxima cumulativas quando luta contra Campeões e monstros.
- **Q · Golpe Destruidor** — O próximo Ataque de Wukong e de seu Clone recebem 135/145/155/165/175 de alcance, causam 20/45/70/95/120 + 0.5 attack_damage (bonus) de Dano Físico adicional e removem 10/15/20/25/30% da Armadura do alvo por 3/3/3/3/3s. O Tempo de Recarga da Habilidade é reduzido em 0.5/0.5/0.5/0.5/0.5s sempre que Wukong ou o Clone atingir um inimigo com um Ataque ou uma Habilidade. Esta Habilidade ativa efeitos ao causar dano.(?)
- **W · Guerreiro Trapaceiro** — Wukong avança e fica Invisível por 1/1/1/1/1s, deixando para trás um Clone imóvel que permanece no lugar por 4/4/4/4/4s. O Clone imita a ultimate de Wukong e Ataca inimigos próximos a quem Wukong tenha causado dano recentemente, causando 40/45/50/55/60% do dano normal.(?)
- **E · Resplendor das Nuvens** — Wukong avança até um inimigo, enviando Clones que imitam o avanço até 2/2/2/2/2 inimigos próximos adicionais. Cada inimigo atingido sofre 80/120/160/200/240 + 1 ability_power de Dano Mágico. Ele e o Clone recebem 40/45/50/55/60% de Velocidade de Ataque por 5/5/5/5/5s. (?)
- **R · Ciclone** — Wukong recebe 20/20/20% de Velocidade de Movimento e gira seu cajado, Arremessando ao ar os inimigos próximos por 0.6/0.6/0.6s e causando 0/0/0 + 2.75 attack_damage mais 0.08/0.12/0.16 da Vida máxima como Dano Físico adicional ao longo de 2/2/2s. A Habilidade pode ser conjurada uma segunda vez dentro de 8/8/8s antes de entrar em Tempo de Recarga.(?)

## Brand, a Vingança Flamejante

`Brand` · id 63 · a distância · mage, support

- **P · Labareda** — As habilidades de Brand deixam seus alvos incandescentes, causando dano ao longo de 4s, acumulando até 3 vezes. Se Brand abate um inimigo que estiver incandescente, ele recupera Mana. Quando Labareda chegar ao máximo de acúmulos em um Campeão ou monstro grande, ela se torna instável. Detona depois de 2s, aplicando efeitos de habilidades e causando muito dano em uma área ao redor da vítima.
- **Q · Cauterizar** — Brand arremessa uma bola de fogo que causa 70/100/130/160/190 + 0.65 ability_power de Dano Mágico ao primeiro inimigo atingido. Se o alvo estiver Incandescente, ele será Atordoado por 1.75/1.75/1.75/1.75/1.75s.(?)
- **W · Pilar de Chamas** — Brand cria um pilar de fogo puro, causando 75/120/165/210/255 + 0.7 ability_power de Dano Mágico. Unidades Incandescentes sofrem 93.75/150/206.25/262.5/318.75 + 0.875 ability_power de dano.(?)
- **E · Conflagração** — Brand conjura uma poderosa explosão no alvo, causando 55/80/105/130/155 + 0.6 ability_power de Dano Mágico a unidades próximas. Se o alvo estiver Incandescente, o alcance da dispersão de Conflagração é dobrado.(?)
- **R · Piroclasma** — Brand desfere uma torrente de fogo devastadora que ressalta para Brand ou para outro inimigo até 5 vezes, causando 100/175/250 + 0.3 ability_power de Dano Mágico aos inimigos cada vez que ressaltar. Os ressaltos priorizam o acúmulo máximo de Labareda em Campeões. Se o alvo estiver Incandescente, ele sofrerá 30/45/60% de Lentidão por um breve momento.(?)

## Lee Sin, o Monge Cego

`Lee Sin` · id 64 · corpo a corpo · fighter, assassin

- **P · Agitação** — Após Lee Sin usar uma habilidade, seus próximos 2 ataques básicos recebem Velocidade de Ataque e recuperam Energia.
- **Q · Onda Sônica / Ataque Ressonante** — Lee Sin projeta uma onda sonora dissonante que causa 60/90/120/150/180 + 0.9 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido, concede Visão Mágica dele e permite que Lee Sin Reconjure dentro de 3/3/3/3/3s. Reconjuração: Lee Sin avança até o inimigo atingido pela onda sonora, causando 60/90/120/150/180 + 0.9 attack_damage (bonus) - 120/180/240/300/360 + 1.8 attack_damage (bonus) de Dano Físico que escala com a Vida perdida do alvo. (?)
- **W · Proteger / Vontade de Ferro** — Lee Sin avança até um aliado ou sentinela. Se o alvo for um Campeão, Lee Sin concede a ambos 60/105/150/195/240 + 0.8 ability_power de Escudo por 2/2/2/2/2s. Lee Sin pode Reconjurar pelos próximos 3/3/3/3/3s. Reconjuração: Lee Sin recebe 10/14/18/22/26% de Vampirismo Universal por 4/4/4/4/4s.(?)
- **E · Tempestade / Mutilar** — Lee Sin golpeia o chão, enviando uma onda de choque que causa 35/60/85/110/135 + 0.9 attack_damage de Dano Mágico e revela unidades inimigas atingidas por 4/4/4/4/4s. Se atingir um inimigo, Lee Sin pode Reconjurar pelos próximos 3/3/3/3/3s. Reconjuração: Lee Sin causa 35/45/55/65/75% de Lentidão aos inimigos próximos atingidos pela onda de choque, decaindo ao longo de 4/4/4/4/4s.(?)
- **R · Fúria do Dragão** — Lee Sin desfere um poderoso chute giratório, Empurrando um Campeão inimigo e causando 175/400/625 + 2 attack_damage (bonus) de Dano Físico. Inimigos contra os quais o alvo colidir são brevemente Arremessados ao ar e sofrem 175/400/625 + 2 attack_damage (bonus) mais 12/15/18% da Vida adicional do primeiro inimigo atingido como Dano Físico.(?)

## Vayne, a Caçadora Noturna

`Vayne` · id 67 · a distância · marksman, assassin

- **P · Caçadora Noturna** — Vayne caça implacavelmente os malfeitores, ganhando Velocidade de Movimento ao se deslocar em direção a Campeões inimigos próximos.
- **Q · Rolamento** — Vayne faz um rolamento curto e causa 0/0/0/0/0 + 0.5 ability_power + 0.75 attack_damage de Dano Físico adicional no próximo Ataque. Esta Habilidade ativa efeitos ao causar dano.(?)
- **W · Dardos de Prata** — Passivo: cada terceiro Ataque ou Habilidade consecutivo contra um inimigo causa 0.06/0.07/0.08/0.09/0.1 da Vida máxima como Dano Verdadeiro adicional.(?)
- **E · Condenar** — Vayne dispara um projétil que Empurra e causa 50/85/120/155/190 + 0.5 attack_damage (bonus) de Dano Físico. Se o alvo colidir com um terreno, ele sofrerá 75/127.5/180/232.5/285 + 0.75 attack_damage (bonus) de Dano Físico adicional e será Atordoado por 1.5/1.5/1.5/1.5/1.5s. (?)
- **R · Hora Final** — Vayne recebe 35/50/65 de Dano de Ataque por 8/10/12s, estendendo-se em 4/4/4s sempre que um Campeão a quem ela tiver causado dano for abatido dentro de 3/3/3s. Além disso, durante esse efeito:Caçadora Noturna concede 90/90/90 de Velocidade de Movimento.O Tempo de Recarga de Rolamento é reduzido em 30/40/50% e concede Invisibilidade por 1/1/1s.(?)

## Rumble, a Ameaça Mecânica

`Rumble` · id 68 · corpo a corpo · fighter, mage

- **P · Titã do Ferro-Velho** — Cada habilidade que Rumble conjura lhe concede Aquecimento. Ao chegar em 50% de Aquecimento, ele entra na Zona de Perigo, concedendo efeitos adicionais a todas as suas habilidades. Ao chegar a 100% de Aquecimento, ele começa a superaquecer, recebendo Velocidade de Ataque adicional e dano adicional em seus ataques básicos, mas fica incapaz de conjurar habilidades por alguns segundos.
- **Q · Cospe-Fogo** — Rumble acende seu lançador de chamas, causando 50/75/100/125/150 + 1.05 ability_power e 6/6.5/7/7.5/8% de Dano Mágico com base na Vida máxima por 3/3/3/3/3s, menos 70/70/70/70/70% contra tropas. Zona de Perigo: o dano aumenta para 75/112.5/150/187.5/225 + 1.5749999 ability_power e 0.06/0.065/0.07/0.075/0.08 da Vida máxima. A porcentagem de dano é limitada a 300/300/300/300/300 de dano a monstros. (?)
- **W · Escudo de Sucata** — Rumble ergue uma barreira, ganhando 25/55/85/115/145 + 0.3 ability_power + 0.04 max_health de Escudo por (?)s e 10/15/20/25/30% de Velocidade de Movimento por 1/1/1/1/1s. Zona de Perigo: em vez disso, a barreira concede 37.5/82.5/127.5/172.5/217.5 + 0.45000002 ability_power + 0.06 max_health de Escudo e 0.1/0.15/0.2/0.25/0.3 de Velocidade de Movimento.(?)
- **E · Arpão Elétrico** — Rumble lança um arpão elétrico, causando 55/80/105/130/155 + 0.5 ability_power de Dano Mágico, 15/20/25/30/35% de Lentidão por 2/2/2/2/2s e reduzindo a Resistência Mágica do inimigo em 10/12/14/16/18% por 4/4/4/4/4s. Atingir inimigos sob Lentidão causada por esta Habilidade aumenta a Lentidão em 30/40/50/60/70% e reduz a Resistência Mágica do inimigo em 20/24/28/32/36% Zona de Perigo: o arpão causa 82.5/120/157.5/195/232.5 + 0.75 ability_power de Dano Mágico e a Lentidão e Resistência Mágica reduzidas aumentam em 50%.(?)
- **R · O Equalizador** — Rumble lança foguetes em linha reta, criando um rastro flamejante que dura 4.5/4.5/4.5s, causando 35/35/35% de Lentidão e 120/200/280 + 0.35 ability_power de Dano Mágico por segundo. Ele pode controlar a direção do rastro clicando e arrastando enquanto conjura a Habilidade.(?)

## Cassiopeia, o Abraço da Serpente

`Cassiopeia` · id 69 · a distância · mage

- **P · Graça Serpentina** — Todos os efeitos de Velocidade de Movimento são mais eficazes em Cassiopeia.
- **Q · Explosão Venenosa** — Cassiopeia causa uma explosão de gás venenoso, Envenenando inimigos e causando 75/110/145/180/215 + 0.65 ability_power de Dano Mágico ao longo de 3/3/3/3/3s. Atingir um Campeão concede a ela 30/35/40/45/50% de Velocidade de Movimento que decai ao longo de 3/3/3/3/3s.(?)
- **W · Miasma** — Cassiopeia cospe veneno, deixando nuvens tóxicas no chão por 5/5/5/5/5s. Inimigos dentro das nuvens sofrem 20/25/30/35/40 + 0.1 ability_power de Dano Mágico por segundo, são Envenenados, Presos ao chão e sofrem 40/50/60/70/80% de Lentidão.(?)
- **E · Presas Duplas** — Cassiopeia lança presas mortais, causando 52/52/52/52/52 + 0.1 ability_power de Dano Mágico. Se o inimigo estiver Envenenado, ele sofrerá 20/45/70/95/120 + 0.55 ability_power de Dano Mágico adicional e Cassiopeia restaurará 0/0/0/0/0 + 0.1 ability_power de Vida, valor esse reduzido para 0/0/0/0/0 + 0.025 ability_power de Vida contra tropas e monstros pequenos. Se a Habilidade abater o alvo, Cassiopeia restaurará 40/40/40/40/40 de Mana. (?)
- **R · Olhar Petrificador** — Cassiopeia lança um Olhar Petrificador, causando 150/250/350 + 0.5 ability_power de Dano Mágico e Atordoando inimigos que estiverem de frente para ela por 2/2/2s. Inimigos que não estiverem de frente para ela sofrem 40/40/40% de Lentidão que decai ao longo da mesma duração.(?)

## Skarner, o Soberano Primordial

`Skarner` · id 72 · corpo a corpo · tank, fighter

- **P · Linhas de Vibração** — Os Ataques do Skarner, Terra Arrasada, Sublevação e Empalar, aplicam Tremor. Com o máximo de acúmulos de Tremor, os inimigos sofrem Dano Mágico com base na Vida máxima ao longo da duração.
- **Q · Terra Arrasada/Sublevação** — Skarner arranca um pedregulho do chão, fortalecendo seus próximos 3 Ataques com 20/25/30/35/40% de Velocidade de Ataque e causando 10/20/30/40/50 + 0.9 attack_damage (bonus) + 0.03 max_health (bonus) de Dano Físico a inimigos próximos. O último Ataque causa 11/11/11/11/11% da Vida máxima como Dano Físico adicional e 40/40/40/40/40% de Lentidão por 1/1/1/1/1s a inimigos atingidos. Reconjurar: Skarner encerra esta Habilidade e arremessa o pedregulho, causando @spell.SkarnerQ:AbilityDamage@ + @spell.SkarnerQ:MaxHPPercent*100@% da Vida máxima como Dano Físico, além de @spell.SkarnerQ:SlowPercent*100@% de Lentidão ao primeiro inimigo atingido e aos inimigos próximos por @spell.SkarnerQ:SlowDuration(?)SpellModifierDescriptionAppend@
- **W · Bastião Sísmico** — Skarner recebe 0/0/0/0/0 + 0.08 max_health de Escudo por 2.5/2.5/2.5/2.5/2.5s e gera um terremoto que causa 50/70/90/110/130 + 0.8 ability_power de Dano Mágico e @SlowEffect*-100@% de Lentidão aos inimigos próximos por 1/1/1/1/1s.(?)
- **E · Impacto de Ixtal** — Skarner avança, seguindo na direção escolhida e ignorando o terreno. Se cruzar com um Campeão ou monstro grande, Skarner o arrastará pelo restante da investida. Colidir com uma parede enquanto segura um inimigo causa 30/60/90/120/150 + 1.2 attack_damage (bonus) + 0.06 max_health de Dano Físico ao inimigo e o Atordoa por 1.1/1.1/1.1/1.1/1.1s. Skarner pode Reconjurar esta Habilidade para encerrar a investida mais cedo.(?)
- **R · Empalar** — As caudas de Skarner atacam, causando 150/250/350 + 1 ability_power de Dano Mágico e Suprimindo os 3 primeiros Campeões atingidos por 1.5/1.5/1.5s. Os Campeões atingidos são arrastados por Skarner pela duração da Habilidade. Se Skarner atingir ao menos um Campeão, ele recebe 40/40/40% de Velocidade de Movimento por 1.5/1.5/1.5s. Se Terra Arrasada estiver ativa, Skarner conjurará Sublevação primeiro.(?)

## Heimerdinger, o Inventor Idolatrado

`Heimerdinger` · id 74 · a distância · mage, support

- **P · Afinidade Hextec** — Recebe Velocidade de Movimento enquanto próxima a torres aliadas e torres posicionadas por Heimerdinger.
- **Q · Torre Evolutiva H-28G** — Heimerdinger constrói uma Torre que ataca inimigos próximos. Heimerdinger pode ter até 3/3/3/3/3 torres ativas ao mesmo tempo. As Torres são carregadas lentamente. Com o máximo de cargas, elas disparam um ataque mais forte. Se Heimerdinger estiver muito longe, suas Torres são desativadas depois de 8s. A Habilidade tem 3/3/3/3/3 cargas.(?)
- **W · Micro-Mísseis Hextec** — Heimerdinger dispara uma barragem de 5/5/5/5/5 mísseis, causando 50/75/100/125/150 + 0.55 ability_power de Dano Mágico ao primeiro inimigo atingido. Mísseis adicionais que atingirem o alvo causam dano reduzido. Dano máximo: 90/135/180/225/270 + 1.03 ability_power de Dano Mágico. Torres próximas recebem 20% de carga por cada míssil que atingir um Campeão.(?)
- **E · Granada de Tempestade de Elétrons CH-2** — Heimerdinger lança uma granada em uma área, causando 60/100/140/180/220 + 0.6 ability_power de Dano Mágico e (?)% de Lentidão por 2/2/2/2/2s. Inimigos no centro também são Atordoados por 1.5/1.5/1.5/1.5/1.5s. Atingir um Campeão carrega totalmente as Torres próximas.(?)
- **R · MELHORIA!!!** — Heimerdinger aprimora sua próxima Habilidade (exceto ultimate). Torre Apex H-28Q: posiciona uma Torre aprimorada por 8s (não influencia o limite de torres) que causa 80/100/120 + 0.35 ability_power de Dano Mágico por disparo comum e 100/140/180 + 0.7 ability_power de Dano Mágico por disparo carregado. A Torre é imune a Controle de Grupo. Além disso, seu ataque causa dano em área e 25% de Lentidão por 2s. Enxame de Mísseis Hextec: dispara 4 remessas de mísseis, e cada uma causa 135/180/225 + 0.45 ability_power de Dano Mágico. Campeões e monstros da selva atingidos por mísseis adicionais sofrem dano reduzido, mas tropas sofrem dano aumentado. Dano máximo: 503/697.5/892 + 1.83 ability_power de Dano Mágico. Granada de Relâmpagos CH-3X: arremessa uma granada saltitante que descarrega três vezes, causando 100/200/300 + 0.6 ability_power de Dano Mágico. As áreas de Atordoamento e Lentidão são maiores. Reconjuração: cancela a Habilidade.(?)

## Nasus, o Curador das Areias

`Nasus` · id 75 · corpo a corpo · fighter, tank

- **P · Devorador de Almas** — Nasus drena a energia espiritual de seu inimigo, acumulando Roubo de Vida adicional.
- **Q · Ataque Sifão** — O próximo Ataque de Nasus causa (?) de Dano Físico. Abater um inimigo com esse Ataque aumenta permanentemente seu dano em 4/4/4/4/4, aumentado para 10/10/10/10/10 contra Campeões, tropas grandes e monstros grandes da selva.(?)
- **W · Murchar** — Nasus envelhece um Campeão inimigo, causando 35/35/35/35/35% de Lentidão, que aumenta para 47/59/71/83/95% ao longo de 5/5/5/5/5s. A Velocidade de Ataque do alvo também é reduzida em 75/75/75/75/75% da Lentidão.(?)
- **E · Fogo Espiritual** — Nasus acende uma chama espiritual, causando 50/80/110/140/170 + 0.6 ability_power de Dano Mágico. Inimigos na área perdem @ArmorShredPercent*-100@% de Armadura e sofrem 10/16/22/28/34 + 0.12 ability_power de Dano Mágico ao longo de 5/5/5/5/5s.(?)
- **R · Fúria das Areias** — Nasus se fortalece dentro de uma tempestade de areia por 15s, aumentando sua Vida máxima em 300/450/600, além de receber 40/55/70 de Armadura e Resistência Mágica. Enquanto a tempestade estiver ativa, os inimigos próximos sofrerão 0.03/0.04/0.05 + 0.0001 ability_power da Vida máxima deles como Dano Mágico por segundo e o Tempo de Recarga de Ataque Sifão será reduzido em 50/50/50%.(?)

## Nidalee, a Caçadora Bestial

`Nidalee` · id 76 · a distância · assassin, mage

- **P · Espreitar** — Mover-se pelos arbustos aumenta a Velocidade de Movimento de Nidalee em 10% por 2 segundos, aumentando para 30% em direção a campeões inimigos visíveis dentre 1400 de alcance. Atingir Campeões ou monstros com Arremessar Lança ou Arapuca aciona uma Caçada, que concede Visão Mágica deles por 4 segundos. Durante esse tempo, Nidalee recebe 10% de Velocidade de Movimento (aumentando para 30% em direção ao alvo Caçado) e o uso de Bote ou Investida contra o alvo é aprimorado.
- **Q · Arremessar Lança / Bote** — Forma Humana: Nidalee arremessa uma lança, causando 70/90/110/130/150 + 0.5 ability_power de Dano Mágico, aumentando até 227.5/292.5/357.5/422.5/487.5 + 1.625 ability_power de Dano Mágico com base na distância percorrida.(?)
- **W · Arapuca / Investida** — Forma Humana: Nidalee posiciona uma armadilha invisível por 2min. Quando um inimigo passa por cima da armadilha, ele sofre 10/20/30/40/50 + 0.05 ability_power de Dano Mágico por segundo durante 4/4/4/4/4s. Nidalee pode ativar até 10/10/10/10/10 armadilhas ao mesmo tempo.(?)
- **E · Ímpeto Selvagem / Patada** — Forma Humana: Nidalee restaura 50/75/100/125/150 + 0.35 ability_power de Vida, aumentado para até 100/150/200/250/300 + 0.7 ability_power com base na Vida perdida, e concede 30/40/50/60/70% de Velocidade de Ataque por 7/7/7/7/7s.(?)
- **R · Aspecto do Puma** — Passivo: enquanto estiver na Forma Humana, aplicar Caça redefine o Tempo de Recarga da Habilidade. Forma Humana: Nidalee se transforma na Forma de Puma, recebendo alcance corpo a corpo nos Ataques e substituindo suas Habilidades ativas. Forma de Puma: Nidalee se transforma na Forma Humana, recebendo alcance à distância nos Ataques e substituindo suas Habilidades ativas.(?)

## Udyr, o Andarilho Espiritual

`Udyr` · id 77 · corpo a corpo · fighter, tank

- **P · Ponte Espiritual** — Udyr tem quatro Habilidades básicas que o fazem alternar entre Posturas e pode Reconjurar uma Habilidade para renová-la com efeitos adicionais na ultimate. Além disso, após usar uma Habilidade, os próximos dois Ataques de Udyr recebem Velocidade de Ataque.
- **Q · Garra Selvagem** — Postura da Garra: Udyr recebe 20/32/44/56/68% de Velocidade de Ataque e seus Ataques causam 6/12/18/24/30 + 0.2 attack_damage (bonus) + 0.01 max_health (bonus) de Dano Físico %i:OnHit% ao contato por 4/4/4/4/4s. Além disso, os próximos dois Ataques de Udyr nessa postura causam 0.03/0.04/0.05/0.06/0.07 + 0.00035 attack_damage (bonus) da Vida máxima como Dano Físico adicional e ganham 50/50/50/50/50 de alcance. Despertar: aumenta a Velocidade de Ataque adicional para (?) e o dano da Vida máxima para 0.07/0.08/0.09/0.099999994/0.11 + 0.0005 attack_damage (bonus) + 0.00001 max_health (bonus). Os próximos dois Ataques de Udyr invocam relâmpagos seis vezes, causando um total de (?) da Vida máxima como Dano Mágico a alvos isolados (os relâmpagos se propagam para outros alvos próximos quando possível). (?)
- **W · Manto de Ferro** — Postura do Manto: Udyr recebe 45/65/85/105/125 + 0.4 ability_power + 0.5 attack_damage (bonus) + 0.02 max_health de Escudo por 4/4/4/4/4s. Além disso, os próximos dois Ataques de Udyr recebem 15/16/17/18/19% de Roubo de Vida e restauram 0/0/0/0/0 + 0.08 ability_power + 0.012 max_health de Vida. Despertar: recebe 195/215/235/255/275 + 0.65 ability_power + 1 attack_damage (bonus) + 0.08 max_health de Escudo, restaura 97.5/107.5/117.5/127.5/137.5 + 0.325 ability_power + 0.5 attack_damage (bonus) + 0.04 max_health de Vida ao longo de 4/4/4/4/4s, e os próximos dois Ataques de Udyr recebem 30/32/34/36/38% de Roubo de Vida e restauram 0/0/0/0/0 + 0.16 ability_power + 0.024 max_health de Vida. (?)
- **E · Investida Ardente** — Postura da Investida: Udyr recebe 0.25/0.31/0.37/0.43/0.49 + 0.0005 attack_damage (bonus) de Velocidade de Movimento que decai ao longo de 4/4/4/4/4s. Além disso, os Ataques de Udyr o fazem avançar até o alvo e causam Atordoamento por 0.75/0.75/0.75/0.75/0.75s (6/5.6/5.2/4.8/4.4s de Tempo de Recarga por alvo). Despertar: concede imunidade a efeitos Imobilizadores e Debilitantes e 0.4/0.4/0.4/0.4/0.4 + 0.001 attack_damage (bonus) de Velocidade de Movimento adicional por 1.5/1.5/1.5/1.5/1.5s.(?)
- **R · Tempestade Alada** — Postura da Tormenta: Udyr cerca a si mesmo com uma tempestade glacial por 4/4/4s, causando 15/18/21% de Lentidão e 20/36/52 + 0.35 ability_power de Dano Mágico por segundo aos inimigos próximos. Além disso, os próximos dois Ataques de Udyr nessa postura causam 40/40/40 + 0.35 ability_power de Dano Mágico aos inimigos dentro da tempestade. Despertar: desencadeia a tempestade, fazendo-a perseguir o último inimigo Atacado por Udyr, causando 0.14/0.14/0.14 + 0.00035 ability_power da Vida máxima como Dano Mágico adicional ao longo da duração e 0.05/0.05/0.05 de Lentidão adicional.(?)

## Poppy, a Guardiã do Martelo

`Poppy` · id 78 · corpo a corpo · tank, fighter

- **P · Embaixadora de Ferro** — Poppy arremessa seu broquel, que bate no alvo e cai no chão. Ela pode apanhá-lo para receber um escudo temporário.
- **Q · Choque do Martelo** — Poppy golpeia o chão, causando 30/55/80/105/130 + 0.75 attack_damage (bonus) mais 7/7.5/8/8.5/9% da Vida máxima como Dano Físico e tornando o solo instável. A área instável causa 0.2/0.23/0.26/0.29/0.32 + 0.00008 max_health (bonus)% de Lentidão aos inimigos e explode depois de 1/1/1/1/1s, causando 30/55/80/105/130 + 0.75 attack_damage (bonus) mais 7/7.5/8/8.5/9% da Vida máxima como Dano Físico novamente.(?)
- **W · Presença Inabalável** — Passivo: Poppy recebe 0/0/0/0/0 + 0.16 armor de Armadura e 0/0/0/0/0 + 0.16 magic_resist de Resistência Mágica. Esse efeito será dobrado se Poppy estiver com menos de 40/40/40/40/40% de Vida. Ativo: Poppy recebe 40/40/40/40/40% de Velocidade de Movimento e projeta um campo ao seu redor por 2/2/2/2/2s que impede avanços de inimigos. Inimigos impedidos pelo campo são Presos ao chão e sofrem @SlowAmount*-100@% de Lentidão por 2/2/2/2/2s, além de 70/110/150/190/230 + 0.7 ability_power de Dano Mágico.(?)
- **E · Investida Heroica** — Poppy avança sobre um inimigo, causando 40/60/80/100/120 + 0.6 attack_damage (bonus) de Dano Físico e empurrando-o para a frente. Se ela o fizer colidir contra um terreno, o inimigo será Atordoado por 1.6/1.7/1.8/1.9/2s e sofrerá 40/60/80/100/120 + 0.6 attack_damage (bonus) de Dano Físico adicional.(?)
- **R · Veredito da Guardiã** — Início do carregamento: Poppy começa a canalizar com o martelo por até 4/4/4s, sofrendo 15/15/15% de Lentidão. Liberar: Poppy golpeia o chão, criando uma onda de choque que causa 200/300/400 + 0.9 attack_damage (bonus) de Dano Físico a inimigos ao redor do primeiro Campeão atingido, Arremessando-os em direção ao seu Nexus e tornando-os Inalvejáveis enquanto estiverem suspensos. O alcance e a distância do Arremesso aumentam de acordo com a duração do carregamento. Um arremesso sem carregamento causa 100/150/200 + 0.45 attack_damage (bonus) de Dano Físico e Arremessa ao ar por 1/1/1s.(?)

## Gragas, o Badernista

`Gragas` · id 79 · corpo a corpo · fighter, mage

- **P · Happy Hour** — Gragas cura a si mesmo periodicamente após utilizar uma habilidade.
- **Q · Jogar o Barril** — Gragas arremessa um barril que explode após 4/4/4/4/4s, causando 80/120/160/200/240 + 0.8 ability_power - 120/180/240/300/360 + 1.2 ability_power de Dano Mágico, além de 40/45/50/55/60 - 60/67.5/75/82.5/90% de Lentidão por 2/2/2/2/2s. O dano e a Lentidão aumentam de acordo com o tempo que o barril levou para explodir. Gragas pode Reconjurar para detonar o barril antecipadamente.(?)
- **W · Fúria da Bebedeira** — Gragas prova sua bebida, sofrendo menos 10/14/18/22/26 + 0.04 ability_power de dano por 2.5/2.5/2.5/2.5/2.5s e causando 20/50/80/110/140 + 0.7 ability_power mais 7/7/7/7/7% da Vida máxima como Dano Mágico adicional ao alvo e a inimigos ao redor no próximo Ataque.(?)
- **E · Barrigada** — Gragas avança, colidindo com o primeiro inimigo, Arremessando ao ar os inimigos próximos por 1/1/1/1/1s e causando 80/125/170/215/260 + 0.6 ability_power de Dano Mágico a eles. O Tempo de Recarga da Habilidade é reduzido em 40/40/40/40/40% se Gragas colidir com um inimigo.(?)
- **R · Barril Explosivo** — Gragas arremessa um barril, causando 200/300/400 + 0.8 ability_power de Dano Mágico e Empurrando inimigos da zona de impacto.(?)

## Pantheon, a Lança Indestrutível

`Pantheon` · id 80 · corpo a corpo · fighter, assassin

- **P · Determinação Mortal** — Após algumas habilidades ou ataques, a próxima habilidade de Pantheon será fortalecida.
- **Q · Lança Meteórica** — Pressionar: Pantheon desfere uma estocada com a lança, causando 70/100/130/160/190 + 1.15 attack_damage (bonus) de Dano Físico aos inimigos atingidos. Restitui 60/60/60/60/60% do Tempo de Recarga desta Habilidade. Segurar: Pantheon arremessa a lança, causando 70/100/130/160/190 + 0.5 ability_power + 1.15 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido e 50/50/50/50/50% a menos de dano contra os próximos alvos. Contra inimigos com menos de 20/20/20/20/20% de Vida, esta Habilidade é aprimorada e causa 155/230/305/380/455 + 2.3 attack_damage (bonus) de Dano Físico. Efeito adicional de Determinação Mortal: causa 240/240/240/240/240 + 1.15 attack_damage (bonus) de Dano Físico adicional.(?)
- **W · Escudo-Cometa** — Pantheon salta sobre o alvo, Atordoando-o por 1/1/1/1/1s e causando 0.06/0.065/0.07/0.075/0.08 + 0.00015 ability_power + 0.00004 max_health (bonus) de Dano Físico com base na Vida máxima. Efeito adicional de Determinação Mortal: o próximo Ataque de Pantheon golpeia 3/3/3/3/3 vezes, causando 0/0/0/0/0 + 1.6500001 attack_damage de Dano Físico.(?)
- **E · Égide Impetuosa** — Pantheon prepara o escudo e enfrenta os inimigos na direção escolhida por 1.5/1.5/1.5/1.5/1.5s, tornando-se imune a dano vindo daquela direção (exceto de torres) e causando 0/0/0/0/0 + 1 attack_damage de Dano Físico ao longo da duração. Após canalizar, Pantheon golpeia com o escudo, causando 55/105/155/205/255 + 1.5 attack_damage (bonus) de Dano Físico. Efeito adicional de Determinação Mortal: ao golpear com o escudo, ele recebe 30/30/30/30/30 + 0.025 max_health (bonus) de Armadura e 30/30/30/30/30 + 0.025 max_health (bonus) de Resistência Mágica por 4/4/4/4/4s, além de 60/60/60/60/60% de Velocidade de Movimento por 1.5/1.5/1.5/1.5/1.5s. (?)
- **R · Constelação Cadente** — Passivo: Pantheon recebe 10/20/30% de Penetração de Armadura. Ativo: Pantheon acumula toda a sua força e salta em direção ao céu. Ele arremessa a lança do alto em uma pequena área, causando @spell.PantheonQ:HoldDamageCalc@ de Dano Físico e 50/50/50% de Lentidão por 2/2/2s. Depois disso, Pantheon aterrissa como um meteoro na área-alvo. Causa até 300/500/700 + 1 ability_power de Dano Mágico aos inimigos em linha reta (reduzido em até 50/50/50% nos limites da área). Esta Habilidade deixa Determinação Mortal instantaneamente pronta para ser utilizada.(?)

## Ezreal, o Explorador Pródigo

`Ezreal` · id 81 · a distância · marksman, mage

- **P · Feitiço do Poder Crescente** — Ezreal recebe Velocidade de Ataque crescente a cada vez que acerta uma habilidade, acumulando-se até 5 vezes.
- **Q · Disparo Místico** — Ezreal dispara um projétil de energia, causando 20/45/70/95/120 + 0.4 ability_power + 1.3 attack_damage de Dano Físico ao primeiro inimigo atingido e reduzindo os Tempos de Recarga das Habilidades dele em 1.5/1.5/1.5/1.5/1.5s.(?)
- **W · Fluxo Essencial** — Ezreal dispara um orbe mágico que se prende por 4/4/4/4/4s ao primeiro Campeão, estrutura ou monstro épico da selva que atingir. Ezreal detona o orbe se atingir o alvo com um Ataque ou Habilidade, causando 80/135/190/245/300 + 0.9 ability_power + 1 attack_damage (bonus) de Dano Mágico. Detoná-lo com uma Habilidade recupera o custo de Mana dela mais 60/60/60/60/60 de Mana.(?)
- **E · Translocação Arcana** — Ezreal se teleporta e dispara um projétil no inimigo mais próximo, causando 80/130/180/230/280 + 0.75 ability_power + 0.6 attack_damage (bonus) de Dano Mágico. O projétil prioriza inimigos afetados por Fluxo Essencial.(?)
- **R · Barragem Incendiária** — Ezreal dispara um grande arco de energia que causa 350/550/750 + 1.1 ability_power + 1 attack_damage (bonus) de Dano Mágico. O dano causado a tropas e a monstros não épicos da selva é reduzido a 150/225/300 + 1.1 ability_power + 1 attack_damage (bonus) de Dano Mágico.(?)

## Mordekaiser, o Revenã de Ferro

`Mordekaiser` · id 82 · corpo a corpo · fighter, mage

- **P · Ascensão das Trevas** — Mordekaiser recebe uma poderosa aura de dano e Velocidade de Movimento depois de acertar 3 Ataques ou Habilidades contra Campeões inimigos ou monstros.
- **Q · Obliterar** — Mordekaiser golpeia o chão com a Véu da Noite, causando (?) de Dano Mágico, que aumenta para (?) caso atinja apenas um inimigo.(?)
- **W · Indestrutível** — Passivo: Mordekaiser armazena 45/45/45/45/45% do dano que causa e 7.5/7.5/7.5/7.5/7.5% do dano que sofre. Ativo: Mordekaiser recebe um Escudo equivalente ao dano acumulado. Ele pode Reconjurar a Habilidade para restaurar 35/37.5/40/42.5/45% do Escudo restante como Vida. Escudo mínimo: 0/0/0/0/0 + 0.05 max_health Escudo máximo: 0/0/0/0/0 + 0.3 max_health(?)
- **E · Aperto Mortal** — Passivo: Mordekaiser recebe 5/7.5/10/12.5/15% de Penetração Mágica. Ativo: puxa os inimigos para perto, causando 60/80/100/120/140 + 0.45 ability_power de Dano Mágico.(?)
- **R · Reino da Morte** — Mordekaiser bane um Campeão para o Reino da Morte com ele por 7/7/7s, roubando 10/10/10% de seus atributos principais durante a duração. Caso abata esse inimigo dentro do Reino da Morte, Mordekaiser consome sua alma, mantendo os atributos que roubou até que o alvo ressurja.(?)

## Yorick, o Pastor de Almas

`Yorick` · id 83 · corpo a corpo · fighter, tank

- **P · Pastor de Almas** — A Horda Maldita: Yorick pode invocar Andarilhos da Névoa e atacar inimigos próximos.
- **Q · Extrema-Unção** — O próximo ataque de Yorick causa 30/50/70/90/110 + 0.5 attack_damage de Dano Físico adicional e restaura 10/10/10/10/10 mais 6/7/8/9/10% da Vida perdida, reduzido em 50/50/50/50/50% contra não Campeões. Se esse Ataque atingir um Campeão/monstro grande ou abater o alvo, ele deixará uma cova. Quando houver 3 ou mais covas por perto e esta Habilidade já tiver sido usada, Yorick poderá Reconjurá-la para invocar Andarilhos da Névoa de todas as covas próximas.(?)
- **W · Procissão Sombria** — Yorick invoca uma parede de espíritos, bloqueando o trajeto de inimigos, mas não de aliados. A parede tem 2/2/3/3/4 de Vida e desaparece depois de 4/4/4/4/4s.(?)
- **E · Névoa dos Lamentos** — Yorick arremessa um glóbulo de Névoa Negra que causa 6/6.5/7/7.5/8 + 0.03 ability_power da Vida máxima como Dano Mágico, 0.3/0.3/0.3/0.3/0.3 de Lentidão por 1.5/1.5/1.5/1.5/1.5s e marca Campeões e monstros por 4/4/4/4/4s. Inimigos marcados Despertam covas próximas continuamente (máximo de @Spell.YorickPassive:YorickPassiveGhoulMax@) e sofrem 13/16/19/22/25% de redução de Armadura. Yorick e as unidades que ele invocou recebem 18/21/24/27/30% de Velocidade de Movimento em direção à marca. Os Andarilhos da Névoa saltarão uma vez sobre os inimigos marcados que se afastarem deles.(?)
- **R · Louvor das Ilhas** — Yorick invoca a Donzela da Névoa com 1050/1050/1050 + 0.6 max_health (bonus) de Vida e 50/75/100 + 0.3 attack_damage (bonus) de Dano de Ataque mais 2/3/4 Andarilhos da Névoa. Automaticamente, a Donzela invoca Andarilhos da Névoa de inimigos que tiverem morrido ao redor e marca Campeões inimigos ao atacá-los. Ao causar dano ao alvo da Donzela, Yorick também causará 2/2.5/3% da Vida máxima como Dano Mágico. Após 10s, Yorick pode Reconjurar a Habilidade para libertar a Donzela, enviando-a pela rota mais próxima.(?)

## Akali, a Assassina Renegada

`Akali` · id 84 · corpo a corpo · assassin

- **P · Marca do Assassino** — Causar dano a Campeões com habilidades cria um círculo de energia em volta deles. Sair desse círculo fortalece o próximo Ataque de Akali com alcance e dano adicionais.
- **Q · Golpe dos Cinco Pontos** — Akali arremessa um arco de kunais, causando 45/70/95/120/145 + 0.6 ability_power + 0.65 attack_damage de Dano Mágico e 50/50/50/50/50% de Lentidão aos inimigos atingidos na ponta por 0.5/0.5/0.5/0.5/0.5s.(?)
- **W · Proteção do Crepúsculo** — Akali solta uma bomba de fumaça, liberando uma proteção que se espalha e dura 5/5.5/6/6.5/7s. Além disso, ela recebe 30/35/40/45/50% de Velocidade de Movimento que decai ao longo de 2/2/2/2/2s. Akali aumenta a Energia máxima em 100/100/100/100/100 enquanto a proteção estiver ativa. Enquanto estiver dentro da fumaça, Akali ficará Invisível.(?)
- **E · Investida Shuriken** — Akali salta para trás e atira uma shuriken, causando 70/140/210/280/350 + 1.1 ability_power + 1 attack_damage de Dano Mágico e marcando o primeiro inimigo ou nuvem de fumaça que atingir. É possível Reconjurar uma vez para avançar até o alvo marcado, causando 70/140/210/280/350 + 1.1 ability_power + 1 attack_damage de Dano Mágico.(?)
- **R · Execução Perfeita** — Akali salta sobre um Campeão inimigo, causando 110/220/330 + 0.3 ability_power + 0.5 attack_damage (bonus) de Dano Mágico a todos os inimigos no trajeto. Akali pode Reconjurar depois de 2.5/2.5/2.5s, executando um avanço perfurante e causando 70/140/210 + 0.3 ability_power - 210/420/630 + 0.90000004 ability_power de Dano Mágico com base na Vida perdida.(?)

## Kennen, o Coração da Tempestade

`Kennen` · id 85 · a distância · mage

- **P · Marca da Tormenta** — Kennen atordoa os inimigos que atingir 3 vezes com suas habilidades.
- **Q · Shuriken Trovejante** — Kennen arremessa uma shuriken, causando 75/125/175/225/275 + 0.75 ability_power de Dano Mágico ao primeiro inimigo que atingir.(?)
- **W · Surto Elétrico** — Passivo: a cada 5º Ataque, causa 35/45/55/65/75 + 0.35 ability_power + 0.8 attack_damage (bonus) de Dano Mágico adicional %i:OnHit% ao contato. Ativo: Kennen envia uma onda de eletricidade, causando 70/95/120/145/170 + 0.8 ability_power de Dano Mágico a inimigos próximos afetados por Marca da Tormenta. (?)
- **E · Investida Relâmpago** — Kennen vira uma esfera de raios, recebendo 100/100/100/100/100% de Velocidade de Movimento, causando 80/120/160/200/240 + 0.8 ability_power de Dano Mágico aos inimigos atravessados e ganhando efeito Fantasma por 2/2/2/2/2s. Se Kennen causar dano a pelo menos um inimigo, ele recebe 40/40/40/40/40 de Energia. Depois que a Habilidade termina, Kennen recebe 40/50/60/70/80% de Velocidade de Ataque por 4/4/4/4/4s. Acertos Críticos prolongam a duração do efeito em 1/1/1/1/1s. É possível Reconjurar para finalizar a Habilidade antecipadamente.(?)
- **R · Turbilhão Cortante** — Kennen libera uma tempestade mágica, causando 40/80/120 + 0.25 ability_power de Dano Mágico a todos os inimigos próximos a cada 0.5/0.5/0.5s e recebendo 25/50/75 de Armadura e 25/50/75 de Resistência Mágica por 3/3/3s. Acertos sucessivos causam 10/10/10% de dano aumentado para cada acerto já sofrido pelo inimigo.(?)

## Garen, o Poder de Demacia

`Garen` · id 86 · corpo a corpo · fighter, tank

- **P · Perseverança** — Se Garen não tiver sido atingido recentemente por dano ou habilidades inimigas, ele regenera um percentual de sua Vida total a cada segundo.
- **Q · Acerto Decisivo** — Garen remove todos os efeitos de Lentidão de si e recebe 35/35/35/35/35% de Velocidade de Movimento por 1.4/1.95/2.5/3.05/3.6s. Seu próximo Ataque Silencia por 1.5/1.5/1.5/1.5/1.5s e causa 30/60/90/120/150 + 1.5 attack_damage de Dano Físico.(?)
- **W · Coragem** — Passivo: Garen fica com (?) de Armadura e (?) de Resistência Mágica adicionais. Abater unidades concede 0.2/0.2/0.2/0.2/0.2 de resistências permanentemente, até um máximo de 30/30/30/30/30. Ativo: Garen reúne toda sua coragem por 4/4/4/4/4s, reduzindo todo o dano sofrido em 25/29/33/37/41%. Ele também recebe 65/85/105/125/145 + 0.18 max_health (bonus) de Escudo e 60/60/60/60/60% de Tenacidade por 0.75/0.75/0.75/0.75/0.75s.(?)
- **E · Julgamento** — Garen gira a espada rapidamente por 3/3/3/3/3s, causando 4/7/10/13/16 + 0.4 attack_damage de Dano Físico (?) vezes ao longo da duração. O inimigo mais próximo sofre 25/25/25/25/25% a mais de dano. Campeões atingidos por 6/6/6/6/6 golpes perdem 25/25/25/25/25% de Armadura por 6/6/6/6/6s. Reconjuração: Garen finaliza a Habilidade antecipadamente.(?)
- **R · Justiça Demaciana** — Garen evoca o Poder de Demacia para executar um inimigo, causando 125/200/275 mais 25/30/35% da Vida perdida como Dano Verdadeiro.(?)

## Leona, a Alvorada Radiante

`Leona` · id 89 · corpo a corpo · tank, support

- **P · Luz do Sol** — Habilidades que causam dano atingem os inimigos com Luz do Sol por 1,5s. Quando Campeões aliados causam dano a esses alvos, a Luz do Sol é consumida para causar Dano Mágico adicional.
- **Q · Proteção da Aurora** — O próximo Ataque de Leona Atordoa por 1/1/1/1/1s e causa 10/35/60/85/110 + 0.3 ability_power a mais de Dano Mágico.(?)
- **W · Eclipse** — Leona ergue o escudo, reduzindo todo o dano causado a ela em 8/12/16/20/24 e recebendo 20/27.5/35/42.5/50 + 0.2 armor (bonus) de Armadura e 20/27.5/35/42.5/50 + 0.2 magic_resist (bonus) de Resistência Mágica por 3/3/3/3/3s. O escudo então detona, causando 55/85/115/145/175 + 0.4 ability_power de Dano Mágico a inimigos próximos. Se atingir pelo menos um inimigo, ela mantém a Armadura e Resistência Mágica adicionais por mais 3/3/3/3/3s. (?)
- **E · Lâmina Zênite** — Leona golpeia com uma espada de luz, causando 50/90/130/170/210 + 0.4 ability_power de Dano Mágico. O último Campeão inimigo atingido fica Enraizado por 0.5/0.5/0.5/0.5/0.5s e Leona avança em sua direção.(?)
- **R · Labareda Solar** — Leona conjura um intenso raio de energia solar, causando 150/225/300 + 0.8 ability_power de Dano Mágico e 80/80/80% de Lentidão aos inimigos por 1.75/1.75/1.75s. Inimigos no centro da explosão sofrerão Atordoamento em vez de Lentidão.(?)

## Malzahar, o Profeta do Vazio

`Malzahar` · id 90 · a distância · mage

- **P · Oscilação do Vazio** — Caso não tenha sofrido dano ou efeitos de Controle de Grupo recentemente, Malzahar recebe uma enorme redução de dano e imunidade a Controles de Grupo, que duram um curto período após sofrer dano.
- **Q · Chamado do Vazio** — Malzahar abre dois portais para o Vazio que disparam projéteis entre si, causando 70/105/140/175/210 + 0.55 ability_power de Dano Mágico e Silenciando por 1/1.25/1.5/1.75/2s.(?)
- **W · Enxame do Vazio** — Passivo: as outras Habilidades de Malzahar concedem um acúmulo quando conjuradas (máximo de 2/2/2/2/2). Ativo: Malzahar invoca um Voidling e mais um Voidling adicional por acúmulo. Voidlings duram 8/8/9/9/10s e causam 76.5/78.5/80.5/82.5/84.5 + 0.2 ability_power + 0.4 attack_damage (bonus) de Dano Mágico a cada acerto.(?)
- **E · Visões Maléficas** — Malzahar inflige visões terríveis, causando 80/115/150/185/220 + 0.8 ability_power de Dano Mágico ao longo de 4/4/4/4/4s. Aplicar Chamado do Vazio ou Aperto Ínfero à vítima durante esse período reinicia as visões. Se a vítima morrer, Malzahar recebe 0/0/0/0/0 + 0.02 mana de Mana e as visões se propagam para o inimigo mais próximo. Visões Maléficas executa tropas com menos de 10/10/10/10/10 de Vida.(?)
- **R · Aperto Ínfero** — Malzahar Suprime um Campeão inimigo, causando 125/200/275 + 0.8 ability_power de Dano Mágico ao longo de 2.5/2.5/2.5s. Ele cria uma zona de energia negativa ao redor do alvo, causando 2/3/4 + 0.005 ability_power da Vida máxima como Dano Mágico ao longo de 5/5/5s.(?)

## Talon, a Sombra da Lâmina

`Talon` · id 91 · corpo a corpo · assassin

- **P · Limiar da Lâmina** — As habilidades de Talon causam Ferimentos em Campeões e monstros grandes, acumulando até 3 vezes. Quando Talon ataca um Campeão com 3 acúmulos de Ferimento, o alvo sofre intenso dano de sangramento ao longo do tempo.
- **Q · Diplomacia Noxiana** — Talon salta em direção ao alvo e causa 65/85/105/125/145 + 1 attack_damage (bonus) de Dano Físico. Se a Habilidade for usada em alcance corpo a corpo, ela causará Acerto Crítico de (?) de Dano Físico. Se a Habilidade abater o alvo, Talon restaurará 55/55/55/55/55 de Vida e terá 50/50/50/50/50% do Tempo de Recarga restituído.(?)
- **W · Ancinho** — Talon lança uma rajada de lâminas em uma direção, causando 50/60/70/80/90 + 0.4 attack_damage (bonus) de Dano Físico. Então, as lâminas voltam para ele, causando 60/90/120/150/180 + 0.9 attack_damage (bonus) de Dano Físico e 40/45/50/55/60% de Lentidão por 1/1/1/1/1s.(?)
- **E · Caminho do Assassino** — Talon salta sobre o terreno ou estrutura que estiver mais perto. Ele não pode avançar sobre a mesma seção de terreno por mais de uma vez a cada 160/135/110/85/60s.(?)
- **R · Ataque das Sombras** — Talon dispara um anel de lâminas que causa 90/135/180 + 1 attack_damage (bonus) de Dano Físico, recebe 40/55/70% de Velocidade de Movimento e fica Invisível por 2.5/2.5/2.5s. Quando a Invisibilidade acaba, as lâminas retornam para Talon, causando 90/135/180 + 1 attack_damage (bonus) de Dano Físico novamente. Se ele cancelar a Invisibilidade com um Ataque ou com Diplomacia Noxiana, as lâminas retornam para o alvo.(?)

## Riven, a Exilada

`Riven` · id 92 · corpo a corpo · fighter, assassin

- **P · Lâmina Rúnica** — As habilidades de Riven carregam sua lâmina, e seus ataques básicos consomem cargas para causar dano adicional.
- **Q · Asas Quebradas** — Riven avança para a frente, causando 45/75/105/135/165 + 0.6 attack_damage (bonus) de Dano Físico aos oponentes. Ela pode Reconjurar a Habilidade duas vezes. A primeira Reconjuração é igual à original, mas a segunda tem um efeito diferente: Reconjuração: Riven salta à frente e golpeia o chão, causando 45/75/105/135/165 + 0.6 attack_damage (bonus) de Dano Físico e Arremessando ao ar os inimigos ao redor por 0,75s.(?)
- **W · Explosão de Ki** — A espada de Riven emite uma explosão de energia rúnica, causando 65/95/125/155/185 + 1 attack_damage (bonus) de Dano Físico e Atordoando os inimigos por 0.75/0.75/0.75/0.75/0.75s.(?)
- **E · Valentia** — Riven avança rapidamente e recebe 70/95/120/145/170 + 1.1 attack_damage (bonus) de Escudo por 1,5s.(?)
- **R · Lâmina do Exílio** — A arma de Riven exala energia espiritual, concedendo 0/0/0 + 0.2 attack_damage de Dano de Ataque e alcance aumentado em Ataques e Habilidades de dano por 15/15/15s. Enquanto estiver ativa, é possível Reconjurar. Reconjuração: Riven dispara um golpe de vento que causa 100/150/200 + 0.55 attack_damage (bonus) - 300/450/600 + 1.65 attack_damage (bonus) de Dano Físico com base na Vida perdida dos inimigos.(?)

## Kog'Maw, a Boca do Abismo

`Kog'Maw` · id 96 · a distância · marksman, mage

- **P · Surpresa Icathiana** — Ao morrer, Kog'Maw explode depois de 4s, causando Dano Verdadeiro a inimigos próximos.
- **Q · Cusparada Cáustica** — Passivo: Kog'Maw recebe 5/10/15/20/25% de Velocidade de Ataque. Ativo: Kog'Maw vomita um projétil corrosivo que causa 80/125/170/215/260 + 0.9 ability_power de Dano Mágico ao primeiro inimigo atingido e fragmenta 16/20/24/28/32% de Armadura e Resistência Mágica por 4/4/4/4/4s.(?)
- **W · Barragem Bio-Arcana** — Kog'Maw recebe 130/150/170/190/210 de Alcance de Ataque e seus Ataques causam 3/3.75/4.5/5.25/6 + 0.015 ability_power da Vida máxima como Dano Mágico %i:OnHit% ao contato adicional por 8/8/8/8/8s.(?)
- **E · Gosma do Vazio** — Kog'Maw cospe bile, causando 70/110/150/190/230 + 0.65 ability_power de Dano Mágico e deixando um rastro de gosma por 3/3/3/3/3s. Inimigos que estiverem na gosma sofrem 40/45/50/55/60% de Lentidão.(?)
- **R · Artilharia Viva** — Kog'Maw dispara ácido em uma área, causando 100/140/180 + 0.35 ability_power + 0.75 attack_damage (bonus) mais 0.83333/0.83333/0.83333% a cada 1% da Vida perdida como Dano Mágico e revelando inimigos atingidos por 2s. Inimigos com menos de 40% de Vida sofrem 200/280/360 + 0.7 ability_power + 1.5 attack_damage (bonus) de Dano Mágico. Disparos subsequentes dentro de 8/8/8s custam 40/40/40 de Mana adicional (máximo de 400/400/400 de Mana).(?)

## Shen, o Olho do Crepúsculo

`Shen` · id 98 · corpo a corpo · tank

- **P · Barreira de Ki** — Após conjurar uma habilidade, Shen recebe um escudo. Afetar outros Campeões reduz o Tempo de Recarga desse efeito.
- **Q · Ataque Crepúsculo** — Shen chama sua Espada Espiritual. Inimigos atingidos no trajeto da arma sofrem 25/30/35/40/45% de Lentidão ao se distanciarem de Shen pelos próximos 2/2/2/2/2s. Além disso, os próximos 3/3/3/3/3 ataques de Shen causam 40/40/40/40/40 mais 2/2.5/3/3.5/4 + 0.015 ability_power de Dano Mágico adicional com base na Vida máxima. Se Shen atingir um Campeão inimigo com a Espada Espiritual, o dano é aumentado para 40/40/40/40/40 mais 5/5.5/6/6.5/7 + 0.02 ability_power da Vida máxima como Dano Mágico e ele recebe 50/50/50/50/50% de Velocidade de Ataque para esses Ataques.(?)
- **W · Refúgio Espiritual** — Shen cria uma área defensiva com sua Espada Espiritual por 1.75/1.75/1.75/1.75/1.75s. Ataques contra Campeões aliados dentro da área são bloqueados. Se não houver Campeões a serem protegidos quando a área for criada, ela permanecerá por até 2/2/2/2/2s.(?)
- **E · Corrida das Sombras** — Passivo: causar dano com Ataque Crepúsculo ou com esta Habilidade regenera 50/50/50/50/50 de Energia. Ativo: Shen avança, Provocando Campeões e monstros da selva por 1.5/1.5/1.5/1.5/1.5s e causando 60/85/110/135/160 + 0.11 max_health (bonus) de Dano Físico.(?)
- **R · Manter a União** — Shen concede 120/220/320 + 1.35 ability_power + 0.15 max_health (bonus) - 192/352/512 + 2.16 ability_power + 0.24000001 max_health (bonus) de Escudo a um Campeão aliado em qualquer lugar do mapa com base na Vida perdida do aliado por 5/5/5s (máximo de Escudo com 60% da Vida perdida). Após canalizar por 3/3/3s, Shen teleporta-se para a localização de seu aliado com sua Espada Espiritual.(?)

## Lux, a Dama da Luz

`Lux` · id 99 · a distância · mage, support

- **P · Iluminação** — As habilidades de dano de Lux carregam o alvo com energia por alguns segundos. O próximo ataque de Lux incendeia a energia, causando Dano Mágico adicional (com base no nível de Lux) ao alvo.
- **Q · Ligação da Luz** — Lux dispara uma esfera de luz que Enraíza os dois primeiros inimigos atingidos por 2/2/2/2/2s e causa 80/120/160/200/240 + 0.75 ability_power de Dano Mágico a cada um.(?)
- **W · Barreira Prismática** — Lux arremessa o báculo e concede 40/55/70/85/100 + 0.4 ability_power de Escudo por 2.5/2.5/2.5/2.5/2.5s aos aliados atravessados. Depois, o báculo retorna, concedendo na volta a mesma quantidade de Escudo que concedeu na ida.(?)
- **E · Singularidade Lucente** — Lux cria uma zona iluminada que causa 25/30/35/40/45% de Lentidão e revela a área. Depois de 5/5/5/5/5s ou ao Reconjurar a Habilidade, a zona é detonada e causa 65/115/165/215/265 + 0.8 ability_power de Dano Mágico e Lentidão por mais 1/1/1/1/1s.(?)
- **R · Centelha Final** — Lux dispara um raio de luz ofuscante, causando 300/400/500 + 1.2 ability_power de Dano Mágico a todos os inimigos em linha reta.(?)

## Xerath, o Mago Ascendente

`Xerath` · id 101 · a distância · mage, support

- **P · Oscilação de Mana** — Os ataques básicos de Xerath restauram Mana periodicamente. Sempre que Xerath abate uma unidade, esse Tempo de Recarga é reduzido.
- **Q · Pulso Arcano** — Início do carregamento: Xerath começa a carregar um raio arcano, sofrendo Lentidão gradual até atingir 50%. Liberar: Xerath dispara o raio, causando 75/115/155/195/235 + 0.9 ability_power de Dano Mágico. O alcance aumenta de acordo com o tempo de carregamento.(?)
- **W · Olho da Destruição** — Xerath conjura uma explosão de energia arcana, causando 50/85/120/155/190 + 0.65 ability_power de Dano Mágico e 25/25/25/25/25% de Lentidão por 2.5/2.5/2.5/2.5/2.5s. Inimigos no centro da explosão sofrem 83.350006/141.695/200.04001/258.385/316.73 + 1.08355 ability_power de Dano Mágico e 60/65/70/75/80% de Lentidão, decaindo ao longo de 2.5/2.5/2.5/2.5/2.5s.(?)
- **E · Orbe Eletrizante** — Xerath dispara um orbe de pura magia, Atordoando o primeiro inimigo atingido por até 2.25/2.25/2.25/2.25/2.25s, com base na distância percorrida, e causando 70/100/130/160/190 + 0.45 ability_power de Dano Mágico. (?)
- **R · Ritual Arcano** — Xerath ascende para a sua verdadeira forma e canaliza por 10/10/10s. Durante esse tempo, ele pode Reconjurar até 4/5/6 vezes. Reconjuração: Xerath lança um disparo mágico, causando 170/220/270 + 0.45 ability_power de Dano Mágico. Para cada Campeão atingido, o disparo causa 20/25/30 + 0.05 ability_power de Dano Mágico adicional.(?)

## Shyvana, a Meio-Dragão

`Shyvana` · id 102 · corpo a corpo · fighter, tank

- **P · Armadura de Escamas** — Eliminar Campeões inimigos, tropas grandes e monstros grandes concede a Shyvana acúmulos de Armadura de Escamas, melhorando suas defesas.
- **Q · Golpe de Brasas** — Passivo: os Ataques de Shyvana causam 0.01/0.01/0.01/0.01/0.01 + 0.00011 attack_damage (bonus) de Dano Mágico com base na Vida máxima %i:OnHit% ao contato e reduzem em 0.5/0.5/0.5/0.5/0.5s o Tempo de Recarga desta Habilidade. Ativo: o próximo Ataque de Shyvana atinge o alvo e a área ao redor, causando 10/15/20/25/30 + 0.3 ability_power + 1.1 attack_damage de Dano Físico. Esta Habilidade pode ser reconjurada após um Ataque ou um breve intervalo nos próximos 4/4/4/4/4s. Forma de Dragão: Shyvana ganha uma reconjuração adicional, fazendo com que ela morda o alvo do próximo Ataque e cause 15/22.5/30/37.5/45 + 0.45000002 ability_power + 1.6500001 attack_damage de Dano Verdadeiro.(?)
- **W · Égide Infernal** — Shyvana se envolve em chamas por 2.5/2.5/2.5/2.5/2.5s, recebendo 60/80/100/120/140 + 0.12 max_health (bonus) de Escudo, que aumenta em 18/24/30.000002/36/42 + 0.036000002 max_health (bonus) por Campeão inimigo próximo, e 0.25/0.25/0.25/0.25/0.25 de Velocidade de Movimento, que aumenta para 0.4375/0.4375/0.4375/0.4375/0.4375 ao se mover em direção a Campeões inimigos. Quando o efeito termina, o Escudo é quebrado ou, ao reconjurar, a área ao redor dela explode, causando 80/100/120/140/160 + 0.65 ability_power de Dano Mágico. Forma de Dragão: caso a detonação atinja um Campeão inimigo, Shyvana cura 100/100/100/100/100 + 0.08/0.08/0.08/0.08/0.08 de Vida perdida.(?)
- **E · Explosão Incandescente** — Shyvana lança uma bola de fogo no local-alvo, causando 50/65/80/95/110 + 0.6 ability_power + 0.05/0.05/0.05/0.05/0.05 da Vida máxima como Dano Mágico e 0.3/0.3/0.3/0.3/0.3 de Lentidão por 2/2/2/2/2s. Explode ao atingir um inimigo ou ao chegar ao destino final. Forma de Dragão: a bola de fogo da Shyvana fica maior e atravessa os inimigos. Ao atingir um Campeão, monstro grande ou seu destino final, a habilidade aciona um pulso de fogo que causa 62.5/81.25/100/118.75/137.5 + 0.75 ability_power + 0.0625/0.0625/0.0625/0.0625/0.0625 da Vida máxima como Dano Mágico e 0.3/0.3/0.3/0.3/0.3 de Lentidão por 2/2/2/2/2s. Por 2/2/2/2/2s, a bola de fogo deixa um rastro que causa 25/25/25/25/25 + 0.05 ability_power de Dano Mágico por segundo. (?)
- **R · Descida do Dragão** — Passivo: Shyvana gera 1.25/1.25/1.25 de Fúria do Dragão ao atingir inimigos com Ataques e Habilidades. O efeito aumenta em 2/2/2 na Forma de Dragão e é reduzido em 0.75/0.75/0.75 ao atingir não Campeões em uma área. Ativo: Shyvana se transforma em sua Forma de Dragão, tornando-se Implacável, e voa até o local-alvo cuspindo fogo nos inimigos, causando 150/250/350 + 1 ability_power de Dano Mágico e fazendo com que eles Fujam por 0.75/0.75/0.75s. Enquanto está transformada, Shyvana tem suas Habilidades básicas aprimoradas e ganha 150/250/350 de Vida, tamanho e Alcance de Ataque adicionais. Shyvana perde Fúria do Dragão ao longo do tempo, e sua transformação termina quando ela não tem mais Fúria do Dragão.(?)

## Ahri, a Raposa de Nove Caudas

`Ahri` · id 103 · a distância · mage, assassin

- **P · Furto de Essência** — Depois de abater 9 tropas ou monstros, Ahri se cura. Depois de eliminar um Campeão inimigo, Ahri se cura em uma quantidade ainda maior.
- **Q · Orbe da Ilusão** — Ahri lança seu orbe e puxa de volta, causando 35/60/85/110/135 + 0.5 ability_power de Dano Mágico na ida e 35/60/85/110/135 + 0.5 ability_power de Dano Verdadeiro na volta.(?)
- **W · Fogo de Raposa** — Ahri invoca 3 Fogos de Raposa que perseguem inimigos próximos e causam 40/60/80/100/120 + 0.4 ability_power de Dano Mágico, reduzido a 16/24/32/40/48 + 0.16000001 ability_power do dano após o primeiro. Ela também recebe 40/40/40/40/40% de Velocidade de Movimento que decai ao longo de 2/2/2/2/2s.(?)
- **E · Encanto** — Ahri manda um beijo que Encanta o primeiro inimigo atingido por 1.2/1.35/1.5/1.65/1.8s e causa 80/120/160/200/240 + 0.85 ability_power de Dano Mágico.(?)
- **R · Ímpeto Espiritual** — Ahri avança rapidamente, disparando 3/3/3 raios de essência em inimigos próximos, priorizando Campeões. Cada raio causa 75/125/175 + 0.35 ability_power de Dano Mágico. Ímpeto Espiritual pode ser Reconjurada mais 2 vezes dentro de 15/15/15s. Consumir a Essência de um Campeão com Furto de Essência durante esse período estende o período de reconjuração em 10/10/10s e concede uma reconjuração adicional de Ímpeto Espiritual (limite de armazenamento: 3/3/3).(?)

## Graves, o Foragido

`Graves` · id 104 · a distância · marksman

- **P · Nova Destino** — A escopeta de Graves tem algumas propriedades únicas. Ele precisa recarregar quando sua munição acaba. Ataques disparam 4 projéteis que não podem atravessar unidades. Não Campeões atingidos por mais de um projétil são empurrados para trás.
- **Q · Fim da Linha** — Graves dispara um projétil de pólvora, causando 50/75/100/125/150 + 0.65 attack_damage (bonus) de Dano Físico. Depois de 1s ou ao colidir com um terreno, o projétil é detonado, causando 80/125/170/215/260 + 0.55 attack_damage (bonus) de Dano Físico ao longo do trajeto novamente e aos inimigos próximos à explosão.(?)
- **W · Cortina de Fumaça** — Graves cria uma nuvem de fumaça que dura 4s e causa 50/50/50/50/50% de Lentidão aos inimigos dentro dela, bloqueando a visão deles da área externa. O impacto inicial causa 60/110/160/210/260 + 0.6 ability_power de Dano Mágico.(?)
- **E · Saque Rápido** — Graves avança e recarrega uma Cápsula da escopeta. Ele também recebe um acúmulo por 4/4/4/4/4s (máximo de 8/8/8/8/8 acúmulos) ou dois acúmulos se avançar em direção a um Campeão inimigo. Os acúmulos concedem a ele 7/10/13/16/19 de Armadura e 7/10/13/16/19 de Resistência Mágica. Acúmulos são redefinidos ao causar dano a uma não tropa. Cada acerto de projétil dos Ataques de Graves reduz 0.5/0.5/0.5/0.5/0.5s do Tempo de Recarga desta Habilidade.(?)
- **R · Efeito Colateral** — Graves dispara um cartucho explosivo que o empurra para trás. O cartucho causa 275/425/575 + 1.5 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido. Depois de atingir um Campeão inimigo ou chegar ao seu alcance máximo, o cartucho explode, causando 200/320/440 + 1.2 attack_damage (bonus) de Dano Físico.(?)

## Fizz, o Trapaceiro das Marés

`Fizz` · id 105 · corpo a corpo · assassin, fighter

- **P · Lutador Ligeiro** — Fizz pode se mover através de unidades e sofre uma quantidade fixa de dano reduzido de todas as origens.
- **Q · Ataque do Ouriço** — Fizz avança através do inimigo, causando 0/0/0/0/0 + 1 attack_damage de Dano Físico mais (?) de Dano Mágico.(?)
- **W · Tridente da Pedra do Mar** — Passivo: os Ataques de Fizz fazem com que os inimigos sangrem, causando 30/45/60/75/90 + 0.25 ability_power de Dano Mágico por 3/3/3/3/3s. Ativo: o próximo Ataque de Fizz causa 50/75/100/125/150 + 0.45 ability_power de Dano Mágico adicional. Se o Ataque abater o alvo, Fizz recupera 30/40/50/60/70 de Mana e reduz o Tempo de Recarga da Habilidade para 1/1/1/1/1s. Se não o abater, os Ataques de Fizz causam 20/25/30/35/40 + 0.3 ability_power de Dano Mágico por 5/5/5/5/5s.(?)
- **E · Brincalhão / Trapaceiro** — Fizz salta sobre seu tridente, tornando-se Inalvejável por 0,75s. Depois, causa 80/130/180/230/280 + 0.95 ability_power de Dano Mágico e 40/45/50/55/60% de Lentidão aos inimigos próximos por 2/2/2/2/2s. Fizz pode Reconjurar a Habilidade enquanto estiver Inalvejável para avançar novamente, encerrando o efeito antes do tempo e causando dano em uma área menor (sem causar Lentidão).(?)
- **R · Lançar Isca** — Fizz arremessa um peixe que se prende ao primeiro Campeão atingido. A vítima é afetada por Visão Mágica e 40% - 80% de Lentidão com base na distância percorrida pelo peixe antes de prender-se. Depois de 2/2/2s, um tubarão surge sob o alvo, Arremessando-o ao ar por 1s, Empurrando todo o resto e causando 180/300/420 + 0.6 ability_power - 270/450/630 + 0.90000004 ability_power de Dano Mágico com base na distância percorrida pelo peixe antes de prender-se.(?)

## Volibear, a Tempestade Implacável

`Volibear` · id 106 · corpo a corpo · fighter, tank

- **P · A Tempestade Implacável** — Os Ataques e habilidades de Volibear concedem Velocidade de Ataque e, às vezes, fazem com que seus Ataques causem Dano Mágico adicional a inimigos próximos.
- **Q · Esmagamento Trovejante** — Volibear recebe 0.12/0.155/0.19/0.225/0.26 de Velocidade de Movimento, dobrada para 0.24/0.31/0.38/0.45/0.52 pelos próximos 4/4/4/4/4s quando ele estiver indo em direção a Campeões inimigos. Enquanto a Habilidade estiver ativa, os próximos Ataques de Volibear causam 10/20/30/40/50 + 1.6 attack_damage (bonus) + 1 attack_damage de Dano Físico e Atordoam o alvo por 1/1/1/1/1s. Volibear ficará enfurecido se for Imobilizado por um inimigo antes de Atordoar um alvo, encerrando a Habilidade e redefinindo o Tempo de Recarga dela.(?)
- **W · Fúria Selvagem** — Volibear estraçalha um inimigo, causando 5/30/55/80/105 + 1.1 attack_damage + 0.06 max_health (bonus) de Dano Físico e marcando-o por 8/8/8/8/8s. Caso a Habilidade seja usada em um alvo marcado, o dano aumenta para (?) e Volibear recupera 20/35/50/65/80 mais 0.08/0.11/0.14/0.17/0.2 da Vida perdida.(?)
- **E · Divisor de Céus** — Volibear invoca uma nuvem carregada para disparar um relâmpago, causando 80/110/140/170/200 + 0.7 ability_power mais 11/12/13/14/15% da Vida máxima como Dano Mágico e 40/40/40/40/40% de Lentidão por 2/2/2/2/2s. Se estiver dentro da zona de explosão, Volibear receberá um Escudo de 0/0/0/0/0 + 0.75 ability_power mais 14/14/14/14/14% da Vida máxima por 3/3/3/3/3s.(?)
- **R · Emissário da Tempestade** — Volibear se transforma e salta, recebendo 175/350/525 de Vida e 50/50/50 de Alcance de Ataque pelos próximos 12/12/12s. Ao pousar, ele fende a terra, Desabilitando torres inimigas próximas por 2/3/4s e causando 300/500/700 + 1.25 ability_power + 2.5 attack_damage (bonus) de Dano Físico a elas. Inimigos próximos sofrem 50/50/50% de Lentidão, decaindo ao longo de 1s. Inimigos que estiverem abaixo de Volibear sofrem 300/500/700 + 1.25 ability_power + 2.5 attack_damage (bonus) de Dano Físico.(?)

## Rengar, o Acossador da Alcateia

`Rengar` · id 107 · corpo a corpo · assassin, fighter

- **P · Predador Oculto** — Enquanto estiver em um arbusto, Rengar salta até seu alvo ao usar um ataque básico. Rengar gera Ferocidade sempre que conjura uma habilidade. Ao atingir Ferocidade máxima, sua próxima Habilidade é fortalecida. Abater Campeões inimigos acrescenta troféus ao Colar de Presas de Rengar, concedendo Dano de Ataque adicional.
- **Q · Selvageria** — Os próximos 2 Ataques de Rengar recebem 40/40/40/40/40% de Velocidade de Ataque. O primeiro Ataque causa 20/55/90/125/160 + 1.05 attack_damage de Dano Físico. Ferocidade Máxima: o primeiro Ataque causa 35/35/35/35/35 + 1.2 attack_damage de Dano Físico e concede a Rengar 50/50/50/50/50 de Velocidade de Ataque por 5/5/5/5/5s.(?)
- **W · Rugido de Batalha** — Rengar ruge, causando 50/80/110/140/170 + 0.8 ability_power de Dano Mágico a inimigos próximos e restaurando em forma de Vida 50/50/50/50/50% do dano sofrido nos últimos 1.5/1.5/1.5/1.5/1.5s. Ferocidade Máxima: causa 220/220/220/220/220 + 0.8 ability_power de Dano Mágico e purifica Controles de Grupo que estejam afetando Rengar.(?)
- **E · Boleadeiras** — Rengar arremessa uma boleadeira, causando 55/100/145/190/235 + 0.8 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido, revelando-o e causando 30/45/60/75/90% de Lentidão por 1.75/1.75/1.75/1.75/1.75s. Ferocidade Máxima: causa 50/50/50/50/50 + 0.8 attack_damage (bonus) de Dano Físico e Enraíza o alvo por 1.75/1.75/1.75/1.75/1.75s.(?)
- **R · Furor da Caçada** — Passivo: Rengar também saltará para Atacar quando estiver Camuflado. Ativo: Rengar recebe 40/50/60% de Velocidade de Movimento e Visão Mágica de uma pequena área ao redor do Campeão inimigo mais próximo por 12/16/20s. Depois de 2/2/2s, Rengar fica Camuflado e pode saltar sem estar dentro de um arbusto. Saltar no Campeão inimigo mais próximo causa 0/0/0 + 1 attack_damage de Dano Físico adicional, fragmenta 15/20/25 de Armadura por 4/4/4s e finaliza a Habilidade.(?)

## Varus, a Flecha da Vingança

`Varus` · id 110 · a distância · marksman, mage

- **P · Vingança Viva** — Ao conseguir um abate ou assistência, Varus recebe Dano de Ataque e Poder de Habilidade temporariamente. O efeito é maior se o inimigo for um Campeão.
- **Q · Flecha Perfurante** — Início do carregamento: Varus prepara o próximo disparo, sofrendo @MoveSpeedMod*-100@% de Lentidão. Após 4/4/4/4/4s, se não realizar o disparo, Varus cancelará a Habilidade e restituirá 50/50/50/50/50% do Custo de Mana dela. Liberar: Varus dispara a flecha, causando 53.333603/100.0005/146.6674/193.3343/240.0012 + 0.80000407 attack_damage (bonus) de Dano Físico, que diminui em 15/15/15/15/15% por inimigo atingido (mínimo de 33/33/33/33/33%). O dano e os efeitos de detonação de Arruinar aumentam em até 50/50/50/50/50% de acordo com o tempo de carregamento (máximo de 80/150/220/290/360 + 1.2 attack_damage (bonus) de Dano Físico).(?)
- **W · Aljava da Ruína** — Passivo: os Ataques de Varus causam 4/13/22/31/40 + 0.25 ability_power + 0.15 attack_damage (bonus) de Dano Mágico adicional e aplicam um acúmulo de Arruinar por 6/6/6/6/6s (máx. de 3/3/3/3/3 acúmulos). As outras Habilidades de Varus detonam os acúmulos de Arruinar, causando 0.03/0.035/0.04/0.045/0.05 + 0.00013 ability_power da Vida máxima como Dano Mágico por acúmulo (máximo de 0.089999996/0.105000004/0.12/0.135/0.15 + 0.00039 ability_power de dano com base na Vida máxima). Detonar Arruinar em Campeões e monstros épicos também reduz os Tempos de Recarga das Habilidades básicas em 13/13/13/13/13% do Máximo a cada acúmulo. Ativo: a próxima Flecha Perfurante de Varus causa 0.06/0.08/0.1/0.12/0.14 da Vida perdida como Dano Mágico adicional, podendo aumentar até 0.089999996/0.12/0.15/0.17999999/0.21000001 de dano com base na Vida perdida dependendo do tempo de carregamento.(?)
- **E · Chuva de Flechas** — Varus dispara uma chuva de flechas que causa 60/90/120/150/180 + 0.9 attack_damage (bonus) de Dano Físico e profana o terreno por 4/4/4/4/4s, causando @SlowPercent*-100@% de Lentidão e aplicando 40/40/40/40/40% de Feridas Dolorosas.(?)
- **R · Corrente da Corrupção** — Varus lança um tentáculo de corrupção, Enraizando o primeiro Campeão atingido por 2/2/2s e causando a ele 150/250/350 + 1 ability_power de Dano Mágico. Inimigos Enraizados recebem 3/3/3 acúmulos de Arruinar ao longo da duração. A corrupção se espalha do alvo original para Campeões inimigos não infectados. Caso os alcance, eles sofrem o mesmo dano e são Enraizados.(?)

## Nautilus, o Titã das Profundezas

`Nautilus` · id 111 · corpo a corpo · tank, support

- **P · Âncora Impactante** — O primeiro Ataque de Nautilus a um alvo causa Dano Físico aumentado e Enraíza-o brevemente.
- **Q · Lançar Âncora** — Nautilus arremessa sua âncora à frente. Se acertar um inimigo, ele puxa a si mesmo e ao alvo para perto um do outro, causando 85/130/175/220/265 + 0.9 ability_power de Dano Mágico e Atordoando-o brevemente. Se a âncora atingir um terreno, Nautilus avançará em direção a ele.(?)
- **W · Ira do Titã** — Nautilus recebe 50/60/70/80/90 + 0.08 max_health de Escudo por 6/6/6/6/6s. Enquanto o Escudo persistir, os Ataques de Nautilus causam 30/40/50/60/70 + 0.4 ability_power de Dano Mágico adicional ao alvo e a inimigos próximos por 2s.(?)
- **E · Correnteza** — Nautilus cria três ondas explosivas ao redor de si mesmo. Cada uma causa aos inimigos na área 55/90/125/160/195 + 0.5 ability_power de Dano Mágico e 30/35/40/45/50% de Lentidão que decai ao longo de 1.25/1.25/1.25/1.25/1.25s.(?)
- **R · Carga de Profundidade** — Nautilus dispara uma onda de choque que persegue um Campeão inimigo, causando 150/275/400 + 0.8 ability_power de Dano Mágico, Arremessando-o ao ar e Atordoando-o por 1/1.5/2s. Outros inimigos atingidos pela onda de choque também são Arremessados ao ar e Atordoados, sofrendo 125/175/225 + 0.4 ability_power de Dano Mágico.(?)

## Viktor, o Arauto do Arcano

`Viktor` · id 112 · a distância · mage

- **P · Gloriosa Evolução** — Viktor receberá Fragmentos Hex sempre que abater um inimigo. Com 100 Fragmentos Hex obtidos, ele receberá permanentemente o Aprimoramento de uma habilidade ativa. Depois de aprimorar todas as habilidades básicas, ele pode obter 100 Fragmentos Hex para aprimorar a ultimate.
- **Q · Poder do Sifão** — Viktor explode um inimigo, causando 60/75/90/105/120 + 0.4 ability_power de Dano Mágico e recebendo 140/140/140/140/140 + 0.25 ability_power de Escudo por 2.5/2.5/2.5/2.5/2.5s. O próximo ataque do Viktor em até 4s causa 20/45/70/95/120 + 0.5 ability_power + 1 attack_damage de Dano Mágico. Aprimoramento: Viktor recebe 224/224/224/224/224 + 0.4 ability_power de Escudo e mais 30/30/30/30/30% de Velocidade de Movimento por 2.5/2.5/2.5/2.5/2.5s.(?)
- **W · Campo Gravítico** — Viktor posiciona um dispositivo de aprisionamento gravitacional por 4/4/4/4/4s, causando @SlowPotency*-1@% de Lentidão aos inimigos dentro dele. Inimigos que permanecem dentro do raio de alcance dele por 1,25s ficam Atordoados por 1.5/1.5/1.5/1.5/1.5s. Aprimoramento – Passivo: as habilidades de Viktor causam 20/20/20/20/20% de Lentidão por 1s. (?)
- **E · Raio Hextec** — Viktor dispara um Raio Hextec na direção escolhida, causando 70/110/150/190/230 + 0.5 ability_power de Dano Mágico a qualquer inimigo atingido. Aprimoramento: uma onda de choque segue o Raio Hextec, causando 20/50/80/110/140 + 0.8 ability_power de Dano Mágico. (?)
- **R · Tempestade Arcana** — Viktor conjura uma Tempestade Arcana em uma área por 6.5/6.5/6.5s, causando instantaneamente 100/175/250 + 0.5 ability_power de Dano Mágico e, depois, 65/105/145 + 0.35 ability_power de Dano Mágico por segundo aos inimigos ao redor. A tempestade segue automaticamente os Campeões que tiverem sofrido dano dela recentemente. Reconjuração: Viktor pode mover a tempestade manualmente. Aprimorada: a tempestade se move 25/25/25% mais rapidamente. Se um Campeão que tiver sofrido dano da tempestade morre, ela aumenta de tamanho e duração por 3/3/3s (máximo de 6/6/6 vezes). (?)

## Sejuani, a Fúria do Norte

`Sejuani` · id 113 · corpo a corpo · tank

- **P · Fúria do Norte** — Após ficar fora de combate, Sejuani recebe Armadura Congelada, que concede Armadura, Resistência Mágica e imunidade a reduções de velocidade. Armadura Congelada persiste por um curto período após Sejuani sofrer dano e Sejuani pode causar dano a um inimigo atordoado para estilhaçá-la, causando muito Dano Mágico.
- **Q · Ataque do Ártico** — Sejuani realiza uma investida, causando 90/140/190/240/290 + 0.75 ability_power de Dano Mágico aos inimigos e Arremessando-os ao ar por 0.5/0.5/0.5/0.5/0.5s. A investida se encerra após atingir um Campeão inimigo.(?)
- **W · Ira do Inverno** — Sejuani brande seu mangual, causando 5/15/25/35/45 + 0.3 ability_power + 0.04 max_health de Dano Físico e Empurrando para trás tropas e monstros da selva. Ela então o brande novamente, causando 5/25/45/65/85 + 0.6 ability_power + 0.08 max_health de Dano Físico e uma breve Lentidão aos alvos. Ambos os movimentos aplicam acúmulos de Congelação Permanente.(?)
- **E · Congelamento Permanente** — Passivo: os Ataques de Campeões corpo a corpo aliados que estiverem por perto aplicam um acúmulo a Campeões e monstros da selva. Ativo: Sejuani causa 55/105/155/205/255 + 0.7 ability_power de Dano Mágico ao inimigo-alvo com 4 acúmulos e o Atordoa por 1/1/1/1/1s. (?)
- **R · Prisão Glacial** — Sejuani arremessa sua boleadeira de Gelo Verdadeiro, Atordoando e revelando o primeiro Campeão inimigo que atingir por 1/1/1s e causando 125/150/175 + 0.4 ability_power de Dano Mágico. Se a boleadeira percorrer pelo menos 25% do alcance, ela Atordoa e revela por 1.5/1.5/1.5s. Ela também cria uma tempestade de gelo que causa 80/80/80% de Lentidão aos inimigos próximos por 2/2/2s. Todos os inimigos atingidos sofrem 200/300/400 + 0.8 ability_power de Dano Mágico.(?)

## Fiora, a Grande Duelista

`Fiora` · id 114 · corpo a corpo · fighter, assassin

- **P · Dança da Duelista** — Fiora revelou um Ponto Vital neste Campeão. Se atingir o Ponto Vital, ela restaurará Vida e receberá Velocidade de Movimento.
- **Q · Estocada** — Fiora avança em uma direção e golpeia o inimigo, sentinela ou estrutura mais perto, causando 70/80/90/100/110 + 0.9 attack_damage (bonus) de Dano Físico. O golpe prioriza Pontos Vitais e inimigos que podem ser abatidos. Se Fiora atingir um inimigo, o Tempo de Recarga da Habilidade é reduzido em 50/50/50/50/50%.(?)
- **W · Ripostar** — Fiora bloqueia todo o dano, os desarmes e os efeitos negativos que receberia pelos próximos 0.75/0.75/0.75/0.75/0.75s e depois golpeia. O golpe causa 110/150/190/230/270 + 1 ability_power de Dano Mágico ao primeiro Campeão atingido e Lentidão de Velocidade de Movimento em @MSSlowPercent*-100@% e de Velocidade de Ataque em @AttackSlowPercent*-100@% por 2/2/2/2/2s. Se Fiora bloquear um efeito Imobilizador, o inimigo que sofreu o golpe é Atordoado em vez de sofrer Lentidão.(?)
- **E · Esgrima** — Fiora recebe 50/60/70/80/90% de Velocidade de Ataque pelos próximos dois Ataques. O primeiro Ataque causa @SlowPercent*-100@% de Lentidão por 1/1/1/1/1s. O segundo Ataque sempre é um Acerto Crítico que causa 160/170/180/190/200% de Dano Físico.(?)
- **R · Desafio Grandioso** — Passivo: aumenta a Velocidade de Movimento adicional concedida por Dança da Duelista em 30/40/50%. Ativo: Fiora revela os quatro Pontos Vitais de um Campeão, podendo causar até @spell.FioraPassive:RDamageTotal@ da Vida máxima como Dano Verdadeiro e recebendo a Velocidade de Movimento adicional de Dança da Duelista enquanto estiver perto do alvo. Se Fiora atingir todos os quatro Pontos Vitais dentro de 8/8/8s, ou se o alvo morrer depois que pelo menos um dos Pontos tiver sido atingido, ela restaurará 75/100/125 + 0.6 attack_damage (bonus) de Vida por segundo aos Campeões aliados próximos por 5/5/5s.(?)

## Ziggs, o Especialista em Hexplosivos

`Ziggs` · id 115 · a distância · mage

- **P · Pavio Curto** — Periodicamente, o próximo ataque básico de Ziggs causa Dano Mágico adicional. Esse Tempo de Recarga é reduzido sempre que Ziggs usa uma habilidade.
- **Q · Bomba Saltitante** — Ziggs arremessa uma bomba saltitante que causa 80/130/180/230/280 + 0.6 ability_power de Dano Mágico.(?)
- **W · Carga Concentrada** — Ziggs arremessa uma carga explosiva que detona depois de 4/4/4/4/4s ou quando a Habilidade for Reconjurada. Ela causa 70/105/140/175/210 + 0.5 ability_power de Dano Mágico a inimigos e os Empurra. Ziggs também é lançado para longe, mas não sofre dano. Carga Concentrada destruirá torres com menos de 25/27.5/30/32.5/35% de Vida automaticamente.(?)
- **E · Campo Minado de Hexplosivos** — Ziggs espalha minas de proximidade que detonam ao contato com inimigos, causando 30/70/110/150/190 + 0.25 ability_power de Dano Mágico e @Slow*-100@% de Lentidão por 1.5/1.5/1.5/1.5/1.5s. As minas duram 10/10/10/10/10s.(?)
- **R · Bomba Megainfernal** — Ziggs arremessa sua criação suprema, causando 300/500/700 + 1 ability_power de Dano Mágico no centro do raio da explosão ou 195/325/454.99997 + 0.65 ability_power de Dano Mágico na borda.(?)

## Lulu, a Fada Feiticeira

`Lulu` · id 117 · a distância · support, mage

- **P · Pix, o Silfo Companheiro** — Pix faz disparos de energia mágica sempre que o Campeão que estiver seguindo atacar outra unidade inimiga. Os disparos seguem o inimigo, mas podem ser interceptados por outras unidades.
- **Q · Lança-Purpurina** — Lulu e Pix disparam, cada um, um projétil perfurante que causa 60/95/130/165/200 + 0.5 ability_power de Dano Mágico e @SlowAmount*-100@% de Lentidão, e esse valor decai ao longo de 2/2/2/2/2s. Os inimigos sofrem 60/95/130/165/200 + 0.5 ability_power de Dano Mágico de projéteis adicionais.(?)
- **W · Caprichos** — Quando usado em um aliado, Lulu concede 0.25/0.25/0.25/0.25/0.25 + 0.0005 ability_power de Velocidade de Movimento e 20/22.5/25/27.5/30% de Velocidade de Ataque por 3/3.25/3.5/3.75/4s. Quando usado em um inimigo, Lulu Polimorfa o alvo por 1.2/1.4/1.6/1.8/2s.(?)
- **E · Socorro, Pix!** — Quando usada em um aliado, Pix salta nele e concede Pix, o Silfo Companheiro por 6/6/6/6/6s. Se o aliado for um Campeão, Pix também concede 70/110/150/190/230 + 0.5 ability_power de Escudo por 2.5/2.5/2.5/2.5/2.5s. Quando usada em um inimigo, Pix salta nele e o perturba, causando 70/110/150/190/230 + 0.5 ability_power de Dano Mágico e concedendo Visão Mágica dele por 4/4/4/4/4s.(?)
- **R · Crescimento Virente** — Lulu faz um aliado crescer, Arremessando ao ar inimigos ao redor por 1/1/1s. O aliado gigante recebe 275/425/575 + 0.55 ability_power de Vida máxima e causa 30/45/60% de Lentidão aos inimigos ao redor por 7/7/7s.(?)

## Draven, o Carrasco de Noxus

`Draven` · id 119 · a distância · marksman

- **P · League of Draven** — Draven ganha Adoração de seus fãs quando apanha uma Revolução do Machado ou abate uma tropa, monstro ou torre. Abater Campeões inimigos concede ouro adicional a Draven, com base em quanta Adoração ele tem.
- **Q · Revolução do Machado** — Draven prepara uma Revolução do Machado, fazendo com que seu próximo Ataque cause 40/45/50/55/60 + 0.75 attack_damage (bonus) de Dano Físico adicional e ricocheteie no ar. Se Draven apanhar o machado, ele preparará outra Revolução do Machado. Draven pode ter duas Revoluções do Machado simultaneamente em ação.(?)
- **W · Adrenalina** — Draven recebe efeito Fantasma, 50/55/60/65/70% de Velocidade de Movimento que decai ao longo de 1.5/1.5/1.5/1.5/1.5s e 20/25/30/35/40% de Velocidade de Ataque por 3/3/3/3/3s. Quando Draven apanha uma Revolução do Machado, o Tempo de Recarga desta Habilidade é zerado.(?)
- **E · Sai da Frente** — Draven arremessa um machado lateralmente, causando 75/110/145/180/215 + 0.5 attack_damage (bonus) de Dano Físico, Empurrando e causando 20/25/30/35/40% de Lentidão por 2/2/2/2/2s.(?)
- **R · Reta da Morte** — Draven arremessa dois machados gigantes que causam 200/300/400 + 1.1 attack_damage (bonus) de Dano Físico. Ao atingir um Campeão ou Reconjurar, eles invertem a direção e retornam para Draven. Os machados causam 5/5/5% a menos de dano a cada inimigo atingido, chegando a um mínimo de 50/50/50%. Se Reta da Morte causar dano o bastante para deixar um Campeão inimigo com a Vida menor do que 100/100/100% da quantidade atual de acúmulos de League of Draven ((?)), Draven o executará.(?)

## Hecarim, a Sombra da Guerra

`Hecarim` · id 120 · corpo a corpo · fighter, tank

- **P · Caminho da Guerra** — Hecarim recebe Dano de Ataque equivalente a um percentual de sua Velocidade de Movimento adicional.
- **Q · Enfurecido** — Hecarim golpeia inimigos próximos, causando 60/90/120/150/180 + 0.9 attack_damage (bonus) de Dano Físico. Se a habilidade atinge inimigos, ele ganha um acúmulo que reduz o Tempo de Recarga em 0.75/0.75/0.75/0.75/0.75s e aumenta o dano dela em 3/3/3/3/3 + 0.03 attack_damage (bonus)% por 8/8/8/8/8s. Acumula até 3/3/3/3/3 vezes. (?)
- **W · Espírito do Pavor** — Hecarim causa 80/120/160/200/240 + 0.8 ability_power de Dano Mágico ao longo de 4/4/4/4/4s aos inimigos próximos. Hecarim recebe 5/10/15/20/25 de Armadura e Resistência Mágica e se cura em 25/25/25/25/25% do dano que ele causar aos inimigos próximos e 12.5/12.5/12.5/12.5/12.5% do dano que seus aliados causarem aos inimigos.(?)
- **E · Ataque Devastador** — Hecarim recebe efeito Fantasma e 25/25/25/25/25% de Velocidade de Movimento que aumenta para 65/65/65/65/65% ao longo de 4/4/4/4/4s. O próximo ataque do Campeão Empurra e causa 30/45/60/75/90 + 0.5 attack_damage (bonus) - 60/90/120/150/180 + 1 attack_damage (bonus) de Dano Físico. A distância e o dano do Empurrão aumentam conforme a distância percorrida ao longo da habilidade.(?)
- **R · Massacre das Sombras** — Hecarim invoca cavaleiros espectrais e avança, causando 150/250/350 + 1 ability_power de Dano Mágico a inimigos atingidos. Ao final do avanço, ele libera uma onda de choque que causa Temor por 0.75/0.75/0.75s - 1.5/1.5/1.5s, aumentado com base na distância do avanço.(?)

## Kha'Zix, o Ceifador do Vazio

`Kha'Zix` · id 121 · corpo a corpo · assassin

- **P · Ameaça Invisível** — Inimigos próximos que estiverem isolados de seus aliados são marcados. As habilidades de Kha'Zix possuem interações com alvos isolados. Quando Kha'Zix não estiver visível para a equipe inimiga, ele recebe Ameaça Invisível, fazendo com que seu próximo ataque básico contra um Campeão inimigo cause Dano Mágico adicional e reduza a velocidade do alvo por alguns segundos.
- **Q · Sabor de Medo** — Kha'Zix corta um inimigo próximo, causando @spell.KhazixQ:BaseDamage@ de Dano Físico. O corte causa @spell.KhazixQ:IsoDamage@ de Dano Físico se o inimigo estiver Isolado dos aliados dele. (?)
- **W · Espinho do Vazio** — Kha'Zix dispara um espinho, causando (?) de Dano Físico ao primeiro inimigo atingido em uma pequena área. Se Kha'Zix estiver dentro da área, ele restaura 55/75/95/115/135 + 0.5 ability_power de Vida.(?)
- **E · Pulo** — Kha'Zix salta, causando (?) de Dano Físico ao aterrissar.(?)
- **R · Massacre do Vazio** — Ativo: Kha'Zix fica Invisível por 1.25/1.25/1.25s e recebe 40/40/40% de Velocidade de Movimento. A Habilidade pode ser Reconjurada uma vez dentro de 12/12/12s. Passivo: subir o nível dessa Habilidade permite que uma das Habilidades de Kha'Zix Evolua, concedendo a ela efeitos adicionais.Sabor de Medo: recebe alcance de Habilidade, Alcance de Ataque e reduz o Tempo de Recarga em @spell.KhazixQ:Effect4Amount@% contra alvos Isolados.Espinho do Vazio: dispara 3 espinhos e causa @spell.KhazixW:Effect3Amount@% de Lentidão, aumentada contra alvos Isolados.Pulo: recebe alcance aumentado e redefine o Tempo de Recarga ao eliminar Campeões.Massacre do Vazio: a Invisibilidade dura 2/2/2s e pode ser Reconjurada uma segunda vez.(?)

## Darius, a Mão de Noxus

`Darius` · id 122 · corpo a corpo · fighter, tank

- **P · Hemorragia** — Os Ataques e Habilidades de dano de Darius fazem os inimigos sangrarem, causando Dano Físico ao longo de 5s e acumulando até 5 vezes. Darius se enfurece e recebe uma grande quantidade de Dano de Ataque quando o alvo alcança o máximo de acúmulos.
- **Q · Dizimar** — Darius se prepara e gira seu machado, causando 50/80/110/140/170 + 1 attack_damage de Dano Físico com a lâmina e 17.5/28/38.5/49/59.5 + 0.35 attack_damage de Dano Físico com o cabo. Inimigos atingidos pelo cabo não recebem um acúmulo de Hemorragia. Darius restaura 17/17/17/17/17% da Vida perdida por Campeão inimigo ou monstro grande da selva atingido pela lâmina, até um máximo de 51/51/51/51/51%.(?)
- **W · Ataque Mutilador** — O próximo ataque de Darius causa 0/0/0/0/0 + 1.4 attack_damage de Dano Físico e 90/90/90/90/90% de Lentidão por 1/1/1/1/1s. Caso a habilidade abata o alvo, o Custo de Mana é restituído e o Tempo de Recarga é reduzido em 50/50/50/50/50%. Esta Habilidade ativa efeitos ao causar dano.(?)
- **E · Apreender** — Passivo: Darius recebe 20/25/30/35/40% de Penetração de Armadura. Ativo: Darius maneja seu machado, Puxando, Arremessando ao ar e causando 40/40/40/40/40% de Lentidão por 1/1/1/1/1s.(?)
- **R · Guilhotina de Noxus** — Darius salta na direção de um inimigo e o atinge com um golpe letal, causando 125/250/375 + 0.75 attack_damage (bonus) de Dano Verdadeiro. Para cada acúmulo de Hemorragia no alvo, a Habilidade causa 20/20/20% de dano adicional, até um máximo de 250/500/750 + 1.5 attack_damage (bonus) de Dano Verdadeiro. Se isso abater o alvo, Darius pode Reconjurar a Habilidade uma vez dentro de 20/20/20s. No ranque 3, a Habilidade não tem custo de Mana e abates redefinem completamente o Tempo de Recarga.(?)

## Jayce, o Defensor do Amanhã

`Jayce` · id 126 · corpo a corpo · fighter, marksman

- **P · Capacitor Hextec** — Quando Jayce troca de arma, ele recebe Velocidade de Movimento por um curto período.
- **Q · Aos Céus! / Disparo Chocante** — Martelo de Mercúrio: Jayce salta em um inimigo, causando @spell.JayceToTheSkies:Damage@ de Dano Físico e @spell.JayceToTheSkies:Slow*-100@% de Lentidão aos inimigos ao redor por @spell.JayceToTheSkies:SlowDuration(?)SpellModifierDescriptionAppend@
- **W · Campo Elétrico / Hipercarga** — Martelo de Mercúrio – Passivo: os Ataques do Martelo de Jayce concedem @spell.JayceStaticField:ManaGain@ de Mana. Martelo de Mercúrio – Ativo: Jayce cria uma aura elétrica, causando @spell.JayceStaticField:Damage@ de Dano Mágico ao longo de @spell.JayceStaticField:Duration(?)SpellModifierDescriptionAppend@
- **E · Golpe Trovejante / Portão Acelerador** — Forma de Martelo: Jayce brande o martelo, Empurrando o alvo e causando @spell.JayceThunderingBlow:FlatDamage@ mais @spell.JayceThunderingBlow:PercHPDamage*100@% da Vida máxima como Dano Mágico.(?)
- **R · Canhão de Mercúrio / Martelo de Mercúrio** — Martelo de Mercúrio: Jayce transforma sua arma no Canhão de Mercúrio, recebendo Alcance de Ataque e novas Habilidades. O próximo Ataque de Jayce remove @spell.JayceStanceHtG:RangedFormShred@ de Armadura e de Resistência Mágica por @spell.JayceStanceHtG:ShredDuration(?)SpellModifierDescriptionAppend@

## Lissandra, a Bruxa Gélida

`Lissandra` · id 127 · a distância · mage

- **P · Submissão Glacinata** — Quando um Campeão inimigo morre próximo a Lissandra, ele vira um Servo Congelado. Servos Congelados reduzem a velocidade de inimigos próximos e, depois de um intervalo, estilhaçam devido ao frio intenso, causando Dano Mágico a alvos próximos.
- **Q · Estilhaço de Gelo** — Lissandra arremessa uma lança de gelo que se estilhaça no primeiro inimigo atingido, causando 80/115/150/185/220 + 0.75 ability_power de Dano Mágico e @SlowPercentage*-100@% de Lentidão por 1.5/1.5/1.5/1.5/1.5s. Além disso, quem estiver atrás também sofrerá dano e Lentidão.(?)
- **W · Círculo Ártico** — Lissandra cria um campo de gelo, Enraizando inimigos próximos por 1.25/1.35/1.45/1.55/1.65s e causando 70/105/140/175/210 + 0.7 ability_power de Dano Mágico.(?)
- **E · Caminho Glacial** — Lissandra envia à frente uma garra de gelo, causando 70/105/140/175/210 + 0.6 ability_power de Dano Mágico. A Habilidade pode ser Reconjurada enquanto a garra percorre o trajeto, teleportando Lissandra até o local em que a garra estiver.(?)
- **R · Túmulo Congelado** — Lissandra se protege ou aprisiona um Campeão inimigo dentro de um bloco de gelo. Inimigos são Atordoados por 1.5/1.5/1.5s. Se a Habilidade for conjurada em si mesma, Lissandra entra em Estase por 2.5/2.5/2.5s e restaura 100/150/200 + 0.55 ability_power de Vida, aumentando 1/1/1% a cada 1/1/1% de Vida perdida. Nos dois casos, o alvo propaga gelo sombrio, causando 150/250/350 + 0.75 ability_power de Dano Mágico. O gelo persiste por 3/3/3s, causando @SlowAmount*-100@% de Lentidão aos inimigos.(?)

## Diana, o Escárnio da Lua

`Diana` · id 131 · corpo a corpo · fighter, assassin

- **P · Espada de Prata Lunar** — Cada terceiro Ataque atinge os inimigos próximos, causando Dano Mágico adicional. Depois de conjurar uma Habilidade, Diana recebe Velocidade de Ataque por 5s.
- **Q · Golpe Crescente** — Diana libera um arco de energia lunar, causando 70/105/140/175/210 + 0.7 ability_power de Dano Mágico e marcando inimigos com Plenilúnio por 3/3/3/3/3s. Plenilúnio revela inimigos que não estão em Furtividade.(?)
- **W · Cascata Lívida** — Por 5/5/5/5/5s, Diana cria três esferas orbitantes que explodem ao entrar em contato com inimigos, causando 20/32/44/56/68 + 0.18 ability_power de Dano Mágico, até um máximo de 60/96/132/168/204 + 0.54 ability_power de Dano Mágico. Além disso, ela recebe 45/60/75/90/105 + 0.3 ability_power + 0.11 max_health (bonus) de Escudo pela mesma duração. Quando a última esfera for detonada, Diana receberá 45/60/75/90/105 + 0.3 ability_power + 0.11 max_health (bonus) de Escudo adicional e redefinirá a duração.(?)
- **E · Zênite Lunar** — Diana se torna a encarnação viva da lua vingativa, avançando até um inimigo e causando a ele 50/70/90/110/130 + 0.6 ability_power de Dano Mágico. Se o alvo estiver afetado por Plenilúnio, o Tempo de Recarga da Habilidade é redefinido.(?)
- **R · Colapso Minguante** — Diana Puxa os inimigos próximos, causando 40/50/60% de Lentidão e revelando-os por 2/2/2s. Se Diana atingir pelo menos um Campeão inimigo, ela invoca a lua, causando 200/300/400 + 0.6 ability_power de Dano Mágico e mais 35/60/85 + 0.15 ability_power para cada inimigo puxado além do primeiro, até um máximo de 175/300/425 + 0.75 ability_power de Dano Mágico adicional.(?)

## Quinn, as Asas de Demacia

`Quinn` · id 133 · a distância · marksman, assassin

- **P · Rapina** — Valor, a águia demaciana de Quinn, periodicamente marca inimigos com Rapina. O primeiro Ataque básico de Quinn contra alvos marcados com Rapina causará Dano Físico adicional.
- **Q · Investida Anuviante** — Quinn comanda Valor para que mergulhe, marcando com Rapina o primeiro inimigo atingido, que tem seu raio de visão reduzido por 1.75/1.75/1.75/1.75/1.75s, e causando 65/100/135/170/205 + 0.5 ability_power + 0.8 attack_damage (bonus) de Dano Físico a todos os inimigos ao redor. Se o primeiro inimigo atingido não for um Campeão, ele sofrerá Desarme por 1.75/1.75/1.75/1.75/1.75s.(?)
- **W · Sentidos Apurados** — Passivo: atacar um alvo marcado com Rapina concede a Quinn 28/41/54/67/80% de Velocidade de Ataque e 20/25/30/35/40% de Velocidade de Movimento por 2/2/2/2/2s. Ativo: Valor revela uma grande área próxima por 2/2/2/2/2s.(?)
- **E · Salto** — Quinn avança até um inimigo, causando 40/65/90/115/140 + 0.2 attack_damage (bonus) de Dano Físico e marcando-o com Rapina. Então, ela salta para o sentido oposto do alvo, Empurrando-o brevemente e causando 50/50/50/50/50% de Lentidão, que decai ao longo de 1.5/1.5/1.5/1.5/1.5s.(?)
- **R · Retaguarda do Inimigo** — Quinn convoca Valor, unindo-se a ele após 2s de canalização e recebendo 70/100/130% de Velocidade de Movimento e a capacidade de reconjurar a Habilidade. Atacar ou conjurar Investida Anuviante ou Salto reconjura automaticamente essa Habilidade. Reconjuração: Quinn e Valor realizam uma manobra aérea, causando 60/90/120 + 0.35 attack_damage (bonus) de Dano Físico, marcando Campeões com Rapina e encerrando a Habilidade.(?)

## Syndra, a Soberana Sombria

`Syndra` · id 134 · a distância · mage

- **P · Transcender** — Syndra coleta Farpas de Ira ao subir de nível ou causar dano a inimigos, aprimorando as próprias habilidades. Esfera Negra: Syndra pode segurar uma carga adicional. Força de Vontade: Dano Verdadeiro adicional. Dispersar os Fracos: largura aumentada e causa Lentidão a todos os alvos. Poder Irrestrito: executa alvos com a Vida baixa.
- **Q · Esfera Negra** — Syndra conjura uma Esfera Negra, causando 90/125/160/195/230 + 0.7 ability_power de Dano Mágico. A Esfera Negra dura 6/6/6/6/6s e pode ser manipulada por outras Habilidades de Syndra. @spell.SyndraPassive:Q1UpgradeThreshold@ Farpas de Ira: Syndra pode armazenar até 2/2/2/2/2 cargas de Esfera Negra.(?)
- **W · Força de Vontade** — Syndra apanha uma Esfera Negra, tropa inimiga ou monstro não épico. Ela arrasta consigo o alvo apanhado e pode Reconjurar em até 5s. Reconjuração: Syndra arremessa o alvo apanhado, causando 70/100/130/160/190 + 0.65 ability_power de Dano Mágico e 25/25/25/25/25% de Lentidão por (?)s. @spell.SyndraPassive:WUpgradeThreshold@ Farpas de Ira: a Habilidade causa 12/12/12/12/12 + 0.02 ability_power de Dano Verdadeiro adicional. (?)
- **E · Dispersar os Fracos** — Syndra projeta uma onda de força, Empurrando inimigos e Esferas Negras e causando 60/95/130/165/200 + 0.6 ability_power de Dano Mágico. Esferas Negras impelidas Atordoam inimigos por 1.25/1.25/1.25/1.25/1.25s e causam 60/95/130/165/200 + 0.6 ability_power de Dano Mágico. @spell.SyndraPassive:EUpgradeThreshold@ Farpas de Ira: a Habilidade ganha largura aumentada e causa 70/70/70/70/70% de Lentidão a inimigos por 1.25/1.25/1.25/1.25/1.25s.(?)
- **R · Poder Irrestrito** — Passivo: Esfera Negra recebe 10/10/10 de Aceleração de Habilidade adicional por cada ranque de Poder Irrestrito. Syndra libera todo o seu poder cataclísmico, arremessando em um Campeão inimigo as 3 Esferas Negras que a orbitam e até outras 4 que estejam próximas. Cada Esfera Negra causa 80/120/160 + 0.2 ability_power de Dano Mágico (máximo de 560/840/1120 + 1.4 ability_power de Dano Mágico). @spell.SyndraPassive:RUpgradeThreshold@ Farpas de Ira: a Habilidade executa inimigos com menos de 15/15/15% de Vida.(?)

## Aurelion Sol, o Forjador de Estrelas

`Aurelion Sol` · id 136 · a distância · mage

- **P · Criador Cósmico** — As Habilidades de dano de Aurelion Sol reduzem inimigos a acúmulos de Poeira Estelar, melhorando permanentemente cada uma das Habilidades dele.
- **Q · Sopro de Luz** — Aurelion Sol dispara fogo estelar por até 3.25/3.25/3.25/3.25/9999s, causando 45/60/75/90/105 + 0.55 ability_power de Dano Mágico por segundo ao primeiro inimigo atingido e 50/50/50/50/50% do dano aos inimigos próximos. Cada segundo completo do disparo que atingir o mesmo inimigo causará uma explosão de 60/70/80/90/100 + 0.3 ability_power de Dano Mágico mais (?) da Vida máxima como Dano Mágico, além de absorver 2/2/2/2/2 de Poeira Estelar se for um Campeão.(?)
- **W · Voo Astral** — Aurelion Sol voa em uma direção. Durante o voo, Sopro de Luz não tem Tempo de Recarga nem duração máxima de canalização, e o dano fixo é aumentado em 8/9/10/11/12%. Eliminações de Campeões em até 3/3/3/3/3s após causar dano a eles restituem 90/90/90/90/90% do Tempo de Recarga desta Habilidade. Reconjuração: encerra o voo antecipadamente.(?)
- **E · Singularidade** — Aurelion Sol invoca um buraco negro, causando 10/15/20/25/30 + 0.12 ability_power de Dano Mágico por segundo e Arrastando inimigos em direção ao centro por 5/5/5/5/5s. Inimigos no centro e com menos de (?)% de Vida máxima morrem instantaneamente. O buraco negro absorve Poeira Estelar quando inimigos morrem dentro dele e a cada segundo em que há Campeões inimigos dentro da área.(?)
- **R · Estrela Cadente/Os Céus Caem** — Aurelion Sol arranca uma estrela dos céus e a arremessa na Terra, causando 150/250/350 + 0.75 ability_power de Dano Mágico, Atordoando inimigos por 1/1/1s e absorvendo 5/5/5 de Poeira Estelar para cada Campeão atingido. Obter 75/75/75 de Poeira Estelar transforma a próxima instância de Estrela Cadente em Os Céus Caem. Os Céus Caem: Aurelion Sol arrasta a fúria de uma constelação inteira do cosmo, causando 187.5/312.5/437.5 + 0.9375 ability_power de Dano Mágico a uma área maior, Arremessando ao ar inimigos atingidos por 1/1/1s e emitindo uma onda de choque gigantesca que causa 135/225/315 + 0.67499995 ability_power de Dano Mágico a Campeões e monstros épicos, além de causar Lentidão a todos os inimigos atingidos em 50/50/50% por 1s.(?)

## Kayn, o Ceifador das Sombras

`Kayn` · id 141 · corpo a corpo · fighter, assassin

- **P · A Foice Darkin** — Kayn empunha uma arma ancestral, lutando contra Rhaast, o darkin que vive dentro dela, pelo controle. Ou o Darkin triunfará ou Kayn dominará Rhaast para se tornar o Assassino das Sombras. Darkin: Cura uma porcentagem do dano de habilidades causado a Campeões. Assassino das Sombras: Causa dano adicional nos primeiros segundos em combate com Campeões inimigos.
- **Q · Corte Ceifador** — Kayn avança e gira sua foice, causando 75/105/135/165/195 + 0.85 attack_damage (bonus) de Dano Físico aos inimigos atravessados e, depois, aos inimigos próximos. Darkin: causa 0/0/0/0/0 + 0.65 attack_damage mais 6/6/6/6/6 + 0.035 attack_damage (bonus) da Vida máxima como Dano Físico.(?)
- **W · Alcance da Lâmina** — Kayn ataca com a foice para cima, causando 85/130/175/220/265 + 1.1 attack_damage (bonus) de Dano Físico e @Effect3Amount*-100@% de Lentidão aos inimigos atingidos, decaindo ao longo de (?)s. Assassino das Sombras: o alcance da Habilidade é aumentado e Kayn pode se mover enquanto a utiliza. Darkin: a Habilidade também Arremessa ao ar os inimigos atingidos por (?)s.(?)
- **E · Passo das Sombras** — Kayn recebe (?)% de Velocidade de Movimento, efeito Fantasma e pode se mover através de terrenos por (?)s. Ao entrar em um terreno pela primeira vez, ele restaura 90/100/110/120/130 + 0.45 attack_damage (bonus) de Vida. Ser Imobilizado ou passar mais de (?)s consecutivos fora de um terreno encerra a Habilidade antecipadamente. Assassino das Sombras: recebe (?)% de Velocidade de Movimento, imunidade a Lentidão e reduz o Tempo de Recarga em 10/10/10/10/10s.(?)
- **R · Transgressão do Umbral** — Passivo: Campeões que sofreram dano de Kayn são marcados por 3,15s. Kayn infesta um Campeão marcado, tornando-se Inalvejável. Depois de 2.5/2.5/2.5s ou de Reconjurar, Kayn irrompe, causando 150/250/350 + 1.5 attack_damage (bonus) de Dano Físico ao inimigo. Assassino das Sombras: aumenta o alcance da Habilidade, a distância que Kayn percorre ao sair e redefine o Tempo de Recarga de A Foice Darkin. Darkin: causa 0.15/0.15/0.15 + 0.001 attack_damage (bonus) da Vida máxima como Dano Físico e restaura 0.112500004/0.112500004/0.112500004 + 0.00075 attack_damage (bonus) de Vida (75/75/75% da quantidade de dano).(?)

## Zoe, o Aspecto do Crepúsculo

`Zoe` · id 142 · a distância · mage

- **P · Brililim!** — O próximo ataque básico após Zoe conjurar uma habilidade causa Dano Mágico adicional.
- **Q · Estrela Desviada!** — Zoe dispara uma estrela que causa dano crescente ao primeiro inimigo atingido e aos inimigos ao redor de acordo com a distância percorrida (52/82/112/142/172 + 0.6 ability_power - 130/205/280/355/430 + 1.5 ability_power de Dano Mágico). Zoe pode reconjurar a Habilidade para redirecionar o projétil para uma nova posição próxima a ela.(?)
- **W · Roubo Arcano** — Passivo: os inimigos derrubam fragmentos de feitiço quando conjuram um Feitiço de Invocador ou usam um item ativo. Tropas específicas também deixam cair um fragmento de feitiço quando são abatidas por Zoe ou um aliado próximo. Ela pode coletar o fragmento para conjurar a Habilidade correspondente uma vez. Passivo: quando Zoe conjura esta Habilidade ou qualquer Feitiço de Invocador, ela recebe 30/40/50/60/70% de Velocidade de Movimento por 2/2.25/2.5/2.75/3s e dispara 3 projéteis no último alvo que ela Atacou. Cada projétil causa 15/25/35/45/55 + 0.1 ability_power de Dano Mágico. Ativo: conjura a Habilidade de um fragmento de feitiço que Zoe recolheu.(?)
- **E · Bolha do Soninho** — Zoe chuta uma bolha que causa 70/110/150/190/230 + 0.45 ability_power de Dano Mágico, tornando-se uma armadilha caso não atinja ninguém. O alcance da bolha aumenta quando ela atravessa um terreno. Se a bolha ou a armadilha atinge um Campeão inimigo, ela reduz em 16/19.5/23/26.5/30% o Tempo de Recarga. Depois de um intervalo, a vítima Adormece e sofre 30/30/30/30/30% de redução de Resistência Mágica por 2s. Ataques e Habilidades interrompem o sono e causam o dobro de dano, até 70/110/150/190/230 + 0.45 ability_power de Dano Verdadeiro. (?)
- **R · Salto Dimensional** — Zoe se teleporta para uma posição próxima por 1s e retorna para a posição anterior logo em seguida. Durante esse período, ela pode conjurar Habilidades e Atacar, mas não pode se mover.(?)

## Zyra, a Ascensão dos Espinhos

`Zyra` · id 143 · a distância · mage, support

- **P · Jardim de Espinhos** — Sementes surgem periodicamente ao redor de Zyra, mais rapidamente dependendo do nível. Zyra pode conjurar Farpas Mortais ou Pântano das Raízes próximo a sementes para fazer surgir plantas que lutarão por ela.
- **Q · Farpas Mortais** — Zyra faz com que vinhas grossas se espalhem e explodam em espinhos, causando 60/100/140/180/220 + 0.65 ability_power de Dano Mágico. Se a Habilidade for conjurada perto de uma Semente, a Semente se transformará em um Cospe-Espinhos que causa @spell.ZyraP:PlantDamage@ de Dano Mágico e dura @spell.ZyraP:PlantDuration@s. Os Cospe-Espinhos têm 575 de alcance.(?)
- **W · Crescimento Desenfreado** — Zyra planta uma Semente que dura 60/60/60/60/60s. Essas Sementes concedem Visão Mágica dos Campeões inimigos que pisam nelas por 2/2/2/2/2s, mas são destruídas. Esta Habilidade tem 2 cargas e (?)s de Tempo de Recarga. Abater um monstro ou tropa inimiga reduz o tempo de recarga em 35/35/35/35/35%. Eliminar Campeões reduz o tempo de recarga em 100/100/100/100/100%.(?)
- **E · Pântano das Raízes** — Zyra lança vinhas à frente, Enraizando por 1/1.25/1.5/1.75/2s e causando 60/95/130/165/200 + 0.6 ability_power de Dano Mágico. Se a Habilidade passar perto de uma Semente, a Semente se transforma em uma Vinha Áspera que causa @spell.ZyraP:PlantDamage@ de Dano Mágico e dura @spell.ZyraP:PlantDuration@s. Vinhas Ásperas têm 400 de alcance e seus Ataques causam 30/30/30/30/30% de Lentidão por 2/2/2/2/2s. A Lentidão causada por várias Vinhas Ásperas acumula até 2/2/2/2/2 vezes.(?)
- **R · Espinhos Sufocantes** — Zyra invoca a fúria da natureza e faz surgir uma grande raiz retorcida, causando 200/300/400 + 0.7 ability_power de Dano Mágico. Após 2 segundos, as vinhas rebatem para cima, Arremessando ao Ar por 1/1/1s. As plantas de Zyra dentro do alcance da raiz ficam enfurecidas, redefinindo suas durações, recebendo 50/50/50% de Vida e causando 50/50/50% de dano adicional.(?)

## Kai'Sa, a Filha do Vazio

`Kai'Sa` · id 145 · a distância · marksman, mage

- **P · Segunda Pele** — Os ataques básicos de Kai'Sa acumulam Plasma, causando Dano Mágico adicional crescente. Efeitos imobilizadores de aliados ajudam a acumular Plasma. Além disso, as aquisições de itens de Kai'Sa aprimoram suas habilidades básicas, deixando-as mais poderosas.
- **Q · Chuva Icathiana** — Kai'Sa lança (?) projéteis que se dividem entre os inimigos próximos, cada um causando 40/55/70/85/100 + 0.2 ability_power + 0.55 attack_damage (bonus) de Dano Físico, até um máximo de 150/206.25/262.5/318.75/375 + 0.75 ability_power + 2.0625 attack_damage (bonus). Acertos de projéteis adicionais em Campeões ou monstros causam 25/25/25/25/25% do dano. Evolução: Kai'Sa dispara (?) projéteis. Valor atual: (?)/(?) de Dano de Ataque adicional(?)
- **W · Exploradora do Vazio** — Kai'Sa dispara uma explosão do Vazio que causa (?) de Dano Mágico, aplica (?) acúmulos de Plasma e concede Visão Mágica do primeiro inimigo atingido por @spell.KaisaPassive:PDuration@s. Evolução: Kai'Sa aplica (?) acúmulos de Plasma, e atingir um Campeão reduz o Tempo de Recarga em (?)%. Valor atual: (?)/(?) de Poder de Habilidade(?)
- **E · Sobrecarga** — Kai'Sa se sobrecarrega com energia do Vazio, recebendo (?) de Velocidade de Movimento e efeito Fantasma durante o carregamento, além de (?)% de Velocidade de Ataque por (?)s. Ataques reduzem o Tempo de Recarga desta Habilidade em (?)s. Evolução: Kai'Sa também fica Invisível por (?)s. Valor atual: (?)%/(?)% de Velocidade de Ataque adicional(?)
- **R · Instinto Assassino** — Kai'Sa se teletransporta para perto de um Campeão inimigo afetado por Plasma e recebe 100/150/200 + 1.2 ability_power + 0.9 attack_damage de Escudo por 2/2/2s.(?)

## Seraphine, a Cantora Sonhadora

`Seraphine` · id 147 · a distância · support, mage

- **P · Presença de Palco** — Cada terceira habilidade básica será conjurada duas vezes por Seraphine. Conjurar habilidades próximas a aliados também concede dano e alcance adicionais a ela no próximo ataque básico.
- **Q · Nota Aguda** — Seraphine projeta uma nota pura, causando 60/85/110/135/160 + 0.4 ability_power de Dano Mágico, que aumenta contra Campeões e monstros de acordo com a porcentagem de Vida perdida do alvo até 105/148.75/192.5/236.25/280 + 0.7 ability_power de dano quando abaixo de 25/25/25/25/25% de Vida.(?)
- **W · Som Envolvente** — Seraphine incentiva os Campeões aliados próximos com uma canção, concedendo 0.080000006/0.080000006/0.080000006/0.080000006/0.080000006 + 0.00008 ability_power de Velocidade de Movimento aos aliados, 0.2/0.2/0.2/0.2/0.2 + 0.0002 ability_power de Velocidade de Movimento a si e 60/80/100/120/140 + 0.2 ability_power de Escudo a todos por 2.5/2.5/2.5/2.5/2.5s. Se Seraphine já estiver protegida por um Escudo, convidará os aliados a se juntarem a ela, restaurando 8/10/12/14/16% da Vida perdida deles após 2.5/2.5/2.5/2.5/2.5s de intervalo.(?)
- **E · Ritmo Contagiante** — Seraphine libera uma poderosa onda de som, causando 70/100/130/160/190 + 0.5 ability_power de Dano Mágico a inimigos em uma linha reta e 99/99/99/99/99% de Lentidão por 1.1/1.2/1.3/1.4/1.5s. Inimigos que já estejam sob Lentidão serão Enraizados. Caso já estejam Imobilizados, serão Atordoados.(?)
- **R · Bis** — Seraphine assume o palco e projeta uma força cativante que Encanta inimigos por 1.25/1.5/1.75s e causa 150/200/250 + 0.4 ability_power de Dano Mágico. Campeões atingidos (até mesmo aliados) se tornarão parte da apresentação, estendendo o alcance da Habilidade. Campeões aliados recebem Notas máximas.(?)

## Gnar, o Yordle Pré-Histórico

`Gnar` · id 150 · corpo a corpo · fighter, tank

- **P · Fúria Genética** — Enquanto está em combate, Gnar gera Fúria. Quando atinge o máximo de Fúria, sua próxima habilidade o transformará em Mega-Gnar, concedendo mais condições de sobrevivência e acesso a novas habilidades.
- **Q · Bumerangue / Pedregulho** — Mini-Gnar: Gnar arremessa um bumerangue que causa @spell.GnarQ:MiniTotalDamage@ de Dano Físico e @spell.GnarQ:SlowAmount*100@% de Lentidão aos inimigos por @spell.GnarQ:SlowDuration@s. O bumerangue volta depois de atingir um inimigo, causando dano reduzido a alvos subsequentes. Cada inimigo pode ser atingido somente uma vez. Apanhar o bumerangue reduz o Tempo de Recarga em @spell.GnarQ:MiniCDRefund*100@%.(?)
- **W · Hiperativo / Safanão** — Mini-Gnar – Passivo: cada terceiro Ataque ou Habilidade contra o mesmo inimigo causa @spell.GnarW:MiniTotalDamage@ mais @spell.GnarW:MiniPercentHPDamage*100@% da Vida máxima como Dano Mágico e concede @spell.GnarR:RHyperMovementSpeedPercent@% de Velocidade de Movimento que decai ao longo de @spell.GnarW:MiniHasteDuration(?)SpellModifierDescriptionAppend@
- **E · Pirueta / Encontrão** — Mini-Gnar: Gnar dá uma pirueta, recebendo @spell.GnarE:MinibAS*100@% de Velocidade de Ataque por @spell.GnarE:MiniASDuration@s. Se Gnar aterrissar em uma unidade, ele dá outra pirueta, indo mais longe. Pular em um inimigo causa @spell.GnarE:MiniTotalDamage@ de Dano Físico e @spell.GnarE:MoveSpeedMod*-100@% de Lentidão por um breve período.(?)
- **R · GNAR!** — Mini-Gnar – Passivo: aumenta a Velocidade de Movimento de Hiperativo. Mega-Gnar: Gnar arremessa inimigos próximos, Empurrando-os e causando a eles 200/300/400 + 1 ability_power + 0.5 attack_damage (bonus) de Dano Físico e 45/45/45% de Lentidão por 1.25/1.5/1.75s. Inimigos que atingirem um terreno sofrerão 300/450/600 + 1.5 ability_power + 0.75 attack_damage (bonus) de Dano Físico e serão Atordoados.(?)

## Zac, a Arma Secreta

`Zac` · id 154 · corpo a corpo · tank, fighter

- **P · Divisão Celular** — Sempre que Zac atinge um inimigo com uma habilidade, ele perde parte de si mesmo que pode ser reabsorvida para restaurar Vida. Ao receber dano letal, Zac divide-se em 4 pedaços que tentam se recombinar. Se qualquer pedaço permanecer, ele ressuscitará uma quantia de Vida, dependendo da vida dos pedaços sobreviventes. Cada pedaço tem um percentual da Vida máxima de Zac, de sua Armadura e de Resistência Mágica. Essa habilidade tem 5 minutos de Tempo de Recarga.
- **Q · Esticada** — Zac estica um braço que se prende ao primeiro inimigo atingido, causando 60/90/120/150/180 + 0.3 ability_power + 0.03 max_health (bonus) de Dano Mágico e uma breve Lentidão. O próximo Ataque de Zac ganha um alcance maior e causa a mesma quantidade de dano e Lentidão. Se Zac atingir um inimigo diferente com o Ataque, ele Arremessará ao ar ambos os alvos na direção um do outro. Caso colidam, eles e outros inimigos ao redor sofrerão 60/90/120/150/180 + 0.3 ability_power + 0.03 max_health (bonus) de Dano Mágico e uma breve Lentidão.(?)
- **W · Matéria Instável** — O corpo de Zac entra em erupção, causando 40/50/60/70/80 + 4/5/6/7/8 + 0.03 ability_power da Vida máxima como Dano Mágico a inimigos próximos. Absorver Gosma reduz o Tempo de Recarga da Habilidade em 1s.(?)
- **E · Estilingue Elástico** — Início do carregamento: Zac se estica para trás, carregando um avanço por 0.9/1/1.1/1.2/1.3s. Liberar: Zac se catapulta e Arremessa ao ar inimigos onde aterrissa por até 1/1/1/1/1s (com base no tempo de carregamento), causando 60/105/150/195/240 + 0.8 ability_power de Dano Mágico. Zac faz surgir um pedaço adicional de Gosma para cada Campeão inimigo atingido.(?)
- **R · Vamos pular!** — Zac ricocheteia 4/4/4 vezes. O primeiro pulo a atingir cada inimigo os Arremessa para trás e causa 120/190/260 + 0.4 ability_power de Dano Mágico. Pulos subsequentes causam 60/95/130 + 0.2 ability_power de Dano Mágico e 20/20/20% de Lentidão por 1/1/1s. Zac recebe até 50/50/50% de Velocidade de Movimento ao longo do tempo e pode conjurar Matéria Instável enquanto pula.(?)

## Yasuo, o Imperdoável

`Yasuo` · id 157 · corpo a corpo · fighter, assassin

- **P · Estilo do Errante** — A Chance de Acerto Crítico do Yasuo é aumentada. Além disso, Yasuo vai gerando um Escudo sempre que se movimenta. O Escudo é ativado quando Yasuo sofre dano de um Campeão ou monstro.
- **Q · Tempestade de Aço** — Yasuo golpeia, causando 20/45/70/95/120 + 1.05 attack_damage de Dano Físico. Se acertar um inimigo, recebe um acúmulo por 6/6/6/6/6s. Enquanto tiver 2 acúmulos, seu próximo uso da Habilidade desfere um tornado à distância, causando o mesmo dano a inimigos e Arremessando-os ao ar por 1/1/1/1/1s. Se utilizada enquanto avança, a Habilidade atinge em um círculo em vez de desferir uma estocada.(?)
- **W · Parede de Vento** — Yasuo cria uma parede movediça que bloqueia todos os projéteis inimigos por 4s.(?)
- **E · Espada Ágil** — Yasuo avança e atravessa um alvo, causando 70/85/100/115/130 + 0.6 ability_power + 0.2 attack_damage (bonus) de Dano Mágico. Cada uso da Habilidade concede 17.5/21.25/25/28.75/32.5 + 0.15 ability_power + 0.05 attack_damage (bonus) de dano adicional para usos subsequentes por 5/5/5/5/5s, acumulando até 4/4/4/4/4x. Esta Habilidade tem 10/9/8/7/6s de Tempo de Recarga por alvo.(?)
- **R · Último Suspiro** — Yasuo teleporta-se para um Campeão inimigo Arremessado ao ar, causando 200/350/500 + 1.5 attack_damage (bonus) de Dano Físico e mantendo todos os inimigos da área que foram Arremessados ao ar por mais 1/1/1s. Ele também recebe Fluxo máximo, mas perde todos os acúmulos de Tempestade de Aço. Depois, seus Acertos Críticos ignoram 60/60/60% de Armadura adicional por 15/15/15s.(?)

## Vel'Koz, o Olho do Vazio

`Vel'Koz` · id 161 · a distância · mage, support

- **P · Desconstrução Orgânica** — As habilidades de Vel'Koz aplicam Desconstrução Orgânica nos inimigos ao contato. Ao chegar a 3 acúmulos, o inimigo receberá uma rajada de Dano Verdadeiro.
- **Q · Fissão Plasmática** — Vel'Koz dispara um projétil de plasma que causa 80/120/160/200/240 + 0.9 ability_power de Dano Mágico e 70/70/70/70/70% de Lentidão, que decai ao longo de 1/1.4/1.8/2.2/2.6s. Quando chega ao alcance máximo, seja ao atingir um inimigo ou ao ser Reconjurado, o projétil se divide em dois, traçando uma trajetória perpendicular. Abater uma unidade com a Habilidade restaura 20/22.5/25/27.5/30 de Mana.(?)
- **W · Fenda do Vazio** — Vel'Koz abre uma fenda ao Vazio, causando 30/50/70/90/110 + 0.2 ability_power de Dano Mágico. A fenda detona, causando 45/75/105/135/165 + 0.25 ability_power de Dano Mágico. A Habilidade tem 2 cargas (recarrega a cada (?)s).(?)
- **E · Ruptura Tectônica** — Vel'Koz rompe uma área próxima, fazendo com que ela exploda, Arremessando ao ar por 0.75/0.75/0.75/0.75/0.75s e causando 70/100/130/160/190 + 0.3 ability_power de Dano Mágico. Inimigos próximos a ele são Empurrados em vez de Arremessados ao ar.(?)
- **R · Raio Desintegrador de Formas de Vida** — Vel'Koz canaliza um raio de energia que segue o cursor, causando até (?) de Dano Mágico por 2,5s e (?)% de Lentidão. Causa Dano Verdadeiro ao atingir inimigos que sofreram dano de Desconstrução Orgânica recentemente. Inimigos ganham acúmulos periódicos de Desconstrução Orgânica enquanto estiverem dentro do raio.(?)

## Taliyah, a Tecelã de Pedras

`Taliyah` · id 163 · a distância · mage, support

- **P · Deslizar em Pedras** — Taliyah recebe Velocidade de Movimento próxima a paredes.
- **Q · Voleio Entrelaçado** — Taliyah arremessa 5 pedras. Cada uma causa 55/72.5/90/107.5/125 + 0.5 ability_power de Dano Mágico em uma área ao redor do primeiro inimigo atingido e cria Terreno Manipulado naquela área. Acertos seguintes contra o mesmo inimigo causarão 60/60/60/60/60% a menos de dano. Conjurações em terreno manipulado custam 10/10/10/10/10 de Mana, têm 50/50/50/50/50% a menos de Tempo de Recarga, consomem o Terreno Manipulado e arremessam um pedregulho que causa 20/25/30/35/40% de Lentidão por 1.5/1.5/1.5/1.5/1.5s aos inimigos atingidos e 99/130.5/162/193.5/225 + 0.9 ability_power de Dano Mágico ao alvo primário. Monstros atingidos pelo pedregulho são Atordoados por 3/3/3/3/3s.(?)
- **W · Empurrão Sísmico** — Taliyah agita a terra, Arremessando inimigos dentro de uma área na direção escolhida.(?)
- **E · Terra Desfiada** — Taliyah espalha pedras soltas em uma área, causando 60/105/150/195/240 + 0.6 ability_power de Dano Mágico e 20/20/20/20/20% de Lentidão aos inimigos atingidos. As pedras detonam quando inimigos avançam ou são Arremessados no meio delas, Atordoando pelo restante da duração do movimento + 0.75/0.75/0.75/0.75/0.75s e causando 25/40/55/70/85 + 0.3 ability_power de Dano Mágico.(?)
- **R · Muro da Tecelã** — Taliyah cria uma enorme parede de terra por 4/4/4s. Se Reconjurar imediatamente, Taliyah também percorre a parede conforme ela é criada. Mover-se ou sofrer imobilização interromperá o trajeto. A habilidade não pode ser conjurada se Taliyah tiver sofrido dano de Campeões ou estruturas nos últimos 3/3/3s.(?)

## Camille, a Sombra de Aço

`Camille` · id 164 · corpo a corpo · fighter, assassin

- **P · Defesa Adaptativa** — Ataques básicos contra Campeões concedem um Escudo equivalente a uma porcentagem da Vida máxima de Camille contra o tipo de dano deles (Físico ou Mágico) por um breve período.
- **Q · Protocolo de Precisão** — O próximo Ataque de Camille causa 0/0/0/0/0 + 0.2 attack_damage de Dano Físico adicional e concede a ela 25/30/35/40/45% de Velocidade de Movimento por 1/1/1/1/1s. A habilidade pode ser Reconjurada nos próximos 3.5/3.5/3.5/3.5/3.5s. Se o Ataque da Reconjuração atingir um alvo em ao menos 1.5/1.5/1.5/1.5/1.5s depois do primeiro, o dano adicional é aumentado para 0/0/0/0/0 + 0.4 attack_damage, e 0.4/0.4/0.4/0.4/0.4 do dano do Ataque é convertido em Dano Verdadeiro. Esta Habilidade ativa efeitos ao causar dano.(?)
- **W · Varredura Tática** — Camille se prepara e depois golpeia, causando 60/85/110/135/160 + 0.6 attack_damage (bonus) de Dano Físico. Inimigos atingidos pela metade exterior sofrem 80/80/80/80/80% de Lentidão, que decai ao longo de 2/2/2/2/2s, e sofrem 0.07/0.075/0.08/0.085/0.09 + 0.00025 attack_damage (bonus) da Vida máxima como Dano Físico adicional. Camille restaura 100/100/100/100/100% do dano adicional causado a Campeões como Vida.(?)
- **E · Disparo de Gancho** — Camille dispara um gancho que se prende a um terreno, puxando-se na direção dele por 1s e podendo Reconjurar. Reconjuração: Camille avança a partir da parede, colidindo com o primeiro Campeão inimigo atingido. Ao aterrissar, ela recebe 40/45/50/55/60% Velocidade de Ataque por 5/5/5/5/5s, causa 60/90/120/150/180 + 0.75 attack_damage (bonus) de Dano Físico a inimigos próximos e Atordoa Campeões inimigos por 0.75/0.75/0.75/0.75/0.75s. Avanços em direção a Campeões inimigos têm alcance dobrado.(?)
- **R · Ultimato Hextec** — Camille fica brevemente Inalvejável e salta sobre um Campeão inimigo, interrompendo canalizações e prendendo-o em uma área da qual ele não pode escapar durante 2.5/3.25/4s. Outros inimigos próximos são Empurrados. Os Ataques dela contra o inimigo preso causam 4/6/8% da Vida atual como Dano Mágico adicional.(?)

## Akshan, o Sentinela Rebelde

`Akshan` · id 166 · a distância · marksman, assassin

- **P · Lutando Sujo** — Os Ataques e Habilidades de Akshan causam dano adicional a cada terceiro acerto e concedem a ele um Escudo se o alvo for um Campeão. Ao Atacar, ele desfere um Ataque adicional que causa dano reduzido. Se cancelar o Ataque adicional, ele recebe Velocidade de Movimento.
- **Q · Bumerangue Vingativo** — Akshan arremessa um bumerangue que causa 45/75/105/135/165 + 0.7 attack_damage (bonus) de Dano Físico e ganha alcance cada vez que atinge um alvo. Campeões atingidos concedem a Akshan 0.2/0.2/0.2/0.2/0.2 + 0.0005 ability_power de Velocidade de Movimento, que decai ao longo de 1/1/1/1/1s.(?)
- **W · Rebeldia** — {{ Spell_AkshanW_Tooltip_1/1/1/1/1 }}(?)
- **E · Impulso Heroico** — Primeira conjuração: Akshan dispara um arpéu que se prende ao primeiro terreno atingido. Segunda conjuração: pendurado, Akshan se balança pelo terreno, disparando repetidamente contra o inimigo mais próximo e causando 8/16/24/32/40 + 0.25 attack_damage de Dano Físico por disparo. Terceira conjuração: Akshan salta da corda e faz um último disparo. Colidir com um Campeão inimigo ou terreno interrompe o balanço mais cedo. Eliminar Campeões redefine o Tempo de Recarga desta Habilidade.(?)
- **R · Punição** — Akshan mira em um Campeão e começa a sobrecarregar a arma por até 2.5/2.5/2.5s, armazenando um máximo de 5/6/7 balas. Reconjuração: Akshan dispara as balas armazenadas, e cada uma causa um mínimo de 25/35/45 + 0.15 attack_damage de Dano Físico ao primeiro inimigo ou estrutura que atingir, podendo aumentar para até 75/105/135 + 0.45000002 attack_damage de Dano Físico com base na Vida perdida do alvo.(?)

## Bel'Veth, a Imperatriz do Vazio

`Bel'Veth` · id 200 · corpo a corpo · fighter

- **P · Morte em Lavanda ** — Bel'Veth recebe acúmulos de Velocidade de Ataque permanente depois de eliminar tropas, monstros grandes e Campeões. Além disso, ela recebe Velocidade de Ataque adicional temporária depois de usar uma Habilidade.
- **Q · Impulso do Vazio** — Bel'Veth avança, causando 12/14/16/18/20 + 1.05 attack_damage de Dano Físico aos inimigos que atravessar. Cada direção tem um Tempo de Recarga único de (?)s (reduzido com base na Velocidade de Ataque). (?)
- **W · Acima e Abaixo** — Bel'Veth bate com a cauda no chão, causando 80/140/200/260/320 + 1.5 ability_power (bonus) de Dano Mágico, Arremessando ao ar por 0.6/0.7/0.8/0.9/1s e causando 30/30/30/30/30% de Lentidão a inimigos atingidos por 2/2/2/2/2s. Se um Campeão for atingido, o Tempo de Recarga de Impulso do Vazio é redefinido naquela direção.(?)
- **E · Turbilhão da Realeza** — Bel'Veth canaliza cortes em volta de si, recebendo 20/30/40/50/60% de redução de dano, 0.2/0.25/0.3/0.35/0.4 de Roubo de Vida e atacando (?) vezes ao longo de 1.5/1.5/1.5/1.5/1.5s, aumentando o número de ataques de acordo com a Velocidade de Ataque. Cada Ataque atinge o inimigo com a Vida mais baixa, causando 10/12/14/16/18 + 0.12 attack_damage - 20/24/28/32/36 + 0.24 attack_damage de Dano Físico com base na Vida perdida do alvo. Usar outra Habilidade ou Reconjurar finaliza esta Habilidade antecipadamente. (?)
- **R · Banquete Eterno** — Passivo: Ataques causam 2/4/6 + 0.03 attack_damage (bonus) de Dano Verdadeiro adicional, acumulando infinitamente. Eliminações de Campeões e monstros épicos deixam um Coral do Vazio. Ativo: Bel'Veth consome um Coral do Vazio, recebendo 1/1/1 acúmulo de Lavanda e ativando a Forma Verdadeira da Campeã. Coral do Vazio de monstros épicos do Vazio transforma tropas próximas que morrem em Rêmoras do Vazio. Durante a conjuração, Bel'Veth aplica Lentidão nos inimigos próximos antes de explodir, causando 150/200/250 + 1.5 ability_power + 20/20/20% da Vida perdida como Dano Verdadeiro. Na Forma Verdadeira, Bel'Veth recebe 100/250/400 + 1.5 ability_power + 1.5 attack_damage (bonus) de Vida máxima, 25/75/125 de Alcance de Ataque e 6/13/20% de Velocidade de Ataque total, e Impulso do Vazio pode atravessar paredes. A Forma Verdadeira dura 45/45/45s, aumentando para 90/90/90s com 40/40/40 acúmulos de Lavanda. Com 80/80/80 acúmulos de Lavanda, a Forma Verdadeira dura até a morte. (?)

## Braum, o Coração de Freljord

`Braum` · id 201 · corpo a corpo · tank, support

- **P · Golpes Concussivos** — Os ataques básicos de Braum aplicam Golpes Concussivos. Uma vez que o primeiro acúmulo é aplicado, os ataques básicos de aliados também acumulam Golpes Concussivos. Ao atingir 4 acúmulos, o alvo é atordoado e sofre Dano Mágico. Pelos próximos segundos, ele não pode receber acúmulos, mas recebe Dano Mágico adicional dos ataques de Braum.
- **Q · Mordida do Inverno** — Braum dispara gelo de seu escudo, causando ao primeiro inimigo atingido 75/120/165/210/255 + 0.025 max_health de Dano Mágico e 70/70/70/70/70% de Lentidão que decai ao longo de 2/2/2/2/2s. Aplica um acúmulo de Golpes Concussivos. (?)
- **W · Eu te Protejo** — Braum salta até um Campeão aliado ou tropa aliada. Ao aterrissar, ele concede ao alvo 20/25/30/35/40 + 0.12 armor (bonus) de Armadura e 20/25/30/35/40 + 0.12 magic_resist (bonus) de Resistência Mágica por 3/3/3/3/3s. Braum recebe 20/25/30/35/40 + 0.36 armor (bonus) de Armadura e 20/25/30/35/40 + 0.36 magic_resist (bonus) de Resistência Mágica pela mesma duração.(?)
- **E · Inquebrável** — Braum ergue seu escudo por 3/3.25/3.5/3.75/4s, bloqueando projéteis inimigos vindos da direção escolhida e fazendo com que atinjam Braum antes de serem destruídos. O primeiro projétil que Braum bloqueia não causa dano, e projéteis subsequentes causam 35/40/45/50/55% de dano reduzido. Braum receberá 10/10/10/10/10% de Velocidade de Movimento enquanto o escudo permanecer erguido.(?)
- **R · Fissura Glacial** — Braum bate com seu escudo no chão, enviando uma fissura que Arremessa ao ar os inimigos no trajeto e próximos a Braum, além de causar 150/250/350 + 0.6 ability_power de Dano Mágico. O primeiro alvo atingido é Arremessado ao ar por 0.6/0.6/0.6s - 1/1.5/2s, dependendo da distância entre ele e Braum. Todos os outros alvos atingidos são Arremessado ao ar por 0.6/0.6/0.6s. A fissura cria uma zona que dura 4/4/4s e causa 40/50/60% de Lentidão. (?)

## Jhin, o Virtuoso

`Jhin` · id 202 · a distância · marksman, mage

- **P · Sussurro** — Sussurro, o canhão de mão de Jhin, é um instrumento preciso que foi projetado para causar muito dano. Ele dispara com frequência fixa e carrega apenas quatro projéteis. Jhin encanta o projétil final com magia sombria para causar Acerto Crítico e dano de execução adicional. Sempre que causa um Acerto Crítico, Sussurro inspira Jhin com um impulso de Velocidade de Movimento.
- **Q · Granada Dançante** — Jhin lança um cartucho que causa 44/69/94/119/144 + 0.6 ability_power + 0.44 attack_damage de Dano Físico antes de saltar até um inimigo próximo ainda não atingido. O cartucho pode atingir no máximo 4/4/4/4/4 vezes. Inimigos que morrem logo depois de serem atingidos aumentam o dano de acertos subsequentes em 35/35/35/35/35%.(?)
- **W · Florescer Mortal** — Jhin dispara um tiro de longo alcance, causando 70/105/140/175/210 + 0.5 attack_damage de Dano Físico ao primeiro Campeão atingido e aos outros inimigos pelo caminho. Caso essa Habilidade atinja um Campeão que tenha sofrido dano de um Campeão aliado nos últimos 4/4/4/4/4s, ela o Enraizará por 1.25/1.5/1.75/2/2.25s e concederá a Jhin a Velocidade de Movimento de Sussurro.(?)
- **E · Audiência Cativa** — Passivo: Campeões abatidos por Jhin criarão e detonarão uma Armadilha de Lótus onde estiverem. Ativo: Jhin posiciona uma Armadilha de Lótus invisível que dura 3/3/3/3/3min, criando uma área que causa 35/35/35/35/35% de Lentidão ao ser pisada por um inimigo. Depois de 2/2/2/2/2s, a armadilha detona, causando 20/80/140/200/260 + 1 ability_power + 1.2 attack_damage de Dano Mágico. A Habilidade tem 2 cargas (recarrega a cada (?)s).(?)
- **R · Aclamação** — Jhin se prepara e canaliza, habilitando 4 superdisparos. Cada um deles causa de 64/128/192 + 0.25 attack_damage a 256/512/768 + 1 attack_damage de Dano Físico ao primeiro Campeão atingido com base na porcentagem de Vida perdida do inimigo e reduz a Velocidade de Movimento dele em 80/80/80% por 0.5/0.5/0.5s. O quarto disparo é um Acerto Crítico que causa 200/200/200% de dano.(?)

## Kindred, os Caçadores Eternos

`Kindred` · id 203 · a distância · marksman

- **P · Marca Familiar** — Os Kindred marcam alvos para Caçar. Concluir uma Caçada fortalece permanentemente as habilidades básicas deles. A cada 4 caçadas concluídas, o Alcance de Ataque também aumenta.
- **Q · Dança de Flechas** — Os Kindred saltam, disparando uma flecha em até 3 inimigos, causando 40/65/90/115/140 + 0.75 attack_damage (bonus) de Dano Físico e recebendo (?) de Velocidade de Ataque por 4/4/4/4/4s. Enquanto estiver dentro de Frenesi do Lobo, o Tempo de Recarga desta Habilidade é reduzido para 4/3.5/3/2.5/2s.(?)
- **W · Frenesi do Lobo** — Passivo: Os Kindred recebem acúmulos enquanto se movem e Atacam. Com 100 acúmulos, o próximo Ataque restaura até 47/47/47/47/47 de Vida com base na Vida perdida deles. Ativo: Os Kindred marcam um território, ordenando ao Lobo que morda o último inimigo que a Ovelha Atacou. As mordidas do Lobo causam 25/30/35/40/45 + 0.2 ability_power + 0.2 attack_damage (bonus) mais (?) da Vida atual do alvo como Dano Mágico.(?)
- **E · Pesar Crescente** — Os Kindred enfraquecem um inimigo, causando 30/30/30/30/30 + 0.05 ability_power% de Lentidão por 1/1/1/1/1s. O terceiro Ataque dos Kindred contra o alvo dentro de 4/4/4/4/4s envia o Lobo para investir contra o inimigo, causando 80/110/140/170/200 + 1 attack_damage (bonus) + (?) da Vida perdida como Dano Físico.(?)
- **R · Refúgio da Ovelha** — Os Kindred abençoam o terreno por 4/4/4s, impedindo que qualquer unidade (seja ela aliada, inimiga ou neutra) seja abatida enquanto permanecer ali. Ao atingir 10% de Vida, as unidades não podem sofrer dano ou ser curadas enquanto permanecerem na área. Quando a bênção termina, todas as unidades dentro dela se curam em 225/300/375 de Vida.(?)

## Zeri, A Faísca de Zaun

`Zeri` · id 221 · a distância · marksman

- **P · Bateria Viva** — Os Ataques de Zeri causam Dano Mágico e são considerados Habilidades. Mover-se e conjurar Rajada Reluzente armazena energia na Eletrobolsa. Quando estiver totalmente carregada, o próximo Ataque da Campeã causará dano adicional.
- **Q · Rajada Reluzente** — Zeri dispara uma rajada de 7/7/7/7/7 projéteis que causa 22/26/30/34/38 + 1.02 attack_damage de Dano Físico ao primeiro inimigo atingido. Esta Habilidade é considerada um Ataque. (?)
- **W · Laser de Ultrachoque** — Zeri dispara um pulso elétrico que causa 30/70/110/150/190 + 0.5 ability_power + 1.2 attack_damage de Dano Físico e 30/35/40/45/50% de Lentidão ao primeiro inimigo atingido por 2/2/2/2/2s. Se o pulso atingir um terreno, ele se expande em um laser que aplica os efeitos em uma área e causa Acerto Crítico igual a (?) de Dano Físico a Campeões e monstros.(?)
- **E · Faísca Acelerada** — Zeri avança por uma pequena distância e salta sobre qualquer terreno que tocar, aumentando bastante o alcance do avanço. Pelos próximos 5/5/5/5/5s, os disparos de Rajada Reluzente atravessarão os alvos, causando 80/85/90/95/100% de dano a inimigos depois do primeiro e 22/24/26/28/30 + 0.2 ability_power de Dano Mágico ao contato ao primeiro alvo atingido. Atingir um inimigo com um Ataque reduz o Tempo de Recarga desta Habilidade em 0.5/0.5/0.5/0.5/0.5s. Acertos Críticos reduzem o Tempo de Recarga em 1.5/1.5/1.5/1.5/1.5s.(?)
- **R · Impacto Eletrizante** — Zeri descarrega uma grande explosão elétrica, causando 150/250/350 + 1.1 ability_power + 0.6 attack_damage (bonus) de Dano Mágico a inimigos próximos. Se atingir um Campeão inimigo, Zeri recebe 30/30/30% de Velocidade de Ataque e 15/15/15% de Velocidade de Movimento por 5/5/5s. Atingir um Campeão inimigo com um Ataque ou Habilidade aumenta a duração da Habilidade e concede um acúmulo de Sobrecarga por 2.5/2.5/2.5s. Acertos Críticos concedem 2 acúmulos adicionais. Cada acúmulo concede 1.5/1.5/1.5% de Velocidade de Movimento. Durante esse período, Rajada Reluzente se transformará em um disparo triplo mais rápido que causa 0/0/0 + 0.4 attack_damage de Dano Físico a inimigos próximos. (?)

## Jinx, o Gatilho Desenfreado

`Jinx` · id 222 · a distância · marksman

- **P · Anime-se!** — Jinx recebe um aumento drástico de Velocidade de Movimento e de Velocidade de Ataque sempre que ajuda a destruir uma estrutura ou a abater um Campeão inimigo ou monstro épico da selva.
- **Q · Trocando!** — Jinx alterna entre Fishbones (o Lança-Mísseis) e Pow-Pow (a Metralhadora). Quando Jinx usa o lança-mísseis, os Ataques dela causam 0/0/0/0/0 + 1.1 attack_damage de Dano Físico ao alvo e aos inimigos próximos, recebem 100/125/150/175/200 de alcance, custam Mana e escalam 10/10/10/10/10% a menos com Velocidade de Ataque adicional. Quando Jinx usa a metralhadora, os Ataques dela concedem Velocidade de Ataque por 2.5/2.5/2.5/2.5/2.5s, acumulando até 3/3/3/3/3 vezes (máximo de +30/55/80/105/130% %i:scaleAS%).(?)
- **W · Zap!** — Jinx libera um disparo elétrico que causa 10/60/110/160/210 + 1.4 attack_damage de Dano Físico ao primeiro inimigo atingido, causando 40/50/60/70/80% de Lentidão e revelando-o por 2/2/2/2/2s.(?)
- **E · Mordidinha Flamejante!** — Jinx arremessa 3 Mordidinhas que duram 5/5/5/5/5s. Elas explodem ao entrar em contato com Campeões inimigos, Enraizando-os por 1.5/1.5/1.5/1.5/1.5s e causando 90/140/190/240/290 + 1 ability_power de Dano Mágico a inimigos próximos.(?)
- **R · Super Mega Míssil da Morte!** — Jinx dispara um míssil que explode no primeiro Campeão inimigo que atingir, causando de 20/35/50 + 0.12 attack_damage (bonus) a 200/350/500 + 1.2 attack_damage (bonus) + 25/30/35% da Vida perdida como Dano Físico, acumulando dano ao longo do primeiro segundo do trajeto. Inimigos próximos sofrem 80/80/80% do dano. O dano com base na Vida perdida não pode exceder 1200/1200/1200 contra monstros.(?)

## Tahm Kench, o Rei do Rio

`Tahm Kench` · id 223 · corpo a corpo · tank, support

- **P · Um Gosto Adquirido** — Tahm Kench usa o próprio corpo para impulsionar os Ataques, causando dano adicional com base na Vida total dele. Causar dano a Campeões inimigos aumenta os acúmulos de Um Gosto Adquirido. Com três acúmulos, ele pode usar Devorar em um Campeão inimigo.
- **Q · Língua-chicote** — Causa 75/120/165/210/255 + 1 ability_power de Dano Mágico e 50/50/50/50/50% de Lentidão ao primeiro inimigo atingido por 2/2/2/2/2s. Ao acertar um Campeão, cura Tahm Kench em 10/15/20/25/30 + 5/5.5/6/6.5/7% da própria Vida perdida e aplica um acúmulo de Um Gosto Adquirido, causando @Spell.TahmKenchPassive:TotalDamage@ de Dano Mágico adicional. Se o Campeão já tiver 3 acúmulos de Um Gosto Adquirido, ele será Atordoado por 1.5/1.5/1.5/1.5/1.5s e os acúmulos serão consumidos. Ative Devorar enquanto a língua estiver no meio do trajeto para devorar Campeões inimigos que já tiverem 3 acúmulos de Um Gosto Adquirido ao atingi-los.(?)
- **W · Mergulho Abissal** — Mergulha e reaparece no local-alvo, causando 100/135/170/205/240 + 1.5 ability_power de Dano Mágico e Arremessando ao ar todos os inimigos em uma área por 1/1/1/1/1s. Atingir ao menos um Campeão inimigo restaura 40/42.5/45/47.5/50% do Tempo de Recarga e do Custo de Mana. Aliados Devorados podem ser carregados na viagem (mas podem sair antecipadamente).(?)
- **E · Pele Grossa** — Passivo: 15/23/31/39/47% do dano sofrido por Tahm Kench é armazenado por Pele Grossa, aumentado para 42/44/46/48/50% se houver pelo menos 2/2/2/2/2 Campeões inimigos próximos. Se ele não tiver sofrido dano por 4/4/4/4/4s, Pele Grossa será consumida rapidamente para curá-lo em 1/1/1/1/1 do valor dela. Ativo: converte todo o dano armazenado por Pele Grossa em um Escudo que dura 2.5/2.5/2.5/2.5/2.5s.(?)
- **R · Devorar** — Tahm Kench devora um Campeão por alguns segundos. Ele pode Reconjurar para cuspi-lo. Campeões inimigos: requerem 3 acúmulos de Um Gosto Adquirido. São devorados por até 3/3/3s e sofrem 100/250/400 +0.15/0.15/0.15 + 0.0007 ability_power da Vida máxima como Dano Mágico. Tahm Kench sofre 40/40/40% de Lentidão e permanece Preso ao chão durante o efeito. Campeões aliados: são devorados por até 3/3/3s e recebem 650/800/950 + 1 ability_power de Escudo, que decai gradualmente após serem cuspidos. Os aliados podem se libertar antecipadamente. Tahm Kench fica Preso ao chão durante este efeito, mas pode conjurar Mergulho Abissal e receber 60/60/60% de Velocidade de Movimento por 3/3/3s.(?)

## Briar, a Fome Contida

`Briar` · id 233 · corpo a corpo · fighter, assassin

- **P · Maldição Carmesim** — Os Ataques e as Habilidades de Briar acumulam sangramento que cura Briar numa quantidade do dano causado. Constantemente faminta, ela ganha cura aumentada com base na Vida perdida, mas não tem Regeneração de Vida inata.
- **Q · Vertigem** — Briar salta em direção a um alvo, Atordoando-o por 0.85/0.85/0.85/0.85/0.85s, causando 60/85/110/135/160 + 0.6 ability_power + 0.8 attack_damage (bonus) de Dano Físico e reduzindo 10/12.5/15/17.5/20% da Armadura e da Resistência Mágica dele por 5/5/5/5/5s. Briar para de priorizar Campeões se utilizar esta Habilidade em uma tropa ou um monstro durante Frenesi Sanguinário.(?)
- **W · Frenesi Sanguinário/Ataque Faminto** — Briar salta e ativa Frenesi Sanguinário, entrando em provocação contra o inimigo mais próximo por 5/5/5/5/5s, priorizando Campeões. Durante o Frenesi Sanguinário, ela ganha 55/65/75/85/95% de Velocidade de Ataque e 24/33/42/51/60% de Velocidade de Movimento, e seus Ataques causam 0/0/0/0/0 + 0.6 attack_damage de Dano Físico a inimigos ao redor do alvo principal. Briar pode Reconjurar esta Habilidade para fortalecer seu próximo Ataque. Ela causa 5/20/35/50/65 + 1.05 attack_damage + 9/9/9/9/9 + 0.025 attack_damage (bonus)% de Dano Físico com base na Vida perdida e se cura em 0/0/0/0/0 + 0.05 max_health + 24.000002/28/32/36/40% do dano causado. (?)
- **E · Grito Arrepiante** — Início do carregamento: Briar remove Frenesi Sanguinário e reúne energia, recebendo 35/35/35/35/35% de redução de dano e restaurando 0/0/0/0/0 + 0.1 max_health de Vida ao longo de 1s. Libertação: Briar dá um grito que causa até 80/115/150/185/220 + 1 ability_power + 1 attack_damage (bonus) de Dano Mágico com base no tempo de carregamento, além de 80/80/80/80/80% de Lentidão por 0.5/0.5/0.5/0.5/0.5s. Quando o grito está totalmente carregado, ele Empurra inimigos, causando 140/215/290/365/440 + 2.4 ability_power + 2.4 attack_damage (bonus) de Dano Mágico àqueles que atingirem uma parede e Atordoando-os por 1.5/1.5/1.5/1.5/1.5s.(?)
- **R · Morte Certa** — Briar lança o hemólito da berlinda dela e voa até o primeiro Campeão atingido, marcando-o como sua presa. Ao chegar, ela causa 150/250/350 + 1.3 ability_power de Dano Mágico a tudo ao seu redor e faz inimigos que não sejam sua presa Fugirem por 1.5/1.5/1.5s. Depois disso, ela entra num Frenesi Sanguinário fortalecido e persegue a presa até a morte. Nesse período, ela recebe 0/0/0 + 0.2 attack_damage de Armadura e Resistência Mágica, além de 10/15/20% de Roubo de Vida e 10/20/30% de Velocidade de Movimento adicional.(?)

## Viego, O Rei Destruído

`Viego` · id 234 · corpo a corpo · fighter, assassin

- **P · Dominação Monárquica** — Inimigos derrotados por Viego se tornam Espectros. Ao atacá-los, Viego toma controle do corpo do inimigo morto temporariamente, curando-se em um percentual da Vida máxima do alvo e ganhando acesso a seus itens e habilidades básicas. Ele substitui a ultimate da vítima por uma conjuração gratuita de sua própria ultimate.
- **Q · Espada do Rei Destruído** — Passivo: os Ataques de Viego causam 2/3/4/5/6 da Vida atual como Dano Físico adicional. O primeiro Ataque contra um inimigo que tiver sofrido dano de uma Habilidade recentemente atinge o alvo uma segunda vez, causando 0/0/0/0/0 + 0.15 ability_power + 0.2 attack_damage de Dano Físico e restaurando Vida em um valor equivalente a 150/150/150/150/150% do dano causado. Esses efeitos são mantidos durante Possessão. Ativo: Viego golpeia à frente, causando 25/40/55/70/85 + 0.7 attack_damage de Dano Físico.(?)
- **W · Posse Espectral** — Início do carregamento: Viego começa a acumular Névoa, sofrendo 10/10/10/10/10% de Lentidão. Liberar: Viego avança e arremessa a Névoa acumulada. Causa 80/135/190/245/300 + 1 ability_power de Dano Mágico ao primeiro inimigo atingido e o Atordoa por 0.25/0.25/0.25/0.25/0.25s -1.25/1.25/1.25/1.25/1.25s com base no tempo de carregamento. (?)
- **E · Domínio Atormentado** — Viego envia um espectro para assombrar o primeiro terreno atingido, envolvendo-o em Névoa por 8/8/8/8/8s. Viego recebe Camuflagem, 0.25/0.275/0.3/0.325/0.35 + 0.0004 ability_power de Velocidade de Movimento e 30/35/40/45/50% de Velocidade de Ataque enquanto estiver dentro da Névoa.(?)
- **R · Destruidor de Corações** — Viego descarta quaisquer almas que esteja Possuindo no momento e se teleporta. Ao chegar, ele Ataca o Campeão inimigo com a porcentagem de Vida mais baixa, causando 99/99/99% de Lentidão brevemente e 0/0/0 + 1.2 attack_damage + 12/16/20 + 0.05 attack_damage (bonus)% da Vida perdida como Dano Físico. Outros inimigos próximos são Empurrados e sofrem 0/0/0 + 1.2 attack_damage de Dano Físico.(?)

## Senna, A Redentora

`Senna` · id 235 · a distância · support, marksman

- **P · Absolvição** — Quando unidades são abatidas perto de Senna, suas almas são periodicamente aprisionadas na Névoa Negra. Senna pode atacar essas almas para libertá-las, absorvendo a Névoa que as aprisiona na morte. A Névoa é o combustível do poder de seu Canhão Relicário, alimentando-o com Dano de Ataque, Alcance de Ataque e Chance de Acerto Crítico. Ataques do Canhão Relicário de Senna demoram mais para disparar, causam dano adicional e concedem a ela brevemente uma parte da Velocidade de Movimento do alvo.
- **Q · Escuridão Perfurante** — Senna dispara um raio de escuridão perfurante através de um aliado ou inimigo, causando 30/55/80/105/130 + 0.6 attack_damage (bonus) de Dano Físico e 0.15/0.15/0.15/0.15/0.15 + 0.0007 ability_power + 0.0015 attack_damage (bonus) de Lentidão por 1/1.25/1.5/1.75/2s. Restaura 40/60/80/100/120 + 0.35 ability_power + 0.4 attack_damage (bonus) de Vida de Campeões aliados. Ataques reduzem o Tempo de Recarga da Habilidade em 1/1/1/1/1s.(?)
- **W · Abraço Final** — Senna dispara Névoa Negra, causando 70/110/150/190/230 + 0.9 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido. Após um intervalo de 1/1/1/1/1s, o alvo e outros inimigos próximos são Enraizados por 1.25/1.5/1.75/2/2.25s.(?)
- **E · Maldição da Névoa Negra** — Senna se transforma em uma nuvem de Névoa Negra por 6/6.5/7/7.5/8s, tornando-se um Espectro. Campeões aliados que entram na Névoa são Camuflados e se tornam espectros quando saem. Espectros recebem 0.2/0.2/0.2/0.2/0.2 + 0.0005 ability_power de Velocidade de Movimento, não são selecionáveis e escondem suas identidades enquanto não houver Campeões inimigos próximos.(?)
- **R · Sombra da Alvorada** — Senna dispara um raio de luz que causa 250/400/550 + 0.7 ability_power + 1.15 attack_damage (bonus) de Dano Físico a todos os Campeões inimigos atingidos. Campeões aliados atingidos em uma área mais ampla recebem (?) de Escudo com duração de 3/3/3s.(?)

## Lucian, o Purificador

`Lucian` · id 236 · a distância · marksman, assassin

- **P · Disparo Iluminado** — Sempre que Lucian usa uma Habilidade, seu próximo Ataque se transforma em um disparo duplo. Quando Lucian é curado ou protegido por um aliado, ou quando um Campeão inimigo próximo é imobilizado, os próximos dois Ataques básicos dele causam Dano Mágico adicional.
- **Q · Luz Perfurante** — Lucian dispara uma rajada de Luz Perfurante, causando 80/115/150/185/220 + 1 attack_damage (bonus) de Dano Físico. (?)
- **W · Chama Ardente** — Lucian dispara um tiro que explode ao chegar ao fim do alcance ou atingir um inimigo, causando 75/110/145/180/215 + 0.9 ability_power de Dano Mágico, revelando inimigos por um breve momento e marcando-os por 6s. Quando Lucian ou um aliado causa dano a um inimigo marcado, Lucian recebe 60/65/70/75/80 de Velocidade de Movimento por 1s. Aliados que acionam esse efeito também concedem Vigilância a Lucian.(?)
- **E · Perseguição Implacável** — Lucian avança. O Tempo de Recarga é reduzido em 1/1/1/1/1s sempre que ele atinge um inimigo com Disparo Iluminado (2/2/2/2/2s para Campeões). (?)
- **R · O Expurgo** — Lucian atira rapidamente, lançando (?) disparos em uma direção por 3/3/3s ou até Reconjurar. Cada disparo causa 15/30/45 + 0.15 ability_power + 0.25 attack_damage de Dano Físico ao primeiro inimigo atingido. Enquanto estiver atirando, Lucian pode usar Perseguição Implacável. Dano total: (?) de Dano Físico (?)

## Zed, o Mestre das Sombras

`Zed` · id 238 · corpo a corpo · assassin

- **P · Desprezo pelos Fracos** — Os ataques básicos de Zed contra alvos com pouca Vida causam Dano Mágico adicional. Esse efeito pode ocorrer contra o mesmo Campeão inimigo apenas uma vez a cada poucos segundos.
- **Q · Shuriken Laminado** — Zed e suas Sombras arremessam shurikens. Cada uma causa 80/120/160/200/240 + 1 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido e 48/72/96/120.00001/144 + 0.6 attack_damage (bonus) de Dano Físico a cada inimigo adicional.(?)
- **W · Sombra Viva** — Passivo: Zed recebe 30/35/40/45/50 de Energia sempre que ele e suas Sombras atingem um inimigo com a mesma habilidade. Ativo: a Sombra de Zed avança, permanecendo no lugar por 5/5/5/5/5s. Reconjurar esta Habilidade faz Zed trocar de posição com a Sombra.(?)
- **E · Corte Sombrio** — Zed e suas Sombras desferem um corte, causando 70/92.5/115/137.5/160 + 0.7 attack_damage (bonus) de Dano Físico a inimigos próximos. Cada Campeão inimigo que for atingido pelo corte de Zed reduz o Tempo de Recarga de Sombra Viva em 3/3/3/3/3s. Inimigos atingidos pelo corte de uma Sombra sofrem @MoveSpeedMod*-100@% de Lentidão por 1.5/1.5/1.5/1.5/1.5s. Inimigos atingidos por mais de um ataque não sofrem dano adicional, e sim @MoveSpeedModBonus*-100@% de Lentidão.(?)
- **R · Marca Fatal** — Zed se torna Inalvejável, avança até um Campeão inimigo e o marca. Depois de 3/3/3s, a marca é ativada, causando 0/0/0 + 1 attack_damage de Dano Físico e repetindo 25/40/55% de todo o dano causado por Zed ao alvo enquanto a marca estava ativa. O avanço deixa uma Sombra para trás por 7.5/7.5/7.5s. Zed pode Reconjurar a Habilidade para trocar de posição com a Sombra.(?)

## Kled, o Cavaleiro Intratável

`Kled` · id 240 · corpo a corpo · fighter

- **P · Skaarl, a Lagarto Covarde** — Kled fica em sua possante montaria, Skaarl, que toma dano por ele. Quando a Vida dela acaba, Kled desmonta. Enquanto desmontado, as habilidades de Kled mudam e ele causa menos dano a Campeões. Kled pode restaurar a coragem de Skaarl ao enfrentar inimigos. Quando a coragem chega ao máximo, Kled volta a montar em Skaarl com uma parte de sua Vida.
- **Q · Armadilha na Corda** — Montado: Kled arremessa uma armadilha de urso que causa 30/55/80/105/130 + 0.6 attack_damage (bonus) de Dano Físico e se prende ao primeiro Campeão inimigo ou monstro grande da selva atingido. Se ficar próximo a um inimigo preso por 1.75/1.75/1.75/1.75/1.75s, Kled arrancará a armadilha, causando 60/110/160/210/260 + 1.2 attack_damage (bonus) de Dano Físico, Puxando e causando @SlowAmount*-100@% de Lentidão ao alvo por 2.5/2.5/2.5/2.5/2.5s.(?)
- **W · Tendências Violentas** — Passivo: o próximo Ataque de Kled concede 150/150/150/150/150% de Velocidade de Ataque por quatro Ataques ou 4/4/4/4/4s. O quarto golpe causa 20/30/40/50/60 + 4.5/5/5.5/6/6.5 + 0.02 attack_damage (bonus) + 0.004 max_health (bonus) da Vida máxima como Dano Físico.(?)
- **E · Justar** — Kled avança, causando 35/60/85/110/135 + 0.55 attack_damage (bonus) de Dano Físico aos inimigos que ele atravessar no caminho e puxando tropas e monstros pequenos. Se a Habilidade atingir um Campeão inimigo ou monstros grandes da selva, Kled recebe 50/50/50/50/50% de Velocidade de Movimento por 1/1/1/1/1s e pode Reconjurar em até 3/3/3/3/3s para avançar de volta através do mesmo alvo.(?)
- **R · Avançaaaaaaar!!!** — Kled avança em direção a uma área, deixando um rastro que aumenta a Velocidade de Movimento dos aliados. Enquanto avança, e 2s depois de avançar, Kled recebe até 200/300/400 + 3 attack_damage (bonus) de Escudo. Skaarl atropela o primeiro Campeão inimigo no caminho, causando 4/6/8 + 0.03 attack_damage (bonus) - 4/6/8 + 0.03 attack_damage (bonus) da Vida máxima como Dano Mágico (com base na distância percorrida) e Empurrando brevemente.(?)

## Ekko, o Rapaz que Estilhaçou o Tempo

`Ekko` · id 245 · corpo a corpo · assassin, mage

- **P · Ressonância Revo-Z** — Cada terceiro ataque ou habilidade de dano no mesmo alvo causa Dano Mágico adicional e concede a Ekko um impulso de velocidade se o alvo for um Campeão.
- **Q · Giratempo** — Ekko arremessa um dispositivo, causando 80/95/110/125/140 + 0.3 ability_power de Dano Mágico. Ao atingir um Campeão ou chegar ao fim do alcance, ele se expande em um campo que causa @SlowPercent*-100@% de Lentidão aos inimigos atingidos. Depois da expansão, Ekko puxa o dispositivo de volta, causando 40/65/90/115/140 + 0.6 ability_power de Dano Mágico.(?)
- **W · Convergência Paralela** — Passivo: os Ataques de Ekko contra inimigos com menos de 30/30/30/30/30% de Vida causam 3/3/3/3/3 + 0.03 ability_power da Vida perdida como Dano Mágico. Ativo: Ekko dispara uma cronoesfera que dura 1.5/1.5/1.5/1.5/1.5s depois de um intervalo, causando 40/40/40/40/40% de Lentidão aos inimigos dentro da área. Se entrar nela, Ekko a detona, Atordoando inimigos por 2.25/2.25/2.25/2.25/2.25s e recebendo 100/120/140/160/180 + 1.5 ability_power de Escudo.(?)
- **E · Mergulho Fásico** — Ekko avança e fortalece seu próximo Ataque, recebendo alcance adicional, teleportando-se até o alvo e causando a ele 50/75/100/125/150 + 0.4 ability_power de Dano Mágico adicional.(?)
- **R · Cronoquebra** — Ekko volta no tempo, entrando em Estase enquanto se teleporta para onde estava 4s atrás e causando 200/350/500 + 1.75 ability_power de Dano Mágico a inimigos próximos. Além disso, Ekko restaura 100/150/200 + 0.6 ability_power de Vida, aumentado em 3/3/3% para cada 1% da Vida perdida nos últimos 4s.(?)

## Qiyana, a Imperatriz dos Elementos

`Qiyana` · id 246 · corpo a corpo · assassin

- **P · Privilégio da Realeza** — O primeiro ataque básico ou habilidade de Qiyana contra cada inimigo causa dano adicional.
- **Q · Cólera Elemental / Lâmina de Ixtal** — Se Qiyana não tiver um Encantamento, ela golpeia com a arma, causando 70/100/130/160/190 + 0.85 attack_damage (bonus) de Dano Físico a inimigos em uma pequena área. Caso faça isso, a Habilidade ganha maior alcance e efeitos adicionais com base no tipo de Encantamento:Encantamento de Gelo: Enraíza os inimigos atingidos por um curto período, depois causa @SlowPotency*-100@% de Lentidão por 1/1/1/1/1s.Encantamento de Pedra: causa 42/60.000004/78/96/114.00001 + 0.51000005 attack_damage (bonus) de Dano Físico adicional a unidades com menos de 50/50/50/50/50% de Vida.Encantamento da Selva: forma uma trilha que deixa Qiyana Invisível e concede 20/20/20/20/20% de Velocidade de Movimento. (?)
- **W · Terraforme** — Passivo: enquanto a arma de Qiyana estiver Encantada, ela recebe 15/20/25/30/35% de Velocidade de Ataque e seus Ataques causam 8/16/24/32/40 + 0.45 ability_power + 0.2 attack_damage (bonus) de Dano Mágico adicional. Ela também recebe 3/5/7/9/11% de Velocidade de Movimento quando fora de combate e próxima ao tipo do terreno correspondente. Ativo: Qiyana avança em direção ao arbusto, terreno ou rio próximo e Encanta a arma com o tipo específico daquele terreno. Fazer isso redefine o Tempo de Recarga de Cólera Elemental / Lâmina de Ixtal.(?)
- **E · Audácia** — Qiyana avança e atravessa um inimigo, causando 50/90/130/170/210 + 0.5 attack_damage (bonus) de Dano Físico.(?)
- **R · Suprema Demonstração de Talento** — Qiyana lança uma onda de choque que Empurra inimigos e detona quando atingirem um terreno. A explosão se estende até o limite do terreno, Atordoando por 0,5s - 1/1/1s e causando 100/200/300 + 1.25 attack_damage (bonus) mais 0.1/0.1/0.1 da Vida máxima como Dano Físico. A duração do Atordoamento é reduzida de acordo com a distância percorrida pela onda de choque. Qualquer rio ou arbusto que a onda atravessar também explode depois de um intervalo, causando o mesmo dano e Atordoamento.(?)

## Vi, a Defensora de Piltover

`Vi` · id 254 · corpo a corpo · fighter, assassin

- **P · Blindagem** — Vi carrega um escudo ao longo do tempo que pode ser ativado ao atingir um inimigo com uma habilidade.
- **Q · Quebra-Cofres** — Início do carregamento: Vi começa a carregar um soco poderoso e recebe 15/15/15/15/15% de Lentidão. Liberar: Vi avança para a frente, causando 40/60/80/100/120 + 0.6 attack_damage (bonus) - 100/150/200/250/300 + 1.5 attack_damage (bonus) de Dano Físico com base no tempo de carregamento e aplicando Pancada Certeira a todos os inimigos atingidos. O trajeto de Vi é interrompido ao colidir com um Campeão inimigo e o Arremessa para trás.(?)
- **W · Pancada Certeira** — Passivo: cada terceiro Ataque no mesmo alvo causa 4/5/6/7/8 + 0.035 attack_damage (bonus) da Vida máxima como Dano Físico, remove 20/20/20/20/20% de Armadura e concede a Vi 30/35/40/45/50% de Velocidade de Ataque por 4/4/4/4/4s. Isso também reduz o Tempo de Recarga restante de Blindagem em @spell.ViPassive:CDReductionOn3Hit(?)SpellModifierDescriptionAppend@
- **E · Força Implacável** — O próximo Ataque de Vi causa 10/30/50/70/90 + 1 ability_power + 1.1 attack_damage de Dano Físico ao alvo e a inimigos atrás dele. A Habilidade tem 2 cargas. (recarrega a cada (?)s).(?)
- **R · Saque e Enterrada** — Vi mira em um Campeão inimigo, revelando-o e avançando em sua direção sem poder ser interrompida. Ao alcançá-lo, Vi o Arremessa ao ar por 1.3/1.3/1.3s e causa 150/250/350 + 0.9 attack_damage (bonus) de Dano Físico. Qualquer outro inimigo com que Vi colide no trajeto também sofre dano, é arremessado para o lado e fica Atordoado por 0.75/0.75/0.75s.(?)

## Aatrox, a Espada Darkin

`Aatrox` · id 266 · corpo a corpo · fighter

- **P · Postura do Arauto da Morte** — Periodicamente, o próximo ataque básico de Aatrox causa Dano Mágico adicional e o cura com base na Vida máxima do alvo.
- **Q · A Espada Darkin** — Aatrox bate com sua espada no chão, causando 10/25/40/55/70 + 0.6 attack_damage de Dano Físico. Inimigos atingidos na área exterior são Arremessados ao ar brevemente e sofrem 17.5/43.75/70/96.25/122.5 + 1.0500001 attack_damage de Dano Físico. A Habilidade pode ser Reconjurada duas vezes, alterando o formato do golpe e causando 25% de dano a mais do que o golpe anterior.(?)
- **W · Correntes Infernais** — Aatrox arremessa uma corrente, causando ao primeiro inimigo atingido @WSlowPercentage*-100@% de Lentidão por 1.5/1.5/1.5/1.5/1.5s e 30/40/50/60/70 + 0.4 attack_damage de Dano Físico. Campeões e monstros grandes da selva devem sair da área de impacto em até 1.5/1.5/1.5/1.5/1.5s ou serão Puxados de volta ao centro, sofrendo o mesmo dano novamente.(?)
- **E · Avanço Umbral** — Passivo: Aatrox se cura em 16/16/16/16/16 + 0.011 max_health (bonus) do dano causado a Campeões. Ativo: Aatrox avança. Ele pode usar essa Habilidade durante as animações das outras Habilidades.(?)
- **R · Aniquilador de Mundos** — Aatrox revela sua verdadeira forma demoníaca, causando Temor às tropas próximas por 3/3/3s e recebendo 60/80/100% de Velocidade de Movimento que decai ao longo de 10/10/10s. Ele também recebe 20/30/40% de Dano de Ataque e aumenta a própria cura em 50/75/100% ao longo da duração. Eliminar Campeões aumenta a duração desse efeito em 5/5/5s e redefine o aumento de Velocidade de Movimento.(?)

## Nami, a Conjuradora das Marés

`Nami` · id 267 · a distância · support, mage

- **P · Maré Oscilante** — Quando as Habilidades de Nami atingem Campeões aliados, eles recebem Velocidade de Movimento por um curto período.
- **Q · Prisão Aquática** — Nami arremessa uma bolha que Atordoa por 1.5/1.5/1.5/1.5/1.5s e causa 90/145/200/255/310 + 0.5 ability_power de Dano Mágico.(?)
- **W · Vazante e Fluxo** — Nami libera um jato d'água que ricocheteia entre Campeões aliados e inimigos, até um máximo de 3/3/3/3/3 alvos. Cada Campeão pode ser atingido apenas uma vez.Ao atingir aliados, restaura 55/80/105/130/155 + 0.4 ability_power de Vida e ricocheteia para um Campeão inimigo próximo. Ao atingir inimigos, causa 60/95/130/165/200 + 0.5 ability_power de Dano Mágico e ricocheteia para um Campeão aliado próximo. As quantidades de dano e cura são modificadas em -20/-20/-20/-20/-20 + 0.15 ability_power a cada ricochete. (?)
- **E · Bênção da Conjuradora** — Nami fortalece os próximos 3/3/3/3/3 Ataques e Habilidades de um Campeão aliado por 6/6/6/6/6s, fazendo com que causem 15/20/25/30/35 + 0.05 ability_power de Lentidão por 1/1/1/1/1s e 20/35/50/65/80 + 0.2 ability_power de Dano Mágico. (?)
- **R · Maré Violenta** — Nami invoca uma grande onda, Arremessando ao ar os inimigos por 0,5s, causando 70/70/70% de Lentidão e 150/250/350 + 0.6 ability_power de Dano Mágico. A duração da Lentidão aumenta com base na distância percorrida pela onda, até no máximo 4/4/4s. Aliados atingidos pela onda recebem o efeito dobrado de Maré Oscilante.(?)

## Azir, o Imperador das Areias

`Azir` · id 268 · a distância · mage, marksman

- **P · Legado de Shurima** — Azir pode invocar o Disco Solar das ruínas de uma torre aliada ou inimiga.
- **Q · Areias da Conquista** — Azir envia todos os Soldados de Areia até uma área, causando aos inimigos atravessados 75/95/115/135/155 + 0.35 ability_power de Dano Mágico e @SlowAmount*-100@% de Lentidão por 1s.(?)
- **W · Surja!** — Azir invoca um Soldado de Areia por 10/10/10/10/10s. Quando Azir Ataca um inimigo próximo a um Soldado de Areia, ele ordena que o soldado apunhale, causando (?) de Dano Mágico na direção do inimigo. Essa habilidade tem (?) cargas.(?)
- **E · Areias Oscilantes** — Azir concede a si mesmo 70/110/150/190/230 + 0.6 ability_power de Escudo por 1.5/1.5/1.5/1.5/1.5s e avança até um de seus Soldados de Areia, causando 70/110/150/190/230 + 0.6 ability_power de Dano Mágico aos inimigos atravessados. Se atingir um Campeão inimigo, Azir para e recebe uma carga de Soldado de Areia.(?)
- **R · Decreto do Imperador** — Azir invoca uma barreira de soldados que avança, Empurrando inimigos e causando a eles 200/400/600 + 0.75 ability_power de Dano Mágico. Os soldados permanecem no lugar, bloqueando o caminho dos inimigos por 5/5/5s.(?)

## Yuumi, a Gata Mágica

`Yuumi` · id 350 · a distância · support, mage

- **P · Amizade Felina** — Periodicamente, quando Yuumi atinge um Campeão com um Ataque ou Habilidade, ela restaura Vida para si e para o próximo aliado a quem ela se Conectar. Enquanto Conectada, Yuumi cria um laço especial com os aliados. O aliado com o laço mais forte aprimora as Habilidades de Yuumi enquanto ela estiver Conectada a ele.
- **Q · Projétil Errante** — Yuumi invoca um projétil errante que causa 60/95/130/165/200 + 0.2 ability_power de Dano Mágico ao primeiro inimigo atingido, além de 20/20/20/20/20% de Lentidão. Se essa habilidade for conjurada enquanto Yuumi estiver Conectada, ela poderá controlar o projétil com o cursor do mouse por um breve período antes que ele acelere em linha reta. O projétil acelerado causa 80/135/190/245/300 + 0.3 ability_power de Dano Mágico e 50/53/56/59/62% de Lentidão ao alvo por 2/2/2/2/2s. Bônus de Melhor Amigo: a Lentidão de Projétil Errante sempre será aprimorada, e atingir um Campeão inimigo também concede 10/12/14/16/18 + 0.05 ability_power de Dano Mágico ao contato %i:OnHit% por 5/5/5/5/5s. O dano adicional ao contato pode ser aumentado em 75/75/75/75/75% com base na Chance de Acerto Crítico do aliado.(?)
- **W · Você e Eu!** — Passivo: enquanto estiver junto ao Melhor Amigo, Yuumi recebe 4/5/6/7/8% de cura e Resistência do Escudo adicionais, e o aliado dela também restaura 3/4/5/6/7 + 0.03 ability_power de Vida ao contato %i:OnHit%. Ativo: Yuumi avança até um Campeão aliado e se Conecta a ele. Enquanto estiver Conectada, ela seguirá o movimento do companheiro e ficará Inalvejável, exceto por torres. Os efeitos Imobilizadores sobre Yuumi colocam a Habilidade em um Tempo de Recarga de 5/5/5/5/5s.(?)
- **E · Frenética** — Yuumi se protege com um Escudo, bloqueando 65/90/115/140/165 + 0.4 ability_power de dano e recebendo 25/27.5/30/32.5/35 + 0.08 ability_power% de Velocidade de Ataque por 3/3/3/3/3s. Se o Escudo persistir, o alvo também recebe 20/20/20/20/20% de Velocidade de Movimento. Se Yuumi estiver Conectada, esta Habilidade afetará o aliado em vez dela, restaurando 20/24/28/32/36 de Mana para ele, aumentado em 100/100/100/100/100% com base no Mana perdido do alvo.(?)
- **R · Capítulo Final** — Yuumi canaliza por 3.5/3.5/3.5s, disparando 5/5/5 ondas mágicas que afetam ambas as equipes. Se conjurar inicialmente enquanto ela estiver Conectada, Yuumi pode controlar a direção das ondas com o cursor do mouse. Inimigos atingidos sofrem 75/125/175 + 0.25 ability_power de Dano Mágico e @BaseSlow*-100@% de Lentidão por 1.25/1.25/1.25s, aumentando em @BonusSlowPerWave*-100@% por onda atingida. Campeões aliados são curados em 30/50/70 + 0.12 ability_power de Vida por onda. A cura excedente é convertida em um Escudo. Efeito de Melhor Amigo: para o Melhor Amigo, aumenta a cura para 39/65/91 + 0.15599999 ability_power de Vida. Conjurar Você e Eu! travará as ondas na direção atual. Yuumi pode se mover e conjurar Frenética durante a canalização. (?)

## Samira, a Rosa do Deserto

`Samira` · id 360 · a distância · marksman, assassin

- **P · Impulso Audacioso** — Samira gera um combo conforme Ataca ou conjura habilidades de forma alternada. Seus Ataques corpo a corpo causam Dano Mágico adicional. Os Ataques de Samira contra inimigos sob efeitos Imobilizadores a farão avançar ao Alcance de Ataque dela. Se o inimigo estiver afetado por Arremesso ao ar, ela mantém o Arremesso ao ar por um curto período.
- **Q · Talento Natural** — Samira dispara um tiro, causando 0/5/10/15/20 + 1.1 attack_damage de Dano Físico ao primeiro inimigo atingido. Se a Habilidade for conjurada em direção a um inimigo dentro de seu alcance corpo a corpo, ela golpeia com a espada em vez de atirar, causando 0/5/10/15/20 + 1.1 attack_damage de Dano Físico. (?)
- **W · Voragem Afiada** — Samira gira a espada ao seu redor por 0.75/0.75/0.75/0.75/0.75s, causando 20/35/50/65/80 + 0.5 attack_damage (bonus) de Dano Físico aos inimigos duas vezes e destruindo qualquer projétil inimigo que entrar na área. (?)
- **E · Ímpeto Indomável** — Samira avança e atravessa um inimigo (incluindo estruturas), cortando os inimigos atravessados, causando a eles 50/60/70/80/90 + 0.2 attack_damage (bonus) de Dano Mágico e recebendo 20/25/30/35/40% de Velocidade de Ataque por 5/5/5/5/5s. Se um Campeão inimigo for eliminado dentro de 3s depois que Samira causar dano a ele, o Tempo de Recarga desta Habilidade será redefinido.(?)
- **R · Gatilho Infernal** — Samira pode usar esta Habilidade se a pontuação de Estilo atual for S. Usá-la consome toda a pontuação de Estilo. Samira dispara uma saraivada de tiros, atingindo 10 vezes todos os inimigos ao seu redor ao longo de 2s. Cada tiro causa 20/40/60 + 0.3 attack_damage de Dano Físico e aplica Roubo de Vida com 100/100/100% de eficácia. Cada tiro também pode causar Acerto Crítico.(?)

## Thresh, o Guardião das Correntes

`Thresh` · id 412 · a distância · support, tank

- **P · Condenação** — Thresh pode colher a alma de inimigos que morrerem perto dele, recebendo Armadura e Poder de Habilidade permanentemente.
- **Q · Sentença** — Thresh arremessa sua foice, Atordoando a primeira unidade atingida e Puxando-a em sua direção por 1.5/1.5/1.5/1.5/1.5s. A foice causa 100/150/200/250/300 + 0.9 ability_power de Dano Mágico e concede Visão Mágica ao longo da duração. Thresh pode Reconjurar a Habilidade para se puxar para o inimigo. Se acertar a Habilidade, o Tempo de Recarga é reduzido em 2/2/2/2/2s.(?)
- **W · Passagem Sombria** — Thresh lança a lanterna, permitindo que um aliado clique sobre ela para ser puxado até ele. A lanterna também concede (?) de Escudo por 4/4/4/4/4s a Thresh e ao primeiro Campeão aliado que entrar em contato com ela.(?)
- **E · Esfolar** — Passivo: os Ataques de Thresh causam dano adicional com base no tempo desde seu último Ataque. Eles causam entre (?) e (?) de Dano Mágico. Ativo: Thresh açoita com suas correntes, Puxando ou Empurrando inimigos na direção em que agitá-las. Inimigos atingidos também sofrem 20/25/30/35/40% de Lentidão por 1/1/1/1/1s e 75/120/165/210/255 + 0.7 ability_power de Dano Mágico.(?)
- **R · A Caixa** — Thresh cria uma prisão de paredes espectrais, causando 99/99/99% de Lentidão a Campeões por 2/2/2s e 250/400/550 + 1 ability_power de Dano Mágico. As paredes quebram após uma colisão. Depois que uma é quebrada, o restante delas não causa dano e a duração da Lentidão é reduzida pela metade.(?)

## Illaoi, a Sacerdotisa Cráquem

`Illaoi` · id 420 · corpo a corpo · fighter, tank

- **P · Profetisa de um Deus Ancião** — Illaoi e os Recipientes que ela cria fazem surgir Tentáculos em terreno intransponível próximo. Tentáculos golpeiam espíritos, Recipientes e vítimas da Lição Dura de Illaoi. Tentáculos causam Dano Físico a inimigos atingidos e curarão Illaoi caso atinjam um Campeão.
- **Q · Golpe de Tentáculo** — Passivo: o dano de Esmagar aumenta em @spell.IllaoiQ:TentacleDamageAmp*100@% (atualmente @spell.IllaoiQ:TentacleDamageTotal@ de Dano Físico). Ativo: Illaoi brande seu ídolo, fazendo com que um Tentáculo Esmague à frente.(?)
- **W · Lição Dura** — O próximo Ataque de Illaoi faz com que ela salte sobre o alvo, causando 3/3.5/4/4.5/5 + 0.035 attack_damage da Vida máxima como Dano Físico adicional. Ao atacar, os Tentáculos próximos Esmagam o alvo.(?)
- **E · Teste de Espírito** — Illaoi arranca o espírito de um Campeão inimigo por 7/7/7/7/7s. O espírito pode sofrer dano como um Campeão, e 0.25/0.3/0.35/0.4/0.45 + 0.0008 attack_damage desse dano ecoa no dono. Se o espírito morrer ou se o alvo sair do alcance, o alvo ficará marcado por 4/4/4/4/4s e sofrerá 80/80/80/80/80% de Lentidão por 1.5/1.5/1.5/1.5/1.5s. Se possível, inimigos marcados fazem surgir Tentáculos. Os Tentáculos Esmagam automaticamente os espíritos e os inimigos marcados a cada 3/3/3/3/3s.(?)
- **R · Salto de Fé** — Illaoi golpeia o chão com seu ídolo, causando 150/250/350 + 0.5 attack_damage (bonus) de Dano Físico a inimigos próximos e fazendo surgir um Tentáculo para cada Campeão inimigo atingido. Pelos próximos 8/8/8s, Tentáculos ficam inalvejáveis, Esmagam 50% mais rápido, e Lição Dura tem Tempo de Recarga de @spell.IllaoiW:CooldownDuringR(?)SpellModifierDescriptionAppend@

## Rek'Sai, a Escavadora do Vazio

`Rek'Sai` · id 421 · corpo a corpo · fighter, tank

- **P · Fúria dos Xer'Sai** — Rek'Sai gera Fúria ao Atacar e atingir inimigos com Habilidades básicas. Enquanto estiver Escavada, ela consome a Fúria para restaurar Vida.
- **Q · Ira da Rainha / Sondar Presas** — Emergida: os próximos 3 Ataques de Rek'Sai dentro de 3/3/3/3/3s recebem 35/35/35/35/35% de Velocidade de Ataque e causam 0/0/0/0/0 + 0.3 attack_damage de Dano Físico adicional a inimigos próximos. Ataques redefinem a duração da Habilidade.(?)
- **W · Escavar / Emergir** — Emergida: Rek'Sai se enterra no chão, recebendo novas Habilidades, mas não pode Atacar. Rek'Sai recebe 5/10/15/20/25 de Velocidade de Movimento e fica com @VisionRadiusMod*-100@% de alcance de visão reduzido, mas inimigos próximos em movimento e que não seriam vistos normalmente têm suas posições reveladas para ela e seus aliados.(?)
- **E · Mordida Feroz / Túnel** — Emergida: Rek'Sai morde um alvo, causando @spell.RekSaiE:BaseDamageCalculation@ de Dano Físico. Com o máximo de Fúria, a mordida causa @spell.RekSaiE:EmpoweredDamageCalculation@ de Dano Verdadeiro no lugar.(?)
- **R · Investida do Vazio** — Rek'Sai alveja um inimigo a quem tenha causado dano nos últimos 5/5/5s e mergulha no solo, ficando inalvejável. Depois, ela salta sobre o alvo sem poder ser interrompida e causa 150/250/350 + 1 attack_damage (bonus) mais 15/20/25% da Vida máxima como Dano Físico, além de zerar o Tempo de Recarga de Escavar / Emergir.(?)

## Ivern, o Pai do Verde

`Ivern` · id 427 · a distância · support, mage

- **P · Amigo da Floresta** — Ivern não pode atacar ou ser atacado por monstros não épicos. Ele pode criar bosques mágicos que crescem ao longo do tempo em acampamentos da selva. Quando o bosque estiver em seu tamanho final, Ivern pode libertar os monstros para receber ouro e experiência.
- **Q · Encantador de Raízes** — Ivern conjura uma vinha, causando 80/125/170/215/260 + 0.7 ability_power de Dano Mágico e Enraizando o primeiro inimigo atingido por 1.6/1.7/1.8/1.9/2s. Aliados que Atacarem um inimigo Enraizado avançarão até o Alcance de Ataque. Reconjuração: Ivern avança diretamente para o inimigo Enraizado. Atingir monstros não épicos reduz o Tempo de Recarga de Encantador de Raízes em 50%.(?)
- **W · Formação de Arbustos** — Passivo: enquanto Ivern estiver dentro de arbustos e por 3/3/3/3/3s depois de sair, os Ataques dele causam 20/27.5/35/42.5/50 + 0.2 ability_power de Dano Mágico adicional. Aliados próximos recebem esse efeito por 1.5/1.5/1.5/1.5/1.5s e causam 10/15/20/25/30 + 0.1 ability_power de Dano Mágico. Ativo: Ivern cria um pedaço de Arbusto, revelando a área por 8/8/8/8/8s. O Arbusto persiste até que a equipe de Ivern perca a visão dele ou até 45/45/45/45/45s.(?)
- **E · Semente Engatilhada** — Ivern concede 75/115/155/195/235 + 0.5 ability_power de Escudo a um Campeão aliado ou à Margarida. Depois de 2/2/2/2/2s, o Escudo explode, causando 70/90/110/130/150 + 0.8 ability_power de Dano Mágico e 40/45/50/55/60% de Lentidão aos inimigos por 2/2/2/2/2s. Se Semente Engatilhada detonar e nenhum Campeão inimigo for atingido enquanto o Escudo persistir, o aliado receberá 75/115/155/195/235 + 0.5 ability_power de Escudo por 2/2/2/2/2s.(?)
- **R · Margarida!** — Ivern invoca sua amiga Margarida para participar do combate por 45/45/45s. Pancada da Margarida!: o 3º Ataque consecutivo da Margarida no mesmo Campeão ou monstro épico enviará uma onda de choque, causando 90/140/190 + 0.5 ability_power de Dano Mágico a todos os inimigos atingidos e Arremessando-os ao Ar por 1/1/1s. Esse efeito pode ocorrer somente uma vez a cada 3/3/3s. Reconjuração: faz Margarida atacar ou se mover.(?)

## Kalista, a Lança da Vingança

`Kalista` · id 429 · a distância · marksman

- **P · Aprumo Marcial** — Realize um comando de movimentação enquanto Kalista carrega seu ataque básico ou Perfurar para saltar em uma curta distância quando ela lançar o ataque.
- **Q · Perfurar** — Kalista arremessa uma lança, causando 10/75/140/205/270 + 1.05 attack_damage de Dano Físico ao primeiro alvo atingido. Se a habilidade abater o alvo, a lança continuará seguindo adiante, carregando quaisquer acúmulos de Lacerar para o próximo alvo atingido. Kalista pode avançar depois de usar a habilidade utilizando Aprumo Marcial.(?)
- **W · Vigia** — Passivo: quando Kalista e seu aliado Em Juramento Atacam o mesmo alvo, ela causa 10/12/14/16/18% da Vida máxima como Dano Mágico. O efeito tem um Tempo de Recarga de 10/10/10/10/10s por alvo e um limite de 100/125/150/175/200 contra não Campeões. Ativo: Kalista envia um fantasma para patrulhar uma área por três rondas. Campeões vistos são revelados por 4s. A Habilidade tem 2 cargas (recarrega a cada 90/80/70/60/50s).(?)
- **E · Lacerar** — Passivo: as lanças de Kalista permanecem nos alvos por 4s, sem limite de acúmulos. Ativo: Kalista arranca as lanças de inimigos próximos, causando 5/15/25/35/45 + 0.65 ability_power + 0.7 attack_damage mais 7/14/21/28/35 + 0.5 ability_power + 0.2 attack_damage de Dano Físico por lança após a primeira. Causa 0.1/0.18/0.26/0.34/0.42 + 0.0005 ability_power de Lentidão aos inimigos atingidos por 2/2/2/2/2s. Se abater ao menos um alvo, a Habilidade restituirá o Tempo de Recarga e 10/15/20/25/30 de Mana.(?)
- **R · Chamado do Destino** — Kalista coloca o aliado Em Juramento em Estase e o puxa em direção a ela por 4s. O aliado Em Juramento pode clicar para ser lançado, parando no primeiro Campeão atingido e Arremessando ao ar todos os inimigos próximos. Caso o aliado Em Juramento atinja um Campeão, será posicionado no limite máximo de seu Alcance de Ataque.(?)

## Bardo, o Protetor Andarilho

`Bard` · id 432 · a distância · support, mage

- **P · Chamado do Viajante** — Mipes: Bardo atrai pequenos espíritos que ajudam em seus ataques básicos e causam Dano Mágico adicional. Quando Bardo coleta sinos o suficiente, os mipes também causam dano em uma área e reduzem a velocidade de inimigos atingidos. Sinos: sinos ancestrais aparecem aleatoriamente para Bardo coletar. Eles concedem experiência, Velocidade de Movimento fora de combate e restauram Mana.
- **Q · Prisão Cósmica** — Bardo libera um raio de energia espiritual, causando 80/120/160/200/240 + 0.8 ability_power de Dano Mágico aos dois primeiros inimigos atingidos. O primeiro alvo atingido sofre 60/60/60/60/60% de Lentidão por 1/1.2/1.4/1.6/1.8s. Se o raio atingir um segundo inimigo ou uma parede, todos os inimigos atingidos são Atordoados por 1/1.2/1.4/1.6/1.8s. (?)
- **W · Santuário do Protetor** — Bardo cria um santuário de Vida que concede 0.2/0.225/0.25/0.275/0.3 + 0.0006 ability_power de Velocidade de Movimento, decaindo ao longo de 1.5/1.5/1.5/1.5/1.5s, e restaura pelo menos 25/50/75/100/125 + 0.4 ability_power de Vida ao primeiro aliado que entrar nele. O santuário cresce e pode restaurar até 50/87.5/125/162.5/200 + 0.7 ability_power de Vida depois de existir por 5/5/5/5/5s. Bardo pode ter até 3/3/3/3/3 santuários ativos ao mesmo tempo. Se um Campeão inimigo tocar no santuário, o santuário é destruído. Esta Habilidade tem 2/2/2/2/2 cargas. Santuários ativos no momento: (?)/(?)(?)
- **E · Jornada Mágica** — Bardo abre um portal unidirecional através de um terreno por 10/10/10/10/10s. Qualquer Campeão pode entrar no portal ao se mover em direção a ele enquanto está perto da entrada.(?)
- **R · Têmpera do Destino** — Bardo lança um campo mágico de energia protetora em uma área, colocando todas as unidades e estruturas atingidas em Estase por 2.5/2.5/2.5s.(?)

## Rakan, O Charmoso

`Rakan` · id 497 · a distância · support

- **P · Plumas Mágicas** — Rakan recebe um escudo periodicamente.
- **Q · Pena Reluzente** — Rakan arremessa uma pluma mágica que causa 70/115/160/205/250 + 0.7 ability_power de Dano Mágico ao primeiro inimigo atingido. Caso a pluma atinja um Campeão ou um monstro épico da selva, Rakan restaura 210/210/210/210/210 + 0.55 ability_power da própria Vida e da Vida de aliados próximos depois de 3/3/3/3/3s ou ao encostar em um Campeão aliado.(?)
- **W · Entrada Triunfal** — Rakan avança e dá um salto, Arremessando ao ar os inimigos por 1/1/1/1/1s e causando 70/120/170/220/270 + 0.8 ability_power de Dano Mágico.(?)
- **E · Dança da Batalha** — Rakan avança em direção a um Campeão aliado, concedendo a ele 50/75/100/125/150 + 0.7 ability_power de Escudo por 3/3/3/3/3s. Rakan pode Reconjurar a Habilidade mais uma vez dentro de 5/5/5/5/5s.(?)
- **R · Rapidez** — Rakan recebe 75/75/75% de Velocidade de Movimento por 4/4/4s. Rakan causa 100/200/300 + 0.5 ability_power de Dano Mágico e Encanta os inimigos por 1/1.25/1.5s na primeira vez que encosta neles. O primeiro Campeão encantado concede a Rakan 150/150/150% de Velocidade de Movimento, que decai ao longo do tempo.(?)

## Xayah, a Rebelde

`Xayah` · id 498 · a distância · marksman

- **P · Cortes Certeiros** — Depois de usar uma habilidade, os próximos ataques básicos de Xayah acertam todos os alvos pelo caminho e deixam uma Pluma.
- **Q · Punhais Duplos** — Xayah arremessa duas adagas, causando 45/60/75/90/105 + 0.5 attack_damage (bonus) de Dano Físico e deixando duas Plumas. Alvos atingidos depois do primeiro sofrem 22.5/30/37.5/45/52.5 + 0.25 attack_damage (bonus) de dano de cada adaga.(?)
- **W · Plumagem Mortífera** — Xayah cria uma tempestade de lâminas por 4/4/4/4/4s, recebendo 35/40/45/50/55% de Velocidade de Ataque e uma lâmina secundária que causa 25/25/25/25/25% de dano com seus Ataques. Quando a lâmina secundária atinge um Campeão, Xayah recebe 30/30/30/30/30% de Velocidade de Movimento por 1.5/1.5/1.5/1.5/1.5s. Se Rakan está por perto, ele recebe os efeitos da Habilidade, mas, quando Xayah atinge um alvo, ele recebe Velocidade de Movimento.(?)
- **E · Invocadora das Lâminas** — Xayah puxa de volta todas as Plumas, e cada uma causa 50/65/80/95/110 + 0.4 attack_damage (bonus) de Dano Físico. Se 3/3/3/3/3 ou mais Plumas atingirem um inimigo, ele será Enraizado por 1.25/1.25/1.25/1.25/1.25s.(?)
- **R · Tempestade de Plumas** — Xayah salta, tornando-se Inalvejável e recebendo efeito Fantasma por 1,5. Depois, ela arremessa adagas que causam 200/300/400 + 1 attack_damage (bonus) de Dano Físico e deixam para trás uma fileira de Plumas.(?)

## Ornn, O Fogo sob a Montanha

`Ornn` · id 516 · corpo a corpo · tank

- **P · Forja Viva** — Ornn recebe Armadura e Resistência Mágica adicionais de todas as origens. Ele pode gastar ouro para forjar itens não consumíveis em qualquer lugar. Além disso, ele pode criar itens magistrais para si e para aliados.
- **Q · Ruptura Vulcânica** — Ornn bate no chão, criando uma fissura que causa 20/45/70/95/120 + 1.1 attack_damage de Dano Físico e 40/40/40/40/40% de Lentidão por 2/2/2/2/2s. Um pilar de rocha se forma na extremidade da fissura por 4/4/4/4/4s. (?)
- **W · Fôlego do Fole** — Ornn cospe fogo e pisoteia à frente sem poder ser interrompido, causando 12/13/14/15/16% da Vida máxima como Dano Mágico ao longo de 0.75/0.75/0.75/0.75/0.75s. Inimigos atingidos pelo último sopro de fogo se tornam Frágeis por 3/3/3/3/3s. Efeitos Imobilizadores têm a duração aumentada em 30% contra alvos Frágeis e causam 0.17/0.17/0.17/0.17/0.17 da Vida máxima como Dano Mágico adicional. Os Ataques de Ornn contra alvos Frágeis os Empurram, causando dano adicional.(?)
- **E · Investida Calcinante** — Ornn avança, causando 80/125/170/215/260 + 0.4 armor (bonus) + 0.4 magic_resist (bonus) de Dano Físico. Se colidir com um terreno, ele cria uma onda de choque que Arremessa ao ar os inimigos por 1.25/1.25/1.25/1.25/1.25s e aplica o mesmo dano aos que não foram atingidos pelo avanço. O avanço de Ornn destrói pilares de magma e terrenos criados por inimigos.(?)
- **R · Chamado do Deus da Forja** — Ornn invoca um elemental de lava gigantesco que corre em disparada na direção dele, causando 125/175/225 + 0.2 ability_power de Dano Mágico, aplicando Frágil e causando 40/50/60% de Lentidão aos inimigos atingidos por 3/3/3s. Ornn pode Reconjurar para avançar com uma cabeçada. Se Ornn avançar contra o elemental, vai redirecioná-lo e fortalecê-lo. Isso fará o elemental Arremessar ao ar o primeiro Campeão atingido por 1/1/1s e os Campeões subsequentes por 0.5/0.5/0.5s. O elemental também causa 125/175/225 + 0.2 ability_power de Dano Mágico e aplica Frágil novamente.(?)

## Sylas, o Abjugado

`Sylas` · id 517 · corpo a corpo · mage, assassin

- **P · Explosão de Petricita** — Após conjurar uma habilidade, Sylas armazena uma carga de Explosão de Petricita. Seus ataques básicos gastam uma carga e agitam as correntes energizadas ao redor dele, causando Dano Mágico adicional a inimigos atingidos. Enquanto Sylas tiver pelo menos uma carga de Explosão de Petricita, ele recebe Velocidade de Ataque.
- **Q · Correntes-Chicote** — Sylas chicoteia com suas correntes, causando 40/60/80/100/120 + 0.4 ability_power de Dano Mágico e 0.15/0.2/0.25/0.3/0.35 de Lentidão por 1.5/1.5/1.5/1.5/1.5s. O local onde as correntes se cruzam explode, causando mais 60/115/170/225/280 + 0.8 ability_power de Dano Mágico.(?)
- **W · Regicida** — Sylas avança até um inimigo com uma força mágica, causando 75/110/145/180/215 + 0.6 ability_power de Dano Mágico. Sylas restaura 20/40/60/80/100 + 0.3 ability_power + 0.05 max_health (bonus) - 40/80/120/160/200 + 0.6 ability_power + 0.1 max_health (bonus) de Vida contra Campeões com base em sua Vida perdida (a cura máxima é atingida quando estiver com 40/40/40/40/40% ou menos de Vida).(?)
- **E · Evasão / Abdução** — Sylas avança rapidamente e prepara uma Reconjuração por 3,5s. Reconjuração: Sylas arremessa as correntes, puxando-se até o primeiro inimigo atingido. Ele causa 80/130/180/230/280 + 0.8 ability_power de Dano Mágico ao alvo e o Arremessa ao ar por 0.5/0.5/0.5/0.5/0.5s.(?)
- **R · Usurpar** — Sylas usurpa um Campeão inimigo, podendo conjurar uma cópia da ultimate do inimigo com base no nível da ultimate e nos atributos de Sylas. Usurpar um inimigo impõe ao alvo um Tempo de Recarga de 200/200/200% (modificado pela Aceleração de Habilidade de Sylas) do Tempo de Recarga da ultimate usurpada, com um período mínimo de 40/40/40s. Durante esse período, Sylas não poderá usurpar novamente o mesmo inimigo.(?)

## Neeko, a Camaleoa Curiosa

`Neeko` · id 518 · a distância · mage, support

- **P · Encanto Inerente** — Neeko pode imitar um Campeão aliado ou outras unidades no mapa. O disfarce será quebrado ao sofrer efeitos de Controle de Grupo imobilizantes, conjurar habilidades de dano, causar dano a torres inimigas como um não Campeão ou sofrer dano equivalente à barra de Vida do disfarce.
- **Q · Explosão Florescente** — Neeko arremessa uma semente que floresce para causar 60/110/160/210/260 + 0.6 ability_power de Dano Mágico. Se abater uma unidade ou acertar um Campeão ou monstro grande, ela florescerá novamente, causando 35/60/85/110/135 + 0.25 ability_power de Dano Mágico. Máximo de 2 florescências adicionais.(?)
- **W · Metamorfa** — Passivo: cada 3º Ataque causa 30/65/100/135/170 + 0.6 ability_power de Dano Mágico adicional e aumenta a Velocidade de Movimento de Neeko em 10/17.5/25/32.5/40% por 1/1/1/1/1s. Ativo: Neeko fica invisível por 0.5/0.5/0.5/0.5/0.5s e projeta um clone que dura 3/3/3/3/3s. Neeko e o clone recebem 20/25/30/35/40% de Velocidade de Movimento por 3/3/3/3/3s. O clone pode ser controlado usando a tecla de atalho de mover mascote ou ao Reconjurar a Habilidade. O clone repete as conjurações, os emotes e o retorno dela.(?)
- **E · Farpas Emaranhadas** — Neeko lança um emaranhado que causa 70/105/140/175/210 + 0.65 ability_power de Dano Mágico e Enraíza por 0.7/0.9/1.1/1.3/1.5s. O emaranhado é fortalecido depois de atingir um inimigo, aumentando de tamanho, movendo-se mais rápido e Enraizando por 1.8/2.1/2.4/2.7/3s.(?)
- **R · Florescer Repentino** — Após um intervalo, Neeko salta, Arremessando ao ar todos os inimigos próximos por 0.6/0.6/0.6s. Depois, ela bate com força no chão, causando 150/350/550 + 1.2 ability_power de Dano Mágico a todos os inimigos próximos e Atordoando-os por 0.75/0.75/0.75s. Esta Habilidade pode ser preparada em segredo se Neeko estiver disfarçada. 0.5/0.5/0.5s após conjurar esta Habilidade, o disfarce de Neeko se desfaz.(?)

## Aphelios, a Arma dos Devotos

`Aphelios` · id 523 · a distância · marksman

- **P · O Assassino e a Profetisa** — Aphelios empunha 5 Armas Lunari criadas por sua irmã, Alune. Ele tem acesso a duas ao mesmo tempo: uma principal e uma secundária. Cada arma tem um ataque básico e uma Habilidade únicos. Ataques e habilidades consomem uma munição da arma. Quando fica sem munição, Aphelios descarta a arma e Alune invoca a próxima.
- **Q · Habilidades da arma** — (?)
- **W · Fase** — Alterne entre a arma principal e a secundária, equipando Gravitum.(?)
- **E · Sistema de Ordenação de Armas** — (?)
- **R · Vigília do Plenilúnio** — Aphelios dispara uma rajada concentrada de luar que explode ao atingir um Campeão, causando 125/175/225 + 1 ability_power + 0.2 attack_damage (bonus) de Dano Físico a inimigos ao redor. Depois, com a arma principal, Aphelios Ataca todos os Campeões atingidos. {{ Spell_ApheliosR_WeaponMod_(?) }}(?)

## Rell, a Dama de Ferro

`Rell` · id 526 · corpo a corpo · tank, support

- **P · A Ferro e Fogo** — Os Ataques e Habilidades de Rell causam Dano Mágico adicional e roubam Armadura e Resistência Mágica ao contato.
- **Q · Golpe Estilhaçador** — Rell golpeia com a lança, Atordoando os alvos por 0.65/0.65/0.65/0.65/0.65s, destruindo todos os Escudos e causando 60/100/140/180/220 + 0.6 ability_power de Dano Mágico.(?)
- **W · Ferromante: Queda Esmagadora** — Passivo – Espontaneidade Montada: Rell recebe @spell.RellW_Dismount:MountedMoveSpeed@ de Velocidade de Movimento enquanto montada. Ativo – Ferromante: Queda Esmagadora: Rell salta da montaria, Arremessando ao ar os inimigos e causando a eles @spell.RellW_Dismount:DismountDamage@ de Dano Mágico. Rell recebe @spell.RellW_Dismount:Shield@ de Escudo que dura até ela montar novamente. Rell entra na forma blindada, recebendo mais @spell.RellW_Dismount:ResistanceIncrease*100@% de Armadura e Resistência Mágica, @spell.RellW_Dismount:DismountedASBoost*100@% de Velocidade de Ataque e @spell.RellW_Dismount:DismountedRangeBoost@ de Alcance de Ataque. Na forma blindada, ela pode usar Ferromante: Montaria.(?)
- **E · Investida Absoluta** — Rell e um aliado avançam, recebendo 15/15/15/15/15% de Velocidade de Movimento, aumentada para 30/30/30/30/30% na direção de Campeões inimigos ou um do outro por 3/3/3/3/3s. O próximo ataque de Rell ou Golpe Estilhaçador explode em uma área, causando 0.05/0.055/0.06/0.065/0.07 + 0.0003 ability_power da Vida máxima como Dano Mágico.(?)
- **R · Tempestade Magnética** — Rell explode em uma fúria magnética, Arrastando violentamente os inimigos próximos em sua direção. Depois, ela continua Arrastando os inimigos próximos e causa 150/250/350 + 1.1 ability_power de Dano Mágico ao longo dos próximos 2/2/2s.(?)

## Pyke, o Estripador das Águas Sangrentas

`Pyke` · id 555 · corpo a corpo · support, assassin

- **P · Dádiva dos Afogados** — Quando Pyke está escondido dos inimigos, ele regenera o dano recebido recentemente de Campeões. Pyke também não ganha Vida máxima adicional de nenhuma fonte, mas ganha DdA adicional.
- **Q · Espeto de Osso** — Pressionar: Pyke esfaqueia, causando 100/150/200/250/300 + 0.75 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido, priorizando Campeões. Depois, a Habilidade causa 90/90/90/90/90% de Lentidão por 1/1/1/1/1s. Segurar: Pyke lança o arpão, causando 100/150/200/250/300 + 0.75 attack_damage (bonus) de Dano Físico ao primeiro inimigo atingido e Puxando-o em sua direção. Depois, o inimigo sofre 90/90/90/90/90% de Lentidão por 1/1/1/1/1s. Se Pyke atingir um Campeão inimigo ou se a canalização não for concluída, 75/75/75/75/75% do custo de Mana será restituído.(?)
- **W · Mergulho Fantasma** — Pyke recebe Camuflagem e 45/45/45/45/45 + 2 lethality% de Velocidade de Movimento que decai ao longo de 5/5/5/5/5s.(?)
- **E · Ressaca Espectral** — Pyke avança, deixando para trás um fantasma afogado que retorna para ele após um curto período. O fantasma Atordoa por 1.25/1.25/1.25/1.25/1.25 + 0.01 lethalitys e causa 100/150/200/250/300 + 1 attack_damage (bonus) de Dano Físico a Campeões.(?)
- **R · Morte das Profundezas** — Pyke acerta todos os Campeões inimigos em um X, teleportando-se para a área e executando alvos com Vida abaixo de 250/250/250 + 0.8 attack_damage (bonus) + 1.5 lethality. Campeões acima do limite e não Campeões sofrem Dano Físico equivalente a 50/50/50% do valor (125/125/125 + 0.4 attack_damage (bonus) + 0.75 lethality). Quando um Campeão inimigo é abatido dentro do X, Pyke pode Reconjurar a Habilidade sem custo dentro de 20/20/20s. Se ele executar o Campeão, o último aliado que prestou assistência também receberá o ouro do abate. Se não executar, ele ainda recebe o ouro do abate como se tivesse executado. (?)

## Vex, a Melancolista

`Vex` · id 711 · a distância · mage

- **P · Destruição e Escuridão** — Periodicamente, Vex é fortalecida, fazendo com que a próxima Habilidade básica dela cause Temor aos inimigos e interrompa avanços. Sempre que um inimigo próximo avança, Vex aplica uma marca que é consumida para causar dano adicional, além de reduzir o Tempo de Recarga do estado fortalecido.
- **Q · Rajada Mistral** — Vex lança uma onda de névoa que causa 70/115/160/205/250 + 0.7 ability_power de Dano Mágico. Depois de um intervalo, a onda fica menor e mais rápida. Consome a Escuridão dos inimigos atingidos.(?)
- **W · Espaço Pessoal** — Vex recebe 50/75/100/125/150 + 0.75 ability_power de Escudo por 2.5/2.5/2.5/2.5/2.5s e emite uma onda de choque que causa 80/120/160/200/240 + 0.3 ability_power de Dano Mágico. Consome a Escuridão dos inimigos atingidos.(?)
- **E · Penumbra Iminente** — Vex faz a Sombra voar até um determinado local, aumentando de tamanho durante o trajeto. Ao chegar, ela causa 50/70/90/110/130 + 0.4 ability_power de Dano Mágico e 30/35/40/45/50% de Lentidão por 2/2/2/2/2s. Abater um inimigo com esta Habilidade reduz o Tempo de Recarga de Destruição e Escuridão em 10/10/10/10/10%. Aplica Escuridão aos inimigos atingidos.(?)
- **R · Onda Sombria** — A Sombra se empolga e avança, causando @spell.VexR:RDamageCalc@ de Dano Mágico e marcando o primeiro Campeão inimigo atingido por 4s. Reconjuração: avança até o Campeão marcado, causando @spell.VexR:RecastDamageCalc@ de Dano Mágico ao chegar. Se o Campeão marcado morrer dentro de @spell.VexR:TakedownWindow@s depois de sofrer dano desta Habilidade, o Tempo de Recarga dela é temporariamente redefinido.(?)

## Yone, o Inesquecido

`Yone` · id 777 · corpo a corpo · fighter, assassin

- **P · Estilo do Caçador** — Yone causa Dano Mágico a cada segundo Ataque. Além disso, sua Chance de Acerto Crítico é aumentada.
- **Q · Aço Mortal** — Yone golpeia à frente, causando 25/50/75/100/125 + 1.1 attack_damage de Dano Físico. Ao contato, concede um acúmulo por 6/6/6/6/6s. Com 2 acúmulos, essa Habilidade faz Yone avançar com uma corrente de vento que Arremessa ao Ar os inimigos por 0.75/0.75/0.75/0.75/0.75s e causa 25/50/75/100/125 + 1.1 attack_damage de Dano Físico. (?)
- **W · Fenda Espiritual** — Yone desfere um golpe à frente, causando 5/10/15/20/25 + 4/4.5/5/5.5/6% da Vida máxima como Dano Físico e 5/10/15/20/25 + 4/4.5/5/5.5/6% da Vida máxima como Dano Mágico. Caso Yone atinja um inimigo, recebe 90/90/90/90/90 + 0.65 attack_damage (bonus) de Escudo por 1.5/1.5/1.5/1.5/1.5s. O valor total do Escudo aumenta a cada Campeão atingido. (?)
- **E · Desatar da Alma** — Yone assume a forma espiritual por 5/5/5/5/5s, deixando o corpo para trás pela duração da Habilidade e recebendo 10/10/10/10/10% - 30/30/30/30/30% de Velocidade de Movimento crescente. Quando a forma espiritual termina, Yone retorna para o corpo e repete 25/27.5/30/32.5/35% de todo o dano de Ataques e Habilidades que tiver causado a Campeões enquanto estava na forma espiritual. Você pode Reconjurar esta Habilidade durante a forma espiritual. Reconjuração: encerra a forma espiritual antecipadamente.(?)
- **R · Destino Selado** — Yone atinge todos os inimigos que estiverem no caminho, causando 100/200/300 + 0.4 attack_damage (bonus) de Dano Físico e 100/200/300 + 0.4 attack_damage (bonus) de Dano Mágico, teleportando-se para trás do último Campeão atingido e arremessando ao ar os inimigos, puxando-os na direção de Yone.(?)

## Ambessa, Matriarca da Guerra

`Ambessa` · id 799 · corpo a corpo · fighter, assassin

- **P · Passo do Cão-dragão** — Utilizar um comando de ataque ou de movimento enquanto conjura uma habilidade fará com que Ambessa avance a uma curta distância assim que a habilidade for conjurada, recebendo alcance, dano e Velocidade de Ataque adicionais em seu próximo ataque e recuperando energia.
- **Q · Golpe Ardiloso / Pancada Cortante** — Golpe Ardiloso: Ambessa golpeia à frente com suas lâminas, causando 40/60/80/100/120 + 0.6 attack_damage (bonus) + 0.04/0.045/0.05/0.055/0.06 + 0.0003 attack_damage (bonus) de Dano Físico com base na Vida máxima a inimigos no limite do golpe. O restante dos inimigos sofre 0.5/0.5/0.5/0.5/0.5 de dano. Atingir um inimigo prepara uma Pancada Cortante. Pancada Cortante: Ambessa bate as lâminas no chão, causando 50/75/100/125/150 + 0.9 attack_damage (bonus) + 0.04/0.045/0.05/0.055/0.06 + 0.0004 attack_damage (bonus) de Dano Físico com base na Vida máxima contra o primeiro inimigo atingido. Os demais inimigos sofrem 0.5/0.5/0.5/0.5/0.5 de dano.(?)
- **W · Repúdio** — Ambessa recebe 320/320/320/320/320 + 1.5 attack_damage (bonus) de Escudo por 1.5/1.5/1.5/1.5/1.5s e canaliza por 0.5/0.5/0.5/0.5/0.5s. Depois, ela golpeia o chão, causando 50/75/100/125/150 + 0.5 attack_damage (bonus) de Dano Físico a inimigos próximos, aumentado para 75/112.5/150/187.5/225 + 0.75 attack_damage (bonus) de Dano Físico caso o Escudo bloqueie dano de algum Campeão inimigo, monstro grande ou estrutura.(?)
- **E · Lacerar** — Ambessa agita as correntes ao seu redor, causando aos inimigos 40/60/80/100/120 + 0.5 attack_damage (bonus) de Dano Físico e 99/99/99/99/99% de Lentidão que decai ao longo de 1/1/1/1/1s. Iniciar Passo do Cão-dragão com esta Habilidade acionará um golpe adicional.(?)
- **R · Execução Pública** — Passivo: Ambessa recebe 10/20/30% %i:scaleAPen% de Penetração de Armadura e suas Habilidades a curam em 0.15/0.175/0.2 + 0.5 omnivamp do dano causado. Ativo: Ambessa fica Incontrolável e se teleporta até o Campeão inimigo mais distante em linha reta, Suprimindo o alvo por 0.75/0.75/0.75s e arremessando-o no chão, causando 150/250/350 + 0.8 attack_damage (bonus) de Dano Físico e Atordoando-o por 0.4/0.4/0.4s.(?)

## Mel, o Reflexo da Alma

`Mel` · id 800 · a distância · mage, support

- **P · Esplendor Calcinante** — Sempre que Mel usa uma habilidade, ela recebe três projéteis adicionais (até um máximo de nove) para o próximo ataque. Quando Mel causa dano com uma habilidade ou um ataque, ela aplica Deslumbrar, que pode acumular infinitamente. Se o inimigo for atingido por Mel com dano suficiente de Deslumbrar, os acúmulos serão consumidos para executar o alvo.
- **Q · Rajada Radiante** — Mel dispara uma saraivada de 6/7/8/9/10 projéteis que explodem na área ao redor de um local-alvo. O primeiro acerto da explosão causa 60/85/110/135/160 + 0.55 ability_power de Dano Mágico e cada explosão subsequente causa 5/7/9/11/13 + 0.05 ability_power de Dano Mágico, somando até 85/127/173/223/277 + 0.8 ability_power de Dano Mágico no total.(?)
- **W · Refutação** — Mel forma ao redor de si uma barreira que reflete projéteis de Campeões inimigos, recebendo 80/110/140/170/200 + 0.7 ability_power Escudos por 0.75/0.75/0.75/0.75/0.75s e 40/40/40/40/40% de Velocidade de Movimento decrescente por 0.75/0.75/0.75/0.75/0.75s. Projéteis refletidos causam 0.4/0.45/0.5/0.55/0.6 + 0.0005 ability_power do dano original como Dano Mágico.(?)
- **E · Armadilha Solar** — Mel dispara um orbe radiante, Enraizando inimigos pegos no centro por 1.5/1.5/1.5/1.5/1.5s e causando 60/105/150/195/240 + 0.7 ability_power de Dano Mágico. O orbe cria uma área hostil ao redor, causando 30/30/30/30/30% de Lentidão aos inimigos e 16/28/40/52/64 + 0.08 ability_power de Dano Mágico por segundo.(?)
- **R · Eclipse Dourado** — Passivo: o dano de Deslumbrar aumenta para 60/70/80 + 0.1 ability_power de Dano Mágico mais 3/4/5 + 0.0075 ability_power de Dano Mágico por acúmulo. Ativo: Mel usa seus poderes contra todos os inimigos afetados por Deslumbrar, causando 125/200/275 + 0.3 ability_power de Dano Mágico mais 4/7/10 + 0.04 ability_power de Dano Mágico por acúmulo de Deslumbrar. Só é possível conjurar quando tiver algum Campeão inimigo afetado por Deslumbrar.(?)

## Yunara, a Fé Inabalável

`Yunara` · id 804 · a distância · marksman

- **P · Voto às Primeiras Terras** — Os Acertos Críticos de Yunara causam Dano Mágico adicional.
- **Q · Cultivação do Espírito** — Passivo: Yunara causa 5/10/15/20/25 + 0.2 ability_power de Dano Mágico %i:OnHit% ao contato e ataques concedem 1/1/1/1/1 de Liberação (2/2/2/2/2 de Liberação ao atacar Campeões). Ativo: Yunara consome 8/8/8/8/8 de Liberação e recebe 0.2/0.3/0.4/0.5/0.6 de Velocidade de Ataque por 5/5/5/5/5s, além de causar 5/10/15/20/25 + 0.2 ability_power de Dano Mágico adicional %i:OnHit% ao contato. Durante este período, os ataques atingem inimigos próximos, causando 0/0/0/0/0 + 0.3 attack_damage de Dano Físico. Estado de Transcendência: esta Habilidade é ativada instantaneamente por @Spell.YunaraR:Buff_Duration(?)SpellModifierDescriptionAppend@
- **W · Arco do Julgamento | Arco da Ruína** — Yunara lança uma conta de oração que causa 55/95/135/175/215 + 0.5 ability_power + 0.85 attack_damage (bonus) de Dano Mágico e 0.99/0.99/0.99/0.99/0.99 de Lentidão que decai ao longo de 1.5/1.5/1.5/1.5/1.5s. Causa 33/57.000004/81/105.00001/129 + 0.3 ability_power + 0.51000005 attack_damage (bonus) de Dano Mágico adicional por segundo. Estado de Transcendência – Arco da Ruína: Yunara dispara um laser que causa @Spell.YunaraR:Calc_RW_Damage@ de Dano Mágico e @Spell.YunaraR:Calc_RW_Slow_Amount@ de Lentidão que decai ao longo de @Spell.YunaraR:RW_Slow_Duration(?)SpellModifierDescriptionAppend@
- **E · Passos de Kanmei | Sombra Intocável** — Yunara recebe 0.3/0.35/0.4/0.45/0.5 de Velocidade de Movimento, que aumenta para 0.45000002/0.525/0.6/0.67499995/0.75 de Velocidade de Movimento ao se mover em direção a um Campeão inimigo por 1.5/1.5/1.5/1.5/1.5s. Estado de Transcendência – Sombra Intocável: Yunara avança em uma direção.(?)
- **R · Transcendência do Eu** — Yunara entra em um Estado de Transcendência por 15/15/15s, aprimorando suas habilidades básicas durante este período.(?)

## Locke, o Exorcista das Cinzas

`Locke` · id 805 · corpo a corpo · assassin, mage

- **P · Estaca de Prata** — Os ataques de Locke causam Dano Mágico adicional ao contato, aumentado com base na Vida perdida do inimigo.
- **Q · Pregos Ritualísticos** — Locke prepara um conjunto de Pregos Anímicos para arremessar, causando 40/48/56/64/72 + 0.2 ability_power de Dano Mágico e marcando os inimigos atingidos. Os pregos aplicam 25/25/25/25/25/25/25/25/25/25/60/60/60/60/60% de Lentidão por 1/1/1/1/1/1/1/1/1/1/2/2/2/2/2s, acumulando com a quantidade de pregos que acertarem. Atacar o inimigo consome os Pregos Anímicos, causando 18/26/34/42/50 + 0.25 ability_power de Dano Mágico por acúmulo, aumentado em 20/20/20/20/20% e 40/40/40/40/40% para dois e três pregos, respectivamente.(?)
- **W · Ignição da Alma** — Locke incendeia a própria alma, recebendo 0.7/0.7/0.7/0.7/0.7 de Velocidade de Ataque e 0.4/0.4/0.4/0.4/0.4 + 0.0002 ability_power de Velocidade de Movimento que decai ao longo de 2/2/2/2/2s. Por 6/6/6/6/6s, ele sofre 2/2/2/2/2% da Vida atual como Dano Verdadeiro por segundo, mas se cura com base nos últimos 40/60/80/100/120 + 1 ability_power de dano sofrido e recebe 200/200/200/200/200 + 0.2 ability_power de Vida adicional com base na Vida perdida e no tempo decorrido, até um máximo de (?) de Vida restante. É possível reconjurar para encerrar mais cedo.(?)
- **E · Perseguição das Cinzas** — Locke se teleporta para o local-alvo e, ao chegar, causa 40/50/60/70/80 + 0.4 ability_power de Dano Mágico a inimigos próximos. O próximo Ataque dele o faz avançar em direção ao alvo, causando 40/60/80/100/120 + 0.4 ability_power de Dano Mágico a todos os inimigos no caminho. Cada acerto consome Pregos Anímicos. Se Locke conseguir uma eliminação, o Tempo de Recarga desta Habilidade é zerado.(?)
- **R · Purgatório** — Locke arremessa um artefato de restrição no local-alvo, causando 150/225/300 + 0.6 ability_power de Dano Mágico e 99/99/99% de Lentidão a inimigos na área, que decai ao longo de 2/2/2s. O artefato aplica uma marca por 5/5/5s nos inimigos atingidos, selando Campeões inimigos marcados que ficarem com menos de (?)% de Vida e redefinindo a duração da marca nos outros Campeões afetados. Ao final da duração, se pelo menos um Campeão tiver sido selado, o artefato cai no chão. Locke pode pegar o artefato para aumentar permanentemente o limiar de execução em 0.5/0.5/0.5% e restituir 20/20/20% do Tempo de Recarga total a cada Campeão selado. (?)

## Sett, o Chefe

`Sett` · id 875 · corpo a corpo · fighter, tank

- **P · Ousadia da Arena** — Os ataques básicos de Sett alternam entre socos de direita e esquerda. Socos de direita são levemente mais fortes e rápidos. Como Sett odeia perder, recebe Regeneração de Vida adicional com base na Vida perdida.
- **Q · Pancadaria** — Sett anseia por uma luta, recebendo 30/30/30/30/30% de Velocidade de Movimento por 1.5/1.5/1.5/1.5/1.5s ao se mover em direção a Campeões inimigos. Além disso, os dois próximos Ataques de Sett causam 10/20/30/40/50 mais 0.01/0.01/0.01/0.01/0.01 + 0.0001 attack_damage da Vida máxima como Dano Físico.(?)
- **W · Casca-Grossa** — Passivo: Sett armazena 100/100/100/100/100% do dano sofrido como Ousadia, com um máximo de 0/0/0/0/0 + 0.5 max_health. A Ousadia decai rapidamente 4/4/4/4/4s após o dano ter sido sofrido. Ativo: Sett consome toda a Ousadia, recebendo 100/100/100/100/100% da Ousadia consumida como Escudo que decai ao longo de 3/3/3/3/3s. Então, Sett desfere um poderoso soco, causando 80/100/120/140/160 mais 0.25/0.25/0.25/0.25/0.25 + 0.0025 attack_damage (bonus) da Ousadia consumida como Dano Verdadeiro a inimigos no centro (máximo de (?) de dano). Inimigos fora do centro sofrem Dano Físico.(?)
- **E · Quebra-Crânio** — Sett faz os inimigos de cada lado colidirem uns contra os outros, causando 50/70/90/110/130 + 0.6 attack_damage de Dano Físico e 70/70/70/70/70% de Lentidão por 0.5/0.5/0.5/0.5/0.5s. Se Sett agarrar ao menos um inimigo de cada lado, todos os inimigos são Atordoados por 1/1/1/1/1s.(?)
- **R · Hora do Show** — Sett agarra um Campeão inimigo e o Suprime enquanto o carrega para a frente e o arremessa no chão, causando 200/300/400 + 1.2 attack_damage (bonus) mais 40/50/60% da Vida adicional do inimigo agarrado como Dano Físico e 99/99/99% de Lentidão aos inimigos ao redor por 1/1/1s. Quanto mais longe os inimigos estiverem da área onde Sett aterrissar, menos dano eles sofrerão.(?)

## Lillia, o Florir Receoso

`Lillia` · id 876 · a distância · fighter, mage

- **P · Ramo Onírico** — Atingir um Campeão ou monstro com uma habilidade causa dano adicional com base na Vida máxima do alvo ao longo do tempo.
- **Q · Golpes Florescentes** — Passivo: as Habilidades que Lillia acertar concedem 0.03/0.04/0.05/0.06/0.07 + 0.0003 ability_power de Velocidade de Movimento por 6.5/6.5/6.5/6.5/6.5s, acumulando até 4/4/4/4/4 vezes. Ativo: Lillia gira seu cajado no ar, causando 35/45/55/65/75 + 0.35 ability_power de Dano Mágico e mais 35/45/55/65/75 + 0.35 ability_power de Dano Verdadeiro adicional na porção externa do círculo.(?)
- **W · Cuidado! Iiip!** — Lillia se prepara e golpeia com o cajado, causando 80/100/120/140/160 + 0.35 ability_power de Dano Mágico em uma área. Inimigos no centro do impacto sofrem 80/100/120/140/160 + 0.35 ability_power de Dano.(?)
- **E · Semente Espiral** — Lillia lança uma Semente Espiral, causando 60/85/110/135/160 + 0.5 ability_power de Dano Mágico na área onde cair, revelando inimigos e reduzindo a velocidade deles em 40/40/40/40/40% por 3/3/3/3/3s. Caso não atinja inimigos, Semente Espiral continuará avançando até colidir com um inimigo ou terreno.(?)
- **R · Cadência de Ninar** — Lillia faz com que todos os Campeões inimigos afetados por Pó dos Sonhos fiquem sob o efeito de Sonolência por 1.5/1.5/1.5s. Em seguida, eles ficam Adormecidos por 2/2/2s. Se forem despertados por dano, os inimigos sofrerão 100/150/200 + 0.4 ability_power de Dano Mágico adicional.(?)

## Gwen, A Costureira Encantada

`Gwen` · id 887 · corpo a corpo · fighter

- **P · Mil Retalhos** — Os Ataques de Gwen causam Dano Mágico adicional com base na Vida do alvo. Ela se cura em parte do dano causado a Campeões com esse efeito.
- **Q · Corte e Recorte** — Passivo: Gwen recebe 1 acúmulo ao acertar um inimigo com um Ataque (máximo de 4, com duração de 6/6/6/6/6s). Ativo: consome o acúmulo. Gwen corta uma vez, causando 10/14/18/22/26 + 0.05 ability_power de Dano Mágico. Depois, corta novamente para cada munição consumida. Por fim, corta uma última vez para causar 60/85/110/135/160 + 0.35 ability_power de Dano Mágico. O centro de cada um dos golpes converte 50/50/50/50/50% do dano em Dano Verdadeiro e aplica Mil Retalhos aos inimigos atingidos. Causa 80/80/80/80/80% de dano a tropas. Tropas com menos de 20/20/20/20/20% de Vida sofrem 1000/1000/1000/1000/1000% de dano adicional em vez de dano reduzido.(?)
- **W · Névoa Sagrada** — Gwen invoca a Névoa Sagrada, ficando inalvejável perante os inimigos (exceto torres) fora da zona por 4/4/4/4/4s ou até que ela saia da Névoa. Enquanto estiver dentro da Névoa, Gwen receberá 22/24/26/28/30 + 0.07 ability_power de Armadura e Resistência Mágica. Gwen pode Reconjurar a Habilidade uma vez para trazer a Névoa até ela. A Habilidade será Reconjurada automaticamente na primeira vez que Gwen tentar sair da área da Névoa.(?)
- **E · Avanço Afiado** — Gwen avança e fortalece os Ataques dela por 4/4/4/4/4s. Ataques fortalecidos recebem 30/42.5/55/67.5/80 de Velocidade de Ataque, 15/15/15/15/15 + 0.2 ability_power de Dano Mágico %i:OnHit% ao contato e 75/75/75/75/75 de alcance, e o primeiro acerto em um inimigo restitui 25/35/45/55/65% do Tempo de Recarga da Habilidade.(?)
- **R · Ponto-Cruz** — Primeira conjuração: arremessa uma agulha que causa 30/50/70 + 0.1 ability_power de Dano Mágico e @InitialSlow*-100@% de Lentidão por 1.5/1.5/1.5s, aplicando Mil Retalhos a todos os inimigos atingidos. Gwen pode Reconjurar a Habilidade mais 2 vezes nos próximos 6s (1/1/1s de Tempo de Recarga entre conjurações). Segunda conjuração: arremessa três agulhas, causando 90/150/210 + 0.3 ability_power de Dano Mágico. Terceira conjuração: arremessa cinco agulhas, causando 150/250/350 + 0.5 ability_power de Dano Mágico.(?)

## Renata Glasc, a Baronesa da Química

`Renata Glasc` · id 888 · a distância · support, mage

- **P · Financiamento** — Os Ataques de Renata causam dano adicional e marcam inimigos. Os aliados de Renata podem atacar inimigos marcados para causar dano adicional.
- **Q · Negócio Fechado** — Renata envia um foguete de seu braço, Enraizando o primeiro inimigo atingido por 1/1/1/1/1s e causando 80/125/170/215/260 + 0.8 ability_power de Dano Mágico. Reconjuração: Renata Puxa o inimigo em uma direção, causando 80/125/170/215/260 + 0.8 ability_power de Dano Mágico aos inimigos atingidos e Atordoando por 0.5/0.5/0.5/0.5/0.5s se o inimigo arremessado for um Campeão.(?)
- **W · Empréstimo** — Renata infunde um Campeão aliado, concedendo a ele 10/15/20/25/30 + 0.01 ability_power de Velocidade de Ataque e 10/12.5/15/17.5/20 + 0.01 ability_power de Velocidade de Movimento em direção a inimigos, aumentado para 10/15/20/25/30 + 0.01 ability_power de Velocidade de Ataque e 10/12.5/15/17.5/20 + 0.01 ability_power de Velocidade de Movimento ao longo de 5/5/5/5/5s. Eliminações redefinem a duração do efeito. Se o aliado morrer, ele volta a ficar com a Vida cheia, decaindo ao longo de 3s. Se o aliado conseguir uma eliminação enquanto a Vida dele decai, ele ficará com 20/20/20/20/20% da Vida máxima e a Vida dele parará de decair. A morte do Campeão pode ser atrasada ao receber curas ou efeitos similares durante o decaimento, mas só pode ser evitada se o Campeão conseguir uma eliminação. Campeões só podem atrasar a própria morte uma vez.(?)
- **E · Programa de Fidelidade** — Renata envia um par de foguetes quimtec, causando 65/95/125/155/185 + 0.55 ability_power de Dano Mágico e 30% de Lentidão aos inimigos atravessados por 2/2/2/2/2s. Além disso, os foguetes concedem 50/65/80/95/110 + 0.5 ability_power de Escudo aos aliados atingidos por 3/3/3/3/3s.(?)
- **R · Apropriação Agressiva** — Renata envia uma onda de produtos químicos em uma direção, fazendo com que todos os inimigos atingidos entrem em estado de Berserk por 1.25/1.75/2.25s e ataquem a unidade mais próxima, priorizando seus aliados. Enquanto estiverem em estado de Berserk, inimigos recebem 100/100/100% de Velocidade de Ataque.(?)

## Aurora, a Bruxa Entre Mundos

`Aurora` · id 893 · a distância · mage, assassin

- **P · Abjuração Espiritual** — As Habilidades e Ataques da Aurora exorcizam espíritos dos inimigos aos quais ela causar dano. Espíritos exorcizados seguem e curam Aurora.
- **Q · Feitiço Dúplice** — Dispara energia amaldiçoada em uma direção, causando 45/70/95/120/145 + 0.4 ability_power de Dano Mágico aos inimigos e amaldiçoando-os por 3.5/3.5/3.5/3.5/3.5s. Reconjurar: Encerra a maldição, puxando parte do espírito do inimigo para Aurora e causando até 45/70/95/120/145 + 0.4 ability_power de Dano Mágico aos inimigos atravessados com base na Vida perdida deles. O dano causado após o primeiro golpe é reduzido para 20%. Se a duração terminar, Aurora Reconjurará a Habilidade automaticamente.(?)
- **W · Através do Véu** — Salta em uma direção. Ao aterrissar, ingressa no reino espiritual, ficando Invisível por 1/1.15/1.3/1.45/1.6s e entrando no Saltarreinos, ganhando 20/25/30/35/40% de Velocidade de Movimento. Eliminar um Campeão inimigo redefine o Tempo de Recarga desta Habilidade.(?)
- **E · Estranheza** — Converge os reinos, lançando uma explosão de magia espiritual que causa 70/110/150/190/230 + 0.7 ability_power de Dano Mágico aos inimigos em uma área e aplica @SlowPercent*-100@% de Lentidão, valor esse que decai ao longo de 1/1/1/1/1s. Aurora dá um pequeno salto para trás ao conjurar. (?)
- **R · Entre Mundos** — Salta em uma direção. Aurora converge os reinos ao aterrissar, irradiando um pulso de energia espiritual que causa 175/275/375 + 0.7 ability_power de Dano Mágico e @SlowPercent*-100@% de Lentidão aos inimigos atingidos por 2s. A área de convergência permanece ativa por 2.5/3.25/4s, concedendo Saltarreinos a Aurora por 3.5/4.25/5s e permitindo que ela pule de uma borda da área à outra. Inimigos que tentarem entrar na área ou sair dela sofrerão @ExitSlowPercent*-100@% de Lentidão por 1.5/1.75/2s. Aurora pode reconjurar esta Habilidade para encerrar o efeito antecipadamente.(?)

## Nilah, a Alegria Irrestrita

`Nilah` · id 895 · corpo a corpo · fighter, assassin

- **P · Alegria Eterna** — Nilah recebe mais experiência ao abater tropas, além de aprimorar e compartilhar curas e Escudos conjurados por aliados próximos.
- **Q · Lâmina Sem Forma** — Passivo: Nilah recebe 0/0/0/0/0 + 0.3 critical_chance de Penetração de Armadura e seus Ataques contra Campeões restauram 0/0/0/0/0 + 0.2 critical_chance de Vida com base no dano causado. A sobrecura é convertida em Escudo por 4/4/4/4/4s. Ativo: Nilah ataca com sua lâmina-chicote, causando 0/10/20/30/40 + 1 attack_damage de Dano Físico. Se uma unidade ou estrutura inimiga for atingida, Nilah recebe 125 de Alcance de Ataque e 60/60/60/60/60% de Velocidade de Ataque. Além disso, seus Ataques acertam em cone por 4/4/4/4/4s. (?)
- **W · Véu Jubiloso** — Nilah se envolve em névoa por 2.25/2.25/2.25/2.25/2.25s, recebendo efeito Fantasma, 15/17.5/20/22.5/25% de Velocidade de Movimento, esquivando-se de Ataques e reduzindo o Dano Mágico sofrido em 25/25/25/25/25%. Enquanto estiver ativo, tocar Campeões aliados os envolve em névoa, concedendo os mesmos benefícios por 1.5/1.5/1.5/1.5/1.5s. (?)
- **E · Turbilhão** — Nilah avança através de uma unidade, causando 60/70/80/90/100 + 0.2 attack_damage (bonus) de Dano Físico aos inimigos em seu caminho.(?)
- **R · Apoteose** — Nilah gira a lâmina-chicote, causando 60/100/140 + 0.4 attack_damage (bonus) de Dano Físico ao longo de 1s. Depois, ela libera uma explosão de 125/225/325 + 1 attack_damage (bonus) de Dano Físico e Puxa os inimigos próximos para perto. Apoteose cura Nilah e aliados próximos em 0.2/0.2/0.2 + 0.1 critical_chance (+@spell.NilahQ:CritLifesteal@ Lâmina Sem Forma) do dano causado a Campeões inimigos, convertendo qualquer cura excedente em Escudo por 6/6/6s.(?)

## K'Sante, o Orgulho de Nazumah.

`K'Sante` · id 897 · corpo a corpo · tank, fighter

- **P · Instinto Valente** — As habilidades de K'Sante marcam os alvos para que sofram mais dano do próximo ataque dele. Na Forma Irrestrita, K'Sante causa mais dano com todos os ataques e habilidades.
- **Q · Golpes de Ntofo** — K'Sante bate com força a arma no chão, causando 70/100/130/160/190 + 0.4 armor (bonus) + 0.4 magic_resist (bonus) de Dano Físico e 80/80/80/80/80% de Lentidão aos inimigos por 0.5/0.5/0.5/0.5/0.5s. Se acertar, ele recebe um acúmulo de Golpes de Ntofo por 6/6/6/6/6s. Com 2 acúmulos, ele dispara uma onda de choque que Atordoa e Puxa os inimigos por 1/1/1/1/1s. Forma Irrestrita: o Tempo de Recarga é reduzido em (?)%.(?)
- **W · Criador de Caminhos** — K'Sante levanta suas armas defensivamente por 0.4/0.4/0.4/0.4/0.4s - (?)s, tornando-se Imbatível e sofrendo 30/30/30/30/30% a menos de dano. Depois, ele avança, causando 45/75/105/135/165 + 0.08/0.08/0.08/0.08/0.08 + 0.0002 armor (bonus) + 0.0002 magic_resist (bonus) da Vida máxima como Dano Físico. Inimigos atingidos são Empurrados e Atordoados por 0.5/0.5/0.5/0.5/0.5s - 1.75/1.75/1.75/1.75/1.75s (com base no tempo de carregamento). Forma Irrestrita: o Tempo de Recarga é redefinido. Causa 10/10/10/10/10% - 80/80/80/80/80% de dano adicional como Dano Verdadeiro (com base no tempo de carregamento), a redução de dano aumenta para 75/75/75/75/75% e a velocidade do avanço é aumentada, mas não Empurra nem Atordoa mais.(?)
- **E · Passo Forte** — K'Sante avança, recebendo 70/112.5/155/197.5/240 + 0.135 max_health (bonus) de Escudo por 2/2/2/2/2s. Ao avançar até um aliado, a distância do avanço é significativamente aumentada e o aliado também recebe um Escudo. Forma Irrestrita: o Tempo de Recarga é reduzido em 50/50/50/50/50% e a velocidade do avanço é aumentada.(?)
- **R · Forma Irrestrita** — K'Sante quebra seus ntofos para Empurrar um Campeão inimigo, causando 80/115/150 de Dano Físico. Depois, ele avança para trás e assume a Forma Irrestrita por 15/15/15s. Se o inimigo atingir uma parede, ele será Empurrado através dela e K'Sante golpeará novamente, causando 80/115/150 + 0.05 max_health (bonus) de Dano Físico. Enquanto K'Sante estiver na Forma Irrestrita, suas habilidades são melhoradas e ele recebe 40/60/80% de Velocidade de Ataque, 50/50/50% de Penetração de Armadura adicional e 20/20/20% de Vampirismo Universal, mas perde 35/35/35% da Vida máxima, 85/85/85% da Armadura adicional e 85/85/85% da Resistência Mágica adicional.(?)

## Smolder, o Filhote Flamejante

`Smolder` · id 901 · a distância · marksman, mage

- **P · Treinamento Dracônico** — Acertar Campeões com habilidades e abater inimigos com Super-Hálito Flamejante concede um acúmulo de Treinamento Dracônico. Os acúmulos aumentam o dano das habilidades básicas de Smolder.
- **Q · Super-Hálito Flamejante** — Smolder cospe chamas, causando 60/70/80/90/100 + 1.3 attack_damage (bonus) de Dano Físico + @spell.SmolderP:Passive_QDamageIncrease@ de Dano Mágico. Se o alvo morre, Smolder restitui 15/15/15/15/15 de Mana, uma vez por conjuração. Com base nos acúmulos de Treinamento Dracônico, a Habilidade evolui e recebe os seguintes efeitos:25/25/25/25/25 acúmulos: causa dano a todos os inimigos ao redor do alvo.125/125/125/125/125 acúmulos: provoca (?) explosões atrás do alvo que causam 50/50/50/50/50% do dano desta Habilidade.225/225/225/225/225 acúmulos: queima o alvo, causando (?) da Vida máxima como Dano Verdadeiro ao longo de 3/3/3/3/3s. Campeões inimigos que ficarem com menos de 6.5/6.5/6.5/6.5/6.5 de Vida total enquanto estiverem queimando serão abatidos instantaneamente.(?)
- **W · Atchim!** — Smolder dá um adorável espirro flamejante, causando 60/70/80/90/100 + 0.6 attack_damage (bonus) de Dano Físico e 35/35/35/35/35% de Lentidão aos inimigos atingidos por 1.5/1.5/1.5/1.5/1.5s. Atingir Campeões gera uma explosão, causando 10/35/60/85/110 + 0.8 ability_power + 0.5 attack_damage (bonus) de Dano Físico + @spell.SmolderP:Passive_WDamageIncrease@ de Dano Mágico.(?)
- **E · Voa, Voa, Voa** — Smolder levanta voo, recebendo 75/75/75/75/75% de Velocidade de Movimento e ignorando os terrenos por 1.25/1.25/1.25/1.25/1.25s. Durante o voo, Smolder bombardeia o inimigo com a Vida mais baixa (?)x (arredondando para baixo), causando 10/15/20/25/30 + 0.3 attack_damage de Dano Físico + @spell.SmolderP:EBonusDamage@ de Dano Mágico por acerto.(?)
- **R · MANHÊÊÊ!** — A mãe de Smolder sopra fogo das alturas, causando 150/250/350 + 1 ability_power + 1 attack_damage (bonus) de Dano Físico. Inimigos no centro sofrem 225/375/525 + 1.5 ability_power + 1.5 attack_damage (bonus) de Dano Físico e 40/40/40% de Lentidão por 2/2/2s. A mãe de Smolder cura o filho em 100/135/170 + 0.75 ability_power + 0.5 attack_damage (bonus) de Vida se o atingir.(?)

## Milio, A Chama Gentil

`Milio` · id 902 · a distância · support, mage

- **P · A Todo Vapor!** — As habilidades de Milio encantam aliados ao toque, fazendo com que o próximo ataque deles cause uma explosão de dano adicional e queime o alvo.
- **Q · Ultramega Chute Flamejante** — Milio chuta uma bola de fogo, Empurrando o primeiro inimigo atingido. Se atingir um inimigo, a bola ricocheteia e explode, causando 80/140/200/260/320 + 1.2 ability_power de Dano Mágico aos inimigos próximos e 0.4/0.45/0.5/0.55/0.6 + 0.0005 ability_power de Lentidão por 1.5/1.5/1.5/1.5/1.5s. Atingir pelo menos um Campeão inimigo com Ultramega Chute Flamejante restitui 50/50/50/50/50% do Custo de Mana da Habilidade.(?)
- **W · Fogueira Aconchegante** — Milio cria uma fogueira que segue Campeões aliados por 6/6/6/6/6s. Campeões aliados próximos recebem 0.1/0.125/0.15/0.175/0.2 de Alcance de Ataque e restauram 70/90/110/130/150 + 0.15 ability_power de Vida ao longo da duração. A fogueira também aplica A Todo Vapor! a cada 3/3/3/3/3s. Reconjuração: faz a fogueira seguir outro aliado.(?)
- **E · Abraços Quentinhos** — Milio envolve um Campeão aliado em chamas protetoras, concedendo 45/75/105/135/165 + 0.45 ability_power de Escudo e 12/14/16/18/20% de Velocidade de Movimento por 2.5/2.5/2.5/2.5/2.5s. Esta Habilidade tem 2 cargas, e os efeitos se acumulam em alvos repetidos.(?)
- **R · Sopro de Vida** — Milio libera uma onda de chamas reconfortantes em Campeões aliados próximos, purificando efeitos Debilitantes e Imobilizadores, restaurando 150/250/350 + 0.5 ability_power de Vida e concedendo 65/65/65% de Tenacidade por 3/3/3s.(?)

## Zaahen, O Indissociável

`Zaahen` · id 904 · corpo a corpo · fighter, assassin

- **P · Cultivação da Guerra** — Os ataques e as habilidades de Zaahen contra Campeões inimigos concedem a ele acúmulos de Determinação, que proporcionam Dano de Ataque adicional por acúmulo. Enquanto está com Determinação, Zaahen ganha mais Dano de Ataque e é capaz de ressuscitar.
- **Q · A Glaive Darkin** — Zaahen desfere dois cortes no próximo ataque, causando 15/30/45/60/75 + 0.2 attack_damage (bonus) de Dano Físico adicional e curando-se em 5/6/7/8/9% da Vida máxima. Reconjuração: o próximo ataque de Zaahen causa 25/50/75/100/125 + 0.2 attack_damage (bonus) de Dano Físico adicional e Arremessa o alvo ao ar por 0.75/0.75/0.75/0.75/0.75s.(?)
- **W · Temível Retorno** — Zaahen dá uma estocada em uma direção, causando 40/60/80/100/120 + 0.5 attack_damage (bonus) de Dano Físico aos inimigos atingidos, e Puxa-os para si, causando 30/50/70/90/110 + 0.3 attack_damage (bonus) de Dano Físico.(?)
- **E · Ímpeto Áureo** — Zaahen avança e desfere um corte ao redor de si, causando 40/60/80/100/120 + 0.5 attack_damage (bonus) de Dano Físico aos inimigos próximos. Inimigos nas extremidades do corte sofrem 60/90/120/150/180 + 0.75 attack_damage (bonus) de Dano Físico +4/4.5/5/5.5/6% da Vida máxima como Dano Mágico em vez disso.(?)
- **R · Desfecho Fatídico** — Passivo: Zaahen recebe 10/20/30% de %i:scaleAPen% Penetração de Armadura. Ativo: Zaahen se ergue e recebe 50/50/50% de Redução de Dano ao conjurar. Em seguida, Zaahen dá uma estocada para baixo, causando 250/400/550 + 2 attack_damage (bonus) de Dano Físico e restaurando 33/33/33% do dano causado a Campeões inimigos como Vida.(?)

## Hwei, o Visionário

`Hwei` · id 910 · a distância · mage, support

- **P · Assinatura do Visionário** — Hwei prepara os Campeões inimigos que sofrerem dano de suas Habilidades para deixar sua inconfundível assinatura. Atingir um inimigo com uma segunda Habilidade de dano conclui a assinatura, que fica sob ele. A assinatura detona após um breve intervalo, causando Dano Mágico a todos os inimigos dentro do alcance.
- **Q · Tema: Desastre** — Hwei pinta visões de um desastre, causando muito dano a inimigos. Fogo Devastador Hwei dispara uma bola de fogo veloz que explode no primeiro inimigo atingido, causando 50/80/110/140/170 + 0.8 ability_power mais 3/4/5/6/7% da Vida máxima como Dano Mágico. Raio Cortante Hwei envia de longe um raio que demora para cair, causando 60/85/110/135/160 + 0.3 ability_power de Dano Mágico. Atingir um alvo isolado ou Imobilizado aumenta o dano a até 120/201.875/302.5/421.875/560 + 0.6 ability_power de Dano Mágico com base na Vida perdida do alvo. Fissura Derretida Hwei dispara um trajeto de erupções, causando 20/35/50/65/80 + 0.3 ability_power de Dano Mágico em uma área e deixando para trás poças de lava que causam @spell.HweiQE:SlowPercent@% de Lentidão e 20/35/50/65/80 + 0.24 ability_power de Dano Mágico por segundo aos inimigos durante @spell.HweiQE:Duration(?)SpellModifierDescriptionAppend@
- **W · Tema: Serenidade** — Hwei pinta visões de serenidade que proporcionam utilidade para ele e para Campeões aliados. Corrente Fugaz Hwei lança uma corrente de águas turbulentas em linha reta, concedendo 30/32.5/35/37.5/40 + 0.03 ability_power de Velocidade de Movimento aos aliados. Lagoa da Reflexão Hwei forma uma lagoa protetora que concede 100/140/180/220/260 + 0.6 ability_power de Escudo aos Campeões aliados dentro dela ao longo do tempo, reduzido em @spell.HweiWW:ToolTipAllyMod*100@% para aliados. Luzes Vivas Hwei cria três luzes bruxuleantes que fortalecem as próximas três Habilidades ou Ataques dele. Cada uma causa 20/30/40/50/60 + 0.15 ability_power de Dano Mágico adicional e restaura 45/50/55/60/65 de Mana.(?)
- **E · Tema: Tormento** — Hwei pinta visões de tormentos que controlam os inimigos. Semblante Sinistro Hwei lança uma face apavorante que faz o primeiro inimigo atingido Fugir por 1/1.125/1.25/1.375/1.5s e sofrer 70/110/150/190/230 + 0.65 ability_power de Dano Mágico. Olhar do Abismo Hwei forma um olho que permanece no lugar e concede visão do local-alvo, encarando o primeiro Campeão inimigo que aparecer e lançando um projétil guiado que Enraíza o primeiro inimigo atingido por 1.2/1.4/1.6/1.8/2s, causando 70/110/150/190/230 + 0.65 ability_power de Dano Mágico. Gorja Esmagadora Hwei pinta um par de mandíbulas no local-alvo. Elas Arrastam os inimigos em direção ao centro, causando 70/110/150/190/230 + 0.65 ability_power de Dano Mágico e 40/47.5/55/62.5/70% de Lentidão que decai ao longo de 1,25s.(?)
- **R · Desespero Vertiginoso** — Hwei dispara uma visão de puro desespero que se agarra a um Campeão inimigo por 3/3/3s. A visão se expande ao longo do tempo, aplicando 10/10/10% de Lentidão aos inimigos, acumulando a cada 0,25s, e causando 10/20/30 + 0.05 ability_power de Dano Mágico por segundo. A visão se estilhaça ao ser concluída, causando 200/325/450 + 0.8 ability_power de Dano Mágico.(?)
  - **QQ · Fogo Devastador** — Hwei pinta uma bola de fogo flamejante que voa na direção-alvo. Ela explode no primeiro inimigo atingido, causando @spell.HweiQ:Tooltip_QQDamage@ mais @spell.HweiQ:Tooltip_QQBonusDamage@% da Vida máxima como Dano Mágico a todos os inimigos em uma área.(?)
  - **QW · Raio Cortante** — Hwei pinta um raio de longo alcance que demora para cair no local-alvo, causando @spell.HweiQ:Tooltip_QWDamage@ de Dano Mágico. Inimigos isolados ou Imobilizados sofrem dano aumentado com base na Vida perdida, até um máximo de @spell.HweiQ:Tooltip_QWBonusDamage@ de Dano Mágico.(?)
  - **QE · Fissura Derretida** — Hwei pinta um campo de erupções vulcânicas explosivas que deixa uma trilha de lava pelo caminho. Cada erupção causa @spell.HweiQ:Tooltip_QEDamage@ de Dano Mágico aos inimigos atingidos. Os inimigos na área da lava sofrem @spell.HweiQ:Tooltip_QEDamagePerSecond@ de Dano Mágico por segundo e @spell.HweiQE:SlowPercent@% de Lentidão. Cada poça de lava dura @spell.HweiQE:Duration(?)SpellModifierDescriptionAppend@
  - **QR · Lavar Pincel** — Hwei limpa seu pincel e retorna às Habilidades básicas. Isso não custa Mana ou Tempos de Recarga.(?)
  - **WQ · Corrente Fugaz** — Hwei pinta uma corrente de águas turbulentas em linha reta por @spell.HweiW:Tooltip_WQAreaDuration@s, concedendo @spell.HweiW:Tooltip_WQMoveSpeed@ de Velocidade de Movimento a si e aos Campeões aliados.(?)
  - **WW · Lagoa da Reflexão** — Hwei pinta uma lagoa protetora no local-alvo que dura @spell.HweiWW:AreaDuration@s e concede @spell.HweiW:Tooltip_WWStartShieldAmount@ de Escudo imediato, aumentando até @spell.HweiW:Tooltip_WWShieldAmount@ de Escudo ao longo de @spell.HweiWW:SecsToFullStack@s para Campeões aliados na área, reduzido em @spell.HweiWW:ToolTipAllyMod*100@% para aliados.(?)
  - **WE · Luzes Vivas** — Hwei pinta três luzes bruxuleantes que rodeiam-no e fazem com que as próximas três Habilidades ou Ataques dele causem @spell.HweiW:Tooltip_WEOnHitDamage@ de Dano Mágico adicional e concedam @spell.HweiW:Tooltip_WEOnHitManaRestore@ de Mana ao contato.(?)
  - **WR · Lavar Pincel** — Hwei limpa seu pincel e retorna às Habilidades básicas. Isso não custa Mana ou Tempos de Recarga.(?)
  - **EQ · Semblante Sinistro** — Hwei lança uma face apavorante que golpeia o primeiro inimigo atingido, fazendo com que ele Fuja por @spell.HweiEQ:FleeDuration@s e causando @spell.HweiE:Tooltip_EQDamage@ de Dano Mágico.(?)
  - **EW · Olhar do Abismo** — Hwei pinta um olho abissal no local-alvo que concede visão, encarando o Campeão inimigo mais próximo e lançando um projétil guiado que Enraíza o primeiro inimigo no caminho por @spell.HweiE:Tooltip_EWRootDuration@s e causa @spell.HweiE:Tooltip_EWDamage@ de Dano Mágico.(?)
  - **EE · Gorja Esmagadora** — Hwei pinta mandíbulas esmagadoras que causam @spell.HweiE:Tooltip_EEDamage@ de Dano Mágico aos inimigos atingidos e os Arrastam em direção ao centro, causando @spell.HweiE:Tooltip_EESlowAmount@% de Lentidão que decai ao longo de 1,25s.(?)
  - **ER · Lavar Pincel** — Hwei limpa seu pincel e retorna às Habilidades básicas. Isso não custa Mana ou Tempos de Recarga.(?)

## Naafiri, Fera das Cem Mordidas

`Naafiri` · id 950 · corpo a corpo · assassin, fighter

- **P · Em Maior Número** — Uma matilha surge e ataca os alvos dos ataques e habilidades de Naafiri.
- **Q · Adagas Darkin** — Naafiri arremessa lâminas contaminadas pelos Darkin, causando @spell.NaafiriQ:TotalDamageFirstCast@ de Dano Físico e sangramento, que depois causam @spell.NaafiriQ:TotalBleedDamage@ de Dano Físico ao longo de @spell.NaafiriQ:BleedDuration@s. Naafiri pode reconjurar esta Habilidade. Se os inimigos atingidos já estiverem sangrando devido à Habilidade, ela causa o dano restante do sangramento, além de @spell.NaafiriQ:TotalMinDamageSecondCast@ - @spell.NaafiriQ:TotalMaxDamageSecondCast@ de Dano Físico com base na Vida perdida do alvo. Se o alvo for um Campeão ou um monstro grande, Naafiri restaura @spell.NaafiriQ:TotalHealSecondCast@ de Vida. A Matilha salta no primeiro Campeão ou monstro atingido, atacando-o por @spell.NaafiriP:PackmateTauntDuration@s. (?)
- **W · Chamado da Matilha** — Naafiri fica Inalvejável por 1/1/1/1/1s e se prepara para caçar, fazendo surgir 2/2/2/2/2 companheiros de matilha adicionais e recebendo 0/0/0/0/0 + 0.2 attack_damage de Dano de Ataque e 20/22.5/25/27.5/30% de Velocidade de Movimento por 5/5/5/5/5s. A matilha fica Inalvejável e volta para Naafiri. (?)
- **E · Eviscerar** — Naafiri avança, causando 15/25/35/45/55 + 0.4 attack_damage (bonus) de Dano Físico, depois explode em uma tempestade de lâminas, causando 60/85/110/135/160 + 0.8 attack_damage (bonus) de Dano Físico. A matilha fica Inalvejável e volta para Naafiri, restaurando 100% de Vida. (?)
- **R · Caça dos Cães** — Naafiri avança contra um Campeão inimigo, causando 125/200/275 + 1 attack_damage (bonus) de Dano Físico e breve Lentidão. A Matilha fica Inalvejável e avança juntamente com Naafiri, causando 12.5/20/27.5 + 0.1 attack_damage (bonus) de Dano Físico para cada Companheiro. Se Naafiri conseguir uma eliminação em até 7/7/7s, ela revelará inimigos próximos e poderá reconjurar esta habilidade uma vez. A segunda conjuração concederá 100/150/200 + 1.5 attack_damage (bonus) de Escudo por 3/3/3s. (?)

---

_806 ocorrências de `(?)` acima: são números que a fonte não publica._
