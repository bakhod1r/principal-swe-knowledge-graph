---
title: Garbage Collector Implementation (mgc.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Garbage Collector Implementation (mgc.go)

Concurrent tricolor marking, hybrid write barriers, GC pacer math, mark work buffers, dedicated workers, and scavenging.

```text
Garbage Collector Implementation (mgc.go)
│
├── [[Concurrent Tricolor Mark-Sweep Engine]]
├── [[Hybrid Write Barrier Implementation]]
├── [[GC Pacer Algorithm & Target Trigger Math]]
├── [[GC Mark Termination & Sweep Termination STW Phases]]
├── [[GC Dedicated Worker Goroutines (gcBgMarkWorker)]]
├── [[Mutator Assist Allocations]]
├── [[Mark Work Buffers (gcWork & wbuf)]]
├── [[GC Phase Transitions (gcStart, gcMarkDone, gcSweep)]]
└── [[Memory Purging & Scavenging (scavenger.go)]]
```

---

## 🗂️ Topics

- [[Concurrent Tricolor Mark-Sweep Engine]] — White, Grey, Black object sets, concurrent marking phases, and concurrent sweeping in mgc.go.
- [[Hybrid Write Barrier Implementation]] — Dijkstra insertion barrier combined with Yuasa deletion barrier eliminating Stop-The-World stack rescans.
- [[GC Pacer Algorithm & Target Trigger Math]] — Dynamic feedback loop calculating next_gc trigger threshold based on allocation rate, GOGC, and GOMEMLIMIT.
- [[GC Mark Termination & Sweep Termination STW Phases]] — Sub-millisecond STW pause windows: disabling write barriers, flushing local caches, and starting sweep.
- [[GC Dedicated Worker Goroutines (gcBgMarkWorker)]] — Dedicated mark workers: gcMarkWorkerDedicated, gcMarkWorkerFractional, and gcMarkWorkerIdle.
- [[Mutator Assist Allocations]] — Forcing fast-allocating user goroutines to assist GC marking when allocation outpaces GC throughput.
- [[Mark Work Buffers (gcWork & wbuf)]] — Thread-local grey object work buffers (wbuf1, wbuf2) preventing lock contention during concurrent mark.
- [[GC Phase Transitions (gcStart, gcMarkDone, gcSweep)]] — State machine executing STW mark preparation, concurrent mark, STW mark termination, and concurrent sweep.
- [[Memory Purging & Scavenging (scavenger.go)]] — Background page scavenger returning unused virtual memory spans back to the OS via madvise.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]

