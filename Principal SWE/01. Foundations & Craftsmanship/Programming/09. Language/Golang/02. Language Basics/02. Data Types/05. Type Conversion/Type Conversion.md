---
title: Type Conversion
tags:
  - golang
  - types
  - type-conversion
  - principal-swe
parent: "[[Data Types]]"
---

# Type Conversion

Explicit type conversions, string-to-byte casting, strconv parsing, zero-copy unsafe conversions, and compiler allocation eliminations.

```text
Type Conversion
│
├── [[Numeric Conversions]]
├── [[String to Byte Slice Conversions]]
├── [[String to Rune Slice Conversions]]
├── [[strconv Package Conversions]]
├── [[unsafe.String and unsafe.StringData (Go 1.20+)]]
├── [[unsafe.Slice and unsafe.SliceData (Go 1.17+)]]
├── [[Type Assertions vs Type Conversions]]
├── [[Defined Types vs Type Aliases (type T1 T2 vs type T1 = T2)]]
└── [[Compiler Optimization for Map Lookups with string(bytes) Keys]]
```

---

## 🗂️ Topics

- [[Numeric Conversions]] — Explicit casting between `int`, `float`, and unsigned integers without implicit widening.
- [[String to Byte Slice Conversions]] — Converting between `string` and `[]byte`: heap allocation, copying, and mutation safety.
- [[String to Rune Slice Conversions]] — Decoding UTF-8 string into `[]rune` slice of Unicode code points.
- [[strconv Package Conversions]] — `Atoi`, `Itoa`, `ParseFloat`, `FormatInt`, `AppendInt` string-to-number parsing and zero-alloc appending.
- [[unsafe.String and unsafe.StringData (Go 1.20+)]] — Modern standard zero-copy conversion from `*byte` pointer to `string`.
- [[unsafe.Slice and unsafe.SliceData (Go 1.17+)]] — Modern standard zero-copy conversion from raw memory pointers to `[]T` slice headers.
- [[Type Assertions vs Type Conversions]] — Compile-time static type casting vs runtime interface type assertions and type switches.
- [[Defined Types vs Type Aliases (type T1 T2 vs type T1 = T2)]] — Method sets, assignability rules, and underlying type identity.
- [[Compiler Optimization for Map Lookups with string(bytes) Keys]] — Compiler optimization that elides heap allocation when querying maps with byte slices.

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 📚 Module: `Language Basics`
