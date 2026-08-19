---
title: Enterprise & Monorepos
tags:
  - golang
  - dependencies
  - enterprise
  - monorepos
  - security
  - principal-swe
parent: "[[Dependencies]]"
---

# Enterprise & Monorepos

Private corporate setups, authentication, corporate proxies, vendoring, multi-module workspaces, hermetic builds, and supply chain security.

```text
Enterprise & Monorepos
│
├── [[GOPRIVATE Enterprise Repositories]]
├── [[Private Module Authentication (SSH Keys & .netrc vs Git Token Helper)]]
├── [[GONOPROXY & GONOSUMDB Bypass Rules]]
├── [[Athens & JFrog Artifactory Private Go Proxies]]
├── [[Multi-Module Workspaces (go.work)]]
├── [[Vendoring (go mod vendor)]]
├── [[GOVCS Version Control Policy]]
├── [[Hermetic Builds in Bazel and Nix for Go]]
└── [[Supply Chain Security (govulncheck & SBOM)]]
```

---

## 🗂️ Topics

- [[GOPRIVATE Enterprise Repositories]] — Configuring private corporate Git repositories bypassing GOPROXY/GOSUMDB.
- [[Private Module Authentication (SSH Keys & .netrc vs Git Token Helper)]] — Authenticating `go get` against private GitHub, GitLab, and Bitbucket hosts in automated pipelines.
- [[GONOPROXY & GONOSUMDB Bypass Rules]] — Fine-grained glob patterns that bypass the proxy and checksum database independently.
- [[Athens & JFrog Artifactory Private Go Proxies]] — Deploying and maintaining private corporate proxy mirrors and immutable artifact stores.
- [[Multi-Module Workspaces (go.work)]] — Managing multi-module monorepos locally without modifying `go.mod` replace directives.
- [[Vendoring (go mod vendor)]] — Embedding all third-party dependencies in local `vendor/` directory for air-gapped CI/CD.
- [[GOVCS Version Control Policy]] — Restricting allowed VCS tools (`git`, `hg`, `svn`) to prevent remote execution attacks.
- [[Hermetic Builds in Bazel and Nix for Go]] — Rule-based deterministic build caching and hermetic compiler sandboxing for massive Go codebases.
- [[Supply Chain Security (govulncheck & SBOM)]] — Static vulnerability scanning with `govulncheck` and generating Software Bill of Materials (SBOM).

---

## 🔗 References
- ⬆️ Parent: [[Dependencies]]
- 📚 Module: `Go Environment & Commands`
