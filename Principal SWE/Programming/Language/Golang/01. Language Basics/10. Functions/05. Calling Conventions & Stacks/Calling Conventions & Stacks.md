---
title: Calling Conventions & Stacks
tags:
  - golang
  - functions
  - principal-swe
parent: "[[Functions]]"
---

# Calling Conventions & Stacks

Register calling convention, stack splitting, and morestack reallocations.

```text
Calling Conventions & Stacks
│
├── [[Register-Based Calling Convention (ABIInternal)]]
├── [[Stack-Based ABI0 vs ABIInternal]]
└── [[Stack Splitting & morestack Reallocation]]
```

---

## 🗂️ Topics

- [[Register-Based Calling Convention (ABIInternal)]] — Passing up to 9 integer and 15 float arguments in CPU registers for 30-40% speedup.
- [[Stack-Based ABI0 vs ABIInternal]] — Legacy stack frame argument passing vs modern register-based calling convention.
- [[Stack Splitting & morestack Reallocation]] — Stack frame limit checks, morestack preamble, and contiguous stack memory copying.

---

## 🔗 References
- ⬆️ Parent: [[Functions]]

