// Package export gera os arquivos markdown consumidos por um Project do
// claude.ai.
//
// O consumidor e um modelo de linguagem lendo arquivos que o usuario sobe a
// mao — nao ha API publica para gerenciar o conhecimento de um Project. Isso
// torna a CONTAGEM de arquivos uma restricao de primeira classe: poucos
// arquivos consolidados valem mais que uma arvore navegavel, porque cada patch
// exige refazer o upload.
package export

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/luansilvadb/lolbuilder/internal/canonical"
)

// File e um arquivo gerado, com a estimativa de tamanho.
type File struct {
	Name   string
	Body   string
	Tokens int
}

// Result e o conjunto gerado.
type Result struct {
	Files []File
	// Changelog fica FORA do conjunto do Project: e para o operador ler, e nao
	// conhecimento do modelo.
	Changelog *File
	Total     int
}

// Input reune tudo que o export precisa.
type Input struct {
	Dataset *canonical.Dataset
	// Anterior e o dataset do patch anterior, para o changelog. Nil na primeira
	// captura.
	Anterior *canonical.Dataset
	// CapturadoEm e a data da captura do snapshot.
	CapturadoEm string
}

// Generate monta todos os arquivos.
func Generate(in Input) (*Result, error) {
	if in.Dataset == nil {
		return nil, fmt.Errorf("dataset ausente")
	}
	res := &Result{}

	add := func(name, body string) {
		f := File{Name: name, Body: body, Tokens: estimateTokens(body)}
		res.Files = append(res.Files, f)
		res.Total += f.Tokens
	}

	add("01-items.md", renderItems(in.Dataset))
	add("02-runes.md", renderRunes(in.Dataset))
	add("03-summoner-spells.md", renderSpells(in.Dataset))
	add("04-champions.md", renderChampions(in.Dataset))
	add("05-computed.md", renderComputed(in.Dataset))
	add("06-champion-stats.md", renderChampionStats(in.Dataset))

	// O manifesto depende das contagens dos demais, entao e montado por ultimo
	// e inserido na frente.
	manifest := File{Name: "00-MANIFEST.md"}
	manifest.Body = renderManifest(in, res.Files)
	manifest.Tokens = estimateTokens(manifest.Body)
	res.Files = append([]File{manifest}, res.Files...)
	res.Total += manifest.Tokens

	body := renderChangelog(in)
	res.Changelog = &File{
		Name:   in.Dataset.Patch + ".md",
		Body:   body,
		Tokens: estimateTokens(body),
	}
	return res, nil
}

// Write grava o conjunto no destino.
func (r *Result) Write(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, f := range r.Files {
		p := filepath.Join(dir, f.Name)
		if err := os.WriteFile(p, []byte(f.Body), 0o644); err != nil {
			return fmt.Errorf("gravando %s: %w", p, err)
		}
	}
	return nil
}

// WriteChangelog grava o changelog do patch.
func (r *Result) WriteChangelog(dir string) error {
	if r.Changelog == nil {
		return nil
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	p := filepath.Join(dir, r.Changelog.Name)
	return os.WriteFile(p, []byte(r.Changelog.Body), 0o644)
}

// CheckOutputDir compara o conteudo do destino com o conjunto gerado e devolve
// os nomes que o export nao criou.
//
// Nada e removido: o comando escreve num diretorio que o usuario tambem
// manipula a mao, e apagar arquivo que nao criou seria inaceitavel. Avisar
// transfere a decisao para quem pode toma-la.
//
// Isto e o que faz a limpeza durar. Remover o dado antigo resolve o estado de
// hoje; esta verificacao resolve o dia em que um arquivo avulso for salvo aqui
// "so para conferir uma coisa" e esquecido — e for parar no Project junto com
// os corretos.
func (r *Result) CheckOutputDir(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil // destino inexistente ou ilegivel: nao ha o que avisar
	}
	gerados := map[string]bool{}
	for _, f := range r.Files {
		gerados[f.Name] = true
	}
	// Nao e gerado pelo export, mas pertence ao conjunto que vai ao Project.
	gerados["PROJECT-INSTRUCTIONS.md"] = true

	var estranhos []string
	for _, e := range entries {
		if !gerados[e.Name()] {
			estranhos = append(estranhos, e.Name())
		}
	}
	sort.Strings(estranhos)
	return estranhos
}

// estimateTokens aproxima a contagem de tokens por caracteres.
//
// Quatro caracteres por token e a razao usual para texto em alfabeto latino. E
// aproximacao, e nao medida — o teto existe para pegar crescimento de ordem de
// grandeza, e nao para controlar o ultimo por cento.
func estimateTokens(s string) int { return len([]rune(s)) / 4 }
