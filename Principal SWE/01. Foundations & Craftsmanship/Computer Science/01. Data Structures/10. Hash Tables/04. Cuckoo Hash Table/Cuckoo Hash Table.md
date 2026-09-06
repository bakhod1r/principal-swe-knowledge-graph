---
title: Cuckoo Hash Table
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Hash Tables]]"
---

# 📦 Cuckoo Hash Table

Multiple hash locations guaranteeing strict O(1) worst-case lookups.

```text
Cuckoo Hash Table
│
├── [[Cuckoo Hashing and Dual-Hash O(1) Worst-Case Lookups]]
├── [[Cuckoo Hashing Two-Way Choice and Power of Two Choices]]
└── [[Cuckoo Graph Cycles, Stash Mechanics, and Rehash Triggers]]
```

---

## 🗂️ Topics & Implementations

- [[Cuckoo Hashing and Dual-Hash O(1) Worst-Case Lookups]] — Two independent hash functions; lookup inspects at most two deterministic slots.
- [[Cuckoo Hashing Two-Way Choice and Power of Two Choices]] — Inserting into alternate location, kicking existing occupant if occupied.
- [[Cuckoo Graph Cycles, Stash Mechanics, and Rehash Triggers]] — Detecting eviction cycles using cuckoo graphs and handling overflows with small stashes.

---

## 🔗 References
- ⬆️ Parent: [[Hash Tables]]
- 📚 Module: `Data Structures`
