package canon

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// A semantica das runas e curada a mao, nao parseada.
//
// O plugin nao publica bloco estruturado de atributo para runa: os numeros
// existem so dentro do texto, e o texto mistura valor fixo, valor por nivel,
// multiplicador e condicao de combate na mesma frase. Um parser generico
// acertaria a maioria e erraria em silencio justamente nos casos compostos.
//
// O que e automatizado e a DETECCAO. A fonte publica majorChangePatchVersion
// por runa, entao o build sabe exatamente qual runa mudou desde a curadoria e
// precisa de nova revisao — sem precisar comparar texto, que muda por reescrita
// sem mudar numero.

// RuneKind classifica o efeito de uma runa quanto a somabilidade.
type RuneKind string

const (
	// KindSum soma direto no vetor de stats.
	KindSum RuneKind = "sum"
	// KindSumPerLevel soma um valor que cresce linearmente com o nivel.
	KindSumPerLevel RuneKind = "sum_per_level"
	// KindMultiplier multiplica um stat existente. Fica fora do objetivo linear
	// porque quebra a separabilidade, mas e registrado.
	KindMultiplier RuneKind = "multiplier"
	// KindOutOfScope depende de estado de partida ou nao tem stat
	// correspondente no vocabulario.
	KindOutOfScope RuneKind = "out_of_scope"
)

// Degrau e um ganho que so existe a partir de certo nivel do campeao.
//
// Existe porque algumas runas dao o atributo em degraus, e nao linearmente:
// Transcendence concede aceleracao de habilidade no nivel 5 e de novo no 8.
// Sem isto ela cairia em out_of_scope, e o dataset perderia um atributo real e
// incondicional so por nao saber expressa-lo.
type Degrau struct {
	Nivel int              `json:"nivel"`
	Stats map[Stat]float64 `json:"stats"`
}

// CuratedRune e a entrada curada de uma runa.
type CuratedRune struct {
	Name string   `json:"name"`
	Kind RuneKind `json:"kind"`

	// Stats e o ganho incondicional.
	Stats map[Stat]float64 `json:"stats,omitempty"`
	// StatsPorNivel e o ganho que cresce linearmente, por nivel acima do 1.
	StatsPorNivel map[Stat]float64 `json:"stats_por_nivel,omitempty"`
	// Degraus sao ganhos que aparecem a partir de niveis especificos.
	Degraus []Degrau `json:"degraus,omitempty"`

	// Reason e obrigatoria em out_of_scope: uma runa fora do calculo sem
	// motivo escrito e indistinguivel de uma runa que alguem esqueceu de curar.
	Reason string `json:"reason,omitempty"`
	// Note registra a parte do efeito que ficou de fora quando a runa e so
	// parcialmente somavel.
	Note string `json:"note,omitempty"`

	// PatchDaCuradoria e o majorChangePatchVersion da fonte no momento em que a
	// entrada foi escrita. Se a fonte passar a declarar outro, a runa mudou e
	// precisa de revisao.
	PatchDaCuradoria string `json:"patch_da_curadoria"`
}

// RuneCuration e o arquivo de curadoria completo.
type RuneCuration struct {
	CuratedFromPatch string                  `json:"curated_from_patch"`
	Runes            map[string]*CuratedRune `json:"runes"`
}

// LoadRuneCuration le a tabela de curadoria de runas.
func LoadRuneCuration(path string) (*RuneCuration, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("lendo curadoria de runas em %s: %w", path, err)
	}
	var c RuneCuration
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("parseando %s: %w", path, err)
	}
	if len(c.Runes) == 0 {
		return nil, fmt.Errorf("curadoria de runas em %s esta vazia", path)
	}
	for id, r := range c.Runes {
		if err := r.validate(id); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
	}
	return &c, nil
}

