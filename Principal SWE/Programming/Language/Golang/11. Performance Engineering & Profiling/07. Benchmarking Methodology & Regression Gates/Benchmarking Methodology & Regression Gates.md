---
title: Benchmarking Methodology & Regression Gates
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Performance Engineering & Profiling]]"
---

# Benchmarking Methodology & Regression Gates

Statistically valid benchmarking with testing.B, benchstat analysis, environment isolation, and CI gates.

```text
Benchmarking Methodology & Regression Gates
│
├── [[Statistically Valid Benchmarking (testing.B)]]
├── [[Statistical Significance Testing with benchstat]]
├── [[Isolating Benchmarking Environments]]
└── [[Automated Benchmark Regression Gates in CI]]
```

---

## 🗂️ Topics

- [[Statistically Valid Benchmarking (testing.B)]] — b.N, b.ResetTimer(), b.ReportAllocs(), and avoiding compiler loop elimination.
- [[Statistical Significance Testing with benchstat]] — Comparing before-and-after benchmark results with p-value statistical confidence.
- [[Isolating Benchmarking Environments]] — CPU frequency governor locking, core affinity pinning (taskset), and disabling turbo boost.
- [[Automated Benchmark Regression Gates in CI]] — Enforcing maximum allocation count and latency regression thresholds on pull requests.

---

## 🔗 References
- ⬆️ Parent: [[Performance Engineering & Profiling]]

