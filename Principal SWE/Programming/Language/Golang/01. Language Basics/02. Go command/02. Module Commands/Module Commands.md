---
title: Module Commands
tags:
  - golang
  - basics
  - cli
  - toolchain
parent: "[[Go Commands]]"
---

# 📦 Module Commands

Commands that read and write `go.mod`, `go.sum`, and `go.work`.
Concepts live in `Dependencies`.

```text
go.mod
  │
  ├── [[go mod]]   → init, tidy, download, graph, why, verify, vendor, edit
  ├── [[go get]]   → add / upgrade / remove a requirement
  └── [[go work]]  → multi-module workspaces (go.work)
```

## 🗂️ Commands

- **[[go mod]]** — the module maintenance multitool.
- **[[go get]]** — dependency version changes. Does **not** install binaries any more.
- **[[go work]]** — develop several local modules together without `replace`.

## ⚠️ The One Rule

`go build` and `go test` run with `-mod=readonly`: they will never edit `go.mod`.
Only the commands on this page change your dependency graph.

---

## 🔗 References
- ⬆️ Parent: `Go Commands`
