package main

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Card is one question/answer pair cut out of a note.
type Card struct {
	Front  string // "Note title — 3. Section heading"
	Back   string // section body, rendered to HTML
	Tags   []string
	Deck   string
	Title  string // note the card was cut from
	Source string
}

var (
	headingRe  = regexp.MustCompile(`^(#{1,6})\s+(.*)$`)
	numberedRe = regexp.MustCompile(`^(\d+(?:\.\d+)*)[.):]?\s+(.+)$`)
	fenceRe    = regexp.MustCompile("^\\s*(```|~~~)")
	htmlTagRe  = regexp.MustCompile(`<[^>]*>`)
)

// ParseNote cuts one markdown file into cards — one card per numbered heading
// (`## 3. Channels`), holding everything up to the next numbered heading.
func ParseNote(path, deck string, extraTags []string) ([]Card, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	meta, body := splitFrontmatter(strings.Split(string(raw), "\n"))

	title := meta.Title
	if title == "" {
		title = strings.TrimSuffix(filepath.Base(path), ".md")
	}
	tags := dedupe(append(append([]string{}, meta.Tags...), extraTags...))

	var (
		cards   []Card
		trail   []headingRef // the numbered heading and its numbered ancestors
		heading string
		section []string
		inFence bool
	)

	// flush turns the section collected so far into a card.
	flush := func() {
		defer func() { section = nil }()
		if heading == "" {
			return
		}
		back := renderHTML(section)
		if strings.TrimSpace(stripHTML(back)) == "" {
			return // heading with no body — nothing to ask
		}
		cards = append(cards, Card{
			Front:  fmt.Sprintf("%s — %s", title, heading),
			Back:   back,
			Tags:   tags,
			Deck:   deck,
			Title:  title,
			Source: path,
		})
	}

	for _, line := range body {
		if fenceRe.MatchString(line) {
			inFence = !inFence
		}
		if !inFence {
			if level, number, text, ok := numberedHeading(line); ok {
				flush()
				// A sub-heading is a card of its own, but its question only
				// makes sense under the section it sits in: "… — 7. The Four
				// Conditions for Deadlock — 1. Mutual exclusion".
				for len(trail) > 0 && trail[len(trail)-1].level >= level {
					trail = trail[:len(trail)-1]
				}
				trail = append(trail, headingRef{level: level, number: number, text: number + ". " + text})
				heading = joinTrail(trail)
				continue
			}
			// Navigation sections (References, See also, …) carry no content
			// worth reviewing — end the card here.
			if isSkippedHeading(line) {
				flush()
				heading, trail = "", nil
				continue
			}
		}
		if heading != "" {
			section = append(section, line)
		}
	}
	flush()
	cards = dropParentSections(cards)

	return cards, nil
}

// dropParentSections removes the card of a section whose sub-sections became
// cards of their own: what is left of it is only the sentence introducing them
// ("Deadlock requires all four:"), which is nothing to review.
func dropParentSections(cards []Card) []Card {
	kept := cards[:0]
	for i, c := range cards {
		hasChild := i+1 < len(cards) && strings.HasPrefix(cards[i+1].Front, c.Front+" — ")
		if !hasChild {
			kept = append(kept, c)
		}
	}
	return kept
}

// headingRef is one numbered heading on the path down to the current card.
type headingRef struct {
	level  int    // 1 for "#", 2 for "##", …
	number string // "3", or "2.1" for a sub-section
	text   string // "3. Channels"
}

// joinTrail spells out the heading path a card sits under. An ancestor whose
// number the child already repeats ("2. Select" above "2.1. Default case") is
// left out — the child says where it belongs on its own.
func joinTrail(trail []headingRef) string {
	parts := make([]string, 0, len(trail))
	for i, h := range trail {
		if i+1 < len(trail) && strings.HasPrefix(trail[i+1].number, h.number+".") {
			continue
		}
		parts = append(parts, h.text)
	}
	return strings.Join(parts, " — ")
}

// numberedHeading matches "## 3. Channels" and returns (2, "3", "Channels", true).
func numberedHeading(line string) (level int, number, text string, ok bool) {
	h := headingRe.FindStringSubmatch(line)
	if h == nil {
		return 0, "", "", false
	}
	n := numberedRe.FindStringSubmatch(strings.TrimSpace(h[2]))
	if n == nil {
		return 0, "", "", false
	}
	return len(h[1]), n[1], strings.TrimSpace(n[2]), true
}

// skippedHeadings are section titles that only link elsewhere in the vault.
// Their body never becomes part of a card.
var skippedHeadings = []string{
	"references", "reference", "see also", "related", "related notes",
	"links", "sources", "further reading", "backlinks",
}

// isSkippedHeading reports whether the line opens a navigation-only section,
// ignoring any leading emoji or decoration in the title.
func isSkippedHeading(line string) bool {
	h := headingRe.FindStringSubmatch(line)
	if h == nil {
		return false
	}

	title := strings.ToLower(strings.TrimSpace(h[2]))
	title = strings.Map(func(r rune) rune {
		if r > 127 { // drop emoji and other decoration
			return -1
		}
		return r
	}, title)
	title = strings.TrimSpace(strings.Trim(title, ":*_ "))

	for _, skip := range skippedHeadings {
		if title == skip {
			return true
		}
	}
	return false
}

// ─── frontmatter ───

type frontmatter struct {
	Title string
	Tags  []string
}

