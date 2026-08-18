---
title: Inspection Commands
tags:
  - golang
  - basics
  - cli
  - toolchain
parent: "[[Go Commands]]"
---

# 🔍 Inspection Commands

Read-only queries against the toolchain, the module graph, and compiled binaries.

```text
[[go env]]      → toolchain configuration        (see `Settings Environment`)
[[go list]]     → packages and modules, as JSON or templates
[[go doc]]      → documentation, offline, from source
[[go version]]  → toolchain version + build metadata baked into a binary
```

## 🗂️ Commands

- **[[go env]]** — print, persist (`-w`), and reset (`-u`) configuration.
- **[[go list]]** — the query interface; `-m all`, `-deps`, `-json`, `-f`.
- **[[go doc]]** — package and symbol docs without leaving the terminal.
- **[[go version]]** — `-m` reads modules and VCS stamps out of a built binary.

## 💡 Debugging Entry Point

When a build behaves unexpectedly, the answer is usually in one of:

```bash
go env -changed        # what is non-default here
go list -m all         # what version actually got selected
go version -m ./bin/x  # what is really in the artifact
```

---

## 🔗 References
- ⬆️ Parent: `Go Commands`
