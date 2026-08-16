// Package itemgroups le os grupos de exclusividade de item no dump do jogo.
//
// O jogo limita quantos itens de um mesmo grupo o jogador pode ter ao mesmo
// tempo. Botas e o caso conhecido, mas ha vinte outros: Terminus e Lembrete
// Mortal estao ambos em LastWhisper, e comprar o segundo trava o primeiro na
// loja.
//
// O catalogo do plugin NAO publica isso em campo nenhum — conferido item a
// item. So o dump do jogo tem, e por isso esta fonte existe.
//
// Sem ela o otimizador calculava o otimo exato sobre um conjunto viavel errado:
// 5 das 24 builds publicadas no 16.16 nao podiam ser compradas, entre elas a de
// maxima penetracao de armadura, que juntava QUATRO itens do mesmo grupo. Um
// otimo que nao existe e pior que nenhum otimo — quem le nao tem como saber que
// aquilo nao e compravel, e a regra que o proibe nao estava em lugar nenhum do
// conjunto.
//
// Nao ha DecodeStrict aqui, pelo mesmo motivo do dump do mapa: as chaves de
// topo passam de 60 mil e incluem identificadores opacos entre chaves, no
// formato {8875daa8}. A vigilancia desta fonte e a contagem minima do sync.
package itemgroups

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// grupoDefault e o grupo a que TODO item pertence, e por isso nao restringe
// nada. Ele declara max 1 como qualquer outro, e trata-lo como limite tornaria
// impossivel comprar dois itens quaisquer.
const grupoDefault = "Items/ItemGroups/Default"

// Grupo e um limite de posse declarado pelo jogo.
type Grupo struct {
	// ID e o nome curto que o jogo usa, ex: "LastWhisper".
	ID string `json:"id"`
	// Caminho e a chave completa no dump, ex: "Items/ItemGroups/LastWhisper".
	Caminho string `json:"caminho"`
	// Maximo e quantos itens do grupo cabem ao mesmo tempo.
	Maximo int `json:"maximo"`
	// Itens sao os ids que pertencem ao grupo, em ordem crescente.
	Itens []int32 `json:"itens"`
}

type registroGrupo struct {
	Tipo   string `json:"__type"`
	ID     string `json:"mItemGroupID"`
	Maximo int    `json:"mMaxGroupOwnable"`
}

type registroItem struct {
	Tipo   string   `json:"__type"`
	ItemID int32    `json:"itemID"`
	Grupos []string `json:"mItemGroups"`
}

// Ler extrai os grupos que de fato limitam alguma coisa.
//
// Grupo com maximo zero ou negativo nao limita, e o grupo Default vale para
// todo item. Os dois ficam de fora: publicar limite que nao limita ensinaria a
// quem le uma restricao que o jogo nao aplica.
func Ler(raw []byte) ([]Grupo, error) {
	var top map[string]json.RawMessage
	if err := json.Unmarshal(raw, &top); err != nil {
		return nil, fmt.Errorf("parseando o dump de itens do jogo: %w", err)
	}

	limites := map[string]registroGrupo{}
	membros := map[string][]int32{}

	for chave, bruto := range top {
		// Uma sondagem so, no campo que separa os dois tipos que interessam.
		// Decodificar cada uma das ~60 mil entradas nos dois formatos custaria
		// caro e nao acrescentaria nada.
		var tipo struct {
			Tipo string `json:"__type"`
		}
		if err := json.Unmarshal(bruto, &tipo); err != nil {
			// Entrada que nao e objeto: o dump mistura tipos no topo.
			continue
		}

		switch tipo.Tipo {
		case "ItemGroup":
			var g registroGrupo
			if err := json.Unmarshal(bruto, &g); err != nil {
				return nil, fmt.Errorf("parseando o grupo %q: %w", chave, err)
			}
			g.ID = strings.TrimSpace(g.ID)
			limites[chave] = g

		case "ItemData":
			var it registroItem
			if err := json.Unmarshal(bruto, &it); err != nil {
				return nil, fmt.Errorf("parseando o item %q: %w", chave, err)
			}
			for _, g := range it.Grupos {
				if g == grupoDefault {
					continue
				}
				membros[g] = append(membros[g], it.ItemID)
			}
		}
	}

	var out []Grupo
	for caminho, g := range limites {
		if caminho == grupoDefault || g.Maximo <= 0 {
			continue
		}
		itens := membros[caminho]
		if len(itens) == 0 {
			// Grupo declarado e vazio existe aos montes no dump — sao modos e
			// temporadas que sairam do jogo. Sem item, nao restringe nada.
			continue
		}
		sort.Slice(itens, func(i, j int) bool { return itens[i] < itens[j] })
		id := g.ID
		if id == "" {
			id = caminho[strings.LastIndex(caminho, "/")+1:]
		}
		out = append(out, Grupo{ID: id, Caminho: caminho, Maximo: g.Maximo, Itens: dedup(itens)})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Caminho < out[j].Caminho })
	return out, nil
}

func dedup(v []int32) []int32 {
	out := v[:0:0]
	for i, x := range v {
		if i == 0 || x != v[i-1] {
			out = append(out, x)
		}
	}
	return out
}

// Restringir recorta os grupos ao catalogo de um modo e descarta os que deixam
// de restringir alguma coisa.
//
// Sao dois cortes, e os dois importam. O dump mistura os modos: o grupo Botas
// tem 31 itens, dos quais 9 sao variantes de Arena e do modo Jade que a loja do
// Summoner's Rift nunca oferece. E a maior parte dos 270 grupos e auto-grupo de
// um item com suas proprias variantes — dentro de um modo so sobra um membro, e
// um grupo de maximo 1 com um membro so nunca e violado.
//
// Publicar esses grupos vazios de consequencia ensinaria uma restricao que nao
// existe, que e o mesmo defeito que este pacote veio corrigir, so que ao
// contrario.
func Restringir(grupos []Grupo, ids []int32) []Grupo {
	no := map[int32]bool{}
	for _, i := range ids {
		no[i] = true
	}

	var out []Grupo
	for _, g := range grupos {
		var itens []int32
		for _, i := range g.Itens {
			if no[i] {
				itens = append(itens, i)
			}
		}
		if len(itens) <= g.Maximo {
			continue
		}
		g.Itens = itens
		out = append(out, g)
	}
	return out
}

// Excedidos devolve os grupos que uma combinacao de itens viola, e por quanto.
//
// Devolve o grupo e a lista de itens dele presentes, e nao so um booleano: quem
// recebe a violacao precisa saber QUAIS itens brigam para decidir qual tirar.
func Excedidos(grupos []Grupo, itens []int32) map[string][]int32 {
	tem := map[int32]bool{}
	for _, i := range itens {
		tem[i] = true
	}

	viol := map[string][]int32{}
	for _, g := range grupos {
		var presentes []int32
		for _, i := range g.Itens {
			if tem[i] {
				presentes = append(presentes, i)
			}
		}
		if len(presentes) > g.Maximo {
			viol[g.ID] = presentes
		}
	}
	return viol
}

// PorItem indexa os grupos de cada item, para consulta durante a otimizacao.
func PorItem(grupos []Grupo) map[int32][]string {
	out := map[int32][]string{}
	for _, g := range grupos {
		for _, i := range g.Itens {
			out[i] = append(out[i], g.ID)
		}
	}
	for _, gs := range out {
		sort.Strings(gs)
	}
	return out
}
