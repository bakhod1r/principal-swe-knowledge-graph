package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ─── markdown edge cases ───

func TestParseNoteReportsUnreadableFiles(t *testing.T) {
	if _, err := ParseNote(filepath.Join(t.TempDir(), "missing.md"), "D", nil); err == nil {
		t.Fatal("want an error for a missing file")
	}
}

func TestFrontmatterEdgeCases(t *testing.T) {
	t.Run("no frontmatter", func(t *testing.T) {
		meta, body := splitFrontmatter([]string{"# Title", "text"})
		if meta.Title != "" || len(body) != 2 {
			t.Errorf("meta = %+v, body = %v", meta, body)
		}
	})

	t.Run("empty file", func(t *testing.T) {
		meta, body := splitFrontmatter(nil)
		if meta.Title != "" || body != nil {
			t.Errorf("meta = %+v, body = %v", meta, body)
		}
	})

	t.Run("unterminated block is treated as body", func(t *testing.T) {
		lines := []string{"---", "title: X", "still open"}
		meta, body := splitFrontmatter(lines)
		if meta.Title != "" {
			t.Errorf("title = %q, want none — the block never closed", meta.Title)
		}
		if len(body) != 3 {
			t.Errorf("body = %v, want every line kept", body)
		}
	})

	t.Run("a following key ends the tag list", func(t *testing.T) {
		meta, _ := splitFrontmatter([]string{
			"---", "tags:", "  - go", "parent: \"[[X]]\"", "- not-a-tag", "---", "body",
		})
		if strings.Join(meta.Tags, ",") != "go" {
			t.Errorf("tags = %v, want just [go]", meta.Tags)
		}
	})

	t.Run("indented lines stay inside the tag list", func(t *testing.T) {
		meta, _ := splitFrontmatter([]string{
			"---", "tags:", "  - go", "  # a comment", "  - runtime", "---",
		})
		if strings.Join(meta.Tags, ",") != "go,runtime" {
			t.Errorf("tags = %v", meta.Tags)
		}
	})

	t.Run("title after tags", func(t *testing.T) {
		meta, _ := splitFrontmatter([]string{"---", "tags: [a]", "title: 'T'", "- x", "---"})
		if meta.Title != "T" || strings.Join(meta.Tags, ",") != "a" {
			t.Errorf("meta = %+v — the title should end the tag list", meta)
		}
	})
}

func TestUnterminatedCodeFenceIsClosed(t *testing.T) {
	cards := parse(t, "## 1. Broken\n\n```go\nx := 1\n")

	back := cards[0].Back
	if strings.Count(back, "<pre><code") != 1 || !strings.HasSuffix(back, "</code></pre>") {
		t.Errorf("unterminated fence not closed:\n%s", back)
	}
}

func TestBlankLineAfterTheOpeningFenceIsDropped(t *testing.T) {
	cards := parse(t, "## 1. Diagram\n\n```text\n\nA --> B\n```\n")

	back := cards[0].Back
	if strings.Contains(back, "<pre><code class=\"language-text\">\n\n") {
		t.Errorf("leading blank line kept inside the code block:\n%s", back)
	}
	if !strings.Contains(back, "A --&gt; B") {
		t.Errorf("code body lost:\n%s", back)
	}
}

func TestFenceWithoutALanguage(t *testing.T) {
	cards := parse(t, "## 1. Plain\n\n```\nsome text\n```\n")

	if !strings.Contains(cards[0].Back, "<pre><code>") {
		t.Errorf("bare fence should produce an untagged code block:\n%s", cards[0].Back)
	}
}

func TestListClosedByAHorizontalRuleAndByAHeading(t *testing.T) {
	cards := parse(t, `## 1. Mixed

- one
- two

---

### After

- three
`)

	back := cards[0].Back
	if got := strings.Count(back, "<ul>"); got != 2 {
		t.Errorf("got %d lists, want 2 (the rule and heading close the first):\n%s", got, back)
	}
	if got := strings.Count(back, "</ul>"); got != 2 {
		t.Errorf("unbalanced list markup:\n%s", back)
	}
}

func TestStarBulletsAreListsToo(t *testing.T) {
	cards := parse(t, "## 1. Stars\n\n* one\n* two\n")
	if got := strings.Count(cards[0].Back, "<li>"); got != 2 {
		t.Errorf("got %d items:\n%s", got, cards[0].Back)
	}
}

func TestRenderHTMLOfNothing(t *testing.T) {
	if got := renderHTML(nil); got != "" {
		t.Errorf("renderHTML(nil) = %q", got)
	}
}

func TestStripHTML(t *testing.T) {
	if got := stripHTML("<p>a</p><br><b>b</b>"); got != "ab" {
		t.Errorf("stripHTML = %q", got)
	}
}

// ─── vault edge cases ───

func TestDeckFromPathFlattensDeckSeparatorsInFolderNames(t *testing.T) {
	got := DeckFromPath("/x/Principal SWE/A::B/Note.md")
	if got != "Principal SWE::A-B" {
		t.Errorf("DeckFromPath = %q, want the folder's :: flattened", got)
	}
}

