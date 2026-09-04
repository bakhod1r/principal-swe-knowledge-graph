// md-to-anki turns Obsidian notes into Anki cards — one card per numbered
// heading — and pushes them into Anki through the AnkiConnect add-on.
//
// Usage:
//
//	md-to-anki <file.md|dir> [flags]
//
// Flags:
//
//	--deck    target deck (default: mirrored from the note's folder path)
//	--model   Anki note type, created when missing (default "Principal SWE")
//	--tags    extra comma-separated tags
//	--url     AnkiConnect endpoint (default http://127.0.0.1:8765)
//	--dry-run print the cards instead of sending them
//	--json    with --dry-run, dump the AnkiConnect request payload
//	--update  refresh notes that already exist instead of skipping them
//	--flat    one deck per folder (default: one subdeck per note)
//	--html F  write a browser preview of the cards to F instead of importing
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type options struct {
	deck    string
	model   string
	tags    []string
	url     string
	dryRun  bool
	asJSON  bool
	update  bool
	flat    bool
	htmlOut string
	targets []string
}

// out is where the run report is written; tests redirect it.
var out io.Writer = os.Stdout

const usage = "usage: md-to-anki <file.md|dir> [--deck D] [--model M] [--tags a,b] [--dry-run] [--json] [--html F] [--update] [--flat]"

func main() {
	opts, err := parseFlags(os.Args[1:])
	if err != nil {
		fatal(err)
	}
	if err := run(opts); err != nil {
		fatal(err)
	}
}

// run performs the whole job: parse the notes, then preview or import them.
func run(opts options) error {
	cards, files, err := buildCards(opts)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "\n📚 md-to-anki\n")
	fmt.Fprintf(out, "───────────────────────────────────────────────\n")
	fmt.Fprintf(out, "Files: %d   Cards: %d   Model: %s\n\n", len(files), len(cards), opts.model)

	switch {
	case opts.htmlOut != "":
		if err := WritePreviewHTML(opts.htmlOut, cards, opts.model); err != nil {
			return err
		}
		fmt.Fprintf(out, "🖼️  Preview written to %s\n\n", opts.htmlOut)
		return nil

	case opts.dryRun:
		preview(cards, opts)
		return nil

	default:
		return importCards(cards, opts)
	}
}

// ─── command line ───

func parseFlags(argv []string) (options, error) {
	fs := flag.NewFlagSet("md-to-anki", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	var (
		deck   = fs.String("deck", "", "target deck (default: mirrored from the note's folder path)")
		model  = fs.String("model", "Principal SWE", "Anki note type (created automatically when missing)")
		tagCSV = fs.String("tags", "", "extra comma-separated tags")
		url    = fs.String("url", "http://127.0.0.1:8765", "AnkiConnect endpoint")
		dryRun = fs.Bool("dry-run", false, "print the cards instead of sending them")
		asJSON = fs.Bool("json", false, "with --dry-run, dump the AnkiConnect request payload")
		update = fs.Bool("update", false, "refresh notes that already exist instead of skipping them")
		flat   = fs.Bool("flat", false, "put a folder's cards in one deck instead of a subdeck per note")
		htmlP  = fs.String("html", "", "write a browser preview of the cards to this file instead of importing")
	)
	if err := fs.Parse(flagsFirst(argv)); err != nil {
		return options{}, err
	}
	if fs.NArg() < 1 {
		return options{}, fmt.Errorf(usage)
	}

	var tags []string
	for _, t := range strings.Split(*tagCSV, ",") {
		if t = strings.TrimSpace(t); t != "" {
			tags = append(tags, t)
		}
	}

	return options{
		deck:    *deck,
		model:   *model,
		tags:    tags,
		url:     *url,
		dryRun:  *dryRun,
		asJSON:  *asJSON,
		update:  *update,
		flat:    *flat,
		htmlOut: *htmlP,
		targets: fs.Args(),
	}, nil
}

// flagsFirst moves flags ahead of positional arguments so the file or directory
// may appear anywhere — Go's flag package stops parsing at the first positional.
func flagsFirst(args []string) []string {
	takesValue := map[string]bool{"deck": true, "model": true, "tags": true, "url": true, "html": true}

	var flags, positional []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") || arg == "-" {
			positional = append(positional, arg)
			continue
		}

		flags = append(flags, arg)
		name := strings.TrimLeft(arg, "-")
		if !strings.Contains(name, "=") && takesValue[name] && i+1 < len(args) {
			i++
			flags = append(flags, args[i])
		}
	}
	return append(flags, positional...)
}

// ─── card building ───

