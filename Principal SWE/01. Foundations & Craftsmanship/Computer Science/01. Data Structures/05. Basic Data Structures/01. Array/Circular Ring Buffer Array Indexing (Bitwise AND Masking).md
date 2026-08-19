---
title: "Circular Ring Buffer Array Indexing (Bitwise AND Masking)"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Circular Ring Buffer Array Indexing (Bitwise AND Masking)

## 1. Definition
**Circular Ring Buffer Array Indexing** wraps buffer read/write heads around fixed-size arrays without branching or expensive CPU modulo division (`% N`).
When capacity $N$ is chosen as a power of two ($N = 2^k$), index wraparound executes in a **single clock cycle** using a bitwise AND mask ($N-1$):
$$\text{WrappedIndex} = \text{Head} \ \& \ (N - 1)$$

---

## 2. Mental Model
```text
Capacity = 8 (Binary: 0b1000, Mask = 7: 0b0111)
Tail = 6 -> (6 & 7) = Index 6
Tail = 7 -> (7 & 7) = Index 7
Tail = 8 -> (8 & 7) = Index 0 (Instantaneous Wraparound with 0 branches!)
```

---

## 3. Usage
```go
type FastRingBuffer struct {
    buffer []byte
    mask   uint32
    head   uint32
    tail   uint32
}

func NewRingBuffer(powerOfTwo uint32) *FastRingBuffer {
    return &FastRingBuffer{
        buffer: make([]byte, powerOfTwo),
        mask:   powerOfTwo - 1,
    }
}

func (r *FastRingBuffer) Write(b byte) {
    r.buffer[r.tail&r.mask] = b
    r.tail++
}
```

---

## 4. Gotchas
- **Non-Power-of-Two Allocation:** If capacity is not strictly a power of two ($2^k$), the bitwise mask `& (N - 1)` will corrupt indices. Always enforce `(N & (N - 1)) == 0` at initialization.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

