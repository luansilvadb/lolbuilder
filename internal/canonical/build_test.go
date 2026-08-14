package canonical

import (
	"testing"

	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/config"
	"github.com/luansilvadb/lolbuilder/internal/model"
)

// O snapshot de fixture em testdata/16.99 e minusculo de proposito, mas cobre
// todos os casos que a decisao de "compravel" precisa distinguir:
//
//	1001    referenciado pelo modo E inStore  -> compravel
//	3006    referenciado por OUTRA lista do modo, inStore -> compravel
//	1515    referenciado, mas inStore=false (buff de torre) -> nao
//	2052    inStore, mas o modo nao referencia (item de ARAM) -> nao
//	3865    compravel, e a fonte nao publica bloco de atributos
//	771001  referenciado, mas fora da faixa de id do modo (Jade) -> nem entra
func fixture(t *testing.T) *Dataset {
	t.Helper()
	cfg := &config.Config{
		LocaleCanonical: "default",
		LocaleDisplay:   "pt_br",
		Mode: config.Mode{
			Name: "classic", GameMode: "CLASSIC", MapID: 11,
			ItemIDMin: 1000, ItemIDMax: 699999,
			ChampionIDMin: 1, ChampionIDMax: 999,
		},
	}
	ds, err := NewBuilder(cfg, "testdata/16.99").Build("16.99")
	if err != nil {
		t.Fatalf("build da fixture falhou: %v", err)
	}
	return ds
}

func item(t *testing.T, ds *Dataset, id int32) Item {
	t.Helper()
	for _, it := range ds.Items {
		if it.ID == id {
			return it
		}
	}
	t.Fatalf("item %d nao esta no dataset", id)
	return Item{}
}

// TestCompravelExigeAsDuasCondicoes e o teste da decisao central do pacote.
func TestCompravelExigeAsDuasCondicoes(t *testing.T) {
	ds := fixture(t)

	casos := map[int32]bool{
		1001: true,  // referenciado pelo modo e inStore
		3006: true,  // referenciado por outra lista do modo, inStore
		3865: true,  // referenciado e inStore, sem bloco de atributos
		1515: false, // referenciado, mas a fonte nao poe na loja
		2052: false, // inStore, mas o modo nao referencia
	}
	for id, want := range casos {
		if got := item(t, ds, id).Compravel; got != want {
			t.Errorf("item %d compravel = %v, esperado %v", id, got, want)
		}
	}
	if n := len(ds.Purchasable()); n != 3 {
		t.Errorf("Purchasable() devolveu %d itens, esperado 3", n)
	}
}

// TestFaixaDeIDExcluiOOutroJogo: 771001 e referenciado pela loja do modo, mas
// mora na faixa do modo Jade. Nao pode entrar no catalogo.
func TestFaixaDeIDExcluiOOutroJogo(t *testing.T) {
	ds := fixture(t)
	for _, it := range ds.Items {
		if it.ID == 771001 {
			t.Fatal("item do modo Jade entrou no catalogo do LoL moderno")
		}
	}
	if len(ds.Items) != 5 {
		t.Fatalf("catalogo tem %d itens, esperado 5", len(ds.Items))
	}
}

func TestItemBilingueEStats(t *testing.T) {
	it := item(t, fixture(t), 3006)

	if it.Nome != "Grevas do Berserker" || it.NomeCanonico != "Berserker's Greaves" {
		t.Errorf("nomes errados: %q / %q", it.Nome, it.NomeCanonico)
	}
	if it.Custo != 1100 || it.Combina != 800 {
		t.Errorf("custos errados: total=%d combina=%d", it.Custo, it.Combina)
	}
	if len(it.Componentes) != 2 || it.Componentes[0] != 1001 {
		t.Errorf("componentes errados: %v", it.Componentes)
	}
	// O deslocamento das botas esta DENTRO do bloco de atributos no LoL moderno.
	want := canon.Vector{canon.AttackSpeedPct: 30, canon.MoveSpeed: 45}
	for s, v := range want {
		if it.Stats[s] != v {
			t.Errorf("stat %s = %v, esperado %v", s, it.Stats[s], v)
		}
	}
	if len(it.Stats) != len(want) {
		t.Errorf("stats = %+v, esperado %+v", it.Stats, want)
	}
}

