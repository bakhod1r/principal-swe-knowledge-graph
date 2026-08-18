---
title: Runtime & Compiler Mechanics
tags:
  - golang
  - generics
parent: "[[Generics]]"
---

# Runtime & Compiler Mechanics

GcShape stenciling, dictionary passing, monomorphization tradeoffs, and best practices.

```text
Runtime & Compiler Mechanics
│
├── [[GcShape Stenciling & Dictionaries]]
├── [[Generic Performance & Benchmarks]]
├── [[Standard Library Generics (slices, maps, cmp)]]
├── [[Generic Best Practices]]
└── [[Generic Anti-Patterns]]
```

---

## 🗂️ Topics

- [[GcShape Stenciling & Dictionaries]] — How the Go compiler shares code between pointer types while passing dictionaries.
- [[Generic Performance & Benchmarks]] — Comparing generic execution speed vs interface{} vs concrete monomorphization.
- [[Standard Library Generics (slices, maps, cmp)]] — Standard library utility packages powered by type parameters.
- [[Generic Best Practices]] — Writing readable, maintainable generic APIs in Go.
- [[Generic Anti-Patterns]] — Over-parameterization, generic clutter, using generics where simple interfaces suffice.

---

## 🔗 References
- ⬆️ Parent: [[Generics]]
- 🎓 Root: [[Principal SWE]]
