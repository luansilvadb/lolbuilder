package canonical

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/filter"
	"github.com/luansilvadb/lolbuilder/internal/mapdata"
	"github.com/luansilvadb/lolbuilder/internal/model"
)

// Builder monta o modelo canonico a partir de um snapshot ja capturado.
//
// Le apenas do disco. O build nunca vai a rede: o snapshot e a fonte, e um
// build que baixasse alguma coisa deixaria de ser reproduzivel.
type Builder struct {
	cfg      *config.Config
	patchDir string
}

// NewBuilder aponta para o diretorio de um snapshot.
func NewBuilder(cfg *config.Config, patchDir string) *Builder {
	return &Builder{cfg: cfg, patchDir: patchDir}
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
	return ds, nil
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

		// A cobertura mede o que o dataset PUBLICA. Item fora da loja nao entra
		// no arquivo, entao uma linha ilegivel nele nao degrada nada que o
		// consumidor va ler.
		if compravel {
			ds.Coverage.Itens.Acumular(it.Name, leitura)
		}

		ds.Items = append(ds.Items, item)
	}

	sort.Slice(ds.Items, func(i, j int) bool { return ds.Items[i].ID < ds.Items[j].ID })
	sort.Strings(ds.Coverage.Itens.SemBlocoNome)
	return nil
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

	// A posicao da runa na pagina vem dos estilos, e nao do catalogo: perks.json
	// nao diz a que estilo cada runa pertence nem em que linha ela aparece.
	type posicao struct {
		estilo int32
		tipo   string
		linha  int
	}
	posicoes := map[int32]posicao{}
	for _, st := range estilos.Styles {
		for i, sl := range st.Slots {
			for _, id := range sl.Perks {
				// Fragmento de stat aparece nos cinco estilos. A primeira
				// atribuicao vence e o estilo fica zerado, porque dizer que o
				// fragmento pertence a Precision so por ela vir primeiro seria
				// inventar informacao.
				if p, ja := posicoes[id]; ja {
					if p.tipo == SlotStatMod {
						continue
					}
				}
				est := st.ID
				if sl.Type == SlotStatMod {
					est = 0
				}
				posicoes[id] = posicao{estilo: est, tipo: sl.Type, linha: i}
			}
		}
	}

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
			LinhaSlot:            p.linha,
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

		hab := make([]Habilidade, 0, len(detPT.Spells))
		for i, sp := range detPT.Spells {
			nomeCanon := ""
			if i < len(det.Spells) {
				nomeCanon = det.Spells[i].SpellKey
			}
			slot := sp.SpellKey
			if slot == "" {
				slot = nomeCanon
			}
			hab = append(hab, Habilidade{
				Slot:      slot,
				Nome:      sp.Name,
				Descricao: Limpar(sp.DynamicDescription),
			})
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
			Habilidades: hab,
		})
	}
	return nil
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
