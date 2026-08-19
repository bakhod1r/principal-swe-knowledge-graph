---
title: "Bit Manipulation Hacks and Twiddles"
tags:
  - golang
  - bitwise
  - bit-hacks
  - algorithms
  - principal-swe
parent: "[[Bitwise Operations]]"
---

# Bit Manipulation Hacks and Twiddles

## 1. Definition

**Bit Manipulation Hacks** (also known as *Bit Twiddling Hacks*) are constant-time $\mathcal{O}(1)$ algorithmic techniques that compute arithmetic, logic, and bit-level properties using low-level bitwise operations without branching or looping.

### Core Mathematical Invariants:
1. **Two's Complement Negation:** $-x \equiv (\mathbin{\wedge}x) + 1$
2. **Kernighan's Law:** $x \mathbin{\&} (x - 1)$ clears the lowest set (rightmost 1) bit in $x$.
3. **Lowest Set Bit Isolation:** $x \mathbin{\&} (-x)$ extracts only the lowest set bit as an isolated power of 2.
4. **Power-of-Two Invariant:** An integer $x > 0$ is an exact power of 2 if and only if $x \mathbin{\&} (x - 1) == 0$.

---

## 2. Mental Model & Twiddle Catalog

```text
┌──────────────────────────────────────┬────────────────────────┬────────────────────────────────────────────────────────┐
│ Operation                            │ Bit Expression         │ Example (x = 0b00101100 = 44)                          │
├──────────────────────────────────────┼────────────────────────┼────────────────────────────────────────────────────────┤
│ Clear lowest set bit                 │ x & (x - 1)            │ 0b00101100 & 0b00101011 -> 0b00101000 (40)             │
│ Isolate lowest set bit               │ x & (-x)               │ 0b00101100 & 0b11010100 -> 0b00000100 (4)              │
│ Isolate lowest 0 bit                 │ ^x & (x + 1)           │ ^0b00101100 & 0b00101101 -> 0b00000001 (1)             │
│ Set lowest 0 bit                     │ x | (x + 1)            │ 0b00101100 | 0b00101101 -> 0b00101101 (45)             │
│ Clear trailing 1s                    │ x & (x + 1)            │ 0b00101111 & 0b00110000 -> 0b00100000 (32)             │
│ Check power of 2                     │ x > 0 && x&(x-1) == 0  │ 44 &(43) != 0 -> False, 64 &(63) == 0 -> True          │
└──────────────────────────────────────┴────────────────────────┴────────────────────────────────────────────────────────┘
```

---

## 3. Usage

### Production Go Implementation: Comprehensive Bit Manipulation Library

```go
package bitwise

// IsolateLowestSetBit returns the value with only the lowest set bit of x preserved
func IsolateLowestSetBit(x int64) int64 {
	return x & (-x)
}

// ClearLowestSetBit removes the rightmost 1-bit from x
func ClearLowestSetBit(x uint64) uint64 {
	return x & (x - 1)
}

// CountSetBitsKernighan counts set bits in O(number of 1-bits)
func CountSetBitsKernighan(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1 // Clears lowest set bit in each iteration
		count++
	}
	return count
}

// BranchlessMin returns the minimum of two signed integers without conditional branching
func BranchlessMin(x, y int32) int32 {
	// If x < y, diff >> 31 is -1 (all 1s). If x >= y, diff >> 31 is 0.
	diff := x - y
	return y + (diff & (diff >> 31))
}

// BranchlessMax returns the maximum of two signed integers without conditional branching
func BranchlessMax(x, y int32) int32 {
	diff := x - y
	return x - (diff & (diff >> 31))
}

// BranchlessAbs computes absolute value of int32 without conditional jumps
func BranchlessAbs(x int32) int32 {
	mask := x >> 31 // 0 if x >= 0, -1 if x < 0
	return (x + mask) ^ mask
}

// RoundUpToPowerOfTwo32 computes the smallest power of 2 >= x using bit smearing
func RoundUpToPowerOfTwo32(x uint32) uint32 {
	if x <= 1 {
		return 1
	}
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x + 1
}

// InterleaveBits32 (Morton Code / Z-Order Curve 2D coordinate packing)
func InterleaveBits32(x, y uint16) uint32 {
	zx := uint32(x)
	zy := uint32(y)

	zx = (zx | (zx << 8)) & 0x00FF00FF
	zx = (zx | (zx << 4)) & 0x0F0F0F0F
	zx = (zx | (zx << 2)) & 0x33333333
	zx = (zx | (zx << 1)) & 0x55555555

	zy = (zy | (zy << 8)) & 0x00FF00FF
	zy = (zy | (zy << 4)) & 0x0F0F0F0F
	zy = (zy | (zy << 2)) & 0x33333333
	zy = (zy | (zy << 1)) & 0x55555555

	return zx | (zy << 1)
}
```

---

## 4. Gotchas

- **Signed Integer Overflow on `-x`:** In Two's Complement, the minimum value of a signed integer (e.g. `math.MinInt64 = -9223372036854775808`) cannot be negated because `+9223372036854775808` exceeds `math.MaxInt64`. Performing `x & (-x)` on `math.MinInt64` wraps around to `MinInt64`.
- **Zero Input in Power-of-Two Check:** The expression `x & (x - 1) == 0` evaluates to `true` when $x = 0$ (because $0 \ \& \ (-1) == 0$). You must explicitly guard with `x > 0 && (x & (x - 1)) == 0`.
- **Prefer `math/bits` over Manual Twiddles in Go:** While manual bit twiddling algorithms (like parallel bit counters or De Bruijn lookups) are historically famous, in modern Go `bits.OnesCount64()` and `bits.TrailingZeros64()` are directly compiled into native single CPU instructions (`POPCNT`, `TZCNT`), which are substantially faster than software twiddles.

---

## 🔗 References
- ⬆️ Parent: [[Bitwise Operations]]
- 📚 Module: `Language Basics`
