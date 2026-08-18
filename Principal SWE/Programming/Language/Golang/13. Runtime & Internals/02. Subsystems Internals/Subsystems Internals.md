- [[GC Pacer Math & Trigger Ratio]] — Calculating next_gc target based on GOGC, GOMEMLIMIT, and allocation rate.

- [[Stack Splitting & copystack() Implementation]] — Dissecting copystack source in stack.go, pointer fixup tables, and frame relocation.

---
title: Subsystems Internals
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Subsystems Internals

Scheduler source (proc.go), GC source (mgc.go), and memory allocator source (malloc.go).

```text
Subsystems Internals
│
├── [[GMP Scheduler Implementation (proc.go)]]
├── [[Garbage Collector Implementation (mgc.go)]]
├── [[Memory Allocator Implementation (malloc.go)]]
├── [[Panic and Recover Implementation (panic.go)]]
└── [[Channel Implementation (chan.go)]]
```

---

## 🗂️ Topics

- [[GMP Scheduler Implementation (proc.go)]] — findRunnable, schedule loop, work stealing algorithm, sysmon preemption.
- [[Garbage Collector Implementation (mgc.go)]] — Tricolor marking, hybrid write barrier, concurrent sweep, GC pacer math.
- [[Memory Allocator Implementation (malloc.go)]] — mcache span caching, mcentral size classes, mheap arenas, huge pages.
- [[Panic and Recover Implementation (panic.go)]] — _panic and _defer linked list unwinding, gopanic, gorecover.
- [[Channel Implementation (chan.go)]] — makechan, chansend, chanrecv, closechan lock and sudog wait queue algorithms.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
