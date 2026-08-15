package export

import (
	"fmt"
	"sort"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
	"github.com/luansilvadb/lolbuilder/internal/gamedata"
	"github.com/luansilvadb/lolbuilder/internal/optimize"
)

// ---------- 05-computed.md ----------

func renderComputed(ds *canonical.Dataset) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Máximos calculados — patch %s\n\n", ds.Patch)
	b.WriteString("Duas tabelas, os dois resultados **exatos**: são o ótimo global, e não\n")
	b.WriteString("heurística. O que cada um responde está no título dele — e o que ele NÃO\n")
	b.WriteString("responde está logo abaixo.\n\n")

	b.WriteString("## Página de runas de maior atributo\n\n")
	b.WriteString("A página de valor máximo para cada atributo, no nível 18.\n\n")
	b.WriteString("> Só 11 das 69 runas jogáveis somam atributo. As outras 58 dependem de estado\n")
	b.WriteString("> de partida e não entram no cálculo — inclusive todas as pedras fundamentais.\n")
	b.WriteString("> **A página abaixo maximiza um atributo, e não a força do campeão.** Escolher\n")
	b.WriteString("> a pedra fundamental é quase sempre mais importante que o atributo, e essa\n")
	b.WriteString("> escolha o conjunto não calcula. Ver a coluna `escopo` em `02-runes.md`.\n\n")

	if len(ds.Computed.PaginasDeRuna) == 0 {
		b.WriteString("Nenhuma página calculada.\n\n")
	} else {
		b.WriteString("Slot marcado _indiferente_ é slot em que **nenhuma** opção soma nada ao\n")
		b.WriteString("atributo pedido: a escolha fica livre, e nomear uma runa ali seria publicar\n")
		b.WriteString("um desempate como se fosse recomendação. Nas páginas abaixo, a pedra\n")
		b.WriteString("fundamental é sempre indiferente — nenhuma delas concede atributo.\n\n")
		b.WriteString("| atributo | adaptativa | primário | secundário | runas que contribuem | slots livres | total |\n")
		b.WriteString("|---|---|---|---|---|---|---:|\n")
		for _, p := range ds.Computed.PaginasDeRuna {
			contribuem, livres := separarEscolhas(p)
			fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %d de 9 | %s |\n",
				celula(p.Objetivo), string(p.Resolucao),
				celula(p.NomePrimario), celula(p.NomeSecundario),
				celula(strings.Join(contribuem, ", ")), livres,
				celula(vetor(p.Total)))
		}
		b.WriteString("\n")
	}

	b.WriteString("## Máximo de atributo por ouro\n\n")
	b.WriteString("A combinação de até 6 itens que maximiza um atributo dentro de um orçamento,\n")
	b.WriteString("respeitando a regra de botas únicas.\n\n")
	b.WriteString("> **Isto não é uma build boa.** O cálculo ignora passiva e ativa de item, que\n")
	b.WriteString("> é o que faz metade dos itens valerem o que valem. O que está aqui é\n")
	b.WriteString("> literalmente \"o máximo deste atributo que cabe neste ouro\" — útil como piso\n")
	b.WriteString("> de comparação, inútil como recomendação de jogo.\n\n")

	for _, bl := range ds.Computed.BuildsDeItem {
		fmt.Fprintf(&b, "### %s\n\n", bl.Rotulo)
		fmt.Fprintf(&b, "%d de ouro dos %d disponíveis, em %d slots. Força adaptativa resolvida como %s.\n\n",
			bl.Gasto, bl.Orcamento, bl.Slots, bl.Resolucao)
		b.WriteString("| id | item | custo | stats |\n|---|---|---:|---|\n")
		for _, it := range bl.Itens {
			fmt.Fprintf(&b, "| %d | %s | %d | %s |\n",
				it.ID, celula(it.Nome), it.Custo, celula(vetor(it.Stats)))
		}
		fmt.Fprintf(&b, "\n**Total:** %s\n\n", vetor(bl.Total))
	}
	return b.String()
}

// separarEscolhas divide a pagina entre o que contribui e o que fica livre.
//
// Listar so a keystone e os fragmentos escondia de onde vinha o total: a pagina
// de cura e escudo tem os dois indiferentes, e os 5% vinham de uma runa
// secundaria que a tabela nao mostrava. O leitor via um numero sem origem.
func separarEscolhas(p optimize.Pagina) (contribuem []string, livres int) {
	todas := append([]optimize.Escolha{p.Keystone}, p.Menores...)
	todas = append(todas, p.Secundarias...)
	todas = append(todas, p.Fragmentos...)

	for _, e := range todas {
		if e.Indiferente {
			livres++
			continue
		}
		contribuem = append(contribuem, fmt.Sprintf("%s (%s)", e.Nome, vetor(e.Stats)))
	}
	return contribuem, livres
}

// ---------- 06-champion-stats.md ----------

