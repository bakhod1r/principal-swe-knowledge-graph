---
title: GOMEMLIMIT
tags:
  - golang
  - basics
  - environment
  - runtime
  - gc
  - performance
parent: "[[Runtime Tuning]]"
---

# `GOMEMLIMIT`

**Go 1.19+.** A **soft** limit on the total memory the Go runtime may use. As the
limit approaches, the GC runs more aggressively instead of letting the heap grow.

## 1. Usage

```bash
GOMEMLIMIT=2GiB ./server
GOMEMLIMIT=1750MiB ./server
```

Accepts `B`, `KiB`, `MiB`, `GiB`, `TiB`. Default is `math.MaxInt64` (no limit).
At runtime: `debug.SetMemoryLimit(n)`.

## 2. Soft, Not Hard

The runtime will **exceed** the limit rather than fail an allocation. It is a
signal to the GC, not an allocator cap — the program is never killed by it.

```text
memory use rises → GC runs more often → CPU cost rises
                 → if it still cannot keep up, the limit is exceeded
```

That last step is intentional: thrashing the GC forever ("GC death spiral") is
worse than briefly overshooting.

## 3. The Container Recipe

```bash
# container limit 2Gi
GOMEMLIMIT=1800MiB
GOGC=off
```

Accounts for non-heap memory (stacks, the runtime itself, cgo allocations) with
the ~10% margin, and lets the limit alone drive collection. See `GOGC`.

## 4. Gotchas

- It covers memory **the Go runtime manages**. Memory allocated by cgo/`malloc`
  is invisible to it — see `CGO_ENABLED`.
- `GOGC=off` plus no `GOMEMLIMIT` disables collection entirely. Always set both,
  or neither.
- Setting it below the actual live heap makes the GC run continuously at ~100%
  CPU without ever succeeding.

---

## 🔗 References
- ⬆️ Parent: [[Runtime Tuning]]
