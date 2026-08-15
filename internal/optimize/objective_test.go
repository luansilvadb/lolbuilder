package optimize

import (
	"strings"
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
)

func TestParseObjetivo(t *testing.T) {
	obj, err := ParseObjetivo("armor", ResolucaoAD)
	if err != nil {
		t.Fatal(err)
	}
	if obj.Pesos[canon.Armor] != 1 {
		t.Errorf("peso padrao = %v, esperado 1", obj.Pesos[canon.Armor])
	}

	obj, err = ParseObjetivo("attack_damage:1,attack_speed_pct:100", ResolucaoAP)
	if err != nil {
		t.Fatal(err)
	}
	if obj.Pesos[canon.AttackDamage] != 1 || obj.Pesos[canon.AttackSpeedPct] != 100 {
		t.Errorf("pesos = %+v", obj.Pesos)
	}
	if obj.Resolucao != ResolucaoAP {
		t.Errorf("resolucao = %q", obj.Resolucao)
	}
}

// TestForcaAdaptativaNaoEAlvo: ela SEMPRE resolve para outro stat antes de valer
// alguma coisa, entao aceitar o pedido devolveria zero em silencio.
func TestForcaAdaptativaNaoEAlvo(t *testing.T) {
	_, err := ParseObjetivo("adaptive_force", ResolucaoAD)
	if err == nil {
		t.Fatal("objetivo em forca adaptativa foi aceito")
	}
	if !strings.Contains(err.Error(), "attack_damage") {
		t.Fatalf("o erro nao diz o que pedir no lugar: %v", err)
	}
}

func TestObjetivoInvalido(t *testing.T) {
	casos := map[string]string{
		"vazio":             "",
		"stat inventado":    "grandeza_que_nao_existe",
		"peso nao numerico": "armor:muito",
	}
	for nome, s := range casos {
		t.Run(nome, func(t *testing.T) {
			if _, err := ParseObjetivo(s, ResolucaoAD); err == nil {
				t.Fatal("objetivo invalido foi aceito")
			}
		})
	}
	if _, err := ParseObjetivo("armor", "talvez"); err == nil {
		t.Fatal("resolucao invalida foi aceita")
	}
}

// TestValorConverteAdaptativa: quem pede dano de ataque quer que a forca
// adaptativa conte, porque no jogo ela conta.
func TestValorConverteAdaptativa(t *testing.T) {
	v := canon.Vector{canon.AdaptiveForce: 9, canon.Armor: 5}

	ad, _ := ParseObjetivo("attack_damage", ResolucaoAD)
	if got := ad.Valor(v); got != 9 {
		t.Errorf("sob ad, valor = %v, esperado 9", got)
	}
	ap, _ := ParseObjetivo("ability_power", ResolucaoAP)
	if got := ap.Valor(v); got != 9 {
		t.Errorf("sob ap, valor = %v, esperado 9", got)
	}
	cruzado, _ := ParseObjetivo("ability_power", ResolucaoAD)
	if got := cruzado.Valor(v); got != 0 {
		t.Errorf("sob ad, forca adaptativa contou %v para poder de habilidade", got)
	}
}

func TestObjetivoString(t *testing.T) {
	obj, _ := ParseObjetivo("armor,health:2", ResolucaoAD)
	if got := obj.String(); got != "armor,health:2" {
		t.Fatalf("String() = %q", got)
	}
}
