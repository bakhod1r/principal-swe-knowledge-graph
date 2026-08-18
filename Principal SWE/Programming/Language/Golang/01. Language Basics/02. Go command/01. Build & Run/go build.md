---
title: Build Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - build
parent: "[[Build & Run]]"
---

# `go build`

Compiles packages and their dependencies. Writes an executable only for `main`
packages; for library packages it compiles, caches the result, and discards the output.

## 1. Usage

```bash
go build                    # build package in current directory
go build ./...              # build every package in the module
go build -o bin/api ./cmd/api
go build -o /dev/null ./... # type-check everything, keep nothing
```

## 2. Key Flags

| Flag | Meaning |
|---|---|
| `-o <path>` | Output file (or directory, when building multiple packages) |
| `-race` | Enable the race detector (requires cgo on most platforms) |
| `-tags a,b` | Build constraints — selects `//go:build a` files |
| `-ldflags` | Passed to the linker; `-X pkg.Var=value` injects build metadata |
| `-gcflags` | Passed to the compiler; `all=-N -l` disables optimization/inlining for debugging |
| `-trimpath` | Strip absolute paths from the binary — required for reproducible builds |
| `-buildmode` | `exe`, `pie`, `c-archive`, `c-shared`, `plugin`, `shared` |
| `-mod` | `readonly` (default), `mod`, `vendor` |
| `-pgo` | Profile-guided optimization; `auto` picks up `default.pgo` |
| `-a` | Force rebuild of everything, ignoring the cache |
| `-n` / `-x` | Print the commands (without / with running them) |
| `-v` | Print package names as they are compiled |

## 3. Version Stamping

```bash
go build -ldflags "-X main.version=$(git describe --tags) -s -w" -o app
```

`-s -w` drop the symbol table and DWARF data — smaller binary, no stack-trace
symbolization loss for Go panics (Go keeps its own tables).

## 4. Gotchas

- Build output is cached in `GOCACHE` — a second `go build` doing "nothing" is correct.
- `-race` binaries run ~2–10× slower and use far more memory; never ship them.
- `-mod=readonly` is the default since Go 1.16: `go build` will **not** edit `go.mod`.
  Run `go mod` `tidy` or `go get` instead.
- Cross-compiling: set `GOOS` and `GOARCH`; CGO is disabled implicitly unless
  you provide a cross `CC`.

---

## 🔗 References
- ⬆️ Parent: [[Build & Run]]
