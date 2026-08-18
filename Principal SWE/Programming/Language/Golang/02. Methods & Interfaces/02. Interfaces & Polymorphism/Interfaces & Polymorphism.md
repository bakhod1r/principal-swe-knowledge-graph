---
title: Interfaces & Polymorphism
tags:
  - golang
  - methods-and-interfaces
parent: "[[Methods & Interfaces]]"
---

# Interfaces & Polymorphism

Implicit interface satisfaction, empty interface (any), embedding, assertions, and switches.

```text
Interfaces & Polymorphism
│
├── [[Interfaces Basics]]
├── [[Empty Interfaces (any)]]
├── [[Embedding Interfaces]]
├── [[Type Assertions]]
├── [[Type Switch]]
├── [[Common Standard Library Interfaces]]
└── [[Sealed Interfaces]]
```

---

## 🗂️ Topics

- [[Interfaces Basics]] — Implicit satisfaction, single-method interface design philosophy.
- [[Empty Interfaces (any)]] — Working with unknown types, type safety considerations, boxing cost.
- [[Embedding Interfaces]] — Interface composition (io.ReadWriter = io.Reader + io.Writer).
- [[Type Assertions]] — Extracting concrete types from interfaces (x.(T) and comma-ok idiom).
- [[Type Switch]] — Type-based dispatching across multiple interface implementors.
- [[Common Standard Library Interfaces]] — Core contracts: io.Reader, io.Writer, fmt.Stringer, error, sort.Interface.
- [[Sealed Interfaces]] — Restricting implementation to package boundaries via unexported method tags.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
