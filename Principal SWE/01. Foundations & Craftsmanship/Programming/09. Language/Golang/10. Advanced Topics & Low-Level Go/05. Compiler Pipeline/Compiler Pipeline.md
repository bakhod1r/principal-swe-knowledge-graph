---
title: Compiler Pipeline, AST & Static Analysis
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# Compiler Pipeline, AST & Static Analysis

Compiler phases, AST manipulation, custom go/analysis analyzers, SSA passes, BCE optimization, inlining, and compiler pragmas.

```text
Compiler Pipeline, AST & Static Analysis
│
├── [[Compiler Pipeline Architecture (Scanning to Code Generation)]]
├── `AST Programmatic Manipulation (go-ast & go-parser)`
├── `Building Custom Analyzers with go-analysis Framework`
├── [[Compiler SSA (Static Single Assignment) Representation Passes]]
├── [[Bounds Check Elimination (BCE) Compiler Pass]]
├── [[Function Inlining Heuristics & Budget Calculation]]
└── `Compiler Directives & Pragmas (go:noinline, go:nosplit, go:linkname, go:notinheap)`
```

---

## 🗂️ Topics

- [[Compiler Pipeline Architecture (Scanning to Code Generation)]] — Lexer (scanner), Parser (go/ast), Type Checker (go/types), SSA middle-end, Assembler, and Linker.
- `AST Programmatic Manipulation (go-ast & go-parser)` — Parsing, traversing (ast.Walk), inspecting, and rewriting Go Abstract Syntax Trees.
- `Building Custom Analyzers with go-analysis Framework` — Writing custom enterprise static analysis linters checking architectural rules and conventions.
- [[Compiler SSA (Static Single Assignment) Representation Passes]] — Lowering AST into SSA basic blocks, optimization passes, and Plan 9 code emission.
- [[Bounds Check Elimination (BCE) Compiler Pass]] — Proving slice bounds to the compiler (-gcflags="-d=ssa/check_bce") to remove runtime bounds checks.
- [[Function Inlining Heuristics & Budget Calculation]] — Inlining cost scoring algorithm (80 AST node budget), mid-stack inlining, and inlining barriers.
- `Compiler Directives & Pragmas (go:noinline, go:nosplit, go:linkname, go:notinheap)` — Complete deep catalog of compiler pragmas controlling runtime code generation.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]

