package export

import (
	"fmt"
	"sort"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/canonical"
)

// O changelog e para o OPERADOR, e nao para o modelo.
//
// Ele nao vai ao Project: o consumidor quer saber como o jogo e agora, e nao
// como ele mudou. Quem precisa da diferenca e quem mantem o projeto, para
// decidir se uma runa mudou o bastante para exigir nova curadoria e se uma
// queda de cobertura tem explicacao.

func renderChangelog(in Input) string {
	ds := in.Dataset
	var b strings.Builder

	fmt.Fprintf(&b, "# Mudanças no patch %s\n\n", ds.Patch)

	if in.Anterior == nil {
		b.WriteString("Esta é a **captura inicial** — não há patch anterior para comparar.\n\n")
		b.WriteString("A partir do próximo patch, este arquivo passa a listar o que mudou.\n")
		return b.String()
	}
	fmt.Fprintf(&b, "Comparado com o patch %s.\n\n", in.Anterior.Patch)

	seccao := func(titulo string, linhas []string) {
		if len(linhas) == 0 {
			return
		}
		fmt.Fprintf(&b, "## %s\n\n", titulo)
		for _, l := range linhas {
			fmt.Fprintf(&b, "- %s\n", l)
		}
		b.WriteString("\n")
	}

	novos, sumiram, mudaram := diffItens(in.Anterior, ds)
	seccao("Itens novos na loja", novos)
	seccao("Itens que saíram da loja", sumiram)
	seccao("Itens com atributo alterado", mudaram)

	novasR, sumiramR, mudaramR := diffRunas(in.Anterior, ds)
	seccao("Runas novas", novasR)
	seccao("Runas removidas", sumiramR)
	seccao("Runas alteradas pela Riot", mudaramR)

	seccao("Campeões", diffCampeoes(in.Anterior, ds))
	seccao("Cobertura de extração", diffCobertura(in.Anterior, ds))

	if b.Len() == 0 {
		b.WriteString("Nada mudou nas entidades que o conjunto publica.\n")
	}
	return b.String()
}

func diffItens(antes, agora *canonical.Dataset) (novos, sumiram, mudaram []string) {
	a := indexarItens(antes)
	d := indexarItens(agora)

	for id, it := range d {
		velho, ok := a[id]
		if !ok {
			novos = append(novos, fmt.Sprintf("**%s** (%d), %d de ouro", it.Nome, id, it.Custo))
			continue
		}
		if dif := diferencaDeVetor(velho.Stats, it.Stats); dif != "" {
			mudaram = append(mudaram, fmt.Sprintf("**%s** (%d): %s", it.Nome, id, dif))
		}
		if velho.Custo != it.Custo {
			mudaram = append(mudaram,
				fmt.Sprintf("**%s** (%d): custo %d → %d", it.Nome, id, velho.Custo, it.Custo))
		}
	}
	for id, it := range a {
		if _, ok := d[id]; !ok {
			sumiram = append(sumiram, fmt.Sprintf("**%s** (%d)", it.Nome, id))
		}
	}
	sort.Strings(novos)
	sort.Strings(sumiram)
	sort.Strings(mudaram)
	return novos, sumiram, mudaram
}

func indexarItens(ds *canonical.Dataset) map[int32]canonical.Item {
	out := map[int32]canonical.Item{}
	for _, it := range ds.Purchasable() {
		out[it.ID] = it
	}
	return out
}

func diffRunas(antes, agora *canonical.Dataset) (novas, sumiram, mudaram []string) {
	a := indexarRunas(antes)
	d := indexarRunas(agora)

	for id, r := range d {
		velha, ok := a[id]
		if !ok {
			novas = append(novas, fmt.Sprintf("**%s** (%d) — precisa de curadoria", r.NomeCanonico, id))
			continue
		}
		// O gatilho e o campo que a propria fonte publica, e nao o texto: texto
		// muda por reescrita sem mudar numero.
		if velha.PatchDaUltimaMudanca != r.PatchDaUltimaMudanca {
			mudaram = append(mudaram, fmt.Sprintf(
				"**%s** (%d): a Riot marcou mudança em %s (antes %s) — revisar a curadoria",
				r.NomeCanonico, id, r.PatchDaUltimaMudanca, velha.PatchDaUltimaMudanca))
		}
	}
	for id, r := range a {
		if _, ok := d[id]; !ok {
			sumiram = append(sumiram, fmt.Sprintf("**%s** (%d)", r.NomeCanonico, id))
		}
	}
	sort.Strings(novas)
	sort.Strings(sumiram)
	sort.Strings(mudaram)
	return novas, sumiram, mudaram
}

