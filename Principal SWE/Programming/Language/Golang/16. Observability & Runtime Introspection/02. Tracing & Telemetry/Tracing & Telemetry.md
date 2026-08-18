- [[eBPF Observability for Go (Cilium, Pixie)]] — Zero-instrumentation distributed tracing and network inspection using kernel uprobes and kprobes.

- [[Continuous Memory Profiling & OOM Safety]] — Triggering automatic heap dumps before container Linux OOM kills.

---
title: Tracing & Telemetry
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Observability & Runtime Introspection]]"
---

# Tracing & Telemetry

Runtime tracing integration, OpenTelemetry distributed tracing and span propagation.

```text
Tracing & Telemetry
│
├── [[runtime-trace Application Tracing]]
├── [[OpenTelemetry Go SDK (Traces, Metrics, Logs)]]
└── [[OTLP Exporter Integration & Span Propagation]]
```

---

## 🗂️ Topics

- [[runtime-trace Application Tracing]] — runtime/trace User Tasks, Regions, and Logs for in-depth execution timeline analysis.
- [[OpenTelemetry Go SDK (Traces, Metrics, Logs)]] — Instrumenting Go microservices with OpenTelemetry Tracers, Meters, and OTLP exporters.
- [[OTLP Exporter Integration & Span Propagation]] — Propagating trace context across HTTP/gRPC boundaries and exporting to Jaeger/Tempo.

---

## 🔗 References
- ⬆️ Parent: [[Observability & Runtime Introspection]]
- 🎓 Root: [[Principal SWE]]
