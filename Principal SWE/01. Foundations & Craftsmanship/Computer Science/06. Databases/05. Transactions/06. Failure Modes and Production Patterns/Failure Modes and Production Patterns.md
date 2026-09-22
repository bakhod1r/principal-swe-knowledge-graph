---
title: "Failure Modes and Production Patterns"
tags:

  - computer-science
  - databases
  - distributed-systems
  - principal-swe
parent: "[[Transactions, ACID Axioms, and Isolation Levels]]"
---

# 📂 Failure Modes and Production Patterns

Deadlocks, contention, retries, application-level patterns, and anomaly testing.

---

## 🗂️ Topics

- [[Transactions, ACID Axioms, and Isolation Levels Failure Modes and Performance Optimization]] — Failure modes and tuning.
- [[Deadlocks, Detection, and Retry Strategies]] — Wait-for graphs, timeouts, serialization failure retries.
- [[Long-Running Transactions and Lock Contention]] — Bloat, blocking chains, idle in transaction.
- [[Application-Level Patterns (SELECT FOR UPDATE, Optimistic Versioning, Idempotency)]] — Preventing lost updates and write skew in app code.
- [[Testing Isolation Anomalies (Jepsen, Elle, Hermitage)]] — Empirically verifying isolation claims.

---

## 🔗 References
- ⬆️ Parent: [[Transactions, ACID Axioms, and Isolation Levels]]
