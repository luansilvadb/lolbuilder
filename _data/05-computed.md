# Máximos calculados — patch 16.16

Duas tabelas, os dois resultados **exatos**: são o ótimo global, e não
heurística. O que cada um responde está no título dele — e o que ele NÃO
responde está logo abaixo.

## Página de runas de maior atributo

A página de valor máximo para cada atributo, no nível 18.

> Só 11 das 69 runas jogáveis somam atributo. As outras 58 dependem de estado
> de partida e não entram no cálculo — inclusive todas as pedras fundamentais.
> **A página abaixo maximiza um atributo, e não a força do campeão.** Escolher
> a pedra fundamental é quase sempre mais importante que o atributo, e essa
> escolha o conjunto não calcula. Ver a coluna `escopo` em `02-runes.md`.

Slot marcado _indiferente_ é slot em que **nenhuma** opção soma nada ao
atributo pedido: a escolha fica livre, e nomear uma runa ali seria publicar
um desempate como se fosse recomendação. Nas páginas abaixo, a pedra
fundamental é sempre indiferente — nenhuma delas concede atributo.

| atributo | adaptativa | primário | secundário | runas que contribuem | slots livres | total |
|---|---|---|---|---|---|---:|
| ability_haste | ad | Precisão | Feitiçaria | Transcendência (10 acel), Aceleração de Habilidade (8 acel) | 7 de 9 | 18 acel |
| attack_damage | ad | Precisão | Dominação | Força Adaptativa (9 adapt), Força Adaptativa (9 adapt) | 7 de 9 | 18 adapt |
| attack_speed_pct | ad | Precisão | Dominação | Velocidade de Ataque (10 %AS) | 8 de 9 | 10 %AS |
| heal_shield_power_pct | ad | Precisão | Determinação | Revitalizar (5 %cura e escudo) | 8 de 9 | 5 %cura e escudo |
| health | ad | Precisão | Dominação | Escalamento de Vida (180 vida), Escalamento de Vida (180 vida) | 7 de 9 | 360 vida |
| move_speed_pct | ad | Precisão | Feitiçaria | Celeridade (1 %mov), Velocidade de Movimento (2.5 %mov) | 7 de 9 | 3.5 %mov |
| tenacity_pct | ad | Precisão | Dominação | Tenacidade e Resistência a Lentidão (15 %tenacidade) | 8 de 9 | 15 %tenacidade |
| item_haste | ad | Precisão | Inspiração | Perspicácia Cósmica (10 acel item, 18 acel feitico) | 8 de 9 | 10 acel item, 18 acel feitico |
| summoner_spell_haste | ad | Precisão | Inspiração | Perspicácia Cósmica (10 acel item, 18 acel feitico) | 8 de 9 | 10 acel item, 18 acel feitico |

## Máximo de atributo por ouro

A combinação de até 6 itens que maximiza um atributo dentro de um orçamento,
respeitando a regra de botas únicas.

> **Isto não é uma build boa.** O cálculo ignora passiva e ativa de item, que
> é o que faz metade dos itens valerem o que valem. O que está aqui é
> literalmente "o máximo deste atributo que cabe neste ouro" — útil como piso
> de comparação, inútil como recomendação de jogo.

### maximo de ability_haste por ouro, ignorando efeitos de item

15500 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3003 | Cajado do Arcanjo | 2900 | 25 acel, 70 AP, 600 mana |
| 3110 | Coração Congelado | 2500 | 20 acel, 75 armadura, 400 mana |
| 4628 | Foco do Horizonte | 2700 | 25 acel, 75 AP |
| 4629 | Ímpeto Cósmico | 3000 | 25 acel, 70 AP, 350 vida, 4 %mov |
| 6617 | Regenerador de Pedra Lunar | 2200 | 20 acel, 25 AP, 125 %regen mana, 200 vida |
| 6620 | Ecos de Helia | 2200 | 20 acel, 35 AP, 125 %regen mana, 200 vida |

