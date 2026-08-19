---
title: Runtime Tracing & Execution Profiling (trace.go, mprof.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Runtime Tracing & Execution Profiling (trace.go, mprof.go)

Trace event encoding, per-P trace buffers, heap/block/mutex sampling engines, and CPU profiler call-stack collection.

```text
Runtime Tracing & Execution Profiling (trace.go, mprof.go)
│
├── [[Trace Event Buffer Generation & Encoding (trace.go)]]
├── [[Per-P Trace Buffers & Event Flushing Mechanics]]
├── [[Memory Profiling Sampling Engine (mprof.go)]]
├── [[Block & Mutex Profile Recording Internals]]
└── [[CPU Profiler Sampling Clock & Call-Stack Collection]]
```

---

## 🗂️ Topics

- [[Trace Event Buffer Generation & Encoding (trace.go)]] — Binary execution trace format, event types (proc create, goroutine switch, syscall, GC phases).
- [[Per-P Trace Buffers & Event Flushing Mechanics]] — Lock-free event recording into local per-P byte buffers and flushing to global trace consumer.
- [[Memory Profiling Sampling Engine (mprof.go)]] — Poisson sampling distribution (1-in-512KB default) recording allocation call-stacks into mprof bucket tables.
- [[Block & Mutex Profile Recording Internals]] — Capturing stack traces and blocking durations for channel stalls and mutex contention events.
- [[CPU Profiler Sampling Clock & Call-Stack Collection]] — Extracting active goroutine call-stack on every SIGPROF tick and writing to pprof profile buffers.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]

