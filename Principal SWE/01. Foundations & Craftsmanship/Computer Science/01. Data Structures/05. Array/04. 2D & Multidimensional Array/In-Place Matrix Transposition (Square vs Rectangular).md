---
title: "In-Place Matrix Transposition (Square vs Rectangular)"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# In-Place Matrix Transposition (Square vs Rectangular)

## 1. Definition
Matrix transposition reflects elements across their main diagonal:
$$B[j][i] = A[i][j]$$

- **Square Matrix ($N \times N$)**: Trivial in-place swap using $O(1)$ space by swapping elements across the main diagonal:
  $$\text{swap}(A[i][j], A[j][i]) \quad \forall i < j$$
- **Rectangular Matrix ($M \times N, M \ne N$)**: Elements must be permuted within a single flat buffer of size $M \times N$. This requires **cycle-following permutation algorithms** or bit-marking to avoid $O(M \times N)$ auxiliary memory.

## 2. In-Place Square Transpose

```text
Square Transpose (3x3):
[ 1, 2, 3 ]         [ 1, 4, 7 ]
[ 4, 5, 6 ]  ====>  [ 2, 5, 8 ]
[ 7, 8, 9 ]         [ 3, 6, 9 ]

Swap pairs: (0,1)<->(1,0), (0,2)<->(2,0), (1,2)<->(2,1).
Diagonal elements (0,0), (1,1), (2,2) remain untouched.
```

## 3. Rectangular In-Place Cycle Permutation
In a flat buffer of length $M \times N$, an element at linear index $k$ maps to target index:
$$k' = (k \times M) \pmod{M \times N - 1} \quad (0 < k < M \times N - 1)$$
Cycles are tracked using bit vectors or cycle-leader verification to ensure every element is moved exactly once in $O(M \times N)$ time and $O(1)$ extra space.

## 4. Complexity Analysis
- **Square In-Place**: Time $O(N^2)$, Space $O(1)$.
- **Rectangular Out-of-Place**: Time $O(M \times N)$, Space $O(M \times N)$.
- **Rectangular In-Place**: Time $O(M \times N \log(M \times N))$, Space $O(1)$ auxiliary memory.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
