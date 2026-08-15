package export

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/canonical"
)

// rotuloDeStat traduz o stat canonico para o texto curto das tabelas.
//
// E deliberadamente compacto: o mesmo rotulo aparece centenas de vezes num
// arquivo que tem teto de tamanho. A convencao de unidades esta declarada uma
// vez no manifesto, e nao repetida em cada linha.
var rotuloDeStat = map[canon.Stat]string{
	canon.AbilityHaste:        "acel",
	canon.AbilityPower:        "AP",
	canon.AdaptiveForce:       "adapt",
	canon.Armor:               "armadura",
	canon.ArmorPenetrationPct: "%pen armadura",
	canon.AttackDamage:        "AD",
	canon.AttackSpeedPct:      "%AS",
	canon.BaseHealthRegenPct:  "%regen vida",
	canon.BaseManaRegenPct:    "%regen mana",
	canon.CriticalChancePct:   "%crit",
	canon.CriticalDamagePct:   "%dano crit",
	canon.GoldPer10:           "ouro/10",
	canon.HealShieldPowerPct:  "%cura e escudo",
	canon.Health:              "vida",
	canon.ItemHaste:           "acel item",
	canon.SummonerSpellHaste:  "acel feitico",
	canon.Lethality:           "letalidade",
	canon.LifeStealPct:        "%roubo de vida",
	canon.MagicPenetration:    "pen magica",
	canon.MagicPenetrationPct: "%pen magica",
	canon.MagicResist:         "MR",
	canon.Mana:                "mana",
	canon.MoveSpeed:           "mov",
	canon.MoveSpeedPct:        "%mov",
	canon.OmnivampPct:         "%omnivamp",
	canon.TenacityPct:         "%tenacidade",
}

func rotulo(s canon.Stat) string {
	if r, ok := rotuloDeStat[s]; ok {
		return r
	}
	return string(s)
}

// num formata um numero na precisao que a fonte de fato carrega.
//
// O jogo calcula em float32 e o dump serializa a representacao float64 desse
// float32, entao 4.2 chega como 4.199999809265137. Publicar os 17 digitos seria
// duas coisas ruins ao mesmo tempo: mentir sobre a precisao do dado, e gastar
// cinco tokens onde um basta — num arquivo que tem teto de tamanho e dezenas de
// milhares desses numeros.
//
// FormatFloat com bitSize 32 devolve a representacao mais curta que reproduz o
// mesmo float32. Nao e arredondamento: e parar de imprimir digitos que a fonte
// nao tem.
func num(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 32)
}

// vetor formata um conjunto de stats em ordem canonica.
func vetor(v canon.Vector) string {
	if len(v) == 0 {
		return "—"
	}
	partes := make([]string, 0, len(v))
	for _, s := range v.Stats() {
		partes = append(partes, num(v[s])+" "+rotulo(s))
	}
	return strings.Join(partes, ", ")
}

// serie formata uma serie por rank.
func serie(vs []float64) string {
	if len(vs) == 0 {
		return "—"
	}
	partes := make([]string, 0, len(vs))
	for _, v := range vs {
		partes = append(partes, num(v))
	}
	return strings.Join(partes, "/")
}

// ids formata uma lista de ids.
func ids(v []int32) string {
	if len(v) == 0 {
		return "—"
	}
	partes := make([]string, 0, len(v))
	for _, x := range v {
		partes = append(partes, strconv.Itoa(int(x)))
	}
	return strings.Join(partes, " + ")
}

// celula deixa o texto seguro dentro de uma tabela markdown.
func celula(s string) string {
	if s == "" {
		return "—"
	}
	s = strings.ReplaceAll(s, "|", "\\|")
	return strings.Join(strings.Fields(s), " ")
}

// ---------- 01-items.md ----------

