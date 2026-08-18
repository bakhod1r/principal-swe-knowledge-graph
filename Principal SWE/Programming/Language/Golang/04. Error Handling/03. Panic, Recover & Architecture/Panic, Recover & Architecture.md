---
title: Panic, Recover & Architecture
tags:
  - golang
  - error-handling
parent: "[[Error Handling]]"
---

# Panic, Recover & Architecture

Unrecoverable program states, deferred recover handlers, stack traces, and resilient error design.

```text
Panic, Recover & Architecture
│
├── [[panic Semantics]]
├── [[recover in Deferred Functions]]
├── [[Goroutine Panic Isolation]]
├── [[Stack Traces & runtime-debug]]
├── [[Domain vs Infrastructure Errors]]
├── [[Error Design Best Practices]]
└── [[Handle Errors, Don't Just Check]]
```

---

## 🗂️ Topics

- [[panic Semantics]] — Unwinding the goroutine call stack on unrecoverable conditions.
- [[recover in Deferred Functions]] — Catching runtime panics and restoring program execution safely.
- [[Goroutine Panic Isolation]] — Panics inside goroutines crash the entire process unless caught locally.
- [[Stack Traces & runtime-debug]] — Capturing and formatting stack traces for observability.
- [[Domain vs Infrastructure Errors]] — Separating business logic errors from database/network failures.
- [[Error Design Best Practices]] — Decorating errors without losing context, avoid string matching errors.
- [[Handle Errors, Don't Just Check]] — Meaningful error mitigation vs blindly propagating nil errors.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]
- 🎓 Root: [[Principal SWE]]
