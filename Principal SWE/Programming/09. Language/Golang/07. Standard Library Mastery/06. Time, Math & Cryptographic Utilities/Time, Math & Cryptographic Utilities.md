---
title: Time, Math & Cryptographic Utilities
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# Time, Math & Cryptographic Utilities

time package monotonic vs wall clocks, timers, math/big arbitrary precision, and math/rand/v2.

```text
Time, Math & Cryptographic Utilities
│
├── [[time Package Monotonic vs Wall Clocks]]
├── [[time.Ticker & time.Timer Resource Management]]
├── [[math-big Arbitrary-Precision Arithmetic]]
└── [[math-rand-v2 Modern Fast Pseudo-Random Generator]]
```

---

## 🗂️ Topics

- [[time Package Monotonic vs Wall Clocks]] — Understanding wall clock reading (date/time) vs monotonic clock reading for elapsed duration calculations.
- [[time.Ticker & time.Timer Resource Management]] — Preventing resource leaks: time.After in loops, resetting timers, and ticker stop callbacks.
- [[math-big Arbitrary-Precision Arithmetic]] — Calculating arbitrary-precision integers (big.Int), floats (big.Float), and rational numbers.
- [[math-rand-v2 Modern Fast Pseudo-Random Generator]] — Go 1.22+ math/rand/v2: ChaCha8-based high-speed non-cryptographic random generation.

- [[context Package Architecture & Cancellation Trees]] — Cancelation trees, deadline propagation, WithValue key isolation, and Cause propagation (Go 1.20+).
- [[math-cmplx Complex Mathematical Operations]] — Trigonometric, hyperbolic, exponential, and polar operations on complex64/complex128 numbers.
- [[hash Package Hierarchy & CRC32 Hardware Checksums]] — hash.Hash, hash.Hash64 interfaces, and hardware-accelerated CRC32 IEEE/Castagnoli checksums.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]

