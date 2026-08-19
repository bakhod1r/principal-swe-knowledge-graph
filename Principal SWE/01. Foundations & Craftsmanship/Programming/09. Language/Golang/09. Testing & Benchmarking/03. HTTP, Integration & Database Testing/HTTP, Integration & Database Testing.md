---
title: HTTP, Integration & Database Testing
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# HTTP, Integration & Database Testing

httptest package, golden files snapshot testing, Testcontainers-go, database fixtures, and in-memory gRPC testing.

```text
HTTP, Integration & Database Testing
│
├── [[httptest Package (ResponseRecorder & NewServer)]]
├── [[Golden Files Snapshot Testing (-update flag)]]
├── [[Testcontainers-go (Real Postgres, Redis, Kafka)]]
├── [[Database Fixtures & Migration Lifecycle in Tests]]
└── [[Testing gRPC Services (bufconn In-Memory Listeners)]]
```

---

## 🗂️ Topics

- [[httptest Package (ResponseRecorder & NewServer)]] — Testing HTTP handlers, routing, and middlewares without opening real OS network ports.
- [[Golden Files Snapshot Testing (-update flag)]] — Validating complex JSON, HTML, and binary payloads against tracked golden snapshot files.
- [[Testcontainers-go (Real Postgres, Redis, Kafka)]] — Spinning up real Dockerized database and message broker containers in integration tests.
- [[Database Fixtures & Migration Lifecycle in Tests]] — Setting up clean transactional test databases, seeding fixtures, and rolling back migrations.
- [[Testing gRPC Services (bufconn In-Memory Listeners)]] — In-memory gRPC client-server integration testing without TCP network socket overhead.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]

