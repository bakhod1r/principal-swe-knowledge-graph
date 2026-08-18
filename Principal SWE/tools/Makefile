# ─────────────────────────────────────────────────
# Empty Note Finder — Makefile
# ─────────────────────────────────────────────────
#
# Usage:
#   make random PATH_ARG=<dir>     — pick a random empty note
#   make all    PATH_ARG=<dir>     — list all empty notes
#   make topic  PATH_ARG=<dir>     — random topic + its empty notes
#   make stats  PATH_ARG=<dir>     — per-topic statistics
#
# Examples:
#   make random PATH_ARG="./Programming/Language/Golang"
#   make stats  PATH_ARG="."
# ─────────────────────────────────────────────────

SHELL := /bin/bash
SCRIPT := ./find_empty_note.sh

# Default path — current directory
PATH_ARG ?= .

.PHONY: help random all topic stats setup

help: ## Show help
	@echo ""
	@echo "📝 Empty Note Finder"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo ""
	@echo "Commands:"
	@echo "  make random PATH_ARG=<dir>   Pick a random empty note"
	@echo "  make all    PATH_ARG=<dir>   List all empty notes"
	@echo "  make topic  PATH_ARG=<dir>   Random topic + its empty notes"
	@echo "  make stats  PATH_ARG=<dir>   Per-topic statistics"
	@echo "  make setup                   Make script executable"
	@echo ""
	@echo "Examples:"
	@echo "  make random PATH_ARG=\"./Programming/Language/Golang\""
	@echo "  make stats  PATH_ARG=\".\""
	@echo ""

setup: ## Make the script executable
	@chmod +x $(SCRIPT)
	@echo "✅ $(SCRIPT) — ready!"

random: setup ## Find a random empty note
	@$(SCRIPT) "$(PATH_ARG)"

all: setup ## List all empty notes
	@$(SCRIPT) "$(PATH_ARG)" --all

topic: setup ## Random topic + its empty notes
	@$(SCRIPT) "$(PATH_ARG)" --topic

stats: setup ## Per-topic statistics
	@$(SCRIPT) "$(PATH_ARG)" --stats
