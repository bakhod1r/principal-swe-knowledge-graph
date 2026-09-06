---
title: Skip List
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Linked Lists]]"
---

# 📦 Skip List

Multi-layer probabilistic forward pointers providing O(log N) search, insertion, and deletion without tree rebalancing.

```text
Skip List
│
├── [[Skip List Probabilistic Express Towers and O(log N) Search]]
├── [[Skip List Probabilistic Geometric Level Distribution]]
├── [[Skip List Search, Insertion, and Predecessor Array Traversal]]
└── [[Lock-Free Concurrent Skip List (Java ConcurrentSkipListMap Mechanics)]]
```

---

## 🗂️ Topics & Implementations

- [[Skip List Probabilistic Express Towers and O(log N) Search]] — Multi-level express lanes with coin-toss geometric distribution.
- [[Skip List Probabilistic Geometric Level Distribution]] — Mathematical proofs of height p = 1/2 distribution and O(log N) expected bounds.
- [[Skip List Search, Insertion, and Predecessor Array Traversal]] — Tracking update arrays across tower levels during insertions and deletions.
- [[Lock-Free Concurrent Skip List (Java ConcurrentSkipListMap Mechanics)]] — Atomic CAS-based node linking enabling scalable concurrent sorted sets.

---

## 🔗 References
- ⬆️ Parent: [[Linked Lists]]
- 📚 Module: `Data Structures`
