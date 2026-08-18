---
title: Functions & Closures
tags:
  - golang
  - basics
parent: "[[Language Basics]]"
---

# Functions & Closures

Function declarations, multiple return values, named returns, variadic functions, closures, defer, and init.

```text
Functions & Closures
│
├── [[Function Declarations]]
├── [[Multiple Return Values]]
├── [[Named Return Values]]
├── [[Variadic Functions]]
├── [[Anonymous Functions]]
├── [[Closures]]
├── [[Closure Internals (Heap Escape)]]
├── [[Call by Value]]
├── [[defer Statement]]
├── [[defer Ordering and Evaluation]]
└── [[init() Function]]
```

---

## 🗂️ Topics

- [[Function Declarations]] — Signatures, parameter lists, return types, first-class citizen functions.
- [[Multiple Return Values]] — Idiomatic Go function signatures returning (result, error).
- [[Named Return Values]] — Naked returns and return value variable scoping.
- [[Variadic Functions]] — Passing variable numbers of arguments with ...T syntax.
- [[Anonymous Functions]] — Inline lambda functions and immediately invoked function expressions.
- [[Closures]] — Functions capturing variables from enclosing lexical scopes.
- [[Closure Internals (Heap Escape)]] — How the compiler moves captured variables to the heap.
- [[Call by Value]] — Go is strictly pass-by-value: passing pointers copies the pointer address.
- [[defer Statement]] — LIFO deferred function execution on function return.
- [[defer Ordering and Evaluation]] — Argument evaluation at defer time vs execution at return time.
- [[init() Function]] — Package initialization lifecycle and execution order across packages.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]
- 🎓 Root: [[Principal SWE]]
