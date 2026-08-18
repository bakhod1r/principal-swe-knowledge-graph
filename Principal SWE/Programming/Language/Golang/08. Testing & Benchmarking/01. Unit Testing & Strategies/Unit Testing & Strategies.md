---
title: Unit Testing & Strategies
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Unit Testing & Strategies

Testing package fundamentals, table-driven tests, subtests, test helpers, and parallel execution.

```text
Unit Testing & Strategies
│
├── [[Testing Basics]]
├── [[Table-Driven Tests]]
├── [[Subtests (t.Run)]]
├── [[Test Helpers (t.Helper)]]
├── [[TestMain Function]]
├── [[Parallel Tests (t.Parallel)]]
└── [[t.Cleanup Resource Hooks]]
```

---

## 🗂️ Topics

- [[Testing Basics]] — Writing test functions (TestXxx), testing.T assertions, failure logging, t.FailNow.
- [[Table-Driven Tests]] — Idiomatic Go testing structure: slice of test structs with inputs and expected outputs.
- [[Subtests (t.Run)]] — Hierarchical test execution, granular failure reporting, running specific subtests.
- [[Test Helpers (t.Helper)]] — Marking helper functions to keep stack traces pointing to actual test lines.
- [[TestMain Function]] — Global test suite setup and teardown hooks (m.Run()).
- [[Parallel Tests (t.Parallel)]] — Concurrent test execution, isolating test state, detecting races under test.
- [[t.Cleanup Resource Hooks]] — Registering reliable cleanup callbacks executed on test termination.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]
- 🎓 Root: [[Principal SWE]]
