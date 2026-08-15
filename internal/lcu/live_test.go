package lcu

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// respostaDeExemplo tem a forma real: o bloco do jogador ativo NAO traz o nome
// do campeao, que so existe na lista de participantes.
const respostaDeExemplo = `{
  "activePlayer": {
    "level": 6,
    "riotId": "Fulano#BR1",
    "summonerName": "Fulano",
    "championStats": {
      "maxHealth": 1290, "armor": 58, "magicResist": 42,
      "attackDamage": 94, "attackSpeed": 0.73, "attackRange": 175,
      "moveSpeed": 340, "healthRegenRate": 3.1
    }
  },
  "allPlayers": [
    {"riotId": "Outro#BR1", "summonerName": "Outro", "championName": "Darius", "items": []},
    {"riotId": "Fulano#BR1", "summonerName": "Fulano", "championName": "Garen",
     "items": [{"displayName": "Cota de Malha"}, {"displayName": "Botas"}]}
  ],
  "gameData": {"gameTime": 612.5, "gameMode": "CLASSIC"}
}`

func servidor(t *testing.T, h http.HandlerFunc) *LiveClient {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	c := NewLive(2 * time.Second)
	c.base = srv.URL
	return c
}

func TestAmostrar(t *testing.T) {
	c := servidor(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/allgamedata" {
			t.Errorf("caminho = %q", r.URL.Path)
		}
		w.Write([]byte(respostaDeExemplo))
	})

	am, err := c.Amostrar()
	if err != nil {
		t.Fatal(err)
	}
	// O campeao vem da lista de participantes, casada pelo identificador.
	if am.Campeao != "Garen" {
		t.Errorf("campeao = %q, esperado Garen", am.Campeao)
	}
	if am.Nivel != 6 || am.TempoDeJogo != 612.5 {
		t.Errorf("nivel=%d tempo=%v", am.Nivel, am.TempoDeJogo)
	}
	if am.Stats.MaxHealth != 1290 || am.Stats.Armor != 58 {
		t.Errorf("stats = %+v", am.Stats)
	}
	// Os itens acompanham a amostra porque contaminam a comparacao: sem eles,
	// uma divergencia real e uma bota comprada ficam indistinguiveis.
	if len(am.Itens) != 2 || am.Itens[0] != "Cota de Malha" {
		t.Errorf("itens = %v", am.Itens)
	}
}

// TestJogadorNaoIdentificadoEErro: sem saber o campeao, a amostra nao pode ser
// comparada com nada — e adivinhar pelo primeiro da lista compararia o dataset
// com o campeao errado.
func TestJogadorNaoIdentificadoEErro(t *testing.T) {
	semEu := strings.Replace(respostaDeExemplo,
		`{"riotId": "Fulano#BR1", "summonerName": "Fulano", "championName": "Garen",
     "items": [{"displayName": "Cota de Malha"}, {"displayName": "Botas"}]}`,
		`{"riotId": "Terceiro#BR1", "summonerName": "Terceiro", "championName": "Teemo", "items": []}`, 1)

	c := servidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(semEu))
	})
	if _, err := c.Amostrar(); err == nil {
		t.Fatal("amostra sem o jogador identificado passou")
	}
}

// TestSemPartidaNaoEFalhaDeLeitura: a porta fechada e o estado normal fora do
// jogo, e tratar isso como erro faria o comando gritar sempre que ninguem
// estivesse jogando.
func TestSemPartidaNaoEFalhaDeLeitura(t *testing.T) {
	c := NewLive(200 * time.Millisecond)
	c.base = "https://127.0.0.1:1" // porta fechada

	if c.EmPartida() {
		t.Fatal("EmPartida devolveu verdadeiro sem servidor")
	}
	_, err := c.Amostrar()
	if !errors.Is(err, ErrSemPartida) {
		t.Fatalf("erro = %v, esperado ErrSemPartida", err)
	}
}

func TestEmPartida(t *testing.T) {
	c := servidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"gameMode":"CLASSIC"}`))
	})
	if !c.EmPartida() {
		t.Fatal("EmPartida devolveu falso com a partida respondendo")
	}
	modo, err := c.Modo()
	if err != nil || modo != "CLASSIC" {
		t.Fatalf("Modo = %q (%v)", modo, err)
	}
}

func TestStatusNaoOKEErro(t *testing.T) {
	c := servidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	if c.EmPartida() {
		t.Error("status 500 foi lido como partida em andamento")
	}
	if _, err := c.Amostrar(); err == nil {
		t.Error("status 500 passou em silencio")
	}
}

func TestRespostaMalformada(t *testing.T) {
	c := servidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("nao e json"))
	})
	if _, err := c.Amostrar(); err == nil {
		t.Fatal("resposta malformada passou")
	}
}
