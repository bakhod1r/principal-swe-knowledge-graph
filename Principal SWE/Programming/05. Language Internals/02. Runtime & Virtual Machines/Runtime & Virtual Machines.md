---
title: Runtime & Virtual Machines
tags:
  - programming
  - language-internals
  - principal-swe
parent: "[[Language Internals]]"
---

# Runtime & Virtual Machines

Bytecode interpreters and JIT execution engines.

```text
Runtime & Virtual Machines
│
├── [[Stack-Based vs Register-Based Virtual Machines (JVM, V8, BEAM)]]
├── [[Just-In-Time (JIT) Compilation & Tiered Compilation]]
└── [[Interpreter Execution Loops & Bytecode Deserialization]]
```

---

## 🗂️ Topics

- [[Stack-Based vs Register-Based Virtual Machines (JVM, V8, BEAM)]] — Comparing operand stack execution models with virtual register architectures for bytecode interpreters.
- [[Just-In-Time (JIT) Compilation & Tiered Compilation]] — Dynamic profiling, baseline JITs, optimizing JITs (HotSpot C2, V8 TurboFan), and deoptimization bailouts.
- [[Interpreter Execution Loops & Bytecode Deserialization]] — Direct threading, indirect threading, switch-based interpreters, and opcode dispatch optimizations.

---

## 🔗 References
- ⬆️ Parent: [[Language Internals]]
- 🎓 Root: [[Principal SWE]]
