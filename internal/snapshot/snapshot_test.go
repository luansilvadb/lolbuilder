package snapshot

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestBeginWriteCommitEExists(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir)

	if s.Exists("16.16") {
		t.Fatal("snapshot inexistente reportado como existente")
	}

	w, err := s.BeginWrite("16.16")
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Add("items.json", []byte(`[]`), "etag-1"); err != nil {
		t.Fatal(err)
	}
	if err := w.Add("pt_br/items.json", []byte(`[]`), "etag-2"); err != nil {
		t.Fatal(err)
	}
	if err := w.Add("characters/garen/garen.bin.json", []byte(`{}`), "etag-3"); err != nil {
		t.Fatal(err)
	}
	if err := w.Commit("16.16.123+x", "https://exemplo", "latest", map[string]int{"runes": 103}); err != nil {
		t.Fatal(err)
	}

	if !s.Exists("16.16") {
		t.Fatal("snapshot gravado nao foi encontrado")
	}
	cap, err := s.LoadCapture("16.16")
	if err != nil {
		t.Fatal(err)
	}
	if cap.Patch != "16.16" || cap.PatchFull != "16.16.123+x" || cap.Patchline != "latest" {
		t.Fatalf("manifesto gravado errado: %+v", cap)
	}
	if cap.Retroactive() {
		t.Fatal("captura de latest marcada como retroativa")
	}
	if cap.Counts["runes"] != 103 {
		t.Fatalf("contagens perdidas: %+v", cap.Counts)
	}
	if cap.Files["characters/garen/garen.bin.json"].ETag != "etag-3" {
		t.Fatalf("etags perdidos: %+v", cap.Files)
	}

	got, err := s.ReadFile("16.16", "pt_br/items.json")
	if err != nil || string(got) != "[]" {
		t.Fatalf("ReadFile devolveu %q, %v", string(got), err)
	}
}

// TestAbortNaoDeixaResto: um aborto no meio nunca pode deixar snapshot parcial,
// porque snapshot parcial e indistinguivel de "a Riot removeu o conteudo".
func TestAbortNaoDeixaResto(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir)

	w, err := s.BeginWrite("16.16")
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Add("items.json", []byte(`[]`), ""); err != nil {
		t.Fatal(err)
	}
	w.Abort()

	if s.Exists("16.16") {
		t.Fatal("aborto promoveu a captura")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		t.Fatalf("aborto deixou resto em disco: %s", e.Name())
	}
}

func TestNaoSobrescreveCapturaExistente(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir)

	w, _ := s.BeginWrite("16.16")
	if err := w.Commit("16.16.1", "src", "latest", nil); err != nil {
		t.Fatal(err)
	}

	_, err := s.BeginWrite("16.16")
	if !errors.Is(err, ErrAlreadyExists) {
		t.Fatalf("BeginWrite sobre captura existente devolveu %v", err)
	}
}

func TestCapturaRetroativaEMarcada(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir)

	w, _ := s.BeginWrite("15.20")
	if err := w.Commit("15.20.1", "src", "15.20", nil); err != nil {
		t.Fatal(err)
	}
	cap, err := s.LoadCapture("15.20")
	if err != nil {
		t.Fatal(err)
	}
	if !cap.Retroactive() {
		t.Fatal("captura de patchline versionado nao foi marcada como retroativa")
	}
}

// TestLatestPatchOrdenaPorVersaoENaoPorTexto: "16.9" vem antes de "16.10".
func TestLatestPatchOrdenaPorVersao(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir)
	for _, p := range []string{"16.9", "16.10", "16.2"} {
		w, err := s.BeginWrite(p)
		if err != nil {
			t.Fatal(err)
		}
		if err := w.Commit(p, "src", "latest", nil); err != nil {
			t.Fatal(err)
		}
	}
	got, err := s.LatestPatch()
	if err != nil {
		t.Fatal(err)
	}
	if got != "16.10" {
		t.Fatalf("LatestPatch = %q, esperado 16.10 (ordem numerica, nao alfabetica)", got)
	}
}

func TestLatestPatchVazio(t *testing.T) {
	s := NewStore(filepath.Join(t.TempDir(), "naoexiste"))
	got, err := s.LatestPatch()
	if err != nil {
		t.Fatalf("diretorio ausente deveria ser vazio, nao erro: %v", err)
	}
	if got != "" {
		t.Fatalf("LatestPatch = %q em diretorio vazio", got)
	}
}

// TestLatestPatchIgnoraDiretorioSemManifesto: uma area temporaria que sobreviveu
// a um kill -9 nao pode ser confundida com captura.
func TestLatestPatchIgnoraDiretorioSemManifesto(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir)
	if err := os.MkdirAll(filepath.Join(dir, "16.20"), 0o755); err != nil {
		t.Fatal(err)
	}
	got, err := s.LatestPatch()
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Fatalf("diretorio sem capture.json contou como captura: %q", got)
	}
}

func TestLoadCaptureAusenteNaoEErro(t *testing.T) {
	s := NewStore(t.TempDir())
	cap, err := s.LoadCapture("16.16")
	if err != nil || cap != nil {
		t.Fatalf("LoadCapture de captura ausente devolveu %+v, %v", cap, err)
	}
}
