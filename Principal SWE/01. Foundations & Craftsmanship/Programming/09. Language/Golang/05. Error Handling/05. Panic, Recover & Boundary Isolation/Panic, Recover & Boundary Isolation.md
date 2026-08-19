---
title: Panic, Recover & Boundary Isolation
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Error Handling (Clean Code)]]"
---

# Panic, Recover & Boundary Isolation

Panic vs error decision tree, recover boundary middleware, goroutine orphan panics, error conversion, and stack frame dumps.

```text
Panic, Recover & Boundary Isolation
│
├── [[Panic vs Error Decision Tree (Exceptional vs Expected)]]
├── [[Recover Boundary Middleware in HTTP & gRPC Servers]]
├── [[Goroutine Panic Isolation Hazard (Orphan Panics)]]
├── [[Converting Panics to Clean Domain Errors]]
└── [[Structured Stack Frame Dumps during Panic Recovery]]
```

---

## 🗂️ Topics

- [[Panic vs Error Decision Tree (Exceptional vs Expected)]] — Staff-level heuristics: when is panic acceptable (programmer bug, startup failure) vs error.
- [[Recover Boundary Middleware in HTTP & gRPC Servers]] — Capturing uncaught goroutine panics, logging stack traces, and returning 500 Internal Server Error.
- [[Goroutine Panic Isolation Hazard (Orphan Panics)]] — Why a panic in a spawned background goroutine terminates the entire process if unrecovered inside itself.
- [[Converting Panics to Clean Domain Errors]] — Safely capturing third-party panics inside library boundaries and translating them to structured errors.
- [[Structured Stack Frame Dumps during Panic Recovery]] — Using debug.Stack() to capture and sanitize full multi-goroutine traces upon panic.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling (Clean Code)]]

