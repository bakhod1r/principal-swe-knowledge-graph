---
title: Event Driven Programming
tags:
  - programming
  - paradigms
  - principal-swe
parent: "[[Programming Paradigms]]"
---

# Event Driven Programming

EDA broker vs mediator topologies, Event Sourcing, CQRS, delivery semantics (at-least-once, exactly-once), and CDC Outbox.

```text
Event Driven Programming
│
├── [[Event-Driven Architecture (EDA) Core Topology: Broker vs Mediator]]
├── [[Event Sourcing and CQRS (Command Query Responsibility Segregation)]]
├── [[At-Least-Once, At-Most-Once, and Exactly-Once Event Delivery Semantics]]
└── [[Change Data Capture (CDC) and Transactional Outbox Pattern]]
```

---

## 🗂️ Topics

- [[Event-Driven Architecture (EDA) Core Topology: Broker vs Mediator]] — Comparing decentralized event-broker choreographies with centralized orchestrator mediators.
- [[Event Sourcing and CQRS (Command Query Responsibility Segregation)]] — Persisting state as an append-only log of immutable domain events and building read projections.
- [[At-Least-Once, At-Most-Once, and Exactly-Once Event Delivery Semantics]] — Guaranteeing idempotency, deduplication keys, and distributed consumer commit semantics.
- [[Change Data Capture (CDC) and Transactional Outbox Pattern]] — Reliably tailing database write-ahead logs (WAL) to publish events without dual-write bugs.

---

## 🔗 References
- ⬆️ Parent: [[Programming Paradigms]]

