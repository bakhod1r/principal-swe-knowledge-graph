---
title: Go Commands
tags:
  - golang
  - basics
  - cli
  - toolchain
parent: "[[Language Basic]]"
---

# 💻 Go Commands

The `go` command is a build system, package manager, test runner, and toolchain
driver in one binary. Everything below is `go <verb>`.

```text
go
│
├── `Build & Run`        build · run · install · generate · clean
├── `Module Commands`    mod · get · work
├── `Testing & Quality`  test · vet · fmt · fix
├── `Inspection`         env · list · doc · version
└── `Tooling`            tool · telemetry
```

---

## 🗂️ Categories

1. 🔨 **`Build & Run`** — `go build`, `go run`, `go install`, `go generate`, `go clean`.
2. 📦 **`Module Commands`** — `go mod`, `go get`, `go work`. Concepts in `Dependencies`.
3. 🧪 **`Testing & Quality`** — `go test`, `go vet`, `go fmt`, `go fix`.
4. 🔍 **`Inspection`** — `go env`, `go list`, `go doc`, `go version`.
5. 🧰 **`Tooling`** — `go tool`, `go telemetry`.

---

## ⚡ Daily Driver

```bash
go mod tidy                    # dependencies honest
gofmt -l .                     # formatting check
go vet ./...                   # correctness lint
go test -race -count=1 ./...   # tests, no cache, race detector
go build -trimpath -o bin/app ./cmd/app
```

---

## 🧭 Mental Model

```text
go.mod  ──`go get` / `go mod`──►  build list (MVS)
                                          │
                                          ▼
                              GOMODCACHE + GOCACHE
                                          │
                    ┌─────────────────────┼─────────────────────┐
                    ▼                     ▼                     ▼
              `go build`           `go test`           `go install`
                    │                     │                     │
                 binary              test results            GOBIN
```

Every command above reads `Settings Environment` for its configuration and
`GOFLAGS` for implicit default flags.

---

## 📋 Full Verb Reference

| Verb | Category | One-line |
|---|---|---|
| `build` | Build & Run | Compile packages and dependencies |
| `clean` | Build & Run | Remove artifacts and caches |
| `doc` | Inspection | Show documentation |
| `env` | Inspection | Print / set toolchain configuration |
| `fix` | Testing & Quality | Rewrite deprecated API usage |
| `fmt` | Testing & Quality | Run `gofmt -l -w` |
| `generate` | Build & Run | Run `//go:generate` directives |
| `get` | Module | Change dependency versions |
| `install` | Build & Run | Compile and install a binary |
| `list` | Inspection | Query packages and modules |
| `mod` | Module | Module maintenance |
| `run` | Build & Run | Compile and run |
| `telemetry` | Tooling | Configure usage telemetry |
| `test` | Testing & Quality | Run tests and benchmarks |
| `tool` | Tooling | Run a toolchain or pinned tool |
| `version` | Inspection | Toolchain version; binary build info |
| `vet` | Testing & Quality | Static analysis |
| `work` | Module | Multi-module workspaces |

---

## 🔗 References
- ⬆️ Parent: `Language Basic`
