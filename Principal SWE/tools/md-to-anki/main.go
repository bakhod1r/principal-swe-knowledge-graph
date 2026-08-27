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
//	--html F  write a browser preview of the cards to F instead of importing
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
	htmlOut string
	targets []string
}

func main() {
	opts := parseFlags()

	cards, files, err := buildCards(opts)
	if err != nil {
		fatal(err)
	}

	fmt.Printf("\n📚 md-to-anki\n")
	fmt.Printf("───────────────────────────────────────────────\n")
	fmt.Printf("Files: %d   Cards: %d   Model: %s\n\n", len(files), len(cards), opts.model)

	if opts.htmlOut != "" {
		if err := WritePreviewHTML(opts.htmlOut, cards, opts.model); err != nil {
			fatal(err)
		}
		fmt.Printf("🖼️  Preview written to %s\n\n", opts.htmlOut)
		return
	}
	if opts.dryRun {
		preview(cards, opts)
		return
	}
	if err := importCards(cards, opts); err != nil {
		fatal(err)
	}
}

// ─── command line ───

func parseFlags() options {
	var (
		deck   = flag.String("deck", "", "target deck (default: mirrored from the note's folder path)")
		model  = flag.String("model", "Principal SWE", "Anki note type (created automatically when missing)")
		tagCSV = flag.String("tags", "", "extra comma-separated tags")
		url    = flag.String("url", "http://127.0.0.1:8765", "AnkiConnect endpoint")
		dryRun = flag.Bool("dry-run", false, "print the cards instead of sending them")
		asJSON = flag.Bool("json", false, "with --dry-run, dump the AnkiConnect request payload")
		update = flag.Bool("update", false, "refresh notes that already exist instead of skipping them")
		htmlP  = flag.String("html", "", "write a browser preview of the cards to this file instead of importing")
	)
	flag.CommandLine.Parse(flagsFirst(os.Args[1:]))

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: md-to-anki <file.md|dir> [--deck D] [--model M] [--tags a,b] [--dry-run] [--json] [--update]")
		os.Exit(1)
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
		htmlOut: *htmlP,
		targets: flag.Args(),
	}
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

	for _, file := range files {
		deck := opts.deck
		if deck == "" {
			deck = DeckFromPath(file)
		}
		c, err := ParseNote(file, deck, opts.tags)
		if err != nil {
			return nil, nil, err
		}
		cards = append(cards, c...)
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
				"allowDuplicate":        false,
				"duplicateScope":        "deck",
				"duplicateScopeOptions": map[string]any{"checkChildren": false},
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
	var existing []int
	for i, ok := range addable {
		if ok {
			fresh = append(fresh, notes[i])
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

	report(added, len(existing), opts.update)
	return nil
}

func report(added, existing int, updated bool) {
	fmt.Printf("✅ Added: %d\n", added)
	switch {
	case existing == 0:
	case updated:
		fmt.Printf("♻️  Updated (already existed): %d\n", existing)
	default:
		fmt.Printf("⏭️  Skipped duplicates: %d  (use --update to refresh them)\n", existing)
	}
	fmt.Println()
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
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "❌ %v\n", err)
	os.Exit(1)
}
