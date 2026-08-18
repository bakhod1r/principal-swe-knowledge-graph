---
title: Method Values, Expressions & Promotion
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Methods & Interfaces]]"
---

# Method Values, Expressions & Promotion

Method values, method expressions, struct embedding promotion, method shadowing, and compiler auto-referencing.

```text
Method Values, Expressions & Promotion
│
├── [[Method Values (Instance Binding & Closures)]]
├── [[Method Expressions (Static Functions)]]
├── [[Embedded Struct Method Promotion]]
├── [[Method Shadowing & Ambiguity in Nested Structs]]
└── [[Auto-Referencing and Auto-Dereferencing]]
```

---

## 🗂️ Topics

- [[Method Values (Instance Binding & Closures)]] — Binding a struct instance to a method, returning a first-class closure, and escape analysis.
- [[Method Expressions (Static Functions)]] — Calling methods as static unbound functions with explicit receiver parameters (Type.Method).
- [[Embedded Struct Method Promotion]] — Composition over inheritance: promoting inner embedded methods to outer struct method set.
- [[Method Shadowing & Ambiguity in Nested Structs]] — Outer method overriding inner methods and resolving duplicate method collision errors.
- [[Auto-Referencing and Auto-Dereferencing]] — How the compiler automatically inserts & or * when invoking methods across pointer/value boundaries.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]
- 🎓 Root: [[Principal SWE]]
