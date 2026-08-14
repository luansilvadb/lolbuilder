// Package canonical monta o modelo canonico a partir de um snapshot.
//
// E a camada onde a semantica mora: o snapshot guarda o que a fonte serviu, o
// filter recorta o que pertence ao modo, e aqui se decide o que cada coisa
// significa — notadamente o que e compravel, que nenhum campo isolado da fonte
// responde.
package canonical

import "github.com/luansilvadb/lolbuilder/internal/canon"

// Dataset e o modelo canonico de um patch.
type Dataset struct {
	Patch string `json:"patch"`

	Items          []Item          `json:"items"`
	Runes          []Rune          `json:"runes"`
	RuneStyles     []RuneStyle     `json:"rune_styles"`
	Champions      []Champion      `json:"champions"`
	SummonerSpells []SummonerSpell `json:"summoner_spells"`

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

// Habilidade e uma habilidade de slot ou a passiva, so em texto.
type Habilidade struct {
	Slot      string `json:"slot"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao,omitempty"`
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
