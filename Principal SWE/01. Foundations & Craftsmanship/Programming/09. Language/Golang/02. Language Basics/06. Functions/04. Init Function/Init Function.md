---
title: Init Function
tags:
  - golang
  - functions
  - principal-swe
parent: "[[Functions (Clean Code)]]"
---

# Init Function

Package initialization lifecycle, execution ordering, and dependency graphs.

```text
Init Function
│
├── [[init() Function Lifecycle]]
├── [[Package Initialization Dependency Graph]]
├── [[Multiple init() Functions]]
└── [[Side-Effect Imports]]
```

---

## 🗂️ Topics

- [[init() Function Lifecycle]] — Automatic execution during package loading before main.main().
- [[Package Initialization Dependency Graph]] — Topological execution order of variable initializations and init() calls.
- [[Multiple init() Functions]] — Defining multiple init() functions across different files within the same package.
- [[Side-Effect Imports]] — Importing packages exclusively for init() registration (_ "net/http/pprof").

---

## 🔗 References
- ⬆️ Parent: [[Functions (Clean Code)]]

