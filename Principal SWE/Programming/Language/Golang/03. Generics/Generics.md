---
title: Generics
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Golang]]"
---

# 🧬 Generics

Parametric polymorphism in Go: type parameters, constraints, type inference, generic data structures, and GC shape stenciling.

```text
Generics
│
├── [[Core Concepts & Type Parameters|01. Core Concepts & Type Parameters]]
│   ├── [[Why Generics]]
│   ├── [[Generic Functions]]
│   ├── [[Generic Structs & Slices]]
│   ├── [[Type Constraints (any, comparable)]]
│   ├── [[Type Inference]]
│   └── [[Generic Instantiation & Type Arguments]]
├── [[Advanced Constraints & Data Structures|02. Advanced Constraints & Data Structures]]
│   ├── [[Type Sets and Union Constraints]]
│   ├── [[Approximation Element (~T)]]
│   ├── [[Generic Lock-Free Queue]]
│   ├── [[Generic Concurrent Skip List]]
│   ├── [[Generic Data Structures]]
│   ├── [[Generics vs Interfaces]]
│   ├── [[Generic Limitations]]
│   └── [[Recursive Type Constraints]]
└── [[Runtime & Compiler Mechanics|03. Runtime & Compiler Mechanics]]
│   ├── [[GcShape Stenciling & Dictionaries]]
│   ├── [[Generic Performance & Benchmarks]]
│   ├── [[Standard Library Generic Packages]]
│   ├── [[Generic Best Practices]]
│   └── [[Generic Anti-Patterns]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Core Concepts & Type Parameters|01. Core Concepts & Type Parameters]]
- [[Why Generics]] — Eliminating code duplication and dynamic type assertions without reflection overhead.
- [[Generic Functions]] — Syntax for declaring and invoking parameterized functions with [T any].
- [[Generic Structs & Slices]] — Parameterized structs, custom generic slice/map collection wrappers.
- [[Type Constraints (any, comparable)]] — Builtin constraints: any (interface{}), comparable (equality checkable).
- [[Type Inference]] — Function argument type inference and constraint type inference mechanics.
- [[Generic Instantiation & Type Arguments]] — Explicit vs implicit type argument instantiation.
### 2. 📂 [[Advanced Constraints & Data Structures|02. Advanced Constraints & Data Structures]]
- [[Type Sets and Union Constraints]] — Defining type sets with union operator (|) and custom constraint interfaces.
- [[Approximation Element (~T)]] — Allowing defined types with underlying type T using tilde operator (~int).
- [[Generic Lock-Free Queue]] — High-concurrency generic queue using atomic CAS operations.
- [[Generic Concurrent Skip List]] — Probabilistic search structure with concurrent lock-free reads.
- [[Generic Data Structures]] — Building generic Binary Trees, Linked Lists, Ring Buffers, and LRU Caches.
- [[Generics vs Interfaces]] — Architectural decision framework: compile-time parametric polymorphism vs dynamic polymorphism.
- [[Generic Limitations]] — No generic methods on non-generic types, no type assertions on type parameters.
- [[Recursive Type Constraints]] — Self-referential constraints (type Node[T Node[T]] interface).
### 3. 📂 [[Runtime & Compiler Mechanics|03. Runtime & Compiler Mechanics]]
- [[GcShape Stenciling & Dictionaries]] — Go compiler GcShape sharing and runtime dictionary parameter passing.
- [[Generic Performance & Benchmarks]] — Execution speed, binary size, and allocation profile of generics vs interfaces.
- [[Standard Library Generic Packages]] — slices, maps, and cmp standard packages algorithms.
- [[Generic Best Practices]] — When to use generics, API readability, maintaining simplicity.
- [[Generic Anti-Patterns]] — Over-parameterization, generic clutter, building complex type hierarchies.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
