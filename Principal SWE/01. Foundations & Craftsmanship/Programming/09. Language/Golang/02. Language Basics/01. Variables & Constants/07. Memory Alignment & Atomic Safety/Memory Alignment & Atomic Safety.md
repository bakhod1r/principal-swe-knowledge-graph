---
title: Memory Alignment & Atomic Safety
tags:
  - golang
  - variables
  - memory-alignment
  - atomic
  - principal-swe
parent: "[[Variables & Constants]]"
---

# Memory Alignment & Atomic Safety

Memory layout, word boundary alignment, 64-bit alignment constraints on 32-bit architectures, atomic panic prevention, and cache line false sharing.

```text
Memory Alignment & Atomic Safety
│
├── [[Struct Field Reordering for Alignment]]
├── [[64-Bit Alignment on 32-Bit Archs]]
├── [[Atomic Panic Prevention]]
├── [[Go 1.19+ Typed Atomics and atomic.Pointer]]
├── [[False Sharing and Cache Line Alignment]]
├── [[Zero-Sized Types and Trailing Struct Field Padding]]
├── [[Memory Ordering and Go Memory Model]]
├── [[Hardware Alignment and Unaligned Memory Access]]
└── [[Fieldalignment Linter and Struct Optimization Tooling]]
```

---

## 🗂️ Topics

- [[Struct Field Reordering for Alignment]] — Eliminating padding bytes by ordering struct fields from largest to smallest.
- [[64-Bit Alignment on 32-Bit Archs]] — Why 64-bit atomic operations panic on unaligned 32-bit memory boundaries.
- [[Atomic Panic Prevention]] — Preventing runtime panics from unaligned atomics, type mismatches, and nil storage.
- [[Go 1.19+ Typed Atomics and atomic.Pointer]] — Safe 64-bit alignment across all CPUs, generic atomic pointers, and zero-unsafe idioms.
- [[False Sharing and Cache Line Alignment]] — Eliminating multi-core L1/L2 cache line bouncing via padding and memory striped counters.
- [[Zero-Sized Types and Trailing Struct Field Padding]] — Behavior of `struct{}` and trailing word padding to protect heap boundary allocations.
- [[Memory Ordering and Go Memory Model]] — Happens-before relationships, sequential consistency, and store-load barriers.
- [[Hardware Alignment and Unaligned Memory Access]] — Memory bus word boundaries, x86 vs ARM alignment faults, and `encoding/binary` safety.
- [[Fieldalignment Linter and Struct Optimization Tooling]] — Automated struct layout analysis and memory footprint minimization tooling.

---

## 🔗 References
- ⬆️ Parent: [[Variables & Constants]]
- 📚 Module: `Language Basics`
