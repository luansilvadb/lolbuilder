package optimize

import (
	"fmt"
	"sort"

	"github.com/luansilvadb/lolbuilder/internal/canon"
)

// A pagina de runas do LoL moderno tem forma fixa, publicada pela fonte:
//
//	primario:    1 keystone, escolhida entre 3 ou 4
//	             3 runas menores, uma por linha, 3 opcoes cada
//	secundario:  2 runas menores, de linhas DIFERENTES, de um estilo permitido
//	fragmentos:  3, um por linha de kStatMod, 3 opcoes cada
//
// O espaco bruto e da ordem de 1,5 milhao de paginas. Enumera-lo seria viavel,
// mas desnecessario: com objetivo linear, cada escolha e independente das
// outras DADO o par de estilos. Por isso o otimo sai escolhendo o melhor de
// cada slot, e o par de estilos e o unico eixo que precisa ser varrido — sao 20
// combinacoes.
//
// O atalho e verificado contra busca exaustiva num teste, sobre o catalogo
// real. Sem essa verificacao, um erro de raciocinio aqui publicaria pagina
// subotima com cara de exata.

// Catalogo e o que o otimizador precisa saber do dataset.
type Catalogo struct {
	Estilos []Estilo
	// Curadas indexa a semantica curada por id de runa.
	Curadas map[int32]*canon.CuratedRune
	// Nomes indexa o nome publicado, para a saida.
	Nomes map[int32]string
}

// Estilo e um caminho de runa com suas linhas.
type Estilo struct {
	ID                   int32
	Nome                 string
	SubEstilosPermitidos []int32
	// Keystones e as opcoes da primeira linha.
	Keystones []int32
	// Menores sao as tres linhas de runa menor, na ordem da fonte.
	Menores [][]int32
	// Fragmentos sao as tres linhas de fragmento de stat.
	Fragmentos [][]int32
}

// Escolha e uma runa escolhida, com o que ela concede.
type Escolha struct {
	ID    int32        `json:"id"`
	Nome  string       `json:"nome"`
	Stats canon.Vector `json:"stats,omitempty"`
}

// Pagina e uma pagina de runas completa.
type Pagina struct {
	EstiloPrimario   int32  `json:"estilo_primario"`
	NomePrimario     string `json:"nome_primario"`
	EstiloSecundario int32  `json:"estilo_secundario"`
	NomeSecundario   string `json:"nome_secundario"`

	Keystone    Escolha   `json:"keystone"`
	Menores     []Escolha `json:"menores"`
	Secundarias []Escolha `json:"secundarias"`
	Fragmentos  []Escolha `json:"fragmentos"`

	Total canon.Vector `json:"total"`
	Valor float64      `json:"valor"`

	// Nivel e Resolucao ficam no resultado porque mudam o numero. Publicar o
	// total sem eles daria um valor que so vale sob premissas invisiveis.
	Nivel     int       `json:"nivel"`
	Resolucao Resolucao `json:"resolucao_adaptativa"`
}

// MelhorPagina devolve a pagina de valor maximo sob o objetivo.
func MelhorPagina(cat Catalogo, obj Objetivo, nivel int) (Pagina, error) {
	if len(cat.Estilos) == 0 {
		return Pagina{}, fmt.Errorf("catalogo sem estilo algum")
	}
	porID := make(map[int32]Estilo, len(cat.Estilos))
	for _, e := range cat.Estilos {
		porID[e.ID] = e
	}

	var melhor Pagina
	achou := false

	for _, prim := range cat.Estilos {
		for _, subID := range prim.SubEstilosPermitidos {
			sec, ok := porID[subID]
			if !ok {
				return Pagina{}, fmt.Errorf(
					"estilo %d permite o secundario %d, que nao esta no catalogo", prim.ID, subID)
			}
			p := cat.melhorComEstilos(prim, sec, obj, nivel)
			if !achou || p.Valor > melhor.Valor || (p.Valor == melhor.Valor && desempate(p, melhor)) {
				melhor, achou = p, true
			}
		}
	}
	return melhor, nil
}

