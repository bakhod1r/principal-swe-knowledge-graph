---
title: Target OS & Architecture
tags:
  - golang
  - environment
  - architecture
  - principal-swe
parent: "[[Settings Environment]]"
---

# Target OS & Architecture

Cross-compilation matrices, microarchitecture levels, CGO linking semantics, libc/musl resolution, WebAssembly targets, and DNS resolver engines.

```text
Target OS & Architecture
│
├── [[GOOS & GOARCH Matrix]]
├── [[GOARM, GOAMD64, and GOMIPS Microarchitecture Levels (v1-v4)]]
├── [[CGO_ENABLED (Static vs Dynamic Linking)]]
├── [[CGO Cross-Compilation with Zig & musl-cross (Zero-Docker Cross-Compiling)]]
├── [[libc vs musl in Scratch Containers]]
├── [[WebAssembly Targets (GOOS=js vs GOOS=wasip1)]]
└── [[DNS Resolvers (cgo vs netgo)]]
```

---

## 🗂️ Topics

- [[GOOS & GOARCH Matrix]] — Target operating systems (`linux`, `darwin`, `windows`, `wasip1`) and CPU architectures (`amd64`, `arm64`, `riscv64`).
- [[GOARM, GOAMD64, and GOMIPS Microarchitecture Levels (v1-v4)]] — Tuning binary instructions for modern CPU instruction sets (`GOAMD64=v3` for AVX2).
- [[CGO_ENABLED (Static vs Dynamic Linking)]] — Compiling pure statically linked binaries (`CGO_ENABLED=0`) vs dynamically linked C binaries.
- [[CGO Cross-Compilation with Zig & musl-cross (Zero-Docker Cross-Compiling)]] — Cross-compiling CGO dependencies without Docker containers via `zig cc`.
- [[libc vs musl in Scratch Containers]] — Deploying binaries into `scratch`/Alpine containers and `musl` compatibility traps.
- [[WebAssembly Targets (GOOS=js vs GOOS=wasip1)]] — Browser WebAssembly (`GOOS=js`) vs Serverless WASI binaries (`GOOS=wasip1`).
- [[DNS Resolvers (cgo vs netgo)]] — Pure Go `netgo` DNS resolver vs `cgo` host OS `libc` resolver resolution differences.

---

## 🔗 References
- ⬆️ Parent: [[Settings Environment]]
- 📚 Module: `Go Environment & Commands`
