---
title: Core Architecture
tags:
  - golang
  - runtime
parent: "[[Runtime & Internals]]"
---

# Core Architecture

Runtime boot sequence, memory model, and package runtime interface functions.

```text
Core Architecture
│
├── [[Runtime Bootstrapping (rt0_go, schedinit)]]
├── [[sysmon Background Daemon Thread]]
└── [[runtime Package Diagnostic APIs]]
```

---

## 🗂️ Topics

- [[Runtime Bootstrapping (rt0_go, schedinit)]] — Entry point assembly, OS thread creation, runtime initialization, main goroutine start.
- [[sysmon Background Daemon Thread]] — Tick loop, retaking stuck Ps from long syscalls, forcing GC, preemption.
- [[runtime Package Diagnostic APIs]] — runtime.Gosched, runtime.LockOSThread, runtime.NumGoroutine, runtime.GC, runtime.ReadMemStats.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
