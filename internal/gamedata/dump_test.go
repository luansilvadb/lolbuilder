package gamedata

import (
	"strings"
	"testing"
)

// dumpMinimo reproduz a forma real: o registro do modo, um registro de OUTRO
// modo, uma habilidade e uma chave opaca.
const dumpMinimo = `{
  "Characters/Teste/CharacterRecords/Root": {
    "__type": "CharacterRecord",
    "mCharacterName": "Teste",
    "baseHPModifiable": {"baseValue": 690, "__type": "ModifiableFloat"},
    "hpPerLevelModifiable": {"baseValue": 98, "__type": "ModifiableFloat"},
    "baseStaticHPRegenModifiable": {"baseValue": 1.6, "__type": "ModifiableFloat"},
    "hpRegenPerLevelModifiable": {"baseValue": 0.1, "__type": "ModifiableFloat"},
    "baseDamageModifiable": {"baseValue": 69, "__type": "ModifiableFloat"},
    "damagePerLevelModifiable": {"baseValue": 4.5, "__type": "ModifiableFloat"},
    "baseArmorModifiable": {"baseValue": 38, "__type": "ModifiableFloat"},
    "armorPerLevelModifiable": {"baseValue": 4.2, "__type": "ModifiableFloat"},
    "attackSpeedModifiable": {"baseValue": 0.625, "__type": "ModifiableFloat"},
    "attackSpeedRatioModifiable": {"baseValue": 0.625, "__type": "ModifiableFloat"},
    "attackSpeedPerLevelModifiable": {"baseValue": 3.65, "__type": "ModifiableFloat"},
    "baseMR": {"baseValue": 32, "__type": "ModifiableFloat"},
    "mrPerLevel": {"baseValue": 1.55, "__type": "ModifiableFloat"},
    "baseMoveSpeedModifiable": {"baseValue": 340, "__type": "ModifiableFloat"},
    "attackRangeModifiable": {"baseValue": 175, "__type": "ModifiableFloat"},
    "critDamageMultiplier": 2.0,
    "spells": ["Characters/Teste/Spells/Q", "Characters/Teste/Spells/W",
               "Characters/Teste/Spells/E", "Characters/Teste/Spells/R"],
    "spellNames": ["Q", "W", "E", "R"],
    "mCharacterPassiveSpell": "{aabbccdd}",
    "passiveName": "Passiva",
    "{deadbeef}": {"lixo": 1}
  },
  "Characters/Teste/CharacterRecords/URF": {
    "__type": "CharacterRecord",
    "mCharacterName": "Teste",
    "baseHPModifiable": {"baseValue": 9999, "__type": "ModifiableFloat"},
    "spells": []
  },
  "Characters/Teste/Spells/Q": {"__type": "SpellObject", "ObjectName": "TesteQ",
    "mSpell": {"Cooldown": {"values": [0,8,8,8,8,8,8]}, "cooldownTime": [0,9,9,9,9,9,9],
               "manaValues": {"values": [55,65,75,85,95,95]}, "mana": [1,1,1,1,1,1],
               "castRange": [300,300,300,300,300,300]}},
  "Characters/Teste/Spells/W": {"__type": "SpellObject", "ObjectName": "TesteW", "mSpell": {}},
  "Characters/Teste/Spells/E": {"__type": "SpellObject", "ObjectName": "TesteE", "mSpell": {}},
  "Characters/Teste/Spells/R": {"__type": "SpellObject", "ObjectName": "TesteR", "mSpell": {}},
  "{aabbccdd}": {"__type": "SpellObject", "ObjectName": "TestePassiva", "mSpell": {}}
}`

func parse(t *testing.T) *Dump {
	t.Helper()
	d, err := ParseDump("teste", []byte(dumpMinimo))
	if err != nil {
		t.Fatalf("ParseDump falhou: %v", err)
	}
	return d
}

// TestEscolheORegistroDoModo e a guarda que impede publicar estatistica de URF
// como se fosse do Summoner's Rift. Tres campeoes do 16.16 publicam mais de um
// CharacterRecord no mesmo arquivo.
func TestEscolheORegistroDoModo(t *testing.T) {
	d := parse(t)
	if d.Record.BaseHP.BaseValue != 690 {
		t.Fatalf("vida base = %v; pegou o registro de outro modo", d.Record.BaseHP.BaseValue)
	}
}

func TestFalhaSemRegistroDoModo(t *testing.T) {
	semRoot := strings.Replace(dumpMinimo, "/CharacterRecords/Root", "/CharacterRecords/ARAM", 1)
	_, err := ParseDump("teste", []byte(semRoot))
	if err == nil {
		t.Fatal("dump sem o registro do modo passou")
	}
	if !strings.Contains(err.Error(), "outros modos") {
		t.Fatalf("o erro nao explica o que houve: %v", err)
	}
}

// TestChaveOpacaNaoAlarma: um hash novo so significa que o dicionario do
// CommunityDragon avancou, o que nao diz nada sobre o jogo. Um campo com NOME
// novo e que precisa alarmar.
func TestChaveOpacaNaoAlarma(t *testing.T) {
	parse(t) // o dump de teste ja tem {deadbeef} dentro do registro
}

