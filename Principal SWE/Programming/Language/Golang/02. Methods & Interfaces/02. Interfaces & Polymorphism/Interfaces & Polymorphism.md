---
title: Interfaces & Polymorphism
tags:
  - golang
  - interfaces
parent: "[[Methods & Interfaces]]"
---

# Interfaces & Polymorphism

Interface contracts, implicit implementation, polymorphism, type assertions, and standard interfaces.

```text
Interfaces & Polymorphism
│
├── [[Interface Basics]]
├── [[Empty Interface (any)]]
├── [[Embedding Interfaces]]
├── [[Type Assertions]]
├── [[Type Switch]]
├── [[Common Standard Library Interfaces]]
└── [[Sealed Interfaces]]
```

---

## 🗂️ Topics

- [[Interface Basics]] — Implicit satisfaction, structural typing, and consumer-defined interface contracts.
- [[Empty Interface (any)]] — Working with dynamic unknown types, boxing, and type safety tradeoffs.
- [[Embedding Interfaces]] — Composing fine-grained interfaces (io.ReadWriter = io.Reader + io.Writer).
- [[Type Assertions]] — Dynamic type extraction with comma-ok idiom (v, ok := i.(T)).
- [[Type Switch]] — Multi-type branch dispatching using switch v := i.(type).
- [[Common Standard Library Interfaces]] — Core contracts: io.Reader, io.Writer, io.Closer, fmt.Stringer, error, sort.Interface.
- [[Sealed Interfaces]] — Restricting external implementations using unexported method tokens.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
