package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/optimize"
)

type formatoSaida string

const (
	formatoTexto formatoSaida = "text"
	formatoJSON  formatoSaida = "json"
)

const notaBuild = "passivas e ativas dos itens não entram no cálculo"

type envelopeSaida struct {
	Comando             string      `json:"comando"`
	Patch               string      `json:"patch"`
	Objetivo            string      `json:"objetivo"`
	ResolucaoAdaptativa string      `json:"resolucao_adaptativa"`
	Resultado           interface{} `json:"resultado"`
	Avisos              []string    `json:"avisos"`
}

var nomesStat = map[canon.Stat]string{
	canon.AbilityHaste:        "aceleração de habilidade",
	canon.AbilityPower:        "poder de habilidade",
	canon.AdaptiveForce:       "força adaptativa",
	canon.Armor:               "armadura",
	canon.ArmorPenetrationPct: "penetração de armadura",
	canon.AttackDamage:        "dano de ataque",
	canon.AttackSpeedPct:      "velocidade de ataque",
	canon.BaseHealthRegenPct:  "regeneração de vida",
	canon.BaseManaRegenPct:    "regeneração de mana",
	canon.CriticalChancePct:   "chance de crítico",
	canon.CriticalDamagePct:   "dano crítico",
	canon.GoldPer10:           "ouro a cada 10 segundos",
	canon.HealShieldPowerPct:  "poder de cura e escudo",
	canon.Health:              "vida",
	canon.ItemHaste:           "aceleração de item",
	canon.SummonerSpellHaste:  "aceleração de feitiço",
	canon.Lethality:           "letalidade",
	canon.LifeStealPct:        "roubo de vida",
	canon.MagicPenetration:    "penetração mágica",
	canon.MagicPenetrationPct: "penetração mágica",
	canon.MagicResist:         "resistência mágica",
	canon.Mana:                "mana",
	canon.MoveSpeed:           "velocidade de movimento",
	canon.MoveSpeedPct:        "velocidade de movimento",
	canon.OmnivampPct:         "omnivamp",
	canon.TenacityPct:         "tenacidade",
}

func parseFormatoSaida(raw string) (formatoSaida, error) {
	switch formatoSaida(strings.ToLower(strings.TrimSpace(raw))) {
	case formatoTexto:
		return formatoTexto, nil
	case formatoJSON:
		return formatoJSON, nil
	default:
		return "", fmt.Errorf("formato invalido: %q (formatos aceitos: text, json)", raw)
	}
}

func emitirPaginas(formato formatoSaida, patch string, paginas []optimize.Pagina, objetivo *optimize.Objetivo) error {
	resultado := interface{}(paginas)
	if objetivo != nil && len(paginas) == 1 {
		resultado = paginas[0]
	}
	envelope := envelopeSaida{
		Comando:   "runes",
		Patch:     patch,
		Resultado: resultado,
		Avisos:    avisosPaginas(paginas),
	}
	if objetivo != nil {
		envelope.Objetivo = objetivo.String()
		envelope.ResolucaoAdaptativa = string(objetivo.Resolucao)
	}
	if formato == formatoJSON {
		return escreverJSON(envelope)
	}
	return escreverTexto(textoPaginas(patch, paginas))
}

func emitirBuilds(formato formatoSaida, patch string, builds []optimize.Build, objetivo *optimize.Objetivo) error {
	resultado := interface{}(builds)
	if objetivo != nil && len(builds) == 1 {
		resultado = builds[0]
	}
	envelope := envelopeSaida{
		Comando:   "builds",
		Patch:     patch,
		Resultado: resultado,
		Avisos:    []string{notaBuild},
	}
	if objetivo != nil {
		envelope.Objetivo = objetivo.String()
		envelope.ResolucaoAdaptativa = string(objetivo.Resolucao)
	}
	if formato == formatoJSON {
		return escreverJSON(envelope)
	}
	return escreverTexto(textoBuilds(patch, builds))
}

func escreverJSON(v envelopeSaida) error {
	return escreverJSONEm(os.Stdout, v)
}

func escreverJSONEm(w io.Writer, v envelopeSaida) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

func escreverTexto(s string) error {
	_, err := fmt.Fprint(os.Stdout, s)
	return err
}

func cabecalhoTexto(comando, patch string) string {
	return fmt.Sprintf("lolbuilder · %s · patch %s · snapshot local\n", comando, patch)
}

