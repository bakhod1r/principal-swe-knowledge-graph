---
title: Debugging
tags:
  - programming
  - diagnostics
  - principal-swe
parent: "[[Diagnostics]]"
---

# Debugging

Debuggers mechanics (INT 3 traps, PTRACE, DWARF), deterministic time-travel record/replay, live process inspection, and DAP protocol.

```text
Debugging
│
├── [[Debuggers Mechanics (INT 3 Traps, PTRACE, DWARF Location Expressions)]]
├── [[Deterministic Time-Travel Debugging (Record and Replay)]]
├── [[Live Process Inspection and Hot Patching in Production]]
└── [[Remote Debugging Protocols (DAP - Debug Adapter Protocol)]]
```

---

## 🗂️ Topics

- [[Debuggers Mechanics (INT 3 Traps, PTRACE, DWARF Location Expressions)]] — How debuggers set software breakpoints, inspect registers, and evaluate variables in memory.
- [[Deterministic Time-Travel Debugging (Record and Replay)]] — Recording nondeterministic multi-threaded executions and stepping backwards to isolate transient bugs.
- [[Live Process Inspection and Hot Patching in Production]] — Attaching to live running containers without interrupting production client traffic.
- [[Remote Debugging Protocols (DAP - Debug Adapter Protocol)]] — Standardizing IDE debugger clients with remote runtime daemons across heterogeneous language runtimes.

---

## 🔗 References
- ⬆️ Parent: [[Diagnostics]]

