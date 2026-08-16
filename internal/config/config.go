// Package config carrega e valida os parametros de runtime do sync.
//
// Nada especifico do modo fica embutido em codigo: faixas de ID, gameMode e
// mapId vem daqui, para que uma renomeacao pela Riot se resolva editando um
// arquivo.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Mode agrupa tudo que identifica o modo de jogo nas fontes do CDragon.
//
// Nao ha FilePrefix como no modo Jade: os catalogos do LoL moderno sao globais
// e nao carregam prefixo. O recorte usa tres convencoes diferentes, uma por
// entidade, e todas dependem dos campos abaixo — ver o pacote filter.
type Mode struct {
	Name     string `json:"name"`
	GameMode string `json:"game_mode"`
	MapID    int32  `json:"map_id"`

	// As faixas de ID separam o LoL moderno do modo Jade, que a Riot publica
	// no MESMO items.json e no MESMO champion-summary.json. Sem elas o dataset
	// misturaria dois jogos diferentes.
	ItemIDMin     int32 `json:"item_id_min"`
	ItemIDMax     int32 `json:"item_id_max"`
	ChampionIDMin int32 `json:"champion_id_min"`
	ChampionIDMax int32 `json:"champion_id_max"`
}

// Minimums define a contagem minima aceitavel por entidade. Ficar abaixo de
// qualquer uma delas aborta o sync antes de qualquer escrita.
//
// Cada entidade e contada na fonte que alimenta o artefato correspondente:
// Runes vem de perks.json, PerkStyles de perkstyles.json, ItemShop da lista de
// loja do dump do mapa. Sao chaves distintas porque contar uma no lugar da
// outra deixa a outra sem guarda.
type Minimums struct {
	ItemCatalog    int `json:"item_catalog"`
	ItemShop       int `json:"item_shop"`
	Runes          int `json:"runes"`
	PerkStyles     int `json:"perk_styles"`
	Champions      int `json:"champions"`
	SummonerSpells int `json:"summoner_spells"`

	// ItemGroups e o minimo de grupos de exclusividade que restringem a loja do
	// modo. E a unica vigilancia dessa fonte: ela nao passa por DecodeStrict,
	// entao uma renomeacao de campo pela Riot zeraria a extracao em silencio e o
	// otimizador voltaria a publicar build que a loja se recusa a vender.
	ItemGroups int `json:"item_groups"`
}

// CoverageMinimums define a taxa minima de extracao aceitavel, em pontos
// percentuais, para cada eixo do dump de dados do jogo.
//
// Sao eixos separados e nao um numero unico porque uma media diluiria
// justamente o sintoma que o mecanismo precisa detectar: passivas resolvem
// menos que habilidades de slot por natureza da fonte.
//
// Os campos *WithoutFormulaMax sao tetos, e nao pisos: limitam a fracao de
// entidades para as quais a fonte nao publica formula alguma. Existem porque a
// taxa de resolucao sozinha nao detecta esse movimento — uma taxa calibrada
// para caber nas lacunas de hoje cabe tambem no dobro delas.
type CoverageMinimums struct {
	ChampionStats int `json:"champion_stats"`
	Abilities     int `json:"abilities"`
	Passives      int `json:"passives"`

	AbilitiesWithoutFormulaMax int `json:"abilities_without_formula_max"`
	PassivesWithoutFormulaMax  int `json:"passives_without_formula_max"`

	// RankAlignment e a concordancia minima entre o arquivo do plugin e o dump,
	// em pontos percentuais, nas series de recarga e custo.
	RankAlignment int `json:"rank_alignment"`
}

