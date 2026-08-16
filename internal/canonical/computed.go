package canonical

import (
	"fmt"
	"path/filepath"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/optimize"
)

// Computed sao os maximos exatos que o dataset pre-calcula.
//
// Ficam no modelo canonico, e nao no export, porque sao dado derivado do mesmo
// snapshot: o export so os formata.
type Computed struct {
	// PaginasDeRuna traz a pagina de valor maximo por objetivo.
	PaginasDeRuna []optimize.Pagina `json:"paginas_de_runa,omitempty"`
	// BuildsDeItem traz o maximo por ouro por objetivo. Ver o Rotulo de cada
	// uma: NAO sao builds otimas.
	BuildsDeItem []optimize.Build `json:"builds_de_item,omitempty"`
}

// Parametros do pre-calculo. Ficam aqui, e nao em config.json, porque nao sao
// do modo: sao a pergunta que o dataset escolhe responder.
const (
	// nivelDoPreCalculo e o nivel em que as paginas sao avaliadas.
	nivelDoPreCalculo = 18
	// slotsDeInventario e o numero de espacos de item de um campeao.
	slotsDeInventario = 6
	// orcamentoDoPreCalculo e o ouro de uma partida longa. Nao e teto do jogo:
	// e a pergunta "o que cabe com este ouro", e o valor esta publicado junto do
	// resultado para que ninguem o tome por universal.
	orcamentoDoPreCalculo = 20000
)

// buildComputed monta as tabelas pre-calculadas.
func (b *Builder) buildComputed(ds *Dataset) error {
	cur, err := canon.LoadRuneCuration(filepath.Join(b.curationDir, "runes.json"))
	if err != nil {
		return err
	}
	if err := b.conferirCuradoria(ds, cur); err != nil {
		return err
	}

	aplicarCuradoria(ds, cur)
	cat := montarCatalogo(ds, cur)
	itens := CandidatosDeItem(ds)

	for _, stat := range canon.All {
		// Forca adaptativa nao e alvo em si: ela sempre resolve para outro
		// stat antes de valer alguma coisa. Ver ParseObjetivo.
		if stat == canon.AdaptiveForce {
			continue
		}
		for _, res := range []optimize.Resolucao{optimize.ResolucaoAD, optimize.ResolucaoAP} {
			obj, err := optimize.ParseObjetivo(string(stat), res)
			if err != nil {
				return err
			}

			p, err := optimize.MelhorPagina(cat, obj, nivelDoPreCalculo)
			if err != nil {
				return err
			}
			if p.Valor > 0 {
				ds.Computed.PaginasDeRuna = append(ds.Computed.PaginasDeRuna, p)
			}

			bl, err := optimize.MelhorBuild(itens, GruposParaMochila(ds), obj, slotsDeInventario, orcamentoDoPreCalculo)
			if err != nil {
				return err
			}
			if bl.Valor > 0 {
				ds.Computed.BuildsDeItem = append(ds.Computed.BuildsDeItem, bl)
			}

			// Sem forca adaptativa em jogo, as duas resolucoes dao o mesmo
			// resultado. Publicar as duas seria repetir a mesma linha.
			if !envolveAdaptativa(p, bl) {
				break
			}
		}
	}
	return nil
}

// envolveAdaptativa informa se a resolucao declarada mudou alguma coisa.
func envolveAdaptativa(p optimize.Pagina, b optimize.Build) bool {
	if _, ok := p.Total[canon.AdaptiveForce]; ok {
		return true
	}
	_, ok := b.Total[canon.AdaptiveForce]
	return ok
}

