---
title: Primitive Runtime Helpers (slice.go, map.go, iface.go)
tags:
  - golang
  - runtime
  - principal-swe
parent: "[[Runtime & Internals]]"
---

# Primitive Runtime Helpers (slice.go, map.go, iface.go)

growslice reallocation math, legacy hmap vs Swiss Table map engine, string conversions, interface construction, and finalizers.

```text
Primitive Runtime Helpers (slice.go, map.go, iface.go)
│
├── [[Slice Allocation & Dynamic Growth Math (growslice in slice.go)]]
├── [[Legacy Map (hmap & bmap) Runtime Implementation (map.go)]]
├── [[Swiss Table Map Runtime Engine (map_swiss.go Go 1.24+)]]
├── [[String Concatenation & Memory Conversion (string.go)]]
├── [[Interface Construction & Assertion Helpers (convT & assertI2I)]]
└── [[Finalizer Queue & Execution Lifecycle (runtime.SetFinalizer)]]
```

---

## 🗂️ Topics

- [[Slice Allocation & Dynamic Growth Math (growslice in slice.go)]] — Slice growth threshold math (2x below 256 elements, 1.25x + 192 above 256) and memory class rounding.
- [[Legacy Map (hmap & bmap) Runtime Implementation (map.go)]] — Bucket array allocation, tophash comparison, 6.5 load factor evacuation, and overflow buckets.
- [[Swiss Table Map Runtime Engine (map_swiss.go Go 1.24+)]] — Go 1.24+ Swiss Table layout: control bytes, 16-element groups, SIMD vector matching, and slot arrays.
- [[String Concatenation & Memory Conversion (string.go)]] — concatstrings algorithm, small string optimizations, and runtime byte-to-string casting.
- [[Interface Construction & Assertion Helpers (convT & assertI2I)]] — convT scalar boxing, type assertion checks (assertI2I), and dynamic itab caching (itabAdd).
- [[Finalizer Queue & Execution Lifecycle (runtime.SetFinalizer)]] — Special finalizer block allocation (fintab), special records, and the dedicated finq execution goroutine.

---

## 🔗 References
- ⬆️ Parent: [[Runtime & Internals]]

