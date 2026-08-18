---
title: Hybrid Production Sorts
tags:
  - computer-science
  - algorithms
  - sorting
  - principal-swe
parent: "[[Sorting]]"
---

# Hybrid Production Sorts

TimSort, IntroSort, PDQSort (Pattern-Defeating QuickSort).

```text
Hybrid Production Sorts
│
├── [[TimSort Run Finding and Galloping Mode]]
├── [[IntroSort Depth Limiting and Fallback]]
└── [[Pattern-Defeating QuickSort (PDQSort)]]
```

---

## 🗂️ Topics

- [[TimSort Run Finding and Galloping Mode]] — Adaptive stable hybrid of Merge Sort and Insertion Sort optimizing real-world ordered data (Python/Java).
- [[IntroSort Depth Limiting and Fallback]] — Hybrid Quick Sort falling back to Heap Sort upon recursion depth limits to guarantee O(N log N) (C++ std::sort).
- [[Pattern-Defeating QuickSort (PDQSort)]] — Branchless partitioning, fallback to Heap Sort, and detecting pattern repetitions (Go 1.19+ sort).

---

## 🔗 References
- ⬆️ Parent: [[Sorting]]
- 🎓 Root: [[Principal SWE]]
