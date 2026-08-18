---
title: Linear Structures
tags:
  - computer-science
  - algorithms
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Linear Structures

Arrays, dynamic vectors, circular ring buffers, singly/doubly/unrolled linked lists, monotonic stacks, and deques.

```text
Linear Structures
│
├── [[Dynamic Array Contiguity and Growth Policy]]
├── [[Singly and Doubly Linked List Pointer Layouts]]
├── [[Unrolled Linked List and Memory Density]]
├── [[Circular Ring Buffer and Lock-Free Queues]]
└── [[Monotonic Stack and Monotonic Deque]]
```

---

## 🗂️ Topics

- [[Dynamic Array Contiguity and Growth Policy]] — Contiguous memory allocation, geometric reallocation factors, and cache prefetching.
- [[Singly and Doubly Linked List Pointer Layouts]] — Node pointer chasing, intrusive list designs (Linux kernel list_head), and cache misses.
- [[Unrolled Linked List and Memory Density]] — Combining arrays within list nodes to maximize cache line utilization.
- [[Circular Ring Buffer and Lock-Free Queues]] — Fixed-size power-of-two ring buffers with atomic head/tail masking for high-throughput IPC.
- [[Monotonic Stack and Monotonic Deque]] — Maintaining sorted invariant stacks to solve Next Greater Element and Sliding Window Maximum in O(N).

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 🎓 Root: [[Principal SWE]]
