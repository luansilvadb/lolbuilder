package optimize

import (
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
)

// catalogoDeTeste reproduz a forma real: dois estilos, keystone com 3 opcoes,
// tres linhas menores de 3, e tres linhas de fragmento.
func catalogoDeTeste() Catalogo {
	soma := func(s canon.Stat, v float64) *canon.CuratedRune {
		return &canon.CuratedRune{Kind: canon.KindSum, Stats: map[canon.Stat]float64{s: v}}
	}
	fora := func() *canon.CuratedRune {
		return &canon.CuratedRune{Kind: canon.KindOutOfScope, Reason: "condicional"}
	}

	cur := map[int32]*canon.CuratedRune{
		// Estilo 100: keystones
		1001: fora(), 1002: soma(canon.Armor, 5), 1003: fora(),
		// menores do 100
		1101: soma(canon.Armor, 3), 1102: fora(), 1103: soma(canon.Health, 40),
		1201: fora(), 1202: soma(canon.Armor, 7), 1203: soma(canon.Armor, 2),
		1301: soma(canon.Health, 10), 1302: fora(), 1303: fora(),
		// Estilo 200
		2001: soma(canon.Armor, 1), 2002: fora(), 2003: fora(),
		2101: soma(canon.Armor, 9), 2102: fora(), 2103: fora(),
		2201: soma(canon.Armor, 4), 2202: fora(), 2203: fora(),
		2301: soma(canon.Armor, 6), 2302: fora(), 2303: fora(),
		// fragmentos, iguais nos dois estilos
		5001: soma(canon.Armor, 2), 5002: soma(canon.Health, 65), 5003: fora(),
		5004: soma(canon.Armor, 8), 5005: fora(), 5006: soma(canon.Health, 15),
		5007: soma(canon.Armor, 1), 5008: soma(canon.Armor, 3), 5009: fora(),
	}
	nomes := map[int32]string{}
	for id := range cur {
		nomes[id] = "runa"
	}

	frag := [][]int32{{5001, 5002, 5003}, {5004, 5005, 5006}, {5007, 5008, 5009}}
	return Catalogo{
		Curadas: cur, Nomes: nomes,
		Estilos: []Estilo{
			{ID: 100, Nome: "Cem", SubEstilosPermitidos: []int32{200},
				Keystones:  []int32{1001, 1002, 1003},
				Menores:    [][]int32{{1101, 1102, 1103}, {1201, 1202, 1203}, {1301, 1302, 1303}},
				Fragmentos: frag},
			{ID: 200, Nome: "Duzentos", SubEstilosPermitidos: []int32{100},
				Keystones:  []int32{2001, 2002, 2003},
				Menores:    [][]int32{{2101, 2102, 2103}, {2201, 2202, 2203}, {2301, 2302, 2303}},
				Fragmentos: frag},
		},
	}
}

