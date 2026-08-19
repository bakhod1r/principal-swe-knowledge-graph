---
title: Pointers
tags:
  - golang
  - pointers
  - principal-swe
parent: "[[Language Basics]]"
---

# Pointers

Pointers, pointer semantics, reference types, unsafe pointer arithmetic, escape analysis, and memory management.

```text
Pointers
│
├── [[Pointer Basics]]
├── [[Pointers with Types]]
├── [[Unsafe Pointers & Low-Level]]
├── [[Pointer Memory Management]]
└── [[Escape Analysis Internals]]
```

---

## 🗂️ Topics

- [[Pointer Basics]] — Memory addresses, pointer types (`*T`), address-of (`&`), dereference (`*`), and value semantics.
- [[Pointers with Types]] — Pointers with structs, slices, maps, channels, and interface pointer traps.
- [[Unsafe Pointers & Low-Level]] — `unsafe.Pointer`, `uintptr` arithmetic, memory word boundaries, and runtime pointer pinning.
- [[Pointer Memory Management]] — TCMalloc allocation hierarchy (`mcache`, `mcentral`, `mheap`), hybrid write barriers, and tricolor GC.
- [[Escape Analysis Internals]] — Stack vs heap allocation rules, pointer flow graph analysis, and heap escape triggers.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]
- 📚 Module: `Language Basics`
