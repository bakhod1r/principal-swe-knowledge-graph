---
title: Data Types
tags:
  - golang
  - basics
parent: "[[Language Basics]]"
---

# Data Types

Numeric types, booleans, runes, strings, conversions, and memory representations.

```text
Data Types
│
├── [[Integers (Signed and Unsigned)]]
├── [[Floating Points]]
├── [[Complex Numbers]]
├── [[Boolean]]
├── [[Runes & UTF-8]]
├── [[Strings]]
├── [[Raw String Literals]]
├── [[Interpreted String Literals]]
├── [[String Internals (stringStruct)]]
├── [[Numeric Conversions]]
└── [[Overflow and Precision]]
```

---

## 🗂️ Topics

- [[Integers (Signed and Unsigned)]] — int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint, uintptr.
- [[Floating Points]] — IEEE-754 float32, float64 precision and NaN/Infinity behaviors.
- [[Complex Numbers]] — complex64, complex128 arithmetic and real/imag builtins.
- [[Boolean]] — bool type (true, false) and logical operators (&&, ||, !).
- [[Runes & UTF-8]] — rune (int32) representing Unicode code points and UTF-8 encoding.
- [[Strings]] — Immutable byte slices, UTF-8 indexing, len() in bytes vs utf8.RuneCountInString.
- [[Raw String Literals]] — Multi-line unescaped strings enclosed in backticks.
- [[Interpreted String Literals]] — Double-quoted strings with escape sequences (\n, \t, \x).
- [[String Internals (stringStruct)]] — Two-word struct: pointer to byte array and length in bytes.
- [[Numeric Conversions]] — Explicit type casting between numeric types without implicit widening.
- [[Overflow and Precision]] — Integer arithmetic overflow behavior and floating-point precision loss.

---

## 🔗 References
- ⬆️ Parent: [[Language Basics]]
- 🎓 Root: [[Principal SWE]]
