---
title: Prefix Sums Difference Arrays
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Prefix Sums Difference Arrays

Cumulative sum pre-computation and O(1) range updates.

```text
Prefix Sums Difference Arrays
│
├── [[1D Prefix Sums Range Query (O(1) Interval Sums)]]
├── [[2D Matrix Prefix Sums (Inclusion-Exclusion Rectangular Queries)]]
├── [[Difference Array Range Update (O(1) Batch Range Addition)]]
├── [[2D Difference Array (Grid Boundary Updates)]]
├── [[Subarray Sum Equals K (Prefix Sum Frequency Map)]]
├── [[Continuous Subarray Sum (Modulo Prefix Tracking)]]
└── [[Range XOR Prefix Queries (Bitwise Reversibility)]]
```

---

## 🗂️ Operations & Topics

- [[1D Prefix Sums Range Query (O(1) Interval Sums)]] — Pre-computing cumulative sums P[i] = P[i-1] + A[i] for O(1) range queries (P[R] - P[L-1]).
- [[2D Matrix Prefix Sums (Inclusion-Exclusion Rectangular Queries)]] — Inclusion-Exclusion rectangular sum queries in O(1) time after O(N*M) preprocessing.
- [[Difference Array Range Update (O(1) Batch Range Addition)]] — Applying range addition [L, R] += V in O(1) by updating Diff[L] += V and Diff[R+1] -= V.
- [[2D Difference Array (Grid Boundary Updates)]] — Applying 2D rectangular grid additions in O(1) time using 4 corner updates.
- [[Subarray Sum Equals K (Prefix Sum Frequency Map)]] — Counting subarrays summing to K in O(N) time using prefix sum frequency hash map.
- [[Continuous Subarray Sum (Modulo Prefix Tracking)]] — Tracking running prefix sum modulo K in hash map to find multiple-of-K subarrays in O(N).
- [[Range XOR Prefix Queries (Bitwise Reversibility)]] — Exploiting XOR self-inverse property (X ^ X = 0) for O(1) range XOR queries.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: `Data Structures & Algorithms`