// TestEfeitoNaoRepeteOsStats: o bloco de atributos ja virou vetor, e repeti-lo
// no texto daria ao consumidor duas fontes para o mesmo numero.
func TestEfeitoNaoRepeteOsStats(t *testing.T) {
	it := item(t, fixture(t), 3006)
	if it.Efeito != "Passiva Anda depressa." {
		t.Fatalf("efeito = %q", it.Efeito)
	}
}

func TestCoberturaSoContaOQueEPublicado(t *testing.T) {
	c := fixture(t).Coverage.Itens

	// Dos 3 compraveis, 2 tem bloco e os 2 sao lidos por inteiro. World Atlas
	// nao tem bloco e fica fora do denominador.
	if c.ComBloco != 2 || c.LidosPorInteiro != 2 {
		t.Errorf("cobertura = %+v", c)
	}
	if c.Taxa() != 100 {
		t.Errorf("taxa = %v, esperado 100", c.Taxa())
	}
	if c.SemBloco != 1 || len(c.SemBlocoNome) != 1 || c.SemBlocoNome[0] != "World Atlas" {
		t.Errorf("item sem bloco nao foi reportado por nome: %+v", c.SemBlocoNome)
	}
	// Os itens nao compraveis nao entram: 1515 tem bloco e seria lido, 2052
	// tambem, e nenhum dos dois aparece no arquivo publicado.
	if c.ComBloco+c.SemBloco != 3 {
		t.Errorf("item fora da loja entrou na cobertura: %+v", c)
	}
}

func TestRunasSituadasNaPagina(t *testing.T) {
	ds := fixture(t)
	porID := map[int32]Rune{}
	for _, r := range ds.Runes {
		porID[r.ID] = r
	}

	key := porID[8005]
	if !key.Keystone() || key.Estilo != 8000 {
		t.Errorf("keystone situada errado: %+v", key)
	}
	if len(key.LinhasSlot) != 1 || key.LinhasSlot[0] != 0 {
		t.Errorf("linhas da keystone = %v, esperado [0]", key.LinhasSlot)
	}
	if key.Nome != "Pressione o Ataque" || key.NomeCanonico != "Press the Attack" {
		t.Errorf("nomes de runa errados: %q / %q", key.Nome, key.NomeCanonico)
	}
	if key.Resumo != "Atinge tres vezes" {
		t.Errorf("resumo nao foi limpo de marcacao: %q", key.Resumo)
	}
	if key.PatchDaUltimaMudanca != "14.10" {
		t.Errorf("gatilho de revisao perdido: %q", key.PatchDaUltimaMudanca)
	}

	// Fragmento aparece em todos os estilos, entao nao pertence a nenhum.
	frag := porID[5008]
	if !frag.Fragmento() {
		t.Errorf("fragmento nao foi reconhecido: %+v", frag)
	}
	if frag.Estilo != 0 {
		t.Errorf("fragmento foi atribuido ao estilo %d; ele vale em todos", frag.Estilo)
	}
}

// TestFragmentoGuardaTodasAsLinhas e o teste do defeito encontrado na revisao do
// M2: Forca Adaptativa ocupa as linhas 1 e 2 na fixture (4 e 5 no jogo real), e
// guardar so a primeira faria um otimizador nunca coloca-la no slot flexivel.
func TestFragmentoGuardaTodasAsLinhas(t *testing.T) {
	ds := fixture(t)
	for _, r := range ds.Runes {
		if r.ID != 5008 {
			continue
		}
		if len(r.LinhasSlot) != 2 || r.LinhasSlot[0] != 1 || r.LinhasSlot[1] != 2 {
			t.Fatalf("linhas do fragmento = %v, esperado [1 2]", r.LinhasSlot)
		}
		return
	}
	t.Fatal("o fragmento 5008 sumiu do catalogo")
}

// TestHweiPublicaOLivroDeFeiticos: o livro e um segundo conjunto de habilidades
// com dano e recarga proprios. Publicar so as quatro do slot principal
// descreveria um campeao que nao existe.
func TestHweiPublicaOLivroDeFeiticos(t *testing.T) {
	ds := fixture(t)
	for _, c := range ds.Champions {
		if c.ID != 910 {
			continue
		}
		if len(c.Habilidades) != 1 {
			t.Errorf("habilidades de slot = %d", len(c.Habilidades))
		}
		if len(c.SubHabilidades) != 2 {
			t.Fatalf("sub-habilidades = %d, esperado 2: %+v", len(c.SubHabilidades), c.SubHabilidades)
		}
		if c.SubHabilidades[0].Slot != "qq" || c.SubHabilidades[0].Nome != "Fogo Devastador" {
			t.Errorf("sub-habilidade errada: %+v", c.SubHabilidades[0])
		}
		return
	}
	t.Fatal("Hwei sumiu do dataset")
}