func TestCollectFilesReportsWalkErrors(t *testing.T) {
	root := t.TempDir()
	locked := filepath.Join(root, "locked")
	if err := os.Mkdir(locked, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(locked, "a.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(locked, 0o000); err != nil {
		t.Skip("cannot drop directory permissions here")
	}
	t.Cleanup(func() { os.Chmod(locked, 0o755) })

	if os.Geteuid() == 0 {
		t.Skip("running as root — permissions are not enforced")
	}
	if _, err := CollectFiles([]string{root}); err == nil {
		t.Error("want the walk error to surface")
	}
}

// ─── transport edge cases ───

func TestCallReportsAMalformedResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer server.Close()

	if err := NewClient(server.URL).call("version", nil, nil); err == nil {
		t.Fatal("want a decode error")
	}
}

func TestCallReportsAResultOfTheWrongShape(t *testing.T) {
	fake := newFakeAnki(t).reply("modelNames", "a string, not a list")

	var models []string
	if err := fake.client().call("modelNames", nil, &models); err == nil {
		t.Fatal("want an unmarshal error")
	}
}

// ─── import edge cases ───

func TestImportCardsStopsWhenAddNotesFails(t *testing.T) {
	captureOut(t)
	fake := newFakeAnki(t).
		reply("modelNames", []string{"M"}).
		reply("modelFieldNames", []string{"Front", "Back"}).
		reply("modelTemplates", map[string]any{cardTemplateName: map[string]any{}}).
		reply("updateModelStyling", nil).
		reply("updateModelTemplates", nil).
		reply("createDeck", 1).
		reply("canAddNotes", []bool{true})
	// addNotes is not registered, so it fails.

	err := importCards([]Card{{Front: "Q", Back: "A", Deck: "D"}}, options{model: "M", url: fake.server.URL})
	if err == nil {
		t.Fatal("want the addNotes error to stop the import")
	}
}

func TestImportCardsStopsWhenTheModelCannotBeResolved(t *testing.T) {
	captureOut(t)
	fake := newFakeAnki(t) // modelNames is not registered

	err := importCards([]Card{{Front: "Q", Deck: "D"}}, options{model: "M", url: fake.server.URL})
	if err == nil {
		t.Fatal("want the model error to stop the import")
	}
}

func TestImportCardsSurvivesAFailedUpdate(t *testing.T) {
	buf := captureOut(t)
	fake := newFakeAnki(t).
		reply("modelNames", []string{"M"}).
		reply("modelFieldNames", []string{"Front", "Back"}).
		reply("modelTemplates", map[string]any{cardTemplateName: map[string]any{}}).
		reply("updateModelStyling", nil).
		reply("updateModelTemplates", nil).
		reply("createDeck", 1).
		reply("canAddNotes", []bool{false}).
		reply("findNotes", []int64{}) // nothing matches → update fails

	err := importCards(
		[]Card{{Front: "Q", Back: "A", Deck: "D"}},
		options{model: "M", url: fake.server.URL, update: true},
	)
	if err != nil {
		t.Fatalf("a failed update must not abort the run: %v", err)
	}
	if !strings.Contains(buf.String(), "✅ Added: 0") {
		t.Errorf("report missing:\n%s", buf.String())
	}
}

func TestRunReportsAnUnwritableHTMLTarget(t *testing.T) {
	captureOut(t)
	root := vaultWith(t, map[string]string{"Note.md": "## 1. X\n\nBody.\n"})

	err := run(options{targets: []string{root}, htmlOut: filepath.Join(root, "no-dir", "x.html")})
	if err == nil {
		t.Fatal("want an error for an unwritable preview path")
	}
}

func TestBuildCardsReportsUnreadableNotes(t *testing.T) {
	root := t.TempDir()
	note := filepath.Join(root, "a.md")
	if err := os.WriteFile(note, []byte("## 1. X\n\nBody.\n"), 0o000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chmod(note, 0o644) })

	if os.Geteuid() == 0 {
		t.Skip("running as root — permissions are not enforced")
	}
	if _, _, err := buildCards(options{targets: []string{root}}); err == nil {
		t.Error("want the read error to surface")
	}
}

// ─── the last uncovered branches ───

func TestLastWrittenOnAnEmptyRenderer(t *testing.T) {
	r := &htmlRenderer{}
	if got := r.lastWritten(); got != "" {
		t.Errorf("lastWritten = %q, want empty", got)
	}
}

func TestMainRunsEndToEnd(t *testing.T) {
	buf := captureOut(t)

	root := vaultWith(t, map[string]string{"Note.md": "## 1. X\n\nBody.\n"})

	args := os.Args
	os.Args = []string{"md-to-anki", root, "--dry-run"}
	defer func() { os.Args = args }()

	code := 0
	oldExit := exit
	exit = func(c int) { code = c }
	defer func() { exit = oldExit }()

	main()

	if code != 0 {
		t.Errorf("exit code = %d, want 0", code)
	}
	if !strings.Contains(buf.String(), "Front: Note — 1. X") {
		t.Errorf("main did not run the job:\n%s", buf.String())
	}
}

func TestMainExitsOnBadArguments(t *testing.T) {
	captureOut(t)

	args := os.Args
	os.Args = []string{"md-to-anki"} // no target
	defer func() { os.Args = args }()

	stderr := os.Stderr
	devNull, _ := os.Open(os.DevNull)
	os.Stderr = devNull
	defer func() { os.Stderr = stderr; devNull.Close() }()

	code := 0
	oldExit := exit
	exit = func(c int) { code = c }
	defer func() { exit = oldExit }()

	main()

	if code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
}
