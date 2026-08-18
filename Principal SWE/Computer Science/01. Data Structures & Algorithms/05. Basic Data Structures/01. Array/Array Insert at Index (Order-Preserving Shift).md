---
title: "Array Insert at Index (Order-Preserving Shift)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Insert at Index (Order-Preserving Shift)

## 1. Definition
**Array Insert at Index** inserts an element at a specified offset $k$, shifting all elements from index $k$ through $N-1$ one position to the right.
It executes in $\mathcal{O}(N - k)$ time, requiring $\mathcal{O}(N)$ worst-case data movement when inserting at index 0.

---

## 2. Mental Model
```text
Insert 99 at Index 2:
Original (Len=5):  [ 10 ][ 20 ][ 30 ][ 40 ][ 50 ][ _ ]
                                   \     \     \
Shift Right (SIMD):               [ 30 ][ 40 ][ 50 ]
Target Slot Opened: [ 10 ][ 20 ][ 99 ][ 30 ][ 40 ][ 50 ]
```

---

## 3. Usage
```go
// In-place slice insertion with memory move
func InsertAt[T any](slice []T, index int, value T) []T {
    var zero T
    slice = append(slice, zero)         // Extend length by 1
    copy(slice[index+1:], slice[index:]) // Hardware-accelerated shift right
    slice[index] = value
    return slice
}
```

---

## 4. Gotchas
- **Repeated Head Insertions:** Inserting $N$ elements at index 0 in a loop runs in $\mathcal{O}(N^2)$ quadratic time. Use a `Deque` or reverse append for front-heavy insertions.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