func buildCards(opts options) (cards []Card, files []string, err error) {
	files, err = CollectFiles(opts.targets)
	if err != nil {
		return nil, nil, err
	}
	if len(files) == 0 {
		return nil, nil, fmt.Errorf("no .md files found")
	}

	seen := map[string]string{} // front → note it first came from

	for _, file := range files {
		deck := opts.deck
		if deck == "" {
			deck = DeckFromPath(file)
		}
		parsed, err := ParseNote(file, deck, opts.tags)
		if err != nil {
			return nil, nil, err
		}

		// Two notes can carry the same title and section heading. AnkiConnect
		// only checks a new card against the collection, not against the rest
		// of the batch, so identical fronts are dropped here.
		for _, c := range parsed {
			if first, dup := seen[c.Front]; dup {
				fmt.Fprintf(os.Stderr, "⚠️  duplicate card %q in %s (already in %s) — skipped\n",
					c.Front, filepath.Base(file), filepath.Base(first))
				continue
			}
			seen[c.Front] = file
			// One note per deck keeps a folder of a dozen notes from piling
			// its cards into a single 150-card deck.
			if !opts.flat && opts.deck == "" {
				c.Deck += "::" + strings.ReplaceAll(c.Title, "::", "-")
			}
			cards = append(cards, c)
		}
	}
	return cards, files, nil
}

// notesFor converts cards into AnkiConnect notes written into the given fields.
func notesFor(cards []Card, model string, fields Fields) []Note {
	notes := make([]Note, 0, len(cards))
	for _, c := range cards {
		notes = append(notes, Note{
			DeckName:  c.Deck,
			ModelName: model,
			Fields:    map[string]string{fields.Front: c.Front, fields.Back: c.Back},
			Tags:      c.Tags,
			Options: map[string]any{
				// A card is a duplicate anywhere in the collection, not just
				// inside its own deck — moving or renaming a deck must never
				// let the same question come back twice.
				"allowDuplicate": false,
				"duplicateScope": "collection",
			},
		})
	}
	return notes
}

// ─── import ───

func importCards(cards []Card, opts options) error {
	client := NewClient(opts.url)

	fields, err := client.EnsureModel(opts.model)
	if err != nil {
		return err
	}
	notes := notesFor(cards, opts.model, fields)

	for _, deck := range dedupe(deckNames(cards)) {
		if err := client.CreateDeck(deck); err != nil {
			return err
		}
	}

	// AnkiConnect rejects a whole addNotes batch that contains a duplicate,
	// so new notes and existing ones are handled separately.
	addable, err := client.CanAdd(notes)
	if err != nil {
		return err
	}

	var fresh []Note
	var newCards, existing []int
	for i, ok := range addable {
		if ok {
			fresh = append(fresh, notes[i])
			newCards = append(newCards, i)
		} else {
			existing = append(existing, i)
		}
	}

	added, err := client.AddNotes(fresh)
	if err != nil {
		return err
	}

	if opts.update {
		for _, i := range existing {
			if err := client.UpdateNote(notes[i], fields.Front); err != nil {
				fmt.Fprintf(os.Stderr, "⚠️  update failed for %q: %v\n", cards[i].Front, err)
			}
		}
	}

	listNewNotes(cards, newCards)
	report(added, len(existing), opts.update)
	return nil
}

// listNewNotes shows which notes contributed cards that were not in Anki yet.
func listNewNotes(cards []Card, newCards []int) {
	if len(newCards) == 0 {
		return
	}

	counts := map[string]int{}
	var order []string
	for _, i := range newCards {
		src := cards[i].Source
		if counts[src] == 0 {
			order = append(order, src)
		}
		counts[src]++
	}

	fmt.Fprintf(out, "🆕 New notes: %d\n", len(order))
	for _, src := range order {
		fmt.Fprintf(out, "   %-4d %s\n", counts[src], filepath.Base(src))
	}
	fmt.Fprintln(out)
}

func report(added, existing int, updated bool) {
	fmt.Fprintf(out, "✅ Added: %d\n", added)
	switch {
	case existing == 0:
	case updated:
		fmt.Fprintf(out, "♻️  Updated (already existed): %d\n", existing)
	default:
		fmt.Fprintf(out, "⏭️  Already in Anki, skipped: %d  (use --update to refresh them)\n", existing)
	}
	fmt.Fprintln(out)
}

func deckNames(cards []Card) []string {
	names := make([]string, 0, len(cards))
	for _, c := range cards {
		names = append(names, c.Deck)
	}
	return names
}

// ─── preview ───

func preview(cards []Card, opts options) {
	if opts.asJSON {
		notes := notesFor(cards, opts.model, Fields{Front: "Front", Back: "Back"})
		payload, _ := json.MarshalIndent(request{
			Action:  "addNotes",
			Version: 6,
			Params:  map[string]any{"notes": notes},
		}, "", "  ")
		fmt.Fprintln(out, string(payload))
		return
	}

	for i, c := range cards {
		fmt.Fprintf(out, "── Card %d ──────────────────────────────────\n", i+1)
		fmt.Fprintf(out, "Deck : %s\n", c.Deck)
		fmt.Fprintf(out, "Tags : %s\n", strings.Join(c.Tags, ", "))
		fmt.Fprintf(out, "Front: %s\n", c.Front)
		fmt.Fprintf(out, "Back :\n%s\n\n", c.Back)
	}
}

// exit is os.Exit in production; tests replace it.
var exit = os.Exit

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "❌ %v\n", err)
	exit(1)
}
