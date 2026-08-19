---
title: Runtime Metrics & Health Inspection
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Observability & Runtime Introspection]]"
---

# Runtime Metrics & Health Inspection

runtime/metrics framework, expvar public endpoints, Prometheus client integration, GC metrics, and Kubernetes health probes.

```text
Runtime Metrics & Health Inspection
│
├── [[runtime-metrics API Deep Dive]]
├── [[expvar Package & Public HTTP Telemetry]]
├── [[Prometheus Go Client Integration (client_golang)]]
├── [[Memory & GC Metrics Interpretation]]
├── [[Scheduler & Goroutine Contention Metrics]]
└── [[Kubernetes Health Probes (Liveness, Readiness, Startup)]]
```

---

## 🗂️ Topics

- [[runtime-metrics API Deep Dive]] — Structured runtime metrics: /sched/latencies:seconds, /gc/pauses:seconds, /memory/classes/heap:bytes.
- [[expvar Package & Public HTTP Telemetry]] — Publishing public JSON metrics endpoints with standard counters, maps, and custom stats.
- [[Prometheus Go Client Integration (client_golang)]] — Gauge, Counter, Histogram, Summary, collector registration, and scrapers.
- [[Memory & GC Metrics Interpretation]] — Correlating heap in-use, heap sys, GC cycle counts, and GC CPU utilization percentages.
- [[Scheduler & Goroutine Contention Metrics]] — Monitoring runnable goroutines waiting on logical processor queues and starvation.
- [[Kubernetes Health Probes (Liveness, Readiness, Startup)]] — Designing resilient, non-blocking health check endpoints with timeout isolation.

---

## 🔗 References
- ⬆️ Parent: [[Observability & Runtime Introspection]]

