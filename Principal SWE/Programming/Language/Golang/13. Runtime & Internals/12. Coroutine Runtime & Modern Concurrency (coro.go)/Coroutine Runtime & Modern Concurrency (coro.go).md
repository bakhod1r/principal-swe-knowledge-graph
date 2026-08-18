---
title: Coroutine Runtime & Modern Concurrency (coro.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Coroutine Runtime & Modern Concurrency (coro.go)

Go 1.23+ coroutine engine, corostart, coroswitch, GMP scheduler integration, and zero-allocation iterator bridges.

```text
Coroutine Runtime & Modern Concurrency (coro.go)
│
├── [[Coroutine Stack-Switching Engine (coro.go Go 1.23+)]]
├── [[corostart & coroswitch Implementation Details]]
├── [[Integrating Coroutines with GMP Scheduler]]
└── [[Zero-Allocation Push-Pull Iterator Runtime Bridges]]
```

---

## 🗂️ Topics

- [[Coroutine Stack-Switching Engine (coro.go Go 1.23+)]] — Go 1.23+ lightweight asymmetric coroutine runtime implementation powering iter.Pull.
- [[corostart & coroswitch Implementation Details]] — Creating paired coroutine execution contexts and switching execution control between caller and callee.
- [[Integrating Coroutines with GMP Scheduler]] — How coroutines cooperate with goroutines, stack copying, and preemptible scheduler safepoints.
- [[Zero-Allocation Push-Pull Iterator Runtime Bridges]] — Transforming push-based yield functions into stateful pull iterators without heap allocations.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
