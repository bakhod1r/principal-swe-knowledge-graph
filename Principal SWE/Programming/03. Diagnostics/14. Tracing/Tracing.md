---
title: Tracing
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Tracing

Distributed tracing with W3C Trace Context, span lifecycle and links, in-process context propagation, and critical path latency waterfalls.

```text
Tracing
│
├── [[Distributed Tracing Mechanics and W3C Trace Context Standard]]
├── [[Span Lifecycle, Child Spans, and Span Links Architecture]]
├── [[In-Process Context Propagation and Thread-Local Storage (TLS)]]
└── [[Critical Path Latency Attribution and Waterfall Analysis]]
```

---

## 🗂️ Topics

- [[Distributed Tracing Mechanics and W3C Trace Context Standard]] — Propagating traceparent and tracestate headers across HTTP, gRPC, and message brokers.
- [[Span Lifecycle, Child Spans, and Span Links Architecture]] — Modeling asynchronous operations, batch jobs, and fan-out calls using OpenTelemetry span links.
- [[In-Process Context Propagation and Thread-Local Storage (TLS)]] — Propagating trace spans through asynchronous queues, thread pools, and goroutine switches.
- [[Critical Path Latency Attribution and Waterfall Analysis]] — Identifying the true bottlenecks across microservice dependency DAGs using critical path analysis.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]

