// Package optimize resolve os maximos exatos que o dataset publica.
//
// Duas perguntas, dois algoritmos, a mesma disciplina: o resultado e o otimo
// global, e nao uma heuristica. Onde ha atalho, ele e verificado contra busca
// exaustiva num teste — um otimizador que erra publica numero errado com cara
// de exato, que e pior que nao publicar nada.
package optimize

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/luansilvadb/lolbuilder/internal/canon"
)

// Resolucao diz em que atributo a forca adaptativa vira numero.
//
// Ela e deterministica DADA esta declaracao — e o que a distingue de efeito de
// fato condicional. O resultado publica a resolucao junto, nunca a esconde
// dentro do numero.
type Resolucao string

const (
	// ResolucaoAD trata forca adaptativa como dano de ataque. E o desempate do
	// proprio jogo quando os dois bonus estao iguais.
	ResolucaoAD Resolucao = "ad"
	// ResolucaoAP trata forca adaptativa como poder de habilidade.
	ResolucaoAP Resolucao = "ap"
)

// StatResolvido devolve em que stat a forca adaptativa entra.
func (r Resolucao) StatResolvido() canon.Stat {
	if r == ResolucaoAP {
		return canon.AbilityPower
	}
	return canon.AttackDamage
}

// Objetivo e o que maximizar: um peso por stat.
//
// E linear por construcao. Objetivo nao linear — "dano por segundo", "vida
// efetiva" — depende de modelo de combate, que o dataset declara nao ter.
type Objetivo struct {
	Pesos     canon.Vector
	Resolucao Resolucao
}

// ParseObjetivo le a forma de linha de comando.
//
// Aceita "armor" (peso 1) e "attack_damage:1,attack_speed_pct:100", em que o
// peso torna comparaveis grandezas de escala diferente.
func ParseObjetivo(s string, res Resolucao) (Objetivo, error) {
	obj := Objetivo{Pesos: canon.Vector{}, Resolucao: res}
	if res != ResolucaoAD && res != ResolucaoAP {
		return obj, fmt.Errorf("resolucao de forca adaptativa invalida: %q (use ad ou ap)", res)
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return obj, fmt.Errorf("objetivo vazio")
	}
	for _, parte := range strings.Split(s, ",") {
		parte = strings.TrimSpace(parte)
		if parte == "" {
			continue
		}
		nome, pesoTxt, temPeso := strings.Cut(parte, ":")
		stat := canon.Stat(strings.TrimSpace(nome))
		if !statConhecido(stat) {
			return obj, fmt.Errorf("stat %q fora do vocabulario canonico; validos: %s",
				nome, listaDeStats())
		}
		// Pedir o maximo de forca adaptativa nao tem resposta: ela SEMPRE vira
		// dano de ataque ou poder de habilidade antes de valer alguma coisa.
		// Aceitar o pedido devolveria zero em silencio.
		if stat == canon.AdaptiveForce {
			return obj, fmt.Errorf(
				"forca adaptativa resolve para dano de ataque ou poder de habilidade, " +
					"e nao e alvo em si — peca attack_damage ou ability_power, e use " +
					"-adaptive para dizer como resolve-la")
		}
		peso := 1.0
		if temPeso {
			p, err := strconv.ParseFloat(strings.TrimSpace(pesoTxt), 64)
			if err != nil {
				return obj, fmt.Errorf("peso invalido em %q: %w", parte, err)
			}
			peso = p
		}
		obj.Pesos[stat] += peso
	}
	if len(obj.Pesos) == 0 {
		return obj, fmt.Errorf("objetivo sem stat algum")
	}
	return obj, nil
}

// Valor pontua um vetor de stats sob o objetivo.
//
// A forca adaptativa e convertida ANTES de pontuar, e nao pontuada como stat
// proprio: quem pede "maximo de dano de ataque" quer que os 9 de forca
// adaptativa do fragmento contem, porque no jogo eles contam.
func (o Objetivo) Valor(v canon.Vector) float64 {
	total := 0.0
	alvo := o.Resolucao.StatResolvido()
	for stat, qtd := range v {
		if stat == canon.AdaptiveForce {
			stat = alvo
		}
		total += o.Pesos[stat] * qtd
	}
	return total
}

// Stats devolve os stats do objetivo em ordem canonica.
func (o Objetivo) Stats() []canon.Stat { return o.Pesos.Stats() }

func (o Objetivo) String() string {
	partes := make([]string, 0, len(o.Pesos))
	for _, s := range o.Stats() {
		if p := o.Pesos[s]; p == 1 {
			partes = append(partes, string(s))
		} else {
			partes = append(partes, fmt.Sprintf("%s:%g", s, p))
		}
	}
	return strings.Join(partes, ",")
}

func statConhecido(s canon.Stat) bool {
	for _, c := range canon.All {
		if c == s {
			return true
		}
	}
	return false
}

func listaDeStats() string {
	out := make([]string, 0, len(canon.All))
	for _, s := range canon.All {
		out = append(out, string(s))
	}
	sort.Strings(out)
	return strings.Join(out, ", ")
}
