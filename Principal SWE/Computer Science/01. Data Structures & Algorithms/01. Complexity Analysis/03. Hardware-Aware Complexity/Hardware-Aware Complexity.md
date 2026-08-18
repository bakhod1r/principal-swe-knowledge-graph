---
title: Hardware-Aware Complexity
tags:
  - computer-science
  - algorithms
  - complexity
  - principal-swe
parent: "[[Complexity Analysis]]"
---

# Hardware-Aware Complexity

Cache-conscious algorithms, External Memory I/O model, memory latency hierarchies, and branch prediction penalties.

```text
Hardware-Aware Complexity
│
├── [[External Memory Model and IO Complexity]]
├── [[Cache-Oblivious Algorithm Design]]
├── [[Branch Prediction Penalty in Branchless Code]]
└── [[Memory Allocation Overhead and Fragmentation Cost]]
```

---

## 🗂️ Topics

- [[External Memory Model and IO Complexity]] — Aggarwal-Vitter model measuring disk page and cache line block transfers (B and M).
- [[Cache-Oblivious Algorithm Design]] — Optimal cache usage across all memory hierarchy levels without hardware tuning parameters.
- [[Branch Prediction Penalty in Branchless Code]] — Designing branch-free arithmetic to eliminate CPU pipeline stalls on random data.
- [[Memory Allocation Overhead and Fragmentation Cost]] — Analyzing hidden allocator overhead and cache line pollution in pointer-heavy structures.

---

## 🔗 References
- ⬆️ Parent: [[Complexity Analysis]]
- 🎓 Root: [[Principal SWE]]
