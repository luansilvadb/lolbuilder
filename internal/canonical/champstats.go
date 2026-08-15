package canonical

import (
	"fmt"

	"path/filepath"
	"sort"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/gamedata"
	"github.com/luansilvadb/lolbuilder/internal/model"
)

// buildChampStats acrescenta aos campeoes o que so o dump de dados do jogo tem:
// estatistica base e formula de habilidade.
//
// O plugin publica esses campos presentes e zerados. Publicar o zero dele seria
// afirmar que a habilidade nao causa dano, e nao que o dano e desconhecido.
func (b *Builder) buildChampStats(ds *Dataset) error {
	nomes, err := gamedata.LoadStatNames(filepath.Join(b.curationDir, "statenum.json"))
	if err != nil {
		return err
	}
	checker := gamedata.NewAlignChecker(b.cfg.CoverageMinimums.RankAlignment)
	cov := &ds.Coverage.Campeoes

	for i := range ds.Champions {
		c := &ds.Champions[i]

		raw, err := b.ler("characters/" + config.GameDataFile(c.Alias))
		if err != nil {
			return err
		}
		dump, err := gamedata.ParseDump(c.Alias, raw)
		if err != nil {
			return err
		}

		stats, gaps := dump.Stats()
		cov.CampeoesTotal++
		if stats.Complete() {
			cov.CampeoesComStats++
		}
		cov.LacunasDeStat = append(cov.LacunasDeStat, gaps...)
		c.Stats = &stats

		// A classificacao melee/ranged passa a vir do alcance de ataque do dump,
		// e nao do attackType do plugin. Sao a mesma informacao, mas o dump e a
		// fonte que tem o numero — e varias runas rendem valores diferentes para
		// os dois casos.
		c.CorpoACorpo = stats.Melee

		spells, err := dump.Spells()
		if err != nil {
			return err
		}
		if err := b.aplicarHabilidades(c, spells, nomes, cov, checker); err != nil {
			return err
		}
		b.aplicarSubHabilidades(c, dump, nomes, cov)
	}

	sort.Strings(cov.LacunasDeStat)
	ds.Coverage.Alinhamento = checker.Reports()
	sort.Slice(ds.Coverage.Alinhamento, func(i, j int) bool {
		return ds.Coverage.Alinhamento[i].Series < ds.Coverage.Alinhamento[j].Series
	})
	return nil
}

// aplicarHabilidades casa cada habilidade do dump com a do plugin e resolve as
// formulas.
func (b *Builder) aplicarHabilidades(
	c *Champion,
	spells []gamedata.Spell,
	nomes gamedata.StatNames,
	cov *CoberturaDeCampeoes,
	checker *gamedata.AlignChecker,
) error {
	// O plugin ja entregou nome e texto em portugues, na ordem Q, W, E, R. O
	// dump entrega os numeros na mesma ordem, porque os dois leem o mesmo array
	// CharacterRecord.spells.
	for _, sp := range spells {
		var alvo *Habilidade
		switch {
		case sp.Slot == gamedata.SlotPassive:
			alvo = &c.Passiva
		case int(sp.Slot) < len(c.Habilidades):
			alvo = &c.Habilidades[sp.Slot]
		default:
			return fmt.Errorf("%s: o dump tem a habilidade %s e o plugin nao",
				c.Alias, sp.Slot)
		}

		alvo.NomeInterno = sp.Name
		ranks := sp.Ranks()

		alvo.Recarga = recortar(sp.CooldownTime(), ranks)
		alvo.Custo = recortar(sp.Mana(), ranks)
		alvo.Alcance = alcancePublicavel(recortar(sp.CastRange, ranks), cov)

		if sp.Slot != gamedata.SlotPassive {
			who := c.Alias + " " + sp.Slot.String()
			checker.Compare("cooldown", who, sp.CooldownTime(), seriePlugin(c, sp.Slot, "recarga"), ranks)
			checker.Compare("mana", who, sp.Mana(), seriePlugin(c, sp.Slot, "custo"), ranks)
		}

		resolvidos, total := b.resolverEfeitos(alvo, &sp, nomes, ranks)
		alvo.SeriesNomeadas = seriesNaoConsumidas(&sp, ranks)
		alvo.TodasAsSeries = todasAsSeries(&sp, ranks)

		contar(cov, sp.Slot == gamedata.SlotPassive, resolvidos, total)
	}
	return nil
}

