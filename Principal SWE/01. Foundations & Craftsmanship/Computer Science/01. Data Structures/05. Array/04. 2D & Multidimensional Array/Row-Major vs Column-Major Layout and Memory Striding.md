---
title: "Row-Major vs Column-Major Layout and Memory Striding"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# Row-Major vs Column-Major Layout and Memory Striding

## 1. Definition
Physical DRAM and CPU cache hierarchies are inherently 1-dimensional, linearly addressable arrays of bytes. Multi-dimensional matrices $A \in \mathbb{R}^{M \times N}$ must map two-dimensional coordinate pairs $(i, j)$ into a single scalar linear offset.

- **Row-Major Order**: Consecutive elements of a row are placed next to each other in memory. Used by C, C++, Python (NumPy default), Rust, and Java primitives.
- **Column-Major Order**: Consecutive elements of a column are placed next to each other in memory. Used by Fortran, MATLAB, Julia, R, and BLAS/LAPACK.

$$\text{Row-Major Offset: } \text{index}(i, j) = i \times N + j$$
$$\text{Column-Major Offset: } \text{index}(i, j) = j \times M + i$$

## 2. Mental Model & Memory Layout

```text
Matrix (2 rows, 3 cols):
  Row 0: [ A, B, C ]
  Row 1: [ D, E, F ]

Row-Major Physical Layout (C-style):
[ A ][ B ][ C ][ D ][ E ][ F ]
  0    1    2    3    4    5
  ----Row 0----  ----Row 1----

Column-Major Physical Layout (Fortran-style):
[ A ][ D ][ B ][ E ][ C ][ F ]
  0    1    2    3    4    5
  --Col 0--  --Col 1--  --Col 2--
```

## 3. Usage & Striding Mechanics
Strides specify the number of bytes (or element steps) needed to advance one unit in each dimension:
- In row-major with element size $S$: $\text{strides} = (N \times S, S)$.
- In column-major with element size $S$: $\text{strides} = (S, M \times S)$.

## 4. Gotchas & Cache Thrashing
- Iterating columns in the outer loop and rows in the inner loop on a row-major matrix causes stride-$N$ memory access. If $N \times \text{element\_size} \ge 64$ bytes (cache line width), **every single memory access causes an L1 cache miss**.
- BLAS/LAPACK interoperability: Passing C row-major pointers directly to BLAS without transposing or specifying `CblasRowMajor` causes silent mathematical corruption.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
