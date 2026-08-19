---
title: "Array Delete at Index (Order-Preserving Shift)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Delete at Index (Order-Preserving Shift)

## 1. Definition
**Array Delete at Index** removes the element at index $k$ while preserving the sequential relative order of all remaining elements.
It shifts elements from index $k+1$ through $N-1$ one slot to the left:
$$A[i] = A[i+1] \quad \forall i \in [k, N-2]$$
Time complexity is $\mathcal{O}(N - k)$, with $\mathcal{O}(N)$ worst-case cost when deleting index 0.

---

## 2. Mental Model
```text
Delete Index 1 (Item 'B') (Len=5):
Original:     [ A ][ B ][ C ][ D ][ E ]
                      <--  <--  <--
Shift Left:   [ A ][ C ][ D ][ E ][ E (stale) ]
Zero & Trunc: [ A ][ C ][ D ][ E ] (Len=4, Cap=5)
```

---

## 3. Usage
```go
// Safe order-preserving deletion with GC memory leak prevention
func DeletePreservingOrder[T any](s []T, index int) []T {
    copy(s[index:], s[index+1:]) // Shift elements left
    var zero T
    s[len(s)-1] = zero           // Zero out vacated slot for GC!
    return s[:len(s)-1]
}
```

---

## 4. Gotchas
- **GC Memory Leak on Deleted Pointers:** Failing to zero out `s[len(s)-1]` leaves a live pointer in the unused capacity tail, preventing GC collection of heavy objects.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

