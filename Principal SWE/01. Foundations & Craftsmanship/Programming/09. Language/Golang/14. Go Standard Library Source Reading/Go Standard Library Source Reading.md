---
title: Go Standard Library Source Reading
tags:
  - golang
  - source-reading
  - principal-swe
parent: "[[Golang]]"
---

# 📖 Go Standard Library Source Reading

Architectural code walkthroughs of production-proven Go packages: net/http, sync, context, database/sql, and encoding/json.

```text
Go Standard Library Source Reading
│
├── `01. Core & Concurrency Source`
│   ├── `sync.Mutex Source Walkthrough`
│   ├── `sync.WaitGroup and sync.Once Source Walkthrough`
│   ├── `sync.Pool & sync.Map Source Walkthrough`
│   ├── `context.Context Tree Source Walkthrough`
│   └── `Channel Implementation Source (chan.go)`
└── [[Networking & Data Source|02. Networking & Data Source]]
│   ├── `net-http Server.Serve & Transport Source`
│   ├── `database-sql Connection Pool Source`
│   └── `encoding-json Scanner & Encoder Source`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 `01. Core & Concurrency Source`
- `sync.Mutex Source Walkthrough` — Dissecting sync.Mutex fast-path/slow-path starvation, normal vs starvation mode.
- `sync.WaitGroup and sync.Once Source Walkthrough` — Atomic state bitpacking in WaitGroup, double-checked atomic Once.
- `sync.Pool & sync.Map Source Walkthrough` — Dissecting lockless atomic loads and per-P pool caches.
- `context.Context Tree Source Walkthrough` — emptyCtx, cancelCtx tree propagation, timerCtx deadlines, valueCtx lookup.
- `Channel Implementation Source (chan.go)` — makechan, chansend, chanrecv, closechan, direct copy optimizations.
### 2. 📂 [[Networking & Data Source|02. Networking & Data Source]]
- `net-http Server.Serve & Transport Source` — Server.Serve loop, conn.serve, Transport connection pooling, RoundTripper interface.
- `database-sql Connection Pool Source` — DB connection pool management, driver interface, tx isolation level handling.
- `encoding-json Scanner & Encoder Source` — Reflect-based encoder compilation cache, state machine scanner, stream buffer reuse.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Contents

- [[Core & Concurrency Source]]
- [[Networking & Data Source]]
