---
title: Channel Architecture & Operations
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Channel Architecture & Operations

Channel operations, hchan memory layout, sudog wait queues, direct stack copy optimization, and selectgo.

```text
Channel Architecture & Operations
│
├── [[Channel Send Operation (ch <- v)]]
├── [[Channel Receive Operation (v <- ch & v, ok <- ch)]]
├── [[Channel Close Operation (close(ch))]]
├── [[Channel Length & Capacity (len, cap)]]
├── [[Unbuffered Channel Synchronous Rendezvous]]
├── [[Buffered Channel Ring Buffer Pointer Math]]
├── [[Nil Channel Blocking & Deadlock Patterns]]
├── [[Closed Channel Read-Write Behavior Matrix]]
├── [[Channel Memory Layout (hchan Struct)]]
├── [[sudog Wait Queue Architecture]]
├── [[Direct Stack-to-Stack Copy Optimization]]
└── [[selectgo Runtime Multiplexing Algorithm]]
```

---

## 🗂️ Topics

- [[Channel Send Operation (ch <- v)]] — Synchronous vs asynchronous send execution, locking, and waking receivers.
- [[Channel Receive Operation (v <- ch & v, ok <- ch)]] — Reading from channels, comma-ok idiom, and unblocking senders.
- [[Channel Close Operation (close(ch))]] — Closing rules, panic conditions (closing closed or nil channels), and broadcast signaling.
- [[Channel Length & Capacity (len, cap)]] — Inspecting buffer length and total capacity with len() and cap().
- [[Unbuffered Channel Synchronous Rendezvous]] — Direct synchronous rendezvous signaling without intermediate buffer storage.
- [[Buffered Channel Ring Buffer Pointer Math]] — Circular ring buffer pointer math (sendx, recvx, qcount) in hchan.
- [[Nil Channel Blocking & Deadlock Patterns]] — Why reads and writes to nil channels block forever and how to use them in select.
- [[Closed Channel Read-Write Behavior Matrix]] — Matrix of reading zero values from closed channels vs panicking on writes.
- [[Channel Memory Layout (hchan Struct)]] — Dissecting hchan fields: ring buffer, lock mutex, sendq and recvq wait queues.
- [[sudog Wait Queue Architecture]] — How sudog structs wrap waiting goroutines and integrate with runtime pools.
- [[Direct Stack-to-Stack Copy Optimization]] — Lockless direct memmove between goroutine stacks bypassing intermediate channel buffer.
- [[selectgo Runtime Multiplexing Algorithm]] — Pseudo-random case shuffling, lock acquisition ordering, and non-blocking select.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]

