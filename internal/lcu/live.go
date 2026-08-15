// Package lcu fala com o jogo em andamento.
//
// A Live Client Data API roda em porta fixa no proprio computador, nao usa
// lockfile e nao pede autenticacao. Existe apenas enquanto ha partida, e serve
// o que o jogo esta calculando naquele instante.
//
// E o unico oraculo de fidelidade que o projeto tem. Nenhuma fonte publica o
// que o jogo faz em tempo de execucao: o dump declara base e crescimento, e se
// a formula que aplicamos sobre eles estiver errada, nada no dado denuncia. Foi
// assim que a resistencia magica ficou anos publicada como fixa no dataset do
// modo Jade.
package lcu

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// ErrSemPartida indica que nao ha partida em andamento.
//
// Distinto de falha de leitura: a porta fechada e o estado normal fora do jogo,
// e tratar isso como erro faria o comando gritar sempre que ninguem estivesse
// jogando.
var ErrSemPartida = errors.New("nenhuma partida em andamento")

const (
	// PortaLive e fixa, ao contrario da porta do cliente.
	PortaLive = 2999
	baseLive  = "https://127.0.0.1:2999/liveclientdata"
)

// LiveClient le os dados da partida em andamento.
type LiveClient struct {
	http *http.Client
	base string
}

// NewLive monta o cliente da partida.
func NewLive(timeout time.Duration) *LiveClient {
	return &LiveClient{
		base: baseLive,
		http: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				// Certificado autoassinado em 127.0.0.1. A verificacao desligada
				// vale para ESTE cliente e nao para o transporte padrao do
				// projeto — o que fala com o CommunityDragon continua estrito.
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

// EmPartida informa se ha partida em andamento.
func (c *LiveClient) EmPartida() bool {
	_, err := c.get("/gamestats")
	return err == nil
}

func (c *LiveClient) get(caminho string) ([]byte, error) {
	res, err := c.http.Get(c.base + caminho)
	if err != nil {
		return nil, fmt.Errorf("%w — entre numa partida e rode de novo (%v)", ErrSemPartida, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("a partida respondeu %s em %s", res.Status, caminho)
	}
	return io.ReadAll(res.Body)
}

// StatsDoJogo sao as estatisticas que o jogo reporta para o campeao do jogador.
//
// Os nomes sao os da API. Nao ha traducao para o vocabulario canonico aqui: o
// ponto do comando e confrontar duas leituras independentes, e renomear no
// caminho abriria espaco para o mapeamento esconder a divergencia.
type StatsDoJogo struct {
	MaxHealth       float64 `json:"maxHealth"`
	Armor           float64 `json:"armor"`
	MagicResist     float64 `json:"magicResist"`
	AttackDamage    float64 `json:"attackDamage"`
	AttackSpeed     float64 `json:"attackSpeed"`
	AttackRange     float64 `json:"attackRange"`
	MoveSpeed       float64 `json:"moveSpeed"`
	HealthRegenRate float64 `json:"healthRegenRate"`
}

// Amostra e uma leitura das estatisticas num nivel.
type Amostra struct {
	Nivel   int         `json:"nivel"`
	Campeao string      `json:"campeao"`
	Stats   StatsDoJogo `json:"stats"`

	// Itens sao os equipados no momento da leitura. Item que concede atributo
	// contamina a comparacao, e por isso a lista acompanha a amostra em vez de
	// ficar de fora — sem ela, uma divergencia real e uma bota comprada ficam
	// indistinguiveis.
	Itens []string `json:"itens,omitempty"`

	// TempoDeJogo em segundos, para ordenar amostras do mesmo nivel.
	TempoDeJogo float64 `json:"tempo_de_jogo"`

	// Partida identifica a partida pela escalacao. Duas leituras com escalacoes
	// diferentes sao de partidas diferentes, e comparar o crescimento entre elas
	// nao cancela bonus nenhum — as runas e os fragmentos mudam junto.
	//
	// A API nao publica identificador de partida, entao a escalacao e o que da
	// para usar. Ela NAO distingue duas partidas de Treino montadas iguais; quem
	// pega esse caso e o tempo de jogo, que reinicia do zero.
	Partida string `json:"partida,omitempty"`
}

type respostaLive struct {
	ActivePlayer struct {
		Level         int         `json:"level"`
		RiotID        string      `json:"riotId"`
		SummonerName  string      `json:"summonerName"`
		ChampionStats StatsDoJogo `json:"championStats"`
	} `json:"activePlayer"`
	AllPlayers []struct {
		RiotID       string `json:"riotId"`
		SummonerName string `json:"summonerName"`
		ChampionName string `json:"championName"`
		Items        []struct {
			DisplayName string `json:"displayName"`
		} `json:"items"`
	} `json:"allPlayers"`
	GameData struct {
		GameTime float64 `json:"gameTime"`
		GameMode string  `json:"gameMode"`
	} `json:"gameData"`
}

// escalacao resume quem esta na partida, para servir de identidade dela.
//
// Ordenada antes de resumir porque a API nao promete ordem estavel entre
// chamadas, e uma reordenacao viraria "outra partida" sem que nada tivesse
// mudado. O resumo e curto de proposito: ele so precisa distinguir, e o arquivo
// de amostras e lido por gente.
func (r respostaLive) escalacao() string {
	linhas := make([]string, 0, len(r.AllPlayers))
	for _, p := range r.AllPlayers {
		quem := p.RiotID
		if quem == "" {
			quem = p.SummonerName
		}
		linhas = append(linhas, quem+"|"+p.ChampionName)
	}
	if len(linhas) == 0 {
		return ""
	}
	sort.Strings(linhas)
	soma := sha256.Sum256([]byte(strings.Join(linhas, "\n")))
	return hex.EncodeToString(soma[:4])
}

// Amostrar le o estado atual da partida.
func (c *LiveClient) Amostrar() (*Amostra, error) {
	raw, err := c.get("/allgamedata")
	if err != nil {
		return nil, err
	}

	var r respostaLive
	if err := json.Unmarshal(raw, &r); err != nil {
		return nil, fmt.Errorf("lendo os dados da partida: %w", err)
	}

	am := &Amostra{
		Nivel:       r.ActivePlayer.Level,
		Stats:       r.ActivePlayer.ChampionStats,
		TempoDeJogo: r.GameData.GameTime,
		Partida:     r.escalacao(),
	}

	// O nome do campeao nao vem no bloco do jogador ativo: e preciso casar com a
	// lista de participantes pelo identificador.
	eu := r.ActivePlayer.RiotID
	if eu == "" {
		eu = r.ActivePlayer.SummonerName
	}
	for _, p := range r.AllPlayers {
		if p.RiotID != eu && p.SummonerName != eu {
			continue
		}
		am.Campeao = p.ChampionName
		for _, it := range p.Items {
			am.Itens = append(am.Itens, it.DisplayName)
		}
	}
	if am.Campeao == "" {
		return nil, fmt.Errorf("nao foi possivel identificar o campeao do jogador %q "+
			"entre os %d participantes", eu, len(r.AllPlayers))
	}
	return am, nil
}

// Modo devolve o modo da partida, para conferir que a amostra e do Summoner's
// Rift e nao de ARAM.
func (c *LiveClient) Modo() (string, error) {
	raw, err := c.get("/gamestats")
	if err != nil {
		return "", err
	}
	var g struct {
		GameMode string `json:"gameMode"`
	}
	if err := json.Unmarshal(raw, &g); err != nil {
		return "", err
	}
	return g.GameMode, nil
}
