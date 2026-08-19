---
title: Runes
tags:
  - golang
  - types
  - runes
  - unicode
  - principal-swe
parent: "[[Data Types]]"
---

# Runes

Unicode code points, rune literals, UTF-8 encoding, unicode/utf8 package, normalization, and grapheme clusters.

```text
Runes
│
├── [[Rune Type (int32)]]
├── [[Rune Literals & Escapes]]
├── [[Unicode & UTF-8 Relationship]]
├── [[unicode-utf8 Standard Package]]
├── [[Rune Count vs Byte Length (utf8.RuneCountInString)]]
├── [[Unicode Normalization (NFC, NFD, NFKC, NFKD)]]
├── [[Grapheme Clusters and Multi-Rune Characters]]
└── [[unicode Package (Categories, Properties, and Scripts)]]
```

---

## 🗂️ Topics

- [[Rune Type (int32)]] — Rune alias for `int32` representing a single Unicode code point.
- [[Rune Literals & Escapes]] — Single-quoted Unicode characters, `\u` (16-bit), and `\U` (32-bit) escape sequences.
- [[Unicode & UTF-8 Relationship]] — Variable-width UTF-8 encoding (1 to 4 bytes per code point).
- [[unicode-utf8 Standard Package]] — `utf8.DecodeRune`, `utf8.RuneLen`, and fast validation functions.
- [[Rune Count vs Byte Length (utf8.RuneCountInString)]] — Understanding $O(N)$ rune counting vs $O(1)$ byte length indexing traps.
- [[Unicode Normalization (NFC, NFD, NFKC, NFKD)]] — Canonical and compatibility decomposition/composition via `golang.org/x/text/unicode/norm`.
- [[Grapheme Clusters and Multi-Rune Characters]] — Handling multi-rune visual characters (emojis, accent modifiers) with `golang.org/x/exp/unicode/ragel`.
- [[unicode Package (Categories, Properties, and Scripts)]] — Character classifications (`unicode.IsLetter`, `unicode.IsDigit`, `unicode.IsSpace`, and script tables).

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 📚 Module: `Language Basics`