// desempate mantem a saida estavel quando duas paginas valem o mesmo.
//
// Sem isto, a ordem de iteracao de mapa decidiria qual pagina o dataset
// publica, e o arquivo mudaria entre execucoes sem que nada tivesse mudado.
func desempate(a, b Pagina) bool {
	if a.EstiloPrimario != b.EstiloPrimario {
		return a.EstiloPrimario < b.EstiloPrimario
	}
	return a.EstiloSecundario < b.EstiloSecundario
}

// melhorComEstilos resolve a pagina otima para um par de estilos fixo.
//
// Com os estilos fixos, cada slot e independente: a keystone nao restringe a
// runa menor, e nenhuma linha oferece a mesma runa que outra. Por isso o maximo
// da soma e a soma dos maximos.
func (cat Catalogo) melhorComEstilos(prim, sec Estilo, obj Objetivo, nivel int) Pagina {
	p := Pagina{
		EstiloPrimario:   prim.ID,
		NomePrimario:     prim.Nome,
		EstiloSecundario: sec.ID,
		NomeSecundario:   sec.Nome,
		Total:            canon.Vector{},
		Nivel:            nivel,
		Resolucao:        obj.Resolucao,
	}

	p.Keystone = cat.melhorDaLinha(prim.Keystones, obj, nivel)
	for _, linha := range prim.Menores {
		p.Menores = append(p.Menores, cat.melhorDaLinha(linha, obj, nivel))
	}

	// O secundario da DUAS runas, de linhas diferentes. O melhor par e formado
	// pelos melhores de cada linha: como as linhas sao disjuntas, escolher a
	// segunda melhor de uma linha nunca supera a melhor de outra.
	candidatos := make([]Escolha, 0, len(sec.Menores))
	for _, linha := range sec.Menores {
		candidatos = append(candidatos, cat.melhorDaLinha(linha, obj, nivel))
	}
	sort.SliceStable(candidatos, func(i, j int) bool {
		vi, vj := obj.Valor(candidatos[i].Stats), obj.Valor(candidatos[j].Stats)
		if vi != vj {
			return vi > vj
		}
		return candidatos[i].ID < candidatos[j].ID
	})
	if len(candidatos) > 2 {
		candidatos = candidatos[:2]
	}
	sort.Slice(candidatos, func(i, j int) bool { return candidatos[i].ID < candidatos[j].ID })
	p.Secundarias = candidatos

	// Os fragmentos vem do primario, mas sao os mesmos nos cinco estilos.
	for _, linha := range prim.Fragmentos {
		p.Fragmentos = append(p.Fragmentos, cat.melhorDaLinha(linha, obj, nivel))
	}

	for _, e := range append(append(append([]Escolha{p.Keystone}, p.Menores...), p.Secundarias...), p.Fragmentos...) {
		p.Total.Merge(e.Stats)
	}
	p.Valor = obj.Valor(p.Total)
	return p
}

// melhorDaLinha escolhe a runa de maior valor numa linha.
//
// Empate resolve pelo menor id, e nao pela ordem da fonte: a ordem da fonte
// pode mudar entre patches sem que nada de fato tenha mudado, e ai o arquivo
// publicado mudaria junto.
func (cat Catalogo) melhorDaLinha(linha []int32, obj Objetivo, nivel int) Escolha {
	var melhor Escolha
	melhorValor := 0.0
	achou := false

	for _, id := range linha {
		e := Escolha{ID: id, Nome: cat.Nomes[id], Stats: cat.vetor(id, nivel)}
		v := obj.Valor(e.Stats)
		if !achou || v > melhorValor || (v == melhorValor && id < melhor.ID) {
			melhor, melhorValor, achou = e, v, true
		}
	}
	return melhor
}

// vetor devolve o que a runa concede no nivel, ou vazio se ela nao soma.
func (cat Catalogo) vetor(id int32, nivel int) canon.Vector {
	cur, ok := cat.Curadas[id]
	if !ok {
		return canon.Vector{}
	}
	return cur.VetorNoNivel(nivel)
}
