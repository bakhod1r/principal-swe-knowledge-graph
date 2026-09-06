---
title: Cache Eviction Structures
tags:
  - computer-science
  - data-structures
  - caching
  - principal-swe
parent: "[[Data Structures]]"
---

# ⚡ Cache Eviction Structures

Deterministic, frequency-aware, and adaptive caching replacement policies optimizing hit ratios and memory budgets.

```text
Cache Eviction Structures
│
├── [[LRU Cache|01. LRU Cache]]
├── [[LFU Cache|02. LFU Cache]]
├── [[Mfu Cache|03. Mfu Cache]]
└── [[Arc 2q Cache|04. Arc 2q Cache]]
```

---

## 🗂️ Topics

- 📂 [[LRU Cache|01. LRU Cache]] — Least Recently Used cache combining hash map indexing with doubly linked list recency tracking in O(1).
- 📂 [[LFU Cache|02. LFU Cache]] — Least Frequently Used cache tracking access frequencies via frequency buckets and doubly linked lists in O(1).
- 📂 [[Mfu Cache|03. Mfu Cache]] — Most Frequently Used eviction model prioritizing recently introduced transient items over saturated historical keys.
- 📂 [[Arc 2q Cache|04. Arc 2q Cache]] — Adaptive Replacement Cache (ARC) and Two-Queue (2Q) policies self-tuning recency vs frequency in real time.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`
