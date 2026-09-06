---
title: Chunked Block Deque
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Deque]]"
---

# 📦 Chunked Block Deque

std::deque architecture: array of fixed-size chunks.

```text
Chunked Block Deque
│
├── [[Chunked Block Array Deque Architecture (std::deque)]]
├── [[std::deque Fixed-Size Buffer Map and Iterator Invalidation]]
└── [[Chunked Block Deque Random Access Offset Math (Map Index + Buffer Offset)]]
```

---

## 🗂️ Topics & Implementations

- [[Chunked Block Array Deque Architecture (std::deque)]] — Central map of pointers to fixed-size (e.g. 512B) memory blocks.
- [[std::deque Fixed-Size Buffer Map and Iterator Invalidation]] — Why std::deque insertions do not invalidate pointers to existing elements.
- [[Chunked Block Deque Random Access Offset Math (Map Index + Buffer Offset)]] — Two-tier pointer indirection arithmetic achieving O(1) random access: index / block_size.

---

## 🔗 References
- ⬆️ Parent: [[Deque]]
- 📚 Module: `Data Structures`
