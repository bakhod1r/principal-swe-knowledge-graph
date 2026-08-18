---
title: Memory Management & Compiler
tags:
  - golang
  - advanced
parent: "[[Advanced Topics & Low-Level Go]]"
---

# Memory Management & Compiler

Stack vs heap decisions, escape analysis algorithms, compiler SSA optimization passes, and PGO.

```text
Memory Management & Compiler
│
├── [[Memory Allocation Hierarchy (mcache, mcentral, mheap)]]
├── [[Escape Analysis Algorithms]]
├── [[Compiler SSA Optimization Passes]]
├── [[Compiler & Linker Flags]]
└── [[Profile-Guided Optimization (PGO)]]
```

---

## 🗂️ Topics

- [[Memory Allocation Hierarchy (mcache, mcentral, mheap)]] — TCMalloc-inspired per-P caches, central spans, size classes, and page arenas.
- [[Escape Analysis Algorithms]] — Compiler escape analysis rules (go build -gcflags="-m -m"), leaking params, flow analysis.
- [[Compiler SSA Optimization Passes]] — Static Single Assignment intermediate representation, Dead Code Elimination, Bounds Check Elimination (BCE).
- [[Compiler & Linker Flags]] — -gcflags="-N -l", -ldflags="-s -w -X main.version=1.0", -trimpath reproducibility.
- [[Profile-Guided Optimization (PGO)]] — Feeding production CPU profiles into go build for 2-14% compiler speedups.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]
- 🎓 Root: [[Principal SWE]]
