---
title: go work use
tags:
  - golang
  - basics
  - cli
  - toolchain
  - workspaces
parent: "[[Module Commands]]"
---

# `go work use`

Adds or removes module directories in `go.work`.

```bash
go work use ./api ./shared
go work use -r .            # recursively add every module found
go work use -r ./services
```

Removal is via `go work edit -dropuse=./old`.

## 1. What a `use` Line Means

```go
go 1.26.2

use (
    ./api
    ./shared
)
```

Those modules resolve to the **local directory**, not to a published version, for
every build run inside the workspace.

## 2. `-r` on a Monorepo

```bash
go work init
go work use -r .
```

Two commands to make a repository of twenty modules build against each other.
The historical alternative was twenty `replace` lines per `go.mod`.

## 3. Gotchas

- `use` paths are **relative to `go.work`** and are machine-specific — one more
  reason not to commit the file.
- A directory without a `go.mod` is rejected.
- Adding a module does not add its requirements to any `go.mod`; run
  `go work sync` for that.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
