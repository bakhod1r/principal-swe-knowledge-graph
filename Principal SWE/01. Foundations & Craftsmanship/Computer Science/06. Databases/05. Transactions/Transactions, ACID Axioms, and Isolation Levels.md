---
title: Transactions, ACID Axioms, and Isolation Levels
tags:

  - computer-science
  - databases
  - distributed-systems
  - principal-swe
parent: "[[Databases]]"
---

# 📦 Transactions, ACID Axioms, and Isolation Levels

Atomicity, Consistency, Isolation, Durability, ANSI SQL isolation levels, Snapshot Isolation, Write Skew, and Phantom Reads.

```text
Transactions, ACID Axioms, and Isolation Levels
│
├── 01. Transaction Fundamentals
│   ├── [[Transaction Fundamentals]]
│   ├── [[Transactions, ACID Axioms, and Isolation Levels Core Principles and Mechanics]]
│   ├── [[Transaction Lifecycle (BEGIN, COMMIT, ROLLBACK, Savepoints)]]
│   ├── [[Transaction Boundaries, Autocommit, and Session Semantics]]
│   ├── [[Serializability Theory (Schedules, Conflict and View Serializability)]]
│   └── [[Recoverability, Cascading Aborts, and Strict Schedules]]
│
├── 02. ACID Axioms
│   ├── [[ACID Axioms]]
│   ├── [[Atomicity (All-or-Nothing and Undo Semantics)]]
│   ├── [[Consistency (Invariants, Constraints, and Application Responsibility)]]
│   ├── [[Isolation (The Concurrency Illusion)]]
│   ├── [[Durability (fsync, WAL, and Commit Guarantees)]]
│   └── [[ACID vs BASE and Eventual Consistency]]
│
├── 03. Anomalies and Phenomena
│   ├── [[Anomalies and Phenomena]]
│   ├── [[Dirty Writes and Dirty Reads (P0, P1)]]
│   ├── [[Non-Repeatable Reads and Read Skew (P2, A5A)]]
│   ├── [[Phantom Reads and Predicate Conflicts (P3)]]
│   ├── [[Lost Updates (P4)]]
│   ├── [[Write Skew (A5B)]]
│   └── [[Read-Only Transaction Anomaly]]
│
├── 04. Isolation Levels
│   ├── [[Isolation Levels]]
│   ├── [[ANSI SQL Isolation Levels and Their Critique (Berenson et al. 1995)]]
│   ├── [[Read Uncommitted and Read Committed]]
│   ├── [[Repeatable Read (Vendor Semantics Differ)]]
│   ├── [[Snapshot Isolation]]
│   ├── [[Serializable Snapshot Isolation (SSI)]]
│   └── [[Strict Serializability vs Linearizability]]
│
├── 05. Implementation Techniques
│   ├── [[Implementation Techniques]]
│   ├── [[Transactions, ACID Axioms, and Isolation Levels Storage Architecture and Implementation]]
│   ├── [[Lock-Based Isolation (2PL, Predicate, Gap and Next-Key Locks)]]
│   ├── [[MVCC-Based Isolation and Snapshot Visibility]]
│   ├── [[Optimistic Concurrency Control (OCC) and Validation]]
│   └── [[Isolation Implementations in PostgreSQL, MySQL InnoDB, Oracle, and SQL Server]]
│
└── 06. Failure Modes and Production Patterns
    ├── [[Failure Modes and Production Patterns]]
    ├── [[Transactions, ACID Axioms, and Isolation Levels Failure Modes and Performance Optimization]]
    ├── [[Deadlocks, Detection, and Retry Strategies]]
    ├── [[Long-Running Transactions and Lock Contention]]
    ├── [[Application-Level Patterns (SELECT FOR UPDATE, Optimistic Versioning, Idempotency)]]
    └── [[Testing Isolation Anomalies (Jepsen, Elle, Hermitage)]]
```

---

## 🗂️ Core Knowledge Domains

### 1. 📂 [[Transaction Fundamentals|01. Transaction Fundamentals]]
- [[Transaction Fundamentals]] — Master topology: Transaction abstraction, lifecycle, boundaries, and serializability theory.
- [[Transactions, ACID Axioms, and Isolation Levels Core Principles and Mechanics]] — Core transaction mechanics and guarantees.
- [[Transaction Lifecycle (BEGIN, COMMIT, ROLLBACK, Savepoints)]] — States, commit/abort paths, nested savepoints.
- [[Transaction Boundaries, Autocommit, and Session Semantics]] — Implicit vs explicit transactions, driver autocommit, scope bugs.
- [[Serializability Theory (Schedules, Conflict and View Serializability)]] — Schedules, precedence graphs, conflict vs view equivalence.
- [[Recoverability, Cascading Aborts, and Strict Schedules]] — Recoverable, ACA, strict schedules.