// resolverEfeitos avalia cada formula nomeada da habilidade em todos os ranks.
//
// Devolve quantas resolveram e quantas existem. Formula que nao resolve nao
// entra na saida: o efeito simplesmente nao e publicado, e o motivo fica na
// lista de lacunas.
func (b *Builder) resolverEfeitos(
	alvo *Habilidade, sp *gamedata.Spell, nomes gamedata.StatNames, ranks int,
) (resolvidos, total int) {
	chaves := make([]string, 0, len(sp.Calculations))
	for k := range sp.Calculations {
		chaves = append(chaves, k)
	}
	sort.Strings(chaves)

	for _, chave := range chaves {
		total++

		efeito := Efeito{Nome: chave}
		if d, ok := sp.DerivacaoDe(chave); ok {
			efeito.DerivadoDe = d.De
			efeito.Multiplicador = d.Multiplicador
		}

		ok := true
		for rank := 1; rank <= ranks; rank++ {
			ctx := gamedata.Context{Rank: rank, Level: nivelDeReferencia, Spell: sp, Stats: nomes}
			e, err := gamedata.Evaluate(chave, ctx)
			if err != nil {
				efeito.NaoResolvido = err.Error()
				ok = false
				break
			}
			// O nivel so e declarado quando ele MUDA o resultado. Declara-lo em
			// todo efeito acrescentaria ruido a maioria, que nao depende dele.
			if gamedata.DependeDoNivel(chave, ctx) {
				efeito.NivelDeReferencia = nivelDeReferencia
			}
			efeito.PorRank = append(efeito.PorRank, converter(e.Normalize()))
		}
		if ok {
			resolvidos++
		} else {
			efeito.PorRank = nil
		}
		alvo.Efeitos = append(alvo.Efeitos, efeito)
	}
	return resolvidos, total
}

// nivelDeReferencia e o nivel em que as parcelas que dependem do nivel do
// campeao sao avaliadas.
//
// 18 e o nivel maximo, e e o que o dataset original usa. Um efeito que
// interpola entre o nivel 1 e o 18 sai aqui no extremo superior; a interpolacao
// em si esta publicada no proprio dump para quem quiser outro nivel.
const nivelDeReferencia = 18

// recortar corta o preenchimento que a fonte coloca depois do ultimo rank real.
//
// Os arrays do dump sao maiores que o numero de ranks: sobra repeticao do
// ultimo valor ou zero. Publicar o preenchimento inventaria um rank 6 que nao
// existe no jogo.
func recortar(valores []float64, ranks int) []float64 {
	if len(valores) == 0 {
		return nil
	}
	out := make([]float64, 0, ranks)
	for rank := 1; rank <= ranks; rank++ {
		v, ok := gamedata.AtRank(valores, rank)
		if !ok {
			return out
		}
		out = append(out, v)
	}
	return out
}

// seriesNaoConsumidas devolve as series nomeadas que nenhuma formula da
// habilidade referencia.
//
// As consumidas ficam de fora para nao publicar o mesmo numero duas vezes: uma
// habilidade que resolve TotalDamage a partir de BaseDamage e ADRatio
// apresentaria o mesmo dano tres vezes se as duas series saissem ao lado do
// efeito ja resolvido, como se fossem parcelas somaveis.
//
// As NAO consumidas sao o oposto: sao numeros que a fonte publica e que nenhuma
// formula usa — duracao de lentidao, raio, tempo de recarga interno. Sem elas,
// muitas habilidades de utilidade sairiam sem numero nenhum.
func seriesNaoConsumidas(sp *gamedata.Spell, ranks int) []SerieNomeada {
	consumidas := sp.ConsumedDataValues()
	var out []SerieNomeada
	for _, dv := range sp.DataValues {
		if consumidas[strings.ToLower(dv.Name)] {
			continue
		}
		valores := recortar(dv.Values, ranks)
		if len(valores) == 0 || todosZero(valores) {
			continue
		}
		out = append(out, SerieNomeada{Nome: dv.Name, PorRank: valores})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Nome < out[j].Nome })
	return out
}

// todasAsSeries devolve TODAS as series nomeadas, inclusive as consumidas.
//
// Serve so para resolver o texto da habilidade, que referencia series que
// alguma formula ja consome. Nao e publicada: ao lado do efeito resolvido, ela
// apresentaria o mesmo numero duas vezes.
func todasAsSeries(sp *gamedata.Spell, ranks int) []SerieNomeada {
	var out []SerieNomeada
	for _, dv := range sp.DataValues {
		valores := recortar(dv.Values, ranks)
		if len(valores) == 0 {
			continue
		}
		out = append(out, SerieNomeada{Nome: dv.Name, PorRank: valores})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Nome < out[j].Nome })
	return out
}

func todosZero(v []float64) bool {
	for _, x := range v {
		if x != 0 {
			return false
		}
	}
	return true
}

