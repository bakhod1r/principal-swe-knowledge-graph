---
title: Modules, Workspaces & Versioning (go.mod, go.work)
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Code Organization & Architecture]]"
---

# Modules, Workspaces & Versioning (go.mod, go.work)

go.work multi-module workspaces, Semantic Import Versioning (v2+), replace/retract directives, private modules, and vendoring.

```text
Modules, Workspaces & Versioning (go.mod, go.work)
│
├── [[Multi-Module Workspaces with go.work]]
├── [[Semantic Import Versioning (v2+ Module Paths)]]
├── [[replace, exclude, and retract Directives in go.mod]]
├── [[Private Module Configuration (GOPRIVATE, GONOPROXY)]]
└── [[Vendoring Mechanics (go mod vendor & -mod=vendor)]]
```

---

## 🗂️ Topics

- [[Multi-Module Workspaces with go.work]] — Local development across multiple sibling modules without requiring replace directives.
- [[Semantic Import Versioning (v2+ Module Paths)]] — Major version upgrades (/v2), branch strategies, and breaking API boundary management.
- [[replace, exclude, and retract Directives in go.mod]] — Local development overrides, CVE-vulnerable release retraction, and dependency exclusions.
- [[Private Module Configuration (GOPRIVATE, GONOPROXY)]] — Authenticating with enterprise GitLab/GitHub private registries and SSH deploy keys.
- [[Vendoring Mechanics (go mod vendor & -mod=vendor)]] — Hermetic builds, offline CI pipelines, and vendored dependency auditing.

---

## 🔗 References
- ⬆️ Parent: [[Code Organization & Architecture]]

