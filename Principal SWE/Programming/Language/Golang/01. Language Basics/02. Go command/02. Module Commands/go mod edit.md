---
title: go mod edit
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
  - scripting
parent: "[[Module Commands]]"
---

# `go mod edit`

Programmatic, non-interactive edits to `go.mod`. Built for scripts and CI, not
for daily use.

```bash
go mod edit -require=github.com/x/y@v1.2.3
go mod edit -droprequire=github.com/x/y
go mod edit -replace=github.com/x/y=../local/y
go mod edit -dropreplace=github.com/x/y
go mod edit -exclude=github.com/x/y@v1.4.2
go mod edit -retract=v1.0.3
go mod edit -go=1.24.0 -toolchain=go1.24.5
go mod edit -json                # parse go.mod as JSON
go mod edit -fmt                 # reformat only
```

## 1. It Does Not Resolve

`-require` writes the line and stops. It does **not** download, does not update
`go.sum`, and does not run `MVS`. That is the difference from `go get`:

| | `go mod edit -require` | `go get` |
|---|---|---|
| Network | No | Yes |
| `go.sum` updated | No | Yes |
| Transitive deps resolved | No | Yes |

## 2. Scripting Use

```bash
# point every internal dep at a local checkout, for a monorepo CI job
for m in $(go list -m -f '{{.Path}}' all | grep '^git.corp/'); do
  go mod edit -replace="$m=../$(basename "$m")"
done
```

For interactive multi-module work, `go work` is the better tool.

## 3. Gotchas

- Leaves `go.mod` in a state that does not build until you run `go mod tidy`
  or `go mod download`.
- `-json` is the supported way to read `go.mod`; parsing it with `grep` breaks on
  block syntax and comments.
- `-replace` here is identical in effect to writing `replace` by hand.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
