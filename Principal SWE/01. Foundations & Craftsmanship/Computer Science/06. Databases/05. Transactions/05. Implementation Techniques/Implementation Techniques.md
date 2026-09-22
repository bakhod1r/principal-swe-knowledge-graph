---
title: "Implementation Techniques"
tags:

  - computer-science
  - databases
  - distributed-systems
  - principal-swe
parent: "[[Transactions, ACID Axioms, and Isolation Levels]]"
---

# 📂 Implementation Techniques

How engines enforce isolation: locks, MVCC, OCC, and vendor specifics.

---

## 🗂️ Topics

- [[Transactions, ACID Axioms, and Isolation Levels Storage Architecture and Implementation]] — Storage-level implementation of transactions.
- [[Lock-Based Isolation (2PL, Predicate, Gap and Next-Key Locks)]] — Pessimistic enforcement, range locking.
- [[MVCC-Based Isolation and Snapshot Visibility]] — Version chains, xmin/xmax, visibility rules.
- [[Optimistic Concurrency Control (OCC) and Validation]] — Read-validate-write phases, abort rates.
- [[Isolation Implementations in PostgreSQL, MySQL InnoDB, Oracle, and SQL Server]] — Vendor behaviour comparison.

---

## 🔗 References
- ⬆️ Parent: [[Transactions, ACID Axioms, and Isolation Levels]]
