---
title: Generics
tags:
  - golang
  - generics
  - principal-swe
parent: "[[Golang]]"
---

# 🧬 Generics

Parametric polymorphism in Go: type parameters, constraints, type sets, generic data structures, standard generic packages, GcShape stenciling, and architectural patterns.

```text
Generics
│
├── [[Type Parameters & Functions|01. Type Parameters & Functions]]
│   ├── `Parametric Polymorphism Problem Statement`
│   ├── `Generic Function Declarations & Invocation`
│   ├── `Generic Method Limitations`
│   ├── `Type Inference Algorithms`
│   ├── `Instantiation Mechanics & Monomorphization`
│   └── `Generic Zero Value Idioms`
├── `02. Type Constraints & Type Sets`
│   ├── `Predeclared Constraints (any vs comparable)`
│   ├── `comparable Contract & Interface Equality Pitfalls`
│   ├── `Type Sets Mathematical Model`
│   ├── `Union Element (Pipe) & Intersection Sets`
│   ├── `Underlying Types & Tilde Operator (~T)`
│   ├── `cmp.Ordered Constraint`
│   ├── `Recursive Constraint Interfaces`
│   └── `Structural Constraint Interfaces`
├── [[Generic Data Structures & Collections|03. Generic Data Structures & Collections]]
│   ├── `Generic Slice Wrapper & High-Order Functions`
│   ├── `Generic Lock-Free Stack (Treiber Stack)`
│   ├── `Generic Lock-Free Queue (Michael-Scott Queue)`
│   ├── `Generic Concurrent Skip List`
│   ├── `Generic LRU & LFU Cache`
│   ├── `Generic Priority Queue (Binary Heap)`
│   ├── `Generic Ring Buffer & Circular Queue`
│   └── `Generic Result and Option Monads`
├── [[Standard Library Generics|04. Standard Library Generics]]
│   ├── `slices Package Deep Dive`
│   ├── `maps Package Deep Dive`
│   ├── `cmp Package Deep Dive`
│   ├── `sync.Map Typesafe Generic Wrapper`
│   └── [[atomic.Pointer Type Safety (Go 1.19+)]]
├── [[Compiler Mechanics & GcShape Stenciling|05. Compiler Mechanics & GcShape Stenciling]]
│   ├── `GcShape Stenciling Algorithm`
│   ├── `Dictionary Parameter Passing`
│   ├── `Monomorphization vs Type Erasure vs GcShape`
│   ├── `Binary Size Bloat & Compilation Time Analysis`
│   └── `Generics vs Interfaces Performance Benchmarks`
├── [[Generic Architecture & Design Patterns|06. Generic Architecture & Design Patterns]]
│   ├── `Generic Repository Pattern in Clean Architecture`
│   ├── `Generic Unit of Work Pattern`
│   ├── `Generic Builder & Functional Options Pattern`
│   ├── `Generic Event Bus & PubSub Pipeline`
│   ├── `Generic Middleware & Handler Pipeline`
│   └── `Generic Object Pool (sync.Pool Wrapper)`
└── [[Limitations, Gotchas & Anti-Patterns|07. Limitations, Gotchas & Anti-Patterns]]
│   ├── `No Generic Methods on Non-Generic Types Trap`
│   ├── `No Type Assertions on Type Parameters`
│   ├── `The Over-Parameterization Anti-Pattern`
│   ├── `Interface vs Generic Decision Tree`
│   └── `Generic Type Aliases (Go 1.24+)`
```

---

## 🗂️ Core Categories & Topics

### 📂 [[Type Parameters & Functions|01. Type Parameters & Functions]]
- [[Parametric Polymorphism Problem Statement]] — Why Go avoided generics for a decade and the problems type parameters solve.
- [[Generic Function Declarations & Invocation]] — Syntax for declaring and calling parameterized functions with [T any].
- [[Generic Method Limitations]] — Why Go does not allow parameterized methods on non-generic structs and function alternatives.
- [[Type Inference Algorithms]] — Function argument type inference and constraint type inference mechanics.
- [[Instantiation Mechanics & Monomorphization]] — Compile-time type parameter validation, substitution, and instantiation.
- [[Generic Zero Value Idioms]] — Returning zero values in generic functions using var zero T and new(T).
### 📂 [[Type Constraints|02. Type Constraints]]

### 📂 [[Type Sets|08. Type Sets]]

