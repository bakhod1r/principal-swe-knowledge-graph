---
title: Trees and Hierarchies
tags:
  - computer-science
  - algorithms
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Trees and Hierarchies

AVL, Red-Black, Splay, Treap, B-Tree, B+ Tree, Segment Tree, Fenwick Tree, Trie, Radix, and HAMT.

```text
Trees and Hierarchies
│
├── [[AVL Tree Strict Balance and Rotations]]
├── [[Red-Black Tree Invariants and Colored Balancing]]
├── [[Splay Tree and Amortized Self-Balancing]]
├── [[Treap and Randomized Binary Search Tree]]
├── [[B-Tree and B-Plus Tree on Disk Storage]]
├── [[Segment Tree and Lazy Propagation]]
├── [[Fenwick Tree (Binary Indexed Tree) Indexing Trick]]
├── [[Trie, Radix Tree, and Compressed Prefix Tries]]
└── [[Hash Array Mapped Trie (HAMT)]]
```

---

## 🗂️ Topics

- [[AVL Tree Strict Balance and Rotations]] — Rigid balance factor (|h_L - h_R| <= 1) providing optimal read-heavy search performance.
- [[Red-Black Tree Invariants and Colored Balancing]] — Relaxed height bound (2 log n) optimizing insertion and deletion mutation throughput.
- [[Splay Tree and Amortized Self-Balancing]] — Access-driven tree rotation (splaying) providing O(log n) amortized cost for non-uniform access.
- [[Treap and Randomized Binary Search Tree]] — Combining BST keys with randomized max-heap priorities to maintain balance without rotations.
- [[B-Tree and B-Plus Tree on Disk Storage]] — High fan-out multi-way search trees optimized for block storage and database index pages.
- [[Segment Tree and Lazy Propagation]] — O(log N) range queries and range updates using deferred tree node modifications.
- [[Fenwick Tree (Binary Indexed Tree) Indexing Trick]] — Prefix sum queries and point updates using two s complement bitwise operations (i & -i).
- [[Trie, Radix Tree, and Compressed Prefix Tries]] — Prefix-based retrieval structures, edge-compacted Patricia tries, and longest prefix matching.
- [[Hash Array Mapped Trie (HAMT)]] — Persistent, immutable, high-fan-out (32-way) trie structures used in modern functional languages.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 🎓 Root: [[Principal SWE]]