func textoPaginas(patch string, paginas []optimize.Pagina) string {
	var b strings.Builder
	b.WriteString(cabecalhoTexto("runes", patch))
	for i, p := range paginas {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(textoPagina(p))
	}
	return b.String()
}

func textoPagina(p optimize.Pagina) string {
	var b strings.Builder
	obj, ok := objetivoDaPagina(p)
	objetivo := p.Objetivo
	if ok {
		objetivo = nomeObjetivo(obj)
	}
	fmt.Fprintf(&b, "\n%s + %s · objetivo: %s · nivel %d · adaptativa: %s\n",
		p.NomePrimario, p.NomeSecundario, objetivo, p.Nivel, p.Resolucao)

	escolhas := escolhasPagina(p)
	for _, escolha := range escolhas {
		fmt.Fprintf(&b, "  %-11s %s\n", escolha.secao+":", textoEscolha(escolha.escolha))
	}

	contribuem := make([]string, 0)
	livres := 0
	for _, escolha := range escolhas {
		if escolha.escolha.Indiferente || escolha.escolha.Nome == "" {
			livres++
			continue
		}
		contribuem = append(contribuem, textoEscolha(escolha.escolha))
	}
	if len(contribuem) == 0 {
		b.WriteString("  contribuem: nenhuma\n")
		if ok {
			fmt.Fprintf(&b, "  aviso: nenhuma runa contribui para %s\n", nomeObjetivo(obj))
		}
	} else {
		fmt.Fprintf(&b, "  contribuem: %s\n", strings.Join(contribuem, ", "))
	}
	fmt.Fprintf(&b, "  slots livres: %d de %d\n", livres, len(escolhas))
	if ok {
		fmt.Fprintf(&b, "  total: %s\n", resumoObjetivo(obj, p.Total))
	} else {
		fmt.Fprintf(&b, "  total: %s\n", formatVector(p.Total))
	}
	return b.String()
}

type escolhaPagina struct {
	secao   string
	escolha optimize.Escolha
}

func escolhasPagina(p optimize.Pagina) []escolhaPagina {
	out := []escolhaPagina{{secao: "keystone", escolha: p.Keystone}}
	for _, e := range p.Menores {
		out = append(out, escolhaPagina{secao: "primaria", escolha: e})
	}
	for _, e := range p.Secundarias {
		out = append(out, escolhaPagina{secao: "secundaria", escolha: e})
	}
	for _, e := range p.Fragmentos {
		out = append(out, escolhaPagina{secao: "fragmento", escolha: e})
	}
	return out
}

func textoEscolha(e optimize.Escolha) string {
	if e.Indiferente || e.Nome == "" {
		return "indiferente"
	}
	if len(e.Stats) == 0 {
		return e.Nome
	}
	return fmt.Sprintf("%s (%s)", e.Nome, formatVector(e.Stats))
}

func avisosPaginas(paginas []optimize.Pagina) []string {
	avisos := []string{}
	for _, p := range paginas {
		if len(contribuicoesPagina(p)) == 0 {
			obj, ok := objetivoDaPagina(p)
			if ok {
				avisos = append(avisos, "nenhuma runa contribui para "+nomeObjetivo(obj))
			}
		}
	}
	return avisos
}

func contribuicoesPagina(p optimize.Pagina) []optimize.Escolha {
	out := make([]optimize.Escolha, 0)
	for _, e := range escolhasPagina(p) {
		if !e.escolha.Indiferente && e.escolha.Nome != "" {
			out = append(out, e.escolha)
		}
	}
	return out
}

func textoBuilds(patch string, builds []optimize.Build) string {
	var b strings.Builder
	b.WriteString(cabecalhoTexto("builds", patch))
	for _, build := range builds {
		b.WriteString(textoBuild(build))
	}
	return b.String()
}

