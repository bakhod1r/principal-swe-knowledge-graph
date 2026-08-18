---
title: GOHOSTARCH
tags:
  - golang
  - basics
  - environment
  - build
  - platform
  - cross-compile
parent: "[[Build & Platform]]"
---

# `GOHOSTARCH`

**Read-only.** The CPU architecture the toolchain binary itself was built for.
The architecture counterpart of `GOHOSTOS`; contrast with `GOARCH`.

## 1. Usage

```bash
go env GOHOSTARCH        # arm64 on Apple Silicon, amd64 on Intel
```

## 2. Rosetta / Emulation Trap

On Apple Silicon an `amd64` Go toolchain installed by mistake reports:

```text
GOHOSTARCH=amd64
```

and every build runs under Rosetta — roughly 2–4× slower compiles, with no error
message. Checking `GOHOSTARCH` is the fastest diagnosis.

```bash
uname -m          # arm64   (the real CPU)
go env GOHOSTARCH # amd64   → wrong toolchain installed
```

## 3. Bootstrap Relevance

`cmd/dist` uses `GOHOSTOS`/`GOHOSTARCH` to decide what to build first when
compiling Go from source. See `Toolchain & Compiler`.

## 4. Gotchas

- Read-only; not settable with `go env -w`.
- `GOARCH` defaults to `GOHOSTARCH`, which is why builds "just work" until you
  cross-compile.

---

## 🔗 References
- ⬆️ Parent: [[Build & Platform]]
