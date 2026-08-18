---
title: "Array Append and Dynamic Capacity Allocation"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Append and Dynamic Capacity Allocation

## 1. Definition
**Array Append** adds an element to the logical end of a dynamic array.
When $\text{len} < \text{cap}$, it writes the value to the free memory slot at $\text{len}$ in $\mathcal{O}(1)$ time.
When $\text{len} == \text{cap}$, it allocates a new contiguous buffer of size $\text{cap} \times g$, copies all $N$ existing elements, deallocates the old buffer, and appends the element, achieving **$\mathcal{O}(1)$ amortized time**.

---

## 2. Mental Model
```text
Append Sequence (Initial Len=3, Cap=3):
Step 1: Array is Full [ 10 ][ 20 ][ 30 ]
Step 2: Append(40) Triggers Reallocation
        Allocate New Buffer (Cap = 6):
        [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]
Step 3: Copy Elements (Memcpy):
        [ 10 ][ 20 ][ 30 ][ _ ][ _ ][ _ ]
Step 4: Write New Value & Free Old:
        [ 10 ][ 20 ][ 30 ][ 40 ][ _ ][ _ ] (Len=4, Cap=6)
```

---

## 3. Usage
```go
// Go 1.18+ Smooth Growth Algorithm
func calculateNewCap(oldCap, neededCap int) int {
    doublecap := oldCap + oldCap
    if neededCap > doublecap {
        return neededCap
    }
    const threshold = 256
    if oldCap < threshold {
        return doublecap // 2x growth for small slices
    }
    // Smooth transition towards 1.25x + 192 for large buffers
    newcap := oldCap
    for 0 < newcap && newcap < neededCap {
        newcap += (newcap + 3*threshold) / 4
    }
    return newcap
}
```

---

## 4. Gotchas
- **Heap Spikes on Large Arrays:** Appending to a 10 GB array requires a temporary 30 GB memory allocation during copying, triggering OOM panics in memory-constrained containers.
- **Stale Shared Sub-Slice Mutation:** Appending to a sub-slice when $\text{len} < \text{cap}$ mutates the underlying parent array elements in place!

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]
- 🎓 Root: [[Principal SWE]]
