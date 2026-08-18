---
title: Dependency Injection & Decoupling Patterns
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Code Organization & Architecture]]"
---

# Dependency Injection & Decoupling Patterns

Manual constructor DI, compile-time DI with Google Wire, Functional Options pattern, interface decoupling, and Uber Fx.

```text
Dependency Injection & Decoupling Patterns
│
├── [[Manual Constructor Dependency Injection (Idiomatic Go)]]
├── [[Compile-Time Dependency Injection with Google Wire]]
├── [[Functional Options Pattern for Flexible Configuration]]
├── [[Interface-Driven Decoupling & Testability]]
└── [[Uber Fx Framework in Enterprise Go Services]]
```

---

## 🗂️ Topics

- [[Manual Constructor Dependency Injection (Idiomatic Go)]] — Creating explicit NewService(repo, client, logger) constructors without magic frameworks.
- [[Compile-Time Dependency Injection with Google Wire]] — Generating static dependency injection graphs at build time without reflection overhead.
- [[Functional Options Pattern for Flexible Configuration]] — Clean, extensible constructor configurations with functional options (WithTimeout, WithRetries).
- [[Interface-Driven Decoupling & Testability]] — Declaring consumer-side interfaces to enable seamless mock and fake substitutions in tests.
- [[Uber Fx Framework in Enterprise Go Services]] — Lifecycle management (OnStart, OnStop), dependency injection, and modular container bootstrapping.

---

## 🔗 References
- ⬆️ Parent: [[Code Organization & Architecture]]

