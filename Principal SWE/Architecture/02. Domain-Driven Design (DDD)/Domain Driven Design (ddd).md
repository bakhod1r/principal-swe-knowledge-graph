---
title: Domain Driven Design (ddd)
tags:
  - architecture
  - software-design
  - domain-driven-design-(ddd)
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Domain Driven Design (ddd)

Strategic and tactical Domain-Driven Design: Ubiquitous Language, Entity & Value Object invariants, Aggregate boundaries, Repositories, Domain Events, Bounded Contexts, and Context Mapping patterns.

```text
Domain Driven Design (ddd)
│
├── [[Ubiquitous Language and Event Storming|01. Ubiquitous Language and Domain Discovery]]
├── [[Entities, Value Objects, and Aggregate Roots|02. Entities, Value Objects, and Aggregates]]
├── [[Repositories and Domain Factories|03. Repositories and Factories]]
├── [[Domain Services and Asynchronous Domain Events|04. Domain Services and Domain Events]]
├── [[Bounded Contexts and Strategic Context Mapping|05. Bounded Contexts and Context Mapping]]
├── [[Strategic vs Tactical Domain Driven Design|06. Strategic vs Tactical Design]]
└── [[Enterprise Analysis Patterns in Domain Modeling|07. Analysis Patterns in DDD]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Ubiquitous Language and Event Storming|01. Ubiquitous Language and Domain Discovery]] — Collaborative domain discovery, eliminating translation overhead between business domain experts and software engineers.
- 📂 [[Entities, Value Objects, and Aggregate Roots|02. Entities, Value Objects, and Aggregates]] — Identity-based Entities, immutable Value Objects, transactional consistency boundaries, and Aggregate invariants.
- 📂 [[Repositories and Domain Factories|03. Repositories and Factories]] — Encapsulating persistence infrastructure behind collection-oriented repository interfaces and complex aggregate creation.
- 📂 [[Domain Services and Asynchronous Domain Events|04. Domain Services and Domain Events]] — Stateless cross-aggregate business logic and domain event publishing for eventual consistency across bounded contexts.
- 📂 [[Bounded Contexts and Strategic Context Mapping|05. Bounded Contexts and Context Mapping]] — Defining explicit boundary perimeters: Shared Kernel, Customer-Supplier, Conformist, Open Host Service (OHS), and Anti-Corruption Layer (ACL).
- 📂 [[Strategic vs Tactical Domain Driven Design|06. Strategic vs Tactical Design]] — Core Domain vs Supporting Subdomains vs Generic Subdomains; mapping business competitive advantage to technical investment.
- 📂 [[Enterprise Analysis Patterns in Domain Modeling|07. Analysis Patterns in DDD]] — Accountability patterns, Party models, Observation and Measurement patterns, and flexible domain structures.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]
- 🎓 Root: [[Principal SWE]]
