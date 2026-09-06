---
title: Array
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Array Data Structure (Memory Architectures & Concrete Types)

Comprehensive Principal SWE engineering catalog of contiguous memory arrays across 8 concrete types and specialized domains: physical pointer arithmetic, SIB register indexing, dynamic vector allocation, geometric growth policies (1.5x vs 2.0x), lock-free circular ring buffers, multi-dimensional matrices and voxel tensors, compressed sparse formats (COO, CSR, CSC), packed bitsets, system memory-mapped buffers, and in-place algorithms across 66 comprehensive operational blueprints.

```text
Array
│
├── 01. Static Array
│   ├── [[Static Array]]
│   ├── [[Array Memory Layout and Pointer Arithmetic]]
│   ├── [[Array Access and Indexing Mechanics]]
│   ├── [[Array Memory Alignment, Cache Line Striding, and False Sharing]]
│   ├── [[Array Bounds Checking Elimination (BCE) Compiler Optimization]]
│   ├── [[Array Memory Copy (SIMD rep movsq and memmove)]]
│   ├── [[Structure of Arrays (SoA) vs Array of Structures (AoS)]]
│   ├── [[Array Data-Oriented Design (DOD) and Cache Line Utilization]]
│   ├── [[Variable-Length Arrays (VLA) vs Fixed Stack Allocation]]
│   └── [[Array Address Space Layout Randomization (ASLR) and Stack Protection]]
││
├── 02. Dynamic Array
│   ├── [[Dynamic Array]]
│   ├── [[Array Append and Dynamic Capacity Allocation]]
│   ├── [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]]
│   ├── [[Array Pre-Allocation and Capacity Hints]]
│   ├── [[Array Push and Pop (LIFO Dynamic Stack Operations)]]
│   ├── [[Array Memory Leaks and Garbage Collection Truncation]]
│   ├── [[Array Clear and Logical vs Physical Truncation]]
│   ├── [[Array Shrink-to-Fit and Memory Compaction]]
│   └── [[Dynamic Array Allocator Arena Integration (jemalloc, TCMalloc)]]
││
├── 03. Circular Buffer Array
│   ├── [[Circular Buffer Array]]
│   ├── [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]]
│   ├── [[Circular Buffer Overwrite and Head-Tail Pointer Invariants]]
│   ├── [[Single-Producer Single-Consumer (SPSC) Lock-Free Array Ring Buffer]]
│   └── [[Disruptor Ring Buffer Pattern (Cache Line Padded Sequence Tracking)]]
││
├── 04. 2D & Multidimensional Array
│   ├── [[Matrix & Multidimensional Arrays]]
│   ├── [[Row-Major vs Column-Major Layout and Memory Striding]]
│   ├── [[3D Array Layout, Voxel Tensors, and Coordinate Flattening]]
│   ├── [[Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)]]
│   ├── [[Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching]]
│   ├── [[In-Place Matrix Transposition (Square vs Rectangular)]]
│   ├── [[Matrix 90-Degree In-Place Rotation (Transpose and Reverse)]]
│   ├── [[Spiral Matrix Traversal and Boundary Layer Peel]]
│   ├── [[Diagonal and Anti-Diagonal Matrix Traversals]]
│   ├── [[Cache-Oblivious Matrix Multiplication and Tiling]]
│   ├── [[3D Tensor Slicing, Sub-Volumes, and Dimensional Projections]]
│   └── [[Blocked Z-Morton Order (Morton Space-Filling Curve) Matrix Layout]]
││
├── 05. Sparse Matrix
│   ├── [[Sparse Matrix]]
│   ├── [[Sparse Matrix Coordinate List (COO) Representation]]
│   ├── [[Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)]]
│   ├── [[Diagonal Format (DIA) and Banded Matrix Storage]]
│   └── [[Sparse Matrix-Vector Multiplication (SpMV) Cache Optimization]]
││
├── 06. Bit Array
│   ├── [[Bit Array]]
│   ├── [[Bit Array (Packed Bitset Operations)]]
│   ├── [[Bitwise Word-Level Parallelism and Population Count (POPCNT)]]
│   ├── [[Bitset Set Algebra (SIMD 64-Bit Word Operations)]]
│   └── [[Roaring Bitmaps and Run-Length Encoded (RLE) Hybrid Bitsets]]
││
├── 07. Specialized System Arrays
│   ├── [[Specialized System Arrays]]
│   ├── [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]]
│   ├── [[Gap Buffer (Text Editor Cursor Buffer)]]
│   ├── [[Hashed Array Tree (HAT) Cache-Friendly Flat Allocation]]
│   ├── [[Suffix Array and Longest Common Prefix (LCP) Array]]
│   └── [[Non-Temporal Store Arrays (Direct-to-DRAM Streaming Stores)]]
││
└── 08. Array Algorithms & Operations
    ├── [[Array Algorithms & Operations]]
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

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Static Array|01. Static Array]]
- [[Static Array]] — Master topology: Contiguous physical memory buffer, compile-time/fixed capacity, pointer arithmetic, SIB register indexing, and compiler Bounds Checking Elimination (BCE).
- [[Array Memory Layout and Pointer Arithmetic]] — Linear contiguous byte addressing, scale-index-base translation, and pointer offsets.
- [[Array Access and Indexing Mechanics]] — O(1) random-access mechanics, hardware SIB register decoding, and zero-based offset proofs.
- [[Array Memory Alignment, Cache Line Striding, and False Sharing]] — CPU memory bus alignment, 64-byte L1 cache line prefetching, and multi-threaded false sharing.
- [[Array Bounds Checking Elimination (BCE) Compiler Optimization]] — Compiler optimization eliminating runtime bounds check branch overheads via loop induction analysis.
- [[Array Memory Copy (SIMD rep movsq and memmove)]] — Hardware vectorized block transfers, AVX-512 alignment, and overlapping buffer guarantees.
- [[Structure of Arrays (SoA) vs Array of Structures (AoS)]] — Data-oriented design, SIMD memory packing, and cache throughput optimizations.
- [[Array Data-Oriented Design (DOD) and Cache Line Utilization]] — Organizing contiguous arrays for maximum memory bandwidth and CPU pipelining efficiency.
- [[Variable-Length Arrays (VLA) vs Fixed Stack Allocation]] — Stack pointer adjustment runtime risks, alloca mechanics, and stack overflow pitfalls.
- [[Array Address Space Layout Randomization (ASLR) and Stack Protection]] — Buffer overflow exploit mitigations, stack canaries, and hardware memory protection.

### 2. 📂 [[Dynamic Array|02. Dynamic Array]]
- [[Dynamic Array]] — Master topology: Resizable heap-allocated contiguous vector, amortized geometric growth policies (1.5x vs 2.0x), size hints, and allocator arena interactions.
- [[Array Append and Dynamic Capacity Allocation]] — Amortized O(1) growth, allocation doubling proofs, and reallocation copy overhead profiles.
- [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]] — Mathematical amortization proofs, allocator memory reuse, and heap fragmentation economics.
- [[Array Pre-Allocation and Capacity Hints]] — Zero-reallocation ingestion via size hinting and heap reserve sizing.
- [[Array Push and Pop (LIFO Dynamic Stack Operations)]] — Strict O(1) LIFO operations on dynamic arrays with top pointer tracking.
- [[Array Memory Leaks and Garbage Collection Truncation]] — Preventing memory retention by pointer zeroing and sub-slice truncation in managed runtimes.
- [[Array Clear and Logical vs Physical Truncation]] — O(1) length reset vs O(N) memory zeroing, GC reclamation, and security zeroing.
- [[Array Shrink-to-Fit and Memory Compaction]] — Reclaiming unused capacity buffers via reallocation and hysteresis shrink thresholds.
- [[Dynamic Array Allocator Arena Integration (jemalloc, TCMalloc)]] — Interactions between dynamic vector resizing, size classes, and thread-local allocator caches.

### 3. 📂 [[Circular Buffer Array|03. Circular Buffer Array]]
- [[Circular Buffer Array]] — Master topology: Fixed-size contiguous ring array, modular wraparound arithmetic, power-of-two bitwise masking, and lock-free concurrency topologies.
- [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]] — Replacing hardware integer modulo (% N) with single-cycle bitwise AND masking (& (N - 1)).
- [[Circular Buffer Overwrite and Head-Tail Pointer Invariants]] — Disambiguating full vs empty buffer states via sequence counters and pointer offsets.
- [[Single-Producer Single-Consumer (SPSC) Lock-Free Array Ring Buffer]] — Wait-free thread synchronization utilizing atomic memory barriers and head/tail sequence tracking.
- [[Disruptor Ring Buffer Pattern (Cache Line Padded Sequence Tracking)]] — High-throughput inter-thread messaging architecture with cache line padding to eliminate false sharing.

### 4. 📂 [[Matrix & Multidimensional Arrays|04. 2D & Multidimensional Array]]
- [[Matrix & Multidimensional Arrays]] — Master topology: Contiguous flat buffers vs jagged arrays, row-major vs column-major striding, 3D voxel tensors, spatial cache locality, and matrix transformations.
- [[Row-Major vs Column-Major Layout and Memory Striding]] — Linear physical addressing formulas, stride steps, and cache traversal efficiency.
- [[3D Array Layout, Voxel Tensors, and Coordinate Flattening]] — 3D volumetric indexing `D*H*W`, spatial grids, and tensor flattening equations.
- [[Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)]] — Single-allocation buffers vs array-of-pointers (`T**`), pointer overhead, and cache misses.
- [[Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching]] — Hardware L1/L2 streamer prefetching and stride-N cache miss penalties.
- [[In-Place Matrix Transposition (Square vs Rectangular)]] — Diagonal swap symmetry and cycle-following permutation algorithms in O(1) space.
- [[Matrix 90-Degree In-Place Rotation (Transpose and Reverse)]] — Layered rotations combining matrix transpose and row reversal.
- [[Spiral Matrix Traversal and Boundary Layer Peel]] — Four-pointer bounding-box contraction traversal across 2D grids.
- [[Diagonal and Anti-Diagonal Matrix Traversals]] — Index invariants (i + j = k, i - j = k) driving zigzag and diagonal passes.
- [[Cache-Oblivious Matrix Multiplication and Tiling]] — Recursive block partitioning matching hardware cache hierarchy boundaries.
- [[3D Tensor Slicing, Sub-Volumes, and Dimensional Projections]] — Zero-copy strided views and axial/sagittal/coronal hyperplane slicing.
- [[Blocked Z-Morton Order (Morton Space-Filling Curve) Matrix Layout]] — Bit-interleaved Morton coding improving 2D spatial locality for arbitrary access patterns.

### 5. 📂 [[Sparse Matrix|05. Sparse Matrix]]
- [[Sparse Matrix]] — Master topology: Memory-compressed sparse array representations, coordinate triplets (COO), compressed row/column storage (CSR/CSC), and banded matrices.
- [[Sparse Matrix Coordinate List (COO) Representation]] — Triplet array storage (row, col, value) for incremental sparse construction.
- [[Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)]] — Industrial compressed representations enabling high-performance SpMV operations.
- [[Diagonal Format (DIA) and Banded Matrix Storage]] — Compact storage for banded diagonal systems and finite-difference stencils.
- [[Sparse Matrix-Vector Multiplication (SpMV) Cache Optimization]] — Optimizing indirect memory access and vectorization for sparse matrix computations.

### 6. 📂 [[Bit Array|06. Bit Array]]
- [[Bit Array]] — Master topology: Dense 64-to-1 memory compression using 64-bit word arrays, word-level SIMD boolean algebra, hardware population count (POPCNT), and bitmap indexing.
- [[Bit Array (Packed Bitset Operations)]] — Bitwise address translation: word offset (i >> 6) and bit shift (1ULL << (i & 63)).
- [[Bitwise Word-Level Parallelism and Population Count (POPCNT)]] — Processing 64 elements per CPU cycle and hardware Hamming weight calculation.
- [[Bitset Set Algebra (SIMD 64-Bit Word Operations)]] — Vectorized set union (OR), intersection (AND), and difference (AND NOT) across bit arrays.
- [[Roaring Bitmaps and Run-Length Encoded (RLE) Hybrid Bitsets]] — Adaptive compressed bitmaps switching between bitset, array, and RLE containers.

### 7. 📂 [[Specialized System Arrays|07. Specialized System Arrays]]
- [[Specialized System Arrays]] — Master topology: OS virtual memory mapped arrays, text editor gap buffers, hashed array trees, suffix arrays, and non-temporal streaming store arrays.
- [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]] — Virtual memory page mapping for multi-gigabyte files bypassing user-space buffer copies.
- [[Gap Buffer (Text Editor Cursor Buffer)]] — Contiguous buffer with movable cursor gap enabling O(1) text insertions and deletions.
- [[Hashed Array Tree (HAT) Cache-Friendly Flat Allocation]] — Array of leaf chunks maintaining O(1) random access with zero reallocations.
- [[Suffix Array and Longest Common Prefix (LCP) Array]] — Compact sorted suffix index for full-text pattern matching and substring queries.
- [[Non-Temporal Store Arrays (Direct-to-DRAM Streaming Stores)]] — Bypassing CPU cache pollution for massive write-only bulk array initializations.

### 8. 📂 [[Array Algorithms & Operations|08. Array Algorithms & Operations]]
- [[Array Algorithms & Operations]] — Master topology: Essential in-place algorithms, memory shifts, branchless binary search, prefix sums, difference arrays, two-pointer compactions, and partitioning.
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
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`
