---
title: Method Declarations & Receivers
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Methods & Interfaces]]"
---

# Method Declarations & Receivers

Methods vs functions, value receivers, pointer receivers, method sets, and methods on defined types.

```text
Method Declarations & Receivers
│
├── [[Methods vs Functions Architectural Differences]]
├── [[Pointer Receivers (State Mutation & Copy Elimination)]]
├── [[Value Receivers (Immutability & Concurrency Safety)]]
├── [[Receiver Type Consistency & Heuristics Decision Tree]]
├── [[Methods on Defined Non-Struct Types]]
├── [[Method Sets for T and Pointer-T]]
└── [[Cross-Package Method Definition Rules & Locality]]
```

---

## 🗂️ Topics

- [[Methods vs Functions Architectural Differences]] — Contrasting method receivers with first-class function parameters in API design.
- [[Pointer Receivers (State Mutation & Copy Elimination)]] — Mutating receiver state, eliminating large struct copy overhead, and consistency rules.
- [[Value Receivers (Immutability & Concurrency Safety)]] — Value copying semantics, concurrent read safety, and immutability guarantees.
- [[Receiver Type Consistency & Heuristics Decision Tree]] — Staff-level decision tree for choosing between value and pointer receivers.
- [[Methods on Defined Non-Struct Types]] — Attaching methods to custom primitives (type MyDuration int64, type StringSet map[string]struct{}).
- [[Method Sets for T and Pointer-T]] — Rules defining method sets for value T and pointer *T and their interface assignability.
- [[Cross-Package Method Definition Rules & Locality]] — Type locality constraints prohibiting method declarations on types from foreign packages.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]

