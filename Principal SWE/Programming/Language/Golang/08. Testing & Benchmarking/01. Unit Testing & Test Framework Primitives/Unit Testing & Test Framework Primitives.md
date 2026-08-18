---
title: Unit Testing & Test Framework Primitives
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Unit Testing & Test Framework Primitives

testing.T primitives, table-driven tests, subtests (t.Run), test helpers (t.Helper), parallel testing, and resource cleanups.

```text
Unit Testing & Test Framework Primitives
│
├── [[testing.T Lifecycle & Fatal vs Error Reporting]]
├── [[Table-Driven Tests Architecture (Idiomatic Pattern)]]
├── [[Subtests (t.Run) Hierarchical Execution & Filtering]]
├── [[Test Helpers (t.Helper) Stack Frame Stripping]]
├── [[Parallel Tests (t.Parallel) Race Isolation]]
├── [[t.Cleanup Reliable Resource Management]]
├── [[TestMain Function Lifecycle (Global Setup & Teardown)]]
└── [[t.Setenv Environment Isolation]]
```

---

## 🗂️ Topics

- [[testing.T Lifecycle & Fatal vs Error Reporting]] — t.Fail, t.FailNow, t.Fatal, t.Error, and t.Log execution semantics and failure propagation.
- [[Table-Driven Tests Architecture (Idiomatic Pattern)]] — Slices of anonymous structs, subtest encapsulation, and clear descriptive assertion failures.
- [[Subtests (t.Run) Hierarchical Execution & Filtering]] — Running specific subtests via regex (-run=TestRoot/SubCase) and test failure isolation.
- [[Test Helpers (t.Helper) Stack Frame Stripping]] — Marking helper functions to strip wrapper lines from assertion failure stack traces.
- [[Parallel Tests (t.Parallel) Race Isolation]] — Pausing and resuming tests concurrently to catch race conditions and shorten CI wall time.
- [[t.Cleanup Reliable Resource Management]] — Registering reliable LIFO cleanup callbacks replacing fragile defer statements in tests.
- [[TestMain Function Lifecycle (Global Setup & Teardown)]] — Global test suite setup, teardown hooks, flag parsing (flag.Parse), and m.Run() execution.
- [[t.Setenv Environment Isolation]] — Safely mutating environment variables per-test with automatic post-test restoration.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]

