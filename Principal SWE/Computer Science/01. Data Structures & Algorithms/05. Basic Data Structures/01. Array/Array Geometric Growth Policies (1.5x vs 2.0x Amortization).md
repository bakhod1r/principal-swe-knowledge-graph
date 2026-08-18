---
title: "Array Geometric Growth Policies (1.5x vs 2.0x Amortization)"
tags:
  - computer-science
  - data-structures
  - arrays
  - basic-data-structures
  - principal-swe
parent: "[[Array (Basic Data Structures)]]"
---

# Array Geometric Growth Policies (1.5x vs 2.0x Amortization)

## 1. Definition
**Geometric Capacity Growth** guarantees that dynamic array insertions achieve $\mathcal{O}(1)$ amortized cost.
If an array starts at capacity 1 and doubles ($g=2$) upon overflowing, the total cost to append $N = 2^k$ elements is:
$$\text{Total Cost} = N + \sum_{j=0}^{k-1} 2^j = N + (2^k - 1) = 2N - 1 < 2N$$
The amortized cost per single append is $\frac{2N-1}{N} = \mathcal{O}(1)$.

---

## 2. Mental Model
```text
Banker's Accounting Method:
Every normal append pays $3 credits:
- $1 pays for the current insert
- $2 stored on the element to pay for its future copy when the array doubles!

2.0x vs 1.5x Multiplier Memory Fragmentation:
Growth 2.0x: 1, 2, 4, 8, 16, 32 -> Sum(1..16) = 31 < 32 (Can NEVER reuse freed chunks!)
Growth 1.5x: 1, 2, 3, 5, 8, 12, 18, 27 -> Can REUSE previously freed memory blocks!
```

---

## 3. Usage
```cpp
// Folly fbvector (Facebook) uses 1.5x growth factor to maximize jemalloc page reuse!
```

---

## 4. Gotchas
- **Arithmetic Growth Pitfall:** Growing capacity by a fixed constant (e.g. `cap += 1000`) degrades append time to $\mathcal{O}(N^2)$ quadratic cost! Growth MUST be multiplicative (geometric).

---

## 🔗 References
- ⬆️ Parent: [[Array (Basic Data Structures)]]
- 📚 Module: [[Basic Data Structures]]

