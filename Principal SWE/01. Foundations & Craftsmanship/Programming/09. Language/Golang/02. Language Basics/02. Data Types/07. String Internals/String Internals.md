---
title: String Internals
tags:
  - golang
  - types
  - string-internals
  - low-level
  - principal-swe
parent: "[[Data Types]]"
---

# String Internals

`stringStruct` runtime layout, pointer and length fields, zero-copy conversions, `.rodata` deduplication, escape analysis, and memory retention.

```text
String Internals
│
├── [[stringStruct Runtime Layout]]
├── [[String Header vs Slice Header]]
├── [[Sub-String Memory Retention]]
├── [[Zero-Copy String to Byte Slice (unsafe)]]
├── [[RODATA Segment and Compiler String Deduplication]]
├── [[Compiler Zero-Allocation String to Slice Optimizations]]
└── [[Escape Analysis and Heap Allocation of Strings]]
```

---

## 🗂️ Topics

- [[stringStruct Runtime Layout]] — Two-word runtime struct: `unsafe.Pointer Data` and `int Len` (16 bytes on 64-bit).
- [[String Header vs Slice Header]] — Comparing 2-word `stringStruct` (`ptr`, `len`) with 3-word slice (`ptr`, `len`, `cap`).
- [[Sub-String Memory Retention]] — How slicing a tiny piece of a huge string retains the entire backing byte array from GC reclamation.
- [[Zero-Copy String to Byte Slice (unsafe)]] — Modern `unsafe.String` / `unsafe.StringData` idioms and legacy `reflect.StringHeader` pitfalls.
- [[RODATA Segment and Compiler String Deduplication]] — How string constants reside in read-only binary text segments and share backing pointers.
- [[Compiler Zero-Allocation String to Slice Optimizations]] — AST/SSA compiler passes that eliminate heap allocations in string ranges and map lookups.
- [[Escape Analysis and Heap Allocation of Strings]] — Determining when strings escape to the heap vs stack allocation via `go build -gcflags="-m"`.

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 📚 Module: `Language Basics`
