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
│   └── [[Type Inference]]
├── [[Advanced Constraints & Data Structures|02. Advanced Constraints & Data Structures]]
│   ├── [[Type Sets and Union Constraints]]
│   ├── [[Generic Data Structures]]
│   ├── [[Generics vs Interfaces]]
│   └── [[Generic Limitations]]
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
- [[Why Generics]] — Solving code duplication without reflection or interface{} type assertions.
- [[Generic Functions]] — Syntax, type parameter lists, instantiating generic functions.
- [[Generic Structs & Slices]] — Defining generic structs, slices, maps, and interface constraints.
- [[Type Constraints (any, comparable)]] — Predeclared constraints: any, comparable; type set elements.
- [[Type Inference]] — Function argument type inference, constraint type inference.
### 2. 📂 [[Advanced Constraints & Data Structures|02. Advanced Constraints & Data Structures]]
- [[Type Sets and Union Constraints]] — Union constraints (|), approximation elements (~int), custom constraint interfaces.
- [[Generic Data Structures]] — Building generic Trees, Linked Lists, Ring Buffers, and Caches.
- [[Generics vs Interfaces]] — When to use type parameters vs dynamic interface polymorphism.
- [[Generic Limitations]] — No generic methods on non-generic types, no type assertions on type parameters.
### 3. 📂 [[Runtime & Compiler Mechanics|03. Runtime & Compiler Mechanics]]
- [[GcShape Stenciling & Dictionaries]] — How the Go compiler shares code between pointer types while passing dictionaries.
- [[Generic Performance & Benchmarks]] — Comparing generic execution speed vs interface{} vs concrete monomorphization.
- [[Standard Library Generic Packages]] — Standard library utility packages powered by type parameters.
- [[Generic Best Practices]] — Writing readable, maintainable generic APIs in Go.
- [[Generic Anti-Patterns]] — Over-parameterization, generic clutter, using generics where simple interfaces suffice.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

