---
title: go directive
tags:
  - golang
  - basics
  - dependencies
  - modules
  - go.mod
  - compatibility
parent: "[[Core Concepts]]"
---

# `go` directive

Declares the **minimum Go language version** the module requires, and selects the
language semantics the compiler applies.

## 1. Syntax

```go
go 1.24.0
```

Since Go 1.21 the value is an exact version with well-defined meaning, not a hint.

## 2. What It Actually Controls

| Effect | Detail |
|---|---|
| Language semantics | Per-file gating — e.g. Go 1.22 loop-variable scoping only applies at `go 1.22`+ |
| Minimum toolchain | A lower toolchain refuses to build, or upgrades via `GOTOOLCHAIN` |
| `go mod tidy` behaviour | Determines which requirements are kept for older consumers |
| Default `godebug` values | Compatibility switches default to that version's behaviour |

## 3. The Loop Variable Example

```go
for _, v := range items {
    go func() { fmt.Println(v) }()   // go 1.21: all print the last item
}                                     // go 1.22: each prints its own
```

Identical source, different behaviour, decided solely by this directive.

## 4. Gotchas

- Raising it is a **breaking change for consumers** on older toolchains.
- `go mod tidy` uses it to decide graph pruning depth — bumping it can shrink
  `go.mod` significantly. See `Module Graph Pruning`.
- It is a *minimum*, not a pin. Pin the toolchain with `toolchain`.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
