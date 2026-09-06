---
title: Stacks
tags:
  - review
  - computer-science
  - data-structures
  - stacks
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Stacks (LIFO Call Frames & Parsing Topologies)

Comprehensive engineering catalog of Last-In First-Out state management, contiguous dynamic buffers, node-allocated call stacks, auxiliary $O(1)$ Min/Max tracking, dual-stack convergence, and compiler syntax parsing across 5 specialized domains.

```text
Stacks
│
├── 01. Array-Based Stack
│   ├── [[Array-Based Stack]]
│   ├── [[Stack LIFO Invariants and Call Frame Models]]
│   ├── [[Stack Underflow and Overflow Invariants]]
│   ├── [[Stack Push (Top Insertion)]]
│   ├── [[Stack Pop (Top Removal with GC Zeroing)]]
│   ├── [[Stack Peek and Top Inspection]]
│   └── [[Stack Dynamic Array Implementation]]
│
├── 02. Linked List Stack
│   ├── [[Linked List Stack]]
│   └── [[Stack Linked List Implementation]]
│
├── 03. Min-Max Stack
│   ├── [[Min-Max Stack]]
│   └── [[Min Stack and Max Stack with O(1) Auxiliary State]]
│
├── 04. Multi-Stack & Two Stacks
│   ├── [[Multi-Stack]]
│   ├── [[Two Stacks in a Single Array (Converging Tops)]]
│   └── [[Queue Implementation using Two Stacks]]
│
└── 05. Stack Parsing Applications
    ├── [[Stack Parsing Applications]]
    ├── [[Stack Balanced Parentheses and Syntax Parsing]]
    └── [[Evaluating Postfix and Infix Expressions with Stacks]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Array-Based Stack|01. Array-Based Stack]]
- [[Array-Based Stack]] — Master topology of contiguous array-backed stacks.
- [[Stack LIFO Invariants and Call Frame Models]] — Last-In First-Out state axioms and hardware CPU execution stack frames.
- [[Stack Underflow and Overflow Invariants]] — Boundary checking, recursion stack limits, and SIGSEGV stack overflows.
- [[Stack Push (Top Insertion)]] — Amortized O(1) insertion at top index with dynamic reallocation.
- [[Stack Pop (Top Removal with GC Zeroing)]] — O(1) top removal with explicit slot zeroing to prevent GC memory leaks.
- [[Stack Peek and Top Inspection]] — O(1) inspecting top element without mutating stack depth.
- [[Stack Dynamic Array Implementation]] — Contiguous buffer stack with geometric 2.0x growth factor.

### 2. 📂 [[Linked List Stack|02. Linked List Stack]]
- [[Linked List Stack]] — Master topology of heap node-linked stacks.
- [[Stack Linked List Implementation]] — Node allocation stack with zero reallocation spikes and O(1) push/pop.

### 3. 📂 [[Min-Max Stack|03. Min-Max Stack]]
- [[Min-Max Stack]] — Master topology of O(1) auxiliary state stacks.
- [[Min Stack and Max Stack with O(1) Auxiliary State]] — Maintaining running minimum/maximum in O(1) time using paired minimum stack.

### 4. 📂 [[Multi-Stack|04. Multi-Stack & Two Stacks]]
- [[Multi-Stack]] — Master topology of multi-stack buffer packing.
- [[Two Stacks in a Single Array (Converging Tops)]] — Packing two growing stacks at opposite ends of a single fixed array.
- [[Queue Implementation using Two Stacks]] — Simulating FIFO queue using In-Stack and Out-Stack in amortized O(1).

### 5. 📂 [[Stack Parsing Applications|05. Stack Parsing Applications]]
- [[Stack Parsing Applications]] — Master domain for compiler grammar and expression evaluation.
- [[Stack Balanced Parentheses and Syntax Parsing]] — Validating nested bracket syntax and grammar parsing in O(N).
- [[Evaluating Postfix and Infix Expressions with Stacks]] — Dijkstra Shunting-Yard algorithm and postfix Reverse Polish Notation evaluation.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`



