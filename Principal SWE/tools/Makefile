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
	@echo "  make clean                   Remove the binary"
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

clean: ## Remove the binary
	@rm -f $(BINARY)
	@echo "🗑️  Cleaned"
