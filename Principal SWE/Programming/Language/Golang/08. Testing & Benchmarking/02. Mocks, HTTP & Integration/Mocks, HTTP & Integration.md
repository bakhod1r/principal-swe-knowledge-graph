---
title: Mocks, HTTP & Integration
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Mocks, HTTP & Integration

Interface mocking, HTTP test utilities (httptest), golden files, and container-based testing.

```text
Mocks, HTTP & Integration
│
├── [[Interface Mocking & Test Doubles]]
├── [[httptest Package]]
├── [[Golden Files Snapshot Testing]]
├── [[Testcontainers in Go]]
└── [[Property-Based Testing]]
```

---

## 🗂️ Topics

- [[Interface Mocking & Test Doubles]] — Hand-written fakes vs generated mocks (mockery, gomock).
- [[httptest Package]] — httptest.NewServer, httptest.ResponseRecorder for isolated HTTP handler testing.
- [[Golden Files Snapshot Testing]] — Validating complex output payloads (JSON, HTML) against snapshot files (-update flag).
- [[Testcontainers in Go]] — Spinning up real PostgreSQL, Redis, and Kafka containers during integration tests.
- [[Property-Based Testing]] — Generative property testing using testing/quick and gopter.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]
- 🎓 Root: [[Principal SWE]]
