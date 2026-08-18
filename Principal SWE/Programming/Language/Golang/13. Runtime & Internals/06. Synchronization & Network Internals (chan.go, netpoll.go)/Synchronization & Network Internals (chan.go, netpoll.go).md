---
title: Synchronization & Network Internals (chan.go, netpoll.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Synchronization & Network Internals (chan.go, netpoll.go)

makechan/chansend/chanrecv algorithms, closechan, Netpoller epoll/kqueue, futex locks, and runtime APIs.

```text
Synchronization & Network Internals (chan.go, netpoll.go)
│
├── [[makechan, chansend, and chanrecv Implementation (chan.go)]]
├── [[closechan Implementation & Broadcast Signaling]]
├── [[Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)]]
├── [[Futex & OS Mutex Implementation (lock_futex.go)]]
└── [[runtime Package Diagnostic & Inspection APIs]]
```

---

## 🗂️ Topics

- [[makechan, chansend, and chanrecv Implementation (chan.go)]] — Complete lock acquisition, sudog enqueueing, ring-buffer copy, and direct stack transfer algorithms.
- [[closechan Implementation & Broadcast Signaling]] — Locking channel, releasing all waiting receivers with zero values, and panicking senders.
- [[Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)]] — Non-blocking I/O event polling, descriptor registration, and waking parked goroutines.
- [[Futex & OS Mutex Implementation (lock_futex.go)]] — Low-level user-space mutexes, CAS spin-wait, and falling back to Linux SYS_futex kernel sleep.
- [[runtime Package Diagnostic & Inspection APIs]] — runtime.ReadMemStats, runtime.Gosched, runtime.LockOSThread, runtime.GC, runtime.KeepAlive.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
