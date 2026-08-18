---
title: Quality, Security & Debugging
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Toolchain & Developer Experience]]"
---

# Quality, Security & Debugging

Static analysis with golangci-lint, vulnerability scanning with govulncheck, Delve interactive debugging.

```text
Quality, Security & Debugging
│
├── [[go vet Static Analyzer]]
├── [[golangci-lint Architecture]]
├── [[Security (govulncheck)]]
├── [[Debugging with Delve (dlv)]]
├── [[Runtime Diagnostic Flags (GODEBUG)]]
└── [[Live Reloading (Air)]]
```

---

## 🗂️ Topics

- [[go vet Static Analyzer]] — Standard compiler static analysis checks for unreachable code, print format bugs.
- [[golangci-lint Architecture]] — Configuring multi-linter pipelines, fast caching, linters settings (.golangci.yml).
- [[Security (govulncheck)]] — Scanning dependencies for known CVEs using Go vulnerability database and call-graph analysis.
- [[Debugging with Delve (dlv)]] — Setting breakpoints, inspecting goroutine stacks, evaluating variables with dlv CLI/IDE.
- [[Runtime Diagnostic Flags (GODEBUG)]] — gctrace=1, schedtrace=1000, madvdontneed=1, asyncpreemptoff=1 in production.
- [[Live Reloading (Air)]] — Hot-reloading Go services on file changes during local development.

---

## 🔗 References
- ⬆️ Parent: [[Go Toolchain & Developer Experience]]
- 🎓 Root: [[Principal SWE]]
