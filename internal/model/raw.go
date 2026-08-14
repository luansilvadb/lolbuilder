package model

// As structs abaixo descrevem as fontes do plugin rcp-be-lol-game-data. Todas
// passam por DecodeStrict, entao o conjunto de campos aqui precisa ser o
// conjunto completo do arquivo — nao apenas o que o projeto le.
//
// Contagens medidas no patch 16.16, com os campos uniformes em 100% das
// entradas: items.json 868 entradas com 19 campos, perks.json 103 com 10,
// perkstyles.json 5 estilos, champion-summary.json 237 com 7,
// summoner-spells.json 39 com 7, champions/<id>.json 173 arquivos com 20.

// ---------- items.json ----------

// Item e uma entrada do catalogo de itens.
//
// O catalogo e global e mistura o LoL moderno com o modo Jade, separados apenas
// pela faixa de ID. Nao ha campo de mapa: o que a loja do modo vende sai do
// dump do mapa, nao daqui. Ver filter.ItemsInRange e filter.ShopItemIDs.
type Item struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// Active nao quer dizer "compravel". InStore tambem nao: e verdadeiro em
	// 696 dos 868 itens do 16.16, incluindo os de ARAM e Arena.
	Active            bool     `json:"active"`
	InStore           bool     `json:"inStore"`
	From              []int32  `json:"from"`
	To                []int32  `json:"to"`
	Categories        []string `json:"categories"`
	MaxStacks         int32    `json:"maxStacks"`
	RequiredChampion  string   `json:"requiredChampion"`
	RequiredAlly      string   `json:"requiredAlly"`
	RequiredBuffName  string   `json:"requiredBuffCurrencyName"`
	RequiredBuffCost  int32    `json:"requiredBuffCurrencyCost"`
	SpecialRecipe     int32    `json:"specialRecipe"`
	IsEnchantment     bool     `json:"isEnchantment"`
	Price             int32    `json:"price"`
	PriceTotal        int32    `json:"priceTotal"`
	DisplayInItemSets bool     `json:"displayInItemSets"`
	IconPath          string   `json:"iconPath"`
}

// ---------- perks.json ----------

