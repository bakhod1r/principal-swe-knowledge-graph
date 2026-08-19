---
title: Build Caching & Reproducibility
tags:
  - golang
  - environment
  - build-cache
  - reproducibility
  - principal-swe
parent: "[[Settings Environment]]"
---

# Build Caching & Reproducibility

Action cache mechanics, output caches, reproducible builds, binary provenance, runtime tuning flags, and PGO compilation.

```text
Build Caching & Reproducibility
│
├── [[GOCACHE and GOTMPDIR]]
├── [[Action Cache vs Output Cache]]
├── [[Reproducible Builds (-trimpath)]]
├── [[Build Tags & Conditional Compilation Constraints (go build)]]
├── [[PGO (Profile-Guided Optimization) Workflow (default.pgo)]]
├── [[GOMEMLIMIT Memory Balancer]]
├── [[GOMAXPROCS vs CFS Quota Throttling]]
├── [[GODEBUG Environment Flags & Runtime Tracing (gctrace, schedtrace, asyncpreemptoff)]]
└── [[Binary Provenance & Build Info]]
```

---

## 🗂️ Topics

- [[GOCACHE and GOTMPDIR]] — Build artifact cache and temporary compilation directory mechanics.
- [[Action Cache vs Output Cache]] — Content-addressable storage of compilation actions and cache hit validation.
- [[Reproducible Builds (-trimpath)]] — Eliminating machine-specific path metadata for deterministic byte-for-byte binaries.
- [[Build Tags & Conditional Compilation Constraints (go build)]] — Modern `//go:build` syntax, boolean logic constraint combinations, and OS tags.
- [[PGO (Profile-Guided Optimization) Workflow (default.pgo)]] — Feeding production pprof CPU profiles into `go build` for automatic inlining and branch optimizations.
- [[GOMEMLIMIT Memory Balancer]] — Go 1.19+ soft memory limit preventing container OOM kills and tuning GC pacing.
- [[GOMAXPROCS vs CFS Quota Throttling]] — Kubernetes CFS bandwidth throttling and automating core counts with `automaxprocs`.
- [[GODEBUG Environment Flags & Runtime Tracing (gctrace, schedtrace, asyncpreemptoff)]] — Low-level runtime debugging output flags.
- [[Binary Provenance & Build Info]] — Inspecting embedded Git SHA, VCS revision, and build flags via `debug.ReadBuildInfo()`.

---

## 🔗 References
- ⬆️ Parent: [[Settings Environment]]
- 📚 Module: `Go Environment & Commands`
