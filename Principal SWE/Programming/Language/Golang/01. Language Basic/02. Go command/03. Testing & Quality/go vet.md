---
title: Vet Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - static-analysis
parent: "[[Testing & Quality]]"
---

# `go vet`

Static analysis for mistakes the compiler accepts. A subset runs automatically
during `go test`.

## 1. Usage

```bash
go vet ./...
go vet -printf=false ./...           # disable one analyzer
go vet -vettool=$(which shadow) ./... # run an external analyzer
```

## 2. High-Value Analyzers

| Analyzer | Catches |
|---|---|
| `printf` | Format string / argument mismatches |
| `copylocks` | Copying a `sync.Mutex` by value |
| `lostcancel` | `context.WithCancel` whose cancel is never called |
| `loopclosure` | Capturing a loop variable (pre-Go 1.22 semantics) |
| `httpresponse` | `defer resp.Body.Close()` before the error check |
| `nilfunc`, `unusedresult` | Comparisons and calls with no effect |
| `atomic` | Misuse of `sync/atomic` assignment |

## 3. In CI

```bash
go vet ./... || exit 1
```

`go vet` exits non-zero on any finding, so it needs no wrapper.

## 4. Gotchas

- `go test` runs only a safe subset (`atomic`, `bool`, `buildtags`, `errorsas`,
  `ifaceassert`, `nilfunc`, `printf`, `stringintconv`). Full `go vet` finds more.
- It needs the package to type-check; a compile error masks all vet output.
- `go vet` is not a linter. Style belongs to `golangci-lint`; correctness belongs here.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Quality]]
