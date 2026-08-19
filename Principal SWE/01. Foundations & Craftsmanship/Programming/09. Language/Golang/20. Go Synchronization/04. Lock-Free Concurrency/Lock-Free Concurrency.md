---
title: Lock-Free & Atomic Concurrency
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Go Synchronization]]"
---

# Lock-Free & Atomic Concurrency

sync/atomic primitives, CAS loops, atomic.Pointer[T], lock-free queues, stacks, and ring buffers.

```text
Lock-Free & Atomic Concurrency
│
├── `sync-atomic Primitives (Load, Store, CAS, Swap, Add)`
├── `Atomic CAS Loop Pattern`
├── [[atomic.Pointer[T] & atomic.Value Type Safety]]
├── [[Lock-Free Stack (Treiber Stack)]]
├── [[Lock-Free Queue (Michael-Scott Queue)]]
├── [[Lock-Free Ring Buffer]]
└── [[Lock-Free vs Mutex Performance Benchmarks]]
```

---

## 🗂️ Topics

- `sync-atomic Primitives (Load, Store, CAS, Swap, Add)` — Hardware atomic primitives providing sequential consistency without mutex locks.
- `Atomic CAS Loop Pattern` — Optimistic concurrency control with Compare-And-Swap spin loops.
- [[atomic.Pointer[T] & atomic.Value Type Safety]] — Type-safe atomic pointers and atomic value containers in Go 1.19+.
- [[Lock-Free Stack (Treiber Stack)]] — Implementing a concurrent lock-free LIFO stack using atomic pointer CAS.
- [[Lock-Free Queue (Michael-Scott Queue)]] — Implementing a concurrent lock-free FIFO queue with head and tail pointers.
- [[Lock-Free Ring Buffer]] — High-throughput single-producer single-consumer (SPSC) and MPMC lock-free buffers.
- [[Lock-Free vs Mutex Performance Benchmarks]] — Measuring throughput, CPU cache pressure, and scalability tradeoffs.

---

## 🔗 References
- ⬆️ Parent: `Concurrency & Synchronization`

