---
title: Synchronization & Network Internals (chan.go, netpoll.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Synchronization & Network Internals (chan.go, netpoll.go)

makechan/chansend/chanrecv algorithms, selectgo algorithm, sudog nodes, Netpoller epoll/kqueue, and futex locks.

```text
Synchronization & Network Internals (chan.go, netpoll.go)
│
├── [[makechan, chansend, and chanrecv Implementation (chan.go)]]
├── [[closechan Implementation & Broadcast Signaling]]
├── [[selectgo Multi-Channel Select Algorithm (select.go)]]
├── [[sudog Synchronization Node Layout & Caching]]
├── [[Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)]]
├── [[Futex & OS Mutex Implementation (lock_futex.go)]]
└── [[runtime Package Diagnostic & Inspection APIs]]
```

---

## 🗂️ Topics

- [[makechan, chansend, and chanrecv Implementation (chan.go)]] — Complete lock acquisition, sudog enqueueing, ring-buffer copy, and direct stack transfer algorithms.
- [[closechan Implementation & Broadcast Signaling]] — Locking channel, releasing all waiting receivers with zero values, and panicking senders.
- [[selectgo Multi-Channel Select Algorithm (select.go)]] — Case shuffling (fastrand), lock ordering by channel address to prevent deadlocks, and polling.
- [[sudog Synchronization Node Layout & Caching]] — Synchronization waiting node representing parked goroutine on channel or sync primitive.
- [[Netpoller Implementation (netpoll_epoll.go, netpoll_kqueue.go)]] — Non-blocking I/O event polling, descriptor registration, and waking parked goroutines.
- [[Futex & OS Mutex Implementation (lock_futex.go)]] — Low-level user-space mutexes, CAS spin-wait, and falling back to Linux SYS_futex kernel sleep.
- [[runtime Package Diagnostic & Inspection APIs]] — runtime.ReadMemStats, runtime.Gosched, runtime.LockOSThread, runtime.GC, runtime.KeepAlive.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]
- 🎓 Root: [[Principal SWE]]
