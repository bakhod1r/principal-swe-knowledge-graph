---
title: Vendoring
tags:
  - golang
  - basics
  - dependencies
  - modules
  - vendor
parent: "[[Core Concepts]]"
---

# Vendoring

Copies every dependency's source into a `vendor/` directory inside the module, so
builds need no network and no `GOMODCACHE`.

## 1. Usage

```bash
go mod vendor          # create/refresh vendor/
go build ./...         # automatically uses vendor/ — see below
go mod vendor -e       # keep going despite errors
```

## 2. Automatic Activation

When `vendor/modules.txt` exists **and** the `go directive` is ≥ 1.14, the go
command defaults to `-mod=vendor`. No flag needed; also no network access.

```bash
go build -mod=mod ./...   # explicitly ignore vendor/
```

## 3. `vendor/modules.txt`

```text
# github.com/go-chi/chi/v5 v5.1.0
## explicit; go 1.20
github.com/go-chi/chi/v5
github.com/go-chi/chi/v5/middleware
```

Records versions and which packages are actually used. The go command
**verifies** it against `go.mod` and fails on drift — vendoring is not a way to
smuggle in modified dependencies unnoticed.

## 4. Trade-offs

| | Vendored | Not vendored |
|---|---|---|
| Build without network | ✅ | ❌ (needs cache or proxy) |
| Repository size | Large — diffs full of third-party code | Small |
| `go.sum` verification | Skipped at build time | Enforced |
| Reviewing dependency changes | Visible in the diff | Requires reading `go.mod` |

## 5. Gotchas

- `vendor/` **skips checksum verification** at build time; integrity moves to
  code review. Run `go mod verify` before committing a refresh.
- Only packages actually imported are vendored — reflection-based or
  `go:generate` tooling can break.
- Forgetting `go mod vendor` after `go get` produces a confusing
  "inconsistent vendoring" error, which is the check working correctly.

---

## 🔗 References
- ⬆️ Parent: `01. Core Concepts`
