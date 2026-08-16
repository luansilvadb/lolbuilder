package canonical

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/filter"
	"github.com/luansilvadb/lolbuilder/internal/itemgroups"
	"github.com/luansilvadb/lolbuilder/internal/mapdata"
	"github.com/luansilvadb/lolbuilder/internal/model"
)

// Builder monta o modelo canonico a partir de um snapshot ja capturado.
//
// Le apenas do disco. O build nunca vai a rede: o snapshot e a fonte, e um
// build que baixasse alguma coisa deixaria de ser reproduzivel.
type Builder struct {
	cfg         *config.Config
	patchDir    string
	curationDir string
}

// NewBuilder aponta para o diretorio de um snapshot e para o da curadoria.
func NewBuilder(cfg *config.Config, patchDir, curationDir string) *Builder {
	return &Builder{cfg: cfg, patchDir: patchDir, curationDir: curationDir}
}

// ler carrega um arquivo bruto do snapshot.
func (b *Builder) ler(nome string) ([]byte, error) {
	raw, err := os.ReadFile(filepath.Join(b.patchDir, filepath.FromSlash(nome)))
	if err != nil {
		return nil, fmt.Errorf("lendo %s do snapshot: %w", nome, err)
	}
	return raw, nil
}

// decodificar le e decodifica estritamente um arquivo do snapshot.
func decodificar[T any](b *Builder, nome string) (T, error) {
	var zero T
	raw, err := b.ler(nome)
	if err != nil {
		return zero, err
	}
	return model.DecodeStrict[T](nome, raw)
}

// display devolve o caminho do mesmo arquivo no locale de exibicao.
func (b *Builder) display(nome string) string {
	return b.cfg.LocaleDisplay + "/" + nome
}

// Build monta o dataset do patch.
func (b *Builder) Build(patch string) (*Dataset, error) {
	ds := &Dataset{Patch: patch}

	if err := b.buildItems(ds); err != nil {
		return nil, err
	}
	if err := b.buildRunes(ds); err != nil {
		return nil, err
	}
	if err := b.buildChampions(ds); err != nil {
		return nil, err
	}
	if err := b.buildSpells(ds); err != nil {
		return nil, err
	}
	// Por ultimo: o dump de dados do jogo depende dos campeoes ja montados, e e
	// dele que vem estatistica base e formula de habilidade.
	if err := b.buildChampStats(ds); err != nil {
		return nil, err
	}
	// Antes do pre-calculo: o otimizador de build depende dos grupos para nao
	// propor combinacao que a loja recusa.
	if err := b.buildGruposDeItem(ds); err != nil {
		return nil, err
	}
	// Por ultimo: o pre-calculo consome o catalogo ja montado.
	if err := b.buildComputed(ds); err != nil {
		return nil, err
	}
	return ds, nil
}

// buildGruposDeItem carrega os limites de posse do dump de itens do jogo.
//
// Recorta ao conjunto COMPRAVEL, e nao ao catalogo: o dump mistura os modos, e o
// grupo Botas chega com 30 itens dos quais 9 sao variantes de Arena que a loja
// do Summoner's Rift nunca oferece. Um grupo cujos membros compraveis cabem no
// limite nao restringe nada e sai fora — publicar limite sem consequencia ensina
// uma regra que o jogo nao aplica, que e o mesmo defeito ao contrario.
func (b *Builder) buildGruposDeItem(ds *Dataset) error {
	raw, err := b.ler("items.bin.json")
	if err != nil {
		return err
	}
	todos, err := itemgroups.Ler(raw)
	if err != nil {
		return err
	}

	var compraveis []int32
	for _, it := range ds.Items {
		if it.Compravel {
			compraveis = append(compraveis, it.ID)
		}
	}
	restritos := itemgroups.Restringir(todos, compraveis)

	ds.GruposDeItem = make([]GrupoDeItem, 0, len(restritos))
	for _, g := range restritos {
		ds.GruposDeItem = append(ds.GruposDeItem,
			GrupoDeItem{ID: g.ID, Maximo: g.Maximo, Itens: g.Itens})
	}

	porItem := itemgroups.PorItem(restritos)
	for i := range ds.Items {
		ds.Items[i].Grupos = porItem[ds.Items[i].ID]
	}
	return nil
}

