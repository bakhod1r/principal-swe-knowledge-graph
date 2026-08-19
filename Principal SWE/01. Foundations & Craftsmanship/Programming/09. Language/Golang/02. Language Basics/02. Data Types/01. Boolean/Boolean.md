---
title: Boolean
tags:
  - golang
  - types
  - boolean
  - principal-swe
parent: "[[Data Types]]"
---

# Boolean

Boolean primitives, logical operators, truth tables, short-circuit evaluation, memory footprints, and branchless optimizations.

```text
Boolean
│
├── [[Boolean Primitives (bool)]]
├── [[Logical Operators (AND, OR, NOT)]]
├── [[Short-Circuit Evaluation]]
├── [[Boolean in Control Flow]]
├── [[Boolean Memory Layout and unsafe.Sizeof(bool)]]
├── [[Branchless Logic and CMOV Optimization]]
└── [[Bit-Packed Boolean Flags (Bitset vs bool slice)]]
```

---

## 🗂️ Topics

- [[Boolean Primitives (bool)]] — true and false literals, 1-byte memory representation, and zero-value (false).
- [[Logical Operators (AND, OR, NOT)]] — Logical AND (&&), OR (||), NOT (!) and operator precedence rules.
- [[Short-Circuit Evaluation]] — Non-evaluation of right-hand operands when left operand determines outcome.
- [[Boolean in Control Flow]] — Using boolean expressions in if conditions, select statements, and switch cases.
- [[Boolean Memory Layout and unsafe.Sizeof(bool)]] — Why bool requires 1 full byte (8 bits) for byte addressability and alignment padding.
- [[Branchless Logic and CMOV Optimization]] — Writing branchless boolean arithmetic to prevent CPU pipeline flushes and branch mispredictions.
- [[Bit-Packed Boolean Flags (Bitset vs bool slice)]] — Storing 8 booleans per byte via bitfields to eliminate 87.5% memory overhead compared to `[]bool`.

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 📚 Module: `Language Basics`
