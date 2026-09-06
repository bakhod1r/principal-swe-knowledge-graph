---
title: "Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)

## 1. Definition
There are two primary paradigms to represent 2D/3D grids in high-level programming languages:

1. **Contiguous Flat Buffer (`T*` with Strides)**: A single monolithic heap block of size $M \times N \times \text{sizeof}(T)$. Rows are placed contiguously.
2. **Jagged Array / Array of Pointers (`T**`)**: An array of $M$ pointers, where each pointer addresses an independently allocated array of $N$ elements.

## 2. Mental Model & Architecture Comparison

```text
Contiguous Flat Buffer:
Buffer: [ Row 0 ][ Row 1 ][ Row 2 ]  <-- Single malloc(M * N * sizeof(T))
Single base pointer, zero indirection, 100% cache locality.

Jagged Array (Array of Pointers):
Pointers: [ ptr0 ][ ptr1 ][ ptr2 ]
             |       |       |
             v       v       v
         [Row 0]  [Row 1]  [Row 2]  <-- M separate mallocs in scattered heap
Requires 2 memory loads per element access; pointers thrash TLB and L1 cache.
```

## 3. Trade-off Matrix

| Metric | Contiguous Flat Buffer | Jagged Array (`T**`) |
| :--- | :--- | :--- |
| **Allocation Cost** | $O(1)$ single `malloc` | $O(M)$ separate `malloc` calls |
| **Memory Overhead** | $0$ auxiliary bytes | $M \times 8$ bytes for pointers + allocator headers |
| **Dereference Latency** | 1 pointer arithmetic + 1 load | 2 pointer dereferences (dependent load stall) |
| **Hardware Prefetch** | Seamless L1/L2 stream prefetch | Constant prefetcher pipeline invalidation |
| **SIMD Vectorization** | Vectorizable across row boundaries | Cannot vectorize across discontinuous rows |
| **Row Length Flexibility**| Fixed rectangular grid | Supports ragged/variable-length rows |

## 4. Gotchas
- In Java and C#, `int[][]` is always a jagged array of references. In high-performance C# code, prefer `Span2D` or flat `int[]` with arithmetic indexing.
- Freeing jagged memory requires traversing all $M$ pointers. Freeing flat memory is a single `free(buf)`.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
