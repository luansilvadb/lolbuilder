package canon

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Algumas passivas convertem uma estatistica em outra.
//
// Isto importa para UMA coisa so: o oraculo em partida. A Live Client Data API
// reporta o valor TOTAL de cada estatistica, ja com o que a passiva concedeu.
// O dataset publica base e crescimento por nivel, que e outra grandeza. Sem
// saber da conversao, o oraculo acusa divergencia onde os dois estao certos.
//
// Foi assim que o Rammus apareceu: o crescimento de dano de ataque do nivel 1 ao
// 6 veio 3.8809 acima do previsto, numa leitura sem item nenhum. O excesso e
// exatamente 0.15 da armadura mais 0.15 da resistencia magica que ele ganhou no
// caminho — o Casco Espetado, que o dataset ja publicava certo.
//
// A tabela e curada, e nao derivada da formula publicada, porque a fonte NAO
// distingue "passiva que concede atributo" de "passiva que causa dano escalando
// naquele atributo". As duas saem com a mesma forma. Derivar automaticamente
// inventaria bonus inexistente para toda passiva do segundo tipo, e inventar
// bonus e pior que nao conhecer nenhum: o oraculo passaria a confirmar o que
// deveria contestar.

// Conversao e uma passiva que alimenta uma estatistica a partir de outras.
type Conversao struct {
	// Campeao e o nome canonico, como o dataset o guarda.
	Campeao string `json:"campeao"`
	// Eixo e a estatistica que RECEBE.
	Eixo Stat `json:"eixo"`
	// Origem sao as estatisticas que alimentam, e o coeficiente de cada uma.
	Origem map[Stat]float64 `json:"origem"`

	// Habilidade e o slot onde a conversao mora, para quem for conferir no
	// 06-champion-stats.md.
	Habilidade string `json:"habilidade"`

	// Evidencia e obrigatoria: uma conversao sem medicao que a sustente e um
	// palpite, e um palpite aqui faz o oraculo confirmar dataset errado. A tabela
	// existe justamente para ser mais confiavel que a formula publicada.
	Evidencia string `json:"evidencia"`
	// Motivo descreve o efeito em uma linha.
	Motivo string `json:"motivo"`
}

// CuradoriaDeConversoes e o arquivo completo.
type CuradoriaDeConversoes struct {
	PatchDaCuradoria string      `json:"patch_da_curadoria"`
	Conversoes       []Conversao `json:"conversoes"`
}

// LoadConversoes le a tabela de conversoes.
//
// Arquivo ausente NAO e erro: a tabela e opcional, e sem ela o oraculo se
// comporta como antes. Erro seria exigi-la de quem so quer rodar o build.
func LoadConversoes(path string) (*CuradoriaDeConversoes, error) {
	raw, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &CuradoriaDeConversoes{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("lendo curadoria de conversoes em %s: %w", path, err)
	}
	var c CuradoriaDeConversoes
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("parseando %s: %w", path, err)
	}
	for i := range c.Conversoes {
		if err := c.Conversoes[i].validate(); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
	}
	return &c, nil
}

func (c *Conversao) validate() error {
	if c.Campeao == "" {
		return fmt.Errorf("ha uma conversao sem campeao")
	}
	if c.Eixo == "" {
		return fmt.Errorf("conversao de %s nao diz qual eixo recebe", c.Campeao)
	}
	if len(c.Origem) == 0 {
		return fmt.Errorf("conversao de %s nao declara origem alguma", c.Campeao)
	}
	if c.Evidencia == "" {
		return fmt.Errorf(
			"conversao de %s nao tem evidencia — sem medicao que a sustente ela e "+
				"palpite, e palpite aqui faz o oraculo confirmar dataset errado", c.Campeao)
	}
	if err := Vector(c.Origem).Validate(); err != nil {
		return fmt.Errorf("conversao de %s: %w", c.Campeao, err)
	}
	for _, s := range All {
		if s == c.Eixo {
			return nil
		}
	}
	return fmt.Errorf("conversao de %s aponta o eixo %q, fora do vocabulario",
		c.Campeao, c.Eixo)
}

// Do devolve as conversoes que alimentam um eixo de um campeao.
//
// Recebe todos os nomes conhecidos do campeao — traduzido, canonico e alias —
// porque a curadoria e escrita a mao e nao ha razao para exigir de quem a
// escreve qual dos tres o dataset guarda em qual campo.
func (c *CuradoriaDeConversoes) Do(nomes []string, eixo Stat) []Conversao {
	var out []Conversao
	for _, cv := range c.Conversoes {
		if cv.Eixo != eixo {
			continue
		}
		for _, n := range nomes {
			if n != "" && strings.EqualFold(n, cv.Campeao) {
				out = append(out, cv)
				break
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Habilidade < out[j].Habilidade })
	return out
}
