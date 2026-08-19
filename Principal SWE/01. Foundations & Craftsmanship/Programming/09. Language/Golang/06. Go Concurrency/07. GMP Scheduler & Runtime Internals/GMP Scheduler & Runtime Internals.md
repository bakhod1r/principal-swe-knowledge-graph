---
title: GMP Scheduler & Runtime Internals
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Go Concurrency]]"
---

# GMP Scheduler & Runtime Internals

GMP architecture, local/global runqueues, work stealing, sysmon, signal preemption (SIGURG), and CFS quotas.

```text
GMP Scheduler & Runtime Internals
│
├── [[GMP Model (G, M, P, Sched Structs)]]
├── [[Runqueue Architecture (Local vs Global)]]
├── [[Work Stealing Algorithm]]
├── [[Netpoller (epoll, kqueue, IOCP) Integration]]
├── [[Syscall Handling & M Handoff]]
├── [[Sysmon Daemon Thread]]
├── [[Signal-Based Async Preemption (SIGURG)]]
├── [[Thread Parking (notesleep & notewakeup)]]
└── [[GOMAXPROCS & Container CFS Quota Throttling]]
```

---

## 🗂️ Topics

- [[GMP Model (G, M, P, Sched Structs)]] — Goroutines (G), OS Threads (M), Logical Processors (P), and global scheduler state.
- [[Runqueue Architecture (Local vs Global)]] — 256-element lock-free local runqueue per P and mutex-guarded global runqueue.
- [[Work Stealing Algorithm]] — Checking local queue, 1/61 global check, netpoller check, and stealing half from random P.
- [[Netpoller (epoll, kqueue, IOCP) Integration]] — Asynchronous non-blocking I/O event loop integrated directly with scheduler parking.
- [[Syscall Handling & M Handoff]] — entersyscall, exitsyscall, detaching P from blocked M, and waking parked threads.
- [[Sysmon Daemon Thread]] — Background monitoring thread: forcing periodic GC, retaking stuck Ps, and preemption.
- [[Signal-Based Async Preemption (SIGURG)]] — Non-cooperative async preemption of tight compute loops via OS signals.
- [[Thread Parking (notesleep & notewakeup)]] — Low-level OS thread sleeping and futex waking mechanisms in the runtime.
- [[GOMAXPROCS & Container CFS Quota Throttling]] — Kubernetes CPU limits, CFS period/quota calculation, and automaxprocs.

---

## 🔗 References
- ⬆️ Parent: `Concurrency & Synchronization`

