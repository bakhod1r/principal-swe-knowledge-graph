---
title: go mod graph
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
  - debugging
parent: "[[Module Commands]]"
---

# `go mod graph`

Prints the module requirement graph as `from to` edge pairs, one per line.

```bash
go mod graph
go mod graph | grep 'golang.org/x/text'
go mod graph -x
```

## 1. Output Shape

```text
github.com/me/api github.com/go-chi/chi/v5@v5.1.0
github.com/go-chi/chi/v5@v5.1.0 golang.org/x/text@v0.14.0
```

The left side of the first column has no version — that is the main module.

## 2. Answering "Why Is This Version Selected?"

```bash
# every module that requires x/text, and at which version
go mod graph | grep ' golang.org/x/text@' | sort -u
```

`MVS` picks the **highest** of those. Combine with `go mod why` for the
package-level import path.

## 3. Finding Version Conflicts

```bash
go mod graph | awk '{print $2}' | sed 's/@.*//' | sort | uniq -d
```

Modules appearing at more than one version — the ones MVS had to arbitrate.

## 4. Gotchas

- The graph is **pruned** at `go 1.17`+, so it is smaller than the historical
  full graph. See `Module Graph Pruning`.
- Output is large; always pipe it. A mid-sized service easily prints thousands of
  edges.
- Edges show *requirements*, not *selections*. Use `go list -m all` for what was
  actually chosen.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
