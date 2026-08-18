---
title: Static Analysis, Linters & Security Auditing
tags:
  - golang
  - toolchain
  - principal-swe
parent: "[[Go Toolchain & Developer Experience]]"
---

# Static Analysis, Linters & Security Auditing

govulncheck vulnerability database, golangci-lint enterprise rule sets, custom go/analysis linters, and SBOM supply chain generation.

```text
Static Analysis, Linters & Security Auditing
│
├── [[govulncheck & Official Go Vulnerability Database]]
├── [[golangci-lint Configuration & Enterprise Rule Sets]]
├── [[Building Custom Static Linters with go-analysis]]
├── [[go vet Diagnostic Analyzers Suite]]
└── [[Software Bill of Materials (SBOM) Generation (cyclonedx-gomod)]]
```

---

## 🗂️ Topics

- [[govulncheck & Official Go Vulnerability Database]] — Static call-graph scanning detecting actively reachable CVE vulnerabilities in dependencies.
- [[golangci-lint Configuration & Enterprise Rule Sets]] — Configuring 50+ enterprise linters (errcheck, gocritic, govet, revive, staticcheck, wsl).
- [[Building Custom Static Linters with go-analysis]] — Writing bespoke organizational linters enforcing company API standards and conventions.
- [[go vet Diagnostic Analyzers Suite]] — Built-in compiler analyzers: printf, shadow, structtag, atomic, copylocks, and unreachable code.
- [[Software Bill of Materials (SBOM) Generation (cyclonedx-gomod)]] — Extracting automated SPDX/CycloneDX supply chain dependency manifests from compiled Go binaries.

---

## 🔗 References
- ⬆️ Parent: [[Go Toolchain & Developer Experience]]
- 🎓 Root: [[Principal SWE]]
