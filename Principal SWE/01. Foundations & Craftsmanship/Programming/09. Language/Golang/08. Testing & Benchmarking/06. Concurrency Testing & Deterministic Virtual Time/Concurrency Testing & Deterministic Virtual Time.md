---
title: Concurrency Testing & Deterministic Virtual Time
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Concurrency Testing & Deterministic Virtual Time

Race detector (-race), deterministic virtual time (synctest Go 1.24+), goroutine leak detection (goleak), and coverage gates.

```text
Concurrency Testing & Deterministic Virtual Time
│
├── [[Race Detector (-race ThreadSanitizer Mechanics)]]
├── [[Deterministic Virtual Time Testing (synctest Go 1.24+)]]
├── [[Testing Goroutine Leaks (goleak)]]
├── [[Deadlock & Starvation Testing in Concurrent Code]]
└── [[Code Coverage Profiling & Quality Gates (-cover, -coverprofile)]]
```

---

## 🗂️ Topics

- [[Race Detector (-race ThreadSanitizer Mechanics)]] — Runtime data race detection, 10x CPU / 20x memory overhead, and CI pipeline integration.
- [[Deterministic Virtual Time Testing (synctest Go 1.24+)]] — Testing concurrent goroutines and timers instantly without time.Sleep using virtual time bubbles.
- [[Testing Goroutine Leaks (goleak)]] — Uber goleak library verifying that no background goroutines leaked after test suite execution.
- [[Deadlock & Starvation Testing in Concurrent Code]] — Simulating high lock contention and verifying timeout-based deadlock recovery in tests.
- [[Code Coverage Profiling & Quality Gates (-cover, -coverprofile)]] — Generating HTML coverage reports, filtering generated code, and enforcing CI coverage thresholds.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]

