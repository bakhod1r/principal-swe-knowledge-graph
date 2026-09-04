---
title: Goroutines
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Go Concurrency]]"
---

# Goroutines

Goroutine fundamentals, CSP model, spawning mechanics, runtime states, leak diagnostics, deadlock patterns, and thread pinning.

```text
Goroutines
│
├── [[CSP Concurrency Model in Go]]
├── [[Goroutine Spawning Mechanics (go keyword)]]
├── [[Goroutines vs OS Threads]]
├── [[12 Goroutine Runtime States]]
├── [[Goroutine Leaks Diagnostics]]
├── [[Deadlock]]
└── [[Thread Pinning (LockOSThread & UnlockOSThread)]]
```

---

## 🗂️ Topics

- [[CSP Concurrency Model in Go]] — Communicating Sequential Processes: Do not communicate by sharing memory; share memory by communicating.
- [[Goroutine Spawning Mechanics (go keyword)]] — Spawning lightweight user-space execution threads with go keyword and 2KB stack.
- [[Goroutines vs OS Threads]] — Comparing memory footprints (2KB vs 2MB-8MB), creation costs, and context switch overhead.
- [[12 Goroutine Runtime States]] — Dissecting _Gidle, _Grunnable, _Grunning, _Gsyscall, _Gwaiting, _Gdead, _Gcopystack, _Gpreempted.
- [[Goroutine Leaks Diagnostics]] — Identifying blocked goroutines using pprof, runtime.NumGoroutine, and trace dumps.
- [[Deadlock]] — A state where goroutines are permanently blocked, each waiting for an event or resource that can never occur.
- [[Thread Pinning (LockOSThread & UnlockOSThread)]] — Binding goroutines to dedicated OS threads for Cgo and graphics libraries.

---

## 🔗 References
- ⬆️ Parent: [[Go Concurrency]]
- 🔄 Related: [[Goroutine Memory Lifecycle]]
- 📚 Module: `Concurrency & Synchronization`
