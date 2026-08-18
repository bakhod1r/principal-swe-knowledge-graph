---
title: FFI & Cgo Architecture
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# FFI & Cgo Architecture

Cgo stack switching architecture, call overhead benchmarks, pointer passing rules, pure Go static builds, and exported C APIs.

```text
FFI & Cgo Architecture
│
├── [[Cgo Architecture & Cross-Language Stack Switching]]
├── [[Cgo Performance Overhead & Call Overhead Benchmarks]]
├── [[Passing Pointers between Go and C (Pointer Passing Rules)]]
├── [[Pure Go vs Cgo Compilation (CGO_ENABLED=0 vs 1)]]
└── [[Calling Go Functions from C (export Directives)]]
```

---

## 🗂️ Topics

- [[Cgo Architecture & Cross-Language Stack Switching]] — How Go executes C code: switching from 2KB Go stack to OS C thread stack and back.
- [[Cgo Performance Overhead & Call Overhead Benchmarks]] — Measuring the 100ns+ CPU overhead per Cgo call and optimizing via batching.
- [[Passing Pointers between Go and C (Pointer Passing Rules)]] — Rules governing passing Go pointers to C, C.CString, C.free, and memory leak prevention.
- [[Pure Go vs Cgo Compilation (CGO_ENABLED=0 vs 1)]] — Building pure statically linked binaries, musl vs glibc compatibility, and Docker scratch deployment.
- [[Calling Go Functions from C (export Directives)]] — Exporting Go functions with //export and building .so/.dylib C-shared dynamic libraries.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]