// buildItems monta o catalogo e decide o que e compravel.
func (b *Builder) buildItems(ds *Dataset) error {
	canonicos, err := decodificar[[]model.Item](b, "items.json")
	if err != nil {
		return err
	}
	exibicao, err := decodificar[[]model.Item](b, b.display("items.json"))
	if err != nil {
		return err
	}
	porID := filter.ItemsByID(exibicao)

	mapa, err := b.ler(fmt.Sprintf("map%d.bin.json", b.cfg.Mode.MapID))
	if err != nil {
		return err
	}
	referenciados, err := mapdata.ShopItemIDs(mapa, b.cfg.Mode.MapID, b.cfg.Mode.GameMode)
	if err != nil {
		return err
	}
	naLoja := make(map[int32]bool, len(referenciados))
	for _, id := range referenciados {
		naLoja[id] = true
	}

	doModo := filter.ItemsInRange(canonicos, b.cfg.Mode)
	ds.Items = make([]Item, 0, len(doModo))

	for _, it := range doModo {
		pt, ok := porID[it.ID]
		if !ok {
			return fmt.Errorf("item %d (%s) nao existe no locale de exibicao — "+
				"o snapshot esta inconsistente", it.ID, it.Name)
		}

		// A conjuncao E a decisao. Ver o comentario de Item.Compravel.
		compravel := naLoja[it.ID] && it.InStore

		item := Item{
			ID:           it.ID,
			Nome:         pt.Name,
			NomeCanonico: it.Name,
			Custo:        it.PriceTotal,
			Combina:      it.Price,
			Componentes:  it.From,
			Efeito:       TextoDeEfeito(pt.Description),
			Compravel:    compravel,

			categoriaBotas: temCategoria(it.Categories, categoriaBotas),
		}

		// Os stats saem do locale CANONICO, e nao do traduzido: o vocabulario de
		// rotulos e em ingles, e ler o traduzido exigiria uma segunda tabela de
		// rotulos que so poderia divergir da primeira.
		leitura := canon.LerStatsDeItem(it.Name, it.Description)
		if len(leitura.Stats) > 0 {
			if err := leitura.Stats.Validate(); err != nil {
				return fmt.Errorf("item %d (%s): %w", it.ID, it.Name, err)
			}
			item.Stats = leitura.Stats
		}

		// A taxa mede o que o dataset PUBLICA: item fora da loja nao entra no
		// arquivo, e uma linha ilegivel nele nao degrada nada que o consumidor
		// va ler. Mas a linha ilegivel e registrada assim mesmo, num segundo
		// numero, porque ela e o aviso antecipado de forma nova na fonte.
		if compravel {
			ds.Coverage.Itens.Acumular(it.Name, leitura)
		} else {
			ds.Coverage.VocabularioForaDaLoja = append(ds.Coverage.VocabularioForaDaLoja, leitura.NaoLidas...)
		}

		ds.Items = append(ds.Items, item)
	}

	sort.Slice(ds.Items, func(i, j int) bool { return ds.Items[i].ID < ds.Items[j].ID })
	sort.Strings(ds.Coverage.Itens.SemBlocoNome)
	marcarBotas(ds.Items)
	return nil
}

// categoriaBotas e como a fonte marca calcado no campo categories.
const categoriaBotas = "Boots"

// marcarBotas identifica calcado percorrendo a arvore de componentes.
//
// A categoria direta NAO basta: Gunmetal Greaves evolui de Berserker's Greaves e
// a fonte a categoriza como AttackSpeed, LifeSteal e NonbootsMovement, sem
// Boots. Sem a transitividade o otimizador montaria uma build com dois pares de
// botas — impossivel no jogo, e com cara de otima.
//
// Aprimoramento de botas e botas: e a unica leitura que sobrevive a fonte
// esquecer a etiqueta em mais um item no proximo patch.
func marcarBotas(itens []Item) {
	porID := make(map[int32]*Item, len(itens))
	for i := range itens {
		porID[itens[i].ID] = &itens[i]
	}

	memo := map[int32]bool{}
	var ehBotas func(id int32, visitando map[int32]bool) bool
	ehBotas = func(id int32, visitando map[int32]bool) bool {
		if v, ok := memo[id]; ok {
			return v
		}
		it, ok := porID[id]
		if !ok || visitando[id] {
			return false
		}
		visitando[id] = true
		defer delete(visitando, id)

		r := it.categoriaBotas
		if !r {
			for _, c := range it.Componentes {
				if ehBotas(c, visitando) {
					r = true
					break
				}
			}
		}
		memo[id] = r
		return r
	}

	for i := range itens {
		itens[i].Botas = ehBotas(itens[i].ID, map[int32]bool{})
	}
}

func temCategoria(cats []string, alvo string) bool {
	for _, c := range cats {
		if c == alvo {
			return true
		}
	}
	return false
}

