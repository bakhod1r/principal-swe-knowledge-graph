---
title: Matrix & Multidimensional Arrays
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Matrix & Multidimensional Arrays

Contiguous flat buffers vs jagged arrays, row-major vs column-major striding, 3D voxel tensors, spatial cache locality, and matrix transformations.

```text
Matrix & Multidimensional Arrays
│
├── [[Row-Major vs Column-Major Layout and Memory Striding]]
├── [[3D Array Layout, Voxel Tensors, and Coordinate Flattening]]
├── [[Contiguous Flat Buffers vs Jagged Arrays (Array of Pointers)]]
├── [[Cache Locality, Row-Wise vs Column-Wise Traversal Prefetching]]
├── [[In-Place Matrix Transposition (Square vs Rectangular)]]
├── [[Matrix 90-Degree In-Place Rotation (Transpose and Reverse)]]
├── [[Spiral Matrix Traversal and Boundary Layer Peel]]
├── [[Diagonal and Anti-Diagonal Matrix Traversals]]
├── [[Cache-Oblivious Matrix Multiplication and Tiling]]
├── [[3D Tensor Slicing, Sub-Volumes, and Dimensional Projections]]
└── [[Blocked Z-Morton Order (Morton Space-Filling Curve) Matrix Layout]]
```

---

## 🗂️ Topics & Implementations

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

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
