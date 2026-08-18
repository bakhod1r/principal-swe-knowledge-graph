---
title: Limitations, Gotchas & Anti-Patterns
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Limitations, Gotchas & Anti-Patterns

Generic method restrictions, type assertions on type parameters, over-parameterization, and Go 1.24 generic aliases.

```text
Limitations, Gotchas & Anti-Patterns
│
├── [[No Generic Methods on Non-Generic Types Trap]]
├── [[No Type Assertions on Type Parameters]]
├── [[The Over-Parameterization Anti-Pattern]]
├── [[Interface vs Generic Decision Tree]]
└── [[Generic Type Aliases (Go 1.24+)]]
```

---

## 🗂️ Topics

- [[No Generic Methods on Non-Generic Types Trap]] — Why methods cannot introduce new type parameters and package-level function workarounds.
- [[No Type Assertions on Type Parameters]] — Why T.(int) is illegal and how to perform type dispatching via interface boxing.
- [[The Over-Parameterization Anti-Pattern]] — Avoiding unnecessary type parameter clutter that harms API readability.
- [[Interface vs Generic Decision Tree]] — Staff-level decision framework: when to choose dynamic interfaces vs static generics.
- [[Generic Type Aliases (Go 1.24+)]] — type Set[T] = map[T]struct{} generic type alias syntax and codebase refactoring.

---

## 🔗 References
- ⬆️ Parent: [[Generics]]
- 🎓 Root: [[Principal SWE]]
