---
title: Pointer Memory Management
tags:
  - golang
  - pointers
  - principal-swe
parent: "[[Pointers]]"
---

# Pointer Memory Management & Escape Analysis

Stack vs heap allocation, escape analysis algorithms, and GC mechanics.

```text
Memory Management & Escape Analysis
│
├── `Stack vs Heap Allocation Rules`
├── `Escape Analysis Compiler Passes`
├── `Heap Escape Triggers`
├── [[TCMalloc Allocation Hierarchy (mcache, mcentral, mheap)]]
└── [[Garbage Collection Tricolor Mark-Sweep]]
```

---

## 🗂️ Topics

- `Stack vs Heap Allocation Rules` — Fast stack allocation (bump allocator) vs heap allocation (TCMalloc).
- `Escape Analysis Compiler Passes` — Escape analysis algorithms (go build -gcflags="-m -m") and pointer flow graphs.
- `Heap Escape Triggers` — Pointers returned from functions, interface boxing, dynamic slice sizes, large structs.
- [[TCMalloc Allocation Hierarchy (mcache, mcentral, mheap)]] — Thread-local span caches, size classes, and central memory arenas.
- [[Garbage Collection Tricolor Mark-Sweep]] — Concurrent tricolor marking, hybrid write barrier, and STW pause elimination.


## 🗂️ Contents

- [[Function Inlining Budget (Max 80 AST Nodes)]]
- [[Garbage Collection Tricolor Mark-Sweep]]
- [[TCMalloc Allocation Hierarchy (mcache, mcentral, mheap)]]
- [[Tricolor GC Hybrid Write Barrier Deep Dive]]

---

## 🔗 References
- ⬆️ Parent: [[Pointers]]

