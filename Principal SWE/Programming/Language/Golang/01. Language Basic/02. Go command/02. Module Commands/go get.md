---
title: Get Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
parent: "[[Module Commands]]"
---

# `go get`

Adds, upgrades, or removes dependencies in `go.mod`. Since Go 1.18 it **no
longer installs binaries** — use `go install` for that.

## 1. Usage

```bash
go get github.com/go-chi/chi/v5           # add / upgrade to latest
go get github.com/go-chi/chi/v5@v5.1.0    # pin exact version
go get github.com/go-chi/chi/v5@none      # remove the requirement
go get -u ./...                           # upgrade direct + indirect deps
go get -u=patch ./...                     # patch releases only
go get go@1.24.0                          # change the go directive
go get toolchain@go1.24.5                 # change GOTOOLCHAIN pin
```

## 2. Version Queries

| Query | Meaning |
|---|---|
| `@latest` | Highest release tag (skips pre-releases) |
| `@upgrade` | Latest, but never downgrade the current selection |
| `@patch` | Highest patch of the current minor |
| `@v1.2.3` | Exact version |
| `@<branch>` / `@<commit>` | Pseudo-version derived from VCS |
| `@none` | Remove |

## 3. Gotchas

- `go get` writes to `go.mod` and `go.sum` — it is the one command that ignores
  `-mod=readonly`.
- Upgrading with `-u` upgrades **test dependencies of dependencies** too. Prefer
  naming modules explicitly.
- A major version bump changes the import path (`/v2`) — `go get` will not do
  that rewrite for you. See `Semantic Versioning`.
- Private repositories need `GOPRIVATE`; see `Private Modules`.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
