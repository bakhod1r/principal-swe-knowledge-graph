---
title: "Bit Clear Operator (&^)"
tags:
  - golang
  - bitwise
  - variables
  - bit-clear
  - principal-swe
parent: "[[Bitwise Operations]]"
---

# Bit Clear Operator (`&^`)

## 1. Definition

In Go, the **Bit Clear** operator (`&^`), also known as **AND NOT**, is a unique built-in binary operator designed to clear (zero out) specific bits in a bitfield according to a mask.

For expression `z = x &^ y`:
- For each bit position $i$, if $y_i == 1$, then $z_i = 0$ (the bit is cleared).
- If $y_i == 0$, then $z_i = x_i$ (the original bit in $x$ is retained unchanged).

Mathematically and logically:
$$x \mathbin{\&^\wedge} y \equiv x \mathbin{\&} (\mathbin{\wedge} y)$$

### Hardware Mechanical Sympathy
- On **ARM architectures**, `&^` maps directly to the native single-cycle **`BIC` (Bit Clear)** assembly instruction (`BIC Rd, Rn, Rm`).
- On **x86-64 with BMI1 extension**, `&^` maps directly to the native **`ANDN` (Logical AND NOT)** instruction (`ANDN r32a, r32b, r/m32`), avoiding a separate `NOT` instruction and saving register execution cycles.

---

## 2. Mental Model

### Truth Table for `&^` (AND NOT)
```text
┌───┬───┬──────────────┬────────────────────────────────────────────────────────┐
│ X │ Y │ X &^ Y       │ Operational Behavior                                   │
├───┼───┼──────────────┼────────────────────────────────────────────────────────┤
│ 0 │ 0 │      0       │ Y is 0 -> Preserve X (0 remains 0)                     │
│ 0 │ 1 │      0       │ Y is 1 -> Clear bit (becomes 0)                        │
│ 1 │ 0 │      1       │ Y is 0 -> Preserve X (1 remains 1)                     │
│ 1 │ 1 │      0       │ Y is 1 -> Clear bit (1 is cleared to 0)                │
└───┴───┴──────────────┴────────────────────────────────────────────────────────┘
```

### Visual Bitfield Filtering
```text
Original Value X:    1 1 0 1 1 0 1 0  (0xDA)
Mask to Clear Y:     0 0 0 1 1 0 0 0  (Bits 3 and 4 must be cleared)
───────────────────────────────────────────────────────────────────
Result Z (X &^ Y):   1 1 0 0 0 0 1 0  (0xC2)
                          ▲ ▲
                          └─── Bits 3 & 4 explicitly cleared to 0
```

---

## 3. Usage

### Production Go Implementation: Atomic Flag Management and Subsystem Status

```go
package bitwise

import (
	"fmt"
	"sync/atomic"
)

type NodeState uint32

const (
	StateInitialized NodeState = 1 << 0 // 0001
	StateLeader      NodeState = 1 << 1 // 0010
	StateSyncing     NodeState = 1 << 2 // 0100
	StateDraining    NodeState = 1 << 3 // 1000
)

// ClearFlag removes specific flags from the state using Bit Clear (&^)
func ClearFlag(current NodeState, toRemove NodeState) NodeState {
	// z = current &^ toRemove
	return current &^ toRemove
}

// AtomicClearFlag atomically clears bits in a shared state using CAS loop
func AtomicClearFlag(statePtr *atomic.Uint32, flagMask uint32) {
	for {
		oldVal := statePtr.Load()
		newVal := oldVal &^ flagMask // Clear target bits
		if statePtr.CompareAndSwap(oldVal, newVal) {
			return
		}
	}
}

// MemoryPageMasking demonstrates clearing page offset to find page base address
func MemoryPageBaseAddress(addr uint64, pageSize uint64) uint64 {
	// Page size must be power of two (e.g. 4096 = 0x1000)
	pageOffsetMask := pageSize - 1 // 0x0FFF
	return addr &^ pageOffsetMask  // Clears offset bits, yielding page base
}

func ExampleBitClear() {
	var state NodeState = StateInitialized | StateLeader | StateSyncing // 0111 (7)
	fmt.Printf("Initial state: %04b\n", state)

	// Demote from leader: clear StateLeader bit
	state = state &^ StateLeader
	fmt.Printf("After demoting leader: %04b\n", state) // 0101 (5)

	// Step-down entirely: clear Syncing and Initialized
	state = state &^ (StateSyncing | StateInitialized)
	fmt.Printf("After full teardown: %04b\n", state) // 0000 (0)
}
```

---

## 4. Gotchas

- **Asymmetry of Operands:** Unlike `&`, `|`, and `^`, the `&^` operator is **non-commutative**:
  $$x \mathbin{\&^\wedge} y \neq y \mathbin{\&^\wedge} x$$
  - `x &^ y` clears bits in `x` where `y` has 1s.
  - `y &^ x` clears bits in `y` where `x` has 1s.
- **Syntactic Difference with C/C++:** Developers coming from C/C++ frequently write `x = x & ~y`, which does not compile in Go. You must either write `x = x &^ y` or `x = x & (^y)`. The idiomatic and cleaner Go form is always `x &^ y`.
- **Precedence in Complex Expressions:** `&^` has the highest binary operator precedence in Go (same level as `*`, `/`, `%`, `<<`, `>>`, `&`). In expressions like `a + b &^ c`, Go evaluates `(b &^ c)` first, then adds `a`. Always use parentheses when mixing arithmetic and bit-clearing.

---

## 🔗 References
- ⬆️ Parent: [[Bitwise Operations]]
- 📚 Module: `Language Basics`
