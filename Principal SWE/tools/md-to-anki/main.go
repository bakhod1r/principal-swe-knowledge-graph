// md-to-anki — split an Obsidian note into Anki notes, one per numbered heading,
// and push them into Anki through the AnkiConnect add-on (http://127.0.0.1:8765).
//
// Usage:
//
//	md-to-anki <file.md|dir> [flags]
//
// Flags:
//
//	--deck    target deck name (default: derived from vault folder path)
//	--model   Anki note type (default "Basic")
//	--tags    extra comma-separated tags
//	--url     AnkiConnect endpoint (default http://127.0.0.1:8765)
//	--dry-run print the cards (and the JSON payload) instead of sending
//	--json    with --dry-run, dump the full AnkiConnect request
//	--update  update the note when a card with the same Front already exists
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// ─── Card ───

type Card struct {
	Front  string
	Back   string
	Tags   []string
	Deck   string
	Source string
}

// ─── Markdown parsing ───

var (
	headingRe  = regexp.MustCompile(`^(#{1,6})\s+(.*)$`)
	numberedRe = regexp.MustCompile(`^(\d+(?:\.\d+)*)[.):]?\s+(.+)$`)
	fenceRe    = regexp.MustCompile("^\\s*(```|~~~)")
	wikiRe     = regexp.MustCompile(`\[\[([^\]|]+)(?:\|([^\]]+))?\]\]`)
	linkRe     = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	boldRe     = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	italicRe   = regexp.MustCompile(`(^|[^*])\*([^*\n]+)\*`)
	codeRe     = regexp.MustCompile("`([^`\n]+)`")
)

type frontmatter struct {
	Title string
	Tags  []string
}

// splitFrontmatter returns the YAML frontmatter fields we care about and the body lines.
func splitFrontmatter(lines []string) (frontmatter, []string) {
	fm := frontmatter{}
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return fm, lines
	}
	end := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			end = i
			break
		}
	}
	if end == -1 {
		return fm, lines
	}

	inTags := false
	for _, l := range lines[1:end] {
		trimmed := strings.TrimSpace(l)
		switch {
		case strings.HasPrefix(trimmed, "title:"):
			fm.Title = unquote(strings.TrimSpace(strings.TrimPrefix(trimmed, "title:")))
			inTags = false
		case strings.HasPrefix(trimmed, "tags:"):
			inTags = true
			rest := strings.TrimSpace(strings.TrimPrefix(trimmed, "tags:"))
			rest = strings.Trim(rest, "[]")
			for _, t := range strings.Split(rest, ",") {
				if t = strings.TrimSpace(unquote(t)); t != "" {
					fm.Tags = append(fm.Tags, t)
				}
			}
		case inTags && strings.HasPrefix(trimmed, "- "):
			fm.Tags = append(fm.Tags, unquote(strings.TrimSpace(strings.TrimPrefix(trimmed, "- "))))
		default:
			if !strings.HasPrefix(l, " ") {
				inTags = false
			}
		}
	}
	return fm, lines[end+1:]
}

func unquote(s string) string {
	return strings.Trim(strings.TrimSpace(s), `"'`)
}

// parseNote splits one markdown file into cards — one card per numbered heading.
func parseNote(path string, deck string, extraTags []string) ([]Card, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(raw), "\n")
	fm, body := splitFrontmatter(lines)

	title := fm.Title
	if title == "" {
		title = strings.TrimSuffix(filepath.Base(path), ".md")
	}

	tags := append([]string{}, fm.Tags...)
	tags = append(tags, extraTags...)

	var (
		cards   []Card
		curHead string
		buf     []string
		inFence bool
	)

	flush := func() {
		if curHead == "" {
			return
		}
		back := renderHTML(buf)
		if strings.TrimSpace(stripTags(back)) == "" {
			curHead, buf = "", nil
			return
		}
		cards = append(cards, Card{
			Front:  fmt.Sprintf("%s — %s", title, curHead),
			Back:   back,
			Tags:   dedupe(tags),
			Deck:   deck,
			Source: path,
		})
		curHead, buf = "", nil
	}

	for _, l := range body {
		if fenceRe.MatchString(l) {
			inFence = !inFence
		}
		if !inFence {
			if m := headingRe.FindStringSubmatch(l); m != nil {
				if n := numberedRe.FindStringSubmatch(strings.TrimSpace(m[2])); n != nil {
					flush()
					curHead = fmt.Sprintf("%s. %s", n[1], strings.TrimSpace(n[2]))
					continue
				}
			}
		}
		if curHead != "" {
			buf = append(buf, l)
		}
	}
	flush()
	return cards, nil
}

