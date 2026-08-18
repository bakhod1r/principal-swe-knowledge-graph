---
title: Build & Run Commands
tags:
  - golang
  - basics
  - cli
  - toolchain
parent: "[[Go Commands]]"
---

# 🔨 Build & Run Commands

Commands that turn source into artifacts and execute them.

```text
source
  │
  ├── [[go build]]     → binary on disk (cached objects in GOCACHE)
  ├── [[go run]]       → binary in a temp dir, executed, discarded
  ├── [[go install]]   → binary in GOBIN / $GOPATH/bin
  ├── [[go generate]]  → runs //go:generate directives (before building)
  └── [[go clean]]     → removes artifacts and caches
```

## 🗂️ Commands

- **[[go build]]** — compile packages; flags for tags, ldflags, trimpath, race, PGO.
- **[[go run]]** — compile and run in one step, including `pkg@version` tools.
- **[[go install]]** — install executables to `GOBIN`; the `@version` form for tooling.
- **[[go generate]]** — run code generators declared in source comments.
- **[[go clean]]** — remove object files, build cache, module cache, test cache.

---

## 🔗 References
- ⬆️ Parent: `Go Commands`
