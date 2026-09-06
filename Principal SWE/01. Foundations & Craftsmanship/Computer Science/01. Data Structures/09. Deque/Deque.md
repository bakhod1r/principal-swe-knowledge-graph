---
title: Deque
tags:
  - review
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Deque

Double-Ended Queue architectures: circular buffers, chunked memory blocks, work-stealing scheduling, and sliding window monotonic queues.

```text
Deque
│
├── 01. Circular Array Deque
│   ├── [[Circular Array Deque]]
│   ├── [[Deque Invariants and Double-Ended Contract]]
│   ├── [[Deque Push Front and Pop Front Operations]]
│   ├── [[Deque Push Back and Pop Back Operations]]
│   ├── [[Deque Circular Array Buffer Implementation]]
│   ├── [[Circular Deque Power-of-Two Bitwise Indexing for Head and Tail]]
│   └── [[Circular Deque Dynamic Geometric Reallocation and Unwrapping]]
││
├── 02. Chunked Block Deque
│   ├── [[Chunked Block Deque]]
│   ├── [[Chunked Block Array Deque Architecture (std::deque)]]
│   ├── [[std::deque Fixed-Size Buffer Map and Iterator Invalidation]]
│   └── [[Chunked Block Deque Random Access Offset Math (Map Index + Buffer Offset)]]
││
├── 03. Work-Stealing Deque
│   ├── [[Work-Stealing Deque]]
│   ├── [[Work-Stealing Deque Architecture (Chase-Lev Algorithm)]]
│   ├── [[Chase-Lev Work-Stealing Deque Lock-Free Algorithm]]
│   └── [[Fork-Join Frameworks and Work Balancing via Bottom-Push and Top-Steal]]
││
└── 04. Monotonic Deque Algorithms
    ├── [[Monotonic Deque]]
    ├── [[Sliding Window Maximum in O(N) via Monotonic Decreasing Deque]]
    └── [[Shortest Subarray with Sum at Least K via Monotonic Deque]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Circular Array Deque|01. Circular Array Deque]]
- [[Circular Array Deque]] — Master topology: Circular buffer allowing O(1) push and pop at both ends.
- [[Deque Invariants and Double-Ended Contract]] — Formal invariant: O(1) insertions and deletions at both head and tail.
- [[Deque Push Front and Pop Front Operations]] — O(1) head manipulations with circular index decrement and wraparound.
- [[Deque Push Back and Pop Back Operations]] — O(1) tail manipulations with circular index increment and wraparound.
- [[Deque Circular Array Buffer Implementation]] — Bidirectional modulo arithmetic on contiguous arrays.
- [[Circular Deque Power-of-Two Bitwise Indexing for Head and Tail]] — Masking head and tail pointers with (capacity - 1) for zero-cost wraparound.
- [[Circular Deque Dynamic Geometric Reallocation and Unwrapping]] — Reordering wrapped head-tail elements into linear order during array doubling.

### 2. 📂 [[Chunked Block Deque|02. Chunked Block Deque]]
- [[Chunked Block Deque]] — Master topology: std::deque architecture: array of fixed-size chunks.
- [[Chunked Block Array Deque Architecture (std::deque)]] — Central map of pointers to fixed-size (e.g. 512B) memory blocks.
- [[std::deque Fixed-Size Buffer Map and Iterator Invalidation]] — Why std::deque insertions do not invalidate pointers to existing elements.
- [[Chunked Block Deque Random Access Offset Math (Map Index + Buffer Offset)]] — Two-tier pointer indirection arithmetic achieving O(1) random access: index / block_size.

### 3. 📂 [[Work-Stealing Deque|03. Work-Stealing Deque]]
- [[Work-Stealing Deque]] — Master topology: Multi-threaded task scheduling deque (Chase-Lev).
- [[Work-Stealing Deque Architecture (Chase-Lev Algorithm)]] — Worker thread pushes/pops LIFO from bottom; victim threads steal FIFO from top.
- [[Chase-Lev Work-Stealing Deque Lock-Free Algorithm]] — Memory fences and atomic CAS preventing race conditions during concurrent steals.
- [[Fork-Join Frameworks and Work Balancing via Bottom-Push and Top-Steal]] — Load balancing across multicore CPUs in modern runtimes (Go, Java ForkJoinPool).

### 4. 📂 [[Monotonic Deque|04. Monotonic Deque Algorithms]]
- [[Monotonic Deque]] — Master topology: Monotonically ordered deque for linear time sliding window optimizations.
- [[Sliding Window Maximum in O(N) via Monotonic Decreasing Deque]] — Maintaining potential maxima in a sliding window with O(1) amortized cost per element.
- [[Shortest Subarray with Sum at Least K via Monotonic Deque]] — Prefix sum monotonicity with double-ended queue pruning in O(N) time.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`
