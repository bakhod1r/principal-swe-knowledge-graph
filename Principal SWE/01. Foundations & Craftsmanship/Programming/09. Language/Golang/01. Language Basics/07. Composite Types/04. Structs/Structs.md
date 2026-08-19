---
title: Structs
parent: "[[Composite Types]]"
---

- [[Empty Struct struct{} Memory Address (zerobase pointer)]] — Why all empty structs struct{} point to the same runtime.zerobase memory address.

- [[Struct Alignment, Offset Calculation & False Sharing]] — Field alignment rules, 64-byte cache lines, and eliminating multi-core false sharing.

---
title: Structs
tags:
  - golang
  - structs
  - principal-swe
parent: "`Composite Types`"
---

# Structs

Struct field alignment, memory padding optimization, tags reflection, composition embedding, and empty structs.

```text
Structs
│
├── [[Struct Memory Layout & Padding]]
├── [[Struct Tags & Serialization]]
├── [[Embedding Structs & Promotion]]
├── [[Empty Struct (struct{}) Idioms]]
├── [[Anonymous Structs & Inline Definitions]]
├── [[Struct Comparison & Equality Rules]]
└── [[Struct Pass-by-Value vs Pointer]]
```

---

## 🗂️ Topics

- [[Struct Memory Layout & Padding]] — CPU word alignment (8-byte on 64-bit), padding bytes, and field ordering optimization from largest to smallest.
- [[Struct Tags & Serialization]] — Field reflection metadata string literals (`json:"name,omitempty"`), encoding/json, and validator libraries.
- [[Embedding Structs & Promotion]] — Composition over inheritance: embedded inner struct fields and methods promoted to outer struct.
- [[Empty Struct (struct{}) Idioms]] — Zero-byte memory allocation type used for set membership (map[T]struct{}) and signal channels (chan struct{}).
- [[Anonymous Structs & Inline Definitions]] — Ad-hoc unexported structs for one-off table-driven test cases, local grouping, or json parsing.
- [[Struct Comparison & Equality Rules]] — Compile-time struct comparability rules (all fields must be comparable), comparing with ==.
- [[Struct Pass-by-Value vs Pointer]] — Memory copy overhead of large structs vs escape to heap and GC pressure with pointers.
- [[Cache Line Padding (cpu.CacheLineSize)]]
- [[Empty Struct struct{} Memory Address (zerobase pointer)]]
- [[False Sharing & CPU Cache Lines (64-byte)]]
- [[Struct Alignment, Offset Calculation & False Sharing]]

---

## 🔗 References
- ⬆️ Parent: `Composite Types`

