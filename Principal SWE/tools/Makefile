# ─────────────────────────────────────────────────
# Empty Note Finder — Makefile
# ─────────────────────────────────────────────────
#
# Usage:
#   make random PATH_ARG=<dir>     — pick a random empty note (random topic → random note)
#   make all    PATH_ARG=<dir>     — list all empty notes
#   make topic  PATH_ARG=<dir>     — random topic → all its empty notes
#   make stats  PATH_ARG=<dir>     — per-topic statistics
#
# Examples:
#   make random PATH_ARG="../Programming/Language/Golang"
#   make stats  PATH_ARG=".."
# ─────────────────────────────────────────────────

SHELL := /bin/bash
BINARY := enf
SRC_DIR := ./empty-note-finder

ANKI_BINARY := md2anki
ANKI_SRC_DIR := ./md-to-anki
DECK ?=
TAGS ?=

# Default path — parent directory (vault root)
PATH_ARG ?= ..

.PHONY: help build random all topic stats clean

help: ## Show help
	@echo ""
	@echo "📝 Empty Note Finder"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo ""
	@echo "Commands:"
	@echo "  make build                   Build the binary"
	@echo "  make random PATH_ARG=<dir>   Random topic → random empty note"
	@echo "  make all    PATH_ARG=<dir>   List all empty notes"
	@echo "  make topic  PATH_ARG=<dir>   Random topic → all its empty notes"
	@echo "  make stats  PATH_ARG=<dir>   Per-topic statistics"
	@echo "  make clean                   Remove the binaries"
	@echo ""
	@echo "📚 md-to-anki (AnkiConnect)"
	@echo "  make anki-sync               Import every note not in Anki yet"
	@echo "  make anki         NOTE=<md>  Import one note or folder"
	@echo "  make anki-update  NOTE=<md>  Import + refresh existing cards"
	@echo "  make anki-preview NOTE=<md>  Print cards, touch nothing"
	@echo "  make anki-html    NOTE=<md>  Browser preview"
	@echo "  make anki-json    NOTE=<md>  Dump the AnkiConnect payload"
	@echo "  make anki-test               Run the unit tests"
	@echo "  (NOTE defaults to the whole vault; optional: DECK=\"Name\" TAGS=\"a,b\")"
	@echo ""
	@echo "Examples:"
	@echo "  make random PATH_ARG=\"../Programming/Language/Golang\""
	@echo "  make stats  PATH_ARG=\"..\""
	@echo ""

build: ## Build the Go binary
	@cd $(SRC_DIR) && go build -buildvcs=false -o ../$(BINARY) .
	@echo "✅ Built ./$(BINARY)"

random: build ## Random topic → random empty note
	@./$(BINARY) "$(PATH_ARG)"

all: build ## List all empty notes
	@./$(BINARY) "$(PATH_ARG)" --all

topic: build ## Random topic → all its empty notes
	@./$(BINARY) "$(PATH_ARG)" --topic

stats: build ## Per-topic statistics
	@./$(BINARY) "$(PATH_ARG)" --stats

clean: ## Remove the binaries
	@rm -f $(BINARY) $(ANKI_BINARY)
	@echo "🗑️  Cleaned"

# ─────────────────────────────────────────────────
# md-to-anki — one Anki card per numbered heading
#
#   make anki-sync                             — import every note not in Anki yet
#   make anki         NOTE="../path/Note.md"   — import one note (or folder)
#   make anki-preview NOTE="../path/Note.md"   — show the cards, touch nothing
#   make anki-html    NOTE="../path/Note.md"   — open a browser preview
#   make anki-update  NOTE="../path/Note.md"   — also refresh existing cards
#   make anki NOTE="../01. Foundations & Craftsmanship" DECK="Go" TAGS="drill"
#
# Cards already in Anki are never added twice: duplicates are detected across
# the whole collection, so a re-run only picks up what is new.
# ─────────────────────────────────────────────────

.PHONY: anki-build anki-test anki-sync anki anki-preview anki-html anki-json anki-update anki-check

anki-build: ## Build the md-to-anki binary
	@cd $(ANKI_SRC_DIR) && go build -buildvcs=false -o ../$(ANKI_BINARY) .
	@echo "✅ Built ./$(ANKI_BINARY)"

ANKI_FLAGS = $(if $(DECK),--deck "$(DECK)",) $(if $(TAGS),--tags "$(TAGS)",)
ANKI_HTML ?= /tmp/md-to-anki-preview.html

# Whole vault by default; NOTE=<path> narrows it to one note or folder.
NOTE ?= ..

anki-test: ## Run the md-to-anki unit tests
	@cd $(ANKI_SRC_DIR) && go test ./...

anki-sync: anki-build anki-check ## Import every note that is not in Anki yet
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS)

anki: anki-sync ## Alias for anki-sync

anki-update: anki-build anki-check ## Import new cards and refresh existing ones
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS) --update

anki-preview: anki-build ## Print the cards without touching Anki
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS) --dry-run

anki-html: anki-build ## Render the cards to a browser preview and open it
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS) --html "$(ANKI_HTML)"
	@open "$(ANKI_HTML)"

anki-json: anki-build ## Dump the AnkiConnect request payload
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS) --dry-run --json

anki-check: ## Verify Anki is running with the AnkiConnect add-on
	@curl -s -m 3 http://127.0.0.1:8765 -X POST \
		-d '{"action":"version","version":6}' > /dev/null || { \
		echo "❌ AnkiConnect is not answering on http://127.0.0.1:8765"; \
		echo "   Open Anki, then: Tools > Add-ons > Get Add-ons… > 2055492159"; \
		exit 1; }
