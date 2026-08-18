---
title: Memory Management Engines
tags:
  - programming
  - language-internals
  - principal-swe
parent: "[[Language Internals]]"
---

# Memory Management Engines

Garbage collection collectors and low-level heap allocators.

```text
Memory Management Engines
│
├── [[Garbage Collection Algorithms (Tricolor, Generational, G1, ZGC)]]
├── [[Reference Counting & Cycle Collection (ARC, Python GC)]]
├── [[Stack vs Heap Allocation & Escape Analysis]]
└── [[Memory Allocators (TCMalloc, jemalloc, Buddy Allocator)]]
```

---

## 🗂️ Topics

- [[Garbage Collection Algorithms (Tricolor, Generational, G1, ZGC)]] — Mark-sweep, copying collectors, generational hypothesis, concurrent colored marking, and region-based collectors.
- [[Reference Counting & Cycle Collection (ARC, Python GC)]] — Deterministic object destruction, reference counting overhead, and resolving cyclic references via trial deletion.
- [[Stack vs Heap Allocation & Escape Analysis]] — Compiler flow-graph algorithms determining when object lifetimes outlive stack frames.
- [[Memory Allocators (TCMalloc, jemalloc, Buddy Allocator)]] — Segregated size classes, thread-local caching, virtual memory arenas, and mitigating heap fragmentation.

---

## 🔗 References
- ⬆️ Parent: [[Language Internals]]
- 🎓 Root: [[Principal SWE]]
