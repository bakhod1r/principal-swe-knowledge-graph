---
title: Queues
tags:
  - review
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Queues

First-In First-Out buffering mechanisms, fixed circular ring arrays, power-of-two bitmasking, priority queues, and lock-free concurrent queues.

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
│   ├── [[Queue Linked List Implementation]]
│   ├── [[Queue Underflow and Overflow Exceptions vs Non-Blocking Options]]
│   └── [[Two-Stack FIFO Queue Implementation and Amortized O(1) Proof]]
││
├── 02. Circular Queue
│   ├── [[Circular Queue]]
│   ├── [[Queue Circular Array Buffer Implementation]]
│   ├── [[Queue Power-of-Two Bitwise Masking Indexing]]
│   ├── [[Queue Capacity Growth and Dynamic Resizing]]
│   ├── [[Circular Queue Monotonic Sequence Number Indexing]]
│   └── [[Bounded Queue Backpressure and Drop Policies (Drop-Head vs Drop-Tail)]]
││
├── 03. Priority Queue
│   ├── [[Priority Queue]]
│   ├── [[Priority Queue vs FIFO Queue Tradeoffs]]
│   ├── [[Binary Heap Array Representation for Priority Queue]]
│   ├── [[Indexed Priority Queue and O(log N) Decrease-Key Operation]]
│   └── [[D-ary Heap (4-ary) Priority Queue for Cache Optimization]]
││
└── 04. Concurrent & Lock-Free Queue
    ├── [[Concurrent & Lock-Free Queue]]
    ├── [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]]
    ├── [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]]
    ├── [[Michael-Scott Lock-Free FIFO Queue (CAS on Head and Tail)]]
    ├── [[The ABA Problem and Hazard Pointers in Lock-Free Queues]]
    └── [[Bounded Blocking Queue (Condition Variables and Mutex Striping)]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Linear FIFO Queue|01. Linear FIFO Queue]]
- [[Linear FIFO Queue]] — Master topology: Master topology of linear sequential FIFO queues.
- [[Queue FIFO Invariants and State Transitions]] — First-In First-Out operational axioms and queue state machine bounds.
- [[Queue Empty and Full Invariant Conditions]] — Distinguishing full vs empty states via element count or reserved index.
- [[Queue Enqueue (Tail Insertion)]] — O(1) tail element insertion with capacity overflow validation.
- [[Queue Dequeue (Head Removal)]] — O(1) head element extraction and front pointer advancement.
- [[Queue Peek (Front Element Inspection)]] — O(1) viewing head element without mutating queue state.
- [[Queue Linked List Implementation]] — Heap-allocated node queue with head and tail pointers avoiding capacity limits.
- [[Queue Underflow and Overflow Exceptions vs Non-Blocking Options]] — Handling bounded queue states: throwing exceptions, blocking, or dropping.
- [[Two-Stack FIFO Queue Implementation and Amortized O(1) Proof]] — Implementing FIFO queues using two LIFO stacks with amortized O(1) operations.

### 2. 📂 [[Circular Queue|02. Circular Queue]]
- [[Circular Queue]] — Master topology: Master topology of bounded circular ring buffers.
- [[Queue Circular Array Buffer Implementation]] — Reusing fixed array memory via circular head/tail wraparound.
- [[Queue Power-of-Two Bitwise Masking Indexing]] — Replacing modulo division (% N) with ultra-fast bitwise AND masking (& (N - 1)).
- [[Queue Capacity Growth and Dynamic Resizing]] — Unwrapping circular buffers into linear arrays during geometric doubling.
- [[Circular Queue Monotonic Sequence Number Indexing]] — 64-bit monotonic sequences preventing counter overflow and index wrapping errors.
- [[Bounded Queue Backpressure and Drop Policies (Drop-Head vs Drop-Tail)]] — Shedding load in real-time pipelines: evicting oldest data vs rejecting newest.

### 3. 📂 [[Priority Queue|03. Priority Queue]]
- [[Priority Queue]] — Master topology: Master topology of priority-ordered queues.
- [[Priority Queue vs FIFO Queue Tradeoffs]] — O(log N) priority-ordered heap queues vs O(1) temporal FIFO queues.
- [[Binary Heap Array Representation for Priority Queue]] — Implicit tree representation in a contiguous array with 2i+1 and 2i+2 indices.
- [[Indexed Priority Queue and O(log N) Decrease-Key Operation]] — Tracking value-to-index mapping for instant priority updates in Dijkstra's algorithm.
- [[D-ary Heap (4-ary) Priority Queue for Cache Optimization]] — Wider branching factors matching cache lines and reducing tree height.

### 4. 📂 [[Concurrent & Lock-Free Queue|04. Concurrent & Lock-Free Queue]]
- [[Concurrent & Lock-Free Queue]] — Master topology: Master topology of concurrent thread-safe queues.
- [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]] — Single-producer single-consumer ring buffer with atomic memory barriers and cache line padding.
- [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]] — Lock-free concurrent queue using atomic compare-and-swap on head and tail sequences.
- [[Michael-Scott Lock-Free FIFO Queue (CAS on Head and Tail)]] — Standard concurrent linked queue using atomic pointer swings on node chains.
- [[The ABA Problem and Hazard Pointers in Lock-Free Queues]] — Preventing stale pointer reuse using generation tags and hazard pointer registries.
- [[Bounded Blocking Queue (Condition Variables and Mutex Striping)]] — Producer-consumer synchronization with pthread cond/wait and semaphore throttling.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`
