---
title: "Array Delete Fast (Unordered Swap-and-Pop)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Delete Fast (Unordered Swap-and-Pop)

## 1. Definition
**Array Delete Fast (Unordered Swap-and-Pop)** achieves **strict $\mathcal{O}(1)$ removal** when element order does not need to be preserved.
It overwrites the target slot $k$ with the element from the tail $N-1$, and then truncates the length by 1:
$$A[k] = A[N-1]; \quad N = N - 1$$

---

## 2. Mental Model
```text
Delete Index 1 (Item 'B') in O(1):
Step 1: Read Last Item 'E' at Index 4
Step 2: Overwrite Index 1 with 'E':  [ A ][ E ][ C ][ D ][ E ]
Step 3: Truncate Length:             [ A ][ E ][ C ][ D ] (Zero shifts executed!)
```

---

## 3. Usage
```go
// O(1) Unordered Slice Deletion
func DeleteFastUnordered[T any](s []T, index int) []T {
    n := len(s)
    s[index] = s[n-1] // Overwrite target with last element
    var zero T
    s[n-1] = zero     // Prevent memory leak
    return s[:n-1]
}
```

---

## 4. Gotchas
- **Destroys Sorted Invariants:** Cannot be used on sorted arrays, binary search buffers, or monotonic priority queues.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

