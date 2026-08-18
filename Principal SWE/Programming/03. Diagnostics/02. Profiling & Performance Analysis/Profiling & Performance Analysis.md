---
title: Profiling & Performance Analysis
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Profiling & Performance Analysis

CPU sampling, heap graphs, and lock contention analysis.

```text
Profiling & Performance Analysis
│
├── [[Sampling vs Instrumentation Profilers]]
├── [[CPU Profiling & Flame Graphs Interpretation]]
├── [[Memory Profiling, Allocations & Leak Detection]]
└── [[Off-CPU Analysis & Lock Contention Profiling]]
```

---

## 🗂️ Topics

- [[Sampling vs Instrumentation Profilers]] — Understanding overhead trade-offs: statistical timer-based sampling vs intrusive function-wrapping.
- [[CPU Profiling & Flame Graphs Interpretation]] — Generating, reading, and navigating Brendan Gregg flame graphs to identify hot code paths.
- [[Memory Profiling, Allocations & Leak Detection]] — Tracking heap growth, allocation rates, object retention trees, and finding memory leaks.
- [[Off-CPU Analysis & Lock Contention Profiling]] — Measuring time spent sleeping, waiting for mutexes, blocking on I/O, and thread context switching.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]
- 🎓 Root: [[Principal SWE]]
