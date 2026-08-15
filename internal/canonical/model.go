// Package canonical monta o modelo canonico a partir de um snapshot.
//
// E a camada onde a semantica mora: o snapshot guarda o que a fonte serviu, o
// filter recorta o que pertence ao modo, e aqui se decide o que cada coisa
// significa — notadamente o que e compravel, que nenhum campo isolado da fonte
// responde.
package canonical

import (
	"github.com/luansilvadb/lolbuilder/internal/canon"
	"github.com/luansilvadb/lolbuilder/internal/gamedata"
)

// Dataset e o modelo canonico de um patch.
type Dataset struct {
	Patch string `json:"patch"`

	Items          []Item          `json:"items"`
	Runes          []Rune          `json:"runes"`
	RuneStyles     []RuneStyle     `json:"rune_styles"`
	Champions      []Champion      `json:"champions"`
	SummonerSpells []SummonerSpell `json:"summoner_spells"`

	Computed Computed `json:"computed"`
	Coverage Coverage `json:"coverage"`
}

// Item e um item do catalogo do modo.
//
// Nome vem do locale de exibicao e NomeCanonico do canonico. Os dois sao
// publicados lado a lado porque o consumidor pergunta em portugues, mas todo o
// corpus externo — notas de patch, guias, discussao — usa o nome em ingles.
type Item struct {
	ID           int32  `json:"id"`
	Nome         string `json:"nome"`
	NomeCanonico string `json:"nome_canonico"`

	// Custo e o total acumulado; Combina e o que se paga a partir dos
	// componentes. Os nomes seguem o dataset original.
	Custo       int32   `json:"custo"`
	Combina     int32   `json:"combina"`
	Componentes []int32 `json:"componentes,omitempty"`

	Stats  canon.Vector `json:"stats,omitempty"`
	Efeito string       `json:"efeito,omitempty"`

	// Compravel e a decisao central deste pacote: o item e referenciado pelas
	// listas de loja do modo E declarado na loja pela fonte.
	//
	// Nenhum dos dois sozinho serve. No 16.16 sao 333 itens com InStore
	// verdadeiro que o modo nao referencia — de ARAM e Arena — e 55 que o modo
	// referencia mas a fonte nao poe na loja, que sao buffs de torre,
	// marcadores de estrutura e placeholders.
	Compravel bool `json:"compravel"`

	// Botas marca o item como calcado, contando os aprimoramentos. O jogo nao
	// deixa carregar dois, e o otimizador de build precisa da restricao — sem
	// ela ele proporia uma build impossivel com cara de otima.
	Botas bool `json:"botas,omitempty"`

	// categoriaBotas guarda a etiqueta DIRETA da fonte, antes da propagacao
	// pela arvore de componentes. Ver marcarBotas.
	categoriaBotas bool
}

