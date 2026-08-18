---
title: Golang
tags:
  - golang
  - backend
  - language
  - principal-swe
  - architecture
parent: "[[Language]]"
---

# 🐹 Golang (Principal Software Engineering)

Go (Golang) is an open-source, statically typed, compiled programming language designed at Google for ultra-high performance, massive concurrency, robust scalability, and resilient distributed systems.

```text
Golang Knowledge Universe
│
├── [[Language Basics|01. Language Basics]]
├── [[Methods & Interfaces|02. Methods & Interfaces]]
├── [[Generics|03. Generics]]
├── [[Error Handling|04. Error Handling]]
├── [[Code Organization & Architecture|05. Code Organization & Architecture]]
├── [[Concurrency & Synchronization|06. Concurrency & Synchronization]]
├── [[Standard Library Mastery|07. Standard Library Mastery]]
├── [[Testing & Benchmarking|08. Testing & Benchmarking]]
├── [[Go Toolchain & Developer Experience|09. Go Toolchain & Developer Experience]]
├── [[Advanced Topics & Low-Level Go|10. Advanced Topics & Low-Level Go]]
├── [[Performance Engineering & Profiling|11. Performance Engineering & Profiling]]
├── [[Design Patterns in Go|12. Design Patterns in Go]]
├── [[Runtime & Internals|13. Runtime & Internals]]
├── [[Go Standard Library Source Reading|14. Go Standard Library Source Reading]]
├── [[WebAssembly & Alternative Targets|15. WebAssembly & Alternative Targets]]
├── [[Observability & Runtime Introspection|16. Observability & Runtime Introspection]]
└── [[Modern Language Features|17. Modern Language Features]]
```

---

## 🗺️ Golang Comprehensive Roadmap (17 Pillars)

### 1. 🔤 [[Language Basics|01. Language Basics]]
- Environment & PATH, Go CLI Toolchain, Dependencies & Go Modules, GOROOT Layout, Variables & Constants, Data Types, Composite Types, Conditionals, Loops, Functions, and Pointers.

### 2. 🧩 [[Methods & Interfaces|02. Methods & Interfaces]]
- Methods vs Functions, Receivers, Interface Basics, Empty Interfaces (any), Embedding, Type Assertions, Type Switch, iface/eface Internals.

### 3. 🧬 [[Generics|03. Generics]]
- Type Parameters, Constraints (comparable, any), Generic Functions & Types, Type Inference, Performance & Stenciling.

### 4. 🚨 [[Error Handling|04. Error Handling]]
- Error Interface, Wrapping (%w), errors.Is/As/Join, Custom Errors, Sentinel Errors, panic & recover.

### 5. 📁 [[Code Organization & Architecture|05. Code Organization & Architecture]]
- Go Modules & Dependencies, Packages, Standard Layout, internal/ packages, Workspaces (go.work), DI & Architecture.

### 6. ⚡ [[Concurrency & Synchronization|06. Concurrency & Synchronization]]
- Goroutines, Channels (hchan), sync package (Mutex, RWMutex, Pool), context, GMP Scheduler, Concurrency Patterns.

### 7. 📚 [[Standard Library Mastery|07. Standard Library Mastery]]
- io, os, net/http, encoding/json, slog, time, bufio, crypto, database/sql, embed.

### 8. 🧪 [[Testing & Benchmarking|08. Testing & Benchmarking]]
- Unit Testing, Table-driven tests, Mocks/Stubs, httptest, Benchmarks, Coverage, Fuzzing, Testcontainers.

### 9. 🛠️ [[Go Toolchain & Developer Experience|09. Go Toolchain & Developer Experience]]
- Core commands, Code generation (go:generate), Linting (golangci-lint), Security (govulncheck), Delve debugging.

### 10. 🔬 [[Advanced Topics & Low-Level Go|10. Advanced Topics & Low-Level Go]]
- Memory Management, Escape Analysis, unsafe, reflect, Cgo & FFI, Assembly (Plan 9), PGO, Linker flags.

### 11. 🏎️ [[Performance Engineering & Profiling|11. Performance Engineering & Profiling]]
- CPU & Memory Profiling (pprof), Mutex/Block Profiling, Execution Tracer (go tool trace), Zero-Allocation patterns.

### 12. 🏛️ [[Design Patterns in Go|12. Design Patterns in Go]]
- Functional Options, Builder, Strategy, Decorator, Adapter, Factory, Observer, Singleton, Iterator, Facade, Middleware.

### 13. ⚙️ [[Runtime & Internals|13. Runtime & Internals]]
- Runtime Architecture, Scheduler Source (proc.go), GC Source (mgc.go), Memory Allocator (malloc.go), runtime package.

### 14. 📖 [[Go Standard Library Source Reading|14. Go Standard Library Source Reading]]
- Deep architectural walkthroughs of stdlib source: net/http, sync, runtime, context, database/sql, encoding/json.

### 15. 🌐 [[WebAssembly & Alternative Targets|15. WebAssembly & Alternative Targets]]
- GOOS=js GOARCH=wasm, WASI (wasip1), TinyGo embedded & WebAssembly compilation, interop & production deployment.

### 16. 📊 [[Observability & Runtime Introspection|16. Observability & Runtime Introspection]]
- runtime/metrics, expvar, Runtime Execution Tracing, OpenTelemetry Go SDK, GODEBUG flags.

### 17. ✨ [[Modern Language Features|17. Modern Language Features]]
- Iterators (iter, range over func), loopvar semantics (Go 1.22+), Builtins (min, max, clear), Generic Type Aliases.

---

## 🔗 Navigation
- ⬆️ Parent: [[Language]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
