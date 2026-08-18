---
title: Methods & Receivers
tags:
  - golang
  - methods
parent: "[[Methods & Interfaces]]"
---

# Methods & Receivers

Method declarations, receiver semantics, method sets, and method expressions.

```text
Methods & Receivers
│
├── [[Methods vs Functions]]
├── [[Pointer Receivers]]
├── [[Value Receivers]]
├── [[Receiver Choice Heuristics]]
├── [[Method Sets]]
├── [[Method Values]]
├── [[Method Expressions]]
├── [[Methods on Defined Types]]
├── [[Cross-Package Method Rules]]
└── [[Struct Method Promotion]]
```

---

## 🗂️ Topics

- [[Methods vs Functions]] — Comparing method receivers with first-class function parameters.
- [[Pointer Receivers]] — Mutating receiver state, avoiding large struct copies, receiver consistency rules.
- [[Value Receivers]] — Immutability semantics, concurrent read safety, and value copying overhead.
- [[Receiver Choice Heuristics]] — Decision tree for pointer vs value receiver selection.
- [[Method Sets]] — Rules governing method sets for T and *T and interface satisfaction.
- [[Method Values]] — Binding an instance to a method returning a first-class function (obj.Method).
- [[Method Expressions]] — Treating methods as static functions with explicit receiver parameter (Type.Method).
- [[Methods on Defined Types]] — Attaching methods to custom non-struct types (type MyDuration int64).
- [[Cross-Package Method Rules]] — Type locality rules prohibiting method definitions on foreign package types.
- [[Struct Method Promotion]] — Transparently promoting embedded inner struct methods to outer struct.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
