---
title: Stack Traces & Diagnostic Enrichment
tags:
  - golang
  - error-handling
  - principal-swe
parent: "[[Error Handling (Clean Code)]]"
---

# Stack Traces & Diagnostic Enrichment

Stack trace capture, structured error logging with slog, trace ID / span attachment, and monitoring integration.

```text
Stack Traces & Diagnostic Enrichment
│
├── [[Stack Trace Capture Mechanics (pkg-errors & custom)]]
├── [[Structured Error Logging Integration with slog]]
├── [[Distributed Trace ID & Span Attachment to Errors]]
└── [[Error Monitoring & Crash Reporting Integration (Sentry, Rollbar)]]
```

---

## 🗂️ Topics

- [[Stack Trace Capture Mechanics (pkg-errors & custom)]] — Capturing caller program counters (runtime.Callers) at error creation sites with minimal CPU overhead.
- [[Structured Error Logging Integration with slog]] — Emitting error causes, stack frames, and contextual key-value pairs to structured log streams.
- [[Distributed Trace ID & Span Attachment to Errors]] — Correlating errors with active OpenTelemetry span contexts for instant observability triaging.
- [[Error Monitoring & Crash Reporting Integration (Sentry, Rollbar)]] — Automated capture, fingerprinting, and grouping of unhandled production errors.

---

## 🔗 References
- ⬆️ Parent: [[Error Handling (Clean Code)]]

