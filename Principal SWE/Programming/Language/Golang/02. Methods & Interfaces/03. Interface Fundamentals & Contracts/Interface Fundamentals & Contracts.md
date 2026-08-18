---
title: Interface Fundamentals & Contracts
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Methods & Interfaces]]"
---

# Interface Fundamentals & Contracts

Structural typing, implicit satisfaction, consumer-defined contracts, embedding, empty interface (any), and sealed interfaces.

```text
Interface Fundamentals & Contracts
│
├── [[Structural Typing & Implicit Satisfaction]]
├── [[Consumer-Defined Interfaces (Accept Interfaces, Return Structs)]]
├── [[Interface Composition & Embedding]]
├── [[Empty Interface (any) & Type Boxing Mechanics]]
├── [[Common Standard Library Contracts]]
└── [[Sealed Interfaces (Restricting External Implementations)]]
```

---

## 🗂️ Topics

- [[Structural Typing & Implicit Satisfaction]] — Why Go rejected explicit implements keywords in favor of compile-time duck typing.
- [[Consumer-Defined Interfaces (Accept Interfaces, Return Structs)]] — Staff design principle: defining minimal interfaces in consumer packages.
- [[Interface Composition & Embedding]] — Composing fine-grained interfaces (io.ReadWriteCloser = io.Reader + io.Writer + io.Closer).
- [[Empty Interface (any) & Type Boxing Mechanics]] — Working with arbitrary types, boxing overhead, and type safety tradeoffs.
- [[Common Standard Library Contracts]] — Core contracts: io.Reader, io.Writer, io.Closer, fmt.Stringer, error, sort.Interface.
- [[Sealed Interfaces (Restricting External Implementations)]] — Restricting interface implementations to internal packages via unexported method tokens.

---

## 🔗 References
- ⬆️ Parent: [[Methods & Interfaces]]

