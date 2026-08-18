---
title: Array (Basic Data Structures)
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Array Data Structure & Complete Operations Catalog

Exhaustive, hardware-conscious mastery of contiguous memory arrays: physical pointer arithmetic, SIB register indexing, geometric capacity growth (1.5x vs 2.0x), amortized append, order-preserving vs fast unordered removals, zero-copy slicing, SIMD memory copies, in-place rotations/compactions, deduplication, Lomuto/Hoare partitioning, compiler bounds checking elimination (BCE), CPU cache line prefetching, False Sharing, SoA vs AoS, bit arrays, and memory-mapped files across 25 granular operational blueprints.

```text
Array
│
├── [[Array Memory Layout and Pointer Arithmetic]]
├── [[Array Access and Indexing Mechanics]]
├── [[Array Append and Dynamic Capacity Allocation]]
├── [[Array Pre-Allocation and Capacity Hints]]
├── [[Array Insert at Index (Order-Preserving Shift)]]
├── [[Array Delete at Index (Order-Preserving Shift)]]
├── [[Array Delete Fast (Unordered Swap-and-Pop)]]
├── [[Array Clear and Logical vs Physical Truncation]]
├── [[Array Search (Linear vs Branchless Binary Search)]]
├── [[Array Slicing, Windows, and Full 3-Index Expressions]]
├── [[Array Memory Copy (SIMD rep movsq and memmove)]]
├── [[Array Reverse and In-Place 3-Reversal Rotation]]
├── [[Array Filter and In-Place Two-Pointer Compaction]]
├── [[Array Deduplication on Sorted Buffers]]
├── [[Array Partitioning (Lomuto vs Hoare Partition)]]
├── [[Array Push and Pop (LIFO Dynamic Stack Operations)]]
├── [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]]
├── [[Array Memory Alignment, Cache Line Striding, and False Sharing]]
├── [[Array Bounds Checking Elimination (BCE) Compiler Optimization]]
├── [[Array Memory Leaks and Garbage Collection Truncation]]
├── [[Multidimensional Array Striding (Row-Major vs Column-Major)]]
├── [[Structure of Arrays (SoA) vs Array of Structures (AoS)]]
├── [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]]
├── [[Bit Array (Packed Bitset Operations)]]
└── [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]]
```

---

## 🗂️ Array Operations & Topics

- [[Array Memory Layout and Pointer Arithmetic]] — Linear contiguous byte addressing, scale-index-base translation, and memory offsets.
- [[Array Access and Indexing Mechanics]] — O(1) random-access mechanics, SIB register decoding, and zero-overhead lookups.
- [[Array Append and Dynamic Capacity Allocation]] — Amortized O(1) growth, allocation doubling, and copy overhead profiles.
- [[Array Pre-Allocation and Capacity Hints]] — Zero-reallocation ingestion via size hinting and heap reserve sizing.
- [[Array Insert at Index (Order-Preserving Shift)]] — O(N) middle insertion requiring rightward memory block displacement.
- [[Array Delete at Index (Order-Preserving Shift)]] — O(N) element removal preserving sequential ordering via leftward shift.
- [[Array Delete Fast (Unordered Swap-and-Pop)]] — O(1) constant time deletion by swapping target slot with the tail element.
- [[Array Clear and Logical vs Physical Truncation]] — O(1) length reset vs O(N) memory zeroing and GC reclamation.
- [[Array Search (Linear vs Branchless Binary Search)]] — O(N) sequential scans vs O(log N) branchless binary search on sorted data.
- [[Array Slicing, Windows, and Full 3-Index Expressions]] — Zero-copy pointer windowing and 3-index capacity boundary capping.
- [[Array Memory Copy (SIMD rep movsq and memmove)]] — Hardware vectorized block transfers, AVX-512 alignment, and overlapping buffers.
- [[Array Reverse and In-Place 3-Reversal Rotation]] — Two-pointer reversal and O(1) space block rotation algorithm.
- [[Array Filter and In-Place Two-Pointer Compaction]] — Zero-allocation compaction using fast-read and slow-write pointer pairs.
- [[Array Deduplication on Sorted Buffers]] — O(N) single-pass deduplication with zero allocations on sorted sequences.
- [[Array Partitioning (Lomuto vs Hoare Partition)]] — In-place pivot partitioning schemes for QuickSort and QuickSelect.
- [[Array Push and Pop (LIFO Dynamic Stack Operations)]] — Strict O(1) LIFO operations on dynamic arrays with top pointer tracking.
- [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]] — Mathematical amortization proofs, allocator reuse, and memory fragmentation.
- [[Array Memory Alignment, Cache Line Striding, and False Sharing]] — Hardware memory bus alignment, 64-byte L1 cache line prefetching, and false sharing.
- [[Array Bounds Checking Elimination (BCE) Compiler Optimization]] — Compiler optimization eliminating runtime bounds check branch overheads.
- [[Array Memory Leaks and Garbage Collection Truncation]] — Preventing memory retention by pointer zeroing and sub-slice truncation.
- [[Multidimensional Array Striding (Row-Major vs Column-Major)]] — Index mapping formulas, stride penalties, and hardware prefetcher optimization.
- [[Structure of Arrays (SoA) vs Array of Structures (AoS)]] — Data-oriented design, SIMD memory packing, and cache efficiency.
- [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]] — Modulo-free circular queue wrapping using power-of-two bitwise masks.
- [[Bit Array (Packed Bitset Operations)]] — Dense 64-to-1 memory compression using 64-bit word bit manipulation.
- [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]] — Operating system virtual memory page table mapping for multi-gigabyte datasets.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: [[Data Structures & Algorithms]]
- 🎓 Root: [[Principal SWE]]
