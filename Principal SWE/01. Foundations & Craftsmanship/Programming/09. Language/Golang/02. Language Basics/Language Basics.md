---
title: Language Basics
type: index
tags:
  - golang
  - basics
  - principal-swe
parent: "[[Golang]]"
cssclasses:
  - index
---

# 🔤 Language Basics

Core language syntax: variables, constants, data types, composite types (Array, Slices, Maps, Structs), control flow, functions, calling conventions, pointers, and memory architecture.

```text
Language Basics
│
├── [[Variables & Constants|01. Variables & Constants]]
│   ├── `Var vs Short Declaration`
│   ├── `Zero Values`
│   ├── `Const & Iota`
│   ├── `Scope & Shadowing`
│   ├── `Blank Identifier`
│   ├── `Bitwise Operations`
│   └── `Memory Alignment & Atomic Safety`
├── [[Data Types|02. Data Types]]
│   ├── `Boolean`
│   ├── `Numeric Types`
│   ├── `Runes`
│   ├── `Strings`
│   ├── `Type Conversion`
│   ├── `Commands and Docs`
│   └── `String Internals`
├── [[Composite Types|03. Composite Types]]
│   ├── `Array`
│   ├── `Slices`
│   ├── `Maps`
│   └── `Structs`
├── [[Conditionals|04. Conditionals]]
│   ├── `If Statement`
│   ├── `Short Statement in If`
│   ├── `If-Else Chains`
│   └── `Switch Statement`
├── [[Loops|05. Loops]]
│   ├── `For Loop`
│   ├── `For Range`
│   └── `Control Flow & Jumps`
├── [[Functions|06. Functions]]
│   ├── `Function Basics`
│   ├── `Anonymous Functions & Closures`
│   ├── `Defer`
│   ├── `Init Function`
│   └── `Calling Conventions & Stacks`
└── [[Pointers|07. Pointers]]
    ├── `Pointer Basics`
    ├── `Pointers with Types`
    ├── `Unsafe Pointers & Low-Level`
    └── `Memory Management & Escape Analysis`
```

---

## 🗂️ Core Categories & Topics
### 1. 📂 [[Variables & Constants|01. Variables & Constants]]
- `Var vs Short Declaration` — Explicit var vs short := declaration syntax, package vs block scoping.
- `Zero Values` — Default zero initialization (0, false, "", nil) and useful zero values.
- `Const & Iota` — Untyped constant precision (256-bit), iota enumerators, and bitmask flag idioms.
- `Scope & Shadowing` — Lexical block scoping rules and shadowing traps.
- `Blank Identifier` — Discarding returns, side-effect imports, and compile-time interface assertion checks.
- `Bitwise Operations` — Bitwise AND/OR/XOR, bit clear (&^), bit shifts, and math/bits intrinsics.
- `Memory Alignment & Atomic Safety` — 64-bit alignment constraints on 32-bit architectures, atomic panics, and struct field reordering.

### 2. 📂 [[Data Types|02. Data Types]]
- `Boolean` — bool primitives, logical operators, and short-circuit evaluation.
- `Numeric Types` — Signed/unsigned integers, float32/float64 IEEE-754, NaN semantics, epsilon comparisons, and fixed-point decimals.
- `Runes` — Unicode code points, UTF-8 variable-width encoding, and unicode/utf8 package.
- `Strings` — String immutability, raw/interpreted literals, strings.Builder, and UTF-8 streaming decoders.
- `Type Conversion` — Explicit numeric casting, string-to-byte conversions, and strconv package.
- `Commands and Docs` — go doc, godoc, and package documentation standards.
- `String Internals` — stringStruct 2-word runtime layout and zero-copy unsafe conversions.

### 3. 📂 [[Composite Types|03. Composite Types]]
- `Array` — Fixed-length arrays, contiguous memory layout, value semantics, and comparison.
- `Slices` — Dynamic sliceHeader, capacity growth formula, 3-index slicing, and full slice tricks catalog.
- `Maps` — Hash table internals (hmap/bmap), Swiss Tables (Go 1.24+ SIMD control bytes), make preallocation, comma-ok, and concurrency crash safety.
- `Structs` — Memory layout, 8-byte word alignment padding, CPU cache lines, false sharing elimination, and empty structs.

### 4. 📂 [[Conditionals|04. Conditionals]]
- `If Statement` — if condition syntax, nested if flattening, and early return guard clauses.
- `Short Statement in If` — Scoped variable initialization preceding condition.
- `If-Else Chains` — Multi-way branching and happy path left-alignment.
- `Switch Statement` — Expression switch, tagless switch, fallthrough, type switches, and branch performance.

### 5. 📂 [[Loops|05. Loops]]
- `For Loop` — Three-component loop, while-style condition loop, and infinite loop.
- `For Range` — Iterating slices, maps, strings, channels, integers (1.22+), and push/pull iterators (1.23+ coroutines).
- `Control Flow & Jumps` — break, continue, goto cleanup idioms, and labeled jumps.

### 6. 📂 [[Functions|06. Functions]]
- `Function Basics` — Function declarations, multiple return values, named returns, variadics, pass-by-value.
- `Anonymous Functions & Closures` — Function literals, closures, variable capture, and heap escape.
- `Defer` — LIFO defer mechanics, argument evaluation timing, open-coded defers (Go 1.14+ zero-cost), and heap escapes.
- `Init Function` — Package initialization lifecycle and dependency graph execution ordering.
- `Calling Conventions & Stacks` — Register-based calling convention (ABIInternal), stack splitting, and morestack reallocations.

### 7. 📂 [[Pointers|07. Pointers]]
- `Pointer Basics` — Memory addresses, pointer types (*T), address-of (&), dereferencing (*), and nil safety.
- `Pointers with Types` — Pointers with structs, slices, maps, and pointer-to-interface traps.
- `Unsafe Pointers & Low-Level` — unsafe.Pointer, uintptr arithmetic, memory alignment, and zero-copy conversions.
- `Memory Management & Escape Analysis` — Stack vs heap allocation, escape analysis pointer flow graphs, inlining budget (max 80 AST nodes), and tricolor GC hybrid write barriers.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Topics

- [[Variables & Constants]]
- [[Data Types]]
- [[Composite Types]]
- [[Conditionals]]
- [[Loops]]
- [[Functions]]
- [[Pointers]]
