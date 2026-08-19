---
title: Array
tags:
  - golang
  - arrays
  - principal-swe
parent: "[[Composite Types]]"
---

# Array

Fixed-length arrays, contiguous memory layout, value semantics, and comparison.

```text
Array
│
├── [[Fixed Length & Contiguous Memory]]
├── [[Array Memory Layout]]
├── [[Array Pass-by-Value Semantics]]
├── [[Multi-Dimensional Arrays]]
├── [[Array Comparison (==)]]
├── [[Array Slicing to Slice Header]]
└── [[Zero-Length Arrays ([0]T)]]
```

---

## 🗂️ Topics

- [[Fixed Length & Contiguous Memory]] — Fixed-size contiguous memory sequences where length is part of the type ([N]T).
- [[Array Memory Layout]] — Zero-overhead sequential layout on stack or heap with zero padding between elements.
- [[Array Pass-by-Value Semantics]] — Copying the entire memory contents of the array upon function call or assignment.
- [[Multi-Dimensional Arrays]] — Matrices and multi-dimensional grids ([M][N]T) in row-major contiguous memory.
- [[Array Comparison (==)]] — Direct value equality comparison (arr1 == arr2) when element types are comparable.
- [[Array Slicing to Slice Header]] — Creating a slice header referencing an array (arr[:] and &arr).
- [[Zero-Length Arrays ([0]T)]] — Zero-byte array types and their memory allocation characteristics.


## 🗂️ Contents

- [[Array Comparison (==)]]
- [[Array Memory Layout]]
- [[Array Pass-by-Value Semantics]]
- [[Array Slicing to Slice Header]]
- [[Fixed Length & Contiguous Memory]]
- [[Multi-Dimensional Arrays]]
- [[Zero-Length Arrays ([0]T)]]

---

## 🔗 References
- ⬆️ Parent: [[Composite Types]]

