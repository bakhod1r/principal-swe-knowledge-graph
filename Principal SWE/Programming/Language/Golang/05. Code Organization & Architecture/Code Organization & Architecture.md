---
title: Code Organization & Architecture
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Golang]]"
---

# 📁 Code Organization & Architecture

Go modules, package boundaries, standard project layout, internal packages, multi-module workspaces, and dependency management.

```text
Code Organization & Architecture
│
├── [[Modules & Packages|01. Modules & Packages]]
│   ├── [[Modules & go.mod Directives]]
│   ├── [[go.sum Checksum Verification]]
│   ├── [[Packages & Imports]]
│   ├── [[Package Naming Conventions]]
│   ├── [[Module Versioning (SemVer)]]
│   ├── [[Private Modules & Enterprise Setup]]
│   └── [[Vendoring Mechanics (go mod vendor)]]
├── [[Project Layout & Boundaries|02. Project Layout & Boundaries]]
│   ├── [[Standard Go Project Layout]]
│   ├── [[internal- Visibility Enforcement]]
│   ├── [[Workspaces (go.work)]]
│   ├── [[Circular Dependency Prevention]]
│   ├── [[Package Cohesion & Coupling]]
│   └── [[Layered Architecture vs Package by Feature]]
└── [[Architecture & Dependency Injection|03. Architecture & Dependency Injection]]
│   ├── [[Clean Architecture in Go]]
│   ├── [[Hexagonal Architecture (Ports & Adapters)]]
│   ├── [[Domain-Driven Design (DDD) in Go]]
│   ├── [[Dependency Injection Principles]]
│   ├── [[Wire Compile-Time DI]]
│   ├── [[Fx Runtime DI]]
│   └── [[Modular Monoliths vs Microservices in Go]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Modules & Packages|01. Modules & Packages]]
- [[Modules & go.mod Directives]] — require, replace, exclude, and retract directives in go.mod.
- [[go.sum Checksum Verification]] — Cryptographic hashing, go.sum integrity verification, and tamper detection.
- [[Packages & Imports]] — Package namespace rules, single package per directory, dot imports, and blank imports.
- [[Package Naming Conventions]] — Short, concise, lowercase, singular package names without underscores.
- [[Module Versioning (SemVer)]] — Semantic Versioning rules and v2+ major version import path suffixes.
- [[Private Modules & Enterprise Setup]] — GOPRIVATE, authentication, enterprise git repository integration.
- [[Vendoring Mechanics (go mod vendor)]] — Embedding dependencies for reproducible offline enterprise builds.
### 2. 📂 [[Project Layout & Boundaries|02. Project Layout & Boundaries]]
- [[Standard Go Project Layout]] — cmd/, internal/, pkg/, api/ directory layout conventions.
- [[internal- Visibility Enforcement]] — Compiler-enforced access control preventing unauthorized external imports.
- [[Workspaces (go.work)]] — Multi-module local development without modifying go.mod replace directives.
- [[Circular Dependency Prevention]] — Avoiding import cycles through interface abstraction and layered architecture.
- [[Package Cohesion & Coupling]] — Designing focused, loosely-coupled packages with high cohesion.
- [[Layered Architecture vs Package by Feature]] — Comparing horizontal layer packaging with vertical domain feature packaging.
### 3. 📂 [[Architecture & Dependency Injection|03. Architecture & Dependency Injection]]
- [[Clean Architecture in Go]] — Separating domain entities, use cases, controllers, and database adapters.
- [[Hexagonal Architecture (Ports & Adapters)]] — Defining domain ports via interfaces and adapter implementations.
- [[Domain-Driven Design (DDD) in Go]] — Entities, Value Objects, Aggregates, and Repositories in Go.
- [[Dependency Injection Principles]] — Constructor injection vs functional options injection.
- [[Wire Compile-Time DI]] — Google Wire automated compile-time dependency injection code generator.
- [[Fx Runtime DI]] — Uber Fx reflection-based runtime dependency injection framework.
- [[Modular Monoliths vs Microservices in Go]] — Architectural boundaries, shared domain models, and RPC service boundaries.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
