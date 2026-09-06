---
title: Linked Lists
tags:
  - review
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Linked Lists

Sequential pointer-linked nodes, bidirectional handles, cyclic loops, Linux kernel intrusive structures, unrolled cache chunks, and multi-level probabilistic skip lists.

```text
Linked Lists
│
├── 01. Singly Linked List
│   ├── [[Singly Linked List]]
│   ├── [[Singly Linked List Node Memory Layout and Heap Pointers]]
│   ├── [[Singly Linked List Dummy Sentinel Head Pattern]]
│   ├── [[Singly Linked List Insert at Head (O(1) Prepend)]]
│   ├── [[Singly Linked List Insert at Tail (O(1) with Tail Pointer)]]
│   ├── [[Singly Linked List Insert at Index (Order-Preserving)]]
│   ├── [[Singly Linked List Delete Head and Tail Elements]]
│   ├── [[Singly Linked List Delete by Value (First Occurrence)]]
│   ├── [[Singly Linked List Delete Node with O(1) Pointer Handle]]
│   ├── [[Singly Linked List Memory Overhead and Cache Miss Economics]]
│   └── [[Forward Iterator Mechanics and Traversal Bounds]]
││
├── 02. Doubly Linked List
│   ├── [[Doubly Linked List]]
│   ├── [[Doubly Linked List and Bidirectional Links]]
│   ├── [[Doubly Linked List Sentinel Nodes (Circular Dummy Invariant)]]
│   ├── [[Doubly Linked List O(1) Node Splice and Transfer]]
│   └── [[LRU Cache Eviction Policy using Doubly Linked List and Hash Map]]
││
├── 03. Circular Linked List
│   ├── [[Circular Linked List]]
│   ├── [[Circular Linked List and Cyclic Looping Mechanics]]
│   ├── [[Circular Linked List Round-Robin Scheduler Implementation]]
│   ├── [[Josephus Problem and Cyclic Elimination Mechanics]]
│   └── [[Splitting and Concatenating Circular Linked Lists]]
││
├── 04. Intrusive Linked List
│   ├── [[Intrusive Linked List]]
│   ├── [[Intrusive Linked List Architecture (Linux Kernel list_head)]]
│   ├── [[Linux Kernel list_head Topology and container_of Macro]]
│   ├── [[Intrusive vs Non-Intrusive Zero-Allocation Performance Profiles]]
│   └── [[Embedding Multiple Intrusive Lists in a Single Struct]]
││
├── 05. Unrolled Linked List
│   ├── [[Unrolled Linked List]]
│   ├── [[Unrolled Linked List Hybrid Cache Topology]]
│   ├── [[Unrolled Node Capacity and Cache Line Alignment (64-Byte Matching)]]
│   ├── [[Unrolled List Element Insertion and Node Splitting Mechanics]]
│   └── [[Defragmentation and Node Merging in Unrolled Lists]]
││
├── 06. Skip List
│   ├── [[Skip List]]
│   ├── [[Skip List Probabilistic Express Towers and O(log N) Search]]
│   ├── [[Skip List Probabilistic Geometric Level Distribution]]
│   ├── [[Skip List Search, Insertion, and Predecessor Array Traversal]]
│   └── [[Lock-Free Concurrent Skip List (Java ConcurrentSkipListMap Mechanics)]]
││
└── 07. Pointer Algorithms & Techniques
    ├── [[Pointer Algorithms & Techniques]]
    ├── [[Reverse Linked List (Iterative 3-Pointer vs Recursive)]]
    ├── [[Middle of Linked List (Fast and Slow Pointer Floyd Pattern)]]
    ├── [[N-th Node from End (Two-Pointer Fixed-Gap Sliding Window)]]
    ├── [[Detect and Break Loop in Linked List (Floyd Cycle Finding)]]
    ├── [[Merge Two Sorted Linked Lists (In-Place Pointer Splicing)]]
    ├── [[Palindrome Linked List Verification (In-Place Half Reverse)]]
    ├── [[Intersection of Two Linked Lists (Difference-in-Length vs Cycle)]]
    └── [[Merge K Sorted Linked Lists (Min-Heap vs Divide-and-Conquer)]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Singly Linked List|01. Singly Linked List]]
- [[Singly Linked List]] — Master topology: Linear single-pointer sequence, dynamic heap nodes, dummy head sentinel pattern, and fundamental node insertions/deletions.
- [[Singly Linked List Node Memory Layout and Heap Pointers]] — Heap allocation, next pointer memory overhead, and 64-bit address alignment.
- [[Singly Linked List Dummy Sentinel Head Pattern]] — Eliminating edge-case branching for empty lists and head insertions.
- [[Singly Linked List Insert at Head (O(1) Prepend)]] — O(1) pointer redirection prepend operation without list traversal.
- [[Singly Linked List Insert at Tail (O(1) with Tail Pointer)]] — Maintaining tail pointer for instant O(1) appending vs O(N) scans.
- [[Singly Linked List Insert at Index (Order-Preserving)]] — Traversing to index i - 1 and splicing node in O(N) time.
- [[Singly Linked List Delete Head and Tail Elements]] — O(1) head removal vs O(N) tail removal requiring predecessor search.
- [[Singly Linked List Delete by Value (First Occurrence)]] — Locating target node and bypassing predecessor pointer.
- [[Singly Linked List Delete Node with O(1) Pointer Handle]] — Copying successor value into target slot to delete without predecessor pointer.
- [[Singly Linked List Memory Overhead and Cache Miss Economics]] — Analyzing pointer-to-payload ratio and pointer-chasing latency costs.
- [[Forward Iterator Mechanics and Traversal Bounds]] — Single-pass forward iteration invariants and termination checks.

### 2. 📂 [[Doubly Linked List|02. Doubly Linked List]]
- [[Doubly Linked List]] — Master topology: Bidirectional next and previous node links, O(1) arbitrary node deletion with handle, and sentinel-guarded boundaries.
- [[Doubly Linked List and Bidirectional Links]] — Bidirectional next/prev node links, symmetric pointer traversal, and memory overhead.
- [[Doubly Linked List Sentinel Nodes (Circular Dummy Invariant)]] — Using a single dummy sentinel node with circular next/prev links to eliminate null checks.
- [[Doubly Linked List O(1) Node Splice and Transfer]] — Zero-allocation transfer of sub-chains between doubly linked lists in constant time.
- [[LRU Cache Eviction Policy using Doubly Linked List and Hash Map]] — O(1) get/put cache architecture coupling a hash table with a doubly linked list.

### 3. 📂 [[Circular Linked List|03. Circular Linked List]]
- [[Circular Linked List]] — Master topology: Cyclic pointer chains where the tail references head, enabling infinite continuous traversals.
- [[Circular Linked List and Cyclic Looping Mechanics]] — Tail-to-head circular reference invariant and termination conditions.
- [[Circular Linked List Round-Robin Scheduler Implementation]] — Fair CPU time-slice distribution across active threads via circular traversal.
- [[Josephus Problem and Cyclic Elimination Mechanics]] — Mathematical elimination intervals in cyclic structures in O(N) time.
- [[Splitting and Concatenating Circular Linked Lists]] — O(1) concatenation and midpoint splitting of circular chains.

### 4. 📂 [[Intrusive Linked List|04. Intrusive Linked List]]
- [[Intrusive Linked List]] — Master topology: Nodes embedding pointer heads directly within data payloads, eliminating dynamic wrapper allocations.
- [[Intrusive Linked List Architecture (Linux Kernel list_head)]] — Embedding pointer hooks inside payloads and zero wrapper allocations.
- [[Linux Kernel list_head Topology and container_of Macro]] — Pointer arithmetic recovering enclosing struct pointers via offsetof calculations.
- [[Intrusive vs Non-Intrusive Zero-Allocation Performance Profiles]] — Eliminating cache misses and allocator pressure in systems engines.
- [[Embedding Multiple Intrusive Lists in a Single Struct]] — Simultaneously tracking an object in priority, hash, and FIFO queues without multiple wrappers.

### 5. 📂 [[Unrolled Linked List|05. Unrolled Linked List]]
- [[Unrolled Linked List]] — Master topology: Hybrid cache-conscious structure packing short contiguous arrays inside each linked list node.
- [[Unrolled Linked List Hybrid Cache Topology]] — Packing flat array chunks into nodes to match 64-byte CPU cache lines.
- [[Unrolled Node Capacity and Cache Line Alignment (64-Byte Matching)]] — Sizing node array buffers to perfectly fit L1 cache lines and SIMD lanes.
- [[Unrolled List Element Insertion and Node Splitting Mechanics]] — Maintaining balance through node splitting when chunk capacity exceeds threshold.
- [[Defragmentation and Node Merging in Unrolled Lists]] — Underflow handling: coalescing under-filled adjacent nodes to save memory.

### 6. 📂 [[Skip List|06. Skip List]]
- [[Skip List]] — Master topology: Multi-layer probabilistic forward pointers providing O(log N) search, insertion, and deletion without tree rebalancing.
- [[Skip List Probabilistic Express Towers and O(log N) Search]] — Multi-level express lanes with coin-toss geometric distribution.
- [[Skip List Probabilistic Geometric Level Distribution]] — Mathematical proofs of height p = 1/2 distribution and O(log N) expected bounds.
- [[Skip List Search, Insertion, and Predecessor Array Traversal]] — Tracking update arrays across tower levels during insertions and deletions.
- [[Lock-Free Concurrent Skip List (Java ConcurrentSkipListMap Mechanics)]] — Atomic CAS-based node linking enabling scalable concurrent sorted sets.

### 7. 📂 [[Pointer Algorithms & Techniques|07. Pointer Algorithms & Techniques]]
- [[Pointer Algorithms & Techniques]] — Master topology: Fundamental algorithmic primitives for linked pointer manipulation and traversal invariants.
- [[Reverse Linked List (Iterative 3-Pointer vs Recursive)]] — Inverting pointer direction in O(N) time and O(1) space.
- [[Middle of Linked List (Fast and Slow Pointer Floyd Pattern)]] — Finding midpoint with two-pointer tortoise and hare pattern.
- [[N-th Node from End (Two-Pointer Fixed-Gap Sliding Window)]] — Finding trailing elements in a single pass using a gap of size N.
- [[Detect and Break Loop in Linked List (Floyd Cycle Finding)]] — Detecting cycles, finding cycle entry point, and severing circular pointer.
- [[Merge Two Sorted Linked Lists (In-Place Pointer Splicing)]] — Interleaving two sorted lists into one sorted list in O(1) space.
- [[Palindrome Linked List Verification (In-Place Half Reverse)]] — Combining midpoint detection, half-reverse, and two-way scan.
- [[Intersection of Two Linked Lists (Difference-in-Length vs Cycle)]] — Finding merge points of convergent lists via length alignment.
- [[Merge K Sorted Linked Lists (Min-Heap vs Divide-and-Conquer)]] — Optimal O(N log K) merging using priority queues and pairwise merges.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`
