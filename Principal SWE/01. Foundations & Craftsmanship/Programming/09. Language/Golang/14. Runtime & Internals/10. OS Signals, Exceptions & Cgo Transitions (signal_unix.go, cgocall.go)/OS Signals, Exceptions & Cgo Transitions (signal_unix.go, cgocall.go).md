---
title: OS Signals, Exceptions & Cgo Transitions (signal_unix.go, cgocall.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# OS Signals, Exceptions & Cgo Transitions (signal_unix.go, cgocall.go)

Signal handling on gsignal stack, SIGSEGV recovery, SIGPROF profiler, cross-ABI cgocall transitions, and extraM threads.

```text
OS Signals, Exceptions & Cgo Transitions (signal_unix.go, cgocall.go)
│
├── [[Signal Handling Architecture & gsignal Stack (signal_unix.go)]]
├── [[SIGSEGV & SIGBUS Recovery to Runtime Panics]]
├── [[Profiling Signal Generation (SIGPROF Sampling Engine)]]
├── [[Cross-ABI Cgo Transition Mechanics (cgocall & cgocallback)]]
└── [[Extra OS Threads Management for Cgo (extraM)]]
```

---

## 🗂️ Topics

- [[Signal Handling Architecture & gsignal Stack (signal_unix.go)]] — Registering POSIX signals (initsig), dedicated signal stack (gsignal), and sighandler execution.
- [[SIGSEGV & SIGBUS Recovery to Runtime Panics]] — Catching hardware null-pointer dereferences and memory faults and converting them into catchable panics.
- [[Profiling Signal Generation (SIGPROF Sampling Engine)]] — OS interval timer delivering SIGPROF (100Hz) to capture thread program counters into trace buffers.
- [[Cross-ABI Cgo Transition Mechanics (cgocall & cgocallback)]] — Switching from Go stack to OS C thread stack, saving registers, calling C, and transitioning back.
- [[Extra OS Threads Management for Cgo (extraM)]] — Managing background OS threads (extraM) to handle incoming callbacks from C into Go.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]

