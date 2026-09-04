---
title: Channel Operations
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Go Concurrency]]"
---

# Channel Operations

Channel send and receive mechanics, closing rules and broadcast signaling, length/capacity inspection, nil channel blocking, and closed channel behavior matrix.

```text
Channel Operations
│
├── [[Channel Send Operation (ch <- v)]]
├── [[Channel Receive Operation (v <- ch & v, ok <- ch)]]
├── [[Channel Close Operation (close(ch))]]
├── [[Channel Length & Capacity (len, cap)]]
├── [[Nil Channel Blocking & Deadlock Patterns]]
└── [[Closed Channel Read-Write Behavior Matrix]]
```

---

## 🗂️ Topics

- [[Channel Send Operation (ch <- v)]] — Synchronous vs asynchronous send execution, locking, and waking receivers.
- [[Channel Receive Operation (v <- ch & v, ok <- ch)]] — Reading from channels, comma-ok idiom, and unblocking senders.
- [[Channel Close Operation (close(ch))]] — Closing rules, panic conditions (closing closed or nil channels), and broadcast signaling.
- [[Channel Length & Capacity (len, cap)]] — Inspecting buffer length and total capacity with len() and cap().
- [[Nil Channel Blocking & Deadlock Patterns]] — Why reads and writes to nil channels block forever and how to use them in select.
- [[Closed Channel Read-Write Behavior Matrix]] — Matrix of reading zero values from closed channels vs panicking on writes.

---

## 🔗 References
- ⬆️ Parent: [[Go Concurrency]]
- 🔄 Related: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
