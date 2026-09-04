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
├── 01. Go Environment & Commands
├── 02. Language Basics
├── 03. Methods
├── 04. Generics
├── 05. Error Handling
├── 06. Code Organization
├── 07. Go Concurrency
├── 08. Standard Library Mastery
├── 09. Testing & Benchmarking
├── 10. Go Toolchain & Developer Experience
├── 11. Advanced Topics & Low-Level Go
├── 12. Go Performance Engineering
├── 13. Design Patterns in Go
├── 14. Runtime & Internals
├── 15. Go Standard Library Source Reading
├── 16. Observability & Runtime Introspection
├── 17. Modern Language Features
├── 18. Go Security
├── 19. Profiling Tooling
├── 20. Go Interfaces
├── 21. Go Synchronization
├── 22. Cryptography
├── 23. Hardening
└── 24. Application Architecture
```

---

## 🗺️ Core Knowledge Pillars

### 1. 📂 [[Go Environment & Commands|01. Go Environment & Commands]]
Environment variables, the `go` command surface, CGO linking, the module system and MVS algorithm, cross-compilation targets, and the layout of the Go distribution itself.

### 2. 📂 [[Language Basics|02. Language Basics]]
Variables, atomic memory alignment, data types, IEEE-754 precision, slices, maps (Swiss Tables), structs (padding/false sharing), conditionals, loops, functions, and escape analysis.

### 3. 📂 [[Methods|03. Methods]]
Method receivers, value vs pointer semantics, method sets, dynamic dispatch, `iface`/`eface` memory layouts, `itab` virtual tables, and compiler devirtualization.

### 4. 📂 [[Generics|04. Generics]]
Parametric polymorphism, type parameters, type constraints, type sets (`~T`), generic data structures, GcShape stenciling, dictionary passing, and generic design patterns.

### 5. 📂 [[Error Handling|05. Error Handling]]
Error interface mechanics, sentinel errors, error wrapping (`%w`), multi-errors (`errors.Join`), domain error codes, stack traces, panic/recover boundary isolation, and error anti-patterns.

### 6. 📂 [[Code Organization|06. Code Organization]]
Package-Oriented Design (POD), `internal/` visibility enforcement, Clean/Hexagonal architecture, multi-module workspaces (`go.work`), compile-time DI (Google Wire), REST/gRPC API evolution.

### 7. 📂 [[Go Concurrency|07. Go Concurrency]]
GMP scheduler mechanics, channels (hchan), sync primitives, lock-free atomics, memory model, hardware cache lines, context cancellation trees, and distributed resilience patterns.

### 8. 📂 [[Standard Library Mastery|08. Standard Library Mastery]]
Deep production mastery of `net/http`, `io`, `os`, `context`, `sync`, `database/sql`, `crypto`, `encoding/json`, `slog`, and `syscall`.

### 9. 📂 [[Testing & Benchmarking|09. Testing & Benchmarking]]
Unit testing primitives, table-driven tests, mocks/fakes, Testcontainers integration, benchmarking allocation profiles (`benchstat`), native fuzzing, mutation testing, deterministic virtual time (`synctest`), and race detection.

### 10. 📂 [[Go Toolchain & Developer Experience|10. Go Toolchain & Developer Experience]]
Compiler CLI flags, linker flags (`-s -w`), assembly generation (`-S`), vet, govulncheck, custom linter development, goreleaser, and multi-stage Docker builds.

### 11. 📂 [[Advanced Topics & Low-Level Go|11. Advanced Topics & Low-Level Go]]
`unsafe.Pointer` mechanics, runtime reflection, Cgo FFI, Plan 9 assembly, compiler SSA pipeline, custom analyzers, linker internals, WebAssembly (GOOS=js), WASI (GOOS=wasip1), and TinyGo embedded targets.

### 12. 📂 [[Go Performance Engineering|12. Go Performance Engineering]]
CPU/memory profiling with pprof, execution tracing with `go tool trace`, zero-allocation patterns, hardware cache optimization, continuous profiling fleets, PGO, sub-microsecond latency, and SIMD vectorization.

### 13. 📂 [[Design Patterns in Go|13. Design Patterns in Go]]
Idiomatic Go implementations of Creational, Structural, Behavioral, Concurrency, and Cloud-Native Microservice design patterns.

### 14. 📂 [[Runtime & Internals|14. Runtime & Internals]]
Deep source code dive into Go runtime internals: boot sequence, GMP scheduler (`proc.go`), Garbage Collector (`mgc.go`), TCMalloc allocator (`malloc.go`), panic/defer (`panic.go`), channels (`chan.go`), timers (`time.go`), and stack growth (`stack.go`).

### 15. 📂 [[Go Standard Library Source Reading|15. Go Standard Library Source Reading]]
Line-by-line architectural source code analysis of `net/http`, `io`, `sync`, `context`, `database/sql`, `bytes`, and `reflect`.

### 16. 📂 [[Observability & Runtime Introspection|16. Observability & Runtime Introspection]]
Cloud-native observability in Go: `runtime/metrics`, OpenTelemetry distributed tracing, structured logging (`slog`), Linux eBPF kernel inspection, and GODEBUG diagnostics.

### 17. 📂 [[Modern Language Features|17. Modern Language Features]]
Modern Go evolution (Go 1.21 through Go 1.24+): custom iterators (`iter.Seq`), loop variable scoping, generic aliases, weak pointers, deterministic virtual time (`synctest`), and post-quantum cryptography.

### 18. 📂 [[Go Security|18. Go Security]]
TLS 1.3/mTLS zero-trust architecture, AEAD encryption (AES-GCM, ChaCha20), constant-time operations, memory zeroing, JWT/PASETO tokens, SSRF/SQLi/XSS defense, and SLSA supply chain security.

### 19. 📂 [[Profiling Tooling|19. Profiling Tooling]]
Profiling subsystems (pprof), execution tracing and scheduler latency, and production profiling with fleet telemetry.

### 20. 📂 [[Go Interfaces|20. Go Interfaces]]
Interface fundamentals and contracts, type assertions and type switches, runtime internals (`iface`, `eface`, `itab`), and interface architecture and design patterns.

### 21. 📂 [[Go Synchronization|21. Go Synchronization]]
Synchronization primitives (`sync`), atomic operations, and lock-free concurrency.

### 22. 📂 [[Cryptography|22. Cryptography]]
Transport Layer Security and certificates (`crypto/tls`, `crypto/x509`), symmetric and asymmetric cryptographic primitives, and enterprise secrets management with key vaults (HSM, KMS).

### 23. 📂 [[Hardening|23. Hardening]]
Side-channel attacks and memory hardening, application hardening and OWASP Top 10 defense, Linux kernel isolation and sandboxing, and zero trust network architecture with service mesh security.

### 24. 📂 [[Application Architecture|24. Application Architecture]]
Project layouts and repository architectures, dependency injection and decoupling patterns, enterprise API design and evolution, and code architecture anti-patterns and code smells.

---

## 🔗 Global References
- ⬆️ Parent: [[Programming]]

---

## 🗂️ Topics

- [[Go Environment & Commands]]
- [[Language Basics]]
- [[Methods]]
- [[Generics]]
- [[Error Handling]]
- [[Code Organization]]
- [[Go Concurrency]]
- [[Standard Library Mastery]]
- [[Testing & Benchmarking]]
- [[Go Toolchain & Developer Experience]]
- [[Advanced Topics & Low-Level Go]]
- [[Go Performance Engineering]]
- [[Design Patterns in Go]]
- [[Runtime & Internals]]
- [[Go Standard Library Source Reading]]
- [[Observability & Runtime Introspection]]
- [[Modern Language Features]]
- [[Go Security]]
- [[Profiling Tooling]]
- [[Go Interfaces]]
- [[Go Synchronization]]
- [[Cryptography]]
- [[Hardening]]
- [[Application Architecture]]
