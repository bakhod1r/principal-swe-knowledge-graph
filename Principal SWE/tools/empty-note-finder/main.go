// empty-note-finder is a CLI tool that scans an Obsidian vault
// and finds empty markdown notes (detected by line count).
//
// Default behavior: picks a random topic, then picks a random
// empty note from that topic.
//
// Usage:
//
//	empty-note-finder <path>            random topic → random note
//	empty-note-finder <path> --all      list all empty notes
//	empty-note-finder <path> --topic    random topic → all its empty notes
//	empty-note-finder <path> --stats    per-topic statistics
package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/mrb/empty-note-finder/ui"
	"github.com/mrb/empty-note-finder/vault"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	searchPath := os.Args[1]
	mode := "random"
	if len(os.Args) >= 3 {
		mode = os.Args[2]
	}

	// Validate path
	info, err := os.Stat(searchPath)
	if err != nil || !info.IsDir() {
		fmt.Fprintf(os.Stderr, "\033[0;31m❌ Directory not found: %s\033[0m\n", searchPath)
		os.Exit(1)
	}

	// Scan
	scanner := vault.NewScanner()
	result, err := scanner.Scan(searchPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[0;31m❌ Scan error: %v\033[0m\n", err)
		os.Exit(1)
	}

	// Print header and summary
	ui.PrintHeader(searchPath, scanner.MaxLines)
	ui.PrintSummary(result)

	if len(result.EmptyNotes) == 0 {
		return
	}

	// Dispatch by mode
	switch mode {
	case "--all":
		ui.PrintAll(result.EmptyNotes)

	case "--topic":
		topic, notes := result.RandomTopic()
		if topic != "" {
			ui.PrintTopic(topic, notes)
		}

	case "--stats":
		stats := result.TopicStats()
		ui.PrintStats(stats)

	default:
		// Default: pick a random topic, then a random note from it
		topic, notes := result.RandomTopic()
		if topic == "" {
			return
		}
		// Pick one random note from the selected topic
		note := notes[rand.IntN(len(notes))]
		ui.PrintRandom(note)
	}

	fmt.Println()
}

func printUsage() {
	fmt.Println()
	fmt.Println("\033[0;31m❌ Error: Path required!\033[0m")
	fmt.Println()
	fmt.Println("\033[2mUsage: empty-note-finder <path> [--all|--topic|--stats]\033[0m")
	fmt.Println()
	fmt.Println("  \033[0;36m<path>\033[0m              Search directory")
	fmt.Println("  \033[0;36m--all\033[0m               List all empty notes")
	fmt.Println("  \033[0;36m--topic\033[0m             Random topic → all its empty notes")
	fmt.Println("  \033[0;36m--stats\033[0m             Per-topic statistics")
	fmt.Println("  \033[2m(none)\033[0m              Random topic → random note from it")
	fmt.Println()
}
