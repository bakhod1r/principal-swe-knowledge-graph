---
title: GORACE
tags:
  - golang
  - basics
  - environment
  - runtime
  - testing
  - concurrency
parent: "[[Runtime Tuning]]"
---

# `GORACE`

Options for binaries built with `-race`. Ignored by non-race builds.

## 1. Syntax

Space-separated `key=value` pairs:

```bash
GORACE="halt_on_error=1 log_path=/tmp/race" go test -race ./...
```

## 2. Options

| Option | Default | Meaning |
|---|---|---|
| `halt_on_error` | `0` | Exit on the **first** race instead of continuing |
| `log_path` | stderr | Write reports to `<path>.<pid>` |
| `history_size` | `1` | Per-goroutine memory-access history; `n` → `32K × 2^n` entries |
| `atexit_sleep_ms` | `1000` | Wait at exit so background goroutines can report |
| `strip_path_prefix` | — | Trim a prefix from paths in reports |
| `exitcode` | `66` | Process exit code when a race is found |

## 3. Practical Settings

```bash
# CI: fail fast, fail loud
GORACE="halt_on_error=1" go test -race ./...

# Deep debugging: longer history for a race with a truncated stack
GORACE="history_size=7" go test -race -run TestFlaky ./...
```

`history_size=7` costs substantial memory — use it on one test, not a whole suite.

## 4. Gotchas

- **Without `halt_on_error=1`, `go test -race` still fails the run** on a detected
  race; the option only controls whether it stops at the first one.
- The race detector finds races that **actually happened** during the run. A clean
  run is not proof of correctness.
- `-race` costs ~2–10× CPU and ~5–10× memory. Never ship a race binary.
- See `GOTRACEBACK` for what gets printed when a race-built binary crashes.

---

## 🔗 References
- ⬆️ Parent: [[Runtime Tuning]]
