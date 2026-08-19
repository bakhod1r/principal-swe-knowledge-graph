---
title: Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation
tags:
  - review
  - architecture
  - systems-architecture
  - resilience,-fault-tolerance-and-chaos-engineering
  - principal-swe
parent: "[[Resilience, Fault Tolerance & Chaos Engineering]]"
---

# 📦 Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation

Isolating thread pools and connection pools per downstream dependency so a slow external API does not exhaust resources for the entire application.

```text
Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation
│
├── [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation Architectural Foundations and Invariants]]
├── [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation Production Implementation and Patterns]]
└── [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation Structural Anti Patterns and Gotchas]]
```

---

## 🗂️ Architectural Blueprints & Patterns

- [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation Architectural Foundations and Invariants]]
- [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation Production Implementation and Patterns]]
- [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation Structural Anti Patterns and Gotchas]]

---

## 🔗 References
- ⬆️ Parent: [[Resilience, Fault Tolerance & Chaos Engineering]]
- 📚 Module: `Architecture`

