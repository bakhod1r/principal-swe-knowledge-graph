---
title: "Array Memory Copy (SIMD rep movsq and memmove)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Memory Copy (SIMD rep movsq and memmove)

## 1. Definition
**Array Memory Copy** transfers a contiguous block of bytes from a source address to a destination address.
Modern runtimes optimize memory copying using SIMD AVX-512 / NEON vectorized instructions (`rep movsq` / `memmove`), reaching memory bandwidth limits of over $50\text{ GB/s}$.

---

## 2. Mental Model
```text
Vectorized AVX-512 Copy:
Single Instruction Copies 64 Bytes (512 bits) per clock cycle:
Source: [ 64 Bytes Block ] ─── VMOVDQU64 ───> Destination: [ 64 Bytes Block ]

Overlapping Buffer Safety:
Source:      [ A ][ B ][ C ][ D ]
Destination:      [ A ][ B ][ C ][ D ]
                    ^
                    +--- memcpy corrupts data! memmove copies backwards safely!
```

---

## 3. Usage
```go
// Fast slice clone in Go
func CloneSlice[T any](src []T) []T {
    dst := make([]T, len(src))
    copy(dst, src) // Hardware-accelerated memory copy
    return dst
}
```

---

## 4. Gotchas
- **Shallow Copy Trap:** `copy` duplicates only pointer values. If elements are pointers, mutating an element in the copy modifies the original object.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