func textoBuild(build optimize.Build) string {
	var b strings.Builder
	obj, ok := objetivoDaBuild(build)
	rotulo := build.Rotulo
	if ok {
		rotulo = fmt.Sprintf("máximo de %s por ouro, ignorando efeitos de item", nomeObjetivo(obj))
	}
	fmt.Fprintf(&b, "\n%s\n", rotulo)
	if ok {
		fmt.Fprintf(&b, "  objetivo: %s\n", resumoObjetivo(obj, build.Total))
	}
	sobra := build.Orcamento - build.Gasto
	fmt.Fprintf(&b, "  gasto: %s / %s ouro (sobra %s)\n",
		formatNumber(float64(build.Gasto)), formatNumber(float64(build.Orcamento)), formatNumber(float64(sobra)))
	fmt.Fprintf(&b, "  força adaptativa: %s\n", build.Resolucao)
	b.WriteString("  itens:\n")
	for _, it := range build.Itens {
		fmt.Fprintf(&b, "    %-6d %-32s %s ouro\n", it.ID, it.Nome, formatNumber(float64(it.Custo)))
	}
	if ok {
		adicionais := vetorAdicional(build.Total, obj)
		if len(adicionais) == 0 {
			b.WriteString("  atributos adicionais: nenhum\n")
		} else {
			fmt.Fprintf(&b, "  atributos adicionais: %s\n", formatVector(adicionais))
		}
	} else {
		fmt.Fprintf(&b, "  total: %s\n", formatVector(build.Total))
	}
	fmt.Fprintf(&b, "  nota: %s\n", notaBuild)
	return b.String()
}

func objetivoDaPagina(p optimize.Pagina) (optimize.Objetivo, bool) {
	obj, err := optimize.ParseObjetivo(p.Objetivo, p.Resolucao)
	return obj, err == nil
}

func objetivoDaBuild(build optimize.Build) (optimize.Objetivo, bool) {
	const prefixo = "maximo de "
	const sufixo = " por ouro, ignorando efeitos de item"
	if !strings.HasPrefix(build.Rotulo, prefixo) || !strings.HasSuffix(build.Rotulo, sufixo) {
		return optimize.Objetivo{}, false
	}
	raw := strings.TrimSuffix(strings.TrimPrefix(build.Rotulo, prefixo), sufixo)
	obj, err := optimize.ParseObjetivo(raw, build.Resolucao)
	return obj, err == nil
}

func nomeObjetivo(obj optimize.Objetivo) string {
	stats := obj.Stats()
	if len(stats) == 1 && obj.Pesos[stats[0]] == 1 {
		return nomeStat(stats[0])
	}
	return obj.String()
}

func resumoObjetivo(obj optimize.Objetivo, total canon.Vector) string {
	stats := obj.Stats()
	if len(stats) == 1 && obj.Pesos[stats[0]] == 1 {
		s := stats[0]
		valor := total[s]
		if s == obj.Resolucao.StatResolvido() {
			valor += total[canon.AdaptiveForce]
		}
		return formatValorStat(s, valor)
	}
	return formatNumber(obj.Valor(total)) + " pontos de objetivo"
}

func vetorAdicional(total canon.Vector, obj optimize.Objetivo) canon.Vector {
	out := canon.Vector{}
	stats := obj.Stats()
	for s, valor := range total {
		remover := false
		for _, alvo := range stats {
			if s == alvo || (s == canon.AdaptiveForce && alvo == obj.Resolucao.StatResolvido()) {
				remover = true
				break
			}
		}
		if !remover {
			out[s] = valor
		}
	}
	return out
}

func formatVector(v canon.Vector) string {
	if len(v) == 0 {
		return "nenhum"
	}
	partes := make([]string, 0, len(v))
	for _, s := range v.Stats() {
		partes = append(partes, formatValorStat(s, v[s]))
	}
	return strings.Join(partes, ", ")
}

func formatValorStat(s canon.Stat, valor float64) string {
	numero := formatNumber(valor)
	if s.Percentual() {
		return numero + "% " + nomeStat(s)
	}
	return numero + " " + nomeStat(s)
}

func nomeStat(s canon.Stat) string {
	if nome, ok := nomesStat[s]; ok {
		return nome
	}
	return string(s)
}

func formatNumber(valor float64) string {
	if valor == 0 {
		return "0"
	}
	raw := strconv.FormatFloat(valor, 'f', -1, 64)
	sinal := ""
	if strings.HasPrefix(raw, "-") {
		sinal = "-"
		raw = strings.TrimPrefix(raw, "-")
	}
	partes := strings.SplitN(raw, ".", 2)
	inteiro := partes[0]
	for i := len(inteiro) - 3; i > 0; i -= 3 {
		inteiro = inteiro[:i] + "." + inteiro[i:]
	}
	if len(partes) == 1 {
		return sinal + inteiro
	}
	return sinal + inteiro + "," + partes[1]
}
