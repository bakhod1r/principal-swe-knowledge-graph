---
title: Intrusive Linked List
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Linked Lists]]"
---

# 📦 Intrusive Linked List

Nodes embedding pointer heads directly within data payloads, eliminating dynamic wrapper allocations.

```text
Intrusive Linked List
│
├── [[Intrusive Linked List Architecture (Linux Kernel list_head)]]
├── [[Linux Kernel list_head Topology and container_of Macro]]
├── [[Intrusive vs Non-Intrusive Zero-Allocation Performance Profiles]]
└── [[Embedding Multiple Intrusive Lists in a Single Struct]]
```

---

## 🗂️ Topics & Implementations

- [[Intrusive Linked List Architecture (Linux Kernel list_head)]] — Embedding pointer hooks inside payloads and zero wrapper allocations.
- [[Linux Kernel list_head Topology and container_of Macro]] — Pointer arithmetic recovering enclosing struct pointers via offsetof calculations.
- [[Intrusive vs Non-Intrusive Zero-Allocation Performance Profiles]] — Eliminating cache misses and allocator pressure in systems engines.
- [[Embedding Multiple Intrusive Lists in a Single Struct]] — Simultaneously tracking an object in priority, hash, and FIFO queues without multiple wrappers.

---

## 🔗 References
- ⬆️ Parent: [[Linked Lists]]
- 📚 Module: `Data Structures`
