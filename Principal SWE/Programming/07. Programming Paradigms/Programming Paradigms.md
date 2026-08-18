---
title: Programming Paradigms
tags:
  - programming
  - paradigms
  - principal-swe
parent: "[[Programming]]"
---

# 💻 Programming Paradigms

Imperative, functional, declarative, reactive streams, and Actor model concurrency.

```text
Programming Paradigms
│
├── [[Imperative & Procedural Paradigm|01. Imperative & Procedural Paradigm]]
│   ├── [[Structured Programming (Dijkstra, Structured Control Flow)]]
│   ├── [[Procedural Abstraction & Call Stacks]]
│   └── [[State Mutation, Side Effects & Memory Mutation]]
├── [[Functional Programming (FP)|02. Functional Programming (FP)]]
│   ├── [[Pure Functions, Immutability & Referential Transparency]]
│   ├── [[First-Class & Higher-Order Functions (Map, Filter, Reduce)]]
│   ├── [[Currying, Partial Application & Function Composition]]
│   ├── [[Monads, Functors & Algebraic Data Types (ADT)]]
│   └── [[Lazy Evaluation & Infinite Streams]]
└── [[Declarative & Reactive Programming|03. Declarative & Reactive Programming]]
│   ├── [[Declarative vs Imperative UI & State (React, Flutter, SwiftUI)]]
│   ├── [[Reactive Streams & Event-Driven Dataflows (Rx, Observers)]]
│   ├── [[Actor Model Concurrency (Erlang, Akka)]]
│   └── [[Dataflow & Pipeline Architectures]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Imperative & Procedural Paradigm|01. Imperative & Procedural Paradigm]]
- [[Structured Programming (Dijkstra, Structured Control Flow)]] — Eliminating arbitrary jumps (goto) in favor of sequencing, selection (if/else), and iteration (while/for).
- [[Procedural Abstraction & Call Stacks]] — Decomposing algorithms into subroutines, activation records, stack frames, and parameter passing conventions.
- [[State Mutation, Side Effects & Memory Mutation]] — Managing mutable shared memory, temporal coupling, and side-effect visibility across procedure calls.
### 2. 📂 [[Functional Programming (FP)|02. Functional Programming (FP)]]
- [[Pure Functions, Immutability & Referential Transparency]] — Functions with zero side effects where output depends strictly on input parameters, enabling effortless caching and concurrency.
- [[First-Class & Higher-Order Functions (Map, Filter, Reduce)]] — Treating functions as data values, passing functions as parameters, and composing transformation pipelines.
- [[Currying, Partial Application & Function Composition]] — Transforming multi-argument functions into unary function chains and combining functions into pipelines.
- [[Monads, Functors & Algebraic Data Types (ADT)]] — Sum types (enums/unions), product types (structs/tuples), Functor mapping, and Monadic sequencing (Option, Result).
- [[Lazy Evaluation & Infinite Streams]] — Deferring computation until values are explicitly required and processing unbounded data streams.
### 3. 📂 [[Declarative & Reactive Programming|03. Declarative & Reactive Programming]]
- [[Declarative vs Imperative UI & State (React, Flutter, SwiftUI)]] — Expressing what the UI should look like for a given state rather than issuing step-by-step DOM mutations.
- [[Reactive Streams & Event-Driven Dataflows (Rx, Observers)]] — Asynchronous push-based data streams with backpressure handling, transformation operators, and subscriber models.
- [[Actor Model Concurrency (Erlang, Akka)]] — Share-nothing concurrency: isolated actors communicating strictly via asynchronous message passing.
- [[Dataflow & Pipeline Architectures]] — Building streaming computation DAGs where nodes process inputs as tokens arrive on data channels.

---

## 🔗 Navigation
- ⬆️ Parent: [[Programming]]
- 🎓 Root: [[Principal SWE]]
