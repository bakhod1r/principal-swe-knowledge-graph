---
title: Specialized System Arrays
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Specialized System Arrays

OS virtual memory mapped arrays, text editor gap buffers, hashed array trees, suffix arrays, and non-temporal streaming store arrays.

```text
Specialized System Arrays
│
├── [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]]
├── [[Gap Buffer (Text Editor Cursor Buffer)]]
├── [[Hashed Array Tree (HAT) Cache-Friendly Flat Allocation]]
├── [[Suffix Array and Longest Common Prefix (LCP) Array]]
└── [[Non-Temporal Store Arrays (Direct-to-DRAM Streaming Stores)]]
```

---

## 🗂️ Topics & Implementations

- [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]] — Virtual memory page mapping for multi-gigabyte files bypassing user-space buffer copies.
- [[Gap Buffer (Text Editor Cursor Buffer)]] — Contiguous buffer with movable cursor gap enabling O(1) text insertions and deletions.
- [[Hashed Array Tree (HAT) Cache-Friendly Flat Allocation]] — Array of leaf chunks maintaining O(1) random access with zero reallocations.
- [[Suffix Array and Longest Common Prefix (LCP) Array]] — Compact sorted suffix index for full-text pattern matching and substring queries.
- [[Non-Temporal Store Arrays (Direct-to-DRAM Streaming Stores)]] — Bypassing CPU cache pollution for massive write-only bulk array initializations.

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
