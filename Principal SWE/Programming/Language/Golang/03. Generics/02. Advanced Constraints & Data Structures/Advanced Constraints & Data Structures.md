---
title: Advanced Constraints & Data Structures
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Advanced Constraints & Data Structures

Type sets, approximation elements (~T), generic collections, and architectural tradeoffs.

```text
Advanced Constraints & Data Structures
│
├── [[Type Sets and Union Constraints]]
├── [[Approximation Element (~T)]]
├── [[Generic Lock-Free Queue]]
├── [[Generic Concurrent Skip List]]
├── [[Generic Data Structures]]
├── [[Generics vs Interfaces]]
├── [[Generic Limitations]]
└── [[Recursive Type Constraints]]
```

---

## 🗂️ Topics

- [[Type Sets and Union Constraints]] — Defining type sets with union operator (|) and custom constraint interfaces.
- [[Approximation Element (~T)]] — Allowing defined types with underlying type T using tilde operator (~int).
- [[Generic Lock-Free Queue]] — High-concurrency generic queue using atomic CAS operations.
- [[Generic Concurrent Skip List]] — Probabilistic search structure with concurrent lock-free reads.
- [[Generic Data Structures]] — Building generic Binary Trees, Linked Lists, Ring Buffers, and LRU Caches.
- [[Generics vs Interfaces]] — Architectural decision framework: compile-time parametric polymorphism vs dynamic polymorphism.
- [[Generic Limitations]] — No generic methods on non-generic types, no type assertions on type parameters.
- [[Recursive Type Constraints]] — Self-referential constraints (type Node[T Node[T]] interface).

---

## 🔗 References
- ⬆️ Parent: [[Generics]]
- 🎓 Root: [[Principal SWE]]
