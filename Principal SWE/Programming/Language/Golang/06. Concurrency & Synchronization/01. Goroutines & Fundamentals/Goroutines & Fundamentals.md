---
title: Goroutines & Fundamentals
tags:
  - golang
  - concurrency
parent: "[[Concurrency & Synchronization]]"
---

# Goroutines & Fundamentals

Lightweight threads, goroutine lifecycle, memory footprint (2KB initial stack), and channel basics.

```text
Goroutines & Fundamentals
│
├── [[CSP Concurrency Model]]
├── [[Goroutine Mechanics]]
├── [[Stack Growth & Segmented Stacks]]
├── [[Goroutine Lifecycle & States]]
└── [[Goroutine Leaks & Diagnostics]]
```

---

## 🗂️ Topics

- [[CSP Concurrency Model]] — Communicating Sequential Processes: Do not communicate by sharing memory; share memory by communicating.
- [[Goroutine Mechanics]] — Spawning concurrent execution threads with go keyword, 2KB initial stack allocation.
- [[Stack Growth & Segmented Stacks]] — Contiguous stack growth (2KB to 1GB) and stack copying mechanics.
- [[Goroutine Lifecycle & States]] — _Gidle, _Grunnable, _Grunning, _Gwaiting, _Gsyscall, _Gdead states.
- [[Goroutine Leaks & Diagnostics]] — Identifying blocked goroutines, leaked channels, and diagnostic tools (pprof, runtime.NumGoroutine).

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
