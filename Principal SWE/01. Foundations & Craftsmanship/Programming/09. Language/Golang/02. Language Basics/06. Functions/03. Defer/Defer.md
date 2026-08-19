---
title: Defer
tags:
  - golang
  - functions
  - defer
  - principal-swe
parent: "[[Functions]]"
---

# Defer

LIFO deferred execution, argument evaluation, named returns mutation, open-coded defers, and stack allocation.

```text
Defer
│
├── [[defer Statement Mechanics]]
├── [[Argument Evaluation at Defer Time]]
├── [[Modifying Named Return Values via Defer]]
├── [[Stack vs Heap Defer Allocation]]
├── [[Open-Coded Defers (Go 1.14+ Zero-Cost)]]
└── [[defer in Loops Resource Leak Trap]]
```

---

## 🗂️ Topics

- [[defer Statement Mechanics]] — LIFO deferred function execution on function return or panic.
- [[Argument Evaluation at Defer Time]] — Immediate evaluation of defer arguments vs delayed execution of body.
- [[Modifying Named Return Values via Defer]] — Mutating named return variables inside deferred closures.
- [[Stack vs Heap Defer Allocation]] — Stack-allocated `_defer` structs vs heap-allocated defer records.
- [[Open-Coded Defers (Go 1.14+ Zero-Cost)]] — Open-coded defers (Go 1.14+) inlining defer calls directly into exit points (zero overhead).
- [[defer in Loops Resource Leak Trap]] — Accumulating unexecuted defers inside long-running loops and wrapping in worker functions.

---

## 🔗 References
- ⬆️ Parent: [[Functions]]
- 📚 Module: `Language Basics`
