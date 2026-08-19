---
title: CPU Cache & Hardware-Conscious Design
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Go Performance Engineering]]"
---

# CPU Cache & Hardware-Conscious Design

CPU cache hierarchy, 64-byte cache lines, false sharing elimination, sequential memory layouts, and BCE.

```text
CPU Cache & Hardware-Conscious Design
│
├── [[CPU Cache Hierarchy (L1, L2, L3, Cache Lines)]]
├── [[False Sharing Elimination with Cache Line Padding]]
├── [[Sequential Memory Access & Cache Locality]]
├── [[SIMD Vectorization & Bounds Check Elimination (BCE)]]
└── [[Branch Prediction & Hot-Path Code Alignment]]
```

---

## 🗂️ Topics

- [[CPU Cache Hierarchy (L1, L2, L3, Cache Lines)]] — 64-byte cache line structure, latency tiers (1ns L1 vs 100ns RAM), and temporal/spatial locality.
- [[False Sharing Elimination with Cache Line Padding]] — Placing [64]byte padding between high-contention atomic counters on multicore CPUs.
- [[Sequential Memory Access & Cache Locality]] — Array-of-Structs (AoS) vs Struct-of-Arrays (SoA) data layouts for CPU prefetchers.
- [[SIMD Vectorization & Bounds Check Elimination (BCE)]] — Helping the compiler eliminate slice bounds checks and emit vectorized CPU instructions.
- [[Branch Prediction & Hot-Path Code Alignment]] — Structuring error branches and happy paths to maximize CPU branch prediction hit rates.

---

## 🔗 References
- ⬆️ Parent: `Performance Engineering & Profiling`

