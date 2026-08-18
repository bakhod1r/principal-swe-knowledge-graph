- [[Continuous Profiling in Production (Pyroscope, Parca)]] — Fleet-wide continuous low-overhead CPU, heap, and goroutine profiling.

- [[Off-CPU Analysis with eBPF in Go]] — Measuring off-CPU blocking (I-O wait, lock contention, context switching) using Linux eBPF.

---
title: Profiling Subsystems
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Performance Engineering & Profiling]]"
---

# Profiling Subsystems

Capturing CPU flamegraphs, heap allocation profiles, mutex contention, and blocking profiles.

```text
Profiling Subsystems
│
├── [[CPU Profiling & Flamegraphs]]
├── [[Memory & Heap Profiling]]
├── [[Mutex Contention Profiling]]
├── [[Block Profiling]]
└── [[pprof Interactive Visualizer]]
```

---

## 🗂️ Topics

- [[CPU Profiling & Flamegraphs]] — Sampling CPU profiling (runtime/pprof, net/http/pprof), interpreting Flamegraphs and hot paths.
- [[Memory & Heap Profiling]] — inuse_space, inuse_objects, alloc_space, alloc_objects, diagnosing memory leaks.
- [[Mutex Contention Profiling]] — runtime.SetMutexProfileFraction, measuring lock acquisition delay and contention hotspots.
- [[Block Profiling]] — runtime.SetBlockProfileRate, identifying goroutines blocked on channels and network I/O.
- [[pprof Interactive Visualizer]] — pprof CLI commands (top, list, web, peek, disasm), web UI visualization.

---

## 🔗 References
- ⬆️ Parent: [[Performance Engineering & Profiling]]
- 🎓 Root: [[Principal SWE]]
