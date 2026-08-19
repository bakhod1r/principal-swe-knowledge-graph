---
title: Profiling Subsystems (pprof)
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Profiling Tooling]]"
---

# Profiling Subsystems (pprof)

Sampling CPU profiling, heap/memory profiling, mutex contention, block profiling, and interactive pprof analysis.

```text
Profiling Subsystems (pprof)
│
├── [[CPU Profiling & Flamegraph Analysis]]
├── [[Heap & Memory Profiling (Alloc vs Inuse)]]
├── [[Mutex Contention Profiling (runtime.SetMutexProfileFraction)]]
├── [[Block Profiling (runtime.SetBlockProfileRate)]]
├── [[Goroutine Stack Dump Profiling]]
└── [[Interactive pprof CLI & Web Visualizer]]
```

---

## 🗂️ Topics

- [[CPU Profiling & Flamegraph Analysis]] — Sampling CPU profiler (100Hz default), interpreting flamegraphs, hot path identification, and overhead.
- [[Heap & Memory Profiling (Alloc vs Inuse)]] — alloc_space, alloc_objects, inuse_space, inuse_objects, diagnosing steady-state vs transient memory bloat.
- [[Mutex Contention Profiling (runtime.SetMutexProfileFraction)]] — Measuring lock wait times and lock acquisition contention hotspots in production.
- [[Block Profiling (runtime.SetBlockProfileRate)]] — Tracking goroutines blocked on channels, network I/O, and system calls.
- [[Goroutine Stack Dump Profiling]] — Capturing full stack traces of all live goroutines to detect deadlocks and goroutine leaks.
- [[Interactive pprof CLI & Web Visualizer]] — go tool pprof interactive commands (top, list, web, peek, disasm, source) and web visualizer.

---

## 🔗 References
- ⬆️ Parent: `Performance Engineering & Profiling`

