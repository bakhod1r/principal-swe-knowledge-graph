---
title: Runtime Scheduler & Internals
tags:
  - golang
  - concurrency
parent: "[[Concurrency & Synchronization]]"
---

# Runtime Scheduler & Internals

GMP scheduler, work stealing, sysmon daemon, signal preemption, and lock-free CAS.

```text
Runtime Scheduler & Internals
│
├── [[GMP Model (G, M, P)]]
├── [[Work Stealing Algorithm]]
├── [[Sysmon Background Daemon]]
├── [[Signal-Based Async Preemption (SIGURG)]]
├── [[Netpoller]]
├── [[Lock-Free Programming & CAS]]
└── [[GOMAXPROCS Tuning]]
```

---

## 🗂️ Topics

- [[GMP Model (G, M, P)]] — Goroutines (G), OS Threads (M), Logical Processors (P), and scheduling queues.
- [[Work Stealing Algorithm]] — O(1) local run queue checking, global run queue starvation check, and random P stealing.
- [[Sysmon Background Daemon]] — Monitoring blocked syscalls, forcing periodic GC, and triggering preemption.
- [[Signal-Based Async Preemption (SIGURG)]] — Non-cooperative async signal preemption of tight loops without function calls.
- [[Netpoller]] — Non-blocking I/O event multiplexing (epoll, kqueue, IOCP) integrated with GMP scheduler.
- [[Lock-Free Programming & CAS]] — sync/atomic primitives, Compare-And-Swap algorithms, and memory fences.
- [[GOMAXPROCS Tuning]] — Tuning CPU core allocation and container CPU quota limits (automaxprocs).

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
