---
title: cmd/cgo
tags:
  - golang
  - goroot
  - cgo
  - interop
parent: "[[Toolchain & Compiler]]"
---

# `cmd/cgo`

Generates the Go and C glue that lets a Go package call C, and vice versa. Invoked
automatically for any package importing `"C"`.

## 1. What Triggers It

```go
/*
#include <stdlib.h>
*/
import "C"
```

The comment immediately above `import "C"` is the **preamble** — real C compiled
by `CC` with `CGO_CFLAGS`.

## 2. What It Produces

```text
user.go  ──cgo──►  _cgo_gotypes.go   (Go declarations for C types)
                   user.cgo1.go      (rewritten Go)
                   user.cgo2.c       (C stubs)
                   _cgo_export.h     (C header for //export'ed Go funcs)
```

Inspect with:

```bash
go build -work -x ./...        # keeps the work dir, prints commands
go tool cgo -godefs types.go   # generate Go structs from C headers
```

## 3. The Cost of a cgo Call

```text
Go goroutine on a P
     │  cgo call
     ▼
runtime.entersyscall — the P is released to run other goroutines
     ▼
C code on the OS thread (a full OS stack, not a goroutine stack)
     ▼
runtime.exitsyscall — reacquire a P
```

Roughly 50–100× the cost of a Go call. Batch at the boundary; never call C in a
tight loop.

## 4. Gotchas

- **Cross-compiling silently disables cgo** unless `CC` is a cross toolchain —
  see `CGO_ENABLED` and `GOHOSTOS`.
- Go pointers may not be stored in C memory; violations are caught by
  `GODEBUG=cgocheck=2` — see `GODEBUG`.
- cgo memory is invisible to the GC and to `GOMEMLIMIT`.
- `import "C"` makes the package uncompilable by other Go implementations and
  blocks the pure-Go cross-compilation story entirely.

---

## 🔗 References
- ⬆️ Parent: [[Toolchain & Compiler]]