### 2. 📂 [[ACID Axioms|02. ACID Axioms]]
- [[ACID Axioms]] — Master topology: Formal meaning of each ACID property and where guarantees actually come from.
- [[Atomicity (All-or-Nothing and Undo Semantics)]] — Abort handling, undo logs, partial failure.
- [[Consistency (Invariants, Constraints, and Application Responsibility)]] — DB constraints vs app invariants; the weakest letter.
- [[Isolation (The Concurrency Illusion)]] — Concurrent transactions appear serial; spectrum of weakness.
- [[Durability (fsync, WAL, and Commit Guarantees)]] — Commit acknowledgement, fsync, group commit, replication durability.
- [[ACID vs BASE and Eventual Consistency]] — ACID vs BASE trade-offs in distributed stores.

### 3. 📂 [[Anomalies and Phenomena|03. Anomalies and Phenomena]]
- [[Anomalies and Phenomena]] — Master topology: Concurrency anomalies (P0-P4, A5A, A5B) that isolation levels permit or forbid.
- [[Dirty Writes and Dirty Reads (P0, P1)]] — Reading or overwriting uncommitted data.
- [[Non-Repeatable Reads and Read Skew (P2, A5A)]] — Same row, different values within one transaction.
- [[Phantom Reads and Predicate Conflicts (P3)]] — Range queries see new rows.
- [[Lost Updates (P4)]] — Read-modify-write races.
- [[Write Skew (A5B)]] — Disjoint writes break shared invariant under SI.
- [[Read-Only Transaction Anomaly]] — Fekete read-only anomaly under Snapshot Isolation.

### 4. 📂 [[Isolation Levels|04. Isolation Levels]]
- [[Isolation Levels]] — Master topology: ANSI levels, their critique, Snapshot Isolation, SSI, and strict serializability.
- [[ANSI SQL Isolation Levels and Their Critique (Berenson et al. 1995)]] — ANSI definitions, ambiguity, Generalized Isolation (Adya).
- [[Read Uncommitted and Read Committed]] — Weakest levels; statement-level snapshots.
- [[Repeatable Read (Vendor Semantics Differ)]] — Repeatable Read meaning in PostgreSQL vs MySQL.
- [[Snapshot Isolation]] — First-committer-wins, snapshot visibility, write skew.
- [[Serializable Snapshot Isolation (SSI)]] — rw-antidependency tracking, dangerous structures (Cahill).
- [[Strict Serializability vs Linearizability]] — Real-time ordering; Spanner, CockroachDB guarantees.

### 5. 📂 [[Implementation Techniques|05. Implementation Techniques]]
- [[Implementation Techniques]] — Master topology: How engines enforce isolation: locks, MVCC, OCC, and vendor specifics.
- [[Transactions, ACID Axioms, and Isolation Levels Storage Architecture and Implementation]] — Storage-level implementation of transactions.
- [[Lock-Based Isolation (2PL, Predicate, Gap and Next-Key Locks)]] — Pessimistic enforcement, range locking.
- [[MVCC-Based Isolation and Snapshot Visibility]] — Version chains, xmin/xmax, visibility rules.
- [[Optimistic Concurrency Control (OCC) and Validation]] — Read-validate-write phases, abort rates.
- [[Isolation Implementations in PostgreSQL, MySQL InnoDB, Oracle, and SQL Server]] — Vendor behaviour comparison.

### 6. 📂 [[Failure Modes and Production Patterns|06. Failure Modes and Production Patterns]]
- [[Failure Modes and Production Patterns]] — Master topology: Deadlocks, contention, retries, application-level patterns, and anomaly testing.
- [[Transactions, ACID Axioms, and Isolation Levels Failure Modes and Performance Optimization]] — Failure modes and tuning.
- [[Deadlocks, Detection, and Retry Strategies]] — Wait-for graphs, timeouts, serialization failure retries.
- [[Long-Running Transactions and Lock Contention]] — Bloat, blocking chains, idle in transaction.
- [[Application-Level Patterns (SELECT FOR UPDATE, Optimistic Versioning, Idempotency)]] — Preventing lost updates and write skew in app code.
- [[Testing Isolation Anomalies (Jepsen, Elle, Hermitage)]] — Empirically verifying isolation claims.

---

## 🔗 References
- ⬆️ Parent: [[Databases]]
- 📚 Module: `Computer Science`
