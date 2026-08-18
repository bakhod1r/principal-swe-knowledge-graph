---
title: Distributed Tracing & OpenTelemetry (OTel)
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Observability & Runtime Introspection]]"
---

# Distributed Tracing & OpenTelemetry (OTel)

OpenTelemetry Go SDK, distributed context propagation (W3C), OTLP exporters, auto-instrumentation, and span lifecycles.

```text
Distributed Tracing & OpenTelemetry (OTel)
│
├── [[OpenTelemetry Go SDK Architecture (TracerProvider & Spans)]]
├── [[Distributed Context Propagation (W3C TraceContext & B3)]]
├── [[OTLP Exporters (gRPC & HTTP Protocols)]]
├── [[Database & HTTP Middleware Instrumentation]]
└── [[Span Events, Status Codes & Error Recording]]
```

---

## 🗂️ Topics

- [[OpenTelemetry Go SDK Architecture (TracerProvider & Spans)]] — Initializing OpenTelemetry TracerProvider, Samplers, Span processors, and resource detectors.
- [[Distributed Context Propagation (W3C TraceContext & B3)]] — Injecting and extracting trace IDs across HTTP (propagation.TraceContext) and gRPC metadata.
- [[OTLP Exporters (gRPC & HTTP Protocols)]] — Exporting telemetry streams to OpenTelemetry Collector, Jaeger, Tempo, and Honeycomb.
- [[Database & HTTP Middleware Instrumentation]] — Auto-instrumenting net/http handlers and SQL drivers (otelhttp, otelsql) with parent span binding.
- [[Span Events, Status Codes & Error Recording]] — span.RecordError(), span.SetStatus(), and attaching structured semantic attributes to spans.

---

## 🔗 References
- ⬆️ Parent: [[Observability & Runtime Introspection]]

