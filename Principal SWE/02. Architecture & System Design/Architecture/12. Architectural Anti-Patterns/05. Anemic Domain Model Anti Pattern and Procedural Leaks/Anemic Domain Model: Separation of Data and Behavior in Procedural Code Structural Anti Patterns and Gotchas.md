---
title: "Anemic Domain Model: Separation of Data and Behavior in Procedural Code Structural Anti Patterns and Gotchas"
tags:
  - review
  - architecture
  - systems-architecture
  - architectural-anti-patterns-and-technical-debt-refactoring
  - principal-swe
parent: "[[Anemic Domain Model: Separation of Data and Behavior in Procedural Code]]"
---

# Anemic Domain Model: Separation of Data and Behavior in Procedural Code Structural Anti Patterns and Gotchas

## 1. Definition
**Anemic Domain Model: Separation of Data and Behavior in Procedural Code Structural Anti Patterns and Gotchas** represents a foundational architectural blueprint, structural invariant, and enterprise engineering standard within **Architectural Anti-Patterns & Technical Debt Refactoring**.
Recognizing anemic domain classes (getters/setters only) with business logic scattered in procedural services, and refactoring to Rich Domain Models. Covering Critical architectural anti-patterns, failure modes, trade-offs, and refactoring strategies.
It establishes rigorous system boundaries, decoupling mechanisms, high-availability guarantees, and structural integrity across large-scale software systems:
- **Architectural Invariants:** Enforces single responsibility at scale, strict boundary isolation, clear contract definitions, and verifiable resilience.
- **Enterprise Leverage:** Maximizes maintainability, eliminates brittle coupling, enables autonomous team delivery, and protects the system against catastrophic failure modes.

---

## 2. Mental Model
```text
Architectural Boundary & Invariant Flow for Anemic Domain Model: Separation of Data and Behavior in Procedural Code Structural Anti Patterns and Gotchas:
[ Inbound Consumer / External Client ] ───> [ Strict Boundary Adapter / API Gateway ]
                                                              │
                    ┌─────────────────────────────────────────┴─────────────────────────────────────────┐
                    ▼                                                                                   ▼
     [ Core Domain Logic / Business Policy ]                                             [ Asynchronous Event / Integration Outbox ]
                    │                                                                                   │
                    └─────────────────────────────────────────┬─────────────────────────────────────────┘
                                                              ▼
                                  [ Isolated Persistent Storage / External Enterprise Service ]
```
- **Architectural Law:** The cost of changing a software boundary increases by an order of magnitude at each subsequent phase of development. Design boundaries deliberately.

---

## 3. Usage
```go
// Production Go architectural implementation and boundary pattern for Anemic Domain Model: Separation of Data and Behavior in Procedural Code Structural Anti Patterns and Gotchas
package main

import (
    "context"
    "fmt"
    "time"
)

type AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasPort interface {
    Execute(ctx context.Context, req AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasRequest) (*AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasResponse, error)
}

type AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasRequest struct {
    ID        string
    Timestamp time.Time
    Payload   map[string]any
}

type AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasResponse struct {
    Success bool
    Message string
}

type AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasService struct {
    adapter AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasPort
}

func NewAnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasService(adapter AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasPort) *AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasService {
    return &AnemicDomainModelSeparationofDataandBehaviorinProceduralCodeStructuralAntiPatternsandGotchasService{adapter: adapter}
}
```

---

## 4. Gotchas
- **Leaky Domain Abstractions:** Exposing internal database entities directly across API boundaries allows database schema changes to break external clients, creating tight cross-system coupling.
- **Unbounded Synchronous Cascading Calls:** Chaining multiple synchronous RPC calls across microservices causes latency accumulation, multiplies failure probabilities, and leads to system-wide distributed deadlocks.

---

## 🔗 References
- ⬆️ Parent: [[Anemic Domain Model: Separation of Data and Behavior in Procedural Code]]
- 📚 Module: `Architectural Anti Patterns & Technical Debt Refactoring`

