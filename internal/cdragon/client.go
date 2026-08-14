// Package cdragon fala com o CommunityDragon: descoberta de patch e download
// condicional das fontes de dados.
package cdragon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const userAgent = "lolbuilder/1 (+ferramenta pessoal de dataset do League of Legends)"

// Client e um wrapper fino sobre net/http com timeout e User-Agent proprios.
type Client struct {
	http *http.Client
}

// New cria um cliente com o timeout informado.
func New(timeout time.Duration) *Client {
	return &Client{http: &http.Client{Timeout: timeout}}
}

// Result e o desfecho de um download condicional.
type Result struct {
	Body        []byte // vazio quando NotModified
	ETag        string
	NotModified bool
}

// Fetch busca uma URL. Se etag nao for vazio, envia If-None-Match e trata 304
// como "inalterado" em vez de erro.
func (c *Client) Fetch(url, etag string) (*Result, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("montando requisicao para %s: %w", url, err)
	}
	req.Header.Set("User-Agent", userAgent)
	if etag != "" {
		req.Header.Set("If-None-Match", etag)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("buscando %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return &Result{ETag: etag, NotModified: true}, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("buscando %s: status %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("lendo corpo de %s: %w", url, err)
	}
	return &Result{Body: body, ETag: resp.Header.Get("ETag")}, nil
}

// Status e o conteudo de status.live.txt: um carimbo de tempo e um estado.
type Status struct {
	Timestamp string
	State     string
}

// Done informa se o CDragon terminou de processar o patchline. Baixar durante o
// processamento pode render arquivos incompletos.
func (s Status) Done() bool {
	return strings.EqualFold(s.State, "done")
}

// FetchStatus le status.live.txt. Formato observado: "<timestamp> <estado>".
func (c *Client) FetchStatus(url string) (*Status, error) {
	res, err := c.Fetch(url, "")
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(string(res.Body))
	if len(fields) < 2 {
		return nil, fmt.Errorf("status.live.txt em formato inesperado: %q", string(res.Body))
	}
	return &Status{Timestamp: fields[0], State: fields[len(fields)-1]}, nil
}

// patchRe extrai "16.16" de "16.16.8049184+branch.releases-16-16.content.release".
var patchRe = regexp.MustCompile(`^(\d+)\.(\d+)`)

// FetchPatch descobre a versao do patch servida pelo patchline configurado.
// Devolve a forma curta (ex "16.16") e a string completa.
func (c *Client) FetchPatch(metadataURL string) (short, full string, err error) {
	res, err := c.Fetch(metadataURL, "")
	if err != nil {
		return "", "", err
	}
	var meta struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(res.Body, &meta); err != nil {
		return "", "", fmt.Errorf("parseando content-metadata.json: %w", err)
	}
	if meta.Version == "" {
		return "", "", fmt.Errorf("content-metadata.json sem campo version")
	}
	m := patchRe.FindStringSubmatch(meta.Version)
	if m == nil {
		return "", "", fmt.Errorf("versao em formato inesperado: %q", meta.Version)
	}
	return m[1] + "." + m[2], meta.Version, nil
}
