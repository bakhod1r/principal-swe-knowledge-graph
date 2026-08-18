---
title: FFI & Low-Level Assembly
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# FFI & Low-Level Assembly

Cgo foreign function interface, Plan 9 assembly language, SIMD intrinsics, and build constraints.

```text
FFI & Low-Level Assembly
│
├── [[Cgo Architecture & Overhead]]
├── [[Plan 9 Go Assembly]]
├── [[Go ABIInternal Register Calling Convention]]
└── [[CPU Feature Detection (internal-cpu)]]
```

---

## 🗂️ Topics

- [[Cgo Architecture & Overhead]] — Calling C libraries from Go, Cgo stack switching overhead, memory pinning rules (runtime.Pinner).
- [[Plan 9 Go Assembly]] — Writing Go assembly functions, pseudo-registers (FP, SP, SB, PC), instruction syntax.
- [[Go ABIInternal Register Calling Convention]] — Passing function arguments and returns in CPU registers instead of stack frames.
- [[CPU Feature Detection (internal-cpu)]] — Detecting hardware AVX, SSE, AES-NI CPU instructions at runtime.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]
- 🎓 Root: [[Principal SWE]]
