package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeNote drops a markdown file into a temp dir and returns its path.
func writeNote(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Note.md")
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func parse(t *testing.T, body string) []Card {
	t.Helper()
	cards, err := ParseNote(writeNote(t, body), "Deck", nil)
	if err != nil {
		t.Fatal(err)
	}
	return cards
}

func TestOneCardPerNumberedHeading(t *testing.T) {
	cards := parse(t, `---
title: "CSP in Go"
tags:
  - golang
  - concurrency
---
# CSP in Go

Intro text that belongs to no card.

## 1. Channels

Channels carry values.

## 2. Select

Select waits on many channels.

### 2.1 Default case

A default makes it non-blocking.
`)

	if len(cards) != 3 {
		t.Fatalf("got %d cards, want 3", len(cards))
	}

	wantFronts := []string{
		"CSP in Go — 1. Channels",
		"CSP in Go — 2. Select",
		"CSP in Go — 2.1. Default case",
	}
	for i, want := range wantFronts {
		if cards[i].Front != want {
			t.Errorf("card %d front = %q, want %q", i, cards[i].Front, want)
		}
	}

	// Text above the first numbered heading is not part of any card.
	for _, c := range cards {
		if strings.Contains(c.Back, "belongs to no card") {
			t.Errorf("intro text leaked into %q", c.Front)
		}
	}

	// A section stops at the next numbered heading.
	if strings.Contains(cards[1].Back, "non-blocking") {
		t.Error("card 2 swallowed the body of card 2.1")
	}

	wantTags := []string{"golang", "concurrency"}
	for i, want := range wantTags {
		if cards[0].Tags[i] != want {
			t.Errorf("tag %d = %q, want %q", i, cards[0].Tags[i], want)
		}
	}
}

func TestFrontmatterInlineTagsAndFallbackTitle(t *testing.T) {
	cards := parse(t, `---
tags: [go, "runtime"]
---
## 1. Scheduler

The scheduler multiplexes goroutines.
`)

	if len(cards) != 1 {
		t.Fatalf("got %d cards, want 1", len(cards))
	}
	// No title in the frontmatter — the file name is used.
	if want := "Note — 1. Scheduler"; cards[0].Front != want {
		t.Errorf("front = %q, want %q", cards[0].Front, want)
	}
	if got := strings.Join(cards[0].Tags, ","); got != "go,runtime" {
		t.Errorf("tags = %q, want %q", got, "go,runtime")
	}
}

func TestReferencesSectionIsDropped(t *testing.T) {
	cards := parse(t, `## 1. Goroutines

A goroutine is cheap.

## 🔗 References
- ⬆️ Parent: `+"`Goroutines & Memory Lifecycle`"+`
- 📚 Module: `+"`Concurrency & Synchronization`"+`
`)

	if len(cards) != 1 {
		t.Fatalf("got %d cards, want 1", len(cards))
	}
	for _, bad := range []string{"Parent:", "Module:", "References"} {
		if strings.Contains(cards[0].Back, bad) {
			t.Errorf("reference section leaked %q into the card", bad)
		}
	}
}

func TestSkippedHeadingVariants(t *testing.T) {
	skipped := []string{
		"## References", "### see also", "## 🔗 Related Notes",
		"## **Sources**", "# Further Reading",
	}
	for _, line := range skipped {
		if !isSkippedHeading(line) {
			t.Errorf("isSkippedHeading(%q) = false, want true", line)
		}
	}

	kept := []string{
		"## 1. References in Go", "## Reference Counting", "plain text", "## Channels",
	}
	for _, line := range kept {
		if isSkippedHeading(line) {
			t.Errorf("isSkippedHeading(%q) = true, want false", line)
		}
	}
}

func TestHeadingWithEmptyBodyMakesNoCard(t *testing.T) {
	cards := parse(t, `## 1. Empty

## 2. Real

Body here.
`)
	if len(cards) != 1 {
		t.Fatalf("got %d cards, want 1", len(cards))
	}
	if !strings.Contains(cards[0].Front, "2. Real") {
		t.Errorf("kept the wrong card: %q", cards[0].Front)
	}
}

func TestCodeFenceIsNotParsedAsMarkdown(t *testing.T) {
	cards := parse(t, "## 1. Comments\n\n```go\n// 2. This is a comment, not a heading\nx := *p\nch <- 1 // a < b\n```\n")

	if len(cards) != 1 {
		t.Fatalf("got %d cards, want 1 — a heading inside a fence started a new card", len(cards))
	}

	back := cards[0].Back
	if !strings.Contains(back, `<pre><code class="language-go">`) {
		t.Errorf("missing language-tagged code block:\n%s", back)
	}
	if strings.Contains(back, "<i>p") {
		t.Error("italic markdown was applied inside a code block")
	}
	if !strings.Contains(back, "ch &lt;- 1") {
		t.Errorf("code was not HTML-escaped:\n%s", back)
	}
}

func TestLooseListBecomesOneList(t *testing.T) {
	cards := parse(t, `## 1. Sync primitives

- mutexes

- atomics

- channels
`)

	back := cards[0].Back
	if got := strings.Count(back, "<ul>"); got != 1 {
		t.Errorf("got %d <ul> elements, want 1:\n%s", got, back)
	}
	if got := strings.Count(back, "<li>"); got != 3 {
		t.Errorf("got %d list items, want 3", got)
	}
}

func TestInlineMarkdown(t *testing.T) {
	cases := map[string]string{
		"**bold**":                  "<b>bold</b>",
		"*slanted*":                 "<i>slanted</i>",
		"`make(chan int)`":          "<code>make(chan int)</code>",
		"[[Goroutines]]":            "Goroutines",
		"[[Goroutines|the note]]":   "the note",
		"[docs](https://go.dev)":    `<a href="https://go.dev">docs</a>`,
		"a < b & c > d":             "a &lt; b &amp; c &gt; d",
		"see [[A|B]] and **C** now": "see B and <b>C</b> now",
	}
	for in, want := range cases {
		if got := inline(in); !strings.Contains(got, want) {
			t.Errorf("inline(%q) = %q, want it to contain %q", in, got, want)
		}
	}
}

func TestBlockquoteAndUnnumberedHeading(t *testing.T) {
	cards := parse(t, `## 1. Principle

### Rule of thumb

> Share memory by communicating.
`)

	back := cards[0].Back
	if !strings.Contains(back, "<b>Rule of thumb</b>") {
		t.Errorf("unnumbered heading should render bold inside the card:\n%s", back)
	}
	if !strings.Contains(back, "<blockquote>Share memory by communicating.</blockquote>") {
		t.Errorf("blockquote not rendered:\n%s", back)
	}
}
