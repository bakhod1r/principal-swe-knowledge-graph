---
title: tool directive
tags:
  - golang
  - basics
  - dependencies
  - modules
  - go.mod
  - tooling
parent: "[[Core Concepts]]"
---

# `tool`

**Go 1.24+.** Records executable tools the module depends on, so their versions
are pinned in `go.mod` like any other dependency.

## 1. Syntax

```go
tool (
    golang.org/x/tools/cmd/stringer
    github.com/golangci/golangci-lint/cmd/golangci-lint
)
```

## 2. Usage

```bash
go get -tool golang.org/x/tools/cmd/stringer   # add
go tool stringer -type=Pill                    # run the pinned version
go tool                                        # list module tools too
go get -tool golang.org/x/tools/cmd/stringer@none   # remove
```

## 3. What It Replaced

The old `tools.go` hack:

```go
//go:build tools

package tools

import _ "golang.org/x/tools/cmd/stringer"   // blank import to pin the version
```

A fake package with a build tag, existing only to make `go mod tidy` keep the
requirement. The `tool` directive makes this first-class.

## 4. Gotchas

- Tools become **real module requirements** — their dependencies enter your
  module graph and can affect `MVS` for your library consumers.
- `go tool <name>` builds on first use; the binary is cached in `GOCACHE`, not
  installed to `GOBIN`.
- For tools with no version constraint, `go install` `pkg@version` is still the
  lighter option — it touches nothing in `go.mod`.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