func dedupe(in []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, s := range in {
		if s = strings.TrimSpace(s); s != "" && !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

var tagStripRe = regexp.MustCompile(`<[^>]*>`)

func stripTags(s string) string { return tagStripRe.ReplaceAllString(s, "") }

// renderHTML turns a markdown block into the small HTML subset Anki renders well.
func renderHTML(lines []string) string {
	var (
		out     []string
		inFence bool
		inList  bool
	)
	closeList := func() {
		if inList {
			out = append(out, "</ul>")
			inList = false
		}
	}

	for _, l := range lines {
		if fenceRe.MatchString(l) {
			if !inFence {
				closeList()
				lang := strings.TrimSpace(strings.TrimLeft(strings.TrimSpace(l), "`~"))
				cls := ""
				if lang != "" {
					cls = fmt.Sprintf(` class="language-%s"`, html.EscapeString(lang))
				}
				out = append(out, fmt.Sprintf("<pre><code%s>", cls))
				inFence = true
			} else {
				out = append(out, "</code></pre>")
				inFence = false
			}
			continue
		}
		if inFence {
			// Skip a blank line right after the opening fence.
			if strings.TrimSpace(l) == "" && strings.Contains(out[len(out)-1], "<pre><code") {
				continue
			}
			out = append(out, html.EscapeString(l))
			continue
		}

		trimmed := strings.TrimSpace(l)
		switch {
		case trimmed == "":
			// A blank line does not end a list — Obsidian writes loose lists.
		case trimmed == "---":
			closeList()
		case strings.HasPrefix(trimmed, "> "):
			closeList()
			out = append(out, "<blockquote>"+inline(strings.TrimPrefix(trimmed, "> "))+"</blockquote>")
		case strings.HasPrefix(trimmed, "- "), strings.HasPrefix(trimmed, "* "):
			if !inList {
				out = append(out, "<ul>")
				inList = true
			}
			out = append(out, "<li>"+inline(trimmed[2:])+"</li>")
		case headingRe.MatchString(trimmed):
			closeList()
			m := headingRe.FindStringSubmatch(trimmed)
			out = append(out, "<b>"+inline(m[2])+"</b>")
		default:
			closeList()
			out = append(out, "<p>"+inline(trimmed)+"</p>")
		}
	}
	closeList()
	if inFence {
		out = append(out, "</code></pre>")
	}
	return strings.Join(out, "\n")
}

// inline escapes text, then re-applies the inline markdown constructs.
func inline(s string) string {
	s = html.EscapeString(s)
	s = wikiRe.ReplaceAllStringFunc(s, func(m string) string {
		g := wikiRe.FindStringSubmatch(m)
		if g[2] != "" {
			return g[2]
		}
		return g[1]
	})
	s = linkRe.ReplaceAllString(s, `<a href="$2">$1</a>`)
	s = codeRe.ReplaceAllString(s, "<code>$1</code>")
	s = boldRe.ReplaceAllString(s, "<b>$1</b>")
	s = italicRe.ReplaceAllString(s, "$1<i>$2</i>")
	return s
}

// ─── AnkiConnect ───

type ankiRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  any    `json:"params,omitempty"`
}

type ankiResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *string         `json:"error"`
}

