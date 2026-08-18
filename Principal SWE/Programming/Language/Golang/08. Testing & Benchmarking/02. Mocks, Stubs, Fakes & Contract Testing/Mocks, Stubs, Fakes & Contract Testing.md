---
title: Mocks, Stubs, Fakes & Contract Testing
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Mocks, Stubs, Fakes & Contract Testing

Test double taxonomy, hand-written in-memory fakes, mock generators (mockery, gomock), and contract testing.

```text
Mocks, Stubs, Fakes & Contract Testing
│
├── [[Test Doubles Taxonomy (Dummies, Stubs, Fakes, Spies, Mocks)]]
├── [[Hand-Written In-Memory Fakes (Zero-Dependency Testing)]]
├── [[Mock Generators (mockery, gomock, moq)]]
├── [[Mocking What You Do Not Own Anti-Pattern in Testing]]
└── [[Consumer-Driven Contract Testing (Pact in Go)]]
```

---

## 🗂️ Topics

- [[Test Doubles Taxonomy (Dummies, Stubs, Fakes, Spies, Mocks)]] — Architectural definitions of test doubles and when to choose fakes vs generated mocks.
- [[Hand-Written In-Memory Fakes (Zero-Dependency Testing)]] — Writing high-performance in-memory repository fakes for deterministic unit testing.
- [[Mock Generators (mockery, gomock, moq)]] — Automated interface mock generation, call expectations, argument matchers, and assertions.
- [[Mocking What You Do Not Own Anti-Pattern in Testing]] — Why third-party libraries should never be mocked directly and must be wrapped in domain interfaces.
- [[Consumer-Driven Contract Testing (Pact in Go)]] — Validating API schemas, contracts, and backwards compatibility across distributed microservices.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]
- 🎓 Root: [[Principal SWE]]
