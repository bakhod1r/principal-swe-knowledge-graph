---
title: Linked List Stack
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Stacks]]"
---

# 📦 Linked List Stack

Node-based stack with head pointer as top.

```text
Linked List Stack
│
├── [[Stack Linked List Implementation]]
├── [[Linked List Stack Memory Fragmentation and Free-List Recycling]]
└── [[Lock-Free Concurrent Stack (Treiber Stack with Atomic CAS)]]
```

---

## 🗂️ Topics & Implementations

- [[Stack Linked List Implementation]] — O(1) push and pop at list head avoiding array reallocation pauses.
- [[Linked List Stack Memory Fragmentation and Free-List Recycling]] — Mitigating heap allocation overhead via thread-local node freelists.
- [[Lock-Free Concurrent Stack (Treiber Stack with Atomic CAS)]] — Atomic CAS pointer swing on top pointer enabling wait-free concurrent pushes.

---

## 🔗 References
- ⬆️ Parent: [[Stacks]]
- 📚 Module: `Data Structures`
