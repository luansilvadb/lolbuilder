package cdragon

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func novoServidor(t *testing.T, h http.HandlerFunc) (*Client, string) {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	return New(2 * time.Second), srv.URL
}

func TestFetchDevolveCorpoEETag(t *testing.T) {
	cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("User-Agent"); !strings.HasPrefix(got, "lolbuilder/") {
			t.Errorf("User-Agent = %q, esperado prefixo lolbuilder/", got)
		}
		if r.Header.Get("If-None-Match") != "" {
			t.Error("If-None-Match enviado sem etag anterior")
		}
		w.Header().Set("ETag", `"abc"`)
		w.Write([]byte(`{"ok":true}`))
	})

	res, err := cli.Fetch(url, "")
	if err != nil {
		t.Fatal(err)
	}
	if res.NotModified {
		t.Fatal("200 foi tratado como 304")
	}
	if string(res.Body) != `{"ok":true}` || res.ETag != `"abc"` {
		t.Fatalf("resultado inesperado: %+v", res)
	}
}

// TestFetch304NaoEErro protege o mecanismo que evita rebaixar 12 MB de dump a
// cada patch: 304 tem de voltar como "inalterado", preservando o etag para o
// chamador reaproveitar os bytes da captura anterior.
func TestFetch304NaoEErro(t *testing.T) {
	cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("If-None-Match") != `"abc"` {
			t.Errorf("If-None-Match = %q, esperado \"abc\"", r.Header.Get("If-None-Match"))
		}
		w.WriteHeader(http.StatusNotModified)
	})

	res, err := cli.Fetch(url, `"abc"`)
	if err != nil {
		t.Fatalf("304 virou erro: %v", err)
	}
	if !res.NotModified {
		t.Fatal("304 nao foi marcado como inalterado")
	}
	if res.ETag != `"abc"` {
		t.Fatalf("304 perdeu o etag: %q", res.ETag)
	}
	if len(res.Body) != 0 {
		t.Fatalf("304 trouxe corpo: %q", string(res.Body))
	}
}

func TestFetchStatusNaoOKEErro(t *testing.T) {
	for _, code := range []int{http.StatusNotFound, http.StatusInternalServerError, http.StatusForbidden} {
		cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
		})
		_, err := cli.Fetch(url, "")
		if err == nil {
			t.Errorf("status %d passou em silencio", code)
			continue
		}
		if !strings.Contains(err.Error(), url) {
			t.Errorf("o erro de %d nao diz qual URL falhou: %v", code, err)
		}
	}
}

func TestFetchStatusLive(t *testing.T) {
	cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("2026-08-14T09:00:03Z done\n"))
	})

	st, err := cli.FetchStatus(url)
	if err != nil {
		t.Fatal(err)
	}
	if st.Timestamp != "2026-08-14T09:00:03Z" || st.State != "done" {
		t.Fatalf("status lido errado: %+v", st)
	}
	if !st.Done() {
		t.Fatal("estado done nao foi reconhecido")
	}
}

// TestStatusNaoDoneBloqueia: baixar durante o processamento do CDragon pode
// render arquivos incompletos, que e pior que nao baixar.
func TestStatusNaoDoneBloqueia(t *testing.T) {
	cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("2026-08-14T09:00:03Z running"))
	})

	st, err := cli.FetchStatus(url)
	if err != nil {
		t.Fatal(err)
	}
	if st.Done() {
		t.Fatal("estado running foi tratado como pronto")
	}
}

func TestStatusMalformadoEErro(t *testing.T) {
	cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("soUmaPalavra"))
	})
	if _, err := cli.FetchStatus(url); err == nil {
		t.Fatal("status.live.txt em formato inesperado passou")
	}
}

func TestFetchPatchExtraiFormaCurta(t *testing.T) {
	cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"version":"16.16.8049184+branch.releases-16-16.content.release"}`))
	})

	short, full, err := cli.FetchPatch(url)
	if err != nil {
		t.Fatal(err)
	}
	if short != "16.16" {
		t.Fatalf("forma curta = %q, esperado 16.16", short)
	}
	if !strings.HasPrefix(full, "16.16.8049184") {
		t.Fatalf("forma completa perdida: %q", full)
	}
}

// TestFetchPatchNaoConfundeMinorComPatch: "16.9" e "16.10" precisam sair
// distintos, ou duas capturas diferentes cairiam no mesmo diretorio.
func TestFetchPatchFormasValidas(t *testing.T) {
	casos := map[string]string{
		"16.9.1+x":  "16.9",
		"16.10.1+x": "16.10",
		"15.20.999": "15.20",
		"1.0":       "1.0",
	}
	for versao, want := range casos {
		cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"version":"` + versao + `"}`))
		})
		got, _, err := cli.FetchPatch(url)
		if err != nil {
			t.Errorf("versao %q falhou: %v", versao, err)
			continue
		}
		if got != want {
			t.Errorf("versao %q virou %q, esperado %q", versao, got, want)
		}
	}
}

func TestFetchPatchInvalido(t *testing.T) {
	casos := map[string]string{
		"json quebrado":   `nao e json`,
		"sem version":     `{"outra":"coisa"}`,
		"version vazia":   `{"version":""}`,
		"version sem num": `{"version":"branch.releases"}`,
	}
	for nome, corpo := range casos {
		t.Run(nome, func(t *testing.T) {
			cli, url := novoServidor(t, func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(corpo))
			})
			if _, _, err := cli.FetchPatch(url); err == nil {
				t.Fatal("content-metadata.json invalido passou")
			}
		})
	}
}

func TestFetchURLInvalida(t *testing.T) {
	cli := New(time.Second)
	if _, err := cli.Fetch("://sem-esquema", ""); err == nil {
		t.Fatal("URL malformada passou")
	}
}
