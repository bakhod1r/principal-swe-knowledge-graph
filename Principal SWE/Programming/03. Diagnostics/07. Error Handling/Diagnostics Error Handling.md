---
title: Diagnostics Error Handling
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Diagnostics Error Handling

Hierarchical domain error architecture, context chaining, stack trace preservation, panic-safe boundaries, and error budget alerting.

```text
Diagnostics Error Handling
│
├── [[Hierarchical Domain Error Architecture (Sentinel vs Typed Errors)]]
├── [[Error Context Chaining and Stack Trace Preservation]]
├── [[Panic-Safe Boundary Isolation in Goroutines and Worker Pools]]
└── [[Error Budget Tracking and SLA Error Rate Monitoring]]
```

---

## 🗂️ Topics

- [[Hierarchical Domain Error Architecture (Sentinel vs Typed Errors)]] — Designing domain error taxonomies with distinct machine codes and user-safe messages.
- [[Error Context Chaining and Stack Trace Preservation]] — Wrapping low-level errors while preserving root cause transparency and debugging context.
- [[Panic-Safe Boundary Isolation in Goroutines and Worker Pools]] — Containing worker goroutine panics from crashing entire parent processes in background pools.
- [[Error Budget Tracking and SLA Error Rate Monitoring]] — SRE error budget consumption alerting and automated SLO degradation mitigation triggers.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]

