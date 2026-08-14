// Package gamedata le o dump de dados do jogo (game/data/characters), a unica
// fonte com estatistica base de campeao e dano de habilidade.
//
// O arquivo do plugin (champions/<id>.json) traz esses campos presentes e
// zerados; aqui eles existem de verdade.
//
// O dump nao e uma API. E a arvore de objetos do jogo serializada com nomes de
// campo resolvidos por um dicionario incompleto — o que o dicionario nao
// resolve vira uma chave opaca como {c1984296}. Por isso a leitura e em duas
// camadas: mapa cru sobre o arquivo inteiro, e decodificacao tipada so no que
// entendemos.
package gamedata

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/model"
)

// opaqueKey casa com os identificadores que o dicionario do CommunityDragon
// ainda nao resolveu, no formato {a1b2c3d4}.
var opaqueKey = regexp.MustCompile(`^\{[0-9a-f]{8}\}$`)

// ModFloat e o envelope que o dump usa para quase todo numero do
// CharacterRecord.
type ModFloat struct {
	BaseValue float64 `json:"baseValue"`
	Type      string  `json:"__type"`
}

// Record e o CharacterRecord: as estatisticas base do campeao e os ponteiros
// para as habilidades.
//
// Todos os 80 campos nao opacos observados nos 173 campeoes estao mapeados,
// mesmo os que nao usamos — e o que faz um campo novo virar erro em vez de dado
// perdido em silencio. Os que nao usamos ficam como json.RawMessage de
// proposito: mapear o nome basta para o alarme funcionar, e modelar o conteudo
// seria manutencao paga por dado que ninguem le.
type Record struct {
	CharacterName string `json:"mCharacterName"`
	Name          string `json:"name"`

	// Estatisticas escalaveis: valor no nivel 1 mais crescimento por nivel.
	BaseHP            ModFloat  `json:"baseHPModifiable"`
	HPPerLevel        ModFloat  `json:"hpPerLevelModifiable"`
	BaseHPRegen       *ModFloat `json:"baseStaticHPRegenModifiable"`
	HPRegenPerLevel   ModFloat  `json:"hpRegenPerLevelModifiable"`
	BaseDamage        ModFloat  `json:"baseDamageModifiable"`
	DamagePerLevel    ModFloat  `json:"damagePerLevelModifiable"`
	BaseArmor         ModFloat  `json:"baseArmorModifiable"`
	ArmorPerLevel     ModFloat  `json:"armorPerLevelModifiable"`
	BaseAttackSpeed   ModFloat  `json:"attackSpeedModifiable"`
	AttackSpeedRatio  ModFloat  `json:"attackSpeedRatioModifiable"`
	AttackSpeedPerLvl ModFloat  `json:"attackSpeedPerLevelModifiable"`

	// A resistencia magica cresce por nivel para o elenco inteiro no LoL
	// moderno. A fonte omite o campo para quem nao cresce, e a omissao vale
	// zero — nao vale "ausente". Por isso ModFloat e nao *ModFloat.
	BaseMR     ModFloat `json:"baseMR"`
	MrPerLevel ModFloat `json:"mrPerLevel"`

	// Estatisticas fixas.
	BaseMoveSpeed        ModFloat `json:"baseMoveSpeedModifiable"`
	AttackRange          ModFloat `json:"attackRangeModifiable"`
	CritDamageMultiplier *float64 `json:"critDamageMultiplier"`

	// Ponteiros para as habilidades. Spells e ordenado: Q, W, E, R.
	Spells       []string `json:"spells"`
	SpellNames   []string `json:"spellNames"`
	PassiveSpell string   `json:"mCharacterPassiveSpell"`
	PassiveName  string   `json:"passiveName"`
	PassiveTip   string   `json:"passiveToolTip"`

	// Mapeados para o alarme de campo novo, nao lidos.
	PackmanagerData                     json.RawMessage `json:"PackmanagerData"`
	RecSpellRankUpInfolist              json.RawMessage `json:"RecSpellRankUpInfolist"`
	TreatAutoAttacksAsNormalSpells      json.RawMessage `json:"TreatAutoAttacksAsNormalSpells"`
	AcquisitionRange                    json.RawMessage `json:"acquisitionRange"`
	AllyChampSpecificHealthSuffix       json.RawMessage `json:"allyChampSpecificHealthSuffix"`
	BasicAttack                         json.RawMessage `json:"basicAttack"`
	CharAudioNameOverride               json.RawMessage `json:"charAudioNameOverride"`
	CharacterToolData                   json.RawMessage `json:"characterToolData"`
	CritAttacks                         json.RawMessage `json:"critAttacks"`
	CriticalAttack                      json.RawMessage `json:"criticalAttack"`
	DeathEventListeningRadius           json.RawMessage `json:"deathEventListeningRadius"`
	EnemyChampSpecificHealthSuffix      json.RawMessage `json:"enemyChampSpecificHealthSuffix"`
	EnemyTooltip                        json.RawMessage `json:"enemyTooltip"`
	EvolutionData                       json.RawMessage `json:"evolutionData"`
	ExpGivenOnDeath                     json.RawMessage `json:"expGivenOnDeath"`
	ExtraAttacks                        json.RawMessage `json:"extraAttacks"`
	ExtraSpells                         json.RawMessage `json:"extraSpells"`
	ExtraSpellsList                     json.RawMessage `json:"extraSpellsList"`
	Flags                               json.RawMessage `json:"flags"`
	FriendlyTooltip                     json.RawMessage `json:"friendlyTooltip"`
	FriendlyUxOverrideExcludeTagsString json.RawMessage `json:"friendlyUxOverrideExcludeTagsString"`
	FriendlyUxOverrideIncludeTagsString json.RawMessage `json:"friendlyUxOverrideIncludeTagsString"`
	FriendlyUxOverrideTeam              json.RawMessage `json:"friendlyUxOverrideTeam"`
	GoldGivenOnDeath                    json.RawMessage `json:"goldGivenOnDeath"`
	HealthBarHeight                     json.RawMessage `json:"healthBarHeight"`
	HighlightHealthbarIcons             json.RawMessage `json:"highlightHealthbarIcons"`
	MAbilities                          json.RawMessage `json:"mAbilities"`
	MAbilitySlotCC                      json.RawMessage `json:"mAbilitySlotCC"`
	MAdaptiveForceToAbilityPowerWeight  json.RawMessage `json:"mAdaptiveForceToAbilityPowerWeight"`
	MCharacterCalculations              json.RawMessage `json:"mCharacterCalculations"`
	MCharacterPassiveBuffs              json.RawMessage `json:"mCharacterPassiveBuffs"`
	MEducationToolData                  json.RawMessage `json:"mEducationToolData"`
	MFallbackCharacterName              json.RawMessage `json:"mFallbackCharacterName"`
	MPerkReplacements                   json.RawMessage `json:"mPerkReplacements"`
	MPreferredPerkStyle                 json.RawMessage `json:"mPreferredPerkStyle"`
	MUseCCAnimations                    json.RawMessage `json:"mUseCCAnimations"`
	OutlineBBoxExpansion                json.RawMessage `json:"outlineBBoxExpansion"`
	OverrideGameplayCollisionRadius     json.RawMessage `json:"overrideGameplayCollisionRadius"`
	ParName                             json.RawMessage `json:"parName"`
	Passive1IconName                    json.RawMessage `json:"passive1IconName"`
	PassiveLuaName                      json.RawMessage `json:"passiveLuaName"`
	PassiveRange                        json.RawMessage `json:"passiveRange"`
	PassiveSpellRef                     json.RawMessage `json:"passiveSpell"`
	PathfindingCollisionRadius          json.RawMessage `json:"pathfindingCollisionRadius"`
	PlatformEnabled                     json.RawMessage `json:"platformEnabled"`
	PrimaryAbilityResource              json.RawMessage `json:"primaryAbilityResource"`
	PurchaseIdentities                  json.RawMessage `json:"purchaseIdentities"`
	SecondaryAbilityResource            json.RawMessage `json:"secondaryAbilityResource"`
	SelectionHeight                     json.RawMessage `json:"selectionHeight"`
	SelectionRadius                     json.RawMessage `json:"selectionRadius"`
	SelfCBChampSpecificHealthSuffix     json.RawMessage `json:"selfCBChampSpecificHealthSuffix"`
	SelfChampSpecificHealthSuffix       json.RawMessage `json:"selfChampSpecificHealthSuffix"`
	SilhouetteAttachmentAnim            json.RawMessage `json:"silhouetteAttachmentAnim"`
	UnitTagsString                      json.RawMessage `json:"unitTagsString"`
	UseRiotRelationships                json.RawMessage `json:"useRiotRelationships"`
	WeaponMaterials                     json.RawMessage `json:"weaponMaterials"`

	Type string `json:"__type"`
}

