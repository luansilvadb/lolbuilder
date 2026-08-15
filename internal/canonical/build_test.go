package canonical

import (
	"strings"
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
	ds, err := NewBuilder(cfg, "testdata/16.99", "testdata/curation").Build("16.99")
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
		if len(c.Habilidades) != 4 {
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
	if len(ds.RuneStyles) != 2 {
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
	if len(c.Habilidades) != 4 {
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

// TestEstatisticasVemDoDump: o plugin publica esses campos presentes e ZERADOS.
// Publicar o zero dele afirmaria que o campeao nao tem vida.
func TestEstatisticasVemDoDump(t *testing.T) {
	c := fixture(t).Champions[0]
	if c.Stats == nil {
		t.Fatal("o campeao saiu sem estatisticas")
	}
	if c.Stats.HP.Base != 600 || c.Stats.HP.PerLevel != 98 {
		t.Errorf("vida = %+v", c.Stats.HP)
	}
	if got := c.Stats.HP.At(18); got != 600+98*17 {
		t.Errorf("vida no nivel 18 = %v", got)
	}
	// Alcance 625 no dump: a distancia decide melee, e nao o attackType do
	// plugin.
	if c.CorpoACorpo {
		t.Error("alcance 625 foi classificado como corpo a corpo")
	}
}

// TestFormulaResolvidaPorRank confere o caminho inteiro: dump -> avaliador ->
// modelo publicado.
func TestFormulaResolvidaPorRank(t *testing.T) {
	q := fixture(t).Champions[0].Habilidades[0]

	if len(q.Recarga) != 5 || q.Recarga[0] != 8 {
		t.Errorf("recarga = %v", q.Recarga)
	}
	if len(q.Custo) != 5 || q.Custo[0] != 55 || q.Custo[4] != 95 {
		t.Errorf("custo = %v", q.Custo)
	}
	if len(q.Efeitos) != 1 || q.Efeitos[0].Nome != "TotalDamage" {
		t.Fatalf("efeitos = %+v", q.Efeitos)
	}
	e := q.Efeitos[0]
	if len(e.PorRank) != 5 {
		t.Fatalf("a formula resolveu em %d ranks, esperado 5", len(e.PorRank))
	}
	if e.PorRank[0].Fixo != 30 || e.PorRank[4].Fixo != 150 {
		t.Errorf("parcela fixa por rank errada: %v e %v", e.PorRank[0].Fixo, e.PorRank[4].Fixo)
	}
	if len(e.PorRank[0].Escalas) != 1 || e.PorRank[0].Escalas[0].Stat != "ability_power" {
		t.Errorf("escala errada: %+v", e.PorRank[0].Escalas)
	}
}

// TestFormulaNaoResolvidaNaoPublicaZero e a regra central do projeto: dano zero
// mente, ausencia so nao informa, e so a segunda e recuperavel pelo leitor.
func TestFormulaNaoResolvidaNaoPublicaZero(t *testing.T) {
	r := fixture(t).Champions[0].Habilidades[3]
	if len(r.Efeitos) != 1 {
		t.Fatalf("efeitos = %+v", r.Efeitos)
	}
	if len(r.Efeitos[0].PorRank) != 0 {
		t.Fatalf("formula sem valor honesto publicou numero: %+v", r.Efeitos[0].PorRank)
	}
	if r.Efeitos[0].NaoResolvido == "" {
		t.Error("a formula ficou sem numero e sem motivo declarado")
	}
}

// TestSeriesNaoConsumidasSaem: BaseDamage e consumida por TotalDamage e nao pode
// aparecer de novo; SlowDuration nao e consumida por formula nenhuma e sem ela a
// habilidade sairia sem esse numero.
func TestSeriesNaoConsumidasSaem(t *testing.T) {
	q := fixture(t).Champions[0].Habilidades[0]
	nomes := map[string]bool{}
	for _, s := range q.SeriesNomeadas {
		nomes[s.Nome] = true
	}
	if !nomes["SlowDuration"] {
		t.Errorf("serie nao consumida sumiu: %+v", q.SeriesNomeadas)
	}
	if nomes["BaseDamage"] {
		t.Errorf("serie ja consumida foi publicada de novo: %+v", q.SeriesNomeadas)
	}
}

// TestLacunaDeStatEReportadaPorNome: Hwei na fixture nao tem regeneracao base.
func TestLacunaDeStatEReportada(t *testing.T) {
	cov := fixture(t).Coverage.Campeoes
	if cov.CampeoesTotal != 2 || cov.CampeoesComStats != 2 {
		t.Errorf("cobertura de estatisticas = %+v", cov)
	}
	if len(cov.LacunasDeStat) != 1 || !strings.Contains(cov.LacunasDeStat[0], "Hwei") {
		t.Fatalf("a lacuna nao foi reportada por nome: %v", cov.LacunasDeStat)
	}
}

// TestAlinhamentoDeRank cruza as series do dump com as do plugin. Sem esse
// cruzamento, um erro de indexacao publicaria TODO valor por rank deslocado, em
// silencio e sem sintoma no dado.
func TestAlinhamentoDeRank(t *testing.T) {
	rel := fixture(t).Coverage.Alinhamento
	if len(rel) != 2 {
		t.Fatalf("relatorios de alinhamento = %d, esperado cooldown e mana", len(rel))
	}
	for _, r := range rel {
		if !r.OK() {
			t.Errorf("serie %s: %v", r.Series, r.Err())
		}
		if r.Agreement() != 100 {
			t.Errorf("serie %s concordou %.1f%%, esperado 100", r.Series, r.Agreement())
		}
	}
}

// TestBotasSaoTransitivas e o teste do defeito que o dado real revelou: a fonte
// categoriza Gunmetal Greaves como AttackSpeed, LifeSteal e NonbootsMovement,
// sem Boots, apesar de ela evoluir de Berserker's Greaves. Sem a propagacao
// pela arvore de componentes, o otimizador montaria uma build com dois pares de
// botas — impossivel no jogo, e com cara de otima.
func TestBotasSaoTransitivas(t *testing.T) {
	itens := []Item{
		{ID: 1001, categoriaBotas: true},       // botas base
		{ID: 3006, Componentes: []int32{1001}}, // aprimoramento, sem etiqueta
		{ID: 3172, Componentes: []int32{3006}}, // aprimoramento do aprimoramento
		{ID: 3075, Componentes: []int32{1029}}, // nada a ver com calcado
		{ID: 1029},                             // componente comum
	}
	marcarBotas(itens)

	want := map[int32]bool{1001: true, 3006: true, 3172: true, 3075: false, 1029: false}
	for _, it := range itens {
		if it.Botas != want[it.ID] {
			t.Errorf("item %d botas = %v, esperado %v", it.ID, it.Botas, want[it.ID])
		}
	}
}

// TestBotasNaoTravaEmCiclo: a arvore de componentes vem da fonte, e uma
// referencia circular travaria o build em vez de acusar.
func TestBotasNaoTravaEmCiclo(t *testing.T) {
	itens := []Item{
		{ID: 1, Componentes: []int32{2}},
		{ID: 2, Componentes: []int32{1}},
	}
	marcarBotas(itens)
	for _, it := range itens {
		if it.Botas {
			t.Errorf("ciclo sem calcado marcou o item %d como botas", it.ID)
		}
	}
}

// TestCuradoriaChegaAoModelo: sem isto o arquivo de runas listaria as 69 sem
// explicar por que a maioria nao aparece no pre-calculo.
func TestCuradoriaChegaAoModelo(t *testing.T) {
	ds := fixture(t)
	porID := map[int32]Rune{}
	for _, r := range ds.Runes {
		porID[r.ID] = r
	}

	fora := porID[8005]
	if fora.Escopo != "out_of_scope" || fora.MotivoDoEscopo == "" {
		t.Errorf("runa fora do calculo saiu sem escopo ou sem motivo: %+v", fora)
	}
	if len(fora.StatsDaRuna) != 0 {
		t.Errorf("runa fora do calculo publicou stats: %v", fora.StatsDaRuna)
	}

	frag := porID[5001]
	if frag.Escopo != "sum_per_level" {
		t.Errorf("escopo do fragmento = %q", frag.Escopo)
	}
	// 10 no nivel 1 mais 10 por nivel: 180 no nivel 18.
	if got := frag.StatsDaRuna[canon.Health]; got != 180 {
		t.Errorf("stats da runa no nivel de pre-calculo = %v, esperado 180", got)
	}
}
