---
title: Loop Variable Scoping & Range Over Int (Go 1.22+)
tags:
  - golang
  - modern-go
  - principal-swe
parent: "[[Modern Language Features]]"
---

# Loop Variable Scoping & Range Over Int (Go 1.22+)

Per-iteration variable scoping eliminating goroutine closure capture bugs, for i := range n syntax, and loopvar migration.

```text
Loop Variable Scoping & Range Over Int (Go 1.22+)
│
├── [[Per-Iteration Loop Scoping Semantics]]
├── [[Range Over Integer Syntax (for i := range n)]]
└── [[Loopvar Migration & Bisect Tooling]]
```

---

## 🗂️ Topics

- [[Per-Iteration Loop Scoping Semantics]] — Creating a new lexical variable instance per iteration, eliminating the classic goroutine closure bug.
- [[Range Over Integer Syntax (for i := range n)]] — for i := range 10 syntactic simplification and compiler bounds optimizations.
- [[Loopvar Migration & Bisect Tooling]] — Using go test -gcflags=-d=loopvar=2 and go vet to detect semantic changes in legacy code.

---

## 🔗 References
- ⬆️ Parent: [[Modern Language Features]]

