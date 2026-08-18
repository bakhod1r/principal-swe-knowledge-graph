---
title: Mod Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - modules
parent: "[[Module Commands]]"
---

# `go mod`

Module maintenance. See `Dependencies` for the concepts.

## 1. Subcommands

| Subcommand | Purpose |
|---|---|
| `go mod init <path>` | Create `go.mod` with the given module path |
| `go mod tidy` | Add missing / drop unused requirements, refresh `go.sum` |
| `go mod download` | Fetch modules into `GOMODCACHE` without building |
| `go mod verify` | Check cached modules against their recorded hashes |
| `go mod graph` | Print the full module dependency graph |
| `go mod why <pkg>` | Explain the shortest import path to a package |
| `go mod edit` | Programmatic edits to `go.mod` (`-require`, `-replace`, `-droprequire`, `-json`) |
| `go mod vendor` | Copy dependencies into `vendor/` |

## 2. Everyday Sequence

```bash
go mod init github.com/me/api
go get github.com/go-chi/chi/v5
go mod tidy
go mod verify
```

## 3. Debugging a Dependency

```bash
go mod why -m golang.org/x/text     # who pulls it in
go mod graph | grep 'x/text'        # every edge that reaches it
go mod edit -replace=old=new@v1.2.3 # local override — see `replace`
```

## 4. Gotchas

- `go mod tidy` respects the `go` directive: it keeps requirements needed by the
  **oldest** supported version, which is why it sometimes adds things back.
- `-mod=vendor` becomes automatic when a `vendor/` directory exists and the `go`
  directive is ≥1.14.
- `go mod tidy` needs network access unless everything is already in
  `GOMODCACHE`; `GOFLAGS=-mod=mod` does not change that.

---

## 🔗 References
- ⬆️ Parent: [[Module Commands]]
