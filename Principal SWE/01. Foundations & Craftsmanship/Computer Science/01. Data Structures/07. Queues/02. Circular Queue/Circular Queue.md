---
title: Circular Queue
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Queues]]"
---

# 📦 Circular Queue

Master topology of bounded circular ring buffers.

```text
Circular Queue
│
├── [[Queue Circular Array Buffer Implementation]]
├── [[Queue Power-of-Two Bitwise Masking Indexing]]
├── [[Queue Capacity Growth and Dynamic Resizing]]
├── [[Circular Queue Monotonic Sequence Number Indexing]]
└── [[Bounded Queue Backpressure and Drop Policies (Drop-Head vs Drop-Tail)]]
```

---

## 🗂️ Topics & Implementations

- [[Queue Circular Array Buffer Implementation]] — Reusing fixed array memory via circular head/tail wraparound.
- [[Queue Power-of-Two Bitwise Masking Indexing]] — Replacing modulo division (% N) with ultra-fast bitwise AND masking (& (N - 1)).
- [[Queue Capacity Growth and Dynamic Resizing]] — Unwrapping circular buffers into linear arrays during geometric doubling.
- [[Circular Queue Monotonic Sequence Number Indexing]] — 64-bit monotonic sequences preventing counter overflow and index wrapping errors.
- [[Bounded Queue Backpressure and Drop Policies (Drop-Head vs Drop-Tail)]] — Shedding load in real-time pipelines: evicting oldest data vs rejecting newest.

---

## 🔗 References
- ⬆️ Parent: [[Queues]]
- 📚 Module: `Data Structures`
