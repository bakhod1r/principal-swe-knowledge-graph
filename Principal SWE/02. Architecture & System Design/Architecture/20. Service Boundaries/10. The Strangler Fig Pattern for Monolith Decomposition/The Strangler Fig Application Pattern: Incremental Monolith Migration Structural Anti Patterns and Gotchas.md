---
title: "The Strangler Fig Application Pattern: Incremental Monolith Migration Structural Anti Patterns and Gotchas"
tags:
  - architecture
  - systems-architecture
  - microservice-architecture-and-service-boundaries
  - principal-swe
parent: "[[The Strangler Fig Application Pattern: Incremental Monolith Migration]]"
---

# The Strangler Fig Application Pattern: Incremental Monolith Migration Structural Anti Patterns and Gotchas

## 1. Definition
**The Strangler Fig Application Pattern: Incremental Monolith Migration Structural Anti Patterns and Gotchas** represents a foundational architectural blueprint, structural invariant, and enterprise engineering standard within **Microservice Architecture & Service Boundaries**.
Intercepting inbound traffic at the edge proxy/API gateway, routing legacy endpoints to the monolith while routing new features to microservices until sunset. Covering Critical architectural anti-patterns, failure modes, trade-offs, and refactoring strategies.
It establishes rigorous system boundaries, decoupling mechanisms, high-availability guarantees, and structural integrity across large-scale software systems:
- **Architectural Invariants:** Enforces single responsibility at scale, strict boundary isolation, clear contract definitions, and verifiable resilience.
- **Enterprise Leverage:** Maximizes maintainability, eliminates brittle coupling, enables autonomous team delivery, and protects the system against catastrophic failure modes.

---

## 2. Mental Model
```text
Architectural Boundary & Invariant Flow for The Strangler Fig Application Pattern: Incremental Monolith Migration Structural Anti Patterns and Gotchas:
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
// Production Go architectural implementation and boundary pattern for The Strangler Fig Application Pattern: Incremental Monolith Migration Structural Anti Patterns and Gotchas
package main

import (
    "context"
    "fmt"
    "time"
)

type TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasPort interface {
    Execute(ctx context.Context, req TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasRequest) (*TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasResponse, error)
}

type TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasRequest struct {
    ID        string
    Timestamp time.Time
    Payload   map[string]any
}

type TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasResponse struct {
    Success bool
    Message string
}

type TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasService struct {
    adapter TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasPort
}

func NewTheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasService(adapter TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasPort) *TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasService {
    return &TheStranglerFigApplicationPatternIncrementalMonolithMigrationStructuralAntiPatternsandGotchasService{adapter: adapter}
}
```

---

## 4. Gotchas
- **Leaky Domain Abstractions:** Exposing internal database entities directly across API boundaries allows database schema changes to break external clients, creating tight cross-system coupling.
- **Unbounded Synchronous Cascading Calls:** Chaining multiple synchronous RPC calls across microservices causes latency accumulation, multiplies failure probabilities, and leads to system-wide distributed deadlocks.

---

## 🔗 References
- ⬆️ Parent: [[The Strangler Fig Application Pattern: Incremental Monolith Migration]]
- 📚 Module: `Microservice Architecture & Service Boundaries`

