---
title: Tool Command
tags:
  - golang
  - basics
  - cli
  - toolchain
  - toolchain
parent: "[[Tooling]]"
---

# `go tool`

Runs the binaries in `GOTOOLDIR`, plus — since Go 1.24 — the tools your module
pins in `go.mod`.

## 1. Listing

```bash
go tool                      # tools the toolchain advertises
go tool -n pprof             # print the path instead of running
```

The bare listing is deliberately short (`asm cgo compile cover fix link
preprofile vet` on Go 1.26). Analysis tools such as `pprof`, `trace`, `nm`,
`objdump`, `addr2line`, `covdata`, `test2json`, `pack`, `buildid`, and `dist`
are **not listed but still run** — probe with `go tool -n <name>`.

## 2. Analysis Tools

```bash
go tool cover -html=cover.out
go tool pprof -http=:8080 cpu.prof
go tool trace trace.out
go tool objdump -s 'main\.hot' ./bin/api
go tool nm ./bin/api | sort -k2
```

## 3. Compiler Internals

`compile`, `link`, `asm`, `cgo`, `dist`, `buildid`, `addr2line`, `pack`.
These are the pieces `go build` orchestrates — see `cmd_compile`, `cmd_link`,
`cmd_asm`, `Toolchain & Compiler`.

```bash
go tool compile -S main.go       # assembly listing
go build -gcflags='-m -m' ./...  # escape analysis decisions
```

## 4. Module Tool Dependencies (Go 1.24+)

```bash
go get -tool golang.org/x/tools/cmd/stringer
go tool stringer -type=Pill      # version-pinned by go.mod
go tool                          # now lists it too
```

Replaces the old `tools.go` + blank-import trick. See `go install`.

## 5. Gotchas

- Tools under `go tool` are unversioned toolchain internals; their flags can
  change between Go releases without a compatibility promise.
- `go tool pprof` needs `graphviz` for the SVG/web views.

---

## 🔗 References
- ⬆️ Parent: [[Tooling]]
