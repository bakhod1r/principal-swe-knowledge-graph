---
title: Continuous Profiling
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Continuous Profiling

Always-on fleet-wide profiling (Pyroscope, Parca), flame graph navigation, off-CPU analysis, and eBPF sampling.

```text
Continuous Profiling
│
├── [[Continuous Fleet-Wide Production Profiling (Pyroscope, Parca)]]
├── [[Navigating and Diffing Flame Graphs across Production Deployments]]
├── [[Off-CPU and Blocking Profiling at Scale]]
└── [[eBPF-Based Kernel and User-Space Continuous Sampling]]
```

---

## 🗂️ Topics

- [[Continuous Fleet-Wide Production Profiling (Pyroscope, Parca)]] — Always-on low-overhead (<1% CPU) profiling across thousands of production microservices.
- [[Navigating and Diffing Flame Graphs across Production Deployments]] — Identifying subtle CPU and memory regressions between git commits via differential flame graphs.
- [[Off-CPU and Blocking Profiling at Scale]] — Measuring thread parking, lock contention, and I/O wait times in production services.
- [[eBPF-Based Kernel and User-Space Continuous Sampling]] — Native stack walking across compiled binaries and kernel syscalls without runtime-specific agents.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]
- 🎓 Root: [[Principal SWE]]
