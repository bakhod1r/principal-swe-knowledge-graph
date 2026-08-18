---
title: GOMAXPROCS
tags:
  - golang
  - basics
  - environment
  - runtime
  - scheduler
  - performance
parent: "[[Runtime Tuning]]"
---

# `GOMAXPROCS`

The maximum number of OS threads that may execute Go code simultaneously — the
number of `P`s in the scheduler. See `proc (Scheduler)`.

## 1. Default

| Go version | Default |
|---|---|
| ≤ 1.24 | `runtime.NumCPU()` — the **host** CPU count |
| 1.25+ | Container-aware: derived from the cgroup CPU limit when lower |

```bash
GOMAXPROCS=4 ./server
```

At runtime: `runtime.GOMAXPROCS(4)`; read the current value with
`runtime.GOMAXPROCS(0)`.

## 2. The Classic Container Bug (pre-1.25)

```text
Kubernetes limit: cpu: "2"     → cgroup quota = 2 cores
Host machine:     64 cores
Go ≤1.24:         GOMAXPROCS = 64
```

64 `P`s competing for 2 cores' worth of quota causes heavy context switching and
CFS throttling — latency spikes with idle-looking CPU graphs. The historic fix is
the `automaxprocs` library; Go 1.25 makes it unnecessary.

## 3. It Does Not Limit Goroutines or Threads

```text
GOMAXPROCS   = goroutines running Go code at once
threads       = unbounded; blocking syscalls spawn more
goroutines    = unbounded; millions are normal
```

A blocking syscall detaches its `M` from the `P`, so thread count regularly
exceeds `GOMAXPROCS`.

## 4. Gotchas

- Setting it to `1` does **not** make code race-free; goroutines still interleave
  at every preemption point.
- Lowering it can help latency under CFS quota, but hurts throughput on
  parallel workloads — measure, do not guess.
- `runtime.NumCPU()` still reports the host count; only the default changed.

---

## 🔗 References
- ⬆️ Parent: [[Runtime Tuning]]
