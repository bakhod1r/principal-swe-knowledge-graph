---
title: Language Ergonomics & Iterators
tags:
  - golang
  - modern-go
  - principal-swe
parent: "[[Modern Language Features]]"
---

# Language Ergonomics & Iterators

Go 1.23+ iterators, range-over-func patterns, loopvar scoping changes, and built-in functions.

```text
Language Ergonomics & Iterators
│
├── [[Go 1.23 Iterators (iter.Seq, iter.Seq2)]]
├── [[Go 1.23 Pull Iterators (iter.Pull, iter.Pull2)]]
├── [[Go 1.22 Loop Variable Scoping (Loopvar)]]
└── [[min, max, and clear Builtin Functions]]
```

---

## 🗂️ Topics

- [[Go 1.23 Iterators (iter.Seq, iter.Seq2)]] — Standard iterator types, yielding values, writing custom range iterator functions.
- [[Go 1.23 Pull Iterators (iter.Pull, iter.Pull2)]] — Coroutine-based pull iterators yielding elements on demand with stop() cleanup.
- [[Go 1.22 Loop Variable Scoping (Loopvar)]] — Per-iteration variable scoping eliminating goroutine loop capture bugs.
- [[min, max, and clear Builtin Functions]] — Predeclared min/max for ordered types, clear() for zeroing slices and maps.

---

## 🔗 References
- ⬆️ Parent: [[Modern Language Features]]
- 🎓 Root: [[Principal SWE]]
