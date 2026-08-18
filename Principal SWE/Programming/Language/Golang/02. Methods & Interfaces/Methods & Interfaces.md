---
title: Methods & Interfaces
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Golang]]"
---

# 🧩 Methods & Interfaces

Methods vs Functions, Receivers, Interface Basics, Empty Interfaces (any), Embedding, Type Assertions, Type Switch, and iface/eface Internals.

```text
Methods & Interfaces
│
├── [[Methods & Receivers|01. Methods & Receivers]]
│   ├── [[Methods vs Functions]]
│   ├── [[Pointer Receivers]]
│   ├── [[Value Receivers]]
│   ├── [[Receiver Choice Heuristics]]
│   ├── [[Method Sets]]
│   ├── [[Method Values]]
│   ├── [[Methods on Defined Types]]
│   ├── [[Cross-Package Method Rules]]
│   └── [[Struct Method Promotion]]
├── [[Interfaces & Polymorphism|02. Interfaces & Polymorphism]]
│   ├── [[Interface Basics]]
│   ├── [[Empty Interface (any)]]
│   ├── [[Embedding Interfaces]]
│   ├── [[Type Assertions]]
│   ├── [[Type Switch]]
│   ├── [[Common Standard Library Interfaces]]
│   └── [[Sealed Interfaces]]
└── [[Runtime Internals & Architecture|03. Runtime Internals & Architecture]]
│   ├── [[iface and eface Structs]]
│   ├── [[itab Dynamic Dispatch Table]]
│   ├── [[Interface Allocation Cost]]
│   ├── [[Interface Best Practices]]
│   └── [[Interface Anti-Patterns]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Methods & Receivers|01. Methods & Receivers]]
- [[Methods vs Functions]] — Comparison of pure functions vs methods with receivers.
- [[Pointer Receivers]] — Mutating state, avoiding copying large structs, consistency rules.
- [[Value Receivers]] — Immutability, value semantics, copy overhead considerations.
- [[Receiver Choice Heuristics]] — Guidelines on when to choose pointer vs value receiver.
- [[Method Sets]] — Rules governing which methods belong to T and *T.
- [[Method Values]] — Treating methods as first-class functions (T.Method vs instance.Method).
- [[Methods on Defined Types]] — Attaching methods to non-struct defined types (type MyInt int).
- [[Cross-Package Method Rules]] — Receiver type locality rules (cannot define methods on foreign types).
- [[Struct Method Promotion]] — Method inheritance-like behavior through struct embedding.
### 2. 📂 [[Interfaces & Polymorphism|02. Interfaces & Polymorphism]]
- [[Interface Basics]] — Implicit satisfaction, single-method interface design philosophy.
- [[Empty Interface (any)]] — Working with unknown types, type safety considerations, boxing cost.
- [[Embedding Interfaces]] — Interface composition (io.ReadWriter = io.Reader + io.Writer).
- [[Type Assertions]] — Extracting concrete types from interfaces (x.(T) and comma-ok idiom).
- [[Type Switch]] — Type-based dispatching across multiple interface implementors.
- [[Common Standard Library Interfaces]] — Core contracts: io.Reader, io.Writer, fmt.Stringer, error, sort.Interface.
- [[Sealed Interfaces]] — Restricting implementation to package boundaries via unexported method tags.
### 3. 📂 [[Runtime Internals & Architecture|03. Runtime Internals & Architecture]]
- [[iface and eface Structs]] — Two-word structure: _type/itab pointer + data pointer.
- [[itab Dynamic Dispatch Table]] — Virtual method table construction, caching, and dispatch cost.
- [[Interface Allocation Cost]] — When assigning a concrete value to an interface causes a heap allocation.
- [[Interface Best Practices]] — Accept interfaces, return structs; small interfaces; consumer-defined interfaces.
- [[Interface Anti-Patterns]] — Premature abstraction, interface pollution, returning interfaces.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