// buildRunes monta o catalogo de runas e os cinco estilos.
func (b *Builder) buildRunes(ds *Dataset) error {
	canonicas, err := decodificar[[]model.Rune](b, "perks.json")
	if err != nil {
		return err
	}
	exibicao, err := decodificar[[]model.Rune](b, b.display("perks.json"))
	if err != nil {
		return err
	}
	estilos, err := decodificar[model.PerkStyles](b, "perkstyles.json")
	if err != nil {
		return err
	}
	estilosPT, err := decodificar[model.PerkStyles](b, b.display("perkstyles.json"))
	if err != nil {
		return err
	}

	porID := make(map[int32]model.Rune, len(exibicao))
	for _, r := range exibicao {
		porID[r.ID] = r
	}

	posicoes := situarRunas(estilos)

	ds.Runes = make([]Rune, 0, len(canonicas))
	for _, r := range canonicas {
		pt, ok := porID[r.ID]
		if !ok {
			return fmt.Errorf("runa %d (%s) nao existe no locale de exibicao — "+
				"o snapshot esta inconsistente", r.ID, r.Name)
		}
		p := posicoes[r.ID]
		ds.Runes = append(ds.Runes, Rune{
			ID:                   r.ID,
			Nome:                 pt.Name,
			NomeCanonico:         r.Name,
			Resumo:               Limpar(pt.ShortDoc),
			Descricao:            Limpar(pt.LongDoc),
			PatchDaUltimaMudanca: r.MajorChangePatchVersion,
			Estilo:               p.estilo,
			TipoSlot:             p.tipo,
			LinhasSlot:           p.linhas,
		})
	}
	sort.Slice(ds.Runes, func(i, j int) bool { return ds.Runes[i].ID < ds.Runes[j].ID })

	nomePT := make(map[int32]string, len(estilosPT.Styles))
	for _, st := range estilosPT.Styles {
		nomePT[st.ID] = st.Name
	}

	ds.RuneStyles = make([]RuneStyle, 0, len(estilos.Styles))
	for _, st := range estilos.Styles {
		linhas := make([]Linha, 0, len(st.Slots))
		for _, sl := range st.Slots {
			linhas = append(linhas, Linha{Tipo: sl.Type, Runas: sl.Perks})
		}
		bonus := make(map[int32]int32, len(st.SubStyleBonus))
		for _, sb := range st.SubStyleBonus {
			bonus[sb.StyleID] = sb.PerkID
		}
		ds.RuneStyles = append(ds.RuneStyles, RuneStyle{
			ID:                   st.ID,
			Nome:                 nomePT[st.ID],
			NomeCanonico:         st.Name,
			SubEstilosPermitidos: st.AllowedSubStyles,
			Linhas:               linhas,
			BonusPorSubEstilo:    bonus,
		})
	}
	sort.Slice(ds.RuneStyles, func(i, j int) bool { return ds.RuneStyles[i].ID < ds.RuneStyles[j].ID })
	return nil
}

// posicaoDeRuna e onde uma runa aparece na pagina.
type posicaoDeRuna struct {
	estilo int32
	tipo   string
	linhas []int
}

// situarRunas descobre onde cada runa aparece, varrendo os estilos.
//
// A posicao nao esta no catalogo: perks.json nao diz a que estilo cada runa
// pertence nem em que linha ela aparece. So perkstyles.json diz.
//
// Fragmento de stat aparece nos cinco estilos e pode ocupar mais de uma linha.
// Por isso as linhas sao um conjunto, e o estilo fica zerado: os 7 fragmentos
// do 16.16 valem em qualquer caminho.
func situarRunas(estilos model.PerkStyles) map[int32]posicaoDeRuna {
	vistas := map[int32]map[int]bool{}
	out := map[int32]posicaoDeRuna{}

	for _, st := range estilos.Styles {
		for i, sl := range st.Slots {
			for _, id := range sl.Perks {
				p, ja := out[id]
				if !ja {
					p = posicaoDeRuna{tipo: sl.Type}
					vistas[id] = map[int]bool{}
					if sl.Type != SlotStatMod {
						p.estilo = st.ID
					}
				}
				if !vistas[id][i] {
					vistas[id][i] = true
					p.linhas = append(p.linhas, i)
				}
				out[id] = p
			}
		}
	}
	for id, p := range out {
		sort.Ints(p.linhas)
		out[id] = p
	}
	return out
}

