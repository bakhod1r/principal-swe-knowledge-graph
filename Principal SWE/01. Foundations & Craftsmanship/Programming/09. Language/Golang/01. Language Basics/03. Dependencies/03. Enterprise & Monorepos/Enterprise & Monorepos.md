---
title: Enterprise & Monorepos
tags:
  - golang
  - dependencies
  - principal-swe
parent: "[[Dependencies]]"
---

# Enterprise & Monorepos

Private corporate setups, vendoring, multi-module workspaces, and supply chain security.

```text
Enterprise & Monorepos
│
├── [[GOPRIVATE Enterprise Repositories]]
├── [[Vendoring (go mod vendor)]]
├── [[Multi-Module Workspaces (go.work)]]
├── [[GONOPROXY & GONOSUMDB Bypass Rules]]
├── [[GOVCS Version Control Policy]]
└── [[Supply Chain Security (govulncheck & SBOM)]]
```

---

## 🗂️ Topics

- [[GOPRIVATE Enterprise Repositories]] — Configuring private corporate Git repositories bypassing GOPROXY/GOSUMDB.
- [[Vendoring (go mod vendor)]] — Embedding all third-party dependencies in local vendor/ directory for air-gapped CI/CD.
- [[Multi-Module Workspaces (go.work)]] — Managing multi-module monorepos locally without modifying go.mod replace directives.
- [[GONOPROXY & GONOSUMDB Bypass Rules]] — Glob patterns that bypass the proxy and the checksum database independently of GOPRIVATE.
- [[GOVCS Version Control Policy]] — Which version control commands `go get` may run per module path pattern.
- [[Supply Chain Security (govulncheck & SBOM)]] — Static vulnerability scanning with govulncheck and generating Software Bill of Materials.

---

## 🔗 References
- ⬆️ Parent: [[Dependencies]]

