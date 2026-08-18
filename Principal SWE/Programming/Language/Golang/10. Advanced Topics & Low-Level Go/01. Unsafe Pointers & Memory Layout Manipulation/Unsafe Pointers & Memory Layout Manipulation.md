---
title: Unsafe Pointers & Memory Layout Manipulation
tags:
  - golang
  - advanced
  - principal-swe
parent: "[[Advanced Topics & Low-Level Go]]"
---

# Unsafe Pointers & Memory Layout Manipulation

unsafe.Pointer casting, uintptr pointer arithmetic, zero-copy unsafe slicing, struct offsets, word alignment, and runtime.Pinner.

```text
Unsafe Pointers & Memory Layout Manipulation
│
├── [[unsafe.Pointer Mechanics & Arbitrary Pointer Casting]]
├── [[uintptr Pointer Arithmetic Rules & GC Safety Constraints]]
├── [[unsafe.Slice and unsafe.String Zero-Copy Constructors]]
├── [[Struct Field Offsets (unsafe.Offsetof & unsafe.Alignof)]]
├── [[Memory Alignment & Word Boundaries (8-byte Alignment)]]
└── [[Pointer Pinning (runtime.Pinner) for Cgo Integration]]
```

---

## 🗂️ Topics

- [[unsafe.Pointer Mechanics & Arbitrary Pointer Casting]] — Bypassing Go static type system, casting between arbitrary pointer types, and pointer safety rules.
- [[uintptr Pointer Arithmetic Rules & GC Safety Constraints]] — Safe conversion patterns: unsafe.Pointer to uintptr for offset arithmetic without GC pointer relocation traps.
- [[unsafe.Slice and unsafe.String Zero-Copy Constructors]] — Go 1.20+ safe zero-copy slice and string constructors replacing legacy reflect headers.
- [[Struct Field Offsets (unsafe.Offsetof & unsafe.Alignof)]] — Computing struct memory field byte offsets dynamically for high-speed binary serialization.
- [[Memory Alignment & Word Boundaries (8-byte Alignment)]] — Natural word boundaries, CPU memory alignment rules, struct padding, and unaligned panic prevention.
- [[Pointer Pinning (runtime.Pinner) for Cgo Integration]] — Pinning Go heap memory addresses to prevent GC movement during concurrent foreign C function calls.

---

## 🔗 References
- ⬆️ Parent: [[Advanced Topics & Low-Level Go]]

