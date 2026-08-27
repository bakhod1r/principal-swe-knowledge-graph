package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// captureOut redirects the run report for one test.
func captureOut(t *testing.T) *bytes.Buffer {
	t.Helper()
	buf := &bytes.Buffer{}
	old := out
	out = buf
	t.Cleanup(func() { out = old })
	return buf
}

// vaultWith writes notes into a temp vault and returns its root.
func vaultWith(t *testing.T, notes map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for rel, body := range notes {
		full := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

// ─── flags ───

func TestParseFlags(t *testing.T) {
	opts, err := parseFlags([]string{"note.md", "--deck", "Go", "--tags", " a , b ,,", "--update"})
	if err != nil {
		t.Fatal(err)
	}

	if opts.deck != "Go" || !opts.update {
		t.Errorf("opts = %+v", opts)
	}
	if strings.Join(opts.tags, ",") != "a,b" {
		t.Errorf("tags = %v, want [a b] with blanks dropped", opts.tags)
	}
	if opts.model != "Principal SWE" || opts.url != "http://127.0.0.1:8765" {
		t.Errorf("defaults lost: %+v", opts)
	}
	if len(opts.targets) != 1 || opts.targets[0] != "note.md" {
		t.Errorf("targets = %v", opts.targets)
	}
}

func TestParseFlagsWithoutATargetFails(t *testing.T) {
	if _, err := parseFlags([]string{"--update"}); err == nil {
		t.Fatal("want an error when no note is given")
	}
}

func TestParseFlagsRejectsAnUnknownFlag(t *testing.T) {
	devNull, _ := os.Open(os.DevNull)
	defer devNull.Close()

	stderr := os.Stderr
	os.Stderr = devNull
	defer func() { os.Stderr = stderr }()

	if _, err := parseFlags([]string{"note.md", "--nope"}); err == nil {
		t.Fatal("want an error for an unknown flag")
	}
}

// ─── card building ───

func TestBuildCardsDropsDuplicateFronts(t *testing.T) {
	root := vaultWith(t, map[string]string{
		"a/Note.md": "## 1. Same\n\nBody A.\n",
		"b/Note.md": "## 1. Same\n\nBody B.\n",
	})

	cards, files, err := buildCards(options{targets: []string{root}})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("files = %d, want 2", len(files))
	}
	if len(cards) != 1 {
		t.Fatalf("cards = %d, want 1 — the identical front should be dropped", len(cards))
	}
}

func TestBuildCardsHonoursTheDeckOverride(t *testing.T) {
	root := vaultWith(t, map[string]string{"Note.md": "## 1. X\n\nBody.\n"})

	cards, _, err := buildCards(options{targets: []string{root}, deck: "Custom::Deck"})
	if err != nil {
		t.Fatal(err)
	}
	if cards[0].Deck != "Custom::Deck" {
		t.Errorf("deck = %q", cards[0].Deck)
	}
}

func TestBuildCardsErrors(t *testing.T) {
	if _, _, err := buildCards(options{targets: []string{"/no/such/path"}}); err == nil {
		t.Error("want an error for a missing path")
	}

	empty := t.TempDir()
	if _, _, err := buildCards(options{targets: []string{empty}}); err == nil {
		t.Error("want an error when the folder holds no notes")
	}
}

func TestNotesForBlocksDuplicatesAcrossTheCollection(t *testing.T) {
	notes := notesFor(
		[]Card{{Front: "Q", Back: "A", Deck: "D", Tags: []string{"go"}}},
		"Principal SWE",
		Fields{Front: "Front", Back: "Back"},
	)

	n := notes[0]
	if n.Options["allowDuplicate"] != false {
		t.Error("duplicates must not be allowed")
	}
	if n.Options["duplicateScope"] != "collection" {
		t.Errorf("duplicateScope = %v, want collection-wide", n.Options["duplicateScope"])
	}
	if n.Fields["Front"] != "Q" || n.ModelName != "Principal SWE" || n.DeckName != "D" {
		t.Errorf("note = %+v", n)
	}
}

func TestNotesForWritesIntoAForeignNoteTypesFields(t *testing.T) {
	notes := notesFor([]Card{{Front: "Q", Back: "A"}}, "LeetCode", Fields{Front: "Title", Back: "Kind"})

	if notes[0].Fields["Title"] != "Q" || notes[0].Fields["Kind"] != "A" {
		t.Errorf("fields = %v", notes[0].Fields)
	}
}

func TestDeckNames(t *testing.T) {
	got := deckNames([]Card{{Deck: "A"}, {Deck: "B"}, {Deck: "A"}})
	if strings.Join(got, ",") != "A,B,A" {
		t.Errorf("deckNames = %v", got)
	}
}

// ─── reporting ───

func TestReport(t *testing.T) {
	cases := []struct {
		name     string
		added    int
		existing int
		updated  bool
		want     string
	}{
		{"all new", 5, 0, false, "✅ Added: 5"},
		{"skipped", 0, 3, false, "Already in Anki, skipped: 3"},
		{"updated", 1, 3, true, "Updated (already existed): 3"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buf := captureOut(t)
			report(tc.added, tc.existing, tc.updated)
			if !strings.Contains(buf.String(), tc.want) {
				t.Errorf("report = %q, want it to mention %q", buf.String(), tc.want)
			}
		})
	}
}

