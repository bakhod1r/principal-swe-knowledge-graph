---
title: GMP Scheduler Implementation (proc.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# GMP Scheduler Implementation (proc.go)

Scheduler execution loop, findrunnable work search hierarchy, lock-free work stealing, context switching, goroutine creation, and sysmon.

```text
GMP Scheduler Implementation (proc.go)
│
├── [[Scheduler Loop (schedule() function)]]
├── [[findrunnable() Work Search Hierarchy]]
├── [[Work Stealing Algorithm Implementation (runqsteal)]]
├── [[Context Switch Mechanics (gogo & mcall)]]
├── [[Goroutine Creation Mechanics (newproc1 & newproc)]]
├── [[M Parking, Waking & Thread Pool (startm & stopm)]]
├── [[Syscall Handling (entersyscall & exitsyscall)]]
├── [[Sysmon Daemon Implementation]]
└── [[Non-Cooperative Async Preemption (preemptone & asyncPreempt)]]
```

---

## 🗂️ Topics

- [[Scheduler Loop (schedule() function)]] — Step-by-step walkthrough of schedule() in proc.go, finding runnable work, and context switching.
- [[findrunnable() Work Search Hierarchy]] — Checking local runqueue, global runqueue (1/61 check), netpoller, work stealing from peers, and idle parking.
- [[Work Stealing Algorithm Implementation (runqsteal)]] — Lock-free stealing half of peer logical processor local runqueue using atomic CAS on head/tail indices.
- [[Context Switch Mechanics (gogo & mcall)]] — Assembly context switching: saving callee-saved registers into g.sched and restoring target goroutine state.
- [[Goroutine Creation Mechanics (newproc1 & newproc)]] — Allocating new g struct from per-P gFree cache or heap, setting up entry PC, and enqueueing to runq.
- [[M Parking, Waking & Thread Pool (startm & stopm)]] — Spinning threads, notesleep parking, waking idle threads, and managing OS thread lifecycle.
- [[Syscall Handling (entersyscall & exitsyscall)]] — Detaching logical processor P from blocking OS thread M, retaking P, and thread parking on notesleep.
- [[Sysmon Daemon Implementation]] — Background monitoring thread: retaking stuck Ps from long syscalls, forcing periodic GC, and preemption.
- [[Non-Cooperative Async Preemption (preemptone & asyncPreempt)]] — Injecting SIGURG POSIX signal handler into running thread to force safe-point preemption.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
