package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// vaultRoot is the vault folder the deck tree is mirrored from.
const vaultRoot = "Principal SWE"

// skipDirs are never walked when a directory is given on the command line.
var skipDirs = map[string]bool{".obsidian": true, ".git": true, "tools": true}

var numberPrefixRe = regexp.MustCompile(`^\d+\.\s*`)

// DeckFromPath mirrors the note's folder path into an Anki deck name:
//
//	.../Principal SWE/01. Foundations/Programming/Golang/Note.md
//	→ Principal SWE::Foundations::Programming::Golang
//
// The "01. " ordering prefixes Obsidian uses are dropped.
func DeckFromPath(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		abs = path
	}

	parts := strings.Split(filepath.ToSlash(filepath.Dir(abs)), "/")
	root := -1
	for i, p := range parts {
		if p == vaultRoot {
			root = i
		}
	}
	if root == -1 {
		return vaultRoot
	}

	deck := []string{vaultRoot}
	for _, p := range parts[root+1:] {
		p = numberPrefixRe.ReplaceAllString(p, "")
		p = strings.ReplaceAll(p, "::", "-") // "::" separates decks in Anki
		if p != "" {
			deck = append(deck, p)
		}
	}
	return strings.Join(deck, "::")
}

// IsIndexNote reports whether a note is the index (map-of-content) note of its
// folder — the note named after the folder it sits in, ordering prefix ignored:
//
//	.../09. Language/Golang/Golang.md
//
// Those notes only link out to the real material, so they make no cards.
func IsIndexNote(path string) bool {
	if !strings.HasSuffix(path, ".md") {
		return false
	}
	name := numberPrefixRe.ReplaceAllString(strings.TrimSuffix(filepath.Base(path), ".md"), "")
	dir := numberPrefixRe.ReplaceAllString(filepath.Base(filepath.Dir(path)), "")
	return name == dir
}

// CollectFiles expands the command-line arguments into a list of .md files.
func CollectFiles(args []string) ([]string, error) {
	var files []string
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			files = append(files, arg)
			continue
		}

		err = filepath.WalkDir(arg, func(path string, d fs.DirEntry, err error) error {
			switch {
			case err != nil:
				return err
			case d.IsDir() && skipDirs[d.Name()]:
				return fs.SkipDir
			case !d.IsDir() && IsIndexNote(path):
				return nil
			case !d.IsDir() && strings.HasSuffix(path, ".md"):
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}