func TestListNewNotesGroupsByNote(t *testing.T) {
	buf := captureOut(t)
	cards := []Card{
		{Source: "/v/A.md"}, {Source: "/v/A.md"}, {Source: "/v/B.md"},
	}
	listNewNotes(cards, []int{0, 1, 2})

	got := buf.String()
	if !strings.Contains(got, "🆕 New notes: 2") {
		t.Errorf("missing note count:\n%s", got)
	}
	if !strings.Contains(got, "2    A.md") || !strings.Contains(got, "1    B.md") {
		t.Errorf("per-note counts wrong:\n%s", got)
	}
}

func TestListNewNotesStaysQuietWhenNothingIsNew(t *testing.T) {
	buf := captureOut(t)
	listNewNotes([]Card{{Source: "/v/A.md"}}, nil)

	if buf.Len() != 0 {
		t.Errorf("printed %q for an empty import", buf.String())
	}
}

// ─── preview ───

func TestPreviewPrintsEveryCard(t *testing.T) {
	buf := captureOut(t)
	preview([]Card{
		{Front: "Q1", Back: "<p>A1</p>", Deck: "D", Tags: []string{"go", "x"}},
		{Front: "Q2", Back: "<p>A2</p>", Deck: "D"},
	}, options{})

	got := buf.String()
	for _, want := range []string{"── Card 1", "── Card 2", "Q1", "<p>A2</p>", "go, x"} {
		if !strings.Contains(got, want) {
			t.Errorf("preview missing %q:\n%s", want, got)
		}
	}
}

func TestPreviewAsJSONIsAValidAddNotesRequest(t *testing.T) {
	buf := captureOut(t)
	preview([]Card{{Front: "Q", Back: "A", Deck: "D"}}, options{asJSON: true, model: "Principal SWE"})

	var req struct {
		Action  string `json:"action"`
		Version int    `json:"version"`
		Params  struct {
			Notes []Note `json:"notes"`
		} `json:"params"`
	}
	if err := json.Unmarshal(buf.Bytes(), &req); err != nil {
		t.Fatalf("payload is not valid JSON: %v\n%s", err, buf.String())
	}
	if req.Action != "addNotes" || req.Version != 6 || len(req.Params.Notes) != 1 {
		t.Errorf("payload = %+v", req)
	}
}

// ─── run ───

func TestRunWritesAnHTMLPreview(t *testing.T) {
	captureOut(t)
	root := vaultWith(t, map[string]string{"Note.md": "## 1. X\n\n```go\nx := 1\n```\n"})
	target := filepath.Join(t.TempDir(), "preview.html")

	err := run(options{targets: []string{root}, htmlOut: target, model: "Principal SWE"})
	if err != nil {
		t.Fatal(err)
	}

	page, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"<!doctype html>", "language-go", ".tok-kw", "answer-body"} {
		if !strings.Contains(string(page), want) {
			t.Errorf("preview page missing %q", want)
		}
	}
}

