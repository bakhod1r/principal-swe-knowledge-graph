---
title: Metrics & Diagnostics
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Observability & Runtime Introspection]]"
---

# Metrics & Diagnostics

Reading runtime metrics counters, exposing HTTP diagnostics, and runtime debug flags.

```text
Metrics & Diagnostics
│
├── [[runtime-metrics Structured Metrics]]
├── [[expvar Public HTTP Metrics]]
├── [[GODEBUG Runtime Diagnostic Flags]]
└── [[runtime-debug Build Info & Heap Dumps]]
```

---

## 🗂️ Topics

- [[runtime-metrics Structured Metrics]] — Modern structured /sched/latencies, /memory/classes, /gc/pauses metrics inspection.
- [[expvar Public HTTP Metrics]] — Publishing public HTTP JSON metrics endpoints for standard Go counters and custom stats.
- [[GODEBUG Runtime Diagnostic Flags]] — gctrace=1, schedtrace=1000, madvdontneed=1, asyncpreemptoff=1 in production.
- [[runtime-debug Build Info & Heap Dumps]] — debug.ReadBuildInfo(), debug.WriteHeapDump(), debug.SetMemoryLimit() in code.

---

## 🔗 References
- ⬆️ Parent: [[Observability & Runtime Introspection]]
- 🎓 Root: [[Principal SWE]]
