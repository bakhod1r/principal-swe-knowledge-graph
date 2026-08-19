---
title: Production Diagnostics & Introspection (GODEBUG)
tags:
  - golang
  - observability
  - principal-swe
parent: "[[Observability & Runtime Introspection]]"
---

# Production Diagnostics & Introspection (GODEBUG)

GODEBUG flags (gctrace, schedtrace), proactive heap dumps on memory spikes, core dumps (GOTRACEBACK), and dynamic memory limits.

```text
Production Diagnostics & Introspection (GODEBUG)
│
├── [[GODEBUG Environment Flags Deep Catalog]]
├── [[Automatic Heap Dumps on Memory Spikes (debug.WriteHeapDump)]]
├── [[Core Dumps & Crash Interception (GOTRACEBACK)]]
└── [[Dynamic Memory Limit Tuning (debug.SetMemoryLimit)]]
```

---

## 🗂️ Topics

- [[GODEBUG Environment Flags Deep Catalog]] — gctrace=1 (GC cycle stats), schedtrace=1000 (GMP states), scheddetail=1, asyncpreemptoff=1.
- [[Automatic Heap Dumps on Memory Spikes (debug.WriteHeapDump)]] — Programmatically dumping full heap snapshots to disk before container Linux OOM kills.
- [[Core Dumps & Crash Interception (GOTRACEBACK)]] — Configuring GOTRACEBACK=crash to generate Linux core dumps on fatal panics for Delve post-mortem.
- [[Dynamic Memory Limit Tuning (debug.SetMemoryLimit)]] — Dynamically adjusting GOMEMLIMIT at runtime based on container cgroup memory limits.

---

## 🔗 References
- ⬆️ Parent: [[Observability & Runtime Introspection]]

