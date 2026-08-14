// Package snapshot grava as capturas por patch.
//
// Ao contrario do modo Jade, o CDragon versiona os patchlines do jogo base:
// uma captura perdida pode ser refeita depois apontando -patch para o patchline
// versionado. Isso muda a urgencia, nao a disciplina — um snapshot gravado
// continua nunca sendo reescrito, e a gravacao continua atomica, porque um
// snapshot parcial promovido a definitivo e indistinguivel de "a Riot removeu o
// conteudo" e o erro so apareceria semanas depois, na saida.
package snapshot

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// FileMeta e o registro de um arquivo dentro da captura.
type FileMeta struct {
	ETag  string `json:"etag"`
	Bytes int    `json:"bytes"`
}

// Capture e o manifesto da captura, gravado junto dos dados brutos.
type Capture struct {
	Patch      string `json:"patch"`
	PatchFull  string `json:"patch_full"`
	CapturedAt string `json:"captured_at"`
	Source     string `json:"source"`

	// Patchline registra de onde a captura veio: "latest" numa captura ao vivo,
	// ou o patchline versionado numa captura retroativa.
	//
	// E a procedencia, e existe porque patchlines antigos do CDragon tem
	// completude variavel. Sem isso, um snapshot retroativo degradado ficaria
	// indistinguivel de um bom, e poderia acabar calibrando os minimos de
	// cobertura — que e o unico numero do projeto que nao pode nascer de dado
	// suspeito.
	Patchline string `json:"patchline"`

	Counts map[string]int      `json:"counts"`
	Files  map[string]FileMeta `json:"files"`
}

// Retroactive informa se a captura veio de um patchline versionado.
func (c *Capture) Retroactive() bool {
	return c.Patchline != "" && c.Patchline != "latest"
}

// ErrAlreadyExists sinaliza que o patch ja foi capturado.
var ErrAlreadyExists = errors.New("snapshot ja existe para este patch")

// Store gerencia o diretorio de snapshots.
type Store struct {
	dir string
}

// NewStore aponta para o diretorio raiz das capturas.
func NewStore(dir string) *Store { return &Store{dir: dir} }

// PatchDir e o caminho da captura de um patch.
func (s *Store) PatchDir(patch string) string { return filepath.Join(s.dir, patch) }

// Exists informa se o patch ja foi capturado.
func (s *Store) Exists(patch string) bool {
	_, err := os.Stat(filepath.Join(s.PatchDir(patch), "capture.json"))
	return err == nil
}

// LoadCapture le o manifesto de uma captura anterior. Devolve nil sem erro
// quando a captura nao existe.
func (s *Store) LoadCapture(patch string) (*Capture, error) {
	raw, err := os.ReadFile(filepath.Join(s.PatchDir(patch), "capture.json"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var c Capture
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("capture.json de %s corrompido: %w", patch, err)
	}
	return &c, nil
}

// LatestPatch devolve o patch da captura mais recente, comparando pelos nomes
// de diretorio em ordem natural de versao. Devolve string vazia se nao houver
// nenhuma.
func (s *Store) LatestPatch() (string, error) {
	entries, err := os.ReadDir(s.dir)
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	var patches []string
	for _, e := range entries {
		if e.IsDir() && s.Exists(e.Name()) {
			patches = append(patches, e.Name())
		}
	}
	if len(patches) == 0 {
		return "", nil
	}
	sort.Slice(patches, func(i, j int) bool { return lessVersion(patches[i], patches[j]) })
	return patches[len(patches)-1], nil
}

// ReadFile le um arquivo bruto de uma captura anterior — usado quando o
// servidor responde 304 e o conteudo pode ser reaproveitado.
func (s *Store) ReadFile(patch, name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.PatchDir(patch), name))
}

// Writer acumula os arquivos de uma captura em area temporaria e so os promove
// ao diretorio final no Commit. Um aborto no meio nunca deixa snapshot parcial.
type Writer struct {
	store    *Store
	patch    string
	tmpDir   string
	files    map[string]FileMeta
	commited bool
}

// BeginWrite prepara uma captura nova. Falha se o patch ja existe.
func (s *Store) BeginWrite(patch string) (*Writer, error) {
	if s.Exists(patch) {
		return nil, fmt.Errorf("%w: %s", ErrAlreadyExists, patch)
	}
	tmp := filepath.Join(s.dir, ".tmp-"+patch)
	if err := os.RemoveAll(tmp); err != nil {
		return nil, fmt.Errorf("limpando area temporaria: %w", err)
	}
	if err := os.MkdirAll(tmp, 0o755); err != nil {
		return nil, fmt.Errorf("criando area temporaria: %w", err)
	}
	return &Writer{store: s, patch: patch, tmpDir: tmp, files: map[string]FileMeta{}}, nil
}

// Add grava um arquivo bruto na area temporaria, sem transformacao alguma.
func (w *Writer) Add(name string, data []byte, etag string) error {
	dest := filepath.Join(w.tmpDir, name)
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(dest, data, 0o644); err != nil {
		return fmt.Errorf("gravando %s: %w", name, err)
	}
	w.files[name] = FileMeta{ETag: etag, Bytes: len(data)}
	return nil
}

// Commit grava o manifesto e promove a area temporaria ao diretorio definitivo.
func (w *Writer) Commit(patchFull, source, patchline string, counts map[string]int) error {
	cap := Capture{
		Patch:      w.patch,
		PatchFull:  patchFull,
		CapturedAt: time.Now().UTC().Format(time.RFC3339),
		Source:     source,
		Patchline:  patchline,
		Counts:     counts,
		Files:      w.files,
	}
	raw, err := json.MarshalIndent(cap, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(w.tmpDir, "capture.json"), append(raw, '\n'), 0o644); err != nil {
		return fmt.Errorf("gravando capture.json: %w", err)
	}

	final := w.store.PatchDir(w.patch)
	if _, err := os.Stat(final); err == nil {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, w.patch)
	}
	if err := os.MkdirAll(filepath.Dir(final), 0o755); err != nil {
		return err
	}
	if err := os.Rename(w.tmpDir, final); err != nil {
		return fmt.Errorf("promovendo captura de %s: %w", w.patch, err)
	}
	w.commited = true
	return nil
}

// Abort descarta a area temporaria. Seguro de chamar apos Commit.
func (w *Writer) Abort() {
	if !w.commited {
		_ = os.RemoveAll(w.tmpDir)
	}
}

// lessVersion compara "16.9" < "16.10" numericamente, e nao alfabeticamente.
func lessVersion(a, b string) bool {
	pa, pb := splitVersion(a), splitVersion(b)
	for i := 0; i < len(pa) && i < len(pb); i++ {
		if pa[i] != pb[i] {
			return pa[i] < pb[i]
		}
	}
	return len(pa) < len(pb)
}

func splitVersion(v string) []int {
	var out []int
	cur, has := 0, false
	for _, r := range v {
		if r >= '0' && r <= '9' {
			cur, has = cur*10+int(r-'0'), true
			continue
		}
		if has {
			out = append(out, cur)
		}
		cur, has = 0, false
	}
	if has {
		out = append(out, cur)
	}
	return out
}
