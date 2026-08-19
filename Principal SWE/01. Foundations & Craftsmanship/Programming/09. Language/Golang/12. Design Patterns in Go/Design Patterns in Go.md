---
title: Design Patterns in Go
tags:
  - golang
  - design-patterns
  - principal-swe
parent: "[[Golang]]"
---

# 🏛️ Design Patterns in Go

Idiomatic Go implementations of creational, structural, behavioral, and distributed cloud-native design patterns.

```text
Design Patterns in Go
│
├── [[Creational Patterns|01. Creational Patterns]]
│   ├── `Functional Options Pattern`
│   ├── `Builder Pattern`
│   ├── `Factory Pattern`
│   ├── `Singleton Pattern (sync.Once)`
│   ├── `Object Pool Pattern (sync.Pool)`
│   └── `Registry Pattern`
├── [[Structural Patterns|02. Structural Patterns]]
│   ├── `Adapter Pattern`
│   ├── `Decorator Pattern`
│   ├── `Facade Pattern`
│   ├── `Proxy Pattern`
│   └── `Composite Pattern`
├── `03. Behavioral & Concurrency Patterns`
│   ├── `Strategy Pattern`
│   ├── `Observer Pattern`
│   ├── `Iterator Pattern`
│   ├── `Chain of Responsibility Pattern`
│   ├── `Command Pattern`
│   ├── `State Pattern`
│   ├── `PubSub Pattern`
│   ├── `Futures and Promises Pattern`
│   └── `Fail-Fast Pattern`
└── `04. Microservice & Cloud-Native Patterns`
│   ├── `Outbox Pattern`
│   ├── `Saga Pattern (Orchestration vs Choreography)`
│   ├── `Dead Letter Queue (DLQ) Pattern`
│   ├── `Idempotent Consumer Pattern`
│   ├── `Sidecar Communication Pattern`
│   └── `Graceful Degradation Pattern`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Creational Patterns|01. Creational Patterns]]
- `Functional Options Pattern` — Clean, extensible struct initialization with default values and option functions.
- `Builder Pattern` — Step-by-step construction of complex objects with validation.
- `Factory Pattern` — Encapsulating object creation behind interface contracts.
- `Singleton Pattern (sync.Once)` — Thread-safe lazy initialization using sync.Once.
- `Object Pool Pattern (sync.Pool)` — Reusing heavy allocations with sync.Pool and custom ring buffers.
- `Registry Pattern` — Thread-safe dynamic registration and lookup of plugins or handlers.
### 2. 📂 [[Structural Patterns|02. Structural Patterns]]
- `Adapter Pattern` — Bridging incompatible interfaces without modifying existing structs.
- `Decorator Pattern` — Wrapping structs to augment functionality (e.g. logging, metrics wrappers).
- `Facade Pattern` — Providing a simplified high-level interface over a complex subsystem.
- `Proxy Pattern` — Controlling access to an underlying object (e.g. caching proxy, auth proxy).
- `Composite Pattern` — Treating individual objects and compositions of objects uniformly.
### 3. 📂 `03. Behavioral & Concurrency Patterns`
- `Strategy Pattern` — Swapping business algorithms at runtime via interface injection.
- `Observer Pattern` — Event notification system using channels and listener registries.
- `Iterator Pattern` — Iterating custom collections using callbacks, channels, and Go 1.23 iterators.
- `Chain of Responsibility Pattern` — Passing requests along a dynamic chain of handlers.
- `Command Pattern` — Encapsulating requests as objects with undo/redo execution capabilities.
- `State Pattern` — Encapsulating object behavior transitions based on internal state machines.
- `PubSub Pattern` — Decoupled publisher-subscriber messaging with topic routing.
- `Futures and Promises Pattern` — Async value computation using channels and read-only signaling.
- `Fail-Fast Pattern` — Early validation and fail-fast assertions in distributed systems.
### 4. 📂 `04. Microservice & Cloud-Native Patterns`
- `Outbox Pattern` — Guaranteed at-least-once message publishing using relational database transaction logs.
- `Saga Pattern (Orchestration vs Choreography)` — Managing distributed multi-service transactions with compensating rollback actions.
- `Dead Letter Queue (DLQ) Pattern` — Isolating unprocessable poison messages for inspection and replay.
- `Idempotent Consumer Pattern` — Deduplicating incoming message delivery using persistent transaction IDs.
- `Sidecar Communication Pattern` — Interacting with local Envoy/Dapr sidecars over gRPC/UDS Unix Domain Sockets.
- `Graceful Degradation Pattern` — Serving stale cached data or partial responses during downstream dependency outages.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

