---
title: Linked Lists (Basic Data Structures)
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Linked Lists (Basic Data Structures)

Node pointer chasing, intrusive lists, sentinel nodes, reversals, fast/slow pointers, and cycle detection.

```text
Linked Lists (Basic Data Structures)
│
├── [[Singly Linked List Memory Layout and Pointer Chasing]]
├── [[Doubly Linked List and Bidirectional Links]]
├── [[Circular Linked List and Ring Traversal]]
├── [[Linked List Insert at Head (O(1) Prepend)]]
├── [[Linked List Insert at Tail (O(1) Append with Tail Reference)]]
├── [[Linked List Insert at Index (O(N) Positional Traversal)]]
├── [[Linked List Delete Head and Tail]]
├── [[Linked List Delete Node by Value]]
├── [[Linked List Delete Node in O(1) without Prev Pointer]]
├── [[Linked List Sentinel Dummy Node Pattern]]
├── [[Linked List Reverse (Iterative 3-Pointer Algorithm)]]
├── [[Linked List Reverse (Recursive Call Stack Unwinding)]]
├── [[Linked List Find Middle (Fast and Slow Pointers)]]
├── [[Linked List Detect Cycle and Entry Point (Floyd Algorithm)]]
├── [[Linked List Merge Two Sorted Lists (In-Place Relinking)]]
├── [[Linked List Remove Nth Node From End (Fast-Slow Window)]]
├── [[Linked List Palindrome Verification]]
├── [[Unrolled Linked List (Cache Line Chunking)]]
├── [[Intrusive Linked List (Linux Kernel list_head container_of)]]
└── [[Skip List (Probabilistic Multi-Level Express Lanes)]]
```

---

## 🗂️ Operations & Topics

- [[Singly Linked List Memory Layout and Pointer Chasing]] — Disjoint heap node allocation, next pointer offsets, and L1 cache miss penalties.
- [[Doubly Linked List and Bidirectional Links]] — Bidirectional next/prev node links, O(1) deletion with node handle, and 16-byte pointer overhead.
- [[Circular Linked List and Ring Traversal]] — Continuous cyclic node linking for round-robin CPU scheduling and cyclic ring buffers.
- [[Linked List Insert at Head (O(1) Prepend)]] — O(1) constant time insertion at head by updating new node next pointer to current head.
- [[Linked List Insert at Tail (O(1) Append with Tail Reference)]] — O(1) tail appending via direct tail pointer tracking without traversing entire list.
- [[Linked List Insert at Index (O(N) Positional Traversal)]] — Traversing to index k-1 and relinking pointer chains in O(k) time.
- [[Linked List Delete Head and Tail]] — O(1) head removal and O(N) tail removal (or O(1) on doubly linked lists).
- [[Linked List Delete Node by Value]] — Linear search to locate target node and bypass its reference in O(N).
- [[Linked List Delete Node in O(1) without Prev Pointer]] — O(1) deletion trick by copying value and next pointer from next node.
- [[Linked List Sentinel Dummy Node Pattern]] — Eliminating null boundary edge cases via dummy head and tail sentinel nodes.
- [[Linked List Reverse (Iterative 3-Pointer Algorithm)]] — O(N) in-place pointer reversal using prev, curr, and next tracking pointers.
- [[Linked List Reverse (Recursive Call Stack Unwinding)]] — Reversing linked list via recursive call stack unwinding in O(N) space.
- [[Linked List Find Middle (Fast and Slow Pointers)]] — Tortoise and Hare 2x speed pointer algorithm to locate exact middle in single pass.
- [[Linked List Detect Cycle and Entry Point (Floyd Algorithm)]] — Locating cycle meeting point and resetting slow pointer to head to find cycle origin.
- [[Linked List Merge Two Sorted Lists (In-Place Relinking)]] — Merging two sorted linked chains in O(N+M) time and O(1) auxiliary memory.
- [[Linked List Remove Nth Node From End (Fast-Slow Window)]] — Maintaining gap window of size N between pointers to delete target in single pass.
- [[Linked List Palindrome Verification]] — Finding middle, reversing second half in-place, and comparing halves in O(N) time and O(1) space.
- [[Unrolled Linked List (Cache Line Chunking)]] — Embedding fixed arrays within linked nodes to match 64-byte L1 cache lines.
- [[Intrusive Linked List (Linux Kernel list_head container_of)]] — Zero-allocation node embedding directly inside payload data structures.
- [[Skip List (Probabilistic Multi-Level Express Lanes)]] — Layered forward pointer levels delivering O(log N) concurrent search and insertion.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: [[Data Structures & Algorithms]]

