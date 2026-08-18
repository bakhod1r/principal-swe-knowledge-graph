---
title: Fmt Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - formatting
parent: "[[Testing & Quality]]"
---

# `go fmt`

A thin wrapper that runs `gofmt -l -w` on the named packages.

## 1. Usage

```bash
go fmt ./...           # rewrite files in place, print the ones changed
gofmt -l .             # list unformatted files, change nothing
gofmt -d main.go       # unified diff of what would change
gofmt -s -w ./...      # also apply simplifications (not done by `go fmt`)
```

## 2. `go fmt` vs `gofmt`

| | `go fmt` | `gofmt` |
|---|---|---|
| Argument | package patterns (`./...`) | files and directories |
| Simplify `-s` | no | yes |
| Rewrite rules `-r` | no | yes |

For imports grouping/pruning use `goimports`, which `go fmt` does not include.

## 3. In CI

```bash
test -z "$(gofmt -l .)" || { gofmt -l .; exit 1; }
```

## 4. Gotchas

- Formatting is non-negotiable in Go; there are no options for style.
- `go fmt` writes files — in a check-only pipeline use `gofmt -l`.
- It skips directories named `testdata` and anything ignored by the build.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Quality]]
