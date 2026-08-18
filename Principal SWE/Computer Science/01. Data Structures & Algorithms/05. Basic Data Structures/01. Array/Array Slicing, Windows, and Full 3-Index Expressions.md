---
title: "Array Slicing, Windows, and Full 3-Index Expressions"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Slicing, Windows, and Full 3-Index Expressions

## 1. Definition
**Array Slicing** constructs a lightweight view (`struct { Data *T; Len int; Cap int }`) over a contiguous sub-range of an existing array in strict $\mathcal{O}(1)$ time without memory copies.
**3-Index Slicing** (`s[low:high:max]`) limits the capacity of the resulting slice to `max - low`, preventing append operations from modifying the parent array.

---

## 2. Mental Model
```text
Parent Array (Cap = 8):
Index:    0    1    2    3    4    5    6    7
Data:   [ 10 ][ 20 ][ 30 ][ 40 ][ 50 ][ 60 ][ 70 ][ 80 ]
                     ^              ^          ^
                     |              |          |
Child Slice:         +--- Low = 2   |          |
                          High = 4 -+          |
                          Max = 5 -------------+
Len = High - Low = 2, Cap = Max - Low = 3
```

---

## 3. Usage
```go
// 3-Index Slicing to Protect Parent Buffers
func SafeSubSlice(parent []byte) []byte {
    // Limits child slice capacity so append cannot overwrite parent[5..]
    return parent[2:5:5] 
}
```

---

## 4. Gotchas
- **Sub-Slice Memory Leaks:** Retaining a 2-byte slice from a 500 MB array keeps the entire 500 MB allocated in memory because the GC cannot free partial backing buffers.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