func indexarRunas(ds *canonical.Dataset) map[int32]canonical.Rune {
	out := map[int32]canonical.Rune{}
	for _, r := range ds.Runes {
		if r.TipoSlot != "" {
			out[r.ID] = r
		}
	}
	return out
}

func diffCampeoes(antes, agora *canonical.Dataset) []string {
	a := map[int32]canonical.Champion{}
	for _, c := range antes.Champions {
		a[c.ID] = c
	}
	var out []string
	for _, c := range agora.Champions {
		velho, ok := a[c.ID]
		if !ok {
			out = append(out, fmt.Sprintf("**%s** (%d) é novo", c.Nome, c.ID))
			continue
		}
		if velho.Stats == nil || c.Stats == nil {
			continue
		}
		if velho.Stats.HP.Base != c.Stats.HP.Base {
			out = append(out, fmt.Sprintf("**%s**: vida base %s → %s",
				c.Nome, num(velho.Stats.HP.Base), num(c.Stats.HP.Base)))
		}
		if velho.Stats.AttackDamage.Base != c.Stats.AttackDamage.Base {
			out = append(out, fmt.Sprintf("**%s**: dano base %s → %s",
				c.Nome, num(velho.Stats.AttackDamage.Base), num(c.Stats.AttackDamage.Base)))
		}
	}
	delete(a, 0)
	for _, c := range antes.Champions {
		encontrado := false
		for _, n := range agora.Champions {
			if n.ID == c.ID {
				encontrado = true
				break
			}
		}
		if !encontrado {
			out = append(out, fmt.Sprintf("**%s** (%d) saiu", c.Nome, c.ID))
		}
	}
	sort.Strings(out)
	return out
}

// diffCobertura compara as taxas. Queda sem explicacao e o sinal de que a fonte
// mudou de forma.
func diffCobertura(antes, agora *canonical.Dataset) []string {
	var out []string
	cmp := func(nome string, a, d canonical.CoberturaDeEntidade) {
		ta, td := taxa(a.Resolvidas, a.Total), taxa(d.Resolvidas, d.Total)
		if ta == td {
			return
		}
		seta := "subiu"
		if td < ta {
			seta = "CAIU"
		}
		out = append(out, fmt.Sprintf("%s: %.1f%% → %.1f%% (%s)", nome, ta, td, seta))
	}
	cmp("habilidades resolvidas", antes.Coverage.Campeoes.Habilidades, agora.Coverage.Campeoes.Habilidades)
	cmp("passivas resolvidas", antes.Coverage.Campeoes.Passivas, agora.Coverage.Campeoes.Passivas)

	if a, d := antes.Coverage.Itens.Taxa(), agora.Coverage.Itens.Taxa(); a != d {
		out = append(out, fmt.Sprintf("leitura de atributo de item: %.1f%% → %.1f%%", a, d))
	}
	return out
}

func taxa(parte, total int) float64 {
	if total == 0 {
		return 0
	}
	return 100 * float64(parte) / float64(total)
}

// diferencaDeVetor descreve o que mudou entre dois conjuntos de stats.
func diferencaDeVetor(antes, agora canon.Vector) string {
	todos := map[canon.Stat]bool{}
	for s := range antes {
		todos[s] = true
	}
	for s := range agora {
		todos[s] = true
	}

	var partes []string
	for _, s := range canon.All {
		if !todos[s] {
			continue
		}
		a, d := antes[s], agora[s]
		if a == d {
			continue
		}
		partes = append(partes, fmt.Sprintf("%s %s → %s", rotulo(s), num(a), num(d)))
	}
	return strings.Join(partes, ", ")
}