func (r *CuratedRune) validate(id string) error {
	switch r.Kind {
	case KindSum, KindSumPerLevel, KindMultiplier, KindOutOfScope:
	default:
		return fmt.Errorf("runa %s tem kind %q, fora do vocabulario", id, r.Kind)
	}

	if r.Kind == KindOutOfScope && r.Reason == "" {
		return fmt.Errorf(
			"runa %s esta fora do calculo sem motivo escrito — sem ele, ela fica "+
				"indistinguivel de uma runa que alguem esqueceu de curar", id)
	}
	if r.Kind != KindOutOfScope && len(r.Stats) == 0 &&
		len(r.StatsPorNivel) == 0 && len(r.Degraus) == 0 {
		return fmt.Errorf("runa %s e somavel mas nao declara stat algum", id)
	}
	for _, m := range []map[Stat]float64{r.Stats, r.StatsPorNivel} {
		if err := Vector(m).Validate(); err != nil {
			return fmt.Errorf("runa %s: %w", id, err)
		}
	}
	for _, d := range r.Degraus {
		if d.Nivel < 1 || d.Nivel > 18 {
			return fmt.Errorf("runa %s tem degrau no nivel %d, fora de 1..18", id, d.Nivel)
		}
		if err := Vector(d.Stats).Validate(); err != nil {
			return fmt.Errorf("runa %s, degrau do nivel %d: %w", id, d.Nivel, err)
		}
	}
	return nil
}

// Somavel informa se a runa entra no objetivo linear.
func (r *CuratedRune) Somavel() bool {
	return r.Kind == KindSum || r.Kind == KindSumPerLevel
}

// VetorNoNivel devolve o que a runa concede no nivel pedido.
func (r *CuratedRune) VetorNoNivel(nivel int) Vector {
	out := Vector{}
	if !r.Somavel() {
		return out
	}
	for s, v := range r.Stats {
		out.Add(s, v)
	}
	for s, v := range r.StatsPorNivel {
		out.Add(s, v*float64(nivel-1))
	}
	for _, d := range r.Degraus {
		if nivel >= d.Nivel {
			for s, v := range d.Stats {
				out.Add(s, v)
			}
		}
	}
	return out
}

// Get devolve a entrada curada de um id.
func (c *RuneCuration) Get(id int32) (*CuratedRune, bool) {
	r, ok := c.Runes[fmt.Sprint(id)]
	return r, ok
}

// RunaDesatualizada aponta uma runa cujo patch de mudanca avancou desde a
// curadoria.
type RunaDesatualizada struct {
	ID     int32
	Nome   string
	Curado string
	Atual  string
}

// RunaNaoCurada aponta uma runa jogavel sem entrada na curadoria.
type RunaNaoCurada struct {
	ID   int32
	Nome string
}

// Conferir compara a curadoria com o catalogo publicado.
//
// Devolve as runas jogaveis sem entrada e as que mudaram desde a curadoria. As
// duas listas abortam o build: sob curadoria parcial, uma runa nova que soma
// stat entra como zero silencioso, que e o modo de falha que este mecanismo
// existe para tornar impossivel.
func (c *RuneCuration) Conferir(jogaveis []RunaPublicada) (naoCuradas []RunaNaoCurada, desatualizadas []RunaDesatualizada) {
	for _, r := range jogaveis {
		cur, ok := c.Get(r.ID)
		if !ok {
			naoCuradas = append(naoCuradas, RunaNaoCurada{ID: r.ID, Nome: r.Nome})
			continue
		}
		if cur.PatchDaCuradoria != r.PatchDaUltimaMudanca {
			desatualizadas = append(desatualizadas, RunaDesatualizada{
				ID: r.ID, Nome: r.Nome,
				Curado: cur.PatchDaCuradoria, Atual: r.PatchDaUltimaMudanca,
			})
		}
	}
	sort.Slice(naoCuradas, func(i, j int) bool { return naoCuradas[i].ID < naoCuradas[j].ID })
	sort.Slice(desatualizadas, func(i, j int) bool { return desatualizadas[i].ID < desatualizadas[j].ID })
	return naoCuradas, desatualizadas
}

// RunaPublicada e o minimo que a conferencia precisa saber de uma runa.
//
// E uma struct propria, e nao o tipo do modelo canonico, para que o pacote
// canon nao dependa do canonical — a dependencia correria no sentido errado.
type RunaPublicada struct {
	ID                   int32
	Nome                 string
	PatchDaUltimaMudanca string
}
