- [[GOMEMLIMIT Memory Balancer]] — Go 1.19+ soft memory limit preventing container OOM kills and tuning GC pacing.

- [[GOMAXPROCS vs CFS Quota Throttling]] — Kubernetes CFS bandwidth throttling and automating core counts with automaxprocs.

---
title: Build Caching & Reproducibility
tags:
  - golang
  - environment
  - principal-swe
parent: "[[Settings Environment]]"
---

# Build Caching & Reproducibility

Action cache mechanics, output caches, reproducible builds, and binary provenance.

```text
Build Caching & Reproducibility
│
├── [[GOCACHE and GOTMPDIR]]
├── [[Action Cache vs Output Cache]]
├── [[Reproducible Builds (-trimpath)]]
└── [[Binary Provenance & Build Info]]
```

---

## 🗂️ Topics

- [[GOCACHE and GOTMPDIR]] — Build artifact cache and temporary compilation directory mechanics.
- [[Action Cache vs Output Cache]] — Content-addressable storage of compilation actions and cache hit validation.
- [[Reproducible Builds (-trimpath)]] — Eliminating machine-specific path metadata for deterministic byte-for-byte binaries.
- [[Binary Provenance & Build Info]] — Inspecting embedded Git SHA, VCS revision, and build flags via debug.ReadBuildInfo().

---

## 🔗 References
- ⬆️ Parent: [[Settings Environment]]

