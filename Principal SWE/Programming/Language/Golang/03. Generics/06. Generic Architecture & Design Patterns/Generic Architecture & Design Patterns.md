---
title: Generic Architecture & Design Patterns
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Generic Architecture & Design Patterns

Clean Architecture generic repository, unit of work, builders, event buses, and object pools.

```text
Generic Architecture & Design Patterns
│
├── [[Generic Repository Pattern in Clean Architecture]]
├── [[Generic Unit of Work Pattern]]
├── [[Generic Builder & Functional Options Pattern]]
├── [[Generic Event Bus & PubSub Pipeline]]
├── [[Generic Middleware & Handler Pipeline]]
└── [[Generic Object Pool (sync.Pool Wrapper)]]
```

---

## 🗂️ Topics

- [[Generic Repository Pattern in Clean Architecture]] — Universal CRUD database repository interfaces (Repository[T, ID]).
- [[Generic Unit of Work Pattern]] — Managing multi-repository database transactions with generic contracts.
- [[Generic Builder & Functional Options Pattern]] — Constructing complex typed domain objects with validation.
- [[Generic Event Bus & PubSub Pipeline]] — Type-safe event dispatcher and topic subscriber routing in Go.
- [[Generic Middleware & Handler Pipeline]] — Composable request/response middleware chains without interface boxing.
- [[Generic Object Pool (sync.Pool Wrapper)]] — Zero-allocation typed object reuse wrapping sync.Pool.

---

## 🔗 References
- ⬆️ Parent: [[Generics]]
- 🎓 Root: [[Principal SWE]]
