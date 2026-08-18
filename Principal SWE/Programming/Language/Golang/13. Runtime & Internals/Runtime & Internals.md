---
title: Runtime & Internals
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Golang]]"
---

# ⚙️ Runtime & Internals

Source code dive into Go runtime internals: GMP scheduler (proc.go), Garbage Collector (mgc.go), TCMalloc-style allocator (malloc.go), and channels (chan.go).

```text
Runtime & Internals
│
├── [[Core Architecture|01. Core Architecture]]
│   ├── [[Runtime Bootstrapping (rt0_go, schedinit)]]
│   ├── [[sysmon Background Daemon Thread]]
│   └── [[runtime Package Diagnostic APIs]]
└── [[Subsystems Internals|02. Subsystems Internals]]
│   ├── [[GMP Scheduler Implementation (proc.go)]]
│   ├── [[Garbage Collector Implementation (mgc.go)]]
│   ├── [[Memory Allocator Implementation (malloc.go)]]
│   └── [[Panic and Recover Implementation (panic.go)]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Core Architecture|01. Core Architecture]]
- [[Runtime Bootstrapping (rt0_go, schedinit)]] — Entry point assembly, OS thread creation, runtime initialization, main goroutine start.
- [[sysmon Background Daemon Thread]] — Tick loop, retaking stuck Ps from long syscalls, forcing GC, preemption.
- [[runtime Package Diagnostic APIs]] — runtime.Gosched, runtime.LockOSThread, runtime.NumGoroutine, runtime.GC, runtime.ReadMemStats.
### 2. 📂 [[Subsystems Internals|02. Subsystems Internals]]
- [[GMP Scheduler Implementation (proc.go)]] — findRunnable, schedule loop, work stealing algorithm, sysmon preemption.
- [[Garbage Collector Implementation (mgc.go)]] — Tricolor marking, hybrid write barrier, concurrent sweep, GC pacer math.
- [[Memory Allocator Implementation (malloc.go)]] — mcache span caching, mcentral size classes, mheap arenas, huge pages.
- [[Panic and Recover Implementation (panic.go)]] — _panic and _defer linked list unwinding, gopanic, gorecover.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

