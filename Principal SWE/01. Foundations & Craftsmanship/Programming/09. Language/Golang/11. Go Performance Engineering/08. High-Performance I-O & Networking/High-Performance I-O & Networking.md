---
title: High-Performance I-O & Networking
tags:
  - golang
  - performance
  - principal-swe
parent: "[[Go Performance Engineering]]"
---

# High-Performance I-O & Networking

Buffered I/O stream tuning, zero-copy kernel transfers, connection pool optimization, and fast serialization.

```text
High-Performance I-O & Networking
│
├── [[Buffered Stream Processing with bufio]]
├── [[Zero-Copy Network I-O with sendfile & splice]]
├── [[HTTP Connection Pooling & Keep-Alive Tuning]]
└── [[Fast-Path Serialization (Protobuf vs Sonic vs Stdlib JSON)]]
```

---

## 🗂️ Topics

- [[Buffered Stream Processing with bufio]] — Custom buffer sizing (bufio.NewReaderSize) for minimizing OS read/write syscalls.
- [[Zero-Copy Network I-O with sendfile & splice]] — Bypassing user-space memory buffers for direct kernel-to-socket transfers.
- [[HTTP Connection Pooling & Keep-Alive Tuning]] — Optimizing MaxIdleConns, MaxIdleConnsPerHost, and IdleConnTimeout in high-load clients.
- [[Fast-Path Serialization (Protobuf vs Sonic vs Stdlib JSON)]] — Performance matrix and allocation profiles of modern serialization formats.

---

## 🔗 References
- ⬆️ Parent: `Performance Engineering & Profiling`

