- [[Distributed Consensus & Raft in Go]] — Leader election, log replication, and state machine replication with hashicorp-raft and etcd-raft.

---
title: Distributed Concurrency & Resilience
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Distributed Concurrency & Resilience

Distributed concurrency patterns, circuit breakers, rate limiters, and outbox patterns.

```text
Distributed Concurrency & Resilience
│
├── [[Circuit Breaker Pattern in Go]]
├── [[Distributed Rate Limiting (Redis & Token Bucket)]]
├── [[Idempotency Keys & Deduplication]]
├── [[Distributed Lock (Redlock & Etcd)]]
├── [[Bulkhead Pattern]]
└── [[Outbox Pattern in Go]]
```

---

## 🗂️ Topics

- [[Circuit Breaker Pattern in Go]] — State machine (Closed, Open, Half-Open) protecting downstream services from cascading failure.
- [[Distributed Rate Limiting (Redis & Token Bucket)]] — Multi-instance rate limiting with Redis sliding window and token buckets.
- [[Idempotency Keys & Deduplication]] — Preventing duplicate execution of concurrent mutation requests.
- [[Distributed Lock (Redlock & Etcd)]] — Safe distributed mutual exclusion with TTL lease renewal and fencing tokens.
- [[Bulkhead Pattern]] — Isolating critical resource pools to prevent total system saturation.
- [[Outbox Pattern in Go]] — Transactional database outbox pattern ensuring atomic database write and message publish.

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
