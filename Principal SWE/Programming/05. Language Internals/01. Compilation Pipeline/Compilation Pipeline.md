---
title: Compilation Pipeline
tags:
  - programming
  - language-internals
  - principal-swe
parent: "[[Language Internals]]"
---

# Compilation Pipeline

From source code to machine assembly.

```text
Compilation Pipeline
│
├── [[Lexical Analysis, Tokenization & Parsing (AST)]]
├── [[Semantic Analysis, Type Checking & Symbol Tables]]
├── [[Intermediate Representation (IR) & Static Single Assignment (SSA)]]
├── [[Code Optimization Passes (Inlining, DCE, Loop Unrolling)]]
└── [[Machine Code Generation, Assemblers & Linkers]]
```

---

## 🗂️ Topics

- [[Lexical Analysis, Tokenization & Parsing (AST)]] — Transforming character streams into tokens and building Abstract Syntax Trees via recursive descent or LR parsers.
- [[Semantic Analysis, Type Checking & Symbol Tables]] — Scope resolution, static type validation, type inference algorithms (Hindley-Milner), and symbol table management.
- [[Intermediate Representation (IR) & Static Single Assignment (SSA)]] — SSA form, control flow graphs (CFG), and phi-nodes enabling advanced compiler optimizations.
- [[Code Optimization Passes (Inlining, DCE, Loop Unrolling)]] — Compiler transformations: function inlining, dead-code elimination, escape analysis, and loop vectorization.
- [[Machine Code Generation, Assemblers & Linkers]] — Instruction selection, register allocation (graph coloring), relocations, and final executable linking.

---

## 🔗 References
- ⬆️ Parent: [[Language Internals]]
- 🎓 Root: [[Principal SWE]]