// Rune e uma runa ou fragmento de stat do sistema Runes Reforged.
//
// Sem classificacao de escopo ainda: a curadoria entra no M4. Aqui esta o que a
// fonte publica.
type Rune struct {
	ID           int32  `json:"id"`
	Nome         string `json:"nome"`
	NomeCanonico string `json:"nome_canonico"`

	Resumo    string `json:"resumo"`
	Descricao string `json:"descricao"`

	// PatchDaUltimaMudanca e o gatilho de revisao da curadoria. A fonte publica
	// esse campo por runa, o que dispensa comparar o texto entre patches.
	PatchDaUltimaMudanca string `json:"patch_da_ultima_mudanca"`

	// Estilo e o caminho a que a runa pertence. Zero em fragmento de stat, que
	// aparece nos cinco — atribuir um fragmento a Precision so porque ela vem
	// primeiro na lista seria inventar informacao.
	Estilo   int32  `json:"estilo,omitempty"`
	TipoSlot string `json:"tipo_slot,omitempty"`

	// LinhasSlot sao TODAS as linhas em que a runa aparece, e nao a primeira.
	//
	// Keystone e runa menor ocupam uma linha so. Fragmento de stat pode ocupar
	// duas: no 16.16, Forca Adaptativa esta na linha 4 e na 5, e Vida Escalavel
	// na 5 e na 6. Guardar so a primeira faria um otimizador nunca colocar
	// Forca Adaptativa no slot flexivel — perdendo opcao real e publicando
	// pagina subotima com cara de exata, que e o pior erro possivel num sistema
	// que promete exatidao.
	LinhasSlot []int `json:"linhas_slot,omitempty"`

	// Escopo e a classificacao curada: se a runa soma atributo e, quando nao
	// soma, por que.
	//
	// Sai no dataset porque e o que explica ao leitor a ausencia de 58 das 69
	// runas do pre-calculo. Sem isso o arquivo lista as runas e deixa a
	// ausencia sem resposta, que e o mesmo que nao declarar limite nenhum.
	Escopo string `json:"escopo,omitempty"`
	// MotivoDoEscopo diz por que a runa fica fora do calculo.
	MotivoDoEscopo string `json:"motivo_do_escopo,omitempty"`
	// RessalvaDoEscopo registra a parte do efeito que ficou de fora quando a
	// runa e so parcialmente somavel.
	RessalvaDoEscopo string `json:"ressalva_do_escopo,omitempty"`
	// StatsDaRuna e o que ela concede no nivel 18, quando soma.
	StatsDaRuna canon.Vector `json:"stats,omitempty"`
}

// Fragmento informa se a runa e um fragmento de stat.
func (r Rune) Fragmento() bool { return r.TipoSlot == SlotStatMod }

// Keystone informa se a runa e uma pedra fundamental.
func (r Rune) Keystone() bool { return r.TipoSlot == SlotKeyStone }

// Os tipos de slot que perkstyles.json publica.
const (
	SlotKeyStone = "kKeyStone"
	SlotMenor    = "kMixedRegularSplashable"
	SlotStatMod  = "kStatMod"
)

// RuneStyle e um dos cinco caminhos de runa.
//
// A estrutura vem publicada pela fonte, ao contrario das maestrias do modo
// Jade, que precisavam de constantes curadas a mao. E daqui que o otimizador
// enumera as paginas possiveis.
type RuneStyle struct {
	ID           int32  `json:"id"`
	Nome         string `json:"nome"`
	NomeCanonico string `json:"nome_canonico"`

	SubEstilosPermitidos []int32 `json:"subestilos_permitidos"`
	Linhas               []Linha `json:"linhas"`

	// BonusPorSubEstilo e a runa concedida por escolher determinado estilo
	// secundario.
	BonusPorSubEstilo map[int32]int32 `json:"bonus_por_subestilo,omitempty"`
}

// Linha e um slot do caminho, com as runas que ele oferece.
type Linha struct {
	Tipo  string  `json:"tipo"`
	Runas []int32 `json:"runas"`
}

// Champion e um campeao do modo.
//
// Estatistica base e formula de habilidade NAO estao aqui: vem do dump de dados
// do jogo e entram no M3. O que ha aqui e o que o plugin publica.
type Champion struct {
	ID           int32  `json:"id"`
	Nome         string `json:"nome"`
	NomeCanonico string `json:"nome_canonico"`
	Alias        string `json:"alias"`
	Titulo       string `json:"titulo"`

	Papeis []string `json:"papeis,omitempty"`

	// CorpoACorpo sai de tacticalInfo.attackType. Importa porque varias runas
	// rendem valores diferentes para melee e ranged.
	CorpoACorpo bool   `json:"corpo_a_corpo"`
	TipoDeDano  string `json:"tipo_de_dano,omitempty"`

	// Stats vem do dump de dados do jogo, nunca do plugin.
	Stats *gamedata.Stats `json:"stats,omitempty"`

	Passiva     Habilidade   `json:"passiva"`
	Habilidades []Habilidade `json:"habilidades"`

	// SubHabilidades sao as habilidades que um livro de feiticos acrescenta.
	//
	// No 16.16 so Hwei tem: tres grupos de quatro, com slots qq, qw, qe, qr e
	// assim por diante. Sao 12 habilidades de verdade, com dano e recarga
	// proprios — publicar Hwei so com as quatro do slot principal descreveria
	// um campeao que nao existe.
	SubHabilidades []Habilidade `json:"sub_habilidades,omitempty"`
}

