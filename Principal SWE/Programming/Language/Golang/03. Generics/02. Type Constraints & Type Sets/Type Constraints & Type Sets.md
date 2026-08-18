---
title: Type Constraints & Type Sets
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Type Constraints & Type Sets

Predeclared constraints (any, comparable), type sets, union elements, tilde operator (~T), and recursive constraints.

```text
Type Constraints & Type Sets
│
├── [[Predeclared Constraints (any vs comparable)]]
├── [[comparable Contract & Interface Equality Pitfalls]]
├── [[Type Sets Mathematical Model]]
├── [[Union Element (Pipe) & Intersection Sets]]
├── [[Underlying Types & Tilde Operator (~T)]]
├── [[cmp.Ordered Constraint]]
├── [[Recursive Constraint Interfaces]]
└── [[Structural Constraint Interfaces]]
```

---

## 🗂️ Topics

- [[Predeclared Constraints (any vs comparable)]] — Contrasting any with comparable and interface equality constraints.
- [[comparable Contract & Interface Equality Pitfalls]] — Why comparing non-comparable concrete types stored in interface fields panics at runtime.
- [[Type Sets Mathematical Model]] — Interfaces as mathematical type sets rather than purely method sets.
- [[Union Element (Pipe) & Intersection Sets]] — Defining type unions (int | int64 | float64) and intersection rules.
- [[Underlying Types & Tilde Operator (~T)]] — Matching user-defined custom types with underlying primitives using ~T.
- [[cmp.Ordered Constraint]] — Standard constraint for types supporting <, <=, >, and >= comparison operators.
- [[Recursive Constraint Interfaces]] — Self-referential constraints (type Node[T Node[T]] interface).
- [[Structural Constraint Interfaces]] — Combining method signatures and type elements within a single interface.

---

## 🔗 References
- ⬆️ Parent: [[Generics]]

