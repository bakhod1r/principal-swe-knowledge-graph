---
title: Modules & Packages
tags:
  - golang
  - architecture
parent: "[[Code Organization & Architecture]]"
---

# Modules & Packages

Go module system, semantic import versioning, package declarations, and private repos.

```text
Modules & Packages
│
├── [[Modules & go.mod Directives]]
├── [[go.sum Checksum Verification]]
├── [[Packages & Imports]]
├── [[Package Naming Conventions]]
├── [[Module Versioning (SemVer)]]
└── [[Private Modules & Enterprise Setup]]
```

---

## 🗂️ Topics

- [[Modules & go.mod Directives]] — require, replace, exclude, retract directives and toolchain maintenance.
- [[go.sum Checksum Verification]] — Cryptographic hashing and tamper detection in module downloads.
- [[Packages & Imports]] — Package namespace rules, single package per directory, dot imports, blank imports.
- [[Package Naming Conventions]] — Short, concise, lowercase, singular package names without underscores.
- [[Module Versioning (SemVer)]] — Semantic Versioning rules and v2+ major version import path suffixes.
- [[Private Modules & Enterprise Setup]] — GOPRIVATE, authentication, enterprise git repository integration.

---

## 🔗 References
- ⬆️ Parent: [[Code Organization & Architecture]]
- 🎓 Root: [[Principal SWE]]