// Habilidade e uma habilidade de slot ou a passiva.
//
// Nome e Descricao vem do plugin, em portugues. Todo o resto vem do dump de
// dados do jogo: o plugin publica esses campos presentes e ZERADOS, e publicar
// o zero dele afirmaria que a habilidade nao causa dano em vez de dizer que o
// dano e desconhecido.
type Habilidade struct {
	Slot      string `json:"slot"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao,omitempty"`

	// NomeInterno e como o dump chama o objeto da habilidade. Serve para casar
	// uma lacuna reportada com a entidade que a produziu.
	NomeInterno string `json:"nome_interno,omitempty"`

	Recarga []float64 `json:"recarga,omitempty"`
	Custo   []float64 `json:"custo,omitempty"`

	// Alcance sai vazio quando a fonte publica valor de sentinela.
	//
	// A banda legitima do 16.16 termina em 7500, nos ultimates globais, e ha uma
	// lacuna limpa ate 10000, onde comeca a sentinela: 37 habilidades em 10000,
	// 132 em 25000, e casos ate 200000. Publicar "alcance 25000" para o Q do
	// Galio seria pior que nao publicar alcance nenhum. Ver alcanceMaximoPlausivel.
	Alcance []float64 `json:"alcance,omitempty"`

	Efeitos []Efeito `json:"efeitos,omitempty"`

	// SeriesNomeadas sao numeros que a fonte publica e que nenhuma formula da
	// habilidade consome — duracao de lentidao, raio, recarga interna.
	//
	// Sem elas, boa parte das habilidades de utilidade sairia sem numero
	// nenhum. As series JA consumidas ficam de fora, para nao apresentar o
	// mesmo dano duas vezes como se fossem parcelas somaveis.
	SeriesNomeadas []SerieNomeada `json:"series_nomeadas,omitempty"`

	// Series do plugin, guardadas so para o cruzamento de alinhamento de rank.
	recargaDoPlugin []float64
	custoDoPlugin   []float64
}

// Efeito e uma formula nomeada da habilidade, resolvida por rank.
type Efeito struct {
	Nome string `json:"nome"`

	// PorRank fica vazio quando a formula nao resolve. Efeito sem numero nunca
	// e publicado com zero: zero afirma que nao causa dano, ausencia so nao
	// informa, e so a segunda e recuperavel pelo leitor.
	PorRank []Expressao `json:"por_rank,omitempty"`

	// NaoResolvido diz por que a formula ficou sem numero.
	NaoResolvido string `json:"nao_resolvido,omitempty"`

	// NivelDeReferencia so aparece quando o efeito MUDA com o nivel do campeao.
	//
	// Sao 407 parcelas assim no 16.16. Sem esse campo, quem le um numero toma
	// por constante o que na verdade e o valor no nivel 18 — no nivel 1 varios
	// deles valem menos de um quarto disso.
	NivelDeReferencia int `json:"nivel_de_referencia,omitempty"`

	// DerivadoDe rotula um calculo cujo nome a fonte nao publica, transcrevendo
	// a relacao que ela declara: "TotalDPS x 0.25" nao inventa nome nenhum.
	DerivadoDe    string  `json:"derivado_de,omitempty"`
	Multiplicador float64 `json:"multiplicador,omitempty"`
}

// Expressao e o efeito em funcao das estatisticas do campeao.
//
// Nao e um numero: o dataset publica quanto a habilidade causa DADAS as
// estatisticas, e quem monta a build faz a conta. Simular combate esta fora do
// escopo.
type Expressao struct {
	Fixo    float64  `json:"fixo"`
	Escalas []Escala `json:"escalas,omitempty"`
}

