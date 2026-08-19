---
title: Clean, Hexagonal & Onion Architecture
tags:
  - architecture
  - systems-architecture
  - clean,-hexagonal-and-onion-architecture
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Clean, Hexagonal & Onion Architecture

Inversion-of-Control and boundary architectures: The Dependency Rule, Hexagonal (Ports & Adapters), Onion Architecture, Clean Architecture, Entities & Use Cases, Repositories, and Component Cohesion principles (REP, CCP, CRP).

```text
Clean, Hexagonal & Onion Architecture
│
├── [[The Dependency Rule, Concentric Circles, and Inversion of Control|01. The Dependency Rule and Inversion of Control]]
├── [[Entities, Enterprise Business Rules, and Critical Invariants|02. Entities, Business Rules, and Enterprise Invariants]]
├── [[Use Cases (interactors), Application Boundaries, and Request Response Models|03. Use Cases and Interactors Application Business Rules]]
├── [[Hexagonal Architecture (ports & Adapters) by Alistair Cockburn|04. Hexagonal Architecture Ports and Adapters]]
├── [[Onion Architecture (jeffrey Palermo), Domain Core, and Infrastructure Outer Layers|05. Onion Architecture and Domain Core Purity]]
├── [[Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels|06. Interface Adapters, Presenters, and Gateways]]
├── [[Component Cohesion Principles: Reuse Release (rep), Common Closure (ccp), Common Reuse (crp)|07. Component Cohesion Principles Rep, Ccp, Crp]]
├── [[Component Coupling Principles: Acyclic Dependencies (adp), Stable Dependencies (sdp), Stable Abstractions (sap)|08. Component Coupling Principles Adp, Sdp, Sap]]
├── [[Architectural Boundary Crossing, Dynamic Typing vs Static Dtos, and Serialization|09. Boundaries Crossing and Data Transfer Objects Dtos]]
├── [[Testing Strategies in Clean Architecture: Unit, Interactor, and Sociable Tests|10. Testing Strategies in Clean and Hexagonal Architectures]]
├── [[Framework Independence: Avoiding Framework Poisoning in Business Logic|11. Clean Architecture Framework Independence and Pragmatism]]
└── [[Refactoring Legacy Spaghetti Code to Clean Hexagonal Architecture|12. Refactoring Monoliths to Hexagonal Clean Architectures]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[The Dependency Rule, Concentric Circles, and Inversion of Control|01. The Dependency Rule and Inversion of Control]] — Source code dependencies must point only inward toward higher-level policies, decoupling business rules from UI, databases, and frameworks.
- 📂 [[Entities, Enterprise Business Rules, and Critical Invariants|02. Entities, Business Rules, and Enterprise Invariants]] — Encapsulating pure enterprise business logic, rich domain models vs anemic structures, and ensuring business invariants hold true independent of technology.
- 📂 [[Use Cases (interactors), Application Boundaries, and Request Response Models|03. Use Cases and Interactors Application Business Rules]] — Orchestrating domain entities, defining application-specific workflows, handling input/output boundary DTOs, and keeping use cases framework-free.
- 📂 [[Hexagonal Architecture (ports & Adapters) by Alistair Cockburn|04. Hexagonal Architecture Ports and Adapters]] — Driving (Primary) Ports and Adapters (HTTP, CLI) vs Driven (Secondary) Ports and Adapters (PostgreSQL, Kafka), mockability, and swappable infrastructure.
- 📂 [[Onion Architecture (jeffrey Palermo), Domain Core, and Infrastructure Outer Layers|05. Onion Architecture and Domain Core Purity]] — Domain Model at the center, Domain Services, Application Services, and Outer UI/Infrastructure ring, ensuring zero database dependencies in the domain core.
- 📂 [[Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels|06. Interface Adapters, Presenters, and Gateways]] — Converting data from use case format to web/UI format, implementing repository interfaces in data access gateways, and separating presentation logic.
- 📂 [[Component Cohesion Principles: Reuse Release (rep), Common Closure (ccp), Common Reuse (crp)|07. Component Cohesion Principles Rep, Ccp, Crp]] — Release-Reuse Equivalence Principle, Common Closure Principle (classes that change together belong together), and Common Reuse Principle.
- 📂 [[Component Coupling Principles: Acyclic Dependencies (adp), Stable Dependencies (sdp), Stable Abstractions (sap)|08. Component Coupling Principles Adp, Sdp, Sap]] — Acyclic dependency graphs, the Stable Dependencies Principle (depend in the direction of stability), and the Main Sequence (Abstractness vs Instability).
- 📂 [[Architectural Boundary Crossing, Dynamic Typing vs Static Dtos, and Serialization|09. Boundaries Crossing and Data Transfer Objects Dtos]] — Preventing domain entity leaks into web responses, strict boundary DTO mapping, and eliminating accidental coupling between database schema and API payloads.
- 📂 [[Testing Strategies in Clean Architecture: Unit, Interactor, and Sociable Tests|10. Testing Strategies in Clean and Hexagonal Architectures]] — Testing use cases in complete isolation using in-memory mock/fake adapters, testing driving adapters with contract tests, and achieving 100% test velocity.
- 📂 [[Framework Independence: Avoiding Framework Poisoning in Business Logic|11. Clean Architecture Framework Independence and Pragmatism]] — Treating frameworks as implementation details, protecting core business code from ORM/Web framework breaking upgrades, and pragmatic exceptions.
- 📂 [[Refactoring Legacy Spaghetti Code to Clean Hexagonal Architecture|12. Refactoring Monoliths to Hexagonal Clean Architectures]] — Identifying domain boundaries in legacy code, introducing ports around database calls, extracting use case interactors, and verifying with golden master tests.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

