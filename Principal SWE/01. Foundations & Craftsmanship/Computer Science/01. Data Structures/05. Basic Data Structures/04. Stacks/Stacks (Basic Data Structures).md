---
title: Stacks (Basic Data Structures)
tags:
  - review
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Stacks (Basic Data Structures)

LIFO state management, dynamic array stacks, call stacks, and O(1) Min/Max stacks.

```text
Stacks (Basic Data Structures)
│
├── [[Stack LIFO Invariants and Call Frame Models]]
├── [[Stack Push (Top Insertion)]]
├── [[Stack Pop (Top Removal with GC Zeroing)]]
├── [[Stack Peek and Top Inspection]]
├── [[Stack Dynamic Array Implementation]]
├── [[Stack Linked List Implementation]]
├── [[Stack Underflow and Overflow Invariants]]
├── [[Stack Balanced Parentheses and Syntax Parsing]]
├── [[Min Stack and Max Stack with O(1) Auxiliary State]]
├── [[Two Stacks in a Single Array (Converging Tops)]]
├── [[Queue Implementation using Two Stacks]]
└── [[Evaluating Postfix and Infix Expressions with Stacks]]
```

---

## 🗂️ Operations & Topics

- [[Stack LIFO Invariants and Call Frame Models]] — Last-In First-Out state axioms and hardware CPU execution stack frames.
- [[Stack Push (Top Insertion)]] — Amortized O(1) insertion at top index with dynamic reallocation.
- [[Stack Pop (Top Removal with GC Zeroing)]] — O(1) top removal with explicit slot zeroing to prevent GC memory leaks.
- [[Stack Peek and Top Inspection]] — O(1) inspecting top element without mutating stack depth.
- [[Stack Dynamic Array Implementation]] — Contiguous buffer stack with geometric 2.0x growth factor.
- [[Stack Linked List Implementation]] — Node allocation stack with zero reallocation spikes and O(1) push/pop.
- [[Stack Underflow and Overflow Invariants]] — Boundary checking, recursion stack limits, and SIGSEGV stack overflows.
- [[Stack Balanced Parentheses and Syntax Parsing]] — Validating nested bracket syntax and grammar parsing in O(N).
- [[Min Stack and Max Stack with O(1) Auxiliary State]] — Maintaining running minimum/maximum in O(1) time using paired minimum stack.
- [[Two Stacks in a Single Array (Converging Tops)]] — Packing two growing stacks at opposite ends of a single fixed array.
- [[Queue Implementation using Two Stacks]] — Simulating FIFO queue using In-Stack and Out-Stack in amortized O(1).
- [[Evaluating Postfix and Infix Expressions with Stacks]] — Dijkstra Shunting-Yard algorithm and postfix Reverse Polish Notation evaluation.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: `Data Structures & Algorithms`

