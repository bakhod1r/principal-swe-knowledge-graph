---
title: Modules & Packages
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Code Organization & Architecture]]"
---

# Modules & Packages

Go module system, semantic import versioning, package declarations, and private repos.

```text
Modules & Packages
│
├── [[Modules & go.mod Directives]]
├── [[go.sum Dependency Verification & Integrity]]
├── [[Packages & Imports]]
├── [[Package Naming Conventions]]
├── [[Module Versioning (SemVer)]]
├── [[Private Modules & Enterprise Setup]]
└── [[Vendoring Mechanics (go mod vendor)]]
```

---

## 🗂️ Topics

- [[Modules & go.mod Directives]] — require, replace, exclude, and retract directives in go.mod.
- [[go.sum Dependency Verification & Integrity]] — Cryptographic hashing, go.sum integrity verification, and tamper detection.
- [[Packages & Imports]] — Package namespace rules, single package per directory, dot imports, and blank imports.
- [[Package Naming Conventions]] — Short, concise, lowercase, singular package names without underscores.
- [[Module Versioning (SemVer)]] — Semantic Versioning rules and v2+ major version import path suffixes.
- [[Private Modules & Enterprise Setup]] — GOPRIVATE, authentication, enterprise git repository integration.
- [[Vendoring Mechanics (go mod vendor)]] — Embedding dependencies for reproducible offline enterprise builds.

---

## 🔗 References
- ⬆️ Parent: [[Code Organization & Architecture]]
- 🎓 Root: [[Principal SWE]]
