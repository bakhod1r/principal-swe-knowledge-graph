---
title: Modern Language Features
tags:
  - golang
  - modern-go
  - principal-swe
parent: "[[Golang]]"
---

# ✨ Modern Language Features

Recent Go enhancements: iterators (iter, range over func), Go 1.22 loop variable scoping, min/max/clear builtins, and generic type aliases.

```text
Modern Language Features
│
├── [[Language Ergonomics & Iterators|01. Language Ergonomics & Iterators]]
│   ├── [[Go 1.23 Iterators (iter.Seq, iter.Seq2)]]
│   ├── [[Go 1.22 Loop Variable Scoping (Loopvar)]]
│   └── [[min, max, and clear Builtin Functions]]
└── [[Type System & Stdlib Additions|02. Type System & Stdlib Additions]]
│   ├── [[Generic Type Aliases (Go 1.24+)]]
│   ├── [[slices and maps Standard Packages]]
│   ├── [[cmp.Ordered and cmp.Compare]]
│   └── [[synctest Experimental Testing Package]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Language Ergonomics & Iterators|01. Language Ergonomics & Iterators]]
- [[Go 1.23 Iterators (iter.Seq, iter.Seq2)]] — Standard iterator types, yielding values, writing custom range iterator functions.
- [[Go 1.22 Loop Variable Scoping (Loopvar)]] — Per-iteration variable scoping eliminating goroutine loop capture bugs.
- [[min, max, and clear Builtin Functions]] — Predeclared min/max for ordered types, clear() for zeroing slices and maps.
### 2. 📂 [[Type System & Stdlib Additions|02. Type System & Stdlib Additions]]
- [[Generic Type Aliases (Go 1.24+)]] — type MyList[T] = other.List[T] generic alias syntax and migration patterns.
- [[slices and maps Standard Packages]] — Standard library algorithm functions for slices and maps without external dependencies.
- [[cmp.Ordered and cmp.Compare]] — Standard ordering comparison interface and three-way compare function.
- [[synctest Experimental Testing Package]] — Virtual time testing package for deterministic testing of concurrent code.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

