---
title: Methods & Interfaces
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Golang]]"
---

# 🧩 Methods & Interfaces

Methods, receiver semantics, method sets, structural interfaces, type assertions, iface/eface runtime layout, itab dispatch tables, devirtualization, and interface architecture.

```text
Methods & Interfaces
│
├── [[Method Declarations & Receivers|01. Method Declarations & Receivers]]
│   ├── [[Methods vs Functions Architectural Differences]]
│   ├── [[Pointer Receivers (State Mutation & Copy Elimination)]]
│   ├── [[Value Receivers (Immutability & Concurrency Safety)]]
│   ├── [[Receiver Type Consistency & Heuristics Decision Tree]]
│   ├── [[Methods on Defined Non-Struct Types]]
│   ├── [[Method Sets for T and Pointer-T]]
│   └── [[Cross-Package Method Definition Rules & Locality]]
├── [[Method Values, Expressions & Promotion|02. Method Values, Expressions & Promotion]]
│   ├── [[Method Values (Instance Binding & Closures)]]
│   ├── [[Method Expressions (Static Functions)]]
│   ├── [[Embedded Struct Method Promotion]]
│   ├── [[Method Shadowing & Ambiguity in Nested Structs]]
│   └── [[Auto-Referencing and Auto-Dereferencing]]
├── [[Interface Fundamentals & Contracts|03. Interface Fundamentals & Contracts]]
│   ├── [[Structural Typing & Implicit Satisfaction]]
│   ├── [[Consumer-Defined Interfaces (Accept Interfaces, Return Structs)]]
│   ├── [[Interface Composition & Embedding]]
│   ├── [[Empty Interface (any) & Type Boxing Mechanics]]
│   ├── [[Common Standard Library Contracts]]
│   └── [[Sealed Interfaces (Restricting External Implementations)]]
├── [[Type Assertions, Switches & Inspection|04. Type Assertions, Switches & Inspection]]
│   ├── [[Type Assertions & Comma-Ok Idiom]]
│   ├── [[Unchecked Type Assertions & Panic Hazards]]
│   ├── [[Type Switch Dynamic Dispatching]]
│   ├── [[The Typed Nil Interface Trap (*MyType(nil) != nil)]]
│   └── [[Defensive Nil Checking Patterns]]
├── [[Runtime Internals (iface, eface, itab)|05. Runtime Internals (iface, eface, itab)]]
│   ├── [[iface Runtime Memory Layout (tab, data)]]
│   ├── [[eface Runtime Memory Layout (_type, data)]]
│   ├── [[itab Virtual Method Table & Function Offsets]]
│   ├── [[itab Global Cache & Dynamic Type Hash Tables]]
│   ├── [[Direct Interface Values (Small Word Optimization)]]
│   └── [[Interface Boxing & Heap Allocation Triggers]]
├── [[Compiler Optimizations & Devirtualization|06. Compiler Optimizations & Devirtualization]]
│   ├── [[Devirtualization Compiler Pass]]
│   ├── [[Dynamic Interface Dispatch Overhead]]
│   ├── [[Mid-Stack Inlining Barriers with Interfaces]]
│   ├── [[Escape Analysis with Interface Boxing]]
│   └── [[Benchmarking Concrete Calls vs Interface Virtual Calls]]
└── [[Interface Architecture & Design Patterns|07. Interface Architecture & Design Patterns]]
│   ├── [[Interface Segregation Principle (ISP) in Go]]
│   ├── [[Adapter Pattern with Interfaces]]
│   ├── [[Decorator & Middleware Pattern with Interfaces]]
│   ├── [[Strategy Pattern via Functional Interfaces]]
│   ├── [[The Interface Pollution Anti-Pattern]]
│   ├── [[Mocking What You Do Not Own Anti-Pattern]]
│   └── [[Staff-Level Interface Design Guidelines]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Method Declarations & Receivers|01. Method Declarations & Receivers]]
- [[Methods vs Functions Architectural Differences]] — Contrasting method receivers with first-class function parameters in API design.
- [[Pointer Receivers (State Mutation & Copy Elimination)]] — Mutating receiver state, eliminating large struct copy overhead, and consistency rules.
- [[Value Receivers (Immutability & Concurrency Safety)]] — Value copying semantics, concurrent read safety, and immutability guarantees.
- [[Receiver Type Consistency & Heuristics Decision Tree]] — Staff-level decision tree for choosing between value and pointer receivers.
- [[Methods on Defined Non-Struct Types]] — Attaching methods to custom primitives (type MyDuration int64, type StringSet map[string]struct{}).
- [[Method Sets for T and Pointer-T]] — Rules defining method sets for value T and pointer *T and their interface assignability.
- [[Cross-Package Method Definition Rules & Locality]] — Type locality constraints prohibiting method declarations on types from foreign packages.
### 2. 📂 [[Method Values, Expressions & Promotion|02. Method Values, Expressions & Promotion]]
- [[Method Values (Instance Binding & Closures)]] — Binding a struct instance to a method, returning a first-class closure, and escape analysis.
- [[Method Expressions (Static Functions)]] — Calling methods as static unbound functions with explicit receiver parameters (Type.Method).
- [[Embedded Struct Method Promotion]] — Composition over inheritance: promoting inner embedded methods to outer struct method set.
- [[Method Shadowing & Ambiguity in Nested Structs]] — Outer method overriding inner methods and resolving duplicate method collision errors.
- [[Auto-Referencing and Auto-Dereferencing]] — How the compiler automatically inserts & or * when invoking methods across pointer/value boundaries.
### 3. 📂 [[Interface Fundamentals & Contracts|03. Interface Fundamentals & Contracts]]
- [[Structural Typing & Implicit Satisfaction]] — Why Go rejected explicit implements keywords in favor of compile-time duck typing.
- [[Consumer-Defined Interfaces (Accept Interfaces, Return Structs)]] — Staff design principle: defining minimal interfaces in consumer packages.
- [[Interface Composition & Embedding]] — Composing fine-grained interfaces (io.ReadWriteCloser = io.Reader + io.Writer + io.Closer).
- [[Empty Interface (any) & Type Boxing Mechanics]] — Working with arbitrary types, boxing overhead, and type safety tradeoffs.
- [[Common Standard Library Contracts]] — Core contracts: io.Reader, io.Writer, io.Closer, fmt.Stringer, error, sort.Interface.
- [[Sealed Interfaces (Restricting External Implementations)]] — Restricting interface implementations to internal packages via unexported method tokens.
### 4. 📂 [[Type Assertions, Switches & Inspection|04. Type Assertions, Switches & Inspection]]
- [[Type Assertions & Comma-Ok Idiom]] — Extracting concrete types from interfaces safely using val, ok := i.(ConcreteType).
- [[Unchecked Type Assertions & Panic Hazards]] — Why unchecked assertions (val := i.(ConcreteType)) cause fatal runtime panics on mismatch.
- [[Type Switch Dynamic Dispatching]] — Multi-way type branching using switch v := i.(type) syntax.
- [[The Typed Nil Interface Trap (*MyType(nil) != nil)]] — Why an interface holding a typed nil pointer is not equal to nil (iface.tab != nil).
- [[Defensive Nil Checking Patterns]] — Architectural patterns for safely checking both interface nil and underlying pointer nil.
### 5. 📂 [[Runtime Internals (iface, eface, itab)|05. Runtime Internals (iface, eface, itab)]]
- [[iface Runtime Memory Layout (tab, data)]] — Non-empty interface layout: *itab table pointer and unsafe.Pointer data pointer (16 bytes).
- [[eface Runtime Memory Layout (_type, data)]] — Empty interface layout: *_type metadata pointer and unsafe.Pointer data pointer (16 bytes).
- [[itab Virtual Method Table & Function Offsets]] — Layout of itab: interface type, concrete type, hash code, and function pointers array [1]uintptr.
- [[itab Global Cache & Dynamic Type Hash Tables]] — How runtime dynamically computes and caches itab instances for type pairs in a global hash table.
- [[Direct Interface Values (Small Word Optimization)]] — Compiler optimization storing pointers and small scalar values directly inside the data word.
- [[Interface Boxing & Heap Allocation Triggers]] — When assigning a value to an interface forces heap allocation vs stack escape.
### 6. 📂 [[Compiler Optimizations & Devirtualization|06. Compiler Optimizations & Devirtualization]]
- [[Devirtualization Compiler Pass]] — How compiler statically proves concrete types and transforms indirect itab calls into direct static calls.
- [[Dynamic Interface Dispatch Overhead]] — Measuring the CPU penalty of indirect function pointer calls and CPU branch predictor misses.
- [[Mid-Stack Inlining Barriers with Interfaces]] — Why virtual interface calls prevent compiler function inlining across architectural boundaries.
- [[Escape Analysis with Interface Boxing]] — Why passing values to interface parameters (e.g. fmt.Println) frequently triggers heap escapes.
- [[Benchmarking Concrete Calls vs Interface Virtual Calls]] — Microbenchmarks comparing direct struct calls, devirtualized calls, and dynamic interface calls.
### 7. 📂 [[Interface Architecture & Design Patterns|07. Interface Architecture & Design Patterns]]
- [[Interface Segregation Principle (ISP) in Go]] — Keeping interfaces small (1-2 methods) and tailored to specific consumer requirements.
- [[Adapter Pattern with Interfaces]] — Bridging incompatible third-party libraries behind clean domain interfaces.
- [[Decorator & Middleware Pattern with Interfaces]] — Composing cross-cutting concerns (logging, metrics, retries) via interface wrappers.
- [[Strategy Pattern via Functional Interfaces]] — Injecting interchangeable algorithms at runtime using single-method interfaces or function types.
- [[The Interface Pollution Anti-Pattern]] — Prematurely defining interfaces for every struct before having multiple implementations.
- [[Mocking What You Do Not Own Anti-Pattern]] — Why you should only define interfaces and mocks for code you own and consume.
- [[Staff-Level Interface Design Guidelines]] — Best practices for designing scalable, maintainable Go APIs with clean interface boundaries.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