**Total:** 135 acel, 275 AP, 75 armadura, 250 %regen mana, 750 vida, 1000 mana, 4 %mov

### maximo de ability_power por ouro, ignorando efeitos de item

18600 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3089 | Capuz da Morte de Rabadon | 3500 | 130 AP |
| 3100 | Perdição de Lich | 2900 | 10 acel, 100 AP, 6 %mov |
| 3102 | Véu da Banshee | 3000 | 105 AP, 40 MR |
| 3157 | Ampulheta de Zhonya | 3250 | 105 AP, 50 armadura |
| 4645 | Chama Sombria | 3200 | 110 AP, 15 pen magica |
| 6655 | Eco de Luden | 2750 | 10 acel, 100 AP, 600 mana |

**Total:** 20 acel, 650 AP, 50 armadura, 15 pen magica, 40 MR, 600 mana, 6 %mov

### maximo de armor por ouro, ignorando efeitos de item

16150 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2502 | Desespero Eterno | 2800 | 15 acel, 50 armadura, 400 vida |
| 3068 | Égide de Fogo Solar | 2800 | 10 acel, 50 armadura, 350 vida |
| 3075 | Armadura de Espinhos | 2450 | 75 armadura, 150 vida |
| 3110 | Coração Congelado | 2500 | 20 acel, 75 armadura, 400 mana |
| 3143 | Presságio de Randuin | 2700 | 75 armadura, 350 vida |
| 3742 | Couraça do Defunto | 2900 | 55 armadura, 350 vida, 4 %mov |

**Total:** 45 acel, 380 armadura, 1600 vida, 400 mana, 4 %mov

### maximo de armor_penetration_pct por ouro, ignorando efeitos de item

3000 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 6694 | Rancor de Serylda | 3000 | 15 acel, 35 %pen armadura, 45 AD |

**Total:** 15 acel, 35 %pen armadura, 45 AD

### maximo de attack_damage por ouro, ignorando efeitos de item

19000 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2517 | Fome Eterna | 3100 | 65 AD, 5 %omnivamp, 20 %tenacidade |
| 3031 | Gume do Infinito | 3500 | 75 AD, 25 %crit, 30 %dano crit |
| 3072 | Sedenta por Sangue | 3400 | 80 AD, 15 %roubo de vida |
| 3074 | Hidra Raivosa | 3300 | 15 acel, 65 AD, 12 %roubo de vida |
| 3179 | Glaive Sombria | 2800 | 15 acel, 60 AD, 18 letalidade |
| 6692 | Eclipse | 2900 | 15 acel, 60 AD |

**Total:** 45 acel, 405 AD, 25 %crit, 30 %dano crit, 18 letalidade, 27 %roubo de vida, 5 %omnivamp, 20 %tenacidade

### maximo de attack_damage por ouro, ignorando efeitos de item

19000 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ap.

| id | item | custo | stats |
|---|---|---:|---|
| 2517 | Fome Eterna | 3100 | 65 AD, 5 %omnivamp, 20 %tenacidade |
| 3031 | Gume do Infinito | 3500 | 75 AD, 25 %crit, 30 %dano crit |
| 3072 | Sedenta por Sangue | 3400 | 80 AD, 15 %roubo de vida |
| 3074 | Hidra Raivosa | 3300 | 15 acel, 65 AD, 12 %roubo de vida |
| 3179 | Glaive Sombria | 2800 | 15 acel, 60 AD, 18 letalidade |
| 6692 | Eclipse | 2900 | 15 acel, 60 AD |

**Total:** 45 acel, 405 AD, 25 %crit, 30 %dano crit, 18 letalidade, 27 %roubo de vida, 5 %omnivamp, 20 %tenacidade

### maximo de attack_speed_pct por ouro, ignorando efeitos de item

