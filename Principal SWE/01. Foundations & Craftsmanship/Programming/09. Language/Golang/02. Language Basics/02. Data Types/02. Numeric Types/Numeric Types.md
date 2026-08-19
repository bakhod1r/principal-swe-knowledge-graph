---
title: Numeric Types
tags:
  - golang
  - types
  - numerics
  - principal-swe
parent: "[[Data Types]]"
---

# Numeric Types

Signed/unsigned integers, floating-point numbers, IEEE-754 precision, complex numbers, arbitrary precision, and SIMD vectorization.

```text
Numeric Types
│
├── [[Signed Integers (int8 to int64, int)]]
├── [[Unsigned Integers (uint8 to uint64, uint, uintptr)]]
├── [[Floating Point Numbers (float32, float64)]]
├── [[Complex Numbers (complex64, complex128)]]
├── [[Overflow and Precision Loss]]
├── [[Two's Complement Representation & Integer Overflow Wrap]]
├── [[IEEE-754 Precision & NaN Semantics]]
├── [[Epsilon Floating-Point Comparisons]]
├── [[Float Fast-Math, Subnormals, and Denormals]]
├── [[uintptr vs unsafe.Pointer Semantics]]
├── [[math-big Package (Arbitrary-Precision Arithmetic)]]
├── [[Financial Arithmetic Alternatives (shopspring-decimal)]]
└── [[SIMD and Vectorized Math Operations]]
```

---

## 🗂️ Topics

- [[Signed Integers (int8 to int64, int)]] — Architecture-dependent `int` vs fixed-width `int8`, `int16`, `int32`, `int64`.
- [[Unsigned Integers (uint8 to uint64, uint, uintptr)]] — `byte` (`uint8`), `uint`, and `uintptr` unsigned types.
- [[Floating Point Numbers (float32, float64)]] — IEEE-754 floating point representation and precision limits.
- [[Complex Numbers (complex64, complex128)]] — Complex number arithmetic, `real()` and `imag()` built-in functions.
- [[Overflow and Precision Loss]] — Silent integer wrap-around overflow behavior and truncation risks.
- [[Two's Complement Representation & Integer Overflow Wrap]] — Hardware binary integer representation and two's complement sign arithmetic.
- [[IEEE-754 Precision & NaN Semantics]] — Special values: `+Inf`, `-Inf`, `NaN`, signed zeros, and IEEE-754 bit representations.
- [[Epsilon Floating-Point Comparisons]] — Safe floating point equality testing using machine epsilon thresholds.
- [[Float Fast-Math, Subnormals, and Denormals]] — Subnormal numbers, denormals flush-to-zero, and CPU pipeline penalties.
- [[uintptr vs unsafe.Pointer Semantics]] — Why `uintptr` is an unmanaged integer that escape analysis and GC ignore.
- [[math-big Package (Arbitrary-Precision Arithmetic)]] — Multi-precision arithmetic for huge integers (`big.Int`) and high-precision floats (`big.Float`).
- [[Financial Arithmetic Alternatives (shopspring-decimal)]] — Exact base-10 fixed-point arithmetic for currency and financial calculations.
- [[SIMD and Vectorized Math Operations]] — Hardware vector registers (AVX2/NEON) and batch numeric slice processing.

---

## 🔗 References
- ⬆️ Parent: [[Data Types]]
- 📚 Module: `Language Basics`
