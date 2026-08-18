---
title: Blank Identifier
tags:
  - golang
  - variables
  - principal-swe
parent: "[[Variables & Constants]]"
---

# Blank Identifier

Discarding unused values, side-effect imports, and compile-time interface assertion checks.

```text
Blank Identifier
│
├── [[Discarding Unused Returns]]
├── [[Side-Effect Package Imports]]
└── [[Compile-Time Interface Verification]]
```

---

## 🗂️ Topics

- [[Discarding Unused Returns]] — Ignoring multiple return values or loop indices with blank identifier (_).
- [[Side-Effect Package Imports]] — Importing packages solely for their init() side-effects (_ "pkg").
- [[Compile-Time Interface Verification]] — Type assertion checks ensuring struct implements interface (var _ Interface = (*Struct)(nil)).

---

## 🔗 References
- ⬆️ Parent: [[Variables & Constants]]
- 🎓 Root: [[Principal SWE]]
