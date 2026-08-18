---
title: Module Files & Checksums
tags:
  - golang
  - dependencies
  - principal-swe
parent: "[[Dependencies & Go Modules]]"
---

# Module Files & Checksums

Module declarations, direct/indirect dependencies, cryptographic checksum verification, and mirrors.

```text
Module Files & Checksums
│
├── [[go.mod Directives (require, replace, exclude, retract)]]
├── [[go.sum Checksum Verification]]
├── [[GOPROXY & Module Mirrors]]
└── [[GOSUMDB Notary Verification]]
```

---

## 🗂️ Topics

- [[go.mod Directives (require, replace, exclude, retract)]] — Module path, minimum Go version, and requirement manipulation directives.
- [[go.sum Checksum Verification]] — Cryptographic hashing, SHA-256 tree hashes, and tamper detection in module downloads.
- [[GOPROXY & Module Mirrors]] — HTTP download mirror and proxy architecture for public and caching Go modules.
- [[GOSUMDB Notary Verification]] — Cryptographic transparency log notary database verifying module checksum consistency.

---

## 🔗 References
- ⬆️ Parent: [[Dependencies & Go Modules]]
- 🎓 Root: [[Principal SWE]]
