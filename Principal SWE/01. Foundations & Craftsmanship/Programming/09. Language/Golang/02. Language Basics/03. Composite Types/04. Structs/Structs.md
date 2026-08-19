---
title: Structs
tags:
  - golang
  - structs
  - composite-types
  - principal-swe
parent: "[[Composite Types]]"
---

# Structs

Struct field alignment, memory padding optimization, tags reflection, composition embedding, empty structs, and encapsulation patterns.

```text
Structs
│
├── [[Struct Memory Layout & Padding]]
├── [[Struct Alignment, Offset Calculation & False Sharing]]
├── [[Cache Line Padding (cpu.CacheLineSize)]]
├── [[Struct Pass-by-Value vs Pointer]]
├── [[Struct Comparison & Equality Rules]]
├── [[Embedding Structs & Promotion]]
├── [[Struct Embedding vs Composition (Inheritance vs Delegation)]]
├── [[Anonymous Structs & Inline Definitions]]
├── [[Struct Tags & Serialization]]
├── [[reflect.StructTag Parsing & Custom Validation Engines]]
├── [[Empty Struct (struct{}) Idioms]]
├── [[Zero-Allocation Struct Packing with Bitfields]]
└── [[Unexported Fields & Opaque Type Encapsulation (Pimpl Idiom in Go)]]
```

---

## 🗂️ Topics

- [[Struct Memory Layout & Padding]] — CPU word alignment (8-byte on 64-bit), padding bytes, and field ordering optimization from largest to smallest.
- [[Struct Alignment, Offset Calculation & False Sharing]] — Field alignment rules, 64-byte cache lines, and calculating struct offsets.
- [[Cache Line Padding (cpu.CacheLineSize)]] — Eliminating cache line bouncing using hardware architecture padding.
- [[Struct Pass-by-Value vs Pointer]] — Memory copy overhead of large structs vs escape to heap and GC pressure with pointers.
- [[Struct Comparison & Equality Rules]] — Compile-time struct comparability rules (all fields must be comparable), comparing with `==`.
- [[Embedding Structs & Promotion]] — Composition over inheritance: embedded inner struct fields and methods promoted to outer struct.
- [[Struct Embedding vs Composition (Inheritance vs Delegation)]] — Method shadowing, dispatch resolution, and why embedding is not classical OOP inheritance.
- [[Anonymous Structs & Inline Definitions]] — Ad-hoc unexported structs for one-off table-driven test cases, local grouping, or json parsing.
- [[Struct Tags & Serialization]] — Field reflection metadata string literals (`json:"name,omitempty"`), `encoding/json`, and tag conventions.
- [[reflect.StructTag Parsing & Custom Validation Engines]] — Zero-allocation struct tag extraction, custom DSL tag parsing, and struct validator implementations.
- [[Empty Struct (struct{}) Idioms]] — Zero-byte memory allocation type used for set membership (`map[T]struct{}`) and signal channels (`chan struct{}`).
- [[Zero-Allocation Struct Packing with Bitfields]] — Packing boolean flags and sub-word integers into single bitfields for memory-constrained caches.
- [[Unexported Fields & Opaque Type Encapsulation (Pimpl Idiom in Go)]] — Information hiding, constructor validation functions, and preventing uninitialized struct states.

---

## 🔗 References
- ⬆️ Parent: [[Composite Types]]
- 📚 Module: `Language Basics`
