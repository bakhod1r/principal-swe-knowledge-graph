---
title: Observability & Runtime Introspection
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Golang]]"
---

# 📊 Observability & Runtime Introspection

Cloud-native observability in Go: runtime/metrics, OpenTelemetry distributed tracing, structured logging (slog), Linux eBPF kernel inspection, and GODEBUG diagnostics.

```text
Observability & Runtime Introspection
│
├── [[Runtime Metrics & Health Inspection|01. Runtime Metrics & Health Inspection]]
│   ├── [[runtime-metrics API Deep Dive]]
│   ├── [[expvar Package & Public HTTP Telemetry]]
│   ├── [[Prometheus Go Client Integration (client_golang)]]
│   ├── [[Memory & GC Metrics Interpretation]]
│   ├── [[Scheduler & Goroutine Contention Metrics]]
│   └── [[Kubernetes Health Probes (Liveness, Readiness, Startup)]]
├── [[Distributed Tracing & OpenTelemetry (OTel)|02. Distributed Tracing & OpenTelemetry (OTel)]]
│   ├── [[OpenTelemetry Go SDK Architecture (TracerProvider & Spans)]]
│   ├── [[Distributed Context Propagation (W3C TraceContext & B3)]]
│   ├── [[OTLP Exporters (gRPC & HTTP Protocols)]]
│   ├── [[Database & HTTP Middleware Instrumentation]]
│   └── [[Span Events, Status Codes & Error Recording]]
├── [[Structured Logging Architecture (slog)|03. Structured Logging Architecture (slog)]]
│   ├── [[log-slog Core Architecture (Logger, Handler, Record)]]
│   ├── [[JSON vs Text Handlers & Custom Handler Pipelines]]
│   ├── [[Context-Aware Logging & Trace ID Correlation]]
│   ├── [[Log Levels, Dynamic Level Filtering & Groups]]
│   └── [[High-Performance Zero-Allocation Logging (Zap vs slog)]]
├── [[eBPF & Kernel-Level Observability|04. eBPF & Kernel-Level Observability]]
│   ├── [[eBPF Zero-Code Instrumentation Architecture]]
│   ├── [[eBPF Distributed Tracing with Cilium & Pixie]]
│   ├── [[Continuous Off-CPU Profiling with eBPF]]
│   └── [[Kernel Network Packet Latency & Drop Analysis]]
└── [[Production Diagnostics & Introspection (GODEBUG)|05. Production Diagnostics & Introspection (GODEBUG)]]
│   ├── [[GODEBUG Environment Flags Deep Catalog]]
│   ├── [[Automatic Heap Dumps on Memory Spikes (debug.WriteHeapDump)]]
│   ├── [[Core Dumps & Crash Interception (GOTRACEBACK)]]
│   └── [[Dynamic Memory Limit Tuning (debug.SetMemoryLimit)]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Runtime Metrics & Health Inspection|01. Runtime Metrics & Health Inspection]]
- [[runtime-metrics API Deep Dive]] — Structured runtime metrics: /sched/latencies:seconds, /gc/pauses:seconds, /memory/classes/heap:bytes.
- [[expvar Package & Public HTTP Telemetry]] — Publishing public JSON metrics endpoints with standard counters, maps, and custom stats.
- [[Prometheus Go Client Integration (client_golang)]] — Gauge, Counter, Histogram, Summary, collector registration, and scrapers.
- [[Memory & GC Metrics Interpretation]] — Correlating heap in-use, heap sys, GC cycle counts, and GC CPU utilization percentages.
- [[Scheduler & Goroutine Contention Metrics]] — Monitoring runnable goroutines waiting on logical processor queues and starvation.
- [[Kubernetes Health Probes (Liveness, Readiness, Startup)]] — Designing resilient, non-blocking health check endpoints with timeout isolation.
### 2. 📂 [[Distributed Tracing & OpenTelemetry (OTel)|02. Distributed Tracing & OpenTelemetry (OTel)]]
- [[OpenTelemetry Go SDK Architecture (TracerProvider & Spans)]] — Initializing OpenTelemetry TracerProvider, Samplers, Span processors, and resource detectors.
- [[Distributed Context Propagation (W3C TraceContext & B3)]] — Injecting and extracting trace IDs across HTTP (propagation.TraceContext) and gRPC metadata.
- [[OTLP Exporters (gRPC & HTTP Protocols)]] — Exporting telemetry streams to OpenTelemetry Collector, Jaeger, Tempo, and Honeycomb.
- [[Database & HTTP Middleware Instrumentation]] — Auto-instrumenting net/http handlers and SQL drivers (otelhttp, otelsql) with parent span binding.
- [[Span Events, Status Codes & Error Recording]] — span.RecordError(), span.SetStatus(), and attaching structured semantic attributes to spans.
### 3. 📂 [[Structured Logging Architecture (slog)|03. Structured Logging Architecture (slog)]]
- [[log-slog Core Architecture (Logger, Handler, Record)]] — Go 1.21+ structured logging engine, Handler interface contract, and Record lifecycle.
- [[JSON vs Text Handlers & Custom Handler Pipelines]] — Building high-performance, non-blocking asynchronous log writers with buffering.
- [[Context-Aware Logging & Trace ID Correlation]] — Automatically injecting OpenTelemetry trace IDs and span IDs into slog attributes via middleware.
- [[Log Levels, Dynamic Level Filtering & Groups]] — slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError, and grouped attributes (slog.Group).
- [[High-Performance Zero-Allocation Logging (Zap vs slog)]] — Comparing Uber Zap strongly typed zero-alloc fields with slog Handler performance.
### 4. 📂 [[eBPF & Kernel-Level Observability|04. eBPF & Kernel-Level Observability]]
- [[eBPF Zero-Code Instrumentation Architecture]] — Attaching Linux eBPF kernel kprobes and user-space uprobes to Go binaries without code changes.
- [[eBPF Distributed Tracing with Cilium & Pixie]] — Automatic protocol parsing (HTTP, gRPC, Kafka) at the Linux socket layer via eBPF.
- [[Continuous Off-CPU Profiling with eBPF]] — Measuring nanoseconds spent sleeping on lock mutexes, epoll waits, and disk I/O.
- [[Kernel Network Packet Latency & Drop Analysis]] — Diagnosing TCP retransmits, SYN queue drops, and connection resets using eBPF TC/XDP filters.
### 5. 📂 [[Production Diagnostics & Introspection (GODEBUG)|05. Production Diagnostics & Introspection (GODEBUG)]]
- [[GODEBUG Environment Flags Deep Catalog]] — gctrace=1 (GC cycle stats), schedtrace=1000 (GMP states), scheddetail=1, asyncpreemptoff=1.
- [[Automatic Heap Dumps on Memory Spikes (debug.WriteHeapDump)]] — Programmatically dumping full heap snapshots to disk before container Linux OOM kills.
- [[Core Dumps & Crash Interception (GOTRACEBACK)]] — Configuring GOTRACEBACK=crash to generate Linux core dumps on fatal panics for Delve post-mortem.
- [[Dynamic Memory Limit Tuning (debug.SetMemoryLimit)]] — Dynamically adjusting GOMEMLIMIT at runtime based on container cgroup memory limits.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

