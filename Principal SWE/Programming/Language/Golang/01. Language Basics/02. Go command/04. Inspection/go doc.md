---
title: Doc Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - documentation
parent: "[[Inspection]]"
---

# `go doc`

Prints documentation for a package, symbol, or method — offline, from source.

## 1. Usage

```bash
go doc                         # current package summary
go doc fmt                     # package docs
go doc fmt.Printf              # one symbol
go doc -all net/http           # everything, including unexported headings
go doc -src sync.Once.Do       # show the source
go doc -u ./internal/user      # include unexported identifiers
```

## 2. Resolution Order

```text
symbol
  ↓
current package  →  module dependencies  →  standard library
```

Third-party packages must already be in `GOMODCACHE`.

## 3. Gotchas

- `go doc` reads the module cache, not the network — `go mod download` first if
  a package is missing. See `go mod`.
- `godoc` (the web server) is a separate tool, now `pkg.go.dev` or `gopls`.
- Doc comments must sit immediately above the declaration with no blank line.

---

## 🔗 References
- ⬆️ Parent: [[Inspection]]
