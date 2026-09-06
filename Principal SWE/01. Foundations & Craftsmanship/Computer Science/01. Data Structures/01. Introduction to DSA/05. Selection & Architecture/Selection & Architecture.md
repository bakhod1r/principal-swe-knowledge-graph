---
title: Selection & Architecture
tags:
  - computer-science
  - data-structures
  - principal-swe
parent: "[[Introduction to DSA]]"
---

# 📦 Selection & Architecture

Production engineering methodology for data structure selection based on empirical profiling and workloads.

```text
Selection & Architecture
│
├── [[Workload-Driven Data Structure Selection Framework]]
├── [[Introduction to DSA Real-World Applications and Systems Architecture]]
├── [[Concurrency and Thread-Safety in Data Structures]]
├── [[In-Memory Layouts vs Serialization and Wire Formats]]
├── [[Access Pattern Profiling (Read-Heavy, Write-Heavy, Scan-Heavy)]]
└── [[Data Structure Evolution and Zero-Downtime Migration Patterns]]
```

---

## 🗂️ Topics & Implementations

- [[Workload-Driven Data Structure Selection Framework]] — Step-by-step architectural decision trees matching access patterns to data structures.
- [[Introduction to DSA Real-World Applications and Systems Architecture]] — Mapping core data structures to production distributed systems and infrastructure.
- [[Concurrency and Thread-Safety in Data Structures]] — Lock contention, reader-writer locks, lock-free atomics, and hazard pointers.
- [[In-Memory Layouts vs Serialization and Wire Formats]] — FlatBuffers, Protobuf, and Cap'n Proto zero-copy deserialization vs in-memory pointer graphs.
- [[Access Pattern Profiling (Read-Heavy, Write-Heavy, Scan-Heavy)]] — Tailoring data layout to empirical hardware access profiles and query workloads.
- [[Data Structure Evolution and Zero-Downtime Migration Patterns]] — Upgrading running data topologies in high-availability online production services.

---

## 🔗 References
- ⬆️ Parent: [[Introduction to DSA]]
- 📚 Module: `Data Structures`
