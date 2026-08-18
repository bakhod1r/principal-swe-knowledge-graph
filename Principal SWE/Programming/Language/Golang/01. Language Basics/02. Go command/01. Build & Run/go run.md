---
title: Run Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - run
parent: "[[Build & Run]]"
---

# `go run`

Compiles and immediately runs a `main` package. The binary goes to a temporary
directory and is deleted afterwards.

## 1. Usage

```bash
go run .                       # current directory
go run ./cmd/api
go run main.go helper.go       # explicit file list — must be the whole package
go run . --port=8080           # args after the package go to the program
```

## 2. Run Without Installing

```bash
go run golang.org/x/tools/cmd/stringer@latest -type=Pill
```

Module-aware: fetches, builds, and runs a tool at a pinned version without
touching your `go.mod` or `GOBIN`.

## 3. Exit Codes

`go run` propagates the program's exit code. Build failures exit with 1 and print
to stderr, so distinguishing "build broke" from "program failed" needs care in scripts.

## 4. Gotchas

- The binary is temporary — signals like `SIGTERM` reach the child, but the
  process tree has an extra layer. For signal-sensitive work use `go build` first.
- `go run main.go` silently ignores other files in the package. Prefer `go run .`.
- Not for production. No caching of the final link step benefit over `go build`.

---

## 🔗 References
- ⬆️ Parent: [[Build & Run]]
