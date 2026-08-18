---
title: src
tags:
  - golang
  - goroot
  - stdlib
  - source
parent: "[[Installation Directories]]"
---

# `$GOROOT/src`

The complete source of the standard library, the runtime, and the toolchain
commands. The largest and most valuable directory in `GOROOT`.

## 1. Layout

```text
$GOROOT/src/
├── fmt/  net/  os/  sync/  ...   ← standard library (`Standard Library`)
├── runtime/                       ← the runtime (`Runtime Internals`)
├── cmd/                           ← the toolchain (`Toolchain & Compiler`)
│   ├── compile/  link/  go/  asm/  vet/  ...
├── internal/                      ← shared, unimportable (`Internal Packages`)
└── vendor/                        ← golang.org/x deps used by the stdlib
```

## 2. Reading It Is the Point

```bash
cd "$(go env GOROOT)/src"
less sync/once.go          # ~70 lines, teaches the whole memory-ordering idea
less runtime/proc.go       # the scheduler
grep -rn "func Marshal" encoding/json/
```

Go's standard library is deliberately readable — it is the best available
reference for idiomatic Go, and it is already on your disk.

## 3. It Is Compiled From Here

There are no pre-compiled `.a` archives shipped since Go 1.20; the standard
library is built from `src/` and cached in `GOCACHE` on first use. That is why
the very first build after installing Go is slow.

## 4. Gotchas

- **Never edit these files.** They are inputs to a hashed build; a local change
  silently affects every project and vanishes on upgrade.
- `src/vendor/` holds `golang.org/x/...` packages the stdlib itself uses; they
  are not importable from your code.
- `GOPATH``/src` is unrelated — that is the legacy workspace, see `GO111MODULE`.

---

## 🔗 References
- ⬆️ Parent: [[Installation Directories]]
