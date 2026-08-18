---
title: Observability & Logging
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Observability & Logging

Structured logging, distributed tracing, and metrics.

```text
Observability & Logging
│
├── [[Structured Logging (JSON, slog, Logfmt)]]
├── [[Distributed Tracing & Context Propagation (OpenTelemetry)]]
├── [[High-Cardinality Metrics (Counters, Gauges, Histograms)]]
└── [[Production Profiling Fleets & Continuous Profiling]]
```

---

## 🗂️ Topics

- [[Structured Logging (JSON, slog, Logfmt)]] — Building machine-readable, indexable, high-throughput log streams with contextual attributes.
- [[Distributed Tracing & Context Propagation (OpenTelemetry)]] — Propagating trace context (W3C Trace Context) across network boundaries for latency attribution.
- [[High-Cardinality Metrics (Counters, Gauges, Histograms)]] — Designing metric models, percentiles (p95, p99), and managing cardinality explosion.
- [[Production Profiling Fleets & Continuous Profiling]] — Running continuous low-overhead profiling across production clusters (Pyroscope, Parca).

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]
- 🎓 Root: [[Principal SWE]]
