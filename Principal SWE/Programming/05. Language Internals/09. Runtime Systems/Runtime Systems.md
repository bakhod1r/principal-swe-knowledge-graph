---
title: Runtime Systems
tags:
  - programming
  - language-internals
  - principal-swe
parent: "[[Language Internals]]"
---

# Runtime Systems

Comprehensive engineering guide, patterns, and principles for Runtime Systems.

```text
Runtime Systems
│
├── [[Object Model and Layout]]
├── [[Method Dispatch and Inline Caches]]
├── [[Dynamic Linking and Loading]]
├── [[JIT Compilation and Tiering]]
├── [[Deoptimization and Speculation]]
├── [[Stack Management and Unwinding]]
└── [[Runtime GC Integration]]
```

---

## 🗂️ Topics

- [[Object Model and Layout]]
- [[Method Dispatch and Inline Caches]]
- [[Dynamic Linking and Loading]]
- [[JIT Compilation and Tiering]]
- [[Deoptimization and Speculation]]
- [[Stack Management and Unwinding]]
- [[Runtime GC Integration]]

- [[Polymorphic Inline Caching (PIC) and Megamorphic Call Sites]] — How JIT compilers accelerate dynamic virtual dispatch using monomorphic, polymorphic, and megamorphic inline cache stubs.
- [[Safepoints, Stop-the-World Signals, and Thread Suspension]] — How runtimes inject polling checks and OS signals (SIGSEGV/SIGUSR1) to bring threads to consistent execution safepoints.

---

## 🔗 References
- ⬆️ Parent: [[Language Internals]]

