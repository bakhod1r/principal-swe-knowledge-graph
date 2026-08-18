---
title: Go Memory Model & Hardware Concurrency
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Go Memory Model & Hardware Concurrency

Go Memory Model specification, happens-before, instruction reordering, memory barriers, and cache coherency.

```text
Go Memory Model & Hardware Concurrency
│
├── [[Go Memory Model Specification]]
├── [[Happens-Before Relationship Rules]]
├── [[Instruction Reordering (Compiler & CPU Out-of-Order)]]
├── [[Memory Barriers & CPU Store Buffers]]
├── [[CPU Cache Hierarchy & Cache Lines (64-byte)]]
├── [[Cache Coherency Protocols (MESI & MOESI)]]
├── [[False Sharing & Cache Line Invalidation]]
└── [[Data Race vs Race Condition Deep Dive]]
```

---

## 🗂️ Topics

- [[Go Memory Model Specification]] — Formal rules defining when a write to a variable by one goroutine is visible to another.
- [[Happens-Before Relationship Rules]] — Establishing strict happens-before edges via channels, mutexes, sync.Once, and goroutines.
- [[Instruction Reordering (Compiler & CPU Out-of-Order)]] — How compilers and CPU out-of-order execution reorder memory instructions.
- [[Memory Barriers & CPU Store Buffers]] — Hardware memory fences (MFENCE, SFENCE, LFENCE) and store buffer flushing.
- [[CPU Cache Hierarchy & Cache Lines (64-byte)]] — L1/L2/L3 CPU caches, 64-byte cache line granularity, and latency hierarchy.
- [[Cache Coherency Protocols (MESI & MOESI)]] — Modified, Exclusive, Shared, Invalid cache line state transitions across CPU cores.
- [[False Sharing & Cache Line Invalidation]] — Contention between independent variables on the same 64-byte cache line and padding fixes.
- [[Data Race vs Race Condition Deep Dive]] — Distinguishing undefined behavior data races from high-level logical race conditions.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
