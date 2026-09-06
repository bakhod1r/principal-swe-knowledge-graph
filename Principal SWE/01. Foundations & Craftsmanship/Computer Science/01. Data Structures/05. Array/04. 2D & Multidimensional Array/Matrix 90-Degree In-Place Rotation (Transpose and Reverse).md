---
title: "Matrix 90-Degree In-Place Rotation (Transpose and Reverse)"
tags:
  - review
  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# Matrix 90-Degree In-Place Rotation (Transpose and Reverse)

## 1. Definition
Rotating an $N \times N$ square matrix by $90^\circ$ clockwise or counter-clockwise without allocating auxiliary matrix buffers ($O(1)$ space).

Algebraic Decomposition:
- **$90^\circ$ Clockwise Rotation**: $\text{Transpose}(A) \implies \text{Reverse each Row}(A)$
- **$90^\circ$ Counter-Clockwise Rotation**: $\text{Transpose}(A) \implies \text{Reverse each Column}(A)$
- **$180^\circ$ Rotation**: $\text{Reverse Rows} \implies \text{Reverse Columns}$

## 2. Visual Step-by-Step Breakdown

```text
Original:           1. Transpose:          2. Reverse Rows (Clockwise):
[ 1, 2, 3 ]         [ 1, 4, 7 ]            [ 7, 4, 1 ]
[ 4, 5, 6 ]  ====>  [ 2, 5, 8 ]    ====>   [ 8, 5, 2 ]
[ 7, 8, 9 ]         [ 3, 6, 9 ]            [ 9, 6, 3 ]
```

## 3. Implementation Blueprint
```c
void rotate_90_clockwise(int* A, int N) {
    // Step 1: Transpose in-place
    for (int i = 0; i < N; i++) {
        for (int j = i + 1; j < N; j++) {
            int tmp = A[i * N + j];
            A[i * N + j] = A[j * N + i];
            A[j * N + i] = tmp;
        }
    }
    // Step 2: Reverse each row
    for (int i = 0; i < N; i++) {
        int left = 0, right = N - 1;
        while (left < right) {
            int tmp = A[i * N + left];
            A[i * N + left] = A[i * N + right];
            A[i * N + right] = tmp;
            left++;
            right--;
        }
    }
}
```

## 4. Complexity & Gotchas
- **Time Complexity**: $O(N^2)$ — $\approx N^2/2$ swaps for transpose $+ N^2/2$ swaps for row reversal.
- **Space Complexity**: $O(1)$ strict auxiliary space.
- **Gotcha**: Attempting 4-way cyclic shifts directly on concentric rings is error-prone due to boundary index off-by-one errors. Transpose + Reverse is clean, branchless, and cache-friendly.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
