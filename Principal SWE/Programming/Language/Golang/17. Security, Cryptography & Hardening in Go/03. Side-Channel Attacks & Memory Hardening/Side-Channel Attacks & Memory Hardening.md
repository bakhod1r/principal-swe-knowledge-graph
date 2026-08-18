---
title: Side-Channel Attacks & Memory Hardening
tags:
  - golang
  - security
  - principal-swe
parent: "[[Security, Cryptography & Hardening in Go]]"
---

# Side-Channel Attacks & Memory Hardening

Constant-time operations (crypto/subtle), memory zeroing (memclr), secret masking in heap dumps/logs, and unix.Mlock.

```text
Side-Channel Attacks & Memory Hardening
│
├── [[Constant-Time Cryptographic Operations (crypto-subtle)]]
├── [[Sensitive Secret Erasure & Memory Zeroing (memclr)]]
├── [[Preventing Secret Leakage in Heap Dumps & Logs]]
└── [[Memory Locking (unix.Mlock) to Prevent Swap Spills]]
```

---

## 🗂️ Topics

- [[Constant-Time Cryptographic Operations (crypto-subtle)]] — Eliminating CPU timing attack vulnerabilities using subtle.ConstantTimeCompare.
- [[Sensitive Secret Erasure & Memory Zeroing (memclr)]] — Zeroing byte slices containing private keys/passwords before GC deallocation.
- [[Preventing Secret Leakage in Heap Dumps & Logs]] — Implementing fmt.Stringer and slog.LogValuer masking to redact sensitive fields.
- [[Memory Locking (unix.Mlock) to Prevent Swap Spills]] — Pinning cryptographic keys in RAM using unix.Mlock to prevent swapping to disk.

---

## 🔗 References
- ⬆️ Parent: [[Security, Cryptography & Hardening in Go]]
- 🎓 Root: [[Principal SWE]]