// exaustiva enumera TODAS as paginas validas e devolve o maior valor.
//
// E deliberadamente burra e lenta. Ela existe para conferir o atalho de
// MelhorPagina, que escolhe o melhor de cada slot em vez de enumerar — se o
// raciocinio por tras do atalho estiver errado, o dataset publica pagina
// subotima com cara de exata, e nada no dado denuncia.
func exaustiva(cat Catalogo, obj Objetivo, nivel int) float64 {
	porID := map[int32]Estilo{}
	for _, e := range cat.Estilos {
		porID[e.ID] = e
	}
	melhor := 0.0
	primeiro := true

	for _, prim := range cat.Estilos {
		for _, subID := range prim.SubEstilosPermitidos {
			sec := porID[subID]
			for _, ks := range prim.Keystones {
				for _, m1 := range prim.Menores[0] {
					for _, m2 := range prim.Menores[1] {
						for _, m3 := range prim.Menores[2] {
							// Duas runas do secundario, de linhas diferentes.
							for a := 0; a < len(sec.Menores); a++ {
								for b := a + 1; b < len(sec.Menores); b++ {
									for _, s1 := range sec.Menores[a] {
										for _, s2 := range sec.Menores[b] {
											for _, f1 := range prim.Fragmentos[0] {
												for _, f2 := range prim.Fragmentos[1] {
													for _, f3 := range prim.Fragmentos[2] {
														v := canon.Vector{}
														for _, id := range []int32{ks, m1, m2, m3, s1, s2, f1, f2, f3} {
															v.Merge(cat.vetor(id, nivel))
														}
														if val := obj.Valor(v); primeiro || val > melhor {
															melhor, primeiro = val, false
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return melhor
}

// TestMelhorPaginaBateComBuscaExaustiva e a verificacao que sustenta a promessa
// de exatidao. O atalho e o otimo, e nao uma aproximacao.
func TestMelhorPaginaBateComBuscaExaustiva(t *testing.T) {
	cat := catalogoDeTeste()
	objetivos := []string{
		"armor",
		"health",
		"armor:1,health:0.1",
		"health:1,armor:20",
	}
	for _, texto := range objetivos {
		t.Run(texto, func(t *testing.T) {
			obj, err := ParseObjetivo(texto, ResolucaoAD)
			if err != nil {
				t.Fatal(err)
			}
			p, err := MelhorPagina(cat, obj, 18)
			if err != nil {
				t.Fatal(err)
			}
			if want := exaustiva(cat, obj, 18); p.Valor != want {
				t.Fatalf("MelhorPagina deu %v, a busca exaustiva achou %v", p.Valor, want)
			}
		})
	}
}

func TestPaginaTemAFormaCerta(t *testing.T) {
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	p, err := MelhorPagina(catalogoDeTeste(), obj, 18)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Menores) != 3 {
		t.Errorf("menores = %d, esperado 3", len(p.Menores))
	}
	if len(p.Secundarias) != 2 {
		t.Errorf("secundarias = %d, esperado 2", len(p.Secundarias))
	}
	if len(p.Fragmentos) != 3 {
		t.Errorf("fragmentos = %d, esperado 3", len(p.Fragmentos))
	}
	if p.EstiloPrimario == p.EstiloSecundario {
		t.Error("o secundario ficou igual ao primario")
	}
	if p.Nivel != 18 || p.Resolucao != ResolucaoAD {
		t.Errorf("o resultado nao carrega as premissas: nivel=%d resolucao=%q", p.Nivel, p.Resolucao)
	}
}

// TestSecundariasDeLinhasDiferentes: a regra do jogo proibe as duas runas
// secundarias na mesma linha, e ignora-la produziria pagina impossivel.
func TestSecundariasDeLinhasDiferentes(t *testing.T) {
	cat := catalogoDeTeste()
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	p, err := MelhorPagina(cat, obj, 18)
	if err != nil {
		t.Fatal(err)
	}

	var sec Estilo
	for _, e := range cat.Estilos {
		if e.ID == p.EstiloSecundario {
			sec = e
		}
	}
	linhaDe := func(id int32) int {
		for i, linha := range sec.Menores {
			for _, x := range linha {
				if x == id {
					return i
				}
			}
		}
		return -1
	}
	a, b := linhaDe(p.Secundarias[0].ID), linhaDe(p.Secundarias[1].ID)
	if a < 0 || b < 0 {
		t.Fatalf("runa secundaria fora do estilo secundario: %+v", p.Secundarias)
	}
	if a == b {
		t.Fatalf("as duas secundarias vieram da linha %d", a)
	}
}

// TestNivelMudaOResultado: um fragmento que escala com o nivel so vence no
// nivel alto, e publicar o resultado sem declarar o nivel esconderia isso.
func TestNivelMudaOResultado(t *testing.T) {
	cat := catalogoDeTeste()
	// Fragmento que vale 1 no nivel 1 e cresce 10 por nivel.
	cat.Curadas[5001] = &canon.CuratedRune{
		Kind:          canon.KindSumPerLevel,
		Stats:         map[canon.Stat]float64{canon.Armor: 1},
		StatsPorNivel: map[canon.Stat]float64{canon.Armor: 10},
	}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)

	baixo, _ := MelhorPagina(cat, obj, 1)
	alto, _ := MelhorPagina(cat, obj, 18)
	if !(alto.Valor > baixo.Valor) {
		t.Fatalf("o nivel nao mudou o resultado: %v no 1 e %v no 18", baixo.Valor, alto.Valor)
	}
}

// TestForcaAdaptativaResolveNoObjetivo: quem pede dano de ataque quer que os 9
// de forca adaptativa do fragmento contem, porque no jogo eles contam.
func TestForcaAdaptativaResolveNoObjetivo(t *testing.T) {
	cat := catalogoDeTeste()
	cat.Curadas[5003] = &canon.CuratedRune{
		Kind:  canon.KindSum,
		Stats: map[canon.Stat]float64{canon.AdaptiveForce: 9},
	}

	ad, _ := ParseObjetivo("attack_damage", ResolucaoAD)
	pAD, _ := MelhorPagina(cat, ad, 18)
	if pAD.Valor != 9 {
		t.Errorf("sob resolucao ad, forca adaptativa valeu %v, esperado 9", pAD.Valor)
	}

	ap, _ := ParseObjetivo("attack_damage", ResolucaoAP)
	pAP, _ := MelhorPagina(cat, ap, 18)
	if pAP.Valor != 0 {
		t.Errorf("sob resolucao ap, forca adaptativa contou %v para dano de ataque", pAP.Valor)
	}
}

func TestEstiloSecundarioAusenteEErro(t *testing.T) {
	cat := catalogoDeTeste()
	cat.Estilos[0].SubEstilosPermitidos = []int32{999}
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	if _, err := MelhorPagina(cat, obj, 18); err == nil {
		t.Fatal("estilo secundario inexistente passou em silencio")
	}
}

func TestCatalogoVazioEErro(t *testing.T) {
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	if _, err := MelhorPagina(Catalogo{}, obj, 18); err == nil {
		t.Fatal("catalogo vazio passou")
	}
}

// TestSlotSemContribuicaoEIndiferente e o teste do defeito que a revisao do M5
// encontrou: 49 das 81 escolhas publicadas nas paginas do 16.16 nao somavam
// nada ao objetivo, e saiam com nome de runa. As nove paginas mostravam
// "pedra fundamental: Pressione o Ataque" — o menor id, e nao o melhor.
func TestSlotSemContribuicaoEIndiferente(t *testing.T) {
	cat := catalogoDeTeste()
	// Objetivo que NENHUMA runa do catalogo atende.
	obj, err := ParseObjetivo("lethality", ResolucaoAD)
	if err != nil {
		t.Fatal(err)
	}
	p, err := MelhorPagina(cat, obj, 18)
	if err != nil {
		t.Fatal(err)
	}

	todas := append(append(append([]Escolha{p.Keystone}, p.Menores...), p.Secundarias...), p.Fragmentos...)
	for _, e := range todas {
		if !e.Indiferente {
			t.Errorf("slot sem contribuicao saiu como %q (id %d) em vez de indiferente", e.Nome, e.ID)
		}
		if e.Nome != "" || e.ID != 0 {
			t.Errorf("slot indiferente carregou runa: %+v", e)
		}
	}
}

// TestSlotComContribuicaoNaoEIndiferente: a marca so vale quando de fato nao ha
// escolha melhor, senao ela esconderia o resultado.
func TestSlotComContribuicaoNaoEIndiferente(t *testing.T) {
	obj, _ := ParseObjetivo("armor", ResolucaoAD)
	p, err := MelhorPagina(catalogoDeTeste(), obj, 18)
	if err != nil {
		t.Fatal(err)
	}
	if p.Keystone.Indiferente || p.Keystone.Nome == "" {
		t.Errorf("keystone que soma armadura saiu como indiferente: %+v", p.Keystone)
	}
	if p.Valor <= 0 {
		t.Errorf("a pagina perdeu valor: %v", p.Valor)
	}
}
