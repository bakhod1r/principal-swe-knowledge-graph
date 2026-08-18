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

---

## 🔗 References
- ⬆️ Parent: [[Standard Library Mastery]]
- 🎓 Root: [[Principal SWE]]
