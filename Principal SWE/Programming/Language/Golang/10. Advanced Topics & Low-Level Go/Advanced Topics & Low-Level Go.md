---
title: Advanced Topics & Low-Level Go
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Golang]]"
---

# 🔬 Advanced Topics & Low-Level Go

Memory management internals, escape analysis, reflect, unsafe, Cgo FFI, Plan 9 assembly, compiler flags, and profile-guided optimization (PGO).

```text
Advanced Topics & Low-Level Go
│
├── [[Memory Management & Compiler|01. Memory Management & Compiler]]
│   ├── [[Memory Allocation Hierarchy (mcache, mcentral, mheap)]]
│   ├── [[Escape Analysis Algorithms]]
│   ├── [[Compiler SSA Optimization Passes]]
│   ├── [[Compiler & Linker Flags]]
│   └── [[Profile-Guided Optimization (PGO)]]
├── [[Metaprogramming & Low-Level|02. Metaprogramming & Low-Level]]
│   ├── [[reflect.Type and reflect.Value]]
│   ├── [[unsafe.Pointer & uintptr Arithmetic]]
│   ├── [[unsafe.Slice and unsafe.String]]
│   ├── [[go:linkname Compiler Directive]]
│   └── [[Go Plugins (.so) Dynamic Loading]]
├── [[FFI & Low-Level Assembly|03. FFI & Low-Level Assembly]]
│   ├── [[Cgo Architecture & Overhead]]
│   ├── [[Plan 9 Go Assembly]]
│   ├── [[Go ABIInternal Register Calling Convention]]
│   └── [[CPU Feature Detection (internal-cpu)]]
└── [[Compiler Pipeline & Code Generation|04. Compiler Pipeline & Code Generation]]
│   ├── [[Compiler Pipeline Overview]]
│   ├── [[AST Inspection & go-ast]]
│   ├── [[Custom Static Analyzers with go-analysis]]
│   ├── [[Bounds Check Elimination (BCE) Deep Dive]]
│   ├── [[Function Inlining Heuristics]]
│   └── [[Dead Code Elimination & Static Branch Pruning]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Memory Management & Compiler|01. Memory Management & Compiler]]
- [[Memory Allocation Hierarchy (mcache, mcentral, mheap)]] — TCMalloc-inspired per-P caches, central spans, size classes, and page arenas.
- [[Escape Analysis Algorithms]] — Compiler escape analysis rules (go build -gcflags="-m -m"), leaking params, flow analysis.
- [[Compiler SSA Optimization Passes]] — Static Single Assignment intermediate representation, Dead Code Elimination, Bounds Check Elimination (BCE).
- [[Compiler & Linker Flags]] — -gcflags="-N -l", -ldflags="-s -w -X main.version=1.0", -trimpath reproducibility.
- [[Profile-Guided Optimization (PGO)]] — Feeding production CPU profiles into go build for 2-14% compiler speedups.
### 2. 📂 [[Metaprogramming & Low-Level|02. Metaprogramming & Low-Level]]
- [[reflect.Type and reflect.Value]] — Deep type introspection, struct field inspection, method invocation, performance overhead.
- [[unsafe.Pointer & uintptr Arithmetic]] — Direct memory manipulation, computing struct field offsets, casting pointers.
- [[unsafe.Slice and unsafe.String]] — Zero-copy conversion between byte slices and strings without heap allocation.
- [[go:linkname Compiler Directive]] — Linking to unexported runtime and standard library functions across package boundaries.
- [[Go Plugins (.so) Dynamic Loading]] — Compiling and dynamically loading shared object plugins at runtime via plugin package.
### 3. 📂 [[FFI & Low-Level Assembly|03. FFI & Low-Level Assembly]]
- [[Cgo Architecture & Overhead]] — Calling C libraries from Go, Cgo stack switching overhead, memory pinning rules (runtime.Pinner).
- [[Plan 9 Go Assembly]] — Writing Go assembly functions, pseudo-registers (FP, SP, SB, PC), instruction syntax.
- [[Go ABIInternal Register Calling Convention]] — Passing function arguments and returns in CPU registers instead of stack frames.
- [[CPU Feature Detection (internal-cpu)]] — Detecting hardware AVX, SSE, AES-NI CPU instructions at runtime.
### 4. 📂 [[Compiler Pipeline & Code Generation|04. Compiler Pipeline & Code Generation]]
- [[Compiler Pipeline Overview]] — Lexer (scanner), Parser (AST), Type Checking, IR/Middle-end, SSA Backend, Assembler, Linker.
- [[AST Inspection & go-ast]] — Parsing Go source code programmatically with go/parser and go/ast.
- [[Custom Static Analyzers with go-analysis]] — Building custom linter passes with the standard go/analysis framework.
- [[Bounds Check Elimination (BCE) Deep Dive]] — Proving slice bounds to the compiler to remove runtime bounds checking panic branches.
- [[Function Inlining Heuristics]] — Inlining cost budget (max 80 nodes), mid-stack inlining, and //go:noinline directive.
- [[Dead Code Elimination & Static Branch Pruning]] — Compiler pruning of unreached code branches based on constant conditions.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
