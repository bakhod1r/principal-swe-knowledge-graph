- [[synctest Virtual Time Bubble Testing (Go 1.24+)]] — Deterministic, zero-wall-time testing of complex concurrent workflows using synctest.Run().

- [[Generic Type Aliases Migration Strategies]] — Refactoring monolithic shared libraries using Go 1.24 generic type aliases.

---
title: Type System & Stdlib Additions
tags:
  - golang
  - modern-go
  - principal-swe
parent: "[[Modern Language Features]]"
---

# Type System & Stdlib Additions

Generic type aliases (Go 1.24+), sync.Map modern methods, and newly introduced standard packages.

```text
Type System & Stdlib Additions
│
├── [[Generic Type Aliases (Go 1.24+)]]
├── [[slices and maps Standard Packages]]
├── [[cmp.Ordered and cmp.Compare]]
├── [[sync.OnceFunc, OnceValue, OnceValues (Go 1.21+)]]
└── [[synctest Experimental Testing Package]]
```

---

## 🗂️ Topics

- [[Generic Type Aliases (Go 1.24+)]] — type MyList[T] = other.List[T] generic alias syntax and migration patterns.
- [[slices and maps Standard Packages]] — Standard library algorithm functions for slices and maps without external dependencies.
- [[cmp.Ordered and cmp.Compare]] — Standard ordering comparison interface and three-way compare function.
- [[sync.OnceFunc, OnceValue, OnceValues (Go 1.21+)]] — Standard lazy initialization wrapper functions.
- [[synctest Experimental Testing Package]] — Virtual time testing package for deterministic testing of concurrent code.

---

## 🔗 References
- ⬆️ Parent: [[Modern Language Features]]
- 🎓 Root: [[Principal SWE]]