// conferirCuradoria aborta se alguma runa jogavel ficou sem entrada ou mudou.
//
// Sob curadoria parcial, uma runa nova que soma stat entra como zero
// silencioso — o modo de falha que este mecanismo existe para tornar
// impossivel.
func (b *Builder) conferirCuradoria(ds *Dataset, cur *canon.RuneCuration) error {
	var jogaveis []canon.RunaPublicada
	for _, r := range ds.Runes {
		if r.TipoSlot == "" {
			// Runa que o jogo nao oferece: removida, Template, ou titulo de
			// pagina. Curar entidade que nao se pode escolher seria ruido.
			continue
		}
		jogaveis = append(jogaveis, canon.RunaPublicada{
			ID: r.ID, Nome: r.NomeCanonico, PatchDaUltimaMudanca: r.PatchDaUltimaMudanca,
		})
	}

	naoCuradas, desatualizadas := cur.Conferir(jogaveis)
	if len(naoCuradas) > 0 {
		return fmt.Errorf(
			"%d runa(s) jogavel(is) sem entrada em curation/runes.json: %v — "+
				"classifique cada uma antes de publicar, porque runa somavel sem "+
				"curadoria entra no dataset como zero", len(naoCuradas), naoCuradas)
	}
	if len(desatualizadas) > 0 {
		return fmt.Errorf(
			"%d runa(s) mudaram desde a curadoria e precisam de revisao: %v — "+
				"a fonte avancou o majorChangePatchVersion delas",
			len(desatualizadas), desatualizadas)
	}
	return nil
}

// aplicarCuradoria leva a classificacao curada para o modelo publicado.
//
// Sem isto, a curadoria so existiria dentro do otimizador, e o arquivo de runas
// listaria as 69 sem explicar por que 58 delas nao aparecem no pre-calculo — o
// mesmo que nao declarar limite nenhum.
func aplicarCuradoria(ds *Dataset, cur *canon.RuneCuration) {
	for i := range ds.Runes {
		r := &ds.Runes[i]
		c, ok := cur.Get(r.ID)
		if !ok {
			continue
		}
		r.Escopo = string(c.Kind)
		r.MotivoDoEscopo = c.Reason
		r.RessalvaDoEscopo = c.Note
		if v := c.VetorNoNivel(nivelDoPreCalculo); len(v) > 0 {
			r.StatsDaRuna = v
		}
	}
}

// CatalogoDeRunas monta o catalogo do otimizador lendo a curadoria do disco.
//
// Existe para o comando de linha, que resolve um objetivo avulso sem passar
// pelo pre-calculo.
func CatalogoDeRunas(ds *Dataset, curationDir string) (optimize.Catalogo, error) {
	cur, err := canon.LoadRuneCuration(filepath.Join(curationDir, "runes.json"))
	if err != nil {
		return optimize.Catalogo{}, err
	}
	return montarCatalogo(ds, cur), nil
}

// montarCatalogo monta o que o otimizador de pagina precisa.
func montarCatalogo(ds *Dataset, cur *canon.RuneCuration) optimize.Catalogo {
	cat := optimize.Catalogo{
		Curadas: map[int32]*canon.CuratedRune{},
		Nomes:   map[int32]string{},
	}
	for _, r := range ds.Runes {
		cat.Nomes[r.ID] = r.Nome
		if c, ok := cur.Get(r.ID); ok {
			cat.Curadas[r.ID] = c
		}
	}
	for _, st := range ds.RuneStyles {
		e := optimize.Estilo{
			ID: st.ID, Nome: st.Nome,
			SubEstilosPermitidos: st.SubEstilosPermitidos,
		}
		for _, l := range st.Linhas {
			switch l.Tipo {
			case SlotKeyStone:
				e.Keystones = append(e.Keystones, l.Runas...)
			case SlotMenor:
				e.Menores = append(e.Menores, l.Runas)
			case SlotStatMod:
				e.Fragmentos = append(e.Fragmentos, l.Runas)
			}
		}
		cat.Estilos = append(cat.Estilos, e)
	}
	return cat
}

// GruposParaMochila traduz os limites do dataset para o vocabulario da mochila.
func GruposParaMochila(ds *Dataset) []optimize.Grupo {
	out := make([]optimize.Grupo, 0, len(ds.GruposDeItem))
	for _, g := range ds.GruposDeItem {
		out = append(out, optimize.Grupo{ID: g.ID, Maximo: g.Maximo})
	}
	return out
}

// CandidatosDeItem monta o que a mochila precisa, so com o que e compravel.
func CandidatosDeItem(ds *Dataset) []optimize.ItemCandidato {
	out := make([]optimize.ItemCandidato, 0, 256)
	for _, it := range ds.Items {
		if !it.Compravel || len(it.Stats) == 0 {
			continue
		}
		out = append(out, optimize.ItemCandidato{
			ID: it.ID, Nome: it.Nome, Custo: it.Custo,
			Stats: it.Stats, Grupos: it.Grupos,
		})
	}
	return out
}
