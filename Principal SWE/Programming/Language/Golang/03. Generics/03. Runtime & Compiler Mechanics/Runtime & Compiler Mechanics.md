- [[Generic Zero Value Idiom]] — Returning zero values in generic functions using var zero T pattern.

- [[Dictionary Parameter Passing & Pointer Unification]] — How runtime unifies pointer types under single GcShape while passing dictionaries for scalars.

---
title: Runtime & Compiler Mechanics
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Runtime & Compiler Mechanics

GcShape stenciling, dictionary passing, monomorphization tradeoffs, and best practices.

```text
Runtime & Compiler Mechanics
│
├── [[GcShape Stenciling & Dictionaries]]
├── [[Generic Performance & Benchmarks]]
├── [[Standard Library Generic Packages]]
├── [[Generic Best Practices]]
└── [[Generic Anti-Patterns]]
```

---

## 🗂️ Topics

- [[GcShape Stenciling & Dictionaries]] — Go compiler GcShape sharing and runtime dictionary parameter passing.
- [[Generic Performance & Benchmarks]] — Execution speed, binary size, and allocation profile of generics vs interfaces.
- [[Standard Library Generic Packages]] — slices, maps, and cmp standard packages algorithms.
- [[Generic Best Practices]] — When to use generics, API readability, maintaining simplicity.
- [[Generic Anti-Patterns]] — Over-parameterization, generic clutter, building complex type hierarchies.

---

## 🔗 References
- ⬆️ Parent: [[Generics]]
- 🎓 Root: [[Principal SWE]]
