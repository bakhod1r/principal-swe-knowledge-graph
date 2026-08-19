---
title: Project Layouts & Repository Architectures
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Application Architecture]]"
---

# Project Layouts & Repository Architectures

Standard Go project layout (cmd, internal, pkg, api), Clean/Hexagonal Architecture, monorepos, and DDD bounded contexts.

```text
Project Layouts & Repository Architectures
│
├── [[Standard Go Project Layout (cmd, internal, pkg, api, web)]]
├── [[Flat vs Layered vs Modular Repository Architectures]]
├── [[Clean Architecture & Hexagonal Ports-and-Adapters]]
├── [[Enterprise Monorepos vs Multi-Repo Microservices]]
└── [[Domain-Driven Design (DDD) Bounded Contexts in Go]]
```

---

## 🗂️ Topics

- [[Standard Go Project Layout (cmd, internal, pkg, api, web)]] — Enterprise standard directory taxonomy and separation of operational entry points.
- [[Flat vs Layered vs Modular Repository Architectures]] — Staff-level evaluation: when to use flat single-package designs vs multi-module layouts.
- [[Clean Architecture & Hexagonal Ports-and-Adapters]] — Domain entities at the core, use case interactor layers, and pluggable infrastructure adapters.
- [[Enterprise Monorepos vs Multi-Repo Microservices]] — Tooling, workspace orchestration (go.work), shared libraries, and release boundaries.
- [[Domain-Driven Design (DDD) Bounded Contexts in Go]] — Structuring aggregates, value objects, domain events, and repositories within isolated packages.

---

## 🔗 References
- ⬆️ Parent: `Code Organization & Architecture`

