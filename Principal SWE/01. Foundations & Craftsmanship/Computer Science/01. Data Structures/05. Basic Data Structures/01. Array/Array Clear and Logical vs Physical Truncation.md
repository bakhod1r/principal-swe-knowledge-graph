---
title: "Array Clear and Logical vs Physical Truncation"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Clear and Logical vs Physical Truncation

## 1. Definition
**Array Clear** empties all elements from the array:
- **Logical Clear ($\mathcal{O}(1)$):** Sets `len = 0` while retaining allocated capacity for future appends with zero heap allocations.
- **Physical Clear ($\mathcal{O}(N)$):** Zeroes all memory slots (`clear()`, `memset`) to release live pointers for garbage collection.

---

## 2. Mental Model
```text
Logical Reset (arr[:0]):
Before: ptr ---> [ 10 ][ 20 ][ 30 ][ 40 ] (Len=4, Cap=4)
After:  ptr ---> [ 10 ][ 20 ][ 30 ][ 40 ] (Len=0, Cap=4)
- Fast! Reuses memory, but retains underlying bytes.

Physical Clear (clear(arr)):
After:  ptr ---> [ 0  ][ 0  ][ 0  ][ 0  ] (Len=0, Cap=4)
- Releases all GC object references!
```

---

## 3. Usage
```go
// Go 1.21+ Clear Idioms
func ClearSlice[T any](s []T) []T {
    clear(s)     // Overwrites all elements with zero value (O(N))
    return s[:0] // Resets length to 0 (O(1))
}
```

---

## 4. Gotchas
- **Permanent Heap Bloat:** If a buffer grew to 2 GB during a spike, doing `buf = buf[:0]` retains the 2 GB allocated memory indefinitely. Assign `buf = nil` to return memory to the OS.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

