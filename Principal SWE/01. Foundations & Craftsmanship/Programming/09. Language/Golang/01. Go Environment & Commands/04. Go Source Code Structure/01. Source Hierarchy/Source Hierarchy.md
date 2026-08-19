---
title: Source Hierarchy
tags:
  - golang
  - source-structure
  - goroot
  - principal-swe
parent: "[[Go Source Code Structure]]"
---

# Source Hierarchy

Standard library source root (`$GOROOT/src`), runtime engine, compiler, assembly sources, kernel abstractions, and internal packages.

```text
Source Hierarchy
│
├── [[src Directory Layout]]
├── [[src-runtime Architecture]]
├── [[src-sync & src-syscall Kernel Abstractions]]
├── [[Assembly Sources (.s) and Plan 9 Pseudo-Registers (FP, SP, SB, PC)]]
├── [[src-cmd Toolchain Source]]
├── [[src-internal Private Packages]]
└── [[GOROOT bin and pkg Layout]]
```

---

## 🗂️ Topics

- [[src Directory Layout]] — Root source tree for all standard library packages and runtime engine.
- [[src-runtime Architecture]] — Core runtime engine sources: `proc.go`, `mgc.go`, `malloc.go`, `chan.go`, `panic.go`.
- [[src-sync & src-syscall Kernel Abstractions]] — OS primitives, mutex spinning, and Linux `epoll` / Darwin `kqueue` network poller.
- [[Assembly Sources (.s) and Plan 9 Pseudo-Registers (FP, SP, SB, PC)]] — Hand-optimized Plan 9 assembly routines and pseudo-register semantics.
- [[src-cmd Toolchain Source]] — Compiler and linker source code: `cmd/compile`, `cmd/link`, `cmd/asm`, `cmd/go`.
- [[src-internal Private Packages]] — Standard library private unimportable helper packages (`internal/cpu`, `internal/bytealg`).
- [[GOROOT bin and pkg Layout]] — Executable compiler binaries and precompiled package metadata archives.

---

## 🔗 References
- ⬆️ Parent: [[Go Source Code Structure]]
- 📚 Module: `Go Environment & Commands`
