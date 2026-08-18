---
title: GOHOSTOS
tags:
  - golang
  - basics
  - environment
  - build
  - platform
  - cross-compile
parent: "[[Build & Platform]]"
---

# `GOHOSTOS`

**Read-only.** The operating system the **toolchain itself** runs on, as opposed
to `GOOS`, which is the operating system you are building **for**.

## 1. The Distinction

```text
GOHOSTOS / GOHOSTARCH   →  where the compiler runs   (this machine)
GOOS     / GOARCH       →  where the binary runs     (the target)
```

```bash
go env GOHOSTOS GOHOSTARCH    # darwin arm64
go env GOOS GOARCH            # darwin arm64   (no cross-compile)

GOOS=linux GOARCH=amd64 go build ./cmd/app
go env GOOS GOARCH            # still darwin arm64 — the vars were per-command
```

## 2. Detecting a Cross Build

```bash
if [ "$(go env GOOS)" != "$(go env GOHOSTOS)" ]; then
  echo "cross-compiling: CGO will be disabled unless CC is a cross toolchain"
fi
```

This is exactly the condition that silently flips `CGO_ENABLED` to `0`.

## 3. Gotchas

- Both are read-only; `go env -w` rejects them.
- Tests cannot run under a cross build — `go test` needs an executable for the
  host, so `GOOS=linux go test` on macOS builds but does not run.
- `go tool dist list` enumerates every valid `GOOS/GOARCH` pair.

---

## 🔗 References
- ⬆️ Parent: [[Build & Platform]]
