---
title: Module Files & Checksums
tags:
  - golang
  - dependencies
  - modules
  - checksums
  - principal-swe
parent: "[[Dependencies]]"
---

# Module Files & Checksums

Module declarations, direct/indirect dependencies, workspace configurations, cryptographic checksum verification, mirrors, and cache management.

```text
Module Files & Checksums
│
├── [[go.mod Directives (require, replace, exclude, retract)]]
├── [[Indirect Dependencies & indirect Annotations]]
├── [[go.work File Syntax & Directives (use, replace)]]
├── [[go.sum Checksum Verification]]
├── [[GOPROXY & Module Mirrors]]
├── [[GOSUMDB Notary Verification]]
└── [[Module Cache Eviction & Integrity (go clean -modcache)]]
```

---

## 🗂️ Topics

- [[go.mod Directives (require, replace, exclude, retract)]] — Module path, minimum Go version, and requirement manipulation directives.
- [[Indirect Dependencies & indirect Annotations]] — Why transitive dependencies get pinned and how to manage `// indirect` requirements.
- [[go.work File Syntax & Directives (use, replace)]] — Local multi-module development configuration without altering `go.mod` files.
- [[go.sum Checksum Verification]] — Cryptographic hashing, SHA-256 tree hashes, and tamper detection in module downloads.
- [[GOPROXY & Module Mirrors]] — HTTP download mirror and proxy architecture for public and caching Go modules.
- [[GOSUMDB Notary Verification]] — Cryptographic transparency log notary database verifying module checksum consistency.
- [[Module Cache Eviction & Integrity (go clean -modcache)]] — Purging `$GOPATH/pkg/mod` and dealing with read-only cache permissions.

---

## 🔗 References
- ⬆️ Parent: [[Dependencies]]
- 📚 Module: `Go Environment & Commands`
