---
title: Concurrent & Lock-Free Queue
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Queues]]"
---

# 📦 Concurrent & Lock-Free Queue

Master topology of concurrent thread-safe queues.

```text
Concurrent & Lock-Free Queue
│
├── [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]]
├── [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]]
├── [[Michael-Scott Lock-Free FIFO Queue (CAS on Head and Tail)]]
├── [[The ABA Problem and Hazard Pointers in Lock-Free Queues]]
└── [[Bounded Blocking Queue (Condition Variables and Mutex Striping)]]
```

---

## 🗂️ Topics & Implementations

- [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]] — Single-producer single-consumer ring buffer with atomic memory barriers and cache line padding.
- [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]] — Lock-free concurrent queue using atomic compare-and-swap on head and tail sequences.
- [[Michael-Scott Lock-Free FIFO Queue (CAS on Head and Tail)]] — Standard concurrent linked queue using atomic pointer swings on node chains.
- [[The ABA Problem and Hazard Pointers in Lock-Free Queues]] — Preventing stale pointer reuse using generation tags and hazard pointer registries.
- [[Bounded Blocking Queue (Condition Variables and Mutex Striping)]] — Producer-consumer synchronization with pthread cond/wait and semaphore throttling.

---

## 🔗 References
- ⬆️ Parent: [[Queues]]
- 📚 Module: `Data Structures`
