---
title: Deque (Basic Data Structures)
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Deque (Basic Data Structures)

Double-Ended Queue supporting O(1) operations at both ends with chunked arrays.

```text
Deque (Basic Data Structures)
│
├── [[Deque Invariants and Double-Ended Operational Model]]
├── [[Deque Push Front and Push Back]]
├── [[Deque Pop Front and Pop Back]]
├── [[Deque Peek Front and Peek Back]]
├── [[Deque Array-of-Chunks Implementation (std::deque Map of Blocks)]]
├── [[Deque Circular Array Ring Implementation]]
├── [[Deque Dynamic Resizing and Block Allocation]]
└── [[Work-Stealing Deque (Chase-Lev Work Stealing)]]
```

---

## 🗂️ Operations & Topics

- [[Deque Invariants and Double-Ended Operational Model]] — Double-ended queue axioms supporting push/pop at both front and back in O(1).
- [[Deque Push Front and Push Back]] — O(1) insertion at head and tail boundaries.
- [[Deque Pop Front and Pop Back]] — O(1) element removal from head and tail.
- [[Deque Peek Front and Peek Back]] — O(1) inspecting boundary elements without state mutation.
- [[Deque Array-of-Chunks Implementation (std::deque Map of Blocks)]] — Central map of fixed-size 512-byte buffer blocks avoiding giant reallocations.
- [[Deque Circular Array Ring Implementation]] — Fixed-size power-of-two circular buffer deque with head and tail wrapping.
- [[Deque Dynamic Resizing and Block Allocation]] — Allocating new 512B blocks at map boundaries without copying existing elements.
- [[Work-Stealing Deque (Chase-Lev Work Stealing)]] — Lock-free work-stealing deque for multi-core task scheduling runtime engines.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: `Data Structures & Algorithms`

