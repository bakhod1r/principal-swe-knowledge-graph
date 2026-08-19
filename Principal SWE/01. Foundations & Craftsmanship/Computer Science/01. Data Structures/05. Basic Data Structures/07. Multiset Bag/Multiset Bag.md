---
title: Multiset Bag
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Multiset Bag

Collections permitting duplicate values with frequency counters and sorted order.

```text
Multiset Bag
│
├── [[Multiset Bag Invariants and Multiplicity Models]]
├── [[Multiset Add and Frequency Increment]]
├── [[Multiset Remove and Frequency Decrement]]
├── [[Multiset Query Key Frequency and Total Multiplicity]]
├── [[Multiset Distinct Keys Traversal]]
├── [[Hash-Based Multiset vs Balanced Tree Multiset]]
└── [[Multiset Sliding Window Frequency Balancing]]
```

---

## 🗂️ Operations & Topics

- [[Multiset Bag Invariants and Multiplicity Models]] — Formal bag axioms permitting duplicate elements with associated multiplicity counts.
- [[Multiset Add and Frequency Increment]] — Inserting an element and incrementing its occurrence frequency in O(1).
- [[Multiset Remove and Frequency Decrement]] — Decrementing occurrence count and deleting key upon reaching zero.
- [[Multiset Query Key Frequency and Total Multiplicity]] — O(1) querying total duplicate count for a given key and total elements in bag.
- [[Multiset Distinct Keys Traversal]] — Iterating unique key domain without visiting duplicate entries.
- [[Hash-Based Multiset vs Balanced Tree Multiset]] — Hash map frequency bag O(1) vs C++ std::multiset Red-Black tree O(log N).
- [[Multiset Sliding Window Frequency Balancing]] — Maintaining dynamic frequency maps across sliding windows in O(1) amortized time.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: `Data Structures & Algorithms`

