---
title: Dependencies & Go Module System
tags:
  - golang
  - basics
  - dependencies
  - modules
parent: "[[Language Basic]]"
---

# 📦 Dependencies & Go Module System

Go provides a built-in, deterministic, and cryptographically verified dependency management system centered around **Go Modules**.

```text
Dependencies
│
├── [[Core Concepts]]
│   ├── Module, Package
│   ├── go.mod, go.sum
│   └── require, replace, go.work
│
├── [[Version Resolution]]
│   ├── Semantic Versioning
│   ├── MVS (Minimal Version Selection)
│   └── Transitive Dependencies
│
├── [[Distribution & Integrity]]
│   ├── GOPROXY
│   └── GOSUMDB
│
├── [[Private Modules]]
│   ├── GOPRIVATE
│   ├── GONOPROXY
│   └── GONOSUMDB
│
└── [[Security & Auditing]]
    ├── Vulnerabilities (govulncheck)
    ├── Supply Chain Security
    └── Dependency Auditing
```

---

## 🗂️ Categories

1. 🧱 **[[Core Concepts]]** — Modules, Packages, `go.mod`, `go.sum`, `require`, `replace`, and `go.work`.
2. 🎯 **[[Version Resolution]]** — Semantic Versioning, Minimal Version Selection (MVS), and Transitive Dependencies.
3. 🌐 **[[Distribution & Integrity]]** — GOPROXY caching, checksum database (GOSUMDB), and immutability.
4. 🔒 **[[Private Modules]]** — Enterprise private repositories, GOPRIVATE, GONOPROXY, and GONOSUMDB.
5. 🛡️ **[[Security & Auditing]]** — Vulnerability scanning (`govulncheck`), supply chain verification, and auditing.

---

## 🔗 Navigation
- ⬆️ Parent: [[Language Basic]]
- 🛠️ Setup Guide: `Settings Environment`
