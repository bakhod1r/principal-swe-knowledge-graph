---
title: "Diagonal and Anti-Diagonal Matrix Traversals"
tags:

  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# Diagonal and Anti-Diagonal Matrix Traversals

## 1. Definition
Mathematical classification of linear paths across 2D matrices:
- **Main Diagonal**: Elements where row index equals column index ($i = j$).
- **Anti-Diagonals (Secondary Diagonals)**: Lines perpendicular to the main diagonal where the sum of indices is invariant:
  $$i + j = k, \quad k \in [0, M + N - 2]$$
- **Parallel Diagonals**: Lines parallel to the main diagonal where the difference of indices is invariant:
  $$i - j = c, \quad c \in [-(N-1), M-1]$$

## 2. Invariant Geometry

```text
Anti-Diagonal Invariant (i + j = k):
k=0: [0,0]
k=1: [0,1], [1,0]
k=2: [0,2], [1,1], [2,0]
k=3: [1,2], [2,1]
k=4: [2,2]
```

## 3. Industrial Applications
- **Zigzag Traversal in JPEG Compression**: JPEG DCT quantization matrices are serialized into 1D bitstreams along anti-diagonals to group high-frequency zeros together for run-length encoding.
- **Dynamic Programming Space Optimization**: In edit distance or sequence alignment, states depend only on $(i-1, j-1)$, $(i-1, j)$, and $(i, j-1)$, enabling $O(N)$ wavefront evaluation along anti-diagonal wave fronts.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
