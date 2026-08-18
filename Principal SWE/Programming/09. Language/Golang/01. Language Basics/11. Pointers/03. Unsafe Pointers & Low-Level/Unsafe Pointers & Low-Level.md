---
title: Unsafe Pointers & Low-Level
tags:
  - golang
  - pointers
  - principal-swe
parent: "[[Pointers]]"
---

# Unsafe Pointers & Low-Level

unsafe.Pointer, uintptr arithmetic, alignment, and zero-copy casting.

```text
Unsafe Pointers & Low-Level
│
├── [[unsafe.Pointer Mechanics]]
├── [[uintptr Pointer Arithmetic Rules]]
├── [[Memory Alignment & Word Boundaries]]
├── [[Zero-Copy Conversions (unsafe.Slice, unsafe.String)]]
└── [[Pointer Pinning (runtime.Pinner)]]
```

---

## 🗂️ Topics

- [[unsafe.Pointer Mechanics]] — Bypassing Go type safety, casting between arbitrary pointer types.
- [[uintptr Pointer Arithmetic Rules]] — Converting unsafe.Pointer to uintptr for offset arithmetic and GC safety rules.
- [[Memory Alignment & Word Boundaries]] — CPU word alignment (8-byte on 64-bit), unsafe.Alignof, and unsafe.Offsetof.
- [[Zero-Copy Conversions (unsafe.Slice, unsafe.String)]] — Zero-copy slicing and string casting without memory allocation.
- [[Pointer Pinning (runtime.Pinner)]] — Preventing GC from moving or reclaiming memory passed to foreign C code.

---

## 🔗 References
- ⬆️ Parent: [[Pointers]]