// splitFrontmatter pulls title and tags out of the leading YAML block and
// returns the remaining body lines.
func splitFrontmatter(lines []string) (frontmatter, []string) {
	var meta frontmatter
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return meta, lines
	}

	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			end = i
			break
		}
	}
	if end == -1 {
		return meta, lines
	}

	inTagList := false
	for _, line := range lines[1:end] {
		trimmed := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(trimmed, "title:"):
			meta.Title = unquote(strings.TrimPrefix(trimmed, "title:"))
			inTagList = false

		case strings.HasPrefix(trimmed, "tags:"):
			// Inline form: tags: [go, concurrency]
			inTagList = true
			inline := strings.Trim(strings.TrimSpace(strings.TrimPrefix(trimmed, "tags:")), "[]")
			for _, t := range strings.Split(inline, ",") {
				if t = unquote(t); t != "" {
					meta.Tags = append(meta.Tags, t)
				}
			}

		case inTagList && strings.HasPrefix(trimmed, "- "):
			// List form: one "- tag" per line.
			meta.Tags = append(meta.Tags, unquote(strings.TrimPrefix(trimmed, "- ")))

		default:
			if !strings.HasPrefix(line, " ") {
				inTagList = false
			}
		}
	}
	return meta, lines[end+1:]
}

func unquote(s string) string { return strings.Trim(strings.TrimSpace(s), `"'`) }

// ─── markdown → HTML ───

// renderHTML converts a section body into the small HTML subset Anki renders
// well: paragraphs, lists, blockquotes and fenced code blocks.
func renderHTML(lines []string) string {
	r := &htmlRenderer{}
	for _, line := range lines {
		r.line(line)
	}
	return r.done()
}

type htmlRenderer struct {
	out     []string
	inFence bool
	inList  bool
}

func (r *htmlRenderer) line(line string) {
	if fenceRe.MatchString(line) {
		r.toggleFence(line)
		return
	}
	if r.inFence {
		r.codeLine(line)
		return
	}

	trimmed := strings.TrimSpace(line)
	switch {
	case trimmed == "":
		// Blank lines do not end a list — Obsidian writes loose lists.

	case trimmed == "---":
		r.closeList()

	case strings.HasPrefix(trimmed, "> "):
		r.closeList()
		r.write("<blockquote>" + inline(strings.TrimPrefix(trimmed, "> ")) + "</blockquote>")

	case strings.HasPrefix(trimmed, "- "), strings.HasPrefix(trimmed, "* "):
		r.openList()
		r.write("<li>" + inline(trimmed[2:]) + "</li>")

	case headingRe.MatchString(trimmed):
		r.closeList()
		h := headingRe.FindStringSubmatch(trimmed)
		r.write("<b>" + inline(h[2]) + "</b>")

	default:
		r.closeList()
		r.write("<p>" + inline(trimmed) + "</p>")
	}
}

func (r *htmlRenderer) toggleFence(line string) {
	if r.inFence {
		r.write("</code></pre>")
		r.inFence = false
		return
	}
	r.closeList()

	lang := strings.TrimSpace(strings.TrimLeft(strings.TrimSpace(line), "`~"))
	class := ""
	if lang != "" {
		class = fmt.Sprintf(` class="language-%s"`, html.EscapeString(lang))
	}
	r.write("<pre><code" + class + ">")
	r.inFence = true
}

func (r *htmlRenderer) codeLine(line string) {
	// Drop a blank line right after the opening fence.
	if strings.TrimSpace(line) == "" && strings.Contains(r.lastWritten(), "<pre><code") {
		return
	}
	r.write(html.EscapeString(line))
}

func (r *htmlRenderer) openList() {
	if !r.inList {
		r.write("<ul>")
		r.inList = true
	}
}

func (r *htmlRenderer) closeList() {
	if r.inList {
		r.write("</ul>")
		r.inList = false
	}
}

func (r *htmlRenderer) done() string {
	r.closeList()
	if r.inFence {
		r.write("</code></pre>")
	}
	return strings.Join(r.out, "\n")
}

func (r *htmlRenderer) write(s string) { r.out = append(r.out, s) }

// lastWritten is the most recent line of output, or "" when nothing was written.
func (r *htmlRenderer) lastWritten() string {
	if len(r.out) == 0 {
		return ""
	}
	return r.out[len(r.out)-1]
}

// ─── inline markdown ───

var (
	wikiRe   = regexp.MustCompile(`\[\[([^\]|]+)(?:\|([^\]]+))?\]\]`)
	linkRe   = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	boldRe   = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	italicRe = regexp.MustCompile(`(^|[^*])\*([^*\n]+)\*`)
	codeRe   = regexp.MustCompile("`([^`\n]+)`")
)

// inline escapes the text, then re-applies the inline markdown constructs:
// wiki links, links, `code`, **bold**, *italic*.
func inline(s string) string {
	s = html.EscapeString(s)
	s = wikiRe.ReplaceAllStringFunc(s, func(m string) string {
		g := wikiRe.FindStringSubmatch(m)
		if g[2] != "" {
			return g[2] // [[target|alias]] → alias
		}
		return g[1]
	})
	s = linkRe.ReplaceAllString(s, `<a href="$2">$1</a>`)
	s = codeRe.ReplaceAllString(s, "<code>$1</code>")
	s = boldRe.ReplaceAllString(s, "<b>$1</b>")
	s = italicRe.ReplaceAllString(s, "$1<i>$2</i>")
	return s
}

func stripHTML(s string) string { return htmlTagRe.ReplaceAllString(s, "") }

func dedupe(in []string) []string {
	seen := make(map[string]bool, len(in))
	var out []string
	for _, s := range in {
		if s = strings.TrimSpace(s); s != "" && !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}
