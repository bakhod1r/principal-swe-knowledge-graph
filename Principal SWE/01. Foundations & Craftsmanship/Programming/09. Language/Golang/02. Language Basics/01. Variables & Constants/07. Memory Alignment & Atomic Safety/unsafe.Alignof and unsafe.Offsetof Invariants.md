---
title: "unsafe.Alignof and unsafe.Offsetof Invariants"
tags:
  - golang
  - unsafe
  - memory-alignment
  - internals
  - principal-swe
parent: "[[Memory Alignment & Atomic Safety]]"
---

# `unsafe.Alignof` and `unsafe.Offsetof` Invariants

## 1. Definition

Go provides low-level reflection and memory introspection primitives in package `unsafe`:

- **`unsafe.Alignof(x)`:** Returns the required alignment (in bytes) of a variable of type $T$. It is guaranteed to be a power of 2 ($1, 2, 4, 8$).
- **`unsafe.Offsetof(s.f)`:** Returns the byte offset within the struct $s$ of the field $f$. It is evaluated at **compile time** as a constant.
- **`unsafe.Sizeof(x)`:** Returns the total memory footprint (in bytes) occupied by a variable of type $T$, including internal padding.

### Alignment Guarantee Formula:
$$\text{Alignof}(T) = \begin{cases} 
\text{Sizeof}(T) & \text{if } \text{Sizeof}(T) \le \text{WordSize} \\
\text{WordSize} & \text{if } \text{Sizeof}(T) > \text{WordSize} \\
\max_{f \in T} \text{Alignof}(f) & \text{if } T \text{ is a struct}
\end{cases}$$

---

## 2. Mental Model

### Alignment Factor Table by Type on 64-Bit System (WordSize = 8)

```text
┌──────────────────────────────┬──────────────┬───────────────────┐
│ Go Type                      │ Size (Bytes) │ Alignment (Bytes) │
├──────────────────────────────┼──────────────┼───────────────────┤
│ bool, byte, uint8, int8      │      1       │         1         │
│ int16, uint16                │      2       │         2         │
│ int32, uint32, float32       │      4       │         4         │
│ int64, uint64, float64       │      8       │         8         │
│ *T, unsafe.Pointer, uintptr  │      8       │         8         │
│ string (Data, Len)           │     16       │         8         │
│ []T (Data, Len, Cap)         │     24       │         8         │
│ interface{} (Type, Data)     │     16       │         8         │
│ struct{} (Empty Struct)      │      0       │         1         │
└──────────────────────────────┴──────────────┴───────────────────┘
```

---

## 3. Usage

### Production Go Engineering: Zero-Cost Compile-Time Alignment Verification

```go
package alignment

import (
	"fmt"
	"unsafe"
)

type NetworkHeader struct {
	Magic     uint16 // Size: 2, Offset: 0
	_         uint16 // Padding: 2 bytes
	Length    uint32 // Size: 4, Offset: 4
	Timestamp uint64 // Size: 8, Offset: 8
}

// Compile-time struct assertion checking exact size and field offsets
const (
	_ = uint(unsafe.Sizeof(NetworkHeader{}) - 16)   // Fails compile if Size != 16
	_ = uint(unsafe.Offsetof(NetworkHeader{}.Timestamp) - 8) // Fails compile if Offset != 8
)

func PrintMemoryLayout(h NetworkHeader) {
	fmt.Printf("Total Size:      %d bytes\n", unsafe.Sizeof(h))
	fmt.Printf("Struct Alignment:%d bytes\n", unsafe.Alignof(h))
	fmt.Printf("Field Magic:     Offset %d (Align %d)\n", unsafe.Offsetof(h.Magic), unsafe.Alignof(h.Magic))
	fmt.Printf("Field Length:    Offset %d (Align %d)\n", unsafe.Offsetof(h.Length), unsafe.Alignof(h.Length))
	fmt.Printf("Field Timestamp: Offset %d (Align %d)\n", unsafe.Offsetof(h.Timestamp), unsafe.Alignof(h.Timestamp))
}
```

---

## 4. Gotchas

- **Non-Portable Pointer Arithmetic:** Assuming field offsets are constant across different compiler versions or target architectures (`GOARCH=arm` vs `GOARCH=amd64`) will lead to memory corruption. Always calculate dynamic offsets via `unsafe.Offsetof` or use `unsafe.Pointer`.
- **Garbage Collector Escape on Pointer Conversion:** Converting an `unsafe.Pointer` to `uintptr` removes the pointer from the Garbage Collector's root set. If memory is moved during a stack growth or GC cycle while stored as a `uintptr`, the `uintptr` will become a stale, invalid address.

---

## 🔗 References
- ⬆️ Parent: [[Memory Alignment & Atomic Safety]]
- 📚 Module: `Language Basics`
