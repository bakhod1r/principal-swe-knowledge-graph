---
title: Concurrent & Lock-Free Structures
tags:
  - computer-science
  - data-structures
  - concurrency
  - lock-free
  - principal-swe
parent: "[[Data Structures]]"
---

# 🔒 Concurrent & Lock-Free Structures

Hardware-atomic primitives, non-blocking queue/stack topologies, lock-free hash tables, and safe memory reclamation algorithms.

```text
Concurrent & Lock-Free Structures
│
├── [[Cas Atomic Primitives|01. Cas Atomic Primitives]]
├── [[Lock Free Queue Michael Scott|02. Lock Free Queue Michael Scott]]
├── [[Lock Free Stack|03. Lock Free Stack]]
├── [[Concurrent Hash Map|04. Concurrent Hash Map]]
├── [[Rcu|05. Rcu]]
└── [[Hazard Pointers|06. Hazard Pointers]]
```

---

## 🗂️ Topics

- 📂 [[Cas Atomic Primitives|01. Cas Atomic Primitives]] — Hardware Compare-And-Swap (CAS), LL/SC, atomic memory orderings, and ABA prevention.
- 📂 [[Lock Free Queue Michael Scott|02. Lock Free Queue Michael Scott]] — Non-blocking FIFO queue using sentinel nodes and concurrent tail advancement via CAS.
- 📂 [[Lock Free Stack|03. Lock Free Stack]] — Treiber stack utilizing atomic CAS top-pointer updates and backoff elimination under contention.
- 📂 [[Concurrent Hash Map|04. Concurrent Hash Map]] — Striped bucket locking, lock-free split-ordered lists, and lock-free resizing tables.
- 📂 [[Rcu|05. Rcu]] — Read-Copy-Update lock-free read synchronization with quiescent grace period deferred reclamation.
- 📂 [[Hazard Pointers|06. Hazard Pointers]] — Thread-local hazard registers protecting concurrent pointer reads against concurrent deallocation.

---

## 🔗 References
- ⬆️ Parent: [[Data Structures]]
- 📚 Module: `Data Structures`
