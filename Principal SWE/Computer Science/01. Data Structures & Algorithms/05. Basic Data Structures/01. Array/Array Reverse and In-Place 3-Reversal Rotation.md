---
title: "Array Reverse and In-Place 3-Reversal Rotation"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Reverse and In-Place 3-Reversal Rotation

## 1. Definition
- **Array Reverse:** Swaps elements from outside inward using two pointers ($l, r$) in $\mathcal{O}(N)$ time and $\mathcal{O}(1)$ space.
- **Array Rotate by $k$ positions:** Achieved in $\mathcal{O}(N)$ time using the **3-Reversal Algorithm**:
  1. $\text{Reverse}(A[0 \dots N-1])$
  2. $\text{Reverse}(A[0 \dots k-1])$
  3. $\text{Reverse}(A[k \dots N-1])$

---

## 2. Mental Model
```text
Rotate [ 1, 2, 3, 4, 5, 6, 7 ] Right by k=3:
Step 1: Reverse Whole Array: [ 7, 6, 5, 4, 3, 2, 1 ]
Step 2: Reverse First k (3):  [ 5, 6, 7, 4, 3, 2, 1 ]
Step 3: Reverse Rest (4..7): [ 5, 6, 7, 1, 2, 3, 4 ] -> Optimal O(1) Space!
```

---

## 3. Usage
```go
// 3-Reversal Array Rotation in Go
func RotateRight(arr []int, k int) {
    n := len(arr)
    if n == 0 { return }
    k %= n
    reverse(arr, 0, n-1)
    reverse(arr, 0, k-1)
    reverse(arr, k, n-1)
}

func reverse(arr []int, l, r int) {
    for l < r {
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }
}
```

---

## 4. Gotchas
- **Modulus Normalization:** Always compute `k = k % len(arr)`. Otherwise $k > N$ causes out-of-bounds index panics.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