type ankiNote struct {
	DeckName  string            `json:"deckName"`
	ModelName string            `json:"modelName"`
	Fields    map[string]string `json:"fields"`
	Tags      []string          `json:"tags"`
	Options   map[string]any    `json:"options"`
}

func ankiCall(url, action string, params any, out any) error {
	body, err := json.Marshal(ankiRequest{Action: action, Version: 6, Params: params})
	if err != nil {
		return err
	}
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("AnkiConnect unreachable at %s (is Anki running with the AnkiConnect add-on?): %w", url, err)
	}
	defer resp.Body.Close()

	var ar ankiResponse
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		return err
	}
	if ar.Error != nil {
		return fmt.Errorf("AnkiConnect error on %s: %s", action, *ar.Error)
	}
	if out != nil {
		return json.Unmarshal(ar.Result, out)
	}
	return nil
}

// ─── Deck naming ───

// deckFromPath turns "01. Foundations/Programming/Golang/Note.md" into
// "Principal SWE::Foundations::Programming::Golang".
func deckFromPath(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		abs = path
	}
	parts := strings.Split(filepath.ToSlash(filepath.Dir(abs)), "/")
	root := -1
	for i, p := range parts {
		if p == "Principal SWE" {
			root = i
		}
	}
	if root == -1 {
		return "Principal SWE"
	}
	seg := []string{"Principal SWE"}
	numPrefix := regexp.MustCompile(`^\d+\.\s*`)
	for _, p := range parts[root+1:] {
		p = numPrefix.ReplaceAllString(p, "")
		p = strings.ReplaceAll(p, "::", "-")
		if p != "" {
			seg = append(seg, p)
		}
	}
	return strings.Join(seg, "::")
}

// ─── main ───

func main() {
	var (
		deck   = flag.String("deck", "", "target deck (default: derived from the note's folder path)")
		model  = flag.String("model", "Basic", "Anki note type")
		tagCSV = flag.String("tags", "", "extra comma-separated tags")
		url    = flag.String("url", "http://127.0.0.1:8765", "AnkiConnect endpoint")
		dryRun = flag.Bool("dry-run", false, "print the cards instead of sending them")
		asJSON = flag.Bool("json", false, "with --dry-run, dump the AnkiConnect request payload")
		update = flag.Bool("update", false, "update existing notes with the same Front instead of skipping")
	)
	flag.CommandLine.Parse(reorder(os.Args[1:]))

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: md-to-anki <file.md|dir> [--deck D] [--model M] [--tags a,b] [--dry-run] [--json] [--update]")
		os.Exit(1)
	}

	var extraTags []string
	for _, t := range strings.Split(*tagCSV, ",") {
		if t = strings.TrimSpace(t); t != "" {
			extraTags = append(extraTags, t)
		}
	}

	files, err := collectFiles(flag.Args())
	if err != nil {
		fatal(err)
	}
	if len(files) == 0 {
		fatal(fmt.Errorf("no .md files found"))
	}

	var cards []Card
	for _, f := range files {
		d := *deck
		if d == "" {
			d = deckFromPath(f)
		}
		c, err := parseNote(f, d, extraTags)
		if err != nil {
			fatal(err)
		}
		cards = append(cards, c...)
	}

	fmt.Printf("\n📚 md-to-anki\n")
	fmt.Printf("───────────────────────────────────────────────\n")
	fmt.Printf("Files: %d   Cards: %d   Model: %s\n\n", len(files), len(cards), *model)

	notes := make([]ankiNote, 0, len(cards))
	for _, c := range cards {
		notes = append(notes, ankiNote{
			DeckName:  c.Deck,
			ModelName: *model,
			Fields:    map[string]string{"Front": c.Front, "Back": c.Back},
			Tags:      c.Tags,
			Options: map[string]any{
				"allowDuplicate":        false,
				"duplicateScope":        "deck",
				"duplicateScopeOptions": map[string]any{"checkChildren": false},
			},
		})
	}

	if *dryRun {
		if *asJSON {
			payload, _ := json.MarshalIndent(ankiRequest{
				Action: "addNotes", Version: 6,
				Params: map[string]any{"notes": notes},
			}, "", "  ")
			fmt.Println(string(payload))
			return
		}
		for i, c := range cards {
			fmt.Printf("── Card %d ──────────────────────────────────\n", i+1)
			fmt.Printf("Deck : %s\n", c.Deck)
			fmt.Printf("Tags : %s\n", strings.Join(c.Tags, ", "))
			fmt.Printf("Front: %s\n", c.Front)
			fmt.Printf("Back :\n%s\n\n", c.Back)
		}
		return
	}

	// Make sure every deck exists.
	for _, d := range dedupe(deckNames(cards)) {
		if err := ankiCall(*url, "createDeck", map[string]any{"deck": d}, nil); err != nil {
			fatal(err)
		}
	}

	var ids []*int64
	if err := ankiCall(*url, "addNotes", map[string]any{"notes": notes}, &ids); err != nil {
		fatal(err)
	}

	added, dup := 0, 0
	for i, id := range ids {
		if id != nil {
			added++
			continue
		}
		dup++
		if *update {
			if err := updateExisting(*url, notes[i]); err != nil {
				fmt.Fprintf(os.Stderr, "⚠️  update failed for %q: %v\n", cards[i].Front, err)
			}
		}
	}

	fmt.Printf("✅ Added: %d\n", added)
	if dup > 0 {
		if *update {
			fmt.Printf("♻️  Updated (already existed): %d\n", dup)
		} else {
			fmt.Printf("⏭️  Skipped duplicates: %d  (use --update to refresh them)\n", dup)
		}
	}
	fmt.Println()
}