func TestRunDryRunTouchesNoServer(t *testing.T) {
	buf := captureOut(t)
	root := vaultWith(t, map[string]string{"Note.md": "## 1. X\n\nBody.\n"})

	// An unreachable URL proves nothing was sent.
	err := run(options{targets: []string{root}, dryRun: true, url: "http://127.0.0.1:1"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "Front: Note — 1. X") {
		t.Errorf("dry run did not print the card:\n%s", buf.String())
	}
}

func TestRunImportsThroughAnkiConnect(t *testing.T) {
	buf := captureOut(t)

	fake := newFakeAnki(t).
		reply("modelNames", []string{"Principal SWE"}).
		reply("modelFieldNames", []string{"Front", "Back"}).
		reply("modelTemplates", map[string]any{cardTemplateName: map[string]any{}}).
		reply("updateModelStyling", nil).
		reply("updateModelTemplates", nil).
		reply("createDeck", 1).
		reply("canAddNotes", []bool{true, false}).
		reply("addNotes", []any{111}).
		reply("findNotes", []int64{222}).
		reply("updateNoteFields", nil)

	root := vaultWith(t, map[string]string{
		"Note.md": "## 1. New\n\nBody.\n\n## 2. Old\n\nBody.\n",
	})

	err := run(options{
		targets: []string{root},
		model:   "Principal SWE",
		url:     fake.server.URL,
		update:  true,
	})
	if err != nil {
		t.Fatal(err)
	}

	got := buf.String()
	if !strings.Contains(got, "✅ Added: 1") {
		t.Errorf("report = %q", got)
	}
	if !strings.Contains(got, "Updated (already existed): 1") {
		t.Errorf("the existing card was not refreshed:\n%s", got)
	}

	// Only the addable note may reach addNotes.
	sent := fake.paramsOf("addNotes")["notes"].([]any)
	if len(sent) != 1 {
		t.Fatalf("sent %d notes, want 1", len(sent))
	}
	front := sent[0].(map[string]any)["fields"].(map[string]any)["Front"]
	if front != "Note — 1. New" {
		t.Errorf("wrong note sent: %v", front)
	}
}

func TestRunReportsImportFailures(t *testing.T) {
	captureOut(t)
	root := vaultWith(t, map[string]string{"Note.md": "## 1. X\n\nBody.\n"})

	err := run(options{targets: []string{root}, url: "http://127.0.0.1:1"})
	if err == nil || !strings.Contains(err.Error(), "AnkiConnect unreachable") {
		t.Fatalf("err = %v", err)
	}
}

func TestRunPropagatesParseFailures(t *testing.T) {
	captureOut(t)
	if err := run(options{targets: []string{"/no/such/note.md"}}); err == nil {
		t.Fatal("want an error")
	}
}

func TestImportCardsStopsWhenTheDeckCannotBeCreated(t *testing.T) {
	captureOut(t)
	fake := newFakeAnki(t).
		reply("modelNames", []string{"M"}).
		reply("modelFieldNames", []string{"Front", "Back"}).
		reply("modelTemplates", map[string]any{cardTemplateName: map[string]any{}}).
		reply("updateModelStyling", nil).
		reply("updateModelTemplates", nil)
	// createDeck is not registered, so it fails.

	err := importCards([]Card{{Front: "Q", Back: "A", Deck: "D"}}, options{model: "M", url: fake.server.URL})
	if err == nil {
		t.Fatal("want the deck error to stop the import")
	}
}

func TestImportCardsStopsWhenTheDuplicateCheckFails(t *testing.T) {
	captureOut(t)
	fake := newFakeAnki(t).
		reply("modelNames", []string{"M"}).
		reply("modelFieldNames", []string{"Front", "Back"}).
		reply("modelTemplates", map[string]any{cardTemplateName: map[string]any{}}).
		reply("updateModelStyling", nil).
		reply("updateModelTemplates", nil).
		reply("createDeck", 1)

	err := importCards([]Card{{Front: "Q", Back: "A", Deck: "D"}}, options{model: "M", url: fake.server.URL})
	if err == nil {
		t.Fatal("want the canAddNotes error to stop the import")
	}
}

func TestFatalExitsWithAnErrorCode(t *testing.T) {
	stderr := os.Stderr
	devNull, _ := os.Open(os.DevNull)
	os.Stderr = devNull
	defer func() { os.Stderr = stderr; devNull.Close() }()

	code := 0
	oldExit := exit
	exit = func(c int) { code = c }
	defer func() { exit = oldExit }()

	fatal(errString("boom"))

	if code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
}
