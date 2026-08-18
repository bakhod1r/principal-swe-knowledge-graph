---
title: Standard Library Generics
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Generics]]"
---

# Standard Library Generics

Standard library generic packages: slices, maps, cmp, sync.Map wrappers, and atomic.Pointer.

```text
Standard Library Generics
│
├── [[slices Package Deep Dive]]
├── [[maps Package Deep Dive]]
├── [[cmp Package Deep Dive]]
├── [[sync.Map Typesafe Generic Wrapper]]
└── [[atomic.Pointer[T] Type Safety (Go 1.19+)]]
```

---

## 🗂️ Topics

- [[slices Package Deep Dive]] — Generic slice algorithms: slices.Sort, slices.BinarySearch, slices.Contains, slices.Clone, slices.Delete.
- [[maps Package Deep Dive]] — Generic map helpers: maps.Clone, maps.Copy, maps.Equal, maps.DeleteFunc.
- [[cmp Package Deep Dive]] — Ordering functions: cmp.Compare, cmp.Less, and cmp.Or default fallback values.
- [[sync.Map Typesafe Generic Wrapper]] — Building a type-safe generic wrapper over sync.Map without casting.
- [[atomic.Pointer[T] Type Safety (Go 1.19+)]] — Lock-free atomic pointer storage with full generic compile-time type safety.

---

## 🔗 References
- ⬆️ Parent: [[Generics]]
- 🎓 Root: [[Principal SWE]]