func renderChampionStats(ds *canonical.Dataset) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Estatísticas de campeão — patch %s\n\n", ds.Patch)
	b.WriteString("Estatística base com crescimento por nível, e o efeito de cada habilidade\n")
	b.WriteString("por rank. Vem do dump de dados do jogo, e não do arquivo do cliente — o\n")
	b.WriteString("cliente publica esses campos zerados.\n\n")
	b.WriteString("**O efeito é uma expressão, e não um número.** `50 + 0.8 AP` significa 50 de\n")
	b.WriteString("dano fixo mais 80% do poder de habilidade. O conjunto publica os insumos; a\n")
	b.WriteString("conta com a build montada é de quem consome.\n\n")
	b.WriteString("Habilidade sem número na tabela é habilidade cuja fórmula a fonte não\n")
	b.WriteString("resolve — **nunca é zero**. As que ficaram sem número estão listadas no fim.\n\n")

	var semNumero []string

	for _, c := range ds.Champions {
		fmt.Fprintf(&b, "## %s\n\n", c.Nome)
		if c.Stats != nil {
			b.WriteString(tabelaDeStats(*c.Stats))
		}

		todas := append([]canonical.Habilidade{c.Passiva}, c.Habilidades...)
		todas = append(todas, c.SubHabilidades...)
		for _, h := range todas {
			linhas := linhasDeHabilidade(h)
			if len(linhas) == 0 {
				semNumero = append(semNumero,
					fmt.Sprintf("%s %s (%s)", c.Nome, strings.ToUpper(h.Slot), h.Nome))
				continue
			}
			fmt.Fprintf(&b, "**%s · %s**", strings.ToUpper(h.Slot), h.Nome)
			if len(h.Recarga) > 0 {
				fmt.Fprintf(&b, " · recarga %s", serie(h.Recarga))
			}
			if len(h.Custo) > 0 {
				fmt.Fprintf(&b, " · custo %s", serie(h.Custo))
			}
			// Alcance zerado nao e alcance: e habilidade que a fonte nao
			// descreve por distancia, como escudo em si mesmo.
			if len(h.Alcance) > 0 && !soZeros(h.Alcance) {
				fmt.Fprintf(&b, " · alcance %s", serie(h.Alcance))
			}
			b.WriteString("\n\n")
			for _, l := range linhas {
				fmt.Fprintf(&b, "- %s\n", l)
			}
			b.WriteString("\n")
		}
	}

	if len(semNumero) > 0 {
		b.WriteString("## Sem número na fonte\n\n")
		fmt.Fprintf(&b, "%d habilidades cuja fórmula a fonte não publica ou o extrator não\n", len(semNumero))
		b.WriteString("resolve. Estão aqui nomeadas uma a uma, e **não** publicadas com zero:\n")
		b.WriteString("zero afirmaria que a habilidade não causa dano, enquanto a ausência apenas\n")
		b.WriteString("não informa — e só a segunda é recuperável por quem lê.\n\n")
		sort.Strings(semNumero)
		for _, s := range semNumero {
			fmt.Fprintf(&b, "- %s\n", s)
		}
		b.WriteString("\n")
	}
	return b.String()
}

// soZeros informa se a serie inteira e zero.
func soZeros(v []float64) bool {
	for _, x := range v {
		if x != 0 {
			return false
		}
	}
	return true
}

func tabelaDeStats(s gamedata.Stats) string {
	var b strings.Builder
	b.WriteString("| estatística | base | por nível | no 18 |\n|---|---:|---:|---:|\n")
	linha := func(nome string, sc gamedata.Scaling) {
		fmt.Fprintf(&b, "| %s | %s | %s | %s |\n",
			nome, num(sc.Base), num(sc.PerLevel), num(sc.At(18)))
	}
	linha("vida", s.HP)
	linha("dano de ataque", s.AttackDamage)
	linha("armadura", s.Armor)
	linha("resistência mágica", s.MagicResist)
	linha("velocidade de ataque", s.AttackSpeed)
	if s.HPRegen != nil {
		linha("regeneração de vida", *s.HPRegen)
	}
	fmt.Fprintf(&b, "| velocidade de movimento | %s | — | %s |\n", num(s.MoveSpeed), num(s.MoveSpeed))
	fmt.Fprintf(&b, "| alcance de ataque | %s | — | %s |\n", num(s.AttackRange), num(s.AttackRange))
	if s.CritDamageMultiplier != nil {
		fmt.Fprintf(&b, "| multiplicador de crítico | %s | — | %s |\n",
			num(*s.CritDamageMultiplier), num(*s.CritDamageMultiplier))
	}
	b.WriteString("\n")
	return b.String()
}

// linhasDeHabilidade monta as linhas de efeito e de serie nomeada.
func linhasDeHabilidade(h canonical.Habilidade) []string {
	var out []string
	for _, e := range h.Efeitos {
		if len(e.PorRank) == 0 {
			continue
		}
		nome := e.Nome
		if e.DerivadoDe != "" {
			nome = fmt.Sprintf("%s (= %s × %s)", nome, e.DerivadoDe, num(e.Multiplicador))
		}
		linha := fmt.Sprintf("%s: %s", nome, expressoes(e.PorRank))
		if e.NivelDeReferencia > 0 {
			linha += fmt.Sprintf(" _(no nível %d)_", e.NivelDeReferencia)
		}
		out = append(out, linha)
	}
	for _, s := range h.SeriesNomeadas {
		out = append(out, fmt.Sprintf("%s: %s", s.Nome, serie(s.PorRank)))
	}
	return out
}

// expressoes formata a expressao de cada rank.
//
// A parcela fixa varia por rank e a escala nao, entao a escala sai UMA vez no
// fim: repeti-la em cada rank triplicaria o arquivo sem acrescentar informacao.
func expressoes(ranks []canonical.Expressao) string {
	fixos := make([]string, 0, len(ranks))
	for _, e := range ranks {
		fixos = append(fixos, num(e.Fixo))
	}
	out := strings.Join(fixos, "/")

	if len(ranks) > 0 && len(ranks[0].Escalas) > 0 {
		partes := make([]string, 0, len(ranks[0].Escalas))
		for _, es := range ranks[0].Escalas {
			p := fmt.Sprintf("%s %s", num(es.Coeficiente), es.Stat)
			if es.Parcela != "total" {
				p += " (" + es.Parcela + ")"
			}
			partes = append(partes, p)
		}
		out += " + " + strings.Join(partes, " + ")
	}
	return out
}
