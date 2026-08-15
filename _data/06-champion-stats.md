# Estatísticas de campeão — patch 16.16

Estatística base com crescimento por nível, e o efeito de cada habilidade
por rank. Vem do dump de dados do jogo, e não do arquivo do cliente — o
cliente publica esses campos zerados.

**O efeito é uma expressão, e não um número.** `50 + 0.8 AP` significa 50 de
dano fixo mais 80% do poder de habilidade. O conjunto publica os insumos; a
conta com a build montada é de quem consome.

Habilidade sem número na tabela é habilidade cuja fórmula a fonte não
resolve — **nunca é zero**. As que ficaram sem número estão listadas no fim.

## Annie

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 560 | 96 | 2192 |
| dano de ataque | 50 | 2.65 | 95.05 |
| armadura | 23 | 4 | 91 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.61 | 1.36 | 23.73 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 625 | — | 625 |
| multiplicador de crítico | 2 | — | 2 |

**P · Piromania**

- StunDuration: 1.75 _(no nível 18)_
- MaxStacks: 4
- StunBaseDuration: 1.25
- StunDurationPerTier: 0.25

**Q · Desintegrar** · recarga 4/4/4/4/4 · custo 60/65/70/75/80 · alcance 625/625/625/625/625

- TotalDamage: 80/125/170/215/260 + 0.8 ability_power

**W · Incinerar** · recarga 7/7/7/7/7 · custo 70/75/80/85/90 · alcance 600/600/600/600/600

- TotalDamage: 70/110/150/190/230 + 0.8 ability_power

**E · Escudo Fundido** · custo 60/65/70/75/80 · alcance 800/800/800/800/800

- DamageReturn: 25/35/45/55/65 + 0.4 ability_power
- MoveSpeedCalc: 0.5/0.5/0.5/0.5/0.5 _(no nível 18)_
- ShieldBlockTotal: 60/95/130/165/200 + 0.4 ability_power
- CastRangeGenerosity: 225/225/225/225/225
- DamageReduction: 13/17/21/25/29
- MovementSpeed: 0.15/0.2/0.25/0.3/0.35
- MovementSpeedDuration: 1.5/1.5/1.5/1.5/1.5
- ShieldDuration: 3/3/3/3/3

**R · Invocar: Tibbers** · recarga 130/115/100 · custo 100/100/100 · alcance 600/600/600

- InitialBurstDamage: 150/275/400 + 0.75 ability_power
- TibbersAADamage: 30/45/60 + 0.1 ability_power
- TibbersAuraDamage: 8/12/16 + 0.04 ability_power
- TibbersTotalHP: 1150/1150/1150 + 0.5 ability_power
- TibbersTotalResists: 30/30/30
- AvengerEnrageDuration: 10/10/10
- EnrageDuration: 3/3/3
- MovespeedBurst: 1/1/1
- RPercentPenBuff: 0.1/0.15/0.2
- TibbersAttackSpeedDecay: -0.1/-0.1/-0.1
- TibbersBonusMS: 0/25/50
- TibbersLifetime: 45/45/45

## Olaf

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 645 | 119 | 2668 |
| dano de ataque | 68 | 4.7 | 147.9 |
| armadura | 35 | 4.2 | 106.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.72 | 2.7 | 46.620003 |
| regeneração de vida | 1.3 | 0.12 | 3.34 |
| velocidade de movimento | 350 | — | 350 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Fúria Berserker** · recarga 0

- MaxAttackSpeed: 1 _(no nível 18)_
- MaxLifeSteal: 0.25 _(no nível 18)_
- AttackSpeedPerMissingHPPerc: 0.009
- MaxStatsThreshold: 0.3

**Q · Ressaca** · recarga 9/9/9/9/9 · custo 50/55/60/65/70 · alcance 1000/1000/1000/1000/1000

- TooltipCDRefund: 2.5/2.5/2.5/2.5/2.5
- TotalDamage: 70/120/170/220/270 + 1 attack_damage (bonus)
- DebuffDuration: 4/4/4/4/4
- MaxSlowDistance: 800/800/800/800/800
- MaxSlowDuration: 3/3/3/3/3
- MinSlowDistance: 400/400/400/400/400
- MinSlowDuration: 1/1/1/1/1
- MinimumRange: 425/425/425/425/425
- MonsterDamage: 20/45/70/95/120
- ShredAmount: 0.2/0.2/0.2/0.2/0.2
- SlowAmount: 0.3/0.35/0.4/0.45/0.5

**W · Duro na Queda** · recarga 16/15/14/13/12 · custo 50/50/50/50/50 · alcance 700/700/700/700/700

- MaxShieldCalc: 10/40/70/100/130 + 0.122499995 max_health
- Attackspeed: 0.4/0.5/0.6/0.7/0.8
- Duration: 5/5/5/5/5
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5

**E · Balanço Temerário** · recarga 11/10/9/8/7 · alcance 325/325/325/325/325

- HealthCostCalc (= TotalDamage × 1): 28/46/64/82/100 + 0.2 attack_damage
- TotalDamage: 70/115/160/205/250 + 0.5 attack_damage
- {fcf1daee}: 28/46/64/82/100 + 0.2 attack_damage
- ADRatio: 0.5/0.5/0.5/0.5/0.5
- Cast_Time_Attack_Speed_Cap: 125/125/125/125/125
- Cast_Time_Base: 0.25/0.25/0.25/0.25/0.25
- Cast_Time_Min: 0.175/0.175/0.175/0.175/0.175
- ChampionRefresh: 1/1/1/1/1
- MonsterRefresh: 2/2/2/2/2

**R · Ragnarok** · recarga 100/90/80 · custo 100/100/100

- AD: 10/20/30 + 0.25 attack_damage
- Duration: 3/3/3
- DurationExtension: 2.5/2.5/2.5
- Haste: 0.2/0.45/0.7
- HasteDuration: 1/1/1
- Resists: 10/15/20

## Galio

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 126 | 2742 |
| dano de ataque | 59 | 3.5 | 118.5 |
| armadura | 24 | 4.7 | 103.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 1.5 | 26.125 |
| regeneração de vida | 1.6 | 0.16 | 4.32 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Esmagada Colossal**

- PassiveCooldownModified: 5
- TotalDamage: 115 + 0.4 ability_power + 1 attack_damage + 0.6 magic_resist (bonus) _(no nível 18)_
- ChargeRatePerHit: 3

**Q · Ventos de Guerra** · recarga 11/10/9/8/7 · custo 60/65/70/75/80

- PercentSuperQDamageTT (= {85f4c0e9} × 1): 8/8/8/8/8 + 0.04 ability_power
- QMissileDamage: 70/105/140/175/210 + 0.7 ability_power
- SuperQMonsterMaxDamageTotal: 200/200/200/200/200
- {85f4c0e9}: 2/2/2/2/2 + 0.01 ability_power
- SuperQDuration: 2/2/2/2/2

**W · Escudo de Durand** · recarga 18/17/16/15/14 · custo 50/50/50/50/50 · alcance 275/275/275/275/275

- MagicDamageReduction: 0.25/0.3/0.35/0.4/0.45 + 0.0004 ability_power + 0.0008 magic_resist (bonus) + 0.0001 max_health (bonus)
- MaxTotalDamage (= MinTotalDamage × 3): 60/90/120/150/180 + 0.90000004 ability_power
- MinTotalDamage: 20/30/40/50/60 + 0.3 ability_power
- PassiveShieldOOCTimer: 8/8/8/8/8 _(no nível 18)_
- PhysicalDamageReduction (= MagicDamageReduction × 1): 0.125/0.15/0.175/0.2/0.225 + 0.0002 ability_power + 0.0004 magic_resist (bonus) + 0.00005 max_health (bonus)
- TotalPassiveShield: 0/0/0/0/0 + 0.075 max_health
- {7a7521c1} (= MagicDamageReduction × 100): 25/30.000002/35/40/45 + 0.04 ability_power + 0.08 magic_resist (bonus) + 0.01 max_health (bonus)
- {c9a870dd} (= MagicDamageReduction × 50): 12.5/15.000001/17.5/20/22.5 + 0.02 ability_power + 0.04 magic_resist (bonus) + 0.005 max_health (bonus)
- CCDurationMax: 1.5/1.5/1.5/1.5/1.5
- CCDurationMin: 0.5/0.5/0.5/0.5/0.5
- DRLingerDuration: 2/2/2/2/2
- MaxChargeDuration: 2/2/2/2/2
- MaximumWBaseDamage: 60/90/120/150/180
- SelfSlowPercent: 15/15/15/15/15

**E · Soco Justiceiro** · recarga 11/10/9/8/7 · custo 50/50/50/50/50

- PVEDamage (= TotalDamage × 0.8): 80/108/136/164/192 + 0.8 ability_power
- TotalDamage: 100/135/170/205/240 + 1 ability_power
- KnockupDuration: 0.75/0.75/0.75/0.75/0.75
- MinRange: 250/250/250/250/250

**R · Entrada Heroica** · recarga 180/160/140 · custo 100/100/100 · alcance 4000/4750/5500

- TotalDamage: 150/250/350 + 0.7 ability_power + 1 magic_resist (bonus)
- StunDurationOuter: 0.75/0.75/0.75
- TemporaryWShieldDuration: 5/5/5

## Twisted Fate

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 604 | 108 | 2440 |
| dano de ataque | 52 | 2.5 | 94.5 |
| armadura | 24 | 4.35 | 97.95 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.5 | 43.125 |
| regeneração de vida | 1.1 | 0.12 | 3.1399999 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Dados Viciados**

- GameModeInteger: 1
- GoldRatioMaxCherry: 0.08
- GoldRatioMinCherry: 0.02

**Q · Curingas** · recarga 6/5.75/5.5/5.25/5 · custo 60/70/80/90/100

- TotalDamage: 60/105/150/195/240 + 0.85 ability_power + 0.5 attack_damage (bonus)

**W · Escolha uma Carta** · recarga 6/6/6/6/6 · custo 50/55/60/65/70 · alcance 200/200/200/200/200

- ttAD: 1/1/1/1/1
- ttBlueAP: 1/1/1/1/1
- ttGoldAP: 0.5/0.5/0.5/0.5/0.5
- ttRedAP: 0.7/0.7/0.7/0.7/0.7

**E · Baralho Marcado** · recarga 0/0/0/0/0

- BonusDamage: 65/90/115/140/165 + 0.4 ability_power + 0.2 attack_damage (bonus)
- AttackSpeedBonus: 15/25/35/45/55
- TowerEffectiveness: 0.5/0.5/0.5/0.5/0.5

**R · Destino** · recarga 170/140/110 · custo 100/100/100 · alcance 5500/5500/5500

- RecastDuration: 6/8/10

## Xin Zhao

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 106 | 2422 |
| dano de ataque | 63 | 3 | 114 |
| armadura | 35 | 4.4 | 109.8 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.645 | 3.5 | 60.145 |
| regeneração de vida | 1.6 | 0.14 | 3.98 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Determinação**

- HealAPRatio: 0.7 _(no nível 18)_
- HealHPRatio: 0.049999997 _(no nível 18)_
- TotalDamage: 0 + 0.2 ability_power + 0.6 attack_damage _(no nível 18)_
- TotalHealing: 0 + 0.7 ability_power + 0.049999997 max_health _(no nível 18)_

**Q · Golpe de Três Garras** · recarga 7/6.5/6/5.5/5 · custo 30/30/30/30/30 · alcance 375/375/375/375/375

- BonusDamage: 15/30/45/60/75 + 0.4 attack_damage (bonus)
- KnockUpDuration: 0.75/0.75/0.75/0.75/0.75

**W · Vento Vira Relâmpago** · recarga 12/11/10/9/8 · custo 60/60/60/60/60 · alcance 1000/1000/1000/1000/1000

- MinionMod: 53.34/53.34/53.34/53.34/53.34 _(no nível 18)_
- SlashDamage: 30/40/50/60/70 + 0.3 attack_damage
- ThrustDamage: 50/85/120/155/190 + 0.65 ability_power + 0.9 attack_damage
- TotalSlowDuration: 1.5/1.5/1.5/1.5/1.5 + 0.005 ability_power
- DelayTime: 0.05/0.05/0.05/0.05/0.05
- HitNum: 4/4/4/4/4
- MarkDuration: 3/3/3/3/3

**E · Investida Audaciosa** · recarga 11/11/11/11/11 · custo 60/60/60/60/60 · alcance 650/650/650/650/650

- ChargeDamage: 50/75/100/125/150 + 1.2 ability_power
- APToASRatio: 0.001/0.001/0.001/0.001/0.001
- ASDuration: 5/5/5/5/5
- ASMod: 0.38/0.46/0.54/0.62/0.7
- AoERadius: 250/250/250/250/250
- PermanentASToASRatio: 5/5/5/5/5
- SlowAmount: 30/30/30/30/30
- SlowDuration: 0.5/0.5/0.5/0.5/0.5

**R · Guarda Crescente** · recarga 120/110/100 · custo 100/100/100 · alcance 500/500/500

- TotalDamage: 75/175/275 + 1.1 ability_power + 1 attack_damage (bonus)
- MarkDuration: 3/3/3
- MissileDefenseBaseDuration: 4/4/4
- MissileDefenseRadius: 450/450/450
- PercentCurrentHealthDamage: 0.15/0.15/0.15

## Urgot

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 655 | 102 | 2389 |
| dano de ataque | 63 | 4 | 131 |
| armadura | 36 | 5 | 121 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 3.75 | 64.375 |
| regeneração de vida | 1.5 | 0.14 | 3.88 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 350 | — | 350 |
| multiplicador de crítico | 2 | — | 2 |

**P · Chamas Ecoantes**

- ADDamage: 0 + 1 attack_damage _(no nível 18)_
- MonsterCap: 360 _(no nível 18)_
- PerLegCD: 2.5 _(no nível 18)_
- PercentHPRatio: 0.06 _(no nível 18)_
- LegWarningTime: 2.5

**Q · Carga Corrosiva** · recarga 10/9.5/9/8.5/8 · custo 70/70/70/70/70 · alcance 800/800/800/800/800

- TotalDamage: 25/70/115/160/205 + 0.7 attack_damage
- AoERadius: 210/210/210/210/210
- DelayOnMissileLaunch: 0.1/0.1/0.1/0.1/0.1
- ExplosionDelayPostMissile: 0.3/0.3/0.3/0.3/0.3
- SlowAmount: 0.45/0.5/0.55/0.6/0.65
- SlowDuration: 1.25/1.25/1.25/1.25/1.25
- WHalterDuration: 0.3/0.3/0.3/0.3/0.3

**W · Expurgar** · recarga 12/9/6/3/0 · custo 40/30/20/10/0 · alcance 450/450/450/450/450

- DamagePerShot: 12/12/12/12/12 + 0.2 attack_damage
- Duration: 4/4/4/4/25000
- MinionMinimumDamage: 50/50/50/50/50
- MoveSpeedDuration: 0.5/0.5/0.5/0.5/0.5
- MoveSpeedMod: 125/125/125/125/125
- OnHitDamageReduction: 0.5/0.5/0.5/0.5/0.5
- SlowResistance: 40/40/40/40/40
- WAttacksPerSecond: 3/3/3/3/3

**E · Desdém** · recarga 16/15.5/15/14.5/14 · custo 60/70/80/90/100

- EDamage: 90/120/150/180/210 + 1 attack_damage (bonus)
- ETotalShieldHealth: 55/75/95/115/135 + 1.35 attack_damage (bonus) + 0.135 max_health (bonus)
- BaseDashSpeed: 1200/1200/1200/1200/1200
- EShieldDuration: 4/4/4/4/4
- MinimumDashRange: 450/450/450/450/450
- StunDuration: 1.5/1.5/1.5/1.5/1.5
- ThrowDistance: 100/100/100/100/100
- WHalterDuration: 1.5/1.5/1.5/1.5/1.5

**R · Pior que a Morte** · recarga 100/85/70 · custo 100/100/100 · alcance 2500/2500/2500

- RCalculatedDamage: 100/225/350 + 0.5 attack_damage (bonus)
- RChannelTimeReel: 1/1/1
- RFearDuration: 1.5/1.5/1.5
- RFearRadius: 600/600/600
- RHealthThreshold: 25/25/25
- RMoveSpeedMod: 75/75/75
- RSlowDuration: 4/4/4
- RSlowMultiplier: 1/1/1

## LeBlanc

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 598 | 108 | 2434 |
| dano de ataque | 55 | 2.2 | 92.4 |
| armadura | 22 | 4.2 | 93.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 2.35 | 40.607998 |
| regeneração de vida | 1.5 | 0.11 | 3.37 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Imagem-Espelho** · recarga 60

- Cooldown: 60
- TooltipCooldown: 60

**Q · Sigilo de Malícia** · recarga 6/6/6/6/6 · custo 50/50/50/50/50 · alcance 700/700/700/700/700

- BonusMinionDamage: 10/10/10/10/10
- Damage: 65/90/115/140/165 + 0.4 ability_power
- MarkDamage: 65/90/115/140/165 + 0.4 ability_power
- CooldownRefund: 0.3/0.3/0.3/0.3/0.3
- ManaRefund: 1/1/1/1/1
- MarkDuration: 3.5/3.5/3.5/3.5/3.5

**W · Distorção** · recarga 15/13.75/12.5/11.25/10 · custo 60/70/80/90/100

- TotalDamage: 75/115/155/195/235 + 0.8 ability_power
- SnapbackDelay: 0.2/0.2/0.2/0.2/0.2
- SnapbackQueueWindow: 0.2/0.2/0.2/0.2/0.2
- SnapbackTimeAllowed: 4/4/4/4/4
- SpellMaxRange: 600/600/600/600/600

**E · Correntes Etéreas** · recarga 14/13.25/12.5/11.75/11 · custo 50/50/50/50/50

- DelayedDamage: 80/120/160/200/240 + 0.85 ability_power
- InitialDamage: 50/70/90/110/130 + 0.4 ability_power
- RootDuration: 1.5/1.5/1.5/1.5/1.5
- TetherDistance: 865/865/865/865/865
- TetherDuration: 1.5/1.5/1.5/1.5/1.5

**R · Mímica** · recarga 45/35/25

- RE1Damage: 70/150/230 + 0.4 ability_power
- RE2Damage: 140/300/460 + 0.85 ability_power
- RQ1Damage: 70/150/230 + 0.4 ability_power
- RQ2Damage: 140/300/460 + 0.8 ability_power
- RWDamage: 150/315/480 + 0.8 ability_power
- RQMarkDuration: 3.5/3.5/3.5

## Vladimir

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 110 | 2470 |
| dano de ataque | 55 | 3 | 106 |
| armadura | 24 | 4.5 | 100.5 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 2 | 34.658 |
| regeneração de vida | 1.4 | 0.12 | 3.4399998 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 450 | — | 450 |
| multiplicador de crítico | 2 | — | 2 |

**P · Pacto Vermelho**

- ApproximateAPBonusAvoidingRecursion: 0 + -0.053328 ability_power + 0.033329997 max_health (bonus)
- ApproximateHPBonusAvoidingRecursion: 0 + 1.696 ability_power (bonus) + -0.056527674 max_health (bonus)
- HPforAP: 30

**Q · Transfusão** · recarga 9/7.9/6.8/5.7/4.6 · alcance 600/600/600/600/600

- BaseDamageTooltip: 80/100/120/140/160 + 0.6 ability_power
- BaseHealTooltip: 20/25/30/35/40 + 0.35 ability_power
- EmpoweredDamageTooltip (= BaseDamageTooltip × 1.85): 148/185/222/259/296 + 1.11 ability_power
- EmpoweredHealPercent: 5/5/5/5/5 + 0.04 ability_power
- EmpoweredHealPercentTooltip (= EmpoweredHealPercent × 1): 5/5/5/5/5 + 0.04 ability_power
- EmpoweredHealTooltip (= {097cda4f} × 1): 200/200/200/200/200 _(no nível 18)_
- MovementSpeedOnQ2: 40/40/40/40/40 _(no nível 18)_
- {097cda4f}: 200/200/200/200/200 _(no nível 18)_
- DamagePercentAmp: 85/85/85/85/85
- Empowered___HealPer100AP_0_1_: 0.04/0.04/0.04/0.04/0.04
- FrenzyDuration: 2.5/2.5/2.5/2.5/2.5
- MinionHealPercent: 35/35/35/35/35

**W · Poça de Sangue** · recarga 28/25/22/19/16 · alcance 1050/1050/1050/1050/1050

- TotalDamage: 80/135/190/245/300 + 0.15 max_health (bonus)
- TotalHeal (= TotalDamage × 1): 24/40.5/57.000004/73.5/90 + 0.045 max_health (bonus)
- HasteBoost: 0.375/0.375/0.375/0.375/0.375
- HasteDuration: 1/1/1/1/1
- HealthCost: 0.15/0.15/0.15/0.15/0.15
- MinionHealingMod: 0.6/0.6/0.6/0.6/0.6
- MoveSpeedMod: -0.4/-0.4/-0.4/-0.4/-0.4

**E · Maré de Sangue** · recarga 13/11/9/7/5 · alcance 550/550/550/550/550

- ChargeHealthTooltip: 0/0/0/0/0 + 0.08 max_health
- MaxDamageTooltip: 60/90/120/150/180 + 0.8 ability_power + 0.06 max_health
- MinDamageTooltip: 30/45/60/75/90 + 0.35 ability_power + 0.015 max_health
- MaxChannelTime: 1.5/1.5/1.5/1.5/1.5
- MaxHealthCost: 8/8/8/8/8
- SlowPercent: 40/45/50/55/60
- TimetoRampMaxDamage: 1/1/1/1/1
- TotalMaxHPDamage: 6/6/6/6/6

**R · Hemopraga** · recarga 120/120/120 · alcance 625/625/625

- Damage: 150/250/350 + 0.7 ability_power
- SecondaryHealingTooltip (= Damage × 0.4): 60/100/140 + 0.28 ability_power
- DamageAmp: 10/10/10
- Duration: 4/4/4
- VampPercentAdditionalChamp: 40/40/40
- VampPercentFirstChamp: 100/100/100

## Fiddlesticks

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 106 | 2452 |
| dano de ataque | 55 | 2.65 | 100.05 |
| armadura | 34 | 4.7 | 113.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.11 | 36.495 |
| regeneração de vida | 1.1 | 0.12 | 3.1399999 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 480 | — | 480 |
| multiplicador de crítico | 2 | — | 2 |

**Q · Aterrorizar** · recarga 15/14.5/14/13.5/13 · custo 65/65/65/65/65 · alcance 575/575/575/575/575

- TotalPercentHealthDamage: 0.04/0.045/0.05/0.055/0.06 + 0.0003 ability_power
- TotalPercentHealthDamageFeared (= TotalPercentHealthDamage × 2): 0.08/0.09/0.1/0.11/0.12 + 0.0006 ability_power
- FearDuration: 1.2/1.4/1.6/1.8/2
- MinimumDamage: 40/60/80/100/120

**W · Colheita Farta** · recarga 10/9.5/9/8.5/8 · custo 60/65/70/75/80 · alcance 650/650/650/650/650

- DrainDamageCalc: 60/90/120/150/180 + 0.45 ability_power
- DrainDuration: 2/2/2/2/2
- DrainLeashRange: 725/725/725/725/725
- MinionDamageMod: 0.5/0.5/0.5/0.5/0.5
- MinionHealingMod: 15/15/15/15/15
- MonsterDamageMod: 1.35/1.35/1.35/1.35/1.35
- MonsterHealingMod: 0.45/0.45/0.45/0.45/0.45
- MonsterMaxDamage: 400/400/400/400/400
- PercentForTooltip: 12/14.5/17/19.5/22
- PercentMultiplier: 0.12/0.145/0.17/0.195/0.22
- TicksPerSecond: 4/4/4/4/4
- VampPercentage: 25/32.5/40/47.5/55

**E · Ceifar** · recarga 10/9/8/7/6 · custo 40/45/50/55/60 · alcance 850/850/850/850/850

- Damage: 70/105/140/175/210 + 0.5 ability_power
- CastRange: 850/850/850/850/850
- EmpoweredSlowAmount: -0.5/-0.55/-0.6/-0.65/-0.7
- SilenceDuration: 1.25/1.25/1.25/1.25/1.25
- SlowAmount: -0.3/-0.35/-0.4/-0.45/-0.5

**R · Tempestade de Corvos** · recarga 140/110/80 · custo 100/100/100

- APRatioFullDuration: 0/0/0 + 0.5 ability_power
- APRatioTick: 0/0/0 + 0.5 ability_power
- DamageDone: 150/250/350 + 0.5 ability_power
- TotalDamage (= {d0d3db7f} × 1): 750/1250/1750 + 2.5 ability_power
- {d0d3db7f}: 150/250/350 + 0.5 ability_power
- ChannelTime: 1.5/1.5/1.5

## Kayle

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 670 | 92 | 2234 |
| dano de ataque | 50 | 2.5 | 92.5 |
| armadura | 26 | 4.2 | 97.399994 |
| resistência mágica | 22 | 1.3 | 44.1 |
| velocidade de ataque | 0.625 | 1.5 | 26.125 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ascensão Divina** · recarga 0

- EnrageTotalASPerStack: 6
- PassiveWaveDamage: 20 + 0.25 ability_power + 0.1 attack_damage (bonus)
- EnrageDuration: 5
- EnrageMaxStacks: 5
- FinalAttackRange: 625
- LevelForPassiveRank0: 1
- LevelForPassiveRank1: 6
- LevelForPassiveRank2: 11
- LevelForPassiveRank3: 16
- MSTowardsEnemy: 0.1
- MSTowardsEnemyRadius: 2000
- UpgradedAttackRange: 525

**Q · Explosão Radiante** · recarga 12/11/10/9/8 · custo 60/70/80/90/100 · alcance 900/900/900/900/900

- TotalDamage: 60/90/120/150/180 + 0.5 ability_power + 0.6 attack_damage (bonus)
- CastDelay: 0.25/0.25/0.25/0.25/0.25
- ExplosionBackwardDist: 100/100/100/100/100
- ExplosionCenterAoERadius: 100/100/100/100/100
- ExplosionForwardBackWidth: 90/90/90/90/90
- ExplosionForwardDist: 400/400/400/400/400
- ExplosionForwardOffset: 100/100/100/100/100
- ExplosionLeftRightDist: 150/150/150/150/150
- ExplosionLeftRightWidth: 125/125/125/125/125
- ManaRefundPercent: 0.5/0.5/0.5/0.5/0.5
- ShredDuration: 4/4/4/4/4
- ShredPercent: 15/15/15/15/15
- SlowDuration: 2/2/2/2/2
- SlowPercent: 25/30/35/40/45

**W · Bênção Celestial** · recarga 15/15/15/15/15 · custo 70/75/80/85/90

- TotalHaste: 0.24/0.28/0.32/0.36/0.4 + 0.0008 ability_power
- TotalHeal: 55/80/105/130/155 + 0.25 ability_power
- HasteDuration: 2/2/2/2/2

**E · Lâmina de Fogo Estelar** · recarga 8/7.5/7/6.5/6 · alcance 550/550/550/550/550

- ActiveTotalExecuteDamage: 8/8.5/9/9.5/10 + 0.015 ability_power
- EPassiveTotalDamage: 15/20/25/30/35 + 0.2 ability_power + 0.1 attack_damage (bonus)
- {266027bf}: 0/0/0/0/0 + 1 attack_damage
- BuffDuration: 5/5/5/5/5
- ExplosionRadius: 350/350/350/350/350
- MaxExecuteVsMonsters: 400/400/400/400/400

**R · Sentença Divina** · recarga 160/120/80 · custo 100/50/0 · alcance 900/900/900

- TotalDamage: 200/300/400 + 0.7 ability_power + 1 attack_damage (bonus)
- AoERadius: 675/675/775
- InvulnDuration: 2.5/2.5/2.5

## Master Yi

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 105 | 2425 |
| dano de ataque | 65 | 2.5 | 107.5 |
| armadura | 33 | 4.5 | 109.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.679 | 2.5 | 43.179 |
| regeneração de vida | 1.5 | 0.13 | 3.71 |
| velocidade de movimento | 355 | — | 355 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ataque duplo**

- TotalDamage: 0 + 0.5 attack_damage
- AttackCount: 4
- StackDuration: 4

**Q · Ataque Alpha** · recarga 20/19.5/19/18.5/18 · custo 50/55/60/65/70 · alcance 600/600/600/600/600

- BasicAttackCDR: 1/1/1/1/1
- SingleCritTotalDamage: 35/70/105/140/175 + 1.225 attack_damage
- SingleTotalDamage (= TotalDamage × 1): 35/70/105/140/175 + 1.225 attack_damage
- SubesquentDamage: 20/40/60/80/100 + 0.7 attack_damage
- TotalDamage: 20/40/60/80/100 + 0.7 attack_damage
- BaseOnHitMultiplier: 0.75/0.75/0.75/0.75/0.75
- BonusMonsterDamage: 60/85/110/135/160

**W · Meditar** · custo 40/40/40/40/40 · alcance 20/20/20/20/20

- InitialDR: 0.7/0.7/0.7/0.7/0.70000005
- TotalHeal: 120/200/280/360/440 + 1 ability_power
- BaseManaCost: 40/40/40/40/40
- DRLinger: 0.5/0.5/0.5/0.5/0.5
- DamageReductionTowerMod: 0.5/0.5/0.5/0.5/0.5
- HealDuration: 4/4/4/4/4
- InitialExtraDRDuration: 0.5/0.5/0.5/0.5/0.5
- MaxMissingHealthPercent: 1/1/1/1/1
- PercentManaCostPerSecond: 0.06/0.06/0.06/0.06/0.06
- TickFrequency: 0.5/0.5/0.5/0.5/0.5

**E · Estilo Wuju** · recarga 14/14/14/14/14 · alcance 20/20/20/20/20

- TotalDamage: 20/25/30/35/40 + 0.35 attack_damage (bonus)
- Duration: 5/5/5/5/5
- LegacySwordVFXDuration: 3/3/3/3/3

**R · Highlander** · recarga 85/85/85 · custo 100/100/100 · alcance 1/1/1

- RASBonus: 25/45/65
- RCooldownRefund: 0.7/0.7/0.7
- RDuration: 7/7/7
- RKillAssistExtension: 7/7/7
- RMSBonus: 35/45/55

## Alistar

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 685 | 120 | 2725 |
| dano de ataque | 62 | 3.75 | 125.75 |
| armadura | 40 | 4.7 | 119.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2.125 | 36.75 |
| regeneração de vida | 1.7 | 0.17 | 4.59 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Urro Triunfante** · recarga 3 · alcance 950

- AllyHeal (= BaseHeal × 1): 0 + 0.07 max_health
- BaseHeal: 0 + 0.05 max_health
- PassiveAllyNumHeal: 15
- PassiveCooldown: 3
- PassiveMaxStacks: 7
- PassiveStacksChampionKill: 7

**Q · Pulverizar** · recarga 14/13/12/11/10 · custo 50/55/60/65/70 · alcance 330/330/330/330/330

- TotalDamage: 60/100/140/180/220 + 0.8 ability_power
- AoERadius: 375/375/375/375/375
- KnockupDuration: 1/1/1/1/1

**W · Cabeçada** · recarga 14/13/12/11/10 · custo 50/55/60/65/70 · alcance 650/650/650/650/650

- TotalDamage: 55/110/165/220/275 + 1 ability_power
- KnockBackDistance: 700/700/700/700/700
- KnockUpDuration: 0.75/0.75/0.75/0.75/0.75
- StunDuration: 0.75/0.75/0.75/0.75/0.75

**E · Atropelar** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70 · alcance 350/350/350/350/350

- AttackBonusDamage: 20/20/20/20/20
- TotalDamage: 80/110/140/170/200 + 0.7 ability_power
- BonusAARange: 50/50/50/50/50
- Duration: 5/5/5/5/5
- FullChargeDuration: 5/5/5/5/5
- MaxStacks: 5/5/5/5/5
- ProcDamageBase: 20/20/20/20/20
- ProcDamageScale: 15/15/15/15/15
- Radius: 350/350/350/350/350
- StunDuration: 1/1/1/1/1

**R · Vontade Indestrutível** · recarga 120/100/80 · custo 100/100/100 · alcance 1/1/1

- RDamageReduction: 55/65/75
- RDuration: 7/7/7

## Ryze

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 124 | 2728 |
| dano de ataque | 55 | 3 | 106 |
| armadura | 22 | 4.2 | 93.399994 |
| resistência mágica | 32 | 1.3 | 54.1 |
| velocidade de ataque | 0.658 | 2.11 | 36.528 |
| regeneração de vida | 1.6 | 0.16 | 4.32 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Maestria Arcana** · recarga 0 · alcance 625

- PassiveManaCalcTooltip: 0 + 10 ability_power
- {92bc0080}: 0 + 0.05 ability_power
- APAmount: 100

**Q · Sobrecarregar** · recarga 5/5/5/5/5 · custo 40/38/36/34/32 · alcance 1000/1000/1000/1000/1000

- QDamageCalc: 75/95/115/135/155 + 0.55 ability_power + 0.02 mana
- DamageAmp: 40/50/60/70/80
- MaximumRunes: 2/2/2/2/2
- MovementSpeedAmount: 28/32/36/40/44
- MovementSpeedDuration: 2/2/2/2/2
- RuneDuration: 4/4/4/4/4

**W · Prisão de Runa** · recarga 11/10.5/10/9.5/9 · custo 50/60/70/80/90 · alcance 550/550/550/550/550

- WDamageCalc: 60/90/120/150/180 + 0.6 ability_power + 0.03 mana
- CCDuration: 1.5/1.5/1.5/1.5/1.5
- SlowAmount: 0.5/0.5/0.5/0.5/0.5

**E · Fluxo de Feitiço** · recarga 3.5/3.25/3/2.75/2.5 · custo 35/45/55/65/75 · alcance 550/550/550/550/550

- EDamageCalc: 60/90/120/150/180 + 0.5 ability_power + 0.02 mana
- BounceRadius: 350/350/350/350/350
- BounceRadiusLargeUnit: 400/400/400/400/400
- BounceRadiusQLargeUnit: 500/500/500/500/500
- BounceRadiusQSmallUnit: 350/350/350/350/350
- DamageAmp: 10/20/30/40/50
- DebuffDuration: 4/4/4/4/4

**R · Portal de Reinos** · recarga 180/160/140 · custo 100/100/100 · alcance 3000/3000/3000

- AllyIndicatorBuffDuration: 0.75/0.75/0.75
- ChargeTime: 2.1/2.1/2.1
- ChargeTimeTooltip: 2/2/2
- HealPercentMinionMod: 0.33/0.33/0.33
- MinimumCastRange: 1000/1000/1000
- OverloadDamageBonus: 50/75/100
- OverloadHealPercent: 20/25/30
- TeleportTime: 0.65/0.65/0.65

## Sion

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 655 | 87 | 2134 |
| dano de ataque | 68 | 4 | 136 |
| armadura | 36 | 4.2 | 107.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.679 | 1.3 | 22.779 |
| regeneração de vida | 1.8 | 0.16 | 4.52 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · In Gloria Mori**

- Lifesteal: 1
- NonChampCap: 75
- PercentMaxHP: 0.1
- StructureMod: 0.4

**Q · Golpe Demolidor** · recarga 10/9/8/7/6 · custo 45/45/45/45/45

- MaxDamageTotal: 90/155/220/285/350 + 1.2 attack_damage
- MinDamageTotal: 30/45/60/75/90 + 0.4 attack_damage
- BaseStunTime: 1.25/1.25/1.25/1.25/1.25
- MinionRatio: 60/60/60/60/60
- MonsterRatio: 165/165/165/165/165
- SlowAmount: -0.8/-0.8/-0.8/-0.8/-0.8

**W · Fornalha da Alma** · recarga 15/14/13/12/11 · custo 75/80/85/90/95 · alcance 500/500/500/500/500

- TotalDamage: 40/65/90/115/140 + 0.4 ability_power
- TotalShield: 60/75/90/105/120 + 0.4 ability_power + 0.08 max_health
- DamagePercentHealthTooltip: 0.1/0.11/0.12/0.13/0.14
- DetonateRecastCooldown: 3/3/3/3/3
- HPPerChampKill: 15/15/15/15/15
- HPPerKill: 4/4/4/4/4
- HPPerLargeKill: 15/15/15/15/15
- MaxHPDamageRatio: 14/14/14/14/14
- ShieldDuration: 6/6/6/6/6

**E · Urro do Assassino** · recarga 12/11/10/9/8 · custo 35/40/45/50/55

- TotalDamage: 65/100/135/170/205 + 0.55 ability_power
- ArmorShred: 25/25/25/25/25
- ArmorShredDuration: 4/4/4/4/4
- SlowAmount: 40/45/50/55/60
- SlowDuration: 2.5/2.5/2.5/2.5/2.5

**R · Investida Incontrolável** · recarga 140/100/60 · custo 100/100/100

- MaxDamageTotal: 400/800/1200 + 1.2 attack_damage (bonus)
- MinDamageTotal: 150/300/450 + 0.6 attack_damage (bonus)
- MaxStunDuration: 1.75/1.75/1.75
- MinStunDuration: 0.75/0.75/0.75
- MinionDamagePercent: 500/500/500
- MoveSpeedCap: 950/950/950
- SlowAmount: 40/45/50

## Sivir

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 104 | 2368 |
| dano de ataque | 60 | 2.5 | 102.5 |
| armadura | 30 | 4 | 98 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.625 | 1.6 | 27.825 |
| regeneração de vida | 0.65 | 0.11 | 2.52 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Pés Ligeiros**

- FlatMS: 75 _(no nível 18)_
- HasteDuration: 1.5

**Q · Lâmina Bumerangue** · recarga 10/9.5/9/8.5/8 · custo 55/60/65/70/75

- CritScaling: 0/0/0/0/0 + 0.4 critical_chance
- PreCritDamage: 60/85/110/135/160 + 0.6 ability_power + 0.7 attack_damage (bonus)
- TotalDamage: 60/85/110/135/160 + 0.6 ability_power + 0.7 attack_damage (bonus)
- FallOffMinimum: 0.4/0.4/0.4/0.4/0.4
- FallOffRatio: 0.15/0.15/0.15/0.15/0.15
- QAttackSpeedCastReductionMax: 0.6/0.6/0.6/0.6/0.6
- QAttackSpeedCastReductionPercent: 0.5/0.5/0.5/0.5/0.5
- QBaseCastTime: 0.25/0.25/0.25/0.25/0.25

**W · Ricochete** · recarga 12/12/12/12/12 · custo 60/60/60/60/60 · alcance 20/20/20/20/20

- BounceDamage: 0/0/0/0/0 + 0.4 attack_damage
- FirstTargetDamage: 0/0/0/0/0 + 1 attack_damage
- TotalMaxDamage: 0/0/0/0/0 + 4.2 attack_damage
- BounceAttacks: 3/3/3/3/3
- BuffDuration: 4/4/4/4/4
- MinionDamageMod: 0.65/0.65/0.65/0.65/0.65
- RicochetAttackSpeed: 0.2/0.25/0.3/0.35/0.4

**E · Escudo de Feitiço** · recarga 24/22.5/21/19.5/18 · alcance 20/20/20/20/20

- TotalHeal: 0/0/0/0/0 + 0.5 ability_power + 0.6 attack_damage
- SpellShieldDuration: 1.5/1.5/1.5/1.5/1.5

**R · Na Caçada** · recarga 120/100/80 · custo 100/100/100 · alcance 1000/1000/1000

- AttackCooldownRefund: 0.5/0.5/0.5
- AuraRange: 1000/1000/1000
- BuffExtension: 8/10/12
- DamagedMarkerDuration: 3/3/3
- HuntAttackSpeed: 0.05/0.06/0.07
- MaxMS: 0.2/0.25/0.3
- MaxMSDuration: 2/3/4
- MinMS: 0.3/0.35/0.4
- UltDuration: 8/10/12

## Soraka

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 605 | 88 | 2101 |
| dano de ataque | 50 | 3 | 101 |
| armadura | 32 | 5 | 117 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.14 | 37.005 |
| regeneração de vida | 0.5 | 0.1 | 2.2 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Salvação** · alcance 2500

- HealthThreshold: 0.4
- MovementSpeed: 0.9

**Q · Chamado Estelar** · recarga 8/7/6/5/4 · custo 45/50/55/60/65 · alcance 800/800/800/800/800

- TotalDamage: 85/120/155/190/225 + 0.35 ability_power
- TotalHot: 60/75/90/105/120 + 0.3 ability_power
- HotDuration: 2.5/2.5/2.5/2.5/2.5
- MoveSpeedDuration: 2.5/2.5/2.5/2.5/2.5
- MoveSpeedHaste: 0.2/0.225/0.25/0.275/0.3
- MoveSpeedSlow: 0.3/0.3/0.3/0.3/0.3
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**W · Infusão Astral** · recarga 6/5/4/3/2 · custo 40/45/50/55/60 · alcance 550/550/550/550/550

- MinimumHealth: 0/0/0/0/0 + 0.05 max_health
- TotalHeal: 90/110/130/150/170 + 0.5 ability_power
- PercentHealthCost: 0.1/0.1/0.1/0.1/0.1
- PercentHealthCostRefund: 0.8/0.85/0.9/0.95/1

**E · Equinócio** · recarga 20/19/18/17/16 · custo 70/75/80/85/90 · alcance 875/875/875/875/875

- TotalDamage: 70/95/120/145/170 + 0.4 ability_power
- RootDelay: 1.5/1.5/1.5/1.5/1.5
- RootDuration: 1/1.25/1.5/1.75/2

**R · Desejo** · recarga 150/135/120 · custo 100/100/100

- AmpedHealing (= HealingCalc × 1): 225/375/525 + 0.75 ability_power
- HealingCalc: 150/250/350 + 0.5 ability_power

## Teemo

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 615 | 104 | 2383 |
| dano de ataque | 54 | 3 | 105 |
| armadura | 24 | 4.5 | 100.5 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.69 | 3.38 | 58.15 |
| regeneração de vida | 1.1 | 0.13 | 3.31 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Técnicas de Guerrilha**

- BonusAttackSpeed: 0.8 _(no nível 18)_
- AttackSpeedDuration: 5
- StealthCooldownDuration: 1.5

**Q · Dardo Ofuscante** · recarga 7/7/7/7/7 · custo 70/75/80/85/90 · alcance 680/680/680/680/680

- CalculatedDamage: 80/125/170/215/260 + 0.7 ability_power
- {0f87e4ba}: 2/2.25/2.5/2.75/3
- MissChance: 1/1/1/1/1

**W · Mover Depressa** · recarga 14/14/14/14/14 · custo 40/40/40/40/40 · alcance 20/20/20/20/20

- ActiveMoveSpeedBonus: 0.24/0.32/0.4/0.48/0.56
- ActiveMoveSpeedBuffDuration: 3/3/3/3/3
- PassiveCooldownOnDamageTaken: 5/5/5/5/5
- PassiveMoveSpeedBonus: 0.12/0.16/0.2/0.24/0.28

**E · Tiro Tóxico** · recarga 0/0/0/0/0 · alcance 680/680/680/680/680

- ImpactCalculatedDamage: 9/23/37/51/65 + 0.3 ability_power + 0.05 attack_damage (bonus)
- TickCalculatedDamage: 6/12/18/24/30 + 0.1 ability_power + 0.025 attack_damage (bonus)
- TotalDotDamage (= TickCalculatedDamage × 1): 24/48/72/96/120 + 0.4 ability_power + 0.1 attack_damage (bonus)
- {8b70cfd6} (= ImpactCalculatedDamage × 1): 14.400001/36.8/59.2/81.6/104 + 0.48000002 ability_power + 0.080000006 attack_damage (bonus)
- {938cba49} (= TickCalculatedDamage × 1): 9.6/19.2/28.800001/38.4/48 + 0.16000001 ability_power + 0.040000003 attack_damage (bonus)
- TickFrequency: 1/1/1/1/1

**R · Armadilha Venenosa** · recarga 0.25/0.25/0.25 · custo 75/55/35 · alcance 600/750/900

- TotalDamage: 200/325/450 + 0.5 ability_power
- ArmTime: 1/1/1
- DebuffDuration: 4/4/4
- ExplosionRadius: 450/450/450
- MaxAmmo: 3/4/5
- MaxBounceDistance: 360/440/550
- MaxTraps: 25000/25000/25000
- MinTossDistanceBounceThreshold: 220/500/500
- MushroomDuration: 5/5/5
- SlowAmount: 30/40/50
- TriggerRadius: 160/160/160
- VisionRadius: 210/210/210

## Tristana

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 102 | 2374 |
| dano de ataque | 60 | 3.4 | 117.8 |
| armadura | 30 | 4 | 98 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.656 | 1.5 | 26.156 |
| regeneração de vida | 0.8 | 0.1 | 2.5 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Tiro Certeiro** · recarga 0

- BonusPassiveRange: 150 _(no nível 18)_

**Q · Tiro Rápido** · recarga 20/19/18/17/16 · custo 15/20/25/30/35 · alcance 20/20/20/20/20

- AttackSpeedMod: 0.6/0.75/0.9/1.05/1.2
- BuffDuration: 7/7/7/7/7

**W · Salto-foguete** · recarga 22/20/18/16/14 · custo 30/35/40/45/50

- LandingDamage: 70/105/140/175/210 + 0.5 ability_power + 1 attack_damage (bonus)
- DamageRadius: 350/350/350/350/350
- SlowDuration: 2/2/2/2/2
- SlowMod: -0.4/-0.4/-0.4/-0.4/-0.4

**E · Carga Explosiva** · recarga 16/15.5/15/14.5/14 · custo 50/55/60/65/70 · alcance 550/550/550/550/550

- ActiveDamage: 60/85/110/135/160 + 0.5 ability_power + 0.8 attack_damage (bonus)
- ActiveMaxDamage (= ActiveDamage × 1): 120/170/220/270/320 + 1 ability_power + 1.6 attack_damage (bonus)
- PassiveDamage: 45/60/75/90/105 + 0.25 ability_power
- ActiveDuration: 4/4/4/4/4
- ActiveRadius: 300/300/300/300/300
- ActiveTowerRadius: 600/600/600/600/600
- PassiveRadius: 300/300/300/300/300

**R · Tiro Destruidor** · recarga 100/100/100 · custo 100/100/100 · alcance 550/550/550

- DamageCalc: 225/275/325 + 1 ability_power + 0.7 attack_damage (bonus)
- KnockbackArea: 200/200/200
- KnockbackDistance: 600/800/1000
- StunDuration: 0.4/0.55/0.7

## Warwick

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 99 | 2303 |
| dano de ataque | 65 | 2.5 | 107.5 |
| armadura | 33 | 4.4 | 107.8 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.638 | 2 | 34.638 |
| regeneração de vida | 0.8 | 0.15 | 3.3500001 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Fome Eterna**

- OnHitDamage: 55 + 0.1 ability_power + 0.15 attack_damage (bonus) _(no nível 18)_
- EmpoweredHealingRatio: 2.5
- EmpoweredHealingThreshold: 0.25
- HealingRatio: 1
- HealingThreshold: 0.5
- HealthRestoreSlushTime: 1.55

**Q · Presas da Fera** · recarga 8/7.5/7/6.5/6 · custo 80/85/90/95/100 · alcance 365/365/365/365/365

- BaseBiteDamage: 0/0/0/0/0 + 1 ability_power + 1.2 attack_damage
- APRatio: 1/1/1/1/1
- BiteZoneLength: 450/450/450/450/450
- CheatLength: 200/200/200/200/200
- CheatLengthMoveBlock: 300/300/300/300/300
- HPDamageCap: 150/165/180/195/210
- LifestealPercent: 25/37.5/50/62.5/75
- QCastRangeOverride: 425/425/425/425/425
- TargetPercentHPDamage: 6/7/8/9/10

**W · Caçada Sangrenta** · recarga 80/70/60/50/40 · custo 55/55/55/55/55 · alcance 3000/3000/3000/3000/3000

- ASBonus: 70/80/90/100/110
- Duration: 8/8/8/8/8
- FirstHPThreshold: 0.5/0.5/0.5/0.5/0.5
- MSBonus: 35/42.5/50/57.5/65
- MoveSpeedBonus: 15/20/25/30/35
- NoTargetCDReduction: 0.7/0.7/0.7/0.7/0.7
- PassiveASBonus: 70/80/90/100/110
- PassiveASDuration: 1.25/1.25/1.25/1.25/1.25
- PassiveMSBonus: 35/42.5/50/57.5/65
- SecondHPThreshold: 0.25/0.25/0.25/0.25/0.25

**E · Uivo Primitivo** · recarga 15/14/13/12/11 · custo 40/40/40/40/40 · alcance 375/375/375/375/375

- DRAmount: 35/40/45/50/55
- DRDuration: 2.75/2.75/2.75/2.75/2.75
- DRDurationTooltipOnly: 2.5/2.5/2.5/2.5/2.5
- FearDuration: 1/1/1/1/1
- RecastDelay: 1/1/1/1/1

**R · Coerção Infinita** · recarga 110/90/70 · custo 100/100/100

- DamageCumulative: 175/350/525 + 1.67 attack_damage (bonus)
- RDuration: 1.5/1.5/1.5

## Nunu e Willump

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 90 | 2140 |
| dano de ataque | 61 | 3 | 112 |
| armadura | 29 | 4.2 | 100.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2.25 | 38.875 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Chamado de Freljord** · recarga 20 · custo 40 · alcance 20

- CleaveDamage: 0 + 0.3 attack_damage
- ASIncrease: 0.2
- BloodBoilBaseDuration: 4
- LargeRadius: 1000
- MSIncrease: 0.1
- MonsterRadius: 700

**Q · Consumir** · recarga 13/12/11/10/9 · custo 60/60/60/60/60 · alcance 125/125/125/125/125

- ChampionHealing (= MonsterHealing × 1): 39/57.000004/75/93/111.00001 + 0.54 ability_power + 0.060000002 max_health (bonus)
- MonsterHealing: 65/95/125/155/185 + 0.9 ability_power + 0.1 max_health (bonus)
- TotalChampionDamage: 60/100/140/180/220 + 0.65 ability_power + 0.05 max_health (bonus)
- LowHealthHealingScalar: 0.5/0.5/0.5/0.5/0.5
- LowHealthThreshhold: 0.5/0.5/0.5/0.5/0.5
- MonsterMinionDamage: 400/600/800/1000/1200

**W · A Maior Bola de Neve de Todas!** · recarga 14/14/14/14/14 · custo 50/55/60/65/70

- MaximumSnowballDamage: 180/225/270/315/360 + 1.5 ability_power
- MaximumStunDuration: 1.25/1.25/1.25/1.25/1.25
- NoImpactSnowballDamage (= MaximumSnowballDamage × 1): 59.940002/74.925/89.91/104.895004/119.880005 + 0.4995 ability_power
- AdditionalKnockupOverTime: 0.25/0.25/0.25/0.25/0.25
- AdditionalSpeedCapPerLevel: 10/10/10/10/10
- AdditionalSpeedPerLevel: 10/10/10/10/10
- AdditionalStartingSpeed: 75/75/75/75/75
- InitialCastRadiusScalar: 2.5/2.5/2.5/2.5/2.5
- LargeSnowballMissileWidth: 200/200/200/200/200
- LargeSnowballTime: 5/5/5/5/5
- MaxDamageScalar: 1/1/1/1/1
- MaxDamageTime: 5/5/5/5/5
- MaxDuration: 10/10/10/10/10
- MaximumSnowballRadius: 200/200/200/200/200
- MaximumSnowballRollDistance: 1750/1750/1750/1750/1750
- MediumSnowballMissileWidth: 137.5/137.5/137.5/137.5/137.5
- MediumSnowballTime: 2.5/2.5/2.5/2.5/2.5
- MinimumSnowballMissileSpeed: 350/350/350/350/350
- MinimumSnowballRadius: 75/75/75/75/75
- MinimumSnowballRollDistance: 750/750/750/750/750
- MonsterCollisionRadiusScalar: 1.25/1.25/1.25/1.25/1.25
- MonsterImpactRadiusScalar: 1.75/1.75/1.75/1.75/1.75
- SlowAmount: -0.5/-0.5/-0.5/-0.5/-0.5
- SlowDuration: 1/1/1/1/1
- SmallSnowballMissileWidth: 75/75/75/75/75
- TurnRateMultTime1: 1/1/1/1/1
- TurnRateMultTime2: 2/2/2/2/2
- TurnRateMultTime3: 3/3/3/3/3
- TurnRateMultiplier1: 1.75/1.75/1.75/1.75/1.75
- TurnRateMultiplier2: 2.5/2.5/2.5/2.5/2.5
- TurnRateMultiplier3: 3.25/3.25/3.25/3.25/3.25
- WalkDistanceAfterThrow: 350/350/350/350/350

**E · Rajada de Bolas de Neve** · recarga 14/13/12/11/10 · custo 50/55/60/65/70 · alcance 5000/5000/5000/5000/5000

- RootDuration: 1.5/1.5/1.5/1.5/1.5 _(no nível 18)_
- TotalRootDamage: 20/30/40/50/60 + 0.8 ability_power
- TotalSnowballDamage: 15/22.5/30/37.5/45 + 0.12 ability_power
- DelayBetweenSnowballs: 0.15/0.15/0.15/0.15/0.15
- SlowAmount: -0.3/-0.35/-0.4/-0.45/-0.5
- SlowDuration: 1/1/1/1/1
- SplashConeAngle: 22.5/22.5/22.5/22.5/22.5
- SplashConeLength: 225/225/225/225/225
- TotalSpellDuration: 3/3/3/3/3

**R · Zero Absoluto** · recarga 110/100/90 · custo 100/100/100 · alcance 650/650/650

- MaximumDamage: 625/925/1275 + 3 ability_power
- MinDamage (= MaximumDamage × 0.5): 312.5/462.5/637.5 + 1.5 ability_power
- TotalShieldAmount: 65/75/85 + 1.5 ability_power + 0.3 max_health (bonus)
- ChannelDuration: 3/3/3
- MaxShieldDuration: 6/6/6
- MaxSlowAmount: -0.95/-0.95/-0.95
- ShieldDecayDuration: 3/3/3
- SlowDuration: 3/3/3
- SlowStartAmount: -0.5/-0.5/-0.5

## Miss Fortune

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 625 | 100 | 2325 |
| dano de ataque | 55 | 2.4 | 95.8 |
| armadura | 25 | 4 | 93 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.656 | 3 | 51.656 |
| regeneração de vida | 0.75 | 0.13 | 2.96 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Batida do Amor** · recarga 0

- MinionDamage (= TotalDamage × 0.5): 0 + 0.5 attack_damage _(no nível 18)_
- TotalDamage: 0 + 1 attack_damage _(no nível 18)_

**Q · Dois por Um** · recarga 7/6/5/4/3 · custo 40/40/40/40/40 · alcance 550/550/550/550/550

- TotalDamageTooltip: 20/45/70/95/120 + 0.35 ability_power + 1 attack_damage
- {5a1fd353}: 0/0/0/0/0 + 1 critical_damage

**W · Desfilando** · recarga 12/12/12/12/12 · custo 45/45/45/45/45 · alcance 600/600/600/600/600

- LoveTapRefund: 2/2/2/2/2
- ActiveAS: 0.4/0.55/0.7/0.85/1
- ActiveDuration: 4/4/4/4/4
- PassiveBaseMS: 30/35/40/45/50
- PassiveBaseMSOOC: 4/4/4/4/4
- PassiveMaxMS: 60/70/80/90/100
- PassiveMaxMSExtraOOC: 3/3/3/3/3

**E · Chuva de Disparos** · recarga 18/17/16/15/14 · custo 80/80/80/80/80 · alcance 1000/1000/1000/1000/1000

- TotalDamage (= TotalDamagePerSecond × 1): 70/100/130/160/190 + 1.2 ability_power
- TotalDamagePerSecond: 35/50/65/80/95 + 0.6 ability_power
- TotalSlowAmount: 0.4/0.4/0.4/0.4/0.4 + 0.0006 ability_power
- CastDelay: 0.25/0.25/0.25/0.25/0.25
- TicksPerSecond: 4/4/4/4/4

**R · Metendo Bala** · recarga 120/110/100 · custo 100/100/100

- PhysicalDamagePerWave: 20/30/40 + 0.25 ability_power + 0.6 attack_damage
- TotalPhysicalDamage (= PhysicalDamagePerWave × 1): 280/480/720 + 3.5 ability_power + 8.400001 attack_damage
- BaseChannelDuration: 3/3/3

## Ashe

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 101 | 2327 |
| dano de ataque | 59 | 3.5 | 118.5 |
| armadura | 26 | 4.6 | 104.2 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.658 | 3 | 51.658 |
| regeneração de vida | 0.7 | 0.11 | 2.57 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 600 | — | 600 |
| multiplicador de crítico | 1 | — | 1 |

**P · Tiro Congelado**

- EmpoweredSlowAmount: 0.6 _(no nível 18)_
- SlowAmount: 0.3 _(no nível 18)_
- SlowDuration: 2

**Q · Concentração** · recarga 0/0/0/0/0 · custo 30/30/30/30/30

- EmpoweredDamage: 0/0/0/0/0 + 1.1 attack_damage
- BonusAS: 20/30/40/50/60
- BuffDuration: 6/6/6/6/6
- MaxStacks: 4/4/4/4/4
- ShotsPerStrike: 5/5/5/5/5
- StackDuration: 4/4/4/4/4
- StackFalloffDuration: 1/1/1/1/1
- TimerDuration: 1/1/1/1/1

**W · Rajada** · recarga 18/14.5/11/7.5/4 · custo 75/70/65/60/55 · alcance 600/600/600/600/600

- TotalDamage: 60/95/130/165/200 + 1 attack_damage (bonus)
- NumberOfArrowsTooltip: 7/8/9/10/11

**E · Olhar do Falcão** · recarga 5/5/5/5/5

- ChargeCooldown: 90/80/70/60/50
- VisionDuration: 5/5/5/5/5

**R · Flecha de Cristal Encantada** · recarga 100/80/60 · custo 100/100/100

- RMainDamage: 200/400/600 + 1.2 ability_power
- {6b932875} (= RMainDamage × 1): 200/400/600 + 1.2 ability_power
- MaxStunDuration: 3.5/3.5/3.5
- MinStunDuration: 1/1/1
- SlowRadius: 400/400/400

## Tryndamere

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 696 | 108 | 2532 |
| dano de ataque | 66 | 4.5 | 142.5 |
| armadura | 33 | 4.8 | 114.600006 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.67 | 3.4 | 58.47 |
| regeneração de vida | 1.7 | 0.18 | 4.76 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Fúria da Batalha**

- PassiveCritConversionTooltip: 0.5
- {dc63b617}: 0.05
- CherryFuryModifier: 1.5

**Q · Sanguinário** · recarga 12/12/12/12/12

- BaseHeal: 30/40/50/60/70 + 0.3 ability_power
- HealPerFury: 0.5/0.95/1.4/1.85/2.3 + 0.012 ability_power
- MaximumBonusAD: 20/35/50/65/80
- RemainingHealthThreshold: 0.1/0.1/0.1/0.1/0.1

**W · Grito Zombador** · recarga 14/14/14/14/14

- ADReduction: -20/-35/-50/-65/-80
- ReductionDuration: 4/4/4/4/4
- SlowDuration: 3.25/3.25/3.25/3.25/3.25
- SlowPotency: -0.3/-0.35/-0.4/-0.45/-0.5

**E · Corte Giratório** · recarga 12/11/10/9/8

- TotalDamage: 80/120/160/200/240 + 0.8 ability_power + 1 attack_damage (bonus)
- ChampCDRefund: 1.5/1.5/1.5/1.5/1.5
- ChampFuryGain: 5/5/5/5/5
- DamageAoE: 225/225/225/225/225
- NonChampCDRefund: 0.75/0.75/0.75/0.75/0.75
- NonChampFuryGain: 2/2/2/2/2

**R · Fúria Sem Fim** · recarga 120/100/80

- TryndRDuration: 5/5/5
- TryndRFuryGain: 50/75/100
- TryndRMinHealth: 30/50/70

## Jax

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 103 | 2401 |
| dano de ataque | 68 | 4.25 | 140.25 |
| armadura | 36 | 4.2 | 107.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.638 | 3.4 | 58.438004 |
| regeneração de vida | 1.7 | 0.11 | 3.57 |
| velocidade de movimento | 350 | — | 350 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Investida Implacável**

- AttackSpeedPerStack: 0.125 _(no nível 18)_
- BelvethVsJaxQuestRewardAttackSpeedPerStack: 0.0028
- MaxBonusAttackSpeed (= AttackSpeedPerStack × 1): 1 _(no nível 18)_
- BelvethVsJaxQuestRewardStacks: 5
- BuffDuration: 2.5
- FallOffRate: 0.35

**Q · Salto Atacante** · recarga 8/7.5/7/6.5/6 · custo 50/50/50/50/50 · alcance 700/700/700/700/700

- TotalDamage: 65/105/145/185/225 + 1 attack_damage (bonus)

**W · Energizar** · recarga 7/6/5/4/3 · custo 30/30/30/30/30 · alcance 300/300/300/300/300

- FinalDamage: 50/85/120/155/190 + 0.6 ability_power
- TotalDamage: 50/85/120/155/190 + 0.6 ability_power
- StructureMod: 0.5/0.5/0.5/0.5/0.5

**E · Contra-Ataque** · recarga 17/15/13/11/9 · custo 50/60/70/80/90 · alcance 20/20/20/20/20

- MaxDamage (= TotalDamage × 2): 80/140/200/260/320 + 1.4 ability_power (bonus)
- MaxPercentHealthDamage: 4/4/4/4/4
- MonsterDamageCap: 9000/9000/9000/9000/9000
- TotalDamage: 40/70/100/130/160 + 0.7 ability_power (bonus)
- AoEDamageReduction: 25/25/25/25/25
- DodgeDuration: 2/2/2/2/2
- MaxDodgesForDamageIncrease: 5/5/5/5/5
- PercentIncreasedPerDodge: 0.2/0.2/0.2/0.2/0.2
- StunDuration: 1/1/1/1/1

**R · Grão-Mestre de Armas** · recarga 110/100/90 · custo 100/100/100

- BaseArmor: 45/60/75 + 0.4 attack_damage (bonus)
- BaseMR (= BaseArmor × 1): 27.000002/36/45 + 0.24000001 attack_damage (bonus)
- BonusArmor: 20/25/30 + 0.1 attack_damage (bonus)
- BonusMR (= BonusArmor × 1): 12/15.000001/18 + 0.060000002 attack_damage (bonus)
- FinalDamage: 75/130/185 + 0.6 ability_power
- OnHitDamage: 75/130/185 + 0.6 ability_power
- SwingDamageTotal: 100/175/250 + 1 ability_power (bonus)
- AoESize: 375/375/375
- Duration: 8/8/8
- PassiveFallOffTime: 2.5/2.5/2.5
- StructureMod: 0.5/0.5/0.5

## Morgana

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 104 | 2398 |
| dano de ataque | 56 | 3.5 | 115.5 |
| armadura | 25 | 4.2 | 96.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 1.53 | 26.635 |
| regeneração de vida | 1.1 | 0.08 | 2.46 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 450 | — | 450 |
| multiplicador de crítico | 2 | — | 2 |

**P · Sifão da Alma** · recarga 0

- HealPercent: 18

**Q · Ligação das Trevas** · custo 50/55/60/65/70 · alcance 1300/1300/1300/1300/1300

- TotalDamage: 80/135/190/245/300 + 0.9 ability_power
- RootDuration: 2/2.25/2.5/2.75/3

**W · Sombra Atormentada** · recarga 12/12/12/12/12 · custo 70/80/90/100/110 · alcance 900/900/900/900/900

- TotalMaxDamage (= TotalMinDamage × 1): 36/62/88/114/140 + 0.4 ability_power
- TotalMinDamage: 18/31/44/57/70 + 0.2 ability_power
- CDRefundPercent: 0.05/0.05/0.05/0.05/0.05
- MonsterMod: 170/170/170/170/170
- TickRate: 0.5/0.5/0.5/0.5/0.5
- WDuration: 5/5/5/5/5

**E · Escudo Negro** · recarga 26/23.5/21/18.5/16 · custo 80/80/80/80/80 · alcance 800/800/800/800/800

- TotalShieldStrength: 100/155/210/265/320 + 0.7 ability_power
- ShieldDuration: 5/5/5/5/5

**R · Grilhões da Alma** · recarga 120/110/100 · custo 100/100/100 · alcance 625/625/625

- TotalDamage: 200/275/350 + 0.8 ability_power
- ChainDuration: 3/3/3
- HastePercent: 20/40/60
- SlowPercent: 20/20/20
- StunDuration: 1.5/1.75/2

## Zilean

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 574 | 96 | 2206 |
| dano de ataque | 52 | 3 | 103 |
| armadura | 24 | 5 | 109 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 2.13 | 36.868004 |
| regeneração de vida | 1.1 | 0.1 | 2.8 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Tempo Engarrafado** · alcance 825

- XPPer5 (= {e359ec0d} × 2.5): 6 _(no nível 18)_
- {e359ec0d}: 2.4 _(no nível 18)_
- GamemodeInteger: 1
- PassiveCooldown: 120

**Q · Bomba-relógio** · recarga 10/9.5/9/8.5/8 · custo 60/65/70/75/80 · alcance 900/900/900/900/900

- TotalDamage: 75/115/165/230/300 + 0.9 ability_power
- FuseDuration: 3/3/3/3/3
- StunDuration: 1.1/1.2/1.3/1.4/1.5

**W · Retroceder** · recarga 14/12/10/8/6 · custo 35/35/35/35/35 · alcance 600/600/600/600/600

- CooldownReduction: 10/10/10/10/10
- ManaCost: 35/35/35/35/35

**E · Distorção no Tempo** · recarga 15/15/15/15/15 · custo 50/50/50/50/50 · alcance 550/550/550/550/550

- Duration: 2.5/2.5/2.5/2.5/2.5
- SpeedAmount: 40/55/70/85/99

**R · Alteração Temporal** · recarga 120/90/60 · custo 125/150/175 · alcance 900/900/900

- RTotalHeal: 600/850/1100 + 2 ability_power
- RDuration: 5/5/5
- ReviveStateDuration: 3/3/3

## Singed

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 96 | 2282 |
| dano de ataque | 63 | 3.4 | 120.8 |
| armadura | 34 | 4.2 | 105.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.7 | 1.9 | 33 |
| regeneração de vida | 1.9 | 0.11 | 3.77 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Corrente de Ar Nociva** · alcance 225

- MSDuration: 2
- MSPercent: 0.25
- PerTargetCD: 8
- TriggerArea: 225

**Q · Rastro de Veneno** · recarga 0/0/0/0/0 · custo 13/13/13/13/13 · alcance 20/20/20/20/20

- ApproximateTotalDamageTooltip (= DamagePerSecond × 4.75): 95/142.5/190/237.5/285 + 2.01875 ability_power
- DamagePerSecond: 20/30/40/50/60 + 0.425 ability_power
- CloudDuration: 3.25/3.25/3.25/3.25/3.25
- MaxLingerTicks: 8/8/8/8/8
- PoisonDuration: 2/2/2/2/2
- TicksPerSecond: 4/4/4/4/4
- ToggleCooldown: 1/1/1/1/1

**W · Mega Adesivo** · recarga 17/16/15/14/13 · custo 60/70/80/90/100 · alcance 1000/1000/1000/1000/1000

- DelayExecute: 0.375/0.375/0.375/0.375/0.375
- SlowPercent: 50/55/60/65/70
- WDuration: 3/3/3/3/3
- WRadius: 265/265/265/265/265

**E · Lançar** · recarga 10/9.5/9/8.5/8 · custo 60/70/80/90/100 · alcance 125/125/125/125/125

- BaseDamage: 50/60/70/80/90 + 0.55 ability_power
- FlingDistance: 420/420/420/420/420
- MaxHPDamage: 6/6.5/7/7.5/8
- NonChampionDamageCap: 300/300/300/300/300
- RootDuration: 1/1.25/1.5/1.75/2

**R · Poção da Insanidade** · recarga 100/100/100 · custo 100/100/100 · alcance 20/20/20

- Duration: 25/25/25
- GrievousAmount: 0.4/0.4/0.4
- GrievousDuration: 1/1/1
- StatAmount: 25/55/85

## Evelynn

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 642 | 98 | 2308 |
| dano de ataque | 61 | 3 | 112 |
| armadura | 37 | 4.7 | 116.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.667 | 2.1 | 36.366997 |
| regeneração de vida | 1.7 | 0.15 | 4.25 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Sombra Demoníaca**

- HealPerSecondTOOLTIP: 150 _(no nível 18)_
- HealingThresholdTOOLTIP: 250 + 2.5 ability_power
- DemonShadeTimer: 4
- StealthDropTimer: 1.5

**Q · Espinho de Ódio** · recarga 4/4/4/4/4 · custo 40/45/50/55/60 · alcance 800/800/800/800/800

- MissileDamage: 25/30/35/40/45 + 0.25 ability_power
- TotalBonusDamage: 15/25/35/45/55 + 0.25 ability_power
- CooldownRefund: 0.5/0.5/0.5/0.5/0.5
- FakeCD: 4/4/4/4/4
- MissileBaseDamage: 30/35/40/45/50
- QStackCount: 3/3/3/3/3

**W · Fascinação** · recarga 15/14/13/12/11 · custo 60/70/80/90/100 · alcance 1200/1300/1400/1500/1600

- MonsterDamageTotalTOOLTIP: 250/300/350/400/450 + 0.6 ability_power
- ActualCD: 15/14/13/12/11
- BuffDuration: 5/5/5/5/5
- CharmDuration: 1.25/1.5/1.75/2/2.25
- MRShred: 0.35/0.375/0.4/0.425/0.45
- MonsterCharm: 3/3.25/3.5/3.75/4
- MonsterExtraCharm: 2/2/2/2/2
- NonCharmSlowDuration: 1.5/1.5/1.5/1.5/1.5
- ShredDuration: 4/4/4/4/4
- SlowAmount: 0.45/0.45/0.45/0.45/0.45
- SlowDuration: 0.75/0.75/0.75/0.75/0.75
- SlowDuringCharm: -0.45/-0.45/-0.45/-0.45/-0.45

**E · Chicotada** · recarga 8/8/8/8/8 · custo 40/45/50/55/60 · alcance 210/210/210/210/210

- PercentHealthBaseTOOLTIP: 3/3/3/3/3 + 0.015 ability_power
- PercentHealthEmpoweredTOOLTIP: 4/4/4/4/4 + 0.025 ability_power
- BaseDamage: 60/90/120/150/180
- EmpoweredDamage: 80/120/160/200/240
- MSDuration: 2/2/2/2/2
- MonsterCap: 450/450/450/450/450
- MonsterDamageCamp: 450/450/450/450/450
- MoveSpeed: 0.3/0.35/0.4/0.45/0.5
- SpeedAmount: 0.3/0.35/0.4/0.45/0.5
- SpeedDuration: 2/2/2/2/2

**R · Última Carícia** · recarga 120/100/80 · custo 100/100/100

- CritDamage (= Damage × 1): 300/600/900.00006 + 1.8000001 ability_power
- Damage: 125/250/375 + 0.75 ability_power
- CritBonus: 1.4/1.4/1.4
- CritTreshold: 0.3/0.3/0.3
- PassiveReset: 1.25/1.25/1.25

## Twitch

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 98 | 2296 |
| dano de ataque | 59 | 3 | 110 |
| armadura | 27 | 4 | 95 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.679 | 3 | 51.679 |
| regeneração de vida | 0.75 | 0.12 | 2.79 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Veneno Mortal**

- DamageMaxTotal (= DamagePerSecond × 1): 180 + 1.0799999 ability_power _(no nível 18)_
- DamagePerSecond: 5 + 0.03 ability_power _(no nível 18)_
- DamagePerSecondMax (= DamagePerSecond × 1): 30 + 0.17999999 ability_power _(no nível 18)_

**Q · Emboscada** · recarga 16/16/16/16/16 · custo 40/40/40/40/40 · alcance 20/20/20/20/20

- AttackSpeedDuration: 6/6/6/6/6
- AttackSpeedMod: 0.4/0.45/0.5/0.55/0.6
- HiddenSpeed: 30/30/30/30/30
- MaxFadeTime: 1/1/1/1/1
- MoveBonusRange: 1000/1000/1000/1000/1000
- MoveSpeedMod: 10/10/10/10/10
- StealthDetectionRange: 500/500/500/500/500
- StealthDuration: 10/11/12/13/14

**W · Tonel de Veneno** · recarga 13/12/11/10/9 · custo 70/70/70/70/70 · alcance 950/950/950/950/950

- TotalSlowAmount: 30/35/40/45/50 + 0.06 ability_power
- AoERadius: 300/300/300/300/300
- Duration: 3/3/3/3/3

**E · Contaminar** · recarga 12/11/10/9/8 · custo 50/60/70/80/90 · alcance 1200/1200/1200/1200/1200

- MagicDamagePerStack: 0/0/0/0/0 + 0.35 ability_power
- MaxMagicDamage (= MagicDamagePerStack × 1): 0/0/0/0/0 + 2.1 ability_power
- MaxPhysicalDamage: 110/150/190/230/270 + 2.1 attack_damage (bonus)
- PhysicalDamagePerStack: 15/20/25/30/35 + 0.35 attack_damage (bonus)

**R · Passando Fogo** · recarga 90/90/90 · custo 100/100/100 · alcance 1200/1200/1200

- BonusAD: 30/45/60
- BonusRange: 300/300/300
- Duration: 6/6/6
- FallOffDamage: 0.1/0.1/0.1
- MinimumFallOffDamage: 0.6/0.6/0.6
- OvershootMin: 250/250/250
- OvershootMult: 0.3/0.3/0.3

## Karthus

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 110 | 2490 |
| dano de ataque | 46 | 3.25 | 101.25 |
| armadura | 21 | 4.7 | 100.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.11 | 36.495 |
| regeneração de vida | 1.3 | 0.11 | 3.1699998 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 450 | — | 450 |
| multiplicador de crítico | 2 | — | 2 |

**P · Desafio da Morte**

- PassiveDuration: 7

**Q · Devastar** · recarga 0/0/0/0/0 · custo 20/25/30/35/40 · alcance 875/875/875/875/875

- QDamage: 40/59/78/97/116 + 0.35 ability_power
- QSingleTargetDamage (= QDamage × 2): 80/118/156/194/232 + 0.7 ability_power
- MonsterMod: 1/1/1/1/1

**W · Barreira da Dor** · recarga 15/15/15/15/15 · custo 70/70/70/70/70 · alcance 1000/1000/1000/1000/1000

- DebuffDuration: 5/5/5/5/5
- MagicResistShred: 25/25/25/25/25
- SlowPercent: 40/50/60/70/80
- TT_WallWidth: 800/900/1000/1100/1200
- WallDuration: 5/5/5/5/5

**E · Perverter** · recarga 0.5/0.5/0.5/0.5/0.5 · custo 30/42/54/66/78 · alcance 550/550/550/550/550

- TotalDPS: 30/50/70/90/110 + 0.2 ability_power
- {57456bbc} (= TotalDPS × 0.25): 7.5/12.5/17.5/22.5/27.5 + 0.05 ability_power
- ManaRestoreOnKill: 10/20/30/40/50

**R · Réquiem** · recarga 200/180/160 · custo 100/100/100

- TotalDamage: 200/350/500 + 0.7 ability_power

## Cho'Gath

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 644 | 94 | 2242 |
| dano de ataque | 69 | 4.2 | 140.4 |
| armadura | 38 | 4.5 | 114.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 1.44 | 25.138 |
| regeneração de vida | 1.8 | 0.17 | 4.69 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Carnívoro**

- ChogathCarnivoreHeal: 18
- ChogathCarnivoreMana: 4.72

**Q · Ruptura** · recarga 6/6/6/6/6 · custo 50/50/50/50/50 · alcance 950/950/950/950/950

- TotalDamageTooltip: 80/135/190/245/300 + 1 ability_power
- SlowAmountPercentage: 60/60/60/60/60

**E · Espinhos Vorpais** · recarga 8/7/6/5/4 · custo 30/30/30/30/30 · alcance 40/40/40/40/40

- FlatDamageCalc: 20/40/60/80/100 + 0.3 ability_power
- AttackRangeIncrease: 50/50/50/50/50
- BuffDuration: 6/6/6/6/6
- MaximumAttacks: 3/3/3/3/3
- ModifiedMonsterCap: 80/110/140/170/200
- MonsterDamageCap: 200/200/200/200/200
- SlowAmountPercentage: 30/35/40/45/50
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**R · Banquete** · recarga 80/70/60 · custo 100/100/100 · alcance 175/175/175

- RDamage: 300/475/650 + 0.5 ability_power + 0.1 max_health (bonus)
- RMonsterDamage: 1200/1200/1200 + 0.5 ability_power + 0.1 max_health (bonus)
- MaxBonusAttackRange: 75/75/75
- MaxBonusCastRange: 25/25/25
- RHealthPerStack: 80/120/160
- RMinionMaxStacks: 6/6/6
- RStoneplateRatio: 1/1/1

## Amumu

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 685 | 94 | 2283 |
| dano de ataque | 57 | 3.8 | 121.6 |
| armadura | 33 | 4 | 101 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.736 | 2.18 | 37.796 |
| regeneração de vida | 1.8 | 0.17 | 4.69 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Toque Amaldiçoado** · recarga 0

- DamageAmp: 0.1
- DebuffDuration: 3

**Q · Lançar Bandagens** · recarga 3/3/3/3/3 · custo 50/50/50/50/50 · alcance 1100/1100/1100/1100/1100

- TotalDamage: 70/95/120/145/170 + 0.85 ability_power
- DashSpeed: 1800/1800/1800/1800/1800
- StunDuration: 1/1/1/1/1

**W · Desespero** · recarga 1/1/1/1/1 · custo 8/8/8/8/8 · alcance 300/300/300/300/300

- TotalHealthDamage: 1/1.25/1.5/1.75/2 + 0.005 ability_power
- {8a96509c} (= TotalHealthDamage × 0.005): 0.005/0.0062499996/0.0075/0.00875/0.01 + 0.000025 ability_power
- {c8e45bc3}: 10/10/10/10/10

**E · Chilique** · recarga 9/8/7/6/5 · custo 35/35/35/35/35 · alcance 350/350/350/350/350

- DamageReduction: 5/7/9/11/13 + 0.03 armor (bonus) + 0.03 magic_resist (bonus)
- TantrumDamage: 65/95/125/155/185 + 0.5 ability_power
- CDROnHit: 0.75/0.75/0.75/0.75/0.75
- FlatDamageReductionMax: 0.5/0.5/0.5/0.5/0.5
- PassiveScaling: 0.03/0.03/0.03/0.03/0.03

**R · A Maldição da Múmia Triste** · recarga 150/125/100 · custo 100/150/200 · alcance 550/550/550

- RCalculatedDamage: 200/300/400 + 0.8 ability_power
- RDuration: 1.5/1.5/1.5

## Rammus

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 645 | 100 | 2345 |
| dano de ataque | 65 | 2.75 | 111.75 |
| armadura | 35 | 4.5 | 111.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.7 | 2.215 | 38.355 |
| regeneração de vida | 1.6 | 0.11 | 3.47 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Casco Espetado**

- TotalDamage: 0 + 0.15 armor + 0.15 magic_resist
- BaseDamage: 10

**Q · Bola do Poder** · recarga 12/10.5/9/7.5/6 · custo 60/60/60/60/60 · alcance 300/300/300/300/300

- MaximumMoveSpeed (= MinimumMoveSpeed × 1): 2.346/2.346/2.346/2.346/2.346 _(no nível 18)_
- MinimumMoveSpeed: 0.391/0.391/0.391/0.391/0.391 _(no nível 18)_
- PowerBallDamage: 80/120/160/200/240 + 1 ability_power
- HeightVariable: 25/25/25/25/25
- KnockbackDistance: 125/125/125/125/125
- SelfStunTiming: 0.25/0.25/0.25/0.25/0.25
- SlowDuration: 1/1/1/1/1
- SlowPercent: 40/50/60/70/80
- StunDuration: 0.35/0.35/0.35/0.35/0.35

**W · Bola Curva Defensiva** · recarga 7/7/7/7/7 · custo 40/40/40/40/40 · alcance 300/300/300/300/300

- BonusArmorTooltip: 35.1/44/53.649998/64.049995/75.200005 + 0.3 armor
- BonusMRTooltip: 26/34.375/43.5/53.375/64 + 0.3 magic_resist
- RecastDamageTooltip: 15/15/15/15/15 + 0.1 armor + 0.1 magic_resist
- ReturnDamageCalc: 15/15/15/15/15 + 0.1 armor + 0.1 magic_resist
- BuffDuration: 7/7/7/7/7

**E · Provocação Enlouquecedora** · recarga 12/12/12/12/12 · custo 50/50/50/50/50 · alcance 325/325/325/325/325

- MonsterDamageCalc: 80/100/120/140/160 + 0.7 ability_power
- Duration: 1.2/1.4/1.6/1.8/2

**R · Colisão Ascendente** · recarga 120/105/90 · custo 100/100/100

- InitialDamageCalc: 150/250/350 + 0.6 ability_power
- MaxSlow: 30.000002/40/50
- PulseDamageCalc: 0/0/0 + 0.1 ability_power
- TooltipMaxDamageCalc: 150/250/350 + 0.6 ability_power
- BaseCastRange: 800/800/800
- BuffDuration: 3.5/3.5/3.5
- DashRangeGrowth: 1.5/1.5/1.5
- DashSpeedGrowth: 1.1/1.1/1.1
- KnockupDuration: 0.75/0.75/0.75
- MaxDashSpeed: 2000/2000/2000
- MaxRangeForDmgGrowth: 1700/1700/1700
- MinDashSpeed: 900/900/900
- SlowDuration: 1.5/1.5/1.5
- TremorsAoERange: 400/400/400
- TremorsKnockupRange: 200/200/200
- TurretDamageModifier: 2/2/2

## Anivia

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 550 | 92 | 2114 |
| dano de ataque | 51 | 3.2 | 105.4 |
| armadura | 19 | 4.1 | 88.7 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 1.68 | 29.217999 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 600 | — | 600 |
| multiplicador de crítico | 2 | — | 2 |

**P · Renascimento** · recarga 0

- BonusResists: 20 _(no nível 18)_
- BonusResistsTooltip (= BonusResists × -1): -20 _(no nível 18)_
- Cooldown: 240

**Q · Lampejo Gelado** · recarga 11/10/9/8/7 · custo 80/85/90/95/100

- TotalExplosionDamage: 60/95/130/165/200 + 0.45 ability_power
- TotalPassthroughDamage: 50/70/90/110/130 + 0.25 ability_power
- SlowDuration: 3/3/3/3/3
- StunDuration: 1.1/1.2/1.3/1.4/1.5

**W · Cristalizar** · recarga 17/17/17/17/17 · custo 70/70/70/70/70 · alcance 1000/1000/1000/1000/1000

- ChampPushDistance: 120/120/120/120/120
- NonChampPushDistance: 250/250/250/250/250
- WallChunks: 4/5/6/7/8
- WallDuration: 5/5/5/5/5
- WallWidth: 400/500/600/700/800

**E · Congelamento** · recarga 4/4/4/4/4 · custo 50/50/50/50/50 · alcance 600/600/600/600/600

- EmpoweredDamage (= TotalDamage × 2): 110/160/210/260/310 + 1.1 ability_power
- TotalDamage: 55/80/105/130/155 + 0.55 ability_power

**R · Tempestade Glacial** · recarga 4/3/2 · custo 60/60/60 · alcance 750/750/750

- EmpoweredDamagePerSecondTooltipOnly (= TotalDamagePerSecond × 3): 90/135/180 + 0.375 ability_power
- EnhancedSlow: 20/30/40
- TotalDamagePerSecond: 30/45/60 + 0.125 ability_power
- BonusMultiplier: 300/300/300
- ChillDuration: 1/1/1
- GrowthTime: 1.5/1.5/1.5
- LeashBreak: 1000/1000/1000
- LeashWarning: 1000/1000/1000
- ManaCostPerSecond: 35/45/55
- MinCooldown: 1/1/1
- SlowDurationAtMaxMultiplier: 1.5/1.5/1.5
- SlowPercentEmpoweredTT: 30/45/60
- TickRate: 0.5/0.5/0.5

## Shaco

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 99 | 2313 |
| dano de ataque | 63 | 3 | 114 |
| armadura | 30 | 4 | 98 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.694 | 3 | 51.694 |
| regeneração de vida | 1.7 | 0.11 | 3.57 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Apunhalar** · recarga 0

- BasicAttackDamage: 35 + 0.2 attack_damage (bonus) _(no nível 18)_
- ShivDamage: 50 + 0.1 ability_power _(no nível 18)_
- ShivDamageExecute (= ShivDamage × 1): 75 + 0.15 ability_power _(no nível 18)_
- CloneMult: 0.75
- PerUnitCD: 3
- ShivSlowAmount: -0.4
- ShredDuration: 4

**Q · Enganar** · recarga 13/12.5/12/11.5/11 · custo 40/40/40/40/40

- QCritDamageMod: 0.39999998/0.39999998/0.39999998/0.39999998/0.39999998 + 0.6 critical_damage
- TotalDamage: 25/35/45/55/65 + 0.6 attack_damage (bonus)
- CDRefund: 2.5/2.5/2.5/2.5/2.5
- ExtraAATime: 0.25/0.25/0.25/0.25/0.25
- PseudoCastTime: 0.125/0.125/0.125/0.125/0.125
- StealthDuration: 2.5/2.75/3/3.25/3.5

**W · Caixinha-Surpresa** · recarga 15/15/15/15/15 · custo 70/65/60/55/50 · alcance 500/500/500/500/500

- AoEDamage: 10/15/20/25/30 + 0.12 ability_power
- STDamage: 25/40/55/70/85 + 0.18 ability_power
- TrapDuration: 40/40/40/40/40 + 0.1 ability_power
- ArmTime: 2/2/2/2/2
- FearDuration: 0.5/0.75/1/1.25/1.5
- MaxTraps: 25000/25000/25000/25000/25000
- MinionFearDuration: 2/2/2/2/2
- MonsterBonusDamage: 20/35/50/65/80

**E · Veneno de Dois Gumes** · recarga 8/8/8/8/8 · custo 75/75/75/75/75 · alcance 625/625/625/625/625

- TotalDamage: 70/95/120/145/170 + 0.6 ability_power + 0.8 attack_damage (bonus)
- TotalExecuteDamage (= TotalDamage × 1): 105/142.5/180/217.5/255 + 0.90000004 ability_power + 1.2 attack_damage (bonus)
- ExecuteHealthThreshold: 0.3/0.3/0.3/0.3/0.3
- SlowAmount: -0.2/-0.225/-0.25/-0.275/-0.3
- SlowDurationActive: 3/3/3/3/3
- SlowDurationPassive: 2/2/2/2/2

**R · Alucinações** · recarga 100/90/80 · custo 100/100/100 · alcance 200/200/200

- AoEDamage: 10/20/30 + 0.1 ability_power
- ExplosionTotalDamage: 150/225/300 + 0.7 ability_power
- STDamage: 25/50/75 + 0.15 ability_power
- BoxArmTime: 2/2/2
- BoxFearDuration: 1/1/1
- BoxLifetime: 6/6/6
- CloneAADamagePercent: 0.6/0.6/0.6
- CloneIncomingDamagePercent: 0.5/0.5/0.5
- CloneLifetime: 18/18/18
- TeleportRange: 150/150/150

## Dr. Mundo

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 103 | 2391 |
| dano de ataque | 61 | 2.5 | 103.5 |
| armadura | 32 | 4.5 | 108.5 |
| resistência mágica | 29 | 2.3 | 68.1 |
| velocidade de ataque | 0.67 | 3.3 | 56.77 |
| regeneração de vida | 1.4 | 0.1 | 3.1 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Vai Para Onde Quer** · recarga 0

- MaxHealthRegen: 0.004
- PassiveCooldown: 15 _(no nível 18)_
- CannisterDistanceAway: 525
- CannisterGroundDuration: 7
- CannisterMaxAngle: 70
- CannisterPickupRadius: 115
- CurrentHealthLoss: 0.04
- MaxHealthGain: 0.04
- PassiveCooldownRefund: 15
- VFXLineIndicatorRange: 1000

**Q · Serra Infectada** · recarga 4/4/4/4/4 · alcance 1050/1050/1050/1050/1050

- HealthRestoreOnHitChampionMonster: 50/60/70/80/90
- HealthRestoreOnHitMinion: 25/30/35/40/45
- CurrentHealthDamage: 0.2/0.225/0.25/0.275/0.3
- MaximumMonsterDamage: 250/325/400/475/550
- MinimumDamage: 80/130/180/230/280
- SlowAmount: 0.4/0.4/0.4/0.4/0.4
- SlowDuration: 2/2/2/2/2

**W · Choquinho Cardíaco** · recarga 17/16.5/16/15.5/15 · alcance 325/325/325/325/325

- GrayHealthStorageInitial: 0.95/0.95/0.95/0.95/0.95 _(no nível 18)_
- TotalDamage: 20/35/50/65/80 + 0.07 max_health (bonus)
- CurrentHealthCost: 0.08/0.08/0.08/0.08/0.08
- DamagePerTick: 5/8.75/12.5/16.25/20
- Duration: 3/3/3/3/3
- GrayHealthBigMod: 1/1/1/1/1
- GrayHealthInitialDuration: 0.75/0.75/0.75/0.75/0.75
- GrayHealthSmallMod: 0.5/0.5/0.5/0.5/0.5
- GrayHealthStorage: 0.25/0.25/0.25/0.25/0.25
- SecondCastLockout: 0.5/0.5/0.5/0.5/0.5
- SecondCastRange: 325/325/325/325/325

**E · Traumatismo** · recarga 9/8.25/7.5/6.75/6

- AdditionalDamage: 5/15/25/35/45 + 0.05 max_health (bonus)
- MaxDamageAmpTooltip: 0.39999998/0.39999998/0.39999998/0.39999998/0.39999998
- PassiveBonusAD: 0/0/0/0/0 + 2 max_health
- AttackOverrideDuration: 4/4/4/4/4
- FlatHealthCost: 10/25/40/55/70
- MaxMissingHealthThreshold: 0.7/0.7/0.7/0.7/0.7
- MinionMod: 1.4/1.4/1.4/1.4/1.4
- MissileDistance: 800/800/800/800/800
- MonsterMod: 1.4/1.4/1.4/1.4/1.4

**R · Dosagem Máxima** · recarga 120/120/120 · alcance 20/20/20

- BonusPerNearbyChampion: 0.05/0.05/0.05
- Duration: 10/10/10
- MaxHealthHoT: 0.2/0.4/0.6
- MissingHealthHeal: 0.15/0.2/0.25
- SpeedBoostAmount: 0.15/0.25/0.35
- TakedownDurationExtension: 2/2/2

## Sona

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 550 | 91 | 2097 |
| dano de ataque | 49 | 3 | 100 |
| armadura | 26 | 4.2 | 97.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.644 | 2.3 | 39.744 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Power Chord**

- PowerChordDamage: 20 + 0.2 ability_power
- AccelerandoCap: 60
- AccelerandoUltCDR: 1.5
- PowerChordPassiveCountMax: 3
- TempoDuration: 2

**Q · Hino do Valor** · recarga 8/8/8/8/8 · custo 50/55/60/65/70 · alcance 825/825/825/825/825

- TotalDamage: 50/85/120/155/190 + 0.4 ability_power
- TotalOnHitDamage: 10/15/20/25/30 + 0.1 ability_power
- TotalStaccatoDamage: 30/30/30/30/30 + 0.3 ability_power
- AuraDuration: 3/3/3/3/3
- AuraRange: 400/400/400/400/400
- BaseGlobalCD: 0.5/0.5/0.5/0.5/0.5
- EnemiesToHit: 2/2/2/2/2
- OnHitDuration: 5/5/5/5/5
- OnHitProcs: 1/1/1/1/1

**W · Ária da Perseverança** · custo 80/85/90/95/100 · alcance 1000/1000/1000/1000/1000

- TotalDiminuendoWeakenPercent: 0.25/0.25/0.25/0.25/0.25 + 0.0004 ability_power
- TotalHeal: 30/45/60/75/90 + 0.3 ability_power
- TotalShield: 25/45/65/85/105 + 0.25 ability_power
- AccelerandoShieldBreakpoint: 25/45/65/85/105
- AuraDuration: 3/3/3/3/3
- AuraRange: 400/400/400/400/400
- BaseGlobalCD: 0.5/0.5/0.5/0.5/0.5
- DiminuendoDuration: 3/3/3/3/3
- HealRange: 1000/1000/1000/1000/1000
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5

**E · Canção da Celeridade** · recarga 14/14/14/14/14 · custo 65/65/65/65/65 · alcance 1000/1000/1000/1000/1000

- TotalAllyMovementSpeed: 0.1/0.12/0.14/0.16/0.18 + 0.0002 ability_power
- TotalSelfMovementSpeed: 0.2/0.2/0.2/0.2/0.2 + 0.0002 ability_power
- TotalTempoMoveSpeedSlow: 0.5/0.5/0.5/0.5/0.5 + 0.0004 ability_power
- AllyMovementSpeedDuration: 3/3/3/3/3
- AuraDuration: 3/3/3/3/3
- AuraRange: 400/400/400/400/400
- BaseGlobalCD: 0.5/0.5/0.5/0.5/0.5
- SelfMovementSpeedDurationMax: 7/7/7/7/7
- SelfMovementSpeedDurationMin: 3/3/3/3/3
- TempoDuration: 2/2/2/2/2

**R · Crescendo** · recarga 140/120/100 · custo 100/100/100 · alcance 1000/1000/1000

- TotalDamage: 150/250/350 + 0.5 ability_power
- StunDuration: 1.5/1.5/1.5

## Kassadin

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 646 | 113 | 2567 |
| dano de ataque | 59 | 3.9 | 125.3 |
| armadura | 21 | 4 | 89 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.64 | 3.7 | 63.54 |
| regeneração de vida | 1.2 | 0.1 | 2.9 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Pedra do Vazio**

- DamageReductionPercent: 0.1

**Q · Esfera Nula** · recarga 9/8.5/8/7.5/7 · custo 60/65/70/75/80 · alcance 650/650/650/650/650

- TotalDamage: 65/95/125/155/185 + 0.7 ability_power
- TotalShield: 80/110/140/170/200 + 0.3 ability_power
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5

**W · Lâmina Ínfera** · recarga 7/7/7/7/7 · custo 1/1/1/1/1 · alcance 1/1/1/1/1

- ActiveDamage: 50/75/100/125/150 + 0.8 ability_power
- OnHitDamage: 25/25/25/25/25 + 0.1 ability_power
- ChampionMissingManaRatio: 20/22.5/25/27.5/30
- MissingManaRatio: 4/4.5/5/5.5/6

**E · Força de Pulso** · recarga 21/20/19/18/17 · custo 60/65/70/75/80

- TotalDamage: 70/100/130/160/190 + 0.7 ability_power
- ReductionPerSpellCast: 0.75/0.75/0.75/0.75/0.75
- SlowAmount: 50/55/60/65/70
- SlowDuration: 1/1/1/1/1

**R · Caminhar na Fenda** · recarga 5/3.5/2 · custo 40/40/40

- BaseDamage: 70/90/110 + 0.5 ability_power + 0.02 mana
- BonusDamage: 35/45/55 + 0.07 ability_power + 0.01 mana
- BaseCD: 5/3.5/2
- CastRange: 500/500/500
- MaxStacks: 4/4/4
- RBaseCost: 40/40/40
- RManaRatio: 2/2/2
- RStackDuration: 15/15/15
- RStackManaRatio: 1/1/1
- RiftWalkBaseDamage: 70/90/110

## Irelia

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 115 | 2585 |
| dano de ataque | 65 | 3.5 | 124.5 |
| armadura | 36 | 4.7 | 115.899994 |
| resistência mágica | 30 | 2.05 | 64.85 |
| velocidade de ataque | 0.656 | 2.5 | 43.156 |
| regeneração de vida | 0.7 | 0.17 | 3.59 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 200 | — | 200 |
| multiplicador de crítico | 2 | — | 2 |

**P · Fervor Ioniano** · recarga 0

- SingleStackAS: 25 _(no nível 18)_
- BuffDuration: 6
- MaxStacks: 4
- OnHitBaseDamage: 10
- OnHitPerLevel: 3
- OnHitStructureMod: 0.5

**Q · Surto da Lâmina** · recarga 10/9/8/7/6 · custo 15/15/15/15/15 · alcance 600/600/600/600/600

- ChampionDamage: 5/25/45/65/85 + 0.7 attack_damage
- HealAmount: 0/0/0/0/0 + 0.09 attack_damage
- MinionDamage: 55/75/95/115/135 + 0.7 attack_damage
- DashSpeedBonus: 1400/1400/1400/1400/1400

**W · Dança Desafiadora** · recarga 20/18/16/14/12 · custo 70/75/80/85/90 · alcance 775/775/775/775/775

- FinalMagicDR (= FinalPhysicalDR × 1): 35/35/35/35/35 + 0.035 ability_power _(no nível 18)_
- FinalPhysicalDR: 70/70/70/70/70 + 0.07 ability_power _(no nível 18)_
- MaxDamageCalc: 30/60/90/120/150 + 1.5 ability_power + 1.2 attack_damage
- MinDamageCalc: 10/20/30/40/50 + 0.5 ability_power + 0.4 attack_damage
- BasePhysicalDR: 50/50/50/50/50
- ChargeTimeForMax: 0.75/0.75/0.75/0.75/0.75
- MaxBonusRatio: 2/2/2/2/2
- MaxDuration: 1.5/1.5/1.5/1.5/1.5

**E · Dueto Impecável** · recarga 16/14.5/13/11.5/10 · custo 50/50/50/50/50

- TotalDamage: 70/110/150/190/230 + 1 ability_power
- BuffDuration: 3.5/3.5/3.5/3.5/3.5
- CDBetweenCast: 0.25/0.25/0.25/0.25/0.25
- MarkDuration: 5/5/5/5/5
- MaxRange: 775/775/775/775/775
- MinRange: 50/50/50/50/50
- RevTime: 0.35/0.35/0.35/0.35/0.35
- StunDuration: 0.75/0.75/0.75/0.75/0.75

**R · Lâmina da Vanguarda** · recarga 125/105/85 · custo 100/100/100 · alcance 1000/1000/1000

- MissileDamage: 125/200/275 + 0.7 ability_power
- ZoneDamage: 125/200/275 + 0.7 ability_power
- CCDuration: 1.5/1.5/1.5
- CooldownAmount: 0.5/1/1.5
- MarkDuration: 5/5/5
- SlowAmount: 90/90/90
- ZoneDuration: 2.5/2.5/2.5

## Janna

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 570 | 90 | 2100 |
| dano de ataque | 47 | 2.5 | 89.5 |
| armadura | 28 | 4.5 | 104.5 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 3 | 51.625 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Brisa de Impulso** · alcance 1200

- BonusDamage: 0 + 0.3 move_speed (bonus)
- MSToOnHitConversionRate: 0.3
- MSPercentAlly: 0.06

**Q · Ventania Uivante** · recarga 14/14/14/14/14 · custo 90/95/100/105/110

- ExtraDamagePerSecondCharged: 10/15/20/25/30 + 0.1 ability_power
- MaxDamage: 85/135/185/235/285 + 0.8 ability_power
- MaxKnockup: 1.25/1.25/1.25/1.25/1.25
- MinimumDamage: 55/90/125/160/195 + 0.5 ability_power
- BaseRange: 1100/1100/1100/1100/1100
- ChargeDistancePercent: 20/20/20/20/20
- MinionMod: 1/1/1/1/1
- MissileTravelTime: 1.25/1.25/1.25/1.25/1.25

**W · Zéfiro** · recarga 8/7.5/7/6.5/6 · custo 50/55/60/65/70 · alcance 550/550/550/550/550

- TotalDamage: 55/85/115/145/175 + 0.5 ability_power
- TotalMS: 0.06/0.07/0.08/0.09/0.1 + 0.0002 ability_power
- TotalSlow: 20/24/28/32/36 + 0.06 ability_power
- SlowDuration: 2/2/2/2/2

**E · Olho da Tempestade** · recarga 16/15/14/13/12 · custo 70/75/80/85/90 · alcance 800/800/800/800/800

- TotalAD: 10/15/20/25/30 + 0.1 ability_power
- TotalShield: 80/120/160/200/240 + 0.55 ability_power
- DecayGracePeriod: 4/4/4/4/4
- ECDRefundforCC: 0.2/0.2/0.2/0.2/0.2
- EmpowerDuration: 5/5/5/5/5
- ShieldDuration: 4/4/4/4/4

**R · Monção** · recarga 130/115/100 · custo 100/100/100 · alcance 725/725/725

- HealPerSecond: 100/150/200 + 0.5 ability_power
- TotalHeal (= HealPerSecond × 1): 300/450/600 + 1.5 ability_power
- KnockbackBaseRange: 875/875/875
- KnockbackDuration: 0.5/0.5/0.5
- KnockbackGravity: 10/10/10
- KnockbackSpeed: 1200/1200/1200
- MaxKnockback: 875/875/875
- Range: 700/700/700

## Gangplank

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 114 | 2568 |
| dano de ataque | 64 | 4.2 | 135.4 |
| armadura | 31 | 4.7 | 110.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 3.2 | 55.058002 |
| regeneração de vida | 1.2 | 0.12 | 3.24 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Julgamento de Fogo**

- MoveSpeed: 0.3 _(no nível 18)_
- TotalDamage: 250 + 1 attack_damage (bonus) _(no nível 18)_
- Cooldown: 15
- DoTDuration: 2.5
- MoveSpeedDuration: 2
- TurretDamageMult: 0.5

**Q · Negociarrr** · recarga 4.5/4.5/4.5/4.5/4.5 · custo 50/45/40/35/30 · alcance 625/625/625/625/625

- ShotDamage: 10/40/70/100/130 + 1 attack_damage
- CherryRatio: 0.07/0.07/0.07/0.07/0.07
- GameModeInteger: 1/1/1/1/1
- GoldProc: 3/4/5/6/7
- PlunderThreshold: 500/500/500/500/500
- SSProc: 4/5/6/7/8

**W · Remover Escorbuto** · recarga 22/20/18/16/14 · custo 60/70/80/90/100

- BaseHealth: 45/70/95/120/145 + 0.9 ability_power
- BuffDuration_unused: 0.25/0.25/0.25/0.25/0.25
- HasteAmount: 200/200/200/200/200
- PercentHeal: 13/13/13/13/13
- SteroidAmount_unused: 30/40/50/60/70

**E · Barril de Pólvora** · recarga 0/0/0/0/0 · alcance 1000/1000/1000/1000/1000

- BarrelDecayTime: 0.5/0.5/0.5/0.5/0.5 _(no nível 18)_
- FinalSlowAmount: 40/50/60/70/80
- BarrelArmorPenetration: 40/40/40/40/40
- BarrelBaseDecay: 2/2/2/2/2
- BarrelDuration: 25/25/25/25/25
- BarrelEnemyGoldBounty: 10/10/10/10/10
- BonusDamageToChampions: 75/95/115/135/155
- CooldownBetweenBarrels: 0.5/0.5/0.5/0.5/0.5
- DebuffDuration: 2/2/2/2/2
- MaxBarrels: 3/3/4/4/5
- MinionDamagePercentage: 100/100/100/100/100

**R · Barragem de Canhão** · recarga 160/140/120 · custo 100/100/100

- DeathsDaughterDamage: 120/210/300 + 0.3 ability_power
- OneWaveDamage: 40/70/100 + 0.1 ability_power
- TotalDamageTooltip (= OneWaveDamage × 1): 480/840/1200 + 1.2 ability_power
- CannonDelay: 0.5/0.5/0.5
- CannonInterval: 2/2/2
- DeathsDaughterSlow: 75/75/75
- DeathsDaughterSlowDuration: 1/1/1
- RaiseMoraleHaste: 40/40/40
- RaiseMoraleHasteDuration: 2/2/2
- SlowDuration: 0.5/0.5/0.5
- SlowPercent: 30/30/30
- ZoneDuration: 8/8/8

## Corki

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 100 | 2310 |
| dano de ataque | 52 | 2.5 | 94.5 |
| armadura | 27 | 4.5 | 103.5 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.644 | 2.8 | 48.244 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Munição Hextec**

- BasicAttackTOOLTIP: 0 + 0.2 attack_damage
- BonusMS: 0.4
- InitialCD: 10
- PackageDuration: 45
- SubsequentCD: 5

**Q · Bomba de Fósforo** · recarga 9/8.5/8/7.5/7 · custo 60/65/70/75/80 · alcance 825/825/825/825/825

- TotalDamage: 60/105/150/195/240 + 1 ability_power + 1.25 attack_damage (bonus)
- RevealDuration: 6/6/6/6/6

**W · Valquíria** · recarga 20/18/16/14/12 · custo 80/85/90/95/100

- DashSpeed: 650/650/650/650/650 + 1 move_speed
- MaximumDamage: 150/225/300/375/450 + 1.5 ability_power + 2 attack_damage (bonus)
- DamageRadius: 200/200/200/200/200
- MaximumRange: 600/600/600/600/600
- MaximumTicks: 5/5/5/5/5
- MinimumRange: 300/300/300/300/300
- TicksPerSecond: 2/2/2/2/2
- TrailDuration: 2.5/2.5/2.5/2.5/2.5

**E · Metralhadora** · recarga 12/12/12/12/12 · custo 50/55/60/65/70 · alcance 600/600/600/600/600

- TotalDamage: 80/130/180/230/280 + 2.4 attack_damage (bonus)
- ShredCap: 4/4/4/4/4
- ShredDuration: 2/2/2/2/2
- ShredMax: -12/-14/-16/-18/-20
- SprayDuration: 4/4/4/4/4
- TicksPerSecond: 4/4/4/4/4

**R · Barragem de Mísseis** · recarga 2/2/2 · custo 35/35/35

- AttackRefund: 1/1/1 + 2 critical_chance
- RBigMissileDamage (= RSmallMissileDamage × 1): 180/340/500 + 1.7 attack_damage (bonus)
- RSmallMissileDamage: 90/170/250 + 0.85 attack_damage (bonus)
- AmmoWhenLearned: 2/2/2
- MaxAmmoTOOLTIP: 4/4/4
- RBaseMissileRadius: 200/200/200
- RBigOneMissileRadius: 300/300/300

## Karma

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 109 | 2483 |
| dano de ataque | 49 | 3 | 100 |
| armadura | 28 | 5 | 113 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.3 | 39.725 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ímpeto Ardente** · recarga 8 · custo 40 · alcance 300

- SpellMantraRefund: 4

**Q · Chama Interior** · recarga 9/8/7/6/5 · custo 40/50/60/70/80

- IsEmpowered: 1/1/1/1/1
- TotalDamage: 60/110/160/210/260 + 0.7 ability_power
- {087bc3ca}: 0/0/0/0/0
- {5d5198eb}: 1/1/1/1/1
- SlowAmount: -0.4/-0.4/-0.4/-0.4/-0.4
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**W · Decisão Absorta** · recarga 12/12/12/12/12 · custo 50/55/60/65/70 · alcance 675/675/675/675/675

- InitialDamage: 40/65/90/115/140 + 0.45 ability_power
- LeashBreakRange: 825/825/825/825/825
- RootDuration: 1.6/1.7/1.8/1.9/2
- TetherDuration: 2/2/2/2/2

**E · Inspiração** · recarga 10/9.5/9/8.5/8 · custo 60/65/70/75/80 · alcance 800/800/800/800/800

- TotalShield: 80/130/180/230/280 + 0.6 ability_power
- MoveSpeed: 0.4/0.4/0.4/0.4/0.4
- MoveSpeedDuration: 2/2/2/2/2
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5

**R · Mantra** · recarga 40/38/36 · alcance 1100/1100/1100

- REBonusShield: 45/85/125 + 0.45 ability_power
- REBonusShieldArea (= REBonusShield × 1): 45/85/125 + 0.45 ability_power
- RQFieldDamage: 40/130/220 + 0.5 ability_power
- RQImpactDamage: 40/100/160 + 0.3 ability_power
- RWHealAmount: 17/17/17 + 0.01 ability_power
- REMoveSpeed: 0.15/0.15/0.15
- RQSlow: -0.5/-0.5/-0.5
- RQSlowDuration: 1.5/1.5/1.5
- RWBonusRoot: 0.5/0.75/1

## Taric

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 645 | 99 | 2328 |
| dano de ataque | 55 | 3.5 | 114.5 |
| armadura | 40 | 4.3 | 113.100006 |
| resistência mágica | 28 | 2.05 | 62.85 |
| velocidade de ataque | 0.625 | 2 | 34.625 |
| regeneração de vida | 1.2 | 0.1 | 2.9 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Bravata**

- CDR: 1
- TotalDamage: 93 + 0.15 armor (bonus) _(no nível 18)_
- Duration: 5

**Q · Fulgor Estelar** · recarga 3/3/3/3/3 · custo 60/60/60/60/60 · alcance 325/325/325/325/325

- HealingPerStack: 25/25/25/25/25 + 0.15 ability_power + 0.01 max_health
- MaxStackHealing: 25/25/25/25/25 + 0.15 ability_power + 0.01 max_health
- StackCooldown: 15/15/15/15/15
- BaseHealMax: 25/50/75/100/125

**W · Bastião** · recarga 15/15/15/15/15 · custo 60/60/60/60/60 · alcance 800/800/800/800/800

- BonusArmor: 0/0/0/0/0 + 0.06 armor
- LeashBreakActualRange: 1300/1300/1300/1300/1300
- LeashBreakWarningRange: 1000/1000/1000/1000/1000
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5
- ShieldHPRatio: 7/8/9/10/11

**E · Deslumbrar** · recarga 16/15/14/13/12 · custo 40/40/40/40/40

- TotalDamage: 90/130/170/210/250 + 0.5 ability_power + 0.5 armor (bonus)
- ChargeDuration: 1/1/1/1/1
- StunDuration: 1.5/1.5/1.5/1.5/1.5

**R · Resplendor Cósmico** · recarga 180/150/120 · custo 100/100/100

- InitialDelay: 2.5/2.5/2.5
- InvulnDuration: 2.5/2.5/2.5

## Veigar

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 580 | 108 | 2416 |
| dano de ataque | 52 | 2.7 | 97.9 |
| armadura | 18 | 5.2 | 106.399994 |
| resistência mágica | 32 | 1.3 | 54.1 |
| velocidade de ataque | 0.625 | 2.24 | 38.705 |
| regeneração de vida | 1.3 | 0.12 | 3.34 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Poder Maligno Fenomenal**

- APPerStack: 1
- DarkMatterCDRIncrement: 0.1
- PStacksPerDarkMatterCDR: 50
- dAbilityStacks: 1
- dQKillStacks: 1
- dQKillStacksLarge: 3
- dTakedownStacks: 5

**Q · Golpe Maligno** · recarga 6/5.5/5/4.5/4 · custo 30/35/40/45/50

- TotalDamage: 80/120/160/200/240 + 0.5 ability_power

**W · Matéria Escura** · recarga 0/0/0/0/0 · custo 60/65/70/75/80 · alcance 950/950/950/950/950

- TotalDamage: 85/140/195/250/305 + 0.7 ability_power
- BaseCooldown: 8/8/8/8/8
- ImpactDelay: 1.2/1.2/1.2/1.2/1.2

**E · Horizonte de Eventos** · recarga 20/18.5/17/15.5/14 · custo 70/75/80/85/90 · alcance 700/700/700/700/700

- CageDelay: 0.5/0.5/0.5/0.5/0.5
- StunDuration: 1.5/1.75/2/2.25/2.5

**R · Explosão Primordial** · recarga 120/90/60 · custo 100/100/100 · alcance 650/650/650

- MaxDamage (= MinDamage × 2): 350/500/650 + 1.3 ability_power
- MinDamage: 175/250/325 + 0.65 ability_power
- MaxExecuteMult: 2/2/2

## Trundle

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 110 | 2520 |
| dano de ataque | 68 | 4.5 | 144.5 |
| armadura | 37 | 4.5 | 113.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.67 | 2.9 | 49.97 |
| regeneração de vida | 1.2 | 0.15 | 3.7500002 |
| velocidade de movimento | 350 | — | 350 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Tributo do Rei** · alcance 1400

- RegenPercent: 0.055 _(no nível 18)_

**Q · Mordida** · recarga 3.5/3.5/3.5/3.5/3.5 · custo 20/20/20/20/20 · alcance 300/300/300/300/300

- TotalDamage: 10/30/50/70/90 + 1.15 attack_damage
- BonusAD: 20/25/30/35/40
- BonusRange: 25/25/25/25/25
- PrepDuration: 7/7/7/7/7
- SapDebuffDuration: 5/5/5/5/5
- SappedAD: -10/-12.5/-15/-17.5/-20
- SlowAmount: 0.75/0.75/0.75/0.75/0.75
- SlowDuration: 0.1/0.1/0.1/0.1/0.1
- TooltipQADRatio: 0.15/0.25/0.35/0.45/0.55

**W · Domínio Congelado** · recarga 18/17/16/15/14 · custo 40/40/40/40/40

- ASBonus: 0.3/0.45/0.6/0.75/0.9
- Duration: 8/8/8/8/8
- HealingBonus: 0.25/0.25/0.25/0.25/0.25
- MSBonus: 0.2/0.28/0.36/0.44/0.52

**E · Pilar de Gelo** · recarga 21/19.5/18/16.5/15 · custo 75/75/75/75/75 · alcance 1000/1000/1000/1000/1000

- AllyDisplacementRange: 150/150/150/150/150
- EnemyDisplacementGravity: 60/60/60/60/60
- EnemyDisplacementRange: 225/225/225/225/225
- EnemyDisplacementSpeed: 400/400/400/400/400
- KnockbackRadius: 225/225/225/225/225
- PillarDuration: 6/6/6/6/6
- SlowAmount: 34/38/42/46/50
- SlowRadius: 360/360/360/360/360

**R · Subjugar** · recarga 120/100/80 · custo 100/100/100 · alcance 650/650/650

- TotalPercentHPDamage: 0.2/0.25/0.3 + 0.0002 ability_power
- ActualDurationOfDrainBuff: 5/5/5
- ArmorMRShred: 0.4/0.4/0.4
- DurationOfDrainForTooltip: 4/4/4
- NumberOfTicks: 4/4/4
- PercentOfDamageAndShredUpfront: 0.5/0.5/0.5
- ScaleModelAmount: 0.02/0.02/0.02
- ShredDuration: 8/8/8

## Swain

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 595 | 99 | 2278 |
| dano de ataque | 58 | 2.7 | 103.9 |
| armadura | 25 | 4.7 | 104.899994 |
| resistência mágica | 31 | 1.55 | 57.35 |
| velocidade de ataque | 0.625 | 2.11 | 36.495 |
| regeneração de vida | 0.6 | 0.1 | 2.3 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Bando Voraz** · alcance 1150

- PassiveHealCalculated (= PassiveHealPercent × 1): 0 + 0.06 max_health
- PassiveHealPercent: 0.06
- SoulCollectionRange: 1100

**Q · Mão da Morte** · recarga 7/6/5/4/3 · custo 40/45/50/55/60

- ExtraBoltDamage (= InitialDamage × 1): 15/22.5/30/37.5/45 + 0.1125 ability_power
- InitialDamage: 60/90/120/150/180 + 0.45 ability_power
- MaxDamage (= InitialDamage × 1): 120/180/240/300/360 + 0.9 ability_power
- BoundaryCheckRange: 725/725/725/725/725

**W · Visão do Império** · recarga 22/21/20/19/18 · custo 60/65/70/75/80 · alcance 5500/6000/6500/7000/7500

- MinionDamage (= TotalDamage × 1): 35/52.5/70/87.5/105 + 0.3 ability_power
- TotalDamage: 70/105/140/175/210 + 0.6 ability_power
- EffectRadius: 325/325/325/325/325
- RevealDuration: 6/6/6/6/6
- Slow: -0.5/-0.5/-0.5/-0.5/-0.5
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- VisionRadius: 420/420/420/420/420

**E · Nuncamova** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70 · alcance 850/850/850/850/850

- SecondaryDamage: 90/130/170/210/250 + 0.7 ability_power
- CooldownRefund: 0.2/0.2/0.2/0.2/0.2
- ExplosionRadius: 205/205/205/205/205
- PullDistance: 290/290/290/290/290
- RootDuration: 1.5/1.5/1.5/1.5/1.5

**R · Ascensão Demoníaca** · recarga 120/120/120 · custo 100/100/100 · alcance 600/600/600

- DamageCalc: 15/25/35 + 0.04 ability_power
- DemonflareDamageTotal: 150/250/350 + 0.4 ability_power
- HealingCalc: 15/30/45 + 0.05 ability_power + 0.015 max_health (bonus)
- MinionMonsterHeal (= HealingCalc × 1): 1.5/3/4.5 + 0.0050000004 ability_power + 0.0015 max_health (bonus)
- AmpTime: 5/5/5
- DegenAmpAmount: 15/15/15
- DemonPowerDegen: 10/10/10
- DemonPowerMax: 50/50/50
- DemonPowerRegen: 20/20/20
- DemonflareCastDelay: 2/2/2
- DemonflareCooldownTooltip: 8/8/8
- DemonflareSlowAmount: 0.5/0.5/0.5
- DemonflareSlowDuration: 1.5/1.5/1.5
- MaxDemonflareCast: 100/100/100
- TimeBetweenTicks: 0.5/0.5/0.5

## Caitlyn

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 580 | 107 | 2399 |
| dano de ataque | 62 | 3.8 | 126.6 |
| armadura | 27 | 4.7 | 106.899994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.681 | 4 | 68.681 |
| regeneração de vida | 0.7 | 0.11 | 2.57 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 650 | — | 650 |
| multiplicador de crítico | 2 | — | 2 |

**P · Bem na Mira**

- BrushAttackTotal: 2
- HeadShotMinionBonusDamage: 0 + 0.099999994 attack_damage _(no nível 18)_
- {36ebd3cf}: 0 + 1 critical_damage
- AttacksPerHeadshot: 5
- TowerCap: 2.5

**Q · Pacificadora de Piltover** · recarga 10/9/8/7/6 · custo 55/60/65/70/75 · alcance 1300/1300/1300/1300/1300

- InitialDamage: 50/90/130/170/210 + 1.25 attack_damage
- SecondaryDamage (= InitialDamage × 1): 30.000002/54.000004/78/102.00001/126.00001 + 0.75 attack_damage

**W · Armadilha Mecânica Yordle** · recarga 0.5/0.5/0.5/0.5/0.5 · custo 20/20/20/20/20 · alcance 800/800/800/800/800

- HeadShotBonusDamage: 35/80/125/170/215 + 0.3 attack_damage (bonus)
- MaximumCharges: 3/3/4/4/5
- MaximumTraps: 3/3/4/4/5
- RootDuration: 1.5/1.5/1.5/1.5/1.5
- TrapDuration: 30/35/40/45/50

**E · Rede Calibre 90** · recarga 16/14/12/10/8 · custo 75/75/75/75/75

- NetDamage: 80/130/180/230/280 + 0.8 ability_power
- SlowAmount: 50/50/50/50/50
- SlowDuration: 1/1/1/1/1

**R · Ás na Manga** · recarga 90/90/90 · custo 100/100/100 · alcance 3500/3500/3500

- RTotalDamage: 300/475/650 + 1 attack_damage (bonus)

## Blitzcrank

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 109 | 2453 |
| dano de ataque | 62 | 3.5 | 121.5 |
| armadura | 37 | 4.7 | 116.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 1.13 | 19.835 |
| regeneração de vida | 1.5 | 0.15 | 4.05 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Barreira de Mana**

- ShieldAmount: 0 + 0.35 mana
- Cooldown: 90
- HealthThreshold: 0.3
- ManaPercent: 0.35
- ManaRatio: 1
- ShieldDuration: 10

**Q · Puxão Biônico** · recarga 20/19/18/17/16 · custo 100/100/100/100/100

- TotalDamage: 110/160/210/260/310 + 1.2 ability_power
- LolipopLength: 70/70/70/70/70

**W · Turbo** · recarga 15/15/15/15/15 · custo 75/75/75/75/75 · alcance 1/1/1/1/1

- {c01903c1}: 60/60/60/60/60
- AttackSpeedMod: 0.3/0.4/0.5/0.6/0.7
- Duration: 5/5/5/5/5
- MoveSpeedMod: 0.6/0.65/0.7/0.75/0.8
- MoveSpeedModMin: 0.1/0.1/0.1/0.1/0.1
- MoveSpeedModMinTime: 2.5/2.5/2.5/2.5/2.5
- MoveSpeedModReduction: 0.3/0.3/0.3/0.3/0.3
- PercentHealthDamage: 0.01/0.01/0.01/0.01/0.01
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**E · Punho do Poder** · recarga 7/6.5/6/5.5/5 · custo 25/25/25/25/25 · alcance 300/300/300/300/300

- TotalDamage: 0/0/0/0/0 + 0.25 ability_power + 2 attack_damage
- {a9d6b924}: 0/0/0/0/0 + 1.25 ability_power + 1.75 attack_damage
- CCDuration: 1/1/1/1/1

**R · Campo Estático** · recarga 60/40/20 · custo 100/100/100 · alcance 600/600/600

- ActiveDamage: 275/400/525 + 1 ability_power
- PassiveDamage: 50/100/150 + 0.3 ability_power + 0.02 mana
- ActiveRange: 600/600/600
- PassiveManaRatio: 0.05/0.05/0.05
- SilenceDuration: 0.5/0.5/0.5
- ZapCountdown: 1/1/1

## Malphite

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 665 | 104 | 2433 |
| dano de ataque | 62 | 4 | 130 |
| armadura | 40 | 4.95 | 124.149994 |
| resistência mágica | 28 | 2.05 | 62.85 |
| velocidade de ataque | 0.736 | 3.4 | 58.536003 |
| regeneração de vida | 1.4 | 0.11 | 3.27 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Escudo de Granito**

- PassiveCooldown: 6 _(no nível 18)_
- TotalShield: 0 + 0.1 max_health

**Q · Fragmento Sísmico** · recarga 8/8/8/8/8 · custo 70/75/80/85/90 · alcance 625/625/625/625/625

- QDamageCalc: 70/120/170/220/270 + 0.6 ability_power
- SlowDuration: 3/3/3/3/3
- SpeedSteal: 20/25/30/35/40

**W · Trovoada** · recarga 10/9.5/9/8.5/8 · custo 30/35/40/45/50

- ThunderclapSplash: 15/25/35/45/55 + 0.3 ability_power + 0.15 armor
- TotalBonusDamage: 30/40/50/60/70 + 0.2 ability_power + 0.15 armor
- {185669f9} (= TotalBonusDamage × 1): 60/80/100/120/140 + 0.4 ability_power + 0.3 armor
- {71941255}: 0.3/0.45000002/0.6/0.75/0.90000004
- MonsterDamageMod: 1.8/1.8/1.8/1.8/1.8
- ThunderClapSplashRange: 400/400/400/400/400
- ThunderclapBuffDuration: 5/5/5/5/5

**E · Estrondar Terreno** · recarga 7/7/7/7/7 · custo 50/55/60/65/70

- EDamageCalc: 60/95/130/165/200 + 0.6 ability_power + 0.4 armor
- ASReduction: 30/35/40/45/50
- Duration: 3/3/3/3/3

**R · Força Incontrolável** · recarga 130/115/100 · custo 100/100/100 · alcance 1000/1000/1000

- TotalDamage: 200/300/400 + 0.9 ability_power
- KnockupDuration: 1.5/1.5/1.5

## Katarina

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 672 | 108 | 2508 |
| dano de ataque | 58 | 3.2 | 112.4 |
| armadura | 32 | 4.7 | 111.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 2.74 | 47.238 |
| regeneração de vida | 1.5 | 0.14 | 3.88 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Voracidade** · alcance 340

- TotalDamage: 240.01538 + 1 ability_power + 0.6 attack_damage (bonus) _(no nível 18)_
- {e1ba6b0f}: 0.96 _(no nível 18)_
- DaggerDuration: 4
- DaggerRadius: 340
- ResetCDR: 15
- ResetWindow: 3

**Q · Lâmina Saltitante** · recarga 11/10/9/8/7 · alcance 625/625/625/625/625

- TotalDamage: 80/115/150/185/220 + 0.4 ability_power
- BounceOffset: 350/350/350/350/350
- BounceRadius: 450/450/450/450/450
- MaxBounces: 2/2/2/2/2

**W · Preparação** · recarga 15/14/13/12/11

- MovespeedAmount: 50/60/70/80/90
- MovespeedDuration: 1.25/1.25/1.25/1.25/1.25

**E · Shunpo** · recarga 12/11/10/9/8

- DaggerCooldownReduction (= {016cd4b3} × 1): 11.5199995/10.559999/9.599999/8.639999/7.68 _(no nível 18)_
- TooltipDaggerReduction: 0.96/0.96/0.96/0.96/0.96 _(no nível 18)_
- TotalDamage: 20/30/40/50/60 + 0.25 ability_power + 0.4 attack_damage
- {016cd4b3}: 0.96/0.96/0.96/0.96/0.96 _(no nível 18)_

**R · Lótus da Morte** · recarga 75/60/45 · alcance 550/550/550

- DamageCalc: 25/37.5/50 + 0.19 ability_power
- TotalDamageCalc (= DamageCalc × 1): 375/562.5/750 + 2.85 ability_power
- GrievousAmount: 0.4/0.4/0.4
- GrievousDuration: 3/3/3
- OnHitRatio: 0.25/0.3/0.35

## Nocturne

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 655 | 109 | 2508 |
| dano de ataque | 62 | 2.6 | 106.2 |
| armadura | 38 | 4.2 | 109.399994 |
| resistência mágica | 32 | 1.55 | 58.35 |
| velocidade de ataque | 0.721 | 2.7 | 46.621002 |
| regeneração de vida | 1.4 | 0.15 | 3.95 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Lâminas Sombrias** · recarga 12

- TotalDamageCrit: 0 + 1.1 attack_damage
- TotalDamageNoCrit: 0 + 1.2 attack_damage
- AACDR: 1
- AAChampMonsterCDR: 3
- Cooldown: 12
- DamageMinionMod: 0.5
- HealLevel18Value: 32
- HealLevel1Value: 13
- HealMinionMod: 0.5
- Radius: 360

**Q · Portador do Anoitecer** · recarga 8/8/8/8/8 · custo 60/65/70/75/80 · alcance 1200/1200/1200/1200/1200

- TotalDamage: 65/105/145/185/225 + 0.85 attack_damage (bonus)
- BonusTrailAD: 15/25/35/45/55
- MoveSpeed: 20/25/30/35/40
- TrailDuration: 5/5/5/5/5

**W · Proteção das Trevas** · recarga 20/18/16/14/12 · custo 50/50/50/50/50 · alcance 20/20/20/20/20

- ActiveAS: 30/35/40/45/50
- DoubleASDuration: 5/5/5/5/5
- PassiveAS: 0.3/0.05/0.05/0.05/0.05
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5

**E · Horror Indescritível** · recarga 15/14/13/12/11 · custo 60/65/70/75/80 · alcance 425/425/425/425/425

- TotalDamage: 80/125/170/215/260 + 1 ability_power
- CCDuration: 1.25/1.5/1.75/2/2.25
- LeashBreakRange: 465/465/465/465/465
- LeashDuration: 2/2/2/2/2
- TooltipFearMS: 0.9/0.9/0.9/0.9/0.9

**R · Paranoia** · recarga 140/115/90 · custo 100/100/100 · alcance 2500/3250/4000

- Damage: 150/275/400 + 1.2 attack_damage (bonus)
- DashSpeed: 1800/1800/1800
- ParanoiaDuration: 6/6/6
- TrackDistance: 3500/4250/5000

## Maokai

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 665 | 109 | 2518 |
| dano de ataque | 64 | 3.3 | 120.1 |
| armadura | 35 | 5.2 | 123.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.8 | 2.125 | 36.925 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Seiva Mágica**

- PassiveCooldown: 20 _(no nível 18)_
- PassiveHealingTotal: 0 + 0.04 max_health
- JungPassCooldownReduction: 1.5
- PassiveCooldownReduction: 4
- PassiveHealthThreshold: 0.95

**Q · Esmagamento Espinhoso** · recarga 7/6.5/6/5.5/5 · custo 40/40/40/40/40

- TotalDamage: 75/120/165/210/255 + 0.5 ability_power
- TotalPercentHealth: 0.02/0.025/0.03/0.035/0.04
- BonusMonsterDamage: 150/160/170/180/190
- KnockbackDistance: 400/400/400/400/400
- KnockbackSpeed: 950/950/950/950/950
- LolipopRadius: 325/325/325/325/325
- NonChampMax: 9999/9999/9999/9999/9999
- SlowAmount: 0.99/0.99/0.99/0.99/0.99
- SlowDuration: 0.25/0.25/0.25/0.25/0.25

**W · Avanço Retorcido** · recarga 14/13/12/11/10 · custo 60/60/60/60/60 · alcance 525/525/525/525/525

- TotalDamage: 60/85/110/135/160 + 0.4 ability_power
- DashSpeed: 1300/1300/1300/1300/1300
- RootDuration: 1/1.1/1.2/1.3/1.4

**E · Atirar Mudas** · recarga 18/17/16/15/14 · custo 60/65/70/75/80 · alcance 1100/1100/1100/1100/1100

- EmpoweredSaplingDuration: 30/30/30/30/30 + 0.015 max_health (bonus)
- EmpoweredSlowAmount: 0.45/0.45/0.45/0.45/0.45 + 0.0001 ability_power + 0.0001 max_health (bonus)
- SaplingMoveSpeed: 400/400/400/400/400
- TotalDamage: 50/75/100/125/150 + 0.25 ability_power + 0.05 max_health (bonus)
- TotalEmpoweredDamage: 100/150/200/250/300 + 0.5 ability_power + 0.1 max_health (bonus)
- EmpoweredDoTDuration: 2/2/2/2/2
- GrowthDuration: 0.5/0.5/0.5/0.5/0.5
- MaxTraps: 25000/25000/25000/25000/25000
- SaplingDetectionRadius: 550/550/550/550/550
- SlowAmount: 0.45/0.45/0.45/0.45/0.45
- SlowDuration: 2/2/2/2/2

**R · Garras da Natureza** · recarga 130/110/90 · custo 100/100/100

- TotalDamage: 150/225/300 + 0.75 ability_power
- HasteDuration: 2/2/2
- MaxRootDuration: 2.25/2.25/2.25
- MinRootDuration: 0.75/0.75/0.75
- MoveHaste: 0.4/0.5/0.6
- RootDuration: 0.6/0.6/0.6

## Renekton

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 660 | 111 | 2547 |
| dano de ataque | 69 | 4.15 | 139.55 |
| armadura | 35 | 5.2 | 123.399994 |
| resistência mágica | 28 | 2.05 | 62.85 |
| velocidade de ataque | 0.665 | 2.75 | 47.415 |
| regeneração de vida | 1.6 | 0.15 | 4.15 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Domínio da Ira** · recarga 0

- FuryCost: 50
- FuryIncreasePercent: 0.5
- FuryPerAttack: 5
- InCombatDuration: 12
- LowHealthPercentThreshold: 0.5

**Q · Abater os Indefesos** · recarga 7/7/7/7/7 · alcance 325/325/325/325/325

- BasicDamage: 60/90/120/150/180 + 1 attack_damage (bonus)
- ChampHealing: 12/20/28/36/44 + 0.17 attack_damage (bonus)
- EmpChampHealing (= ChampHealing × 1): 36/60/84/108/132 + 0.51 attack_damage (bonus)
- EmpDamage: 90/135/180/225/270 + 1.4 attack_damage (bonus)
- EmpNonChampHealing (= NonChampHealing × 1): 6/9/12/15/18 + 0.06 attack_damage (bonus)
- NonChampHealing: 2/3/4/5/6 + 0.02 attack_damage (bonus)
- BasicHealCap: 50/75/100/125/150
- ChampionFuryGain: 10/10/10/10/10
- EmpoweredHealCap: 200/300/400/500/600
- FuryGainCap: 30/30/30/30/30
- MinionFuryGain: 2.5/2.5/2.5/2.5/2.5
- Range: 400/400/400/400/400
- UltimateRangeMod: 1.2/1.2/1.2/1.2/1.2

**W · Predador Desumano** · recarga 16/14/12/10/8 · alcance 300/300/300/300/300

- BasicTotalDamage (= HitDamage × 2): 10/40/70/100/130 + 1.5 attack_damage
- EmpTotalDamage (= HitDamage × 3): 15/60/105/150/195 + 2.25 attack_damage
- HitDamage: 5/20/35/50/65 + 0.75 attack_damage
- BonusFuryVsChamps: 10/10/10/10/10
- BuffDuration: 6/6/6/6/6
- EnragedStunDuration: 1.5/1.5/1.5/1.5/1.5
- StunDuration: 0.75/0.75/0.75/0.75/0.75

**E · Fatiar e Picar** · recarga 16/14.5/13/11.5/10

- BasicDamage: 40/70/100/130/160 + 0.9 attack_damage (bonus)
- DashSpeed: 750/750/750/750/750 + 1 move_speed
- EmpDamage: 70/115/160/205/250 + 1.35 attack_damage (bonus)
- ChampionRageGeneration: 10/10/10/10/10
- DiceTimer: 4/4/4/4/4
- EnragedArmorShred: 25/27.5/30/32.5/35
- FuryMax: 30/30/30/30/30
- MinionRageGeneration: 2/2/2/2/2
- ShredTimer: 4/4/4/4/4

**R · Dominus** · recarga 120/100/80 · alcance 20/20/20

- TotalDamagePerSecond: 60/150/240 + 0.1 ability_power + 0.1 attack_damage (bonus)
- BuffDuration: 15/15/15
- FuryOnCast: 20/20/20
- FuryPerSecond: 5/5/5
- HealthGain: 300/500/700
- TickRate: 0.5/0.5/0.5

## Jarvan IV

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 104 | 2408 |
| dano de ataque | 64 | 3 | 115 |
| armadura | 36 | 4.6 | 114.2 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 2.5 | 43.158 |
| regeneração de vida | 1.6 | 0.14 | 3.98 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Cadência Marcial**

- TooltipCooldown: 3 _(no nível 18)_
- MaximumCadenceDamage: 400
- MinimumCadenceDamage: 20
- TooltipCurrentHealthDamage: 0.08

**Q · Ataque do Dragão** · recarga 10/9/8/7/6 · custo 45/50/55/60/65 · alcance 770/770/770/770/770

- TotalDamage: 90/130/170/210/250 + 1.45 attack_damage (bonus)
- BaseARShred: 0.1/0.14/0.18/0.22/0.26

**W · Égide de Ouro** · recarga 9/9/9/9/9 · custo 30/30/30/30/30 · alcance 625/625/625/625/625

- BonusShield: 0/0/0/0/0 + 0.013 max_health
- BaseSlowAmount: 0.15/0.2/0.25/0.3/0.35

**E · Estandarte Demaciano** · recarga 12/11.5/11/10.5/10 · custo 55/55/55/55/55

- BaseAuraAS: 0.2/0.225/0.25/0.275/0.3
- EBehindJarvanCheck: 400/400/400/400/400
- EDashSpeed: 1400/1400/1400/1400/1400
- EKnockUpAoE: 180/180/180/180/180
- PermanentAttackSpeed: 0.2/0.225/0.25/0.275/0.3

**R · Cataclisma** · recarga 120/105/90 · custo 100/100/100 · alcance 650/650/650

- DamageCalc: 200/325/450 + 1.8 attack_damage (bonus)
- WallDuration: 3.5/3.5/3.5

## Elise

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 109 | 2473 |
| dano de ataque | 55 | 3 | 106 |
| armadura | 30 | 4.5 | 106.5 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 1.75 | 30.375 |
| regeneração de vida | 1.1 | 0.12 | 3.1399999 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Aranha Rainha**

- {38943004}: 5
- SpiderlingAPRatio: 0.15
- SpiderlingAoEReduction: 0.75

**Q · Neurotoxina / Mordida Venenosa** · recarga 6/6/6/6/6 · custo 80/85/90/95/100 · alcance 575/575/575/575/575

- HumanPercentHealth: 4/4/4/4/4 + 0.03 ability_power
- MonsterDamageCapCalc: 65/85/105/125/145 + 0.9 ability_power
- BaseDamage: 40/70/100/130/160
- SpiderBaseDamage: 50/80/110/140/170

**W · Cria Volátil / Frenesi Aracnídeo** · recarga 12/12/12/12/12 · custo 60/70/80/90/100 · alcance 950/950/950/950/950

- TotalDamage: 60/100/140/180/220 + 0.75 ability_power
- ExplosionSize: 275/275/275/275/275
- TimeToExplode: 3/3/3/3/3

**E · Casulo / Rapel** · recarga 12/11.5/11/10.5/10 · custo 50/50/50/50/50 · alcance 1100/1100/1100/1100/1100

- TotalStunDuration: 1.6/1.8/2/2.2/2.4
- RappelCooldown: 22/21/20/19/18
- RappelDamageAmp: 40/55/70/85/100
- VisionRadius: 250/250/250/250/250

**R · Forma de Aranha** · recarga 3/3/3 · alcance 20/20/20

- PassiveTotalDamage: 12/22/32 + 0.15 ability_power
- PassiveTotalHealing: 6/8/10 + 0.08 ability_power
- SpiderlingTotalDamage: 10/20/30 + 0.15 ability_power
- BaseSpiderlingsStored: 2/3/4
- PassiveMSOnHit: 0.1/0.1/0.1
- SpiderlingArmor: 30/50/70
- SpiderlingBonusResist: 0/20/40
- SpiderlingMagicResist: 50/70/90
- SpiderlingsStoredCap: 20/20/20

## Orianna

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 565 | 110 | 2435 |
| dano de ataque | 44 | 2.6 | 88.2 |
| armadura | 20 | 4.2 | 91.399994 |
| resistência mágica | 26 | 1.3 | 48.1 |
| velocidade de ataque | 0.658 | 3.5 | 60.158 |
| regeneração de vida | 1.4 | 0.11 | 3.27 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Dando Corda**

- StackDamage (= TotalDamage × 1): 7.5000005 + 0.0225 ability_power _(no nível 18)_
- StackDamageMax (= TotalDamage × 1): 15.000001 + 0.045 ability_power _(no nível 18)_
- TotalDamage: 50 + 0.15 ability_power _(no nível 18)_
- StackCount: 2
- StackDuration: 4

**Q · Comando: Atacar** · recarga 7/6/5/4/3 · custo 35/35/35/35/35 · alcance 825/825/825/825/825

- MinimumDamageTooltip (= TotalDamageTooltip × 1): 42/63/84/105/126 + 0.385 ability_power
- TotalDamageTooltip: 60/90/120/150/180 + 0.55 ability_power
- ReducedDamagePercent: 30/30/30/30/30

**W · Comando: Dissonância** · recarga 7/7/7/7/7 · custo 60/65/70/75/80

- TotalDamage: 70/110/150/190/230 + 0.8 ability_power
- FieldDuration: 3/3/3/3/3
- HasteAmount: 0.2/0.25/0.3/0.35/0.4
- SlowAmount: 0.2/0.25/0.3/0.35/0.4
- SlowAndHasteDuration: 2/2/2/2/2

**E · Comando: Proteger** · recarga 9/9/9/9/9 · custo 60/60/60/60/60 · alcance 1120/1120/1120/1120/1120

- TotalDamageTooltip: 60/90/120/150/180 + 0.3 ability_power
- TotalShieldTooltip: 55/90/125/160/195 + 0.45 ability_power
- DefenseBonus: 6/12/18/24/30
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5

**R · Comando: Onda de Choque** · recarga 110/95/80 · custo 100/100/100

- TotalDamage: 225/350/475 + 1.1 ability_power
- AoERadius: 415/415/415
- StunDuration: 0.75/0.75/0.75

## Wukong

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 99 | 2293 |
| dano de ataque | 66 | 3.5 | 125.5 |
| armadura | 31 | 4.7 | 110.899994 |
| resistência mágica | 28 | 2.05 | 62.85 |
| velocidade de ataque | 0.69 | 3 | 51.69 |
| regeneração de vida | 0.7 | 0.13 | 2.9099998 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Pele de Pedra** · recarga 20

- BonusArmor: 10 _(no nível 18)_
- TooltipMaxArmor (= BonusArmor × 1): 60 _(no nível 18)_
- CombatDuration: 3
- FallOffRate: 1
- HealthPercentPer5: 0.0035
- MaxStacks: 5
- StackDuration: 5
- StackMultiplier: 1
- TooltipMaxHealthPercentPer5: 2.1

**Q · Golpe Destruidor** · recarga 8/7.5/7/6.5/6 · custo 20/20/20/20/20

- BonusDamageTT: 20/45/70/95/120 + 0.5 attack_damage (bonus)
- TotalDamage: 20/45/70/95/120 + 0.5 attack_damage (bonus)
- ArmorShredPercent: 0.1/0.15/0.2/0.25/0.3
- AttackRangeBonus: 135/145/155/165/175
- BuffDuration: 6/6/6/6/6
- CooldownDecrease: 0.5/0.5/0.5/0.5/0.5
- ShredDuration: 3/3/3/3/3

**W · Guerreiro Trapaceiro** · recarga 22/21/20/19/18 · custo 60/55/50/45/40

- CloneDamageMod: 0.4/0.45/0.5/0.55/0.6
- CloneDuration: 4/4/4/4/4
- DashSpeed: 900/900/900/900/900
- MinRange: 100/100/100/100/100
- RangeClamp: 300/300/300/300/300
- StealthDuration: 1/1/1/1/1

**E · Resplendor das Nuvens** · recarga 10/9.25/8.5/7.75/7 · custo 30/35/40/45/50 · alcance 650/650/650/650/650

- TotalDamage: 80/120/160/200/240 + 1 ability_power
- TotalDamageMonsters (= TotalDamage × 1): 80/120/160/200/240 + 1 ability_power
- AttackSpeed: 0.4/0.45/0.5/0.55/0.6
- AttackSpeedDuration: 5/5/5/5/5
- DashSpeed: 1050/1050/1050/1050/1050
- EndPointOffsetDistance: 75/75/75/75/75
- ExtraTargetRange: 700/700/700/700/700
- ExtraTargets: 2/2/2/2/2

**R · Ciclone** · recarga 130/110/90 · custo 100/100/100 · alcance 660/660/660

- DamagePerSecondTotal: 0/0/0 + 1.375 attack_damage
- MonsterCap: 600/600/600 _(no nível 18)_
- MonsterCapTT (= MonsterCap × 1): 1200/1200/1200 _(no nível 18)_
- PercentHPDamageTT (= {26c20668} × 1): 0.08/0.12/0.16
- TotalDamageTT (= DamagePerSecondTotal × 1): 0/0/0 + 2.75 attack_damage
- {26c20668}: 0.04/0.06/0.08
- KnockupDuration: 0.6/0.6/0.6
- LockoutTimeBetweenCasts: 1/1/1
- MoveSpeed: 0.2/0.2/0.2
- RecastWindow: 8/8/8
- SecondsPerTick: 0.25/0.25/0.25

## Brand

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 570 | 105 | 2355 |
| dano de ataque | 57 | 3 | 108 |
| armadura | 24 | 4.2 | 95.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.681 | 2 | 34.681 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Labareda**

- ExplosionDamage: 12 + 0.02 ability_power _(no nível 18)_
- JungleMonsterExplosionFlatCap: 525 _(no nível 18)_
- ManaRestore: 39.982998 _(no nível 18)_
- {4124b2bf}: 20
- {afbf404a}: 12.5
- ManaRestoration: 20
- MonsterMod: 270
- PercentHealthDamage: 2

**Q · Cauterizar** · recarga 8/7.5/7/6.5/6 · custo 70/70/70/70/70

- TotalDamage: 70/100/130/160/190 + 0.65 ability_power
- StunDuration: 1.75/1.75/1.75/1.75/1.75

**W · Pilar de Chamas** · recarga 10/9.5/9/8.5/8 · custo 60/70/80/90/100 · alcance 900/900/900/900/900

- EmpoweredDamage (= TotalDamage × 1.25): 93.75/150/206.25/262.5/318.75 + 0.875 ability_power
- TotalDamage: 75/120/165/210/255 + 0.7 ability_power
- PDamageBonus: 0.25/0.25/0.25/0.25/0.25

**E · Conflagração** · recarga 13/12/11/10/9 · custo 90/90/90/90/90 · alcance 675/675/675/675/675

- EDamageCalc: 55/80/105/130/155 + 0.6 ability_power

**R · Piroclasma** · recarga 100/90/80 · custo 100/100/100 · alcance 750/750/750

- TotalDamage: 100/175/250 + 0.3 ability_power
- SlowAmount: 30/45/60
- SlowDuration: 0.25/0.25/0.25

## Lee Sin

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 645 | 108 | 2481 |
| dano de ataque | 66 | 3.4 | 123.8 |
| armadura | 36 | 4.5 | 112.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.651 | 3 | 51.651 |
| regeneração de vida | 1.5 | 0.14 | 3.88 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Agitação**

- EnergyReturn: 20 _(no nível 18)_
- TTFirstHitEnergy (= EnergyReturn × 1): 40 _(no nível 18)_
- PassiveAS: 40

**Q · Onda Sônica / Ataque Ressonante** · recarga 10/9/8/7/6 · custo 50/50/50/50/50 · alcance 1200/1200/1200/1200/1200

- EmpoweredDamage (= RecastDamage × 2): 120/180/240/300/360 + 1.8 attack_damage (bonus)
- InitialDamage: 60/90/120/150/180 + 0.9 attack_damage (bonus)
- RecastDamage: 60/90/120/150/180 + 0.9 attack_damage (bonus)
- CherryBonusHaste: 30/30/30/30/30
- DashSpeed: 1350/1350/1350/1350/1350
- Q2MaxMissingHealthMod: 1/1/1/1/1
- ReactivateTime: 3/3/3/3/3

**W · Proteger / Vontade de Ferro** · recarga 7/7/7/7/7 · custo 50/50/50/50/50 · alcance 700/700/700/700/700

- ShieldAmount: 60/105/150/195/240 + 0.8 ability_power
- DashSpeed: 1350/1350/1350/1350/1350
- LifestealAndSpellVamp: 10/14/18/22/26
- LifestealAndSpellVampTime: 4/4/4/4/4
- ShieldDuration: 2/2/2/2/2
- W1CooldownRecovered: 0.5/0.5/0.5/0.5/0.5
- W1ReactivateTime: 3/3/3/3/3

**E · Tempestade / Mutilar** · recarga 8/8/8/8/8 · custo 50/50/50/50/50 · alcance 450/450/450/450/450

- InitialDamage: 35/60/85/110/135 + 0.9 attack_damage
- ReactivateTime: 3/3/3/3/3
- SlowAmount: 35/45/55/65/75
- SlowDuration: 4/4/4/4/4

**R · Fúria do Dragão** · recarga 110/85/60 · alcance 375/375/375

- Damage: 175/400/625 + 2 attack_damage (bonus)
- KickDistance: 800/800/800
- PercentHPCarryThrough: 12/15/18

## Vayne

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 550 | 103 | 2301 |
| dano de ataque | 60 | 2.35 | 99.95 |
| armadura | 23 | 4.6 | 101.2 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.658 | 3.3 | 56.758 |
| regeneração de vida | 0.7 | 0.11 | 2.57 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Caçadora Noturna** · alcance 2000

- MovementSpeed: 30
- Range: 2000

**Q · Rolamento** · recarga 6/5/4/3/2 · custo 30/30/30/30/30

- ADRatioBonus: 0/0/0/0/0 + 0.5 ability_power + 0.75 attack_damage
- Duration: 3/3/3/3/3

**W · Dardos de Prata** · recarga 0/0/0/0/0

- TotalDamage: 0.06/0.07/0.08/0.09/0.1
- DamageFloor: 50/65/80/95/110
- DamageVsMonsters: 140/155/170/185/200
- DebuffDuration: 3.5/3.5/3.5/3.5/3.5

**E · Condenar** · recarga 20/18/16/14/12 · custo 90/90/90/90/90 · alcance 550/550/550/550/550

- EmpoweredDamageTT (= TotalDamage × 1): 75/127.5/180/232.5/285 + 0.75 attack_damage (bonus)
- TotalDamage: 50/85/120/155/190 + 0.5 attack_damage (bonus)
- KnockbackDistance: 475/475/475/475/475
- StunDuration: 1.5/1.5/1.5/1.5/1.5

**R · Hora Final** · recarga 100/85/70 · custo 80/80/80 · alcance 1/1/1

- BaseDuration: 8/10/12
- BonusAttackDamage: 35/50/65
- DamagedMarkerDuration: 3/3/3
- DurationToAdd: 4/4/4
- MovementSpeed: 90/90/90
- TumbleCDReduction: 30/40/50
- TumbleStealthDuration: 1/1/1

## Rumble

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 105 | 2425 |
| dano de ataque | 64 | 3.2 | 118.4 |
| armadura | 36 | 4.7 | 115.899994 |
| resistência mágica | 28 | 1.55 | 54.35 |
| velocidade de ataque | 0.644 | 1.85 | 32.094 |
| regeneração de vida | 1.4 | 0.12 | 3.4399998 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Titã do Ferro-Velho**

- MonsterCapScaling: 150 _(no nível 18)_
- OverheatAS: 1.3 _(no nível 18)_
- TotalBaseDamage: 40 + 0.25 ability_power _(no nível 18)_
- DangerZoneHeat: 50
- HeatDecayPerS: -10
- MonsterCap: 80
- OverheatDuration: 4
- OverheatPercBonusDamage: 0.04
- OverheatingHeat: 150

**Q · Cospe-Fogo** · recarga 10/9/8/7/6 · alcance 600/600/600/600/600

- EmpoweredDamage (= FlatDamage × 1): 75/112.5/150/187.5/225 + 1.5749999 ability_power
- EmpoweredHealth: 0.06/0.065/0.07/0.075/0.08
- FlatDamage: 50/75/100/125/150 + 1.05 ability_power
- MinionDamage (= FlatDamage × 1): 35/52.5/70/87.5/105 + 0.73499995 ability_power
- MonsterCap: 300/300/300/300/300 _(no nível 18)_
- BurnDuration: 0.6/0.6/0.6/0.6/0.6
- DebuffTicks: 3/3/3/3/3
- FlamespitterDuration: 3/3/3/3/3
- InitialHeatCost: 20/20/20/20/20
- TickRate: 0.25/0.25/0.25/0.25/0.25
- TotalTicks: 12/12/12/12/12

**W · Escudo de Sucata** · recarga 6/6/6/6/6 · alcance 20/20/20/20/20

- EmpoweredMS: 0.1/0.15/0.2/0.25/0.3
- EmpoweredShield (= TotalShield × 1.5): 37.5/82.5/127.5/172.5/217.5 + 0.45000002 ability_power + 0.06 max_health
- TotalShield: 25/55/85/115/145 + 0.3 ability_power + 0.04 max_health
- HeatCost: 20/20/20/20/20
- MoveSpeedDuration: 1/1/1/1/1
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5

**E · Arpão Elétrico** · recarga 0.5/0.5/0.5/0.5/0.5

- EmpDamage (= TotalDamage × 1.5): 82.5/120/157.5/195/232.5 + 0.75 ability_power
- TotalDamage: 55/80/105/130/155 + 0.5 ability_power
- BaseSlowAmount: 15/20/25/30/35
- EmpoweredSlowAmount: 30/40/50/60/70
- EnhancedMagicPen: 0.2/0.24/0.28/0.32/0.36
- FirstCastHeatCost: 20/20/20/20/20
- PercMagicPen: 0.1/0.12/0.14/0.16/0.18
- ShredDuration: 4/4/4/4/4
- SlowDuration: 2/2/2/2/2

**R · O Equalizador** · recarga 130/105/80 · alcance 1700/1700/1700

- DamagePerSecond: 120/200/280 + 0.35 ability_power
- TotalDamage (= DamagePerSecond × 1): 540/900/1260 + 1.5749999 ability_power
- BurnLingerTime: 1/1/1
- MaxBurnSeconds: 5/5/5
- NumMissiles: 6/6/6
- SlowAmount: 35/35/35
- TickRate: 0.25/0.25/0.25

## Cassiopeia

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 98 | 2296 |
| dano de ataque | 53 | 3 | 104 |
| armadura | 18 | 4.7 | 97.899994 |
| resistência mágica | 32 | 1.3 | 54.1 |
| velocidade de ataque | 0.647 | 1.5 | 26.147 |
| regeneração de vida | 1.1 | 0.1 | 2.8 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Graça Serpentina**

- PercentHasteMod: 0.06

**Q · Explosão Venenosa** · recarga 3.5/3.5/3.5/3.5/3.5 · custo 50/55/60/65/70 · alcance 850/850/850/850/850

- TooltipTotalDamage: 75/110/145/180/215 + 0.65 ability_power
- ChampHitMSBonus: 30/35/40/45/50
- ChampHitMSDuration: 3/3/3/3/3
- NumDamageTicks: 7/7/7/7/7
- PoisonDuration: 3/3/3/3/3

**W · Miasma** · recarga 24/22/20/18/16 · custo 70/75/80/85/90

- DamagePerSecond: 20/25/30/35/40 + 0.1 ability_power
- CloudDuration: 5/5/5/5/5
- PoisonDuration: 1/1/1/1/1
- SlowPercent: 40/50/60/70/80

**E · Presas Duplas** · recarga 0.75/0.75/0.75/0.75/0.75 · custo 40/40/40/40/40 · alcance 700/700/700/700/700

- BasicDamage: 52/52/52/52/52 + 0.1 ability_power
- BonusPoisonedDamage: 20/45/70/95/120 + 0.55 ability_power
- HealCalc: 0/0/0/0/0 + 0.1 ability_power
- HealCalcMinion (= HealCalc × 1): 0/0/0/0/0 + 0.025 ability_power

**R · Olhar Petrificador** · recarga 120/100/80 · custo 100/100/100 · alcance 825/825/825

- RDamage: 150/250/350 + 0.5 ability_power
- RCCDuration: 2/2/2
- RSlowPercent: 40/40/40

## Skarner

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 110 | 2500 |
| dano de ataque | 63 | 5 | 148 |
| armadura | 33 | 4.5 | 109.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2 | 34.625 |
| regeneração de vida | 1.5 | 0.15 | 4.05 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Linhas de Vibração**

- DamageMonsterCap: 300 _(no nível 18)_
- PercentHealthDamage: 9 _(no nível 18)_
- Duration: 4
- StacksToTriggerPassive: 3
- TickFrequency: 0.5

**Q · Terra Arrasada/Sublevação** · recarga 8/6.75/5.5/4.25/3 · custo 30/30/30/30/30

- AbilityDamage: 10/20/30/40/50 + 0.9 attack_damage (bonus) + 0.03 max_health (bonus)
- AttackRadius: 300/300/300/300/300
- AttackRange: 25/25/25/25/25
- AttackSpeed: 0.2/0.25/0.3/0.35/0.4
- MaxHPPercent: 0.11/0.11/0.11/0.11/0.11
- MonsterDamageCap: 150/200/250/300/350
- RockHoldDuration: 5/5/5/5/5
- SlowDuration: 1/1/1/1/1
- SlowPercent: 0.4/0.4/0.4/0.4/0.4
- TurretDamageMod: 1/1/1/1/1

**W · Bastião Sísmico** · recarga 10/9/8/7/6 · custo 60/65/70/75/80 · alcance 650/650/650/650/650

- Damage: 50/70/90/110/130 + 0.8 ability_power
- InitialShield: 0/0/0/0/0 + 0.08 max_health
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5
- SlowDuration: 1/1/1/1/1
- SlowEffect: -0.2/-0.2/-0.2/-0.2/-0.2

**E · Impacto de Ixtal** · recarga 22/21/20/19/18 · custo 50/55/60/65/70 · alcance 1700/1700/1700/1700/1700

- PinDamage: 30/60/90/120/150 + 1.2 attack_damage (bonus) + 0.06 max_health
- Acceleration: 100/100/100/100/100
- ChargeDuration: 2.75/2.75/2.75/2.75/2.75
- DistancePostGrab: 675/675/675/675/675
- FinalHitboxDelay: 0.4/0.4/0.4/0.4/0.4
- GrabHitboxOffset: 90/90/90/90/90
- GrabHitboxRadius: 160/160/160/160/160
- GrabbedMoveSpeedBonus: 200/200/200/200/200
- InitialHitboxOffset: 90/90/90/90/90
- InitialHitboxRadius: 160/160/160/160/160
- InitialSpeed: 150/150/150/150/150
- MaximumSpeed: 950/950/950/950/950
- MinGrabbedSpeed: 650/650/650/650/650
- MinimumQRenew: 1.5/1.5/1.5/1.5/1.5
- RefundPercent: 0.65/0.65/0.65/0.65/0.65
- SecondHitboxDelay: 0.3/0.3/0.3/0.3/0.3
- SecondHitboxOffset: 150/150/150/150/150
- SecondHitboxRadius: 100/100/100/100/100
- StunDuration: 1.1/1.1/1.1/1.1/1.1
- VictimOffsetIncrease: 200/200/200/200/200
- WallRevealRadius: 650/650/650/650/650
- WallWarningRadius: 1750/1750/1750/1750/1750

**R · Empalar** · recarga 120/105/90 · custo 100/100/100 · alcance 625/625/625

- Damage: 150/250/350 + 1 ability_power
- BackwardOffset: 50/50/50
- DropDistance: 200/200/200
- Duration: 1.5/1.5/1.5
- HitboxEndHalfWidth: 100/100/100
- HitboxStartHalfWidth: 175/175/175
- HoldDistance: 300/300/300
- SpeedBoostAmount: 0.4/0.4/0.4
- SpeedBoostDuration: 1.5/1.5/1.5

## Heimerdinger

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 558 | 105 | 2343 |
| dano de ataque | 56 | 2.7 | 101.9 |
| armadura | 19 | 4.2 | 90.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 1.36 | 23.778 |
| regeneração de vida | 1.4 | 0.11 | 3.27 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Afinidade Hextec**

- MovementSpeed: 0.2

**Q · Torre Evolutiva H-28G** · recarga 1/1/1/1/1 · custo 20/20/20/20/20 · alcance 350/350/350/350/350

- BeamCooldown: 90/90/90/90/90
- Damage: 7/11/15/19/23 + 0.35 ability_power
- DamageBeam: 40/60/80/100/120 + 0.55 ability_power
- KitCooldown: 20/20/20/20/20
- TurretHealth: 640/640/640/640/640 + 0.05 ability_power _(no nível 18)_
- DistanceToHeimer: 900/900/900/900/900
- KitCost: 1/1/1/1/1
- MaxKits: 3/3/3/3/3
- MaxTurrets: 3/3/3/3/3

**W · Micro-Mísseis Hextec** · recarga 11/10/9/8/7 · custo 50/60/70/80/90

- Damage: 50/75/100/125/150 + 0.55 ability_power
- ExtraHitDamage: 10/15/20/25/30 + 0.12 ability_power
- ExtraHitDamageMinions (= ExtraHitDamage × 3): 30/45/60/75/90 + 0.35999998 ability_power
- TotalDamage: 90/135/180/225/270 + 1.03 ability_power
- Rockets: 5/5/5/5/5

**E · Granada de Tempestade de Elétrons CH-2** · recarga 11/11/11/11/11 · custo 85/85/85/85/85 · alcance 925/925/925/925/925

- Damage: 60/100/140/180/220 + 0.6 ability_power
- SlowDuration: 2/2/2/2/2
- SlowPercent: 0.35/0.35/0.35/0.35/0.35
- StunDuration: 1.5/1.5/1.5/1.5/1.5

**R · MELHORIA!!!** · recarga 100/85/70 · custo 100/100/100 · alcance 280/280/280

- EUltDamage: 100/200/300 + 0.6 ability_power
- QUltDamage: 80/100/120 + 0.35 ability_power
- QUltDamageBeam: 100/140/180 + 0.7 ability_power
- QUltTurretHealth: 725/725/725 + 0.5 ability_power
- RQTurretResists: 30/30/30
- WUltDamage: 135/180/225 + 0.45 ability_power
- WUltTotalDamage: 503/697.5/892 + 1.83 ability_power
- {0df0e2e2} (= WUltTotalDamage × 20): 10060/13950/17840 + 36.600002 ability_power
- {a8af07ef}: 32/45/58 + 0.12 ability_power
- {f058914b} (= {a8af07ef} × 0.5): 16/22.5/29 + 0.06 ability_power
- WUltPrimaryAPRatio: 0.45/0.45/0.45
- WUltPrimaryBaseDamage: 135/180/225
- WUltSecondaryAPRatio: 0.12/0.12/0.12
- WUltSecondaryBaseDamage: 28/39/49

## Nasus

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 104 | 2418 |
| dano de ataque | 67 | 4 | 135 |
| armadura | 34 | 4.7 | 113.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.638 | 3.48 | 59.798 |
| regeneração de vida | 1.8 | 0.18 | 4.86 |
| velocidade de movimento | 350 | — | 350 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Devorador de Almas**

- LifestealTooltip: 24 _(no nível 18)_

**Q · Ataque Sifão** · recarga 7.5/6.5/5.5/4.5/3.5 · custo 20/20/20/20/20 · alcance 255/255/255/255/255

- BasicStacks: 4/4/4/4/4
- BigStacks: 10/10/10/10/10
- BonusRange: 50/50/50/50/50
- BuffDuration: 10/10/10/10/10

**W · Murchar** · recarga 15/14/13/12/11 · custo 80/80/80/80/80 · alcance 700/700/700/700/700

- AttackSpeedSlowMult: 0.75/0.75/0.75/0.75/0.75
- Duration: 5/5/5/5/5
- MaxSlowTooltipOnly: 47/59/71/83/95
- SlowBase: 35/35/35/35/35
- SlowPerTick: 3/6/9/12/15

**E · Fogo Espiritual** · recarga 12/12/12/12/12 · custo 60/70/80/90/100 · alcance 650/650/650/650/650

- InitialDamage: 50/80/110/140/170 + 0.6 ability_power
- TotalDotDamage: 10/16/22/28/34 + 0.12 ability_power
- ArmorShredPercent: -0.3/-0.35/-0.4/-0.45/-0.5

**R · Fúria das Areias** · recarga 120/100/80 · custo 100/100/100

- DamageCalc: 0.03/0.04/0.05 + 0.0001 ability_power
- AttackRangeIncrease: 50/50/50
- BaseAoERadius: 376/376/376
- BonusHealth: 300/450/600
- Duration: 15/15/15
- InitialResistGain: 40/55/70
- MaxDamageCap: 240/240/240
- MinAoERadius: 400/400/400
- QCDR: 0.5/0.5/0.5
- SizeIncreasePercent: 0.3/0.35/0.4
- TickRate: 0.5/0.5/0.5

## Nidalee

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 109 | 2463 |
| dano de ataque | 58 | 3.5 | 117.5 |
| armadura | 32 | 5 | 117 |
| resistência mágica | 30 | 1.45 | 54.65 |
| velocidade de ataque | 0.638 | 3.22 | 55.378002 |
| regeneração de vida | 1.2 | 0.12 | 3.24 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**Q · Arremessar Lança / Bote** · recarga 6/6/6/6/6 · custo 50/55/60/65/70 · alcance 1500/1500/1500/1500/1500

- HumanMaximumDamage (= HumanMinimumDamage × 1): 227.5/292.5/357.5/422.5/487.5 + 1.625 ability_power
- HumanMinimumDamage: 70/90/110/130/150 + 0.5 ability_power
- {38f95809}: 5/30/55/80/105
- RangeThreshold: 525/525/525/525/525
- SpearMaximumDamage: 227.5/292.5/357.5/422.5/487.5

**W · Arapuca / Investida** · recarga 13/12/11/10/9 · custo 30/35/40/45/50 · alcance 900/900/900/900/900

- DamagePerSecond: 10/20/30/40/50 + 0.05 ability_power
- MaxTraps: 10/10/10/10/10 _(no nível 18)_
- DotDuration: 4/4/4/4/4
- TotalFlatDamage: 40/80/120/160/200
- TrapLifetime: 120/120/120/120/120

**E · Ímpeto Selvagem / Patada** · recarga 12/12/12/12/12 · custo 50/55/60/65/70 · alcance 900/900/900/900/900

- MaxHealing (= TotalHealing × 1): 100/150/200/250/300 + 0.7 ability_power
- TotalHealing: 50/75/100/125/150 + 0.35 ability_power
- ASDuration: 7/7/7/7/7
- BonusAS: 0.3/0.4/0.5/0.6/0.7
- MaxHealThreshold: 0.05/0.05/0.05/0.05/0.05

**R · Aspecto do Puma** · recarga 3/3/3 · alcance 20/20/20

- TotalPounceDamage: 55/100/145 + 0.3 ability_power + 0.5 attack_damage (bonus)
- TotalSwipeDamage: 70/130/190 + 0.55 ability_power + 0.7 attack_damage (bonus)
- TotalTakedownDamage: 5/30/55 + 0.4 ability_power + 0.75 attack_damage
- PassivePercentMS: 10/10/10
- PounceCooldown: 3/2.5/2
- TakedownDamageAmp: 1/1.25/1.5

## Udyr

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 664 | 92 | 2228 |
| dano de ataque | 62 | 4 | 130 |
| armadura | 31 | 4.7 | 110.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.65 | 3 | 51.65 |
| regeneração de vida | 1.2 | 0.15 | 3.7500002 |
| velocidade de movimento | 350 | — | 350 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ponte Espiritual**

- AttackSpeed: 0.3
- LightningDamage: 0 + 0.2 ability_power + 0.1 attack_damage
- UltCD: 20 _(no nível 18)_
- {01dc7936}: 40 + 0.4 attack_damage _(no nível 18)_
- {1800f4cb}: 0.3 _(no nível 18)_
- {1cfcbfbe}: 0.8 _(no nível 18)_
- {5e1424c3}: 20 + 0.05 max_health
- {6a8aaac9}: 15
- {73fab8fe}: 10 + 0.4 ability_power + 0.3 attack_damage
- AttackSpeedDuration: 4
- GlobalCD: 1.5
- GlobalCDEmpowered: 0.5
- MaxCharge: 1000
- SwapChargeBonus: 1
- UdyrAHAmountPerCloudDragon: 6
- UltCDReduction: 0.05

**Q · Garra Selvagem** · recarga 6/6/6/6/6 · custo 20/20/20/20/20 · alcance 600/600/600/600/600

- LightningDamageToMinionsMin: 160/160/160/160/160 _(no nível 18)_
- MaxHPOnHit1: 0.03/0.04/0.05/0.06/0.07 + 0.00035 attack_damage (bonus)
- MonsterCap: 15/15/15/15/15 + 0.5 ability_power + 1 attack_damage (bonus)
- MonsterNukeMinRCalc (= {af259c48} × 8): 280/280/280/280/280 _(no nível 18)_
- OnHitDamage: 6/12/18/24/30 + 0.2 attack_damage (bonus) + 0.01 max_health (bonus)
- Q2TotalOnHitHPDamage: 0.07/0.08/0.09/0.099999994/0.11 + 0.0005 attack_damage (bonus) + 0.00001 max_health (bonus) _(no nível 18)_
- {af259c48}: 35/35/35/35/35 _(no nível 18)_
- {ba540e22} (= MonsterCap × 1): 15/15/15/15/15 + 0.5 ability_power + 1 attack_damage (bonus)
- {ec11b8c6}: 1/1/1/1/1 _(no nível 18)_
- AttackRange: 50/50/50/50/50
- AttackSpeedDurationBase: 4/4/4/4/4
- AttackSpeedDurationEmpowered: 4/4/4/4/4
- BounceRange: 450/450/450/450/450
- EmpoweredBonusASLevel1: 0.2/0.2/0.2/0.2/0.2
- EmpoweredBonusASLevel18: 0.7/0.7/0.7/0.7/0.7
- LightningDamageLevel1: 0.015/0.015/0.015/0.015/0.015
- LightningDamageLevel18: 0.03/0.03/0.03/0.03/0.03

**W · Manto de Ferro** · recarga 6/6/6/6/6 · custo 40/40/40/40/40

- HealPerSecond: 0/0/0/0/0 + 0.05 max_health
- LifeOnHit: 0/0/0/0/0 + 0.08 ability_power + 0.012 max_health
- LifeOnHitAwakened (= LifeOnHit × 1): 0/0/0/0/0 + 0.16 ability_power + 0.024 max_health
- Omnivamp: 0.25/0.25/0.25/0.25/0.25
- RecastHeal (= RecastShield × 1): 97.5/107.5/117.5/127.5/137.5 + 0.325 ability_power + 0.5 attack_damage (bonus) + 0.04 max_health _(no nível 18)_
- RecastShield: 195/215/235/255/275 + 0.65 ability_power + 1 attack_damage (bonus) + 0.08 max_health _(no nível 18)_
- TotalShield: 45/65/85/105/125 + 0.4 ability_power + 0.5 attack_damage (bonus) + 0.02 max_health
- {5fcd3b84}: 0.3/0.32/0.34/0.36/0.38
- {b038be53} (= {bbb9794b} × 1): 0/0/0/0/0
- {bbb9794b}: 0.2/0.2/0.2/0.2/0.2 _(no nível 18)_
- HealOnHitMinionPenalty: 0.4/0.4/0.4/0.4/0.4
- ShieldDuration: 4/4/4/4/4

**E · Investida Ardente** · recarga 6/6/6/6/6 · custo 40/40/40/40/40 · alcance 600/600/600/600/600

- MoveSpeed: 0.25/0.31/0.37/0.43/0.49 + 0.0005 attack_damage (bonus)
- MoveSpeedBonus: 0.4/0.4/0.4/0.4/0.4 + 0.001 attack_damage (bonus) _(no nível 18)_
- EmpoweredBonusRange: 75/75/75/75/75
- ICD: 6/5.6/5.2/4.8/4.4
- MoveSpeedDuration: 4/4/4/4/4
- StunDuration: 0.75/0.75/0.75/0.75/0.75
- UnstoppableDuration: 1.5/1.5/1.5/1.5/1.5

**R · Tempestade Alada** · recarga 6/6/6 · custo 40/40/40 · alcance 370/370/370

- DamageToMinions_Scaling: 0.8/0.8/0.8 _(no nível 18)_
- EmpoweredSlow: 0.05/0.05/0.05
- MonsterCap: 50/50/50 _(no nível 18)_
- PercentHPBlast: 0.14/0.14/0.14 + 0.00035 ability_power _(no nível 18)_
- PulseDamage: 40/40/40 + 0.35 ability_power _(no nível 18)_
- StormDamage: 20/36/52 + 0.35 ability_power
- {817f8a73}: -100/-100/-100 _(no nível 18)_
- {af259c48}: 35/35/35 _(no nível 18)_
- BuffDuration: 4/4/4
- SlowDuration: 2/2/2
- SlowPotency: 0.15/0.18/0.21

## Poppy

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 110 | 2480 |
| dano de ataque | 56 | 4 | 124 |
| armadura | 35 | 5 | 120 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 2.5 | 43.158 |
| regeneração de vida | 1.8 | 0.16 | 4.52 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Embaixadora de Ferro**

- ActualCooldown: 8 _(no nível 18)_
- ShieldValue: 0 + 0.2 max_health _(no nível 18)_
- TotalDamage: 180 _(no nível 18)_
- BonusRange: 350
- ShieldDuration: 3
- ShieldPickUpTime: 4

**Q · Choque do Martelo** · recarga 8/7/6/5/4 · custo 35/40/45/50/55 · alcance 430/430/430/430/430

- BaseDamage: 30/55/80/105/130 + 0.75 attack_damage (bonus)
- MoveSpeedMod: 0.2/0.23/0.26/0.29/0.32 + 0.00008 max_health (bonus)
- DelayBetweenTwoHits: 1/1/1/1/1
- HealthDamagePercent: 7/7.5/8/8.5/9
- MaxHealthDamageToNonHeroes: 85/120/155/190/225
- SlowDuration: 0.5/0.5/0.5/0.5/0.5

**W · Presença Inabalável** · recarga 20/18/16/14/12 · custo 50/50/50/50/50

- BonusArmor: 0/0/0/0/0 + 0.16 armor
- BonusMR: 0/0/0/0/0 + 0.16 magic_resist
- InterruptDamage: 70/110/150/190/230 + 0.7 ability_power
- Duration: 2/2/2/2/2
- GroundingDuration: 2/2/2/2/2
- Haste: 40/40/40/40/40
- KnockupDuration: 0.5/0.5/0.5/0.5/0.5
- PassiveEmpoweredHealthPercent: 0.4/0.4/0.4/0.4/0.4
- SlowAmount: -0.25/-0.25/-0.25/-0.25/-0.25

**E · Investida Heroica** · recarga 14/13/12/11/10 · custo 70/70/70/70/70 · alcance 475/475/475/475/475

- TackleDamage: 40/60/80/100/120 + 0.6 attack_damage (bonus)
- DashSpeed: 1800/1800/1800/1800/1800
- MaxEnemyPushbackDistance: 400/400/400/400/400
- StunDuration: 1.6/1.7/1.8/1.9/2
- WallDamage: 40/60/80/100/120

**R · Veredito da Guardiã** · recarga 140/120/100 · custo 100/100/100 · alcance 500/500/500

- Damage: 200/300/400 + 0.9 attack_damage (bonus)
- HalfDamage (= Damage × 1): 100/150/200 + 0.45 attack_damage (bonus)
- CancelCDRefund: 30/30/30
- ChannelDistanceUpdateTick: 0.25/0.25/0.25
- ChannelMaxDuration: 4/4/4
- DurationUntilFullyCharged: 1/1/1
- KnockupDurationCharged: 2/2/2
- KnockupDurationSnap: 1/1/1
- MaxKnockbackDistance: 3400/3400/3400
- SelfSlow: 15/15/15

## Gragas

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 115 | 2595 |
| dano de ataque | 64 | 3.5 | 123.5 |
| armadura | 38 | 5 | 123 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.675 | 2.05 | 35.524998 |
| regeneração de vida | 1.1 | 0.1 | 2.8 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Happy Hour** · recarga 12

- HealAmount: 0 + 0.055 max_health
- HealCooldown: 6 _(no nível 18)_
- ArenaCooldownMod: 0.5

**Q · Jogar o Barril** · recarga 10/9/8/7/6 · custo 80/80/80/80/80 · alcance 850/850/850/850/850

- MaxDamage (= MinDamage × 1.5): 120/180/240/300/360 + 1.2 ability_power
- MinDamage: 80/120/160/200/240 + 0.8 ability_power
- BarrelCookMaxAmp: 150/150/150/150/150
- BarrelCookMaxTime: 2/2/2/2/2
- BarrelMaxDuration: 4/4/4/4/4
- MinionDamageMod: 70/70/70/70/70
- SlowDuration: 2/2/2/2/2
- SlowPercent: 40/45/50/55/60

**W · Fúria da Bebedeira** · recarga 5/5/5/5/5 · custo 30/30/30/30/30 · alcance 20/20/20/20/20

- DamageReduction: 10/14/18/22/26 + 0.04 ability_power
- TotalDamage: 20/50/80/110/140 + 0.7 ability_power
- AttackDuration: 5/5/5/5/5
- AttackRangeBuff: 50/50/50/50/50
- DefenseDuration: 2.5/2.5/2.5/2.5/2.5
- MaxHPPercentDamage: 7/7/7/7/7
- MonsterDamageCap: 300/300/300/300/300
- TurretDamageMod: 0.5/0.5/0.5/0.5/0.5

**E · Barrigada** · recarga 14/13.5/13/12.5/12 · custo 50/50/50/50/50

- TotalDamage: 80/125/170/215/260 + 0.6 ability_power
- CooldownRefund: 0.4/0.4/0.4/0.4/0.4
- DashSpeed: 900/900/900/900/900
- ForwardHitbox: 40/40/40/40/40
- HitboxArea: 180/180/180/180/180
- StunDuration: 1/1/1/1/1

**R · Barril Explosivo** · recarga 100/85/70 · custo 100/100/100 · alcance 1050/1050/1050

- DamageDone: 200/300/400 + 0.8 ability_power
- KnockbackDistance: 900/900/900

## Pantheon

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 109 | 2503 |
| dano de ataque | 64 | 3.3 | 120.1 |
| armadura | 40 | 4.95 | 124.149994 |
| resistência mágica | 28 | 2.05 | 62.85 |
| velocidade de ataque | 0.658 | 2.95 | 50.808002 |
| regeneração de vida | 1.2 | 0.13 | 3.4099998 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Determinação Mortal**

- ActionsToEmpower: 5
- CritThreshold: -1
- OutOfCombatTimer: 15
- PhalanxDuration: 3
- QDamageAmp: 0.4
- QDamageAmpMelee: 0.7
- QExtraAttacks: 1

**Q · Lança Meteórica** · recarga 11/10.25/9.5/8.75/8 · custo 25/25/25/25/25

- EmpoweredDamageCalc: 240/240/240/240/240 + 1.15 attack_damage (bonus) _(no nível 18)_
- ExecuteDamageCalcModified: 155/230/305/380/455 + 2.3 attack_damage (bonus)
- HoldDamageCalc: 70/100/130/160/190 + 0.5 ability_power + 1.15 attack_damage (bonus)
- TapDamageCalc: 70/100/130/160/190 + 1.15 attack_damage (bonus)
- {6fc04dc0}: 85/130/175/220/265 + 1.15 attack_damage (bonus)
- CritHealthThreshold: 0.2/0.2/0.2/0.2/0.2
- DamageFalloff: 0.5/0.5/0.5/0.5/0.5
- ExecuteBaseDamage: 155/230/305/380/455
- HoldRange: 1200/1200/1200/1200/1200
- MinTimeHoldCast: 0.35/0.35/0.35/0.35/0.35
- MinionDamageMod: 0.7/0.7/0.7/0.7/0.7
- MonsterDamageMod: 0.8/0.8/0.8/0.8/0.8
- SelfSlow: 0.1/0.1/0.1/0.1/0.1
- TapCooldownRefund: 0.6/0.6/0.6/0.6/0.6

**W · Escudo-Cometa** · recarga 13/12/11/10/9 · custo 55/55/55/55/55 · alcance 600/600/600/600/600

- DamageCalc: 0/0/0/0/0 + 1 ability_power
- EmpoweredDamageMultCalcModified (= {1b016817} × 3): 0/0/0/0/0 + 1.6500001 attack_damage _(no nível 18)_
- MaxHealthDamageCalc: 0.06/0.065/0.07/0.075/0.08 + 0.00015 ability_power + 0.00004 max_health (bonus)
- {1b016817}: 0/0/0/0/0 + 0.55 attack_damage _(no nível 18)_
- BuffDuration: 4/4/4/4/4
- EmpoweredDamageMult: 0.6/0.6/0.6/0.6/0.6
- EmpoweredNumHits: 3/3/3/3/3
- MonsterDamageCap: 150/150/150/150/150
- MonsterDamageMin: 60/60/60/60/60
- StunDuration: 1/1/1/1/1

**E · Égide Impetuosa** · recarga 22/21/20/19/18 · custo 80/80/80/80/80

- DamageCalc: 0/0/0/0/0 + 1 attack_damage
- EmpoweredDamageCalc (= DamageCalc × 1.67): 0/0/0/0/0 + 1.67 attack_damage
- HealCalc: 40/40/40/40/40 + 1.5 attack_damage (bonus)
- ResistsCalc: 30/30/30/30/30 + 0.025 max_health (bonus) _(no nível 18)_
- ShieldDamageCalc: 55/105/155/205/255 + 1.5 attack_damage (bonus)
- {e62bc5e9} (= DamageCalc × 0.167): 0/0/0/0/0 + 0.167 attack_damage
- AttacksPerSecond: 4/4/4/4/4
- MinionDamageReduction: 0.5/0.5/0.5/0.5/0.5
- RecastLockout: 0.3/0.3/0.3/0.3/0.3
- ResistsDuration: 4/4/4/4/4
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5
- ShieldSwipeRadius: 375/375/375/375/375
- SpearStrikesRadius: 525/525/525/525/525
- SpeedAmount: 0.6/0.6/0.6/0.6/0.6
- SpeedDuration: 1.5/1.5/1.5/1.5/1.5

**R · Constelação Cadente** · recarga 180/165/150 · custo 100/100/100 · alcance 5500/5500/5500

- DamageCalc: 300/500/700 + 1 ability_power
- MinDamage (= DamageCalc × 1): 150/250/350 + 0.5 ability_power
- ArmorPenetration: 0.1/0.2/0.3
- CancelCooldown: 30/30/30
- Radius: 450/450/450
- SpearSlow: 0.5/0.5/0.5
- SpearSlowDuration: 2/2/2
- SweetSpotRadius: 125/125/125

## Ezreal

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 102 | 2334 |
| dano de ataque | 60 | 3.75 | 123.75 |
| armadura | 24 | 4.2 | 95.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.625 | 2.5 | 43.125 |
| regeneração de vida | 0.8 | 0.13 | 3.01 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Feitiço do Poder Crescente**

- AttackSpeedPerStack: 0.1
- MaxStacks: 5
- StackDuration: 6

**Q · Disparo Místico** · recarga 5.5/5.25/5/4.75/4.5 · custo 28/31/34/37/40 · alcance 1200/1200/1200/1200/1200

- Damage: 20/45/70/95/120 + 0.4 ability_power + 1.3 attack_damage
- CDRefund: 1.5/1.5/1.5/1.5/1.5

**W · Fluxo Essencial** · recarga 8/8/8/8/8 · custo 50/50/50/50/50 · alcance 1200/1200/1200/1200/1200

- Damage: 80/135/190/245/300 + 0.9 ability_power + 1 attack_damage (bonus)
- DetonationTimeout: 4/4/4/4/4
- ManaReturn: 60/60/60/60/60

**E · Translocação Arcana** · recarga 26/23/20/17/14 · custo 70/70/70/70/70

- Damage: 80/130/180/230/280 + 0.75 ability_power + 0.6 attack_damage (bonus)
- CherryBonusHaste: 15/15/15/15/15
- MissileRange: 750/750/750/750/750

**R · Barragem Incendiária** · recarga 120/105/90 · custo 100/100/100

- Damage: 350/550/750 + 1.1 ability_power + 1 attack_damage (bonus)
- DamageMinionMonster: 150/225/300 + 1.1 ability_power + 1 attack_damage (bonus)
- DamageReductionPerHit: 0.1/0.1/0.1
- MinimumDamagePercent: 0.3/0.3/0.3

## Mordekaiser

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 645 | 104 | 2413 |
| dano de ataque | 61 | 4 | 129 |
| armadura | 37 | 4.2 | 108.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 1 | 17.625 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ascensão das Trevas** · recarga 0

- AuraDamagePerStack: 5 + 0.3 ability_power
- BonusAPAuto: 0 + 0.4 ability_power
- MaxAuraMonsterDamage: 200 _(no nível 18)_
- PassiveMovementSpeed: 0.089999996 _(no nível 18)_
- PercentHealthForAura: 5 _(no nível 18)_
- CombatTrackingDuration: 4
- DoTRadius: 375
- FuryForAutoHit: 1
- FuryForEHit: 1
- FuryForQHit: 1
- MaximumStacks: 3
- MovementSpeed: 0.03
- PercentAPAddedToAutos: 0.4

**Q · Obliterar** · recarga 8/7/6/5/4

- IsolationTooltip: 0.3/0.35/0.4/0.45/0.5
- MaceLength: 625/625/625/625/625
- MaceStartDistance: 400/400/400/400/400
- RectangleWidth: 160/160/160/160/160

**W · Indestrutível** · recarga 12/11/10/9/8

- HealthDecay: 0/0/0/0/0 + 0.005 max_health (base)
- MaxHealthTooltip: 0/0/0/0/0 + 0.3 max_health
- MinHealthTooltip: 0/0/0/0/0 + 0.05 max_health
- BaseShield: 0.05/0.05/0.05/0.05/0.05
- DamageConversion: 0.45/0.45/0.45/0.45/0.45
- DamageTakenConversion: 0.075/0.075/0.075/0.075/0.075
- Duration: 5/5/5/5/5
- HealingPercent: 0.35/0.375/0.4/0.425/0.45
- MaxHealthCap: 0.3/0.3/0.3/0.3/0.3
- MinionPenalty: 0.25/0.25/0.25/0.25/0.25
- TimeBeforeDecay: 1/1/1/1/1

**E · Aperto Mortal** · recarga 16/14/12/10/8 · alcance 700/700/700/700/700

- TotalDamage: 60/80/100/120/140 + 0.45 ability_power
- DelayBeforeMovement: 0.5/0.5/0.5/0.5/0.5
- KnockTowardsDistance: 250/250/250/250/250
- MagicPen: 0.05/0.075/0.1/0.125/0.15
- MaxDistance: 550/550/550/550/550
- StandAsideRange: 900/900/900/900/900

**R · Reino da Morte** · recarga 140/120/100 · alcance 650/650/650

- GhostAPRatio: 0.6/0.6/0.6
- SpiritRealmDuration: 7/7/7
- StatStealPercentScalar: 0.1/0.1/0.1
- ZoneRadius: 1200/1200/1200

## Yorick

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 114 | 2588 |
| dano de ataque | 62 | 5 | 147 |
| armadura | 36 | 4.5 | 112.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2 | 34.625 |
| regeneração de vida | 1.6 | 0.16 | 4.32 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Pastor de Almas**

- AOEDamageModifier: 0.66
- BonusAttackSpeed: 0.8 _(no nível 18)_
- YorickPassiveGhoulDamage: 100 + 0.2 attack_damage (bonus) _(no nível 18)_
- YorickPassiveGhoulHealth: 300 + 0.15 max_health (bonus) _(no nível 18)_
- YorickPassiveSpawnThreshold: 2 _(no nível 18)_
- ARAMPassiveSpawnThreshold: 5
- CherryGhoulHPMod: 0.7
- GhoulMonsterDamageMod: 0.6
- GraveDuration: 300
- MeleeAttackDamageModifier: 2
- MinionDamageTakenMod: 0.6
- MinionMoveBonus: 40
- MonsterDamageTakenMod: 1
- ProcDamageModifier: 0.5
- YorickGhoulDamageReductionVsTowers: 0.5
- YorickGhoulLeashRange: 1600
- YorickGhoulSpawnDelay: 1
- YorickPassiveGhoulMax: 4

**Q · Extrema-Unção** · recarga 6/5.5/5/4.5/4 · custo 20/20/20/20/20

- BonusDamage: 30/50/70/90/110 + 0.5 attack_damage
- QHeal: 10/10/10/10/10
- BuffDuration: 5/5/5/5/5
- CherryGhoulPassiveSpawnRate: 8/8/8/8/8
- HealReduction: 50/50/50/50/50
- MissingHealthRatio: 6/7/8/9/10

**W · Procissão Sombria** · recarga 20/18/16/14/12 · custo 70/70/70/70/70 · alcance 600/600/600/600/600

- CircleDuration: 4/4/4/4/4
- Delay: 0.75/0.75/0.75/0.75/0.75
- GhoulNumber: 18/18/18/18/18
- HPBonus: 0/0/1/1/2
- WallHealthTooltip: 2/2/3/3/4

**E · Névoa dos Lamentos** · recarga 12/11/10/9/8 · custo 50/55/60/65/70 · alcance 700/700/700/700/700

- Calc_HealthDamage: 6/6.5/7/7.5/8 + 0.03 ability_power
- Calc_MinimumDamage: 70/105/140/175/210 + 1 ability_power
- Calc_MonsterCapDamage: 50/75/100/125/150 + 1 ability_power
- Calc_Slow: 0.3/0.3/0.3/0.3/0.3
- ArmorShred: 0.13/0.16/0.19/0.22/0.25
- HasteAmount: 0.18/0.21/0.24/0.27/0.3
- MarkDuration: 4/4/4/4/4
- MarkRange: 1500/1500/1500/1500/1500
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**R · Louvor das Ilhas** · recarga 160/130/100 · custo 100/100/100 · alcance 600/600/600

- YorickBigGhoulDamage: 50/75/100 + 0.3 attack_damage (bonus)
- YorickBigGhoulHealth: 1050/1050/1050 + 0.6 max_health (bonus)
- YorickMaidenResists: 30/30/30
- MinionDamageTakenModifier: 0.3/0.3/0.3
- RControlTime: 25000/25000/25000
- RGhoulNumbers: 2/3/4
- RMarkDamagePercent: 2/2.5/3
- RMarkMaxDamage: 30/30/30
- RRadius: 700/700/700
- RSummonDelay: 1/1/1

## Akali

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 119 | 2623 |
| dano de ataque | 62 | 3.3 | 118.1 |
| armadura | 23 | 4.7 | 102.899994 |
| resistência mágica | 37 | 2.05 | 71.85 |
| velocidade de ataque | 0.625 | 3.2 | 55.025 |
| regeneração de vida | 1.8 | 0.18 | 4.86 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Marca do Assassino**

- Damage: 35 + 0.55 ability_power + 0.6 attack_damage (bonus)
- Level1MS: 0.3
- MSAdditionalBonusAtThreshold: 0.1

**Q · Golpe dos Cinco Pontos** · recarga 1.5/1.5/1.5/1.5/1.5 · custo 110/100/90/80/70 · alcance 550/550/550/550/550

- Damage: 45/70/95/120/145 + 0.6 ability_power + 0.65 attack_damage
- MinionDamage (= Damage × 1): 45/70/95/120/145 + 0.6 ability_power + 0.65 attack_damage
- SlowDuration: 0.5/0.5/0.5/0.5/0.5
- SlowPercentage: 0.5/0.5/0.5/0.5/0.5

**W · Proteção do Crepúsculo** · recarga 20/19/18/17/16 · alcance 350/350/350/350/350

- {955fbaa3}: 0.625/0.625/0.625/0.625/0.625 _(no nível 18)_
- BaseDuration: 5/5.5/6/6.5/7
- CherryBonusHaste: 30/30/30/30/30
- CloudRadius: 140/140/140/140/140
- EnergyRestore: 100/100/100/100/100
- MaxDuration: 7/7.5/8/8.5/9
- MovementSpeed: 30/35/40/45/50
- MovementSpeedDuration: 2/2/2/2/2

**E · Investida Shuriken** · recarga 16/14.5/13/11.5/10 · custo 30/30/30/30/30 · alcance 825/825/825/825/825

- E1Damage: 70/140/210/280/350 + 1.1 ability_power + 1 attack_damage
- E2DamageCalc: 70/140/210/280/350 + 1.1 ability_power + 1 attack_damage
- SelfKnockbackRange: 400/400/400/400/400

**R · Execução Perfeita** · recarga 120/90/60 · alcance 675/675/675

- Cast1Damage: 110/220/330 + 0.3 ability_power + 0.5 attack_damage (bonus)
- Cast2DamageMax (= Cast2DamageMin × 3): 210/420/630 + 0.90000004 ability_power
- Cast2DamageMin: 70/140/210 + 0.3 ability_power
- CooldownBetweenCasts: 2.5/2.5/2.5
- DashDistance: 715/715/715
- MaxExecuteThreshold: 0.3/0.3/0.3

## Kennen

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 580 | 98 | 2246 |
| dano de ataque | 48 | 3.75 | 111.75 |
| armadura | 29 | 4.95 | 113.149994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 3.4 | 58.425003 |
| regeneração de vida | 1.1 | 0.13 | 3.31 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Marca da Tormenta**

- DiminishingReturnDuration: 6
- EnergyRestore: 25
- MarkDuration: 6
- ReducedStunDuration: 0.5
- StunDuration: 1.25

**Q · Shuriken Trovejante** · recarga 7/6.25/5.5/4.75/4 · custo 60/55/50/45/40 · alcance 1050/1050/1050/1050/1050

- TotalDamage: 75/125/175/225/275 + 0.75 ability_power

**W · Surto Elétrico** · recarga 13/11.25/9.5/7.75/6 · custo 40/40/40/40/40 · alcance 725/725/725/725/725

- TotalDamageActive: 70/95/120/145/170 + 0.8 ability_power
- TotalDamagePassive: 35/45/55/65/75 + 0.35 ability_power + 0.8 attack_damage (bonus)

**E · Investida Relâmpago** · recarga 10/9/8/7/6 · custo 80/80/80/80/80 · alcance 200/200/200/200/200

- TotalDamage: 80/120/160/200/240 + 0.8 ability_power
- CritDurationBonus: 1/1/1/1/1
- DamageRadius: 200/200/200/200/200
- DamageToMinions: 0.65/0.65/0.65/0.65/0.65
- DurationAfterBall: 4/4/4/4/4
- DurationAsBall: 2/2/2/2/2
- EnergyRefund: 40/40/40/40/40
- MovementSpeed: 1/1/1/1/1
- TotalAS: 0.4/0.5/0.6/0.7/0.8
- linger: 0.15/0.15/0.15/0.15/0.15

**R · Turbilhão Cortante** · recarga 120/120/120 · alcance 550/550/550

- PerTickDamageCalculated: 40/80/120 + 0.25 ability_power
- DamageAmp: 0.1/0.1/0.1
- KennenRDefenses: 25/50/75
- KennenRDuration: 3/3/3
- KennenRTickRate: 0.5/0.5/0.5

## Garen

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 690 | 98 | 2356 |
| dano de ataque | 69 | 4.5 | 145.5 |
| armadura | 38 | 4.2 | 109.399994 |
| resistência mágica | 32 | 1.55 | 58.35 |
| velocidade de ataque | 0.625 | 3.65 | 62.675003 |
| regeneração de vida | 1.6 | 0.1 | 3.3 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Perseverança**

- RegenCalc: 1.5
- DamageTimer: 8

**Q · Acerto Decisivo** · recarga 8/8/8/8/8 · alcance 300/300/300/300/300

- TotalDamage: 30/60/90/120/150 + 1.5 attack_damage
- AttackWindow: 4.5/4.5/4.5/4.5/4.5
- MovementSpeedAmount: 0.35/0.35/0.35/0.35/0.35
- MovementSpeedDuration: 1.4/1.95/2.5/3.05/3.6
- SilenceDuration: 1.5/1.5/1.5/1.5/1.5

**W · Coragem** · recarga 22/19.5/17/14.5/12

- TotalShield: 65/85/105/125/145 + 0.18 max_health (bonus)
- BuffCounterPerKill: 1/1/1/1/1
- DRDuration: 4/4/4/4/4
- DRPercent: 0.25/0.29/0.33/0.37/0.41
- EpicMonsterStacks: 1/1/1/1/1
- LargeMonsterStacks: 1/1/1/1/1
- ResistGainOnKillTooltip: 0.2/0.2/0.2/0.2/0.2
- ResistMax: 30/30/30/30/30
- UpfrontDuration: 0.75/0.75/0.75/0.75/0.75
- UpfrontTenacity: 0.6/0.6/0.6/0.6/0.6

**E · Julgamento** · recarga 9/8.25/7.5/6.75/6 · alcance 660/660/660/660/660

- NumberOfStrikes: 7/7/7/7/7
- TotalDamage: 4/7/10/13/16 + 0.4 attack_damage
- ASPerTick: 0.25/0.25/0.25/0.25/0.25
- Duration: 3/3/3/3/3
- NearestEnemyBonus: 0.25/0.25/0.25/0.25/0.25
- ShredAmount: 0.25/0.25/0.25/0.25/0.25
- ShredDuration: 6/6/6/6/6
- StacksToShred: 6/6/6/6/6

**R · Justiça Demaciana** · recarga 120/100/80

- BaseDamage: 125/200/275
- ExecuteDamage: 0.25/0.3/0.35
- RevealDuration: 1/1/1

## Leona

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 646 | 101 | 2363 |
| dano de ataque | 60 | 3 | 111 |
| armadura | 43 | 4.7999997 | 124.6 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2.9 | 49.925003 |
| regeneração de vida | 1.7 | 0.17 | 4.59 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Luz do Sol**

- TotalDamage: 151 _(no nível 18)_
- MarkDuration: 2.5

**Q · Proteção da Aurora** · recarga 5/5/5/5/5 · custo 30/35/40/45/50 · alcance 100/100/100/100/100

- TotalDamageTooltip: 10/35/60/85/110 + 0.3 ability_power
- StunDuration: 1/1/1/1/1

**W · Eclipse** · recarga 14/13/12/11/10 · custo 60/60/60/60/60

- BonusArmorTooltip: 20/27.5/35/42.5/50 + 0.2 armor (bonus)
- BonusMRTooltip: 20/27.5/35/42.5/50 + 0.2 magic_resist (bonus)
- TotalDamageTooltip: 55/85/115/145/175 + 0.4 ability_power
- ArmorMRDuration: 3/3/3/3/3
- FlatDamageReduction: 8/12/16/20/24
- FlatDamageReductionMax: 0.5/0.5/0.5/0.5/0.5
- MRBaseBonus: 20/27.5/35/42.5/50

**E · Lâmina Zênite** · recarga 12/10.5/9/7.5/6 · custo 40/45/50/55/60

- TotalDamageTooltip: 50/90/130/170/210 + 0.4 ability_power
- RootDuration: 0.5/0.5/0.5/0.5/0.5

**R · Labareda Solar** · recarga 90/75/60 · custo 100/100/100 · alcance 1200/1200/1200

- ExplosionCalculatedDamage: 150/225/300 + 0.8 ability_power
- {f4c90e5c}: 30/40/50 + 0.15 ability_power
- AttackRangeIncrease: 100/100/100
- CCDuration: 1.75/1.75/1.75
- OnHitDuration: 5/5/5
- OnHitStacks: 3/4/5
- SlowPercent: 80/80/80

## Malzahar

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 580 | 101 | 2297 |
| dano de ataque | 55 | 3 | 106 |
| armadura | 18 | 4.7 | 97.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 1.5 | 26.125 |
| regeneração de vida | 1.2 | 0.12 | 3.24 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Oscilação do Vazio**

- ShieldCooldown: 12 _(no nível 18)_
- DRPercent: 90
- LingerDuration: 0.25

**Q · Chamado do Vazio** · recarga 6/6/6/6/6 · custo 60/65/70/75/80 · alcance 900/900/900/900/900

- TotalDamageTooltip: 70/105/140/175/210 + 0.55 ability_power
- DelayPostCast: 0.4/0.4/0.4/0.4/0.4
- SilenceDuration: 1/1.25/1.5/1.75/2

**W · Enxame do Vazio** · recarga 8/8/8/8/8 · custo 40/45/50/55/60

- VoidlingBonusDamageTooltip: 76.5/78.5/80.5/82.5/84.5 + 0.2 ability_power + 0.4 attack_damage (bonus) _(no nível 18)_
- {c2505620}: 64.5/64.5/64.5/64.5/64.5 _(no nível 18)_
- EpicMonsterMod: 0.5/0.5/0.5/0.5/0.5
- LaneMinionMod: 3/3/3/3/3
- MaxStacks: 2/2/2/2/2
- StackCap: 2/2/2/2/2
- StackDuration: 25000/25000/25000/25000/25000
- SummonDelay: 0.5/0.5/0.5/0.5/0.5
- VoidlingDuration: 8/8/9/9/10

**E · Visões Maléficas** · recarga 11/10/9/8/7 · custo 60/70/80/90/100 · alcance 650/650/650/650/650

- ManaRestore: 0/0/0/0/0 + 0.02 mana
- MinionExecuteThreshold: 10/10/10/10/10
- TotalDamage: 80/115/150/185/220 + 0.8 ability_power
- Duration: 4/4/4/4/4
- SecondsPerTick: 0.25/0.25/0.25/0.25/0.25

**R · Aperto Ínfero** · recarga 140/110/80 · custo 100/100/100 · alcance 700/700/700

- TotalDamageTooltip: 125/200/275 + 0.8 ability_power
- ZoneDamageTooltip: 2/3/4 + 0.005 ability_power
- BeamDamageTicks: 10/10/10
- BeamTetherRange: 1250/1250/1250
- CCDuration: 2.5/2.5/2.5
- NeutralMonsterDamageCap: 120/120/120
- PoolDuration: 5/5/5
- SuppressDuration: 2.5/2.5/2.5
- ZoneDuration: 5/5/5

## Talon

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 658 | 109 | 2511 |
| dano de ataque | 68 | 3.1 | 120.7 |
| armadura | 30 | 4.7 | 109.899994 |
| resistência mágica | 36 | 2.05 | 70.85 |
| velocidade de ataque | 0.625 | 2.9 | 49.925003 |
| regeneração de vida | 1.7 | 0.15 | 4.25 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Limiar da Lâmina** · recarga 0

- BleedDamage: 280 + 2.1 attack_damage (bonus) _(no nível 18)_
- BleedDuration: 2
- MonsterMod: 1
- StackDuration: 6

**Q · Diplomacia Noxiana** · recarga 8/7.5/7/6.5/6 · custo 40/40/40/40/40 · alcance 575/575/575/575/575

- LeapDamage: 65/85/105/125/145 + 1 attack_damage (bonus)
- TotalHealing: 55/55/55/55/55 _(no nível 18)_
- CooldownRefund: 0.5/0.5/0.5/0.5/0.5

**W · Ancinho** · recarga 9/8.5/8/7.5/7 · custo 50/55/60/65/70 · alcance 650/650/650/650/650

- TotalInitialDamage: 50/60/70/80/90 + 0.4 attack_damage (bonus)
- TotalReturnDamage: 60/90/120/150/180 + 0.9 attack_damage (bonus)
- MonsterDamageMod: 1/1/1/1/1
- MovespeedSlow: 0.4/0.45/0.5/0.55/0.6
- ReturnDelay: 0.7/0.7/0.7/0.7/0.7
- SlowDuration: 1/1/1/1/1

**E · Caminho do Assassino** · recarga 0/0/0/0/0 · alcance 6000/6000/6000/6000/6000

- DisplayCD: 2/2/2/2/2
- WallCD: 160/135/110/85/60
- WallJumpDistance: 625/625/625/625/625
- WallLockoutWidth: 1250/1250/1250/1250/1250

**R · Ataque das Sombras** · recarga 100/80/60 · custo 100/100/100 · alcance 550/550/550

- Damage: 90/135/180 + 1 attack_damage (bonus)
- Duration: 2.5/2.5/2.5
- MoveSpeed: 0.4/0.55/0.7

## Riven

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 100 | 2330 |
| dano de ataque | 64 | 3 | 115 |
| armadura | 33 | 4.4 | 107.8 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 3.5 | 60.125 |
| regeneração de vida | 1.7 | 0.1 | 3.4 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Lâmina Rúnica**

- TotalDamage: 0 + 0.45 attack_damage _(no nível 18)_
- Charges: 3
- MonsterBonusDamage: 50
- TowerRatio: 0.5

**Q · Asas Quebradas** · recarga 13/13/13/13/13

- FirstSlashDamage: 45/75/105/135/165 + 0.6 attack_damage (bonus)

**W · Explosão de Ki** · recarga 11/10/9/8/7 · alcance 650/650/650/650/650

- TotalDamage: 65/95/125/155/185 + 1 attack_damage (bonus)
- StunDuration: 0.75/0.75/0.75/0.75/0.75

**E · Valentia** · recarga 10/9/8/7/6

- TotalShield: 70/95/120/145/170 + 1.1 attack_damage (bonus)
- BaseShieldAmount: 70/95/120/145/170
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5

**R · Lâmina do Exílio** · recarga 120/90/60 · alcance 200/200/200

- BonusAD: 0/0/0 + 0.2 attack_damage
- MaxDamage: 300/450/600 + 1.65 attack_damage (bonus)
- MinDamage: 100/150/200 + 0.55 attack_damage (bonus)
- Duration: 15/15/15
- TooltipAttackRange: 75/75/75

## Kog'Maw

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 635 | 99 | 2318 |
| dano de ataque | 61 | 3.1 | 113.7 |
| armadura | 24 | 4.45 | 99.649994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.665 | 2.65 | 45.715 |
| regeneração de vida | 0.75 | 0.11 | 2.62 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Surpresa Icathiana**

- PassiveDamage: 650 _(no nível 18)_
- TooltipPassiveDuration: 4
- TooltipPassiveMS: 0.5

**Q · Cusparada Cáustica** · recarga 7/7/7/7/7 · custo 40/40/40/40/40 · alcance 1200/1200/1200/1200/1200

- TotalDamage: 80/125/170/215/260 + 0.9 ability_power
- AttackSpeed: 0.05/0.1/0.15/0.2/0.25
- ShredAmount: 16/20/24/28/32
- ShredDuration: 4/4/4/4/4

**W · Barragem Bio-Arcana** · recarga 17/17/17/17/17 · custo 40/40/40/40/40 · alcance 530/530/530/530/530

- TotalHealthDamage: 3/3.75/4.5/5.25/6 + 0.015 ability_power
- Duration: 8/8/8/8/8
- MonsterDamageCap: 100/100/100/100/100
- Range: 130/150/170/190/210

**E · Gosma do Vazio** · recarga 12/12/12/12/12 · custo 40/55/70/85/100

- TotalDamage: 70/110/150/190/230 + 0.65 ability_power
- SlowAmount: 40/45/50/55/60
- SlowDuration: 0.25/0.25/0.25/0.25/0.25
- TrailDuration: 3/3/3/3/3

**R · Artilharia Viva** · recarga 2/1.5/1 · custo 40/40/40 · alcance 1300/1550/1800

- BaseDamageCalc: 100/140/180 + 0.35 ability_power + 0.75 attack_damage (bonus)
- MaxDamageCalc (= BaseDamageCalc × 1): 200/280/360 + 0.7 ability_power + 1.5 attack_damage (bonus)
- MidDamageCalc (= BaseDamageCalc × 1): 150/210/270 + 0.525 ability_power + 1.125 attack_damage (bonus)
- BaseCost: 40/40/40
- ManaCostCap: 400/400/400
- ManaCostDuration: 8/8/8
- MaxExtraCost: 360/360/360
- Range: 130/155/180
- TooltipMissingHealthDamageAmp: 0.83333/0.83333/0.83333
- VisionDebuffDuration: 2/2/2
- VisionRadius: 400/400/400

## Shen

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 99 | 2293 |
| dano de ataque | 64 | 3 | 115 |
| armadura | 34 | 4.2 | 105.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.751 | 3 | 51.751 |
| regeneração de vida | 1.7 | 0.15 | 4.25 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Barreira de Ki** · recarga 11

- ShieldCooldownReduction: 4
- ShieldValue: 120 + 0.13 max_health (bonus) _(no nível 18)_
- {58a09e24}: 120 + 0.13 max_health (bonus) _(no nível 18)_
- {7e912bc4}: 120 + 0.13 max_health (bonus) _(no nível 18)_
- {87290189}: 101 + 0.07 max_health (bonus) _(no nível 18)_
- {a7305d7f}: 101 + 0.07 max_health (bonus) _(no nível 18)_
- {ff986284}: 101 + 0.07 max_health (bonus) _(no nível 18)_
- DirectionIndicatorMinimumDistance: 150
- LeashDistanceThreshold: 3500
- LeashPreferredResultDistance: 2250
- PassiveCDFlexAmount: 0.125
- ShieldCooldown: 11
- ShieldDuration: 2.5

**Q · Ataque Crepúsculo** · recarga 8/7.25/6.5/5.75/5 · custo 140/130/120/110/100

- BaseFlatDamage: 40/40/40/40/40 _(no nível 18)_
- BasePercentHealth: 2/2.5/3/3.5/4 + 0.015 ability_power
- EmpPercentHealth: 5/5.5/6/6.5/7 + 0.02 ability_power
- TowerDamageCalc (= BaseFlatDamage × 1): 40/40/40/40/40 _(no nível 18)_
- AttackBuffDuration: 8/8/8/8/8
- MinionDamageCap: 100/125/150/175/200
- MonsterAmp: 1/1/1/1/1
- NumEnhancedAttacks: 3/3/3/3/3
- QBaseOnHitDamage: 10/10/10/10/10
- QCollisionRadiusEnd: 130/130/130/130/130
- QCollisionRadiusStart: 110/110/110/110/110
- QPerLevelOnHitDamage: 6/6/6/6/6
- SlowDuration: 2/2/2/2/2
- SlowPercent: 25/30/35/40/45
- SteroidAS: 50/50/50/50/50
- SteroidDuration: 75/75/75/75/75

**W · Refúgio Espiritual** · recarga 16/14.5/13/11.5/10 · custo 40/40/40/40/40

- ZoneDelay: 2/2/2/2/2
- ZoneDuration: 1.75/1.75/1.75/1.75/1.75

**E · Corrida das Sombras** · recarga 18/16/14/12/10 · custo 150/150/150/150/150

- TauntDamage: 60/85/110/135/160 + 0.11 max_health (bonus)
- energyrefund: 50/50/50/50/50 _(no nível 18)_
- ArrivalCollisionRadius: 150/150/150/150/150
- CCDuration: 1.5/1.5/1.5/1.5/1.5
- CollisionPadding: 25/25/25/25/25
- DashBonusSpeed: 800/800/800/800/800
- DepartureCollisionRadius: 125/125/125/125/125
- MinimumDistance: 300/300/300/300/300

**R · Manter a União** · recarga 200/180/160

- MaxShield (= Shield × 1): 192/352/512 + 2.16 ability_power + 0.24000001 max_health (bonus)
- Shield: 120/220/320 + 1.35 ability_power + 0.15 max_health (bonus)
- ArrivalCheckRadiusInvisible: 600/600/600
- ArrivalCheckRadiusVisible: 2000/2000/2000
- ArrivalPersonalSpace: 175/175/175
- ChannelDuration: 3/3/3
- MaxShieldThreshold: 0.4/0.4/0.4
- ShieldDuration: 5/5/5

## Lux

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 580 | 99 | 2263 |
| dano de ataque | 54 | 3.3 | 110.1 |
| armadura | 21 | 5.2 | 109.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.669 | 3 | 51.669 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Iluminação** · recarga 0

- TotalDamage: 200 + 0.35 ability_power _(no nível 18)_
- DebuffDuration: 6

**Q · Ligação da Luz** · custo 50/50/50/50/50

- TotalDamageTT: 80/120/160/200/240 + 0.75 ability_power
- RootDuration: 2/2/2/2/2

**W · Barreira Prismática** · recarga 12/11.5/11/10.5/10 · custo 60/65/70/75/80

- TotalShieldTT: 40/55/70/85/100 + 0.4 ability_power
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5

**E · Singularidade Lucente** · recarga 10/9.5/9/8.5/8 · custo 70/80/90/100/110 · alcance 1100/1100/1100/1100/1100

- TotalDamageTT: 65/115/165/215/265 + 0.8 ability_power
- {ee9714e4}: 25/30/35/40/45
- SlowLingerDuration: 1/1/1/1/1
- SlowZoneDuration: 5/5/5/5/5
- VisionRadius: 650/650/650/650/650

**R · Centelha Final** · recarga 60/50/40 · custo 100/100/100

- TotalDamage: 300/400/500 + 1.2 ability_power
- ResetAssistWindow: 1.75/1.75/1.75
- ResetPercent: 0.3/0.4/0.5

## Xerath

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 596 | 106 | 2398 |
| dano de ataque | 55 | 3 | 106 |
| armadura | 22 | 4.7 | 101.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 1.36 | 23.778 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Oscilação de Mana**

- ChampionManaRestoreTT (= MinionManaRestoreTT × 2): 60
- MinionManaRestoreTT: 30
- CooldownKillRefund: 3.5

**Q · Pulso Arcano** · recarga 9/8/7/6/5 · custo 80/90/100/110/120 · alcance 750/750/750/750/750

- TooltipTotalDamage: 75/115/155/195/235 + 0.9 ability_power
- ChargeTime: 4/4/4/4/4
- ManaRefundFail: 0.5/0.5/0.5/0.5/0.5
- RampingSlow: 0.1/0.1/0.1/0.1/0.1
- RangeGrowthMult: 1.5/1.5/1.5/1.5/1.5
- RectangleWidth: 145/145/145/145/145
- StartRangePercent: 0.5/0.5/0.5/0.5/0.5
- StartingSelfSlow: -0.2/-0.2/-0.2/-0.2/-0.2

**W · Olho da Destruição** · recarga 14/13/12/11/10 · custo 80/90/100/110/120 · alcance 1000/1000/1000/1000/1000

- SweetSpotTotalDamage (= TotalDamage × 1): 83.350006/141.695/200.04001/258.385/316.73 + 1.08355 ability_power
- TotalDamage: 50/85/120/155/190 + 0.65 ability_power
- DamageDelay: 0.5/0.5/0.5/0.5/0.5
- MultiplicativeSlowDecayRate: 0.8/0.8/0.8/0.8/0.8
- SlowAmount: 0.25/0.25/0.25/0.25/0.25
- SlowDuration: 2.5/2.5/2.5/2.5/2.5
- SweetSpotSlowAmount: 0.6/0.65/0.7/0.75/0.8

**E · Orbe Eletrizante** · recarga 13/12.5/12/11.5/11 · custo 60/65/70/75/80

- TooltipTotalDamage: 70/100/130/160/190 + 0.45 ability_power
- MaxStunDuration: 2.25/2.25/2.25/2.25/2.25
- MinStunDuration: 0.75/0.75/0.75/0.75/0.75
- Range: 1125/1125/1125/1125/1125
- StunScaling_RoughAverageWillBeThisNumber_Every100RangeForTheAdditionalStun_AddedTo_5_: 0.17/0.17/0.17/0.17/0.17

**R · Ritual Arcano** · recarga 130/115/100 · custo 100/100/100 · alcance 5000/5000/5000

- RampDamageCalc: 20/25/30 + 0.05 ability_power
- TooltipTotalDamage: 170/220/270 + 0.45 ability_power
- AoESize: 200/200/200
- CDPerShot: 0.6/0.6/0.6
- Duration: 10/10/10
- FailCastRefund: 0.5/0.5/0.5
- NumberOfShots: 4/5/6

## Shyvana

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 625 | 95 | 2240 |
| dano de ataque | 62 | 4 | 130 |
| armadura | 35 | 4 | 103 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.638 | 2 | 34.638 |
| regeneração de vida | 1.4 | 0.13 | 3.61 |
| velocidade de movimento | 350 | — | 350 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Armadura de Escamas**

- PassiveScale: 0.3
- Stacks_Per_Buff_Monster: 1
- Stacks_Per_Champion: 3
- Stacks_Per_Epic_Monster: 3
- Stacks_Per_Large_Monster: 1

**Q · Golpe de Brasas** · recarga 10/9/8/7/6

- Calc_Damage: 10/15/20/25/30 + 0.3 ability_power + 1.1 attack_damage
- Calc_Dragon_Form_Damage (= Calc_Damage × 1.5): 15/22.5/30/37.5/45 + 0.45000002 ability_power + 1.6500001 attack_damage
- Calc_Max_Health_Damage: 0.01/0.01/0.01/0.01/0.01 + 0.00011 attack_damage (bonus)
- Calc_Max_Health_Monster_Maximum: 200/200/200/200/200
- Calc_Max_Health_Monster_Minimum: 10/10/10/10/10
- Calc_Monster_Bonus: 10/10/10/10/10
- {aee2a1e8}: 1/1/1/1/1
- Bonus_Attack_Range: 50/50/50/50/50
- Cooldown_Reduction: 0.5/0.5/0.5/0.5/0.5
- EnhanceDuration: 6/6/6/6/6
- LockoutDuration: 1.5/1.5/1.5/1.5/1.5
- Q1Length: 450/450/450/450/450
- Q2Radius: 275/275/275/275/275
- RecastDuration: 4/4/4/4/4

**W · Égide Infernal** · recarga 13/12.25/11.5/10.75/10 · alcance 350/350/350/350/350

- Calc_Base_Heal: 100/100/100/100/100 _(no nível 18)_
- Calc_Heal_Max (= Calc_Heal_Min × 1): 400/400/400/400/400 + 0.1 ability_power + 0.2 attack_damage (bonus) _(no nível 18)_
- Calc_Heal_Min: 200/200/200/200/200 + 0.05 ability_power + 0.1 attack_damage (bonus) _(no nível 18)_
- Calc_Missing_Health_Heal: 0.08/0.08/0.08/0.08/0.08 _(no nível 18)_
- Calc_Shield: 60/80/100/120/140 + 0.12 max_health (bonus)
- Calc_Shield_Per_Nearby_Champion (= Calc_Shield × 0.3): 18/24/30.000002/36/42 + 0.036000002 max_health (bonus)
- Damage: 80/100/120/140/160 + 0.65 ability_power
- MoveSpeed: 0.25/0.25/0.25/0.25/0.25
- MoveSpeedTowardsEnemies (= MoveSpeed × 1): 0.4375/0.4375/0.4375/0.4375/0.4375
- Duration: 2.5/2.5/2.5/2.5/2.5
- Radius: 350/350/350/350/350
- Shield_Radius: 600/600/600/600/600

**E · Explosão Incandescente** · recarga 12/11/10/9/8 · alcance 800/800/800/800/800

- Calc_Dragon_Damage (= Damage × 1.25): 62.5/81.25/100/118.75/137.5 + 0.75 ability_power
- Calc_Max_Health_Damage: 0.05/0.05/0.05/0.05/0.05
- Calc_Max_Health_Dragon_Damage (= Calc_Max_Health_Damage × 1): 0.0625/0.0625/0.0625/0.0625/0.0625
- Calc_Max_Health_Monster_Maximum: 200/200/200/200/200
- Calc_Max_Health_Monster_Minimum: 0/0/0/0/0
- Calc_Multihit_Efficacy: 0.4/0.4/0.4/0.4/0.4
- Calc_Slow: 0.3/0.3/0.3/0.3/0.3
- Calc_Slow_Dragon (= Calc_Slow × 1): 0.3/0.3/0.3/0.3/0.3
- Damage: 50/65/80/95/110 + 0.6 ability_power
- DamagePerSecond: 25/25/25/25/25 + 0.05 ability_power _(no nível 18)_
- BaseDamagePerSecond: 10/20/30/40/50
- DragonCastRangeScale: 0.4/0.4/0.4/0.4/0.4
- DragonGroundRadius: 300/300/300/300/300
- DragonTrailRadius: 100/100/100/100/100
- GroundLingerDuration: 2/2/2/2/2
- GroundRadius: 150/150/150/150/150
- RepeatExplosionDelay: 0.25/0.25/0.25/0.25/0.25
- SlowDuration: 2/2/2/2/2
- TrailRadius: 75/75/75/75/75

**R · Descida do Dragão** · recarga 0/0/0 · alcance 1050/1050/1050

- Calc_Bonus_Health: 150/250/350
- Damage: 150/250/350 + 1 ability_power
- TT_Fury_AoE_Penalty: 0.75/0.75/0.75
- TT_Fury_Mult: 2/2/2
- {57d26a55}: 1050/1050/1050
- BonusAttackRange: 25/50/75
- FearDuration: 0.75/0.75/0.75
- FuryDrainPerSecond: 6.66666/6.66666/6.66666
- FuryGain: 0.5/0.75/1
- FuryGainPerUltHaste: 0.01/0.01/0.01
- Fury_Generation: 1.25/1.25/1.25
- JumpDampening: 70/70/70
- JumpDelay: 0.15/0.15/0.15
- JumpForwardMultiplier: 100/100/100
- JumpGravity: 25/25/25
- MinCastRange: 400/400/400
- SizeScaling: 0.05/0.125/0.2
- Width: 225/225/225

## Ahri

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 104 | 2358 |
| dano de ataque | 53 | 3 | 104 |
| armadura | 21 | 4.2 | 92.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.668 | 2.2 | 38.068 |
| regeneração de vida | 0.5 | 0.12 | 2.54 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Furto de Essência** · recarga 0

- ChampionHeal: 165 + 0.3 ability_power _(no nível 18)_
- MinionHeal: 95 + 0.2 ability_power _(no nível 18)_
- MaxStacks: 9
- TakedownWindow: 3

**Q · Orbe da Ilusão** · recarga 7/7/7/7/7 · custo 55/65/75/85/95

- TotalDamage: 35/60/85/110/135 + 0.5 ability_power

**W · Fogo de Raposa** · recarga 9/8/7/6/5 · custo 30/30/30/30/30 · alcance 600/600/600/600/600

- MultiFireDamage (= SingleFireDamage × 1): 16/24/32/40/48 + 0.16000001 ability_power
- SingleFireDamage: 40/60/80/100/120 + 0.4 ability_power
- AcquisitionRange: 550/550/550/550/550
- BonusAcquisitionRange: 725/725/725/725/725
- FlameDuration: 2.5/2.5/2.5/2.5/2.5
- InitialDelay: 0.25/0.25/0.25/0.25/0.25
- MinionBonusDamageMultiplier: 2/2/2/2/2
- MinionBonusDamageThreshold: 0.2/0.2/0.2/0.2/0.2
- MovementSpeed: 0.4/0.4/0.4/0.4/0.4
- MovementSpeedDuration: 2/2/2/2/2
- OrbitRadius: 150/150/150/150/150
- SecondaryDelay: 0.4/0.4/0.4/0.4/0.4

**E · Encanto** · recarga 12/12/12/12/12 · custo 60/60/60/60/60

- TotalDamage: 80/120/160/200/240 + 0.85 ability_power
- CharmDuration: 1.2/1.35/1.5/1.65/1.8
- SlowPercent: -0.65/-0.65/-0.65/-0.65/-0.65

**R · Ímpeto Espiritual** · recarga 140/120/100 · custo 100/100/100

- RCalculatedDamage: 75/125/175 + 0.35 ability_power
- PDurationExtension: 10/10/10
- RAcquisitionRange: 600/600/600
- RBaseDashSpeed: 1200/1200/1200
- RDashCooldown: 1/1/1
- RDashDistance: 500/500/500
- RMaxCasts: 3/3/3
- RMaxTargetsPerCast: 3/3/3
- RRecastWindow: 15/15/15
- RResetCasts: 1/1/1

## Graves

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 625 | 106 | 2427 |
| dano de ataque | 66 | 4 | 134 |
| armadura | 33 | 4.6 | 111.2 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.475 | 3 | 51.475 |
| regeneração de vida | 1.6 | 0.14 | 3.98 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 425 | — | 425 |
| multiplicador de crítico | 2 | — | 2 |

**P · Nova Destino**

- CritDamageMult: -0.5 + 0.5 critical_damage
- MultiBulletDamage (= SingleBulletDamage × 0.333): 0 + 0.33301723 attack_damage _(no nível 18)_
- SingleBulletDamage: 0 + 1.0000517 attack_damage _(no nível 18)_
- {549cda4d}: 0 + 1 critical_damage
- StructureDamageReduction: 0.25

**Q · Fim da Linha** · recarga 13/11.25/9.5/7.75/6 · custo 80/80/80/80/80 · alcance 800/800/800/800/800

- TotalDamage: 50/75/100/125/150 + 0.65 attack_damage (bonus)
- TotalDetonationDamage: 80/125/170/215/260 + 0.55 attack_damage (bonus)
- TerrainCollisionDelay: 0.2/0.2/0.2/0.2/0.2

**W · Cortina de Fumaça** · recarga 26/24/22/20/18 · custo 70/75/80/85/90 · alcance 950/950/950/950/950

- ImpactDamage: 60/110/160/210/260 + 0.6 ability_power
- SlowAmount: 50/50/50/50/50
- SlowDuration: 0.5/0.5/0.5/0.5/0.5
- SmokeDuration: 4/4/4/4/4
- SmokeRadius: 200/200/200/200/200

**E · Saque Rápido** · recarga 16/15/14/13/12 · custo 40/40/40/40/40

- MRGrant: 7/10/13/16/19
- BuffDuration: 4/4/4/4/4
- CooldownPerHit: 0.5/0.5/0.5/0.5/0.5
- DashAngle: 60/60/60/60/60
- DashMaxDistance: 375/375/375/375/375
- DashMinDistance: 275/275/275/275/275
- DashSpeed: 750/750/750/750/750
- MaxStacks: 8/8/8/8/8

**R · Efeito Colateral** · recarga 100/80/60 · custo 100/100/100

- Damage: 275/425/575 + 1.5 attack_damage (bonus)
- FalloffDamage: 200/320/440 + 1.2 attack_damage (bonus)
- RKnockbackDistance: 400/400/400

## Fizz

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 106 | 2442 |
| dano de ataque | 58 | 3 | 109 |
| armadura | 26 | 4.6 | 104.2 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 3.1 | 53.357998 |
| regeneração de vida | 1.6 | 0.14 | 3.98 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Lutador Ligeiro** · custo 40 · alcance 600

- DamageReductionCalc: 4 + 0.01 ability_power
- MonsterDamageReductionCalc: 14 + 0.01 ability_power
- DamageReductionMax: 0.5

**Q · Ataque do Ouriço** · recarga 8/7.5/7/6.5/6 · custo 50/50/50/50/50 · alcance 550/550/550/550/550

- TotalDamage: 0/0/0/0/0 + 1 attack_damage

**W · Tridente da Pedra do Mar** · recarga 7/6/5/4/3 · custo 30/40/50/60/70 · alcance 600/600/600/600/600

- ActiveDamage: 50/75/100/125/150 + 0.45 ability_power
- DoTDamage: 30/45/60/75/90 + 0.25 ability_power
- OnHitBuffDamage: 20/25/30/35/40 + 0.3 ability_power
- ActiveDuration: 4/4/4/4/4
- BleedDuration: 3/3/3/3/3
- BonusMonsterDamage: 90/90/90/90/90
- DoTTicksPerSecond: 2/2/2/2/2
- OnHitBuffDuration: 5/5/5/5/5
- OnKillManaRefund: 30/40/50/60/70
- OnKillNewCooldown: 1/1/1/1/1
- PassiveDoTDuration: 3/3/3/3/3
- TurretMod: 0.5/0.5/0.5/0.5/0.5

**E · Brincalhão / Trapaceiro** · recarga 16/14/12/10/8 · custo 75/80/85/90/95

- EDamage: 80/130/180/230/280 + 0.95 ability_power
- LargeAoESize: 375/375/375/375/375
- SlowAmount: 0.4/0.45/0.5/0.55/0.6
- SlowDuration: 2/2/2/2/2
- SmallAoESize: 225/225/225/225/225

**R · Lançar Isca** · recarga 120/100/80 · custo 100/100/100

- BigSharkDamage (= SmallSharkDamage × 1): 270/450/630 + 0.90000004 ability_power
- SmallSharkDamage: 180/300/420 + 0.6 ability_power
- {262c3c02} (= SmallSharkDamage × 1): 225/375/525 + 0.75 ability_power
- BigFishSize: 450/450/450
- DetonationTime: 2/2/2
- MaxDistance: 1300/1300/1300
- MidFishSize: 325/325/325
- MissileWidthTiers: 80/120/160
- SlowAmount: 60/60/60
- SlowDuration: 2/2/2
- SmallFishSize: 200/200/200

## Volibear

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 104 | 2418 |
| dano de ataque | 65 | 3.5 | 124.5 |
| armadura | 35 | 5.2 | 123.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2 | 34.625 |
| regeneração de vida | 1.8 | 0.15 | 4.35 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · A Tempestade Implacável**

- AttackSpeedCalc: 0.05 + 0.0003 ability_power
- ChainLightningDamage: 11 + 0.45 ability_power
- {ebe98aa3}: 0.05 + 0.04 ability_power
- BounceCounterMax: 5
- BuffDuration: 6
- PDamageRatio: 0.2

**Q · Esmagamento Trovejante** · recarga 12/11.5/11/10.5/10 · custo 50/50/50/50/50 · alcance 300/300/300/300/300

- ADRatioTooltip: 0/0/0/0/0 + 1.6 attack_damage (bonus)
- CalculatedDamage: 10/20/30/40/50 + 1.6 attack_damage (bonus) + 1 attack_damage
- MaxSpeedCalc: 0.24/0.31/0.38/0.45/0.52
- MinSpeedCalc: 0.12/0.155/0.19/0.225/0.26
- {10404951}: 10/20/30/40/50 + 1.6 attack_damage (bonus)
- BonusRange: 25/25/25/25/25
- Duration: 4/4/4/4/4
- StunDuration: 1/1/1/1/1

**W · Fúria Selvagem** · recarga 5/5/5/5/5 · custo 20/25/30/35/40 · alcance 325/325/325/325/325

- PercentMissingHealthHealingRatio: 0.08/0.11/0.14/0.17/0.2
- TotalDamage: 5/30/55/80/105 + 1.1 attack_damage + 0.06 max_health (bonus)
- BaseHeal: 20/35/50/65/80
- MarkDuration: 8/8/8/8/8
- MinionAndMonsterMod: 0.5/0.5/0.5/0.5/0.5

**E · Divisor de Céus** · recarga 16/16/16/16/16 · custo 50/50/50/50/50 · alcance 1200/1200/1200/1200/1200

- APRatioTooltip: 0/0/0/0/0 + 0.7 ability_power
- CalculatedDamage: 80/110/140/170/200 + 0.7 ability_power
- MonsterDamageCap: 150/265/380/495/610 + 0.7 ability_power
- ShieldAPRatioTooltip: 0/0/0/0/0 + 0.75 ability_power
- ShieldValue: 0/0/0/0/0 + 0.75 ability_power + 0.14 max_health
- TotalDamageTooltip: 80/110/140/170/200 + 0.7 ability_power
- PercentDamage: 0.11/0.12/0.13/0.14/0.15
- ShieldDuration: 3/3/3/3/3
- SlowAmount: 0.4/0.4/0.4/0.4/0.4
- SlowDuration: 2/2/2/2/2

**R · Emissário da Tempestade** · recarga 160/135/110 · custo 100/100/100 · alcance 700/700/700

- ADRatio: 0/0/0 + 2.5 attack_damage (bonus)
- APRatioTooltip: 0/0/0 + 1.25 ability_power
- SweetSpotDamageTooltip: 300/500/700 + 1.25 ability_power + 2.5 attack_damage (bonus)
- TowerDamageTooltip: 300/500/700 + 1.25 ability_power + 2.5 attack_damage (bonus)
- {4402b0fa}: 300/500/700 + 1.25 ability_power + 2.5 attack_damage (bonus)
- {d9a2963e}: 175/350/525
- BonusAttackRange: 50/50/50
- DashSpeed: 750/750/750
- SlowAmount: 0.5/0.5/0.5
- SlowDuration: 1/1/1
- TowerDisableDuration: 2/3/4
- TransformDuration: 12/12/12

## Rengar

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 104 | 2358 |
| dano de ataque | 68 | 3 | 119 |
| armadura | 34 | 4.2 | 105.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.667 | 3 | 51.667 |
| regeneração de vida | 1.2 | 0.1 | 2.9 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Predador Oculto**

- EmpoweredMS: 0.5 _(no nível 18)_
- BonusLeapRange: 620
- EmpoweredMSDuration: 1.5
- InCombatTimer: 10
- InCombatTimerVisual: 10
- LeapFerocityGeneration: 1
- MaxFerocity: 4
- RengarPassiveRangeIncrease: 620

**Q · Selvageria** · recarga 0.25/0.25/0.25/0.25/0.25

- EmpoweredQAS: 50/50/50/50/50
- EmpoweredQTotalDamage: 35/35/35/35/35 + 1.2 attack_damage
- QBonusDamage: 20/55/90/125/160 + 0.05 attack_damage
- QTotalDamage: 20/55/90/125/160 + 1.05 attack_damage
- {95c7ca04}: 12/33/54.000004/75/96 + 0.030000001 attack_damage
- {d4a43abe}: 35/35/35/35/35 + 0.2 attack_damage
- {f2ca2500}: 21/21/21/21/21 + 0.120000005 attack_damage
- ASBonus: 40/40/40/40/40
- ASDuration: 5/5/5/5/5
- BaseADRatioTooltip: 5/5/5/5/5
- BuffDuration: 3/3/3/3/3

**W · Rugido de Batalha** · recarga 0.25/0.25/0.25/0.25/0.25 · alcance 450/450/450/450/450

- BonusMonsterDamage: 130/130/130/130/130 _(no nível 18)_
- TotalDamage: 50/80/110/140/170 + 0.8 ability_power
- TotalDamageEmpowered: 220/220/220/220/220 + 0.8 ability_power _(no nível 18)_
- CCImmuneDuration: 1.5/1.5/1.5/1.5/1.5
- DamagePercentageHealed: 50/50/50/50/50
- HealingWindow: 1.5/1.5/1.5/1.5/1.5
- MonsterHealingMod: 100/100/100/100/100

**E · Boleadeiras** · recarga 0.25/0.25/0.25/0.25/0.25

- TotalDamage: 55/100/145/190/235 + 0.8 attack_damage (bonus)
- TotalEmpoweredDamage: 50/50/50/50/50 + 0.8 attack_damage (bonus)
- CCDuration: 1.75/1.75/1.75/1.75/1.75
- RevealDuration: 2/2/2/2/2
- SlowAmount: 30/45/60/75/90

**R · Furor da Caçada** · recarga 100/90/80 · alcance 2500/3000/3500

- BonusDamage: 0/0/0 + 1 attack_damage
- ArmorShred: 15/20/25
- ArmorShredDuration: 4/4/4
- EnemyDetectionRange: 1600/1600/1600
- FadeTime: 2/2/2
- LeapRange: 725/725/725
- SelfRevealRange: 750/750/750
- SelfVisionRange: 2500/3000/3500
- StealthDuration: 12/16/20
- StealthMS: 40/50/60

## Varus

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 105 | 2385 |
| dano de ataque | 59 | 3.4 | 116.8 |
| armadura | 24 | 4 | 92 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.658 | 3.5 | 60.158 |
| regeneração de vida | 0.7 | 0.11 | 2.57 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 575 | — | 575 |
| multiplicador de crítico | 2 | — | 2 |

**P · Vingança Viva** · alcance 650

- ASDuration: 11 _(no nível 18)_
- ChampionAD: 0 + 33 attack_speed (bonus)
- ChampionAP: 0 + 33 attack_speed (bonus)
- ChampionAS: 0.3
- MinionAD: 0 + 11 attack_speed (bonus)
- MinionAP: 0 + 11 attack_speed (bonus)
- MinionAS: 0.2 _(no nível 18)_
- NewASCap: 3.33
- PassiveASMinion: 0.1
- PassiveASRatio: 0.4
- PassiveASRatioMinion: 0.2

**Q · Flecha Perfurante** · recarga 16/15/14/13/12 · custo 50/55/60/65/70 · alcance 925/925/925/925/925

- TotalDamageMax: 80/150/220/290/360 + 1.2 attack_damage (bonus)
- TotalDamageMinTooltip (= TotalDamageMax × 1): 53.333603/100.0005/146.6674/193.3343/240.0012 + 0.80000407 attack_damage (bonus)
- {b3ef3984} (= TotalDamageMax × 1): 53.333603/100.0005/146.6674/193.3343/240.0012 + 0.80000407 attack_damage (bonus)
- CDRefund: 3/3/3/3/3
- FalloffPercent: 0.15/0.15/0.15/0.15/0.15
- ManaRefund: 0.5/0.5/0.5/0.5/0.5
- MaxChannelDuration: 4/4/4/4/4
- MaxChargeAmp: 0.5/0.5/0.5/0.5/0.5
- MinDamagePercent: 0.33/0.33/0.33/0.33/0.33
- MoveSpeedMod: -0.2/-0.2/-0.2/-0.2/-0.2

**W · Aljava da Ruína** · recarga 40/40/40/40/40

- MaxPercentHPPerStack (= PercentHPPerStack × 3): 0.089999996/0.105000004/0.12/0.135/0.15 + 0.00039 ability_power
- MaxQEmpowerPercentHP (= QEmpowerPercentHP × 1): 0.089999996/0.12/0.15/0.17999999/0.21000001
- OnHitDamage: 4/13/22/31/40 + 0.25 ability_power + 0.15 attack_damage (bonus)
- PercentHPPerStack: 0.03/0.035/0.04/0.045/0.05 + 0.00013 ability_power
- QEmpowerPercentHP: 0.06/0.08/0.1/0.12/0.14
- CDRPerBlightStack: 0.13/0.13/0.13/0.13/0.13
- DebuffDuration: 6/6/6/6/6
- MaxMonsterDamage: 120/120/120/120/120
- MaxStacks: 3/3/3/3/3
- QEmpowerMonsterCap: 360/360/360/360/360
- VarusWDebuffDuration: 6/6/6/6/6
- VarusWMaxStacks: 3/3/3/3/3

**E · Chuva de Flechas** · recarga 18/16/14/12/10 · custo 90/90/90/90/90 · alcance 925/925/925/925/925

- TotalDamage: 60/90/120/150/180 + 0.9 attack_damage (bonus)
- AoERange: 300/300/300/300/300
- DebuffDuration: 4/4/4/4/4
- GrievousAmount: 0.4/0.4/0.4/0.4/0.4
- GrievousDuration: 0.5/0.5/0.5/0.5/0.5
- GroundDuration: 4/4/4/4/4
- SlowPercent: -0.3/-0.35/-0.4/-0.45/-0.5

**R · Corrente da Corrupção** · recarga 100/80/60 · custo 100/100/100

- TotalDamage: 150/250/350 + 1 ability_power
- PassiveStacks: 3/3/3
- PassiveStacksAdded: 3/3/3
- RootDuration: 2/2/2
- TimeBetweenStackAddition: 0.5/0.5/0.5
- VarusRLeashTimer: 1.75/1.75/1.75
- VarusRTargetAcquisitionRange: 650/650/650
- VarusRTargetLeashRange: 600/600/600
- VarusWDebuffDuration: 6/6/6
- VarusWMaxStacks: 3/3/3

## Nautilus

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 646 | 100 | 2346 |
| dano de ataque | 61 | 3.3 | 117.1 |
| armadura | 39 | 4.95 | 123.149994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.706 | 1 | 17.706 |
| regeneração de vida | 1.7 | 0.11 | 3.57 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Âncora Impactante**

- BonusDamage: 116 _(no nível 18)_
- RootDuration: 1.5 _(no nível 18)_
- TotalDamageTooltip: 116 + 1 attack_damage _(no nível 18)_
- PerTargetCD: 6

**Q · Lançar Âncora** · recarga 14/13/12/11/10 · custo 60/60/60/60/60

- QDamageCalc: 85/130/175/220/265 + 0.9 ability_power
- TerrainCDR: 0.5/0.5/0.5/0.5/0.5
- TerrainMana: 0.5/0.5/0.5/0.5/0.5
- TerrainManaRefund: 0.5/0.5/0.5/0.5/0.5
- TerrainReduction: 0.5/0.5/0.5/0.5/0.5

**W · Ira do Titã** · recarga 12/12/12/12/12 · custo 60/60/60/60/60 · alcance 700/800/900/1000/1100

- DotDamageCalc: 30/40/50/60/70 + 0.4 ability_power
- ShieldCalc: 50/60/70/80/90 + 0.08 max_health
- ShieldDuration: 6/6/6/6/6

**E · Correnteza** · recarga 7/6.5/6/5.5/5 · custo 50/60/70/80/90

- DamageCalc: 55/90/125/160/195 + 0.5 ability_power
- MonsterBonusCalc: 125/165/205/245/285 + 0.5 ability_power
- DamageMonsterMod: 1.5/1.5/1.5/1.5/1.5
- DamageRadius: 125/125/125/125/125
- ExtraWavePenalty: 0.5/0.5/0.5/0.5/0.5
- MultiHitReduction: 0.5/0.5/0.5/0.5/0.5
- SlowDuration: 1.25/1.25/1.25/1.25/1.25
- SlowPercent: 0.3/0.35/0.4/0.45/0.5

**R · Carga de Profundidade** · recarga 120/100/80 · custo 100/100/100 · alcance 825/825/825

- PrimaryTargetDamage: 150/275/400 + 0.8 ability_power
- SecondaryTargetDamage: 125/175/225 + 0.4 ability_power
- StunDuration: 1/1.5/2

## Viktor

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 100 | 2300 |
| dano de ataque | 53 | 3 | 104 |
| armadura | 23 | 4.4 | 97.8 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 2.11 | 36.528 |
| regeneração de vida | 1.6 | 0.13 | 3.81 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Gloriosa Evolução**

- ARAMStackCadence: 5
- ARAMStacksoverTime: 1
- CannonStacks: 10
- ChampionStacks: 20
- EvolutionStackBreakpoint: 100
- MinionStacks: 1

**Q · Poder do Sifão** · recarga 9/8/7/6/5 · custo 45/50/55/60/65 · alcance 600/600/600/600/600

- AttackTotalDMG: 20/45/70/95/120 + 0.5 ability_power + 1 attack_damage
- ShieldLevelScaling: 140/140/140/140/140 + 0.25 ability_power _(no nível 18)_
- TotalAugmentedShieldValue (= ShieldLevelScaling × 1): 224/224/224/224/224 + 0.4 ability_power _(no nível 18)_
- TotalMissileDamage: 60/75/90/105/120 + 0.4 ability_power
- {147c7de9}: 0/0/0/0/0 + 1 attack_damage
- AugmentMoveSpeedBonus: 30/30/30/30/30
- BuffDuration: 2.5/2.5/2.5/2.5/2.5
- ShieldManaRatio: 0.08/0.08/0.08/0.08/0.08
- TurretDamageMod: 1/1/1/1/1

**W · Campo Gravítico** · recarga 17/16/15/14/13 · custo 65/65/65/65/65 · alcance 800/800/800/800/800

- AugmentSlow: 20/20/20/20/20
- FieldDuration: 4/4/4/4/4
- Size: 275/275/275/275/275
- SlowDuration: 1/1/1/1/1
- SlowPotency: -33/-36/-39/-42/-45
- StackCadence: 0.5/0.5/0.5/0.5/0.5
- StacksForStun: 6/6/6/6/6
- StunDuration: 1.5/1.5/1.5/1.5/1.5

**E · Raio Hextec** · recarga 12/11/10/9/8 · custo 60/70/80/90/100 · alcance 550/550/550/550/550

- AftershockDamage: 20/50/80/110/140 + 0.8 ability_power
- LaserDamage: 70/110/150/190/230 + 0.5 ability_power
- {86b88e5d}: 20/50/80/110/140 + 0.8 ability_power
- Delay: 1/1/1/1/1
- LaserRange: 700/700/700/700/700

**R · Tempestade Arcana** · recarga 120/100/80 · custo 100/100/100 · alcance 700/700/700

- InitialBurstDamage: 100/175/250 + 0.5 ability_power
- SubsequentBurstDamage: 65/105/145 + 0.35 ability_power
- {d3c6a7eb}: 3/3/3
- AugmentBoost: 0.25/0.25/0.25
- EnhancedSizeMult: 0.4/0.4/0.4
- EnhancedSlow: 40/40/40
- EnhancedSlowDuration: 2/2/2
- MaxDistance: 1000/1000/1000
- MaxGrowths: 6/6/6
- MaxStormSpeed: 275/275/275
- MaxTicks: 6/6/6
- MinDistance: 300/300/300
- MinStormSpeed: 200/200/200
- StormDuration: 6.5/6.5/6.5
- StormRadius: 325/325/325
- TickCadence: 1/1/1
- Tooltip_DurationExtension: 3/3/3

## Sejuani

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 114 | 2568 |
| dano de ataque | 66 | 4 | 134 |
| armadura | 34 | 5.45 | 126.649994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.688 | 3.5 | 60.188 |
| regeneração de vida | 1.7 | 0.2 | 5.1 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Fúria do Norte**

- FrostArmorOOC: 6 _(no nível 18)_
- PercentHPDamage: 0.1
- TotalArmorTooltip: 10 + 0.75 armor (bonus)
- TotalMRTooltip: 10 + 0.75 magic_resist (bonus)
- EpicMonsterCap: 250
- FrostArmorDuration: 3

**Q · Ataque do Ártico** · recarga 18/16.5/15/13.5/12 · custo 60/65/70/75/80

- TotalDamageTooltip: 90/140/190/240/290 + 0.75 ability_power
- DashSpeed: 1000/1000/1000/1000/1000
- KnockupDurationTOOLTIPONLY: 0.5/0.5/0.5/0.5/0.5
- MaxTravelDistance: 625/625/625/625/625

**W · Ira do Inverno** · recarga 9/8/7/6/5 · custo 60/60/60/60/60

- FirstHitDamageTooltip: 5/15/25/35/45 + 0.3 ability_power + 0.04 max_health
- SecondHitDamageTooltip: 5/25/45/65/85 + 0.6 ability_power + 0.08 max_health
- SlowAmount: 0.75/0.75/0.75/0.75/0.75
- SlowDuration: 0.25/0.25/0.25/0.25/0.25

**E · Congelamento Permanente** · recarga 1.5/1.5/1.5/1.5/1.5 · custo 20/20/20/20/20 · alcance 600/600/600/600/600

- TotalDamage: 55/105/155/205/255 + 0.7 ability_power
- AuraRange: 1100/1100/1100/1100/1100
- CCDuration: 1/1/1/1/1
- MaxStacks: 4/4/4/4/4
- PerChampionCD: 8/8/8/8/8
- StackDuration: 5/5/5/5/5

**R · Prisão Glacial** · recarga 120/105/90 · custo 100/100/100

- MinorDamageTooltip: 125/150/175 + 0.4 ability_power
- TotalDamageTooltip: 200/300/400 + 0.8 ability_power
- BaseStunDuration: 1/1/1
- EmpoweredStunDuration: 1.5/1.5/1.5
- ExplosionSlowAmount: 80/80/80
- ExplosionSlowDuration: 1/1/1
- ZoneDuration: 2/2/2
- ZoneSlowAmount: 30/30/30

## Fiora

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 99 | 2303 |
| dano de ataque | 66 | 3.3 | 122.1 |
| armadura | 33 | 4.7 | 112.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.69 | 3.2 | 55.09 |
| regeneração de vida | 1.7 | 0.11 | 3.57 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Dança da Duelista** · alcance 1100

- PassiveDamageTotal: 0.03 + 0.0004 attack_damage (bonus)
- PassiveHealAmount: 100 _(no nível 18)_
- RDamageTotal (= PassiveDamageTotal × 4): 0.12 + 0.0016 attack_damage (bonus)
- MovementSpeedDuration: 2

**Q · Estocada** · recarga 13/11.25/9.5/7.75/6 · custo 20/20/20/20/20

- TotalDamage: 70/80/90/100/110 + 0.9 attack_damage (bonus)
- CDRefundPercent: 0.5/0.5/0.5/0.5/0.5

**W · Ripostar** · recarga 24/22/20/18/16 · custo 50/50/50/50/50 · alcance 3000/3000/3000/3000/3000

- StabDamage: 110/150/190/230/270 + 1 ability_power
- AttackSlowPercent: -0.25/-0.25/-0.25/-0.25/-0.25
- BigBlockPercent: 0.15/0.15/0.15/0.15/0.15
- CCDuration: 2/2/2/2/2
- MSSlowPercent: -0.5/-0.5/-0.5/-0.5/-0.5
- ParryDuration: 0.75/0.75/0.75/0.75/0.75
- SmallBlockPercent: 0.03/0.03/0.03/0.03/0.03

**E · Esgrima** · recarga 11/10/9/8/7 · custo 40/40/40/40/40 · alcance 425/425/425/425/425

- ASPercent: 0.5/0.6/0.7/0.8/0.9
- AttackOnePercentTAD: 1/1/1/1/1
- AttackTwoPercentTAD: 1.6/1.7/1.8/1.9/2
- BuffDuration: 3/3/3/3/3
- SlowDuration: 1/1/1/1/1
- SlowPercent: -0.3/-0.3/-0.3/-0.3/-0.3

**R · Desafio Grandioso** · recarga 110/90/70 · custo 100/100/100 · alcance 500/500/500

- HealPerSecondCalc: 75/100/125 + 0.6 attack_damage (bonus)
- HealDuration: 5/5/5
- HealDurationExtension: 1/1/1
- HealRingRadius: 550/550/550
- MSRingRadius: 500/500/500
- MarkDuration: 8/8/8
- PassiveAS: 0.3/0.4/0.5
- PercentMS: 0.3/0.4/0.5
- Ratio: 0.6/0.6/0.6

## Ziggs

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 606 | 106 | 2408 |
| dano de ataque | 55 | 3.1 | 107.7 |
| armadura | 21 | 4.7 | 100.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.656 | 2 | 34.656 |
| regeneração de vida | 1.3 | 0.12 | 3.34 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Pavio Curto** · recarga 12

- SpellCdr: 6 _(no nível 18)_
- StructureDamage (= TotalDamage × 1): 35 + 0.875 ability_power
- TotalDamage: 20 + 0.5 ability_power

**Q · Bomba Saltitante** · recarga 6/5.5/5/4.5/4 · custo 50/55/60/65/70

- TotalDamage: 80/130/180/230/280 + 0.6 ability_power
- BombTriggerRange: 180/180/180/180/180
- BoundingBoxTrigger: 70/70/70/70/70
- ExplosionRadius: 240/240/240/240/240
- MaxRange: 850/850/850/850/850
- MaxRangeFirstBounce: 325/325/325/325/325
- MaxRangeSecondBounce: 225/225/225/225/225

**W · Carga Concentrada** · recarga 20/18/16/14/12 · custo 80/80/80/80/80 · alcance 1000/1000/1000/1000/1000

- TotalDamage: 70/105/140/175/210 + 0.5 ability_power
- BombDuration: 4/4/4/4/4
- CherryBonusHaste: 30/30/30/30/30
- ExplosionRadius: 325/325/325/325/325
- KnockbackDistance: 500/500/500/500/500
- KnockbackDistanceAlly: 775/775/775/775/775
- KnockbackGravity: 50/50/50/50/50
- KnockbackSpeed: 600/600/600/600/600
- PerceptionBubbleDuration: 6/6/6/6/6
- PerceptionBubbleRadius: 400/400/400/400/400
- TurretDestroyPercent: 0.25/0.275/0.3/0.325/0.35

**E · Campo Minado de Hexplosivos** · recarga 16/16/16/16/16 · custo 70/80/90/100/110 · alcance 900/900/900/900/900

- ReducedDamage (= TotalDamage × 1): 12/28/44/60/76 + 0.1 ability_power
- TotalDamage: 30/70/110/150/190 + 0.25 ability_power
- MineDuration: 10/10/10/10/10
- MinePerceptionBubbleSize: 150/150/150/150/150
- Slow: -0.1/-0.2/-0.3/-0.4/-0.5
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- TriggerRange: 135/135/135/135/135

**R · Bomba Megainfernal** · recarga 120/95/70 · custo 100/100/100 · alcance 5000/5000/5000

- BlastDamage (= EmpoweredDamage × 1): 195/325/454.99997 + 0.65 ability_power
- EmpoweredDamage: 300/500/700 + 1 ability_power
- EmpoweredRadius: 250/250/250
- MaximumRadius: 525/525/525

## Lulu

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 565 | 92 | 2129 |
| dano de ataque | 47 | 2.6 | 91.2 |
| armadura | 26 | 4.6 | 104.2 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.25 | 38.875 |
| regeneração de vida | 1.2 | 0.12 | 3.24 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Pix, o Silfo Companheiro**

- CombinedDamage (= TotalDamage × 1): 117 + 0.15 ability_power _(no nível 18)_
- TotalDamage: 39 + 0.05 ability_power _(no nível 18)_

**Q · Lança-Purpurina** · recarga 7/7/7/7/7 · custo 50/55/60/65/70

- BonusMissileDamage: 60/95/130/165/200 + 0.5 ability_power
- TotalDamage: 60/95/130/165/200 + 0.5 ability_power
- MinionMod: 0.7/0.7/0.7/0.7/0.7
- SlowAmount: -0.8/-0.8/-0.8/-0.8/-0.8
- SlowDecayTicks: 8/8/8/8/8
- SlowDuration: 2/2/2/2/2

**W · Caprichos** · recarga 18/18/18/18/18 · custo 65/65/65/65/65 · alcance 650/650/650/650/650

- TotalMS: 0.25/0.25/0.25/0.25/0.25 + 0.0005 ability_power
- ASBonus: 0.2/0.225/0.25/0.275/0.3
- AmountToMultiplyCoefficient2ByForPurposesOfCalculatingMovementSpeed: 0.01/0.01/0.01/0.01/0.01
- CCDuration: 1.2/1.4/1.6/1.8/2
- MSDuration: 3/3.25/3.5/3.75/4
- SlowAmount: -60/-60/-60/-60/-60

**E · Socorro, Pix!** · recarga 10/9.5/9/8.5/8 · custo 60/65/70/75/80 · alcance 650/650/650/650/650

- TotalDamage: 70/110/150/190/230 + 0.5 ability_power
- TotalShield: 70/110/150/190/230 + 0.5 ability_power
- PixOnAllyDuration: 6/6/6/6/6
- PixOnEnemyDuration: 4/4/4/4/4
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5

**R · Crescimento Virente** · recarga 120/100/80 · custo 100/100/100 · alcance 900/900/900

- TotalBonusHealth: 275/425/575 + 0.55 ability_power
- AoERadius: 400/400/400
- BuffDuration: 7/7/7
- KnockbackDistance: 275/275/275
- KnockbackDuration: 1/1/1
- SlowPercent: 30/45/60

## Draven

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 675 | 104 | 2443 |
| dano de ataque | 62 | 3.4 | 119.8 |
| armadura | 29 | 4.5 | 105.5 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.679 | 2.7 | 46.579002 |
| regeneração de vida | 0.75 | 0.14 | 3.13 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · League of Draven**

- CherryStackMultiplier: 5
- CherryStacksPerRoundBeforeMultiplier: 10
- PassiveGoldBase: 25
- PassiveGoldPerStack: 2
- PercentOfStacksLost: 50
- StackGain: 1
- StackGainEnemy: 10

**Q · Revolução do Machado** · recarga 12/11/10/9/8 · custo 45/45/45/45/45

- TotalDamage: 40/45/50/55/60 + 0.75 attack_damage (bonus)
- AxeDuration: 5.75/5.75/5.75/5.75/5.75
- DurationTOOLTIP: 6/6/6/6/6

**W · Adrenalina** · recarga 12/12/12/12/12 · custo 40/35/30/25/20 · alcance 1000/1000/1000/1000/1000

- AttackSpeed: 0.2/0.25/0.3/0.35/0.4
- AttackSpeedDuration: 3/3/3/3/3
- MoveSpeed: 0.5/0.55/0.6/0.65/0.7
- MoveSpeedDuration: 1.5/1.5/1.5/1.5/1.5
- Temp_AS: 20/25/30/35/40
- Temp_ASDuration: 3/3/3/3/3
- Temp_MSDecay: -0.062/-0.069/-0.075/-0.081/-0.087
- Temp_MSDuration: 1.5/1.5/1.5/1.5/1.5
- Temp_MSMod: 50/55/60/65/70

**E · Sai da Frente** · recarga 16/15/14/13/12 · custo 70/70/70/70/70

- TotalDamage: 75/110/145/180/215 + 0.5 attack_damage (bonus)
- KnockbackDuration: 0.5/0.5/0.5/0.5/0.5
- SlowAmount: 20/25/30/35/40
- SlowDuration: 2/2/2/2/2

**R · Reta da Morte** · recarga 100/90/80 · custo 100/100/100

- RCalculatedDamage: 200/300/400 + 1.1 attack_damage (bonus)
- RDamageReductionPerHit: 0.05/0.05/0.05
- RMinDamagePercent: 50/50/50

## Hecarim

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 625 | 106 | 2427 |
| dano de ataque | 66 | 3.7 | 128.9 |
| armadura | 32 | 5.45 | 124.649994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.67 | 2.5 | 43.17 |
| regeneração de vida | 1.4 | 0.15 | 3.95 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Caminho da Guerra**

- BonusAD: 0 + 0.24 move_speed (bonus) _(no nível 18)_

**Q · Enfurecido** · recarga 4/4/4/4/4 · custo 28/26/24/22/20 · alcance 375/375/375/375/375

- Damage: 60/90/120/150/180 + 0.9 attack_damage (bonus)
- MinionDamage (= Damage × 1): 36/54.000004/72/90/108.00001 + 0.54 attack_damage (bonus)
- RampageBonusDamagePerc: 3/3/3/3/3 + 0.03 attack_damage (bonus)
- BuffDuration: 8/8/8/8/8
- MaxStacks: 3/3/3/3/3
- RampageCooldownReduction: 0.75/0.75/0.75/0.75/0.75
- RampageDuration: 6/6/6/6/6

**W · Espírito do Pavor** · recarga 14/14/14/14/14 · custo 50/55/60/65/70 · alcance 525/525/525/525/525

- AllyTooltipLeachValue (= LeechAmount × 1): 12.5/12.5/12.5/12.5/12.5
- LeechAmount: 25/25/25/25/25
- TotalDamage: 80/120/160/200/240 + 0.8 ability_power
- BuffDuration: 4/4/4/4/4
- HealBonusADRatio: 0.02/0.02/0.02/0.02/0.02
- MinionHealCap: 120/150/180/210/240
- ResistAmount: 5/10/15/20/25
- TickTime: 0.5/0.5/0.5/0.5/0.5

**E · Ataque Devastador** · recarga 20/19/18/17/16 · custo 60/60/60/60/60 · alcance 300/300/300/300/300

- MaxDamage (= MinDamage × 2): 60/90/120/150/180 + 1 attack_damage (bonus)
- MinDamage: 30/45/60/75/90 + 0.5 attack_damage (bonus)
- DistanceToMaxDamage: 1200/1200/1200/1200/1200
- Duration: 4/4/4/4/4
- MaxBaseDamage: 60/90/120/150/180
- MaxKnockback: 350/350/350/350/350
- MaxMoveSpeed: 0.65/0.65/0.65/0.65/0.65
- MinKnockback: 150/150/150/150/150
- MinMoveSpeed: 0.25/0.25/0.25/0.25/0.25
- TimeToMaxMoveSpeed: 2.5/2.5/2.5/2.5/2.5

**R · Massacre das Sombras** · recarga 140/120/100 · custo 100/100/100

- DamageDone: 150/250/350 + 1 ability_power
- DashSpeed: 1100/1100/1100
- FearDurationMax: 1.5/1.5/1.5
- FearDurationMin: 0.75/0.75/0.75
- MaxDashRange: 1000/1000/1000
- RangeForMaxFearDuration: 950/950/950

## Kha'Zix

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 643 | 99 | 2326 |
| dano de ataque | 60 | 3.1 | 112.7 |
| armadura | 32 | 4.2 | 103.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.668 | 2.7 | 46.568 |
| regeneração de vida | 1.5 | 0.15 | 4.05 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ameaça Invisível** · recarga 0

- TotalDamage: 136 + 0.5 attack_damage (bonus) _(no nível 18)_
- IsolationRange: 375
- SlowAmount: 0.25
- SlowDuration: 2

**Q · Sabor de Medo** · recarga 4/4/4/4/4 · custo 20/20/20/20/20 · alcance 325/325/325/325/325

- IsEvolved: 1/1/1/1/1
- {8b6fe763}: 0/0/0/0/0
- {ba2719ef}: 1/1/1/1/1
- EvolutionIsolationCDRPercentage: 45/45/45/45/45

**W · Espinho do Vazio** · recarga 9/9/9/9/9 · custo 55/60/65/70/75

- HealAmount: 55/75/95/115/135 + 0.5 ability_power
- IsEvolved: 1/1/1/1/1
- {8b6fe763}: 0/0/0/0/0
- {ba2719ef}: 1/1/1/1/1
- IsolatedSlowDuration: 2/2/2/2/2
- IsolatedSlowPercentage: 60/60/60/60/60
- SlowDuration: 2/2/2/2/2
- SlowPercentage: 40/40/40/40/40

**E · Pulo** · recarga 20/18/16/14/12 · custo 50/50/50/50/50

- IsEvolved: 1/1/1/1/1
- {8b6fe763}: 0/0/0/0/0
- {ba2719ef}: 1/1/1/1/1
- ADRatio: 0.4/0.4/0.4/0.4/0.4
- BaseDamage: 65/100/135/170/205
- EvolvedLeapRange: 900/900/900/900/900

**R · Massacre do Vazio** · recarga 100/85/70 · custo 100/100/100

- IsEvolved: 1/1/1
- {8b6fe763}: 0/0/0
- {ba2719ef}: 1/1/1
- BonusMovementSpeedPercent: 0.4/0.4/0.4
- EvolutionsAvailable: 1/2/3
- EvolvedNumberOfCasts: 3/3/3
- EvolvedStealthDuration: 2/2/2
- NumberOfCasts: 2/2/2
- RecastCD: 2/2/2
- RecastWindow: 12/12/12
- StealthDuration: 1.25/1.25/1.25

## Darius

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 652 | 114 | 2590 |
| dano de ataque | 64 | 5 | 149 |
| armadura | 37 | 5.2 | 125.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 1 | 17.625 |
| regeneração de vida | 2 | 0.19 | 5.23 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Hemorragia**

- BleedDamagePerStack: 30 + 0.3 attack_damage (bonus) _(no nível 18)_
- NoxianMightBonusAD: 30
- BleedDuration: 5
- MaxStacks: 5
- MonsterMod: 2.5
- SecondsPerTick: 1.26

**Q · Dizimar** · recarga 9/8/7/6/5 · custo 25/30/35/40/45 · alcance 270/270/270/270/270

- BladeDamage: 50/80/110/140/170 + 1 attack_damage
- HandleDamage (= BladeDamage × 0.35): 17.5/28/38.5/49/59.5 + 0.35 attack_damage
- InnerCircleDamagePercent: 35/35/35/35/35
- MissingHealPercent: 51/51/51/51/51
- MissingHealthHeal: 17/17/17/17/17

**W · Ataque Mutilador** · recarga 5/5/5/5/5 · custo 40/40/40/40/40 · alcance 300/300/300/300/300

- EmpoweredAttackDamage: 0/0/0/0/0 + 1.4 attack_damage
- PercentCDRefund: 50/50/50/50/50
- SlowDuration: 1/1/1/1/1
- SlowPercent: 90/90/90/90/90

**E · Apreender** · recarga 26/23.5/21/18.5/16 · custo 70/60/50/40/30 · alcance 550/550/550/550/550

- PassivePercentArmorPen: 20/25/30/35/40
- SlowDuration: 1/1/1/1/1
- SlowPercent: 40/40/40/40/40

**R · Guilhotina de Noxus** · recarga 120/100/80 · custo 100/100/0 · alcance 475/475/475

- Damage: 125/250/375 + 0.75 attack_damage (bonus)
- MaximumDamage (= Damage × 2): 250/500/750 + 1.5 attack_damage (bonus)
- RDamagePercentPerHemoStack: 0.2/0.2/0.2
- RRecastDuration: 20/20/20

## Jayce

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 109 | 2443 |
| dano de ataque | 59 | 4.25 | 131.25 |
| armadura | 22 | 5 | 107 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 3 | 51.658 |
| regeneração de vida | 1.2 | 0.12 | 3.24 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Capacitor Hextec** · recarga 0

- FlatMovementSpeed: 30
- MovementSpeedDuration: 0.75

**Q · Aos Céus! / Disparo Chocante** · recarga 16/14/12/10/8 · custo 40/40/40/40/40 · alcance 600/600/600/600/600

- Damage: 60/110/160/210/260 + 1.35 attack_damage (bonus)
- MonsterBonusDamage: 10/10/10/10/10
- Slow: -0.35/-0.4/-0.45/-0.5/-0.55
- SlowDuration: 2/2/2/2/2

**W · Campo Elétrico / Hipercarga** · custo 40/40/40/40/40 · alcance 285/285/285/285/285

- Damage: 140/200/260/320/380 + 1 ability_power
- Duration: 4/4/4/4/4
- ManaGain: 15/17/19/21/23

**E · Golpe Trovejante / Portão Acelerador** · recarga 20/18/16/14/12 · custo 55/55/55/55/55 · alcance 240/240/240/240/240

- FlatDamage: 0/0/0/0/0 + 1 attack_damage (bonus)
- KnockbackDistance: 600/600/600/600/600
- KnockbackDuration: 0.35/0.35/0.35/0.35/0.35
- MonsterCap: 200/300/400/500/600
- PercHPDamage: 0.08/0.108/0.136/0.164/0.192

**R · Canhão de Mercúrio / Martelo de Mercúrio** · recarga 6/6/6 · alcance 600/600/600

- Damage: 130/130/130 + 0.3 attack_damage (bonus) _(no nível 18)_
- RangedFormShred: 0.35/0.35/0.35 _(no nível 18)_
- Resists: 26/26/26 + 0.075 attack_damage (bonus) _(no nível 18)_
- RangedFormRangeIncrease: 375/375/375
- ShredDuration: 5/5/5

## Lissandra

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 110 | 2490 |
| dano de ataque | 55 | 2.7 | 100.9 |
| armadura | 22 | 4.2 | 93.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.656 | 1.5 | 26.156 |
| regeneração de vida | 1.4 | 0.11 | 3.27 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Submissão Glacinata**

- TotalDamage: 120 + 0.5 ability_power
- ExplosionDelay: 4
- MoveSpeedMod: -0.25
- Radius: 450
- Range: 1350

**Q · Estilhaço de Gelo** · recarga 7/6/5/4/3 · custo 55/60/65/70/75

- TotalDamage: 80/115/150/185/220 + 0.75 ability_power
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- SlowPercentage: -0.2/-0.24/-0.28/-0.32/-0.36
- slowDuration: 1.5/1.5/1.5/1.5/1.5
- slowPercentage: -0.2/-0.24/-0.28/-0.32/-0.36

**W · Círculo Ártico** · recarga 10/9.5/9/8.5/8 · custo 40/40/40/40/40 · alcance 450/450/450/450/450

- TotalDamage: 70/105/140/175/210 + 0.7 ability_power
- FreeCastCD: 8/8/8/8/8
- SnareDuration: 1.25/1.35/1.45/1.55/1.65
- TargetsToHit: 3/3/3/3/3

**E · Caminho Glacial** · recarga 24/21/18/15/12 · custo 80/85/90/95/100

- TotalDamage: 70/105/140/175/210 + 0.6 ability_power

**R · Túmulo Congelado** · recarga 120/100/80 · custo 100/100/100 · alcance 550/550/550

- CalculatedDamage: 150/250/350 + 0.75 ability_power
- HealAmount: 100/150/200 + 0.55 ability_power
- EnemyCastDuration: 1.5/1.5/1.5
- SelfCastDuration: 2.5/2.5/2.5
- SelfCastMissingHPPerAbove: 1/1/1
- SelfCastMissingHPRatio: 1/1/1
- SlowAmount: -0.45/-0.6/-0.75
- SlowDuration: 3/3/3
- StartingAoESize: 292/292/292

## Diana

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 109 | 2493 |
| dano de ataque | 57 | 3 | 108 |
| armadura | 31 | 4.2999997 | 104.1 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2 | 34.625 |
| regeneração de vida | 1.3 | 0.17 | 4.19 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Espada de Prata Lunar** · recarga 28

- BonusAS: 0.35 _(no nível 18)_
- CleaveDamage: 20 + 0.5 ability_power
- EmpoweredAS (= BonusAS × 3): 1.05 _(no nível 18)_
- AttackCount: 3
- AttackSpeedValue: 0.3
- BuffDuration: 5
- MonsterMod: 2.7

**Q · Golpe Crescente** · recarga 8/7.5/7/6.5/6 · custo 50/50/50/50/50

- TotalDamage: 70/105/140/175/210 + 0.7 ability_power
- MoonlightDuration: 3/3/3/3/3

**W · Cascata Lívida** · recarga 15/13.5/12/10.5/9 · custo 40/45/50/55/60 · alcance 750/750/750/750/750

- ShieldValue: 45/60/75/90/105 + 0.3 ability_power + 0.11 max_health (bonus)
- TotalDamage: 20/32/44/56/68 + 0.18 ability_power
- TotalMaxDamage (= TotalDamage × 3): 60/96/132/168/204 + 0.54 ability_power
- ShieldDuration: 5/5/5/5/5

**E · Zênite Lunar** · recarga 22/20/18/16/14 · custo 40/45/50/55/60 · alcance 825/825/825/825/825

- TotalDamage: 50/70/90/110/130 + 0.6 ability_power
- CDResetTime: 0.25/0.25/0.25/0.25/0.25

**R · Colapso Minguante** · recarga 100/90/80 · custo 100/100/100 · alcance 475/475/475

- MaxDamage (= RMultiHitAmplification × 5): 175/300/425 + 0.75 ability_power
- RExplosionDamage: 200/300/400 + 0.6 ability_power
- RMultiHitAmplification: 35/60/85 + 0.15 ability_power
- Delay: 1/1/1
- SlowDuration: 2/2/2
- SlowPotency: -0.4/-0.5/-0.6
- SlowTooltip: 40/50/60

## Quinn

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 565 | 107 | 2384 |
| dano de ataque | 59 | 3.2 | 113.4 |
| armadura | 28 | 4.7 | 107.899994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.668 | 3.1 | 53.368 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Rapina**

- BonusDamage: 120 + 0.4 attack_damage (bonus) _(no nível 18)_
- BonusMonsterDmg: 75
- RevealDuration: 4

**Q · Investida Anuviante** · recarga 11/10.5/10/9.5/9 · custo 50/55/60/65/70 · alcance 1050/1050/1050/1050/1050

- TotalDamage: 65/100/135/170/205 + 0.5 ability_power + 0.8 attack_damage (bonus)
- TotalDamageMonster (= TotalDamage × 1): 130/200/270/340/410 + 1 ability_power + 1.6 attack_damage (bonus)
- VisionReductionDuration: 1.75/1.75/1.75/1.75/1.75
- VisionReductionRadius: -1000/-1000/-1000/-1000/-1000

**W · Sentidos Apurados** · recarga 50/45/40/35/30 · alcance 2500/2500/2500/2500/2500

- AttackSpeedBonus: 0.28/0.41/0.54/0.67/0.8
- BuffDuration: 2/2/2/2/2
- CherryBonusHaste: 100/100/100/100/100
- MovespeedAmount: 0.2/0.25/0.3/0.35/0.4
- VisionDuration: 2/2/2/2/2
- VisionRadius: 2100/2100/2100/2100/2100

**E · Salto** · recarga 12/11/10/9/8 · custo 50/50/50/50/50 · alcance 600/600/600/600/600

- TotalDamage: 40/65/90/115/140 + 0.2 attack_damage (bonus)
- SlowAmount: 0.5/0.5/0.5/0.5/0.5
- SlowDecayTime: 1.5/1.5/1.5/1.5/1.5

**R · Retaguarda do Inimigo** · recarga 3/3/3 · custo 50/25/0 · alcance 650/650/650

- Damage: 60/90/120 + 0.35 attack_damage (bonus)
- HarrierCooldown: 7/7/7
- MovementSpeedMod: 0.7/1/1.3
- RHarrierMarkDuration: 4.5/4.5/4.5
- SlowDuration: 3/3/3

## Syndra

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 583 | 100 | 2283 |
| dano de ataque | 54 | 2.9 | 103.3 |
| armadura | 25 | 4 | 93 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 2 | 34.658 |
| regeneração de vida | 1.3 | 0.12 | 3.34 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Transcender**

- ManaPerProc: 20
- StacksPerProc: 3 _(no nível 18)_
- CapstoneAPPerc: 0.15
- EUpgradeThreshold: 80
- MarkDuration: 4
- MaxStackAmount: 120
- PassiveMarkCooldown: 8
- PassiveStacksPerLevel: 5
- Q1UpgradeThreshold: 40
- RUpgradeThreshold: 100
- StackPerSiege: 1
- WUpgradeThreshold: 60

**Q · Esfera Negra** · recarga 7/7/7/7/7 · custo 40/45/50/55/60 · alcance 800/800/800/800/800

- TotalDamage: 90/125/160/195/230 + 0.7 ability_power
- SphereDuration: 6/6/6/6/6
- Upgrade1MaxAmmo: 2/2/2/2/2
- Upgrade2MaxAmmo: 3/3/3/3/3

**W · Força de Vontade** · recarga 12/11/10/9/8 · custo 60/70/80/90/100 · alcance 925/925/925/925/925

- TOOLTIPONLYPassiveBonusPercent: 12/12/12/12/12 + 0.02 ability_power
- ThrowDamage: 70/100/130/160/190 + 0.65 ability_power
- TotalSlowAmount: 25/25/25/25/25
- ExtraSpherePickupRadius: 450/450/450/450/450
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**E · Dispersar os Fracos** · recarga 15/15/15/15/15 · custo 50/50/50/50/50 · alcance 650/650/650/650/650

- TotalDamage: 60/95/130/165/200 + 0.6 ability_power
- ConeAngle: 56/56/56/56/56
- StunDuration: 1.25/1.25/1.25/1.25/1.25
- UpgradedConeAngle: 84/84/84/84/84
- UpgradedSlowAmount: 0.7/0.7/0.7/0.7/0.7
- UpgradedSlowDuration: 1.25/1.25/1.25/1.25/1.25

**R · Poder Irrestrito** · recarga 120/100/80 · custo 100/100/100 · alcance 675/675/675

- DamageCalc: 80/120/160 + 0.2 ability_power
- MaxDamageCalc (= DamageCalc × 1): 560/840/1120 + 1.4 ability_power
- MinDamageCalc (= DamageCalc × 1): 240/360/480 + 0.6 ability_power
- QHastePerRank: 10/10/10
- TOOLTIPONLYTotalQHaste: 10/20/30
- UpgradeExecuteThreshold: 0.15/0.15/0.15

## Aurelion Sol

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 90 | 2130 |
| dano de ataque | 58 | 3.2 | 112.4 |
| armadura | 22 | 4 | 90 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 1.5 | 26.125 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Criador Cósmico** · recarga 0

- BaseExecutionThreshold: 5

**Q · Sopro de Luz** · recarga 3/3/3/3/3 · custo 35/40/45/50/55 · alcance 750/750/750/750/750

- BurstDamage: 60/70/80/90/100 + 0.3 ability_power
- DamagePerSecond: 45/60/75/90/105 + 0.55 ability_power
- LevelBasedRangeScaling: 750/750/750/750/750
- AOEModifier: 0.5/0.5/0.5/0.5/0.5
- AngularDegreesPerSec: 180/180/180/180/180
- BeamSegments: 6/6/6/6/6
- BeamWidth: 140/140/140/140/140
- BurstAfter: 1/1/1/1/1
- DistancePerMass: 0.375/0.375/0.375/0.375/0.375
- ManaCostPerSecond: 35/40/45/50/55
- MaxChannelDuration: 3.25/3.25/3.25/3.25/9999
- MonsterDamageCap: 300/300/300/300/300
- QMassStolen: 2/2/2/2/2
- WProcCooldown: 1/1/1/1/1

**W · Voo Astral** · recarga 22/20.5/19/17.5/16 · custo 50/55/60/65/70 · alcance 1500/1500/1500/1500/1500

- DashSpeed: 340/340/340/340/340 + 1 move_speed
- {70212d39} (= {ffd78035} × 0.5): 170/170/170/170/170 + 0.5 move_speed
- {e3ddb2ca}: 1.2/1.2/1.2/1.2/1.2 _(no nível 18)_
- {ffd78035}: 340/340/340/340/340 + 1 move_speed
- BaseCruiseMS: 340/340/340/340/340
- CD: 22/20.5/19/17.5/16
- DistancePerMass: 7.5/7.5/7.5/7.5/7.5
- FakeCastTime: 0.375/0.375/0.375/0.375/0.375
- OverrideECastRange: 1100/1100/1100/1100/1100
- ResetWindow: 3/3/3/3/3
- TakedownCooldownMultiplier: 0.1/0.1/0.1/0.1/0.1
- TooltipTakedownCooldownMultiplier: 90/90/90/90/90
- TrueDamageBonus: 0.08/0.09/0.1/0.11/0.12

**E · Singularidade** · recarga 12/12/12/12/12 · custo 90/90/90/90/90 · alcance 750/750/750/750/750

- DamagePerSecond: 10/15/20/25/30 + 0.12 ability_power
- LevelBasedRangeScaling: 750/750/750/750/750
- {7741d788}: 150/150/150/150/150 _(no nível 18)_
- ChampionCountBonus: 2/2/2/2/2
- ChampionMassPerSecond: 1/1/1/1/1
- Duration: 5/5/5/5/5
- EpicMonsterCountBonus: 2/2/2/2/2
- GrowthBreakPoint: 1/1/1/1/1
- InnerRadiusAreaPerStack: 180/180/180/180/180
- LargeMinionCountBonus: 2/2/2/2/2
- LargeMonsterCountBonus: 2/2/2/2/2
- MinionMassDeath: 1/1/1/1/1
- NonChampGravity: 300/300/300/300/300
- StartingInnerRadius: 120/120/120/120/120

**R · Estrela Cadente/Os Céus Caem** · recarga 120/110/100 · custo 100/100/100 · alcance 1250/1250/1250

- MaxDamageTooltip: 150/250/350 + 0.75 ability_power
- R2Damage (= MaxDamageTooltip × 1): 187.5/312.5/437.5 + 0.9375 ability_power
- ShockwaveDamage (= MaxDamageTooltip × 1): 135/225/315 + 0.67499995 ability_power
- CalamitySizeBonus: 2/2/2
- CalamityStacks: 75/75/75
- MassStolen: 5/5/5
- ShockwaveSlow: 0.5/0.5/0.5
- SkiesDescendStartingRadius: 550/550/550
- StunDuration: 1/1/1

## Kayn

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 655 | 103 | 2406 |
| dano de ataque | 68 | 3.3 | 124.1 |
| armadura | 38 | 4.5 | 114.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.669 | 2.7 | 46.569 |
| regeneração de vida | 1.6 | 0.15 | 4.15 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · A Foice Darkin**

- KaynSlayerHealing: 0.25 + 0.00005 max_health (bonus)
- KaynAssMaxValue: 0.4
- KaynAssMinValue: 0.2
- PAmpCooldownAss: 8
- PAmpDurationAss: 3
- PAmpPercentBaseAss: 18.823
- PAmpPercentPerLvlAss: 1.1777
- PassiveSecondFormDelayTooltip: 4
- TransformationVariable: 0.6

**Q · Corte Ceifador** · recarga 7/6.5/6/5.5/5 · custo 40/40/40/40/40

- DarkinFlatDamage: 0/0/0/0/0 + 0.65 attack_damage
- DarkinPercentDamage: 6/6/6/6/6 + 0.035 attack_damage (bonus)
- TotalDamage: 75/105/135/165/195 + 0.85 attack_damage (bonus)
- AoERadius: 300/300/300/300/300
- FlatBonusDmgToMonsters: 40/40/40/40/40
- MaxDmgToMonsters: 200/250/300/350/400
- SlayTADRatio: 0.65/0.65/0.65/0.65/0.65
- SlayerMaxHPPer100BAD: 3.5/3.5/3.5/3.5/3.5

**W · Alcance da Lâmina** · recarga 13/12/11/10/9 · custo 40/45/50/55/60 · alcance 700/700/700/700/700

- TotalDamage: 85/130/175/220/265 + 1.1 attack_damage (bonus)
- AssRange: 900/900/900/900/900
- BoxWidth: 160/160/160/160/160
- KnockupDuration: 1/1/1/1/1
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- SlowIntensity: -0.9/-0.9/-0.9/-0.9/-0.9

**E · Passo das Sombras** · recarga 21/19/17/15/13 · custo 90/90/90/90/90

- TotalHealing: 90/100/110/120/130 + 0.45 attack_damage (bonus)
- AssMS: 70/70/70/70/70
- AssassinCDReduction: 10/10/10/10/10
- HealBADRatio: 0.45/0.45/0.45/0.45/0.45
- LingerTime: 1.5/1.5/1.5/1.5/1.5
- MS: 40/40/40/40/40
- MaxInCombatTime: 1.5/1.5/1.5/1.5/1.5
- WallWalkDuration: 7/7.5/8/8.5/9
- WarningParticleRange: 1250/1250/1250/1250/1250

**R · Transgressão do Umbral** · recarga 120/100/80 · custo 100/100/100 · alcance 550/550/550

- Damage: 150/250/350 + 1.5 attack_damage (bonus)
- HealValue (= SlayerDamage × 1): 0.112500004/0.112500004/0.112500004 + 0.00075 attack_damage (bonus)
- SlayerDamage: 0.15/0.15/0.15 + 0.001 attack_damage (bonus)
- AssCastRange: 750/750/750
- AssassinCastRange: 750/750/750
- AssassinJumpOutDistance: 500/500/500
- BaseCastRange: 550/550/550
- InfestDuration: 2.5/2.5/2.5
- JumpOutDistance: 300/300/300
- MinimumInfestTime: 0.5/0.5/0.5

## Zoe

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 106 | 2432 |
| dano de ataque | 58 | 3.3 | 114.1 |
| armadura | 21 | 4.7 | 100.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 2.5 | 43.158 |
| regeneração de vida | 1.5 | 0.12 | 3.54 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Brililim!**

- PassiveDamage: 16 + 0.2 ability_power
- WPickupDurationChampion: 40
- WPickupDurationMinion: 20

**Q · Estrela Desviada!** · recarga 8.5/8/7.5/7/6.5 · custo 40/45/50/55/60

- MaxDamageTooltip (= TotalDamageTooltip × 2.5): 130/205/280/355/430 + 1.5 ability_power
- TotalDamageTooltip: 52/82/112/142/172 + 0.6 ability_power
- MinionDamageMod: 1/1/1/1/1

**W · Roubo Arcano** · recarga 0.25/0.25/0.25/0.25/0.25

- MissileDamageTooltip: 15/25/35/45/55 + 0.1 ability_power
- DurationOfSummonerCanCast: 60/60/60/60/60
- MSDuration: 2/2.25/2.5/2.75/3
- MoveSpeedDuration: 2/2.25/2.5/2.75/3
- MoveSpeedMod: 0.3/0.4/0.5/0.6/0.7
- MovementSpeed: 0.3/0.4/0.5/0.6/0.7

**E · Bolha do Soninho** · recarga 18/17/16/15/14 · custo 80/80/80/80/80 · alcance 800/800/800/800/800

- BreakDamageTooltip: 70/110/150/190/230 + 0.45 ability_power
- TotalDamageTooltip: 70/110/150/190/230 + 0.45 ability_power
- BrittleLinger: 1/1/1/1/1
- CooldownRefresh: 0.16/0.195/0.23/0.265/0.3
- DrowsyDuration: 1.4/1.4/1.4/1.4/1.4
- DrowsySlow: 0.1/0.15/0.2/0.25/0.3
- PercentPen: 0.3/0.3/0.3/0.3/0.3
- SleepDuration: 2.25/2.25/2.25/2.25/2.25
- TrapLife: 5/5/5/5/5

## Zyra

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 574 | 93 | 2155 |
| dano de ataque | 53 | 3.2 | 107.4 |
| armadura | 29 | 4.2 | 100.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.681 | 2.11 | 36.551 |
| regeneração de vida | 1.1 | 0.1 | 2.8 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 575 | — | 575 |
| multiplicador de crítico | 2 | — | 2 |

**P · Jardim de Espinhos** · recarga 8 · custo 70 · alcance 600

- PlantDamage: 75 + 0.2 ability_power _(no nível 18)_
- SeedCooldown: 9.05 _(no nível 18)_
- monstermod: 100 _(no nível 18)_
- {c951f1e1}: 50
- PlantBaseDmg: 12
- PlantDuration: 8
- PlantLvlScaling: 4

**Q · Farpas Mortais** · recarga 7/6.5/6/5.5/5 · custo 55/55/55/55/55 · alcance 800/800/800/800/800

- InitialDamage: 60/100/140/180/220 + 0.65 ability_power

**W · Crescimento Desenfreado** · recarga 0/0/0/0/0 · alcance 850/850/850/850/850

- KillAmmoRefundChamp: 1/1/1/1/1
- KillAmmoRefundMinion: 0.35/0.35/0.35/0.35/0.35
- SeedDuration: 60/60/60/60/60
- VisionGranted: 2/2/2/2/2

**E · Pântano das Raízes** · recarga 11/11/11/11/11 · custo 70/75/80/85/90 · alcance 1150/1150/1150/1150/1150

- TotalDamage: 60/95/130/165/200 + 0.6 ability_power
- MaxSlowStacks: 2/2/2/2/2
- PheromoneDuration: 4/4/4/4/4
- RootDuration: 1/1.25/1.5/1.75/2
- SlowAmountPlantAttack: 30/30/30/30/30
- SlowDurationPlantAttack: 2/2/2/2/2

**R · Espinhos Sufocantes** · recarga 110/100/90 · custo 100/100/100 · alcance 700/700/700

- TotalDamage: 200/300/400 + 0.7 ability_power
- EnragePlantsBonusDamage: 150/150/150
- EnragedBonusHealthPercent: 0.5/0.5/0.5
- KnockupDuration: 1/1/1
- PlantDamageBonus: 50/50/50

## Kai'Sa

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 102 | 2374 |
| dano de ataque | 59 | 2.6 | 103.2 |
| armadura | 25 | 4.2 | 96.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.644 | 1.8 | 31.244 |
| regeneração de vida | 0.8 | 0.11 | 2.67 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Segunda Pele**

- PBaseDamage: 30 + 0.12 ability_power _(no nível 18)_
- PCurrentPerStackDamage: 8 + 0.03 ability_power _(no nível 18)_
- PExecutePercentage: 0.15 + 0.0006 ability_power
- PAllyStacks: 1
- PCoefficient: 0.25
- PDamageCap: 400
- PDuration: 4
- PMaxStacks: 4

**Q · Chuva Icathiana** · recarga 10/9/8/7/6 · custo 55/55/55/55/55 · alcance 600/600/600/600/600

- MaxDamageDisplay: 150/206.25/262.5/318.75/375 + 0.75 ability_power + 2.0625 attack_damage (bonus)
- MaxDamageTotal: 90/123.75/157.5/191.25/225 + 0.45000002 ability_power + 1.2375001 attack_damage (bonus)
- TotalIndividualMissileDamage: 40/55/70/85/100 + 0.2 ability_power + 0.55 attack_damage (bonus)
- {b2bd0d2f}: 150/206.25/262.5/318.75/375 + 0.75 ability_power + 2.0625 attack_damage (bonus)
- MaxDamage: 90/123.75/157.5/191.25/225
- MaxDamageEvolved: 150/206.25/262.5/318.75/375
- MaxDamageEvolvedNL: 206.25/262.5/318.75/375/431.25
- MaxDamageNL: 123.75/157.5/191.25/225/258.75
- MinionDamageMultiplier: 2/2/2/2/2
- MinionHPThreshold: 0.35/0.35/0.35/0.35/0.35

**R · Instinto Assassino** · recarga 130/100/70 · custo 100/100/100

- RCalculatedShieldValue: 100/150/200 + 1.2 ability_power + 0.9 attack_damage
- RRange: 2000/2500/3000
- RShieldDuration: 2/2/2

## Seraphine

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 570 | 95 | 2185 |
| dano de ataque | 50 | 3 | 101 |
| armadura | 26 | 4.2 | 97.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.669 | 2 | 34.669 |
| regeneração de vida | 1.3 | 0.12 | 3.34 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Presença de Palco** · alcance 800

- AutoDamage: 25 + 0.04 ability_power _(no nível 18)_
- AllyNoteDamagePercent: 0.25
- AutoAPRatio: 0.075
- BonusAARange: 25
- DelayBetweenNoteCasts: 0.075
- DelayBetweenNoteCastsDecayRate: 0.95
- EchoCastDelay: 0.01
- MaxNotes: 4
- MinionDamageMultiplier: 1
- NoteDuration: 6
- NoteRadius: 100

**Q · Nota Aguda** · recarga 8/7.5/7/6.5/6 · custo 60/70/80/90/100 · alcance 900/900/900/900/900

- ExplosionDamage: 60/85/110/135/160 + 0.4 ability_power
- TotalEmpoweredDamage (= ExplosionDamage × 1.75): 105/148.75/192.5/236.25/280 + 0.7 ability_power
- DamageAmp: 75/75/75/75/75
- ExecuteThreshold: 0.25/0.25/0.25/0.25/0.25

**W · Som Envolvente** · recarga 22/22/22/22/22 · custo 70/75/80/85/90 · alcance 800/800/800/800/800

- HasteValueAllies (= WMSBonusTotal × 1): 0.080000006/0.080000006/0.080000006/0.080000006/0.080000006 + 0.00008 ability_power
- ShieldValueSeraphine (= TotalShieldStrength × 1): 60/80/100/120/140 + 0.2 ability_power
- TotalShieldStrength: 60/80/100/120/140 + 0.2 ability_power
- WMSBonusTotal: 0.2/0.2/0.2/0.2/0.2 + 0.0002 ability_power
- WMissingHPHeal: 8/10/12/14/16
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5
- WHealSplitDelay: 2.5/2.5/2.5/2.5/2.5

**E · Ritmo Contagiante** · recarga 11/10.5/10/9.5/9 · custo 60/60/60/60/60 · alcance 1300/1300/1300/1300/1300

- FinalDamage: 70/100/130/160/190 + 0.5 ability_power
- MinionDamageMod: 0.7/0.7/0.7/0.7/0.7
- SlowDuration: 1.1/1.2/1.3/1.4/1.5
- SlowValue: 99/99/99/99/99

**R · Bis** · recarga 160/140/120 · custo 100/100/100

- R1TotalDamage: 150/200/250 + 0.4 ability_power
- RChannelDuration: 1.25/1.5/1.75
- RMaxSlow: -0.99/-0.99/-0.99
- RMinSlow: -0.4/-0.4/-0.4
- RRange: 1200/1200/1200

## Gnar

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 540 | 79 | 1883 |
| dano de ataque | 60 | 3.2 | 114.4 |
| armadura | 32 | 3.7 | 94.9 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 6 | 102.625 |
| regeneração de vida | 0.9 | 0.25 | 5.15 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Fúria Genética**

- TotalAS: 0.99 _(no nível 18)_
- TotalAttackRange: 100 _(no nível 18)_
- TotalMegaGnarAD: 48.5 _(no nível 18)_
- TotalMegaGnarArmor: 54.5 _(no nível 18)_
- TotalMegaGnarMR: 63 _(no nível 18)_
- MegaHealthEndingValue: 831
- MegaHealthStartingValue: 100
- TiredDuration: 15

**Q · Bumerangue / Pedregulho** · recarga 16/14.5/13/11.5/10

- MegaTotalDamage: 45/90/135/180/225 + 1.4 attack_damage
- MiniTotalDamage: 5/45/85/125/165 + 1.25 attack_damage
- {4b371c72} (= MiniTotalDamage × 1): 2.5/22.5/42.5/62.5/82.5 + 0.625 attack_damage
- MegaAoERadius: 280/280/280/280/280
- MegaBoulderLifetime: 6/6/6/6/6
- MegaCDRefund: 0.7/0.7/0.7/0.7/0.7
- MegaPickupAndLolipopRadius: 200/200/200/200/200
- MegaSlowAmount: 0.3/0.35/0.4/0.45/0.5
- MegaSlowDuration: 2/2/2/2/2
- MiniCDRefund: 0.4/0.4/0.4/0.4/0.4
- MiniReturnRange: 3000/3000/3000/3000/3000
- SlowAmount: 0.15/0.2/0.25/0.3/0.35
- SlowDuration: 2/2/2/2/2

**W · Hiperativo / Safanão** · recarga 7/7/7/7/7

- MegaTotalDamage: 45/75/105/135/165 + 1 attack_damage
- MiniTotalDamage: 0/10/20/30/40 + 1 ability_power
- MegaLength: 550/550/550/550/550
- MegaStunDuration: 1.25/1.25/1.25/1.25/1.25
- MegaWidth: 200/200/200/200/200
- MiniHasteDuration: 3/3/3/3/3
- MiniMarkDuration: 3.5/3.5/3.5/3.5/3.5
- MiniMaxStacks: 2/2/2/2/2
- MiniMonsterCap: 300/300/300/300/300
- MiniPercentHPDamage: 0.06/0.08/0.1/0.12/0.14

**E · Pirueta / Encontrão** · recarga 22/19.5/17/14.5/12

- MegaTotalDamage: 80/115/150/185/220 + 0.06 max_health
- MiniTotalDamage: 50/85/120/155/190 + 0.06 max_health
- MegaDamageAoE: 375/375/375/375/375
- MegaRange: 675/675/675/675/675
- MegaSlowAoE: 200/200/200/200/200
- MiniASDuration: 6/6/6/6/6
- MiniBounceRange: 525/525/525/525/525
- MiniRange: 475/475/475/475/475
- MinibAS: 0.4/0.45/0.5/0.55/0.6
- MoveSpeedMod: -0.8/-0.8/-0.8/-0.8/-0.8
- SlowDuration: 0.5/0.5/0.5/0.5/0.5
- TravelTime: 0.6/0.6/0.6/0.6/0.6

**R · GNAR!** · recarga 90/60/30

- Damage: 200/300/400 + 1 ability_power + 0.5 attack_damage (bonus)
- WallDamage (= Damage × 1): 300/450/600 + 1.5 ability_power + 0.75 attack_damage (bonus)
- RAoESize: 475/475/475
- RCCDuration: 1.25/1.5/1.75
- RHyperMovementSpeedPercent: 40/60/80
- RKnockbackDistance: 500/500/500
- RSlowPercent: 45/45/45

## Zac

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 685 | 109 | 2538 |
| dano de ataque | 60 | 3.4 | 117.8 |
| armadura | 33 | 5.2 | 121.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.736 | 1.6 | 27.936 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Divisão Celular** · recarga 300

- HealPercent: 0.08 _(no nível 18)_
- ReviveBlobletDuration: 4 _(no nível 18)_
- TotalHeal (= HealPercent × 1): 0 + 0.08 max_health _(no nível 18)_
- ReviveCooldown: 300

**Q · Esticada** · recarga 14/12.5/11/9.5/8

- MaxDamageTooltip: 120/180/240/300/360 + 0.6 ability_power + 0.06 max_health (bonus)
- TotalDamage: 60/90/120/150/180 + 0.3 ability_power + 0.03 max_health (bonus)
- CollisionAoE: 300/300/300/300/300
- MaxSlamDistance: 700/700/700/700/700
- MissileRange: 800/800/800/800/800
- SlowAmount: -0.4/-0.4/-0.4/-0.4/-0.4
- SlowDuration: 0.5/0.5/0.5/0.5/0.5
- TetherDuration: 2.5/2.5/2.5/2.5/2.5
- TetherFalloffDistance: 900/900/900/900/900
- YankDistance: 300/300/300/300/300

**W · Matéria Instável** · recarga 5/5/5/5/5 · alcance 350/350/350/350/350

- DisplayPercentDamage: 4/5/6/7/8 + 0.03 ability_power
- BaseDamage: 40/50/60/70/80
- BaseMaxHealthDamage: 0.04/0.05/0.06/0.07/0.08
- MaxMinionDamage: 200/200/200/200/200
- PercentHealthDamage: 0.04/0.05/0.06/0.07/0.08
- RefundOnBlobPickup: 1/1/1/1/1

**E · Estilingue Elástico** · recarga 21/18/15/12/9 · alcance 300/300/300/300/300

- Damage: 60/105/150/195/240 + 0.8 ability_power
- MaxStun: 1/1/1/1/1
- ChannelTime: 0.9/1/1.1/1.2/1.3
- DistanceToSpeedRatio: 0.6/0.6/0.6/0.6/0.6
- KnockupRadius: 265/265/265/265/265
- MaxRange: 1200/1350/1500/1650/1800
- MaximumSpeed: 1350/1350/1350/1350/1350
- MinimumSpeed: 500/500/500/500/500
- VFXWarningTime: 1/1/1/1/1

**R · Vamos pular!** · recarga 120/105/90 · alcance 300/300/300

- DamagePerBounce: 120/190/260 + 0.4 ability_power
- DamagePerSubsequentBounce (= DamagePerBounce × 1): 60/95/130 + 0.2 ability_power
- DurationTooltip: 3/3/3
- MaxDamageTooltip: 300/475/650 + 1 ability_power
- BeginningMS: 0.2/0.2/0.2
- EndingMS: 0.5/0.5/0.5
- KnockbackDistance: 250/250/250
- KnockbackRange: 300/300/300
- KnockupDuration: 1/1/1
- SlowAmount: 0.2/0.2/0.2
- SlowDuration: 1/1/1

## Yasuo

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 110 | 2460 |
| dano de ataque | 60 | 2.5 | 102.5 |
| armadura | 32 | 4.6 | 110.2 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.697 | 3.5 | 60.197 |
| regeneração de vida | 1.3 | 0.18 | 4.36 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Estilo do Errante**

- CurrentCritDamage: 0 + 1 critical_damage
- ShieldValue: 600 _(no nível 18)_
- {74ce5438}: 45 _(no nível 18)_
- CritChanceMultiplier: 1
- CritDamageMod: 0.9
- YasuoCritToAD: 50

**Q · Tempestade de Aço** · recarga 4/4/4/4/4

- TotalDamage: 20/45/70/95/120 + 1.05 attack_damage
- Cooldown: 4/4/4/4/4
- GatheringStormDuration: 6/6/6/6/6
- KnockUpDurationTOOLTIPONLY: 1/1/1/1/1
- SpinWidth: 260/260/260/260/260
- StabRange: 450/450/450/450/450
- StabWidth: 80/80/80/80/80
- TornadoRange: 1100/1100/1100/1100/1100

**W · Parede de Vento** · recarga 25/23/21/19/17 · alcance 5000/5000/5000/5000/5000

- BubbleRadius: 300/300/300/300/300
- Thickness: 100/100/100/100/100
- TravelRange: 450/450/450/450/450
- WallLife: 4/4/4/4/4
- Width: 320/390/460/530/600

**E · Espada Ágil** · recarga 0.5/0.4/0.3/0.2/0.1 · alcance 475/475/475/475/475

- BonusDamagePerStack (= TotalDamage × 0.25): 17.5/21.25/25/28.75/32.5 + 0.15 ability_power + 0.05 attack_damage (bonus)
- TotalDamage: 70/85/100/115/130 + 0.6 ability_power + 0.2 attack_damage (bonus)
- BaseDashSpeed: 750/750/750/750/750
- MaxStacks: 4/4/4/4/4
- PerTargetCooldown: 10/9/8/7/6
- StackDuration: 5/5/5/5/5

**R · Último Suspiro** · recarga 70/50/30 · alcance 1400/1400/1400

- Damage: 200/350/500 + 1.5 attack_damage (bonus)
- RBuffDuration: 15/15/15
- RCastSearchRadius: 1100/1100/1100
- RKnockupDuration: 1/1/1
- RPercentArmorPen: 60/60/60
- RRadius: 400/400/400

## Vel'Koz

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 102 | 2324 |
| dano de ataque | 55 | 3.1415927 | 108.407074 |
| armadura | 22 | 4.7 | 101.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.643 | 1.59 | 27.673 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Desconstrução Orgânica**

- TotalDamage: 180 + 0.6 ability_power _(no nível 18)_
- Duration: 7
- MaxStacks: 3

**Q · Fissão Plasmática** · recarga 7/7/7/7/7 · custo 40/45/50/55/60

- TotalDamage: 80/120/160/200/240 + 0.9 ability_power
- ManaRefund: 0.5/0.5/0.5/0.5/0.5
- PassiveStacksToAdd: 1/1/1/1/1
- SlowAmount: 0.7/0.7/0.7/0.7/0.7
- SlowDuration: 1/1.4/1.8/2.2/2.6
- SplitTelegraphTime: 0.25/0.25/0.25/0.25/0.25
- TooltipManaRefund: 20/22.5/25/27.5/30

**W · Fenda do Vazio** · recarga 1.5/1.5/1.5/1.5/1.5 · custo 50/55/60/65/70

- InitialDamage: 30/50/70/90/110 + 0.2 ability_power
- SecondaryDamage: 45/75/105/135/165 + 0.25 ability_power
- PassiveStacksToAdd: 1/1/1/1/1

**E · Ruptura Tectônica** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70 · alcance 800/800/800/800/800

- TotalDamage: 70/100/130/160/190 + 0.3 ability_power
- MaxDetonationTime: 0.55/0.55/0.55/0.55/0.55
- MinDetonationTime: 0.25/0.25/0.25/0.25/0.25
- MoveDistance: 225/225/225/225/225
- PassiveStacksToAdd: 1/1/1/1/1
- StunDuration: 0.75/0.75/0.75/0.75/0.75

**R · Raio Desintegrador de Formas de Vida** · recarga 100/90/80 · custo 100/100/100 · alcance 1550/1550/1550

- PassiveStacksToAddPerTick: 1/1/1

## Taliyah

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 550 | 104 | 2318 |
| dano de ataque | 58 | 3.3 | 114.1 |
| armadura | 18 | 4.7 | 97.899994 |
| resistência mágica | 28 | 1.3 | 50.1 |
| velocidade de ataque | 0.658 | 1.36 | 23.778 |
| regeneração de vida | 1.3 | 0.13 | 3.5099998 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Deslizar em Pedras**

- TotalMS: 0.4 _(no nível 18)_
- FallOffTime: 3

**Q · Voleio Entrelaçado** · recarga 7/6/5/4/3 · custo 55/60/65/70/75 · alcance 1000/1000/1000/1000/1000

- BigRockDamage (= RockDamage × 1): 99/130.5/162/193.5/225 + 0.9 ability_power
- MaxDamageTooltip (= RockDamage × 2.6): 143/188.5/233.99998/279.5/325 + 1.3 ability_power
- RockDamage: 55/72.5/90/107.5/125 + 0.5 ability_power
- TotalBonusFlatMonsterDamage: 20/25/30/35/40
- {c363b92e} (= TotalBonusFlatMonsterDamage × 1): 36/45/54/63/72
- AoERadius: 175/175/175/175/175
- AoERadiusBig: 225/225/225/225/225
- BigRockManaCost: 10/10/10/10/10
- ExtraMissileReducedDamagePercent: 60/60/60/60/60
- GroundExhaustDuration: 30/30/30/30/30
- GroundExhaustRadius: 400/400/400/400/400
- MinimumWorkedGroundCD: 0.75/0.75/0.75/0.75/0.75
- MonsterStunDuration: 3/3/3/3/3
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- SlowPercent: 0.2/0.25/0.3/0.35/0.4
- WorkedGroundCDR: 0.5/0.5/0.5/0.5/0.5

**W · Empurrão Sísmico** · recarga 14/12.5/11/9.5/8 · custo 40/30/20/10/0 · alcance 900/900/900/900/900

- KnockupDelay: 0.5/0.5/0.5/0.5/0.5
- ThrowDistance: 400/400/400/400/400

**E · Terra Desfiada** · recarga 14/14/14/14/14 · custo 90/90/90/90/90 · alcance 950/950/950/950/950

- DetonationDamage: 25/40/55/70/85 + 0.3 ability_power
- MaxDetonationDamageTooltip (= DetonationDamage × 2.5): 62.5/100/137.5/175/212.5 + 0.75 ability_power
- ScatterDamage: 60/105/150/195/240 + 0.6 ability_power
- DelayBetweenRows: 0.17/0.17/0.17/0.17/0.17
- MaxStunDuration: 2/2/2/2/2
- MineDamageFalloff: 0.25/0.25/0.25/0.25/0.25
- MineLifetime: 4/4/4/4/4
- MineRadius: 85/85/85/85/85
- MonsterModPercent: 2.25/2.25/2.25/2.25/2.25
- SlowPercent: 0.2/0.2/0.2/0.2/0.2
- StunDuration: 0.75/0.75/0.75/0.75/0.75

**R · Muro da Tecelã** · recarga 180/150/120 · custo 100/100/100

- DCapPassive: 0.1/0.1/0.1
- DamageLockoutTime: 3/3/3
- ManaCost: 100/100/100
- MaxJumpRange: 700/700/700
- MaxJumpRangeOverWalls: 1000/1000/1000
- MinJumpRange: 150/150/150
- MissileSpeed: 2000/2000/2000
- WallDuration: 4/4/4
- WallLength: 2500/4500/6500

## Camille

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 99 | 2333 |
| dano de ataque | 68 | 3.8 | 132.6 |
| armadura | 35 | 4.5 | 111.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.644 | 2.5 | 43.144 |
| regeneração de vida | 1.7 | 0.16 | 4.42 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Defesa Adaptativa**

- PassiveCooldown: 8 _(no nível 18)_
- ShieldAmount: 0 + 0.2 max_health _(no nível 18)_
- ShieldDuration: 2

**Q · Protocolo de Precisão** · recarga 9/8/7/6/5 · custo 25/25/25/25/25 · alcance 325/325/325/325/325

- BonusDamage: 0/0/0/0/0 + 0.2 attack_damage
- DamageConversionPercentage: 0.4/0.4/0.4/0.4/0.4
- EmpoweredBonusDamage (= BonusDamage × 1): 0/0/0/0/0 + 0.4 attack_damage
- BonusAARange: 50/50/50/50/50
- MSBonus: 0.25/0.3/0.35/0.4/0.45
- MSDuration: 1/1/1/1/1
- Q2Duration: 2/2/2/2/2
- QRampUpTime: 1.5/1.5/1.5/1.5/1.5
- QTotalRecastTime: 3.5/3.5/3.5/3.5/3.5

**W · Varredura Tática** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70

- BaseDamageTotal: 60/85/110/135/160 + 0.6 attack_damage (bonus)
- MonsterHealthDamageCap: 300/300/300/300/300
- OuterEdgeTooltip: 0.07/0.075/0.08/0.085/0.09 + 0.00025 attack_damage (bonus)
- ADRequiredFor1PercentDamage: 40/40/40/40/40
- BlastLength: 650/650/650/650/650
- ChargeDuration: 0.75/0.75/0.75/0.75/0.75
- ConeAngle: 35/35/35/35/35
- MonsterDamageReduction: 50/50/50/50/50
- OuterConeHealingRatio: 100/100/100/100/100
- SlowDuration: 2/2/2/2/2
- SlowPercentage: 80/80/80/80/80

**E · Disparo de Gancho** · recarga 16/15/14/13/12 · custo 70/70/70/70/70

- TotalDamage: 60/90/120/150/180 + 0.75 attack_damage (bonus)
- ASBuff: 0.4/0.45/0.5/0.55/0.6
- ASDuration: 5/5/5/5/5
- DashSpeed: 1050/1050/1050/1050/1050
- E2LongDashRange: 800/800/800/800/800
- E2ShortDashRange: 400/400/400/400/400
- ECollisionRange: 130/130/130/130/130
- EMinionMonsterGhostDuration: 3/3/3/3/3
- KnockupDuration: 0.75/0.75/0.75/0.75/0.75
- WallHangDuration: 1/1/1/1/1

**R · Ultimato Hextec** · recarga 140/115/90 · custo 100/100/100 · alcance 475/475/475

- RCircleRadius: 425/425/425
- RDuration: 2.5/3.25/4
- RPercentCurrentHPDamage: 4/6/8

## Akshan

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 107 | 2429 |
| dano de ataque | 52 | 3 | 103 |
| armadura | 26 | 4.7 | 105.899994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.638 | 4 | 68.638 |
| regeneração de vida | 0.75 | 0.13 | 2.96 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Lutando Sujo**

- ASModdedMS (= MS × 1): 75 + 75 attack_speed (bonus) _(no nível 18)_
- MS: 75 _(no nível 18)_
- PassiveCooldown: 4 _(no nível 18)_
- PassiveProcDamage: 150 + 0.6 ability_power _(no nível 18)_
- SecondAutoDamage: 0 + 0.5 attack_damage
- TotalShieldAmount: 280 + 0.35 attack_damage (bonus) _(no nível 18)_
- DebuffDuration: 4.5
- HasteDuration: 1
- MaxStacks: 3
- SecondAutoFlexedAttackRange: 200
- ShieldDuration: 2

**Q · Bumerangue Vingativo** · recarga 9/8/7/6/5 · custo 60/65/70/75/80 · alcance 750/750/750/750/750

- FinalDamage: 45/75/105/135/165 + 0.7 attack_damage (bonus)
- TotalHaste: 0.2/0.2/0.2/0.2/0.2 + 0.0005 ability_power
- ExtensionDistance: 500/500/500/500/500
- HasteDuration: 1/1/1/1/1
- RevealDuration: 1/1/1/1/1
- SecondaryTargetDamage: 0.4/0.5/0.6/0.7/0.8

**W · Rebeldia** · recarga 18/14/10/6/2 · custo 40/30/20/10/0 · alcance 5000/5000/5000/5000/5000

- SecondAutoDamage: 0/0/0/0/0 + 0.46 attack_damage
- DeathTimer: 1/1/1/1/1
- GameModeInteger: 1/1/1/1/1
- GoldReward: 100/100/100/100/100
- MSValue: 80/90/100/110/120
- MissingManaRegen: 0.12/0.12/0.12/0.12/0.12
- PassiveWallRange: 175/175/175/175/175
- RecentlyDamagedWindow: 3/3/3/3/3
- ReviveHealthPercent: 1/1/1/1/1
- ShredValue: 15/17.5/20/22.5/25
- StealthDetectionRange: 800/800/800/800/800
- VFXTrailSpeed: 3/3/3/3/3
- VengeanceDuration: 60/60/60/60/60
- WallGracePeriod: 2/2/2/2/2

**E · Impulso Heroico** · recarga 18/16.5/15/13.5/12 · custo 70/70/70/70/70 · alcance 800/800/800/800/800

- DamageToDeal: 8/16/24/32/40 + 0.25 attack_damage
- {36ebd3cf}: 0/0/0/0/0 + 1 critical_damage
- AnchorLifetime: 3/3/3/3/3
- AttackFrequency: 0.2/0.2/0.2/0.2/0.2
- DistanceShiftTowardsWall: 175/175/175/175/175
- E3LockoutDuration: 0.5/0.5/0.5/0.5/0.5
- EdgeToEdgeCollisionRadius: 50/50/50/50/50
- MinJumpbackDistance: 250/250/250/250/250
- OnHitDamageReduction: 0.25/0.25/0.25/0.25/0.25
- ResetCooldownToSet: 0.5/0.5/0.5/0.5/0.5
- SpinDuration: 25000/25000/25000/25000/25000
- TimeToCastE2: 2/2/2/2/2

**R · Punição** · recarga 100/85/70 · custo 100/100/100 · alcance 2500/2500/2500

- DamagePerBulletWithCrit: 25/35/45 + 0.15 attack_damage
- MaxDamagePerBullet (= DamagePerBulletWithCrit × 1): 75/105/135 + 0.45000002 attack_damage
- CastTimePerBullet: 0.15/0.15/0.15
- ChannelDuration: 2.5/2.5/2.5
- GraceCooldown: 5/5/5
- MinimumChannelDuration: 0.5/0.5/0.5
- NumberOfBullets: 5/6/7

## Bel'Veth

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 580 | 105 | 2365 |
| dano de ataque | 55 | 1.5 | 80.5 |
| armadura | 28 | 5 | 113 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.67 | 0 | 0.67 |
| regeneração de vida | 1.2 | 0.12 | 3.24 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Morte em Lavanda **

- SheenSpeedPerStack: 0.2
- ARAMBonusStacks: 1
- ASPerStackInitialBonusPerLevel: 0.05
- ASPerStackLevel11BonusPerLevel: 0.15
- ASPerStackLevel1Value: 0.1
- ASPerStackLevel6BonusPerLevel: 0.1
- AttackADRatio: 1
- BaronBuffDuration: 180
- BrawlBonusStacks: 2
- ChampionStacks: 2
- MonsterStacks: 1
- OnHitRatio: 1
- SheenDuration: 3
- SheenNumberOfAttacks: 2
- TakedownDuration: 4

**Q · Impulso do Vazio** · recarga 4/3.25/2.5/1.75/1 · alcance 450/450/450/450/450

- BaseDamage: 12/14/16/18/20 + 1.05 attack_damage
- {290ee56e}: 0/0/0/0/0 + 1 attack_speed
- AttackSpeed: 0.4/0.4/0.4/0.4/0.4
- AttackSpeedDuration: 2/2/2/2/2
- BaseDashSpeed: 850/850/850/850/850
- BaseHealthSteal: 0.01/0.02/0.03/0.04/0.05
- CDReduction: 0.5/0.5/0.5/0.5/0.5
- DashDistance: 400/400/400/400/400
- MaxDistanceOverWalls: 625/625/625/625/625
- MinonMod: 1/1/1/1/1
- MonsterMod: 70/70/70/70/70
- NumberOfAttacks: 2/2/2/2/2
- PerSideCDAttackSpeedMultiplier: 0.2/0.2/0.2/0.2/0.2
- PerSideCooldown: 16/15/14/13/12
- RiftDashSpeed: 1500/1500/1500/1500/1500
- SpeedDuration: 4/4/4/4/4

**W · Acima e Abaixo** · recarga 12/11/10/9/8 · alcance 715/715/715/715/715

- Damage: 80/140/200/260/320 + 1.5 ability_power (bonus)
- Duration: 0.6/0.7/0.8/0.9/1
- SlowDuration: 2/2/2/2/2
- SlowPercent: 0.3/0.3/0.3/0.3/0.3

**E · Turbilhão da Realeza** · recarga 24/21/18/15/12 · alcance 500/500/500/500/500

- DamagePerStrike: 10/12/14/16/18 + 0.12 attack_damage
- MaxDamagePerStrikeTooltip (= DamagePerStrike × 2): 20/24/28/32/36 + 0.24 attack_damage
- TotalLifesteal: 0.2/0.25/0.3/0.35/0.4
- TotalStrikes: 6/6/6/6/6 + 2.4 attack_speed (bonus)
- AttackSpeedMult: 8/8/8/8/8
- DRPercent: 0.2/0.3/0.4/0.5/0.6
- MissingHealthMult: 1/1/1/1/1
- MonsterMod: 2/2/2/2/2
- OnHitRatio: 0.12/0.12/0.12/0.12/0.12
- TotalDuration: 1.5/1.5/1.5/1.5/1.5

**R · Banquete Eterno** · recarga 1/1/1 · alcance 450/450/450

- FinalOnHitDamage: 2/4/6 + 0.03 attack_damage (bonus)
- MaxHealthOnDevour: 100/250/400 + 1.5 ability_power + 1.5 attack_damage (bonus)
- MaxMonsterOnHitTooltip (= FinalOnHitDamage × 1): 16/32/48 + 0.24 attack_damage (bonus)
- TotalExplosionDamage: 150/200/250 + 1.5 ability_power
- AoERange: 500/500/500
- BonusAARange: 25/75/125
- CorpseDuration: 15/15/15
- MaxExecuteVsMonsters: 1500/1500/1500
- MeleeVoidlingSkinScaleMod: 0.15/0.15/0.15
- MinionRadius: 1000/1000/1000
- MissingHealthDamage: 0.2/0.2/0.2
- PassiveStacksOnDevour: 1/1/1
- RPassiveStackDuration: 5/5/5
- SizeIncreasePercent: 0.05/0.13/0.21
- StackThresholdForPermanent: 80/80/80
- StackThresholdForUpgrade: 40/40/40
- SteroidDuration: 45/45/45
- SteroidDurationUpgrade: 90/90/90
- TotalASMod: 0.06/0.13/0.2
- VoidDuration: 180/180/180
- VoidlingADScale: 1.1/1.1/1.1
- VoidlingHPScale: 0.2/0.45/0.7
- VoidlingMovementSpeedMod: 100/100/100

## Braum

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 112 | 2514 |
| dano de ataque | 55 | 3.2 | 109.4 |
| armadura | 35 | 5 | 120 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.644 | 3.5 | 60.144 |
| regeneração de vida | 1.7 | 0.2 | 5.1 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Golpes Concussivos**

- OnHitDamage (= TotalDamage × 1): 78.4 _(no nível 18)_
- StunCD: 4 _(no nível 18)_
- StunDuration: 1.75 _(no nível 18)_
- TotalDamage: 196 _(no nível 18)_
- StackCap: 4
- StackDuration: 4

**Q · Mordida do Inverno** · recarga 8/7.5/7/6.5/6 · custo 45/50/55/60/65

- TotalDamage: 75/120/165/210/255 + 0.025 max_health
- DebuffDuration: 4/4/4/4/4
- InitialSlow: 70/70/70/70/70
- MinSlow: 30/30/30/30/30
- MissileRange: 1050/1050/1050/1050/1050
- SlowDuration: 2/2/2/2/2

**W · Eu te Protejo** · recarga 12/11/10/9/8 · custo 40/40/40/40/40 · alcance 650/650/650/650/650

- GrantedAllyArmor: 20/25/30/35/40 + 0.12 armor (bonus)
- GrantedAllyMR: 20/25/30/35/40 + 0.12 magic_resist (bonus)
- GrantedBraumArmor: 20/25/30/35/40 + 0.36 armor (bonus)
- GrantedBraumMR: 20/25/30/35/40 + 0.36 magic_resist (bonus)
- BonusDashSpeed: 750/750/750/750/750
- Duration: 3/3/3/3/3
- RangedDamageThreshold: 400/400/400/400/400

**E · Inquebrável** · recarga 16/14/12/10/8 · custo 30/35/40/45/50

- MoveSpeedPercent: 10/10/10/10/10
- RangedReflect: 0.15/0.15/0.15/0.15/0.15
- ShieldFacingDRAmount: 35/40/45/50/55
- ShieldHoldDuration: 3/3.25/3.5/3.75/4

**R · Fissura Glacial** · recarga 130/115/100 · custo 100/100/100

- TotalDamage: 150/250/350 + 0.6 ability_power
- FirstKnockupDuration: 1/1.5/2
- KnockupDuration: 0.6/0.6/0.6
- MaxKnockup: 1/1.5/2
- MaxKnockupRatio: 0.9/0.9/0.9
- MinKnockup: 0.6/0.6/0.6
- MoveSpeedMod: 40/50/60
- SlowDebuffDuration: 0.25/0.25/0.25
- SlowZoneDuration: 4/4/4

## Jhin

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 655 | 107 | 2474 |
| dano de ataque | 61 | 4.4 | 135.8 |
| armadura | 24 | 4.7 | 103.899994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.625 | 0 | 0.625 |
| regeneração de vida | 0.75 | 0.11 | 2.62 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Sussurro**

- CritMoveSpeedPercent: 0.14 + 0.44 attack_speed (bonus)
- FourthShotExecutePercent: 0.25 _(no nível 18)_
- TotalADPercent: 0.04 + 0.3 attack_speed (bonus) + 0.35 critical_chance
- BaseAttackSpeed: 0.625
- CritReductionPercent: 0.25
- FourthShotDamageMult: 1.5
- HasteDuration: 2
- MaxAmmo: 4
- OutOfCombatTimeBeforeReload: 8
- PercentAttackSpeedPerLevel: 0.03
- ReloadTime: 2.5

**Q · Granada Dançante** · recarga 7/6.5/6/5.5/5 · custo 40/45/50/55/60 · alcance 550/550/550/550/550

- TooltipMaxTargetsHit: 4/4/4/4/4
- TotalDamage: 44/69/94/119/144 + 0.6 ability_power + 0.44 attack_damage
- BounceRange: 450/450/450/450/450
- NumberOfBounces: 3/3/3/3/3
- PercentAmpOnKill: 0.35/0.35/0.35/0.35/0.35

**W · Florescer Mortal** · recarga 12/12/12/12/12 · custo 50/55/60/65/70

- TotalDamage: 70/105/140/175/210 + 0.5 attack_damage
- MinionMod: 0.75/0.75/0.75/0.75/0.75
- RootDuration: 1.25/1.5/1.75/2/2.25
- SpottingDuration: 4/4/4/4/4

**E · Audiência Cativa** · recarga 2/2/2/2/2 · custo 30/30/30/30/30 · alcance 750/750/750/750/750

- TotalDamage: 20/80/140/200/260 + 1 ability_power + 1.2 attack_damage
- ReducedDamagePercent: 0.65/0.65/0.65/0.65/0.65
- TrapAoERadius: 260/260/260/260/260
- TrapArmTime: 0.75/0.75/0.75/0.75/0.75
- TrapDetonationTime: 2/2/2/2/2
- TrapDuration: 3/3/3/3/3
- TrapRevealDuration: 4/4/4/4/4
- TrapSlowAmount: 0.35/0.35/0.35/0.35/0.35
- TrapTriggerRadius: 160/160/160/160/160

**R · Aclamação** · recarga 120/105/90 · custo 100/100/100

- DamageCalc: 64/128/192 + 0.25 attack_damage
- MaxIncreaseCalc (= DamageCalc × 1): 256/512/768 + 1 attack_damage
- FourthShotMultiplier: 2/2/2
- MInimumDelayBetweenShots: 1/1/1
- SlowDuration: 0.5/0.5/0.5
- SlowPercent: 0.8/0.8/0.8

## Kindred

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 595 | 104 | 2363 |
| dano de ataque | 65 | 3.25 | 120.25 |
| armadura | 29 | 4.7 | 108.899994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.625 | 3.5 | 60.125 |
| regeneração de vida | 1.4 | 0.11 | 3.27 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Marca Familiar**

- {d34fc902}: 25
- ASPerMark: 0.05
- AdditionalMarkThreshold: 3
- BaseHealAmount: 45
- EDamagePerMark: 0.5
- HealAmountPerLevel: 2
- InitialMarkThreshold: 4

**Q · Dança de Flechas** · recarga 9/9/9/9/9 · custo 35/35/35/35/35

- TotalDamage: 40/65/90/115/140 + 0.75 attack_damage (bonus)
- BaseASDuration: 4/4/4/4/4
- CDNewValue: 4/3.5/3/2.5/2
- DashSpeed: 500/500/500/500/500
- DashSpeedScaling: 12/12/12/12/12

**W · Frenesi do Lobo** · recarga 18/17/16/15/14 · custo 40/40/40/40/40

- AttackHeal: 47/47/47/47/47
- BaseWolfDamage: 25/30/35/40/45 + 0.2 ability_power + 0.2 attack_damage (bonus)
- BaseCurrentHealthDamage: 1.5/1.5/1.5/1.5/1.5
- LambToWolfAttackSpeedConversionPercent: 0.25/0.25/0.25/0.25/0.25
- MonsterBonusDmg: 0.5/0.5/0.5/0.5/0.5
- MonsterSlowAmount: 0.5/0.5/0.5/0.5/0.5
- RingRange: 800/800/800/800/800
- ZoneDuration: 8.5/8.5/8.5/8.5/8.5

**E · Pesar Crescente** · recarga 14/12.5/11/9.5/8 · alcance 1200/1200/1200/1200/1200

- BaseBiteDamage: 80/110/140/170/200 + 1 attack_damage (bonus)
- TotalSlow: 30/30/30/30/30 + 0.05 ability_power
- MonsterCap: 200/200/200/200/200
- SlowDuration: 1/1/1/1/1
- StacksToProc: 4/4/4/4/4
- TotalDuration: 4/4/4/4/4

**R · Refúgio da Ovelha** · recarga 160/140/120 · custo 100/100/100 · alcance 500/500/500

- AoERadius: 530/530/530
- BuffDuration: 4/4/4
- HealFlat: 225/300/375

## Zeri

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 110 | 2470 |
| dano de ataque | 56 | 2 | 90 |
| armadura | 24 | 4.2 | 95.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.658 | 2 | 34.658 |
| regeneração de vida | 0.65 | 0.14 | 3.03 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Bateria Viva**

- PassiveAttackSpeed: 0.3 _(no nível 18)_
- ChampHitHasteDuration: 1.5
- EnergizedDuration: 2
- HastePercent: 0.1
- PercentShieldSteal: 0.45
- ShieldDuration: 3

**Q · Rajada Reluzente** · recarga 0/0/0/0/0

- ActiveDamageThatCanCrit: 22/26/30/34/38 + 1.02 attack_damage
- CurrentCritDamage: 0/0/0/0/0 + 1 critical_damage
- MinDamage: 25/25/25/25/25 + 0.03 ability_power _(no nível 18)_
- PassiveExecuteThreshold: 160/160/160/160/160 + 0.2 ability_power _(no nível 18)_
- PassiveMaxChargePercentHealth: 0.11/0.11/0.11/0.11/0.11 _(no nível 18)_
- PassiveMaxDamage: 160/160/160/160/160 + 1.1 ability_power _(no nível 18)_
- ActivePercentDamageToTowers: 1/1/1/1/1
- AttackSpeedCap: 1.5/1.5/1.5/1.5/1.5
- ChargeDistanceMultiplier: 0.025/0.025/0.025/0.025/0.025
- ChargePerAttack: 10/10/10/10/10
- EnergizedMissiles: 3/3/3/3/3
- ExcessAttackSpeedToADMult: 0.6/0.6/0.6/0.6/0.6
- MaxPercentHealthToMonsters: 300/300/300/300/300
- MinCooldown: 0.5/0.5/0.5/0.5/0.5
- NumberOfMissiles: 7/7/7/7/7
- SlowDuration: 0.25/0.25/0.25/0.25/0.25
- SlowPercent: 0.99/0.99/0.99/0.99/0.99

**W · Laser de Ultrachoque** · recarga 12/11/10/9/8 · custo 50/60/70/80/90 · alcance 1200/1200/1200/1200/1200

- TotalDamage: 30/70/110/150/190 + 0.5 ability_power + 1.2 attack_damage
- ASpeedCastTimeScalarPerHundrethSecond: 0.09/0.09/0.09/0.09/0.09
- BaseCastTime: 0.55/0.55/0.55/0.55/0.55
- BeamCastTime: 0.85/0.85/0.85/0.85/0.85
- BeamWidth: 200/200/200/200/200
- MinCastTime: 0.3/0.3/0.3/0.3/0.3
- SlowDuration: 2/2/2/2/2
- SlowPercent: 0.3/0.35/0.4/0.45/0.5
- WallBeamRange: 1500/1500/1500/1500/1500

**E · Faísca Acelerada** · recarga 24/22.5/21/19.5/18 · custo 90/85/80/75/70

- BonusDamageTotal: 22/24/26/28/30 + 0.2 ability_power
- DashSpeed: 600/600/600/600/600 + 1 move_speed
- {93e3e1a0}: 600/600/600/600/600 + 1 move_speed
- BuffDuration: 5/5/5/5/5
- CDReductionModNonChamp: 1/1/1/1/1
- CDReductionPerHit: 0.5/0.5/0.5/0.5/0.5
- CritCDReductionPerHit: 1.5/1.5/1.5/1.5/1.5
- IndicatorRevealTime: 0.75/0.75/0.75/0.75/0.75
- MaxDistance: 300/300/300/300/300
- MaxDistanceOverWalls: 3000/3000/3000/3000/3000
- MediumHopAnimCutoff: 550/550/550/550/550
- NumberOfPierceShots: 3/3/3/3/3
- PenDamagePercent: 0.8/0.85/0.9/0.95/1
- RevealRangeOverWalls: 1500/1500/1500/1500/1500
- ShortHopAnimCutoff: 250/250/250/250/250

**R · Impacto Eletrizante** · recarga 80/75/70 · custo 100/100/100

- ChainPhysicalDamage: 0/0/0 + 0.4 attack_damage
- TotalActiveDamage: 150/250/350 + 1.1 ability_power + 0.6 attack_damage (bonus)
- TotalBonusDamage: 5/10/15 + 0.15 ability_power
- BaseASPercent: 0.3/0.3/0.3
- BaseBonusMS: 0.15/0.15/0.15
- DurationIncreasePerHit: 2.5/2.5/2.5
- MSPercent: 0.015/0.015/0.015
- MaxChainTargets: 5/5/5
- MaxHyperchargeDuration: 2.5/2.5/2.5
- MaxStacks: 100000/100000/100000
- NovaRange: 825/825/825
- RDuration: 5/5/5
- StackDuration: 5/5/5

## Jinx

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 105 | 2415 |
| dano de ataque | 59 | 3.25 | 114.25 |
| armadura | 26 | 4.2 | 97.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.625 | 1 | 17.625 |
| regeneração de vida | 0.75 | 0.1 | 2.45 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Anime-se!**

- ASBuff: 25
- AssistMarkerDuration: 3
- BuffDuration: 6
- MSBuff: 175
- MSDecayRate: 0.875

**Q · Trocando!** · recarga 0.9/0.9/0.9/0.9/0.9 · custo 20/20/20/20/20 · alcance 600/600/600/600/600

- RocketDamage: 0/0/0/0/0 + 1.1 attack_damage
- MinigunAttackSpeedDuration: 2.5/2.5/2.5/2.5/2.5
- MinigunAttackSpeedMax: 30/55/80/105/130
- MinigunAttackSpeedStacks: 3/3/3/3/3
- RocketASPDPenalty: 0.1/0.1/0.1/0.1/0.1
- RocketAoERadius: 250/250/250/250/250
- RocketBonusRange: 100/125/150/175/200

**W · Zap!** · recarga 8/7/6/5/4 · custo 40/45/50/55/60

- TotalDamage: 10/60/110/160/210 + 1.4 attack_damage
- JinxLowEndWCastTime: 0.4/0.4/0.4/0.4/0.4
- JinxWASpeedCastTimeScalarPerHundrethSecond: 0.08/0.08/0.08/0.08/0.08
- SlowDuration: 2/2/2/2/2
- SlowPercent: 40/50/60/70/80

**E · Mordidinha Flamejante!** · recarga 24/20.5/17/13.5/10 · custo 90/90/90/90/90 · alcance 925/925/925/925/925

- TotalDamage: 90/140/190/240/290 + 1 ability_power
- GrenadeArmTime: 0.5/0.5/0.5/0.5/0.5
- GrenadeDuration: 5/5/5/5/5
- RootDuration: 1.5/1.5/1.5/1.5/1.5

**R · Super Mega Míssil da Morte!** · recarga 85/65/45 · custo 100/100/100

- DamageFloor: 20/35/50 + 0.12 attack_damage (bonus)
- DamageMax: 200/350/500 + 1.2 attack_damage (bonus)
- AoEDamageMult: 0.8/0.8/0.8
- AoERadius: 400/400/400
- MonsterExecuteMax: 1200/1200/1200
- PercentDamage: 25/30/35

## Tahm Kench

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 103 | 2391 |
| dano de ataque | 56 | 3.2 | 110.4 |
| armadura | 39 | 4.7 | 118.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 2.5 | 43.158 |
| regeneração de vida | 1.3 | 0.11 | 3.1699998 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Um Gosto Adquirido**

- Duration: 5
- MaxStacks: 3
- PerStackDecayTimer: 1

**Q · Língua-chicote** · recarga 7/6.5/6/5.5/5 · custo 50/46/42/38/34 · alcance 900/900/900/900/900

- TotalDamage: 75/120/165/210/255 + 1 ability_power
- BaseHeal: 10/15/20/25/30
- PercentHealthHealing: 0.05/0.055/0.06/0.065/0.07
- SlowAmount: 0.5/0.5/0.5/0.5/0.5
- SlowDuration: 2/2/2/2/2
- StunDuration: 1.5/1.5/1.5/1.5/1.5

**W · Mergulho Abissal** · recarga 21/20/19/18/17 · custo 60/75/90/105/120 · alcance 1000/1050/1100/1150/1200

- TotalDamage: 100/135/170/205/240 + 1.5 ability_power
- AoEDelayFromChannelStart: 1.5/1.5/1.5/1.5/1.5
- ChampRefund: 0.4/0.425/0.45/0.475/0.5
- ChannelTime: 1.35/1.35/1.35/1.35/1.35
- EnemyWarningDelayFromChannelStart: 0.75/0.75/0.75/0.75/0.75
- KnockUpDuration: 1/1/1/1/1
- TotalTime: 2/2/2/2/2

**E · Pele Grossa** · recarga 3/3/3/3/3 · alcance 2400/2400/2400/2400/2400

- GreyHealthHealingRatio: 1/1/1/1/1 _(no nível 18)_
- GreyHealthMaximum: 0/0/0/0/0 + 3 max_health
- EnhancedThreshold: 2/2/2/2/2
- GreyHealthRatio: 0.15/0.23/0.31/0.39/0.47
- GreyHealthRatioEnhanced: 0.42/0.44/0.46/0.48/0.5
- MaxHPHealRatio: 0.1/0.1/0.1/0.1/0.1
- OOCTimer: 4/4/4/4/4
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5
- ThresholdRadius: 2400/2400/2400/2400/2400

**R · Devorar** · recarga 0/0/0 · custo 100/100/100

- CDCalc: 120/100/80
- PercentHPDamage: 0.15/0.15/0.15 + 0.0007 ability_power
- TotalShield: 650/800/950 + 1 ability_power
- AllyDuration: 3/3/3
- AllySpeedAmount: 0.6/0.6/0.6
- BaseDamage: 100/250/400
- CastTime: 0.25/0.25/0.25
- DataManaCost: 100/100/100
- EnemyDuration: 3/3/3
- MinimumAllyDuration: 1/1/1
- ShieldDecayPerSecond: 200/200/200
- SlowAmount: 0.4/0.4/0.4
- SpitOutRangeAlly: 300/300/300
- SpitOutRangeEnemy: 300/300/300

## Briar

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 625 | 95 | 2240 |
| dano de ataque | 60 | 2.5 | 102.5 |
| armadura | 30 | 4.5 | 106.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.644 | 2 | 34.644 |
| regeneração de vida | 0 | 0 | 0 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Maldição Carmesim**

- BleedDamageOverDurationTooltip (= {15da76ad} × 1): 50 + 0.5 attack_damage (bonus) _(no nível 18)_
- BleedMaxDamageOverDurationTooltip (= {15da76ad} × 1): 100 + 1 attack_damage (bonus) _(no nível 18)_
- TotalHealPerMissingHPPercentTooltip (= {b38f3487} × 100): 40 + 0.025 max_health (bonus)
- {15da76ad}: 10 + 0.1 attack_damage (bonus) _(no nível 18)_
- {b38f3487}: 0.4 + 0.00025 max_health (bonus)
- {f06e5a79} (= {15da76ad} × 3): 30 + 0.3 attack_damage (bonus) _(no nível 18)_
- BleedPercentAdd: 0.25
- BleedTickRate: 0.5
- CurrentHealthPercentCost: 0.05
- HealPercent: 0.25
- MaxBleedStacks: 5
- PercentOfBleedHealedOnKill: 1.25

**Q · Vertigem** · recarga 13/12/11/10/9 · alcance 475/475/475/475/475

- TotalDamage: 60/85/110/135/160 + 0.6 ability_power + 0.8 attack_damage (bonus)
- DashSpeedBase: 600/600/600/600/600
- ExtraDashSpeedBasedOnDistance: 300/300/300/300/300
- ShredDuration: 5/5/5/5/5
- ShredPercent: 0.1/0.125/0.15/0.175/0.2
- StunDuration: 0.85/0.85/0.85/0.85/0.85

**W · Frenesi Sanguinário/Ataque Faminto** · recarga 14/13/12/11/10 · alcance 350/350/350/350/350

- AttackMaxHPHeal: 0/0/0/0/0 + 0.05 max_health
- TotalAoEDamage: 0/0/0/0/0 + 0.6 attack_damage
- TotalAttackBonusDamage: 5/20/35/50/65 + 1.05 attack_damage
- TotalAttackPercentMissingHealth: 9/9/9/9/9 + 0.025 attack_damage (bonus)
- AoEAttackRadius: 275/275/275/275/275
- AttackHealPercent: 0.24000001/0.28/0.32/0.36/0.4
- BerserkAS: 0.55/0.65/0.75/0.85/0.95
- BerserkDuration: 5/5/5/5/5
- BerserkMS: 0.24/0.33/0.42/0.51/0.6
- DashSpeed: 1200/1200/1200/1200/1200
- MaxDashDistance: 300/300/300/300/300
- MaxMonsterDamage: 400/400/400/400/400
- MaxRangeOverWalls: 650/650/650/650/650
- MinDashDistance: 75/75/75/75/75
- MonsterAndMinionPercentMod: 0.1/0.1/0.1/0.1/0.1
- TauntRange: 1000/1000/1000/1000/1000

**E · Grito Arrepiante** · recarga 16/16/16/16/16

- Damage: 80/115/150/185/220 + 1 ability_power + 1 attack_damage (bonus)
- PercentMaxHPHeal: 0/0/0/0/0 + 0.1 max_health
- WallHitDamage: 140/215/290/365/440 + 2.4 ability_power + 2.4 attack_damage (bonus)
- DRPercent: 35/35/35/35/35
- KnockbackDistance: 575/575/575/575/575
- MaxMissileRange: 600/600/600/600/600
- MinMissileRange: 400/400/400/400/400
- RKnockTowardsSpeed: 1800/1800/1800/1800/1800
- RKnockupDuration: 0.5/0.5/0.5/0.5/0.5
- RKnockupHeight: 2/2/2/2/2
- RectangleDistanceBehindCaster: 100/100/100/100/100
- RectangleDistanceInFrontOfCaster: 10/10/10/10/10
- SlowDuration: 0.5/0.5/0.5/0.5/0.5
- SlowPercent: 0.8/0.8/0.8/0.8/0.8
- WallStunDuration: 1.5/1.5/1.5/1.5/1.5
- Width: 380/380/380/380/380

**R · Morte Certa** · recarga 120/100/80

- Damage: 150/250/350 + 1.3 ability_power
- TotalResists: 0/0/0 + 0.2 attack_damage
- AoERadius: 575/575/575
- ExtraMoveSpeedPercent: 0.1/0.2/0.3
- FearDuration: 1.5/1.5/1.5
- FleeSlowPercent: 0.35/0.35/0.35
- HuntDuration: 25000/25000/25000
- IdealDashTime: 0.5/0.5/0.5
- LifestealPercent: 0.1/0.15/0.2
- MaxDashSpeed: 5000/5000/5000
- MinDashSpeed: 2500/2500/2500

## Viego

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 109 | 2483 |
| dano de ataque | 57 | 3.5 | 116.5 |
| armadura | 34 | 4.6 | 112.2 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.658 | 2.75 | 47.408 |
| regeneração de vida | 1.4 | 0.14 | 3.78 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 200 | — | 200 |
| multiplicador de crítico | 2 | — | 2 |

**P · Dominação Monárquica**

- PercentHealthHeal: 2 + 0.02 ability_power + 0.025 attack_damage (bonus) + 5 attack_speed (bonus)
- InvulnTimeWhileAttackingSoul: 1
- MoveSpeedPercent: 0.1
- SoulDuration: 8
- TakedownWindow: 3
- TransformDuration: 10

**Q · Espada do Rei Destruído** · recarga 5/4.5/4/3.5/3 · alcance 600/600/600/600/600

- HealthCritDamage: 0.3/0.3/0.3/0.3/0.3 + 0.7 critical_damage
- SecondAttackDamage: 0/0/0/0/0 + 0.15 ability_power + 0.2 attack_damage
- TotalDamage: 25/40/55/70/85 + 0.7 attack_damage
- TotalPercentHealthOnHit: 2/3/4/5/6
- HealModVsChamps: 1.5/1.5/1.5/1.5/1.5
- HealModVsMinions: 1/1/1/1/1
- HealModVsMonsters: 1.25/1.25/1.25/1.25/1.25
- MarkDuration: 4/4/4/4/4
- MinDamageOnHit: 10/15/20/25/30
- MonsterCapOnHit: 80/80/80/80/80
- RectangleRange: 600/600/600/600/600
- RectangleWidth: 125/125/125/125/125

**W · Posse Espectral** · recarga 8/8/8/8/8

- TotalDamage: 80/135/190/245/300 + 1 ability_power
- CDWheninterrupted: 3/3/3/3/3
- DashDistance: 300/300/300/300/300
- DashSpeed: 1000/1000/1000/1000/1000
- MaxBonusMissileRange: 400/400/400/400/400
- MaxChargeTime: 3/3/3/3/3
- MaxStunTT: 1.25/1.25/1.25/1.25/1.25
- MinMissileRange: 500/500/500/500/500
- SelfSlowPercent: 0.1/0.1/0.1/0.1/0.1
- StunDuration: 0.25/0.25/0.25/0.25/0.25
- StunMultPercent: 4/4/4/4/4

**E · Domínio Atormentado** · recarga 14/12/10/8/6

- TotalMovespeed: 0.25/0.275/0.3/0.325/0.35 + 0.0004 ability_power
- AttackSpeed: 0.3/0.35/0.4/0.45/0.5
- CamoRadius: 450/450/450/450/450
- MistDuration: 8/8/8/8/8
- RestealthTime: 0.6/0.6/0.6/0.6/0.6
- RestealthTimeAfterHittingLargeMonster: 1.5/1.5/1.5/1.5/1.5

**R · Destruidor de Corações** · recarga 120/100/80

- TotalDamage: 0/0/0 + 1.2 attack_damage
- TotalPercentHealth: 12/16/20 + 0.05 attack_damage (bonus)
- KnockbackBonusRange: 100/100/100
- KnockbackGravity: 15/15/15
- KnockbackSpeed: 1000/1000/1000
- LockoutTime: 1.5/1.5/1.5
- SlowDuration: 0.25/0.25/0.25
- SlowPercent: 0.99/0.99/0.99

## Senna

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 530 | 89 | 2043 |
| dano de ataque | 50 | 0 | 50 |
| armadura | 25 | 4 | 93 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.625 | 2.6 | 44.824997 |
| regeneração de vida | 0.7 | 0.11 | 2.57 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 600 | — | 600 |
| multiplicador de crítico | 2 | — | 2 |

**P · Absolvição**

- BonusCurentHealthDamage: 1
- BonusOnHitDamage: 0 + 0.2 attack_damage
- CriticalDamage: 0 + 1 critical_damage
- DebuffDuration: 4 _(no nível 18)_
- MSSteal: 0.2 _(no nível 18)_
- ArmorShredDuration: 6
- BonusCritChance: 10
- BonusRange: 20
- ChanceForGoodSoul: 0.2
- CritDamageMod: 0.8
- CritToLifestealConversionPercent: 0.35
- MSStealDuration: 0.5
- MarkDuration: 4
- QuestAssistWindow: 3
- QuestRewardForSennaStacks: 40
- QuestRewardForThreshStacks: 40
- QuestRewardSennaAPPerStack: 1
- QuestRewardSennaArmorPerStack: 1
- SennaSoulsToTriggerQuest: 100
- SoulDropPercentChance: 0.32
- SoulDropPercentChanceOnSennaKill: 0.05
- SoulDuration: 8
- SoulTravelDistanceBad: 350
- SoulTravelDistanceGood: -50
- SoulTravelVarianceRadiusInner: 130
- SoulTravelVarianceRadiusOuter: 150
- StacksForBonus: 20
- ThreshSoulsToTriggerQuest: 50

**Q · Escuridão Perfurante** · recarga 15/15/15/15/15 · custo 70/80/90/100/110 · alcance 600/600/600/600/600

- TotalDamage: 30/55/80/105/130 + 0.6 attack_damage (bonus)
- TotalHeal: 40/60/80/100/120 + 0.35 ability_power + 0.4 attack_damage (bonus)
- TotalSlow: 0.15/0.15/0.15/0.15/0.15 + 0.0007 ability_power + 0.0015 attack_damage (bonus)
- CDReductionOnHit: 1/1/1/1/1
- MaxTargetingRange: 1100/1100/1100/1100/1100
- RectangleLeadDistance: 80/80/80/80/80
- RectangleRange: 1300/1300/1300/1300/1300
- RectangleWidthAlly: 280/280/280/280/280
- RectangleWidthEnemy: 100/100/100/100/100
- SlowDuration: 1/1.25/1.5/1.75/2

**W · Abraço Final** · recarga 11/11/11/11/11 · custo 50/55/60/65/70 · alcance 1300/1300/1300/1300/1300

- Damage: 70/110/150/190/230 + 0.9 attack_damage (bonus)
- DelayTime: 1/1/1/1/1
- RootDuration: 1.25/1.5/1.75/2/2.25

**E · Maldição da Névoa Negra** · recarga 26/24.5/23/21.5/20 · custo 70/70/70/70/70

- TotalMS: 0.2/0.2/0.2/0.2/0.2 + 0.0005 ability_power
- AllyMSMult: 1/1/1/1/1
- BuffDuration: 6/6.5/7/7.5/8
- DefenseRadius: 400/400/400/400/400
- DurationBeforeRestealth: 1.5/1.5/1.5/1.5/1.5
- MinAllyduration: 2/2/2/2/2
- SmokeArmTime: 0.25/0.25/0.25/0.25/0.25
- SmokeLinger: 3.5/3.5/3.5/3.5/3.5
- SmokeRadius: 150/150/150/150/150

**R · Sombra da Alvorada** · recarga 140/120/100 · custo 100/100/100

- TotalDamage: 250/400/550 + 0.7 ability_power + 1.15 attack_damage (bonus)
- PostCastAnimLock: 0.5/0.5/0.5
- ShieldDuration: 3/3/3

## Lucian

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 641 | 100 | 2341 |
| dano de ataque | 60 | 2.5 | 102.5 |
| armadura | 28 | 4.2 | 99.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.638 | 2.5 | 43.138 |
| regeneração de vida | 0.75 | 0.13 | 2.96 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Disparo Iluminado**

- MinionDamage: 0 + 1 attack_damage
- PassiveTotalDamage: 15 + 0.2 attack_damage
- TotalDamage: 0 + 0.6 attack_damage _(no nível 18)_
- CCRange: 1000
- MaxAutos: 4
- NumAuto: 2
- OnhitPassiveDuration: 6
- PassiveDuration: 3.5

**Q · Luz Perfurante** · recarga 9/8/7/6/5 · custo 48/56/64/72/80 · alcance 500/500/500/500/500

- TotalDamage: 80/115/150/185/220 + 1 attack_damage (bonus)
- Range: 1000/1000/1000/1000/1000
- Width: 100/100/100/100/100

**W · Chama Ardente** · recarga 14/13/12/11/10 · custo 60/60/60/60/60

- TotalDamage: 75/110/145/180/215 + 0.9 ability_power
- BuffDuration: 1/1/1/1/1
- DebuffDuration: 6/6/6/6/6
- DebuffMarkerDuration: 1/1/1/1/1
- ImpactRangeCheck: 700/700/700/700/700
- MoveSpeedAmount: 60/65/70/75/80
- VisionBubbleDuration: 1/1/1/1/1
- VisionBubbleSize: 200/200/200/200/200

**E · Perseguição Implacável** · recarga 16/15.5/15/14.5/14 · custo 32/24/16/8/0

- CDRefundBase: 1/1/1/1/1
- CDRefundChampion: 2/2/2/2/2
- DashSpeed: 1350/1350/1350/1350/1350
- MaxDistance: 425/425/425/425/425
- MinDistance: 200/200/200/200/200

**R · O Expurgo** · recarga 110/100/90 · custo 100/100/100

- DamagePerBullet: 15/30/45 + 0.15 ability_power + 0.25 attack_damage
- Duration: 3/3/3
- PercentDamageAmpToMinions: 200/200/200
- RangeCheck: 1050/1050/1050
- RecastLockout: 0.75/0.75/0.75
- VisionBubbleDuration: 1/1/1
- VisionBubbleRadius: 100/100/100

## Zed

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 654 | 99 | 2337 |
| dano de ataque | 63 | 3.4 | 120.8 |
| armadura | 32 | 4.7 | 111.899994 |
| resistência mágica | 29 | 2.05 | 63.85 |
| velocidade de ataque | 0.651 | 3.3 | 56.751 |
| regeneração de vida | 1.4 | 0.13 | 3.61 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Desprezo pelos Fracos**

- FinalDamage: 0 + 0.120000005 max_health _(no nível 18)_
- MaxHPDamage: 0.120000005 _(no nível 18)_
- TotalDamage: 0 + 0.120000005 max_health _(no nível 18)_
- {6a1a62e9}: 0.1 _(no nível 18)_
- {ad3d785f}: 0.120000005 _(no nível 18)_
- CurrentHealthThreshold: 0.5
- MonsterDamageCap: 175
- MonsterDamageMod: 0.75
- PerUnitCD: 10
- ShenZedQuestKillCount: 2

**Q · Shuriken Laminado** · recarga 6/6/6/6/6 · custo 75/70/65/60/55

- PassThroughDamage (= TotalDamage × 0.6): 48/72/96/120.00001/144 + 0.6 attack_damage (bonus)
- TotalDamage: 80/120/160/200/240 + 1 attack_damage (bonus)

**W · Sombra Viva** · recarga 20/19/18/17/16 · custo 40/35/30/25/20 · alcance 700/700/700/700/700

- EnergyRestoreDoubleHit: 30/35/40/45/50
- MimickRange: 25000/25000/25000/25000/25000
- RecastRange: 2000/2000/2000/2000/2000
- ShadowDuration: 5.25/5.25/5.25/5.25/5.25
- ShadowDurationTooltip: 5/5/5/5/5
- SwapTimeLimit: 5.25/5.25/5.25/5.25/5.25

**E · Corte Sombrio** · recarga 5/4.5/4/3.5/3 · custo 40/40/40/40/40 · alcance 290/290/290/290/290

- TotalDamage: 70/92.5/115/137.5/160 + 0.7 attack_damage (bonus)
- MoveSpeedMod: -0.2/-0.25/-0.3/-0.35/-0.4
- MoveSpeedModBonus: -0.3/-0.375/-0.45/-0.525/-0.6
- ShadowHitCDR: 3/3/3/3/3
- ShadowRadius: 290/290/290/290/290
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- ZedRadius: 315/315/315/315/315

**R · Marca Fatal** · recarga 120/110/100 · alcance 625/625/625

- RCalculatedDamage: 0/0/0 + 1 attack_damage
- {43a019eb}: 8/8/8
- RCanSwapBuffDuration: 7.5/7.5/7.5
- RDamageAmp: 0.25/0.4/0.55
- RDeathMarkDuration: 3/3/3
- RGhostedDuration: 4/4/4
- RReactivateCD: 0.5/0.5/0.5
- RVisionDuration: 0.5/0.5/0.5

## Kled

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 410 | 84 | 1838 |
| dano de ataque | 65 | 3.5 | 124.5 |
| armadura | 35 | 5.2 | 123.399994 |
| resistência mágica | 28 | 2.05 | 62.85 |
| velocidade de ataque | 0.625 | 3.5 | 60.125 |
| regeneração de vida | 1.2 | 0.15 | 3.7500002 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Skaarl, a Lagarto Covarde** · recarga 0

- DismountedAttackPenalty: 0.0000000037252903 _(no nível 18)_
- DismountedMS: 70
- DismountedResistBonus: 4 + 0.01 max_health (bonus)
- DismountedResistBonusMax (= DismountedResistBonus × 2.5): 10 + 0.024999999 max_health (bonus)
- MountCooldown: 30
- SkaarlHealth: 1400 _(no nível 18)_
- SkaarlRemountHealth: 0.7 _(no nível 18)_
- CourageLastHit: 5
- CourageVsChamps: 15
- CourageVsOther: 5
- DismountDistanceBase: 200
- DismountInvulnWindow: 0.5
- DismountSpeedRadius: 1200
- DismountedMSPenalty: 40
- MountReturnTime: 0.25
- RemountDashDuration: 0.25
- ResistBonusPerEnemy: 0.3
- ResistBonusRange: 1400

**Q · Armadilha na Corda** · recarga 11/10/9/8/7 · alcance 800/800/800/800/800

- DismountedQDamage: 35/50/65/80/95 + 0.65 attack_damage (bonus)
- TotalDamage: 30/55/80/105/130 + 0.6 attack_damage (bonus)
- TotalYankDamage (= TotalDamage × 1): 60/110/160/210/260 + 1.2 attack_damage (bonus)
- CouragePerPellet: 5/5/5/5/5
- GrievousAmount: 0.4/0.4/0.4/0.4/0.4
- GrievousDuration: 5/5/5/5/5
- MinionDamageMultiplier: 1.5/1.5/1.5/1.5/1.5
- MissileSpawnCheckRadius: 60/60/60/60/60
- SlowAmount: -0.3/-0.35/-0.4/-0.45/-0.5
- SlowDuration: 2.5/2.5/2.5/2.5/2.5
- TetherPopTime: 1.75/1.75/1.75/1.75/1.75

**W · Tendências Violentas** · recarga 13/12/11/10/9

- PercentDamage: 4.5/5/5.5/6/6.5 + 0.02 attack_damage (bonus) + 0.004 max_health (bonus)
- ActiveDuration: 4/4/4/4/4
- AttackSpeed: 1.5/1.5/1.5/1.5/1.5
- BaseFlatDamage: 20/30/40/50/60
- ChampCooldownRefund: 1.5/1.5/1.5/1.5/1.5
- MonsterCap: 200/200/200/200/200
- NonChampCooldownRefund: 0.5/0.5/0.5/0.5/0.5
- WCooldown: 13/12/11/10/9

**E · Justar** · recarga 13/12/11/10/9 · alcance 550/550/550/550/550

- TotalDamage: 35/60/85/110/135 + 0.55 attack_damage (bonus)
- DashSpeed: 600/600/600/600/600
- MoveSpeed: 0.5/0.5/0.5/0.5/0.5
- MoveSpeedDuration: 1/1/1/1/1
- PassthroughDistance: 350/350/350/350/350
- RecastWindow: 3/3/3/3/3
- TetherRange: 700/700/700/700/700

**R · Avançaaaaaaar!!!** · recarga 140/125/110 · alcance 3500/4000/4500

- MaximumChargeDamage: 4/6/8 + 0.03 attack_damage (bonus)
- MaximumShield: 200/300/400 + 3 attack_damage (bonus)
- MinimumDamageTooltip: 4/6/8 + 0.03 attack_damage (bonus)
- {31b91b32}: 0/0/0 + 3 attack_damage (bonus)
- {598e3ed3}: 4/6/8 + 0.03 attack_damage (bonus)
- AreaTriggerRadius: 500/500/500
- BackwardsCheckRange: 150/150/150
- DashAreaRange: 700/700/700
- SecondsToMaxPower: 3/3/3
- TOOLTIPRange: 3500/4000/4500

## Ekko

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 655 | 99 | 2338 |
| dano de ataque | 58 | 3 | 109 |
| armadura | 32 | 4.2 | 103.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.688 | 3.3 | 56.788 |
| regeneração de vida | 1.8 | 0.18 | 4.86 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ressonância Revo-Z**

- BonusMS: 0.8 _(no nível 18)_
- SpeedDuration: 3 _(no nível 18)_
- ThreeHitDamage: 30 + 0.8 ability_power
- LockoutTime: 4
- MonsterMod: 2.7

**Q · Giratempo** · recarga 9/8.5/8/7.5/7 · custo 50/60/70/80/90

- InitialDamage: 80/95/110/125/140 + 0.3 ability_power
- RecallDamage: 40/65/90/115/140 + 0.6 ability_power
- ExpandDuration: 1.75/1.75/1.75/1.75/1.75
- SlowPercent: -0.4/-0.45/-0.5/-0.55/-0.6
- SlowZoneRadius: 165/165/165/165/165

**W · Convergência Paralela** · recarga 22/20/18/16/14 · custo 30/35/40/45/50 · alcance 1600/1600/1600/1600/1600

- MissingHealthPercent: 3/3/3/3/3 + 0.03 ability_power
- TotalShield: 100/120/140/160/180 + 1.5 ability_power
- AoERadius: 375/375/375/375/375
- BelowHealthThreshold: 0.3/0.3/0.3/0.3/0.3
- DelayBeforeDetonation: 3/3/3/3/3
- DelayBeforeEnemyCanSee: 2/2/2/2/2
- OnHitMaxMinionDamage: 150/150/150/150/150
- OnHitMinMinionDamage: 15/15/15/15/15
- SlowPercent: 40/40/40/40/40
- SlowZoneDuration: 1.5/1.5/1.5/1.5/1.5
- StunDuration: 2.25/2.25/2.25/2.25/2.25

**E · Mergulho Fásico** · recarga 9/8.5/8/7.5/7 · custo 40/45/50/55/60

- TotalDamage: 50/75/100/125/150 + 0.4 ability_power
- AttackRangeIncrease: 300/300/300/300/300
- BuffDuration: 3/3/3/3/3
- DashDistance: 350/350/350/350/350

**R · Cronoquebra** · recarga 110/80/50 · custo 100/100/100

- TotalBaseHeal: 100/150/200 + 0.6 ability_power
- TotalDamage: 200/350/500 + 1.75 ability_power
- AoERadius: 375/375/375
- PercentHealAmpPerPercentMissingHealth: 3/3/3

## Qiyana

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 115 | 2545 |
| dano de ataque | 64 | 3.1 | 116.7 |
| armadura | 31 | 4.5 | 107.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.688 | 2.1 | 36.387997 |
| regeneração de vida | 1.6 | 0.18 | 4.6600003 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Privilégio da Realeza**

- FinalDamage: 15 + 0.3 ability_power + 0.25 attack_damage (bonus)
- ICD: 14
- {222614ec}: 1
- BurnDuration: 4
- DoubleHitBuffDuration: 6
- HasteDuration: 1.5
- HealDuration: 4
- HealRadius: 360
- SlowDuration: 1.5
- StoneExecuteBase: 0.06
- StunDuration: 0.5

**Q · Cólera Elemental / Lâmina de Ixtal** · recarga 7/7/7/7/7 · custo 35/35/35/35/35 · alcance 525/525/525/525/525

- EnchantedDamage: 70/100/130/160/190 + 0.85 attack_damage (bonus)
- EnchantedFalloff (= EnchantedDamage × 1): 52.5/75/97.5/120/142.5 + 0.63750005 attack_damage (bonus)
- FalloffDamage: 0.75/0.75/0.75/0.75/0.75
- TremorDamage (= EnchantedDamage × 1): 42/60.000004/78/96/114.00001 + 0.51000005 attack_damage (bonus)
- VanillaDamage: 70/100/130/160/190 + 0.85 attack_damage (bonus)
- VanillaFalloff (= VanillaDamage × 1): 52.5/75/97.5/120/142.5 + 0.63750005 attack_damage (bonus)
- {7dc40c72}: 0.06/0.06/0.06/0.06/0.06 + 0.0006 attack_damage (bonus)
- CritThreshold: 0.5/0.5/0.5/0.5/0.5
- Haste: 0.2/0.2/0.2/0.2/0.2
- JungleDamageAmp: 1.75/1.75/1.75/1.75/1.75
- RockExplosionDelay: 0.4/0.4/0.4/0.4/0.4
- RootDuration: 0.5/0.5/0.5/0.5/0.5
- SlowDuration: 1/1/1/1/1
- SlowPotency: -0.2/-0.2/-0.2/-0.2/-0.2
- StealthDuration: 3/3/3/3/3
- TremorBonus: 0.4/0.4/0.4/0.4/0.4

**W · Terraforme** · recarga 7/7/7/7/7 · custo 25/30/35/40/45 · alcance 1100/1100/1100/1100/1100

- CleaveDamage: 8/16/24/32/40 + 1 attack_damage (bonus)
- OnHitDamage: 8/16/24/32/40 + 0.45 ability_power + 0.2 attack_damage (bonus)
- TremorDamage (= CleaveDamage × 1): 2.8/5.6/8.4/11.2/14 + 0.35 attack_damage (bonus)
- AttackSpeed: 0.15/0.2/0.25/0.3/0.35
- DashSpeed: 440/440/440/440/440
- PassiveMS: 0.03/0.05/0.07/0.09/0.11
- Range: 425/425/425/425/425
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- StealthDuration: 3/3/3/3/3
- TravelDistance: 300/300/300/300/300

**E · Audácia** · recarga 11/10/9/8/7 · custo 40/45/50/55/60 · alcance 650/650/650/650/650

- Damage: 50/90/130/170/210 + 0.5 attack_damage (bonus)
- DashSpeed: 600/600/600/600/600
- DashSpeedMax: 1200/1200/1200/1200/1200

**R · Suprema Demonstração de Talento** · recarga 120/120/120 · custo 100/100/100 · alcance 950/950/950

- Damage: 100/200/300 + 1.25 attack_damage (bonus)
- MissingHealthDamageRock: 0.1/0.1/0.1
- MonsterCap: 500/750/1000 + 1.25 attack_damage (bonus)
- BonusDamageBase: 50/100/150
- RecastWindow: 10/10/10
- RootDuration: 1.5/1.5/1.5
- Slow: 0.4/0.5/0.6
- SlowDuration: 2/2/2
- StunDuration: 1/1/1
- StunDurationMin: 0.5/0.5/0.5

## Vi

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 655 | 105 | 2440 |
| dano de ataque | 63 | 3.5 | 122.5 |
| armadura | 30 | 4.7 | 109.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.644 | 2 | 34.644 |
| regeneração de vida | 2 | 0.2 | 5.4 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Blindagem**

- ShieldCooldown: 16
- TotalShield: 0 + 0.12 max_health
- CDReductionOn3Hit: 4
- ShieldDuration: 3

**Q · Quebra-Cofres** · recarga 12/10.5/9/7.5/6 · custo 50/60/70/80/90 · alcance 250/250/250/250/250

- MaxDamageTooltip (= TotalDamage × 1): 100/150/200/250/300 + 1.5 attack_damage (bonus)
- TotalDamage: 40/60/80/100/120 + 0.6 attack_damage (bonus)
- CanceledRefundTime: 3/3/3/3/3
- ChargeDuration: 6/6/6/6/6
- DashWallCheating: 50/50/50/50/50
- ExtraDashRangeAtMaxCharge: 475/475/475/475/475
- KnockbackDuration: 0.75/0.75/0.75/0.75/0.75
- ManaRefundPercent: 0.5/0.5/0.5/0.5/0.5
- MinDashRange: 250/250/250/250/250
- MinionDamageMult: 1/1/1/1/1
- SelfSlow: 15/15/15/15/15
- VFXChargeDuration: 5/5/5/5/5

**W · Pancada Certeira** · recarga 0/0/0/0/0

- TotalDamageTooltip: 4/5/6/7/8 + 0.035 attack_damage (bonus)
- AttackSpeed: 30/35/40/45/50
- MarkerBuffDuration: 4/4/4/4/4
- MonsterDamageCap: 300/300/300/300/300
- SharedBuffsDuration: 4/4/4/4/4
- ShredAmount: 20/20/20/20/20
- StacksBeforeEffect: 2/2/2/2/2

**E · Força Implacável** · recarga 1/1/1/1/1 · custo 26/32/38/44/50

- TotalDamageTooltip: 10/30/50/70/90 + 1 ability_power + 1.1 attack_damage
- {36ebd3cf}: -1/-1/-1/-1/-1 + 1 critical_damage
- AttackBuffDuration: 6/6/6/6/6
- BonusRange: 50/50/50/50/50
- StaticCooldown: 1/1/1/1/1

**R · Saque e Enterrada** · recarga 140/115/90 · custo 100/100/100 · alcance 800/800/800

- Damage: 150/250/350 + 0.9 attack_damage (bonus)
- SecondaryTargetDamage (= Damage × 1): 150/250/350 + 0.9 attack_damage (bonus)
- RBaseSpeed: 800/800/800
- RStunDuration: 1.3/1.3/1.3
- SecondaryTargetKnockbackDistance: 350/350/350
- SecondaryTargetKnockbackDuration: 0.25/0.25/0.25
- SecondaryTargetStunDuration: 0.75/0.75/0.75

## Aatrox

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 650 | 114 | 2588 |
| dano de ataque | 60 | 5 | 145 |
| armadura | 38 | 4.8 | 119.600006 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.651 | 2.5 | 43.151 |
| regeneração de vida | 0.6 | 0.1 | 2.3 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Postura do Arauto da Morte**

- MonsterDamageCap: 320 _(no nível 18)_
- PCooldown: 10 _(no nível 18)_
- PDamage: 0.1 _(no nível 18)_
- PBonusAARange: 50
- PChargeRate: 2
- PDamageTADBonusLevel1: 0.15
- PDamageTADBonusLevel18: 0.25
- PDebuffDuration: 3
- PHealShieldReduction: 0.4
- PHealingMinionMod: 0.25
- PHealingRatio: 1
- PStructureCap: 100

**Q · A Espada Darkin** · recarga 14/12/10/8/6

- QDamage: 10/25/40/55/70 + 0.6 attack_damage
- QEdgeDamage (= QDamage × 1): 17.5/43.75/70/96.25/122.5 + 1.0500001 attack_damage
- QMinionDamage: 0.7/0.7/0.7/0.7/0.7 _(no nível 18)_
- QExtensionTime: 4/4/4/4/4
- QKnockupDuration: 0.25/0.25/0.25/0.25/0.25
- QMonsterBonus: 25/25/25/25/25
- QRampBonus: 0.25/0.25/0.25/0.25/0.25

**W · Correntes Infernais** · recarga 20/18/16/14/12 · alcance 825/825/825/825/825

- WDamage: 30/40/50/60/70 + 0.4 attack_damage
- WBonusADRatio: 0.4/0.4/0.4/0.4/0.4
- WSlowDuration: 1.5/1.5/1.5/1.5/1.5
- WSlowPercentage: -0.25/-0.275/-0.3/-0.325/-0.35

**E · Avanço Umbral** · recarga 9/8/7/6/5

- TotalEVamp: 16/16/16/16/16 + 0.011 max_health (bonus)
- EBaseRechargeTime: 26/22/18/14/10
- EBonusAD: 15/25/35/45/55
- EDashSpeed: 800/800/800/800/800
- EDuration: 1.5/1.5/1.5/1.5/1.5
- EMaxRange: 300/300/300/300/300
- EMinRange: 75/75/75/75/75
- EResistancePerAD: 10/15/20/25/30
- ESpellVampEmpowered: 16/16/16/16/16

**R · Aniquilador de Mundos** · recarga 120/100/80

- RBonusADRange: 75/75/75
- RChampionFearDuration: 1/1/1
- RDuration: 10/10/10
- RERechargeBonus: 0.5/0.5/0.5
- RExtension: 5/5/5
- RFXExtraDuration: 0.6/0.6/0.6
- RFearRadius: 600/600/600
- RHealingAmp: 0.5/0.75/1
- RMinionFearDuration: 3/3/3
- RMovementSpeed: 0.15/0.2/0.25
- RMovementSpeedBonus: 0.6/0.8/1
- RPercentBloodWellToBoostToMax: 0.8/0.8/0.8
- RSpeedMinDuration: 3/3/3
- RTotalADAmp: 0.2/0.3/0.4

## Nami

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 560 | 88 | 2056 |
| dano de ataque | 54 | 3.1 | 106.7 |
| armadura | 29 | 5.2 | 117.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.644 | 2.61 | 45.014 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Maré Oscilante**

- TotalMSBonus: 100 + 0.25 ability_power
- BuffDuration: 1.5
- UltMult: 2

**Q · Prisão Aquática** · recarga 12/11/10/9/8 · custo 60/60/60/60/60 · alcance 850/850/850/850/850

- TotalDamageTT: 90/145/200/255/310 + 0.5 ability_power
- StunDuration: 1.5/1.5/1.5/1.5/1.5

**W · Vazante e Fluxo** · custo 70/75/80/85/90 · alcance 725/725/725/725/725

- BounceScaling: -20/-20/-20/-20/-20 + 0.15 ability_power
- TotalDamage: 60/95/130/165/200 + 0.5 ability_power
- TotalHeal: 55/80/105/130/155 + 0.4 ability_power
- MaxTargets: 3/3/3/3/3

**E · Bênção da Conjuradora** · recarga 11/11/11/11/11 · custo 55/60/65/70/75 · alcance 800/800/800/800/800

- AoeMod: 0.66/0.66/0.66/0.66/0.66 _(no nível 18)_
- TotalDamage: 20/35/50/65/80 + 0.2 ability_power
- TotalSlow: 15/20/25/30/35 + 0.05 ability_power
- BuffDuration: 6/6/6/6/6
- HitCount: 3/3/3/3/3
- SlowDuration: 1/1/1/1/1

**R · Maré Violenta** · recarga 120/110/100 · custo 100/100/100 · alcance 2200/2200/2200

- TotalDamage: 150/250/350 + 0.6 ability_power
- DistToSlowRatio: 0.002/0.002/0.002
- KnockupDuration: 0.5/0.5/0.5
- MaxSlowDuration: 4/4/4
- MinSlowDuration: 2/2/2
- PassiveUpdateAreaLength: 200/200/200
- SlowAmount: 70/70/70

## Azir

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 575 | 108 | 2411 |
| dano de ataque | 56 | 3.5 | 115.5 |
| armadura | 25 | 5 | 110 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 5 | 85.625 |
| regeneração de vida | 1.4 | 0.15 | 3.95 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Legado de Shurima** · recarga 90 · alcance 700

- BonusResists: 30
- TowerDamage: 230 + 0.4 ability_power
- SoldierDeactivationRange: 660
- SoldierStabRange: 350
- TowerDisintegrationTime: 45

**Q · Areias da Conquista** · recarga 14/12/10/8/6 · custo 70/80/90/100/110 · alcance 720/720/720/720/720

- TotalDamage: 75/95/115/135/155 + 0.35 ability_power
- SlowAmount: -0.25/-0.25/-0.25/-0.25/-0.25
- SoldierDashSpeed: 1600/1600/1600/1600/1600
- SoldierFormationClumpDistance: 200/200/200/200/200
- SpearLength: 325/325/325/325/325

**W · Surja!** · recarga 1.5/1.5/1.5/1.5/1.5 · custo 40/35/30/25/20 · alcance 2500/2500/2500/2500/2500

- SecondaryTargetDamageMod: 0.2/0.2/0.2/0.2/0.2
- InputBufferProtection: 0.5/0.5/0.5/0.5/0.5
- OnHitMultiplier: 0.5/0.5/0.5/0.5/0.5
- SoldierDuration: 10/10/10/10/10
- SubsequentDamageMod: 25/25/25/25/25

**E · Areias Oscilantes** · recarga 22/20.5/19/17.5/16 · custo 60/60/60/60/60 · alcance 3000/3000/3000/3000/3000

- TotalDamage: 70/110/150/190/230 + 0.6 ability_power
- TotalShield: 70/110/150/190/230 + 0.6 ability_power
- CastRange: 1100/1100/1100/1100/1100
- DashSpeed: 1700/1700/1700/1700/1700
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5

**R · Decreto do Imperador** · recarga 120/105/90 · custo 100/100/100

- TotalDamage: 200/400/600 + 0.75 ability_power
- NumberOfSoldiers: 6/7/8
- WallDuration: 5/5/5

## Yuumi

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 500 | 69 | 1673 |
| dano de ataque | 49 | 3.1 | 101.7 |
| armadura | 25 | 4.2 | 96.399994 |
| resistência mágica | 25 | 1.1 | 43.7 |
| velocidade de ataque | 0.625 | 1 | 17.625 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 425 | — | 425 |
| multiplicador de crítico | 2 | — | 2 |

**P · Amizade Felina**

- HealAmount: 110 + 0.3 ability_power _(no nível 18)_
- PassiveCooldown: 20
- {656dde8a}: 50
- HealDelayTime: 4
- MinionKillCompanionship: 2
- PityTimer: 6
- TakedownCompanionship: 2

**Q · Projétil Errante** · recarga 6.5/6.5/6.5/6.5/6.5 · custo 50/55/60/65/70

- OnHitDamageCalc: 10/12/14/16/18 + 0.05 ability_power
- TotalMissileDamage: 60/95/130/165/200 + 0.2 ability_power
- TotalMissileDamageEmpowered: 80/135/190/245/300 + 0.3 ability_power
- AfterburnerSpeed: 950/950/950/950/950
- AllyCritChanceMaxAmp: 0.75/0.75/0.75/0.75/0.75
- BuffDuration: 5/5/5/5/5
- EmpoweredSlowAmount: 50/53/56/59/62
- EmpoweredSlowDuration: 2/2/2/2/2
- InitialTimeDelay: 0.2/0.2/0.2/0.2/0.2
- MissileLifetime: 1.85/1.85/1.85/1.85/1.85
- MissileStartingOffset: 75/75/75/75/75
- SlowAmount: 20/20/20/20/20
- SlowDuration: 1/1/1/1/1
- StartingSpeed: 650/650/650/650/650
- TimeToAfterburners: 1.35/1.35/1.35/1.35/1.35

**W · Você e Eu!** · recarga 0/0/0/0/0

- AttachCooldown: 0/0/0/0/0 _(no nível 18)_
- CCAttachLockout: 5/5/5/5/5
- HealthOnHit: 3/4/5/6/7 + 0.03 ability_power
- CastForgivenessRange: 200/200/200/200/200
- CastRange: 700/700/700/700/700
- DashSpeed: 1300/1350/1400/1450/1500
- HealAndShieldPower: 0.04/0.05/0.06/0.07/0.08
- MaxBonusHealShieldPower: 0.15/0.15/0.15/0.15/0.15
- MinBonusHealShieldPower: 0.05/0.05/0.05/0.05/0.05

**E · Frenética** · custo 80/90/100/110/120

- TotalAttackSpeed: 25/27.5/30/32.5/35 + 0.08 ability_power
- TotalShielding: 65/90/115/140/165 + 0.4 ability_power
- MSAmount: 20/20/20/20/20
- MSDuration: 3/3/3/3/3
- ManaRestore: 20/24/28/32/36
- MaxManaPercIncrease: 1/1/1/1/1

**R · Capítulo Final** · recarga 120/110/100 · custo 100/100/100 · alcance 1100/1100/1100

- AllyHealingPerc: 1.3/1.3/1.3
- EnhancedHealPerWave (= TotalHealPerWave × 1): 39/65/91 + 0.15599999 ability_power
- MultiMissileTotal (= TotalMissileDamage × 1): 18.75/31.25/43.75 + 0.0625 ability_power
- TotalHealPerWave: 30/50/70 + 0.12 ability_power
- TotalMissileDamage: 75/125/175 + 0.25 ability_power
- TotalSingleTargetDamage (= TotalMissileDamage × 1): 150/250/350 + 0.5 ability_power
- BaseSlow: -0.1/-0.1/-0.1
- BonusSlowPerWave: -0.1/-0.1/-0.1
- CCDuration: 1.25/1.25/1.25
- DegreesPerTick: 10/10/10
- MissileSlowPercent: -0.4/-0.4/-0.4
- OvershieldBonusDuration: 1.5/1.5/1.5
- SpaceBehindWave: 200/200/200
- TimeBetweenWaves: 0.75/0.75/0.75
- UltDuration: 3.5/3.5/3.5
- UltLength: 1300/1300/1300
- WavesToRoot: 3/3/3

## Samira

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 108 | 2466 |
| dano de ataque | 57 | 3 | 108 |
| armadura | 26 | 4.7 | 105.899994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.658 | 3.3 | 56.758 |
| regeneração de vida | 0.65 | 0.11 | 2.52 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Impulso Audacioso**

- BonusMeleeDamage: 2 + 0.105 attack_damage _(no nível 18)_
- EmpoweredMeleeDamageTooltip (= BonusMeleeDamage × 2): 4 + 0.21 attack_damage _(no nível 18)_
- MSBonusNew: 0.035 _(no nível 18)_
- QueenMaxDashRange: 950 _(no nível 18)_
- DurationTOOLTIP: 6
- SwordRange: 200
- VisualBonusRange: 400

**Q · Talento Natural** · recarga 6/5/4/3/2 · custo 30/30/30/30/30

- DamageCalc: 0/5/10/15/20 + 1.1 attack_damage
- CheatingThreshold: 20/20/20/20/20
- LifestealMod: 1/1/1/1/1

**W · Voragem Afiada** · recarga 30/28/26/24/22 · custo 60/60/60/60/60 · alcance 325/325/325/325/325

- DamageCalc: 20/35/50/65/80 + 0.5 attack_damage (bonus)
- SlashDuration: 0.75/0.75/0.75/0.75/0.75

**E · Ímpeto Indomável** · recarga 20/18/16/14/12 · custo 40/40/40/40/40 · alcance 600/600/600/600/600

- DashDamage: 50/60/70/80/90 + 0.2 attack_damage (bonus)
- AttackSpeedDuration: 5/5/5/5/5
- BonusAttackSpeed: 0.2/0.25/0.3/0.35/0.4
- DashSpeed: 1600/1600/1600/1600/1600
- DashTotalDistance: 650/650/650/650/650

**R · Gatilho Infernal** · recarga 5/5/5

- DamageCalc: 20/40/60 + 0.3 attack_damage
- BaseMSBonus: 0.02/0.02/0.02
- LifestealMod: 1/1/1
- MinionDamageMod: 0.25/0.25/0.25

## Thresh

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 120 | 2660 |
| dano de ataque | 56 | 2.2 | 93.4 |
| armadura | 33 | 0 | 33 |
| resistência mágica | 30 | 1.55 | 56.35 |
| velocidade de ataque | 0.625 | 3.5 | 60.125 |
| regeneração de vida | 1.4 | 0.11 | 3.27 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 450 | — | 450 |
| multiplicador de crítico | 2 | — | 2 |

**P · Condenação**

- SoulsToGainOnPickUp: 1

**Q · Sentença** · recarga 19/16.5/14/11.5/9 · custo 70/70/70/70/70

- TotalDamage: 100/150/200/250/300 + 0.9 ability_power
- HitBonusCooldown: 2/2/2/2/2
- SlowPercent: 75/75/75/75/75
- TauntLength: 1.5/1.5/1.5/1.5/1.5

**W · Passagem Sombria** · recarga 21/20/19/18/17 · custo 50/55/60/65/70 · alcance 950/950/950/950/950

- LanternDuration: 6/6/6/6/6
- ShieldDuration: 4/4/4/4/4

**E · Esfolar** · recarga 13/12.25/11.5/10.75/10 · custo 60/65/70/75/80

- TotalDamage: 75/120/165/210/255 + 0.7 ability_power
- ActiveSlowPercentage: 20/25/30/35/40
- DmgPerSoul: 1.7/1.7/1.7/1.7/1.7
- FullChargeDuration: 10/10/10/10/10
- Icon1ThresholdPercentage: 0.3/0.3/0.3/0.3/0.3
- Icon2ThresholdPercentage: 0.6/0.6/0.6/0.6/0.6
- Icon3ThresholdPercentage: 0.975/0.975/0.975/0.975/0.975
- PassiveADRatio: 90/120/150/180/210
- PassiveDmgPerSoul: 1.5/1.5/1.5/1.5/1.5
- SlowDuration: 1/1/1/1/1

**R · A Caixa** · recarga 120/100/80 · custo 100/100/100 · alcance 850/850/850

- TotalDamage: 250/400/550 + 1 ability_power
- SlowAmount: 99/99/99
- SlowDuration: 2/2/2
- ThreshRCageRadius: 400/400/400
- ThreshRWallLength: 470.228/470.228/470.228
- ThreshRWallThickness: 30/30/30
- WallDuration: 4/4/4

## Illaoi

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 656 | 115 | 2611 |
| dano de ataque | 65 | 5 | 150 |
| armadura | 35 | 5 | 120 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2.5 | 43.125 |
| regeneração de vida | 1.9 | 0.16 | 4.62 |
| velocidade de movimento | 350 | — | 350 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Profetisa de um Deus Ancião**

- SpawnCD: 7 _(no nível 18)_
- MissingHPPercentHeal: 0.05
- PassiveTentacleLength: 925
- PassiveTentacleWidth: 200
- SpawnDensity: 1000
- SpawnRadius: 1200
- TentacleDisabledLifetime: 30

**Q · Golpe de Tentáculo** · recarga 10/9/8/7/6 · custo 40/45/50/55/60

- DamageIncreaseTooltip: 162/162/162/162/162 + 0.4 ability_power + 1.1 attack_damage _(no nível 18)_
- TentacleDamageTotal: 180/180/180/180/180 + 0.4 ability_power + 1.1 attack_damage _(no nível 18)_
- TentacleLength: 800/800/800/800/800
- TentacleWidth: 200/200/200/200/200

**W · Lição Dura** · recarga 4/4/4/4/4 · custo 30/30/30/30/30 · alcance 350/350/350/350/350

- DashSpeed: 600/600/600/600/600 + 1 move_speed
- HealthPercentTotal: 3/3.5/4/4.5/5 + 0.035 attack_damage
- BuffDuration: 6/6/6/6/6
- CooldownDuringR: 2/2/2/2/2
- MonsterDamageCap: 300/300/300/300/300
- WMinDamage: 20/30/40/50/60

**E · Teste de Espírito** · recarga 16/15/14/13/12 · custo 35/40/45/50/55

- EchoPercent: 0.25/0.3/0.35/0.4/0.45 + 0.0008 attack_damage
- TimeBetweenVesselTentacleSlams: 3/3/3/3/3 _(no nível 18)_
- {87aff6dd}: 3/3/3/3/3 _(no nível 18)_
- Density: 700/700/700/700/700
- GraceBeforeDormantAfterELand: 1.5/1.5/1.5/1.5/1.5
- SlowAmount: 0.8/0.8/0.8/0.8/0.8
- SlowDuration: 1.5/1.5/1.5/1.5/1.5
- SpiritDuration: 7/7/7/7/7
- SpiritLeashRange: 1500/1500/1500/1500/1500
- VesselDuration: 4/4/4/4/4

**R · Salto de Fé** · recarga 120/95/70 · custo 100/100/100 · alcance 450/450/450

- DamageCalc: 150/250/350 + 0.5 attack_damage (bonus)
- BuffRadius: 500/500/500
- Duration: 8/8/8

## Rek'Sai

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 99 | 2283 |
| dano de ataque | 62 | 3 | 113 |
| armadura | 35 | 4.5 | 111.5 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.667 | 2 | 34.667 |
| regeneração de vida | 0.5 | 0.1 | 2.2 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Fúria dos Xer'Sai** · recarga 0

- FuryFromAbilities: 25
- FuryFromAttacks: 25
- FuryMinionMod: 0.16
- HealDuration: 3
- HealPercentEndValue: 0.2
- HealPercentStartValue: 0.09
- PauseDuration: 5

**Q · Ira da Rainha / Sondar Presas** · recarga 4/3.5/3/2.5/2 · alcance 300/300/300/300/300

- BurrowDamageTooltip: 50/80/110/140/170 + 0.7 ability_power + 0.25 attack_damage (bonus)
- TotalDamageTooltip: 0/0/0/0/0 + 0.3 attack_damage
- AoERange: 300/300/300/300/300
- AttackSpeed: 0.35/0.35/0.35/0.35/0.35
- BuffDuration: 3/3/3/3/3
- BuffDuration: 3/3/3/3/3
- BurrowedAoE: 200/200/200/200/200
- BurrowedCooldown: 12/11.5/11/10.5/10
- BurrowedDebuffDuration: 5/5/5/5/5
- UnburrowedBaseDamage: 5/10/15/20/25

**W · Escavar / Emergir** · recarga 4/4/4/4/4 · alcance 1500/1500/1500/1500/1500

- UnburrowDamage: 30/55/80/105/130 + 0.8 ability_power
- ADRatio: 0.5/0.5/0.5/0.5/0.5
- BurrowedMoveSpeed: 5/10/15/20/25
- ChampKnockbackRange: 250/250/250/250/250
- DistanceCheck: 85/85/85/85/85
- KnockbackDuration: 0.3/0.3/0.3/0.3/0.3
- KnockupDuration: 1/1/1/1/1
- KnockupImmunity: 10/9/8/7/6
- KnockupRange: 220/220/220/220/220
- LockoutTime: 0.15/0.15/0.15/0.15/0.15
- OtherKnockbackRange: 175/175/175/175/175
- SearchRange: 1500/1500/1500/1500/1500
- Slow: -0.3/-0.35/-0.4/-0.45/-0.5
- SlowDuration: 1.25/1.25/1.25/1.25/1.25
- TremorTickTime: 1/1/1/1/1
- VisionRadiusMod: -0.65/-0.65/-0.65/-0.65/-0.65

**E · Mordida Feroz / Túnel** · recarga 6/6/6/6/6 · alcance 225/225/225/225/225

- BaseDamageCalculation: 70/95/120/145/170 + 0.6 attack_damage (bonus)
- DashSpeed: 500/500/500/500/500 + 1 move_speed
- EmpoweredDamageCalculation (= BaseDamageCalculation × 1): 84/114.00001/144/174/204.00002 + 0.72 attack_damage (bonus)
- {009db1c8}: 10/10/10/10/10
- {438d87f9} (= DashSpeed × 1): 325/325/325/325/325 + 0.65 move_speed
- DashCooldown: 18/17/16/15/14
- DashDistance: 850/850/850/850/850
- MaximumTunnels: 8/8/8/8/8
- MinimumRange: 850/850/850/850/850
- TunnelReuseCooldown: 6/5/4/3/2

**R · Investida do Vazio** · recarga 120/100/80 · alcance 1500/1500/1500

- RBaseDamageCalc: 150/250/350 + 1 attack_damage (bonus)
- CastBuffDuration: 1.1/1.1/1.1
- DamageCheckRadius: 1000/1000/1000
- ExtraDistance: 125/125/125
- LeapMovementSpeed: 1400/1400/1400
- PercentHealthDamage: 15/20/25
- PreyMarkDuration: 5/5/5
- TravelDistancePercent: 0.4/0.4/0.4
- VODistance: 4000/4000/4000

## Ivern

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 99 | 2313 |
| dano de ataque | 50 | 3 | 101 |
| armadura | 27 | 4.7 | 106.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.644 | 3.4 | 58.444 |
| regeneração de vida | 1.4 | 0.17 | 4.29 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 475 | — | 475 |
| multiplicador de crítico | 2 | — | 2 |

**P · Amigo da Floresta**

- HarvestDuration: 40
- HealthCostPercent: 15
- HealthTooltip (= HealthCostPercent × 1): 0 + 0.14999999 max_health (base)
- ManaTooltip (= {c6433908} × 1): 0 + 0.19999999 mana
- {c6433908}: 20
- KrugExperience: 58
- KrugGold: 84
- LevelForBuffShare: 5
- PickupDuration: 60

**Q · Encantador de Raízes** · recarga 14/13/12/11/10 · custo 60/60/60/60/60 · alcance 1150/1150/1150/1150/1150

- TotalDamage: 80/125/170/215/260 + 0.7 ability_power
- DashSpeed: 1500/1500/1500/1500/1500
- RootDuration: 1.6/1.7/1.8/1.9/2

**W · Formação de Arbustos** · recarga 0.5/0.5/0.5/0.5/0.5 · custo 30/30/30/30/30 · alcance 1000/1000/1000/1000/1000

- TotalAllyDamage: 10/15/20/25/30 + 0.1 ability_power
- TotalDamage: 20/27.5/35/42.5/50 + 0.2 ability_power
- AllyBuffDuration: 1.5/1.5/1.5/1.5/1.5
- BrushRadius: 190/190/190/190/190
- BrushSpacing: 90/90/90/90/90
- BuffDuration: 3/3/3/3/3
- MaxBrushDuration: 45/45/45/45/45
- RevealDuration: 8/8/8/8/8

**E · Semente Engatilhada** · recarga 11/10/9/8/7 · custo 70/70/70/70/70 · alcance 750/750/750/750/750

- TotalDamage: 70/90/110/130/150 + 0.8 ability_power
- TotalShield: 75/115/155/195/235 + 0.5 ability_power
- ShieldDuration: 2/2/2/2/2
- SlowAmount: 0.4/0.45/0.5/0.55/0.6
- SlowDuration: 2/2/2/2/2

**R · Margarida!** · recarga 140/130/120 · custo 100/100/100

- TotalBonusResists: 30/30/30
- TotalDaisyAD: 70/100/130 + 0.15 ability_power
- TotalDaisyHP: 1000/1000/1000 + 0.5 ability_power
- TotalShockwaveDamage: 90/140/190 + 0.5 ability_power
- DROnSpawn: 25/25/25
- DaisyAS: 30/45/60
- DaisyAoeDR: 25/25/25
- DaisyDuration: 45/45/45
- MSOnSpawn: 40/40/40
- ShockwaveCCDuration: 1/1/1
- ShockwaveCD: 3/3/3
- SpawnBuffDuration: 5/5/5

## Kalista

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 560 | 114 | 2498 |
| dano de ataque | 57 | 4.75 | 137.75 |
| armadura | 24 | 5.2 | 112.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.694 | 4.5 | 77.194 |
| regeneração de vida | 0.8 | 0.15 | 3.3500001 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Aprumo Marcial**

- MoveSpeedMultiplier: 280 + 2 move_speed
- {091295b4}: 0 + 1 attack_speed
- {c3169619}: 1 + 1 attack_speed (bonus)
- GameModeInteger: 1

**Q · Perfurar** · recarga 9/9/9/9/9 · custo 60/65/70/75/80

- TotalDamage: 10/75/140/205/270 + 1.05 attack_damage

**W · Vigia** · recarga 30/30/30/30/30

- AmmoRechargeTooltip: 90/80/70/60/50
- MarkDuration: 4/4/4/4/4
- MaxHealthDamage: 0.1/0.12/0.14/0.16/0.18
- MaximumMonsterDamage: 100/125/150/175/200
- PerTargetCooldown: 10/10/10/10/10

**E · Lacerar** · recarga 0/0/0/0/0 · custo 30/30/30/30/30 · alcance 1000/1000/1000/1000/1000

- AdditionalDamage: 7/14/21/28/35 + 0.5 ability_power + 0.2 attack_damage
- NormalDamage: 5/15/25/35/45 + 0.65 ability_power + 0.7 attack_damage
- TotalSlowAmount: 0.1/0.18/0.26/0.34/0.42 + 0.0005 ability_power
- EpicMonsterDamageMod: 0.5/0.5/0.5/0.5/0.5
- FakedCooldown: 10/9.5/9/8.5/8
- ManaRefund: 10/15/20/25/30
- MaxSpears: 254/254/254/254/254
- SlowDuration: 2/2/2/2/2

**R · Chamado do Destino** · recarga 160/140/120 · custo 100/100/100 · alcance 1000/1000/1000

- KnockupDuration: 1/1.5/2

## Bardo

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 103 | 2381 |
| dano de ataque | 52 | 3 | 103 |
| armadura | 34 | 5 | 119 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.658 | 2 | 34.658 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 500 | — | 500 |
| multiplicador de crítico | 2 | — | 2 |

**P · Chamado do Viajante**

- ChimesForSlowUpgrade: 5
- ChimesForSplashAreaUpgrade: 5
- ChimesForSplashDamageUpgrade: 5
- MeepDamageNoChime: 30 + 0.4 ability_power
- BaseMeepSpawnCD: 8
- DamagePerCheckpoint: 6
- MaxSpeedStacks: 10
- SlowDuration: 1
- SpeedStackDuration: 20
- TooltipChimeDamageCheckpoint: 5
- TooltipMSMax: 150
- TooltipMSPerStack: 24
- TooltipManaRestore: 12

**Q · Prisão Cósmica** · recarga 11/10/9/8/7 · custo 60/60/60/60/60

- TotalDamage: 80/120/160/200/240 + 0.8 ability_power
- SlowAmountPercentage: 60/60/60/60/60
- SlowDuration: 1/1.2/1.4/1.6/1.8
- StunDuration: 1/1.2/1.4/1.6/1.8

**W · Santuário do Protetor** · recarga 0/0/0/0/0 · custo 70/70/70/70/70

- Calc_MoveSpeed: 0.2/0.225/0.25/0.275/0.3 + 0.0006 ability_power
- InitialHeal: 25/50/75/100/125 + 0.4 ability_power
- MaxHeal: 50/87.5/125/162.5/200 + 0.7 ability_power
- Ammo_Cooldown: 18/18/18/18/18
- Ammo_Limit: 2/2/2/2/2
- ChargeupTime: 5/5/5/5/5
- MaxPacks: 3/3/3/3/3
- MoveSpeed_Duration: 1.5/1.5/1.5/1.5/1.5

**E · Jornada Mágica** · recarga 22/20.5/19/17.5/16 · custo 30/30/30/30/30 · alcance 900/900/900/900/900

- BaseTravelSpeed: 900/900/900/900/900
- DoorDuration: 10/10/10/10/10
- FriendlyMovementBonusPercentage: 33/33/33/33/33

**R · Têmpera do Destino** · recarga 110/95/80 · custo 100/100/100 · alcance 3400/3400/3400

- RStasisDuration: 2.5/2.5/2.5

## Rakan

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 99 | 2293 |
| dano de ataque | 62 | 3.5 | 121.5 |
| armadura | 30 | 4.9 | 113.3 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.635 | 3 | 51.635 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 300 | — | 300 |
| multiplicador de crítico | 2 | — | 2 |

**P · Plumas Mágicas** · recarga 0

- ShieldCooldown: 40
- TotalShield: 225 + 0.95 ability_power _(no nível 18)_
- HitCooldown: 1

**Q · Pena Reluzente** · recarga 11/10/9/8/7 · custo 45/45/45/45/45

- TotalDamage: 70/115/160/205/250 + 0.7 ability_power
- TotalHeal: 210/210/210/210/210 + 0.55 ability_power _(no nível 18)_
- {b88e1204}: 11/10/9/8/7
- HealDelay: 3/3/3/3/3

**W · Entrada Triunfal** · recarga 14/13/12/11/10 · custo 50/60/70/80/90 · alcance 650/650/650/650/650

- TotalDamage: 70/120/170/220/270 + 0.8 ability_power
- BaseDashSpeed: 1700/1700/1700/1700/1700
- KnockupDuration: 1/1/1/1/1

**E · Dança da Batalha** · recarga 0/0/0/0/0 · custo 40/45/50/55/60

- TotalShield: 50/75/100/125/150 + 0.7 ability_power
- BonusDashSpeed: 1150/1150/1150/1150/1150
- Cooldown: 20/18/16/14/12
- Duration: 3/3/3/3/3
- RecastWindow: 5/5/5/5/5
- XayahBonusDashSpeed: 1250/1250/1250/1250/1250

**R · Rapidez** · recarga 130/110/90 · custo 100/100/100 · alcance 150/150/150

- TotalDamageTooltip: 100/200/300 + 0.5 ability_power
- BuffExtensionDuration: 0.25/0.25/0.25
- CharmDuration: 1/1.25/1.5
- Duration: 4/4/4
- InitialCastSpeed: 75/75/75
- TouchRadius: 150/150/150
- TouchSpeed: 150/150/150

## Xayah

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 107 | 2449 |
| dano de ataque | 60 | 3.5 | 119.5 |
| armadura | 25 | 4.2 | 96.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.658 | 3.9 | 66.958 |
| regeneração de vida | 0.65 | 0.15 | 3.2 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · Cortes Certeiros** · recarga 0 · alcance 1000

- PDamageFalloffMax: 0.55
- PDamageFalloffMid: 0.45
- PDamageFalloffMin: 0.35
- PEmpoweredDuration: 8
- PFeatherDuration: 6
- PFeatherSpacingMax: 40
- PFeatherSpacingMin: 25
- PStackMax: 5
- PStacksPerCast: 3

**Q · Punhais Duplos** · recarga 10/9.5/9/8.5/8 · custo 35/35/35/35/35 · alcance 1100/1100/1100/1100/1100

- MultiHitDamage (= TotalDamage × 1): 22.5/30/37.5/45/52.5 + 0.25 attack_damage (bonus)
- TotalDamage: 45/60/75/90/105 + 0.5 attack_damage (bonus)
- ASpeedCastTimeScalarPerHundrethSecond: 0.07/0.07/0.07/0.07/0.07
- MinCastTime: 0.1/0.1/0.1/0.1/0.1
- Missile1DelaySpeed: 3500/3500/3500/3500/3500
- Missile1DelayTime: 0.334/0.334/0.334/0.334/0.334
- Missile2DelaySpeed: 3500/3500/3500/3500/3500
- Missile2DelayTime: 0.584/0.584/0.584/0.584/0.584
- StartingCastTime: 0.25/0.25/0.25/0.25/0.25

**W · Plumagem Mortífera** · recarga 18/17/16/15/14 · custo 60/55/50/45/40 · alcance 1000/1000/1000/1000/1000

- BonusDamagePercent: 25/25/25/25/25
- WAttackSpeedAmount: 35/40/45/50/55
- WAttackSpeedDuration: 4/4/4/4/4
- WMoveSpeedAmount: 30/30/30/30/30
- WMoveSpeedDuration: 1.5/1.5/1.5/1.5/1.5
- WRakanSeekDistance: 1000/1000/1000/1000/1000

**E · Invocadora das Lâminas** · recarga 12/11/10/9/8 · custo 20/20/20/20/20 · alcance 2000/2000/2000/2000/2000

- FeatherDamage: 50/65/80/95/110 + 0.4 attack_damage (bonus)
- MinionDamage (= FeatherDamage × 1): 25/32.5/40/47.5/55 + 0.2 attack_damage (bonus)
- FeatherFalloff: 0.05/0.05/0.05/0.05/0.05
- FeatherReturnDelay: 0.25/0.25/0.25/0.25/0.25
- FeatherThreshold: 3/3/3/3/3
- RootDuration: 1.25/1.25/1.25/1.25/1.25

**R · Tempestade de Plumas** · recarga 140/120/100 · custo 100/100/100

- Damage: 200/300/400 + 1 attack_damage (bonus)
- RAttackDelay: 1/1/1
- RUntargetable: 1.25/1.25/1.25

## Ornn

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 660 | 109 | 2513 |
| dano de ataque | 69 | 3.5 | 128.5 |
| armadura | 33 | 5.2 | 121.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2 | 34.625 |
| regeneração de vida | 1.8 | 0.18 | 4.86 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Forja Viva** · recarga 0 · alcance 600

- {7325018d}: 10 _(no nível 18)_
- AdditionalMythicStatAmp: 0.04
- BaseStatAmp: 0.1
- GameModeInteger: 1
- MasterworkLevel: 13
- UpgradeADEff: 35
- UpgradeAHEff: 31.25
- UpgradeAPEff: 20
- UpgradeAREff: 20
- UpgradeASEff: 30
- UpgradeHPEff: 2.66
- UpgradeMREff: 18
- UpgradeTotalGoldValue: 1000

**Q · Ruptura Vulcânica** · recarga 9/8.5/8/7.5/7 · custo 45/45/45/45/45 · alcance 750/750/750/750/750

- TotalDamage: 20/45/70/95/120 + 1.1 attack_damage
- PillarDuration: 4/4/4/4/4
- SlowAmount: 40/40/40/40/40
- SlowDuration: 2/2/2/2/2

**W · Fôlego do Fole** · recarga 12/11.5/11/10.5/10 · custo 45/50/55/60/65

- BrittlePercentMaxHPCalc: 0.17/0.17/0.17/0.17/0.17 _(no nível 18)_
- TotalMinimumDamage: 80/130/180/230/280
- TotalMonsterDamageCap: 52/64/76/88/100
- BaseDamageTooltip: 100/150/200/250/300
- BreathDuration: 0.75/0.75/0.75/0.75/0.75
- BrittleDuration: 3/3/3/3/3
- FinalTickRectangleLength: 560/560/560/560/560
- MaxPercentHPPerTickTooltip: 12/13/14/15/16
- MovementSpeedReduction: 0.35/0.35/0.35/0.35/0.35
- PercentHPPerTick: 0.024/0.026/0.028/0.03/0.032
- RectangleLength: 500/500/500/500/500
- RectangleWidth: 175/175/175/175/175

**E · Investida Calcinante** · recarga 14/13.5/13/12.5/12 · custo 35/40/45/50/55

- TotalDamage: 80/125/170/215/260 + 0.4 armor (bonus) + 0.4 magic_resist (bonus)
- DashRadius: 175/175/175/175/175
- DashRange: 650/650/650/650/650
- DashSpeed: 1600/1600/1600/1600/1600
- DashWallCheating: 150/150/150/150/150
- KnockupDuration: 1.25/1.25/1.25/1.25/1.25
- ShockwaveRadius: 360/360/360/360/360

**R · Chamado do Deus da Forja** · recarga 140/120/100 · custo 100/100/100

- MinStun: 0.5/0.5/0.5
- RDamageCalc: 125/175/225 + 0.2 ability_power
- BrittleDurationTOOLTIPONLY: 3/3/3
- RCastRange: 2500/2500/2500
- RMissileAccelerationBase: 100/100/100
- RMissileMaxSpeed: 1200/1200/1200
- RSlowDuration: 2/2/2
- RSlowPercentBasePreMath: 40/50/60

## Sylas

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 122 | 2674 |
| dano de ataque | 61 | 3 | 112 |
| armadura | 29 | 5.2 | 117.399994 |
| resistência mágica | 32 | 2.55 | 75.35 |
| velocidade de ataque | 0.645 | 3.5 | 60.145 |
| regeneração de vida | 1.8 | 0.18 | 4.86 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Explosão de Petricita** · recarga 0

- PassiveAoEDamage: 0 + 0.2 ability_power + 0.4 attack_damage
- PassiveDamage: 0 + 0.3 ability_power + 1.3 attack_damage
- AoE: 300
- CheatingThreshold: 25
- MonsterDamageMulti: 1.15
- PassiveAttackSpeed: 1.25
- PassiveCharges: 3
- PassiveDuration: 4

**Q · Correntes-Chicote** · recarga 10/9/8/7/6 · custo 55/55/55/55/55 · alcance 850/850/850/850/850

- Damage: 40/60/80/100/120 + 0.4 ability_power
- ExplosionDamage: 60/115/170/225/280 + 0.8 ability_power
- SlowAmountCalc: 0.15/0.2/0.25/0.3/0.35
- DetonationDelay: 0.6/0.6/0.6/0.6/0.6
- MinionMod: 0.4/0.4/0.4/0.4/0.4
- MonsterMod: 1/1/1/1/1
- Q2AoE: 180/180/180/180/180
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**W · Regicida** · recarga 12/10.5/9/7.5/6 · custo 50/60/70/80/90

- MaxHealing (= MinHealing × 1): 40/80/120/160/200 + 0.6 ability_power + 0.1 max_health (bonus)
- MinDamage: 75/110/145/180/215 + 0.6 ability_power
- MinHealing: 20/40/60/80/100 + 0.3 ability_power + 0.05 max_health (bonus)
- MaxExecuteThreshold: 0.4/0.4/0.4/0.4/0.4

**E · Evasão / Abdução** · recarga 13/12/11/10/9 · custo 65/65/65/65/65

- Damage: 80/130/180/230/280 + 0.8 ability_power
- KnockUpDuration: 0.5/0.5/0.5/0.5/0.5

**R · Usurpar** · recarga 80/55/30 · custo 75/75/75 · alcance 950/950/950

- PerTargetCooldown: 200/200/200
- EnemyCooldownPercent: 2/2/2
- EnemyCooldownTooltip: 200/200/200
- MinimumEnemyCooldown: 40/40/40
- UltHoldDuration: 90/90/90

## Neeko

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 104 | 2378 |
| dano de ataque | 48 | 2.5 | 90.5 |
| armadura | 21 | 5.2 | 109.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 3.5 | 60.125 |
| regeneração de vida | 1.5 | 0.15 | 4.05 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Encanto Inerente**

- ChargeTime: 2
- MagicPen: 0.2
- MagicPenDuration: 4
- PassiveCooldown: 6

**Q · Explosão Florescente** · recarga 9/8.5/8/7.5/7 · custo 50/60/70/80/90 · alcance 800/800/800/800/800

- ExplosionDamage: 60/110/160/210/260 + 0.6 ability_power
- SecondDamage: 35/60/85/110/135 + 0.25 ability_power
- MonsterBonus: 35/50/65/80/95
- RepeatDelay: 0.75/0.75/0.75/0.75/0.75
- SlowAmount: 40/40/40/40/40
- SlowDuration: 1/1/1/1/1

**W · Metamorfa** · recarga 16/15/14/13/12 · alcance 900/900/900/900/900

- PassiveBonusDamageCalc: 30/65/100/135/170 + 0.6 ability_power
- CloneDamageTimer: 0.75/0.75/0.75/0.75/0.75
- CloneDuration: 3/3/3/3/3
- Haste: 20/25/30/35/40
- HasteDuration: 3/3/3/3/3
- MonsterBonus: 75/75/75/75/75
- PassiveHaste: 10/17.5/25/32.5/40
- PassiveHasteDuration: 1/1/1/1/1
- StealthDuration: 0.5/0.5/0.5/0.5/0.5

**E · Farpas Emaranhadas** · recarga 12/11.5/11/10.5/10 · custo 60/65/70/75/80 · alcance 1000/1000/1000/1000/1000

- BaseDamage: 70/105/140/175/210 + 0.65 ability_power
- BaseRootDuration: 0.5/0.5/0.5/0.5/0.5
- EmpoweredSpeed: 200/200/200/200/200
- EmpoweredWidth: 100/100/100/100/100
- MaxRootDuration: 1.8/2.1/2.4/2.7/3
- MinRootDuration: 0.7/0.9/1.1/1.3/1.5
- TargetsRequiredForMaxValue: 2/2/2/2/2

**R · Florescer Repentino** · recarga 120/105/90 · custo 100/100/100 · alcance 600/600/600

- BaseShield: 75/100/125 + 0.75 ability_power
- ShieldMultiplier: 40/60/80 + 0.4 ability_power
- TotalDamage: 150/350/550 + 1.2 ability_power
- DelayBeforePassiveRemoval: 0.5/0.5/0.5
- DelayUntilExplosion: 0.6/0.6/0.6
- DelayUntilReveal: 1.25/1.25/1.25
- Duration: 2.5/2.5/2.5
- ShieldDuration: 2/2/2
- SlowAmount: 40/40/40
- StunDuration: 0.75/0.75/0.75

## Aphelios

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 600 | 102 | 2334 |
| dano de ataque | 55 | 2.3 | 94.1 |
| armadura | 26 | 4.2 | 97.399994 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.665 | 2.1 | 36.364998 |
| regeneração de vida | 0.65 | 0.11 | 2.52 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · O Assassino e a Profetisa** · recarga 2 · custo 20

- ArPenBonusMax: 27
- AttackDamageMax: 24
- AttackSpeedMax: 0.54

**R · Vigília do Plenilúnio** · recarga 120/110/100 · custo 100/100/100 · alcance 1300/1300/1300

- GravitumRBonusSlow: 0.99/0.99/0.99
- InfernumRBonusDamage: 50/100/150 + 0.25 attack_damage (bonus)
- InfernumRPenalty: 0.9/0.9/0.9
- MaxDamage: 125/175/225 + 1 ability_power + 0.2 attack_damage (bonus)
- {36ebd3cf}: 0/0/0 + 1 critical_damage
- CalibrumRMarkBonusDamage: 50/80/110
- CrescendumRBonusMinis: 5/5/5
- GravitumRRootDuration: 1.35/1.35/1.35
- SeverumRHealBonus: 250/350/450

## Rell

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 104 | 2388 |
| dano de ataque | 55 | 3 | 106 |
| armadura | 30 | 4 | 98 |
| resistência mágica | 28 | 1.8 | 58.6 |
| velocidade de ataque | 0.625 | 2 | 34.625 |
| regeneração de vida | 1.5 | 0.17 | 4.39 |
| velocidade de movimento | 315 | — | 315 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · A Ferro e Fogo** · alcance 1500

- OnHitDamage: 0 + 0.05 armor + 0.05 magic_resist
- StealFloor: 3 _(no nível 18)_
- MaxPercentTooltipOnly: 15
- MaxStacks: 5
- RefreshDuration: 3.5
- ShredDuration: 5
- StealPercent: 0.03

**Q · Golpe Estilhaçador** · recarga 11/10.5/10/9.5/9 · custo 50/50/50/50/50 · alcance 5000/5000/5000/5000/5000

- Damage: 60/100/140/180/220 + 0.6 ability_power
- StunDuration: 0.65/0.65/0.65/0.65/0.65

**W · Ferromante: Queda Esmagadora** · custo 40/40/40/40/40

- DismountDamage: 60/90/120/150/180 + 0.6 ability_power
- FlipDamage: 10/25/40/55/70 + 0.4 ability_power
- Shield: 20/40/60/80/100 + 0.11 max_health
- CrashDownKnockupDuration: 0.4/0.4/0.4/0.4/0.4
- CrashDownStunDuration: 0.8/0.8/0.8/0.8/0.8
- DismountedASBoost: 0.2/0.2/0.2/0.2/0.2
- DismountedRangeBoost: 75/75/75/75/75
- FlipBaseDamage: 10/25/40/55/70
- FlipKnockupDuration: 0.4/0.4/0.4/0.4/0.4
- FlipStunDuration: 0.6/0.6/0.6/0.6/0.6
- MountUpSpeed: 0.3/0.3/0.3/0.3/0.3
- MountUpSpeedDuration: 3.5/3.5/3.5/3.5/3.5
- MountedMoveSpeed: 20/25/30/35/40
- ResistanceIncrease: 0.15/0.15/0.15/0.15/0.15
- SlideDistance: 320/320/320/320/320

**E · Investida Absoluta** · recarga 14/13/12/11/10 · custo 40/40/40/40/40 · alcance 1200/1200/1200/1200/1200

- MaxHealthDamageCalc: 0.05/0.055/0.06/0.065/0.07 + 0.0003 ability_power
- PercentHealthDamageCap: 300/300/300/300/300 _(no nível 18)_
- Duration: 3/3/3/3/3
- EmpowermentBuffer: 2/2/2/2/2
- ExplosionRadius: 300/300/300/300/300
- MSModPhase1: 1/1/1/1/1
- MaxMS: 0.3/0.3/0.3/0.3/0.3
- MinMS: 0.15/0.15/0.15/0.15/0.15
- PassiveMSCombatPenalty: 0.5/0.5/0.5/0.5/0.5
- Phase3Start: 0.5/0.5/0.5/0.5/0.5

**R · Tempestade Magnética** · recarga 120/100/80 · custo 100/100/100 · alcance 200/200/200

- DamagePerSecond: 75/125/175 + 0.55 ability_power
- TotalDamage (= DamagePerSecond × 1): 150/250/350 + 1.1 ability_power
- GravityForce: 300/300/300
- GravityHorizon: 225/225/225
- GravityRadius: 375/375/375
- InitialRadius: 425/425/425

## Pyke

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 670 | 110 | 2540 |
| dano de ataque | 62 | 2 | 96 |
| armadura | 37 | 4.7 | 116.899994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.667 | 2.5 | 43.167 |
| regeneração de vida | 1.4 | 0.1 | 3.1 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Dádiva dos Afogados** · recarga 0

- AdditionalBonusCalc: 0.4 + 0.004 lethality
- DamageStorageMax: 80 + 8 attack_damage (bonus)
- OneEnemyCalc: 0.09 + 0.002 lethality
- AmmoCooldown: 30
- CooldownPerChamp: 1
- HPPerBAD: 14

**Q · Espeto de Osso** · recarga 10/9.5/9/8.5/8 · custo 70/75/80/85/90

- TotalDamage: 100/150/200/250/300 + 0.75 attack_damage (bonus)
- ChargeTimeToGetMaxRange: 1/1/1/1/1
- ManaRefund: 0.75/0.75/0.75/0.75/0.75
- MaxChargeTime: 3/3/3/3/3
- MaxRangeDistance: 1100/1100/1100/1100/1100
- MinChargeTime: 0.5/0.5/0.5/0.5/0.5
- MinRangeDistance: 500/500/500/500/500
- PullDistance: 500/500/500/500/500
- SelfSlow: 0.2/0.2/0.2/0.2/0.2
- SlowAmount: 0.9/0.9/0.9/0.9/0.9
- SlowDuration: 1/1/1/1/1

**W · Mergulho Fantasma** · recarga 14/13/12/11/10 · custo 65/65/65/65/65 · alcance 600/600/600/600/600

- MoveSpeed: 45/45/45/45/45 + 2 lethality
- {601654a1}: 0.25/0.25/0.25/0.25/0.25
- BaseHealCap: 80/80/80/80/80
- CamoDuration: 5/5/5/5/5
- HealADRatio: 8/8/8/8/8
- HealRampBase: 0.01/0.01/0.01/0.01/0.01
- HealRampPerTick: 1.15/1.15/1.15/1.15/1.15
- MaxHPHealCap: 0.55/0.55/0.55/0.55/0.55

**E · Ressaca Espectral** · recarga 15/14/13/12/11 · custo 40/40/40/40/40 · alcance 550/550/550/550/550

- StunDuration: 1.25/1.25/1.25/1.25/1.25 + 0.01 lethality
- TotalDamage: 100/150/200/250/300 + 1 attack_damage (bonus)
- BonusADRatio: 1/1/1/1/1
- DashDistance: 550/550/550/550/550
- ELethalityRatio: 0.01/0.01/0.01/0.01/0.01
- StunDelay: 1/1/1/1/1

**R · Morte das Profundezas** · recarga 100/85/70 · custo 100/100/100 · alcance 750/750/750

- RADDamage: 0/0/0 + 0.8 attack_damage (bonus)
- RBaseDamage: 250/250/250
- RDamage: 250/250/250 + 0.8 attack_damage (bonus) + 1.5 lethality
- RLethalityDamage: 0/0/0 + 1.5 lethality
- ReducedDamageFinal (= RDamage × 1): 125/125/125 + 0.4 attack_damage (bonus) + 0.75 lethality
- DamageCap: 0.5/0.5/0.5
- RRecastDuration: 20/20/20

## Vex

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 104 | 2358 |
| dano de ataque | 54 | 2.75 | 100.75 |
| armadura | 23 | 4.45 | 98.649994 |
| resistência mágica | 28 | 1.3 | 50.1 |
| velocidade de ataque | 0.669 | 1 | 17.669 |
| regeneração de vida | 1.3 | 0.12 | 3.34 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Destruição e Escuridão** · recarga 0 · alcance 1600

- DoomCD: 16 _(no nível 18)_
- FearDuration: 1.5 _(no nível 18)_
- GloomMinionMod: 0.6 _(no nível 18)_
- GloomProcCalc: 150 + 0.25 ability_power _(no nível 18)_
- {c21c264c}: 25000
- DashListenerRange: 1600
- DashLockoutTime: 0.5
- EFearedMinionSpeed: 5
- FearSlowPercentMax: 0.99
- FearSlowPercentMin: 0.6
- FearedMinionSpeed: 25
- GloomCDChamp: 0.25
- GloomCDNonChamp: 0.1
- GloomDuration: 6
- MaxDoomCD: 15
- RangeOfMaxFearSlow: 800
- RangeOfMinFearSlow: 250

**Q · Rajada Mistral** · recarga 8/7/6/5/4 · custo 45/50/55/60/65 · alcance 1200/1200/1200/1200/1200

- QDamageCalc: 70/115/160/205/250 + 0.7 ability_power
- {7b4a4737}: 70/115/160/205/250 + 0.5 ability_power

**W · Espaço Pessoal** · recarga 16/15/14/13/12 · custo 75/75/75/75/75 · alcance 475/475/475/475/475

- ShieldCalc: 50/75/100/125/150 + 0.75 ability_power
- WDamageCalc: 80/120/160/200/240 + 0.3 ability_power
- AOERadius: 475/475/475/475/475
- BonusAOERadiusVsDashes: 75/75/75/75/75
- ShieldDuration: 2.5/2.5/2.5/2.5/2.5

**E · Penumbra Iminente** · recarga 12/12/12/12/12 · custo 70/80/90/100/110 · alcance 800/800/800/800/800

- EDamageCalc: 50/70/90/110/130 + 0.4 ability_power
- GloomCDChampTooltip: 0.25/0.25/0.25/0.25/0.25
- GloomCDNonChampTooltip: 0.1/0.1/0.1/0.1/0.1
- MaxRadius: 300/300/300/300/300
- RadiusGrowthRate: 13/13/13/13/13
- SlowAmount: 0.3/0.35/0.4/0.45/0.5
- SlowDuration: 2/2/2/2/2
- StartingRadius: 200/200/200/200/200

**R · Onda Sombria** · recarga 140/120/100 · custo 100/100/100 · alcance 2000/2500/3000

- RDamageCalc: 75/125/175 + 0.2 ability_power
- RecastDamageCalc: 150/250/350 + 0.5 ability_power
- R2Duration: 4/4/4
- ResetWindow: 12/12/12
- TakedownWindow: 8/8/8

## Yone

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 105 | 2405 |
| dano de ataque | 62 | 2 | 96 |
| armadura | 33 | 4.6 | 111.2 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 3.5 | 60.125 |
| regeneração de vida | 1.5 | 0.15 | 4.05 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Estilo do Caçador** · recarga 0

- CurrentCritDamage: 0 + 1 critical_damage
- CritChanceMultiplier: 1
- CritDamageMod: 0.9
- MagicDamageSplit: 0.5
- YoneCritToAD: 50

**Q · Aço Mortal** · recarga 4/4/4/4/4 · alcance 450/450/450/450/450

- QDamage: 25/50/75/100/125 + 1.1 attack_damage
- BuffDuration: 6/6/6/6/6
- Q3AoESize: 100/100/100/100/100
- Q3DashSpeed: 1500/1500/1500/1500/1500
- Q3KnockupDuration: 0.75/0.75/0.75/0.75/0.75
- Q3KnockupHeight: 0.5/0.5/0.5/0.5/0.5
- Q3MaxDistance: 450/450/450/450/450
- Q3MinDistance: 450/450/450/450/450
- QAttackSpeedCDMax: 0.66667/0.66667/0.66667/0.66667/0.66667
- QAttackSpeedCDPercent: 0.6/0.6/0.6/0.6/0.6
- QAttackSpeedCastReductionMax: 0.5/0.5/0.5/0.5/0.5
- QAttackSpeedCastReductionPercent: 0.41667/0.41667/0.41667/0.41667/0.41667
- QBaseCastTime: 0.35/0.35/0.35/0.35/0.35
- Width: 80/80/80/80/80

**W · Fenda Espiritual** · recarga 14/14/14/14/14 · alcance 700/700/700/700/700

- MinimumDamageMinions: 30/30/30/30/30
- MonsterDamageCap: 320/320/320/320/320 _(no nível 18)_
- WDamage: 10/20/30/40/50 + 1 attack_damage
- WShield: 90/90/90/90/90 + 0.65 attack_damage (bonus) _(no nível 18)_
- {8dfeb810}: 0/0/0/0/0 + 1 attack_damage
- BaseShield: 10/20/30/40/50
- ConeWidth: 40/40/40/40/40
- FirstChampShieldMultiplier: 1/1/1/1/1
- HealthStealMultiplier: 1/1/1/1/1
- MaxHealthDamage: 0.08/0.09/0.1/0.11/0.12
- MinionMod: 0.25/0.25/0.25/0.25/0.25
- SecondChampShieldMultiplier: 0.5/0.5/0.5/0.5/0.5
- ShieldDuration: 1.5/1.5/1.5/1.5/1.5
- WAttackSpeedCDMax: 0.57142/0.57142/0.57142/0.57142/0.57142
- WAttackSpeedCDPercent: 0.6/0.6/0.6/0.6/0.6
- WAttackSpeedCastReductionMax: 0.5/0.5/0.5/0.5/0.5
- WAttackSpeedCastReductionPercent: 0.41667/0.41667/0.41667/0.41667/0.41667
- WBaseCastTime: 0.5/0.5/0.5/0.5/0.5
- WCritMultiplier: 0.8/0.8/0.8/0.8/0.8

**E · Desatar da Alma** · recarga 22/19/16/13/10

- {df232269}: 20/35/50/65/80 + 0.25 attack_damage
- DeathmarkPercent: 0.25/0.275/0.3/0.325/0.35
- EDashRange: 300/300/300/300/300
- EDashSpeed: 1200/1200/1200/1200/1200
- MissingHealthPercent: 0.25/0.275/0.3/0.325/0.35
- MovementSpeed: 0.3/0.3/0.3/0.3/0.3
- OnHitIncrease: 0.25/0.25/0.25/0.25/0.25
- RecastLockout: 0.5/0.5/0.5/0.5/0.5
- ReturnTimer: 5/5/5/5/5
- StartingMS: 0.1/0.1/0.1/0.1/0.1

**R · Destino Selado** · recarga 120/100/80 · alcance 1000/1000/1000

- Damage: 200/400/600 + 0.8 attack_damage (bonus)
- TooltipDamage (= Damage × 0.5): 100/200/300 + 0.4 attack_damage (bonus)
- BlinkDistanceBehindTarget: 200/200/200
- RKnockTowardsDampener: 0.8/0.8/0.8
- RKnockTowardsSpeed: 3000/3000/3000
- RKnockupDuration: 0.75/0.75/0.75
- RKnockupHeight: 0.5/0.5/0.5
- Width: 225/225/225

## Ambessa

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 110 | 2500 |
| dano de ataque | 63 | 3 | 114 |
| armadura | 35 | 4.9 | 118.3 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2.5 | 43.125 |
| regeneração de vida | 1.7 | 0.15 | 4.25 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Passo do Cão-dragão** · recarga 0 · alcance 350

- Calc_Attack_Speed: 0.5
- Calc_OnHit_Damage_Flat: 30 + 0.25 attack_damage (bonus) _(no nível 18)_
- Calc_OnHit_Energy_Refund: 70 _(no nível 18)_
- {57d26a55}: 950 + 1 move_speed _(no nível 18)_
- Attack_Buff_Duration: 4
- Attack_Buff_Max_Stacks: 3
- Attack_Range_Amount: 75
- Attack_Range_Total: 200
- Buffer_Dash_Distance: 350
- Buffer_Dash_Forgiveness_Window: 0.275
- Buffer_Dash_Min_Distance: 175
- Buffer_Dash_Time: 0.3
- Buffer_Dash_W_Forgiveness_Window: 0.25
- Buffer_Dash_Wall_Forgiveness: 500

**Q · Golpe Ardiloso / Pancada Cortante** · recarga 14/13/12/11/10 · custo 70/70/70/70/70 · alcance 650/650/650/650/650

- Calc_Damage_1_Max: 40/60/80/100/120 + 0.6 attack_damage (bonus)
- Calc_Damage_1_Min_Ratio: 0.5/0.5/0.5/0.5/0.5
- Calc_Damage_1_Percent_Max: 0.04/0.045/0.05/0.055/0.06 + 0.0003 attack_damage (bonus)
- Calc_Damage_2_Max: 50/75/100/125/150 + 0.9 attack_damage (bonus)
- Calc_Damage_2_Min_Ratio: 0.5/0.5/0.5/0.5/0.5
- Calc_Damage_2_Percent_Max: 0.04/0.045/0.05/0.055/0.06 + 0.0004 attack_damage (bonus)
- Calc_Damage_Monster_Flat_Bonus: 75/75/75/75/75
- Calc_Damage_Monster_Percent_Cap: 300/300/300/300/300 _(no nível 18)_
- {1442dbe0} (= Calc_Damage_2_Percent_Max × 1): 0.02/0.0225/0.025/0.0275/0.03 + 0.0002 attack_damage (bonus)
- {17768f39} (= Calc_Damage_1_Percent_Max × 1): 0.02/0.0225/0.025/0.0275/0.03 + 0.00015 attack_damage (bonus)
- {3ce89b9e} (= Calc_Damage_2_Max × 1): 25/37.5/50/62.5/75 + 0.45 attack_damage (bonus)
- {a7907b77} (= Calc_Damage_1_Max × 1): 20/30/40/50/60 + 0.3 attack_damage (bonus)
- Swap_Duration: 4/4/4/4/4
- Swap_Static_Cooldown: 0.5/0.5/0.5/0.5/0.5

**W · Repúdio** · recarga 18/17/16/15/14 · custo 70/70/70/70/70 · alcance 325/325/325/325/325

- Calc_Damage_High (= Calc_Damage_Low × 1): 75/112.5/150/187.5/225 + 0.75 attack_damage (bonus)
- Calc_Damage_Low: 50/75/100/125/150 + 0.5 attack_damage (bonus)
- Calc_Shield: 320/320/320/320/320 + 1.5 attack_damage (bonus) _(no nível 18)_
- Buff_Duration: 0.5/0.5/0.5/0.5/0.5
- Dash_Delay: 0.225/0.225/0.225/0.225/0.225
- Shield_Duration: 1.5/1.5/1.5/1.5/1.5

**E · Lacerar** · recarga 13/12/11/10/9 · custo 70/70/70/70/70 · alcance 325/325/325/325/325

- Calc_Damage_Flat: 40/60/80/100/120 + 0.5 attack_damage (bonus)
- Slow_Amount: 0.99/0.99/0.99/0.99/0.99
- Slow_Duration: 1/1/1/1/1

**R · Execução Pública** · recarga 130/115/100 · alcance 1250/1250/1250

- Calc_Damage: 150/250/350 + 0.8 attack_damage (bonus)
- Calc_Omnivamp: 0.15/0.175/0.2 + 0.5 omnivamp
- Armor_Penetration: 0.1/0.2/0.3
- Omnivamp_MinionMod: 0.25/0.25/0.25
- Omnivamp_MonsterMod: 0.25/0.25/0.25
- Stun_Duration: 0.4/0.4/0.4
- Suppress_Duration: 0.75/0.75/0.75

## Mel

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 630 | 99 | 2313 |
| dano de ataque | 54 | 3.3 | 110.1 |
| armadura | 21 | 5.2 | 109.399994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.5 | 43.125 |
| regeneração de vida | 1.2 | 0.11 | 3.07 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Esplendor Calcinante**

- PassiveBonusMissileDamage: 30 + 0.04 ability_power _(no nível 18)_
- BonusAttackDuration: 5
- MaxPassiveBonusMissiles: 9
- OverwhelmDuration: 5
- PassiveBonusMissiles: 3

**Q · Rajada Radiante** · recarga 10/9/8/7/6 · custo 70/75/80/85/90 · alcance 950/950/950/950/950

- AllDamageHit: 85/127/173/223/277 + 0.8 ability_power
- InitialExplosionDamage: 60/85/110/135/160 + 0.55 ability_power
- MonsterModTooltip: 0/0/0/0/0
- TotalExplosionDamage: 5/7/9/11/13 + 0.05 ability_power
- ChannelDuration: 0.5/0.5/0.5/0.5/0.5
- ExplosionRadius: 200/200/200/200/200
- Spread: 25/25/25/25/25

**W · Refutação** · recarga 38/35/33/29/26 · custo 80/60/40/20/0 · alcance 250/250/250/250/250

- DamagePercent: 0.4/0.45/0.5/0.55/0.6 + 0.0005 ability_power
- ShieldAmount: 80/110/140/170/200 + 0.7 ability_power
- Duration: 0.75/0.75/0.75/0.75/0.75
- MoveSpeed: 0.4/0.4/0.4/0.4/0.4
- MoveSpeedDuration: 0.75/0.75/0.75/0.75/0.75
- PhysDamageMod: 0.3/0.3/0.3/0.3/0.3
- ShieldRadius: 175/175/175/175/175

**E · Armadilha Solar** · recarga 11/10.5/10/9.5/9 · custo 50/60/70/80/90 · alcance 1000/1000/1000/1000/1000

- AreaDamagePerSecond: 16/28/40/52/64 + 0.08 ability_power
- Damage: 60/105/150/195/240 + 0.7 ability_power
- MinionModTooltip: 0.5/0.5/0.5/0.5/0.5
- AreaSlowAmount: 0.3/0.3/0.3/0.3/0.3
- AreaSlowDuration: 0.25/0.25/0.25/0.25/0.25
- AreaTicksPerSecond: 8/8/8/8/8
- DoTDuration: 0.5/0.5/0.5/0.5/0.5
- MaxAreaRadius: 230/230/230/230/230
- RootDuration: 1.5/1.5/1.5/1.5/1.5
- TimeToMaxRadius: 0.5/0.5/0.5/0.5/0.5

**R · Eclipse Dourado** · recarga 120/100/80 · custo 100/100/100

- MinionModTooltip: 0.5/0.5/0.5
- PassiveFlatDamage: 60/70/80 + 0.1 ability_power
- PassiveStackDamage: 3/4/5 + 0.0075 ability_power
- UltFlatDamage: 125/200/275 + 0.3 ability_power
- UltStackDamage: 4/7/10 + 0.04 ability_power
- DisintegrateDelay: 0.9/0.9/0.9

## Yunara

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 590 | 110 | 2460 |
| dano de ataque | 55 | 3 | 106 |
| armadura | 25 | 4.4 | 99.8 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.65 | 2 | 34.65 |
| regeneração de vida | 0.8 | 0.11 | 2.67 |
| velocidade de movimento | 325 | — | 325 |
| alcance de ataque | 575 | — | 575 |
| multiplicador de crítico | 2 | — | 2 |

**P · Voto às Primeiras Terras** · recarga 0

- Calc_Damage_Amp: 0.1 + 0.001 ability_power

**Q · Cultivação do Espírito** · recarga 0/0/0/0/0 · custo 30/30/30/30/30

- Calc_Attack_Speed: 0.2/0.3/0.4/0.5/0.6
- Calc_Damage: 5/10/15/20/25 + 0.2 ability_power
- Calc_Damage_Spread: 0/0/0/0/0 + 0.3 attack_damage
- Calc_Minion_Execute_Amp: 2.5/2.5/2.5/2.5/2.5
- Calc_Minion_Execute_Threshold: 0.3/0.3/0.3/0.3/0.3
- Calc_Passive_Damage: 5/10/15/20/25 + 0.2 ability_power
- Buff_Duration: 5/5/5/5/5
- Resource_Champion: 2/2/2/2/2
- Resource_Duration: 6/6/6/6/6
- Resource_Loss: 1/1/1/1/1
- Resource_Loss_Interval: 0.5/0.5/0.5/0.5/0.5
- Resource_Max: 8/8/8/8/8
- Resource_Nonchampion: 1/1/1/1/1
- Spread_Lifesteal_Efficacy: 1/1/1/1/1
- Spread_Onhit_Efficacy: 0.3/0.3/0.3/0.3/0.3
- Spread_Radius: 300/300/300/300/300

**W · Arco do Julgamento | Arco da Ruína** · custo 60/60/60/60/60 · alcance 1150/1150/1150/1150/1150

- Calc_Damage_Initial: 55/95/135/175/215 + 0.5 ability_power + 0.85 attack_damage (bonus)
- Calc_Damage_Per_Second (= Calc_Damage_Initial × 1): 33/57.000004/81/105.00001/129 + 0.3 ability_power + 0.51000005 attack_damage (bonus)
- Calc_Minion_Damage_Mod: 1/1/1/1/1 _(no nível 18)_
- Calc_Slow: 0.99/0.99/0.99/0.99/0.99
- {669adf33} (= Calc_Damage_Initial × 1): 55/95/135/175/215 + 0.5 ability_power + 0.85 attack_damage (bonus) _(no nível 18)_
- Cast_Time_Attack_Speed_Cap: 100/100/100/100/100
- Cast_Time_Attack_Speed_Cap_Empowered: 100/100/100/100/100
- Cast_Time_Base: 0.45/0.45/0.45/0.45/0.45
- Cast_Time_Base_Empowered: 0.6/0.6/0.6/0.6/0.6
- Cast_Time_Min: 0.225/0.225/0.225/0.225/0.225
- Cast_Time_Min_Empowered: 0.45/0.45/0.45/0.45/0.45
- Linger_Duration: 1/1/1/1/1
- Minion_Execute_Tick_Count: 6/6/6/6/6
- Slow_Duration: 1.5/1.5/1.5/1.5/1.5

**E · Passos de Kanmei | Sombra Intocável** · recarga 9/9/9/9/9 · custo 40/40/40/40/40

- Calc_Move_Speed: 0.3/0.35/0.4/0.45/0.5
- Calc_Move_Speed_Enhanced (= Calc_Move_Speed × 1): 0.45000002/0.525/0.6/0.67499995/0.75
- Buff_Duration: 1.5/1.5/1.5/1.5/1.5

**R · Transcendência do Eu** · recarga 100/90/80 · custo 100/100/100

- Calc_RW_Cooldown_Reduction: 0.8/0.8/0.8
- Calc_RW_Damage: 160/320/480 + 0.75 ability_power + 1.2 attack_damage (bonus)
- Calc_RW_Slow_Amount: 0.99/0.99/0.99
- {42964db5}: 1350/1500/1650
- Buff_Duration: 15/15/15
- RW_Slow_Duration: 1/1/1

## Locke

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 109 | 2473 |
| dano de ataque | 58 | 3 | 109 |
| armadura | 32 | 4.2 | 103.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.688 | 3.3 | 56.788 |
| regeneração de vida | 1.8 | 0.18 | 4.86 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Estaca de Prata** · recarga 0

- MaxOnHitEndValue: 80
- MaxOnHitStartValue: 10
- MinOnHitEndValue: 40
- MinOnHitStartValue: 5

**Q · Pregos Ritualísticos** · recarga 10/9/8/7/6 · custo 70/70/70/70/70 · alcance 950/950/950/950/950

- AmmoCooldownReset: 0.35/0.35/0.35/0.35/0.35 _(no nível 18)_
- MissileDamage: 40/48/56/64/72 + 0.2 ability_power
- NailDamage: 18/26/34/42/50 + 0.25 ability_power
- Ammo: 3/3/3/3/3
- Duration: 4/4/4/4/4
- MaxStacks: 3/3/3/3/3
- NailDuration: 4/4/4/4/4
- QMarkMinionMod: 1/1/1/1/1
- QMissileMinionMod: 1/1/1/1/1
- SlowAmount1: 0.25/0.25/0.25/0.25/0.25
- SlowAmount2: 0.25/0.25/0.25/0.25/0.25
- SlowAmount3: 0.6/0.6/0.6/0.6/0.6
- SlowDuration1: 1/1/1/1/1
- SlowDuration2: 1/1/1/1/1
- SlowDuration3: 2/2/2/2/2
- ThreeMarkBonusPercent: 40/40/40/40/40
- TwoMarkBonusPercent: 20/20/20/20/20

**W · Ignição da Alma** · recarga 18/17/16/15/14 · custo 50/55/60/65/70 · alcance 250/250/250/250/250

- AdditionalHeal: 200/200/200/200/200 + 0.2 ability_power _(no nível 18)_
- AttackSpeed: 0.7/0.7/0.7/0.7/0.7 _(no nível 18)_
- DamageRestoreAmount: 40/60/80/100/120 + 1 ability_power
- DecayTimeHelper: 2/2/2/2/2
- Duration: 6/6/6/6/6
- MoveSpeed: 0.4/0.4/0.4/0.4/0.4 + 0.0002 ability_power
- HealCapLevel18Value: 450/450/450/450/450
- HealCapLevel1Value: 150/150/150/150/150
- HealthCost: 0.02/0.02/0.02/0.02/0.02
- MinMoveSpeed: 0.2/0.2/0.2/0.2/0.2
- Radius: 250/250/250/250/250

**E · Perseguição das Cinzas** · custo 30/40/50/60/70 · alcance 425/425/425/425/425

- DashDamage: 40/60/80/100/120 + 0.4 ability_power
- OnHitDamage: 40/50/60/70/80 + 0.4 ability_power
- BonusRange: 275/275/275/275/275
- E2AttackDelay: 0.1/0.1/0.1/0.1/0.1
- MinCastRange: 150/150/150/150/150
- Radius: 100/100/100/100/100

**R · Purgatório** · recarga 120/100/80 · custo 100/100/100 · alcance 1000/1000/1000

- Damage: 150/225/300 + 0.6 ability_power
- AbilityHastePerStack: 6/6/6
- Anim_BiteTime: 0.39/0.39/0.39
- CooldownReduction: 20/20/20
- Duration: 5/5/5
- Radius: 400/400/400
- SlowAmount: 0.99/0.99/0.99
- SlowDuration: 2/2/2

## Sett

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 670 | 114 | 2608 |
| dano de ataque | 60 | 4 | 128 |
| armadura | 33 | 4.7 | 112.899994 |
| resistência mágica | 28 | 2.05 | 62.85 |
| velocidade de ataque | 0.625 | 1.75 | 30.375 |
| regeneração de vida | 1.4 | 0.1 | 3.1 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ousadia da Arena**

- HealingCalc: 0.021 _(no nível 18)_
- RightPunchBonus: 90 + 0.55 attack_damage (bonus) _(no nível 18)_
- tooltipregenpermissinghealthcalc (= {7733e08c} × 5): 200 _(no nível 18)_
- {7733e08c}: 40 _(no nível 18)_
- {a156fbdc}: 1.5 _(no nível 18)_
- AttackSpeed: 1.5
- GrabDuration: 2
- HookDamagePercentFinal: 1.5
- JabDamagePercent: 1
- LockoutTimer: 10
- MarkerDuration: 3
- MaxHPOnHit: 0.06
- MissingHealthUnit: 0.01
- NumHits: 2
- PassiveHealthThreshold: 0.75
- PassiveHealthThresholdAmp: 0.02
- SelfSlowPercent: 0.5

**Q · Pancadaria** · recarga 9/8/7/6/5

- BonusEnemyMaxHealthDamage: 0.01/0.015000001/0.02/0.025/0.030000001
- MaxHealthDamageCalc: 0.01/0.01/0.01/0.01/0.01 + 0.0001 attack_damage
- ADRatio: 1.2/1.2/1.2/1.2/1.2
- AdrenalineDecayLockTimer: 1/1/1/1/1
- BaseDamage: 10/20/30/40/50
- BaseHealthCost: 20/20/20/20/20
- Duration: 4/4/4/4/4
- HealthCost: 0.1/0.1/0.1/0.1/0.1
- MSAmount: 0.3/0.3/0.3/0.3/0.3
- MSDuration: 1.5/1.5/1.5/1.5/1.5
- MaxHealthTADRatioTOOLTIP: 1/1.5/2/2.5/3
- MonsterCap: 400/400/400/400/400
- QMinimumDamage: 10/15/20/25/30
- StackCount: 1/1/1/1/1

**W · Casca-Grossa** · recarga 18/16.5/15/13.5/12

- DamageCalc: 80/100/120/140/160
- DamageConversion: 0.25/0.25/0.25/0.25/0.25 + 0.0025 attack_damage (bonus)
- MaxGrit: 0/0/0/0/0 + 0.5 max_health
- AdrenalineStorageWindow: 4/4/4/4/4
- DamageAmp: 1/1/1/1/1
- DamageStored: 1/1/1/1/1
- Duration: 0.75/0.75/0.75/0.75/0.75
- EnrageADMult: 0.1/0.1/0.1/0.1/0.1
- EnrageArmorPen: 0.5/0.5/0.5/0.5/0.5
- EnrageDuration: 6/6/6/6/6
- EnrageHealingMult: 0.5/0.5/0.5/0.5/0.5
- GreyHealthDuration: 3/3/3/3/3
- LockTime: 2/2/2/2/2
- MSDuration: 2/2/2/2/2
- MovementSpeed: 0.5/0.5/0.5/0.5/0.5
- ResourceDecayRate: 0.3/0.3/0.3/0.3/0.3
- ShieldConversion: 1/1/1/1/1
- ShieldDecayDelay: 0.75/0.75/0.75/0.75/0.75
- ShieldMaxDuration: 3/3/3/3/3

**E · Quebra-Crânio** · recarga 16/14.5/13/11.5/10

- DamageCalc: 50/70/90/110/130 + 0.6 attack_damage
- MonsterBonus: 250/250/250/250/250 _(no nível 18)_
- tooltipadratio_e: 0/0/0/0/0 + 0.6 attack_damage
- Duration: 0.25/0.25/0.25/0.25/0.25
- LeapDistance: 250/250/250/250/250
- MaxHealthDamage: 0.1/0.1/0.1/0.1/0.1
- MonsterDamage: 150/175/200/225/250
- SingleBaseDamage: 10/20/30/40/50
- SlowAmount: 0.7/0.7/0.7/0.7/0.7
- SlowDuration: 0.5/0.5/0.5/0.5/0.5
- StunDuration: 1/1/1/1/1
- ThrowRange: 500/500/500/500/500

**R · Hora do Show** · recarga 120/100/80

- DamageCalc: 200/300/400 + 1.2 attack_damage (bonus)
- MaxHealthDamageCalc: 0.4/0.5/0.6
- tooltipadratio_r: 0/0/0 + 1 attack_damage (bonus)
- CarryDistancePerStack: 200/200/200
- CraterDuration: 10/10/10
- CraterEdgeDamageReduction: 0.5/0.5/0.5
- CraterMaxHaste: 0.6/0.6/0.6
- CraterMaxSlow: 0.99/0.99/0.99
- CraterRadius: 600/600/600
- CraterSweetSpotRadius: 125/125/125
- DelayBeforeDecay: 3/3/3
- DistancePastFirstHit: 600/600/600
- DistanceToFirstHit: 400/400/400
- GrabDuration: 1.5/1.5/1.5
- MaxPassiveStacks: 4/4/4
- MoveSpeedBonus: 800/800/800
- ResistDuration: 6/6/6
- SlowAmount: 0.99/0.99/0.99
- SlowDuration: 1/1/1

## Lillia

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 605 | 105 | 2390 |
| dano de ataque | 61 | 3.1 | 113.7 |
| armadura | 22 | 4.5 | 98.5 |
| resistência mágica | 32 | 1.55 | 58.35 |
| velocidade de ataque | 0.625 | 2.7 | 46.525 |
| regeneração de vida | 0.5 | 0.11 | 2.37 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 325 | — | 325 |
| multiplicador de crítico | 2 | — | 2 |

**P · Ramo Onírico**

- ChampionHeal: 15 + 0.05 ability_power _(no nível 18)_
- ChampionHealTT (= ChampionHeal × 1): 90 + 0.3 ability_power _(no nível 18)_
- DotPercentTooltip (= DotPercentTotal × 0.01): 0.049999997 + 0.000125 ability_power
- DotPercentTotal: 5 + 0.0125 ability_power
- MonsterDamageCap: 180 _(no nível 18)_
- MonsterHealTT (= {b5d0316f} × 1): 39 + 0.15 ability_power
- {395ffd44} (= {b5d0316f} × 1): 1.625 + 0.00625 ability_power
- {7969ffc2} (= {b5d0316f} × 1): 9.75 + 0.0375 ability_power
- {b5d0316f}: 6.5 + 0.025 ability_power
- ChampionHealingPercent: 1
- Duration: 3
- LargeMonsterHealMod: 1.5
- MonsterHealingPercent: 0.25
- MultiHealMod: 0.15

**Q · Golpes Florescentes** · recarga 6/5.5/5/4.5/4 · custo 65/65/65/65/65 · alcance 500/500/500/500/500

- BonusTrueDamage: 35/45/55/65/75 + 0.35 ability_power
- PranceSpeed: 0.03/0.04/0.05/0.06/0.07 + 0.0003 ability_power
- TotalDamage: 35/45/55/65/75 + 0.35 ability_power
- MaximumRadius: 485/485/485/485/485
- MinimumRadius: 225/225/225/225/225
- MinionDamageMod: 1/1/1/1/1
- PranceDuration: 6.5/6.5/6.5/6.5/6.5
- PranceFalloffTime: 1/1/1/1/1
- PranceMaxStacks: 4/4/4/4/4

**W · Cuidado! Iiip!** · recarga 14/13/12/11/10 · custo 50/50/50/50/50

- FlatDamage: 80/100/120/140/160 + 0.35 ability_power
- FlatDamageSweetSpot: 80/100/120/140/160 + 0.35 ability_power
- AnimRangeCutoff: 300/300/300/300/300
- LandingOffset: 150/150/150/150/150
- MaxSpeed: 500/500/500/500/500
- MaximumCastTime: 0.75/0.75/0.75/0.75/0.75
- MaximumDistance: 700/700/700/700/700
- MinimumCastTime: 0.6/0.6/0.6/0.6/0.6
- MinimumDistance: 500/500/500/500/500
- MinimumRadius: 250/250/250/250/250
- MinionDamageMod: 0.5/0.5/0.5/0.5/0.5
- SweetSpotRadius: 65/65/65/65/65

**E · Semente Espiral** · recarga 12/12/12/12/12 · custo 70/70/70/70/70 · alcance 700/700/700/700/700

- ImpactDamageTotal: 60/85/110/135/160 + 0.5 ability_power
- AnimCutoffRange: 350/350/350/350/350
- BounceRange: 25000/25000/25000/25000/25000
- HitboxLength: 200/200/200/200/200
- HitboxWidth: 350/350/350/350/350
- ImpactRangeCheck: 115/115/115/115/115
- SlowAmount: 0.4/0.4/0.4/0.4/0.4
- SlowDuration: 3/3/3/3/3

**R · Cadência de Ninar** · recarga 150/130/110 · custo 50/50/50 · alcance 1600/1600/1600

- TotalDamage: 100/150/200 + 0.4 ability_power
- DrowsyDuration: 1.5/1.5/1.5
- ImpactDamage: 25/25/25
- InitialSlow: -0.1/-0.1/-0.1
- SleepDuration: 2/2/2

## Gwen

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 620 | 115 | 2575 |
| dano de ataque | 63 | 3 | 114 |
| armadura | 39 | 4.9 | 122.3 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.69 | 2.25 | 38.94 |
| regeneração de vida | 1.8 | 0.18 | 4.86 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Mil Retalhos** · recarga 0

- ExecuteDamage: 30 _(no nível 18)_
- HealCap: 40 + 0.07 ability_power _(no nível 18)_
- HealingPercent: 0.67
- MonsterDamageCap: 3 + 0.05 ability_power
- PassiveMaxQTooltip (= {540bce84} × 0.06): 0.06 + 0.00036 ability_power
- PassiveMaxRTooltip (= {540bce84} × 0.09): 0.09 + 0.00054000004 ability_power
- PercentHealth1000Cuts (= {540bce84} × 0.01): 0.01 + 0.00006 ability_power
- {4097817a} (= {540bce84} × 0.05): 0.05 + 0.0003 ability_power
- {4197830d} (= {540bce84} × 0.03): 0.03 + 0.00018 ability_power
- {540bce84}: 1 + 0.006 ability_power
- ExecuteThreshold: 0.4

**Q · Corte e Recorte** · recarga 6.5/5.75/5/4.25/3.5 · custo 40/40/40/40/40 · alcance 500/500/500/500/500

- FinalSwipeDamage: 60/85/110/135/160 + 0.35 ability_power
- MaxDamage: 110/155/200/245/290 + 0.6 ability_power
- MiniSwipeDamage: 10/14/18/22/26 + 0.05 ability_power
- BuffDuration: 6/6/6/6/6
- ExecuteBonus: 1000/1000/1000/1000/1000
- ExecuteThreshold: 0.2/0.2/0.2/0.2/0.2
- InitialArcLength: 475/475/475/475/475
- MiniDamageRatio: 0.2/0.2/0.2/0.2/0.2
- MinionMod: 0.8/0.8/0.8/0.8/0.8
- TrueDamageConversion: 0.5/0.5/0.5/0.5/0.5

**W · Névoa Sagrada** · recarga 22/21/20/19/18 · custo 60/60/60/60/60

- TotalResists: 22/24/26/28/30 + 0.07 ability_power
- DashTargetForgiveness: 125/125/125/125/125
- NewDashCooldown: 2/2/2/2/2
- PullOffset: 75/75/75/75/75
- RecastDelay: 0.5/0.5/0.5/0.5/0.5
- ZoneDuration: 4/4/4/4/4
- ZonePullSpeed: 2000/2000/2000/2000/2000
- ZoneRadius: 370/370/370/370/370
- ZoneWallOffset: -100/-100/-100/-100/-100

**E · Avanço Afiado** · recarga 13/12.5/12/11.5/11 · custo 35/35/35/35/35

- BonusAttackSpeed: 30/42.5/55/67.5/80
- OnHitDamage: 15/15/15/15/15 + 0.2 ability_power
- BonusAttackRange: 75/75/75/75/75
- BuffDuration: 4/4/4/4/4
- CDRefund: 0.25/0.35/0.45/0.55/0.65
- DashRange: 350/350/350/350/350
- DashSpeed: 800/800/800/800/800
- WallCheatDistance: 100/100/100/100/100

**R · Ponto-Cruz** · recarga 120/100/80 · custo 100/100/100 · alcance 1300/1300/1300

- MaxDamage (= TotalDamage × 9): 270/450/630 + 0.90000004 ability_power
- TotalDamage: 30/50/70 + 0.1 ability_power
- TotalDamage3 (= TotalDamage × 3): 90/150/210 + 0.3 ability_power
- TotalDamage5 (= TotalDamage × 5): 150/250/350 + 0.5 ability_power
- DebuffDuration: 1.5/1.5/1.5
- InitialSlow: -0.4/-0.5/-0.6
- LockoutTime: 1/1/1
- SlowAmount: 0.15/0.15/0.15
- SubsequentSlow: -0.15/-0.2/-0.25

## Renata Glasc

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 545 | 94 | 2143 |
| dano de ataque | 49 | 3 | 100 |
| armadura | 27 | 4.7 | 106.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 2.11 | 36.495 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Financiamento** · recarga 0

- PercentAmpCalc: 0.011299999 + 0.0002 ability_power _(no nível 18)_
- PercentAmpCalcSelf: 0.011299999 + 0.0002 ability_power _(no nível 18)_
- MonsterDamageCap: 150
- PassiveDuration: 6
- PercentAmp: 0.1

**Q · Negócio Fechado** · recarga 16/16/16/16/16 · custo 80/80/80/80/80 · alcance 900/900/900/900/900

- TotalDamage: 80/125/170/215/260 + 0.8 ability_power
- PullDistance: 275/275/275/275/275
- RootDuration: 1/1/1/1/1
- SelfSlow: -0.3/-0.3/-0.3/-0.3/-0.3
- StunDuration: 0.5/0.5/0.5/0.5/0.5

**W · Empréstimo** · recarga 28/27/26/25/24 · custo 80/80/80/80/80 · alcance 800/800/800/800/800

- ASCalc: 10/15/20/25/30 + 0.01 ability_power
- FinalASCalc: 10/15/20/25/30 + 0.01 ability_power
- FinalMSCalc: 10/12.5/15/17.5/20 + 0.01 ability_power
- MSCalc: 10/12.5/15/17.5/20 + 0.01 ability_power
- {85d7d7f0}: 10/10/10/10/10
- Duration: 5/5/5/5/5
- MaxStatMultiplier: 2/2/2/2/2
- TagDuration: 6/6/6/6/6
- TicksPerSecond: 4/4/4/4/4
- TriumphPercent: 20/20/20/20/20

**E · Programa de Fidelidade** · recarga 14/13/12/11/10 · custo 70/80/90/100/110 · alcance 800/800/800/800/800

- ShieldCalc: 50/65/80/95/110 + 0.5 ability_power
- TotalDamage: 65/95/125/155/185 + 0.55 ability_power
- {fe737c05}: 30/30/30/30/30
- EndAoE: 225/225/225/225/225
- InitialAoE: 325/325/325/325/325
- ShieldDuration: 3/3/3/3/3
- SlowDuration: 2/2/2/2/2

**R · Apropriação Agressiva** · recarga 150/130/110 · custo 100/100/100 · alcance 2000/2000/2000

- BerserkDuration: 1.25/1.75/2.25
- BonusAttackSpeed: 1/1/1

## Aurora

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 607 | 110 | 2477 |
| dano de ataque | 53 | 3 | 104 |
| armadura | 23 | 4.5 | 99.5 |
| resistência mágica | 32 | 1.3 | 54.1 |
| velocidade de ataque | 0.668 | 2 | 34.668 |
| regeneração de vida | 1.2 | 0.11 | 3.07 |
| velocidade de movimento | 335 | — | 335 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Abjuração Espiritual**

- BonusMSCalc: 0.02 + 0.00015 ability_power
- CooldownDuration: 5
- HealCalc: 20 + 0.02 ability_power _(no nível 18)_
- MoveSpeedCalc: 0.052 + 0.0003 ability_power _(no nível 18)_
- ProcDamage: 0.01 + 0.00027 ability_power
- {394d1781}: 100
- DebuffDuration: 4
- MSIncrease: 5
- SpiritModeDuration: 4

**Q · Feitiço Dúplice** · recarga 9/8.5/8/7.5/7 · custo 60/60/60/60/60 · alcance 900/900/900/900/900

- Damage: 45/70/95/120/145 + 0.4 ability_power
- Q2DamageMax: 45/70/95/120/145 + 0.4 ability_power
- {f7c018c7}: 45/70/95/120/145 + 0.4 ability_power
- DamageReduction: 0.2/0.2/0.2/0.2/0.2
- MarkDuration: 3.5/3.5/3.5/3.5/3.5
- MinionDamageReduction: 0.1/0.1/0.1/0.1/0.1
- MissingHealthPercentMod: 0.5/0.5/0.5/0.5/0.5
- Q2MinionMod: 0.4/0.4/0.4/0.4/0.4

**W · Através do Véu** · recarga 22/21/20/19/18 · custo 80/80/80/80/80

- DashBonusSpeed: 350/350/350/350/350
- InvisDuration: 1/1.15/1.3/1.45/1.6
- JumpDistance: 300/300/300/300/300
- MoveSpeedBonus: 20/25/30/35/40
- WallCheatDistance: 450/450/450/450/450

**E · Estranheza** · recarga 15/14/13/12/11 · custo 80/80/80/80/80 · alcance 825/825/825/825/825

- DamageCalc: 70/110/150/190/230 + 0.7 ability_power
- AoELength: 825/825/825/825/825
- AoEWidth: 175/175/175/175/175
- SelfKnockBackDistance: 250/250/250/250/250
- SlowDecay: 0.85/0.85/0.85/0.85/0.85
- SlowDuration: 1/1/1/1/1
- SlowPercent: -0.8/-0.8/-0.8/-0.8/-0.8

**R · Entre Mundos** · recarga 140/120/100 · custo 100/100/100 · alcance 700/700/700

- DamageCalc: 175/275/375 + 0.7 ability_power
- AoESize: 700/700/700
- AreaDuration: 2.5/3.25/4
- ExitBaseDamage: 50/100/150
- ExitSlowPercent: -0.5/-0.5/-0.5
- JumpMaxDistance: 250/250/250
- JumpMinDistance: 25/25/25
- MSMultiplier: 2/2/2
- RBuffDuration: 3.5/4.25/5
- SlowPercent: -0.3/-0.3/-0.3
- StunDuration: 1.5/1.75/2

## Nilah

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 570 | 101 | 2287 |
| dano de ataque | 58 | 2 | 92 |
| armadura | 27 | 4.2 | 98.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.697 | 1.25 | 21.947 |
| regeneração de vida | 1.2 | 0.18 | 4.26 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 225 | — | 225 |
| multiplicador de crítico | 2 | — | 2 |

**P · Alegria Eterna** · alcance 1000

- ExperiencePercentage: 0.5
- HealDistance: 1000
- HealingIncrease: 0.075
- ShieldIncrease: 0.15
- ShieldMaxDuration: 4

**Q · Lâmina Sem Forma** · recarga 4/4/4/4/4 · custo 30/30/30/30/30

- AttackTotalDamageTooltip: 0/0/0/0/0 + 1 attack_damage
- BonusAttackSpeedCalc: 60/60/60/60/60 _(no nível 18)_
- CritArmorPen: 0/0/0/0/0 + 0.3 critical_chance
- CritLifesteal: 0/0/0/0/0 + 0.2 critical_chance
- DamageCalc: 0/10/20/30/40 + 1 attack_damage
- BuffDuration: 4/4/4/4/4
- CastRange: 600/600/600/600/600
- CheatingThreshold: 20/20/20/20/20
- MinionMod: 0.33/0.33/0.33/0.33/0.33
- MonsterMod: 1/1/1/1/1
- RangeIncrease: 125/125/125/125/125
- ShieldDuration: 4/4/4/4/4

**W · Véu Jubiloso** · recarga 26/25/24/23/22 · custo 60/45/30/15/0 · alcance 150/150/150/150/150

- AreaTriggerSize: 150/150/150/150/150
- BaseDuration: 2.25/2.25/2.25/2.25/2.25
- ExtensionBonus: 0.75/0.75/0.75/0.75/0.75
- MagicDamageReduction: 0.25/0.25/0.25/0.25/0.25
- MoveSpeedPercent: 0.15/0.175/0.2/0.225/0.25
- ShareBaseDuration: 1.5/1.5/1.5/1.5/1.5

**E · Turbilhão** · recarga 0.5/0.5/0.5/0.5/0.5 · custo 40/40/40/40/40 · alcance 550/550/550/550/550

- DashDamage: 60/70/80/90/100 + 0.2 attack_damage (bonus)
- DashSpeed: 2200/2200/2200/2200/2200
- DashTotalDistance: 600/600/600/600/600

**R · Apoteose** · recarga 110/95/80 · custo 100/100/100

- ChampHealingPercent: 0.2/0.2/0.2 + 0.1 critical_chance
- DamageCalc: 125/225/325 + 1 attack_damage (bonus)
- DamagePerTickCalc: 15/25/35 + 0.1 attack_damage (bonus)
- DamagePerTickCalcTooltip (= DamagePerTickCalc × 4): 60/100/140 + 0.4 attack_damage (bonus)
- ShieldCalc: 65/115/165 + 1 attack_damage (bonus)
- Duration: 6/6/6
- Microslow: -0.1/-0.1/-0.1
- NumTicks: 6/6/6
- OtherHealingPercent: 0.1/0.1/0.1
- TickDelay: 0.2/0.2/0.2

## K'Sante

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 625 | 120 | 2665 |
| dano de ataque | 64 | 3.5 | 123.5 |
| armadura | 36 | 5.2 | 124.399994 |
| resistência mágica | 30 | 2.1 | 65.7 |
| velocidade de ataque | 0.688 | 2.5 | 43.188 |
| regeneração de vida | 1.9 | 0.2 | 5.3 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 150 | — | 150 |
| multiplicador de crítico | 2 | — | 2 |

**P · Instinto Valente**

- MaxHealthDamagePercent: 0.01 + 0.0001 armor (bonus) + 0.0001 magic_resist (bonus)
- {281dd874}: 0.19 _(no nível 18)_
- ArmorMRMultiplier: 0.002
- FlatDamage: 12
- MarkDamagePercentMax: 0.02
- MarkDamagePercentMin: 0.01
- MarkDuration: 4
- PassiveStartingDamage: 0.0035
- RecoveryPercent: 0.8

**Q · Golpes de Ntofo** · recarga 3.5/3.5/3.5/3.5/3.5 · custo 20/20/20/20/20 · alcance 450/450/450/450/450

- BaseDamage: 70/100/130/160/190 + 0.4 armor (bonus) + 0.4 magic_resist (bonus)
- BaseCD: 3.5/3.5/3.5/3.5/3.5
- DefenseCapforCooldown: 120/120/120/120/120
- KnockGrav: 30/30/30/30/30
- KnockupDuration: 0.65/0.65/0.65/0.65/0.65
- LifetimeMax: 0.45/0.45/0.45/0.45/0.45
- MaxCastTime: 0.45/0.45/0.45/0.45/0.45
- MinCD: 2/2/2/2/2
- MinCastTime: 0.35/0.35/0.35/0.35/0.35
- MinCastTimeQ3: 0.35/0.35/0.35/0.35/0.35
- PullDistance: 300/300/300/300/300
- PullDistanceMax: 400/400/400/400/400
- PullDistanceMin: 100/100/100/100/100
- PullSpeed: 500/500/500/500/500
- PullSpeedMax: 600/600/600/600/600
- PullSpeedMin: 110/110/110/110/110
- RCooldownReduction: 0.33333/0.33333/0.33333/0.33333/0.33333
- Range: 700/700/700/700/700
- RecastLockoutAO: 1.25/1.25/1.25/1.25/1.25
- RecastWindow: 6/6/6/6/6
- RectangleLength: 465/465/465/465/465
- RectangleWidth: 100/100/100/100/100
- SlowDuration: 0.5/0.5/0.5/0.5/0.5
- SlowPercent: 0.8/0.8/0.8/0.8/0.8
- StunDuration: 1/1/1/1/1

**W · Criador de Caminhos** · recarga 14/13/12/11/10 · custo 40/45/50/55/60 · alcance 250/250/250/250/250

- TotalMaxHealthDamage: 0.08/0.08/0.08/0.08/0.08 + 0.0002 armor (bonus) + 0.0002 magic_resist (bonus)
- AOMinDurationTOOLTIP: 0.75/0.75/0.75/0.75/0.75
- BaseDamage: 45/75/105/135/165
- DamageReduction: 0.3/0.3/0.3/0.3/0.3
- DashSpeedAO: 1800/1800/1800/1800/1800
- DashSpeedBase: 1400/1400/1400/1400/1400
- HexflashBugfixLockout: 0.35/0.35/0.35/0.35/0.35
- KnockbackDist: 125/125/125/125/125
- KnockbackDistOnTopOfMod: 0.6/0.6/0.6/0.6/0.6
- KnockbackSpeed: 800/800/800/800/800
- MaxDashBase: 450/450/450/450/450
- MaxDuration: 1/1/1/1/1
- MaxKnockbackDuration: 1.75/1.75/1.75/1.75/1.75
- MaxMonsterDamage: 180/260/340/420/500
- MinChargeTime: 0.4/0.4/0.4/0.4/0.4
- MinDashBase: 100/100/100/100/100
- MinDurationTOOLTIP: 0.4/0.4/0.4/0.4/0.4
- MinKnockbackDuration: 0.5/0.5/0.5/0.5/0.5
- MinimumMinionDamage: 80/130/180/230/280
- RDamageIncreaseMax: 0.8/0.8/0.8/0.8/0.8
- RDamageIncreaseMin: 0.1/0.1/0.1/0.1/0.1
- RDamageReduction: 0.75/0.75/0.75/0.75/0.75
- TimeToFullCharge: 0.9/0.9/0.9/0.9/0.9
- VictimDashMod: 1.3/1.3/1.3/1.3/1.3

**E · Passo Forte** · recarga 10/9.5/9/8.5/8 · custo 45/50/55/60/65

- TotalShield: 70/112.5/155/197.5/240 + 0.135 max_health (bonus)
- AllySpeedAO: 1400/1400/1400/1400/1400
- AllySpeedBase: 1100/1100/1100/1100/1100
- CooldownModAO: 0.5/0.5/0.5/0.5/0.5
- DashArrivial: 500/500/500/500/500
- DashDistance: 550/550/550/550/550
- FreeTargetRangeAO: 400/400/400/400/400
- FreeTargetRangeBase: 250/250/250/250/250
- FreeTargetSpeedAO: 1250/1250/1250/1250/1250
- FreeTargetSpeedBase: 550/550/550/550/550
- ShieldDuration: 2/2/2/2/2
- TargetingForgiveness: 150/150/150/150/150

**R · Forma Irrestrita** · recarga 120/100/80 · custo 100/100/100 · alcance 350/350/350

- TotalDamageSlamDown: 80/115/150 + 0.05 max_health (bonus)
- AfterImageArrivesFirstBySec: 0.2/0.2/0.2
- AllOutDuration: 15/15/15
- AnimationDashTime: 0.3/0.3/0.3
- ArmorPenPercent: 0.5/0.5/0.5
- AttackSpeed: 0.4/0.6/0.8
- AuraOffset: 200/200/200
- BaseDamage: 80/115/150
- CameraHoldRatioMax: 0.7/0.7/0.7
- CameraHoldRatioMin: 0.2/0.2/0.2
- CarrySpeed: 1000/1000/1000
- CarrySpeedNoWall: 800/800/800
- CarrySpeedRatioMax: 1700/1700/1700
- CarrySpeedRatioMin: 900/900/900
- CarrySpeedUnderstatedNoWall: 1600/1600/1600
- CastTime: 0.35/0.35/0.35
- DefensesLost: 0.85/0.85/0.85
- DistanceRatioMax: 2000/2000/2000
- DistanceRatioMin: 600/600/600
- DistanceTriggerFasterCarrySpeed: 1400/1400/1400
- FasterCarrySpeed: 1500/1500/1500
- HealthLost: 0.35/0.35/0.35
- InitialAnimHoldMax: 0.27/0.27/0.27
- KSanteLandingOffset: 250/250/250
- KnockGrav: 14/14/14
- KsanteAnimLockWall: 1/1/1
- KsanteArrivesFirstBySec: 1/1/1
- LingeringPerceptionPostStun: 1.5/1.5/1.5
- ManaRegenPercent: 5/5/5
- MaxSuppressTime: 10/10/10
- MaxWallCarry: 25000/25000/25000
- MaxWallCarrySR: 5000/5000/5000
- NoWallStunDuration: 0.25/0.25/0.25
- Omnivamp: 0.2/0.2/0.2
- TackleRange: 300/300/300
- TackleRangeOverstatedNoWall: 3000/3000/3000
- TackleSpeed: 2000/2000/2000
- TimeToLast: 0.18/0.18/0.18
- VfxInWallPerDistance: 200/200/200
- VictimCameraZoomMax: 0.8/0.8/0.8
- VictimCameraZoomMin: 0.5/0.5/0.5
- WallBuffer: 350/350/350
- WallBufferFailCase: 100/100/100
- WallCheckDist: -25/-25/-25
- WallLookAhead: 100/100/100
- WallStunDuration: 0.5/0.5/0.5

## Smolder

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 575 | 100 | 2275 |
| dano de ataque | 58 | 2.3 | 97.1 |
| armadura | 24 | 4 | 92 |
| resistência mágica | 33 | 1.1 | 51.7 |
| velocidade de ataque | 0.638 | 4 | 68.638 |
| regeneração de vida | 0.75 | 0.12 | 2.79 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Treinamento Dracônico** · recarga 0

- EStacksPerAttackTooltip: 100

**Q · Super-Hálito Flamejante** · recarga 5.5/5/4.5/4/3.5 · custo 25/25/25/25/25 · alcance 550/550/550/550/550

- Tier3_ExecuteThreshold: 6.5/6.5/6.5/6.5/6.5
- TotalDamage: 60/70/80/90/100 + 1.3 attack_damage (bonus)
- LifestealMod: 0.5/0.5/0.5/0.5/0.5
- ManaRestore: 15/15/15/15/15
- StackTier1: 25/25/25/25/25
- StackTier2: 125/125/125/125/125
- StackTier3: 225/225/225/225/225
- Tier1_ExplosionRadius: 285/285/285/285/285
- Tier2_BlowbackDistance: 500/500/500/500/500
- Tier2_BlowbackExplosionRadius: 150/150/150/150/150
- Tier2_BlowbackPercentageDamage: 50/50/50/50/50
- Tier2_BlowbackSpreadAngle: 15/15/15/15/15
- Tier3_DotLength: 3/3/3/3/3
- Tier3_MonsterCap: 300/300/300/300/300
- Tier3_TrueDamagePercent: 0.01/0.01/0.01/0.01/0.01
- minionmod: 1/1/1/1/1

**W · Atchim!** · recarga 14/13/12/11/10 · custo 50/55/60/65/70 · alcance 1500/1500/1500/1500/1500

- ExplosionDamage: 10/35/60/85/110 + 0.8 ability_power + 0.5 attack_damage (bonus)
- InitialDamage: 60/70/80/90/100 + 0.6 attack_damage (bonus)
- ExplosionDamageMult: 1/1/1/1/1
- ExplosionDamageMultihitPenalty: 0.75/0.75/0.75/0.75/0.75
- ExplosionRadius: 385/385/385/385/385
- MinionMod: 1/1/1/1/1
- SlowAmount: 0.35/0.35/0.35/0.35/0.35
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**E · Voa, Voa, Voa** · recarga 24/22/20/18/16 · custo 65/65/65/65/65 · alcance 700/700/700/700/700

- DamagePerHit: 10/15/20/25/30 + 0.3 attack_damage
- Duration: 1.25/1.25/1.25/1.25/1.25
- MoveSpeed: 0.75/0.75/0.75/0.75/0.75
- Range: 700/700/700/700/700

**R · MANHÊÊÊ!** · recarga 120/110/100 · custo 100/100/100 · alcance 4200/4200/4200

- MomHealCalc: 100/135/170 + 0.75 ability_power + 0.5 attack_damage (bonus)
- TooltipOnly_TotalSweetspotDamage (= TotalDamage × 1): 225/375/525 + 1.5 ability_power + 1.5 attack_damage (bonus)
- TotalDamage: 150/250/350 + 1 ability_power + 1 attack_damage (bonus)
- LingeringSimmerSoundAmount: 3/3/3
- LingeringSimmerSoundDelay: 1/1/1
- LingeringSimmerSoundOffset: 1500/1500/1500
- MinionMod: 0.5/0.5/0.5
- MissileLandingSound: 500/500/500
- MomDistBehindOffset: 600/600/600
- MomSpawnDelay: 1/1/1
- SlowAmount: 0.4/0.4/0.4
- SlowDuration: 2/2/2
- SoundDistanceBehindSmolder: 100/100/100
- SoundDistanceBehindSmolderSelfOnly: 1000/1000/1000
- VFXNorthAngleReplacement: 90/90/90

## Milio

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 560 | 88 | 2056 |
| dano de ataque | 48 | 3.2 | 102.4 |
| armadura | 26 | 4.6 | 104.2 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.625 | 3 | 51.625 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 525 | — | 525 |
| multiplicador de crítico | 2 | — | 2 |

**P · A Todo Vapor!** · recarga 0

- ADBurstRatio: 0.15 _(no nível 18)_
- BaseDamageEnd: 50
- BaseDamageStart: 10
- BurnDuration: 1.5
- MaxMonsterDamage: 100

**Q · Ultramega Chute Flamejante** · custo 50/55/60/65/70 · alcance 1200/1200/1200/1200/1200

- Damage: 80/140/200/260/320 + 1.2 ability_power
- SlowAmountPercent: 0.4/0.45/0.5/0.55/0.6 + 0.0005 ability_power
- ExtraVisionDuration: 1.25/1.25/1.25/1.25/1.25
- FallRadius: 250/250/250/250/250
- FallTime: 0.8/0.8/0.8/0.8/0.8
- ImpactAdjustment: 50/50/50/50/50
- KnockbackDistance: 140/140/140/140/140
- KnockbackMinionDistance: 340/340/340/340/340
- MinionFallRadius: 275/275/275/275/275
- MinionFallTime: 0.9/0.9/0.9/0.9/0.9
- RefundRatio: 0.5/0.5/0.5/0.5/0.5
- SlowDuration: 1.5/1.5/1.5/1.5/1.5

**W · Fogueira Aconchegante** · recarga 29/27/25/23/21 · custo 90/100/110/120/130 · alcance 650/650/650/650/650

- HealingOverTime: 70/90/110/130/150 + 0.15 ability_power
- RangePercent: 0.1/0.125/0.15/0.175/0.2
- HealFrequencySeconds: 3/3/3/3/3
- Radius: 350/350/350/350/350
- ZoneDuration: 6/6/6/6/6

**E · Abraços Quentinhos** · recarga 0.5/0.5/0.5/0.5/0.5 · custo 50/60/70/80/90 · alcance 650/650/650/650/650

- ShieldCalc: 45/75/105/135/165 + 0.45 ability_power
- DecayDelay: 0.15/0.15/0.15/0.15/0.15
- MoveSpeedAmount: 0.12/0.14/0.16/0.18/0.2
- MoveSpeedDuration: 2.5/2.5/2.5/2.5/2.5

**R · Sopro de Vida** · recarga 160/145/130 · custo 100/100/100 · alcance 700/700/700

- HealCalc: 150/250/350 + 0.5 ability_power
- HealFXDelay: 0.1/0.1/0.1
- TenacityAmount: 0.65/0.65/0.65
- TenacityDuration: 3/3/3

## Zaahen

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 640 | 114 | 2578 |
| dano de ataque | 63 | 4 | 131 |
| armadura | 36 | 5 | 121 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.625 | 2.5 | 43.125 |
| regeneração de vida | 1.5 | 0.16 | 4.22 |
| velocidade de movimento | 345 | — | 345 |
| alcance de ataque | 175 | — | 175 |
| multiplicador de crítico | 2 | — | 2 |

**P · Cultivação da Guerra** · recarga 100

- ReviveCooldownCalc: 120 _(no nível 18)_
- FallOffDuration: 5
- MaxStacks: 12
- MaxStacksMultiplier: 2
- NumberOfTicks: 32
- PBonusADMax: 0.028
- PBonusADMin: 0.015
- PercentBonusAD: 0.025
- ReviveAdditionalBonusAtThisLevel: 0.15
- ReviveDuration: 4
- ReviveLevel1Value: 0.3
- RevivePercent: 0.6

**Q · A Glaive Darkin** · recarga 10/9/8/7/6 · custo 25/25/25/25/25

- InitialDamage: 15/30/45/60/75 + 0.2 attack_damage (bonus)
- SecondHitDamage: 25/50/75/100/125 + 0.2 attack_damage (bonus)
- HealPercent: 0.05/0.06/0.07/0.08/0.09
- KnockUpDuration: 0.75/0.75/0.75/0.75/0.75
- MinionHealPercent: 0.5/0.5/0.5/0.5/0.5
- MonsterDamagePercent: 1/1/1/1/1
- QBonusRange: 25/25/25/25/25
- RecastWindow: 4/4/4/4/4
- TimeBetweenAttacks: 1.5/1.5/1.5/1.5/1.5

**W · Temível Retorno** · recarga 14/13.5/13/12.5/12 · custo 50/50/50/50/50 · alcance 850/850/850/850/850

- InitialDamage: 40/60/80/100/120 + 0.5 attack_damage (bonus)
- SecondaryDamage: 30/50/70/90/110 + 0.3 attack_damage (bonus)
- MaximumPullDistance: 225/225/225/225/225
- PullSpeed: 900/900/900/900/900

**E · Ímpeto Áureo** · recarga 10/9.5/9/8.5/8 · custo 40/40/40/40/40

- BaseDamageCalc: 40/60/80/100/120 + 0.5 attack_damage (bonus)
- BonusDamageCalc (= BaseDamageCalc × 1): 60/90/120/150/180 + 0.75 attack_damage (bonus)
- DamageBonusBase: 60/90/120/150/180
- DashDistance: 350/350/350/350/350
- DashSpeed: 900/900/900/900/900
- MaximumRadius: 375/375/375/375/375
- MinimumRadius: 200/200/200/200/200
- MonsterDamageBonus: 50/50/50/50/50
- MonsterDamageCap: 400/400/400/400/400
- PercentHPDamage: 0.04/0.045/0.05/0.055/0.06

**R · Desfecho Fatídico** · recarga 110/95/80 · custo 100/100/100 · alcance 600/600/600

- DamageEndCalc: 250/400/550 + 2 attack_damage (bonus)
- AoESize: 550/550/550
- ArmorPen: 0.1/0.2/0.3
- DamageReduction: 0.5/0.5/0.5
- DashMaxDistance: 600/600/600
- DashMinDistance: 10/10/10
- DashSpeed: 2800/2800/2800
- HealPercent: 0.33/0.33/0.33

## Hwei

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 580 | 109 | 2433 |
| dano de ataque | 54 | 3.3 | 110.1 |
| armadura | 21 | 4.7 | 100.899994 |
| resistência mágica | 30 | 1.3 | 52.1 |
| velocidade de ataque | 0.69 | 2.5 | 43.19 |
| regeneração de vida | 1.1 | 0.11 | 2.97 |
| velocidade de movimento | 330 | — | 330 |
| alcance de ataque | 550 | — | 550 |
| multiplicador de crítico | 2 | — | 2 |

**P · Assinatura do Visionário**

- TotalDamage: 285 + 0.35 ability_power _(no nível 18)_
- Duration: 4
- InkTimeout: 25000

**Q · Tema: Desastre** · recarga 10/9/8/7/6 · custo 80/90/100/110/120

- Tooltip_QEDamage: 20/35/50/65/80 + 0.3 ability_power
- Tooltip_QEDamagePerSecond: 20/35/50/65/80 + 0.24 ability_power
- Tooltip_QQDamage: 50/80/110/140/170 + 0.8 ability_power
- Tooltip_QWBonusDamage (= Tooltip_QWDamage × 1): 120/201.875/302.5/421.875/560 + 0.6 ability_power
- Tooltip_QWDamage: 60/85/110/135/160 + 0.3 ability_power
- Tooltip_QQBonusDamage: 3/4/5/6/7

**W · Tema: Serenidade** · recarga 18/17.5/17/16.5/16 · custo 90/95/100/105/110

- Tooltip_WEOnHitDamage: 20/30/40/50/60 + 0.15 ability_power
- Tooltip_WQMoveSpeed: 30/32.5/35/37.5/40 + 0.03 ability_power
- Tooltip_WWShieldAmount: 100/140/180/220/260 + 0.6 ability_power
- Tooltip_WWStartShieldAmount: 100/140/180/220/260 + 0.6 ability_power
- Tooltip_WEOnHitManaRestore: 45/50/55/60/65
- Tooltip_WQAreaDuration: 4/4.5/5/5.5/6

**E · Tema: Tormento** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70

- Tooltip_EEDamage: 70/110/150/190/230 + 0.65 ability_power
- Tooltip_EQDamage: 70/110/150/190/230 + 0.65 ability_power
- Tooltip_EWDamage: 70/110/150/190/230 + 0.65 ability_power
- Tooltip_EESlowAmount: 40/47.5/55/62.5/70
- Tooltip_EQFleeDuration: 1/1.125/1.25/1.375/1.5
- Tooltip_EWRootDuration: 1.2/1.4/1.6/1.8/2

**R · Desespero Vertiginoso** · recarga 120/100/80 · custo 100/100/100 · alcance 1300/1300/1300

- Damage: 200/325/450 + 0.8 ability_power
- DamageOverTime: 10/20/30 + 0.05 ability_power
- TotalMaxDamage: 230/385/540 + 0.95 ability_power
- MaxDamageHelper: 230/360/490
- SlowDuration: 0.35/0.35/0.35
- SlowPercentPerStack: 10/10/10

**QQ · Fogo Devastador** · recarga 10/9/8/7/6 · custo 80/90/100/110/120 · alcance 800/800/800/800/800

- Damage: 50/80/110/140/170 + 0.8 ability_power
- BonusDamage: 3/4/5/6/7
- MaxBonusDamage: 250/250/250/250/250
- Radius: 175/175/175/175/175

**QW · Raio Cortante** · recarga 10/9/8/7/6 · custo 80/90/100/110/120 · alcance 1900/1900/1900/1900/1900

- BonusDamage (= Damage × 1): 120/201.875/302.5/421.875/560 + 0.6 ability_power
- Damage: 60/85/110/135/160 + 0.3 ability_power
- Delay: 1/1/1/1/1
- MinionDamageMod: 0.5/0.5/0.5/0.5/0.5
- Radius: 225/225/225/225/225

**QE · Fissura Derretida** · recarga 10/9/8/7/6 · custo 80/90/100/110/120 · alcance 1200/1200/1200/1200/1200

- Damage: 20/35/50/65/80 + 0.3 ability_power
- DamagePerSecond: 20/35/50/65/80 + 0.24 ability_power
- {a8f705ac}: 30/30/30/30/30 _(no nível 18)_
- DamageDuration: 0.5/0.5/0.5/0.5/0.5
- Duration: 2.5/2.5/2.5/2.5/2.5
- JungleDamageRatio: 1.35/1.35/1.35/1.35/1.35
- MinionDamageRatio: 0.6/0.6/0.6/0.6/0.6
- Radius: 225/225/225/225/225
- SlowPercent: 35/35/35/35/35

**WQ · Corrente Fugaz** · recarga 18/17.5/17/16.5/16 · custo 90/95/100/105/110 · alcance 1200/1200/1200/1200/1200

- MoveSpeed: 30/32.5/35/37.5/40 + 0.03 ability_power
- AreaDuration: 4/4.5/5/5.5/6
- AreaWidth: 300/300/300/300/300
- Duration: 1/1/1/1/1

**WW · Lagoa da Reflexão** · recarga 18/17.5/17/16.5/16 · custo 90/95/100/105/110 · alcance 650/650/650/650/650

- ShieldAmount: 100/140/180/220/260 + 0.6 ability_power
- {20a3d822}: 100/140/180/220/260 + 0.6 ability_power
- AllyMod: 0.85/0.85/0.85/0.85/0.85
- AreaDuration: 3/3/3/3/3
- SecsToFullStack: 3/3/3/3/3
- ShieldDuration: 1/1/1/1/1
- ToolTipAllyMod: 0.15/0.15/0.15/0.15/0.15
- ZoneRadius: 350/350/350/350/350

**WE · Luzes Vivas** · recarga 18/17.5/17/16.5/16 · custo 90/95/100/105/110 · alcance 500/500/500/500/500

- OnHitDamage: 20/30/40/50/60 + 0.15 ability_power
- TotalOnHitDamage (= OnHitDamage × 3): 60/90/120/150/180 + 0.45000002 ability_power
- {a14e3318}: 48/53/58/63/68
- Duration: 9/9/9/9/9
- ManaCostHelper: -90/-95/-100/-105/-110

**EQ · Semblante Sinistro** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70 · alcance 1100/1100/1100/1100/1100

- Damage: 70/110/150/190/230 + 0.65 ability_power
- FleeDuration: 1/1.125/1.25/1.375/1.5
- MaxSlow: 99/99/99/99/99
- MinSlow: 70/70/70/70/70

**EW · Olhar do Abismo** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70 · alcance 900/900/900/900/900

- Damage: 70/110/150/190/230 + 0.65 ability_power
- FireDelay: 0.3/0.3/0.3/0.3/0.3
- MinSetupDelay: 0.65/0.65/0.65/0.65/0.65
- MinStareDuration: 3/3/3/3/3
- RootDuration: 1.2/1.4/1.6/1.8/2
- SetupDelay: 0.65/0.65/0.65/0.65/0.65
- StareDuration: 3/3/3/3/3
- StareRadius: 350/350/350/350/350
- VisibilityRadius: 450/450/450/450/450

**EE · Gorja Esmagadora** · recarga 12/11.5/11/10.5/10 · custo 50/55/60/65/70 · alcance 800/800/800/800/800

- Damage: 70/110/150/190/230 + 0.65 ability_power
- Delay: 0.6/0.6/0.6/0.6/0.6
- Depth: 340/340/340/340/340
- SlowAmount: 40/47.5/55/62.5/70
- Width: 320/320/320/320/320

## Naafiri

| estatística | base | por nível | no 18 |
|---|---:|---:|---:|
| vida | 610 | 105 | 2395 |
| dano de ataque | 55 | 2 | 89 |
| armadura | 28 | 4.2 | 99.399994 |
| resistência mágica | 32 | 2.05 | 66.85 |
| velocidade de ataque | 0.663 | 2.1 | 36.363 |
| regeneração de vida | 1.25 | 0.12 | 3.29 |
| velocidade de movimento | 340 | — | 340 |
| alcance de ataque | 125 | — | 125 |
| multiplicador de crítico | 2 | — | 2 |

**P · Em Maior Número**

- AOEDamageModifier: 0.76
- FrenzyDamageTooltipOnly (= PackmateTotalDamage × 1): 26 + 0.051999997 attack_damage (bonus) _(no nível 18)_
- MinionExecuteThreshold: 25
- PackmateCap: 5 _(no nível 18)_
- PackmateSpawnCooldown: 10 _(no nível 18)_
- PackmateTotalDamage: 20 + 0.04 attack_damage (bonus) _(no nível 18)_
- CooldownReduceOnAbilityHit: 4
- CooldownReduceOnKill: 1
- HealthRegenDelay: 5
- HealthRegenPercent: 0.25
- LeapMax: 1200
- LeapMin: 75
- LeapSpeed: 1250
- MeleeAttackDamageModifier: 2
- NonChampDamageModifier: 0.5
- PackmateBaseAS: 0.688
- PackmateBaseMS: 340
- PackmateCSHelpRadiusMax: 800
- PackmateCSHelpRadiusMin: 300
- PackmateFrenzyDuration: 3
- PackmateFrenzyMaxCounters: 3
- PackmateLastHitHelpCD: 10
- PackmateModifierTooltipOnly: 0.08
- PackmateMonsterMod: 1.55
- PackmateTauntDuration: 2
- PackmateTowerModifier: 0.5
- ProcDamageModifier: 0.5

**Q · Adagas Darkin** · recarga 9/8.5/8/7.5/7 · custo 50/60/70/80/90 · alcance 900/900/900/900/900

- MinionExecuteThreshold: 30/30/30/30/30
- TotalBleedDamage: 35/60/85/110/135 + 0.8 attack_damage (bonus)
- TotalDamageFirstCast: 35/40/45/50/55 + 0.2 attack_damage (bonus)
- TotalHealSecondCast: 45/60/75/90/105 + 0.4 attack_damage (bonus)
- TotalMaxDamageSecondCast: 30/42.5/55/67.5/80 + 0.7 attack_damage (bonus)
- TotalMinDamageSecondCast: 30/42.5/55/67.5/80 + 0.4 attack_damage (bonus)
- BleedDuration: 5/5/5/5/5
- BleedInterval: 0.5/0.5/0.5/0.5/0.5
- MinionMod: 1/1/1/1/1
- RecastLockout: 0.75/0.75/0.75/0.75/0.75
- RecastWindow: 4/4/4/4/4

**W · Chamado da Matilha** · recarga 26/24/22/20/18 · custo 60/60/60/60/60

- BonusAD: 0/0/0/0/0 + 0.2 attack_damage
- Duration: 5/5/5/5/5
- MoveSpeedAmount: 0.2/0.225/0.25/0.275/0.3
- MoveSpeedDuration: 5/5/5/5/5
- PackmateSpawnRateIncrease: 0.5/0.5/0.5/0.5/0.5
- PackmatesToAdd: 2/2/2/2/2
- UntargetableDuration: 1/1/1/1/1

**E · Eviscerar** · recarga 11/10/9/8/7 · custo 40/40/40/40/40 · alcance 450/450/450/450/450

- TotalDamageFirstSlash: 15/25/35/45/55 + 0.4 attack_damage (bonus)
- TotalDamageSecondSlash: 60/85/110/135/160 + 0.8 attack_damage (bonus)
- DashDistance: 450/450/450/450/450
- DashSpeed: 900/900/900/900/900
- MaxDashCheatDistance: 650/650/650/650/650
- MinDashDistance: 250/250/250/250/250
- RadiusSecondSlash: 230/230/230/230/230

**R · Caça dos Cães** · recarga 110/95/80 · custo 100/100/100 · alcance 900/900/900

- ArmorShred: 6/6/6
- PackmateDamage (= TotalDamage × 1): 12.5/20/27.5 + 0.1 attack_damage (bonus)
- ShieldTotal: 100/150/200 + 1.5 attack_damage (bonus)
- TotalDamage: 125/200/275 + 1 attack_damage (bonus)
- DashSpeed: 1800/1800/1800
- FailCastCDRefund: 0.5/0.5/0.5
- RecastWindow: 12/12/12
- ShieldDuration: 3/3/3
- ShredDuration: 3/3/3
- SlowDuration: 0.25/0.25/0.25
- SlowPercent: -0.99/-0.99/-0.99
- TakedownWindow: 7/7/7
- VisionDuration: 4/4/4
- VisionPulseDuration: 1/1/1
- VisionRadius: 2100/2100/2100
- WRefreshOnCast: 1.75/1.75/1.75

## Sem número na fonte

12 habilidades cuja fórmula a fonte não publica ou o extrator não
resolve. Estão aqui nomeadas uma a uma, e **não** publicadas com zero:
zero afirmaria que a habilidade não causa dano, enquanto a ausência apenas
não informa — e só a segunda é recuperável por quem lê.

- Aphelios E (Sistema de Ordenação de Armas)
- Aphelios Q (Habilidades da arma)
- Aphelios W (Fase)
- Cho'Gath W (Grito Selvagem)
- Fiddlesticks P (Um Espantalho Inofensivo)
- Hwei ER (Lavar Pincel)
- Hwei QR (Lavar Pincel)
- Hwei WR (Lavar Pincel)
- Kai'Sa E (Sobrecarga)
- Kai'Sa W (Exploradora do Vazio)
- Nidalee P (Espreitar)
- Zoe R (Salto Dimensional)

