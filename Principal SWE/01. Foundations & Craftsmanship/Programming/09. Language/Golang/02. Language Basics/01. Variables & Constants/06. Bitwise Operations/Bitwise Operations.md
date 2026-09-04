---
title: Bitwise Operations
tags:
  - golang
  - variables
  - bitwise
  - low-level
  - principal-swe
parent: "[[Variables & Constants]]"
---

# ⚡ Bitwise Operations

Comprehensive guide to low-level bitwise manipulation, binary arithmetic, hardware intrinsics, and bitmasking patterns in Go.

```mermaid
graph TD
    classDef op fill:#1e1e2e,stroke:#89b4fa,stroke-width:2px,color:#cdd6f4;
    classDef pkg fill:#181825,stroke:#a6e3a1,stroke-width:2px,color:#cdd6f4;

    BitOps["Go Bitwise Primitives"]:::op --> Logic["Logical Bitwise<br/>(&, |, ^, Unary ^)"]:::op
    BitOps --> Clear["Bit Clear (AND NOT)<br/>(&^) -> Maps to ARM BIC / x86 ANDN"]:::op
    BitOps --> Shifts["Bit Shifts<br/>(<< Left Shift, >> Right Shift)"]:::op
    BitOps --> Intrinsics["Hardware Intrinsics<br/>(math/bits: POPCNT, LZCNT, TZCNT)"]:::pkg
    
    Logic & Clear & Shifts --> Patterns["Production Systems Patterns<br/>(Bitmasks, Bitsets, SWAR, Endianness, Varints)"]:::op
```

---

## 🗂️ Knowledge Deep-Dive

```text
Bitwise Operations
│
├── [[Bitwise AND (&)]]
├── [[Bitwise OR (Pipe)]]
├── [[Bitwise XOR and NOT (Caret)]]
├── [[Bit Clear Operator (AND NOT)]]
├── [[Left Shift Operator (<<)]]
├── [[Right Shift Operator (>>)]]
├── [[math-bits Standard Package]]
├── [[Bitmasking and Bit Flags Patterns]]
├── [[Bit Manipulation Hacks and Twiddles]]
├── [[Bitsets and Bitmaps Implementation]]
├── [[Endianness and Byte Order (encoding-binary)]]
├── [[SWAR and Bit-Level Parallelism]]
└── [[Varints and Variable-Length Encoding]]
```

---

## 🗂️ Topics

- [[Bitwise AND (&)]] — Bitwise AND operator for masking, filtering, and intersection testing.
- [[Bitwise OR (Pipe)]] — Bitwise OR operator for combining flags, setting bits, and union operations.
- [[Bitwise XOR and NOT (Caret)]] — Bitwise XOR for toggling and unary `^` for bitwise inversion (complement).
- [[Bit Clear Operator (AND NOT)]] — Go-specific AND NOT operator for clearing specific bits in a mask.
- [[Left Shift Operator (<<)]] — Logical left shift for multiplication by powers of two and bit packing.
- [[Right Shift Operator (>>)]] — Logical (unsigned) vs Arithmetic (signed) right shifts for division and sign extension.
- [[math-bits Standard Package]] — Hardware compiler intrinsics for leading zeros, trailing zeros, popcount, and multi-word arithmetic.
- [[Bitmasking and Bit Flags Patterns]] — Production idioms for permission bitfields, state management, and atomic flag transitions.
- [[Bit Manipulation Hacks and Twiddles]] — High-performance branchless algorithms, power-of-two checks, and low-level bit twiddling.
- [[Bitsets and Bitmaps Implementation]] — High-density dynamic bitsets, bit vectors, and $O(N/64)$ set algebra in Go.
- [[Endianness and Byte Order (encoding-binary)]] — BigEndian vs LittleEndian memory layout, byte packing, and network wire protocol encoding.
- [[SWAR and Bit-Level Parallelism]] — SIMD Within A Register techniques for 8-byte parallel search and bitfield manipulation.
- [[Varints and Variable-Length Encoding]] — Protocol Buffers variable-length integer compression and bit-continuation encoding.

---

## 🔗 References
- ⬆️ Parent: [[Variables & Constants]]
- 📚 Module: `Language Basics`
