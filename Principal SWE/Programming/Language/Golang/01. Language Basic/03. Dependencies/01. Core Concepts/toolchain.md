---
title: toolchain directive
tags:
  - golang
  - basics
  - dependencies
  - modules
  - go.mod
  - toolchain
parent: "[[Core Concepts]]"
---

# `toolchain`

Names the **minimum Go toolchain** used to build this module — distinct from the
`go` directive, which declares language version. Go 1.21+.

## 1. Syntax

```go
go 1.23.0
toolchain go1.24.5
```

Read as: *the language level is 1.23, but build me with at least 1.24.5.*

## 2. How Switching Works

```text
local toolchain 1.23.1  runs `go build`
        │
        ├── go.mod says toolchain go1.24.5
        ▼
   GOTOOLCHAIN=auto (default)
        │
        ├── downloads go1.24.5 as a module from GOPROXY
        └── re-executes the build with it
```

Controlled by `GOTOOLCHAIN`: `auto` (default) switches, `local` refuses and
errors instead.

## 3. Managing It

```bash
go get toolchain@go1.24.5     # set the directive
go get go@1.23.0              # set the go directive
go get toolchain@none         # remove the directive
```

## 4. Gotchas

- The toolchain is fetched over `GOPROXY` like any module — air-gapped builds
  need `GOTOOLCHAIN=local` and a preinstalled version.
- `toolchain` must be ≥ the `go` directive; a lower value is rejected.
- A `toolchain` line makes builds reproducible across developer machines; without
  it, whoever has the newest Go silently builds differently.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
