---
title: Type Assertions, Switches & Inspection
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Go Interfaces]]"
---

# Type Assertions, Switches & Inspection

Dynamic type extraction, comma-ok idiom, type switches, typed nil interface trap, and defensive nil checking.

```text
Type Assertions, Switches & Inspection
│
├── [[Type Assertions & Comma-Ok Idiom]]
├── [[Unchecked Type Assertions & Panic Hazards]]
├── [[Type Switch Dynamic Dispatching]]
├── [[The Typed Nil Interface Trap (*MyType(nil) != nil)]]
└── [[Defensive Nil Checking Patterns]]
```

---

## 🗂️ Topics

- [[Type Assertions & Comma-Ok Idiom]] — Extracting concrete types from interfaces safely using val, ok := i.(ConcreteType).
- [[Unchecked Type Assertions & Panic Hazards]] — Why unchecked assertions (val := i.(ConcreteType)) cause fatal runtime panics on mismatch.
- [[Type Switch Dynamic Dispatching]] — Multi-way type branching using switch v := i.(type) syntax.
- [[The Typed Nil Interface Trap (*MyType(nil) != nil)]] — Why an interface holding a typed nil pointer is not equal to nil (iface.tab != nil).
- [[Defensive Nil Checking Patterns]] — Architectural patterns for safely checking both interface nil and underlying pointer nil.

---

## 🔗 References
- ⬆️ Parent: `Methods & Interfaces`

