---
title: Work Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - workspaces
parent: "[[Module Commands]]"
---

# `go work`

Multi-module workspaces (Go 1.18+). Lets several local modules resolve against
each other without `replace` edits. See `go.work` and `GOWORK`.

## 1. Subcommands

```bash
go work init ./api ./shared    # create go.work listing both modules
go work use ./worker           # add another module
go work use -r .               # add every module found recursively
go work edit -dropuse ./old
go work sync                   # push workspace build list back into each go.mod
go work vendor                 # vendor the workspace (Go 1.22+)
```

## 2. Why Not `replace`

```text
replace         → committed into go.mod, affects everyone
go.work         → local file, usually gitignored, affects only you
```

## 3. Gotchas

- `go.work` overrides `replace` directives in member modules.
- Do not commit `go.work` for a library; it changes how consumers' builds resolve
  nothing, but it silently changes yours and hides broken `go.mod` files.
- `go work sync` does not add requirements to `go.mod` that only exist because of
  the workspace — verify with `go mod` `tidy` outside the workspace.
- Disable temporarily with `GOWORK=off`.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
