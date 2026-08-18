---
title: Channel Architecture
tags:
  - golang
  - channels
parent: "[[Concurrency & Synchronization]]"
---

# Channel Architecture

Unbuffered/buffered channels, hchan memory layout, sudog wait queues, and select multiplexing.

```text
Channel Architecture
│
├── [[Unbuffered Channels]]
├── [[Buffered Channels]]
├── [[Channel States & Behaviors]]
├── [[Channel Internals (hchan struct)]]
├── [[Channel Send and Receive Flow]]
├── [[select Multiplexing]]
└── [[Closing Channels Rules]]
```

---

## 🗂️ Topics

- [[Unbuffered Channels]] — Synchronous rendezvous signaling with direct stack-to-stack copy optimization.
- [[Buffered Channels]] — Asynchronous FIFO ring buffer queues with bounded capacity.
- [[Channel States & Behaviors]] — Send, receive, and close semantics on nil, open, and closed channels.
- [[Channel Internals (hchan struct)]] — hchan fields: ring buffer, mutex lock, sendq/recvq sudog wait queues.
- [[Channel Send and Receive Flow]] — Dissecting chansend and chanrecv runtime step-by-step mechanics.
- [[select Multiplexing]] — Non-blocking select with default, pseudo-random case evaluation, and selectgo() runtime.
- [[Closing Channels Rules]] — Only sender should close, closing closed panic, signaling with close(ch).

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
