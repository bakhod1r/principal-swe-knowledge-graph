---
title: Fix Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - migration
parent: "[[Testing & Quality]]"
---

# `go fix`

Rewrites source that uses old APIs or syntax to the current form.

## 1. Usage

```bash
go fix ./...
go tool fix -diff ./...     # preview without writing
go tool fix -r <fixname> .  # apply one named rewrite
```

## 2. What It Handles

Historic language and standard-library migrations recorded in the toolchain —
for example `context` moving out of `golang.org/x/net`, and pre-modules import
rewrites. Each fix is a named rule inside `cmd/fix`; see `cmd_go`.

## 3. Gotchas

- Rarely needed on modern code — most Go 1 API changes were additive.
- It rewrites files in place. Commit first, then run, then read the diff.
- It is **not** a general refactoring tool; for that use `gofmt -r` or `gopls rename`.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Quality]]
