---
title: Minimal Version Selection (MVS)
tags:
  - golang
  - dependencies
  - principal-swe
parent: "[[Dependencies & Go Modules]]"
---

# Minimal Version Selection (MVS)

MVS deterministic algorithm, graph resolution, SemVer compatibility rules, and major version suffixes.

```text
Minimal Version Selection (MVS)
│
├── [[MVS Graph Algorithm Mechanics]]
├── [[Semantic Import Versioning (v2+ Rules)]]
└── [[retract Directive & Broken Releases]]
```

---

## 🗂️ Topics

- [[MVS Graph Algorithm Mechanics]] — Why Go rejected SAT-solvers in favor of deterministic Minimal Version Selection.
- [[Semantic Import Versioning (v2+ Rules)]] — Major version suffixes in import paths (pkg/v2) and backwards compatibility contracts.
- [[retract Directive & Broken Releases]] — Retracting broken or compromised module releases in go.mod.

---

## 🔗 References
- ⬆️ Parent: [[Dependencies & Go Modules]]

