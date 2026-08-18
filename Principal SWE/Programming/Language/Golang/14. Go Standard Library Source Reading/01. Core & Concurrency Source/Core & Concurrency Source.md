---
title: Core & Concurrency Source
tags:
  - golang
  - source-reading
parent: "[[Go Standard Library Source Reading]]"
---

# Core & Concurrency Source

Source code deep dives: sync.Mutex, sync.WaitGroup, sync.Once, context, and channels.

```text
Core & Concurrency Source
│
├── [[sync.Mutex Source Walkthrough]]
├── [[sync.WaitGroup and sync.Once Source Walkthrough]]
├── [[context.Context Tree Source Walkthrough]]
└── [[Channel Implementation Source (chan.go)]]
```

---

## 🗂️ Topics

- [[sync.Mutex Source Walkthrough]] — Dissecting sync.Mutex fast-path/slow-path starvation, normal vs starvation mode.
- [[sync.WaitGroup and sync.Once Source Walkthrough]] — Atomic state bitpacking in WaitGroup, double-checked atomic Once.
- [[context.Context Tree Source Walkthrough]] — emptyCtx, cancelCtx tree propagation, timerCtx deadlines, valueCtx lookup.
- [[Channel Implementation Source (chan.go)]] — makechan, chansend, chanrecv, closechan, direct copy optimizations.

---

## 🔗 References
- ⬆️ Parent: [[Go Standard Library Source Reading]]
- 🎓 Root: [[Principal SWE]]
