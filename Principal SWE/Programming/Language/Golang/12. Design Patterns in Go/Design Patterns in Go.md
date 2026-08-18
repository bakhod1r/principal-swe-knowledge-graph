---
title: Design Patterns in Go
tags:
  - golang
  - design-patterns
  - principal-swe
parent: "[[Golang]]"
---

# 🏛️ Design Patterns in Go

Idiomatic Go implementations of creational, structural, and behavioral design patterns: functional options, builder, decorator, observer, and pipelines.

```text
Design Patterns in Go
│
├── [[Creational Patterns|01. Creational Patterns]]
│   ├── [[Functional Options Pattern]]
│   ├── [[Builder Pattern]]
│   ├── [[Factory Pattern]]
│   ├── [[Singleton Pattern (sync.Once)]]
│   └── [[Object Pool Pattern (sync.Pool)]]
├── [[Structural Patterns|02. Structural Patterns]]
│   ├── [[Adapter Pattern]]
│   ├── [[Decorator Pattern]]
│   ├── [[Facade Pattern]]
│   ├── [[Proxy Pattern]]
│   └── [[Composite Pattern]]
└── [[Behavioral & Concurrency Patterns|03. Behavioral & Concurrency Patterns]]
│   ├── [[Strategy Pattern]]
│   ├── [[Observer Pattern]]
│   ├── [[Iterator Pattern]]
│   ├── [[Chain of Responsibility Pattern]]
│   ├── [[State Pattern]]
│   ├── [[Middleware Handler Pattern]]
│   └── [[Pipeline Pattern]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Creational Patterns|01. Creational Patterns]]
- [[Functional Options Pattern]] — Clean, extensible struct initialization with default values and option functions.
- [[Builder Pattern]] — Step-by-step construction of complex objects with validation.
- [[Factory Pattern]] — Encapsulating object creation behind interface contracts.
- [[Singleton Pattern (sync.Once)]] — Thread-safe lazy initialization using sync.Once.
- [[Object Pool Pattern (sync.Pool)]] — Reusing heavy allocations with sync.Pool and custom ring buffers.
### 2. 📂 [[Structural Patterns|02. Structural Patterns]]
- [[Adapter Pattern]] — Bridging incompatible interfaces without modifying existing structs.
- [[Decorator Pattern]] — Wrapping structs to augment functionality (e.g. logging, metrics wrappers).
- [[Facade Pattern]] — Providing a simplified high-level interface over a complex subsystem.
- [[Proxy Pattern]] — Controlling access to an underlying object (e.g. caching proxy, auth proxy).
- [[Composite Pattern]] — Treating individual objects and compositions of objects uniformly.
### 3. 📂 [[Behavioral & Concurrency Patterns|03. Behavioral & Concurrency Patterns]]
- [[Strategy Pattern]] — Swapping business algorithms at runtime via interface injection.
- [[Observer Pattern]] — Event notification system using channels and listener registries.
- [[Iterator Pattern]] — Iterating custom collections using callbacks, channels, and Go 1.23 iterators.
- [[Chain of Responsibility Pattern]] — Passing requests along a dynamic chain of handlers.
- [[State Pattern]] — Encapsulating object behavior transitions based on internal state machines.
- [[Middleware Handler Pattern]] — HTTP handler chaining with onion-layer request/response processing.
- [[Pipeline Pattern]] — Connecting multi-stage concurrent processing steps through channels.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