15100 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2512 | Dardos de Caça-Demônios | 2650 | 45 %AS, 25 %crit, 4 %mov |
| 3032 | Flechatroz de Yun Tal | 3000 | 50 AD, 45 %AS, 0 %crit |
| 3046 | Dançarina Fantasma | 2650 | 65 %AS, 25 %crit, 10 %mov |
| 3091 | Limite da Razão | 2800 | 50 %AS, 45 MR, 20 %tenacidade |
| 3115 | Dente de Na'Shor | 2900 | 15 acel, 80 AP, 50 %AS |
| 3172 | Grevas Bélicas | 1100 | 45 %AS, 5 %roubo de vida, 45 mov |

**Total:** 15 acel, 80 AP, 50 AD, 300 %AS, 50 %crit, 5 %roubo de vida, 45 MR, 45 mov, 14 %mov, 20 %tenacidade

### maximo de base_health_regen_pct por ouro, ignorando efeitos de item

10150 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 1006 | Pérola do Rejuvenescimento | 300 | 100 %regen vida |
| 3065 | Semblante Espiritual | 2700 | 10 acel, 100 %regen vida, 400 vida, 50 MR |
| 3109 | Juramento do Cavaleiro | 2300 | 10 acel, 40 armadura, 100 %regen vida, 200 vida |
| 3211 | Capuz do Espectro | 1250 | 100 %regen vida, 200 vida, 35 MR |
| 3801 | Braçadeira Cristalina | 800 | 100 %regen vida, 200 vida |
| 6664 | Resplendor Vazio | 2800 | 10 acel, 100 %regen vida, 400 vida, 40 MR |

**Total:** 30 acel, 40 armadura, 600 %regen vida, 1400 vida, 125 MR

### maximo de base_mana_regen_pct por ouro, ignorando efeitos de item

13450 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2065 | Hino Bélico de Shurelya | 2200 | 15 acel, 50 AP, 125 %regen mana, 4 %mov |
| 3504 | Turíbulo Ardente | 2200 | 45 AP, 125 %regen mana, 10 %cura e escudo, 4 %mov |
| 4005 | Mandato Imperial | 2400 | 15 acel, 60 AP, 150 %regen mana |
| 6616 | Cajado Aquafluxo | 2250 | 10 acel, 35 AP, 125 %regen mana, 10 %cura e escudo |
| 6617 | Regenerador de Pedra Lunar | 2200 | 20 acel, 25 AP, 125 %regen mana, 200 vida |
| 6620 | Ecos de Helia | 2200 | 20 acel, 35 AP, 125 %regen mana, 200 vida |

**Total:** 80 acel, 250 AP, 775 %regen mana, 20 %cura e escudo, 400 vida, 8 %mov

### maximo de critical_chance_pct por ouro, ignorando efeitos de item

16050 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2512 | Dardos de Caça-Demônios | 2650 | 45 %AS, 25 %crit, 4 %mov |
| 2523 | Hexótica C44 | 2800 | 55 AD, 25 %crit |
| 3046 | Dançarina Fantasma | 2650 | 65 %AS, 25 %crit, 10 %mov |
| 3085 | Furacão de Runaan | 2650 | 40 %AS, 25 %crit, 5 %mov |
| 3094 | Canhão Fumegante | 2650 | 35 %AS, 25 %crit, 4 %mov |
| 6675 | Adaga Oscilante Navori | 2650 | 40 %AS, 25 %crit, 4 %mov |

**Total:** 55 AD, 225 %AS, 150 %crit, 27 %mov

### maximo de critical_damage_pct por ouro, ignorando efeitos de item

3500 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3031 | Gume do Infinito | 3500 | 75 AD, 25 %crit, 30 %dano crit |

**Total:** 75 AD, 25 %crit, 30 %dano crit

### maximo de gold_per_10 por ouro, ignorando efeitos de item

400 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3869 | Oposição Celestial | 400 | 75 %regen vida, 75 %regen mana, 9 ouro/10, 200 vida |

**Total:** 75 %regen vida, 75 %regen mana, 9 ouro/10, 200 vida

### maximo de heal_shield_power_pct por ouro, ignorando efeitos de item

