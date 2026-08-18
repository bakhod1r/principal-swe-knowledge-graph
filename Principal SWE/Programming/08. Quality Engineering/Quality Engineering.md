---
title: Quality Engineering
tags:
  - programming
  - quality
  - principal-swe
parent: "[[Programming]]"
---

# 💻 Quality Engineering

Testing pyramid, TDD/BDD, contract testing, property-based testing, chaos engineering, and CI/CD verification.

```text
Quality Engineering
│
├── [[Testing Hierarchy & Strategies|01. Testing Hierarchy & Strategies]]
│   ├── [[Test Pyramid & Modern Testing Diamond (Unit, Integration, E2E)]]
│   ├── [[Test-Driven Development (TDD) & Behavior-Driven Development (BDD)]]
│   ├── [[Contract Testing & Consumer-Driven Contracts (Pact)]]
│   └── [[Mocking, Stubs, Spies & Fake Implementations]]
├── [[Advanced Testing Methodologies|02. Advanced Testing Methodologies]]
│   ├── [[Property-Based Testing & Generative Testing (QuickCheck)]]
│   ├── [[Mutation Testing & Test Quality Verification]]
│   ├── [[Chaos Engineering & Fault Injection (Chaos Mesh, Toxiproxy)]]
│   └── [[Load, Stress & Soak Testing (k6, Locust, Gatling)]]
└── [[Continuous Verification & QA in CI-CD|03. Continuous Verification & QA in CI-CD]]
│   ├── [[Automated Flaky Test Detection & Quarantine Pipelines]]
│   ├── [[Shift-Left Quality & Pre-Commit Linting-SAST]]
│   └── [[Code Coverage vs Mutation Score vs Assertion Density]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Testing Hierarchy & Strategies|01. Testing Hierarchy & Strategies]]
- [[Test Pyramid & Modern Testing Diamond (Unit, Integration, E2E)]] — Balancing test execution speed, cost, and confidence across testing tiers to prevent inverted pyramid antipatterns.
- [[Test-Driven Development (TDD) & Behavior-Driven Development (BDD)]] — Red-Green-Refactor cycle, test-first architectural exploration, and executable Gherkin specifications.
- [[Contract Testing & Consumer-Driven Contracts (Pact)]] — Validating microservice API integrations independently without spinning up brittle end-to-end environments.
- [[Mocking, Stubs, Spies & Fake Implementations]] — Understanding test doubles: Dummy, Stub, Spy, Mock, and in-memory Fake implementations.
### 2. 📂 [[Advanced Testing Methodologies|02. Advanced Testing Methodologies]]
- [[Property-Based Testing & Generative Testing (QuickCheck)]] — Generating hundreds of random test cases to verify universal invariants and shrinking failing inputs.
- [[Mutation Testing & Test Quality Verification]] — Injecting artificial defects into source code to verify whether the test suite actively detects and catches regressions.
- [[Chaos Engineering & Fault Injection (Chaos Mesh, Toxiproxy)]] — Simulating network partitions, packet loss, disk exhaustion, and kill signals to verify distributed resilience.
- [[Load, Stress & Soak Testing (k6, Locust, Gatling)]] — Simulating extreme traffic spikes, measuring saturation curves, and detecting slow memory/goroutine leaks.
### 3. 📂 [[Continuous Verification & QA in CI-CD|03. Continuous Verification & QA in CI-CD]]
- [[Automated Flaky Test Detection & Quarantine Pipelines]] — Identifying non-deterministic tests, tracking failure probabilities, and quarantining to protect CI velocity.
- [[Shift-Left Quality & Pre-Commit Linting-SAST]] — Catching syntax, lint, security, and architectural violations before code reaches code review.
- [[Code Coverage vs Mutation Score vs Assertion Density]] — Why raw line coverage is a vanity metric and how mutation score and assertion density measure true test strength.

---

## 🔗 Navigation
- ⬆️ Parent: [[Programming]]
- 🎓 Root: [[Principal SWE]]
