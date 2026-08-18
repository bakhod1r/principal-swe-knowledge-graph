#!/usr/bin/env bash
#
# find_empty_note.sh — Find empty .md notes in an Obsidian vault
#
# Detection: Based on LINE COUNT
#   ≤ 30 lines = empty (template only: frontmatter + headings)
#   > 30 lines = has content
#
# Usage:
#   ./find_empty_note.sh <path>              # random empty note
#   ./find_empty_note.sh <path> --all        # show all empty notes
#   ./find_empty_note.sh <path> --topic      # pick a random topic, show its empty notes
#   ./find_empty_note.sh <path> --stats      # per-topic statistics

set -euo pipefail

# ─── Colors ───
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
BOLD='\033[1m'
DIM='\033[2m'
RESET='\033[0m'

# ─── Settings ───
MAX_LINES=30  # Lines at or below this threshold = empty note

# ─── Validate arguments ───
if [[ $# -lt 1 ]]; then
    echo -e "${RED}❌ Error: Path required!${RESET}"
    echo ""
    echo -e "${DIM}Usage: $0 <path> [--all|--topic|--stats]${RESET}"
    echo ""
    echo -e "  ${CYAN}<path>${RESET}              Search directory"
    echo -e "  ${CYAN}--all${RESET}               All empty notes"
    echo -e "  ${CYAN}--topic${RESET}             Random topic with its empty notes"
    echo -e "  ${CYAN}--stats${RESET}             Per-topic statistics"
    echo -e "  ${DIM}(none)${RESET}              Random single empty note"
    echo ""
    echo -e "${DIM}Empty = ${MAX_LINES} lines or fewer (template only)${RESET}"
    exit 1
fi

SEARCH_PATH="$1"
MODE="${2:---random}"

if [[ ! -d "$SEARCH_PATH" ]]; then
    echo -e "${RED}❌ Directory not found: ${SEARCH_PATH}${RESET}"
    exit 1
fi

# ─── Check if a file is empty based on LINE COUNT ───
is_empty_note() {
    local file="$1"
    local line_count
    line_count=$(wc -l < "$file" | tr -d ' ')

    if [[ "$line_count" -le "$MAX_LINES" ]]; then
        return 0  # empty
    else
        return 1  # has content
    fi
}

# ─── Collect all empty notes ───
collect_empty_notes() {
    local empty_notes=()

    while IFS= read -r -d '' file; do
        if is_empty_note "$file"; then
            empty_notes+=("$file")
        fi
    done < <(find "$SEARCH_PATH" -name '*.md' -type f \
        -not -path '*/.obsidian/*' \
        -not -path '*/tools/*' \
        -print0)

    printf '%s\n' "${empty_notes[@]}"
}

# ─── Get topic name (parent directory) ───
get_topic() {
    local file="$1"
    local dir
    dir=$(dirname "$file")
    basename "$dir"
}

# ─── MAIN ───
echo ""
echo -e "${BOLD}${MAGENTA}📝 Empty Note Finder${RESET}"
echo -e "${DIM}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${RESET}"
echo -e "${DIM}📂 Search path: ${SEARCH_PATH}${RESET}"
echo -e "${DIM}📏 Empty threshold: ≤ ${MAX_LINES} lines${RESET}"
echo ""

# Collect empty notes
mapfile -t EMPTY_NOTES < <(collect_empty_notes)

if [[ ${#EMPTY_NOTES[@]} -eq 0 ]]; then
    echo -e "${GREEN}✅ Great! No empty notes found — everything is filled in!${RESET}"
    exit 0
fi

echo -e "${YELLOW}📊 Total empty notes: ${#EMPTY_NOTES[@]}${RESET}"
echo ""

case "$MODE" in
    --all)
        echo -e "${BOLD}${CYAN}📋 All empty notes:${RESET}"
        echo -e "${DIM}─────────────────────────────────────────────────${RESET}"
        for note in "${EMPTY_NOTES[@]}"; do
            local_topic=$(get_topic "$note")
            local_name=$(basename "$note" .md)
            local_lines=$(wc -l < "$note" | tr -d ' ')
            echo -e "  ${BLUE}[${local_topic}]${RESET} ${local_name} ${DIM}(${local_lines} lines)${RESET}"
        done
        ;;

    --topic)
        # Collect all topics
        declare -A TOPICS
        for note in "${EMPTY_NOTES[@]}"; do
            topic=$(get_topic "$note")
            TOPICS["$topic"]+="$note"$'\n'
        done

        # Pick a random topic
        topic_keys=("${!TOPICS[@]}")
        random_index=$((RANDOM % ${#topic_keys[@]}))
        selected_topic="${topic_keys[$random_index]}"

        echo -e "${BOLD}${CYAN}🎯 Selected topic: ${YELLOW}${selected_topic}${RESET}"
        echo -e "${DIM}─────────────────────────────────────────────────${RESET}"

        while IFS= read -r note; do
            [[ -z "$note" ]] && continue
            local_name=$(basename "$note" .md)
            local_lines=$(wc -l < "$note" | tr -d ' ')
            echo -e "  ${GREEN}📄${RESET} ${local_name} ${DIM}(${local_lines} lines)${RESET}"
        done <<< "${TOPICS[$selected_topic]}"
        ;;

    --stats)
        echo -e "${BOLD}${CYAN}📊 Per-topic statistics:${RESET}"
        echo -e "${DIM}─────────────────────────────────────────────────${RESET}"

        declare -A TOPIC_COUNTS
        for note in "${EMPTY_NOTES[@]}"; do
            topic=$(get_topic "$note")
            TOPIC_COUNTS["$topic"]=$(( ${TOPIC_COUNTS["$topic"]:-0} + 1 ))
        done

        # Sort and display
        for topic in $(echo "${!TOPIC_COUNTS[@]}" | tr ' ' '\n' | sort); do
            count=${TOPIC_COUNTS[$topic]}
            bar=""
            for ((i=0; i<count; i++)); do bar+="█"; done
            echo -e "  ${BLUE}${topic}${RESET} ${DIM}(${count})${RESET} ${YELLOW}${bar}${RESET}"
        done
        ;;

    --random|*)
        # Pick a random note
        random_index=$((RANDOM % ${#EMPTY_NOTES[@]}))
        selected="${EMPTY_NOTES[$random_index]}"
        selected_topic=$(get_topic "$selected")
        selected_name=$(basename "$selected" .md)
        selected_lines=$(wc -l < "$selected" | tr -d ' ')

        echo -e "${BOLD}${GREEN}🎲 Random empty note:${RESET}"
        echo -e "${DIM}─────────────────────────────────────────────────${RESET}"
        echo -e "  ${BLUE}Topic:${RESET}  ${selected_topic}"
        echo -e "  ${BLUE}Note:${RESET}   ${BOLD}${selected_name}${RESET}"
        echo -e "  ${BLUE}Lines:${RESET}  ${selected_lines} lines"
        echo -e "  ${BLUE}Path:${RESET}   ${DIM}${selected}${RESET}"
        echo ""
        echo -e "${DIM}  💡 Fill in this note and strengthen your knowledge!${RESET}"
        ;;
esac

echo ""
