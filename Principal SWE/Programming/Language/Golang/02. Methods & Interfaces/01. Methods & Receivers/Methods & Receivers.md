---
title: Methods & Receivers
tags:
  - golang
  - methods-and-interfaces
parent: "[[Methods & Interfaces]]"
---

# Methods & Receivers

Method syntax, value vs pointer receivers, method sets, and method expressions.

```text
Methods & Receivers
│
├── [[Methods vs Functions]]
├── [[Pointer Receivers]]
├── [[Value Receivers]]
├── [[Receiver Choice Heuristics]]
├── [[Method Sets]]
├── [[Method Values and Expressions]]
├── [[Methods on Defined Types]]
├── [[Cross-Package Method Rules]]
└── [[Struct Method Promotion]]
```

---

## 🗂️ Topics

- [[Methods vs Functions]] — Comparison of pure functions vs methods with receivers.
- [[Pointer Receivers]] — Mutating state, avoiding copying large structs, consistency rules.
- [[Value Receivers]] — Immutability, value semantics, copy overhead considerations.
- [[Receiver Choice Heuristics]] — Guidelines on when to choose pointer vs value receiver.
- [[Method Sets]] — Rules governing which methods belong to T and *T.
- [[Method Values and Expressions]] — Treating methods as first-class functions (T.Method vs instance.Method).
- [[Methods on Defined Types]] — Attaching methods to non-struct defined types (type MyInt int).
- [[Cross-Package Method Rules]] — Receiver type locality rules (cannot define methods on foreign types).
- [[Struct Method Promotion]] — Method inheritance-like behavior through struct embedding.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
