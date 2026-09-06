---
title: "Array"
tags:
  - review
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Array

## 1. Definition

## 2. Mental Model

## 3. Usage

## 4. Gotchas

---

## 🗺️ Module Architecture & Sub-Domains

```
Array/
├── Static Array/
├── Dynamic Array/
├── Circular Buffer Array/
├── 2D & Multidimensional Array/
├── Sparse Matrix/
├── Bit Array/
├── Specialized System Arrays/
└── Common Array Algorithms & Operations/
```

---

## 📑 Comprehensive Sub-Domain Index

### [[Static Array]]
  - [[Array Access and Indexing Mechanics]]
  - [[Array Address Space Layout Randomization (ASLR) and Stack Protection]]
  - [[Array Bounds Checking Elimination (BCE) Compiler Optimization]]
  - [[Array Data-Oriented Design (DOD) and Cache Line Utilization]]
  - [[Array Memory Alignment, Cache Line Striding, and False Sharing]]
  - [[Array Memory Copy (SIMD rep movsq and memmove)]]
  - [[Array Memory Layout and Pointer Arithmetic]]
  - [[Static Array Append and Fixed-Capacity Overflow Invariant]]
  - [[Static Array Branchless Binary Search (cmov and BCE Optimization)]]
  - [[Static Array Delete Fast (Unordered Swap-and-Pop in Fixed Buffer)]]
  - [[Static Array Delete at Index (Fixed Buffer Left Shift)]]
  - [[Static Array Insert at Index (Fixed Buffer Right Shift)]]
  - [[Static Array SIMD Vectorized Linear Scan (AVX-512 Equality Matching)]]
  - [[Structure of Arrays (SoA) vs Array of Structures (AoS)]]
  - [[Variable-Length Arrays (VLA) vs Fixed Stack Allocation]]

### [[Dynamic Array]]
  - [[Array Append and Dynamic Capacity Allocation]]
  - [[Array Clear and Logical vs Physical Truncation]]
  - [[Array Geometric Growth Policies (1.5x vs 2.0x Amortization)]]
  - [[Array Memory Leaks and Garbage Collection Truncation]]
  - [[Array Pre-Allocation and Capacity Hints]]
  - [[Array Push and Pop (LIFO Dynamic Stack Operations)]]
  - [[Array Shrink-to-Fit and Memory Compaction]]
  - [[Dynamic Array Allocator Arena Integration (jemalloc, TCMalloc)]]
  - [[Dynamic Array Append and Reallocation Mechanics]]
  - [[Dynamic Array Delete Fast (Unordered Swap-and-Pop + GC Element Nulling)]]
  - [[Dynamic Array Delete at Index (Order-Preserving Shift + GC Loitering Zeroing)]]
  - [[Dynamic Array In-Place Partitioning (Reallocation-Safe Lomuto & Hoare)]]
  - [[Dynamic Array Insert at Index (Reallocation + Memory Shift)]]
  - [[Dynamic Array Slicing, Windows, and 3-Index Capacity Boundaries]]

### [[Circular Buffer Array]]
  - [[Circular Buffer Dequeue and Head Advancement Operations]]
  - [[Circular Buffer Dynamic Geometric Resizing and Linear Unwrapping]]
  - [[Circular Buffer Enqueue and Tail Wraparound Operations]]
  - [[Circular Buffer Overwrite and Head-Tail Pointer Invariants]]
  - [[Circular Buffer Peek and Contiguous Two-Slice Memory Views]]
  - [[Circular Ring Buffer Array Indexing (Bitwise AND Masking)]]
  - [[Disruptor Ring Buffer Pattern (Cache Line Padded Sequence Tracking)]]
  - [[Single-Producer Single-Consumer (SPSC) Lock-Free Array Ring Buffer]]

### [[2D & Multidimensional Array]]
  - [[3D Array Layout, Voxel Tensors, and Coordinate Flattening]]
  - [[3D Tensor Slicing, Sub-Volumes, and Dimensional Projections]]
  - [[Blocked Z-Morton Order (Morton Space-Filling Curve) Matrix Layout]]
  - [[Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching]]
  - [[Cache-Oblivious Matrix Multiplication and Tiling]]
  - [[Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)]]
  - [[Diagonal and Anti-Diagonal Matrix Traversals]]
  - [[In-Place Matrix Transposition (Square vs Rectangular)]]
  - [[Matrix 90-Degree In-Place Rotation (Transpose and Reverse)]]
  - [[Matrix Sub-Grid Slicing and Window Striding Operations]]
  - [[Row-Major vs Column-Major Layout and Memory Striding]]
  - [[Spiral Matrix Traversal and Boundary Layer Peel]]

### [[Sparse Matrix]]
  - [[Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)]]
  - [[Diagonal Format (DIA) and Banded Matrix Storage]]
  - [[Sparse Matrix Coordinate List (COO) Representation]]
  - [[Sparse Matrix Insertion, Value Mutation, and Transposition]]
  - [[Sparse Matrix-Vector Multiplication (SpMV) Cache Optimization]]

### [[Bit Array]]
  - [[Bit Array (Packed Bitset Operations)]]
  - [[Bit Array Bit-Range Slicing, Shift, and Clear Operations]]
  - [[Bitset Set Algebra (SIMD 64-Bit Word Operations)]]
  - [[Bitwise Word-Level Parallelism and Population Count (POPCNT)]]
  - [[Roaring Bitmaps and Run-Length Encoded (RLE) Hybrid Bitsets]]

### [[Specialized System Arrays]]
  - [[Gap Buffer (Text Editor Cursor Buffer)]]
  - [[Gap Buffer Text Insertion, Deletion, and Cursor Movement Operations]]
  - [[Hashed Array Tree (HAT) Cache-Friendly Flat Allocation]]
  - [[Memory-Mapped Arrays (mmap Zero-Copy File IO)]]
  - [[Non-Temporal Store Arrays (Direct-to-DRAM Streaming Stores)]]
  - [[Suffix Array and Longest Common Prefix (LCP) Array]]

### [[Common Array Algorithms & Operations]]
  - [[Array Deduplication on Sorted Buffers]]
  - [[Array Filter and In-Place Two-Pointer Compaction]]
  - [[Array Partitioning (Lomuto vs Hoare Partition)]]
  - [[Array Reverse and In-Place 3-Reversal Rotation]]
  - [[Array Search (Linear Scan vs Branchless Binary Search)]]
  - [[Boyer-Moore Majority Vote In-Place Array Scan]]
  - [[Difference Array and Range Update Range Query (RURQ)]]
  - [[Dutch National Flag 3-Way Array Partitioning]]
  - [[Kadane's Algorithm for Maximum Subarray Sum]]
  - [[Prefix Sum Array and Range Sum Queries]]
  - [[Two-Pointer Sliding Window on Contiguous Buffers]]

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Curriculum: `Computer Science`