// Escala e uma parcela que cresce com uma estatistica.
type Escala struct {
	Stat        string  `json:"stat"`
	Parcela     string  `json:"parcela"` // total, base ou bonus
	Coeficiente float64 `json:"coeficiente"`
}

// SerieNomeada e uma serie de valores por rank que a fonte publica sob um nome.
type SerieNomeada struct {
	Nome    string    `json:"nome"`
	PorRank []float64 `json:"por_rank"`
}

// SummonerSpell e um feitico de invocador valido no modo.
type SummonerSpell struct {
	ID           int64  `json:"id"`
	Nome         string `json:"nome"`
	NomeCanonico string `json:"nome_canonico"`
	Descricao    string `json:"descricao,omitempty"`
	Recarga      int32  `json:"recarga"`
	NivelMinimo  int32  `json:"nivel_minimo"`
}

// Coverage resume o que o build conseguiu ler.
type Coverage struct {
	// Itens mede o subconjunto PUBLICADO — os compraveis. E a taxa que responde
	// "quanto do que o consumidor vai ler foi lido por inteiro".
	Itens canon.CoberturaDeItens `json:"itens"`

	// VocabularioForaDaLoja sao as linhas que o parser nao soube ler em itens
	// do catalogo do modo que NAO estao na loja.
	//
	// Nao degrada a taxa publicada, e por isso e uma lista e nao um percentual.
	// Existe porque medir vocabulario so sobre o publicado esconde a forma nova
	// ate o patch em que ela chega a loja — e ai a correcao vira urgencia em vez
	// de manutencao. Foi assim que forca adaptativa e <ornnBonus> passaram
	// despercebidos com a cobertura marcando 100%.
	VocabularioForaDaLoja []canon.LinhaNaoLida `json:"vocabulario_fora_da_loja,omitempty"`

	Campeoes    CoberturaDeCampeoes    `json:"campeoes"`
	Alinhamento []gamedata.AlignReport `json:"alinhamento,omitempty"`
}

// CoberturaDeCampeoes resume a extracao do dump de dados do jogo.
type CoberturaDeCampeoes struct {
	CampeoesTotal    int `json:"campeoes_total"`
	CampeoesComStats int `json:"campeoes_com_stats"`

	Habilidades    CoberturaDeEntidade `json:"habilidades"`
	Passivas       CoberturaDeEntidade `json:"passivas"`
	SubHabilidades CoberturaDeEntidade `json:"sub_habilidades"`

	// AlcancesDescartados conta as habilidades cujo alcance a fonte publica como
	// sentinela. Nao e defeito de leitura: e informacao que a fonte nao da.
	AlcancesDescartados int `json:"alcances_descartados"`

	// LacunasDeStat lista, uma a uma, as estatisticas que a fonte nao publica.
	// Sao publicadas como ausentes, nunca como zero.
	LacunasDeStat []string `json:"lacunas_de_stat,omitempty"`
}

// CoberturaDeEntidade separa resolvida por inteiro, parcial e sem formula.
//
// Sao tres categorias e nao duas porque a fracao SEM FORMULA precisa de contador
// proprio: a taxa de resolucao sozinha nao a detecta, e uma taxa calibrada para
// caber nas lacunas de hoje cabe tambem no dobro delas.
type CoberturaDeEntidade struct {
	Total         int `json:"total"`
	Resolvidas    int `json:"resolvidas"`
	Parciais      int `json:"parciais"`
	NaoResolvidas int `json:"nao_resolvidas"`
	SemFormula    int `json:"sem_formula"`
}

// Purchasable devolve so os itens compraveis, preservando a ordem.
func (d *Dataset) Purchasable() []Item {
	out := make([]Item, 0, 256)
	for _, it := range d.Items {
		if it.Compravel {
			out = append(out, it)
		}
	}
	return out
}