func renderItems(ds *canonical.Dataset) string {
	var b strings.Builder
	compraveis := ds.Purchasable()

	fmt.Fprintf(&b, "# Itens — patch %s\n\n", ds.Patch)
	fmt.Fprintf(&b, "%d itens compráveis no Summoner's Rift. `custo` é o total acumulado;\n", len(compraveis))
	b.WriteString("`combina` é o que se paga a partir dos componentes. Stats percentuais em\n")
	b.WriteString("pontos percentuais.\n\n")
	b.WriteString("> A coluna `stats` vem do bloco de atributos da fonte, e no LoL moderno ela\n")
	b.WriteString("> **inclui a velocidade de movimento das botas**. O que continua fora dela são\n")
	b.WriteString("> os efeitos: passiva e ativa dependem de estado de combate e estão só na\n")
	b.WriteString("> coluna `efeito`, em texto.\n\n")

	b.WriteString("| id | item | nome canônico | custo | combina | componentes | stats | efeito |\n")
	b.WriteString("|---|---|---|---:|---:|---|---|---|\n")
	for _, it := range compraveis {
		fmt.Fprintf(&b, "| %d | %s | %s | %d | %d | %s | %s | %s |\n",
			it.ID, celula(it.Nome), celula(it.NomeCanonico), it.Custo, it.Combina,
			ids(it.Componentes), celula(vetor(it.Stats)), celula(it.Efeito))
	}
	return b.String()
}

// ---------- 02-runes.md ----------

func renderRunes(ds *canonical.Dataset) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Runas — patch %s\n\n", ds.Patch)
	b.WriteString("A página tem 5 estilos. O primário dá 1 pedra fundamental mais 3 runas\n")
	b.WriteString("menores, uma por linha; o secundário dá 2 runas menores de linhas diferentes;\n")
	b.WriteString("e 3 fragmentos de stat completam a página.\n\n")
	b.WriteString("A coluna `escopo` diz se a runa entra no cálculo de `05-computed.md`:\n")
	b.WriteString("`sum` soma atributo fixo, `sum_per_level` soma um valor que cresce com o\n")
	b.WriteString("nível, e `out_of_scope` depende de estado de partida. **A maioria é\n")
	b.WriteString("`out_of_scope`, e isso não é lacuna do conjunto**: no sistema moderno a pedra\n")
	b.WriteString("fundamental define comportamento, e não atributo. O motivo de cada uma está\n")
	b.WriteString("na coluna `motivo`.\n\n")

	porEstilo := map[int32][]canonical.Rune{}
	var fragmentos []canonical.Rune
	for _, r := range ds.Runes {
		switch {
		case r.TipoSlot == "":
			continue // runa que o jogo nao oferece
		case r.Fragmento():
			fragmentos = append(fragmentos, r)
		default:
			porEstilo[r.Estilo] = append(porEstilo[r.Estilo], r)
		}
	}

	for _, st := range ds.RuneStyles {
		fmt.Fprintf(&b, "## %s (%s)\n\n", st.Nome, st.NomeCanonico)
		rs := porEstilo[st.ID]
		sort.Slice(rs, func(i, j int) bool {
			if rs[i].TipoSlot != rs[j].TipoSlot {
				return rs[i].TipoSlot == canonical.SlotKeyStone
			}
			return rs[i].ID < rs[j].ID
		})
		tabelaDeRunas(&b, rs)
	}

	b.WriteString("## Fragmentos de stat\n\n")
	b.WriteString("Valem em qualquer estilo. São 3 linhas, e uma runa pode aparecer em mais de\n")
	b.WriteString("uma: Força Adaptativa está na linha ofensiva e na flexível.\n\n")
	sort.Slice(fragmentos, func(i, j int) bool { return fragmentos[i].ID < fragmentos[j].ID })
	tabelaDeRunas(&b, fragmentos)
	return b.String()
}

