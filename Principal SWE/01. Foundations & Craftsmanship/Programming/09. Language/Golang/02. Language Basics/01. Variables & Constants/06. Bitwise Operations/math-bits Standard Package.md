---
title: "math/bits Standard Package"
tags:
  - golang
  - bitwise
  - math-bits
  - hardware-intrinsics
  - principal-swe
parent: "[[Bitwise Operations]]"
---

# `math/bits` Standard Package

## 1. Definition

The standard library package **`math/bits`** (introduced in Go 1.9) provides optimized bit-level operations and multi-word arithmetic primitives.

### Direct Compiler Intrinsics (Hardware Acceleration)
Functions in `math/bits` are recognized by the Go compiler (`cmd/compile`) as **compiler intrinsics**. On modern x86-64 and ARM64 processors, calls to these functions are replaced with single-cycle CPU instructions with **zero function call overhead**:

| `math/bits` Function | Purpose | x86-64 Assembly | ARM64 Assembly |
| :--- | :--- | :--- | :--- |
| `bits.OnesCount64(x)` | Population Count (number of set 1-bits) | `POPCNT` | `VCNT` / `CNT` |
| `bits.LeadingZeros64(x)` | Count Leading Zeros from MSB | `LZCNT` / `BSR` | `CLZ` |
| `bits.TrailingZeros64(x)`| Count Trailing Zeros from LSB | `TZCNT` / `BSF` | `RBIT` + `CLZ` |
| `bits.RotateLeft64(x, k)`| Circular Bit Rotation Left / Right | `ROL` / `ROR` | `ROR` |
| `bits.ReverseBytes64(x)` | Endianness byte swap | `BSWAP` | `REV` |
| `bits.Reverse64(x)` | Reverse bit ordering | Fast algorithm | `RBIT` |
| `bits.Len64(x)` | Min bits needed to store $x$ ($\lceil \log_2(x+1) \rceil$) | `BSR` / `LZCNT` | `CLZ` |

---

## 2. Mental Model

### Visualizing Core Bit Operations
```text
Value: 0000 0000 0000 0000 0000 0000 0010 1100 (44)

1. Leading Zeros:   [0000 0000 ... 000] 0 1 0 1 1 0 0 -> LeadingZeros = 26
2. Length:                               [1 0 1 1 0 0] -> Len = 6 (Bits needed)
3. Population Count:                      1   1 1     -> OnesCount = 3 (Number of 1s)
4. Trailing Zeros:                              [0 0] -> TrailingZeros = 2
```

### Multi-Precision Carry Arithmetic
```text
Add with Carry:  (sum, carryOut) = bits.Add64(x, y, carryIn)
Multiplication:  (hi, lo)        = bits.Mul64(x, y)          // Full 128-bit product!
Division:        (quo, rem)      = bits.Div64(hi, lo, y)     // 128-bit / 64-bit division
```

---

## 3. Usage

### Production Go Implementation: Fast Set Operations, Next Power of 2, and 128-Bit Multiplication

```go
package bitwise

import (
	"fmt"
	"math/bits"
)

// NextPowerOfTwo returns the smallest power of 2 >= v using bits.LeadingZeros64
func NextPowerOfTwo(v uint64) uint64 {
	if v <= 1 {
		return 1
	}
	// If already a power of 2, return v
	if (v & (v - 1)) == 0 {
		return v
	}
	// bits.Len64(v) gives the bit position of highest set bit
	return 1 << bits.Len64(v)
}

// IsPowerOfTwo checks if x is a power of 2 in O(1)
func IsPowerOfTwo(x uint64) bool {
	return x > 0 && (x&(x-1)) == 0
}

// Full128BitMultiply multiplies two 64-bit integers and produces a full 128-bit result (hi, lo)
func Full128BitMultiply(a, b uint64) (hi, lo uint64) {
	return bits.Mul64(a, b)
}

// IterateSetBits executes callback for each 1-bit index in a 64-bit bitmask in O(popcount) time
func IterateSetBits(bitmask uint64, fn func(bitIndex int)) {
	for bitmask != 0 {
		// Find lowest set bit position
		trailing := bits.TrailingZeros64(bitmask)
		fn(trailing)
		// Clear lowest set bit (Kernighan's trick)
		bitmask &= bitmask - 1
	}
}

// FastLog2 computes floor(log2(x)) in a single instruction
func FastLog2(x uint64) int {
	if x == 0 {
		return -1 // Undefined for 0
	}
	return bits.Len64(x) - 1
}

func ExampleMathBits() {
	var val uint64 = 44 // 0b00101100

	fmt.Printf("OnesCount: %d\n", bits.OnesCount64(val))       // 3
	fmt.Printf("LeadingZeros: %d\n", bits.LeadingZeros64(val)) // 58
	fmt.Printf("TrailingZeros: %d\n", bits.TrailingZeros64(val))// 2
	fmt.Printf("Bit Length: %d\n", bits.Len64(val))           // 6
	fmt.Printf("Next Pow2(44): %d\n", NextPowerOfTwo(val))     // 64

	fmt.Println("Iterating set bits in 44:")
	IterateSetBits(val, func(idx int) {
		fmt.Printf(" - Bit at index %d is set\n", idx) // indices 2, 3, 5
	})
}
```

---

## 4. Gotchas

- **`bits.Len64(0)` vs `bits.LeadingZeros64(0)`:**
  - `bits.Len64(0) == 0`
  - `bits.LeadingZeros64(0) == 64`
  - `bits.TrailingZeros64(0) == 64`
  - Always guard against $x == 0$ when calculating $\log_2(x) = \text{Len64}(x) - 1$.
- **`bits.Div64` Division Panic / Overflow:**
  - `bits.Div64(hi, lo, y)` panics if $y == 0$ (division by zero) **OR** if $hi \ge y$ (the quotient cannot fit into a 64-bit integer, resulting in integer overflow panic).
- **Shift Count on Rotate:** `bits.RotateLeft64(x, k)` supports both positive and negative values of $k$. Negative $k$ rotates right. The shift amount is automatically reduced modulo 64.

---

## 🔗 References
- ⬆️ Parent: [[Bitwise Operations]]
- 📚 Module: `Language Basics`
