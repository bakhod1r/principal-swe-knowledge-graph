---
title: Slices
tags:
  - golang
  - slices
  - principal-swe
parent: "[[Composite Types]]"
---

# Slices

Dynamic slice headers, capacity growth rules, make preallocation, conversions, operations, slice tricks, and memory safety.

```text
Slices
│
├── [[Slice Header Internals]]
├── [[Slice Capacity and Growth]]
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
└── [[slices Standard Package (Go 1.21+)]]
```

---

## 🗂️ Topics

- [[Slice Header Internals]] — sliceHeader struct: unsafe.Pointer Data, int Len, int Cap (24 bytes on 64-bit).
- [[Slice Capacity and Growth]] — Dynamic append allocation formula, doubling threshold (<256 elements), and 1.25x+192 smooth growth.
- [[make() for Slices]] — Pre-allocating slice length and capacity (make([]T, len, cap)).
- [[Full Slice Expressions (3-Index Slicing)]] — Limiting capacity with s[low:high:max] to prevent unintended mutation of shared backing arrays.
- [[Slice to Array Conversion]] — Pointer-to-array conversions (*[N]T(s)) and Go 1.20 slice-to-array casting ([N]T(s)).
- [[Array to Slice Conversion]] — Creating a slice viewing an entire or partial fixed array (arr[:]).
- [[Slice Append]] — append() builtin mechanics, reallocation, and assigning back to destination slice.
- [[Slice Copy]] — copy(dst, src) builtin, copying elements between overlapping or disjoint slices, min(len(dst), len(src)).
- [[Slice Cut]] — Deleting a range [i:j] and preserving removed elements (append(s[:i], s[j:]...)).
- [[Slice Insert]] — Inserting single or multiple elements at index i with allocation profile.
- [[Slice Delete (Preserving Order)]] — Removing element at index i while keeping slice element order (append(s[:i], s[i+1:]...)).
- [[Slice Delete (Unordered Swap-and-Pop)]] — O(1) deletion by swapping target index with last element (s[i] = s[len(s)-1]; s = s[:len(s)-1]).
- [[Slice Push and Pop]] — Implementing LIFO Stack and FIFO Queue push/pop operations with slices.
- [[Slice Filter (In-Place)]] — Zero-allocation filtering using two-pointer in-place iteration (b := s[:0]; for _, x := range s ...).
- [[Slice Reverse and Rotate]] — Reversing slice in-place and rotating elements left/right by k positions.
- [[Slice Deduplication and Batching]] — Deduplicating adjacent elements on sorted slices and chunking large slices into batches.
- [[Slice Memory Leaks & GC Truncation]] — Preventing memory retention by pointer zeroing during deletion and avoiding sub-slice references to large arrays.
- [[slices Standard Package (Go 1.21+)]] — Type-safe generic slice algorithms: slices.Sort, slices.BinarySearch, slices.Contains, slices.Clone, slices.Delete.

---

## 🔗 References
- ⬆️ Parent: [[Composite Types]]
- 🎓 Root: [[Principal SWE]]
