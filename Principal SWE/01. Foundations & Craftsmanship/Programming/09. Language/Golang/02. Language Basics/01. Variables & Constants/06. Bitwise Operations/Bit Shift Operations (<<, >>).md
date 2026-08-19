---
title: "Bit Shift Operations (<<, >>)"
tags:
  - golang
  - bitwise
  - bit-shift
  - variables
  - principal-swe
parent: "[[Bitwise Operations]]"
---

# Bit Shift Operations (`<<`, `>>`)

## 1. Definition

In Go, bit shift operators shift the binary bit representation of the left operand leftward (`<<`) or rightward (`>>`) by the number of bit positions specified by the right operand.

- **Left Shift (`x << n`):** Shifts the bits of `x` left by `n` positions. The vacated least significant bits (LSBs) on the right are filled with **zeros**. Bits shifted past the type boundary are discarded. Equivalent to arithmetic multiplication:
  $$x \ll n \equiv x \times 2^n \pmod{2^{\text{width}}}$$
- **Right Shift (`x >> n`):** Shifts the bits of `x` right by `n` positions. Bits shifted past the right boundary are discarded.
  - **Logical Right Shift (for Unsigned Types):** The vacated most significant bits (MSBs) on the left are filled with **zeros**.
    $$u \gg n \equiv \lfloor u / 2^n \rfloor$$
  - **Arithmetic Right Shift (for Signed Types):** The vacated MSBs on the left are filled with copies of the original **Sign Bit** (preserving Two's Complement negative value).
    $$s \gg n \equiv \lfloor s / 2^n \rfloor$$

### Strict Go Shift Type Rule
> [!IMPORTANT]
> In Go, the right-hand operand (the shift amount `n`) must be an **unsigned integer** type (e.g. `uint`, `uint8`, `uint64`) or an untyped constant that fits in an unsigned integer. Attempting to shift by a signed variable `var s int = 2; x << s` causes a compile error: `invalid operation: x << s (shift count type int, must be unsigned integer)`.

---

## 2. Mental Model

### Logical vs Arithmetic Shift
```text
1. Left Shift (Unsigned/Signed): uint8(0b00010110) << 2 -> 0b01011000
   [0 0 0 1 0 1 1 0] << 2
   ┌─┬─┬─┬─┬─┬─┬─┬─┐
   │0│1│0│1│1│0│0│0│  <- Zeros always fill the right
   └─┴─┴─┴─┴─┴─┴─┴─┘

2. Logical Right Shift (Unsigned uint8): uint8(0b11001000) >> 2 -> 0b00110010
   [1 1 0 0 1 0 0 0] >> 2
   ┌─┬─┬─┬─┬─┬─┬─┬─┐
   │0│0│1│1│0│0│1│0│  <- Zeros fill the left (Logical Shift)
   └─┴─┴─┴─┴─┴─┴─┴─┘

3. Arithmetic Right Shift (Signed int8): int8(0b11001000 = -56) >> 2 -> 0b11110010 = -14
   [1 1 0 0 1 0 0 0] >> 2  (Sign bit is 1)
   ┌─┬─┬─┬─┬─┬─┬─┬─┐
   │1│1│1│1│0│0│1│0│  <- Sign bit (1) replicates on the left (Arithmetic Shift)
   └─┴─┴─┴─┴─┴─┴─┴─┘
```

---

## 3. Usage

### Production Go Implementation: Fast Binary Serialization, Varints, and Power-of-Two Alignments

```go
package bitwise

import (
	"encoding/binary"
	"fmt"
)

// FastMultiplyByPowerOfTwo computes x * (2^k) via left shift
func FastMultiplyByPowerOfTwo(x uint64, k uint) uint64 {
	return x << k
}

// FastDivideByPowerOfTwo computes x / (2^k) via right shift
func FastDivideByPowerOfTwo(x uint64, k uint) uint64 {
	return x >> k
}

// AlignUpToPowerOfTwo aligns value x up to the nearest multiple of powerOfTwo (e.g. 64B cacheline or 4KB page)
func AlignUpToPowerOfTwo(x uint64, alignment uint64) uint64 {
	// alignment must be power of two: e.g. 64 (0x40) -> mask = 63 (0x3F)
	mask := alignment - 1
	return (x + mask) &^ mask
}

// PackRGBA packs 4 byte color channels into a single 32-bit uint32
func PackRGBA(r, g, b, a uint8) uint32 {
	return (uint32(r) << 24) |
		(uint32(g) << 16) |
		(uint32(b) << 8) |
		uint32(a)
}

// UnpackRGBA unpacks a 32-bit uint32 into individual 8-bit RGBA channels
func UnpackRGBA(packed uint32) (r, g, b, a uint8) {
	r = uint8((packed >> 24) & 0xFF)
	g = uint8((packed >> 16) & 0xFF)
	b = uint8((packed >> 8) & 0xFF)
	a = uint8(packed & 0xFF)
	return r, g, b, a
}

// EncodeVarint32 encodes a 32-bit unsigned int into protobuf-style Varint bytes
func EncodeVarint32(buf []byte, val uint32) int {
	i := 0
	for val >= 0x80 {
		buf[i] = byte(val&0x7F) | 0x80
		val >>= 7
		i++
	}
	buf[i] = byte(val & 0x7F)
	return i + 1
}
```

---

## 4. Gotchas

- **Shift Count Type Requirement:**
  ```go
  var x uint32 = 100
  shift := 2 // shift is int
  // result := x << shift // COMPILE ERROR: shift count type int, must be unsigned integer
  result := x << uint(shift) // CORRECT: Explicit conversion to unsigned
  ```
- **Overshift / Shift Count Exceeding Type Width:**
  - In Go, constant overshifts (e.g. `var x uint32 = 1; y := x << 32`) are detected at compile time as errors: `invalid operation: x << 32 (shift count 32 >= width of uint32)`.
  - At runtime, shifting by $n \ge \text{width}$ results in:
    - For unsigned types: `0`
    - For signed types: `0` (if positive) or `-1` (if negative, due to sign propagation)
- **Arithmetic Right Shift Floor Division of Negative Odd Numbers:**
  - Arithmetic right shift on negative odd numbers rounds towards negative infinity ($\lfloor x/2 \rfloor$), whereas Go's integer division `/` truncates towards zero.
  - Example: `-5 / 2 == -2`, but `int8(-5) >> 1 == -3`.
- **Precedence with Arithmetic:**
  - Because `<<` and `>>` have higher precedence than `+` and `-`, `1 << 2 + 1` evaluates to `(1 << 2) + 1 = 5` in Go, NOT `1 << (2 + 1) = 8`.

---

## 🔗 References
- ⬆️ Parent: [[Bitwise Operations]]
- 📚 Module: `Language Basics`
