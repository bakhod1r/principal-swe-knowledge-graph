---
title: "Array Deduplication on Sorted Buffers"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Deduplication on Sorted Buffers

## 1. Definition
**Array Deduplication on Sorted Buffers** removes all duplicate elements from a sorted array in a single pass in $\mathcal{O}(N)$ time and $\mathcal{O}(1)$ space by comparing adjacent elements.

---

## 2. Mental Model
```text
Sorted Input: [ 1, 1, 2, 2, 2, 3, 4, 4 ]
Read Index:     ^  ^  ^  ^  ^  ^  ^  ^
Write Unique:   1     2        3  4
Result:       [ 1, 2, 3, 4 ] (Len = 4)
```

---

## 3. Usage
```go
// Go 1.21+ slices.Compact style deduplication
func DeduplicateSorted[T comparable](s []T) []T {
    if len(s) == 0 { return s }
    w := 1
    for r := 1; r < len(s); r++ {
        if s[r] != s[w-1] {
            s[w] = s[r]
            w++
        }
    }
    return s[:w]
}
```

---

## 4. Gotchas
- **Requires Sorted Pre-Condition:** Calling this algorithm on an unsorted array fails to remove non-adjacent duplicates.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]
- 🎓 Root: [[Principal SWE]]
