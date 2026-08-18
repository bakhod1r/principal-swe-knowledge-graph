---
title: Queues (Basic Data Structures)
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Queues (Basic Data Structures)

FIFO buffer mechanics, circular arrays, modulo bitmasking, and lock-free ring buffers.

```text
Queues (Basic Data Structures)
│
├── [[Queue FIFO Invariants and State Transitions]]
├── [[Queue Enqueue (Tail Insertion)]]
├── [[Queue Dequeue (Head Removal)]]
├── [[Queue Peek (Front Element Inspection)]]
├── [[Queue Circular Array Buffer Implementation]]
├── [[Queue Power-of-Two Bitwise Masking Indexing]]
├── [[Queue Linked List Implementation]]
├── [[Queue Empty and Full Invariant Conditions]]
├── [[Queue Capacity Growth and Dynamic Resizing]]
├── [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]]
├── [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]]
└── [[Priority Queue vs FIFO Queue Tradeoffs]]
```

---

## 🗂️ Operations & Topics

- [[Queue FIFO Invariants and State Transitions]] — First-In First-Out operational axioms and queue state machine bounds.
- [[Queue Enqueue (Tail Insertion)]] — O(1) tail element insertion with capacity overflow validation.
- [[Queue Dequeue (Head Removal)]] — O(1) head element extraction and front pointer advancement.
- [[Queue Peek (Front Element Inspection)]] — O(1) viewing head element without mutating queue state.
- [[Queue Circular Array Buffer Implementation]] — Reusing fixed array memory via circular head/tail wraparound.
- [[Queue Power-of-Two Bitwise Masking Indexing]] — Replacing modulo division (% N) with ultra-fast bitwise AND masking (& (N - 1)).
- [[Queue Linked List Implementation]] — Heap-allocated node queue with head and tail pointers avoiding capacity limits.
- [[Queue Empty and Full Invariant Conditions]] — Distinguishing full vs empty states via element count or reserved index.
- [[Queue Capacity Growth and Dynamic Resizing]] — Unwrapping circular buffers into linear arrays during geometric doubling.
- [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]] — Single-producer single-consumer ring buffer with atomic memory barriers and cache line padding.
- [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]] — Lock-free concurrent queue using atomic compare-and-swap on head and tail sequences.
- [[Priority Queue vs FIFO Queue Tradeoffs]] — O(log N) priority-ordered heap queues vs O(1) temporal FIFO queues.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: [[Data Structures & Algorithms]]

