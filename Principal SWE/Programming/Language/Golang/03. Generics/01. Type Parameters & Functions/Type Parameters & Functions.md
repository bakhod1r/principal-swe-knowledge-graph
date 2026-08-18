---
title: Type Parameters & Functions
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Type Parameters & Functions

Parametric polymorphism motivation, generic function syntax, instantiation, type inference, and zero values.

```text
Type Parameters & Functions
│
├── [[Parametric Polymorphism Problem Statement]]
├── [[Generic Function Declarations & Invocation]]
├── [[Generic Method Limitations]]
├── [[Type Inference Algorithms]]
├── [[Instantiation Mechanics & Monomorphization]]
└── [[Generic Zero Value Idioms]]
```

---

## 🗂️ Topics

- [[Parametric Polymorphism Problem Statement]] — Why Go avoided generics for a decade and the problems type parameters solve.
- [[Generic Function Declarations & Invocation]] — Syntax for declaring and calling parameterized functions with [T any].
- [[Generic Method Limitations]] — Why Go does not allow parameterized methods on non-generic structs and function alternatives.
- [[Type Inference Algorithms]] — Function argument type inference and constraint type inference mechanics.
- [[Instantiation Mechanics & Monomorphization]] — Compile-time type parameter validation, substitution, and instantiation.
- [[Generic Zero Value Idioms]] — Returning zero values in generic functions using var zero T and new(T).

---

## 🔗 References
- ⬆️ Parent: [[Generics]]
- 🎓 Root: [[Principal SWE]]
