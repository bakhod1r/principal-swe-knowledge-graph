---
title: Goroutines & Fundamentals
tags:
  - golang
  - concurrency
parent: "[[Concurrency & Synchronization]]"
---

# Goroutines & Fundamentals

CSP concurrency, goroutines vs OS threads, stack growth, lifecycle states, and leak diagnostics.

```text
Goroutines & Fundamentals
│
├── [[CSP Concurrency Model]]
├── [[Goroutine Mechanics]]
├── [[Goroutines vs OS Threads]]
├── [[Stack Growth & Segmented Stacks]]
├── [[Goroutine Lifecycle & States]]
└── [[Goroutine Leaks & Diagnostics]]
```

---

## 🗂️ Topics

- [[CSP Concurrency Model]] — Communicating Sequential Processes: Share memory by communicating, not vice versa.
- [[Goroutine Mechanics]] — Spawning lightweight user-space threads with go keyword and 2KB stack allocation.
- [[Goroutines vs OS Threads]] — Memory footprint, creation cost, context switch overhead (user vs kernel).
- [[Stack Growth & Segmented Stacks]] — Contiguous stack growth (2KB to 1GB) and stack reallocation copying.
- [[Goroutine Lifecycle & States]] — _Gidle, _Grunnable, _Grunning, _Gwaiting, _Gsyscall, and _Gdead states.
- [[Goroutine Leaks & Diagnostics]] — Identifying blocked goroutines, leaked channels, pprof goroutine dump analysis.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
