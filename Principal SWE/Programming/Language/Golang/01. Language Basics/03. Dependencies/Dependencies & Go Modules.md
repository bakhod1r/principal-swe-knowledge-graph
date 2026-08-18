---
title: Dependencies & Go Modules
tags:
  - golang
  - basics
parent: "[[Language Basics]]"
---

# Dependencies & Go Modules

Go module system, dependency resolution, checksum integrity, and private repositories.

```text
Dependencies & Go Modules
│
├── [[go.mod]]
├── [[go.sum]]
├── [[require Directive]]
├── [[replace Directive]]
├── [[retract Directive]]
├── [[Semantic Versioning]]
├── [[MVS (Minimal Version Selection)]]
├── [[GOPROXY]]
├── [[GOSUMDB]]
├── [[GOPRIVATE]]
└── [[Vendoring]]
```

---

## 🗂️ Topics

- [[go.mod]] — Module declaration file, module path, and dependency requirements.
- [[go.sum]] — Cryptographic checksums of direct and indirect module dependencies.
- [[require Directive]] — Declaring minimum required versions of external modules.
- [[replace Directive]] — Substituting module dependencies with local forks or paths.
- [[retract Directive]] — Retracting broken or retracted module version releases.
- [[Semantic Versioning]] — SemVer rules and v2+ major version import path suffixes.
- [[MVS (Minimal Version Selection)]] — Deterministic dependency resolution algorithm in Go.
- [[GOPROXY]] — HTTP download mirror and proxy for public Go modules.
- [[GOSUMDB]] — Cryptographic notary database for module checksum verification.
- [[GOPRIVATE]] — Configuring private corporate repositories bypassing GOPROXY/GOSUMDB.
- [[Vendoring]] — Embedding dependencies in local vendor/ directory with go mod vendor.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]
- 🎓 Root: [[Principal SWE]]
