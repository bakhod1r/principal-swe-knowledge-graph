---
title: Dynamic Array
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Dynamic Array

Resizable heap-allocated contiguous vector, amortized geometric growth policies (1.5x vs 2.0x), size hints, and allocator arena interactions.

```text
Dynamic Array
│
├── [[Array Append and Dynamic Capacity Allocation]]
├── [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]]
├── [[Array Pre-Allocation and Capacity Hints]]
├── [[Array Push and Pop (LIFO Dynamic Stack Operations)]]
├── [[Array Memory Leaks and Garbage Collection Truncation]]
├── [[Array Clear and Logical vs Physical Truncation]]
├── [[Array Shrink-to-Fit and Memory Compaction]]
└── [[Dynamic Array Allocator Arena Integration (jemalloc, TCMalloc)]]
```

---

## 🗂️ Topics & Implementations

- [[Array Append and Dynamic Capacity Allocation]] — Amortized O(1) growth, allocation doubling proofs, and reallocation copy overhead profiles.
- [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]] — Mathematical amortization proofs, allocator memory reuse, and heap fragmentation economics.
- [[Array Pre-Allocation and Capacity Hints]] — Zero-reallocation ingestion via size hinting and heap reserve sizing.
- [[Array Push and Pop (LIFO Dynamic Stack Operations)]] — Strict O(1) LIFO operations on dynamic arrays with top pointer tracking.
- [[Array Memory Leaks and Garbage Collection Truncation]] — Preventing memory retention by pointer zeroing and sub-slice truncation in managed runtimes.
- [[Array Clear and Logical vs Physical Truncation]] — O(1) length reset vs O(N) memory zeroing, GC reclamation, and security zeroing.
- [[Array Shrink-to-Fit and Memory Compaction]] — Reclaiming unused capacity buffers via reallocation and hysteresis shrink thresholds.
- [[Dynamic Array Allocator Arena Integration (jemalloc, TCMalloc)]] — Interactions between dynamic vector resizing, size classes, and thread-local allocator caches.

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