func tabelaDeRunas(b *strings.Builder, rs []canonical.Rune) {
	// As duas ultimas colunas sao coisas diferentes e por isso nao se juntam:
	// "fora do calculo porque" explica a exclusao de uma runa out_of_scope, e
	// "ressalva" registra a parte do efeito que ficou de fora numa runa que soma
	// EM PARTE. Sob um rotulo unico, o leitor nao teria como saber qual dos dois
	// esta lendo.
	b.WriteString("| id | runa | nome canônico | tipo | escopo | stats | efeito | fora do cálculo porque | ressalva |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|---|\n")
	for _, r := range rs {
		tipo := "menor"
		switch {
		case r.Keystone():
			tipo = "pedra fundamental"
		case r.Fragmento():
			tipo = "fragmento"
		}
		fmt.Fprintf(b, "| %d | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			r.ID, celula(r.Nome), celula(r.NomeCanonico), tipo, r.Escopo,
			celula(vetor(r.StatsDaRuna)), celula(r.Resumo),
			celula(r.MotivoDoEscopo), celula(r.RessalvaDoEscopo))
	}
	b.WriteString("\n")
}

// ---------- 03-summoner-spells.md ----------

func renderSpells(ds *canonical.Dataset) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Feitiços de invocador — patch %s\n\n", ds.Patch)
	fmt.Fprintf(&b, "%d feitiços válidos no Summoner's Rift. A recarga é a base, sem\n",
		len(ds.SummonerSpells))
	b.WriteString("aceleração de feitiço.\n\n")
	b.WriteString("| id | feitiço | nome canônico | recarga | nível mínimo | efeito |\n")
	b.WriteString("|---|---|---|---:|---:|---|\n")
	for _, s := range ds.SummonerSpells {
		fmt.Fprintf(&b, "| %d | %s | %s | %ds | %d | %s |\n",
			s.ID, celula(s.Nome), celula(s.NomeCanonico), s.Recarga, s.NivelMinimo,
			celula(s.Descricao))
	}
	return b.String()
}

// ---------- 04-champions.md ----------

func renderChampions(ds *canonical.Dataset) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Campeões — patch %s\n\n", ds.Patch)
	fmt.Fprintf(&b, "%d campeões. As estatísticas base e o efeito das habilidades por rank\n",
		len(ds.Champions))
	b.WriteString("estão em `06-champion-stats.md`; aqui está o que cada habilidade faz.\n\n")
	b.WriteString("> O texto vem do cliente, que o publica com marcadores no lugar dos números.\n")
	b.WriteString("> Eles foram **resolvidos** contra as séries do dump: onde a fonte escreve\n")
	b.WriteString("> `@TotalDamage@`, aqui está o valor por rank. Marcador sem valor na fonte foi\n")
	b.WriteString("> removido em vez de publicado cru — um `@Nome@` visível não informa nada e\n")
	b.WriteString("> ainda parece defeito de geração.\n\n")

	naoResolvidos := 0
	for _, c := range ds.Champions {
		alcance := "a distância"
		if c.CorpoACorpo {
			alcance = "corpo a corpo"
		}
		fmt.Fprintf(&b, "## %s, %s\n\n", c.Nome, c.Titulo)
		fmt.Fprintf(&b, "`%s` · id %d · %s · %s\n\n",
			c.NomeCanonico, c.ID, alcance, strings.Join(c.Papeis, ", "))

		linha := func(prefixo string, h canonical.Habilidade) {
			r := novoResolvedor(h)
			texto := r.Resolver(h.Descricao)
			naoResolvidos += r.NaoResolvidos
			fmt.Fprintf(&b, "%s**%s · %s** — %s\n",
				prefixo, strings.ToUpper(h.Slot), h.Nome, celula(texto))
		}
		linha("- ", c.Passiva)
		for _, h := range c.Habilidades {
			linha("- ", h)
		}
		for _, h := range c.SubHabilidades {
			linha("  - ", h)
		}
		b.WriteString("\n")
	}

	if naoResolvidos > 0 {
		fmt.Fprintf(&b,
			"---\n\n_%d ocorrências de `(?)` acima: são números que a fonte não publica._\n",
			naoResolvidos)
	}
	return b.String()
}
