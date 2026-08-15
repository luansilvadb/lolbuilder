package canon

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func escreverRunas(t *testing.T, conteudo string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "runes.json")
	if err := os.WriteFile(p, []byte(conteudo), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

const curadoriaDeRunas = `{
  "curated_from_patch": "16.16",
  "runes": {
    "5008": {"name": "Adaptive Force", "kind": "sum",
             "stats": {"adaptive_force": 9}, "patch_da_curadoria": ""},
    "5001": {"name": "Health Scaling", "kind": "sum_per_level",
             "stats": {"health": 10}, "stats_por_nivel": {"health": 10},
             "patch_da_curadoria": ""},
    "8210": {"name": "Transcendence", "kind": "sum",
             "degraus": [{"nivel": 5, "stats": {"ability_haste": 5}},
                         {"nivel": 8, "stats": {"ability_haste": 5}}],
             "patch_da_curadoria": "14.10"},
    "8005": {"name": "Press the Attack", "kind": "out_of_scope",
             "reason": "depende de estado de combate", "patch_da_curadoria": "14.10"}
  }
}`

func TestLoadRuneCuration(t *testing.T) {
	c, err := LoadRuneCuration(escreverRunas(t, curadoriaDeRunas))
	if err != nil {
		t.Fatal(err)
	}
	if len(c.Runes) != 4 {
		t.Fatalf("runas = %d", len(c.Runes))
	}
	r, ok := c.Get(5008)
	if !ok || !r.Somavel() {
		t.Fatalf("5008 = %+v (%v)", r, ok)
	}
	fora, _ := c.Get(8005)
	if fora.Somavel() {
		t.Error("runa out_of_scope foi tratada como somavel")
	}
}

// TestVetorNoNivelPorNivel: o fragmento de vida vale 10 no nivel 1 e 180 no 18,
// que e como a fonte declara.
func TestVetorNoNivelPorNivel(t *testing.T) {
	c, _ := LoadRuneCuration(escreverRunas(t, curadoriaDeRunas))
	r, _ := c.Get(5001)

	if got := r.VetorNoNivel(1)[Health]; got != 10 {
		t.Errorf("nivel 1 = %v, esperado 10", got)
	}
	if got := r.VetorNoNivel(18)[Health]; got != 180 {
		t.Errorf("nivel 18 = %v, esperado 180", got)
	}
}

// TestVetorNoNivelDegraus: Transcendence da +5 no nivel 5 e mais +5 no 8. Sem
// os degraus ela cairia em out_of_scope, e o dataset perderia um atributo real
// e incondicional so por nao saber expressa-lo.
func TestVetorNoNivelDegraus(t *testing.T) {
	c, _ := LoadRuneCuration(escreverRunas(t, curadoriaDeRunas))
	r, _ := c.Get(8210)

	casos := map[int]float64{1: 0, 4: 0, 5: 5, 7: 5, 8: 10, 18: 10}
	for nivel, want := range casos {
		if got := r.VetorNoNivel(nivel)[AbilityHaste]; got != want {
			t.Errorf("nivel %d = %v, esperado %v", nivel, got, want)
		}
	}
}

func TestRunaForaDoCalculoExigeMotivo(t *testing.T) {
	semMotivo := strings.Replace(curadoriaDeRunas,
		`"kind": "out_of_scope",
             "reason": "depende de estado de combate", `,
		`"kind": "out_of_scope", `, 1)
	_, err := LoadRuneCuration(escreverRunas(t, semMotivo))
	if err == nil {
		t.Fatal("runa fora do calculo sem motivo foi aceita")
	}
	if !strings.Contains(err.Error(), "esqueceu de curar") {
		t.Fatalf("o erro nao explica o risco: %v", err)
	}
}

func TestCuradoriaDeRunaInvalida(t *testing.T) {
	casos := map[string]string{
		"kind desconhecido": strings.Replace(curadoriaDeRunas, `"kind": "sum"`, `"kind": "talvez"`, 1),
		"somavel sem stat": strings.Replace(curadoriaDeRunas, `"kind": "sum",
             "stats": {"adaptive_force": 9}, `, `"kind": "sum", `, 1),
		"stat fora do vocabulario": strings.Replace(curadoriaDeRunas, `"adaptive_force": 9`, `"grandeza_inventada": 9`, 1),
		"degrau fora de 1..18":     strings.Replace(curadoriaDeRunas, `"nivel": 5`, `"nivel": 99`, 1),
		"arquivo vazio":            `{"runes": {}}`,
	}
	for nome, conteudo := range casos {
		t.Run(nome, func(t *testing.T) {
			if _, err := LoadRuneCuration(escreverRunas(t, conteudo)); err == nil {
				t.Fatal("curadoria invalida foi aceita")
			}
		})
	}
}

// TestConferirPegaRunaSemCuradoria: sob curadoria parcial, uma runa nova que
// soma stat entra no dataset como zero silencioso.
func TestConferirPegaRunaSemCuradoria(t *testing.T) {
	c, _ := LoadRuneCuration(escreverRunas(t, curadoriaDeRunas))
	naoCuradas, desatualizadas := c.Conferir([]RunaPublicada{
		{ID: 5008, Nome: "Adaptive Force", PatchDaUltimaMudanca: ""},
		{ID: 9999, Nome: "Runa Nova", PatchDaUltimaMudanca: "16.16"},
	})
	if len(desatualizadas) != 0 {
		t.Errorf("acusou desatualizada onde nao ha: %v", desatualizadas)
	}
	if len(naoCuradas) != 1 || naoCuradas[0].ID != 9999 {
		t.Fatalf("a runa nova nao foi acusada: %v", naoCuradas)
	}
}

// TestConferirPegaRunaAlterada: a fonte publica majorChangePatchVersion por
// runa, entao o gatilho de revisao nao depende de comparar texto — que muda por
// reescrita sem mudar numero.
func TestConferirPegaRunaAlterada(t *testing.T) {
	c, _ := LoadRuneCuration(escreverRunas(t, curadoriaDeRunas))
	_, desatualizadas := c.Conferir([]RunaPublicada{
		{ID: 8005, Nome: "Press the Attack", PatchDaUltimaMudanca: "16.20"},
	})
	if len(desatualizadas) != 1 {
		t.Fatalf("a runa alterada nao foi acusada: %v", desatualizadas)
	}
	d := desatualizadas[0]
	if d.Curado != "14.10" || d.Atual != "16.20" {
		t.Fatalf("o relatorio nao mostra os dois patches: %+v", d)
	}
}

func TestVetorNoNivelDeRunaForaDoCalculoEVazio(t *testing.T) {
	c, _ := LoadRuneCuration(escreverRunas(t, curadoriaDeRunas))
	r, _ := c.Get(8005)
	if len(r.VetorNoNivel(18)) != 0 {
		t.Fatalf("runa fora do calculo devolveu stats: %v", r.VetorNoNivel(18))
	}
}
