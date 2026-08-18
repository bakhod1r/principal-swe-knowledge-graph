---
title: Generic Data Structures & Collections
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Generic Data Structures & Collections

High-order slice functions, lock-free stacks/queues, skip lists, LRU caches, binary heaps, and monads.

```text
Generic Data Structures & Collections
│
├── [[Generic Slice Wrapper & High-Order Functions]]
├── [[Generic Lock-Free Stack (Treiber Stack)]]
├── [[Generic Lock-Free Queue (Michael-Scott Queue)]]
├── [[Generic Concurrent Skip List]]
├── [[Generic LRU & LFU Cache]]
├── [[Generic Priority Queue (Binary Heap)]]
├── [[Generic Ring Buffer & Circular Queue]]
└── [[Generic Result and Option Monads]]
```

---

## 🗂️ Topics

- [[Generic Slice Wrapper & High-Order Functions]] — Building type-safe Map, Filter, Reduce, FlatMap, and Chunk slice helpers.
- [[Generic Lock-Free Stack (Treiber Stack)]] — Concurrent lock-free LIFO stack using atomic pointer CAS operations.
- [[Generic Lock-Free Queue (Michael-Scott Queue)]] — High-throughput lock-free FIFO queue with atomic head and tail pointers.
- [[Generic Concurrent Skip List]] — Probabilistic search and indexing structure with lockless concurrent reads.
- [[Generic LRU & LFU Cache]] — Thread-safe generic cache with O(1) eviction policies and TTL expiration.
- [[Generic Priority Queue (Binary Heap)]] — Type-safe generic priority queue wrapping container/heap.
- [[Generic Ring Buffer & Circular Queue]] — Fixed-capacity circular ring buffer for zero-allocation stream buffering.
- [[Generic Result and Option Monads]] — Functional error and optionality handling patterns (Result[T, E] and Option[T]).

---

## 🔗 References
- ⬆️ Parent: [[Generics]]

