---
title: Channel Architecture
tags:
  - golang
  - concurrency
parent: "[[Concurrency & Synchronization]]"
---

# Channel Architecture

Unbuffered vs buffered channels, hchan struct, sudog wait queues, and selectgo.

```text
Channel Architecture
│
├── [[Unbuffered Channels]]
├── [[Buffered Channels]]
├── [[Channel States & Behaviors]]
├── [[Channel Internals (hchan struct)]]
└── [[select Multiplexing]]
```

---

## 🗂️ Topics

- [[Unbuffered Channels]] — Synchronous rendezvous signaling with direct stack-to-stack copy optimization.
- [[Buffered Channels]] — Asynchronous FIFO ring buffer queues with bounded capacity.
- [[Channel States & Behaviors]] — Send, receive, and close semantics on nil, open, and closed channels.
- [[Channel Internals (hchan struct)]] — hchan fields: ring buffer, lock, sendq/recvq sudog wait queues.
- [[select Multiplexing]] — Non-blocking select with default, pseudo-random case evaluation, and selectgo() implementation.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
