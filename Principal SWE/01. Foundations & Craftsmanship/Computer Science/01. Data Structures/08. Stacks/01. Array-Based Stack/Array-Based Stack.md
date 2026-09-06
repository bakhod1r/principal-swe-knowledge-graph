---
title: Array-Based Stack
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Stacks]]"
---

# 📦 Array-Based Stack

Contiguous array stack with top index pointer.

```text
Array-Based Stack
│
├── [[Stack LIFO Invariants and State Transitions]]
├── [[Stack Push Operation (Element Insertion)]]
├── [[Stack Pop Operation (Element Removal)]]
├── [[Stack Peek Operation (Top Element Inspection)]]
├── [[Stack Overflow and Underflow Bounds Checking]]
├── [[Stack Frame Memory Allocation vs Heap Dynamic Arrays]]
└── [[Array-Based Stack Cache Locality and Vectorization]]
```

---

## 🗂️ Topics & Implementations

- [[Stack LIFO Invariants and State Transitions]] — Last-In First-Out operational axioms, push/pop state transitions, and top-of-stack pointer.
- [[Stack Push Operation (Element Insertion)]] — O(1) push operation at the top index with capacity reallocation check.
- [[Stack Pop Operation (Element Removal)]] — O(1) pop operation returning top element and decrementing top pointer.
- [[Stack Peek Operation (Top Element Inspection)]] — O(1) viewing top element without mutating stack state.
- [[Stack Overflow and Underflow Bounds Checking]] — Validating top pointer against zero and maximum capacity bounds.
- [[Stack Frame Memory Allocation vs Heap Dynamic Arrays]] — CPU hardware stack pointer (rsp) adjustment vs heap memory allocation latency.
- [[Array-Based Stack Cache Locality and Vectorization]] — Cache line benefits of contiguous stack top access in high-frequency operations.

---

## 🔗 References
- ⬆️ Parent: [[Stacks]]
- 📚 Module: `Data Structures`