// Rune e uma runa ou um fragmento de stat do sistema Runes Reforged.
//
// Os fragmentos vivem no mesmo arquivo que as runas, distinguidos so pela faixa
// de ID (5000-5999 no 16.16) e pelo slot em que aparecem em perkstyles.json.
type Rune struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`

	// MajorChangePatchVersion e o patch em que a Riot mudou a runa de forma
	// relevante. E o gatilho nativo de revisao da curadoria — mais confiavel
	// que comparar LongDesc, que muda por reescrita de texto sem mudar numero.
	MajorChangePatchVersion string `json:"majorChangePatchVersion"`

	// Os numeros da runa so existem dentro destes textos: o plugin nao publica
	// bloco estruturado de stats para runa. Dai a curadoria a mao.
	Tooltip  string `json:"tooltip"`
	ShortDoc string `json:"shortDesc"`
	LongDoc  string `json:"longDesc"`

	RecommendationDescriptor string   `json:"recommendationDescriptor"`
	IconPath                 string   `json:"iconPath"`
	EndOfGameStatDescs       []string `json:"endOfGameStatDescs"`

	// ignorado: mapa de chaves variaveis com dicas de recomendacao para a tela
	// de selecao. Nao descreve efeito nem numero.
	RecommendationDescriptorAttributes Ignored `json:"recommendationDescriptorAttributes"`
}

// ---------- perkstyles.json ----------

// PerkStyles e o arquivo inteiro: um envelope com versao de schema.
type PerkStyles struct {
	SchemaVersion int         `json:"schemaVersion"`
	Styles        []PerkStyle `json:"styles"`
}

// PerkStyle e um dos cinco caminhos de runa.
//
// Esta struct e a razao pela qual o projeto nao precisa curar a estrutura da
// pagina a mao, ao contrario do que o Classicorone fazia com as maestrias: a
// Riot publica os slots, quantas opcoes cada um tem e quais estilos podem ser
// secundarios. O otimizador enumera a partir daqui.
type PerkStyle struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Tooltip string `json:"tooltip"`

	IsAdvanced       bool    `json:"isAdvanced"`
	AllowedSubStyles []int32 `json:"allowedSubStyles"`
	Slots            []Slot  `json:"slots"`

	SubStyleBonus              []SubStyleBonus `json:"subStyleBonus"`
	DefaultStatModsPerSubStyle []StatModPreset `json:"defaultStatModsPerSubStyle"`

	DefaultPageName          string  `json:"defaultPageName"`
	DefaultSubStyle          int32   `json:"defaultSubStyle"`
	DefaultPerks             []int32 `json:"defaultPerks"`
	DefaultPerksWhenSplashed []int32 `json:"defaultPerksWhenSplashed"`

	IconPath string `json:"iconPath"`

	// ignorado: mapa de chaves variaveis no formato p<estilo>_s<sub>_k<keystone>
	// apontando imagens de fundo. Sao 22 a 27 chaves por estilo, todas de
	// asset.
	AssetMap Ignored `json:"assetMap"`
}

// Slot e uma linha do caminho. Type distingue a keystone dos menores e dos
// fragmentos de stat: kKeyStone, kMixedRegularSplashable, kStatMod.
type Slot struct {
	Type      string  `json:"type"`
	SlotLabel string  `json:"slotLabel"`
	Perks     []int32 `json:"perks"`
}

// SubStyleBonus e a runa concedida por escolher determinado estilo secundario.
type SubStyleBonus struct {
	StyleID int32 `json:"styleId"`
	PerkID  int32 `json:"perkId"`
}

// StatModPreset sao os fragmentos sugeridos para uma combinacao de estilos.
// O Id vem como string na fonte, ainda que seja um identificador de estilo.
type StatModPreset struct {
	ID    string  `json:"id"`
	Perks []int32 `json:"perks"`
}

// ---------- champion-summary.json ----------

// ChampionSummary e a entrada resumida de um campeao.
//
// O arquivo traz 237 entradas no 16.16: 173 do LoL moderno, 63 do modo Jade e
// uma sentinela de id -1 chamada "None". A faixa de ID separa as tres.
type ChampionSummary struct {
	ID                 int32    `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Alias              string   `json:"alias"`
	ContentID          string   `json:"contentId"`
	SquarePortraitPath string   `json:"squarePortraitPath"`
	Roles              []string `json:"roles"`
}

// ---------- summoner-spells.json ----------

