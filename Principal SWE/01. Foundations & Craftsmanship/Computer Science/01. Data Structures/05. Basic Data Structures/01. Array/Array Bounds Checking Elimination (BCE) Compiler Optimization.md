---
title: "Array Bounds Checking Elimination (BCE) Compiler Optimization"
tags:
  - review
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Bounds Checking Elimination (BCE) Compiler Optimization

## 1. Definition
**Bounds Checking Elimination (BCE)** is a compiler optimization where the compiler statically proves that loop indices will never exceed array bounds, omitting runtime `CMP`/`JAE` check instructions for maximum inner loop performance.

---

## 2. Mental Model
```text
Unoptimized Loop:
For i in [0..N]:
  Check(i < len) -> JAE panic_bounds_error  (2 extra instructions per loop!)
  Load A[i]

BCE Optimized Loop:
Compiler proves 0 <= i < len statically:
For i in [0..N]:
  Direct Load A[i] (Zero branching overhead!)
```

---

## 3. Usage
```go
// Triggering Go BCE via slice length assertion
func ProcessAll(s []int) {
    if len(s) >= 4 {
        _ = s[3] // Single boundary hint: proves s[0], s[1], s[2], s[3] are all valid!
        s[0] = 10
        s[1] = 20
        s[2] = 30
        s[3] = 40 // All 4 assignments compile with ZERO bounds checks!
    }
}
```

---

## 4. Gotchas
- **De-optimizing Function Calls:** Calling un-inlined functions inside loop bodies breaks the compiler's range analysis, forcing bounds checks to be re-inserted.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: `Basic Data Structures`

