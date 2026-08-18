---
title: Structured Logging Architecture (slog)
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Observability & Runtime Introspection]]"
---

# Structured Logging Architecture (slog)

Go 1.21+ slog engine, custom JSON/text handlers, trace correlation, dynamic log level filtering, and Zap comparisons.

```text
Structured Logging Architecture (slog)
│
├── [[log-slog Core Architecture (Logger, Handler, Record)]]
├── [[JSON vs Text Handlers & Custom Handler Pipelines]]
├── [[Context-Aware Logging & Trace ID Correlation]]
├── [[Log Levels, Dynamic Level Filtering & Groups]]
└── [[High-Performance Zero-Allocation Logging (Zap vs slog)]]
```

---

## 🗂️ Topics

- [[log-slog Core Architecture (Logger, Handler, Record)]] — Go 1.21+ structured logging engine, Handler interface contract, and Record lifecycle.
- [[JSON vs Text Handlers & Custom Handler Pipelines]] — Building high-performance, non-blocking asynchronous log writers with buffering.
- [[Context-Aware Logging & Trace ID Correlation]] — Automatically injecting OpenTelemetry trace IDs and span IDs into slog attributes via middleware.
- [[Log Levels, Dynamic Level Filtering & Groups]] — slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError, and grouped attributes (slog.Group).
- [[High-Performance Zero-Allocation Logging (Zap vs slog)]] — Comparing Uber Zap strongly typed zero-alloc fields with slog Handler performance.

---

## 🔗 References
- ⬆️ Parent: [[Observability & Runtime Introspection]]

