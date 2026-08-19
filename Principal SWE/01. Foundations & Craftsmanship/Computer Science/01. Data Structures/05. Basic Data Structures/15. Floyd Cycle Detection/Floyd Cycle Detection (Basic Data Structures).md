---
title: Floyd Cycle Detection (Basic Data Structures)
tags:
  - computer-science
  - data-structures
  - basic-data-structures
  - principal-swe
parent: "[[Basic Data Structures]]"
---

# 📦 Floyd Cycle Detection (Basic Data Structures)

Tortoise and Hare pointer algorithm for cycle detection and start-node location.

```text
Floyd Cycle Detection (Basic Data Structures)
│
├── [[Floyd Tortoise and Hare Cycle Detection Proof (2(F+a) = F+a+kC)]]
├── [[Detect Cycle in Singly Linked List (Fast-Slow Pointers)]]
├── [[Find Cycle Entry Point in Linked List]]
├── [[Calculate Linked List Cycle Length]]
├── [[Find Duplicate Number in Array (Cycle Pointer Indirection)]]
├── [[Happy Number Cycle Detection (State Graph Cycles)]]
└── [[Brent Cycle Detection Algorithm (Teleporting Hare)]]
```

---

## 🗂️ Operations & Topics

- [[Floyd Tortoise and Hare Cycle Detection Proof (2(F+a) = F+a+kC)]] — Mathematical proof showing 2x fast pointer meets slow pointer within cycle in O(N) time.
- [[Detect Cycle in Singly Linked List (Fast-Slow Pointers)]] — Fast and slow pointer traversal determining cycle existence in O(N) time and O(1) space.
- [[Find Cycle Entry Point in Linked List]] — Resetting slow pointer to head and stepping both at 1x speed to meet at exact cycle origin.
- [[Calculate Linked List Cycle Length]] — Freezing slow pointer and advancing fast pointer until it circles back to determine cycle length C.
- [[Find Duplicate Number in Array (Cycle Pointer Indirection)]] — Mapping array values as next pointers (A[i] -> next) to locate duplicate in O(N) time and O(1) space.
- [[Happy Number Cycle Detection (State Graph Cycles)]] — Detecting infinite repeating sum-of-squares cycles via Floyd pointer detection.
- [[Brent Cycle Detection Algorithm (Teleporting Hare)]] — Powers-of-two teleporting fast pointer finding cycles with 24% fewer pointer steps than Floyd.

---

## 🔗 References
- ⬆️ Parent: [[Basic Data Structures]]
- 📚 Module: `Data Structures & Algorithms`

