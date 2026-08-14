// Package mapdata le o dump do mapa, que e a unica fonte que diz quais itens a
// loja de um modo referencia.
//
// items.json nao tem campo de mapa, e InStore nao substitui: e verdadeiro em
// 696 dos 868 itens do 16.16, incluindo os de ARAM e Arena, e ao mesmo tempo e
// falso em itens que a loja do Summoner's Rift referencia. Quem separa e este
// arquivo.
//
// Nao ha DecodeStrict aqui, pelo mesmo motivo do dump de campeao: as chaves de
// topo sao 1657 e incluem identificadores opacos entre chaves, no formato
// {413f2f94}. A vigilancia desta fonte e a contagem minima do sync.
package mapdata

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// modeRecord e o registro de um modo dentro do mapa. So os campos usados.
type modeRecord struct {
	ModeName  string   `json:"mModeName"`
	ItemLists []string `json:"itemLists"`
}

// itemListRecord e uma das listas de item que o modo referencia.
type itemListRecord struct {
	Type  string   `json:"__type"`
	Items []string `json:"mItems"`
}

// ModePath monta a chave do registro de modo dentro do dump.
//
// O formato e posicional e vem do proprio caminho do arquivo no jogo, entao
// deriva-lo e mais seguro do que procurar por sufixo: ha chaves de outros mapas
// no mesmo arquivo.
func ModePath(mapID int32, gameMode string) string {
	return fmt.Sprintf("Maps/Shipping/Map%d/Modes/%s", mapID, strings.ToUpper(gameMode))
}

// ShopItemIDs devolve os ids de item que a loja do modo referencia, em ordem
// crescente e sem repeticao.
//
// Junta TODAS as listas que o modo declara, e nao apenas a primeira. As listas
// sao grupos semanticos — o catalogo principal, os trinkets, os itens de
// Doran, os companheiros de selva, os itens de suporte, as botas de nivel 3 —
// e ficar so com a maior deixaria de fora, entre outros, todas as botas
// avancadas e os itens iniciais.
//
// O que sai daqui e o que a FONTE referencia, nao o que e comprável: parte
// dessas entradas sao buffs de torre e marcadores internos. Separar as duas
// coisas depende de semantica de item e e trabalho do modelo canonico, nao
// deste pacote.
func ShopItemIDs(raw []byte, mapID int32, gameMode string) ([]int32, error) {
	var top map[string]json.RawMessage
	if err := json.Unmarshal(raw, &top); err != nil {
		return nil, fmt.Errorf("parseando dump do mapa %d: %w", mapID, err)
	}

	key := ModePath(mapID, gameMode)
	rawMode, ok := top[key]
	if !ok {
		return nil, fmt.Errorf(
			"dump do mapa %d nao tem o registro %q — o modo saiu do mapa ou a Riot "+
				"mudou o caminho (veja mode.map_id e mode.game_mode em config.json)",
			mapID, key)
	}

	var mode modeRecord
	if err := json.Unmarshal(rawMode, &mode); err != nil {
		return nil, fmt.Errorf("parseando registro %q: %w", key, err)
	}
	if len(mode.ItemLists) == 0 {
		return nil, fmt.Errorf("registro %q nao declara itemLists", key)
	}

	seen := map[int32]bool{}
	var out []int32
	for _, ref := range mode.ItemLists {
		rawList, ok := top[ref]
		if !ok {
			return nil, fmt.Errorf("%q referencia a lista %q, que nao esta no dump", key, ref)
		}
		var list itemListRecord
		if err := json.Unmarshal(rawList, &list); err != nil {
			return nil, fmt.Errorf("parseando lista %q: %w", ref, err)
		}
		if list.Type != "GameModeItemList" {
			return nil, fmt.Errorf("lista %q tem tipo %q, esperado GameModeItemList", ref, list.Type)
		}
		for _, entry := range list.Items {
			id, err := parseItemRef(entry)
			if err != nil {
				return nil, fmt.Errorf("na lista %q: %w", ref, err)
			}
			if !seen[id] {
				seen[id] = true
				out = append(out, id)
			}
		}
	}

	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out, nil
}

// parseItemRef converte "Items/3181" em 3181.
func parseItemRef(ref string) (int32, error) {
	const prefix = "Items/"
	if !strings.HasPrefix(ref, prefix) {
		return 0, fmt.Errorf("referencia de item em formato inesperado: %q", ref)
	}
	n, err := strconv.ParseInt(ref[len(prefix):], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("referencia de item com id nao numerico: %q", ref)
	}
	return int32(n), nil
}
