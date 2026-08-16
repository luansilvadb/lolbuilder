package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/optimize"
)

func paginaTeste(objetivo string, total canon.Vector) optimize.Pagina {
	return optimize.Pagina{
		Objetivo:       objetivo,
		NomePrimario:   "Precisão",
		NomeSecundario: "Dominação",
		Keystone:       optimize.Escolha{Indiferente: true},
		Menores:        []optimize.Escolha{{Indiferente: true}, {Indiferente: true}, {Indiferente: true}},
		Secundarias:    []optimize.Escolha{{Indiferente: true}, {Indiferente: true}},
		Fragmentos:     []optimize.Escolha{{Indiferente: true}, {Indiferente: true}, {Indiferente: true}},
		Total:          total,
		Nivel:          18,
		Resolucao:      optimize.ResolucaoAD,
	}
}

func TestParseFormatoSaida(t *testing.T) {
	for _, tc := range []struct {
		nome string
		raw  string
		want formatoSaida
	}{
		{nome: "texto", raw: "text", want: formatoTexto},
		{nome: "json", raw: "JSON", want: formatoJSON},
	} {
		t.Run(tc.nome, func(t *testing.T) {
			got, err := parseFormatoSaida(tc.raw)
			if err != nil || got != tc.want {
				t.Fatalf("parseFormatoSaida(%q) = %q, %v; want %q", tc.raw, got, err, tc.want)
			}
		})
	}
	if _, err := parseFormatoSaida("yaml"); err == nil || !strings.Contains(err.Error(), "text, json") {
		t.Fatalf("formato invalido nao explicou as opcoes: %v", err)
	}
}

func TestFormatNumberELocale(t *testing.T) {
	casos := map[string]struct {
		valor float64
		want  string
	}{
		"zero":      {valor: 0, want: "0"},
		"inteiro":   {valor: 1000, want: "1.000"},
		"decimal":   {valor: 3.5, want: "3,5"},
		"negativo":  {valor: -1234.25, want: "-1.234,25"},
		"sem zeros": {valor: 10.0, want: "10"},
	}
	for nome, tc := range casos {
		t.Run(nome, func(t *testing.T) {
			if got := formatNumber(tc.valor); got != tc.want {
				t.Fatalf("formatNumber(%v) = %q; want %q", tc.valor, got, tc.want)
			}
		})
	}
}

func TestFormatVectorOrdenaETraduz(t *testing.T) {
	v := canon.Vector{
		canon.Health:       500,
		canon.Armor:        20,
		canon.AbilityHaste: 10,
	}
	got := formatVector(v)
	want := "10 aceleração de habilidade, 20 armadura, 500 vida"
	if got != want {
		t.Fatalf("formatVector = %q; want %q", got, want)
	}
}

func TestTextoPaginaIndiferente(t *testing.T) {
	p := paginaTeste("armor", canon.Vector{})
	got := textoPagina(p)
	for _, want := range []string{
		"keystone:   indiferente",
		"contribuem: nenhuma",
		"aviso: nenhuma runa contribui para armadura",
		"slots livres: 9 de 9",
		"total: 0 armadura",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("texto da pagina nao contem %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "map[") {
		t.Fatalf("texto da pagina expoe mapa interno: %s", got)
	}
}

func TestTextoBuildDestacaObjetivoEExtras(t *testing.T) {
	b := optimize.Build{
		Rotulo:    "maximo de armor por ouro, ignorando efeitos de item",
		Total:     canon.Vector{canon.Armor: 320, canon.Health: 500, canon.Mana: 400},
		Gasto:     9750,
		Orcamento: 10000,
		Resolucao: optimize.ResolucaoAD,
	}
	got := textoBuild(b)
	for _, want := range []string{
		"máximo de armadura por ouro",
		"objetivo: 320 armadura",
		"gasto: 9.750 / 10.000 ouro (sobra 250)",
		"atributos adicionais: 500 vida, 400 mana",
		"passivas e ativas dos itens não entram no cálculo",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("texto da build nao contem %q:\n%s", want, got)
		}
	}
}

func TestJSONEnvelopePreservaContrato(t *testing.T) {
	p := paginaTeste("armor", canon.Vector{})
	env := envelopeSaida{
		Comando:             "runes",
		Patch:               "16.16",
		Objetivo:            "armor",
		ResolucaoAdaptativa: "ad",
		Resultado:           p,
		Avisos:              []string{"nenhuma runa contribui para armadura"},
	}
	var out bytes.Buffer
	if err := escreverJSONEm(&out, env); err != nil {
		t.Fatal(err)
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(out.Bytes(), &decoded); err != nil {
		t.Fatalf("JSON invalido: %v\n%s", err, out.String())
	}
	for _, campo := range []string{"comando", "patch", "objetivo", "resolucao_adaptativa", "resultado", "avisos"} {
		if _, ok := decoded[campo]; !ok {
			t.Errorf("campo %q ausente no envelope", campo)
		}
	}
	var resultado optimize.Pagina
	if err := json.Unmarshal(decoded["resultado"], &resultado); err != nil {
		t.Fatal(err)
	}
	if resultado.Objetivo != "armor" || !resultado.Keystone.Indiferente {
		t.Fatalf("resultado nao preservou os campos canonicos: %+v", resultado)
	}
	if strings.Contains(out.String(), "nenhuma runa") && !strings.Contains(out.String(), "avisos") {
		t.Fatal("aviso apareceu fora do campo avisos")
	}
}
