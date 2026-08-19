---
title: Strings
tags:
  - golang
  - types
  - strings
  - principal-swe
parent: "[[Data Types]]"
---

# Strings

Immutable byte sequences, string literals, indexing, slicing, concatenation performance, zero-allocation operations, UTF-8 streaming, and string deduplication.

```text
Strings
│
├── [[Interpreted String Literals]]
├── [[Raw String Literals]]
├── [[String Indexing and Slicing]]
├── [[String Immutability & Read-Only Memory (RODATA Segment)]]
├── [[String Concatenation Performance]]
├── [[strings.Builder Deep Dive]]
├── [[Zero-Allocation String Operations (strings.Cut, strings.Clone)]]
├── [[String Interning and Deduplication (unique Package)]]
├── [[UTF-8 Multi-Byte Streaming & Validation (utf8.ValidString)]]
├── [[UTF-8 Streaming Decoders]]
├── [[Zero-Allocation Byte Scanners]]
├── [[strings.Reader vs bytes.Reader (Streaming Interfaces)]]
└── [[strings.Replacer and Multi-String Search (Aho-Corasick)]]
```

---

## 🗂️ Topics

- [[Interpreted String Literals]] — Double-quoted strings supporting escape sequences (`\n`, `\t`, `\x`, `\u`).
- [[Raw String Literals]] — Multi-line unescaped strings enclosed in backticks (`` ` ``) for regex and templates.
- [[String Indexing and Slicing]] — Accessing raw byte offsets (`s[i]`) vs sub-slice expressions (`s[i:j]`).
- [[String Immutability & Read-Only Memory (RODATA Segment)]] — String literals compiled into read-only `.rodata` segment and panic on mutation.
- [[String Concatenation Performance]] — Comparing `+`, `fmt.Sprintf`, `strings.Join`, `bytes.Buffer`, and `strings.Builder`.
- [[strings.Builder Deep Dive]] — Zero-allocation string builder pattern and `String()` memory copy elimination.
- [[Zero-Allocation String Operations (strings.Cut, strings.Clone)]] — Modern standard idioms: `strings.Cut` parsing and `strings.Clone` memory retention fix.
- [[String Interning and Deduplication (unique Package)]] — Go 1.23+ `unique.Make` for canonical deduplication and pointer-equality comparisons.
- [[UTF-8 Multi-Byte Streaming & Validation (utf8.ValidString)]] — Validating corrupted multi-byte UTF-8 streams and RuneError boundary handling.
- [[UTF-8 Streaming Decoders]] — Chunked streaming decoding of UTF-8 streams with `bufio.Scanner` and custom split functions.
- [[Zero-Allocation Byte Scanners]] — Parsing protocol streams and text tokens with zero intermediate string heap allocations.
- [[strings.Reader vs bytes.Reader (Streaming Interfaces)]] — In-memory implementations of `io.Reader`, `io.ReaderAt`, `io.Seeker`, and `io.WriterTo`.
- [[strings.Replacer and Multi-String Search (Aho-Corasick)]] — Fast multi-pattern replacement algorithm using trie-based DFA lookups.

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 📚 Module: `Language Basics`
