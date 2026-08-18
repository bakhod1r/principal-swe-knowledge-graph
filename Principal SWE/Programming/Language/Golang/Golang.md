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
├── [[Language Basic|01. Language Basics]]
├── [[Deep Dive Type System & Generics|02. Deep Dive Type System & Generics]]
├── [[Memory Management & Garbage Collection|03. Memory Management & Garbage Collection]]
├── [[Concurrency & Runtime Internals|04. Concurrency & Runtime Internals]]
├── [[Standard Library Mastery|05. Standard Library Mastery]]
├── [[System Design & Architecture in Go|06. System Design & Architecture in Go]]
├── [[Performance Engineering & Optimization|07. Performance Engineering & Optimization]]
├── [[Testing, Quality & Observability|08. Testing, Quality & Observability]]
├── [[Security & Production Hardening|09. Security & Production Hardening]]
└── [[Advanced Go & Low-Level Systems|10. Advanced Go & Low-Level Systems]]
```

---

## 🗺️ Golang Conceptual Roadmap (10 Pillars)

### 1. 📂 [[Language Basic|01. Language Basics]]
- **Setup & Environment**: GOROOT, GOPATH, GOBIN, OS environment variables.
- **Go Toolchain**: `go build`, `go test`, `go mod`, `go work`, `go tool`.
- **Dependency Management**: Go modules, MVS, GOPROXY, checksum DB, private modules.
- **Source Code Structure**: GOROOT anatomy, runtime internals, stdlib layout.
- **Fundamentals**: Declarations, scopes, types, control flow, functions, errors, memory model.

### 2. 🧬 [[Deep Dive Type System & Generics|02. Deep Dive Type System & Generics]]
- **Type System Internals**: Named vs underlying types, type identity, assignability.
- **Interface Internals**: `iface` vs `eface`, `_type`, `itab` dynamic dispatch table.
- **Reflection & Metadata**: `reflect.Type`, `reflect.Value`, 3 Laws of Reflection, overhead.
- **Unsafe Mechanics**: `unsafe.Pointer`, `uintptr`, pointer arithmetic, `unsafe.Slice/String`.
- **Generics**: Type parameters, constraints, `comparable`, `any`, GcShape stenciling.

### 3. 🧠 [[Memory Management & Garbage Collection|03. Memory Management & Garbage Collection]]
- **Stack vs Heap**: Contiguous stack allocation (2KB→1GB), escape analysis algorithms & flags.
- **Memory Allocator**: TCMalloc design, `mcache`, `mcentral`, `mheap`, `mspan`, size classes.
- **Garbage Collector**: Tricolor Mark & Sweep, Concurrent GC, Write Barrier, Hybrid Barrier.
- **GC Pacing & Tuning**: `GOGC`, `GOMEMLIMIT`, pacing heuristics, latency & STW minimization.
- **Memory Diagnostics**: Goroutine leaks, sub-slice memory retention, diagnostic tools.

### 4. ⚡ [[Concurrency & Runtime Internals|04. Concurrency & Runtime Internals]]
- **GMP Scheduler**: Goroutines (`G`), OS threads (`M`), Processors (`P`), Work stealing, Netpoller.
- **Preemption & Sysmon**: `sysmon` thread, cooperative vs asynchronous signal-based preemption.
- **Channel Internals**: `hchan` struct, ring buffer, wait queues (`sudog`), lock mechanisms.
- **Sync Primitives**: `sync.Mutex` (normal vs starvation), `RWMutex`, `WaitGroup`, `Once`, `sync.Pool`, `sync.Map`.
- **Context System**: `context.Context` cancellation trees, deadlines, propagation, values.
- **Concurrency Patterns**: Worker pools, Pipelines, Fan-In/Fan-Out, Errgroup, Graceful shutdown.

### 5. 📚 [[Standard Library Mastery|05. Standard Library Mastery]]
- **I/O Subsystem**: `io.Reader`, `io.Writer`, `bufio`, `bytes.Buffer`, `strings.Builder`, stream processing.
- **Networking & HTTP**: `net`, `net/http` client/server lifecycle, connection pools, HTTP/2 & 3.
- **Serialization**: `encoding/json`, fast JSON parsers, `encoding/binary`, Protobuf integration.
- **Database & Persistence**: `database/sql` connection pooling, transactions, context management.
- **OS & System**: `os`, `os/exec`, POSIX signal handling, subprocesses, syscalls.
- **Cryptography**: `crypto/tls`, `crypto/rand`, AES, SHA256, TLS 1.3 handshake negotiation.

### 6. 🏛️ [[System Design & Architecture in Go|06. System Design & Architecture in Go]]
- **Clean & Hexagonal**: Ports & Adapters, Domain, UseCase, Repository, Dependency Inversion.
- **Domain-Driven Design (DDD)**: Entities, Value Objects, Aggregates, Domain Events in Go.
- **Project Layout**: Standard Go Layout (`cmd/`, `internal/`, `pkg/`, `api/`), package boundary enforcement.
- **Microservices**: gRPC/Protobuf contracts, RESTful design, interceptors, middleware chains.
- **Event-Driven Architecture**: Kafka, RabbitMQ, NATS, Transactional Outbox, Sagas, Idempotency.
- **Resilient Errors**: Sentinel errors, custom error types, `errors.Is`/`As`/`Join`, domain vs infra errors.
- **Configuration & DI**: Functional options pattern, Uber Fx, Google Wire, 12-factor config.

### 7. 🚀 [[Performance Engineering & Optimization|07. Performance Engineering & Optimization]]
- **Profiling**: `pprof` CPU, Heap, Goroutine, Mutex, Block profiles, Flamegraphs.
- **Execution Tracer**: `go tool trace`, scheduler latency, GC pauses, syscall blocking.
- **Benchmarking**: `testing.B`, `-benchmem`, statistical analysis with `benchstat`.
- **Compiler Optimizations**: Function inlining, Bounds Check Elimination (BCE), dead code elimination.
- **Zero-Allocation**: Memory alignment, struct padding, `sync.Pool`, zero-copy string/byte slicing.

### 8. 🧪 [[Testing, Quality & Observability|08. Testing, Quality & Observability]]
- **Testing Strategies**: Table-driven tests, subtests, golden file tests, `t.Helper()`.
- **Mocks & Testcontainers**: Interface mocking, Testcontainers-go, dependency isolation.
- **Fuzz Testing**: `testing.F` engine, corpus management, boundary edge-case detection.
- **Structured Logging**: `log/slog` standard library, custom handlers, attributes, JSON logs.
- **Distributed Tracing**: OpenTelemetry Go SDK, trace/span propagation across RPCs.
- **Metrics**: Prometheus `client_golang`, Counter, Gauge, Histogram, cardinality control.
- **Static Analysis**: `golangci-lint` optimization, custom AST linters with `go/analysis`.

### 9. 🛡️ [[Security & Production Hardening|09. Security & Production Hardening]]
- **Secure Coding**: OWASP Top 10 prevention, memory safety boundaries, injection defense.
- **TLS & Crypto Hardening**: Strict TLS 1.3, mutual TLS (mTLS), secure key management.
- **Supply Chain Security**: `govulncheck`, SBOM generation (CycloneDX/SPDX), SLSA provenance.
- **Production Containerization**: Multi-stage Docker builds, Distroless/Scratch images, non-root users.

### 10. 🔬 [[Advanced Go & Low-Level Systems|10. Advanced Go & Low-Level Systems]]
- **CGO & FFI**: Cgo mechanics, execution overhead, `runtime.Pinner`, memory safety across boundaries.
- **Go Assembly & Compiler**: Plan 9 assembly syntax, SSA passes, `go tool compile -S`, intrinsics.
- **WebAssembly & TinyGo**: Go Wasm compilation, WASI interface, TinyGo embedded systems.
- **Linux Systems & eBPF**: `cilium/ebpf`, `bpf2go`, kernel tracing, high-speed packet filtering.

---

## 🔗 Navigation
- ⬆️ Parent: [[Language]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
