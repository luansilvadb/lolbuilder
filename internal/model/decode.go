// Package model contem as structs das fontes do CommunityDragon.
package model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// DecodeStrict decodifica com DisallowUnknownFields ativo.
//
// Um campo novo numa fonte do CDragon vira erro aqui, e nao dado ignorado em
// silencio. Isso e deliberado: um campo novo e justamente o sinal de que algo
// mudou no jogo, e perde-lo anula a vantagem que o projeto busca. O custo e ter
// de mapear o campo novo em raw.go quando o alerta aparecer.
func DecodeStrict[T any](file string, data []byte) (T, error) {
	var out T
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&out); err != nil {
		return out, fmt.Errorf("%s: %w (campo novo no CDragon? mapeie em internal/model/raw.go)", file, err)
	}
	return out, nil
}

// Ignored engole uma subarvore que a fonte publica e o projeto declara nao
// consumir, preservando os bytes originais sem validar nada dentro deles.
//
// E a valvula de escape da decodificacao estrita, e existe porque o alarme de
// campo novo precisa ficar apontado para onde importa. O plugin do LoL moderno
// publica catalogo de skins, chromas e mapas de asset por keystone — dezenas de
// campos que mudam a cada patch de cosmetico e que o dataset nunca vai ler.
// Estrito ali significaria build quebrado toda semana por motivo cosmetico, e
// build que quebra por bobagem treina o operador a contornar o alerta.
//
// O contrato: chave que nao esta nem numa struct nem sob um Ignored continua
// sendo erro. Cada uso carrega um comentario dizendo POR QUE aquilo nao e
// consumido, e `grep -rn Ignored internal/model` lista a declaracao inteira.
type Ignored []byte

// UnmarshalJSON guarda os bytes como vieram, sem interpreta-los.
func (i *Ignored) UnmarshalJSON(b []byte) error {
	*i = append((*i)[:0], b...)
	return nil
}

// MarshalJSON devolve os bytes originais.
func (i Ignored) MarshalJSON() ([]byte, error) {
	if len(i) == 0 {
		return []byte("null"), nil
	}
	return i, nil
}
