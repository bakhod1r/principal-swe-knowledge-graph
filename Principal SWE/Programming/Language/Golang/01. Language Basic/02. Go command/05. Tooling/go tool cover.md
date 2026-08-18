---
title: cover
tags:
  - golang
  - basics
  - cli
  - toolchain
  - testing
  - coverage
parent: "[[Tooling]]"
---

# `go tool cover`

Renders coverage profiles produced by `go test`.

## 1. The Loop

```bash
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out              # browser, line-by-line
go tool cover -html=cover.out -o cover.html
go tool cover -func=cover.out              # per-function percentages
```

## 2. `-func` Output

```text
github.com/me/api/user/user.go:21:   Create      85.7%
github.com/me/api/user/user.go:48:   Delete       0.0%
total:                               (statements) 72.4%
```

The `0.0%` lines are the useful output — a total percentage is a management
metric, uncovered functions are an engineering one.

## 3. Cross-Package Coverage

```bash
go test -coverpkg=./... -coverprofile=cover.out ./...
```

Without `-coverpkg`, each package only counts coverage of **itself**, so
integration tests appear to cover nothing.

## 4. Coverage Modes

| Mode | Records |
|---|---|
| `set` (default) | Was the statement executed |
| `count` | How many times |
| `atomic` | Same as `count`, safe under `-race` and parallelism |

```bash
go test -covermode=atomic -race -coverprofile=cover.out ./...
```

## 5. Gotchas

- Coverage measures **executed** statements, not asserted behaviour. 100% coverage
  with no assertions proves nothing.
- Binary-level coverage (integration tests over a built binary) uses `GOCOVERDIR`
  and `go tool covdata` — see `GOCOVERDIR`.

---

## 🔗 References
- ⬆️ Parent: [[Tooling]]
