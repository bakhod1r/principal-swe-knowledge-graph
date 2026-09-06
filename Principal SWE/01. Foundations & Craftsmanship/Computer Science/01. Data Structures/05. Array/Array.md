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

# 📦 Array Data Structure & Modular Architecture

Exhaustive, hardware-conscious mastery of contiguous memory arrays across 5 core operational domains: physical pointer arithmetic, SIB register indexing, dynamic vector allocation, geometric growth policies (1.5x vs 2.0x), in-place algorithms, 2D/3D matrices, sparse representations (COO, CSR, CSC), SIMD vectorization, cache lines, bit arrays, and memory-mapped files across 37 comprehensive operational blueprints.

```text
Array
│
├── 01. Core Anatomy & Memory Architecture
│   ├── [[Array Memory Layout and Pointer Arithmetic]]
│   ├── [[Array Access and Indexing Mechanics]]
│   ├── [[Array Memory Alignment, Cache Line Striding, and False Sharing]]
│   ├── [[Array Bounds Checking Elimination (BCE) Compiler Optimization]]
│   ├── [[Array Memory Copy (SIMD rep movsq and memmove)]]
│   └── [[Structure of Arrays (SoA) vs Array of Structures (AoS)]]
│
├── 02. Dynamic Arrays & Capacity Policies
│   ├── [[Array Append and Dynamic Capacity Allocation]]
│   ├── [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]]
│   ├── [[Array Pre-Allocation and Capacity Hints]]
│   ├── [[Array Push and Pop (LIFO Dynamic Stack Operations)]]
│   ├── [[Array Memory Leaks and Garbage Collection Truncation]]
│   └── [[Array Clear and Logical vs Physical Truncation]]
│
├── 03. Operations & In-Place Algorithms
│   ├── [[Array Insert at Index (Order-Preserving Shift)]]
│   ├── [[Array Delete at Index (Order-Preserving Shift)]]
│   ├── [[Array Delete Fast (Unordered Swap-and-Pop)]]
│   ├── [[Array Search (Linear vs Branchless Binary Search)]]
│   ├── [[Array Slicing, Windows, and Full 3-Index Expressions]]
│   ├── [[Array Reverse and In-Place 3-Reversal Rotation]]
│   ├── [[Array Filter and In-Place Two-Pointer Compaction]]
│   ├── [[Array Deduplication on Sorted Buffers]]
│   └── [[Array Partitioning (Lomuto vs Hoare Partition)]]
│
├── 04. Multidimensional Arrays & Matrices
│   ├── [[Matrix & Multidimensional Arrays]]
│   ├── [[Row-Major vs Column-Major Layout and Memory Striding]]
│   ├── [[3D Array Layout, Voxel Tensors, and Coordinate Flattening]]
│   ├── [[Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)]]
│   ├── [[Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching]]
│   ├── [[In-Place Matrix Transposition (Square vs Rectangular)]]
│   ├── [[Matrix 90-Degree In-Place Rotation (Transpose and Reverse)]]
│   ├── [[Spiral Matrix Traversal and Boundary Layer Peel]]
│   ├── [[Diagonal and Anti-Diagonal Matrix Traversals]]
│   ├── [[Sparse Matrix Coordinate List (COO) Representation]]
│   ├── [[Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)]]
│   ├── [[Cache-Oblivious Matrix Multiplication and Tiling]]
│   └── [[3D Tensor Slicing, Sub-Volumes, and Dimensional Projections]]
│
└── 05. Specialized & System Array Buffers
    ├── [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]]
    ├── [[Bit Array (Packed Bitset Operations)]]
    └── [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]]
```

---

## 🗂️ Core Knowledge Domains

### 1. 📂 01. Core Anatomy & Memory Architecture
- [[Array Memory Layout and Pointer Arithmetic]] — Linear contiguous byte addressing, scale-index-base translation, and memory offsets.
- [[Array Access and Indexing Mechanics]] — O(1) random-access mechanics, SIB register decoding, and zero-overhead lookups.
- [[Array Memory Alignment, Cache Line Striding, and False Sharing]] — Hardware memory bus alignment, 64-byte L1 cache line prefetching, and false sharing.
- [[Array Bounds Checking Elimination (BCE) Compiler Optimization]] — Compiler optimization eliminating runtime bounds check branch overheads.
- [[Array Memory Copy (SIMD rep movsq and memmove)]] — Hardware vectorized block transfers, AVX-512 alignment, and overlapping buffers.
- [[Structure of Arrays (SoA) vs Array of Structures (AoS)]] — Data-oriented design, SIMD memory packing, and cache efficiency.

