---
title: String Internals
tags:
  - golang
  - types
  - principal-swe
parent: "[[Data Types]]"
---

# String Internals

stringStruct layout, pointer and length fields, and zero-copy string conversions.

```text
String Internals
│
├── [[stringStruct Runtime Layout]]
├── [[Zero-Copy String to Byte Slice (unsafe)]]
├── [[String Header vs Slice Header]]
└── [[Sub-String Memory Retention]]
```

---

## 🗂️ Topics

- [[stringStruct Runtime Layout]] — Two-word struct: unsafe.Pointer Data and int Len (16 bytes on 64-bit).
- [[Zero-Copy String to Byte Slice (unsafe)]] — Converting []byte to string without allocation using unsafe.String and unsafe.StringData.
- [[String Header vs Slice Header]] — Comparing 2-word stringStruct with 3-word sliceHeader.
- [[Sub-String Memory Retention]] — How slicing a tiny piece of a huge string retains the entire backing byte array.

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 🎓 Root: [[Principal SWE]]
