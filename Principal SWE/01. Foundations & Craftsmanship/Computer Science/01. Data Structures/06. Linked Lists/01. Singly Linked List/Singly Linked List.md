---
title: Singly Linked List
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Linked Lists]]"
---

# 📦 Singly Linked List

Linear single-pointer sequence, dynamic heap nodes, dummy head sentinel pattern, and fundamental node insertions/deletions.

```text
Singly Linked List
│
├── [[Singly Linked List Node Memory Layout and Heap Pointers]]
├── [[Singly Linked List Dummy Sentinel Head Pattern]]
├── [[Singly Linked List Insert at Head (O(1) Prepend)]]
├── [[Singly Linked List Insert at Tail (O(1) with Tail Pointer)]]
├── [[Singly Linked List Insert at Index (Order-Preserving)]]
├── [[Singly Linked List Delete Head and Tail Elements]]
├── [[Singly Linked List Delete by Value (First Occurrence)]]
├── [[Singly Linked List Delete Node with O(1) Pointer Handle]]
├── [[Singly Linked List Memory Overhead and Cache Miss Economics]]
└── [[Forward Iterator Mechanics and Traversal Bounds]]
```

---

## 🗂️ Topics & Implementations

- [[Singly Linked List Node Memory Layout and Heap Pointers]] — Heap allocation, next pointer memory overhead, and 64-bit address alignment.
- [[Singly Linked List Dummy Sentinel Head Pattern]] — Eliminating edge-case branching for empty lists and head insertions.
- [[Singly Linked List Insert at Head (O(1) Prepend)]] — O(1) pointer redirection prepend operation without list traversal.
- [[Singly Linked List Insert at Tail (O(1) with Tail Pointer)]] — Maintaining tail pointer for instant O(1) appending vs O(N) scans.
- [[Singly Linked List Insert at Index (Order-Preserving)]] — Traversing to index i - 1 and splicing node in O(N) time.
- [[Singly Linked List Delete Head and Tail Elements]] — O(1) head removal vs O(N) tail removal requiring predecessor search.
- [[Singly Linked List Delete by Value (First Occurrence)]] — Locating target node and bypassing predecessor pointer.
- [[Singly Linked List Delete Node with O(1) Pointer Handle]] — Copying successor value into target slot to delete without predecessor pointer.
- [[Singly Linked List Memory Overhead and Cache Miss Economics]] — Analyzing pointer-to-payload ratio and pointer-chasing latency costs.
- [[Forward Iterator Mechanics and Traversal Bounds]] — Single-pass forward iteration invariants and termination checks.

---

## 🔗 References
- ⬆️ Parent: [[Linked Lists]]
- 📚 Module: `Data Structures`
