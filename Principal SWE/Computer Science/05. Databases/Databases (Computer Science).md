---
title: Databases (Computer Science)
tags:
  - computer-science
  - databases-(computer-science)
  - principal-swe
parent: "[[Computer Science]]"
---

# 🏛️ Databases (Computer Science) (Foundations & Systems Architecture)

Database internal architectures: storage engine layouts (LSM vs B-Tree), MVCC transaction isolation, WAL crash recovery (ARIES), query optimizer cost models, consensus-backed distributed databases, and vectorized engines.

```text
Databases (Computer Science)
│
├── [[Storage Engines (lsm Tree vs B Plus Tree)|01. Storage Engines LSM vs B Tree]]
├── [[Multi Version Concurrency Control (mvcc)|02. MVCC Internals]]
├── [[Write Ahead Logging (wal) and Aries Recovery|03. Write Ahead Logging and Recovery]]
├── [[Relational Query Planning and Optimization|04. Query Planner and Optimizer]]
├── [[Consensus Protocols in Distributed Databases|05. Consensus in Distributed Databases]]
├── [[CRDTs and Causality Tracking in Databases|06. CRDTs and Vector Clocks]]
├── [[Distributed Transactions (2pc, 3pc, and Sagas)|07. Distributed Transactions 2PC and Sagas]]
├── [[Buffer Pool and Database Page Cache Management|08. Buffer Pool and Page Cache]]
├── [[Columnar Storage and Vectorized Query Engines|09. Columnar and Vectorized Execution]]
└── [[Cost Based Query Optimization (cbo) and Statistics|10. Cost Based Optimization and Statistics]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Storage Engines (lsm Tree vs B Plus Tree)|01. Storage Engines LSM vs B Tree]] — Write-optimized append-only MemTable/SSTable cascades vs read-optimized balanced B+ Tree disk page layouts.
- 📂 [[Multi Version Concurrency Control (mvcc)|02. MVCC Internals]] — Snapshot isolation, undo/redo logs, tuple transaction headers (xmin, xmax), vacuuming, and write skew.
- 📂 [[Write Ahead Logging (wal) and Aries Recovery|03. Write Ahead Logging and Recovery]] — Sequential log-structured disk writes, checkpoints, and ARIES 3-phase crash recovery (Analysis, Redo, Undo).
- 📂 [[Relational Query Planning and Optimization|04. Query Planner and Optimizer]] — AST parsing, logical plan generation, relational algebra rewrite rules, and cost-based plan enumeration.
- 📂 [[Consensus Protocols in Distributed Databases|05. Consensus in Distributed Databases]] — Raft and Multi-Paxos replication groups, lease-based leader reads, and split-brain quorum protection.
- 📂 [[CRDTs and Causality Tracking in Databases|06. CRDTs and Vector Clocks]] — State-based vs Operation-based CRDTs, vector clocks, Lamport timestamps, and conflict-free multi-master writes.
- 📂 [[Distributed Transactions (2pc, 3pc, and Sagas)|07. Distributed Transactions 2PC and Sagas]] — Two-Phase Commit coordinator protocol, blocking hazards, and compensating asynchronous Saga workflows.
- 📂 [[Buffer Pool and Database Page Cache Management|08. Buffer Pool and Page Cache]] — LRU-2, 2Q, Clock page replacement, dirty page flushing, double buffering avoidance, and direct I/O (O_DIRECT).
- 📂 [[Columnar Storage and Vectorized Query Engines|09. Columnar and Vectorized Execution]] — Parquet/ORC columnar compression (RLE, Snappy), SIMD vectorized batch evaluation, and ClickHouse execution.
- 📂 [[Cost Based Query Optimization (cbo) and Statistics|10. Cost Based Optimization and Statistics]] — HyperLogLog cardinality estimation, histograms, join ordering DP (System R), and selectivity calculations.

---

## 🔗 References
- ⬆️ Parent: [[Computer Science]]
- 🎓 Root: [[Principal SWE]]
