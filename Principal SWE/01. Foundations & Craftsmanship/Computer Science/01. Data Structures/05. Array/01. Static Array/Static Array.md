---
title: Static Array
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Static Array

Contiguous physical memory buffer, compile-time/fixed capacity, pointer arithmetic, SIB register indexing, and compiler Bounds Checking Elimination (BCE).

```text
Static Array
│
├── [[Array Memory Layout and Pointer Arithmetic]]
├── [[Array Access and Indexing Mechanics]]
├── [[Array Memory Alignment, Cache Line Striding, and False Sharing]]
├── [[Array Bounds Checking Elimination (BCE) Compiler Optimization]]
├── [[Array Memory Copy (SIMD rep movsq and memmove)]]
├── [[Structure of Arrays (SoA) vs Array of Structures (AoS)]]
├── [[Array Data-Oriented Design (DOD) and Cache Line Utilization]]
├── [[Variable-Length Arrays (VLA) vs Fixed Stack Allocation]]
└── [[Array Address Space Layout Randomization (ASLR) and Stack Protection]]
```

---

## 🗂️ Topics & Implementations

- [[Array Memory Layout and Pointer Arithmetic]] — Linear contiguous byte addressing, scale-index-base translation, and pointer offsets.
- [[Array Access and Indexing Mechanics]] — O(1) random-access mechanics, hardware SIB register decoding, and zero-based offset proofs.
- [[Array Memory Alignment, Cache Line Striding, and False Sharing]] — CPU memory bus alignment, 64-byte L1 cache line prefetching, and multi-threaded false sharing.
- [[Array Bounds Checking Elimination (BCE) Compiler Optimization]] — Compiler optimization eliminating runtime bounds check branch overheads via loop induction analysis.
- [[Array Memory Copy (SIMD rep movsq and memmove)]] — Hardware vectorized block transfers, AVX-512 alignment, and overlapping buffer guarantees.
- [[Structure of Arrays (SoA) vs Array of Structures (AoS)]] — Data-oriented design, SIMD memory packing, and cache throughput optimizations.
- [[Array Data-Oriented Design (DOD) and Cache Line Utilization]] — Organizing contiguous arrays for maximum memory bandwidth and CPU pipelining efficiency.
- [[Variable-Length Arrays (VLA) vs Fixed Stack Allocation]] — Stack pointer adjustment runtime risks, alloca mechanics, and stack overflow pitfalls.
- [[Array Address Space Layout Randomization (ASLR) and Stack Protection]] — Buffer overflow exploit mitigations, stack canaries, and hardware memory protection.

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
