---
title: Benchmarks, Coverage & Fuzzing
tags:
  - golang
  - testing
parent: "[[Testing & Benchmarking]]"
---

# Benchmarks, Coverage & Fuzzing

Performance benchmarks, memory allocation tracking, test coverage, fuzz testing, and race detector.

```text
Benchmarks, Coverage & Fuzzing
│
├── [[Benchmarks (testing.B)]]
├── [[benchstat Statistical Analysis]]
├── [[Coverage Analysis (-cover)]]
├── [[Fuzz Testing (testing.F)]]
└── [[Race Detector (-race)]]
```

---

## 🗂️ Topics

- [[Benchmarks (testing.B)]] — b.N iterations, b.ResetTimer(), b.ReportAllocs(), memory allocation tracking.
- [[benchstat Statistical Analysis]] — Comparing benchmark results before and after code changes with statistical confidence.
- [[Coverage Analysis (-cover)]] — go test -cover, generating HTML coverage profiles, enforcing coverage gates in CI.
- [[Fuzz Testing (testing.F)]] — Randomized input mutation engine for discovering edge-case crashes and vulnerabilities.
- [[Race Detector (-race)]] — ThreadSanitizer runtime data race detection during test runs.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]
- 🎓 Root: [[Principal SWE]]
