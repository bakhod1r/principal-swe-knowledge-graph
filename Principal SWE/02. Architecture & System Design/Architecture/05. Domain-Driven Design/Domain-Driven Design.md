---
title: Domain-Driven Design
tags:
  - architecture
  - systems-architecture
  - domain-driven-design-(ddd)-and-strategic-modeling
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Domain Driven Design (ddd) & Strategic Modeling

Eric Evans Strategic and Tactical DDD: Ubiquitous Language, Bounded Contexts, Context Mapping, Entities, Value Objects, Aggregates & Invariants, Domain Events, Repositories, Domain Services, and Anti-Corruption Layers (ACL).

```text
Domain Driven Design (ddd) & Strategic Modeling
│
├── `01. Ubiquitous Language and Core Domain Identification`
├── `02. Bounded Contexts and Domain Boundaries`
├── `03. Context Mapping Strategies and Integration Relationships`
├── [[Entities: Unique Identity, Lifecycles, and Mutable State Invariants|04. Entities, Identity, and Mutability Lifecycles]]
├── [[Value Objects: Immutability, Structural Equality, and Side Effect Free Functions|05. Value Objects and Immutability Standards]]
├── [[Aggregates and Aggregate Roots: Transactional and Consistency Boundaries|06. Aggregates, Aggregate Roots, and Invariant Boundaries]]
├── [[Domain Events: Capturing Business State Changes and Event Publication|07. Domain Events and State Mutation Propagation]]
├── [[Repositories and Factories: Encapsulating Aggregate Lifecycle and Persistence|08. Repositories, Factories, and Aggregate Persistence]]
├── [[Domain Services vs Application Services: Pure Domain Logic vs Orchestration|09. Domain Services vs Application Services]]
└── `10. Event Storming and Collaborative Domain Modeling`
```

---

## 🗂️ Core Knowledge Domains

- 📂 `01. Ubiquitous Language and Core Domain Identification` — Creating a shared model-driven vocabulary between domain experts and software engineers, eliminating translation ambiguities, and identifying the Core Domain.
- 📂 `02. Bounded Contexts and Domain Boundaries` — Defining explicit boundaries where a specific domain model applies, separating divergent concepts (e.g. User in Auth vs Customer in Billing), and context isolation.
- 📂 `03. Context Mapping Strategies and Integration Relationships` — Modeling relationships between Bounded Contexts: Partnership, Shared Kernel, Customer-Supplier, Conformist, Open Host Service (OHS), Published Language, and Separate Ways.
- 📂 [[Entities: Unique Identity, Lifecycles, and Mutable State Invariants|04. Entities, Identity, and Mutability Lifecycles]] — Distinguishing entities by unique persistent ID rather than attributes, managing lifecycle transitions, and enforcing domain rules during mutations.
- 📂 [[Value Objects: Immutability, Structural Equality, and Side Effect Free Functions|05. Value Objects and Immutability Standards]] — Self-validating immutable domain building blocks (Money, Address, Email), attribute-based equality, whole value encapsulation, and eliminating primitive obsession.
- 📂 [[Aggregates and Aggregate Roots: Transactional and Consistency Boundaries|06. Aggregates, Aggregate Roots, and Invariant Boundaries]] — Enforcing transactional consistency within a single aggregate, selecting the Aggregate Root, referencing other aggregates only by ID, and small aggregate sizing.
- 📂 [[Domain Events: Capturing Business State Changes and Event Publication|07. Domain Events and State Mutation Propagation]] — Modeling significant business occurrences as immutable past events (`OrderPlaced`, `PaymentReceived`), in-process event dispatching, and asynchronous propagation.
- 📂 [[Repositories and Factories: Encapsulating Aggregate Lifecycle and Persistence|08. Repositories, Factories, and Aggregate Persistence]] — Collection-oriented vs persistence-oriented repository interfaces, reconstructing aggregates via Factories, and avoiding leaky database queries in repositories.
- 📂 [[Domain Services vs Application Services: Pure Domain Logic vs Orchestration|09. Domain Services vs Application Services]] — Pure stateless business logic involving multiple aggregates (Domain Services) vs use-case transaction orchestration, security, and DTO conversion (Application Services).
- 📂 `10. Event Storming and Collaborative Domain Modeling` — Facilitating collaborative rapid modeling with domain experts: mapping Domain Events (orange), Commands (blue), Aggregates (yellow), Read Models (green), and Policies (lilac).

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

