---
title: Compiler Optimizations & Devirtualization
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Methods & Interfaces]]"
---

# Compiler Optimizations & Devirtualization

Compiler devirtualization passes, virtual dispatch overhead, inlining barriers, and escape analysis.

```text
Compiler Optimizations & Devirtualization
│
├── [[Devirtualization Compiler Pass]]
├── [[Dynamic Interface Dispatch Overhead]]
├── [[Mid-Stack Inlining Barriers with Interfaces]]
├── [[Escape Analysis with Interface Boxing]]
└── [[Benchmarking Concrete Calls vs Interface Virtual Calls]]
```

---

## 🗂️ Topics

- [[Devirtualization Compiler Pass]] — How compiler statically proves concrete types and transforms indirect itab calls into direct static calls.
- [[Dynamic Interface Dispatch Overhead]] — Measuring the CPU penalty of indirect function pointer calls and CPU branch predictor misses.
- [[Mid-Stack Inlining Barriers with Interfaces]] — Why virtual interface calls prevent compiler function inlining across architectural boundaries.
- [[Escape Analysis with Interface Boxing]] — Why passing values to interface parameters (e.g. fmt.Println) frequently triggers heap escapes.
- [[Benchmarking Concrete Calls vs Interface Virtual Calls]] — Microbenchmarks comparing direct struct calls, devirtualized calls, and dynamic interface calls.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
