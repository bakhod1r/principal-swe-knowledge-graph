---
title: GOGC
tags:
  - golang
  - basics
  - environment
  - runtime
  - gc
  - performance
parent: "[[Runtime Tuning]]"
---

# `GOGC`

Sets the garbage collector's target heap growth, as a percentage of the live heap
retained after the previous cycle.

## 1. The Formula

```text
next GC target = live heap × (1 + GOGC/100)
```

| `GOGC` | Live heap 100 MB | Effect |
|---|---|---|
| `50` | GC at 150 MB | Less memory, more CPU in GC |
| `100` (default) | GC at 200 MB | Balanced |
| `400` | GC at 500 MB | More memory, less GC CPU |
| `off` | never (until `GOMEMLIMIT`) | Memory grows unbounded |

## 2. Usage

```bash
GOGC=400 ./server        # trade RAM for throughput
GOGC=off GOMEMLIMIT=2GiB ./server
```

At runtime: `debug.SetGCPercent(400)` in `runtime/debug`.

## 3. When to Change It

Raise it when GC CPU time is measurably high and RAM is spare — check with:

```bash
GODEBUG=gctrace=1 ./server
```

Each line reports the GC's CPU share. Below ~5% there is nothing to win.

## 4. Gotchas

- `GOGC=off` **without** `GOMEMLIMIT` means the heap only ever grows. In a
  container that ends in an OOM kill with no Go-side warning.
- It scales the *live* heap, so a program with a small live set and high
  allocation rate collects very often regardless of `GOGC`; fix the allocations.
- See `mgc (Garbage Collector)` for the collector itself.

---

## 🔗 References
- ⬆️ Parent: [[Runtime Tuning]]
