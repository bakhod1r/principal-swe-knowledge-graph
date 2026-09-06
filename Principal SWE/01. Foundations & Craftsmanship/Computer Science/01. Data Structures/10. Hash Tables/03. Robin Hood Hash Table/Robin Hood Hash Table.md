---
title: Robin Hood Hash Table
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Hash Tables]]"
---

# 📦 Robin Hood Hash Table

Stealing from the rich to give to the poor to minimize probe variance.

```text
Robin Hood Hash Table
│
├── [[Robin Hood Hashing and Probe Sequence Length (PSL)]]
├── [[Probe Sequence Length (PSL) Variance Reduction Proof]]
└── [[Backward-Shift Deletion in Robin Hood Hashing (Eliminating Tombstones)]]
```

---

## 🗂️ Topics & Implementations

- [[Robin Hood Hashing and Probe Sequence Length (PSL)]] — Tracking distance from home bucket and swapping on insertion when incoming PSL > resident PSL.
- [[Probe Sequence Length (PSL) Variance Reduction Proof]] — Mathematical proof of dramatically reduced search variance and predictable lookup times.
- [[Backward-Shift Deletion in Robin Hood Hashing (Eliminating Tombstones)]] — Shifting subsequent elements backward until PSL is 0, completely avoiding tombstone overhead.

---

## 🔗 References
- ⬆️ Parent: [[Hash Tables]]
- 📚 Module: `Data Structures`