func TestCampoNovoNoRegistroAlarma(t *testing.T) {
	comCampoNovo := strings.Replace(dumpMinimo,
		`"passiveName": "Passiva",`,
		`"passiveName": "Passiva", "campoQueNaoExistia": 1,`, 1)
	_, err := ParseDump("teste", []byte(comCampoNovo))
	if err == nil {
		t.Fatal("campo novo no CharacterRecord passou em silencio")
	}
	if !strings.Contains(err.Error(), "campoQueNaoExistia") {
		t.Fatalf("o erro nao nomeia o campo: %v", err)
	}
}

func TestStats(t *testing.T) {
	s, gaps := parse(t).Stats()
	if len(gaps) != 0 {
		t.Fatalf("dump completo reportou lacunas: %v", gaps)
	}
	if !s.Complete() {
		t.Fatal("estatisticas completas nao foram reconhecidas")
	}
	if s.HP.At(1) != 690 {
		t.Errorf("vida no nivel 1 = %v", s.HP.At(1))
	}
	if got := s.HP.At(18); got != 690+98*17 {
		t.Errorf("vida no nivel 18 = %v, esperado %v", got, 690+98*17)
	}
	if !s.Melee {
		t.Error("alcance 175 nao foi classificado como corpo a corpo")
	}
}

// TestLacunaNaoViraZero: um campeao sem regeneracao publicada e um campeao cuja
// regeneracao nao sabemos. Afirmar "0 de regeneracao" seria pior que nao
// afirmar nada, porque o consumidor somaria esse zero.
func TestLacunaNaoViraZero(t *testing.T) {
	semRegen := strings.Replace(dumpMinimo,
		`"baseStaticHPRegenModifiable": {"baseValue": 1.6, "__type": "ModifiableFloat"},`, "", 1)
	d, err := ParseDump("teste", []byte(semRegen))
	if err != nil {
		t.Fatal(err)
	}
	s, gaps := d.Stats()
	if s.HPRegen != nil {
		t.Fatalf("regeneracao ausente virou %+v", *s.HPRegen)
	}
	if len(gaps) != 1 || !strings.Contains(gaps[0], "regeneracao") {
		t.Fatalf("a lacuna nao foi reportada: %v", gaps)
	}
	// A ausencia de uma opcional nao derruba o campeao.
	if !s.Complete() {
		t.Error("a ausencia de uma opcional foi tratada como campeao incompleto")
	}
}

func TestSpellsResolveSlotsEPassiva(t *testing.T) {
	spells, err := parse(t).Spells()
	if err != nil {
		t.Fatal(err)
	}
	if len(spells) != 5 {
		t.Fatalf("resolveu %d habilidades, esperado 4 slots mais a passiva", len(spells))
	}
	if spells[0].Slot != SlotQ || spells[0].Name != "TesteQ" {
		t.Errorf("slot Q errado: %+v", spells[0])
	}
	// A passiva e referenciada por um identificador opaco, que e ele proprio uma
	// chave de topo do dump.
	if spells[4].Slot != SlotPassive || spells[4].Name != "TestePassiva" {
		t.Errorf("passiva errada: %+v", spells[4])
	}
	if spells[3].Ranks() != 3 || spells[0].Ranks() != 5 || spells[4].Ranks() != 1 {
		t.Error("contagem de ranks errada por slot")
	}
}

// TestSerieRedefinidaVenceAHerdada: medido no 16.16, a redefinida concorda 100%
// com o plugin e a herdada 98.5%. Ler a herdada publica numero que o jogo nao
// usa.
func TestSerieRedefinidaVenceAHerdada(t *testing.T) {
	spells, err := parse(t).Spells()
	if err != nil {
		t.Fatal(err)
	}
	q := spells[0]
	if got := q.CooldownTime(); len(got) != 7 || got[1] != 8 {
		t.Fatalf("recarga = %v; deveria vir de Cooldown, nao de cooldownTime", got)
	}
	if got := q.Mana(); len(got) != 6 || got[0] != 55 {
		t.Fatalf("custo = %v; deveria vir de manaValues, nao de mana", got)
	}
}

// TestAtRankAplicaODeslocamentoDoTamanho: as duas series do mesmo mSpell nao
// concordam entre si. cooldownTime tem 7 posicoes e comeca no indice 1; mana tem
// 6 e comeca no indice 0.
func TestAtRankAplicaODeslocamentoDoTamanho(t *testing.T) {
	sete := []float64{0, 8, 9, 10, 11, 12, 12}
	seis := []float64{55, 65, 75, 85, 95, 95}

	if v, ok := AtRank(sete, 1); !ok || v != 8 {
		t.Errorf("serie de 7: rank 1 = %v (%v), esperado 8", v, ok)
	}
	if v, ok := AtRank(seis, 1); !ok || v != 55 {
		t.Errorf("serie de 6: rank 1 = %v (%v), esperado 55", v, ok)
	}
	if _, ok := AtRank(seis, 9); ok {
		t.Error("rank fora da serie devolveu valor")
	}
	if RankOffset(len(sete)) != 1 || RankOffset(len(seis)) != 0 {
		t.Error("deslocamento errado por tamanho")
	}
}

func TestDataValueEncontraPorNomeIgnorandoCaixa(t *testing.T) {
	sp := &Spell{DataValues: []DataValue{{Name: "AOEDamagePercent", Values: []float64{0, 1, 2}}}}
	if _, ok := sp.DataValue("aoedamagepercent"); !ok {
		t.Error("busca por nome nao ignorou a caixa; a fonte referencia com caixa diferente da declaracao")
	}
	if _, ok := sp.DataValue("naoexiste"); ok {
		t.Error("serie inexistente foi encontrada")
	}
}
