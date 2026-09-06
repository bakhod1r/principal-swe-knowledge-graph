---
title: Array Algorithms & Operations
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Array Algorithms & Operations

Essential in-place algorithms, memory shifts, branchless binary search, prefix sums, difference arrays, two-pointer compactions, and partitioning.

```text
Array Algorithms & Operations
│
├── [[Array Insert at Index (Order-Preserving Shift)]]
├── [[Array Delete at Index (Order-Preserving Shift)]]
├── [[Array Delete Fast (Unordered Swap-and-Pop)]]
├── [[Array Search (Linear vs Branchless Binary Search)]]
├── [[Array Slicing, Windows, and Full 3-Index Expressions]]
├── [[Array Reverse and In-Place 3-Reversal Rotation]]
├── [[Array Filter and In-Place Two-Pointer Compaction]]
├── [[Array Deduplication on Sorted Buffers]]
├── [[Array Partitioning (Lomuto vs Hoare Partition)]]
├── [[Prefix Sum Array and Range Sum Queries]]
├── [[Difference Array and Range Update Range Query (RURQ)]]
├── [[Two-Pointer Sliding Window on Contiguous Buffers]]
└── [[Boyer-Moore Majority Vote In-Place Array Scan]]
```

---

## 🗂️ Topics & Implementations

- [[Array Insert at Index (Order-Preserving Shift)]] — O(N) middle insertion requiring rightward block shift via memmove.
- [[Array Delete at Index (Order-Preserving Shift)]] — O(N) element removal preserving sequential ordering via leftward shift.
- [[Array Delete Fast (Unordered Swap-and-Pop)]] — O(1) constant-time element deletion for order-agnostic collections via tail swap.
- [[Array Search (Linear vs Branchless Binary Search)]] — O(N) SIMD linear scans vs O(log N) branchless binary search eliminating mispredictions.
- [[Array Slicing, Windows, and Full 3-Index Expressions]] — Zero-copy pointer windowing and 3-index capacity boundary protection.
- [[Array Reverse and In-Place 3-Reversal Rotation]] — Two-pointer reversal and O(1) space block rotation algorithm.
- [[Array Filter and In-Place Two-Pointer Compaction]] — Zero-allocation compaction using fast-read and slow-write pointer pairs.
- [[Array Deduplication on Sorted Buffers]] — O(N) single-pass deduplication with zero allocations on sorted sequences.
- [[Array Partitioning (Lomuto vs Hoare Partition)]] — In-place pivot partitioning schemes for QuickSort and QuickSelect algorithms.
- [[Prefix Sum Array and Range Sum Queries]] — O(1) immutable range sum calculations with O(N) preprocessing.
- [[Difference Array and Range Update Range Query (RURQ)]] — O(1) range additions across intervals with O(N) final reconstruction.
- [[Two-Pointer Sliding Window on Contiguous Buffers]] — Dynamic and fixed-size contiguous subarray optimizations with monotonic queues.
- [[Boyer-Moore Majority Vote In-Place Array Scan]] — O(N) time and O(1) space single-pass dominant element identification algorithm.

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
