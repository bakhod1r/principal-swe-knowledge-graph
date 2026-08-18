---
title: Language Basics
tags:
  - golang
  - basics
  - principal-swe
parent: "[[Golang]]"
---

# 🔤 Language Basics

Core language syntax, environment setup, toolchain commands, modules, GOROOT layout, variables, data types, composite types (Array, Slices, Maps, Structs), control flow, functions, calling conventions, pointers, and memory architecture.

```text
Language Basics
│
├── [[Settings Environment|01. Settings Environment]]
│   ├── [[Core Environment Variables]]
│   ├── [[Target OS & Architecture]]
│   └── [[Build Caching & Reproducibility]]
├── [[Go Commands|02. Go command]]
│   ├── [[Core CLI Commands]]
│   ├── [[Compiler Pragmas & Directives (go:)]]
│   └── [[Binary Inspection Tools (go tool)]]
├── [[Dependencies & Go Modules|03. Dependencies]]
│   ├── [[Module Files & Checksums]]
│   ├── [[Minimal Version Selection (MVS)]]
│   └── [[Enterprise & Monorepos]]
├── [[Go Source Code Structure|04. Go Source Code Structure]]
│   ├── [[Source Hierarchy]]
│   └── [[Standard Package Organization]]
├── [[Variables & Constants|05. Variables & Constants]]
│   ├── [[Var vs Short Declaration]]
│   ├── [[Zero Values]]
│   ├── [[Const & Iota]]
│   ├── [[Scope & Shadowing]]
│   ├── [[Blank Identifier]]
│   ├── [[Bitwise Operations]]
│   └── [[Memory Alignment & Atomic Safety]]
├── [[Data Types|06. Data Types]]
│   ├── [[Boolean]]
│   ├── [[Numeric Types]]
│   ├── [[Runes]]
│   ├── [[Strings]]
│   ├── [[Type Conversion]]
│   ├── [[Commands and Docs]]
│   └── [[String Internals]]
├── [[Composite Types|07. Composite Types]]
│   ├── [[Array]]
│   ├── [[Slices]]
│   ├── [[Maps]]
│   └── [[Structs]]
├── [[Conditionals|08. Conditionals]]
│   ├── [[If Statement]]
│   ├── [[Short Statement in If]]
│   ├── [[If-Else Chains]]
│   └── [[Switch Statement]]
├── [[Loops|09. Loops]]
│   ├── [[For Loop]]
│   ├── [[For Range]]
│   └── [[Control Flow & Jumps]]
├── [[Functions|10. Functions]]
│   ├── [[Function Basics]]
│   ├── [[Anonymous Functions & Closures]]
│   ├── [[Defer]]
│   ├── [[Init Function]]
│   └── [[Calling Conventions & Stacks]]
└── [[Pointers|11. Pointers]]
    ├── [[Pointer Basics]]
    ├── [[Pointers with Types]]
    ├── [[Unsafe Pointers & Low-Level]]
    └── [[Memory Management & Escape Analysis]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Settings Environment|01. Settings Environment]]
- [[Core Environment Variables]] — PATH, GOROOT, GOPATH, GOBIN, GOENV, and Shell profile persistence.
- [[Target OS & Architecture]] — Cross-compilation matrix, CGO static vs dynamic linking, libc/musl scratch containers, and DNS resolvers (cgo vs netgo).
- [[Build Caching & Reproducibility]] — GOCACHE action/output caching, reproducible builds (-trimpath), and VCS provenance via debug.ReadBuildInfo().

### 2. 📂 [[Go Commands|02. Go command]]
- [[Core CLI Commands]] — go build, run, install, test, fmt/gofumpt, vet, doc, clean, and go.work.
- [[Compiler Pragmas & Directives (go:)]] — go:noinline, go:nosplit, go:linkname, go:uintptrescapes, go:notinheap, go:generate, and go:embed.
- [[Binary Inspection Tools (go tool)]] — go tool compile -S (Plan 9 assembly), go tool objdump disassembly, go tool nm symbol tables, and go tool cgo.

### 3. 📂 [[Dependencies & Go Modules|03. Dependencies]]
- [[Module Files & Checksums]] — go.mod directives, go.sum SHA-256 tree verification, GOPROXY mirrors, and GOSUMDB notary.
- [[Minimal Version Selection (MVS)]] — MVS graph algorithm mechanics, semantic import versioning (v2+ rules), and retract directives.
- [[Enterprise & Monorepos]] — GOPRIVATE corporate setups, vendoring (go mod vendor), multi-module workspaces, and govulncheck/SBOM.

### 4. 📂 [[Go Source Code Structure|04. Go Source Code Structure]]
- [[Source Hierarchy]] — Root src tree, runtime sources (proc.go, mgc.go, malloc.go, chan.go), compiler toolchain sources (cmd/), and internal packages.
- [[Standard Package Organization]] — Layered dependency trees, cycle avoidance, and internal/ shared helper package design.

### 5. 📂 [[Variables & Constants|05. Variables & Constants]]
- [[Var vs Short Declaration]] — Explicit var vs short := declaration syntax, package vs block scoping.
- [[Zero Values]] — Default zero initialization (0, false, "", nil) and useful zero values.
- [[Const & Iota]] — Untyped constant precision (256-bit), iota enumerators, and bitmask flag idioms.
- [[Scope & Shadowing]] — Lexical block scoping rules and shadowing traps.
- [[Blank Identifier]] — Discarding returns, side-effect imports, and compile-time interface assertion checks.
- [[Bitwise Operations]] — Bitwise AND/OR/XOR, bit clear (&^), bit shifts, and math/bits intrinsics.
- [[Memory Alignment & Atomic Safety]] — 64-bit alignment constraints on 32-bit architectures, atomic panics, and struct field reordering.

### 6. 📂 [[Data Types|06. Data Types]]
- [[Boolean]] — bool primitives, logical operators, and short-circuit evaluation.
- [[Numeric Types]] — Signed/unsigned integers, float32/float64 IEEE-754, NaN semantics, epsilon comparisons, and fixed-point decimals.
- [[Runes]] — Unicode code points, UTF-8 variable-width encoding, and unicode/utf8 package.
- [[Strings]] — String immutability, raw/interpreted literals, strings.Builder, and UTF-8 streaming decoders.
- [[Type Conversion]] — Explicit numeric casting, string-to-byte conversions, and strconv package.
- [[Commands and Docs]] — go doc, godoc, and package documentation standards.
- [[String Internals]] — stringStruct 2-word runtime layout and zero-copy unsafe conversions.

### 7. 📂 [[Composite Types|07. Composite Types]]
- [[Array]] — Fixed-length arrays, contiguous memory layout, value semantics, and comparison.
- [[Slices]] — Dynamic sliceHeader, capacity growth formula, 3-index slicing, and full slice tricks catalog.
- [[Maps]] — Hash table internals (hmap/bmap), Swiss Tables (Go 1.24+ SIMD control bytes), make preallocation, comma-ok, and concurrency crash safety.
- [[Structs]] — Memory layout, 8-byte word alignment padding, CPU cache lines, false sharing elimination, and empty structs.

### 8. 📂 [[Conditionals|08. Conditionals]]
- [[If Statement]] — if condition syntax, nested if flattening, and early return guard clauses.
- [[Short Statement in If]] — Scoped variable initialization preceding condition.
- [[If-Else Chains]] — Multi-way branching and happy path left-alignment.
- [[Switch Statement]] — Expression switch, tagless switch, fallthrough, type switches, and branch performance.

### 9. 📂 [[Loops|09. Loops]]
- [[For Loop]] — Three-component loop, while-style condition loop, and infinite loop.
- [[For Range]] — Iterating slices, maps, strings, channels, integers (1.22+), and push/pull iterators (1.23+ coroutines).
- [[Control Flow & Jumps]] — break, continue, goto cleanup idioms, and labeled jumps.

### 10. 📂 [[Functions|10. Functions]]
- [[Function Basics]] — Function declarations, multiple return values, named returns, variadics, pass-by-value.
- [[Anonymous Functions & Closures]] — Function literals, closures, variable capture, and heap escape.
- [[Defer]] — LIFO defer mechanics, argument evaluation timing, open-coded defers (Go 1.14+ zero-cost), and heap escapes.
- [[Init Function]] — Package initialization lifecycle and dependency graph execution ordering.
- [[Calling Conventions & Stacks]] — Register-based calling convention (ABIInternal), stack splitting, and morestack reallocations.

### 11. 📂 [[Pointers|11. Pointers]]
- [[Pointer Basics]] — Memory addresses, pointer types (*T), address-of (&), dereferencing (*), and nil safety.
- [[Pointers with Types]] — Pointers with structs, slices, maps, and pointer-to-interface traps.
- [[Unsafe Pointers & Low-Level]] — unsafe.Pointer, uintptr arithmetic, memory alignment, and zero-copy conversions.
- [[Memory Management & Escape Analysis]] — Stack vs heap allocation, escape analysis pointer flow graphs, inlining budget (max 80 AST nodes), and tricolor GC hybrid write barriers.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

