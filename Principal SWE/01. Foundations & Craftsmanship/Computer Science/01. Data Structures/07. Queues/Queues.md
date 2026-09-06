---
title: Queues
tags:
  - review
  - computer-science
  - data-structures
  - queues
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Queues (FIFO Buffers & Concurrency Topologies)

Comprehensive engineering catalog of First-In First-Out buffering mechanisms, fixed circular ring arrays, power-of-two bitmasking, priority queues, and lock-free concurrent queues across 4 specialized domains.

```text
Queues
│
├── 01. Linear FIFO Queue
│   ├── [[Linear FIFO Queue]]
│   ├── [[Queue FIFO Invariants and State Transitions]]
│   ├── [[Queue Empty and Full Invariant Conditions]]
│   ├── [[Queue Enqueue (Tail Insertion)]]
│   ├── [[Queue Dequeue (Head Removal)]]
│   ├── [[Queue Peek (Front Element Inspection)]]
│   └── [[Queue Linked List Implementation]]
│
├── 02. Circular Queue
│   ├── [[Circular Queue]]
│   ├── [[Queue Circular Array Buffer Implementation]]
│   ├── [[Queue Power-of-Two Bitwise Masking Indexing]]
│   └── [[Queue Capacity Growth and Dynamic Resizing]]
│
├── 03. Priority Queue
│   ├── [[Priority Queue]]
│   └── [[Priority Queue vs FIFO Queue Tradeoffs]]
│
└── 04. Concurrent & Lock-Free Queue
    ├── [[Concurrent & Lock-Free Queue]]
    ├── [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]]
    └── [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Linear FIFO Queue|01. Linear FIFO Queue]]
- [[Linear FIFO Queue]] — Master topology of linear sequential FIFO queues.
- [[Queue FIFO Invariants and State Transitions]] — First-In First-Out operational axioms and queue state machine bounds.
- [[Queue Empty and Full Invariant Conditions]] — Distinguishing full vs empty states via element count or reserved index.
- [[Queue Enqueue (Tail Insertion)]] — O(1) tail element insertion with capacity overflow validation.
- [[Queue Dequeue (Head Removal)]] — O(1) head element extraction and front pointer advancement.
- [[Queue Peek (Front Element Inspection)]] — O(1) viewing head element without mutating queue state.
- [[Queue Linked List Implementation]] — Heap-allocated node queue with head and tail pointers avoiding capacity limits.

### 2. 📂 [[Circular Queue|02. Circular Queue]]
- [[Circular Queue]] — Master topology of bounded circular ring buffers.
- [[Queue Circular Array Buffer Implementation]] — Reusing fixed array memory via circular head/tail wraparound.
- [[Queue Power-of-Two Bitwise Masking Indexing]] — Replacing modulo division (% N) with ultra-fast bitwise AND masking (& (N - 1)).
- [[Queue Capacity Growth and Dynamic Resizing]] — Unwrapping circular buffers into linear arrays during geometric doubling.

### 3. 📂 [[Priority Queue|03. Priority Queue]]
- [[Priority Queue]] — Master topology of priority-ordered queues.
- [[Priority Queue vs FIFO Queue Tradeoffs]] — O(log N) priority-ordered heap queues vs O(1) temporal FIFO queues.

### 4. 📂 [[Concurrent & Lock-Free Queue|04. Concurrent & Lock-Free Queue]]
- [[Concurrent & Lock-Free Queue]] — Master topology of concurrent thread-safe queues.
- [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]] — Single-producer single-consumer ring buffer with atomic memory barriers and cache line padding.
- [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]] — Lock-free concurrent queue using atomic compare-and-swap on head and tail sequences.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`



