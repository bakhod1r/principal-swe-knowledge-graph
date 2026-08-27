package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWritePreviewHTML(t *testing.T) {
	target := filepath.Join(t.TempDir(), "out.html")

	cards := []Card{
		{
			Front: "Note — 1. Channels",
			Back:  `<p>Body</p><pre><code class="language-go">x := 1</code></pre>`,
			Deck:  "Principal SWE::Go",
		},
		{Front: "Note — 2. Select", Back: "<p>Body</p>", Deck: "Principal SWE::Go"},
	}
	if err := WritePreviewHTML(target, cards, "Principal SWE"); err != nil {
		t.Fatal(err)
	}

	raw, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	page := string(raw)

	for _, want := range []string{
		"<!doctype html>",
		"2 cards · note type Principal SWE",
		"Note — 1. Channels",
		`class="language-go"`,
		".tok-str",         // the card CSS is inlined
		"answer-body",      // the template scripts run over every card
		"querySelectorAll", // the preview driver is present
	} {
		if !strings.Contains(page, want) {
			t.Errorf("preview page is missing %q", want)
		}
	}

	// The deck name is shown once per card, escaped.
	if got := strings.Count(page, "Principal SWE::Go"); got != 2 {
		t.Errorf("deck label appears %d times, want 2", got)
	}
}

func TestWritePreviewHTMLEscapesCardMetadata(t *testing.T) {
	target := filepath.Join(t.TempDir(), "out.html")

	cards := []Card{{Front: `<script>alert(1)</script>`, Back: "<p>ok</p>", Deck: "A & B"}}
	if err := WritePreviewHTML(target, cards, "M"); err != nil {
		t.Fatal(err)
	}

	page, _ := os.ReadFile(target)
	if strings.Contains(string(page), "<script>alert(1)</script>") {
		t.Error("the front field was not escaped into the page")
	}
	if !strings.Contains(string(page), "A &amp; B") {
		t.Error("the deck name was not escaped")
	}
}

func TestWritePreviewHTMLReportsWriteErrors(t *testing.T) {
	err := WritePreviewHTML(filepath.Join(t.TempDir(), "no-such-dir", "out.html"), nil, "M")
	if err == nil {
		t.Fatal("want an error for an unwritable path")
	}
}

func TestStripScriptTags(t *testing.T) {
	if got := stripScriptTags("<div></div><script>let x = 1;</script>"); got != "let x = 1;" {
		t.Errorf("stripScriptTags = %q", got)
	}
	if got := stripScriptTags("no script here"); got != "" {
		t.Errorf("stripScriptTags = %q, want empty", got)
	}
}

func TestPreviewScriptCarriesBothTemplates(t *testing.T) {
	script := previewScript()

	if !strings.Contains(script, "tok-kw") {
		t.Error("the highlighter from the back template is missing")
	}
	if !strings.Contains(script, "note-title") {
		t.Error("the front template's title split is missing")
	}
	if strings.Contains(script, "<script>") {
		t.Error("the template's own <script> tags leaked into the page script")
	}
}

func TestCardTemplatesAreSelfContained(t *testing.T) {
	for name, tpl := range map[string]string{"front": frontTemplate, "back": backTemplate} {
		if strings.Contains(tpl, "http://") || strings.Contains(tpl, "https://") {
			t.Errorf("%s template loads something over the network", name)
		}
	}
	if !strings.Contains(backTemplate, "{{FrontSide}}") || !strings.Contains(backTemplate, "{{Back}}") {
		t.Error("the back template does not render the note's fields")
	}
	if !strings.Contains(frontTemplate, "{{Front}}") {
		t.Error("the front template does not render the question")
	}
}
