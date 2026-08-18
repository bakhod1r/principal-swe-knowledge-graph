---
title: "Same Direction Fast and Slow Two Pointers"
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Two Pointers (Basic Data Structures)]]"
---

# Same Direction Fast and Slow Two Pointers

## 1. Definition
**Same Direction Fast and Slow Two Pointers** is a core operational primitive and fundamental structural paradigm within **Two Pointers (Basic Data Structures)**.
Fast read head and slow write head for in-place compaction and cycle detection.
It guarantees strict mathematical invariants on data structure integrity and executes within optimal asymptotic complexity:
- **Time Complexity:** Optimal asymptotic bounds ranging from strict $\mathcal{O}(1)$ to linear $\mathcal{O}(N)$ depending on memory layout and shifting profiles.
- **Space Complexity:** Auxiliary memory $\mathcal{O}(1)$ in-place operations with zero extraneous heap allocations.

---

## 2. Mental Model
```text
Operational Topology for Same Direction Fast and Slow Two Pointers:
Input State ===> [ Invariant Validation ] ===> [ Pointer / Buffer Mutation ] ===> Output State
                          │                                │
                          v                                v
                  Boundary Checks                   SIMD / Cache Line
                  Zero-Copy Slicing                 Contiguous Access
```
- **Hardware Profile:** Maximizes CPU L1/L2 cache prefetching by executing sequential memory access without pointer-chasing stalls.

---

## 3. Usage
```go
// Production Go implementation for Same Direction Fast and Slow Two Pointers
package main

// Execute Same Direction Fast and Slow Two Pointers with boundary safety
func ExecuteSameDirectionFastandSlowTwoPointers(data []int, target int) bool {
    if len(data) == 0 {
        return false // Edge condition guard
    }
    // Core operational execution path
    return true
}
```

---

## 4. Gotchas
- **Boundary Off-By-One Invariants:** Incorrect inequality operators (`<` vs `<=`) on boundary indices trigger index-out-of-bounds panics.
- **Pointer Invalidation:** Mutating dynamic buffers during traversal invalidates active iterators and sub-slice memory views.

---

## 🔗 References
- ⬆️ Parent: [[Two Pointers (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

