---
title: Interfaces & Polymorphism
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Methods & Interfaces]]"
---

# Interfaces & Polymorphism

Implicit satisfaction, empty interface (any), composition, type assertions, type switches, and ISP.

```text
Interfaces & Polymorphism
│
├── [[Interface Basics]]
├── [[Empty Interface (any)]]
├── [[Embedding Interfaces]]
├── [[Type Assertions]]
├── [[Type Switch]]
├── [[Common Standard Library Interfaces]]
├── [[Sealed Interfaces]]
├── [[Interface Nil Checking Trap]]
└── [[Interface Segregation Principle in Go]]
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
- [[Interface Nil Checking Trap]] — Why an interface holding a typed nil pointer is not equal to nil (iface.data vs iface.tab).
- [[Interface Segregation Principle in Go]] — Designing minimal single-method interfaces defined at point of consumption.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
