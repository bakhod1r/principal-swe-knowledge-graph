// Package vault provides utilities for scanning and analyzing
// Obsidian vault markdown notes.
package vault

import (
	"bufio"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
)

// Note represents a single markdown note in the vault.
type Note struct {
	Path      string // absolute file path
	Name      string // filename without .md extension
	Topic     string // parent directory name (topic group)
	LineCount int    // total number of lines
}

// ScanResult holds the outcome of a vault scan.
type ScanResult struct {
	EmptyNotes []Note            // all notes classified as empty
	TopicMap   map[string][]Note // empty notes grouped by topic
	TotalFiles int               // total .md files scanned
}

// Scanner walks an Obsidian vault directory and identifies empty notes.
type Scanner struct {
	MaxLines     int      // line count threshold (≤ this = empty)
	ExcludeDirs  []string // directory names to skip
}

// NewScanner creates a Scanner with sensible defaults.
func NewScanner() *Scanner {
	return &Scanner{
		MaxLines:    30,
		ExcludeDirs: []string{".obsidian", "tools", ".trash", ".git"},
	}
}

// Scan walks the given root directory and returns a ScanResult.
func (s *Scanner) Scan(root string) (*ScanResult, error) {
	result := &ScanResult{
		TopicMap: make(map[string][]Note),
	}

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // skip inaccessible paths
		}

		// Skip excluded directories
		if d.IsDir() {
			for _, exc := range s.ExcludeDirs {
				if d.Name() == exc {
					return filepath.SkipDir
				}
			}
			return nil
		}

		// Only process .md files
		if !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}

		result.TotalFiles++

		lineCount, err := countLines(path)
		if err != nil {
			return nil // skip unreadable files
		}

		if lineCount <= s.MaxLines {
			note := Note{
				Path:      path,
				Name:      strings.TrimSuffix(d.Name(), ".md"),
				Topic:     filepath.Base(filepath.Dir(path)),
				LineCount: lineCount,
			}
			result.EmptyNotes = append(result.EmptyNotes, note)
			result.TopicMap[note.Topic] = append(result.TopicMap[note.Topic], note)
		}

		return nil
	})

	return result, err
}

// RandomNote picks a random empty note from the scan result.
// Returns the note and true if found, zero Note and false if none exist.
func (r *ScanResult) RandomNote() (Note, bool) {
	if len(r.EmptyNotes) == 0 {
		return Note{}, false
	}
	return r.EmptyNotes[rand.IntN(len(r.EmptyNotes))], true
}

// RandomTopic picks a random topic and returns its name and empty notes.
// Returns empty string and nil if no topics exist.
func (r *ScanResult) RandomTopic() (string, []Note) {
	if len(r.TopicMap) == 0 {
		return "", nil
	}

	topics := make([]string, 0, len(r.TopicMap))
	for t := range r.TopicMap {
		topics = append(topics, t)
	}

	selected := topics[rand.IntN(len(topics))]
	return selected, r.TopicMap[selected]
}

// TopicStats returns a sorted slice of topic names and their empty note counts.
func (r *ScanResult) TopicStats() []TopicStat {
	stats := make([]TopicStat, 0, len(r.TopicMap))
	for topic, notes := range r.TopicMap {
		stats = append(stats, TopicStat{
			Topic: topic,
			Count: len(notes),
		})
	}
	return stats
}

// TopicStat represents the count of empty notes in a single topic.
type TopicStat struct {
	Topic string
	Count int
}

// countLines counts the number of lines in a file efficiently.
func countLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
