---
title: Compiler Mechanics & GcShape Stenciling
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Compiler Mechanics & GcShape Stenciling

GcShape stenciling algorithm, dictionary passing, monomorphization comparisons, and binary size.

```text
Compiler Mechanics & GcShape Stenciling
│
├── [[GcShape Stenciling Algorithm]]
├── [[Dictionary Parameter Passing]]
├── [[Monomorphization vs Type Erasure vs GcShape]]
├── [[Binary Size Bloat & Compilation Time Analysis]]
└── [[Generics vs Interfaces Performance Benchmarks]]
```

---

## 🗂️ Topics

- [[GcShape Stenciling Algorithm]] — How Go shares machine code across all pointer types with identical GC shapes.
- [[Dictionary Parameter Passing]] — How the runtime passes type metadata dictionaries for scalar types at call sites.
- [[Monomorphization vs Type Erasure vs GcShape]] — Comparing C++ monomorphization, Java type erasure, and Go hybrid GcShape stenciling.
- [[Binary Size Bloat & Compilation Time Analysis]] — Analyzing the impact of generic instantiation on binary size and compilation speed.
- [[Generics vs Interfaces Performance Benchmarks]] — Zero-allocation execution, devirtualization, and 3-5x CPU performance gains over reflection.

---

## 🔗 References
- ⬆️ Parent: [[Generics]]

