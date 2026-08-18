---
title: DP Optimizations
tags:
  - computer-science
  - algorithms
  - dynamic-programming
  - principal-swe
parent: "[[Dynamic Programming]]"
---

# DP Optimizations

Convex Hull Trick, Divide and Conquer Optimization, Knuth Optimization, Monotonic Queue Optimization.

```text
DP Optimizations
│
├── [[Convex Hull Trick and Li Chao Tree]]
├── [[Divide and Conquer DP Optimization]]
├── [[Knuth Optimization for Quadrangle Inequality]]
└── [[Monotonic Queue DP Optimization]]
```

---

## 🗂️ Topics

- [[Convex Hull Trick and Li Chao Tree]] — Optimizing DP transitions of the form DP[i] = min(DP[j] + m_j * x_i + c_j) from O(N^2) to O(N log N).
- [[Divide and Conquer DP Optimization]] — Accelerating 2D DP when the optimal transition index opt[i][j] is monotonic in O(K N log N).
- [[Knuth Optimization for Quadrangle Inequality]] — Reducing interval DP from O(N^3) to O(N^2) when cost functions satisfy the quadrangle inequality.
- [[Monotonic Queue DP Optimization]] — Sliding window maximum optimization reducing transitions over sliding index bounds to O(N).

---

## 🔗 References
- ⬆️ Parent: [[Dynamic Programming]]
- 🎓 Root: [[Principal SWE]]
