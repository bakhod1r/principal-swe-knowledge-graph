---
title: Fuzzing, Mutation & Property-Based Testing
tags:
  - golang
  - testing
  - principal-swe
parent: "[[Testing & Benchmarking]]"
---

# Fuzzing, Mutation & Property-Based Testing

Native Go fuzz testing (testing.F), property-based testing, mutation testing (go-mutesting), and fault injection.

```text
Fuzzing, Mutation & Property-Based Testing
│
├── [[Native Go Fuzz Testing (testing.F Mutational Engine)]]
├── [[Property-Based Testing (testing-quick & gopter)]]
├── [[Mutation Testing in Go (go-mutesting)]]
└── [[Fault Injection & Chaos Testing in Go]]
```

---

## 🗂️ Topics

- [[Native Go Fuzz Testing (testing.F Mutational Engine)]] — Corpus management, seed corpus (testdata/fuzz), and automated mutational fuzzing in CI.
- [[Property-Based Testing (testing-quick & gopter)]] — Generating randomized inputs to verify mathematical invariants, commutativity, and idempotency laws.
- [[Mutation Testing in Go (go-mutesting)]] — Injecting AST code mutations (swapping operators, altering branches) to evaluate test suite quality.
- [[Fault Injection & Chaos Testing in Go]] — Injecting artificial network latency, packet loss, and database errors into test pipelines.

---

## 🔗 References
- ⬆️ Parent: [[Testing & Benchmarking]]
- 🎓 Root: [[Principal SWE]]
