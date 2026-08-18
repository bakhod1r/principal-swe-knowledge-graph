---
title: Runtime Tuning
tags:
  - golang
  - basics
  - environment
  - runtime
  - performance
parent: "[[Settings Environment]]"
---

# 🚀 Runtime Tuning

Variables read by the **running program**, not by the `go` command. None of them
appear in `go env` output — they are consumed by the Go runtime at process start.

```text
go env  ──────────► toolchain configuration   (build time)

process environment ──► runtime configuration (run time)
        │
        ├── [[GOGC]]        — heap growth target
        ├── [[GOMEMLIMIT]]  — soft memory ceiling
        ├── [[GOMAXPROCS]]  — OS threads running Go code
        ├── [[GORACE]]      — race detector options
        └── `GODEBUG`     — runtime diagnostics and compatibility switches
```

## 🗂️ Variables

- **[[GOGC]]** — percentage of live heap growth that triggers a GC cycle. Default `100`.
- **[[GOMEMLIMIT]]** — soft limit on total runtime memory; makes the GC work harder
  as the limit approaches. Go 1.19+.
- **[[GOMAXPROCS]]** — how many OS threads may execute Go code simultaneously.
  Container-aware since Go 1.25.
- **[[GORACE]]** — options for binaries built with `-race`.
- **`GODEBUG`** — documented in `Cache & Testing`; also a runtime knob.

## ⚙️ The Standard Container Pair

```bash
GOMEMLIMIT=$((MEM_LIMIT_BYTES * 9 / 10))
GOGC=off        # let GOMEMLIMIT alone drive collection
```

Set the limit ~10% below the container's hard limit so the GC reacts before the
kernel OOM-killer does. See [[GOMEMLIMIT]] for why `GOGC=off` alone is dangerous.

---

## 🔗 References
- ⬆️ Parent: [[Settings Environment]]
