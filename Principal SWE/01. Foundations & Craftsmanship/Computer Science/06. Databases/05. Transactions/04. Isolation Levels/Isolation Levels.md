---
title: "Isolation Levels"
tags:

  - computer-science
  - databases
  - distributed-systems
  - principal-swe
parent: "[[Transactions, ACID Axioms, and Isolation Levels]]"
---

# 📂 Isolation Levels

ANSI levels, their critique, Snapshot Isolation, SSI, and strict serializability.

---

## 🗂️ Topics

- [[ANSI SQL Isolation Levels and Their Critique (Berenson et al. 1995)]] — ANSI definitions, ambiguity, Generalized Isolation (Adya).
- [[Read Uncommitted and Read Committed]] — Weakest levels; statement-level snapshots.
- [[Repeatable Read (Vendor Semantics Differ)]] — Repeatable Read meaning in PostgreSQL vs MySQL.
- [[Snapshot Isolation]] — First-committer-wins, snapshot visibility, write skew.
- [[Serializable Snapshot Isolation (SSI)]] — rw-antidependency tracking, dangerous structures (Cahill).
- [[Strict Serializability vs Linearizability]] — Real-time ordering; Spanner, CockroachDB guarantees.

---

## 🔗 References
- ⬆️ Parent: [[Transactions, ACID Axioms, and Isolation Levels]]
