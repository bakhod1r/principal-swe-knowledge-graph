---
title: Modules & Dependencies Environment
tags:
  - golang
  - basics
  - environment
  - modules
parent: "[[Settings Environment]]"
---

# 📦 Modules & Dependencies Environment

Variables controlling dependency resolution, module proxies, cryptographic verification, and private repositories.

---

## 🗂️ Module Variables

- [[GOPROXY]] — Proxy server chain for fetching modules (`https://proxy.golang.org,direct`).
- [[GOSUMDB]] — Cryptographic checksum database (`sum.golang.org`).
- [[GOPRIVATE]] — Prefix list matching private modules to bypass GOPROXY and GOSUMDB.
- [[GONOPROXY]] — Prefix list for modules that must be fetched directly via VCS.
- [[GONOSUMDB]] — Prefix list for private modules exempted from GOSUMDB verification.
- [[GOVCS]] — Allowed version control systems (`git`, `hg`, `svn`).
- [[GOTOOLCHAIN]] — Toolchain version management strategy (`auto`, `local`, `path`).

---

## 🔗 Navigation
- ⬆️ Parent: [[Settings Environment]]
- 📦 Module Architecture: `Dependencies`
