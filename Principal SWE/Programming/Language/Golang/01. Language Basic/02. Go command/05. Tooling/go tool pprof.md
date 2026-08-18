---
title: pprof
tags:
  - golang
  - basics
  - cli
  - toolchain
  - profiling
  - performance
parent: "[[Tooling]]"
---

# `go tool pprof`

Interactive and web-based analysis of CPU, heap, block, mutex, and goroutine
profiles.

## 1. Getting a Profile

```bash
# from tests / benchmarks
go test -cpuprofile cpu.prof -memprofile mem.prof -bench . ./...

# from a live server with net/http/pprof imported
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
go tool pprof http://localhost:6060/debug/pprof/heap
```

## 2. Viewing

```bash
go tool pprof -http=:8080 cpu.prof     # flame graph, graph, source view
go tool pprof cpu.prof                 # interactive shell
```

Useful interactive commands: `top`, `top -cum`, `list <regexp>`, `web`, `peek`.

## 3. Reading `top`

```text
      flat  flat%   sum%        cum   cum%
     1.20s 40.00% 40.00%      1.20s 40.00%  runtime.memmove
     0.45s 15.00% 55.00%      2.10s 70.00%  encoding/json.Marshal
```

`flat` = time in that function itself. `cum` = including callees. Optimize by
`cum` to find the subsystem, then by `flat` to find the line.

## 4. Profile Types

| Endpoint | Answers |
|---|---|
| `profile` | Where is CPU time going |
| `heap` | What is allocated **now** (`-inuse_space`) or in total (`-alloc_space`) |
| `allocs` | Total allocations since start |
| `block`, `mutex` | Contention — needs `runtime.SetBlockProfileRate` |
| `goroutine` | Stack of every goroutine — the leak-hunting tool |

## 5. Gotchas

- Web views need `graphviz` installed.
- Exposing `net/http/pprof` on a public port leaks internals — bind it to
  localhost or an admin port.
- Comparing runs: `go tool pprof -base old.prof new.prof`.

---

## 🔗 References
- ⬆️ Parent: [[Tooling]]
