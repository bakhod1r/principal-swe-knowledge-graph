---
title: "Array Pre-Allocation and Capacity Hints"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Pre-Allocation and Capacity Hints

## 1. Definition
**Array Pre-Allocation** reserves the exact expected capacity $\text{cap}$ at allocation time (`make([]T, 0, expectedSize)`, `std::vector::reserve(N)`).
It completely eliminates all intermediate memory reallocations, memory copying, and allocator heap fragmentation, reducing ingestion time by **$3\text{x}\text{--}5\text{x}$**.

---

## 2. Mental Model
```text
Without Pre-allocation (Appending 1000 items):
Alloc(1) -> Alloc(2) -> Alloc(4) -> Alloc(8) ... -> Alloc(1024)
Total Memory Copies: 1 + 2 + 4 + 8 + ... + 512 = 1,023 elements copied!
Total Malloc Calls: 11 distinct heap allocations!

With Pre-allocation (reserve 1000):
Alloc(1000) once!
Total Memory Copies: 0 elements copied!
Total Malloc Calls: 1 heap allocation!
```

---

## 3. Usage
```go
// Pre-allocating slice capacity in Go
func ProcessStream(stream <-chan int, batchSize int) []int {
    // Len = 0, Cap = batchSize -> 0 reallocations during append loop!
    result := make([]int, 0, batchSize)
    for item := range stream {
        result = append(result, item) // Always O(1) direct write!
    }
    return result
}
```

---

## 4. Gotchas
- **Make Length vs Capacity Mistake:** Writing `make([]int, 100)` creates 100 zeroed elements. Calling `append` subsequently appends at index 100, resulting in 200 elements! Use `make([]int, 0, 100)` when using `append`.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

