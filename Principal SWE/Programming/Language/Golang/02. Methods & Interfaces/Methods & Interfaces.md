---
title: Methods & Interfaces
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Golang]]"
---

# 🧩 Methods & Interfaces

Methods, pointer/value receivers, method sets, interface contracts, iface/eface runtime layout, itab dynamic dispatch table, and interface architecture.

```text
Methods & Interfaces
│
├── [[Methods & Receivers|01. Methods & Receivers]]
│   ├── [[Methods vs Functions]]
│   ├── [[Pointer Receivers]]
│   ├── [[Value Receivers]]
│   ├── [[Receiver Choice Heuristics]]
│   ├── [[Method Sets & Assignability Rules]]
│   ├── [[Method Values]]
│   ├── [[Method Expressions]]
│   ├── [[Methods on Defined Types]]
│   ├── [[Cross-Package Method Rules]]
│   ├── [[Struct Method Promotion]]
│   └── [[Auto-Referencing and Auto-Dereferencing]]
├── [[Interfaces & Polymorphism|02. Interfaces & Polymorphism]]
│   ├── [[Interface Basics]]
│   ├── [[Empty Interface (any)]]
│   ├── [[Embedding Interfaces]]
│   ├── [[Type Assertions]]
│   ├── [[Type Switch]]
│   ├── [[Common Standard Library Interfaces]]
│   ├── [[Sealed Interfaces]]
│   ├── [[Interface Nil Checking Trap]]
│   └── [[Interface Segregation Principle in Go]]
└── [[Runtime Internals & Architecture|03. Runtime Internals & Architecture]]
│   ├── [[iface and eface Structs]]
│   ├── [[itab Dynamic Dispatch Table]]
│   ├── [[itab Global Cache & Hash Tables]]
│   ├── [[Interface Allocation Cost]]
│   ├── [[Direct Interface Values (Optimization)]]
│   ├── [[Interface Best Practices]]
│   └── [[Interface Anti-Patterns]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Methods & Receivers|01. Methods & Receivers]]
- [[Methods vs Functions]] — Comparing method receivers with first-class function parameters.
- [[Pointer Receivers]] — Mutating receiver state, avoiding large struct copies, receiver consistency rules.
- [[Value Receivers]] — Immutability semantics, concurrent read safety, and value copying overhead.
- [[Receiver Choice Heuristics]] — Decision tree for pointer vs value receiver selection.
- [[Method Sets & Assignability Rules]] — Rules governing method sets for T and *T and interface satisfaction.
- [[Method Values]] — Binding an instance to a method returning a first-class function (obj.Method).
- [[Method Expressions]] — Treating methods as static functions with explicit receiver parameter (Type.Method).
- [[Methods on Defined Types]] — Attaching methods to custom non-struct types (type MyDuration int64).
- [[Cross-Package Method Rules]] — Type locality rules prohibiting method definitions on foreign package types.
- [[Struct Method Promotion]] — Transparently promoting embedded inner struct methods to outer struct.
- [[Auto-Referencing and Auto-Dereferencing]] — How the Go compiler automatically inserts & or * on method calls.
### 2. 📂 [[Interfaces & Polymorphism|02. Interfaces & Polymorphism]]
- [[Interface Basics]] — Implicit satisfaction, structural typing, and consumer-defined interface contracts.
- [[Empty Interface (any)]] — Working with dynamic unknown types, boxing, and type safety tradeoffs.
- [[Embedding Interfaces]] — Composing fine-grained interfaces (io.ReadWriter = io.Reader + io.Writer).
- [[Type Assertions]] — Dynamic type extraction with comma-ok idiom (v, ok := i.(T)).
- [[Type Switch]] — Multi-type branch dispatching using switch v := i.(type).
- [[Common Standard Library Interfaces]] — Core contracts: io.Reader, io.Writer, io.Closer, fmt.Stringer, error, sort.Interface.
- [[Sealed Interfaces]] — Restricting external implementations using unexported method tokens.
- [[Interface Nil Checking Trap]] — Why an interface holding a typed nil pointer is not equal to nil (iface.data vs iface.tab).
- [[Interface Segregation Principle in Go]] — Designing minimal single-method interfaces defined at point of consumption.
### 3. 📂 [[Runtime Internals & Architecture|03. Runtime Internals & Architecture]]
- [[iface and eface Structs]] — Two-word runtime interface representation: tab/type pointer and data pointer.
- [[itab Dynamic Dispatch Table]] — Interface table layout, method offset resolution, and runtime dispatch performance.
- [[itab Global Cache & Hash Tables]] — How runtime dynamically computes and caches itab instances for type pairs.
- [[Interface Allocation Cost]] — When converting concrete values to interfaces triggers heap allocations.
- [[Direct Interface Values (Optimization)]] — Compiler optimization storing pointers and small words directly inside the data word.
- [[Interface Best Practices]] — Accept interfaces, return structs; keep interfaces small; define interfaces in consumer package.
- [[Interface Anti-Patterns]] — Interface pollution, premature abstraction, mocking what you do not own.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
