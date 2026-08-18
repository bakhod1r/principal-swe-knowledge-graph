---
title: Interactive Debugging & Tracing
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Interactive Debugging & Tracing

Breakpoint engines, post-mortem analysis, and binary tracing.

```text
Interactive Debugging & Tracing
│
├── [[Debuggers Mechanics (Breakpoints, Traps, DWARF Tables)]]
├── [[Time-Travel Debugging & Deterministic Record-Replay]]
├── [[Dynamic Binary Instrumentation (eBPF, DTrace)]]
└── [[Core Dump Forensics & Post-Mortem Crash Analysis]]
```

---

## 🗂️ Topics

- [[Debuggers Mechanics (Breakpoints, Traps, DWARF Tables)]] — How debuggers inject INT 3 traps, intercept ptrace signals, and resolve DWARF symbol tables.
- [[Time-Travel Debugging & Deterministic Record-Replay]] — Recording execution snapshots and replaying instruction streams backward to isolate non-deterministic bugs.
- [[Dynamic Binary Instrumentation (eBPF, DTrace)]] — Zero-overhead live kernel and user-space tracing without recompilation using eBPF probes.
- [[Core Dump Forensics & Post-Mortem Crash Analysis]] — Extracting memory state, register dumps, and stack traces from OS core dumps after process crashes.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]
- 🎓 Root: [[Principal SWE]]
