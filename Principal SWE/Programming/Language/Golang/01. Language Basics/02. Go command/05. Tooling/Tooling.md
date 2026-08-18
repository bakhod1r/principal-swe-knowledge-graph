---
title: Tooling Commands
tags:
  - golang
  - basics
  - cli
  - toolchain
parent: "[[Go Commands]]"
---

# 🧰 Tooling Commands

```text
[[go tool]]       → toolchain internals (compile, link, pprof, cover, trace)
                    + module-pinned tools via the go.mod `tool` directive
[[go telemetry]]  → local / on / off
```

## 🗂️ Commands

- **[[go tool]]** — profiling, coverage rendering, disassembly, and the compiler
  and linker themselves. See `Toolchain & Compiler`.
- **[[go telemetry]]** — opt-in usage counters; defaults to `local`.

## 📊 Profiling Loop

```bash
go test -cpuprofile cpu.prof -bench . ./...
go tool pprof -http=:8080 cpu.prof
```

---

## 🔗 References
- ⬆️ Parent: `Go Commands`
