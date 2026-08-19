---
title: Slices
tags:
  - golang
  - slices
  - composite-types
  - principal-swe
parent: "[[Composite Types]]"
---

# Slices

Dynamic slice headers, capacity growth rules, make preallocation, conversions, operations, slice tricks, BCE optimizations, and memory safety.

```text
Slices
│
├── [[Slice Header Internals]]
├── [[Slice Capacity and Growth]]
├── [[Slice Capacity Growth Algorithm Formula (Go 1.19+ vs Legacy)]]
├── [[make() for Slices]]
├── [[Full Slice Expressions (3-Index Slicing)]]
├── [[Slice to Array Conversion]]
├── [[Array to Slice Conversion]]
├── [[Slice Append]]
├── [[Slice Copy]]
├── [[Slice Cut]]
├── [[Slice Insert]]
├── [[Slice Delete (Preserving Order)]]
├── [[Slice Delete (Unordered Swap-and-Pop)]]
├── [[Slice Push and Pop]]
├── [[Slice Filter (In-Place)]]
├── [[Slice Reverse and Rotate]]
├── [[Slice Deduplication and Batching]]
├── [[Slice Memory Leaks & GC Truncation]]
├── [[Zero-Copy Slice Reslicing and Overlap Hazards]]
├── [[Slice Bounds Check Elimination (BCE) Deep Dive]]
├── [[Slices vs Arrays Escape Analysis & Stack Allocation]]
├── [[sync.Pool of Byte Slices ([]byte Reuse & Buffer Pools)]]
└── [[slices Standard Package (Go 1.21+)]]
```

---

## 🗂️ Topics

- [[Slice Header Internals]] — `sliceHeader` struct: `unsafe.Pointer Data`, `int Len`, `int Cap` (24 bytes on 64-bit).
- [[Slice Capacity and Growth]] — Dynamic append allocation formula, doubling threshold, and growth dynamics.
- [[Slice Capacity Growth Algorithm Formula (Go 1.19+ vs Legacy)]] — Detailed breakdown of Go 1.19+ smooth growth formula `(oldCap + 3*256)/4`.
- [[make() for Slices]] — Pre-allocating slice length and capacity (`make([]T, len, cap)`).
- [[Full Slice Expressions (3-Index Slicing)]] — Limiting capacity with `s[low:high:max]` to prevent unintended mutation of shared backing arrays.
- [[Slice to Array Conversion]] — Pointer-to-array conversions (`*[N]T(s)`) and Go 1.20 slice-to-array casting (`[N]T(s)`).
- [[Array to Slice Conversion]] — Creating a slice viewing an entire or partial fixed array (`arr[:]`).
- [[Slice Append]] — `append()` builtin mechanics, reallocation, and assigning back to destination slice.
- [[Slice Copy]] — `copy(dst, src)` builtin, copying elements between overlapping or disjoint slices, `min(len(dst), len(src))`.
- [[Slice Cut]] — Deleting a range `[i:j]` and preserving removed elements (`append(s[:i], s[j:]...)`).
- [[Slice Insert]] — Inserting single or multiple elements at index `i` with allocation profile.
- [[Slice Delete (Preserving Order)]] — Removing element at index `i` while keeping slice element order (`append(s[:i], s[i+1:]...)`).
- [[Slice Delete (Unordered Swap-and-Pop)]] — $O(1)$ deletion by swapping target index with last element.
- [[Slice Push and Pop]] — Implementing LIFO Stack and FIFO Queue push/pop operations with slices.
- [[Slice Filter (In-Place)]] — Zero-allocation filtering using two-pointer in-place iteration.
- [[Slice Reverse and Rotate]] — Reversing slice in-place and rotating elements left/right by `k` positions.
- [[Slice Deduplication and Batching]] — Deduplicating adjacent elements on sorted slices and chunking large slices into batches.
- [[Slice Memory Leaks & GC Truncation]] — Preventing memory retention by pointer zeroing during deletion and avoiding sub-slice references to large arrays.
- [[Zero-Copy Slice Reslicing and Overlap Hazards]] — Concurrency race conditions and silent overwrite bugs when reslicing shared backing arrays.
- [[Slice Bounds Check Elimination (BCE) Deep Dive]] — SSA compiler analysis passes for proving slice boundaries and eliminating runtime panic branches.
- [[Slices vs Arrays Escape Analysis & Stack Allocation]] — Compiler stack allocation vs heap escape analysis for slices.
- [[sync.Pool of Byte Slices ([]byte Reuse & Buffer Pools)]] — High-throughput zero-GC buffer recycling architectures.
- [[slices Standard Package (Go 1.21+)]] — Type-safe generic slice algorithms: `slices.Sort`, `slices.BinarySearch`, `slices.Contains`, `slices.Clone`, `slices.Delete`.

---

## 🔗 References
- ⬆️ Parent: [[Composite Types]]
- 📚 Module: `Language Basics`
