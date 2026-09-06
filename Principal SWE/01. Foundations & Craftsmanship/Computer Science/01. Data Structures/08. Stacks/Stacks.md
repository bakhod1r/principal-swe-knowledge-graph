---
title: Stacks
tags:
  - review
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Data Structures]]"
---

# 📦 Stacks

Last-In First-Out data structures, call frame mechanics, min-max auxiliary state tracking, monotonic stacks, and shunting-yard compilers.

```text
Stacks
│
├── 01. Array-Based Stack
│   ├── [[Array-Based Stack]]
│   ├── [[Stack LIFO Invariants and State Transitions]]
│   ├── [[Stack Push Operation (Element Insertion)]]
│   ├── [[Stack Pop Operation (Element Removal)]]
│   ├── [[Stack Peek Operation (Top Element Inspection)]]
│   ├── [[Stack Overflow and Underflow Bounds Checking]]
│   ├── [[Stack Frame Memory Allocation vs Heap Dynamic Arrays]]
│   └── [[Array-Based Stack Cache Locality and Vectorization]]
││
├── 02. Linked List Stack
│   ├── [[Linked List Stack]]
│   ├── [[Stack Linked List Implementation]]
│   ├── [[Linked List Stack Memory Fragmentation and Free-List Recycling]]
│   └── [[Lock-Free Concurrent Stack (Treiber Stack with Atomic CAS)]]
││
├── 03. Min-Max Stack
│   ├── [[Min-Max Stack]]
│   ├── [[Min-Max Stack Architecture (O(1) Extrema Queries)]]
│   ├── [[Space-Optimized Min-Stack (Single-Value Delta Encoding)]]
│   └── [[Min-Max Stack with Frequency Tracking]]
││
├── 04. Multi-Stack & Two Stacks
│   ├── [[Multi-Stack]]
│   ├── [[Two Stacks in a Single Array (Converging Tops)]]
│   ├── [[Queue Implementation Using Two Stacks]]
│   ├── [[K-Stacks in a Single Dynamic Array (Free-List Index Tracking)]]
│   └── [[Queue Reconstruction via Double Inverted Stacks]]
││
├── 05. Stack Parsing Applications
│   ├── [[Stack Parsing Applications]]
│   ├── [[Parentheses and Delimiter Matching Algorithm]]
│   ├── [[Infix to Postfix Conversion (Shunting-Yard Algorithm)]]
│   ├── [[Postfix Expression Evaluation (Reverse Polish Notation)]]
│   ├── [[Syntax Abstract Syntax Tree (AST) Generation via Operator Stacks]]
│   └── [[Call Stack Frame Unwinding and Exception Handling Mechanics]]
││
└── 06. Monotonic Stack Algorithms
    ├── [[Monotonic Stack]]
    ├── [[Next Greater Element and Previous Greater Element Patterns]]
    ├── [[Largest Rectangle in Histogram (Monotonic Increasing Stack)]]
    └── [[Trapping Rain Water via Monotonic Decreasing Stack]]
```

---

## 🗂️ Concrete Types & Knowledge Domains

### 1. 📂 [[Array-Based Stack|01. Array-Based Stack]]
- [[Array-Based Stack]] — Master topology: Contiguous array stack with top index pointer.
- [[Stack LIFO Invariants and State Transitions]] — Last-In First-Out operational axioms, push/pop state transitions, and top-of-stack pointer.
- [[Stack Push Operation (Element Insertion)]] — O(1) push operation at the top index with capacity reallocation check.
- [[Stack Pop Operation (Element Removal)]] — O(1) pop operation returning top element and decrementing top pointer.
- [[Stack Peek Operation (Top Element Inspection)]] — O(1) viewing top element without mutating stack state.
- [[Stack Overflow and Underflow Bounds Checking]] — Validating top pointer against zero and maximum capacity bounds.
- [[Stack Frame Memory Allocation vs Heap Dynamic Arrays]] — CPU hardware stack pointer (rsp) adjustment vs heap memory allocation latency.
- [[Array-Based Stack Cache Locality and Vectorization]] — Cache line benefits of contiguous stack top access in high-frequency operations.

### 2. 📂 [[Linked List Stack|02. Linked List Stack]]
- [[Linked List Stack]] — Master topology: Node-based stack with head pointer as top.
- [[Stack Linked List Implementation]] — O(1) push and pop at list head avoiding array reallocation pauses.
- [[Linked List Stack Memory Fragmentation and Free-List Recycling]] — Mitigating heap allocation overhead via thread-local node freelists.
- [[Lock-Free Concurrent Stack (Treiber Stack with Atomic CAS)]] — Atomic CAS pointer swing on top pointer enabling wait-free concurrent pushes.

### 3. 📂 [[Min-Max Stack|03. Min-Max Stack]]
- [[Min-Max Stack]] — Master topology: Tracking minimum and maximum elements in O(1) time.
- [[Min-Max Stack Architecture (O(1) Extrema Queries)]] — Dual-value tracking or auxiliary stack maintaining current minimum at each level.
- [[Space-Optimized Min-Stack (Single-Value Delta Encoding)]] — Storing value deltas (2v - min) to achieve O(1) min queries without auxiliary stacks.
- [[Min-Max Stack with Frequency Tracking]] — Counting occurrences of extrema values to avoid duplicate tracking allocations.

### 4. 📂 [[Multi-Stack|04. Multi-Stack & Two Stacks]]
- [[Multi-Stack]] — Master topology: Embedding multiple stacks in a single contiguous buffer.
- [[Two Stacks in a Single Array (Converging Tops)]] — Top-1 grows from left (0), Top-2 grows from right (N-1), maximizing memory utilization.
- [[Queue Implementation Using Two Stacks]] — Amortized O(1) FIFO queue implementation via in-stack and out-stack transfers.
- [[K-Stacks in a Single Dynamic Array (Free-List Index Tracking)]] — Array of next pointers partitioning a single flat buffer among K arbitrary stacks.
- [[Queue Reconstruction via Double Inverted Stacks]] — Mathematical proof of amortized cost and inversion invariants.

### 5. 📂 [[Stack Parsing Applications|05. Stack Parsing Applications]]
- [[Stack Parsing Applications]] — Master topology: Grammar evaluation, bracket balancing, and operator precedence parsing.
- [[Parentheses and Delimiter Matching Algorithm]] — Validating nested braces, brackets, and parenthesis balance via LIFO stack.
- [[Infix to Postfix Conversion (Shunting-Yard Algorithm)]] — Dijkstra's shunting-yard algorithm transforming algebraic infix into postfix notation.
- [[Postfix Expression Evaluation (Reverse Polish Notation)]] — Single-pass arithmetic operand evaluation on RPN expressions.
- [[Syntax Abstract Syntax Tree (AST) Generation via Operator Stacks]] — Building expression trees using operator precedence and node stacks.
- [[Call Stack Frame Unwinding and Exception Handling Mechanics]] — How runtimes unwind stack frames and locate landing pads during exception propagation.

### 6. 📂 [[Monotonic Stack|06. Monotonic Stack Algorithms]]
- [[Monotonic Stack]] — Master topology: Strictly increasing or decreasing stacks solving nearest element queries in O(N) time.
- [[Next Greater Element and Previous Greater Element Patterns]] — Finding first larger elements on left and right in a single linear pass.
- [[Largest Rectangle in Histogram (Monotonic Increasing Stack)]] — Calculating maximum rectangular areas in linear time via boundary tracking.
- [[Trapping Rain Water via Monotonic Decreasing Stack]] — Calculating elevation water capacity by maintaining decreasing height bounds.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`
