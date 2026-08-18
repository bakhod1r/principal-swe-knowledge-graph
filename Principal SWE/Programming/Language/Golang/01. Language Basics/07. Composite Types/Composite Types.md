---
title: Composite Types
tags:
  - golang
  - basics
parent: "[[Language Basics]]"
---

# Composite Types

Arrays, slices, maps, structs, memory layouts, padding, and zero-allocation techniques.

```text
Composite Types
│
├── [[Arrays]]
├── [[Slices]]
├── [[Slice Header Internals]]
├── [[Slice Capacity and Growth]]
├── [[make() for Slices and Maps]]
├── [[Slice to Array Conversion]]
├── [[Slice Tricks]]
├── [[Maps]]
├── [[Map Internals (hmap and bmap)]]
├── [[Comma-Ok Idiom for Maps]]
├── [[Structs]]
├── [[Struct Memory Layout & Padding]]
├── [[Struct Tags & JSON]]
├── [[Embedding Structs]]
├── [[Empty Struct (struct{})]]
└── [[Anonymous Structs]]
```

---

## 🗂️ Topics

- [[Arrays]] — Fixed-length contiguous memory sequences with value semantics.
- [[Slices]] — Dynamic views over arrays with length, capacity, and pointer.
- [[Slice Header Internals]] — sliceHeader struct: unsafe.Pointer, int len, int cap.
- [[Slice Capacity and Growth]] — Slice append allocation formula and doubling thresholds.
- [[make() for Slices and Maps]] — Allocating and initializing slices, maps, and channels.
- [[Slice to Array Conversion]] — Pointer-to-array conversions and Go 1.20 slice-to-array casting.
- [[Slice Tricks]] — Idiomatic cut, delete, insert, push, pop, and copy patterns.
- [[Maps]] — Hash table implementation with O(1) average lookup, insert, and delete.
- [[Map Internals (hmap and bmap)]] — Buckets, overflow buckets, hash seeds, and incremental evacuation.
- [[Comma-Ok Idiom for Maps]] — Distinguishing between zero values and missing keys (v, ok := m[k]).
- [[Structs]] — User-defined aggregate types grouping named fields.
- [[Struct Memory Layout & Padding]] — CPU memory word alignment, padding bytes, and field ordering.
- [[Struct Tags & JSON]] — Field reflection metadata for serialization (json, xml, db).
- [[Embedding Structs]] — Composition and field/method promotion without inheritance.
- [[Empty Struct (struct{})]] — Zero-byte type for signaling channels and set data structures.
- [[Anonymous Structs]] — Ad-hoc inline struct definitions for local grouping or test cases.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]
- 🎓 Root: [[Principal SWE]]