// TestOrdemDivergenteEntreLocalesAborta: parear por posicao sem conferir o
// spellKey publicaria a descricao do W sob o nome do Q, em silencio.
func TestOrdemDivergenteEntreLocalesAborta(t *testing.T) {
	canonicas := []model.ChampionSpell{{SpellKey: "q", Name: "Q"}, {SpellKey: "w", Name: "W"}}
	trocadas := []model.ChampionSpell{{SpellKey: "w", Name: "W"}, {SpellKey: "q", Name: "Q"}}

	if _, err := parearHabilidades("champions/1.json", canonicas, trocadas); err == nil {
		t.Fatal("ordens divergentes foram pareadas em silencio")
	}
	if _, err := parearHabilidades("champions/1.json", canonicas, canonicas[:1]); err == nil {
		t.Fatal("contagens divergentes foram pareadas em silencio")
	}
}

// TestBonusDeSubEstiloNaoERunaDeSlot: 8004 e referenciado por subStyleBonus e
// nao aparece em slot algum. Ele nao pode ganhar posicao de pagina.
func TestBonusDeSubEstiloNaoERunaDeSlot(t *testing.T) {
	ds := fixture(t)
	for _, r := range ds.Runes {
		if r.ID == 8004 {
			if r.TipoSlot != "" {
				t.Fatalf("bonus de subestilo recebeu tipo de slot %q", r.TipoSlot)
			}
			return
		}
	}
	t.Fatal("a runa 8004 sumiu do catalogo")
}

func TestEstilo(t *testing.T) {
	ds := fixture(t)
	if len(ds.RuneStyles) != 1 {
		t.Fatalf("estilos = %d", len(ds.RuneStyles))
	}
	st := ds.RuneStyles[0]
	if st.Nome != "Precisao" || st.NomeCanonico != "Precision" {
		t.Errorf("nomes de estilo errados: %q / %q", st.Nome, st.NomeCanonico)
	}
	if len(st.Linhas) != 3 || st.Linhas[0].Tipo != SlotKeyStone || st.Linhas[2].Tipo != SlotStatMod {
		t.Errorf("linhas erradas: %+v", st.Linhas)
	}
	if st.BonusPorSubEstilo[8100] != 8004 {
		t.Errorf("bonus por subestilo perdido: %+v", st.BonusPorSubEstilo)
	}
}

func TestCampeoes(t *testing.T) {
	ds := fixture(t)
	if len(ds.Champions) != 2 {
		t.Fatalf("campeoes = %d; a sentinela de id -1 e o campeao do modo Jade deviam sair", len(ds.Champions))
	}
	c := ds.Champions[0]
	if c.Nome != "Annie" || c.Titulo != "a Crianca Sombria" {
		t.Errorf("campeao traduzido errado: %q / %q", c.Nome, c.Titulo)
	}
	if c.CorpoACorpo {
		t.Error("Annie e ranged e foi marcada como corpo a corpo")
	}
	if c.Passiva.Nome != "Piromania" || c.Passiva.Descricao != "Atordoa o alvo" {
		t.Errorf("passiva errada: %+v", c.Passiva)
	}
	if len(c.Habilidades) != 1 {
		t.Fatalf("habilidades = %d", len(c.Habilidades))
	}
	q := c.Habilidades[0]
	if q.Slot != "q" || q.Nome != "Desintegrar" {
		t.Errorf("habilidade errada: %+v", q)
	}
	// A marcacao sai, mas o numero que ela envolve fica.
	if q.Descricao != "Causa 60 de Dano Magico" {
		t.Errorf("descricao = %q", q.Descricao)
	}
}

func TestFeiticosRecortadosPeloModo(t *testing.T) {
	ds := fixture(t)
	if len(ds.SummonerSpells) != 1 {
		t.Fatalf("feiticos = %d, esperado so o de CLASSIC: %+v", len(ds.SummonerSpells), ds.SummonerSpells)
	}
	s := ds.SummonerSpells[0]
	if s.ID != 4 || s.Nome != "Flash" || s.Recarga != 300 || s.NivelMinimo != 7 {
		t.Errorf("feitico errado: %+v", s)
	}
}
