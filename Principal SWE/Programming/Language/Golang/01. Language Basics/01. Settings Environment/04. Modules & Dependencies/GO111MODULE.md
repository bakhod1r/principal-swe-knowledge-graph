---
title: GO111MODULE
tags:
  - golang
  - basics
  - environment
  - modules
  - legacy
parent: "[[Modules & Dependencies]]"
---

# `GO111MODULE`

Selects module mode versus the pre-modules `GOPATH` mode.

## 1. Values

| Value | Behaviour |
|---|---|
| `on` | **Default since Go 1.16.** Always module mode; ignores `GOPATH/src` |
| `off` | Never module mode; imports resolve from `GOPATH/src`, `go.mod` ignored |
| `auto` | Module mode only inside a directory tree containing `go.mod` (pre-1.16 default) |

```bash
go env GO111MODULE
GO111MODULE=off go build ./...   # force legacy resolution
```

## 2. Why It Still Exists

Only for building ancient code that predates modules and was never migrated.
Every current workflow assumes `on`.

## 3. Symptoms of `off`

```text
cannot find package "github.com/x/y" in any of:
        /usr/local/go/src/github.com/x/y (from $GOROOT)
        /home/user/go/src/github.com/x/y (from $GOPATH)
```

A `$GOROOT`/`$GOPATH` search list in an error message means module mode is off —
`go.mod` is not being read at all.

## 4. Gotchas

- Setting `off` disables `go.sum` verification entirely: no `GOSUMDB`, no
  checksum guarantees. Never do this to "fix" a dependency error.
- `GO111MODULE=auto` is gone as a default but is still accepted.
- Inside `GOPATH``/src`, module mode still wins when `go.mod` exists.

---

## 🔗 References
- ⬆️ Parent: [[Modules & Dependencies]]
