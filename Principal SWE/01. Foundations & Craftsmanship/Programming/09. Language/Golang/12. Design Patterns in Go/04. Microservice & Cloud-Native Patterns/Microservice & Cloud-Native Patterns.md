---
title: Microservice & Cloud-Native Patterns
tags:
  - hub
parent: "[[Design Patterns in Go]]"
---

- [[Event Sourcing & CQRS in Go]] — Command and Query responsibility segregation with append-only event stores and projections.

- [[Transactional Inbox & Outbox Orchestration]] — Guaranteeing idempotent consumer processing and exactly-once semantics.

---
title: Microservice & Cloud-Native Patterns
tags:
  - golang
  - design-patterns
  - principal-swe
parent: "`Design Patterns in Go`"
---

# Microservice & Cloud-Native Patterns

Distributed systems design patterns, saga orchestration, outbox, and cloud-native resilience.

```text
Microservice & Cloud-Native Patterns
│
├── [[Outbox Pattern]]
├── [[Saga Pattern (Orchestration vs Choreography)]]
├── [[Dead Letter Queue (DLQ) Pattern]]
├── [[Idempotent Consumer Pattern]]
├── [[Sidecar Communication Pattern]]
└── [[Graceful Degradation Pattern]]
```

---

## 🗂️ Topics

- [[Outbox Pattern]] — Guaranteed at-least-once message publishing using relational database transaction logs.
- [[Saga Pattern (Orchestration vs Choreography)]] — Managing distributed multi-service transactions with compensating rollback actions.
- [[Dead Letter Queue (DLQ) Pattern]] — Isolating unprocessable poison messages for inspection and replay.
- [[Idempotent Consumer Pattern]] — Deduplicating incoming message delivery using persistent transaction IDs.
- [[Sidecar Communication Pattern]] — Interacting with local Envoy/Dapr sidecars over gRPC/UDS Unix Domain Sockets.
- [[Graceful Degradation Pattern]] — Serving stale cached data or partial responses during downstream dependency outages.

---

## 🔗 References
- ⬆️ Parent: `Design Patterns in Go`

