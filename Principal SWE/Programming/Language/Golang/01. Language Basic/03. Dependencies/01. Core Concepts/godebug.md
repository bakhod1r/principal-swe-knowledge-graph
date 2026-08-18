---
title: godebug directive
tags:
  - golang
  - basics
  - dependencies
  - modules
  - go.mod
  - compatibility
parent: "[[Core Concepts]]"
---

# `godebug`

**Go 1.23+.** Sets `GODEBUG` defaults for the whole main module, in `go.mod`
or per-file with `//go:debug`.

## 1. Syntax

```go
go 1.24.0

godebug (
    httplaxcontentlength=1
    tls10server=1
)
```

Single form: `godebug default=go1.21`.

## 2. Why It Exists

Go's compatibility promise sometimes requires changing behaviour for security.
Each such change gets a `GODEBUG` key and keeps the **old** behaviour as default
for modules declaring an older `go` version. This directive lets you opt back to
old behaviour **without** downgrading the whole language version.

```text
go 1.24            → new (secure) defaults everywhere
godebug default=go1.21  → 1.21-era defaults, but 1.24 language features
```

## 3. Per-File Form

```go
//go:debug httplaxcontentlength=1

package main
```

Only allowed in `main` packages — a library cannot change its consumers' runtime
behaviour.

## 4. Gotchas

- A **temporary migration tool**, not a config file. Each key has a removal
  timeline announced in the release notes.
- The environment variable `GODEBUG` overrides the directive at run time.
- `go version -m ./bin/app` shows the baked-in settings — see `go version`.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