// Config e o conteudo de config.json.
type Config struct {
	BaseURL    string `json:"base_url"`
	Patchline  string `json:"patchline"`
	PluginPath string `json:"plugin_path"`

	// GameDataPath e o prefixo do dump de dados do jogo, que nao fica sob o
	// plugin do cliente. E dele que vem estatistica base e dano de habilidade.
	GameDataPath string `json:"game_data_path"`

	// MapDataPath e o prefixo do dump dos mapas. E a unica fonte que diz quais
	// itens a loja do modo vende: items.json nao tem campo de mapa, e InStore
	// e verdadeiro para itens de ARAM e Arena tambem.
	MapDataPath string `json:"map_data_path"`

	// ItemDataPath e o dump de itens do jogo. E a unica fonte dos grupos de
	// exclusividade: o catalogo do plugin nao publica isso em campo nenhum, e
	// sem eles o otimizador monta build que a loja se recusa a vender.
	ItemDataPath string `json:"item_data_path"`

	LocaleCanonical string `json:"locale_canonical"`
	LocaleDisplay   string `json:"locale_display"`

	Mode             Mode             `json:"mode"`
	Minimums         Minimums         `json:"minimums"`
	CoverageMinimums CoverageMinimums `json:"coverage_minimums"`

	// Provisional afrouxa a validacao de coverage_minimums e token_budget_max,
	// que nascem zerados porque ainda nao houve medicao.
	//
	// O preco de afrouxar e que o export se recusa a publicar enquanto isso for
	// verdadeiro (ver ErrProvisorio). Um dataset cuja qualidade nunca passou por
	// porteira nenhuma e o pior risco do projeto: o consumidor final e um modelo
	// de linguagem, que nao tem como saber que falta dado.
	Provisional bool `json:"provisional"`

	// IngameTolerance e a diferenca aceita entre o que o jogo reporta e o que o
	// dataset preve, em unidades da estatistica.
	//
	// Existe porque as duas leituras nao sao aritmeticamente identicas: o jogo
	// reporta em float32 e o dataset calcula em float64. Diferenca na terceira
	// casa e representacao, nao defeito.
	IngameTolerance float64 `json:"ingame_tolerance"`

	// TokenBudgetMax e o teto do conjunto exportado, em tokens estimados.
	//
	// Fica aqui, e nao cravado num teste, porque e decisao registrada: o
	// comentario ao lado dele no arquivo diz quanto foi medido e por que mudou.
	// Um teto ajustado em silencio toda vez que estoura para de medir qualquer
	// coisa.
	TokenBudgetMax int `json:"token_budget_max"`

	SnapshotsDir       string `json:"snapshots_dir"`
	HTTPTimeoutSeconds int    `json:"http_timeout_seconds"`
}

// Load le e valida o arquivo de configuracao.
func Load(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("lendo %s: %w", path, err)
	}

	var c Config
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("parseando %s: %w", path, err)
	}
	if err := c.validate(); err != nil {
		return nil, fmt.Errorf("config invalida em %s: %w", path, err)
	}
	return &c, nil
}

func (c *Config) validate() error {
	required := []struct {
		name  string
		value string
	}{
		{"base_url", c.BaseURL},
		{"patchline", c.Patchline},
		{"plugin_path", c.PluginPath},
		{"game_data_path", c.GameDataPath},
		{"map_data_path", c.MapDataPath},
		{"item_data_path", c.ItemDataPath},
		{"locale_canonical", c.LocaleCanonical},
		{"locale_display", c.LocaleDisplay},
		{"snapshots_dir", c.SnapshotsDir},
		{"mode.name", c.Mode.Name},
		{"mode.game_mode", c.Mode.GameMode},
	}
	for _, f := range required {
		if f.value == "" {
			return fmt.Errorf("campo obrigatorio ausente: %s", f.name)
		}
	}

	if c.Mode.ItemIDMin <= 0 || c.Mode.ItemIDMax <= c.Mode.ItemIDMin {
		return fmt.Errorf("faixa de id de item invalida: [%d, %d]", c.Mode.ItemIDMin, c.Mode.ItemIDMax)
	}
	if c.Mode.ChampionIDMin <= 0 || c.Mode.ChampionIDMax <= c.Mode.ChampionIDMin {
		return fmt.Errorf("faixa de id de campeao invalida: [%d, %d]", c.Mode.ChampionIDMin, c.Mode.ChampionIDMax)
	}
	if c.Mode.MapID <= 0 {
		return fmt.Errorf("mode.map_id invalido: %d", c.Mode.MapID)
	}

	mins := map[string]int{
		"item_catalog":    c.Minimums.ItemCatalog,
		"item_shop":       c.Minimums.ItemShop,
		"runes":           c.Minimums.Runes,
		"perk_styles":     c.Minimums.PerkStyles,
		"champions":       c.Minimums.Champions,
		"summoner_spells": c.Minimums.SummonerSpells,
		"item_groups":     c.Minimums.ItemGroups,
	}
	for name, v := range mins {
		if v <= 0 {
			return fmt.Errorf("minimums.%s deve ser maior que zero (esta %d)", name, v)
		}
	}

	if err := c.validateCoverage(); err != nil {
		return err
	}

	if c.HTTPTimeoutSeconds <= 0 {
		return fmt.Errorf("http_timeout_seconds deve ser maior que zero")
	}
	return nil
}

