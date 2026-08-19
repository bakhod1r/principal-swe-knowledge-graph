---
title: Behavioral & Concurrency Patterns
tags:
  - golang
  - design-patterns
  - principal-swe
parent: "[[Design Patterns in Go]]"
---

# Behavioral & Concurrency Patterns

Strategy, observer, iterator, middleware chains, and processing pipelines.

```text
Behavioral & Concurrency Patterns
│
├── [[Strategy Pattern]]
├── [[Observer Pattern]]
├── [[Iterator Pattern]]
├── [[Chain of Responsibility Pattern]]
├── [[Command Pattern]]
├── [[State Pattern]]
├── `PubSub Pattern`
├── `Futures and Promises Pattern`
└── `Fail-Fast Pattern`
```

---

## 🗂️ Topics

- [[Strategy Pattern]] — Swapping business algorithms at runtime via interface injection.
- [[Observer Pattern]] — Event notification system using channels and listener registries.
- [[Iterator Pattern]] — Iterating custom collections using callbacks, channels, and Go 1.23 iterators.
- [[Chain of Responsibility Pattern]] — Passing requests along a dynamic chain of handlers.
- [[Command Pattern]] — Encapsulating requests as objects with undo/redo execution capabilities.
- [[State Pattern]] — Encapsulating object behavior transitions based on internal state machines.
- `PubSub Pattern` — Decoupled publisher-subscriber messaging with topic routing.
- `Futures and Promises Pattern` — Async value computation using channels and read-only signaling.
- `Fail-Fast Pattern` — Early validation and fail-fast assertions in distributed systems.

---

## 🔗 References
- ⬆️ Parent: [[Design Patterns in Go]]

