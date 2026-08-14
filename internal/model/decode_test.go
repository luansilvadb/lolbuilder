package model

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestDecodeStrictRejeitaCampoNovo e o teste que protege a razao de existir do
// pacote: campo novo no CDragon precisa virar erro, e nao dado descartado.
func TestDecodeStrictRejeitaCampoNovo(t *testing.T) {
	raw := []byte(`[{"id":1001,"name":"Boots","description":"","active":true,
		"inStore":true,"from":[],"to":[],"categories":[],"maxStacks":0,
		"requiredChampion":"","requiredAlly":"","requiredBuffCurrencyName":"",
		"requiredBuffCurrencyCost":0,"specialRecipe":0,"isEnchantment":false,
		"price":300,"priceTotal":300,"displayInItemSets":true,"iconPath":"x.png",
		"campoQueNaoExistia":42}]`)

	if _, err := DecodeStrict[[]Item]("items.json", raw); err == nil {
		t.Fatal("campo desconhecido passou em silencio — o alarme de campo novo esta desligado")
	} else if !strings.Contains(err.Error(), "campoQueNaoExistia") {
		t.Fatalf("o erro nao nomeia o campo novo, entao nao diz o que mapear: %v", err)
	}
}

func TestDecodeStrictAceitaCampoConhecido(t *testing.T) {
	raw := []byte(`[{"id":1001,"name":"Boots","description":"d","active":true,
		"inStore":true,"from":[],"to":[1006],"categories":["Boots"],"maxStacks":0,
		"requiredChampion":"","requiredAlly":"","requiredBuffCurrencyName":"",
		"requiredBuffCurrencyCost":0,"specialRecipe":0,"isEnchantment":false,
		"price":300,"priceTotal":300,"displayInItemSets":true,"iconPath":"x.png"}]`)

	items, err := DecodeStrict[[]Item]("items.json", raw)
	if err != nil {
		t.Fatalf("decodificacao valida falhou: %v", err)
	}
	if len(items) != 1 || items[0].ID != 1001 || items[0].PriceTotal != 300 {
		t.Fatalf("decodificou errado: %+v", items)
	}
}

// TestIgnoredEngoleSubarvore garante que a valvula de escape funciona: o que
// esta sob um Ignored nao dispara o alarme, e os bytes sobrevivem intactos.
func TestIgnoredEngoleSubarvore(t *testing.T) {
	type comIgnorado struct {
		ID  int32   `json:"id"`
		Sub Ignored `json:"sub"`
	}
	raw := []byte(`{"id":7,"sub":{"qualquer":1,"coisa":[{"aninhada":true}]}}`)

	got, err := DecodeStrict[comIgnorado]("teste.json", raw)
	if err != nil {
		t.Fatalf("Ignored nao engoliu a subarvore: %v", err)
	}
	if got.ID != 7 {
		t.Fatalf("id decodificado errado: %d", got.ID)
	}
	if !json.Valid(got.Sub) {
		t.Fatalf("Ignored nao preservou JSON valido: %q", string(got.Sub))
	}
	if !strings.Contains(string(got.Sub), "aninhada") {
		t.Fatalf("Ignored perdeu os bytes originais: %q", string(got.Sub))
	}
}

// TestIgnoredNaoEsconderCampoIrmao confere que o escape e local: um campo novo
// FORA do Ignored continua sendo erro.
func TestIgnoredNaoEsconderCampoIrmao(t *testing.T) {
	type comIgnorado struct {
		ID  int32   `json:"id"`
		Sub Ignored `json:"sub"`
	}
	raw := []byte(`{"id":7,"sub":{"x":1},"novo":2}`)

	if _, err := DecodeStrict[comIgnorado]("teste.json", raw); err == nil {
		t.Fatal("campo novo irmao de um Ignored passou — o escape vazou para fora da subarvore")
	}
}

func TestIgnoredSerializaDeVolta(t *testing.T) {
	var i Ignored
	if err := i.UnmarshalJSON([]byte(`{"a":1}`)); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `{"a":1}` {
		t.Fatalf("serializou %q, esperado {\"a\":1}", string(out))
	}

	var vazio Ignored
	out, err = json.Marshal(vazio)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "null" {
		t.Fatalf("Ignored vazio serializou %q, esperado null", string(out))
	}
}
