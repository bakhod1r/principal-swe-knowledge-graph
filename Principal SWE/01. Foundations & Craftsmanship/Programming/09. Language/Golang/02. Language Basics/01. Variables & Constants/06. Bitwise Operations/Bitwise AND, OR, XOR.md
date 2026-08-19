---
title: "Bitwise AND, OR, XOR"
tags:
  - golang
  - bitwise
  - variables
  - low-level
  - principal-swe
parent: "[[Bitwise Operations]]"
---

# Bitwise AND, OR, XOR

## 1. Definition

In Go, bitwise **AND** (`&`), **OR** (`|`), and **XOR** (`^`) operate on integer operands at the binary bit level. They are evaluated bit-by-bit independently across the binary representation of the operands.

- **Bitwise AND (`a & b`):** Result bit is `1` if and only if both corresponding bits in `a` and `b` are `1`. Used for **masking / filtering** specific bits and testing if a bit is set.
- **Bitwise OR (`a | b`):** Result bit is `1` if at least one corresponding bit in `a` or `b` is `1`. Used for **setting / combining** bit flags.
- **Bitwise XOR (`a ^ b`):** Result bit is `1` if the corresponding bits in `a` and `b` differ. Used for **toggling / flipping** bits, computing parities, and symmetric difference.
- **Bitwise NOT (Unary `^a`):** In Go, there is no `~` operator (unlike C/C++/Java). Unary `^` serves as the **bitwise complement / NOT** (inverts all bits: $0 \to 1, 1 \to 0$), equivalent to `^0 ^ a` or `-1 ^ a` for signed integers.

### Type & Precedence Invariants
- In Go, binary bitwise operators require **identical types** (no implicit type coercion).
- Operator Precedence in Go:
  1. `*`, `/`, `%`, `<<`, `>>`, `&`, `&^` (highest precedence among binary ops)
  2. `+`, `-`, `|`, `^` (additive level)
  3. `==`, `!=`, `<`, `<=`, `>`, `>=` (comparison)
  4. `&&`
  5. `||`

> [!IMPORTANT]
> In Go, `&` and `&^` have higher precedence than `+` and `-`, whereas in C/C++, `+` and `-` have higher precedence than bitwise operators!

---

## 2. Mental Model

### Binary Truth Tables
```text
┌───┬───┬─────────┬────────┬─────────┬──────────────┐
│ A │ B │ A & B   │ A | B  │ A ^ B   │ ^A (Unary)   │
│   │   │ (AND)   │ (OR)   │ (XOR)   │ (Bitwise NOT)│
├───┼───┼─────────┼────────┼─────────┼──────────────┤
│ 0 │ 0 │    0    │   0    │    0    │      1       │
│ 0 │ 1 │    0    │   1    │    1    │      1       │
│ 1 │ 0 │    0    │   1    │    1    │      0       │
│ 1 │ 1 │    1    │   1    │    0    │      0       │
└───┴───┴─────────┴────────┴─────────┴──────────────┘
```

### Hardware Register Mapping
```text
Operand A:   1 1 0 0 1 0 1 0  (0xCA)
Operand B:   1 0 1 0 1 1 0 0  (0xAC)
─────────────────────────────────────
A & B (AND): 1 0 0 0 1 0 0 0  (0x88 - Intersection / Mask)
A | B (OR):  1 1 1 0 1 1 1 0  (0xEE - Union / Set)
A ^ B (XOR): 0 1 1 0 0 1 1 0  (0x66 - Symmetric Difference / Toggle)
^A    (NOT): 0 0 1 1 0 1 0 1  (0x35 - Bitwise Inversion)
```

---

## 3. Usage

### Production Go Implementation: High-Performance Bit Flags and Fast Hash Combining

```go
package bitwise

import (
	"fmt"
)

// PermissionFlags represents access control bits using uint8
type PermissionFlags uint8

const (
	PermRead    PermissionFlags = 1 << 0 // 0000 0001 (1)
	PermWrite   PermissionFlags = 1 << 1 // 0000 0010 (2)
	PermExecute PermissionFlags = 1 << 2 // 0000 0100 (4)
	PermDelete  PermissionFlags = 1 << 3 // 0000 1000 (8)
)

// HasPermission checks if the target flag is set using Bitwise AND
func (p PermissionFlags) Has(flag PermissionFlags) bool {
	return (p & flag) == flag
}

// Add sets additional permission flags using Bitwise OR
func (p *PermissionFlags) Add(flag PermissionFlags) {
	*p |= flag
}

// Toggle flips the state of a permission flag using Bitwise XOR
func (p *PermissionFlags) Toggle(flag PermissionFlags) {
	*p ^= flag
}

// BitwiseNotExample demonstrates unary ^ for complement
func BitwiseNotExample() {
	var val uint8 = 0b00001111 // 15
	var inverted uint8 = ^val   // 0b11110000 (240)
	fmt.Printf("Original: %08b, Inverted: %08b\n", val, inverted)
}

// FastHashCombine combines two 64-bit hashes using XOR and prime multiplication (Boost hash_combine in Go)
func FastHashCombine(seed, v uint64) uint64 {
	// Golden ratio constant 0x9e3779b97f4a7c15
	return seed ^ (v + 0x9e3779b97f4a7c15 + (seed << 6) + (seed >> 2))
}

// InPlaceXORSwap swaps two integers without auxiliary memory using XOR
func InPlaceXORSwap(a, b *int) {
	if a != b { // Must not be the same memory address
		*a ^= *b
		*b ^= *a
		*a ^= *b
	}
}
```

---

## 4. Gotchas

- **No `~` Operator in Go:** Attempting to write `~x` causes a compile-time syntax error: `syntax error: unexpected ~`. Always use unary `^x` for bitwise complement.
- **Go Operator Precedence Differences with C:** In C/C++, `a & b == c` is parsed as `a & (b == c)`. In Go, `&` has higher precedence than `==`, so `a & b == c` is parsed as `(a & b) == c`. However, explicit parentheses are strongly advised for code clarity.
- **Signed Integer Bitwise Inversion:** Inverting a signed integer (e.g. `var x int8 = 1; ^x`) produces `-2` due to Two's Complement representation ($\sim x = -x - 1$). Always use **unsigned types** (`uint8`, `uint32`, `uint64`) when manipulating raw bit patterns.
- **In-Place XOR Swap Aliasing:** Performing `*a ^= *b` when `a` and `b` point to the same variable will zero out the value (`x ^ x == 0`).

---

## 🔗 References
- ⬆️ Parent: [[Bitwise Operations]]
- 📚 Module: `Language Basics`
