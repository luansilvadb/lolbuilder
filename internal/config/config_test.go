package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// configValido e o menor config que passa. Os testes partem dele e estragam um
// campo por vez, para que a mensagem de erro seja a coisa testada.
const configValido = `{
  "base_url": "https://exemplo",
  "patchline": "latest",
  "plugin_path": "plugins/p",
  "game_data_path": "game/data/characters",
  "map_data_path": "game/data/maps/shipping",
  "item_data_path": "game/items.cdtb.bin.json",
  "locale_canonical": "default",
  "locale_display": "pt_br",
  "mode": {
    "name": "classic", "game_mode": "CLASSIC", "map_id": 11,
    "item_id_min": 1000, "item_id_max": 699999,
    "champion_id_min": 1, "champion_id_max": 999
  },
  "minimums": {
    "item_catalog": 500, "item_shop": 200, "runes": 80,
    "perk_styles": 5, "champions": 150, "summoner_spells": 7,
    "item_groups": 16
  },
  "provisional": true,
  "coverage_minimums": {
    "champion_stats": 0, "abilities": 0, "passives": 0,
    "abilities_without_formula_max": 0, "passives_without_formula_max": 0,
    "rank_alignment": 0
  },
  "ingame_tolerance": 0.02,
  "token_budget_max": 0,
  "snapshots_dir": "snapshots",
  "http_timeout_seconds": 60
}`

func escrever(t *testing.T, conteudo string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(p, []byte(conteudo), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestLoadAceitaConfigProvisorio(t *testing.T) {
	cfg, err := Load(escrever(t, configValido))
	if err != nil {
		t.Fatalf("config valido rejeitado: %v", err)
	}
	if !cfg.Provisional {
		t.Fatal("provisional nao foi lido")
	}
	if cfg.Mode.MapID != 11 || cfg.Mode.GameMode != "CLASSIC" {
		t.Fatalf("modo lido errado: %+v", cfg.Mode)
	}
}

// TestCoberturaZeradaSoPassaProvisoria e a trava da decisao de bootstrap: sem
// medicao nao ha publicacao, e a unica forma de rodar sem os minimos e
// declarar isso explicitamente no arquivo.
func TestCoberturaZeradaSoPassaProvisoria(t *testing.T) {
	naoProvisorio := strings.Replace(configValido, `"provisional": true`, `"provisional": false`, 1)
	_, err := Load(escrever(t, naoProvisorio))
	if err == nil {
		t.Fatal("cobertura zerada passou sem a flag de provisorio")
	}
	if !strings.Contains(err.Error(), "provisional") {
		t.Fatalf("o erro nao diz como sair da situacao: %v", err)
	}
}

func TestCoberturaForaDaFaixaFalhaMesmoProvisoria(t *testing.T) {
	ruim := strings.Replace(configValido, `"abilities": 0`, `"abilities": 140`, 1)
	_, err := Load(escrever(t, ruim))
	if err == nil {
		t.Fatal("140% de cobertura passou escondido atras da flag de provisorio")
	}
}

func TestFaixasInvalidas(t *testing.T) {
	casos := map[string]string{
		"faixa de item invertida":  strings.Replace(configValido, `"item_id_max": 699999`, `"item_id_max": 10`, 1),
		"faixa de campeao zerada":  strings.Replace(configValido, `"champion_id_min": 1`, `"champion_id_min": 0`, 1),
		"map_id ausente":           strings.Replace(configValido, `"map_id": 11`, `"map_id": 0`, 1),
		"minimo de runas zerado":   strings.Replace(configValido, `"runes": 80`, `"runes": 0`, 1),
		"locale de exibicao vazio": strings.Replace(configValido, `"locale_display": "pt_br"`, `"locale_display": ""`, 1),
		"caminho do mapa ausente":  strings.Replace(configValido, `"map_data_path": "game/data/maps/shipping"`, `"map_data_path": ""`, 1),
		"timeout de http invalido": strings.Replace(configValido, `"http_timeout_seconds": 60`, `"http_timeout_seconds": 0`, 1),
	}
	for nome, conteudo := range casos {
		t.Run(nome, func(t *testing.T) {
			if _, err := Load(escrever(t, conteudo)); err == nil {
				t.Fatal("config invalido foi aceito")
			}
		})
	}
}

func TestURLs(t *testing.T) {
	cfg, err := Load(escrever(t, configValido))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"DataURL":     "https://exemplo/latest/plugins/p/global/default/v1/items.json",
		"DisplayURL":  "https://exemplo/latest/plugins/p/global/pt_br/v1/items.json",
		"GameDataURL": "https://exemplo/latest/game/data/characters/garen/garen.bin.json",
		"MapDataURL":  "https://exemplo/latest/game/data/maps/shipping/map11/map11.bin.json",
		"MetadataURL": "https://exemplo/latest/content-metadata.json",
		"StatusURL":   "https://exemplo/json/status.live.txt",
	}
	got := map[string]string{
		"DataURL":     cfg.DataURL("items.json"),
		"DisplayURL":  cfg.DisplayURL("items.json"),
		"GameDataURL": cfg.GameDataURL(GameDataFile("Garen")),
		"MapDataURL":  cfg.MapDataURL(),
		"MetadataURL": cfg.MetadataURL(),
		"StatusURL":   cfg.StatusURL(),
	}
	for k, w := range want {
		if got[k] != w {
			t.Errorf("%s = %q, esperado %q", k, got[k], w)
		}
	}
}

// TestGameDataFileDerivaDoAlias protege o mecanismo que faz campeao novo entrar
// no sync sozinho: o caminho do dump sai do alias, e nao de uma tabela.
func TestGameDataFileDerivaDoAlias(t *testing.T) {
	casos := map[string]string{
		"Garen":      "garen/garen.bin.json",
		"MonkeyKing": "monkeyking/monkeyking.bin.json",
		"Aphelios":   "aphelios/aphelios.bin.json",
		"KSante":     "ksante/ksante.bin.json",
		"Nunu":       "nunu/nunu.bin.json",
	}
	for alias, want := range casos {
		if got := GameDataFile(alias); got != want {
			t.Errorf("GameDataFile(%q) = %q, esperado %q", alias, got, want)
		}
	}
}

func TestMapDataFile(t *testing.T) {
	if got := MapDataFile(11); got != "map11/map11.bin.json" {
		t.Fatalf("MapDataFile(11) = %q", got)
	}
}
