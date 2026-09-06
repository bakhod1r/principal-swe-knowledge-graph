---
title: Bit Array (Packed Bitset)
tags:
  - computer-science
  - data-structures
  - arrays
  - principal-swe
parent: "[[Array]]"
---

# 📦 Bit Array (Packed Bitset)

Dense 64-to-1 memory compression using 64-bit word arrays, word-level SIMD boolean algebra, hardware population count (POPCNT), and bitmap indexing.

```text
Bit Array (Packed Bitset)
│
├── [[Bit Array (Packed Bitset Operations)]]
├── [[Bitwise Word-Level Parallelism and Population Count (POPCNT)]]
├── [[Bitset Set Algebra (SIMD 64-Bit Word Operations)]]
└── [[Roaring Bitmaps and Run-Length Encoded (RLE) Hybrid Bitsets]]
```

---

## 🗂️ Topics & Implementations

- [[Bit Array (Packed Bitset Operations)]] — Bitwise address translation: word offset (i >> 6) and bit shift (1ULL << (i & 63)).
- [[Bitwise Word-Level Parallelism and Population Count (POPCNT)]] — Processing 64 elements per CPU cycle and hardware Hamming weight calculation.
- [[Bitset Set Algebra (SIMD 64-Bit Word Operations)]] — Vectorized set union (OR), intersection (AND), and difference (AND NOT) across bit arrays.
- [[Roaring Bitmaps and Run-Length Encoded (RLE) Hybrid Bitsets]] — Adaptive compressed bitmaps switching between bitset, array, and RLE containers.

---

## 🔗 References
- ⬆️ Parent: [[Array]]
- 📚 Module: `Array`
