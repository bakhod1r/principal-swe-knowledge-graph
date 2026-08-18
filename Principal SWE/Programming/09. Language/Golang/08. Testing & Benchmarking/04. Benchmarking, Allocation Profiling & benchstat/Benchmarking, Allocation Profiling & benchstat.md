---
title: Benchmarking, Allocation Profiling & benchstat
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Benchmarking, Allocation Profiling & benchstat

testing.B loop mechanics, allocation profiling (-benchmem), parallel benchmarks, benchstat, and compiler dead-code traps.

```text
Benchmarking, Allocation Profiling & benchstat
│
├── [[testing.B Microbenchmarks (b.N & Iteration Scaling)]]
├── [[Allocation Profiling in Benchmarks (-benchmem & b.ReportAllocs)]]
├── [[Parallel Benchmarks (b.RunParallel & PB)]]
├── [[Statistical Analysis with benchstat]]
└── [[Benchmarking Traps (Compiler Inlining & Dead Code Elimination)]]
```

---

## 🗂️ Topics

- [[testing.B Microbenchmarks (b.N & Iteration Scaling)]] — Loop mechanics, b.N scaling, b.ResetTimer(), and avoiding compiler loop optimization optimizations.
- [[Allocation Profiling in Benchmarks (-benchmem & b.ReportAllocs)]] — Tracking heap allocations per operation (allocs/op and B/op) to prevent performance regressions.
- [[Parallel Benchmarks (b.RunParallel & PB)]] — Measuring throughput under multi-threaded concurrency load across all logical CPU cores.
- [[Statistical Analysis with benchstat]] — Comparing benchmark results before and after code changes with p-value statistical confidence.
- [[Benchmarking Traps (Compiler Inlining & Dead Code Elimination)]] — Forcing benchmark results into package-level sinks (var Sink T) to prevent dead-code removal.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]

