---
title: Linked Lists
tags:
  - review
  - computer-science
  - data-structures
  - linked-lists
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Linked Lists (Topologies & Pointer Architectures)

Comprehensive engineering catalog of linear pointer chains, bidirectional graphs, cyclic rings, cache-aligned unrolled chunks, intrusive memory layouts, probabilistic skip lists, and pointer algorithms across 7 specialized domains.

```text
Linked Lists
│
├── 01. Singly Linked List
│   ├── [[Singly Linked List]]
│   ├── [[Singly Linked List Memory Layout and Pointer Chasing]]
│   ├── [[Linked List Sentinel Dummy Node Pattern]]
│   ├── [[Linked List Insert at Head (O(1) Prepend)]]
│   ├── [[Linked List Insert at Tail (O(1) Append with Tail Reference)]]
│   ├── [[Linked List Insert at Index (O(N) Positional Traversal)]]
│   ├── [[Linked List Delete Head and Tail]]
│   ├── [[Linked List Delete Node by Value]]
│   └── [[Linked List Delete Node in O(1) without Prev Pointer]]
│
├── 02. Doubly Linked List
│   ├── [[Doubly Linked List]]
│   └── [[Doubly Linked List and Bidirectional Links]]
│
├── 03. Circular Linked List
│   ├── [[Circular Linked List]]
│   └── [[Circular Linked List and Ring Traversal]]
│
├── 04. Intrusive Linked List
│   ├── [[Intrusive Linked List]]
│   └── [[Intrusive Linked List (Linux Kernel list_head container_of)]]
│
├── 05. Unrolled Linked List
│   ├── [[Unrolled Linked List]]
│   └── [[Unrolled Linked List (Cache Line Chunking)]]
│
├── 06. Skip List
│   ├── [[Skip List]]
│   └── [[Skip List (Probabilistic Multi-Level Express Lanes)]]
│
└── 07. Pointer Algorithms & Techniques
    ├── [[Linked List Reverse (Iterative 3-Pointer Algorithm)]]
    ├── [[Linked List Reverse (Recursive Call Stack Unwinding)]]
    ├── [[Linked List Find Middle (Fast and Slow Pointers)]]
    ├── [[Linked List Remove Nth Node From End (Fast-Slow Window)]]
    ├── [[Linked List Detect Cycle and Entry Point (Floyd Algorithm)]]
    ├── [[Linked List Merge Two Sorted Lists (In-Place Relinking)]]
    └── [[Linked List Palindrome Verification]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Singly Linked List|01. Singly Linked List]]
- [[Singly Linked List]] — Master topology of unidirectional linked chains.
- [[Singly Linked List Memory Layout and Pointer Chasing]] — Disjoint heap node allocation, next pointer offsets, and L1 cache miss penalties.
- [[Linked List Sentinel Dummy Node Pattern]] — Eliminating null boundary edge cases via dummy head and tail sentinel nodes.
- [[Linked List Insert at Head (O(1) Prepend)]] — O(1) constant time insertion at head by updating new node next pointer to current head.
- [[Linked List Insert at Tail (O(1) Append with Tail Reference)]] — O(1) tail appending via direct tail pointer tracking without traversing entire list.
- [[Linked List Insert at Index (O(N) Positional Traversal)]] — Traversing to index k-1 and relinking pointer chains in O(k) time.
- [[Linked List Delete Head and Tail]] — O(1) head removal and O(N) tail removal.
- [[Linked List Delete Node by Value]] — Linear search to locate target node and bypass its reference in O(N).
- [[Linked List Delete Node in O(1) without Prev Pointer]] — O(1) deletion trick by copying value and next pointer from next node.

### 2. 📂 [[Doubly Linked List|02. Doubly Linked List]]
- [[Doubly Linked List]] — Master topology of bidirectional linked chains.
- [[Doubly Linked List and Bidirectional Links]] — Bidirectional next/prev node links, O(1) deletion with node handle, and 16-byte pointer overhead.

### 3. 📂 [[Circular Linked List|03. Circular Linked List]]
- [[Circular Linked List]] — Master topology of cyclic ring buffers.
- [[Circular Linked List and Ring Traversal]] — Continuous cyclic node linking for round-robin CPU scheduling and cyclic ring buffers.

### 4. 📂 [[Intrusive Linked List|04. Intrusive Linked List]]
- [[Intrusive Linked List]] — Master architecture of payload-embedded list heads.
- [[Intrusive Linked List (Linux Kernel list_head container_of)]] — Zero-overhead node embedding via macro container-of calculations.

### 5. 📂 [[Unrolled Linked List|05. Unrolled Linked List]]
- [[Unrolled Linked List]] — Master topology of cache-chunked lists.
- [[Unrolled Linked List (Cache Line Chunking)]] — Array of nodes packed into 64-byte cache lines, mitigating pointer chasing latency.

### 6. 📂 [[Skip List|06. Skip List]]
- [[Skip List]] — Master topology of probabilistic express-lane index layers.
- [[Skip List (Probabilistic Multi-Level Express Lanes)]] — Probabilistic multi-level index layers achieving O(log N) search and insertion.

### 7. 📂 07. Pointer Algorithms & Techniques
- [[Linked List Reverse (Iterative 3-Pointer Algorithm)]] — O(N) in-place pointer reversal using prev, curr, and next tracking pointers.
- [[Linked List Reverse (Recursive Call Stack Unwinding)]] — Reversing linked list via recursive call stack unwinding in O(N) space.
- [[Linked List Find Middle (Fast and Slow Pointers)]] — Tortoise and Hare 2x speed pointer algorithm to locate exact middle in single pass.
- [[Linked List Remove Nth Node From End (Fast-Slow Window)]] — Maintaining gap window of size N between pointers to delete target in single pass.
- [[Linked List Detect Cycle and Entry Point (Floyd Algorithm)]] — Locating cycle meeting point and resetting slow pointer to head to find cycle origin.
- [[Linked List Merge Two Sorted Lists (In-Place Relinking)]] — Merging two sorted linked chains in O(N+M) time and O(1) auxiliary memory.
- [[Linked List Palindrome Verification]] — Finding middle, reversing second half in-place, and comparing halves in O(N) time and O(1) space.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`

