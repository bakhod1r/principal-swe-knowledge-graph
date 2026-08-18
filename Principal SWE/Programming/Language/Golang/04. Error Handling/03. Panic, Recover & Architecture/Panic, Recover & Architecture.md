- [[Structured Domain Error Codes (gRPC & RFC 7807)]] — Mapping internal Go errors to gRPC status codes and HTTP Problem Details.

- [[Error Allocation Cost & Stack Trace Overhead]] — Performance impact of error struct allocations vs sentinel errors in hot paths.

---
title: Panic, Recover & Architecture
tags:
  - golang
  - error-handling
  - principal-swe
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
├── [[Error Classification (Transient vs Permanent)]]
├── [[Error Design Best Practices]]
└── [[Handle Errors, Don't Just Check]]
```

---

## 🗂️ Topics

- [[panic Semantics]] — Unwinding the goroutine call stack on fatal, unrecoverable programmer errors.
- [[recover in Deferred Functions]] — Safely intercepting panics in deferred functions and converting to errors.
- [[Goroutine Panic Isolation]] — Uncaught panics inside spawned goroutines terminate the entire process.
- [[Stack Traces & runtime-debug]] — Capturing, parsing, and logging panic stack traces for observability.
- [[Domain vs Infrastructure Errors]] — Architectural separation of business rule errors vs database/network errors.
- [[Error Classification (Transient vs Permanent)]] — Categorizing errors for intelligent retries, circuit breaking, and alerting.
- [[Error Design Best Practices]] — Enriching errors without losing original context, avoiding string matching.
- [[Handle Errors, Don't Just Check]] — Meaningful error recovery and remediation vs blind error return propagation.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling]]
- 🎓 Root: [[Principal SWE]]
