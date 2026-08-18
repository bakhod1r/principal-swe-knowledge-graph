---
title: GOEXE
tags:
  - golang
  - basics
  - environment
  - build
  - platform
parent: "[[Build & Platform]]"
---

# `GOEXE`

**Read-only.** The filename suffix the toolchain appends to executables for the
target `GOOS`.

## 1. Values

```text
GOOS=windows  →  GOEXE=.exe
everything else →  GOEXE=      (empty)
```

```bash
go env GOEXE
GOOS=windows go env GOEXE     # .exe
```

## 2. Why It Matters

Scripts and Makefiles that build for several platforms need the suffix to name
the artifact correctly:

```bash
EXT=$(GOOS=$TARGET_OS go env GOEXE)
GOOS=$TARGET_OS GOARCH=$TARGET_ARCH go build -o "bin/app-$TARGET_OS$EXT" ./cmd/app
```

Without it, a Windows cross-build silently produces `bin/app` that Windows will
not execute.

## 3. Gotchas

- `go env -w GOEXE=...` is rejected — the value is derived, not configurable.
- `go install` applies `GOEXE` automatically when writing into `GOBIN`.
- `-o` with an explicit name overrides nothing: the `go` command does **not**
  append `GOEXE` to a path you specified yourself.

---

## 🔗 References
- ⬆️ Parent: [[Build & Platform]]
