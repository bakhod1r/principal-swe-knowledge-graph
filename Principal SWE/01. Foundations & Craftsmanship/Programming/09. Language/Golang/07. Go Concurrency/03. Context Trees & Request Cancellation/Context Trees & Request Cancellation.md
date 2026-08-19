---
title: Context Trees & Request Cancellation
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Go Concurrency]]"
---

# Context Trees & Request Cancellation

context package, cancellation trees, deadlines, timeouts, request-scoped values, and Go 1.21 additions.

```text
Context Trees & Request Cancellation
│
├── [[context.Background() vs context.TODO()]]
├── [[cancelCtx & Cancellation Tree Propagation]]
├── [[timerCtx & Deadline Scheduling (time.AfterFunc)]]
├── [[valueCtx & Key-Value Immutability]]
├── [[context.WithoutCancel (Go 1.21+)]]
├── [[context.AfterFunc (Go 1.21+)]]
├── [[Context Memory Leaks & Resource Hygiene]]
└── [[Context Design Rules]]
```

---

## 🗂️ Topics

- [[context.Background() vs context.TODO()]] — Root context initialization and placeholder context semantics.
- [[cancelCtx & Cancellation Tree Propagation]] — Parent-to-child cancellation propagation and child detach mechanics.
- [[timerCtx & Deadline Scheduling (time.AfterFunc)]] — Scheduling automatic cancellations via system timers and deadline expiration.
- [[valueCtx & Key-Value Immutability]] — Thread-safe immutable request-scoped value lookup and O(N) search depth caveats.
- [[context.WithoutCancel (Go 1.21+)]] — Detaching cancellation from parent context while preserving request values.
- [[context.AfterFunc (Go 1.21+)]] — Registering asynchronous callbacks executed when context is canceled.
- [[Context Memory Leaks & Resource Hygiene]] — Preventing goroutine and memory leaks by always deferring cancel() calls.
- [[Context Design Rules]] — Passing context as first parameter, avoiding context in structs, keeping values clean.

---

## 🔗 References
- ⬆️ Parent: `Concurrency & Synchronization`

