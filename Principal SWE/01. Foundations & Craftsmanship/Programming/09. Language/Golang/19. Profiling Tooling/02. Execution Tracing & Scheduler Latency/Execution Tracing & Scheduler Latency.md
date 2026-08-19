---
title: Execution Tracing & Scheduler Latency
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Profiling Tooling]]"
---

# Execution Tracing & Scheduler Latency

Nanosecond execution tracing with go tool trace, GC pause analysis, syscall blocking, and flight recording.

```text
Execution Tracing & Scheduler Latency
│
├── [[Execution Tracer Architecture (go tool trace)]]
├── [[Diagnosing GC Pauses & STW Latency]]
├── [[Diagnosing Network & Syscall Blocking]]
├── [[User-Defined Tasks & Regions (runtime-trace)]]
└── [[Flight Recording & Continuous Trace Dumps]]
```

---

## 🗂️ Topics

- [[Execution Tracer Architecture (go tool trace)]] — Nanosecond-level execution timeline recording: event log buffers, per-P trace buffers, and trace format.
- [[Diagnosing GC Pauses & STW Latency]] — Correlating GC sweep and mark phases with application p99/p999 latency spikes.
- [[Diagnosing Network & Syscall Blocking]] — Tracking goroutines parked on Netpoller and OS thread syscall handoffs.
- [[User-Defined Tasks & Regions (runtime-trace)]] — Instrumenting domain business logic with trace.WithRegion and trace.Log for execution tracing.
- [[Flight Recording & Continuous Trace Dumps]] — Capturing circular in-memory flight recording ring buffers on error triggers.

---

## 🔗 References
- ⬆️ Parent: `Performance Engineering & Profiling`

