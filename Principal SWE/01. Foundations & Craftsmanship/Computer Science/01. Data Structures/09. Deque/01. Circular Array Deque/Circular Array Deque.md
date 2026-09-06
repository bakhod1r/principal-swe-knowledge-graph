---
title: Circular Array Deque
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Deque]]"
---

# 📦 Circular Array Deque

Circular buffer allowing O(1) push and pop at both ends.

```text
Circular Array Deque
│
├── [[Deque Invariants and Double-Ended Contract]]
├── [[Deque Push Front and Pop Front Operations]]
├── [[Deque Push Back and Pop Back Operations]]
├── [[Deque Circular Array Buffer Implementation]]
├── [[Circular Deque Power-of-Two Bitwise Indexing for Head and Tail]]
└── [[Circular Deque Dynamic Geometric Reallocation and Unwrapping]]
```

---

## 🗂️ Topics & Implementations

- [[Deque Invariants and Double-Ended Contract]] — Formal invariant: O(1) insertions and deletions at both head and tail.
- [[Deque Push Front and Pop Front Operations]] — O(1) head manipulations with circular index decrement and wraparound.
- [[Deque Push Back and Pop Back Operations]] — O(1) tail manipulations with circular index increment and wraparound.
- [[Deque Circular Array Buffer Implementation]] — Bidirectional modulo arithmetic on contiguous arrays.
- [[Circular Deque Power-of-Two Bitwise Indexing for Head and Tail]] — Masking head and tail pointers with (capacity - 1) for zero-cost wraparound.
- [[Circular Deque Dynamic Geometric Reallocation and Unwrapping]] — Reordering wrapped head-tail elements into linear order during array doubling.

---

## 🔗 References
- ⬆️ Parent: [[Deque]]
- 📚 Module: `Data Structures`
