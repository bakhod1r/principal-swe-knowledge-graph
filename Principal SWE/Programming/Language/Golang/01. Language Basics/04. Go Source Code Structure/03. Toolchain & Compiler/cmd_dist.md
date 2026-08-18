---
title: cmd/dist
tags:
  - golang
  - goroot
  - compiler
  - bootstrap
parent: "[[Toolchain & Compiler]]"
---

# `cmd/dist`

The bootstrap and build orchestrator for the Go toolchain itself. Not used when
building your programs — used when building **Go**.

## 1. The Bootstrap Problem

The Go compiler is written in Go. Compiling it requires a Go compiler.

```text
Go 1.4 (last C-written toolchain)      ← historical root
        │
        ▼
Go 1.17 bootstrap toolchain            ← current minimum
        │
        ├── cmd/dist compiles cmd/compile, cmd/link, cmd/asm
        ▼
new toolchain rebuilds itself and the standard library
```

`GOROOT_BOOTSTRAP` points at the older toolchain used for stage one.

## 2. Building Go From Source

```bash
git clone https://go.googlesource.com/go && cd go/src
GOROOT_BOOTSTRAP=/usr/local/go ./make.bash
```

`make.bash` is a thin wrapper that compiles and invokes `cmd/dist`.

## 3. Other Duties

```bash
go tool dist list          # every supported GOOS/GOARCH pair
go tool dist list -json
go tool dist env           # the toolchain's own view of the environment
```

`go tool dist list` is the authoritative cross-compilation target list — better
than any table.

## 4. Gotchas

- `dist` is unversioned toolchain internals; its flags change freely.
- A failed bootstrap almost always means `GOROOT_BOOTSTRAP` points at a toolchain
  older than the required minimum.
- See `GOHOSTOS` and `GOHOSTARCH` for the host/target split `dist` manages.

---

## 🔗 References
- ⬆️ Parent: [[Toolchain & Compiler]]
