---
title: Testing & Benchmarking
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Golang]]"
---

# 🧪 Testing & Benchmarking

Unit testing, table-driven tests, mocking, subtests, benchmarks, coverage analysis, fuzzing, golden files, and testcontainers.

```text
Testing & Benchmarking
│
├── [[Unit Testing & Strategies|01. Unit Testing & Strategies]]
│   ├── [[Testing Basics]]
│   ├── [[Table-Driven Tests]]
│   ├── [[Subtests (t.Run)]]
│   ├── [[Test Helpers (t.Helper)]]
│   ├── [[TestMain Function]]
│   └── [[Parallel Tests (t.Parallel)]]
├── [[Mocks, HTTP & Integration|02. Mocks, HTTP & Integration]]
│   ├── [[Interface Mocking & Test Doubles]]
│   ├── [[httptest Package]]
│   ├── [[Golden Files Snapshot Testing]]
│   └── [[Testcontainers in Go]]
└── [[Benchmarks, Coverage & Fuzzing|03. Benchmarks, Coverage & Fuzzing]]
│   ├── [[Benchmarks (testing.B)]]
│   ├── [[benchstat Statistical Analysis]]
│   ├── [[Coverage Analysis (-cover)]]
│   ├── [[Fuzz Testing (testing.F)]]
│   └── [[Race Detector (-race)]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Unit Testing & Strategies|01. Unit Testing & Strategies]]
- [[Testing Basics]] — Writing test functions (TestXxx), testing.T assertions, failure logging, t.FailNow.
- [[Table-Driven Tests]] — Idiomatic Go testing structure: slice of test structs with inputs and expected outputs.
- [[Subtests (t.Run)]] — Hierarchical test execution, granular failure reporting, running specific subtests.
- [[Test Helpers (t.Helper)]] — Marking helper functions to keep stack traces pointing to actual test lines.
- [[TestMain Function]] — Global test suite setup and teardown hooks (m.Run()).
- [[Parallel Tests (t.Parallel)]] — Concurrent test execution, isolating test state, detecting races under test.
### 2. 📂 [[Mocks, HTTP & Integration|02. Mocks, HTTP & Integration]]
- [[Interface Mocking & Test Doubles]] — Hand-written fakes vs generated mocks (mockery, gomock).
- [[httptest Package]] — httptest.NewServer, httptest.ResponseRecorder for isolated HTTP handler testing.
- [[Golden Files Snapshot Testing]] — Validating complex output payloads (JSON, HTML) against snapshot files (-update flag).
- [[Testcontainers in Go]] — Spinning up real PostgreSQL, Redis, and Kafka containers during integration tests.
### 3. 📂 [[Benchmarks, Coverage & Fuzzing|03. Benchmarks, Coverage & Fuzzing]]
- [[Benchmarks (testing.B)]] — b.N iterations, b.ResetTimer(), b.ReportAllocs(), memory allocation tracking.
- [[benchstat Statistical Analysis]] — Comparing benchmark results before and after code changes with statistical confidence.
- [[Coverage Analysis (-cover)]] — go test -cover, generating HTML coverage profiles, enforcing coverage gates in CI.
- [[Fuzz Testing (testing.F)]] — Randomized input mutation engine for discovering edge-case crashes and vulnerabilities.
- [[Race Detector (-race)]] — ThreadSanitizer runtime data race detection during test runs.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
