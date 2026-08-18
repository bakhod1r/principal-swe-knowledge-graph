---
title: Install Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - install
parent: "[[Build & Run]]"
---

# `go install`

Compiles and installs an executable into `GOBIN` (or `$GOPATH/bin` when `GOBIN`
is empty). See `GOPATH`.

## 1. Two Modes

```bash
# A) install from the current module
go install ./cmd/api

# B) install a tool at a version — ignores the current go.mod entirely
go install golang.org/x/tools/gopls@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
```

Mode B is the correct way to install tooling since Go 1.16. It does **not** add
the tool to your dependency graph.

## 2. Where It Lands

```text
go install
     ↓
   GOBIN  (if set)
     ↓
 $GOPATH/bin  (default)
     ↓
   PATH  → the tool is runnable
```

## 3. Tool Dependencies (Go 1.24+)

`go.mod` supports a `tool` directive, so project tooling is version-pinned:

```bash
go get -tool golang.org/x/tools/cmd/stringer
go tool stringer -type=Pill
```

See `go tool`.

## 4. Gotchas

- `@version` form requires an absolute module path — `go install ./x@latest` is invalid.
- If the tool does not appear on your shell, `$GOPATH/bin` is missing from `PATH`.
- `go install` never writes to `GOROOT`.

---

## 🔗 References
- ⬆️ Parent: [[Build & Run]]
