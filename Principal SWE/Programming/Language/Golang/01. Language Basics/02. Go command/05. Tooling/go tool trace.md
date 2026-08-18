---
title: trace
tags:
  - golang
  - basics
  - cli
  - toolchain
  - profiling
  - concurrency
parent: "[[Tooling]]"
---

# `go tool trace`

Renders an execution trace: every goroutine, scheduler event, GC cycle, and
syscall on a timeline. Answers *latency* questions that `go tool pprof` cannot.

## 1. Capturing

```bash
go test -trace trace.out ./...
go tool pprof   # (unrelated — profiles, not traces)
curl -o trace.out 'http://localhost:6060/debug/pprof/trace?seconds=5'
```

In code:

```go
f, _ := os.Create("trace.out")
trace.Start(f)
defer trace.Stop()
```

## 2. Viewing

```bash
go tool trace trace.out       # opens a browser UI
```

Views: **Goroutine analysis**, **Network blocking profile**, **Synchronization
blocking profile**, **Syscall blocking profile**, **Scheduler latency profile**.

## 3. What It Uniquely Shows

| Question | Where |
|---|---|
| Why is p99 latency bad when CPU is idle? | Scheduler latency profile |
| Is the GC pausing us? | Timeline, GC rows |
| Are goroutines blocked on a mutex or on the network? | Blocking profiles |
| Is one `P` starved while others idle? | Timeline per-proc rows |

CPU profiles show *where* time goes; traces show *why nothing was running*.

## 4. Gotchas

- Tracing has real overhead — capture seconds, not minutes. Files grow fast.
- The UI needs a Chromium-based browser for the timeline view.
- Go 1.22+ uses a lower-overhead tracer; older traces are not interchangeable.
- See `proc (Scheduler)` for the `G`/`M`/`P` entities the timeline displays.

---

## 🔗 References
- ⬆️ Parent: [[Tooling]]