// converter traduz a expressao do avaliador para a forma publicada.
func converter(e gamedata.Expr) Expressao {
	out := Expressao{Fixo: e.Flat}
	for _, t := range e.Terms {
		out.Escalas = append(out.Escalas, Escala{
			Stat:        t.Stat,
			Parcela:     t.Formula,
			Coeficiente: t.Coefficient,
		})
	}
	return out
}

// contar acumula a cobertura de extracao.
//
// Tres categorias, e nao duas: resolvida por inteiro, parcial, e sem formula
// alguma. A fracao sem formula precisa de contador proprio porque a taxa de
// resolucao sozinha nao a detecta — uma taxa calibrada para caber nas lacunas
// de hoje cabe tambem no dobro delas.
func contar(cov *CoberturaDeCampeoes, passiva bool, resolvidos, total int) {
	alvo := &cov.Habilidades
	if passiva {
		alvo = &cov.Passivas
	}
	contarEm(alvo, resolvidos, total)
}

func contarEm(alvo *CoberturaDeEntidade, resolvidos, total int) {
	alvo.Total++
	switch {
	case total == 0:
		alvo.SemFormula++
	case resolvidos == total:
		alvo.Resolvidas++
	case resolvidos > 0:
		alvo.Parciais++
	default:
		alvo.NaoResolvidas++
	}
}

// seriePlugin devolve a serie do plugin correspondente, para o cruzamento de
// alinhamento de rank.
func seriePlugin(c *Champion, slot gamedata.Slot, qual string) []float64 {
	if int(slot) >= len(c.Habilidades) {
		return nil
	}
	h := c.Habilidades[slot]
	if qual == "recarga" {
		return h.recargaDoPlugin
	}
	return h.custoDoPlugin
}

// lerSeriesDoPlugin guarda as series do plugin para o cruzamento posterior.
func lerSeriesDoPlugin(hab []Habilidade, spells []model.ChampionSpell) {
	for i := range hab {
		if i >= len(spells) {
			return
		}
		hab[i].recargaDoPlugin = spells[i].CooldownCoefficients
		hab[i].custoDoPlugin = spells[i].CostCoefficients
	}
}

// alcanceMaximoPlausivel separa alcance de uso do valor de sentinela.
//
// O limiar cai na lacuna observada no 16.16: a banda legitima termina em 7500,
// nos ultimates globais como o R do Pantheon, e a proxima ocorrencia e 10000.
// Nada existe entre os dois. Acima disso sao 37 habilidades em 10000, 132 em
// 25000, e casos ate 200000 — limites internos de missil, e nao alcance que o
// jogador enxerga. Conferi tambem se castRangeValues traria o valor real: zero
// casos. A alternativa nao existe, entao a escolha e entre publicar um numero
// errado e nao publicar nenhum.
const alcanceMaximoPlausivel = 10000

// alcancePublicavel descarta o alcance quando a fonte publica sentinela.
func alcancePublicavel(valores []float64, cov *CoberturaDeCampeoes) []float64 {
	for _, v := range valores {
		if v >= alcanceMaximoPlausivel {
			cov.AlcancesDescartados++
			return nil
		}
	}
	return valores
}

// aplicarSubHabilidades resolve as habilidades que um livro de feiticos
// acrescenta.
//
// Elas nao estao no array de slots do CharacterRecord, entao Spells() nao as
// alcanca. Sem isto, Hwei sairia com doze habilidades de texto e nenhum numero —
// e, pior, invisiveis na cobertura, porque o contador nunca soube que existiam.
func (b *Builder) aplicarSubHabilidades(
	c *Champion, dump *gamedata.Dump, nomes gamedata.StatNames, cov *CoberturaDeCampeoes,
) {
	for i := range c.SubHabilidades {
		h := &c.SubHabilidades[i]

		// O plugin chama a sub-habilidade de "qq"; o dump chama o objeto de
		// "HweiQQ". A ligacao e o alias mais o slot.
		sp, ok := dump.SpellByObjectName(c.Alias + strings.ToUpper(h.Slot))
		if !ok {
			cov.SubHabilidades.Total++
			cov.SubHabilidades.SemFormula++
			continue
		}

		h.NomeInterno = sp.Name
		ranks := sp.Ranks()
		h.Recarga = recortar(sp.CooldownTime(), ranks)
		h.Custo = recortar(sp.Mana(), ranks)
		h.Alcance = alcancePublicavel(recortar(sp.CastRange, ranks), cov)

		resolvidos, total := b.resolverEfeitos(h, sp, nomes, ranks)
		h.SeriesNomeadas = seriesNaoConsumidas(sp, ranks)
		h.TodasAsSeries = todasAsSeries(sp, ranks)

		contarEm(&cov.SubHabilidades, resolvidos, total)
	}
}
