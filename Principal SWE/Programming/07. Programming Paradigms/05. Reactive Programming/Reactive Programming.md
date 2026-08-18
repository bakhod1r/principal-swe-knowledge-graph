---
title: Reactive Programming
tags:
  - programming
  - paradigms
  - principal-swe
parent: "[[Programming Paradigms]]"
---

# Reactive Programming

Reactive Streams specification, Observables, marble diagrams, backpressure strategies (Request-N), and hot/cold multicasting.

```text
Reactive Programming
│
├── [[Reactive Streams Specification and Publisher-Subscriber Contracts]]
├── [[Observable Streams, Operators, and Marble Diagrams (RxJS, Project Reactor)]]
├── [[Backpressure Strategies (Buffering, Dropping, Latest, Request-N)]]
└── [[Hot vs Cold Observables and Multicasting Semantics]]
```

---

## 🗂️ Topics

- [[Reactive Streams Specification and Publisher-Subscriber Contracts]] — The asynchronous stream standard defining Subscription, Publisher, Subscriber, and Processor protocols.
- [[Observable Streams, Operators, and Marble Diagrams (RxJS, Project Reactor)]] — Transforming event pipelines with pure operators (map, filter, flatMap, switchMap, debounce).
- [[Backpressure Strategies (Buffering, Dropping, Latest, Request-N)]] — Preventing fast producers from overwhelming slow consumers using demand-driven signaling.
- [[Hot vs Cold Observables and Multicasting Semantics]] — Understanding producer execution lifecycles: on-demand generation vs shared multicasting broadcasts.

---

## 🔗 References
- ⬆️ Parent: [[Programming Paradigms]]
- 🎓 Root: [[Principal SWE]]