// updateExisting rewrites the fields of the note that shares this Front.
func updateExisting(url string, n ankiNote) error {
	query := fmt.Sprintf(`deck:"%s" "Front:%s"`, n.DeckName, escapeQuery(n.Fields["Front"]))
	var found []int64
	if err := ankiCall(url, "findNotes", map[string]any{"query": query}, &found); err != nil {
		return err
	}
	if len(found) == 0 {
		return fmt.Errorf("no existing note matched")
	}
	return ankiCall(url, "updateNoteFields", map[string]any{
		"note": map[string]any{"id": found[0], "fields": n.Fields},
	}, nil)
}

func escapeQuery(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, `*`, `\*`)
	s = strings.ReplaceAll(s, `_`, `\_`)
	return s
}

func deckNames(cards []Card) []string {
	var out []string
	for _, c := range cards {
		out = append(out, c.Deck)
	}
	return out
}

func collectFiles(args []string) ([]string, error) {
	var files []string
	for _, a := range args {
		info, err := os.Stat(a)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			files = append(files, a)
			continue
		}
		err = filepath.WalkDir(a, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() && (d.Name() == ".obsidian" || d.Name() == "tools" || d.Name() == ".git") {
				return fs.SkipDir
			}
			if !d.IsDir() && strings.HasSuffix(p, ".md") {
				files = append(files, p)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

// reorder puts flags first so the file/dir argument may appear anywhere on the
// command line (Go's flag package stops parsing at the first positional).
func reorder(args []string) []string {
	valueFlags := map[string]bool{"deck": true, "model": true, "tags": true, "url": true}
	var flags, positional []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if !strings.HasPrefix(a, "-") || a == "-" {
			positional = append(positional, a)
			continue
		}
		flags = append(flags, a)
		name := strings.TrimLeft(a, "-")
		if strings.Contains(name, "=") {
			continue
		}
		if valueFlags[name] && i+1 < len(args) {
			i++
			flags = append(flags, args[i])
		}
	}
	return append(flags, positional...)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "❌ %v\n", err)
	os.Exit(1)
}
