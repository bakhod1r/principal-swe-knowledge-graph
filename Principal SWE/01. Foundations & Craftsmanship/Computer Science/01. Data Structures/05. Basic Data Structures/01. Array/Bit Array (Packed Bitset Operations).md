---
title: "Bit Array (Packed Bitset Operations)"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Bit Array (Packed Bitset Operations)

## 1. Definition
A **Bit Array** (Bitset) compresses boolean arrays by packing 64 boolean flags into a single 64-bit unsigned integer (`uint64`), achieving **$64\text{x}$ memory compression** and enabling SIMD-speed parallel set operations (AND, OR, XOR) across 64 elements simultaneously.

---

## 2. Mental Model
```text
Standard bool Array (1 byte per bool):
[ 1B ][ 1B ][ 1B ][ 1B ][ 1B ][ 1B ][ 1B ][ 1B ] = 8 Bytes for 8 flags!

Bit Array (1 bit per bool):
[ 0b10110010 ] = 1 Byte for 8 flags!

Bit Arithmetic:
Word Index:  i >> 6        (i / 64)
Bit Offset:  1 << (i & 63) (i % 64)
Set Bit i:   words[i >> 6] |= (1 << (i & 63))
Test Bit i: (words[i >> 6] & (1 << (i & 63))) != 0
```

---

## 3. Usage
```go
type Bitset struct {
    words []uint64
}

func NewBitset(n int) *Bitset {
    return &Bitset{words: make([]uint64, (n+63)/64)}
}

func (b *Bitset) Set(i int)   { b.words[i>>6] |= 1 << (i & 63) }
func (b *Bitset) Clear(i int) { b.words[i>>6] &= ^(1 << (i & 63)) }
func (b *Bitset) Test(i int) bool {
    return (b.words[i>>6] & (1 << (i & 63))) != 0
}
```

---

## 4. Gotchas
- **Concurrent Bit Contention:** Setting bit $A$ and bit $B$ concurrently on the same 64-bit word without atomic instructions (`sync/atomic`) triggers data corruption on both flags.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

