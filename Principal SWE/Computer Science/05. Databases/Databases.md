---
title: Databases
tags:
  - computer-science
  - databases
  - distributed-systems
  - principal-swe
parent: "[[Computer Science]]"
---

# 🗄️ Databases (Unified Systems Architecture & Theoretical Foundations)

Comprehensive, production-grade master architecture covering the complete spectrum of database systems: relational calculus and normalization, ANSI SQL standards, B+ Tree and LSM-Tree storage engines, buffer pool page caches, ACID transactions and serializable isolation, 2PL lock managers, MVCC visibility, ARIES WAL recovery, multi-dimensional indexing (B-Tree, GIN, GiST, BRIN), Cost-Based Query Optimization (CBO), SIMD vectorized columnar execution, distributed replication and sharding, consensus protocols (Raft, Paxos, Spanner), 2PC distributed transactions, CAP/PACELC consistency models, CRDTs, in-memory caching & Redis architectures, PostgreSQL DBA mastery, Elasticsearch full-text search, NoSQL document models & MongoDB, and enterprise disaster recovery across 26 specialized subdomains.

```text
Databases
│
├── [[Relational Model and Database Foundations|01. Relational Model and Foundations]]
├── [[SQL Standards, Ddl, Dml, and Advanced Dialects|02. SQL Standards DDL DML and Dialects]]
├── [[Storage Engines (b Tree and LSM Tree Internals)|03. Storage Engines B Tree and LSM Tree]]
├── [[Buffer Pool and Database Page Cache Management|04. Buffer Pool and Page Cache]]
├── [[Transactions, ACID Axioms, and Isolation Levels|05. Transactions ACID and Isolation Levels]]
├── [[Locking and Concurrency Control (two Phase Locking 2pl)|06. Locking and Concurrency Control 2PL]]
├── [[MVCC Internals and Vacuum Garbage Collection|07. MVCC Internals and Vacuum Gc]]
├── [[Write Ahead Logging (wal) and ARIES Crash Recovery|08. Write Ahead Logging and ARIES Recovery]]
├── [[Database Indexing Structures (b Tree, Gin, Gist, Brin)|09. Indexing Structures and Search]]
├── [[Query Planning, Relational Algebra, and Cost Based Optimization|10. Query Planning and Cost Based Optimization]]
├── [[Columnar Storage and Vectorized Execution Engines|11. Columnar Storage and Vectorized Execution]]
├── [[Distributed Database Replication and Topologies|12. Distributed Replication and Topologies]]
├── [[Database Partitioning, Sharding, and Federation|13. Partitioning Sharding and Federation]]
├── [[Distributed Consensus in Databases (raft, Paxos, Spanner)|14. Distributed Consensus in Databases]]
├── [[Distributed Transactions (2pc, 3pc, and Sagas)|15. Distributed Transactions 2PC and Sagas]]
├── [[CAP Theorem, Pacelc, and Distributed Consistency Models|16. CAP PACELC and Distributed Consistency]]
├── [[CRDTs and Multi Master Causality Tracking|17. CRDTs and Multi Master Causality]]
├── [[In Memory Data Stores and Caching Topologies|18. In Memory Data Stores and Caching]]
├── [[Document and Key Value NoSQL Models|19. Document and Key Value NoSQL Models]]
├── [[Full Text Search and Vector Database Architectures|20. Full Text Search and Vector Databases]]
├── [[OLTP vs OLAP and Data Warehousing Architectures|21. OLTP vs OLAP and Data Warehousing]]
├── [[Database Connection Pooling, High Availability, and Disaster Recovery|22. Connection Pooling HA and Disaster Recovery]]
├── [[Postgresql Mastery & DBA|23. PostgreSQL Mastery & DBA]]
├── [[Redis & in Memory Architectures|24. Redis & In-Memory Architectures]]
├── [[Elasticsearch & Distributed Search|25. Elasticsearch & Distributed Search]]
└── [[Mongodb & Document Stores|26. MongoDB & Document Stores]]
```

---

## 🗂️ Core Database Domains

- 📂 [[Relational Model and Database Foundations|01. Relational Model and Foundations]]
- 📂 [[SQL Standards, Ddl, Dml, and Advanced Dialects|02. SQL Standards DDL DML and Dialects]]
- 📂 [[Storage Engines (b Tree and LSM Tree Internals)|03. Storage Engines B Tree and LSM Tree]]
- 📂 [[Buffer Pool and Database Page Cache Management|04. Buffer Pool and Page Cache]]
- 📂 [[Transactions, ACID Axioms, and Isolation Levels|05. Transactions ACID and Isolation Levels]]
- 📂 [[Locking and Concurrency Control (two Phase Locking 2pl)|06. Locking and Concurrency Control 2PL]]
- 📂 [[MVCC Internals and Vacuum Garbage Collection|07. MVCC Internals and Vacuum Gc]]
- 📂 [[Write Ahead Logging (wal) and ARIES Crash Recovery|08. Write Ahead Logging and ARIES Recovery]]
- 📂 [[Database Indexing Structures (b Tree, Gin, Gist, Brin)|09. Indexing Structures and Search]]
- 📂 [[Query Planning, Relational Algebra, and Cost Based Optimization|10. Query Planning and Cost Based Optimization]]
- 📂 [[Columnar Storage and Vectorized Execution Engines|11. Columnar Storage and Vectorized Execution]]
- 📂 [[Distributed Database Replication and Topologies|12. Distributed Replication and Topologies]]
- 📂 [[Database Partitioning, Sharding, and Federation|13. Partitioning Sharding and Federation]]
- 📂 [[Distributed Consensus in Databases (raft, Paxos, Spanner)|14. Distributed Consensus in Databases]]
- 📂 [[Distributed Transactions (2pc, 3pc, and Sagas)|15. Distributed Transactions 2PC and Sagas]]
- 📂 [[CAP Theorem, Pacelc, and Distributed Consistency Models|16. CAP PACELC and Distributed Consistency]]
- 📂 [[CRDTs and Multi Master Causality Tracking|17. CRDTs and Multi Master Causality]]
- 📂 [[In Memory Data Stores and Caching Topologies|18. In Memory Data Stores and Caching]]
- 📂 [[Document and Key Value NoSQL Models|19. Document and Key Value NoSQL Models]]
- 📂 [[Full Text Search and Vector Database Architectures|20. Full Text Search and Vector Databases]]
- 📂 [[OLTP vs OLAP and Data Warehousing Architectures|21. OLTP vs OLAP and Data Warehousing]]
- 📂 [[Database Connection Pooling, High Availability, and Disaster Recovery|22. Connection Pooling HA and Disaster Recovery]]
- 📂 [[Postgresql Mastery & DBA|23. PostgreSQL Mastery & DBA]]
- 📂 [[Redis & in Memory Architectures|24. Redis & In-Memory Architectures]]
- 📂 [[Elasticsearch & Distributed Search|25. Elasticsearch & Distributed Search]]
- 📂 [[Mongodb & Document Stores|26. MongoDB & Document Stores]]

---

## 🔗 References
- ⬆️ Parent: [[Computer Science]]
- 🎓 Root: [[Principal SWE]]
