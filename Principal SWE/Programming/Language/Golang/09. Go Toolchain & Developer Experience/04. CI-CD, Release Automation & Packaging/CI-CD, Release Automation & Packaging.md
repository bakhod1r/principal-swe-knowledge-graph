---
title: CI-CD, Release Automation & Packaging
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Toolchain & Developer Experience]]"
---

# CI-CD, Release Automation & Packaging

GoReleaser cross-compilation, scratch/distroless container packaging, monorepo CI build caching, and PIE security hardening.

```text
CI-CD, Release Automation & Packaging
│
├── [[GoReleaser Enterprise Pipeline Automation]]
├── [[Multi-Stage Container Builds (distroless & scratch)]]
├── [[Fast Monorepo CI Caching Strategies (Go Build & Module Caches)]]
├── [[Hermetic Builds with Go Vendoring (-mod=vendor)]]
└── [[Binary Security Hardening (PIE, ASLR, Stack Canaries)]]
```

---

## 🗂️ Topics

- [[GoReleaser Enterprise Pipeline Automation]] — Multi-platform cross-compilation, automated changelog generation, GitHub/GitLab releases, and Docker images.
- [[Multi-Stage Container Builds (distroless & scratch)]] — Ultra-compact (<15MB), hardened container images with CA certificates and non-root users.
- [[Fast Monorepo CI Caching Strategies (Go Build & Module Caches)]] — Sharing ~/.cache/go-build and ~/go/pkg/mod across CI runners to reduce build times by 80%.
- [[Hermetic Builds with Go Vendoring (-mod=vendor)]] — Guaranteeing immutable zero-network CI pipeline execution using vendored module dependencies.
- [[Binary Security Hardening (PIE, ASLR, Stack Canaries)]] — Building Position-Independent Executables (-buildmode=pie) for kernel memory protection and ASLR.

---

## 🔗 References
- ⬆️ Parent: [[Go Toolchain & Developer Experience]]
- 🎓 Root: [[Principal SWE]]
