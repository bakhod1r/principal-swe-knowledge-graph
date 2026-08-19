---
title: "Hexagonal Architecture (ports & Adapters) by Alistair Cockburn Architectural Foundations and Invariants"
tags:
  - review
  - architecture
  - systems-architecture
  - clean,-hexagonal-and-onion-architecture
  - principal-swe
parent: "[[Hexagonal Architecture (ports & Adapters) by Alistair Cockburn]]"
---

# Hexagonal Architecture (ports & Adapters) by Alistair Cockburn Architectural Foundations and Invariants

## 1. Definition
**Hexagonal Architecture (ports & Adapters) by Alistair Cockburn Architectural Foundations and Invariants** represents a foundational architectural blueprint, structural invariant, and enterprise engineering standard within **Clean, Hexagonal & Onion Architecture**.
Driving (Primary) Ports and Adapters (HTTP, CLI) vs Driven (Secondary) Ports and Adapters (PostgreSQL, Kafka), mockability, and swappable infrastructure. Covering Core architectural principles, theoretical invariants, and structural rules.
It establishes rigorous system boundaries, decoupling mechanisms, high-availability guarantees, and structural integrity across large-scale software systems:
- **Architectural Invariants:** Enforces single responsibility at scale, strict boundary isolation, clear contract definitions, and verifiable resilience.
- **Enterprise Leverage:** Maximizes maintainability, eliminates brittle coupling, enables autonomous team delivery, and protects the system against catastrophic failure modes.

---

## 2. Mental Model
```text
Architectural Boundary & Invariant Flow for Hexagonal Architecture (ports & Adapters) by Alistair Cockburn Architectural Foundations and Invariants:
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
// Production Go architectural implementation and boundary pattern for Hexagonal Architecture (ports & Adapters) by Alistair Cockburn Architectural Foundations and Invariants
package main

import (
    "context"
    "fmt"
    "time"
)

type HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsPort interface {
    Execute(ctx context.Context, req HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsRequest) (*HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsResponse, error)
}

type HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsRequest struct {
    ID        string
    Timestamp time.Time
    Payload   map[string]any
}

type HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsResponse struct {
    Success bool
    Message string
}

type HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsService struct {
    adapter HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsPort
}

func NewHexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsService(adapter HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsPort) *HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsService {
    return &HexagonalArchitectureportsAdaptersbyAlistairCockburnArchitecturalFoundationsandInvariantsService{adapter: adapter}
}
```

---

## 4. Gotchas
- **Leaky Domain Abstractions:** Exposing internal database entities directly across API boundaries allows database schema changes to break external clients, creating tight cross-system coupling.
- **Unbounded Synchronous Cascading Calls:** Chaining multiple synchronous RPC calls across microservices causes latency accumulation, multiplies failure probabilities, and leads to system-wide distributed deadlocks.

---

## 🔗 References
- ⬆️ Parent: [[Hexagonal Architecture (ports & Adapters) by Alistair Cockburn]]
- 📚 Module: `Clean, Hexagonal & Onion Architecture`

