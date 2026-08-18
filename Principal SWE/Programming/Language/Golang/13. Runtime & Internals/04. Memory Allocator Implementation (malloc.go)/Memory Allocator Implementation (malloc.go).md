---
title: Memory Allocator Implementation (malloc.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Memory Allocator Implementation (malloc.go)

TCMalloc allocation hierarchy, tiny allocator (<16B), small/large object paths, mspan descriptors, mcentral, and 64MB arenas.

```text
Memory Allocator Implementation (malloc.go)
│
├── [[TCMalloc-Based Allocation Hierarchy]]
├── [[Memory Size Classes & Tiny Allocator (<16B)]]
├── [[Small Object Allocation Path (<32KB)]]
├── [[Large Object Allocation Path (>32KB)]]
├── [[mspan & Page Allocator Architecture]]
├── [[mcentral Span Management & Cache Refilling]]
├── [[Arena Management & Virtual Memory Mapping]]
└── [[FixAlloc Fixed-Size Allocator for Runtime Metadata]]
```

---

## 🗂️ Topics

- [[TCMalloc-Based Allocation Hierarchy]] — Three-tier architecture: thread-local mcache, central mcentral, and global heap mheap.
- [[Memory Size Classes & Tiny Allocator (<16B)]] — 67 size classes, allocating objects <16B grouped into single memory blocks without individual metadata.
- [[Small Object Allocation Path (<32KB)]] — Fast-path lockless allocation from mcache.alloc[spanClass] without global heap locks.
- [[Large Object Allocation Path (>32KB)]] — Direct allocation of contiguous memory pages from global mheap page allocator.
- [[mspan & Page Allocator Architecture]] — Radix tree page allocator (pageAlloc) managing 8KB page chunks and span descriptors.
- [[mcentral Span Management & Cache Refilling]] — Central span list (nonempty / empty lists) providing spans to local mcaches under central lock.
- [[Arena Management & Virtual Memory Mapping]] — 64MB memory arena chunks (heapArena), virtual memory reservations, and mmap kernel mapping.
- [[FixAlloc Fixed-Size Allocator for Runtime Metadata]] — Low-level free-list memory allocator used exclusively for runtime internal structs (mspan, mlink).

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
