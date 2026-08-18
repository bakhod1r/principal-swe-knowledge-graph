---
title: "Array Search (Linear vs Branchless Binary Search)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Search (Linear vs Branchless Binary Search)

## 1. Definition
**Array Search** locates the index of a target value:
- **Linear Search:** Scans elements sequentially in $\mathcal{O}(N)$ time.
- **Branchless Binary Search:** Halves the sorted search space in $\mathcal{O}(\log N)$ time using conditional move instructions (`CMOV`), completely eliminating CPU branch mispredictions.

---

## 2. Mental Model
```text
Standard Binary Search Branch Penalty:
if (arr[mid] < target) -> 50% CPU branch misprediction on random queries!

Branchless Binary Search (Paul-Viriyakulkij Algorithm):
base += (arr[base + half] < target) ? half : 0;
// Compiles to CMOV / conditional addition: ZERO branch stalls!
```

---

## 3. Usage
```cpp
// Branchless Binary Search in C++
int branchless_binary_search(const int* arr, int n, int target) {
    const int* base = arr;
    while (n > 1) {
        int half = n / 2;
        base = (base[half] < target) ? base + half : base;
        n -= half;
    }
    return (*base == target) ? (base - arr) : -1;
}
```

---

## 4. Gotchas
- **Midpoint Integer Overflow:** Never write `mid = (low + high) / 2`. Always use `mid = low + (high - low) / 2` or `(uint(low + high)) >> 1` to prevent signed 32-bit integer overflow.

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]
- 🎓 Root: [[Principal SWE]]
