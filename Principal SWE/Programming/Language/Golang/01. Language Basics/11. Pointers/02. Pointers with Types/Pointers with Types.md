---
title: Pointers with Types
tags:
  - golang
  - pointers
  - principal-swe
parent: "[[Pointers]]"
---

# Pointers with Types

Pointers with structs, slices, maps, channels, and reference types.

```text
Pointers with Types
│
├── [[Pointers with Structs]]
├── [[Pointers with Slices & Maps]]
├── [[Pointer to Slice vs Slice of Pointers]]
└── [[Pointers to Interfaces Trap]]
```

---

## 🗂️ Topics

- [[Pointers with Structs]] — Passing large structs by pointer to avoid memory copy overhead.
- [[Pointers with Slices & Maps]] — Why slices and maps already contain internal pointers (reference type mechanics).
- [[Pointer to Slice vs Slice of Pointers]] — Distinguishing *[]T from []*T and memory locality tradeoffs.
- [[Pointers to Interfaces Trap]] — Why pointers to interfaces (*interface{}) are almost always a design bug.

---

## 🔗 References
- ⬆️ Parent: [[Pointers]]

