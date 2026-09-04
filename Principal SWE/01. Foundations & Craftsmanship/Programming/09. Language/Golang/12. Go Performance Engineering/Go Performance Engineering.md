---
title: Go Performance Engineering
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Golang]]"
---

# 🏎️ Performance Engineering & Profiling

CPU/memory profiling with pprof, execution tracing with go tool trace, zero-allocation patterns, hardware cache optimization, continuous profiling fleets, PGO, and high-throughput networking.

```text
Performance Engineering & Profiling
│
├── `01. Profiling Subsystems (pprof)`
│   ├── `CPU Profiling & Flamegraph Analysis`
│   ├── `Heap & Memory Profiling (Alloc vs Inuse)`
│   ├── `Mutex Contention Profiling (runtime.SetMutexProfileFraction)`
│   ├── `Block Profiling (runtime.SetBlockProfileRate)`
│   ├── `Goroutine Stack Dump Profiling`
│   └── `Interactive pprof CLI & Web Visualizer`
├── `02. Execution Tracing & Scheduler Latency`
│   ├── `Execution Tracer Architecture (go tool trace)`
│   ├── `Diagnosing GC Pauses & STW Latency`
│   ├── `Diagnosing Network & Syscall Blocking`
│   ├── `User-Defined Tasks & Regions (runtime-trace)`
│   └── `Flight Recording & Continuous Trace Dumps`
├── [[Zero-Allocation Optimization Patterns|03. Zero-Allocation Optimization Patterns]]
│   ├── `Buffer Recycling with sync.Pool`
│   ├── `Zero-Copy String and Byte Conversions`
│   ├── `Pre-Allocating Capacity in Slices and Maps`
│   ├── `Struct Field Alignment & Padding Elimination`
│   ├── `Avoiding Interface Boxing in Hot Loops`
│   └── `Stack vs Heap Optimization via Inlining`
├── [[CPU Cache & Hardware-Conscious Design|04. CPU Cache & Hardware-Conscious Design]]
│   ├── `CPU Cache Hierarchy (L1, L2, L3, Cache Lines)`
│   ├── `False Sharing Elimination with Cache Line Padding`
│   ├── `Sequential Memory Access & Cache Locality`
│   ├── `SIMD Vectorization & Bounds Check Elimination (BCE)`
│   └── `Branch Prediction & Hot-Path Code Alignment`
├── `05. Production Profiling & Fleet Telemetry`
│   ├── `Continuous Profiling Fleets (Pyroscope, Parca)`
│   ├── `Off-CPU Analysis with eBPF in Go`
│   ├── `Proactive Heap Dumps Before Container OOM Kills`
│   └── `GODEBUG Diagnostic Flags in Production`
├── [[Compiler Optimizations & PGO|06. Compiler Optimizations & PGO]]
│   ├── `Profile-Guided Optimization (PGO) in Go`
│   ├── `Automated PGO Pipeline in CI-CD`
│   ├── `Function Inlining Heuristics & go:noinline Pragma`
│   └── `Compiler Dead Code Elimination & Branch Pruning`
├── [[Benchmarking Methodology & Regression Gates|07. Benchmarking Methodology & Regression Gates]]
│   ├── `Statistically Valid Benchmarking (testing.B)`
│   ├── `Statistical Significance Testing with benchstat`
│   ├── `Isolating Benchmarking Environments`
│   └── `Automated Benchmark Regression Gates in CI`
├── [[High-Performance I-O & Networking|08. High-Performance I-O & Networking]]
├── `09. Primitives & Concurrency Performance`
├── [[Ultra-Low-Latency & Kernel Bypass Systems|10. Ultra-Low-Latency & Kernel Bypass Systems]]
└── `11. SIMD Vectorization & Mechanical Sympathy`
│   ├── `Buffered Stream Processing with bufio`
│   ├── `Zero-Copy Network I-O with sendfile & splice`
│   ├── `HTTP Connection Pooling & Keep-Alive Tuning`
│   └── `Fast-Path Serialization (Protobuf vs Sonic vs Stdlib JSON)`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Profiling Subsystems (pprof)|01. Profiling Subsystems (pprof)]]
- [[CPU Profiling & Flamegraph Analysis]] — Sampling CPU profiler (100Hz default), interpreting flamegraphs, hot path identification, and overhead.
- [[Heap & Memory Profiling (Alloc vs Inuse)]] — alloc_space, alloc_objects, inuse_space, inuse_objects, diagnosing steady-state vs transient memory bloat.
- [[Mutex Contention Profiling (runtime.SetMutexProfileFraction)]] — Measuring lock wait times and lock acquisition contention hotspots in production.
- [[Block Profiling (runtime.SetBlockProfileRate)]] — Tracking goroutines blocked on channels, network I/O, and system calls.
- [[Goroutine Stack Dump Profiling]] — Capturing full stack traces of all live goroutines to detect deadlocks and goroutine leaks.
- [[Interactive pprof CLI & Web Visualizer]] — go tool pprof interactive commands (top, list, web, peek, disasm, source) and web visualizer.
### 2. 📂 [[Execution Tracing & Scheduler Latency|02. Execution Tracing & Scheduler Latency]]
- [[Execution Tracer Architecture (go tool trace)]] — Nanosecond-level execution timeline recording: event log buffers, per-P trace buffers, and trace format.
- [[Diagnosing GC Pauses & STW Latency]] — Correlating GC sweep and mark phases with application p99/p999 latency spikes.
- [[Diagnosing Network & Syscall Blocking]] — Tracking goroutines parked on Netpoller and OS thread syscall handoffs.
- [[User-Defined Tasks & Regions (runtime-trace)]] — Instrumenting domain business logic with trace.WithRegion and trace.Log for execution tracing.
- [[Flight Recording & Continuous Trace Dumps]] — Capturing circular in-memory flight recording ring buffers on error triggers.
### 📂 [[Zero-Allocation Optimization Patterns|01. Zero-Allocation Optimization Patterns]]
- [[Buffer Recycling with sync.Pool]] — Eliminating heap churn with per-P byte slice and struct pools in high-throughput services.
- [[Zero-Copy String and Byte Conversions]] — Safe zero-allocation string/byte conversions via unsafe.String and unsafe.Slice.
- [[Pre-Allocating Capacity in Slices and Maps]] — Eliminating dynamic growth reallocations via make([]T, 0, cap) and make(map[K]V, hint).
- [[Struct Field Alignment & Padding Elimination]] — Ordering fields from largest to smallest to eliminate 8-byte word padding waste.
- [[Avoiding Interface Boxing in Hot Loops]] — Passing concrete types to prevent runtime iface allocations in high-throughput hot paths.
- [[Stack vs Heap Optimization via Inlining]] — Structuring leaf functions within inlining budgets (80 AST nodes) to keep objects on the stack.
### 📂 [[CPU Cache & Hardware-Conscious Design|02. CPU Cache & Hardware-Conscious Design]]
- [[CPU Cache Hierarchy (L1, L2, L3, Cache Lines)]] — 64-byte cache line structure, latency tiers (1ns L1 vs 100ns RAM), and temporal/spatial locality.
- [[False Sharing Elimination with Cache Line Padding]] — Placing [64]byte padding between high-contention atomic counters on multicore CPUs.
- [[Sequential Memory Access & Cache Locality]] — Array-of-Structs (AoS) vs Struct-of-Arrays (SoA) data layouts for CPU prefetchers.
- [[SIMD Vectorization & Bounds Check Elimination (BCE)]] — Helping the compiler eliminate slice bounds checks and emit vectorized CPU instructions.
- [[Branch Prediction & Hot-Path Code Alignment]] — Structuring error branches and happy paths to maximize CPU branch prediction hit rates.
### 5. 📂 [[Production Profiling & Fleet Telemetry|05. Production Profiling & Fleet Telemetry]]
- [[Continuous Profiling Fleets (Pyroscope, Parca)]] — Ultra-low overhead (<1% CPU) continuous fleet-wide profiling architecture.
- [[Off-CPU Analysis with eBPF in Go]] — Measuring time spent off-CPU waiting on kernel locks, disk I/O, and page faults using eBPF uprobes.
- [[Proactive Heap Dumps Before Container OOM Kills]] — Using debug.SetMemoryLimit and memory watchdogs to dump heaps before SIGKILL.
- [[GODEBUG Diagnostic Flags in Production]] — gctrace=1, schedtrace=1000, madvdontneed=1 for real-time runtime diagnostics.
### 📂 [[Compiler Optimizations & PGO|03. Compiler Optimizations & PGO]]
- [[Profile-Guided Optimization (PGO) in Go]] — Feeding production .pgo profiles into go build for branch prediction and inlining gains.
- [[Automated PGO Pipeline in CI-CD]] — Automated collection of Kubernetes CPU profiles and auto-injecting into build artifacts.
- [[Function Inlining Heuristics & go:noinline Pragma]] — Understanding inlining budget calculation and selectively preventing inlining.
- [[Compiler Dead Code Elimination & Branch Pruning]] — How the Go compiler removes unreached code branches during SSA compilation.
### 📂 [[Benchmarking Methodology & Regression Gates|04. Benchmarking Methodology & Regression Gates]]
- [[Statistically Valid Benchmarking (testing.B)]] — b.N, b.ResetTimer(), b.ReportAllocs(), and avoiding compiler loop elimination.
- [[Statistical Significance Testing with benchstat]] — Comparing before-and-after benchmark results with p-value statistical confidence.
- [[Isolating Benchmarking Environments]] — CPU frequency governor locking, core affinity pinning (taskset), and disabling turbo boost.
- [[Automated Benchmark Regression Gates in CI]] — Enforcing maximum allocation count and latency regression thresholds on pull requests.
### 📂 [[High-Performance I-O & Networking|05. High-Performance I-O & Networking]]
- [[Buffered Stream Processing with bufio]] — Custom buffer sizing (bufio.NewReaderSize) for minimizing OS read/write syscalls.
- [[Zero-Copy Network I-O with sendfile & splice]] — Bypassing user-space memory buffers for direct kernel-to-socket transfers.
- [[HTTP Connection Pooling & Keep-Alive Tuning]] — Optimizing MaxIdleConns, MaxIdleConnsPerHost, and IdleConnTimeout in high-load clients.
- [[Fast-Path Serialization (Protobuf vs Sonic vs Stdlib JSON)]] — Performance matrix and allocation profiles of modern serialization formats.

### 📂 [[Concurrency Performance|09. Concurrency Performance]]

### 📂 [[Ultra-Low-Latency & Kernel Bypass Systems|07. Ultra-Low-Latency & Kernel Bypass Systems]]
- [[Kernel-Bypass Networking in Go (io_uring & DPDK)]] — Using Linux io_uring and DPDK for zero-copy, asynchronous zero-syscall network packet processing.
- [[Memory-Mapped I-O (mmap & unix.Mmap)]] — Mapping multi-gigabyte disk files directly into virtual memory address space for instantaneous zero-copy reads.
- [[NUMA-Aware Memory Architecture & CPU Pinning]] — Non-Uniform Memory Access node affinity and pinning goroutines/OS threads to local NUMA memory controllers.
- [[Memory Arena Allocator (arena package & GOEXPERIMENT=arenas)]] — Regional memory arena allocation enabling bulk object creation and instant zero-GC bulk deallocation.
- [[Lockless Disruptor Pattern in Go (LMAX Disruptor)]] — High-throughput circular ring buffer with memory sequence barriers executing 50M+ operations per second.
- [[Sub-Microsecond Latency Optimization Techniques]] — Eliminating all GC pauses, context switches, allocations, and thread parking in mission-critical hot paths.
- [[HugePages Allocation (2MB & 1GB Transparent Huge Pages)]] — Reducing CPU Translation Lookaside Buffer (TLB) misses and page table walking overhead for massive heap systems.

### 📂 [[SIMD Vectorization|08. SIMD Vectorization]]

### 📂 [[Mechanical Sympathy|10. Mechanical Sympathy]]

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Topics

- [[Allocation Primitives]]
- [[Benchmarking Methodology & Regression Gates]]
- [[CPU Cache & Hardware-Conscious Design]]
- [[Compiler Optimizations & PGO]]
- [[Concurrency Performance]]
- [[High-Performance I-O & Networking]]
- [[Mechanical Sympathy]]
- [[SIMD Vectorization]]
- [[Ultra-Low-Latency & Kernel Bypass Systems]]
- [[Zero-Allocation Optimization Patterns]]
