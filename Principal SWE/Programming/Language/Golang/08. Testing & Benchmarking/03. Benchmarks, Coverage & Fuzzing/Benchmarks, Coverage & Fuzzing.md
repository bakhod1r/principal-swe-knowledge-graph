- [[Benchmarking Allocation Tracing (-benchmem)]] — Tracking allocs-op and B-op in CI pipelines to prevent performance regressions.

- [[Mutation Testing in Go (go-mutesting)]] — Testing the quality and coverage of test suites by injecting AST code mutations.

---
title: Benchmarks, Coverage & Fuzzing
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Benchmarks, Coverage & Fuzzing

Performance benchmarks, memory allocation tracking, test coverage, fuzz testing, and race detector.

```text
Benchmarks, Coverage & Fuzzing
│
├── [[Benchmarks (testing.B)]]
├── [[Parallel Benchmarks (b.RunParallel)]]
├── [[benchstat Statistical Analysis]]
├── [[Coverage Analysis (-cover)]]
├── [[Fuzz Testing (testing.F)]]
└── [[Race Detector (-race)]]
```

---

## 🗂️ Topics

- [[Benchmarks (testing.B)]] — b.N iterations, b.ResetTimer(), b.ReportAllocs(), memory allocation tracking.
- [[Parallel Benchmarks (b.RunParallel)]] — Testing concurrent throughput under multi-threaded load.
- [[benchstat Statistical Analysis]] — Comparing benchmark results before and after code changes with statistical confidence.
- [[Coverage Analysis (-cover)]] — go test -cover, generating HTML coverage profiles, enforcing coverage gates in CI.
- [[Fuzz Testing (testing.F)]] — Randomized input mutation engine for discovering edge-case crashes and vulnerabilities.
- [[Race Detector (-race)]] — ThreadSanitizer runtime data race detection during test runs.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]
- 🎓 Root: [[Principal SWE]]
