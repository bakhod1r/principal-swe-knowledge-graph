---
title: "Bitmasking and Bit Flags Patterns"
tags:
  - golang
  - bitwise
  - bitmask
  - design-patterns
  - principal-swe
parent: "[[Bitwise Operations]]"
---

# Bitmasking and Bit Flags Patterns

## 1. Definition

**Bitmasking** is an architectural pattern in systems programming where multiple boolean states, options, or permission flags are packed into the individual bits of a single integer (typically `uint8`, `uint32`, or `uint64`).

### Core Bitmask Algebra:
- **Set a Flag:** `state |= Flag`
- **Clear a Flag:** `state &^= Flag` (or `state &= ^Flag`)
- **Toggle a Flag:** `state ^= Flag`
- **Check Single Flag:** `(state & Flag) != 0` (or `(state & Flag) == Flag`)
- **Check All Flags in Subset:** `(state & Submask) == Submask`
- **Check Any Flag in Subset:** `(state & Submask) != 0`

---

## 2. Mental Model

```text
Bit Position:   [7]   [6]   [5]   [4]   [3]   [2]   [1]   [0]
Flag Mask:      128    64    32    16     8     4     2     1
               ┌─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐
State:         │ Adm │ Del │ Mod │ Wri │ Rea │ Exe │ Vrf │ Act │
               └─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘
                 1     0     1     1     1     0     0     1  = 0xB9 (185)

Operations:
- Set Exec (Bit 2):   0xB9 |  0x04 -> 0xBD (Bit 2 becomes 1)
- Clear Admin (Bit 7):0xB9 &^ 0x80 -> 0x39 (Bit 7 becomes 0)
- Toggle Mod (Bit 5): 0xB9 ^  0x20 -> 0x99 (Bit 5 flipped from 1 to 0)
- Has Read (Bit 3):   0xB9 &  0x08 == 0x08 -> TRUE
```

---

## 3. Usage

### Production Go Implementation: High-Performance Concurrent Session Flags

```go
package bitwise

import (
	"fmt"
	"sync/atomic"
)

type SessionFlags uint32

const (
	FlagActive        SessionFlags = 1 << 0 // 1
	FlagAuthenticated SessionFlags = 1 << 1 // 2
	FlagMFAVerified   SessionFlags = 1 << 2 // 4
	FlagAdmin         SessionFlags = 1 << 3 // 8
	FlagRateLimited   SessionFlags = 1 << 4 // 16
	FlagSuspended     SessionFlags = 1 << 5 // 32
)

// UserSession demonstrates zero-allocation high-density state packing
type UserSession struct {
	ID    uint64
	flags atomic.Uint32
}

func NewUserSession(id uint64) *UserSession {
	s := &UserSession{ID: id}
	s.flags.Store(uint32(FlagActive))
	return s
}

// Has checks if specific flag is enabled
func (s *UserSession) Has(flag SessionFlags) bool {
	return (s.flags.Load() & uint32(flag)) != 0
}

// HasAll checks if all flags in mask are enabled
func (s *UserSession) HasAll(mask SessionFlags) bool {
	return (s.flags.Load() & uint32(mask)) == uint32(mask)
}

// Enable atomically sets one or more flags
func (s *UserSession) Enable(flag SessionFlags) {
	for {
		oldVal := s.flags.Load()
		newVal := oldVal | uint32(flag)
		if s.flags.CompareAndSwap(oldVal, newVal) {
			return
		}
	}
}

// Disable atomically clears one or more flags using Bit Clear (&^)
func (s *UserSession) Disable(flag SessionFlags) {
	for {
		oldVal := s.flags.Load()
		newVal := oldVal &^ uint32(flag)
		if s.flags.CompareAndSwap(oldVal, newVal) {
			return
		}
	}
}

// Toggle atomically flips target flags
func (s *UserSession) Toggle(flag SessionFlags) {
	for {
		oldVal := s.flags.Load()
		newVal := oldVal ^ uint32(flag)
		if s.flags.CompareAndSwap(oldVal, newVal) {
			return
		}
	}
}

func ExampleSessionBitmask() {
	session := NewUserSession(1001)

	// Step 1: User logs in with MFA
	session.Enable(FlagAuthenticated | FlagMFAVerified)

	fmt.Printf("Is Authenticated: %t\n", session.Has(FlagAuthenticated))
	fmt.Printf("Can Access Admin Dashboard: %t\n", session.HasAll(FlagAuthenticated|FlagMFAVerified|FlagAdmin))

	// Step 2: Grant Admin
	session.Enable(FlagAdmin)
	fmt.Printf("Can Access Admin Dashboard after grant: %t\n", session.HasAll(FlagAuthenticated|FlagMFAVerified|FlagAdmin))

	// Step 3: Revoke MFA
	session.Disable(FlagMFAVerified)
	fmt.Printf("Can Access Admin Dashboard after MFA drop: %t\n", session.HasAll(FlagAuthenticated|FlagMFAVerified|FlagAdmin))
}
```

---

## 4. Gotchas

- **Zero Value Collision:** If a flag is defined as `FlagNone = 0`, checking `state & FlagNone == FlagNone` will always evaluate to `true` (since `state & 0 == 0`). Always start bit flags from `1 << 0` ($1$).
- **Bit Shift Limit of Integer Type:** Storing more than 32 flags in a `uint32` (or 64 in `uint64`) overflows the integer width. If more than 64 flags are needed, use a `[]uint64` bitset or `big.Int`.
- **Non-Atomic RMW Races:** In concurrent environments, `session.flags |= FlagAdmin` is a non-atomic Read-Modify-Write (RMW) cycle. Concurrent goroutines will overwrite and drop each other's bit updates. Always use CAS atomic loops or `sync/atomic`.

---

## 🔗 References
- ⬆️ Parent: [[Bitwise Operations]]
- 📚 Module: `Language Basics`
