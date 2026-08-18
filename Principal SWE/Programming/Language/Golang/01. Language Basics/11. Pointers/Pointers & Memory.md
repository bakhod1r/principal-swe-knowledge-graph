---
title: Pointers & Memory
tags:
  - golang
  - basics
parent: "[[Language Basics]]"
---

# Pointers & Memory

Pointer semantics, address-of (&), dereference (*), pointer arithmetic restrictions, unsafe, and GC.

```text
Pointers & Memory
│
├── [[Pointers Basics]]
├── [[Pointer vs Value Semantics]]
├── [[Pointers with Structs]]
├── [[Pointers with Slices & Maps]]
├── [[unsafe.Pointer and uintptr]]
├── [[nil Pointer Dereference]]
├── [[Memory Management & Escape Analysis]]
└── [[Garbage Collection Overview]]
```

---

## 🗂️ Topics

- [[Pointers Basics]] — Memory addresses, pointer types (*T), address-of operator (&), dereference (*).
- [[Pointer vs Value Semantics]] — Copying data vs sharing memory, mutation vs immutability.
- [[Pointers with Structs]] — Dot notation auto-dereferencing (p.Field equivalent to (*p).Field).
- [[Pointers with Slices & Maps]] — Reference types containing internal pointers vs pointer-to-slice.
- [[unsafe.Pointer and uintptr]] — Bypassing Go type safety, pointer arithmetic, and alignment.
- [[nil Pointer Dereference]] — Runtime panics on dereferencing nil and defensive nil checking patterns.
- [[Memory Management & Escape Analysis]] — Stack vs heap allocation rules determined at compile time.
- [[Garbage Collection Overview]] — Automatic memory reclamation with concurrent tricolor GC.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]
- 🎓 Root: [[Principal SWE]]
