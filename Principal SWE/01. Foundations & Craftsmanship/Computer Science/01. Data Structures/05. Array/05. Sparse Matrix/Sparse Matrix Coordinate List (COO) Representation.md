---
title: "Sparse Matrix Coordinate List (COO) Representation"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - sparse-matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# Sparse Matrix Coordinate List (COO) Representation

## 1. Definition
When a large matrix contains predominantly zero values (e.g., in graph adjacency matrices, recommendation embeddings, finite element models), storing dense zeros wastes massive amounts of memory ($M \times N \times 8$ bytes).

The **Coordinate List (COO)** format stores only non-zero entries as independent triplets:
$$\text{Triplet: } (row, col, value)$$

Stored internally as three parallel arrays:
- `row[]`: row coordinates of non-zero items
- `col[]`: column coordinates of non-zero items
- `val[]`: stored values

## 2. Memory Economics

```text
Dense Matrix (4x4, 3 non-zeros):
[ 0,  0, 10,  0 ]
[ 0,  0,  0,  0 ]
[ 5,  0,  0,  0 ]
[ 0, 20,  0,  0 ]

COO Format:
row: [ 0,  2,  3 ]
col: [ 2,  0,  1 ]
val: [ 10, 5, 20 ]
```

- **Dense Memory**: $M \times N \times \text{sizeof}(T) = 16 \times 4 = 64$ bytes.
- **COO Memory**: $NNZ \times (\text{sizeof}(int) + \text{sizeof}(int) + \text{sizeof}(T)) = 3 \times 12 = 36$ bytes.
- **Breakeven Point**: COO saves memory when non-zero density is below $\approx 33\%$.

## 3. Usage & Gotchas
- COO is the ideal write-format for incremental matrix construction.
- Random element lookup $A[i][j]$ requires $O(\log NNZ)$ if sorted or $O(NNZ)$ if unsorted. Before executing matrix arithmetic (matrix multiplication), COO is typically converted to CSR or CSC.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
