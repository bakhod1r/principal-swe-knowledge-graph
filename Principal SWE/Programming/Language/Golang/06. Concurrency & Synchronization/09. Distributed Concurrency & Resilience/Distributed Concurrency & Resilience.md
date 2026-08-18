---
title: Distributed Concurrency & Resilience
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Distributed Concurrency & Resilience

Circuit breakers, distributed locks, rate limiters, outbox patterns, sagas, and Raft consensus.

```text
Distributed Concurrency & Resilience
│
├── [[Circuit Breaker Pattern in Go]]
├── [[Distributed Rate Limiting (Redis Sliding Window & Token Bucket)]]
├── [[Distributed Lock (Redlock Algorithm & Etcd Lease Fencing)]]
├── [[Idempotency Keys & Deduplication]]
├── [[Transactional Outbox Pattern with PostgreSQL & Kafka]]
├── [[Transactional Inbox Pattern for Consumer Idempotency]]
├── [[Saga Pattern in Distributed Concurrency]]
├── [[Bulkhead Isolation Pattern]]
└── [[Distributed Consensus & Raft in Go (etcd-raft & hashicorp-raft)]]
```

---

## 🗂️ Topics

- [[Circuit Breaker Pattern in Go]] — State machine (Closed, Open, Half-Open) protecting downstream dependencies from cascading failures.
- [[Distributed Rate Limiting (Redis Sliding Window & Token Bucket)]] — Multi-instance cluster rate limiting using Redis and Lua script token buckets.
- [[Distributed Lock (Redlock Algorithm & Etcd Lease Fencing)]] — Safe distributed mutual exclusion with TTL lease renewal and fencing tokens.
- [[Idempotency Keys & Deduplication]] — Guaranteeing exactly-once processing of mutating requests using distributed caches.
- [[Transactional Outbox Pattern with PostgreSQL & Kafka]] — Atomic database write and message event publishing via outbox transaction tables.
- [[Transactional Inbox Pattern for Consumer Idempotency]] — Deduplicating incoming message delivery using persistent message ID tracking.
- [[Saga Pattern in Distributed Concurrency]] — Managing distributed transactions across microservices with compensating rollbacks.
- [[Bulkhead Isolation Pattern]] — Isolating critical resource pools to prevent complete application exhaustion.
- [[Distributed Consensus & Raft in Go (etcd-raft & hashicorp-raft)]] — Leader election, log replication, and distributed state machines in Go.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]

