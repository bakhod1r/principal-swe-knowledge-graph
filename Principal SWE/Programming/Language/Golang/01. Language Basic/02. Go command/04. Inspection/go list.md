---
title: List Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - introspection
parent: "[[Inspection]]"
---

# `go list`

The toolchain's query interface. Everything the `go` command knows about packages
and modules is reachable from here.

## 1. Packages

```bash
go list ./...
go list -json ./cmd/api
go list -deps ./cmd/api                       # full transitive package list
go list -f '{{.ImportPath}} {{.Doc}}' ./...
go list -f '{{join .Imports "\n"}}' ./internal/user
```

## 2. Modules

```bash
go list -m                    # this module
go list -m all                # full build list (see `MVS`)
go list -m -u all             # mark available upgrades
go list -m -json all
go list -m -versions github.com/go-chi/chi/v5
```

## 3. Practical Recipes

```bash
# every package that is not a test
go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}' ./...

# direct dependencies only
go list -m -f '{{if not .Indirect}}{{.Path}} {{.Version}}{{end}}' all

# find modules with a newer release
go list -m -u -f '{{if .Update}}{{.Path}} {{.Version}} → {{.Update.Version}}{{end}}' all
```

## 4. Gotchas

- `-m all` in a module means the **build list**, not "everything imported" — it
  includes modules providing no imported package.
- Template fields come from `cmd/go`'s `Package` and `Module` structs; `-json`
  first, then narrow with `-f`.
- `go list ./...` fails the whole run on one broken package unless you add `-e`.

---

## 🔗 References
- ⬆️ Parent: [[Inspection]]
