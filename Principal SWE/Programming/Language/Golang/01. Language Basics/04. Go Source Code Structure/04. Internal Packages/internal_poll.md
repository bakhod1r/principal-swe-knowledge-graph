---
title: internal/poll
tags:
  - golang
  - goroot
  - runtime
  - io
  - network
parent: "[[Internal Packages]]"
---

# `internal/poll`

The bridge between blocking Go I/O calls and the non-blocking OS event loop.
Everything in `net` and file I/O in `os` funnels through here.

## 1. Position in the Stack

```text
net.Conn.Read / os.File.Read
        │
        ▼
internal/poll.FD.Read
        │
        ├── syscall.Read → EAGAIN?
        │        │
        │        └── runtime_pollWait → park the goroutine
        ▼
runtime netpoll (epoll / kqueue / IOCP)  ← `proc (Scheduler)`
        │
        └── readiness → goready(g) → goroutine resumes
```

## 2. What It Buys

A blocking-looking API over a non-blocking kernel interface. `conn.Read` blocks
**the goroutine**, not the OS thread, so a server can hold hundreds of thousands
of connections on a handful of threads.

## 3. Also Responsible For

- Deadline handling — `SetReadDeadline` timers that wake parked goroutines
- The `fdMutex` reference counting that makes `Close` safe during a concurrent
  `Read`
- `sendfile` / `splice` fast paths behind `io.Copy`

## 4. Gotchas

- **Unimportable** — see `internal visibility rule`.
- Files opened outside the poller (regular disk files on Linux) genuinely block
  an OS thread; only sockets and pipes are pollable. This is why heavy disk I/O
  inflates the thread count.
- `os.File` obtained via `os.NewFile` on a non-blocking fd behaves differently
  from one from `os.Open`.

---

## 🔗 References
- ⬆️ Parent: [[Internal Packages]]
