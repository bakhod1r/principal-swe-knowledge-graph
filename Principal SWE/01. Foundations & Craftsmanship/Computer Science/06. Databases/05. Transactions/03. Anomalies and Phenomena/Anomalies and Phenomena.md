---
title: "Anomalies and Phenomena"
tags:

  - computer-science
  - databases
  - distributed-systems
  - principal-swe
parent: "[[Transactions, ACID Axioms, and Isolation Levels]]"
---

# 📂 Anomalies and Phenomena

Concurrency anomalies (P0-P4, A5A, A5B) that isolation levels permit or forbid.

---

## 🗂️ Topics

- [[Dirty Writes and Dirty Reads (P0, P1)]] — Reading or overwriting uncommitted data.
- [[Non-Repeatable Reads and Read Skew (P2, A5A)]] — Same row, different values within one transaction.
- [[Phantom Reads and Predicate Conflicts (P3)]] — Range queries see new rows.
- [[Lost Updates (P4)]] — Read-modify-write races.
- [[Write Skew (A5B)]] — Disjoint writes break shared invariant under SI.
- [[Read-Only Transaction Anomaly]] — Fekete read-only anomaly under Snapshot Isolation.

---

## 🔗 References
- ⬆️ Parent: [[Transactions, ACID Axioms, and Isolation Levels]]