### 📂 [[Generic Data Structures & Collections|03. Generic Data Structures & Collections]]
- [[Generic Slice Wrapper & High-Order Functions]] — Building type-safe Map, Filter, Reduce, FlatMap, and Chunk slice helpers.
- [[Generic Lock-Free Stack (Treiber Stack)]] — Concurrent lock-free LIFO stack using atomic pointer CAS operations.
- [[Generic Lock-Free Queue (Michael-Scott Queue)]] — High-throughput lock-free FIFO queue with atomic head and tail pointers.
- [[Generic Concurrent Skip List]] — Probabilistic search and indexing structure with lockless concurrent reads.
- [[Generic LRU & LFU Cache]] — Thread-safe generic cache with O(1) eviction policies and TTL expiration.
- [[Generic Priority Queue (Binary Heap)]] — Type-safe generic priority queue wrapping container/heap.
- [[Generic Ring Buffer & Circular Queue]] — Fixed-capacity circular ring buffer for zero-allocation stream buffering.
- [[Generic Result and Option Monads]] — Functional error and optionality handling patterns (Result[T, E] and Option[T]).
### 📂 [[Standard Library Generics|04. Standard Library Generics]]
- [[slices Package Deep Dive]] — Generic slice algorithms: slices.Sort, slices.BinarySearch, slices.Contains, slices.Clone, slices.Delete.
- [[maps Package Deep Dive]] — Generic map helpers: maps.Clone, maps.Copy, maps.Equal, maps.DeleteFunc.
- [[cmp Package Deep Dive]] — Ordering functions: cmp.Compare, cmp.Less, and cmp.Or default fallback values.
- [[sync.Map Typesafe Generic Wrapper]] — Building a type-safe generic wrapper over sync.Map without casting.
- [[atomic.Pointer Type Safety (Go 1.19+)]] — Lock-free atomic pointer storage with full generic compile-time type safety.
### 📂 [[Compiler Mechanics & GcShape Stenciling|05. Compiler Mechanics & GcShape Stenciling]]
- [[GcShape Stenciling Algorithm]] — How Go shares machine code across all pointer types with identical GC shapes.
- [[Dictionary Parameter Passing]] — How the runtime passes type metadata dictionaries for scalar types at call sites.
- [[Monomorphization vs Type Erasure vs GcShape]] — Comparing C++ monomorphization, Java type erasure, and Go hybrid GcShape stenciling.
- [[Binary Size Bloat & Compilation Time Analysis]] — Analyzing the impact of generic instantiation on binary size and compilation speed.
- [[Generics vs Interfaces Performance Benchmarks]] — Zero-allocation execution, devirtualization, and 3-5x CPU performance gains over reflection.
### 📂 [[Generic Architecture & Design Patterns|06. Generic Architecture & Design Patterns]]
- [[Generic Repository Pattern in Clean Architecture]] — Universal CRUD database repository interfaces (Repository[T, ID]).
- [[Generic Unit of Work Pattern]] — Managing multi-repository database transactions with generic contracts.
- [[Generic Builder & Functional Options Pattern]] — Constructing complex typed domain objects with validation.
- [[Generic Event Bus & PubSub Pipeline]] — Type-safe event dispatcher and topic subscriber routing in Go.
- [[Generic Middleware & Handler Pipeline]] — Composable request/response middleware chains without interface boxing.
- [[Generic Object Pool (sync.Pool Wrapper)]] — Zero-allocation typed object reuse wrapping sync.Pool.
### 📂 [[Limitations, Gotchas & Anti-Patterns|07. Limitations, Gotchas & Anti-Patterns]]
- [[No Generic Methods on Non-Generic Types Trap]] — Why methods cannot introduce new type parameters and package-level function workarounds.
- [[No Type Assertions on Type Parameters]] — Why T.(int) is illegal and how to perform type dispatching via interface boxing.
- [[The Over-Parameterization Anti-Pattern]] — Avoiding unnecessary type parameter clutter that harms API readability.
- [[Interface vs Generic Decision Tree]] — Staff-level decision framework: when to choose dynamic interfaces vs static generics.
- [[Generic Type Aliases (Go 1.24+)]] — type Set[T] = map[T]struct{} generic type alias syntax and codebase refactoring.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Topics

- [[Compiler Mechanics & GcShape Stenciling]]
- [[Generic Architecture & Design Patterns]]
- [[Generic Data Structures & Collections]]
- [[Limitations, Gotchas & Anti-Patterns]]
- [[Standard Library Generics]]
- [[Type Constraints]]
- [[Type Parameters & Functions]]
- [[Type Sets]]
