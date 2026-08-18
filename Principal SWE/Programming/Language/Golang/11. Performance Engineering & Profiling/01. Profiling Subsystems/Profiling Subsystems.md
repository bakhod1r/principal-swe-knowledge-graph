---
title: Profiling Subsystems
tags:
  - golang
  - performance
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
