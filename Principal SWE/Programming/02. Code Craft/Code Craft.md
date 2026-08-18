---
title: Code Craft
tags:
  - programming
  - code-craft
  - principal-swe
parent: "[[Programming]]"
---

# 💻 Code Craft

Clean code fundamentals, readability, high cohesion, low coupling, and defensive programming contracts.

```text
Code Craft
│
├── [[Readability & Expressiveness|01. Readability & Expressiveness]]
│   ├── [[Intent-Revealing Code & Cognitive Load Minimization]]
│   ├── [[Naming as API Design & Semantic Clarity]]
│   ├── [[Comments as Rationale vs Code as Mechanism]]
│   └── [[Formatting, Uniformity & Zero-Tolerance Broken Windows]]
├── [[Modularization & Cohesion|02. Modularization & Cohesion]]
│   ├── [[High Cohesion, Low Coupling (Coupling Metrics)]]
│   ├── [[Single Responsibility Principle (SRP) at All Scales]]
│   ├── [[Encapsulation & Information Hiding Boundaries]]
│   └── [[Law of Demeter & Deep vs Shallow Modules]]
└── [[Defensive Programming & Contracts|03. Defensive Programming & Contracts]]
│   ├── [[Design by Contract (Preconditions, Postconditions, Invariants)]]
│   ├── [[Fail-Fast Principle & Unrecoverable State Handling]]
│   ├── [[Null Safety, Optionals & Sentinel Values]]
│   └── [[Idempotency & Pure Function Transformations]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Readability & Expressiveness|01. Readability & Expressiveness]]
- [[Intent-Revealing Code & Cognitive Load Minimization]] — Structuring logic so reading code feels like reading clear technical prose with zero mental stack overflow.
- [[Naming as API Design & Semantic Clarity]] — Naming variables, functions, and modules to convey intent, scope, and lifespan without ambiguity.
- [[Comments as Rationale vs Code as Mechanism]] — Writing comments that explain the WHY, constraints, and business intent rather than repeating what code does.
- [[Formatting, Uniformity & Zero-Tolerance Broken Windows]] — Enforcing strict automated formatting and style uniformity across teams to eliminate bike-shedding.
### 2. 📂 [[Modularization & Cohesion|02. Modularization & Cohesion]]
- [[High Cohesion, Low Coupling (Coupling Metrics)]] — Afferent vs efferent coupling, abstractness vs instability metrics, and the distance from main sequence.
- [[Single Responsibility Principle (SRP) at All Scales]] — Ensuring classes, functions, and modules have only one single actor or reason to change.
- [[Encapsulation & Information Hiding Boundaries]] — Hiding implementation volatility behind stable public interfaces and data abstraction.
- [[Law of Demeter & Deep vs Shallow Modules]] — Designing deep interfaces that offer simple interfaces to complex capabilities while preventing coupling.
### 3. 📂 [[Defensive Programming & Contracts|03. Defensive Programming & Contracts]]
- [[Design by Contract (Preconditions, Postconditions, Invariants)]] — Enforcing boundary assertions and state guarantees across subsystem interfaces.
- [[Fail-Fast Principle & Unrecoverable State Handling]] — Crashing or halting immediately upon encountering impossible internal states rather than corrupting data.
- [[Null Safety, Optionals & Sentinel Values]] — Eliminating the billion-dollar mistake using Option/Maybe types, nullable wrappers, and sentinel objects.
- [[Idempotency & Pure Function Transformations]] — Writing deterministic operations where multiple executions produce identical state transformations.

---

## 🔗 Navigation
- ⬆️ Parent: [[Programming]]
- 🎓 Root: [[Principal SWE]]