// SummonerSpell e um feitico de invocador.
//
// Diferente dos itens, esta fonte traz o recorte de modo embutido em GameModes
// — 9 dos 39 feiticos valem em CLASSIC no 16.16.
type SummonerSpell struct {
	// ID e int64, e nao int32 como nas outras entidades, porque a fonte usa
	// 4294967295 (o maximo de um uint32) como sentinela de "sem id". No 16.16
	// sao tres entradas chamadas "Primal Smite", todas com GameModes vazio —
	// o recorte por modo ja as descarta, mas a decodificacao estoura antes se o
	// campo nao couber.
	ID            int64    `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	SummonerLevel int32    `json:"summonerLevel"`
	Cooldown      int32    `json:"cooldown"`
	GameModes     []string `json:"gameModes"`
	IconPath      string   `json:"iconPath"`
}

// ---------- champions/<id>.json ----------

// Champion e o detalhe de um campeao: texto das habilidades, papeis e
// classificacao.
//
// Os numeros de dano NAO vem daqui. O plugin publica coeficientes e
// effectAmounts largamente zerados; a fonte com estatistica base e formula de
// habilidade e o dump de dados do jogo, em game/data/characters.
type Champion struct {
	ID              int32           `json:"id"`
	ContentID       string          `json:"contentId"`
	Name            string          `json:"name"`
	Alias           string          `json:"alias"`
	Title           string          `json:"title"`
	ShortBio        string          `json:"shortBio"`
	IsVisible       bool            `json:"isVisibleInClient"`
	Roles           []string        `json:"roles"`
	TacticalInfo    TacticalInfo    `json:"tacticalInfo"`
	PlaystyleInfo   PlaystyleInfo   `json:"playstyleInfo"`
	ChampionTagInfo ChampionTagInfo `json:"championTagInfo"`

	Passive Passive         `json:"passive"`
	Spells  []ChampionSpell `json:"spells"`

	// SpellbookOverride existe em um unico campeao no 16.16 (Hwei), que troca o
	// livro de feiticos inteiro. Fica mapeado, e nao ignorado, porque e dado de
	// habilidade de verdade.
	SpellbookOverride [][]ChampionSpell `json:"spellbookOverride,omitempty"`

	SquarePortraitPath string `json:"squarePortraitPath"`
	StingerSfxPath     string `json:"stingerSfxPath"`
	ChooseVoPath       string `json:"chooseVoPath"`
	BanVoPath          string `json:"banVoPath"`

	// ignorado: catalogo de skins e chromas. E a maior parte do arquivo, muda a
	// cada patch de cosmetico, e o dataset nao fala de cosmetico.
	Skins Ignored `json:"skins"`
	// ignorado: itens recomendados pela Riot na tela de loja. Vazio em todos os
	// 173 campeoes do 16.16, e o projeto calcula os proprios maximos por ouro.
	RecommendedItemDefaults Ignored `json:"recommendedItemDefaults"`
}

// TacticalInfo e a classificacao de combate que o cliente exibe.
type TacticalInfo struct {
	Style      int32  `json:"style"`
	Difficulty int32  `json:"difficulty"`
	DamageType string `json:"damageType"`
	// AttackType distingue melee de ranged, que muda o efeito de varias runas.
	AttackType string `json:"attackType"`
}

// PlaystyleInfo sao as notas de 1 a 3 que o cliente mostra no seletor.
type PlaystyleInfo struct {
	Damage       int32 `json:"damage"`
	Durability   int32 `json:"durability"`
	CrowdControl int32 `json:"crowdControl"`
	Mobility     int32 `json:"mobility"`
	Utility      int32 `json:"utility"`
}

// ChampionTagInfo sao as duas etiquetas de arquetipo.
type ChampionTagInfo struct {
	Primary   string `json:"championTagPrimary"`
	Secondary string `json:"championTagSecondary"`
}

// Passive e a passiva do campeao, so em texto.
type Passive struct {
	Name                  string `json:"name"`
	Description           string `json:"description"`
	AbilityIconPath       string `json:"abilityIconPath"`
	AbilityVideoPath      string `json:"abilityVideoPath"`
	AbilityVideoImagePath string `json:"abilityVideoImagePath"`
}

// ChampionSpell e uma habilidade de slot como o plugin a publica.
//
// Os campos numericos aqui sao rasos de proposito na fonte: Coefficients traz
// dois coeficientes genericos e EffectAmounts vem quase sempre zerado. Servem
// para conferir indexacao de rank contra o dump, nao para publicar dano.
type ChampionSpell struct {
	SpellKey           string `json:"spellKey"`
	Name               string `json:"name"`
	Cost               string `json:"cost"`
	Cooldown           string `json:"cooldown"`
	Description        string `json:"description"`
	DynamicDescription string `json:"dynamicDescription"`

	Range                []float64 `json:"range"`
	CostCoefficients     []float64 `json:"costCoefficients"`
	CooldownCoefficients []float64 `json:"cooldownCoefficients"`
	MaxLevel             int32     `json:"maxLevel"`

	Coefficients  Coefficients         `json:"coefficients"`
	EffectAmounts map[string][]float64 `json:"effectAmounts"`
	Ammo          Ammo                 `json:"ammo"`

	AbilityIconPath       string `json:"abilityIconPath"`
	AbilityVideoPath      string `json:"abilityVideoPath"`
	AbilityVideoImagePath string `json:"abilityVideoImagePath"`
}

// Coefficients sao os dois coeficientes genericos do plugin.
type Coefficients struct {
	Coefficient1 float64 `json:"coefficient1"`
	Coefficient2 float64 `json:"coefficient2"`
}

// Ammo descreve habilidades com cargas.
type Ammo struct {
	AmmoRechargeTime []float64 `json:"ammoRechargeTime"`
	MaxAmmo          []int32   `json:"maxAmmo"`
}
