---
title: Migrating Between Languages
tags:
  - programming
  - polyglot
  - principal-swe
parent: "[[Choosing a Language & Polyglot]]"
---

# Migrating Between Languages

Strangler Fig pattern, branch by abstraction, automated AST transpilers, shadow traffic dual-running, and wire protocol parity.

```text
Migrating Between Languages
│
├── [[Strangler Fig Pattern for Language Modernization]]
├── [[Branch by Abstraction in Multi-Language Refactoring]]
├── [[Automated Code Translation and AST Transpilers]]
├── [[Dual-Running and Shadow Traffic Verification]]
└── [[Data Migration and Wire Protocol Parity]]
```

---

## 🗂️ Topics

- [[Strangler Fig Pattern for Language Modernization]] — Incrementally replacing legacy language endpoints behind an API gateway with zero downtime.
- [[Branch by Abstraction in Multi-Language Refactoring]] — Decoupling language components inside monoliths using interface seams and intermediate adapters.
- [[Automated Code Translation and AST Transpilers]] — Feasibility and gotchas of automated source-to-source compilers vs manual architectural rewrites.
- [[Dual-Running and Shadow Traffic Verification]] — Replaying production traffic asynchronously to new language implementations to verify correctness and latency.
- [[Data Migration and Wire Protocol Parity]] — Ensuring binary serialization compatibility, timestamp precision, and floating-point equivalence across language implementations.

---

## 🔗 References
- ⬆️ Parent: [[Choosing a Language & Polyglot]]
- 🎓 Root: [[Principal SWE]]
