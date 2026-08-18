---
title: Tactical DDD
tags:
  - programming
  - oop
  - principal-swe
parent: "[[Object-Oriented Programming]]"
---

# Tactical DDD

Comprehensive engineering guide, patterns, and principles for Tactical DDD.

```text
Tactical DDD
│
├── [[Value Objects]]
├── [[Entities]]
├── [[Aggregates]]
├── [[Repository Concept]]
└── [[Domain Services]]
```

---

## 🗂️ Topics

- [[Value Objects]]
- [[Entities]]
- [[Aggregates]]
- [[Repository Concept]]
- [[Domain Services]]

- [[Aggregate Boundaries and Invariant Enforcement]] — Designing strict transactional boundaries around aggregates to guarantee domain invariants under high concurrency.
- [[Domain Events Publishing and Outbox Pattern Integration]] — Decoupling side effects by emitting immutable domain events from aggregates and persisting them transactionally via the Outbox pattern.
- [[Rich Domain Models vs Transaction Script Architecture]] — Encapsulating complex domain logic and state transitions inside entities vs procedural service layers.

---

## 🔗 References
- ⬆️ Parent: [[Object-Oriented Programming]]
- 🎓 Root: [[Principal SWE]]
