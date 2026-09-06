---
title: Circular Buffer Array (Ring Buffer)
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Circular Buffer Array (Ring Buffer)

Fixed-size contiguous ring array, modular wraparound arithmetic, power-of-two bitwise masking, and lock-free concurrency topologies.

```text
Circular Buffer Array (Ring Buffer)
│
├── [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]]
├── [[Circular Buffer Overwrite and Head-Tail Pointer Invariants]]
├── [[Single-Producer Single-Consumer (SPSC) Lock-Free Array Ring Buffer]]
└── [[Disruptor Ring Buffer Pattern (Cache Line Padded Sequence Tracking)]]
```

---

## 🗂️ Topics & Implementations

- [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]] — Replacing hardware integer modulo (% N) with single-cycle bitwise AND masking (& (N - 1)).
- [[Circular Buffer Overwrite and Head-Tail Pointer Invariants]] — Disambiguating full vs empty buffer states via sequence counters and pointer offsets.
- [[Single-Producer Single-Consumer (SPSC) Lock-Free Array Ring Buffer]] — Wait-free thread synchronization utilizing atomic memory barriers and head/tail sequence tracking.
- [[Disruptor Ring Buffer Pattern (Cache Line Padded Sequence Tracking)]] — High-throughput inter-thread messaging architecture with cache line padding to eliminate false sharing.

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
