---
title: "Spiral Matrix Traversal and Boundary Layer Peel"
tags:

  - computer-science
  - data-structures
  - matrix
  - principal-swe
parent: "[[2D & Multidimensional Array]]"
---

# Spiral Matrix Traversal and Boundary Layer Peel

## 1. Definition
Traversing an $M \times N$ rectangular matrix in clockwise inward spiral order starting from coordinate $(0,0)$ to the center.

## 2. Inward Boundary Peeling Model
Maintain 4 boundary pointers: `top`, `bottom`, `left`, `right`.

```text
       left                  right
top    --> --> --> --> --> --> v
       ^   Top Row Traversal   v
       ^                       v  Right Col
       ^                       v  Traversal
       ^   Bottom Row          v
bottom <-- <-- <-- <-- <-- <-- v
```

1. Traverse from `left` to `right` along `top`, then increment `top++`.
2. Traverse from `top` to `bottom` along `right`, then decrement `right--`.
3. If `top <= bottom`: traverse from `right` to `left` along `bottom`, then decrement `bottom--`.
4. If `left <= right`: traverse from `bottom` to `top` along `left`, then increment `left++`.
5. Repeat until `top > bottom` or `left > right`.

## 3. Gotchas & Edge Cases
- Single Row ($1 \times N$) or Single Column ($M \times 1$): Without the conditional checks `top <= bottom` and `left <= right` before steps 3 and 4, the algorithm duplicates the single row/column backwards.
- Time Complexity: Exactly $O(M \times N)$, visiting each element exactly once. Space: $O(1)$ auxiliary space.

---

## 🔗 References
- ⬆️ Parent: [[2D & Multidimensional Array]]
- 📚 Module: `Data Structures`
