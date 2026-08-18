---
title: Module Graph Pruning
tags:
  - golang
  - basics
  - dependencies
  - modules
  - resolution
parent: "[[Version Resolution]]"
---

# Module Graph Pruning

Since Go 1.17, the module graph loaded for `MVS` contains only the dependencies
that actually matter, instead of every `go.mod` transitively reachable.

## 1. Before and After

```text
go 1.16 and earlier            go 1.17 and later
────────────────────           ─────────────────
load every go.mod              load go.mod of direct deps in full
in the whole graph             + only the parts of deeper deps
                                 whose packages are imported
```

## 2. The Visible Consequence

`go.mod` gets much longer — every transitively-needed module is now listed
explicitly with `// indirect`:

```go
require (
    github.com/go-chi/chi/v5 v5.1.0
)

require (
    github.com/x/y v1.2.3 // indirect
    github.com/z/w v0.4.1 // indirect
)
```

This is not bloat. It is the pruned graph made **explicit**, so resolution no
longer needs to fetch deep `go.mod` files at all.

## 3. Why It Was Worth It

```text
go 1.16: `go build` may download dozens of go.mod files just to compute versions
go 1.17: the answer is already written down → far fewer network round trips
```

## 4. Gotchas

- Pruning is gated on the `go directive` of **your** module. Staying at
  `go 1.16` keeps the old, slower behaviour.
- Bumping from `go 1.16` to `go 1.17`+ makes `go mod tidy` rewrite `go.mod`
  substantially — expect a large, correct diff.
- Modules whose own `go` directive is < 1.17 are not pruned; their full graph is
  still loaded.

---

## 🔗 References
- ⬆️ Parent: `02. Version Resolution`
