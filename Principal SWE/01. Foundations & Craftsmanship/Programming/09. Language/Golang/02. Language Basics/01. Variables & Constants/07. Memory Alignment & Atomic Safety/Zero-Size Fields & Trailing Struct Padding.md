---
title: "Zero-Size Fields & Trailing Struct Padding"
tags:
  - golang
  - memory-alignment
  - struct-padding
  - low-level
  - principal-swe
parent: "[[Memory Alignment & Atomic Safety]]"
---

# Zero-Size Fields & Trailing Struct Padding

## 1. Definition

In Go, an empty struct (`struct{}`) or zero-length array (`[0]T`) has a logical size of **0 bytes** (`unsafe.Sizeof(struct{}{}) == 0`).

However, the Go compiler enforces a critical memory layout invariant regarding zero-size fields when placed at the **end** of a struct:

### The Trailing Zero-Size Invariant:
> **Compiler Invariant:** 
> If the final field of a struct is of size 0, the Go compiler **pads the struct with additional trailing bytes** (typically 8 bytes on 64-bit architectures, 4 bytes on 32-bit architectures) so that the struct's size increases.
>
> **The Root Cause:** If the trailing zero-sized field had offset equal to `sizeof(struct)`, taking its memory address `&s.trailingField` would produce a pointer pointing **one byte past the end of the allocated memory block**. Such a pointer can cause the Garbage Collector to misidentify which heap span owns the object, or cause memory safety violations when converting to `uintptr`.

---

## 2. Mental Model

```text
1. Leading or Middle Zero-Size Field (0 Extra Padding):
   type Leading struct {
       z struct{} // Size: 0, Offset: 0
       a int64    // Size: 8, Offset: 0
   }
   ┌─────────────────────────────────┐
   │ a (int64): 8 Bytes              │
   └─────────────────────────────────┘
   Total Size: 8 Bytes ✅ (Zero padding added)

2. Trailing Zero-Size Field (Forces Word-Sized Trailing Padding!):
   type Trailing struct {
       a int64    // Size: 8, Offset: 0
       z struct{} // Size: 0, Offset: 8 (Points past allocated block!)
   }
   ┌─────────────────────────────────┬─────────────────────────────────┐
   │ a (int64): 8 Bytes              │ Trailing Padding: 8 Bytes       │
   └─────────────────────────────────┴─────────────────────────────────┘
   Total Size: 16 Bytes ⚠️ (8 bytes added solely to protect &s.z pointer!)
```

---

## 3. Usage

### Production Go Verification: Measuring Struct Size Inflation

```go
package alignment

import (
	"fmt"
	"unsafe"
)

type LeadingEmpty struct {
	Signal struct{} // 0 bytes at start
	Value  int64    // 8 bytes
}

type TrailingEmpty struct {
	Value  int64    // 8 bytes
	Signal struct{} // 0 bytes at end -> triggers 8 bytes padding!
}

type MiddleEmpty struct {
	Value1 int32    // 4 bytes
	Signal struct{} // 0 bytes
	Value2 int32    // 4 bytes
}

func ExampleZeroSizeFieldPadding() {
	var l LeadingEmpty
	var t TrailingEmpty
	var m MiddleEmpty

	fmt.Printf("Leading Empty Struct Size:  %d Bytes\n", unsafe.Sizeof(l))  // 8 Bytes
	fmt.Printf("Trailing Empty Struct Size: %d Bytes\n", unsafe.Sizeof(t))  // 16 Bytes (50% inflation!)
	fmt.Printf("Middle Empty Struct Size:   %d Bytes\n", unsafe.Sizeof(m))  // 8 Bytes

	// Pointer arithmetic inspection
	ptrT := uintptr(unsafe.Pointer(&t))
	ptrField := uintptr(unsafe.Pointer(&t.Signal))
	fmt.Printf("Trailing struct base: 0x%x, Signal field ptr: 0x%x (Offset: %d)\n",
		ptrT, ptrField, ptrField-ptrT)
}
```

---

## 4. Gotchas

- **High-Density Collections Inflation:** If you create an array or slice of millions of structs containing a trailing `struct{}` (e.g. `[]TrailingEmpty`), your heap allocation will be **$2\times$ larger** than necessary ($16\text{ MB}$ vs $8\text{ MB}$). Always place zero-sized fields (`struct{}`, `[0]byte`) at the **beginning** of the struct.
- **Embedded Structs with Trailing Empty Struct:** If an embedded struct has a trailing empty struct, embedding it will propagate the padding penalty to the outer enclosing struct.

---

## 🔗 References
- ⬆️ Parent: [[Memory Alignment & Atomic Safety]]
- 📚 Module: `Language Basics`
