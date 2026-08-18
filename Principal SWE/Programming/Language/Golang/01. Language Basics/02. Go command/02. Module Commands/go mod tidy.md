---
title: go mod tidy
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
parent: "[[Module Commands]]"
---

# `go mod tidy`

Makes `go.mod` and `go.sum` exactly match the packages the module actually
imports — adding what is missing, removing what is unused.

```bash
go mod tidy
go mod tidy -v            # list removed modules
go mod tidy -go=1.24      # also set the go directive
go mod tidy -diff         # Go 1.23+: show changes, exit non-zero, write nothing
```

## 1. What It Considers

Imports of **all** packages in the module, including tests, and including tests
of dependencies at the pruning depth set by the `go directive`. This is why it
sometimes adds a module you never import directly — see
`Module Graph Pruning`.

## 2. In CI

```bash
go mod tidy -diff     # fails if go.mod/go.sum are stale
```

Before Go 1.23 the idiom was `go mod tidy && git diff --exit-code`.

## 3. Gotchas

- Needs **network access** unless every module is already in `GOMODCACHE`.
- Removing a requirement that only a build-tagged file imports is correct
  behaviour, and breaks that build. `go mod tidy` reads all tags it knows, but not
  arbitrary custom ones — verify with `go build -tags=...`.
- It respects the `go` directive: lowering it re-adds requirements needed by
  older consumers.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
