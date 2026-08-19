---
title: Golang
tags:
  - golang
  - programming-language
  - principal-swe
parent: "[[Programming]]"
---

# 🔷 Go (Golang) — Principal Software Engineer Vault

A structured, battle-tested, high-performance reference repository for Go engineering at the Principal & Staff+ level.

```text
Golang
├── 01. Language Basics
├── 02. Methods & Interfaces
├── 03. Generics
├── 04. Error Handling
├── 05. Code Organization & Architecture
├── 06. Concurrency & Synchronization
├── 07. Standard Library Mastery
├── 08. Testing & Benchmarking
├── 09. Go Toolchain & Developer Experience
├── 10. Advanced Topics & Low-Level Go
├── 11. Performance Engineering & Profiling
├── 12. Design Patterns in Go
├── 13. Runtime & Internals
├── 14. Go Standard Library Source Reading
├── 15. Observability & Runtime Introspection
├── 16. Modern Language Features
└── 17. Security, Cryptography & Hardening in Go
```

---

## 🗺️ Core Knowledge Pillars

### 1. 📂 [[Language Basics|01. Language Basics]]
Environment setup, CGO linking, MVS algorithm, variables, atomic memory alignment, data types, IEEE-754 precision, slices, maps (Swiss Tables), structs (padding/false sharing), conditionals, loops, functions, and escape analysis.

### 2. 📂 `02. Methods & Interfaces`
Method receivers, value vs pointer semantics, method sets, dynamic interface dispatch, `iface`/`eface` memory layouts, `itab` virtual tables, compiler devirtualization, and interface design patterns.

### 3. 📂 [[Generics|03. Generics]]
Parametric polymorphism, type parameters, type constraints, type sets (`~T`), generic data structures, GcShape stenciling, dictionary passing, and generic design patterns.

### 4. 📂 `04. Error Handling`
Error interface mechanics, sentinel errors, error wrapping (`%w`), multi-errors (`errors.Join`), domain error codes, stack traces, panic/recover boundary isolation, and error anti-patterns.

### 5. 📂 `05. Code Organization & Architecture`
Package-Oriented Design (POD), `internal/` visibility enforcement, Clean/Hexagonal architecture, multi-module workspaces (`go.work`), compile-time DI (Google Wire), REST/gRPC API evolution.

### 6. 📂 `06. Concurrency & Synchronization`
GMP scheduler mechanics, channels (hchan), sync primitives, lock-free atomics, memory model, hardware cache lines, context cancellation trees, and distributed resilience patterns.

### 7. 📂 [[Standard Library Mastery|07. Standard Library Mastery]]
Deep production mastery of `net/http`, `io`, `os`, `context`, `sync`, `database/sql`, `crypto`, `encoding/json`, `slog`, and `syscall`.

### 8. 📂 [[Testing & Benchmarking|08. Testing & Benchmarking]]
Unit testing primitives, table-driven tests, mocks/fakes, Testcontainers integration, benchmarking allocation profiles (`benchstat`), native fuzzing, mutation testing, deterministic virtual time (`synctest`), and race detection.

### 9. 📂 [[Go Toolchain & Developer Experience|09. Go Toolchain & Developer Experience]]
Compiler CLI flags, linker flags (`-s -w`), assembly generation (`-S`), vet, govulncheck, custom linter development, goreleaser, and multi-stage Docker builds.

### 10. 📂 [[Advanced Topics & Low-Level Go|10. Advanced Topics & Low-Level Go]]
`unsafe.Pointer` mechanics, runtime reflection, Cgo FFI, Plan 9 assembly, compiler SSA pipeline, custom analyzers, linker internals, WebAssembly (GOOS=js), WASI (GOOS=wasip1), and TinyGo embedded targets.

### 11. 📂 `11. Performance Engineering & Profiling`
CPU/memory profiling with pprof, execution tracing with `go tool trace`, zero-allocation patterns, hardware cache optimization, continuous profiling fleets, PGO, sub-microsecond latency, and SIMD vectorization.

### 12. 📂 [[Design Patterns in Go|12. Design Patterns in Go]]
Idiomatic Go implementations of Creational, Structural, Behavioral, Concurrency, and Cloud-Native Microservice design patterns.

### 13. 📂 [[Runtime & Internals|13. Runtime & Internals]]
Deep source code dive into Go runtime internals: boot sequence, GMP scheduler (`proc.go`), Garbage Collector (`mgc.go`), TCMalloc allocator (`malloc.go`), panic/defer (`panic.go`), channels (`chan.go`), timers (`time.go`), and stack growth (`stack.go`).

### 14. 📂 [[Go Standard Library Source Reading|14. Go Standard Library Source Reading]]
Line-by-line architectural source code analysis of `net/http`, `io`, `sync`, `context`, `database/sql`, `bytes`, and `reflect`.

### 15. 📂 [[Observability & Runtime Introspection|15. Observability & Runtime Introspection]]
Cloud-native observability in Go: `runtime/metrics`, OpenTelemetry distributed tracing, structured logging (`slog`), Linux eBPF kernel inspection, and GODEBUG diagnostics.

### 16. 📂 [[Modern Language Features|16. Modern Language Features]]
Modern Go evolution (Go 1.21 through Go 1.24+): custom iterators (`iter.Seq`), loop variable scoping, generic aliases, weak pointers, deterministic virtual time (`synctest`), and post-quantum cryptography.

### 17. 📂 `17. Security, Cryptography & Hardening in Go`
TLS 1.3/mTLS zero-trust architecture, AEAD encryption (AES-GCM, ChaCha20), constant-time operations, memory zeroing, JWT/PASETO tokens, SSRF/SQLi/XSS defense, and SLSA supply chain security.

---

## 🔗 Global References
- ⬆️ Parent: [[Programming]]

---

## 🗂️ Topics

- [[Advanced Topics & Low-Level Go]]
- [[Application Architecture]]
- [[Code Organization]]
- [[Cryptography]]
- [[Design Patterns in Go]]
- [[Error Handling]]
- [[Generics]]
- [[Go Concurrency]]
- [[Go Interfaces]]
- [[Go Performance Engineering]]
- [[Go Security]]
- [[Go Standard Library Source Reading]]
- [[Go Synchronization]]
- [[Go Toolchain & Developer Experience]]
- [[Hardening]]
- [[Language Basics]]
- [[Methods]]
- [[Modern Language Features]]
- [[Observability & Runtime Introspection]]
- [[Profiling Tooling]]
- [[Runtime & Internals]]
- [[Standard Library Mastery]]
- [[Testing & Benchmarking]]
