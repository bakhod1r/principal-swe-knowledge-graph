---
title: SQL
tags:
  - programming
  - sql
  - language-mastery
  - principal-swe
parent: "[[Language]]"
---

# 💻 SQL (Production Mastery & Engineering Deep Dive)

Declarative SQL relational querying, schema engineering, and performance optimization: ANSI SQL standards, DDL/DML/DCL/TCL, multi-table JOIN relational algebra, Window functions, Recursive CTEs, query planner cost models (EXPLAIN ANALYZE), B-Tree indexing, transactions & isolation levels, and JSON/JSONB document querying.

```text
SQL
│
├── [[Relational Database Concepts and Relational Theory|01. Relational Database Concepts and Theory]]
├── [[SQL Syntax Fundamentals and Universal Data Types|02. Basic SQL Syntax and Data Types]]
├── [[Data Definition Language (ddl) and Schema Migrations|03. Data Definition Language DDL]]
├── [[Data Manipulation Language (dml) and Data Modification|04. Data Manipulation Language DML]]
├── [[SQL Aggregate Queries, Grouping, and Filtering|05. Aggregate Queries and Grouping]]
├── [[SQL Data Integrity, Constraints, and Invariants|06. Data Integrity and Constraints]]
├── [[SQL Subqueries and Common Table Expressions (ctes)|07. Subqueries and CTEs]]
├── [[SQL Join Queries and Relational Set Operations|08. Join Queries and Relational Algebra]]
├── [[SQL Built in Functions and Transformations|09. Advanced SQL Functions]]
├── [[Database Views and Materialized View Refresh|10. Database Views and Materialized Views]]
├── [[SQL Indexing Strategies and B Tree Query Acceleration|11. Indexing Strategies and Query Execution]]
├── [[SQL Transactions and ACID Concurrency Control|12. Transactions and ACID Guarantees]]
├── [[SQL Security, Grants, and Row Level Security (rls)|13. Data Security and Access Control]]
├── [[SQL Stored Procedures, Functions, and Triggers|14. Stored Procedures and Triggers]]
├── [[SQL Performance Optimization and Query Plan Tuning|15. SQL Performance Optimization and Tuning]]
└── [[Advanced SQL Window Functions and JSONB Operations|16. Advanced SQL Window Functions and JSON]]
```

---

## 🗂️ Core Knowledge Pillars

- 📂 [[Relational Database Concepts and Relational Theory|01. Relational Database Concepts and Theory]] — Relational model fundamentals, Codd's rules, primary keys, foreign keys, and relational schema design.
- 📂 [[SQL Syntax Fundamentals and Universal Data Types|02. Basic SQL Syntax and Data Types]] — Standard SQL keywords, numeric, string, boolean, temporal data types, and comparison operators.
- 📂 [[Data Definition Language (ddl) and Schema Migrations|03. Data Definition Language DDL]] — CREATE TABLE, ALTER TABLE, DROP TABLE, TRUNCATE, and declarative schema migration workflows.
- 📂 [[Data Manipulation Language (dml) and Data Modification|04. Data Manipulation Language DML]] — INSERT, UPDATE, DELETE, MERGE/UPSERT (ON CONFLICT DO UPDATE), and batch data loading.
- 📂 [[SQL Aggregate Queries, Grouping, and Filtering|05. Aggregate Queries and Grouping]] — COUNT, SUM, AVG, MIN, MAX, GROUP BY, HAVING, and multi-column aggregation mechanics.
- 📂 [[SQL Data Integrity, Constraints, and Invariants|06. Data Integrity and Constraints]] — PRIMARY KEY, FOREIGN KEY (CASCADE/RESTRICT), UNIQUE, NOT NULL, and custom CHECK constraints.
- 📂 [[SQL Subqueries and Common Table Expressions (ctes)|07. Subqueries and CTEs]] — Scalar subqueries, correlated subqueries, WITH CTEs, and Recursive CTEs for hierarchical tree traversal.
- 📂 [[SQL Join Queries and Relational Set Operations|08. Join Queries and Relational Algebra]] — INNER JOIN, LEFT/RIGHT/FULL OUTER JOIN, CROSS JOIN, LATERAL JOIN, and UNION/INTERSECT/EXCEPT.
- 📂 [[SQL Built in Functions and Transformations|09. Advanced SQL Functions]] — String manipulation, Date/Time math, CAST/CONVERT, COALESCE, NULLIF, and CASE WHEN expressions.
- 📂 [[Database Views and Materialized View Refresh|10. Database Views and Materialized Views]] — Standard virtual views, Materialized views, concurrent refresh, and query rewriting.
- 📂 [[SQL Indexing Strategies and B Tree Query Acceleration|11. Indexing Strategies and Query Execution]] — B-Tree indexes, composite multi-column indexes, partial indexes, covering indexes, and index selectivity.
- 📂 [[SQL Transactions and ACID Concurrency Control|12. Transactions and ACID Guarantees]] — BEGIN, COMMIT, ROLLBACK, SAVEPOINT, and transaction isolation levels (Read Committed, Serializable).
- 📂 [[SQL Security, Grants, and Row Level Security (rls)|13. Data Security and Access Control]] — GRANT, REVOKE, role hierarchies, principle of least privilege, and Row-Level Security policies.
- 📂 [[SQL Stored Procedures, Functions, and Triggers|14. Stored Procedures and Triggers]] — Procedural extensions (PL/pgSQL), stored functions, execution contexts, and automated audit triggers.
- 📂 [[SQL Performance Optimization and Query Plan Tuning|15. SQL Performance Optimization and Tuning]] — EXPLAIN ANALYZE inspection, Cost-Based Optimizer (CBO), avoiding table scans, and SARGable predicates.
- 📂 [[Advanced SQL Window Functions and JSONB Operations|16. Advanced SQL Window Functions and JSON]] — ROW_NUMBER, RANK, DENSE_RANK, LAG, LEAD, SUM OVER PARTITION, and JSON/JSONB indexing.

---

## 🔗 References
- ⬆️ Parent: [[Language]]

