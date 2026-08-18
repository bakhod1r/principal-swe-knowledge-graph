---
title: Synchronization Primitives (sync)
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Synchronization Primitives (sync)

sync.Mutex (Normal vs Starvation), sync.RWMutex, sync.WaitGroup, sync.Once, sync.Pool, sync.Map, and sync.Cond.

```text
Synchronization Primitives (sync)
│
├── [[sync.Mutex (Normal vs Starvation Mode)]]
├── [[sync.RWMutex (Reader Starvation Prevention)]]
├── [[sync.WaitGroup State Bitpacking]]
├── [[sync.Once & Fast-Path Double-Checked Locking]]
├── [[sync.Pool Architecture (poolLocal & Victim Cache)]]
├── [[sync.Map Architecture (readOnly vs dirty Map)]]
├── [[sync.Cond Condition Variables]]
└── [[Mutex vs Channel Selection Matrix]]
```

---

## 🗂️ Topics

- [[sync.Mutex (Normal vs Starvation Mode)]] — Bimodal mutex algorithm: high-throughput CPU spin vs fair FIFO handoff at 1ms latency threshold.
- [[sync.RWMutex (Reader Starvation Prevention)]] — Reader-writer lock with atomic reader counters and writer priority signaling.
- [[sync.WaitGroup State Bitpacking]] — 64-bit/128-bit atomic state bitpacking combining task counter and waiter counter.
- [[sync.Once & Fast-Path Double-Checked Locking]] — Atomic fast-path initialization and OnceFunc, OnceValue, OnceValues wrappers (Go 1.21+).
- [[sync.Pool Architecture (poolLocal & Victim Cache)]] — Per-P private slots, shared lock-free deques, poolVictim caches, and GC cleansing cycles.
- [[sync.Map Architecture (readOnly vs dirty Map)]] — Lockless atomic reads from readOnly map, dirty map fallbacks, and amortized promotions.
- [[sync.Cond Condition Variables]] — Coordinating goroutines with Wait(), Signal(), Broadcast(), and lost signal hazards.
- [[Mutex vs Channel Selection Matrix]] — Staff-level architectural decision tree for shared memory vs message passing.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