// validateCoverage exige cobertura e teto de tokens declarados, exceto enquanto
// a configuracao se declara provisoria.
//
// Provisorio nao e "sem regra": e uma regra diferente, que troca a porteira do
// build pela porteira do export. Ver Provisional.
func (c *Config) validateCoverage() error {
	covs := map[string]int{
		"champion_stats":                c.CoverageMinimums.ChampionStats,
		"abilities":                     c.CoverageMinimums.Abilities,
		"passives":                      c.CoverageMinimums.Passives,
		"abilities_without_formula_max": c.CoverageMinimums.AbilitiesWithoutFormulaMax,
		"passives_without_formula_max":  c.CoverageMinimums.PassivesWithoutFormulaMax,
		"rank_alignment":                c.CoverageMinimums.RankAlignment,
	}

	if c.Provisional {
		// Zero e o valor esperado aqui; qualquer coisa fora da faixa continua
		// sendo erro, para que um numero digitado errado nao passe escondido
		// atras da flag.
		for name, v := range covs {
			if v < 0 || v > 100 {
				return fmt.Errorf("coverage_minimums.%s deve estar entre 0 e 100 (esta %d)", name, v)
			}
		}
		if c.TokenBudgetMax < 0 {
			return fmt.Errorf("token_budget_max nao pode ser negativo (esta %d)", c.TokenBudgetMax)
		}
		return nil
	}

	for name, v := range covs {
		if v <= 0 || v > 100 {
			return fmt.Errorf("coverage_minimums.%s deve estar entre 1 e 100 (esta %d) — "+
				"ou declare provisional=true enquanto nao houver medicao", name, v)
		}
	}
	if c.TokenBudgetMax <= 0 {
		return fmt.Errorf("token_budget_max deve ser maior que zero (esta %d) — "+
			"ou declare provisional=true enquanto nao houver medicao", c.TokenBudgetMax)
	}
	return nil
}

// HTTPTimeout devolve o timeout ja como duracao.
func (c *Config) HTTPTimeout() time.Duration {
	return time.Duration(c.HTTPTimeoutSeconds) * time.Second
}

// DataURL monta a URL de um arquivo de dados no locale canonico.
func (c *Config) DataURL(file string) string {
	return c.localeURL(c.LocaleCanonical, file)
}

// DisplayURL monta a URL do mesmo arquivo no locale de exibicao.
//
// Existe porque o dataset publica nome de item e de runa em pt_br e mantem o
// nome canonico em ingles ao lado: os dois vem do mesmo arquivo em locales
// diferentes, e sao casados pelo id.
func (c *Config) DisplayURL(file string) string {
	return c.localeURL(c.LocaleDisplay, file)
}

func (c *Config) localeURL(locale, file string) string {
	return fmt.Sprintf("%s/%s/%s/global/%s/v1/%s",
		c.BaseURL, c.Patchline, c.PluginPath, locale, file)
}

// GameDataURL monta a URL de um arquivo do dump de dados do jogo.
//
// Note que aqui nao ha locale: o dump nao e traduzido, e o mesmo para todos os
// idiomas.
func (c *Config) GameDataURL(file string) string {
	return fmt.Sprintf("%s/%s/%s/%s", c.BaseURL, c.Patchline, c.GameDataPath, file)
}

// MapDataURL monta a URL do dump do mapa do modo.
func (c *Config) MapDataURL() string {
	return fmt.Sprintf("%s/%s/%s/%s", c.BaseURL, c.Patchline, c.MapDataPath, MapDataFile(c.Mode.MapID))
}

// ItemDataURL e o dump de itens do jogo, de onde saem os grupos de
// exclusividade.
func (c *Config) ItemDataURL() string {
	return fmt.Sprintf("%s/%s/%s", c.BaseURL, c.Patchline, c.ItemDataPath)
}

// StatusURL e o sinal de que o CDragon terminou de processar o patchline.
func (c *Config) StatusURL() string {
	return fmt.Sprintf("%s/json/status.live.txt", c.BaseURL)
}

// MetadataURL e de onde sai a versao do patch servida pelo patchline.
func (c *Config) MetadataURL() string {
	return fmt.Sprintf("%s/%s/content-metadata.json", c.BaseURL, c.Patchline)
}

// MapDataFile deriva o caminho do dump de um mapa (ex: 11 vira
// "map11/map11.bin.json").
func MapDataFile(mapID int32) string {
	slug := fmt.Sprintf("map%d", mapID)
	return slug + "/" + slug + ".bin.json"
}

// GameDataFile deriva o caminho do dump de um campeao a partir do alias que o
// proprio snapshot ja carrega (ex: "MonkeyKing" vira
// "monkeyking/monkeyking.bin.json").
//
// Derivar em vez de manter tabela e o que faz um campeao novo entrar sozinho no
// sync: a Riot adiciona o campeao, o alias aparece em champion-summary.json, e
// o dump e encontrado sem tocar em codigo.
func GameDataFile(alias string) string {
	slug := strings.ToLower(alias)
	return slug + "/" + slug + ".bin.json"
}
