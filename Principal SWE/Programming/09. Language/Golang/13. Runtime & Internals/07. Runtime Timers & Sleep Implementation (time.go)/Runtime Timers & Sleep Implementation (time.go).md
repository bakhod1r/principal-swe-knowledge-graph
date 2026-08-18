---
title: Runtime Timers & Sleep Implementation (time.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Runtime Timers & Sleep Implementation (time.go)

Per-P timer heaps, timerproc execution, time.Sleep parking, and ticker cleanup lifecycle in runtime/time.go.

```text
Runtime Timers & Sleep Implementation (time.go)
│
├── [[4-Heap Timer Wheel Architecture & Per-P Timers]]
├── [[timerproc & Sysmon Timer Stealing Mechanics]]
├── [[time.Sleep & goparkunlock Mechanics]]
├── [[Timer Bucket Resets (time.Reset) & State Machine]]
└── [[Ticker Garbage Collection Pitfalls & Stop Cleanup]]
```

---

## 🗂️ Topics

- [[4-Heap Timer Wheel Architecture & Per-P Timers]] — How Go 1.14+ moved timers from a single central lock to lockless per-P heaps (p.timers).
- [[timerproc & Sysmon Timer Stealing Mechanics]] — How idle processors and sysmon steal expired timers from busy peers and fire callbacks.
- [[time.Sleep & goparkunlock Mechanics]] — Parking goroutines with waitReasonSleep and awakening via timer channel ready callbacks.
- [[Timer Bucket Resets (time.Reset) & State Machine]] — timerModifiedEarliest, timerDeleted, timerRunning state machine transitions in runtime/time.go.
- [[Ticker Garbage Collection Pitfalls & Stop Cleanup]] — Why unstopped tickers leak timers in per-P heaps until manual Stop() invocation.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]

