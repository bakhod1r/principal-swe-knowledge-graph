---
title: Observability & Runtime Introspection
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Golang]]"
---

# 📊 Observability & Runtime Introspection

Production monitoring with runtime/metrics, expvar, execution tracing, OpenTelemetry Go SDK integration, and GODEBUG diagnostics.

```text
Observability & Runtime Introspection
│
├── [[Metrics & Diagnostics|01. Metrics & Diagnostics]]
│   ├── [[runtime-metrics Structured Metrics]]
│   ├── [[expvar Public HTTP Metrics]]
│   ├── [[GODEBUG Runtime Diagnostic Flags]]
│   └── [[runtime-debug Build Info & Heap Dumps]]
└── [[Tracing & Telemetry|02. Tracing & Telemetry]]
│   ├── [[runtime-trace Application Tracing]]
│   ├── [[OpenTelemetry Go SDK (Traces, Metrics, Logs)]]
│   └── [[OTLP Exporter Integration & Span Propagation]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Metrics & Diagnostics|01. Metrics & Diagnostics]]
- [[runtime-metrics Structured Metrics]] — Modern structured /sched/latencies, /memory/classes, /gc/pauses metrics inspection.
- [[expvar Public HTTP Metrics]] — Publishing public HTTP JSON metrics endpoints for standard Go counters and custom stats.
- [[GODEBUG Runtime Diagnostic Flags]] — gctrace=1, schedtrace=1000, madvdontneed=1, asyncpreemptoff=1 in production.
- [[runtime-debug Build Info & Heap Dumps]] — debug.ReadBuildInfo(), debug.WriteHeapDump(), debug.SetMemoryLimit() in code.
### 2. 📂 [[Tracing & Telemetry|02. Tracing & Telemetry]]
- [[runtime-trace Application Tracing]] — runtime/trace User Tasks, Regions, and Logs for in-depth execution timeline analysis.
- [[OpenTelemetry Go SDK (Traces, Metrics, Logs)]] — Instrumenting Go microservices with OpenTelemetry Tracers, Meters, and OTLP exporters.
- [[OTLP Exporter Integration & Span Propagation]] — Propagating trace context across HTTP/gRPC boundaries and exporting to Jaeger/Tempo.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

