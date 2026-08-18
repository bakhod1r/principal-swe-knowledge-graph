---
title: Plan 9 Go Assembly & Hardware Architecture
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# Plan 9 Go Assembly & Hardware Architecture

Plan 9 assembly syntax, stack frame layouts, ABIInternal register calling convention, SIMD vector instructions, and CPU feature detection.

```text
Plan 9 Go Assembly & Hardware Architecture
│
├── [[Plan 9 Assembly Syntax & Architecture Differences]]
├── [[Stack Frame Layout in Go Assembly Functions]]
├── [[Go ABIInternal Register Calling Convention Mechanics]]
├── [[Writing SIMD Vector Instructions in Go Assembly]]
└── [[CPU Feature Detection (internal-cpu & CPUID)]]
```

---

## 🗂️ Topics

- [[Plan 9 Assembly Syntax & Architecture Differences]] — Pseudo-registers (FP, SP, SB, PC), instruction operands, and Plan 9 assembly conventions.
- [[Stack Frame Layout in Go Assembly Functions]] — Writing assembly functions: stack size, argument size declarations (TEXT ·Add(SB), $0-24), and returns.
- [[Go ABIInternal Register Calling Convention Mechanics]] — Passing integer arguments in RAX, RBX, RCX and floating-point in X0-X14 registers.
- [[Writing SIMD Vector Instructions in Go Assembly]] — Utilizing AVX2, AVX-512, and ARM NEON vectorized instructions for high-throughput computing.
- [[CPU Feature Detection (internal-cpu & CPUID)]] — Detecting runtime CPU capabilities (cpu.X86.HasAVX2, cpu.ARM64.HasAES) to dispatch optimized paths.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]