### 2. 📂 02. Dynamic Arrays & Capacity Policies
- [[Array Append and Dynamic Capacity Allocation]] — Amortized O(1) growth, allocation doubling, and copy overhead profiles.
- [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]] — Mathematical amortization proofs, allocator reuse, and memory fragmentation.
- [[Array Pre-Allocation and Capacity Hints]] — Zero-reallocation ingestion via size hinting and heap reserve sizing.
- [[Array Push and Pop (LIFO Dynamic Stack Operations)]] — Strict O(1) LIFO operations on dynamic arrays with top pointer tracking.
- [[Array Memory Leaks and Garbage Collection Truncation]] — Preventing memory retention by pointer zeroing and sub-slice truncation.
- [[Array Clear and Logical vs Physical Truncation]] — O(1) length reset vs O(N) memory zeroing and GC reclamation.

### 3. 📂 03. Operations & In-Place Algorithms
- [[Array Insert at Index (Order-Preserving Shift)]] — O(N) middle insertion requiring rightward memory block displacement.
- [[Array Delete at Index (Order-Preserving Shift)]] — O(N) element removal preserving sequential ordering via leftward shift.
- [[Array Delete Fast (Unordered Swap-and-Pop)]] — O(1) constant time deletion by swapping target slot with the tail element.
- [[Array Search (Linear vs Branchless Binary Search)]] — O(N) sequential scans vs O(log N) branchless binary search on sorted data.
- [[Array Slicing, Windows, and Full 3-Index Expressions]] — Zero-copy pointer windowing and 3-index capacity boundary capping.
- [[Array Reverse and In-Place 3-Reversal Rotation]] — Two-pointer reversal and O(1) space block rotation algorithm.
- [[Array Filter and In-Place Two-Pointer Compaction]] — Zero-allocation compaction using fast-read and slow-write pointer pairs.
- [[Array Deduplication on Sorted Buffers]] — O(N) single-pass deduplication with zero allocations on sorted sequences.
- [[Array Partitioning (Lomuto vs Hoare Partition)]] — In-place pivot partitioning schemes for QuickSort and QuickSelect.

### 4. 📂 04. Multidimensional Arrays & Matrices
- [[Matrix & Multidimensional Arrays]] — Master index for 2D matrices, 3D voxel tensors, strided buffers, and sparse matrices.
- [[Row-Major vs Column-Major Layout and Memory Striding]] — Flat linear addressing equations and memory striding coefficients.
- [[3D Array Layout, Voxel Tensors, and Coordinate Flattening]] — 3D indexing `D*H*W`, spatial volumetric grids, and tensor flattening.
- [[Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)]] — Single-allocation flat arrays vs array-of-pointers (`T**`), pointer overhead, and cache misses.
- [[Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching]] — CPU L1/L2 cache line prefetching, spatial locality, and stride penalties.
- [[In-Place Matrix Transposition (Square vs Rectangular)]] — Diagonal swap symmetry and cycle-following permutation algorithms in O(1) space.
- [[Matrix 90-Degree In-Place Rotation (Transpose and Reverse)]] — 90-degree rotations combining transpose and reverse.
- [[Spiral Matrix Traversal and Boundary Layer Peel]] — 4-pointer contraction boundary traversal in O(M*N) time.
- [[Diagonal and Anti-Diagonal Matrix Traversals]] — Invariants (i+j=k, i-j=k) driving zigzag and diagonal traversals.
- [[Sparse Matrix Coordinate List (COO) Representation]] — Triplet array storage (row, col, val) for hyper-sparse datasets.
- [[Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)]] — Compressed industrial formats for fast SpMV computations.
- [[Cache-Oblivious Matrix Multiplication and Tiling]] — Recursive block partitioning matching cache hierarchy boundaries.
- [[3D Tensor Slicing, Sub-Volumes, and Dimensional Projections]] — Zero-copy strided views and axial/sagittal/coronal hyperplane slicing.

### 5. 📂 05. Specialized & System Array Buffers
- [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]] — Modulo-free circular queue wrapping using power-of-two bitwise masks.
- [[Bit Array (Packed Bitset Operations)]] — Dense 64-to-1 memory compression using 64-bit word bit manipulation.
- [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]] — Operating system virtual memory page table mapping for multi-gigabyte datasets.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`


