---
title: Language Internals
tags:
  - programming
  - language-internals
  - principal-swe
parent: "[[Programming]]"
---

# 💻 Language Internals

Compilers, AST, SSA intermediate representation, virtual machines (JIT), garbage collection, and memory allocators.

```text
Language Internals
│
├── [[Compilation Pipeline|01. Compilation Pipeline]]
│   ├── [[Lexical Analysis, Tokenization & Parsing (AST)]]
│   ├── [[Semantic Analysis, Type Checking & Symbol Tables]]
│   ├── [[Intermediate Representation (IR) & Static Single Assignment (SSA)]]
│   ├── [[Code Optimization Passes (Inlining, DCE, Loop Unrolling)]]
│   └── [[Machine Code Generation, Assemblers & Linkers]]
├── [[Runtime & Virtual Machines|02. Runtime & Virtual Machines]]
│   ├── [[Stack-Based vs Register-Based Virtual Machines (JVM, V8, BEAM)]]
│   ├── [[Just-In-Time (JIT) Compilation & Tiered Compilation]]
│   └── [[Interpreter Execution Loops & Bytecode Deserialization]]
└── [[Memory Management Engines|03. Memory Management Engines]]
│   ├── [[Garbage Collection Algorithms (Tricolor, Generational, G1, ZGC)]]
│   ├── [[Reference Counting & Cycle Collection (ARC, Python GC)]]
│   ├── [[Stack vs Heap Allocation & Escape Analysis]]
│   └── [[Memory Allocators (TCMalloc, jemalloc, Buddy Allocator)]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Compilation Pipeline|01. Compilation Pipeline]]
- [[Lexical Analysis, Tokenization & Parsing (AST)]] — Transforming character streams into tokens and building Abstract Syntax Trees via recursive descent or LR parsers.
- [[Semantic Analysis, Type Checking & Symbol Tables]] — Scope resolution, static type validation, type inference algorithms (Hindley-Milner), and symbol table management.
- [[Intermediate Representation (IR) & Static Single Assignment (SSA)]] — SSA form, control flow graphs (CFG), and phi-nodes enabling advanced compiler optimizations.
- [[Code Optimization Passes (Inlining, DCE, Loop Unrolling)]] — Compiler transformations: function inlining, dead-code elimination, escape analysis, and loop vectorization.
- [[Machine Code Generation, Assemblers & Linkers]] — Instruction selection, register allocation (graph coloring), relocations, and final executable linking.
### 2. 📂 [[Runtime & Virtual Machines|02. Runtime & Virtual Machines]]
- [[Stack-Based vs Register-Based Virtual Machines (JVM, V8, BEAM)]] — Comparing operand stack execution models with virtual register architectures for bytecode interpreters.
- [[Just-In-Time (JIT) Compilation & Tiered Compilation]] — Dynamic profiling, baseline JITs, optimizing JITs (HotSpot C2, V8 TurboFan), and deoptimization bailouts.
- [[Interpreter Execution Loops & Bytecode Deserialization]] — Direct threading, indirect threading, switch-based interpreters, and opcode dispatch optimizations.
### 3. 📂 [[Memory Management Engines|03. Memory Management Engines]]
- [[Garbage Collection Algorithms (Tricolor, Generational, G1, ZGC)]] — Mark-sweep, copying collectors, generational hypothesis, concurrent colored marking, and region-based collectors.
- [[Reference Counting & Cycle Collection (ARC, Python GC)]] — Deterministic object destruction, reference counting overhead, and resolving cyclic references via trial deletion.
- [[Stack vs Heap Allocation & Escape Analysis]] — Compiler flow-graph algorithms determining when object lifetimes outlive stack frames.
- [[Memory Allocators (TCMalloc, jemalloc, Buddy Allocator)]] — Segregated size classes, thread-local caching, virtual memory arenas, and mitigating heap fragmentation.

---

## 🔗 Navigation
- ⬆️ Parent: [[Programming]]
- 🎓 Root: [[Principal SWE]]
