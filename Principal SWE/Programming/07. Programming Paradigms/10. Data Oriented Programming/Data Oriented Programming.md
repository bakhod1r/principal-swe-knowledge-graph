---
title: Data Oriented Programming
tags:
  - programming
  - paradigms
  - principal-swe
parent: "[[Programming Paradigms]]"
---

# Data Oriented Programming

Separating code from immutable data (Clojure-DOP), generic map representation, cache-conscious SOA vs AOS layouts, and SIMD.

```text
Data Oriented Programming
│
├── [[Separating Code from Immutable Data (Clojure-DOP)]]
├── [[Data Representation with Generic Maps, Records, and Schema Validation]]
├── [[Cache-Conscious Data Layout (Struct of Arrays vs Array of Structs)]]
└── [[Performance Sympathy: SIMD Vectorization on DOD Buffers]]
```

---

## 🗂️ Topics

- [[Separating Code from Immutable Data (Clojure-DOP)]] — The 4 principles of Data-Oriented Programming: separate code from data, represent data with generic maps, immutable data, separate schema.
- [[Data Representation with Generic Maps, Records, and Schema Validation]] — Decoupling domain state from rigid classes using plain maps and decoupled schema validators.
- [[Cache-Conscious Data Layout (Struct of Arrays vs Array of Structs)]] — Organizing memory for maximum L1/L2 CPU cache utilization and hardware prefetcher efficiency.
- [[Performance Sympathy: SIMD Vectorization on DOD Buffers]] — Executing parallel vector operations over contiguous memory blocks in data-oriented designs.

---

## 🔗 References
- ⬆️ Parent: [[Programming Paradigms]]
- 🎓 Root: [[Principal SWE]]
