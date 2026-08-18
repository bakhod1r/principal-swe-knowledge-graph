---
title: go mod init
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
parent: "[[Module Commands]]"
---

# `go mod init`

Creates a new `go.mod` in the current directory.

```bash
go mod init github.com/me/api
go mod init                    # infers the path from VCS remote, if possible
```

## 1. What It Writes

```go
module github.com/me/api

go 1.26.2
```

Only the `module` and `go directive` lines. Requirements come later, from
`go get` or `go mod tidy`.

## 2. Choosing the Path

The path must be the location the module is fetched from — see `module`. For
code that will never be published, any unique prefix works, but a real path costs
nothing and avoids a rename later.

## 3. Gotchas

- Fails if `go.mod` already exists — it will not overwrite.
- Inference only works when a Git remote is configured; otherwise the path is
  required as an argument.
- Running it inside an existing module creates a **nested module**, which the
  parent then excludes from its build. Usually a mistake.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
