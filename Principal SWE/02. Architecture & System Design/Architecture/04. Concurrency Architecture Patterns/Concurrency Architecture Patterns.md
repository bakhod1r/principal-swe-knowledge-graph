---
title: Concurrency Architecture Patterns
tags:
  - architecture
  - systems-architecture
  - concurrency-and-high-performance-design-patterns
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Concurrency & High Performance Design Patterns

High-throughput and low-latency concurrency patterns: Reactor, Proactor, Active Object, Monitor Object, Half-Sync/Half-Async, Leader-Followers, Worker Pool, Disruptor Ring Buffer, and non-blocking synchronization.

```text
Concurrency & High Performance Design Patterns
│
├── [[Reactor Pattern: Synchronous Event Demultiplexing and Service Handlers|01. Reactor Pattern and Event Demultiplexing]]
├── `02. Proactor Pattern and Asynchronous I O Completion`
├── [[Active Object Pattern: Decoupling Method Execution From Invocation|03. Active Object Pattern and Asynchronous Method Invocation]]
├── [[Monitor Object Pattern: Synchronizing Concurrent Access to Object State|04. Monitor Object and Thread Synchronization]]
├── [[Half Sync - Half Async Pattern: Bridging Asynchronous and Synchronous Services|05. Half Sync Half Async Architecture]]
├── `06. Leader Followers Thread Pool Pattern`
├── `07. Worker Pool and Task Queue Paradigms`
├── [[Double Checked Locking Pattern and Memory Barrier Invariants|08. Double Checked Locking and Thread Safe Lazy Initialization]]
├── [[Guarded Suspension and Balking: Conditional Execution and Early Exit|09. Guarded Suspension and Balking Patterns]]
└── `10. Lmax Disruptor Ring Buffer and Mechanical Sympathy`
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Reactor Pattern: Synchronous Event Demultiplexing and Service Handlers|01. Reactor Pattern and Event Demultiplexing]] — Event loop architecture (Node.js, Netty, Nginx), synchronous demultiplexing with `epoll`/`kqueue`, non-blocking I/O event dispatching, and handler execution.
- 📂 `02. Proactor Pattern and Asynchronous I O Completion` — Initiating asynchronous I/O operations (Windows IOCP, Linux io_uring), operating system kernel completion notification, and completion handler callbacks.
- 📂 [[Active Object Pattern: Decoupling Method Execution From Invocation|03. Active Object Pattern and Asynchronous Method Invocation]] — Client thread proxy, activation list queue, scheduler thread, servant object execution, and returning asynchronous Future promises.
- 📂 [[Monitor Object Pattern: Synchronizing Concurrent Access to Object State|04. Monitor Object and Thread Synchronization]] — Encapsulating shared state with critical sections, monitor locks, condition variables, wait-notify synchronization, and preventing race conditions.
- 📂 [[Half Sync - Half Async Pattern: Bridging Asynchronous and Synchronous Services|05. Half Sync Half Async Architecture]] — Decoupling high-concurrency async network I/O layers from blocking synchronous database/business processing layers via thread-safe queues.
- 📂 `06. Leader Followers Thread Pool Pattern` — One thread acts as the leader waiting for incoming I/O events while followers wait on a synchronization condition, zero-overhead handoffs, and throughput.
- 📂 `07. Worker Pool and Task Queue Paradigms` — Fixed vs elastic worker pools, backpressure handling on queue saturation, work rejection policies, and thread pool starvation prevention.
- 📂 [[Double Checked Locking Pattern and Memory Barrier Invariants|08. Double Checked Locking and Thread Safe Lazy Initialization]] — Volatile read-write semantics, preventing uninitialized memory reads across CPU cores, and modern language-specific idioms (Go `sync.Once`, Java enum).
- 📂 [[Guarded Suspension and Balking: Conditional Execution and Early Exit|09. Guarded Suspension and Balking Patterns]] — Suspending thread execution until a precondition is met (Guarded Suspension), and immediately returning when object state is inappropriate (Balking).
- 📂 `10. Lmax Disruptor Ring Buffer and Mechanical Sympathy` — Lock-free bounded ring buffers, avoiding cache line false sharing via padding, memory sequence barriers, and single-writer principle for ultra-low latency.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

