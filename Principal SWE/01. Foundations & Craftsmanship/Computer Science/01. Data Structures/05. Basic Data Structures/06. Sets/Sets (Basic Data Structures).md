---
title: Sets (Basic Data Structures)
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Sets (Basic Data Structures)

Unique element collections, mathematical set algebra, bitset operations, and membership testing.

```text
Sets (Basic Data Structures)
│
├── [[Set Invariants and Mathematical Deduplication]]
├── [[Set Add and Insert Elements]]
├── [[Set Remove and Delete Elements]]
├── [[Set Contains and Fast Membership Testing]]
├── [[Set Union Operation (A OR B)]]
├── [[Set Intersection Operation (A AND B)]]
├── [[Set Difference and Symmetric Difference (A XOR B)]]
├── [[Hash Set vs Tree Set (O(1) Unordered vs O(log N) Sorted)]]
├── [[Bitset Set Algebra (SIMD-Accelerated 64-Bit Words)]]
└── [[Disjoint Set Union (DSU) Set Partitioning]]
```

---

## 🗂️ Operations & Topics

- [[Set Invariants and Mathematical Deduplication]] — Set uniqueness axioms and underlying associative table hash key mapping.
- [[Set Add and Insert Elements]] — O(1) expected insertion enforcing unique element constraints.
- [[Set Remove and Delete Elements]] — O(1) expected element deletion from set storage.
- [[Set Contains and Fast Membership Testing]] — O(1) strict membership verification with zero false positives.
- [[Set Union Operation (A OR B)]] — Combining distinct elements from two sets in O(N+M) time.
- [[Set Intersection Operation (A AND B)]] — Extracting common elements across sets in O(min(N, M)) time.
- [[Set Difference and Symmetric Difference (A XOR B)]] — Extracting unique elements belonging to exactly one set.
- [[Hash Set vs Tree Set (O(1) Unordered vs O(log N) Sorted)]] — Hash table O(1) unordered set vs Red-Black tree O(log N) sorted set with range queries.
- [[Bitset Set Algebra (SIMD-Accelerated 64-Bit Words)]] — Packing 64 boolean flags into uint64 words for ultra-fast bitwise AND/OR/XOR algebra.
- [[Disjoint Set Union (DSU) Set Partitioning]] — Maintaining dynamic equivalence classes in near-constant O(alpha(N)) time.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: `Data Structures & Algorithms`

