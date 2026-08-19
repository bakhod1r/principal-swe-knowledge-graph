---
title: Architectural Anti-Patterns
tags:
  - architecture
  - systems-architecture
  - architectural-anti-patterns-and-technical-debt-refactoring
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Architectural Anti Patterns & Technical Debt Refactoring

System shape anti-patterns and legacy refactoring: Distributed Monolith, God Service, Chatty I/O sprawl, Shared Database anti-pattern, Anemic Domain Model, Golden Hammer, Accidental Complexity, Spaghetti Rot, and Strangler Fig refactoring.

```text
Architectural Anti Patterns & Technical Debt Refactoring
│
├── [[The Distributed Monolith: Tightly Coupled Microservices and Lock Step Deployments|01. The Distributed Monolith Anti Pattern]]
├── [[The God Service (megaservice) Pitfall: Giant Services with Too Many Responsibilities|02. God Service and Megaservice Architecture Pitfall]]
├── [[Chatty I O Anti Pattern, Microservice Sprawl, and Deep Dependency Chains|03. Chatty I O and Microservice Sprawl Pitfall]]
├── [[The Shared Database Anti Pattern: Microservices Bypassing API Contracts|04. Shared Database Anti Pattern and Database Coupling]]
├── [[Anemic Domain Model: Separation of Data and Behavior in Procedural Code|05. Anemic Domain Model Anti Pattern and Procedural Leaks]]
├── [[The Golden Hammer Anti Pattern, Resume Driven Development, and Trend Chasing|06. The Golden Hammer and Trend Driven Architecture]]
├── `07. Accidental vs Essential Complexity and Over Engineering`
├── `08. Big Ball of Mud and Spaghetti Code Rot`
├── `09. Vendor Lock In, Cloud Sprawl, and Leaky Abstractions`
└── `10. Systemic Technical Debt Quantification and Refactoring Roadmaps`
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[The Distributed Monolith: Tightly Coupled Microservices and Lock Step Deployments|01. The Distributed Monolith Anti Pattern]] — Identifying distributed monolith symptoms (simultaneous service deployments, shared codebases, cascading synchronous HTTP calls), and refactoring.
- 📂 [[The God Service (megaservice) Pitfall: Giant Services with Too Many Responsibilities|02. God Service and Megaservice Architecture Pitfall]] — Decomposing oversized services that hold 80% of business logic into bounded subdomains, resolving high contention, and restoring team velocity.
- 📂 [[Chatty I O Anti Pattern, Microservice Sprawl, and Deep Dependency Chains|03. Chatty I O and Microservice Sprawl Pitfall]] — Diagnosing high-latency request waterfalls caused by excessive inter-service network hops, aggregating requests with BFF, and batch APIs.
- 📂 [[The Shared Database Anti Pattern: Microservices Bypassing API Contracts|04. Shared Database Anti Pattern and Database Coupling]] — Why sharing database tables across multiple microservices causes schema lock-in and silent data corruption, and transitioning to API-driven access.
- 📂 [[Anemic Domain Model: Separation of Data and Behavior in Procedural Code|05. Anemic Domain Model Anti Pattern and Procedural Leaks]] — Recognizing anemic domain classes (getters/setters only) with business logic scattered in procedural services, and refactoring to Rich Domain Models.
- 📂 [[The Golden Hammer Anti Pattern, Resume Driven Development, and Trend Chasing|06. The Golden Hammer and Trend Driven Architecture]] — Forcing a single technology (Kafka, Kubernetes, GraphQL) onto every problem regardless of fit, evaluating total cost of ownership, and right-tool-for-the-job.
- 📂 `07. Accidental vs Essential Complexity and Over Engineering` — Distinguishing between core business complexity and self-inflicted architectural bloat, applying YAGNI and KISS, and pruning unused abstractions.
- 📂 `08. Big Ball of Mud and Spaghetti Code Rot` — Why systems degenerate into unstructured spaghetti code, defining modular boundaries, creating dependency matrices, and incremental modularization.
- 📂 `09. Vendor Lock In, Cloud Sprawl, and Leaky Abstractions` — Evaluating proprietary cloud service dependencies against open standards, portability trade-offs, and wrapping vendor SDKs in clean domain ports.
- 📂 `10. Systemic Technical Debt Quantification and Refactoring Roadmaps` — Measuring technical debt interest rate, calculating engineering drag, presenting technical debt refactoring ROI to executives, and phasing migrations.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

