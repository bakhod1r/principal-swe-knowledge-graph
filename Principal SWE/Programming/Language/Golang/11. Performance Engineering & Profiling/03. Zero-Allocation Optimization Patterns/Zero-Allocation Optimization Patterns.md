---
title: Zero-Allocation Optimization Patterns
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Performance Engineering & Profiling]]"
---

# Zero-Allocation Optimization Patterns

Buffer recycling with sync.Pool, zero-copy casting, pre-allocation, struct alignment, and boxing elimination.

```text
Zero-Allocation Optimization Patterns
│
├── [[Buffer Recycling with sync.Pool]]
├── [[Zero-Copy String and Byte Conversions]]
├── [[Pre-Allocating Capacity in Slices and Maps]]
├── [[Struct Field Alignment & Padding Elimination]]
├── [[Avoiding Interface Boxing in Hot Loops]]
└── [[Stack vs Heap Optimization via Inlining]]
```

---

## 🗂️ Topics

- [[Buffer Recycling with sync.Pool]] — Eliminating heap churn with per-P byte slice and struct pools in high-throughput services.
- [[Zero-Copy String and Byte Conversions]] — Safe zero-allocation string/byte conversions via unsafe.String and unsafe.Slice.
- [[Pre-Allocating Capacity in Slices and Maps]] — Eliminating dynamic growth reallocations via make([]T, 0, cap) and make(map[K]V, hint).
- [[Struct Field Alignment & Padding Elimination]] — Ordering fields from largest to smallest to eliminate 8-byte word padding waste.
- [[Avoiding Interface Boxing in Hot Loops]] — Passing concrete types to prevent runtime iface allocations in high-throughput hot paths.
- [[Stack vs Heap Optimization via Inlining]] — Structuring leaf functions within inlining budgets (80 AST nodes) to keep objects on the stack.

---

## 🔗 References
- ⬆️ Parent: [[Performance Engineering & Profiling]]
- 🎓 Root: [[Principal SWE]]
