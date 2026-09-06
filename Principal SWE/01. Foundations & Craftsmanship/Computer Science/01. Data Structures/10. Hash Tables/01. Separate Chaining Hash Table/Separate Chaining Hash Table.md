---
title: Separate Chaining Hash Table
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Hash Tables]]"
---

# 📦 Separate Chaining Hash Table

Bucket arrays linked to separate chains of collisions.

```text
Separate Chaining Hash Table
│
├── [[Separate Chaining Architecture and Bucket Array Topology]]
├── [[Hash Table Insertion and Collision Resolution (Chaining)]]
├── [[Hash Table Deletion and Key Unlinking (Chaining)]]
├── [[Separate Chaining Treeification Threshold (Java 8 Red-Black Tree)]]
├── [[Load Factor Thresholds and Rehashing Economics in Chaining]]
└── [[Cache-Conscious Chaining (Unrolled Bucket Arrays)]]
```

---

## 🗂️ Topics & Implementations

- [[Separate Chaining Architecture and Bucket Array Topology]] — Array of bucket heads pointing to collision lists with load factor tracking.
- [[Hash Table Insertion and Collision Resolution (Chaining)]] — Hashing keys to bucket indices and prepending nodes to collision lists in O(1) time.
- [[Hash Table Deletion and Key Unlinking (Chaining)]] — Bypassing node in bucket list and freeing heap memory.
- [[Separate Chaining Treeification Threshold (Java 8 Red-Black Tree)]] — Converting degenerate O(N) linked lists into O(log N) RB-trees when bucket depth exceeds 8.
- [[Load Factor Thresholds and Rehashing Economics in Chaining]] — Balancing memory overhead and search latency by resizing at alpha = 0.75.
- [[Cache-Conscious Chaining (Unrolled Bucket Arrays)]] — Inlining small arrays in buckets to minimize pointer dereference cache misses.

---

## 🔗 References
- ⬆️ Parent: [[Hash Tables]]
- 📚 Module: `Data Structures`
