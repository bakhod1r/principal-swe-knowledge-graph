---
title: Rust
tags:
  - programming
  - languages
  - rust
  - principal-swe
parent: "[[Language]]"
---

# 💻 Rust

Systems programming in Rust: Ownership, borrowing, lifetimes, fearless concurrency, zero-cost abstractions, trait systems, macros, unsafe Rust, FFI, SIMD, and async Tokio runtime.

```text
Rust
│
├── [[Introduction to Rust, Cargo, and Rustup Toolchain|01. Introduction to Rust and Toolchain]]
├── [[Rust Language Basics, Scalar Types, and Pattern Matching|02. Language Basics, Primitives, and Control Flow]]
├── [[Rust Ownership System, Borrowing Rules, and Move Semantics|03. Ownership, Borrowing, and Move Semantics]]
├── [[Rust Lifetimes, Lifetime Annotations, and Borrow Checker Internals|04. Lifetimes and Borrow Checker Internals]]
├── [[Rust Structs, Enums, and Algebraic Data Types|05. Structs, Enums, and Algebraic Data Types]]
├── [[Rust Traits, Generics, and Static vs Dynamic Dispatch|06. Traits, Generics, and Monomorphization]]
├── [[Rust Error Handling: Result, Option, and the Question Mark Operator|07. Error Handling with Result and Option]]
├── [[Rust Modules, Crates, Visibility, and Multi Crate Workspaces|08. Modules, Crates, and Workspace Architectures]]
├── [[Rust Concurrency: Threads, Channels, Mutex, and Send Sync|09. Concurrency, Threads, and Fearless Concurrency]]
├── [[Async Rust, Futures, Pinning, and Tokio Runtime Architecture|10. Async Rust and Tokio Runtime Architecture]]
├── [[Rust Smart Pointers: Box, Rc, Arc, Refcell, and RAII|11. Smart Pointers and RAII Memory Management]]
├── [[Rust Macros Metaprogramming: Declarative and Procedural Macros|12. Declarative and Procedural Macros Metaprogramming]]
├── [[Unsafe Rust, Raw Pointers, and Foreign Function Interface (ffi)|13. Unsafe Rust, Pointers, and C FFI Interoperability]]
└── [[High Performance Rust: Criterion Benchmarking, Perf, and SIMD|14. Performance Profiling, Criterion, and High Performance Rust]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Introduction to Rust, Cargo, and Rustup Toolchain|01. Introduction to Rust and Toolchain]] — Rust compiler architecture (`rustc`), Cargo package manager, Clippy linter, Rustfmt, release profile optimizations (`opt-level = 3`, LTO), and target triples.
- 📂 [[Rust Language Basics, Scalar Types, and Pattern Matching|02. Language Basics, Primitives, and Control Flow]] — Scalar primitives, compound types (tuples, arrays), immutable by default variables, expression-oriented syntax, `match` exhaustiveness, and `if let`/`while let` ergonomics.
- 📂 [[Rust Ownership System, Borrowing Rules, and Move Semantics|03. Ownership, Borrowing, and Move Semantics]] — Affine type system, Single Owner Principle, move semantics, copy semantics (`Copy` vs `Clone`), and preventing double-free/use-after-free at compile time.
- 📂 [[Rust Lifetimes, Lifetime Annotations, and Borrow Checker Internals|04. Lifetimes and Borrow Checker Internals]] — Non-Lexical Lifetimes (NLL), lifetime elision rules, struct lifetime parameters, static lifetime (`'static`), and solving complex borrow checker lifetime bounds.
- 📂 [[Rust Structs, Enums, and Algebraic Data Types|05. Structs, Enums, and Algebraic Data Types]] — Named-field structs, tuple structs, unit structs, tagged unions with enums, payload-carrying variants, and memory layout optimization (null pointer optimization).
- 📂 [[Rust Traits, Generics, and Static vs Dynamic Dispatch|06. Traits, Generics, and Monomorphization]] — Trait definitions, trait bounds, supertraits, zero-cost monomorphization (static dispatch), trait objects (`dyn Trait`, vtables), and associated types.
- 📂 [[Rust Error Handling: Result, Option, and the Question Mark Operator|07. Error Handling with Result and Option]] — Panic vs recoverable errors, `Option<T>` and `Result<T, E>`, bubbling errors with `?`, custom error types with `thiserror`, and context chaining with `anyhow`.
- 📂 [[Rust Modules, Crates, Visibility, and Multi Crate Workspaces|08. Modules, Crates, and Workspace Architectures]] — Module tree (`mod.rs`, path routing), visibility modifiers (`pub`, `pub(crate)`), external crate dependencies (`Cargo.toml`), and multi-package Cargo workspaces.
- 📂 [[Rust Concurrency: Threads, Channels, Mutex, and Send Sync|09. Concurrency, Threads, and Fearless Concurrency]] — OS threads (`std::thread`), message passing with MPSC channels, shared state concurrency (`Arc<Mutex<T>>`, `RwLock<T>`), and compiler-enforced `Send` and `Sync` traits.
- 📂 [[Async Rust, Futures, Pinning, and Tokio Runtime Architecture|10. Async Rust and Tokio Runtime Architecture]] — Asynchronous state machines, cooperative multitasking, `Future` polling contract, `Pin<&mut T>`, Tokio multi-threaded work-stealing scheduler, and select loops.
- 📂 [[Rust Smart Pointers: Box, Rc, Arc, Refcell, and RAII|11. Smart Pointers and RAII Memory Management]] — Heap allocation with `Box<T>`, reference counting with `Rc<T>` and atomic `Arc<T>`, interior mutability with `RefCell<T>` and `Mutex<T>`, and `Drop` trait RAII cleanup.
- 📂 [[Rust Macros Metaprogramming: Declarative and Procedural Macros|12. Declarative and Procedural Macros Metaprogramming]] — `macro_rules!` pattern-matching declarative macros, procedural macros (Derive, Attribute-like, Function-like), `syn`, and `quote` AST manipulation.
- 📂 [[Unsafe Rust, Raw Pointers, and Foreign Function Interface (ffi)|13. Unsafe Rust, Pointers, and C FFI Interoperability]] — The 5 superpowers of Unsafe Rust, raw pointers (`*const T`, `*mut T`), undefined behavior invariants, creating safe abstractions over unsafe code, and `extern "C"` FFI.
- 📂 [[High Performance Rust: Criterion Benchmarking, Perf, and SIMD|14. Performance Profiling, Criterion, and High Performance Rust]] — Statistical micro-benchmarking with Criterion, cache-conscious memory layouts, autovectorization, explicit SIMD intrinsics, and heap profiling with Valgrind/DHAT.

---

## 🔗 References
- ⬆️ Parent: [[Language]]

