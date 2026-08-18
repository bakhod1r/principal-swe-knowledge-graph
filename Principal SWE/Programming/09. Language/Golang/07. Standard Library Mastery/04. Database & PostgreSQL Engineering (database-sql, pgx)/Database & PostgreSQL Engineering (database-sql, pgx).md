---
title: Database & PostgreSQL Engineering (database-sql, pgx)
tags:
  - golang
  - stdlib
  - principal-swe
parent: "[[Standard Library Mastery]]"
---

# Database & PostgreSQL Engineering (database-sql, pgx)

database/sql connection pooling, transactions, native pgx v5 driver architecture, pgxpool, batching, COPY protocol, and sqlc compiler.

```text
Database & PostgreSQL Engineering (database-sql, pgx)
│
├── [[database-sql Connection Pool Architecture]]
├── [[database-sql Transactions, Prepared Statements & Isolation]]
├── [[pgx Driver Architecture & Protocol vs database-sql]]
├── [[pgxpool Connection Pool Architecture & Tuning]]
├── [[High-Performance Batch Queries with pgx.Batch]]
├── [[Bulk Data Ingestion with pgx.CopyFrom (COPY Protocol)]]
├── [[Listen-Notify Asynchronous PubSub with pgx]]
├── [[PostgreSQL Custom Types & JSONB Encoding with pgx]]
├── [[Transactional Pipelines & Savepoints in pgx]]
└── [[sqlc with pgx Integration (Type-Safe SQL Compiler)]]
```

---

## 🗂️ Topics

- [[database-sql Connection Pool Architecture]] — Connection pooling mechanics: SetMaxOpenConns, SetMaxIdleConns, SetConnMaxLifetime, and driver connections.
- [[database-sql Transactions, Prepared Statements & Isolation]] — Managing ACID transactions with tx.BeginTx, context cancellation, prepared statement caching, and rollback defers.
- [[pgx Driver Architecture & Protocol vs database-sql]] — Native PostgreSQL binary wire protocol, zero-alloc type codecs, and comparing pgx native vs database/sql.
- [[pgxpool Connection Pool Architecture & Tuning]] — pgxpool.Pool tuning: MaxConns, MinConns, MaxConnLifetime, MaxConnIdleTime, and lifecycle health hooks.
- [[High-Performance Batch Queries with pgx.Batch]] — Sending pipelined multi-statement queries in a single network round-trip for massive throughput gains.
- [[Bulk Data Ingestion with pgx.CopyFrom (COPY Protocol)]] — Streaming bulk data ingestion directly into PostgreSQL tables bypassing slow INSERT statements.
- [[Listen-Notify Asynchronous PubSub with pgx]] — Real-time PostgreSQL event streaming and Change Data Capture (CDC) using LISTEN / NOTIFY in Go.
- [[PostgreSQL Custom Types & JSONB Encoding with pgx]] — Native decoding of UUID, Hstore, IPNet, and JSONB directly into Go structs without reflection overhead.
- [[Transactional Pipelines & Savepoints in pgx]] — Nested transactions with SQL SAVEPOINT, isolation levels (Serializable, RepeatableRead), and deferred rollbacks.
- [[sqlc with pgx Integration (Type-Safe SQL Compiler)]] — Generating compile-time type-safe Go repository code from raw SQL queries targeting native pgx/v5.

- [[Prepared Statement Caching & Query Planning Overhead]] — Automatic statement description caching in pgx, avoiding repeated parameter parsing and query plan generation on PostgreSQL.
- [[PostgreSQL Logical Replication & CDC in Go (pglogrepl)]] — Streaming PostgreSQL Write-Ahead Logs (WAL) via logical decoding plugins (pgoutput) using pglogrepl in Go.
- [[PgBouncer Connection Pooling & pgx Compatibility]] — Transaction pooling vs session pooling, statement naming conflicts (prefer_simple_protocol), and parameter status handling.
- [[Optimistic vs Pessimistic Locking (SELECT FOR UPDATE)]] — Implementing SELECT ... FOR UPDATE, SKIP LOCKED, row-level locks, and version columns in Go transactions.
- [[Zero-Downtime Database Migrations (golang-migrate)]] — Writing idempotent up/down migrations, transactional migration locks (pg_advisory_lock), and schema versioning.
- [[PostgreSQL Advisory Locks for Distributed Coordination]] — Utilizing pg_try_advisory_lock(key) and pg_advisory_unlock(key) for lightweight distributed mutexes.
- [[Keyset Pagination vs OFFSET Performance in Go]] — High-speed keyset querying (WHERE id > last_seen_id ORDER BY id LIMIT 50) avoiding O(n) OFFSET table scanning penalties.
- [[Database Circuit Breaking & Query Timeout Budgets]] — Enforcing hard deadline propagation (statement_timeout) across HTTP handlers, pool acquisitions, and query executions.
- [[Scanning Dynamic & NULL Values (sql.Null vs pgtype)]] — Comparing generic nullable types, pointer fields, and pgtype.Value zero-allocation scanners.

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
