---
title: Array
tags:
  - golang
  - arrays
  - composite-types
  - principal-swe
parent: "[[Composite Types]]"
---

# Array

Fixed-length arrays, contiguous memory layout, value semantics, comparison, stack buffers, BCE optimizations, and pointer semantics.

```text
Array
│
├── [[Fixed Length & Contiguous Memory]]
├── [[Array Memory Layout]]
├── [[Array Pass-by-Value Semantics]]
├── [[Array Pointer vs Array Value (Pointer-to-Array vs Array)]]
├── [[Multi-Dimensional Arrays]]
├── [[Array Comparison (==)]]
├── [[Array Slicing to Slice Header]]
├── [[Array Key in Maps (Fixed-Size Byte Arrays as High-Performance Map Keys)]]
├── [[Compile-Time Array Bounds Checking & Bound Check Elimination (BCE)]]
├── [[Constant-Sized Buffers & Scratch Memory on Stack]]
└── [[Zero-Length Arrays (Zero-Sized Array Types)]]
```

---

## 🗂️ Topics

- [[Fixed Length & Contiguous Memory]] — Fixed-size contiguous memory sequences where length is part of the type (`[N]T`).
- [[Array Memory Layout]] — Zero-overhead sequential layout on stack or heap with zero padding between elements.
- [[Array Pass-by-Value Semantics]] — Copying the entire memory contents of the array upon function call or assignment.
- [[Array Pointer vs Array Value (Pointer-to-Array vs Array)]] — Passing array pointers (`*[N]T`) to prevent $O(N)$ stack copies and enable in-place mutation.
- [[Multi-Dimensional Arrays]] — Matrices and multi-dimensional grids (`[M][N]T`) in row-major contiguous memory.
- [[Array Comparison (==)]] — Direct value equality comparison (`arr1 == arr2`) when element types are comparable.
- [[Array Slicing to Slice Header]] — Creating a 3-word slice header referencing an array (`arr[:]` and `&arr`).
- [[Array Key in Maps (Fixed-Size Byte Arrays as High-Performance Map Keys)]] — Using fixed-size `[16]byte` or `[32]byte` arrays as zero-alloc hash map keys.
- [[Compile-Time Array Bounds Checking & Bound Check Elimination (BCE)]] — How the compiler eliminates runtime panic checks on array indexing.
- [[Constant-Sized Buffers & Scratch Memory on Stack]] — Allocating temporary scratch buffers on the goroutine stack without heap escapes.
- [[Zero-Length Arrays (Zero-Sized Array Types)]] — Zero-byte array types, `unsafe.Sizeof([0]T{}) == 0`, and struct field behaviors.

---

## 🔗 References
- ⬆️ Parent: [[Composite Types]]
- 📚 Module: `Language Basics`
