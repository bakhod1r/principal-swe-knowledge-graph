---
title: Enterprise & Monorepos
tags:
  - golang
  - dependencies
  - principal-swe
parent: "[[Dependencies & Go Modules]]"
---

# Enterprise & Monorepos

Private corporate setups, vendoring, multi-module workspaces, and supply chain security.

```text
Enterprise & Monorepos
│
├── [[GOPRIVATE Enterprise Repositories]]
├── [[Vendoring (go mod vendor)]]
├── [[Multi-Module Workspaces (go.work)]]
└── [[Supply Chain Security (govulncheck & SBOM)]]
```

---

## 🗂️ Topics

- [[GOPRIVATE Enterprise Repositories]] — Configuring private corporate Git repositories bypassing GOPROXY/GOSUMDB.
- [[Vendoring (go mod vendor)]] — Embedding all third-party dependencies in local vendor/ directory for air-gapped CI/CD.
- [[Multi-Module Workspaces (go.work)]] — Managing multi-module monorepos locally without modifying go.mod replace directives.
- [[Supply Chain Security (govulncheck & SBOM)]] — Static vulnerability scanning with govulncheck and generating Software Bill of Materials.

---

## 🔗 References
- ⬆️ Parent: [[Dependencies & Go Modules]]

