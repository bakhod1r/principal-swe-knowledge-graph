---
title: "Thoughtworks Technology Radars: Managing Tech Sprawl (adopt, Trial, Assess, Hold) Structural Anti Patterns and Gotchas"
tags:
  - architecture
  - systems-architecture
  - software-architect-leadership-and-governance
  - principal-swe
parent: "[[Thoughtworks Technology Radars: Managing Tech Sprawl (adopt, Trial, Assess, Hold)]]"
---

# Thoughtworks Technology Radars: Managing Tech Sprawl (adopt, Trial, Assess, Hold) Structural Anti Patterns and Gotchas

## 1. Definition
**Thoughtworks Technology Radars: Managing Tech Sprawl (adopt, Trial, Assess, Hold) Structural Anti Patterns and Gotchas** represents a foundational architectural blueprint, structural invariant, and enterprise engineering standard within **Software Architect Leadership & Governance**.
Building company-specific tech radars, standardizing language/framework selections, establishing 'paved roads' for developer golden paths, and deprecation. Covering Critical architectural anti-patterns, failure modes, trade-offs, and refactoring strategies.
It establishes rigorous system boundaries, decoupling mechanisms, high-availability guarantees, and structural integrity across large-scale software systems:
- **Architectural Invariants:** Enforces single responsibility at scale, strict boundary isolation, clear contract definitions, and verifiable resilience.
- **Enterprise Leverage:** Maximizes maintainability, eliminates brittle coupling, enables autonomous team delivery, and protects the system against catastrophic failure modes.

---

## 2. Mental Model
```text
Architectural Boundary & Invariant Flow for Thoughtworks Technology Radars: Managing Tech Sprawl (adopt, Trial, Assess, Hold) Structural Anti Patterns and Gotchas:
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
// Production Go architectural implementation and boundary pattern for Thoughtworks Technology Radars: Managing Tech Sprawl (adopt, Trial, Assess, Hold) Structural Anti Patterns and Gotchas
package main

import (
    "context"
    "fmt"
    "time"
)

type ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasPort interface {
    Execute(ctx context.Context, req ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasRequest) (*ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasResponse, error)
}

type ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasRequest struct {
    ID        string
    Timestamp time.Time
    Payload   map[string]any
}

type ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasResponse struct {
    Success bool
    Message string
}

type ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasService struct {
    adapter ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasPort
}

func NewThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasService(adapter ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasPort) *ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasService {
    return &ThoughtworksTechnologyRadarsManagingTechSprawladoptTrialAssessHoldStructuralAntiPatternsandGotchasService{adapter: adapter}
}
```

---

## 4. Gotchas
- **Leaky Domain Abstractions:** Exposing internal database entities directly across API boundaries allows database schema changes to break external clients, creating tight cross-system coupling.
- **Unbounded Synchronous Cascading Calls:** Chaining multiple synchronous RPC calls across microservices causes latency accumulation, multiplies failure probabilities, and leads to system-wide distributed deadlocks.

---

## 🔗 References
- ⬆️ Parent: [[Thoughtworks Technology Radars: Managing Tech Sprawl (adopt, Trial, Assess, Hold)]]
- 📚 Module: `Software Architect Leadership & Governance`

