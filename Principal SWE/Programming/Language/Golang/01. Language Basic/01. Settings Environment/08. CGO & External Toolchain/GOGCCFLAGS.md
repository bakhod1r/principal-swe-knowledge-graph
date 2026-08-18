---
title: GOGCCFLAGS
tags:
  - golang
  - basics
  - environment
  - cgo
  - compiler
parent: "[[CGO & External Toolchain]]"
---

# `GOGCCFLAGS`

**Read-only.** The flag list the `go` command passes to the host C compiler when
compiling cgo-generated code.

## 1. Inspecting It

```bash
go env GOGCCFLAGS
```

Typical macOS/arm64 value:

```text
-fPIC -arch arm64 -pthread -fno-caret-diagnostics -Qunused-arguments
-fmessage-length=0 -ffile-prefix-map=/var/folders/...=/tmp/go-build
-gno-record-gcc-switches -fno-common
```

## 2. Reading the Flags

| Flag | Why |
|---|---|
| `-fPIC` | Go links position-independent code |
| `-pthread` | The runtime is threaded regardless of your code |
| `-ffile-prefix-map=` | Reproducible builds — strips the temp build path |
| `-gno-record-gcc-switches` | Keeps the compiler command line out of the binary |
| `-fno-common` | Duplicate tentative definitions become link errors, not silent merges |

## 3. Gotchas

- Read-only: `go env -w GOGCCFLAGS=...` is rejected. To add flags use
  `CGO_CFLAGS` / `CGO_LDFLAGS`, which are appended to these.
- The value changes with `GOOS`, `GOARCH`, and the detected `CC` — do not
  hard-code it in build scripts.
- Its presence in `go env` output does **not** mean cgo is enabled; check
  `CGO_ENABLED`.

---

## 🔗 References
- ⬆️ Parent: [[CGO & External Toolchain]]
