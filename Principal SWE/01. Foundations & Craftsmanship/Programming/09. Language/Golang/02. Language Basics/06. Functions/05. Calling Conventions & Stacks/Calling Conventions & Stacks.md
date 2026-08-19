---
title: Calling Conventions & Stacks
tags:
  - golang
  - functions
  - calling-conventions
  - stacks
  - low-level
  - principal-swe
parent: "[[Functions]]"
---

# Calling Conventions & Stacks

Register calling convention (ABIInternal), stack splitting, and `morestack` dynamic stack growth.

```text
Calling Conventions & Stacks
│
├── [[Register-Based Calling Convention (ABIInternal)]]
└── [[Stack Splitting & morestack Reallocation]]
```

---

## 🗂️ Topics

- [[Register-Based Calling Convention (ABIInternal)]] — Passing up to 9 integer and 15 float arguments in CPU registers (Go 1.17+ ABIInternal) vs legacy stack-based ABI0.
- [[Stack Splitting & morestack Reallocation]] — Dynamic 2KB goroutine stacks, stack guard checks, `runtime.morestack` preamble, and contiguous stack doubling.

---

## 🔗 References
- ⬆️ Parent: [[Functions]]
- 📚 Module: `Language Basics`
