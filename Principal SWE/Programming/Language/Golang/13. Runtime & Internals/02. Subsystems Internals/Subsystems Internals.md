---
title: Subsystems Internals
tags:
  - golang
  - runtime
parent: "[[Runtime & Internals]]"
---

# Subsystems Internals

GMP scheduler, tricolor GC, TCMalloc memory allocator, and panic/recover implementation.

```text
Subsystems Internals
│
├── [[GMP Scheduler Implementation (proc.go)]]
├── [[Garbage Collector Implementation (mgc.go)]]
├── [[Memory Allocator Implementation (malloc.go)]]
└── [[Panic and Recover Implementation (panic.go)]]
```

---

## 🗂️ Topics

- [[GMP Scheduler Implementation (proc.go)]] — findRunnable, schedule loop, work stealing algorithm, sysmon preemption.
- [[Garbage Collector Implementation (mgc.go)]] — Tricolor marking, hybrid write barrier, concurrent sweep, GC pacer math.
- [[Memory Allocator Implementation (malloc.go)]] — mcache span caching, mcentral size classes, mheap arenas, huge pages.
- [[Panic and Recover Implementation (panic.go)]] — _panic and _defer linked list unwinding, gopanic, gorecover.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
