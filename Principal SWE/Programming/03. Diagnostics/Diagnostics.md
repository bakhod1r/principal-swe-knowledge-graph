---
title: Diagnostics
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Programming]]"
---

# 💻 Diagnostics

Debuggers, DWARF tables, CPU/memory flame graphs, off-CPU profiling, and distributed tracing.

```text
Diagnostics
│
├── [[Interactive Debugging & Tracing|01. Interactive Debugging & Tracing]]
│   ├── [[Debuggers Mechanics (Breakpoints, Traps, DWARF Tables)]]
│   ├── [[Time-Travel Debugging & Deterministic Record-Replay]]
│   ├── [[Dynamic Binary Instrumentation (eBPF, DTrace)]]
│   └── [[Core Dump Forensics & Post-Mortem Crash Analysis]]
├── [[Profiling & Performance Analysis|02. Profiling & Performance Analysis]]
│   ├── [[Sampling vs Instrumentation Profilers]]
│   ├── [[CPU Profiling & Flame Graphs Interpretation]]
│   ├── [[Memory Profiling, Allocations & Leak Detection]]
│   └── [[Off-CPU Analysis & Lock Contention Profiling]]
└── [[Observability & Logging|03. Observability & Logging]]
│   ├── [[Structured Logging (JSON, slog, Logfmt)]]
│   ├── [[Distributed Tracing & Context Propagation (OpenTelemetry)]]
│   ├── [[High-Cardinality Metrics (Counters, Gauges, Histograms)]]
│   └── [[Production Profiling Fleets & Continuous Profiling]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Interactive Debugging & Tracing|01. Interactive Debugging & Tracing]]
- [[Debuggers Mechanics (Breakpoints, Traps, DWARF Tables)]] — How debuggers inject INT 3 traps, intercept ptrace signals, and resolve DWARF symbol tables.
- [[Time-Travel Debugging & Deterministic Record-Replay]] — Recording execution snapshots and replaying instruction streams backward to isolate non-deterministic bugs.
- [[Dynamic Binary Instrumentation (eBPF, DTrace)]] — Zero-overhead live kernel and user-space tracing without recompilation using eBPF probes.
- [[Core Dump Forensics & Post-Mortem Crash Analysis]] — Extracting memory state, register dumps, and stack traces from OS core dumps after process crashes.
### 2. 📂 [[Profiling & Performance Analysis|02. Profiling & Performance Analysis]]
- [[Sampling vs Instrumentation Profilers]] — Understanding overhead trade-offs: statistical timer-based sampling vs intrusive function-wrapping.
- [[CPU Profiling & Flame Graphs Interpretation]] — Generating, reading, and navigating Brendan Gregg flame graphs to identify hot code paths.
- [[Memory Profiling, Allocations & Leak Detection]] — Tracking heap growth, allocation rates, object retention trees, and finding memory leaks.
- [[Off-CPU Analysis & Lock Contention Profiling]] — Measuring time spent sleeping, waiting for mutexes, blocking on I/O, and thread context switching.
### 3. 📂 [[Observability & Logging|03. Observability & Logging]]
- [[Structured Logging (JSON, slog, Logfmt)]] — Building machine-readable, indexable, high-throughput log streams with contextual attributes.
- [[Distributed Tracing & Context Propagation (OpenTelemetry)]] — Propagating trace context (W3C Trace Context) across network boundaries for latency attribution.
- [[High-Cardinality Metrics (Counters, Gauges, Histograms)]] — Designing metric models, percentiles (p95, p99), and managing cardinality explosion.
- [[Production Profiling Fleets & Continuous Profiling]] — Running continuous low-overhead profiling across production clusters (Pyroscope, Parca).

---

## 🔗 Navigation
- ⬆️ Parent: [[Programming]]
- 🎓 Root: [[Principal SWE]]
