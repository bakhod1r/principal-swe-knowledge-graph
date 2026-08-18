---
title: Hash-Based Structures
tags:
  - computer-science
  - algorithms
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Hash-Based Structures

Hash tables, collision resolution, Robin Hood hashing, Cuckoo hashing, consistent hashing, and probabilistic filters.

```text
Hash-Based Structures
│
├── [[Hash Table Chaining vs Open Addressing]]
├── [[Robin Hood Hashing and Low Variance Lookup]]
├── [[Cuckoo Hashing and Constant Worst-Case Lookup]]
├── [[Consistent Hashing and Virtual Nodes]]
├── [[Bloom Filter and Counting Bloom Filter]]
└── [[Cuckoo Filter and Quotient Filter]]
```

---

## 🗂️ Topics

- [[Hash Table Chaining vs Open Addressing]] — Separate chaining bucket arrays vs linear/quadratic probing and cache locality.
- [[Robin Hood Hashing and Low Variance Lookup]] — Stealing from the rich to equalize probe sequence lengths and bound P99 search latency.
- [[Cuckoo Hashing and Constant Worst-Case Lookup]] — Multiple hash functions and displacement chains ensuring true O(1) worst-case lookups.
- [[Consistent Hashing and Virtual Nodes]] — Ring-based key distribution for distributed caching, dynamodb-style partitioning, and zero-churn additions.
- [[Bloom Filter and Counting Bloom Filter]] — Probabilistic set membership testing with zero false negatives and tunable false positive bitsets.
- [[Cuckoo Filter and Quotient Filter]] — Dynamic deletion-supporting probabilistic filters with superior space efficiency over Bloom filters.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 🎓 Root: [[Principal SWE]]
