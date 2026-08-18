---
title: Sync & Context Primitives
tags:
  - golang
  - concurrency
parent: "[[Concurrency & Synchronization]]"
---

# Sync & Context Primitives

Mutexes, condition variables, atomic operations, sync.Pool, and context propagation trees.

```text
Sync & Context Primitives
│
├── [[sync.Mutex (Normal vs Starvation)]]
├── [[sync.RWMutex]]
├── [[sync.WaitGroup]]
├── [[sync.Once]]
├── [[sync.Pool]]
├── [[sync.Map]]
├── [[context.Context Tree]]
├── [[x-sync-errgroup]]
└── [[Mutex vs Channel Selection]]
```

---

## 🗂️ Topics

- [[sync.Mutex (Normal vs Starvation)]] — Bimodal mutex algorithm: high throughput spin vs fair FIFO handoff.
- [[sync.RWMutex]] — Reader-writer lock with writer starvation prevention.
- [[sync.WaitGroup]] — Atomic counter synchronization for coordinating goroutine completion.
- [[sync.Once]] — Atomic fast-path initialization with double-checked locking.
- [[sync.Pool]] — Lock-free per-P cache for allocating and reusing short-lived temporary objects.
- [[sync.Map]] — Concurrent map optimized for append-only keys and disjoint key reads.
- [[context.Context Tree]] — Cancellation propagation, deadlines, timeouts, and request-scoped values.
- [[x-sync-errgroup]] — Managing concurrent subtasks with error propagation and context cancellation.
- [[Mutex vs Channel Selection]] — When to use shared memory synchronization vs message passing.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
