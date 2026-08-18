- [[runtime-mgc.go Deep Source Walkthrough]] — Line-by-line architectural breakdown of GC marking phases and hybrid write barriers in mgc.go.

---
title: Core & Concurrency Source
tags:
  - golang
  - source-reading
  - principal-swe
parent: "[[Go Standard Library Source Reading]]"
---

# Core & Concurrency Source

Analyzing the source code of sync, context, and runtime primitives.

```text
Core & Concurrency Source
│
├── [[sync.Mutex Source Walkthrough]]
├── [[sync.WaitGroup and sync.Once Source Walkthrough]]
├── [[sync.Pool & sync.Map Source Walkthrough]]
├── [[context.Context Tree Source Walkthrough]]
└── [[Channel Implementation Source (chan.go)]]
```

---

## 🗂️ Topics

- [[sync.Mutex Source Walkthrough]] — Dissecting sync.Mutex fast-path/slow-path starvation, normal vs starvation mode.
- [[sync.WaitGroup and sync.Once Source Walkthrough]] — Atomic state bitpacking in WaitGroup, double-checked atomic Once.
- [[sync.Pool & sync.Map Source Walkthrough]] — Dissecting lockless atomic loads and per-P pool caches.
- [[context.Context Tree Source Walkthrough]] — emptyCtx, cancelCtx tree propagation, timerCtx deadlines, valueCtx lookup.
- [[Channel Implementation Source (chan.go)]] — makechan, chansend, chanrecv, closechan, direct copy optimizations.

---

## 🔗 References
- ⬆️ Parent: [[Go Standard Library Source Reading]]

