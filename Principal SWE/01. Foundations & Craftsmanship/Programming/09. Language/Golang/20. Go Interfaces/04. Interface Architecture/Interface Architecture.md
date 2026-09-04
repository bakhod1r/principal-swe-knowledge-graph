---
title: Interface Architecture & Design Patterns
tags:
  - golang
  - methods-and-interfaces
  - principal-swe
parent: "[[Go Interfaces]]"
---

# Interface Architecture & Design Patterns

Interface Segregation Principle, Adapter, Decorator, Strategy, interface pollution anti-patterns, and mocking guidelines.

```text
Interface Architecture & Design Patterns
│
├── [[Interface Segregation Principle (ISP) in Go]]
├── `Adapter Pattern with Interfaces`
├── `Decorator & Middleware Pattern with Interfaces`
├── `Strategy Pattern via Functional Interfaces`
├── [[The Interface Pollution Anti-Pattern]]
├── [[Mocking What You Do Not Own Anti-Pattern]]
└── [[Staff-Level Interface Design Guidelines]]
```

---

## 🗂️ Topics

- [[Interface Segregation Principle (ISP) in Go]] — Keeping interfaces small (1-2 methods) and tailored to specific consumer requirements.
- [[Adapter Pattern with Interfaces]] — Bridging incompatible third-party libraries behind clean domain interfaces.
- [[Decorator & Middleware Pattern with Interfaces]] — Composing cross-cutting concerns (logging, metrics, retries) via interface wrappers.
- [[Strategy Pattern via Functional Interfaces]] — Injecting interchangeable algorithms at runtime using single-method interfaces or function types.
- [[The Interface Pollution Anti-Pattern]] — Prematurely defining interfaces for every struct before having multiple implementations.
- [[Mocking What You Do Not Own Anti-Pattern]] — Why you should only define interfaces and mocks for code you own and consume.
- [[Staff-Level Interface Design Guidelines]] — Best practices for designing scalable, maintainable Go APIs with clean interface boundaries.

---

## 🔗 References
- ⬆️ Parent: `Methods & Interfaces`

