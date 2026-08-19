---
title: Defer
tags:
  - golang
  - functions
  - principal-swe
parent: "[[Functions (Clean Code)]]"
---

# Defer

LIFO deferred execution, argument evaluation, named returns mutation, and stack allocation.

```text
Defer
│
├── [[defer Statement Mechanics]]
├── [[Argument Evaluation at Defer Time]]
├── [[Modifying Named Return Values via Defer]]
├── [[Stack vs Heap Defer Allocation]]
└── [[defer in Loops Resource Leak Trap]]
```

---

## 🗂️ Topics

- [[defer Statement Mechanics]] — LIFO deferred function execution on function return or panic.
- [[Argument Evaluation at Defer Time]] — Immediate evaluation of defer arguments vs delayed execution of body.
- [[Modifying Named Return Values via Defer]] — Mutating named return variables inside deferred closures.
- [[Stack vs Heap Defer Allocation]] — Open-coded defers (Go 1.14+) eliminating heap allocation for small defers.
- [[defer in Loops Resource Leak Trap]] — Accumulating unexecuted defers inside long-running loops.
- [[Open-Coded Defers (Go 1.14+ Zero-Cost)]]
- [[Stack vs Heap Defer Allocations]]

---

## 🔗 References
- ⬆️ Parent: [[Functions (Clean Code)]]

