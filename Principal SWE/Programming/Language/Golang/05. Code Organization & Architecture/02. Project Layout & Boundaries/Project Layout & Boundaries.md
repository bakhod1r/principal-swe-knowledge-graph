---
title: Project Layout & Boundaries
tags:
  - golang
  - architecture
parent: "[[Code Organization & Architecture]]"
---

# Project Layout & Boundaries

Standard Go project layout, internal visibility enforcement, and workspace development.

```text
Project Layout & Boundaries
│
├── [[Standard Go Project Layout]]
├── [[internal/ Visibility Enforcement]]
├── [[Workspaces (go.work)]]
├── [[Circular Dependency Prevention]]
└── [[Package Cohesion & Coupling]]
```

---

## 🗂️ Topics

- [[Standard Go Project Layout]] — cmd/, internal/, pkg/, api/ directory layout conventions.
- [[internal/ Visibility Enforcement]] — Compiler-enforced access control preventing unauthorized external imports.
- [[Workspaces (go.work)]] — Multi-module local development without modifying go.mod replace directives.
- [[Circular Dependency Prevention]] — Avoiding import cycles through interface abstraction and layered architecture.
- [[Package Cohesion & Coupling]] — Designing focused, loosely-coupled packages.

---

## 🔗 References
- ⬆️ Parent: [[Code Organization & Architecture]]
- 🎓 Root: [[Principal SWE]]
