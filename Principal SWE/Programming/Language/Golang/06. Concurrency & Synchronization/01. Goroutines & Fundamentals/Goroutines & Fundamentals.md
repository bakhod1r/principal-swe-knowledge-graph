---
title: Goroutines & Fundamentals
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Goroutines & Fundamentals

Lightweight threads, goroutine lifecycle, memory footprint (2KB initial stack), and channel basics.

```text
Goroutines & Fundamentals
│
├── [[CSP Concurrency Model]]
├── [[Goroutine Mechanics]]
├── [[Goroutines vs OS Threads]]
├── [[Stack Growth & Segmented Stacks]]
├── [[morestack Stack Split & Pointer Reallocation]]
├── [[Goroutine Lifecycle & States]]
└── [[Goroutine Leaks & Diagnostics]]
```

---

## 🗂️ Topics

- [[CSP Concurrency Model]] — Communicating Sequential Processes: Do not communicate by sharing memory; share memory by communicating.
- [[Goroutine Mechanics]] — Spawning concurrent execution threads with go keyword, 2KB initial stack allocation.
- [[Goroutines vs OS Threads]] — Memory footprint, creation cost, context switch overhead (user vs kernel).
- [[Stack Growth & Segmented Stacks]] — Contiguous stack growth (2KB to 1GB) and stack copying mechanics.
- [[morestack Stack Split & Pointer Reallocation]] — Stack frame boundary checks and pointer adjustment during stack copy.
- [[Goroutine Lifecycle & States]] — _Gidle, _Grunnable, _Grunning, _Gwaiting, _Gsyscall, _Gdead states.
- [[Goroutine Leaks & Diagnostics]] — Identifying blocked goroutines, leaked channels, and diagnostic tools (pprof, runtime.NumGoroutine).

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
