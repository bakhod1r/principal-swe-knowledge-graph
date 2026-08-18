---
title: SIMD Vectorization & Mechanical Sympathy
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Performance Engineering & Profiling]]"
---

# SIMD Vectorization & Mechanical Sympathy

Hardware SIMD intrinsics, branchless coding, memory prefetching, GC-free architectures, and mechanical sympathy.

```text
SIMD Vectorization & Mechanical Sympathy
│
├── [[AVX-512 & ARM NEON Vectorization via Go Assembly]]
├── [[Zero-Allocation Base64 & Hex SIMD Encoding]]
├── [[Fast Hash Functions (AHash, WyHash, XXHash)]]
├── [[Branchless Programming Idioms in Go]]
├── [[Memory Prefetching (prefetchnta & prefetcht0)]]
├── [[Garbage Collection-Free Execution Architecture]]
└── [[Mechanical Sympathy in Go Systems Design]]
```

---

## 🗂️ Topics

- [[AVX-512 & ARM NEON Vectorization via Go Assembly]] — Writing vectorized SIMD vector math, string scanning, and matrix operations in Plan 9 assembly.
- [[Zero-Allocation Base64 & Hex SIMD Encoding]] — Hardware-accelerated byte encoding and decoding bypassing scalar standard library loops.
- [[Fast Hash Functions (AHash, WyHash, XXHash)]] — Replacing standard cryptographic/FNV hashes with 10GB/s non-cryptographic hashes for hash tables and caches.
- [[Branchless Programming Idioms in Go]] — Eliminating CPU branch mispredictions using bitwise selection, conditional moves, and arithmetic masking.
- [[Memory Prefetching (prefetchnta & prefetcht0)]] — Explicitly loading upcoming memory cache lines into L1/L2 caches ahead of execution cycles.
- [[Garbage Collection-Free Execution Architecture]] — Designing mission-critical financial/trading systems running with GOGC=off or pre-allocated off-heap memory.
- [[Mechanical Sympathy in Go Systems Design]] — Designing Go software architecture in complete harmony with CPU execution pipelines, caches, and memory buses.

---

## 🔗 References
- ⬆️ Parent: [[Performance Engineering & Profiling]]
- 🎓 Root: [[Principal SWE]]
