---
title: Function Basics
tags:
  - hub
parent: "[[Functions]]"
---

- [[Named Return Values & Naked Returns Gotchas]] — Naked return readability anti-patterns and deferred named return mutation.

- [[Variadic Functions Slice Allocation & Memory Lifecycle]] — How ...T variadic arguments construct an implicit slice on stack or heap.

---
title: Function Basics
tags:
  - golang
  - functions
  - principal-swe
parent: "`Functions (Clean Code)`"
---

# Function Basics

Declarations, signatures, multiple return values, named returns, and variadics.

```text
Function Basics
│
├── [[Function Declarations & Signatures]]
├── [[Multiple Return Values]]
├── [[Named Return Values & Naked Returns]]
├── [[Variadic Functions (...T)]]
└── [[Call by Value Semantics]]
```

---

## 🗂️ Topics

- [[Function Declarations & Signatures]] — Signatures, parameter lists, return types, first-class citizen functions.
- [[Multiple Return Values]] — Idiomatic Go function signatures returning (result, error).
- [[Named Return Values & Naked Returns]] — Named return variables, documentation clarity, and naked return caveats.
- [[Variadic Functions (...T)]] — Passing variable numbers of arguments with ...T syntax and slice expansion.
- [[Call by Value Semantics]] — Go is strictly pass-by-value: passing pointers copies the pointer address.

---

## 🔗 References
- ⬆️ Parent: `Functions (Clean Code)`

