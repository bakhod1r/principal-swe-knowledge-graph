---
title: Diagnostic Endpoints
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Diagnostic Endpoints

Health/readiness probes (/healthz, /readyz), runtime introspection (/debug/pprof), dynamic log level switching, and circuit breaker status.

```text
Diagnostic Endpoints
│
├── [[Standardized Health and Readiness Probes (k8s healthz, readyz)]]
├── [[Runtime Introspection Endpoints (debug-pprof, debug-vars)]]
├── [[Dynamic Log Level Switching Without Service Restarts]]
└── [[Circuit Breaker and Feature Flag Status Endpoints]]
```

---

## 🗂️ Topics

- [[Standardized Health and Readiness Probes (k8s healthz, readyz)]] — Designing deep vs shallow dependency health checks to prevent cascading Kubernetes pod restarts.
- [[Runtime Introspection Endpoints (debug-pprof, debug-vars)]] — Exposing safe, authenticated runtime profiling hooks and metrics on isolated administrative ports.
- [[Dynamic Log Level Switching Without Service Restarts]] — Changing log verbosity (DEBUG/INFO/ERROR) on running pods via administrative APIs during incidents.
- [[Circuit Breaker and Feature Flag Status Endpoints]] — Real-time state inspection of active circuit breakers, rate limiters, and dynamic feature toggles.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]
- 🎓 Root: [[Principal SWE]]
