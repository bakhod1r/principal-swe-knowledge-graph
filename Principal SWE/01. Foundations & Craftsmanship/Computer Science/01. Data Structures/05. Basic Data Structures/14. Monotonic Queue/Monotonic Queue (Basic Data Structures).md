---
title: Monotonic Queue (Basic Data Structures)
tags:
  - review
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Monotonic Queue (Basic Data Structures)

Double-ended queues maintaining candidate sliding window extrema in amortized O(1).

```text
Monotonic Queue (Basic Data Structures)
│
├── [[Monotonic Queue Invariants and Candidate Extrema Maintenance]]
├── [[Sliding Window Maximum and Minimum in O(N)]]
├── [[Shortest Subarray with Sum at Least K (Monotonic Prefix Queue)]]
├── [[Constrained Subsequence Sum (Sliding Window DP with Monotonic Queue)]]
└── [[Jump Game VI DP Optimization via Monotonic Queue]]
```

---

## 🗂️ Operations & Topics

- [[Monotonic Queue Invariants and Candidate Extrema Maintenance]] — Double-ended queue maintaining candidate extrema in sorted order in amortized O(1) per slide.
- [[Sliding Window Maximum and Minimum in O(N)]] — Computing maximum/minimum of every sliding window of size K in O(N) total time.
- [[Shortest Subarray with Sum at Least K (Monotonic Prefix Queue)]] — Monotonic queue over prefix sum array finding shortest positive subarray in O(N).
- [[Constrained Subsequence Sum (Sliding Window DP with Monotonic Queue)]] — Optimizing DP transition DP[i] = A[i] + max(0, max(DP[i-k..i-1])) in O(N) time.
- [[Jump Game VI DP Optimization via Monotonic Queue]] — Reaching end of array with maximum score in O(N) time using monotonic max-queue.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: `Data Structures & Algorithms`