// Dump e um arquivo do dump ja separado entre o registro do campeao e os
// objetos que ele referencia.
type Dump struct {
	Alias  string
	Record Record

	// objects guarda todas as chaves de topo ainda cruas, para resolver as
	// referencias por caminho e por identificador opaco.
	objects map[string]json.RawMessage
}

// sufixoDoRegistroDoModo e o registro que vale no Summoner's Rift.
//
// Nao e detalhe: 3 campeoes do 16.16 publicam mais de um CharacterRecord no
// mesmo arquivo. Braum e Milio tem Root, URF e SLIME; Cassiopeia tem Root e
// SLIME. Escolher por __type pegaria qualquer um deles, e num patch qualquer o
// dataset passaria a publicar estatistica de URF como se fosse do Rift, sem
// sintoma nenhum.
const sufixoDoRegistroDoModo = "/CharacterRecords/Root"

// ParseDump le um arquivo do dump. alias serve so para as mensagens de erro.
func ParseDump(alias string, data []byte) (*Dump, error) {
	var objects map[string]json.RawMessage
	if err := json.Unmarshal(data, &objects); err != nil {
		return nil, fmt.Errorf("dump de %s: %w", alias, err)
	}

	raw, err := findRecord(alias, objects)
	if err != nil {
		return nil, err
	}
	rec, err := decodeRecord(alias, raw)
	if err != nil {
		return nil, err
	}
	return &Dump{Alias: alias, Record: *rec, objects: objects}, nil
}

