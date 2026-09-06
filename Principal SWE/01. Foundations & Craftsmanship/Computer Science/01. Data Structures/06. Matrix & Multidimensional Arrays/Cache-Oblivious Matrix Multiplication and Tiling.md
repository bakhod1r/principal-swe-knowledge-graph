---
title: "Cache-Oblivious Matrix Multiplication and Tiling"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[Matrix & Multidimensional Arrays]]"
---

# Cache-Oblivious Matrix Multiplication and Tiling

## 1. Definition
Naive matrix multiplication $C = A \times B$ takes $O(N^3)$ operations via 3 nested loops:
$$C[i][j] = \sum_{k=0}^{N-1} A[i][k] \cdot B[k][j]$$

In naive execution, $B[k][j]$ iterates down columns, incurring severe cache misses on every step. **Cache Blocking (Tiling)** and **Cache-Oblivious Divide-and-Conquer** algorithms partition large matrices into sub-blocks that fit completely inside the CPU L1/L2 data cache.

## 2. Block Matrix Tiling Paradigm

```text
Full Matrix A (NxN)            Tiled Blocks (Block size BxB):
+---------+---------+          +----+----+
| Block00 | Block01 |          |B00 |B01 |  <-- Loaded into L1 cache once,
+---------+---------+          +----+----+      reused for all dot products!
| Block10 | Block11 |          |B10 |B11 |
+---------+---------+          +----+----+
```

Loop Tiling Algorithm ($B \times B$ tiles):
```c
for (int i0 = 0; i0 < N; i0 += B)
  for (int j0 = 0; j0 < N; j0 += B)
    for (int k0 = 0; k0 < N; k0 += B)
      for (int i = i0; i < min(i0 + B, N); i++)
        for (int k = k0; k < min(k0 + B, N); k++)
          for (int j = j0; j < min(j0 + B, N); j++)
            C[i * N + j] += A[i * N + k] * B[k * N + j];
```

## 3. Cache Complexity & Speedup
- **Naive I/O Complexity**: $O(N^3)$ cache line transfers from DRAM.
- **Tiled I/O Complexity**: $O(N^3 / (B \cdot \sqrt{M}))$, reducing DRAM bus traffic by up to **80-90%**, achieving near-peak CPU GFLOPS.

---

## 🔗 References
- ⬆️ Parent: [[Matrix & Multidimensional Arrays]]
- 📚 Module: `Data Structures`
