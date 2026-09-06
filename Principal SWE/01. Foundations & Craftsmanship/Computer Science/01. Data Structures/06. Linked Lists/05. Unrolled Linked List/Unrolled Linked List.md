---
title: Unrolled Linked List
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Linked Lists]]"
---

# 📦 Unrolled Linked List

Hybrid cache-conscious structure packing short contiguous arrays inside each linked list node.

```text
Unrolled Linked List
│
├── [[Unrolled Linked List Hybrid Cache Topology]]
├── [[Unrolled Node Capacity and Cache Line Alignment (64-Byte Matching)]]
├── [[Unrolled List Element Insertion and Node Splitting Mechanics]]
└── [[Defragmentation and Node Merging in Unrolled Lists]]
```

---

## 🗂️ Topics & Implementations

- [[Unrolled Linked List Hybrid Cache Topology]] — Packing flat array chunks into nodes to match 64-byte CPU cache lines.
- [[Unrolled Node Capacity and Cache Line Alignment (64-Byte Matching)]] — Sizing node array buffers to perfectly fit L1 cache lines and SIMD lanes.
- [[Unrolled List Element Insertion and Node Splitting Mechanics]] — Maintaining balance through node splitting when chunk capacity exceeds threshold.
- [[Defragmentation and Node Merging in Unrolled Lists]] — Underflow handling: coalescing under-filled adjacent nodes to save memory.

---

## 🔗 References
- ⬆️ Parent: [[Linked Lists]]
- 📚 Module: `Data Structures`
