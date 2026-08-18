---
title: GORISCV64
tags:
  - golang
  - basics
  - environment
  - build
  - platform
  - riscv
parent: "[[Build & Platform]]"
---

# `GORISCV64`

Selects the RISC-V **profile** (instruction-set level) the compiler may target
when `GOARCH=riscv64`. Added in Go 1.23. Sibling of `GOAMD64` and `GOARM64`.

## 1. Values

| Value | Meaning |
|---|---|
| `rva20u64` | **Default.** RVA20 user-mode profile — the conservative baseline |
| `rva22u64` | RVA22 — adds Zba/Zbb bit-manipulation, Zicond |
| `rva23u64` | RVA23 — adds the vector extension V, Zvbb, Zicond (Go 1.24+) |

```bash
GOARCH=riscv64 GORISCV64=rva22u64 go build ./cmd/app
```

## 2. Trade-off

Higher profile → better codegen (bit-manip and conditional-move instructions
instead of branch sequences) → **binary faults on older hardware**. Same trade-off
as `GOAMD64=v3` on x86.

## 3. Gotchas

- Only meaningful when `GOARCH=riscv64`; ignored otherwise, and it does not show
  up in `go env` output on other hosts.
- Runtime feature detection is limited on RISC-V, so a wrong profile is a hard
  `SIGILL`, not a graceful fallback.

---

## 🔗 References
- ⬆️ Parent: [[Build & Platform]]
