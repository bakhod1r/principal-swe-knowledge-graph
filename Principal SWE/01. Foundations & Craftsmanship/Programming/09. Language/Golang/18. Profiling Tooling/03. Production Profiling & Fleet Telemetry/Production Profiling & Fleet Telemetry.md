---
title: Production Profiling & Fleet Telemetry
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Profiling Tooling]]"
---

# Production Profiling & Fleet Telemetry

Continuous profiling fleets, off-CPU analysis with eBPF, proactive OOM heap dumps, and runtime diagnostics.

```text
Production Profiling & Fleet Telemetry
│
├── [[Continuous Profiling Fleets (Pyroscope, Parca)]]
├── [[Off-CPU Analysis with eBPF in Go]]
├── [[Proactive Heap Dumps Before Container OOM Kills]]
└── [[GODEBUG Diagnostic Flags in Production]]
```

---

## 🗂️ Topics

- [[Continuous Profiling Fleets (Pyroscope, Parca)]] — Ultra-low overhead (<1% CPU) continuous fleet-wide profiling architecture.
- [[Off-CPU Analysis with eBPF in Go]] — Measuring time spent off-CPU waiting on kernel locks, disk I/O, and page faults using eBPF uprobes.
- [[Proactive Heap Dumps Before Container OOM Kills]] — Using debug.SetMemoryLimit and memory watchdogs to dump heaps before SIGKILL.
- [[GODEBUG Diagnostic Flags in Production]] — gctrace=1, schedtrace=1000, madvdontneed=1 for real-time runtime diagnostics.

---

## 🔗 References
- ⬆️ Parent: `Performance Engineering & Profiling`