12150 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3107 | Redenção | 2300 | 15 acel, 30 AP, 100 %regen mana, 10 %cura e escudo |
| 3114 | Ídolo Proibido | 600 | 50 %regen mana, 8 %cura e escudo |
| 3222 | Bênção de Mikael | 2300 | 15 acel, 100 %regen mana, 12 %cura e escudo, 250 vida |
| 3504 | Turíbulo Ardente | 2200 | 45 AP, 125 %regen mana, 10 %cura e escudo, 4 %mov |
| 6616 | Cajado Aquafluxo | 2250 | 10 acel, 35 AP, 125 %regen mana, 10 %cura e escudo |
| 6621 | Auronúcleo | 2500 | 45 AP, 100 %regen mana, 16 %cura e escudo |

**Total:** 40 acel, 155 AP, 600 %regen mana, 66 %cura e escudo, 250 vida, 4 %mov

### maximo de health por ouro, ignorando efeitos de item

17700 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2501 | Armadura Sangrenta do Suserano | 3300 | 30 AD, 550 vida |
| 2525 | Couraça Protoplasmática | 2600 | 20 acel, 600 vida |
| 3083 | Armadura de Warmog | 3100 | 100 %regen vida, 1000 vida |
| 3084 | Coração de Aço | 3000 | 100 %regen vida, 900 vida |
| 3119 | Aproximação Invernal | 2400 | 15 acel, 550 vida, 500 mana |
| 3748 | Hidra Titânica | 3300 | 40 AD, 600 vida |

**Total:** 35 acel, 70 AD, 200 %regen vida, 4200 vida, 500 mana

### maximo de lethality por ouro, ignorando efeitos de item

17000 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2520 | Quebra-Bastião | 3000 | 15 acel, 55 AD, 22 letalidade |
| 3142 | Lâmina Fantasma de Youmuu | 2800 | 55 AD, 18 letalidade, 4 %mov |
| 3179 | Glaive Sombria | 2800 | 15 acel, 60 AD, 18 letalidade |
| 6696 | Arco do Axioma | 2750 | 20 acel, 55 AD, 18 letalidade |
| 6697 | Húbris | 2800 | 10 acel, 55 AD, 18 letalidade |
| 6698 | Hidra Profana | 2850 | 10 acel, 55 AD, 18 letalidade |

**Total:** 70 acel, 335 AD, 112 letalidade, 4 %mov

### maximo de life_steal_pct por ouro, ignorando efeitos de item

15100 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 1053 | Cetro Vampírico | 900 | 15 AD, 7 %roubo de vida |
| 3072 | Sedenta por Sangue | 3400 | 80 AD, 15 %roubo de vida |
| 3074 | Hidra Raivosa | 3300 | 15 acel, 65 AD, 12 %roubo de vida |
| 3139 | Cimitarra Mercurial | 3200 | 50 AD, 10 %roubo de vida, 35 MR |
| 3153 | Espada do Rei Destruído | 3200 | 40 AD, 25 %AS, 10 %roubo de vida |
| 3172 | Grevas Bélicas | 1100 | 45 %AS, 5 %roubo de vida, 45 mov |

**Total:** 15 acel, 250 AD, 70 %AS, 59 %roubo de vida, 35 MR, 45 mov

### maximo de magic_penetration por ouro, ignorando efeitos de item

7100 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3175 | Sapatos Enfeitiçados | 1100 | 20 pen magica, 8 %pen magica, 45 mov |
| 4645 | Chama Sombria | 3200 | 110 AP, 15 pen magica |
| 4646 | Ápice da Tempestade | 2800 | 90 AP, 15 pen magica, 6 %mov |

**Total:** 200 AP, 50 pen magica, 8 %pen magica, 45 mov, 6 %mov

### maximo de magic_penetration_pct por ouro, ignorando efeitos de item

4100 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3135 | Cajado do Vazio | 3000 | 95 AP, 40 %pen magica |
| 3175 | Sapatos Enfeitiçados | 1100 | 20 pen magica, 8 %pen magica, 45 mov |

**Total:** 95 AP, 20 pen magica, 48 %pen magica, 45 mov

