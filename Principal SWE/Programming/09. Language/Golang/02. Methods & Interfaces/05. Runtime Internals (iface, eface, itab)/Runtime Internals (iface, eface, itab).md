---
title: Runtime Internals (iface, eface, itab)
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Methods & Interfaces]]"
---

# Runtime Internals (iface, eface, itab)

iface and eface 2-word struct layouts, itab dispatch tables, global itab caching, and direct interface values.

```text
Runtime Internals (iface, eface, itab)
│
├── [[iface Runtime Memory Layout (tab, data)]]
├── [[eface Runtime Memory Layout (_type, data)]]
├── [[itab Virtual Method Table & Function Offsets]]
├── [[itab Global Cache & Dynamic Type Hash Tables]]
├── [[Direct Interface Values (Small Word Optimization)]]
└── [[Interface Boxing & Heap Allocation Triggers]]
```

---

## 🗂️ Topics

- [[iface Runtime Memory Layout (tab, data)]] — Non-empty interface layout: *itab table pointer and unsafe.Pointer data pointer (16 bytes).
- [[eface Runtime Memory Layout (_type, data)]] — Empty interface layout: *_type metadata pointer and unsafe.Pointer data pointer (16 bytes).
- [[itab Virtual Method Table & Function Offsets]] — Layout of itab: interface type, concrete type, hash code, and function pointers array [1]uintptr.
- [[itab Global Cache & Dynamic Type Hash Tables]] — How runtime dynamically computes and caches itab instances for type pairs in a global hash table.
- [[Direct Interface Values (Small Word Optimization)]] — Compiler optimization storing pointers and small scalar values directly inside the data word.
- [[Interface Boxing & Heap Allocation Triggers]] — When assigning a value to an interface forces heap allocation vs stack escape.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]

