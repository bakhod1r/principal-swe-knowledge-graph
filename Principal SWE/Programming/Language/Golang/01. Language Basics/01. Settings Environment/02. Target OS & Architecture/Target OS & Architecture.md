---
title: Target OS & Architecture
tags:
  - golang
  - environment
  - principal-swe
parent: "[[Settings Environment]]"
---

# Target OS & Architecture

Cross-compilation matrices, CGO linking semantics, libc/musl resolution, and DNS resolver engines.

```text
Target OS & Architecture
│
├── [[GOOS & GOARCH Matrix]]
├── [[CGO_ENABLED (Static vs Dynamic Linking)]]
├── [[libc vs musl in Scratch Containers]]
└── [[DNS Resolvers (cgo vs netgo)]]
```

---

## 🗂️ Topics

- [[GOOS & GOARCH Matrix]] — Target operating systems (linux, darwin, windows, wasip1) and CPU architectures (amd64, arm64, riscv64).
- [[CGO_ENABLED (Static vs Dynamic Linking)]] — Compiling pure statically linked binaries (CGO_ENABLED=0) vs dynamically linked C binaries.
- [[libc vs musl in Scratch Containers]] — Deploying binaries into scratch/alpine containers and musl compatibility traps.
- [[DNS Resolvers (cgo vs netgo)]] — Pure Go netgo DNS resolver vs cgo host OS libc resolver resolution differences.

---

## 🔗 References
- ⬆️ Parent: [[Settings Environment]]
- 🎓 Root: [[Principal SWE]]
