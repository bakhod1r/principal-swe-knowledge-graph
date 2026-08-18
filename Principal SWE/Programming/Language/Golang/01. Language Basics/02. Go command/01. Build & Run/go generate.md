---
title: Generate Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - codegen
parent: "[[Build & Run]]"
---

# `go generate`

Scans packages for `//go:generate` directives and runs them. It is a **build step
you invoke manually** — no other `go` command runs it for you.

## 1. Directive Syntax

```go
//go:generate stringer -type=Pill
//go:generate mockgen -source=repo.go -destination=mock_repo.go
```

Must start at column 1, `//go:generate` with no space after `//`.

## 2. Usage

```bash
go generate ./...
go generate -run="mockgen" ./...   # only directives matching the regexp
go generate -n ./...               # dry run
```

## 3. Substituted Variables

| Variable | Value |
|---|---|
| `$GOFILE` | Current file name |
| `$GOLINE` | Line number of the directive |
| `$GOPACKAGE` | Package name |
| `$GOARCH` / `$GOOS` | Target platform |
| `$DOLLAR` | A literal `$` |

## 4. Gotchas

- Directives run in file order, per file, in the package directory.
- The tool must already be on `PATH` — pair with `go install` or the `tool`
  directive plus `go tool`.
- Generated files should be committed; CI verifies with
  `go generate ./... && git diff --exit-code`.

---

## 🔗 References
- ⬆️ Parent: [[Build & Run]]
