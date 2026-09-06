---
title: Open Addressing Hash Table
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Hash Tables]]"
---

# 📦 Open Addressing Hash Table

All elements stored directly in the flat bucket array.

```text
Open Addressing Hash Table
│
├── [[Open Addressing and Flat Array Storage Invariants]]
├── [[Linear Probing Collision Resolution and Cluster Formation]]
├── [[Quadratic Probing Mechanics and Secondary Clustering]]
├── [[Double Hashing Collision Resolution]]
├── [[Tombstone Markers and Lazy Deletion in Open Addressing]]
├── [[Primary and Secondary Clustering in Open Addressing]]
└── [[Tombstone Garbage Collection and Probing Chain Compaction]]
```

---

## 🗂️ Topics & Implementations

- [[Open Addressing and Flat Array Storage Invariants]] — Zero pointer indirection: keys and values stored directly in contiguous array slots.
- [[Linear Probing Collision Resolution and Cluster Formation]] — Sequential slot scanning: h(k, i) = (h(k) + i) % N and cache line prefetch benefits.
- [[Quadratic Probing Mechanics and Secondary Clustering]] — Non-linear probing intervals h(k, i) = (h(k) + c1*i + c2*i^2) % N to mitigate clustering.
- [[Double Hashing Collision Resolution]] — Using a second hash function for probe step size: h(k, i) = (h1(k) + i * h2(k)) % N.
- [[Tombstone Markers and Lazy Deletion in Open Addressing]] — Using deleted sentinels to preserve search chains during element removal.
- [[Primary and Secondary Clustering in Open Addressing]] — Mathematical analysis of probe degradation as load factor exceeds 0.7.
- [[Tombstone Garbage Collection and Probing Chain Compaction]] — Reclaiming tombstone slots during insert probes and full table reorganizations.

---

## 🔗 References
- ⬆️ Parent: [[Hash Tables]]
- 📚 Module: `Data Structures`
