---
title: "Array Memory Layout and Pointer Arithmetic"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Memory Layout and Pointer Arithmetic

## 1. Definition
An **Array** is a linear sequence of homogeneous elements stored in strictly contiguous physical memory.
The location of any element $A[i]$ is calculated via direct pointer arithmetic:
$$\text{Address}(A[i]) = \text{BaseAddress} + i \times \text{sizeof}(T)$$
Where $\text{BaseAddress}$ is the address of the 0-th element and $\text{sizeof}(T)$ is the byte width of the element type.
Because the memory addresses are contiguous, addressing requires zero pointer dereferencing and executes in a single CPU instruction.

---

## 2. Mental Model
```text
Physical RAM Topology (64-bit Architecture):
Base: 0x7FFF0000 | sizeof(uint64) = 8 Bytes

Index:          0              1              2              3
Address:    0x7FFF0000     0x7FFF0008     0x7FFF0010     0x7FFF0018
Memory:   [ 0x0000000A ] [ 0x00000014 ] [ 0x0000001E ] [ 0x00000028 ]
                 |              |              |              |
Offset:        +0 B           +8 B          +16 B          +24 B

x86-64 Assembly Mapping:
MOV RAX, [RDI + RSI * 8]   ; RDI = BaseAddress, RSI = Index i, 8 = Scale
```
- **Zero Indirection:** The CPU translates array indexing directly into hardware SIB (Scale-Index-Base) addressing with 0 memory hops.

---

## 3. Usage
```cpp
// Direct Pointer Arithmetic in C++
#include <iostream>
#include <cstdint>

template <typename T>
inline T* get_element_ptr(T* base, size_t index) {
    // Equivalent to &base[index]
    return reinterpret_cast<T*>(reinterpret_cast<uintptr_t>(base) + index * sizeof(T));
}
```

---

## 4. Gotchas
- **Buffer Over-Read Vulnerabilities:** Indexing beyond array boundaries reads adjacent memory addresses, leaking sensitive tokens, passwords, or triggering `SIGSEGV` segmentation faults.
- **Pointer Alignment Hazards:** Casting arbitrary byte buffers to typed pointers without verifying alignment requirements causes hardware traps on architectures with strict alignment rules (e.g., ARM, SPARC).

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

