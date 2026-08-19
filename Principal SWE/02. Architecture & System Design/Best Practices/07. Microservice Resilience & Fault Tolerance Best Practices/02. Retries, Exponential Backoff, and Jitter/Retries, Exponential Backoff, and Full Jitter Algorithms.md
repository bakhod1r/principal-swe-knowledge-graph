---
title: Retries, Exponential Backoff, and Full Jitter Algorithms
tags:
  - review
  - best-practices
  - software-engineering
  - microservice-resilience-and-fault-tolerance-best-practices
  - principal-swe
parent: "[[Microservice Resilience & Fault Tolerance Best Practices]]"
---

# 📦 Retries, Exponential Backoff, and Full Jitter Algorithms

Preventing thundering herd problems on downstream recovery: exponential delay with full jitter (`delay = random_between(0, min(max_delay, base * 2^attempt))`).

```text
Retries, Exponential Backoff, and Full Jitter Algorithms
│
├── [[Retries, Exponential Backoff, and Full Jitter Algorithms Engineering Standards and Principles]]
├── [[Retries, Exponential Backoff, and Full Jitter Algorithms Production Implementation Patterns]]
└── [[Retries, Exponential Backoff, and Full Jitter Algorithms Failure Modes and Anti Pattern Mitigations]]
```

---

## 🗂️ Engineering Standards & Patterns

- [[Retries, Exponential Backoff, and Full Jitter Algorithms Engineering Standards and Principles]]
- [[Retries, Exponential Backoff, and Full Jitter Algorithms Production Implementation Patterns]]
- [[Retries, Exponential Backoff, and Full Jitter Algorithms Failure Modes and Anti Pattern Mitigations]]

---

## 🔗 References
- ⬆️ Parent: [[Microservice Resilience & Fault Tolerance Best Practices]]
- 📚 Module: `Best Practices`

