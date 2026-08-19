---
title: Custom Iterators & Range Over Functions (Go 1.23+)
tags:
  - golang
  - modern-go
  - principal-swe
parent: "[[Modern Language Features]]"
---

# Custom Iterators & Range Over Functions (Go 1.23+)

iter.Seq and iter.Seq2 generic signatures, compiler loop lowering, push vs pull iterators, coroutine runtime (coro.go), and combinators.

```text
Custom Iterators & Range Over Functions (Go 1.23+)
│
├── [[iter.Seq and iter.Seq2 Standard Contracts]]
├── [[Range Over Functions Compiler Lowering]]
├── [[Push vs Pull Iterators (iter.Pull & iter.Pull2)]]
├── [[Runtime Coroutine Architecture (coro.go)]]
├── [[Functional Iterator Combinators (Filter, Map, Zip, Take)]]
└── [[Iterator Inlining & Allocation Profiles]]
```

---

## 🗂️ Topics

- [[iter.Seq and iter.Seq2 Standard Contracts]] — Standard generic iterator type signatures (iter.Seq[V] and iter.Seq2[K, V]) for collections.
- [[Range Over Functions Compiler Lowering]] — How the compiler transforms for v := range fn into callback invocations with early break yields.
- [[Push vs Pull Iterators (iter.Pull & iter.Pull2)]] — Transforming push yield functions into stateful next(), stop() pull iterators.
- [[Runtime Coroutine Architecture (coro.go)]] — The underlying runtime coroutine stack-switching engine powering iter.Pull in Go.
- [[Functional Iterator Combinators (Filter, Map, Zip, Take)]] — Composing zero-allocation streaming data pipelines using iter.Seq.
- [[Iterator Inlining & Allocation Profiles]] — Compiler inlining rules for iterator callback functions to prevent closure heap allocations.

---

## 🔗 References
- ⬆️ Parent: [[Modern Language Features]]

