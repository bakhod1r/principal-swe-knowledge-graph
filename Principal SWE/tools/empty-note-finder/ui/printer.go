// Package ui handles all terminal output formatting with colors and structure.
package ui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mrb/empty-note-finder/vault"
)

// ANSI color codes
const (
	Red     = "\033[0;31m"
	Green   = "\033[0;32m"
	Yellow  = "\033[1;33m"
	Blue    = "\033[0;34m"
	Cyan    = "\033[0;36m"
	Magenta = "\033[0;35m"
	Bold    = "\033[1m"
	Dim     = "\033[2m"
	Reset   = "\033[0m"
)

// PrintHeader prints the tool banner with search path and threshold info.
func PrintHeader(searchPath string, maxLines int) {
	fmt.Println()
	fmt.Printf("%s%s📝 Empty Note Finder%s\n", Bold, Magenta, Reset)
	fmt.Printf("%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", Dim, Reset)
	fmt.Printf("%s📂 Search path: %s%s\n", Dim, searchPath, Reset)
	fmt.Printf("%s📏 Empty threshold: ≤ %d lines%s\n", Dim, maxLines, Reset)
	fmt.Println()
}

// PrintSummary prints the total count of empty notes.
func PrintSummary(result *vault.ScanResult) {
	if len(result.EmptyNotes) == 0 {
		fmt.Printf("%s✅ Great! No empty notes found — everything is filled in!%s\n\n", Green, Reset)
		return
	}
	fmt.Printf("%s📊 Total empty notes: %d / %d files%s\n\n", Yellow, len(result.EmptyNotes), result.TotalFiles, Reset)
}

// PrintRandom prints a single randomly selected empty note.
func PrintRandom(note vault.Note) {
	fmt.Printf("%s%s🎲 Random empty note:%s\n", Bold, Green, Reset)
	fmt.Printf("%s─────────────────────────────────────────────────%s\n", Dim, Reset)
	fmt.Printf("  %sTopic:%s  %s\n", Blue, Reset, note.Topic)
	fmt.Printf("  %sNote:%s   %s%s%s\n", Blue, Reset, Bold, note.Name, Reset)
	fmt.Printf("  %sLines:%s  %d lines\n", Blue, Reset, note.LineCount)
	fmt.Printf("  %sPath:%s   %s%s%s\n", Blue, Reset, Dim, note.Path, Reset)
	fmt.Println()
	fmt.Printf("%s  💡 Fill in this note and strengthen your knowledge!%s\n", Dim, Reset)
}

// PrintAll prints every empty note grouped by display.
func PrintAll(notes []vault.Note) {
	fmt.Printf("%s%s📋 All empty notes:%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s─────────────────────────────────────────────────%s\n", Dim, Reset)
	for _, note := range notes {
		fmt.Printf("  %s[%s]%s %s %s(%d lines)%s\n",
			Blue, note.Topic, Reset, note.Name, Dim, note.LineCount, Reset)
	}
}

// PrintTopic prints a randomly selected topic and its empty notes.
func PrintTopic(topic string, notes []vault.Note) {
	fmt.Printf("%s%s🎯 Selected topic: %s%s%s\n", Bold, Cyan, Yellow, topic, Reset)
	fmt.Printf("%s─────────────────────────────────────────────────%s\n", Dim, Reset)
	for _, note := range notes {
		fmt.Printf("  %s📄%s %s %s(%d lines)%s\n",
			Green, Reset, note.Name, Dim, note.LineCount, Reset)
	}
}

// PrintStats prints per-topic statistics with bar chart visualization.
func PrintStats(stats []vault.TopicStat) {
	fmt.Printf("%s%s📊 Per-topic statistics:%s\n", Bold, Cyan, Reset)
	fmt.Printf("%s─────────────────────────────────────────────────%s\n", Dim, Reset)

	// Sort by count descending
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})

	// Find max topic name length for alignment
	maxLen := 0
	for _, s := range stats {
		if len(s.Topic) > maxLen {
			maxLen = len(s.Topic)
		}
	}

	for _, s := range stats {
		padding := strings.Repeat(" ", maxLen-len(s.Topic))
		bar := strings.Repeat("█", s.Count)
		fmt.Printf("  %s%s%s%s %s(%2d)%s %s%s%s\n",
			Blue, s.Topic, padding, Reset, Dim, s.Count, Reset, Yellow, bar, Reset)
	}
}
