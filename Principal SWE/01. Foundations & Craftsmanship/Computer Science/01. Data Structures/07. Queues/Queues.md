---
title: "Queues"
tags:
  - review
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# Queues

## 1. Definition

## 2. Mental Model

## 3. Usage

## 4. Gotchas

---

## 🗺️ Module Architecture & Sub-Domains

```
Queues/
├── Linear FIFO Queue/
├── Circular Queue/
├── Priority Queue/
├── Concurrent & Lock-Free Queue/
└── Common Queue Algorithms & Operations/
```

---

## 📑 Comprehensive Sub-Domain Index

### [[Linear FIFO Queue]]
  - [[Queue Dequeue (Head Removal)]]
  - [[Queue Empty and Full Invariant Conditions]]
  - [[Queue Enqueue (Tail Insertion)]]
  - [[Queue FIFO Invariants and State Transitions]]
  - [[Queue Linked List Implementation]]
  - [[Queue Peek (Front Element Inspection)]]
  - [[Queue Underflow and Overflow Exceptions vs Non-Blocking Options]]
  - [[Two-Stack FIFO Queue Implementation and Amortized O(1) Proof]]

### [[Circular Queue]]
  - [[Bounded Queue Backpressure and Drop Policies (Drop-Head vs Drop-Tail)]]
  - [[Circular Queue Dequeue and Head Advancement Operations]]
  - [[Circular Queue Enqueue and Tail Wraparound Operations]]
  - [[Circular Queue Monotonic Sequence Number Indexing]]
  - [[Queue Capacity Growth and Dynamic Resizing]]
  - [[Queue Circular Array Buffer Implementation]]
  - [[Queue Power-of-Two Bitwise Masking Indexing]]

### [[Priority Queue]]
  - [[Binary Heap Array Representation for Priority Queue]]
  - [[D-ary Heap (4-ary) Priority Queue for Cache Optimization]]
  - [[Indexed Priority Queue and O(log N) Decrease-Key Operation]]
  - [[Priority Queue Dequeue (Heap Pop and Sift-Down Operation)]]
  - [[Priority Queue Enqueue (Heap Push and Sift-Up Operation)]]
  - [[Priority Queue Peek (O(1) Root Extrema Inspection)]]
  - [[Priority Queue vs FIFO Queue Tradeoffs]]

### [[Concurrent & Lock-Free Queue]]
  - [[Bounded Blocking Queue (Condition Variables and Mutex Striping)]]
  - [[Lock-Free SPSC Ring Buffer Queue (Disruptor Pattern)]]
  - [[Michael-Scott Lock-Free FIFO Queue (CAS on Head and Tail)]]
  - [[Multi-Producer Multi-Consumer (MPMC) Queue (Atomic CAS)]]
  - [[The ABA Problem and Hazard Pointers in Lock-Free Queues]]

### [[Common Queue Algorithms & Operations]]
  - [[Breadth-First Search (BFS) Traversal via FIFO Queue]]
  - [[Level-Order Binary Tree Traversal using Queue]]
  - [[Moving Average from Data Stream using Queue]]
  - [[Shortest Path in Unweighted Graph via Queue]]

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Curriculum: `Computer Science`
