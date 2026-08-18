---
title: "Array Filter and In-Place Two-Pointer Compaction"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Filter and In-Place Two-Pointer Compaction

## 1. Definition
**Array Filter (In-Place)** filters elements matching a predicate without allocating secondary memory.
It utilizes a **Read Pointer (Fast)** that scans every element and a **Write Pointer (Slow)** that writes only matching elements, executing in $\mathcal{O}(N)$ time and **$\mathcal{O}(1)$ auxiliary space**.

---

## 2. Mental Model
```text
Filter Even Numbers in Place:
Input:  [ 1, 2, 3, 4, 5, 6 ]
Read:     ^  ^  ^  ^  ^  ^
Write:       [2]      [4]      [6]
Result: [ 2, 4, 6 ] (Truncate length to write_index = 3)
```

---

## 3. Usage
```go
// Zero-Allocation In-Place Slice Filtering
func FilterInPlace(s []int, keep func(int) bool) []int {
    w := 0
    for _, x := range s {
        if keep(x) {
            s[w] = x
            w++
        }
    }
    return s[:w] // O(1) truncation
}
```

---

## 4. Gotchas
- **Mutates Original Data:** In-place compaction permanently overwrites the original slice elements.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]
- 🎓 Root: [[Principal SWE]]
