---
title: Goroutines & Memory Lifecycle
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Goroutines & Memory Lifecycle

Goroutine spawning, 2KB initial stack allocation, copystack growth, stack shrinking, runtime states, and leak diagnostics.

```text
Goroutines & Memory Lifecycle
│
├── [[CSP Concurrency Model in Go]]
├── [[Goroutine Spawning Mechanics (go keyword)]]
├── [[Goroutines vs OS Threads]]
├── [[Initial 2KB Stack Allocation (stack.go)]]
├── [[Contiguous Stack Growth (copystack)]]
├── [[Stack Shrinking during GC]]
├── [[Stack Splitting Preambles (morestack & nosplit)]]
├── [[12 Goroutine Runtime States]]
├── [[Goroutine Allocation Pool (gfget & gfput)]]
├── [[Goroutine Leaks Diagnostics]]
└── [[Thread Pinning (LockOSThread & UnlockOSThread)]]
```

---

## 🗂️ Topics

- [[CSP Concurrency Model in Go]] — Communicating Sequential Processes: Do not communicate by sharing memory; share memory by communicating.
- [[Goroutine Spawning Mechanics (go keyword)]] — Spawning lightweight user-space execution threads with go keyword and 2KB stack.
- [[Goroutines vs OS Threads]] — Comparing memory footprints (2KB vs 2MB-8MB), creation costs, and context switch overhead.
- [[Initial 2KB Stack Allocation (stack.go)]] — How the Go runtime allocates minimal contiguous stacks for goroutines.
- [[Contiguous Stack Growth (copystack)]] — Stack expansion up to 1GB and pointer fixup translation during memory reallocation.
- [[Stack Shrinking during GC]] — Shrinking underutilized stacks during garbage collection cycles to reclaim memory.
- [[Stack Splitting Preambles (morestack & nosplit)]] — Compiler-inserted stack boundary checks and //go:nosplit leaf function directives.
- [[12 Goroutine Runtime States]] — Dissecting _Gidle, _Grunnable, _Grunning, _Gsyscall, _Gwaiting, _Gdead, _Gcopystack, _Gpreempted.
- [[Goroutine Allocation Pool (gfget & gfput)]] — Reusing dead goroutine memory structures from scheduler free lists.
- [[Goroutine Leaks Diagnostics]] — Identifying blocked goroutines using pprof, runtime.NumGoroutine, and trace dumps.
- [[Thread Pinning (LockOSThread & UnlockOSThread)]] — Binding goroutines to dedicated OS threads for Cgo and graphics libraries.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
