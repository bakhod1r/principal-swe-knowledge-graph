---
title: Test Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - testing
parent: "[[Testing & Quality]]"
---

# `go test`

Builds and runs tests, benchmarks, examples, and fuzz targets in `_test.go` files.

## 1. Usage

```bash
go test ./...
go test -run 'TestUser/create' ./internal/user
go test -v -race -count=1 ./...
go test -cover -coverprofile=cover.out ./... && go tool cover -html=cover.out
```

## 2. Key Flags

| Flag | Meaning |
|---|---|
| `-run <regexp>` | Select tests; `/` separates subtest levels |
| `-v` | Stream test logs |
| `-race` | Race detector — the single highest-value flag in Go |
| `-count=1` | Bypass the test result cache |
| `-timeout 30s` | Panic the whole binary after this long (default 10m) |
| `-short` | Sets `testing.Short()` — used to skip slow tests |
| `-parallel N` | Max concurrent `t.Parallel()` tests (default `GOMAXPROCS`) |
| `-shuffle=on` | Randomize test order — exposes inter-test coupling |
| `-json` | Machine-readable output for CI |
| `-bench . -benchmem` | Run benchmarks, report allocations |
| `-fuzz Fuzz<Name>` | Run a fuzz target (one package at a time) |
| `-coverpkg=./...` | Attribute coverage across package boundaries |

## 3. Caching

Results are cached in `GOCACHE` keyed by the test binary, its inputs, and
environment. A cached run prints `(cached)`. Clear with
`go clean -testcache` — see `go clean`.

## 4. Gotchas

- `-race` changes timing; a test that only passes without it is already broken.
- `-run` matches by regexp, unanchored: `-run TestUser` also runs `TestUserAdmin`.
  Anchor with `-run '^TestUser$'`.
- A `-timeout` panic dumps every goroutine — that stack dump is the debugging
  payload, not noise. See `GOTRACEBACK`.
- Fuzzing corpora live in `testdata/fuzz` and in the cache; see `go clean` `-fuzzcache`.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Quality]]
