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
NOTE ?=
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
	@echo "  make anki-preview NOTE=<md>  Preview cards (one per numbered heading)"
	@echo "  make anki         NOTE=<md>  Send cards to Anki"
	@echo "  make anki-update  NOTE=<md>  Send + refresh existing cards"
	@echo "  make anki-json    NOTE=<md>  Dump the AnkiConnect payload"
	@echo "  (optional: DECK=\"Name\" TAGS=\"a,b\")"
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
# md-to-anki — one Anki note per numbered heading
#
#   make anki-preview NOTE="../path/Note.md"
#   make anki         NOTE="../path/Note.md"
#   make anki-update  NOTE="../path/Note.md"
#   make anki         NOTE="../01. Foundations & Craftsmanship" DECK="Go" TAGS="drill"
# ─────────────────────────────────────────────────

.PHONY: anki-build anki-preview anki anki-update anki-json

anki-build: ## Build the md-to-anki binary
	@cd $(ANKI_SRC_DIR) && go build -buildvcs=false -o ../$(ANKI_BINARY) .
	@echo "✅ Built ./$(ANKI_BINARY)"

ANKI_FLAGS = $(if $(DECK),--deck "$(DECK)",) $(if $(TAGS),--tags "$(TAGS)",)

anki-preview: anki-build ## Show the cards without touching Anki
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS) --dry-run

anki-json: anki-build ## Dump the AnkiConnect request payload
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS) --dry-run --json

anki: anki-build ## Send the cards to Anki via AnkiConnect
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS)

anki-update: anki-build ## Send and refresh cards that already exist
	@./$(ANKI_BINARY) "$(NOTE)" $(ANKI_FLAGS) --update
