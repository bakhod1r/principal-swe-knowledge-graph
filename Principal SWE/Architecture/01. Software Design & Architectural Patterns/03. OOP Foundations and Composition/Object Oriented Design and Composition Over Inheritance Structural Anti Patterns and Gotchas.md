---
title: "Object Oriented Design and Composition Over Inheritance Structural Anti Patterns and Gotchas"
tags:
  - architecture
  - software-design
  - software-design-and-architectural-patterns
  - principal-swe
parent: "[[Object Oriented Design and Composition Over Inheritance]]"
---

# Object Oriented Design and Composition Over Inheritance Structural Anti Patterns and Gotchas

## 1. Definition
**Object Oriented Design and Composition Over Inheritance Structural Anti Patterns and Gotchas** represents a fundamental architectural discipline, structural pattern, and engineering standard within **Software Design & Architectural Patterns**.
Encapsulation, Polymorphism, Abstract Data Types, Liskov Substitution, and avoiding brittle inheritance hierarchies. Covering Critical architectural anti-patterns, coupling traps, and mitigation checklists.
It establishes rigorous engineering guarantees on software maintainability, loose coupling, and long-term evolutionary capability:
- **Structural Invariants:** Enforces clear separation of concerns, explicit domain boundaries, dependency inversion, and bounded contexts.
- **Architectural Leverage:** Maximizes testability, team autonomy, and technical agility while minimizing total cost of ownership (TCO).

---

## 2. Mental Model
```text
Architectural Structure & Boundary Flow for Object Oriented Design and Composition Over Inheritance Structural Anti Patterns and Gotchas:
[ External UI / API Client ] ───> [ Primary Adapters / Controllers ]
                                                 │
                 ┌───────────────────────────────┴───────────────────────────────┐
                 ▼                                                               ▼
   [ Application Use Cases / Orchestration ]                       [ Domain Core (Entities & Invariants) ]
                 │                                                               │
                 └───────────────────────────────┬───────────────────────────────┘
                                                 ▼
                             [ Secondary Adapters / Infrastructure / DB ]
```
- **Architectural Rule:** Dependencies strictly point inward toward the Domain Core. High-level business policies never depend on low-level infrastructure details.

---

## 3. Usage
```go
// Production Go architectural implementation and boundary pattern for Object Oriented Design and Composition Over Inheritance Structural Anti Patterns and Gotchas
package main

import (
    "context"
    "fmt"
)

// Domain Port (Inward Interface)
type ObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasRepository interface {
    FindByID(ctx context.Context, id string) (*DomainModel, error)
    Save(ctx context.Context, entity *DomainModel) error
}

type DomainModel struct {
    ID    string
    State string
}

func (d *DomainModel) ValidateInvariant() error {
    if d.ID == "" {
        return fmt.Errorf("invariant violated: ID cannot be empty")
    }
    return nil
}

// Application Service (Use Case Interactor)
type ObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasUseCase struct {
    repo ObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasRepository
}

func NewObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasUseCase(repo ObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasRepository) *ObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasUseCase {
    return &ObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasUseCase{repo: repo}
}

func (u *ObjectOrientedDesignandCompositionOverInheritanceStructuralAntiPatternsandGotchasUseCase) Execute(ctx context.Context, id string) error {
    entity, err := u.repo.FindByID(ctx, id)
    if err != nil {
        return err
    }
    return entity.ValidateInvariant()
}
```

---

## 4. Gotchas
- **Leaky Abstractions Across Boundaries:** Exposing persistence entities or ORM annotations directly to API layers couples internal database schemas with external clients, breaking boundary isolation.
- **Anemic Domain Model:** Moving business logic into bloated service classes while leaving domain entities as dumb data bags destroys object encapsulation and invariant enforcement.

---

## 🔗 References
- ⬆️ Parent: [[Object Oriented Design and Composition Over Inheritance]]
- 📚 Module: [[Software Design & Architectural Patterns]]
- 🎓 Root: [[Principal SWE]]