// findRecord localiza o CharacterRecord do modo.
//
// Casa pelo sufixo do caminho E pelo __type. O caminho sozinho seria fragil se
// a Riot renomeasse a arvore; o __type sozinho nao distingue o registro do Rift
// dos registros de outros modos.
func findRecord(alias string, objects map[string]json.RawMessage) (json.RawMessage, error) {
	var found json.RawMessage
	var outros int

	for key, raw := range objects {
		var probe struct {
			Type string `json:"__type"`
		}
		if err := json.Unmarshal(raw, &probe); err != nil {
			continue // valores que nao sao objeto, como a lista __linked
		}
		if probe.Type != "CharacterRecord" {
			continue
		}
		if !strings.HasSuffix(key, sufixoDoRegistroDoModo) {
			outros++
			continue
		}
		if found != nil {
			return nil, fmt.Errorf("dump de %s: mais de um CharacterRecord em %s",
				alias, sufixoDoRegistroDoModo)
		}
		found = raw
	}

	if found == nil {
		return nil, fmt.Errorf(
			"dump de %s: nenhum CharacterRecord em %s (%d registro(s) de outros modos) — "+
				"o formato do dump mudou, ou o campeao saiu do Summoner's Rift",
			alias, sufixoDoRegistroDoModo, outros)
	}
	return found, nil
}

// decodeRecord aplica DecodeStrict ao registro depois de remover as chaves
// opacas.
//
// A remocao e o que separa sinal de ruido: um campo com nome novo e mudanca no
// jogo e deve alarmar; um hash novo so significa que o dicionario do
// CommunityDragon avancou, o que nao diz nada sobre o jogo.
func decodeRecord(alias string, raw json.RawMessage) (*Record, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		return nil, fmt.Errorf("CharacterRecord de %s: %w", alias, err)
	}
	for k := range fields {
		if opaqueKey.MatchString(k) {
			delete(fields, k)
		}
	}
	clean, err := json.Marshal(fields)
	if err != nil {
		return nil, err
	}

	rec, err := model.DecodeStrict[Record]("CharacterRecord de "+alias, clean)
	if err != nil {
		return nil, err
	}
	return &rec, nil
}

// Object devolve uma chave de topo do dump, seja ela um caminho ou um
// identificador opaco.
func (d *Dump) Object(key string) (json.RawMessage, bool) {
	raw, ok := d.objects[key]
	return raw, ok
}

// meleeRangeThreshold separa corpo a corpo de a distancia pelo alcance de
// ataque, e nao pelos papeis declarados: Rammus e "tank" e Sona e "support",
// mas o que decide se uma runa rende o dobro e a distancia do ataque.
const meleeRangeThreshold = 300.0

// Melee informa se o campeao ataca corpo a corpo.
func (r *Record) Melee() bool {
	return r.AttackRange.BaseValue < meleeRangeThreshold
}

// ShortName devolve o nome do personagem sem prefixo de modo, se houver.
func (r *Record) ShortName() string {
	if i := strings.Index(r.CharacterName, "_"); i >= 0 {
		return r.CharacterName[i+1:]
	}
	return r.CharacterName
}
