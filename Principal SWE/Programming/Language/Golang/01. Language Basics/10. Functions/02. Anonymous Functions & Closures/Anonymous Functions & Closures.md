---
title: Anonymous Functions & Closures
tags:
  - golang
  - functions
  - principal-swe
parent: "[[Functions]]"
---

# Anonymous Functions & Closures

Function literals, lexical scoping, closure capture, and heap escape.

```text
Anonymous Functions & Closures
│
├── [[Anonymous Functions (Function Literals)]]
├── [[Closures & Lexical Scoping]]
├── [[Variable Capture Mechanics]]
├── [[Loop Variable Capture Trap]]
└── [[Closure Heap Escape Analysis]]
```

---

## 🗂️ Topics

- [[Anonymous Functions (Function Literals)]] — Inline lambda functions and immediately invoked function expressions (IIFE).
- [[Closures & Lexical Scoping]] — Functions capturing and binding variables from enclosing lexical scopes.
- [[Variable Capture Mechanics]] — Capturing variables by reference and heap allocation promotion.
- [[Loop Variable Capture Trap]] — Capturing loop iteration variables in goroutines (pre-Go 1.22 vs Go 1.22+ semantics).
- [[Closure Heap Escape Analysis]] — How compiler moves captured closure environment variables to the heap.

---

## 🔗 References
- ⬆️ Parent: [[Functions]]
- 🎓 Root: [[Principal SWE]]
