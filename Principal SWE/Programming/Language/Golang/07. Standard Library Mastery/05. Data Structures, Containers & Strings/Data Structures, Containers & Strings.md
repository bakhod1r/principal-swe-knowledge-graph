---
title: Data Structures, Containers & Strings
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# Data Structures, Containers & Strings

container/heap priority queue, container/list, container/ring, strings.Builder, bytes.Buffer, fmt internals, and strconv.

```text
Data Structures, Containers & Strings
│
├── [[container-heap Priority Queue & Custom Sorters]]
├── [[container-list Doubly Linked List Mechanics]]
├── [[container-ring Circular Ring Buffer Mechanics]]
├── [[strings and bytes High-Performance Manipulation]]
├── [[strings.Builder vs bytes.Buffer Zero-Alloc Comparison]]
├── [[fmt Package Formatter Internals & Reflection Cost]]
└── [[strconv Package Fast Numeric Conversion]]
```

---

## 🗂️ Topics

- [[container-heap Priority Queue & Custom Sorters]] — Implementing heap.Interface (Len, Less, Swap, Push, Pop) for min/max priority queues.
- [[container-list Doubly Linked List Mechanics]] — Doubly linked list operations: PushBack, MoveToFront, Remove, and pointer memory overhead.
- [[container-ring Circular Ring Buffer Mechanics]] — Fixed-capacity circular linked list for round-robin rotation and fixed window buffers.
- [[strings and bytes High-Performance Manipulation]] — Zero-allocation string searching, splitting, trimming, and bytes slice manipulation algorithms.
- [[strings.Builder vs bytes.Buffer Zero-Alloc Comparison]] — Comparing heap allocation profiles of strings.Builder (String() zero-copy) vs bytes.Buffer.
- [[fmt Package Formatter Internals & Reflection Cost]] — How fmt.Sprintf uses reflection and interface boxing, and writing high-speed custom fmt.Formatter.
- [[strconv Package Fast Numeric Conversion]] — High-performance string-to-number parsing: strconv.ParseInt, strconv.FormatFloat, and zero-alloc byte appenders.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
