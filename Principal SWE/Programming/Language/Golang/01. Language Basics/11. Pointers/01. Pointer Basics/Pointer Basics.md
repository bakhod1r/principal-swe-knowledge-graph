- [[Pointer Aliasing & Optimization Barriers]] — How multiple pointers referencing the same memory location limit compiler optimizations.

- [[Dangling Pointers Immunity in Garbage Collected Go]] — Why Go is immune to classic C use-after-free and dangling pointer memory corruption.

---
title: Pointer Basics
tags:
  - golang
  - pointers
  - principal-swe
parent: "[[Pointers]]"
---

# Pointer Basics

Memory addresses, pointer types (*T), address-of (&), dereference (*), and value semantics.

```text
Pointer Basics
│
├── [[Memory Addresses & Pointer Types (*T)]]
├── [[Address-of (&) and Dereference (*)]]
├── [[Pointer vs Value Semantics]]
├── [[Auto-Dereferencing with Dot Notation]]
└── [[nil Pointer & Panic Dereference]]
```

---

## 🗂️ Topics

- [[Memory Addresses & Pointer Types (*T)]] — Memory locations, pointer types (*T), and zero value (nil).
- [[Address-of (&) and Dereference (*)]] — Taking memory addresses with & and accessing underlying values with *.
- [[Pointer vs Value Semantics]] — Copying data vs sharing memory, mutation semantics, and immutability.
- [[Auto-Dereferencing with Dot Notation]] — Automatic pointer dereferencing on struct field and method access (p.Field).
- [[nil Pointer & Panic Dereference]] — Runtime panics on dereferencing nil pointers and defensive checking patterns.

---

## 🔗 References
- ⬆️ Parent: [[Pointers]]

