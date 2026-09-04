package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDeckFromPath(t *testing.T) {
	cases := map[string]string{
		"/x/Principal SWE/01. Foundations/Programming/09. Language/Golang/Note.md": "Principal SWE::Foundations::Programming::Language::Golang",
		"/x/Principal SWE/Note.md":     "Principal SWE",
		"/somewhere/else/Note.md":      "Principal SWE",
		"/x/Principal SWE/04. AI/Note": "Principal SWE::AI",
	}
	for path, want := range cases {
		if got := DeckFromPath(path); got != want {
			t.Errorf("DeckFromPath(%q) = %q, want %q", path, got, want)
		}
	}
}

func TestCollectFilesSkipsToolingAndNonMarkdown(t *testing.T) {
	root := t.TempDir()

	files := map[string]string{
		"a.md":                   "note",
		"sub/b.md":               "note",
		"sub/c.txt":              "not markdown",
		"tools/d.md":             "tooling",
		".obsidian/plugins/e.md": "config",
	}
	for rel, body := range files {
		full := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	got, err := CollectFiles([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d files, want 2: %v", len(got), got)
	}
	for _, f := range got {
		if strings.Contains(f, "tools/") || strings.Contains(f, ".obsidian") || !strings.HasSuffix(f, ".md") {
			t.Errorf("unexpected file collected: %s", f)
		}
	}
}

func TestCollectFilesSkipsFolderIndexNotes(t *testing.T) {
	root := t.TempDir()

	files := []string{
		"09. Language/Golang/Golang.md", // index note of its folder
		"09. Language/Golang/Slices.md", // real material
		"09. Language/09. Language.md",  // index note, ordering prefix ignored
	}
	for _, rel := range files {
		full := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("note"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	got, err := CollectFiles([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range got {
		if base := filepath.Base(f); base == "Golang.md" || base == "09. Language.md" {
			t.Errorf("index note collected: %s", f)
		}
	}
}

func TestCollectFilesAcceptsASingleFile(t *testing.T) {
	path := writeNote(t, "## 1. X\n\nbody\n")
	got, err := CollectFiles([]string{path})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0] != path {
		t.Errorf("got %v, want [%s]", got, path)
	}
}

func TestFlagsFirstMovesPositionalToTheEnd(t *testing.T) {
	got := flagsFirst([]string{"note.md", "--deck", "Go", "--update", "--tags=a,b"})
	want := []string{"--deck", "Go", "--update", "--tags=a,b", "note.md"}

	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("flagsFirst = %v, want %v", got, want)
	}
}

func TestEscapeQuery(t *testing.T) {
	got := escapeQuery(`a "b" *c* _d_ \e`)
	want := `a \"b\" \*c\* \_d\_ \\e`
	if got != want {
		t.Errorf("escapeQuery = %q, want %q", got, want)
	}
}

func TestOwnedTemplate(t *testing.T) {
	ours := map[string]map[string]string{cardTemplateName: {}}
	if name, ok := ownedTemplate(ours, []string{"Front", "Back"}); !ok || name != cardTemplateName {
		t.Errorf("own template not recognised: %q %v", name, ok)
	}

	// An older import of ours: two fields, one template, any name.
	legacy := map[string]map[string]string{"Card 1": {}}
	if name, ok := ownedTemplate(legacy, []string{"Front", "Back"}); !ok || name != "Card 1" {
		t.Errorf("legacy template not recognised: %q %v", name, ok)
	}

	// Someone else's rich note type must never be restyled.
	theirs := map[string]map[string]string{"Card 1": {}, "Card 2": {}}
	if _, ok := ownedTemplate(theirs, []string{"Front", "Back", "Example", "Output"}); ok {
		t.Error("foreign note type was treated as ours")
	}
}

func TestDedupe(t *testing.T) {
	got := dedupe([]string{"go", " go ", "", "concurrency", "go"})
	if strings.Join(got, ",") != "go,concurrency" {
		t.Errorf("dedupe = %v", got)
	}
}
