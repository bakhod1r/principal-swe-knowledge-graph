---
title: Sparse Matrix
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Sparse Matrix

Memory-compressed sparse array representations, coordinate triplets (COO), compressed row/column storage (CSR/CSC), and banded matrices.

```text
Sparse Matrix
│
├── [[Sparse Matrix Coordinate List (COO) Representation]]
├── [[Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)]]
├── [[Diagonal Format (DIA) and Banded Matrix Storage]]
└── [[Sparse Matrix-Vector Multiplication (SpMV) Cache Optimization]]
```

---

## 🗂️ Topics & Implementations

- [[Sparse Matrix Coordinate List (COO) Representation]] — Triplet array storage (row, col, value) for incremental sparse construction.
- [[Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)]] — Industrial compressed representations enabling high-performance SpMV operations.
- [[Diagonal Format (DIA) and Banded Matrix Storage]] — Compact storage for banded diagonal systems and finite-difference stencils.
- [[Sparse Matrix-Vector Multiplication (SpMV) Cache Optimization]] — Optimizing indirect memory access and vectorization for sparse matrix computations.

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
