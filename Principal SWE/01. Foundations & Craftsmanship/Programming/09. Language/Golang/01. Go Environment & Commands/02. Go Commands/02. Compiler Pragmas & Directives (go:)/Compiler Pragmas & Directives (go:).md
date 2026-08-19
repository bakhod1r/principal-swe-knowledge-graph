---
title: Compiler Pragmas & Directives (go:)
tags:
  - golang
  - compiler-pragmas
  - toolchain
  - low-level
  - principal-swe
parent: "[[Go Commands]]"
---

# Compiler Pragmas & Directives (go:)

Compiler pragma directives controlling inlining, stack allocation, escape analysis, symbol linking, race detection, WebAssembly FFI, and code generation.

```text
Compiler Pragmas & Directives (go:)
│
├── [[go:noinline Directive]]
├── [[go:noescape Directive]]
├── [[go:nosplit Directive]]
├── [[go:norace Directive]]
├── [[go:linkname Directive]]
├── [[go:uintptrescapes Directive]]
├── [[go:notinheap Directive]]
├── [[go:wasmimport & go:wasmexport Directives]]
├── [[go:generate Directive]]
└── [[go:embed Directive]]
```

---

## 🗂️ Topics

- [[go:noinline Directive]] — Preventing compiler from inlining functions for benchmarking and precise stack traces.
- [[go:noescape Directive]] — Informing escape analysis that pointer arguments do not escape to the heap.
- [[go:nosplit Directive]] — Disabling stack overflow check preamble in low-level runtime and assembly leaf functions.
- [[go:norace Directive]] — Suppressing ThreadSanitizer race detector instrumentation on performance-critical code.
- [[go:linkname Directive]] — Linking to unexported runtime and standard library functions across package boundaries.
- [[go:uintptrescapes Directive]] — Informing compiler that uintptr arguments escape to heap as memory pointers.
- [[go:notinheap Directive]] — Restricting type allocation to off-heap memory to bypass GC tracking entirely.
- [[go:wasmimport & go:wasmexport Directives]] — Defining WebAssembly host function imports and module exports.
- [[go:generate Directive]] — Declaring automated code generators (`stringer`, `mockery`, `protoc`) executed via `go generate`.
- [[go:embed Directive]] — Embedding static files, templates, and directory trees directly into binary executables.

---

## 🔗 References
- ⬆️ Parent: [[Go Commands]]
- 📚 Module: `Go Environment & Commands`
