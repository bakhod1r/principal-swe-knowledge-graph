---
title: Module Distribution & Integrity
tags:
  - golang
  - basics
  - dependencies
  - modules
parent: "[[Dependencies]]"
---

# 🌐 Module Distribution & Integrity

Go ensures fast, reproducible, and secure module distribution through two foundational services:

1. **`GOPROXY`** — Centralized caching proxy for downloading immutable module archives.
2. **`GOSUMDB`** — Global cryptographic transparency log verifying module checksums.

---

## 🗺️ Architectural Flow

```text
go get / go build
       │
       ▼
   GOPROXY (proxy.golang.org)
       │
       ├─ Download module .zip and .mod
       │
       ▼
   Local SHA-256 Hash Computed
       │
       ▼
   GOSUMDB (sum.golang.org)
       │
       ├─ Verify hash matches global transparency log
       │
       ▼
   Recorded into `go.sum`
```

---

## 🔗 Environment References
- 🌐 Proxy Configuration: `GOPROXY`
- 🛡️ Checksum Database: `GOSUMDB`
- 🔒 Bypass for Private Repositories: `Private Modules` | `GOPRIVATE`
- ⬆️ Parent: [[Dependencies]]