### maximo de magic_resist por ouro, ignorando efeitos de item

14700 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 1057 | Capa Negatron | 850 | 45 MR |
| 2504 | Rookern Lamúrico | 2900 | 100 %regen vida, 400 vida, 80 MR |
| 3065 | Semblante Espiritual | 2700 | 10 acel, 100 %regen vida, 400 vida, 50 MR |
| 3091 | Limite da Razão | 2800 | 50 %AS, 45 MR, 20 %tenacidade |
| 4401 | Força da Natureza | 2800 | 400 vida, 55 MR, 4 %mov |
| 8020 | Máscara Abissal | 2650 | 15 acel, 350 vida, 45 MR |

**Total:** 25 acel, 50 %AS, 200 %regen vida, 1550 vida, 320 MR, 4 %mov, 20 %tenacidade

### maximo de mana por ouro, ignorando efeitos de item

16250 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2503 | Tocha de Chamas Negras | 2800 | 20 acel, 80 AP, 600 mana |
| 3003 | Cajado do Arcanjo | 2900 | 25 acel, 70 AP, 600 mana |
| 3110 | Coração Congelado | 2500 | 20 acel, 75 armadura, 400 mana |
| 3118 | Malevolência | 2700 | 15 acel, 90 AP, 600 mana |
| 6655 | Eco de Luden | 2750 | 10 acel, 100 AP, 600 mana |
| 6657 | Bastão das Eras | 2600 | 45 AP, 350 vida, 500 mana |

**Total:** 90 acel, 385 AP, 75 armadura, 350 vida, 3300 mana

### maximo de move_speed por ouro, ignorando efeitos de item

1000 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3170 | Marcha Célere | 1000 | 65 mov |

**Total:** 65 mov

### maximo de move_speed_pct por ouro, ignorando efeitos de item

12575 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 3046 | Dançarina Fantasma | 2650 | 65 %AS, 25 %crit, 10 %mov |
| 3066 | Couraça Lunar Alada | 800 | 200 vida, 4 %mov |
| 3085 | Furacão de Runaan | 2650 | 40 %AS, 25 %crit, 5 %mov |
| 3100 | Perdição de Lich | 2900 | 10 acel, 100 AP, 6 %mov |
| 4646 | Ápice da Tempestade | 2800 | 90 AP, 15 pen magica, 6 %mov |
| 6690 | Retriz | 775 | 15 AD, 4 %mov |

**Total:** 10 acel, 190 AP, 15 AD, 105 %AS, 50 %crit, 200 vida, 15 pen magica, 35 %mov

### maximo de omnivamp_pct por ouro, ignorando efeitos de item

7550 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 1055 | Lâmina de Doran | 450 | 10 AD, 80 vida, 2.5 %omnivamp |
| 2517 | Fome Eterna | 3100 | 65 AD, 5 %omnivamp, 20 %tenacidade |
| 3008 | Grevas Vorazes | 1000 | 45 mov, 4 %omnivamp |
| 3146 | Pistola Laminar Hextec | 3000 | 80 AP, 40 AD, 10 %omnivamp |

**Total:** 80 AP, 115 AD, 80 vida, 45 mov, 21.5 %omnivamp, 20 %tenacidade

### maximo de tenacity_pct por ouro, ignorando efeitos de item

10350 de ouro dos 20000 disponíveis, em 6 slots. Força adaptativa resolvida como ad.

| id | item | custo | stats |
|---|---|---:|---|
| 2517 | Fome Eterna | 3100 | 65 AD, 5 %omnivamp, 20 %tenacidade |
| 3053 | Sinal de Sterak | 3200 | 400 vida, 20 %tenacidade |
| 3091 | Limite da Razão | 2800 | 50 %AS, 45 MR, 20 %tenacidade |
| 3111 | Passos de Mercúrio | 1250 | 20 MR, 45 mov, 30 %tenacidade |

**Total:** 65 AD, 50 %AS, 400 vida, 65 MR, 45 mov, 5 %omnivamp, 90 %tenacidade

