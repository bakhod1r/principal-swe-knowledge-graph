---
title: Sync & Context Primitives
tags:
  - golang
  - sync
parent: "[[Concurrency & Synchronization]]"
---

# Sync & Context Primitives

sync package mutexes, WaitGroup, Once, Pool, Map, Cond, and Context cancellation trees.

```text
Sync & Context Primitives
│
├── [[sync.Mutex (Normal vs Starvation)]]
├── [[sync.RWMutex]]
├── [[sync.WaitGroup]]
├── [[sync.Once & sync.OnceFunc]]
├── [[sync.Pool]]
├── [[sync.Map]]
├── [[sync.Cond]]
├── [[context.Context Tree]]
├── [[x-sync-errgroup]]
├── [[x-sync-singleflight]]
├── [[x-sync-semaphore]]
└── [[Mutex vs Channel Selection]]
```

---

## 🗂️ Topics

- [[sync.Mutex (Normal vs Starvation)]] — Bimodal mutex algorithm: high throughput spin vs fair FIFO handoff.
- [[sync.RWMutex]] — Reader-writer lock with writer starvation prevention algorithm.
- [[sync.WaitGroup]] — Atomic counter synchronization for coordinating goroutine completion.
- [[sync.Once & sync.OnceFunc]] — Atomic fast-path initialization and Go 1.21+ OnceFunc/OnceValue wrappers.
- [[sync.Pool]] — Lock-free per-P cache for allocating and reusing short-lived temporary objects.
- [[sync.Map]] — Concurrent map optimized for append-only keys and disjoint key reads.
- [[sync.Cond]] — Condition variables for broadcasting signals to waiting goroutines.
- [[context.Context Tree]] — Cancellation propagation, deadlines, timeouts, and request-scoped values.
- [[x-sync-errgroup]] — Managing concurrent subtasks with error propagation and context cancellation.
- [[x-sync-singleflight]] — Suppressing duplicate in-flight function calls (request coalescing).
- [[x-sync-semaphore]] — Weighted semaphore for limiting concurrent resource access.
- [[Mutex vs Channel Selection]] — Decision framework: shared memory synchronization vs message passing.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
