---
title: Linear FIFO Queue
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Queues]]"
---

# 📦 Linear FIFO Queue

Master topology of linear sequential FIFO queues.

```text
Linear FIFO Queue
│
├── [[Queue FIFO Invariants and State Transitions]]
├── [[Queue Empty and Full Invariant Conditions]]
├── [[Queue Enqueue (Tail Insertion)]]
├── [[Queue Dequeue (Head Removal)]]
├── [[Queue Peek (Front Element Inspection)]]
├── [[Queue Linked List Implementation]]
├── [[Queue Underflow and Overflow Exceptions vs Non-Blocking Options]]
└── [[Two-Stack FIFO Queue Implementation and Amortized O(1) Proof]]
```

---

## 🗂️ Topics & Implementations

- [[Queue FIFO Invariants and State Transitions]] — First-In First-Out operational axioms and queue state machine bounds.
- [[Queue Empty and Full Invariant Conditions]] — Distinguishing full vs empty states via element count or reserved index.
- [[Queue Enqueue (Tail Insertion)]] — O(1) tail element insertion with capacity overflow validation.
- [[Queue Dequeue (Head Removal)]] — O(1) head element extraction and front pointer advancement.
- [[Queue Peek (Front Element Inspection)]] — O(1) viewing head element without mutating queue state.
- [[Queue Linked List Implementation]] — Heap-allocated node queue with head and tail pointers avoiding capacity limits.
- [[Queue Underflow and Overflow Exceptions vs Non-Blocking Options]] — Handling bounded queue states: throwing exceptions, blocking, or dropping.
- [[Two-Stack FIFO Queue Implementation and Amortized O(1) Proof]] — Implementing FIFO queues using two LIFO stacks with amortized O(1) operations.

---

## 🔗 References
- ⬆️ Parent: [[Queues]]
- 📚 Module: `Data Structures`
