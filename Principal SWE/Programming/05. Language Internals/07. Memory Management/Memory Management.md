---
title: Memory Management
tags:
  - programming
  - language-internals
  - principal-swe
parent: "[[Language Internals]]"
---

# Memory Management

Comprehensive engineering guide, patterns, and principles for Memory Management.

```text
Memory Management
│
├── [[Memory Hierarchy (Memory Management)]]
├── [[Stack vs Heap]]
├── [[Manual Memory Management]]
├── [[Reference Counting]]
├── [[Tracing Garbage Collection]]
├── [[Ownership and Borrowing]]
├── [[Allocators]]
├── [[Escape Analysis]]
├── [[Memory Layout]]
├── [[GC Tuning in Production]]
├── [[Memory Safety]]
├── [[Memory Bugs]]
├── [[Weak References]]
├── [[Finalizers and Destructors]]
├── [[Object Pinning]]
├── [[Off Heap Memory]]
└── [[Memory Pressure and Oom]]
```

---

## 🗂️ Topics

- [[Memory Hierarchy (Memory Management)]]
- [[Stack vs Heap]]
- [[Manual Memory Management]]
- [[Reference Counting]]
- [[Tracing Garbage Collection]]
- [[Ownership and Borrowing]]
- [[Allocators]]
- [[Escape Analysis]]
- [[Memory Layout]]
- [[GC Tuning in Production]]
- [[Memory Safety]]
- [[Memory Bugs]]
- [[Weak References]]
- [[Finalizers and Destructors]]
- [[Object Pinning]]
- [[Off Heap Memory]]
- [[Memory Pressure and Oom]]

- [[Garbage Collection Barrier Mechanics (Read and Write Barriers)]] — How concurrent garbage collectors use hardware and software write/read barriers (Yuasa, Dijkstra, SATB) to preserve tri-color invariants.
- [[TCMalloc and Jemalloc Per-Thread Caching Architecture]] — Thread-caching memory allocators, size-class bins, slab allocation, and reducing global heap lock contention.
- [[Arena and Region Allocation Mechanics in High-Load Systems]] — Fast bump-pointer allocation and O(1) bulk deallocation strategies for request lifecycles and compilers.

---

## 🔗 References
- ⬆️ Parent: [[Language Internals]]
- 🎓 Root: [[Principal SWE]]
