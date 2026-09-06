---
title: "Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - sparse-matrix
  - principal-swe
parent: "[[Matrix & Multidimensional Arrays]]"
---

# Compressed Sparse Row (CSR) and Compressed Sparse Column (CSC)

## 1. Definition
**Compressed Sparse Row (CSR)** compresses row indices from COO into a compact pointer array of offsets, eliminating redundant row coordinate storage. It is the de facto industrial standard for high-performance scientific computing and Sparse Matrix-Vector multiplication (SpMV).

CSR uses 3 arrays:
1. `val[]`: Array of non-zero elements of length $NNZ$.
2. `col_idx[]`: Column index of each non-zero element of length $NNZ$.
3. `row_ptr[]`: Array of size $M + 1$. `row_ptr[i]` points to the starting index in `val[]` for row $i$. Row $i$ spans indices `[row_ptr[i], row_ptr[i+1])`.

**Compressed Sparse Column (CSC)** is the dual format compressing column pointers, optimal for column slicing and SpMV in Fortran/Julia environments.

## 2. Structural Memory Model

```text
Matrix (4x4 with 4 non-zeros):
Row 0: [ 5, 0, 0, 0 ]
Row 1: [ 0, 8, 0, 0 ]
Row 2: [ 0, 3, 9, 0 ]
Row 3: [ 0, 0, 0, 0 ]

val:     [ 5, 8, 3, 9 ]
col_idx: [ 0, 1, 1, 2 ]
row_ptr: [ 0, 1, 2, 4, 4 ]
           ^  ^  ^  ^  ^
         Row0 R1 R2 R3 End
```

## 3. SpMV Algorithm (Sparse Matrix-Vector Multiplication)
$$y = A \cdot x$$
```c
for (int i = 0; i < M; i++) {
    double sum = 0.0;
    int start = row_ptr[i];
    int end = row_ptr[i + 1];
    for (int k = start; k < end; k++) {
        sum += val[k] * x[col_idx[k]];
    }
    y[i] = sum;
}
```
- **Complexity**: $O(NNZ)$ time. Completely skips zero elements without evaluating them.

## 4. Gotchas
- Dynamic insertion into CSR is extremely expensive ($O(NNZ)$ memory shifts). Build sparse matrices using COO or hash maps, then finalize into CSR.

---

## 🔗 References
- ⬆️ Parent: [[Matrix & Multidimensional Arrays]]
- 📚 Module: `Data Structures`
