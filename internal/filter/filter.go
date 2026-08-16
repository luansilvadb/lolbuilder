// Package filter isola o modo de jogo dentro das fontes globais do CDragon.
//
// Nao ha um endpoint que devolva "tudo do Summoner's Rift moderno". A
// identificacao usa tres convencoes diferentes conforme a entidade, e todas as
// constantes vem da configuracao — nada de "CLASSIC" ou de faixa de id embutido
// aqui.
package filter

import (
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/model"
)

// Estilo 1 — faixa de ID.
//
// items.json e champion-summary.json sao globais e trazem o LoL moderno junto
// com o modo Jade, que e outro jogo com outros valores. So a faixa de id separa
// os dois.

// ItemsInRange devolve os itens cujo id cai na faixa do modo.
func ItemsInRange(items []model.Item, m config.Mode) []model.Item {
	out := make([]model.Item, 0, 800)
	for _, it := range items {
		if it.ID >= m.ItemIDMin && it.ID <= m.ItemIDMax {
			out = append(out, it)
		}
	}
	return out
}

// ChampionsInRange devolve os campeoes cujo id cai na faixa do modo.
//
// Exclui tambem a sentinela de id -1 chamada "None", que a fonte publica junto
// dos campeoes de verdade.
func ChampionsInRange(cs []model.ChampionSummary, m config.Mode) []model.ChampionSummary {
	out := make([]model.ChampionSummary, 0, 200)
	for _, c := range cs {
		if c.ID >= m.ChampionIDMin && c.ID <= m.ChampionIDMax {
			out = append(out, c)
		}
	}
	return out
}

// Estilo 2 — campo de conteudo.

// SummonerSpellsForMode devolve os feiticos cujo GameModes inclui o modo.
//
// Esta e a unica das tres entidades em que a propria fonte declara o recorte.
func SummonerSpellsForMode(spells []model.SummonerSpell, m config.Mode) []model.SummonerSpell {
	out := make([]model.SummonerSpell, 0, 20)
	for _, s := range spells {
		for _, gm := range s.GameModes {
			if strings.EqualFold(gm, m.GameMode) {
				out = append(out, s)
				break
			}
		}
	}
	return out
}

// Estilo 3 — lista dedicada no dump do mapa. Ver o pacote mapdata.

// ItemsByID indexa um catalogo por id, para casar com a lista de loja.
func ItemsByID(items []model.Item) map[int32]model.Item {
	out := make(map[int32]model.Item, len(items))
	for _, it := range items {
		out[it.ID] = it
	}
	return out
}

// UnknownShopIDs devolve os ids que a loja do modo referencia e que nao existem
// no catalogo COMPLETO.
//
// Deveria ser sempre vazio: a loja referencia itens do mesmo arquivo. Nao
// sendo, a Riot passou a servir a loja de um mapa cujo catalogo nao e este — e
// isso publica item fantasma se passar em silencio.
//
// O catalogo aqui e o completo, e nao o recortado pelo modo, de proposito. Ver
// ShopIDsOutOfRange para o outro caso, que e diferente e nao e defeito.
func UnknownShopIDs(shopIDs []int32, catalog map[int32]model.Item) []int32 {
	var out []int32
	for _, id := range shopIDs {
		if _, ok := catalog[id]; !ok {
			out = append(out, id)
		}
	}
	return out
}

// ShopIDsOutOfRange devolve os ids que a loja do modo referencia e que caem
// fora da faixa do modo.
//
// Nao e defeito, e por isso nao aborta: as listas do dump do mapa misturam o
// catalogo de verdade com buffs de torre, marcadores internos e placeholders, e
// um deles no 16.16 vive na faixa do modo Jade (771500, "Penetrating Bullets").
// Existe para que esse numero apareca no log em vez de sumir — se ele crescer,
// a fonte mudou de forma e alguem precisa olhar.
func ShopIDsOutOfRange(shopIDs []int32, m config.Mode) []int32 {
	var out []int32
	for _, id := range shopIDs {
		if id < m.ItemIDMin || id > m.ItemIDMax {
			out = append(out, id)
		}
	}
	return out
}

// PurchasableIDs devolve os ids que a loja do modo de fato oferece.
//
// A conjuncao e a mesma que o modelo canonico usa em Item.Compravel, e existe
// aqui para que o sync possa contar sobre o mesmo conjunto que o build publica.
// Nenhum dos dois criterios sozinho serve: InStore e verdadeiro para itens de
// ARAM e Arena, e a lista do mapa referencia buffs de torre e marcadores.
func PurchasableIDs(itens []model.Item, referenciados []int32) []int32 {
	naLoja := make(map[int32]bool, len(referenciados))
	for _, id := range referenciados {
		naLoja[id] = true
	}
	var out []int32
	for _, it := range itens {
		if naLoja[it.ID] && it.InStore {
			out = append(out, it.ID)
		}
	}
	return out
}
