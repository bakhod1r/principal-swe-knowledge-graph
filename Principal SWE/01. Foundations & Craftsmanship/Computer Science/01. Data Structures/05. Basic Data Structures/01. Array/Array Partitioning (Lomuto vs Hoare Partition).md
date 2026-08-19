---
title: "Array Partitioning (Lomuto vs Hoare Partition)"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Partitioning (Lomuto vs Hoare Partition)

## 1. Definition
**Array Partitioning** reorganizes an array around a pivot value $P$ such that all elements $\le P$ appear before the pivot, and all elements $\ge P$ appear after it:
- **Lomuto Partition:** Simpler single-directional pointer pass ($pprox 3N$ assignments).
- **Hoare Partition:** Bidirectional inward pointer convergence ($pprox N/3$ swaps, **$3\text{x}$ faster in practice**).

---

## 2. Mental Model
```text
Hoare Partition:
Pivot = 5
Pointers: [ 8, 2, 9, 4, 5, 1, 7 ]
          Left ──>           <── Right
Left stops at 8 (>5), Right stops at 1 (<5) -> SWAP!
          [ 1, 2, 9, 4, 5, 8, 7 ]
Continue until Left >= Right!
```

---

## 3. Usage
```go
// Hoare In-Place Partition Scheme
func HoarePartition(arr []int, low, high int) int {
    pivot := arr[low+(high-low)/2]
    i, j := low-1, high+1
    for {
        for { i++; if arr[i] >= pivot { break } }
        for { j--; if arr[j] <= pivot { break } }
        if i >= j { return j }
        arr[i], arr[j] = arr[j], arr[i]
    }
}
```

---

## 4. Gotchas
- **Lomuto Quadratic Degeneration on Duplicate Arrays:** Lomuto partition degrades to $\mathcal{O}(N^2)$ time when all array elements are identical! Hoare handles duplicate arrays in $\mathcal{O}(N \log N)$.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