// buildChampions monta os campeoes do modo, so com o que o plugin publica.
func (b *Builder) buildChampions(ds *Dataset) error {
	resumos, err := decodificar[[]model.ChampionSummary](b, "champion-summary.json")
	if err != nil {
		return err
	}
	doModo := filter.ChampionsInRange(resumos, b.cfg.Mode)
	sort.Slice(doModo, func(i, j int) bool { return doModo[i].ID < doModo[j].ID })

	ds.Champions = make([]Champion, 0, len(doModo))
	for _, cs := range doModo {
		nome := fmt.Sprintf("champions/%d.json", cs.ID)
		det, err := decodificar[model.Champion](b, nome)
		if err != nil {
			return err
		}
		detPT, err := decodificar[model.Champion](b, b.display(nome))
		if err != nil {
			return err
		}

		hab, err := parearHabilidades(nome, det.Spells, detPT.Spells)
		if err != nil {
			return err
		}

		// O livro de feiticos e um segundo conjunto de habilidades, com dano e
		// recarga proprios. So Hwei tem no 16.16: tres grupos de quatro.
		var sub []Habilidade
		for g, grupo := range det.SpellbookOverride {
			var grupoPT []model.ChampionSpell
			if g < len(detPT.SpellbookOverride) {
				grupoPT = detPT.SpellbookOverride[g]
			}
			habs, err := parearHabilidades(nome, grupo, grupoPT)
			if err != nil {
				return err
			}
			sub = append(sub, habs...)
		}

		ds.Champions = append(ds.Champions, Champion{
			ID:           cs.ID,
			Nome:         detPT.Name,
			NomeCanonico: det.Name,
			Alias:        det.Alias,
			Titulo:       detPT.Title,
			Papeis:       det.Roles,
			CorpoACorpo:  det.TacticalInfo.AttackType == "melee",
			TipoDeDano:   det.TacticalInfo.DamageType,
			Passiva: Habilidade{
				Slot:      "p",
				Nome:      detPT.Passive.Name,
				Descricao: Limpar(detPT.Passive.Description),
			},
			Habilidades:    hab,
			SubHabilidades: sub,
		})
		lerSeriesDoPlugin(ds.Champions[len(ds.Champions)-1].Habilidades, det.Spells)
	}
	return nil
}

// parearHabilidades casa a habilidade do locale canonico com a do traduzido.
//
// O pareamento e posicional, porque e assim que a fonte serve os dois arquivos,
// mas o spellKey e CONFERIDO. Sem essa conferencia, uma reordenacao num dos
// locales publicaria a descricao do W sob o nome do Q, em silencio e em todos os
// campeoes de uma vez. A guarda do sync confere a contagem, nao a ordem.
func parearHabilidades(arquivo string, canonicas, traduzidas []model.ChampionSpell) ([]Habilidade, error) {
	if len(canonicas) != len(traduzidas) {
		return nil, fmt.Errorf("%s tem %d habilidade(s) no locale canonico e %d no traduzido",
			arquivo, len(canonicas), len(traduzidas))
	}
	out := make([]Habilidade, 0, len(canonicas))
	for i := range canonicas {
		if canonicas[i].SpellKey != traduzidas[i].SpellKey {
			return nil, fmt.Errorf(
				"%s: a habilidade %d e %q no locale canonico e %q no traduzido — "+
					"os dois locales estao em ordens diferentes, e parear por posicao "+
					"publicaria a descricao de uma habilidade sob o nome de outra",
				arquivo, i, canonicas[i].SpellKey, traduzidas[i].SpellKey)
		}
		out = append(out, Habilidade{
			Slot:      traduzidas[i].SpellKey,
			Nome:      traduzidas[i].Name,
			Descricao: Limpar(traduzidas[i].DynamicDescription),
		})
	}
	return out, nil
}

// buildSpells monta os feiticos validos no modo.
func (b *Builder) buildSpells(ds *Dataset) error {
	canonicos, err := decodificar[[]model.SummonerSpell](b, "summoner-spells.json")
	if err != nil {
		return err
	}
	exibicao, err := decodificar[[]model.SummonerSpell](b, b.display("summoner-spells.json"))
	if err != nil {
		return err
	}
	porID := make(map[int64]model.SummonerSpell, len(exibicao))
	for _, s := range exibicao {
		porID[s.ID] = s
	}

	doModo := filter.SummonerSpellsForMode(canonicos, b.cfg.Mode)
	ds.SummonerSpells = make([]SummonerSpell, 0, len(doModo))
	for _, s := range doModo {
		pt := porID[s.ID]
		ds.SummonerSpells = append(ds.SummonerSpells, SummonerSpell{
			ID:           s.ID,
			Nome:         pt.Name,
			NomeCanonico: s.Name,
			Descricao:    Limpar(pt.Description),
			Recarga:      s.Cooldown,
			NivelMinimo:  s.SummonerLevel,
		})
	}
	sort.Slice(ds.SummonerSpells, func(i, j int) bool {
		return ds.SummonerSpells[i].ID < ds.SummonerSpells[j].ID
	})
	return nil
}
