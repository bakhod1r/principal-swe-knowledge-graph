---
title: "Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels Structural Anti Patterns and Gotchas"
tags:
  - architecture
  - systems-architecture
  - clean,-hexagonal-and-onion-architecture
  - principal-swe
parent: "[[Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels]]"
---

# Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels Structural Anti Patterns and Gotchas

## 1. Definition
**Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels Structural Anti Patterns and Gotchas** represents a foundational architectural blueprint, structural invariant, and enterprise engineering standard within **Clean, Hexagonal & Onion Architecture**.
Converting data from use case format to web/UI format, implementing repository interfaces in data access gateways, and separating presentation logic. Covering Critical architectural anti-patterns, failure modes, trade-offs, and refactoring strategies.
It establishes rigorous system boundaries, decoupling mechanisms, high-availability guarantees, and structural integrity across large-scale software systems:
- **Architectural Invariants:** Enforces single responsibility at scale, strict boundary isolation, clear contract definitions, and verifiable resilience.
- **Enterprise Leverage:** Maximizes maintainability, eliminates brittle coupling, enables autonomous team delivery, and protects the system against catastrophic failure modes.

---

## 2. Mental Model
```text
Architectural Boundary & Invariant Flow for Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels Structural Anti Patterns and Gotchas:
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
// Production Go architectural implementation and boundary pattern for Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels Structural Anti Patterns and Gotchas
package main

import (
    "context"
    "fmt"
    "time"
)

type InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasPort interface {
    Execute(ctx context.Context, req InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasRequest) (*InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasResponse, error)
}

type InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasRequest struct {
    ID        string
    Timestamp time.Time
    Payload   map[string]any
}

type InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasResponse struct {
    Success bool
    Message string
}

type InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasService struct {
    adapter InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasPort
}

func NewInterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasService(adapter InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasPort) *InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasService {
    return &InterfaceAdaptersControllersPresentersGatewaysandViewmodelsStructuralAntiPatternsandGotchasService{adapter: adapter}
}
```

---

## 4. Gotchas
- **Leaky Domain Abstractions:** Exposing internal database entities directly across API boundaries allows database schema changes to break external clients, creating tight cross-system coupling.
- **Unbounded Synchronous Cascading Calls:** Chaining multiple synchronous RPC calls across microservices causes latency accumulation, multiplies failure probabilities, and leads to system-wide distributed deadlocks.

---

## 🔗 References
- ⬆️ Parent: [[Interface Adapters: Controllers, Presenters, Gateways, and Viewmodels]]
- 📚 Module: `Clean, Hexagonal & Onion Architecture`

