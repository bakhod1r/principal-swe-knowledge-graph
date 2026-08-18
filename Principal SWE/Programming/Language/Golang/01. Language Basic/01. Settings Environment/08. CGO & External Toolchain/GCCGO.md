---
title: GCCGO
tags:
  - golang
  - basics
  - environment
  - cgo
  - compiler
parent: "[[CGO & External Toolchain]]"
---

# `GCCGO`

Names the `gccgo` binary used when building with `-compiler=gccgo` instead of the
standard `gc` toolchain.

## 1. Usage

```bash
go env GCCGO                              # gccgo
go build -compiler=gccgo ./...
GCCGO=/usr/bin/gccgo-14 go build -compiler=gccgo ./...
```

## 2. `gc` vs `gccgo`

| | `gc` (default) | `gccgo` |
|---|---|---|
| Front end | Go's own, in `cmd_compile` | GCC front end |
| Optimizer | Go SSA backend | GCC's optimizer |
| Platforms | What Go ships | Anything GCC targets |
| Go version support | Current | Usually behind |

`gccgo` exists for architectures the `gc` toolchain never supported, and to reuse
GCC's mature optimizer on numeric code.

## 3. Gotchas

- `gccgo` lags the language spec — generics support arrived far later there.
- Escape analysis, inlining, and the runtime differ, so performance conclusions
  do not transfer between compilers.
- Most tooling (`go tool` `compile`, `-gcflags`, PGO) assumes `gc` and is inert
  under `-compiler=gccgo`.

---

## 🔗 References
- ⬆️ Parent: [[CGO & External Toolchain]]
