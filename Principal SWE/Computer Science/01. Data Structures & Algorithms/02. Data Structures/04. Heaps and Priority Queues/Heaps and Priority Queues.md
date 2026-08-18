---
title: Heaps and Priority Queues
tags:
  - computer-science
  - algorithms
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Heaps and Priority Queues

Binary heaps, D-Ary heaps, Fibonacci heaps, Pairing heaps, and Timer Wheels.

```text
Heaps and Priority Queues
│
├── [[Binary Heap Array Representation and Sift Operations]]
├── [[D-Ary Heap Cache Tuning in Production]]
├── [[Fibonacci Heap and O(1) Decrease Key]]
├── [[Pairing Heap Practical Performance]]
└── [[Hierarchical Timer Wheel Scheduling]]
```

---

## 🗂️ Topics

- [[Binary Heap Array Representation and Sift Operations]] — Implicit 1-indexed array representation (parent = i/2, children = 2i, 2i+1) and O(N) heapify.
- [[D-Ary Heap Cache Tuning in Production]] — Optimizing heap fan-out (D=4 or D=8) to match hardware CPU cache line transfers.
- [[Fibonacci Heap and O(1) Decrease Key]] — Lazy tree consolidation enabling theoretical O(1) amortized Decrease-Key for Dijkstra.
- [[Pairing Heap Practical Performance]] — Simpler self-adjusting heap variant outperforming Fibonacci heaps in real-world implementations.
- [[Hierarchical Timer Wheel Scheduling]] — O(1) insert and tick timer scheduler used in Linux kernel and Netty network engines.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 🎓 Root: [[Principal SWE]]
