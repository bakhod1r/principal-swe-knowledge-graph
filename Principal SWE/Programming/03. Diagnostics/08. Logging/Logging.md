---
title: Logging
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Logging

Structured contextual logging (slog, zap), trace correlation IDs, async buffer flushing, and log aggregation pipelines.

```text
Logging
│
├── [[Structured Contextual Logging Architecture (slog, zap, zerolog)]]
├── [[Context-Aware Trace and Correlation ID Propagation]]
├── [[Log Buffering, Async Flushing, and Backpressure Handling]]
└── [[Log Aggregation Architecture (Vector, Fluentbit, Loki, Elasticsearch)]]
```

---

## 🗂️ Topics

- [[Structured Contextual Logging Architecture (slog, zap, zerolog)]] — JSON-formatted key-value logging with zero memory allocations and high throughput.
- [[Context-Aware Trace and Correlation ID Propagation]] — Linking every log entry automatically to active distributed trace spans and request IDs.
- [[Log Buffering, Async Flushing, and Backpressure Handling]] — Ring-buffered log emitters preventing log system calls from stalling application request threads.
- [[Log Aggregation Architecture (Vector, Fluentbit, Loki, Elasticsearch)]] — Designing high-throughput, compressed, index-free log ingestion and search pipelines.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]

