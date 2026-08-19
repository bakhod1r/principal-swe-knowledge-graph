---
title: Code Organization
tags:
  - golang
  - architecture
  - principal-swe
parent: "[[Golang]]"
---

# 🏛️ Code Organization & Architecture

Enterprise architecture in Go: package-oriented design, standard project layouts, Clean Architecture, multi-module workspaces (go.work), compile-time DI (Google Wire), REST/gRPC API evolution, and architectural anti-patterns.

```text
Code Organization & Architecture
│
├── [[Package Design Principles & Encapsulation|01. Package Design Principles & Encapsulation]]
│   ├── `Package-Oriented Design (POD) Architecture`
│   ├── `internal- Visibility Enforcement Mechanics`
│   ├── `Circular Dependency Elimination Strategies`
│   ├── `Exported vs Unexported Identifiers & API Surface Minimalism`
│   └── `Package Naming Conventions & Anti-Patterns`
├── `02. Project Layouts & Repository Architectures`
│   ├── `Standard Go Project Layout (cmd, internal, pkg, api, web)`
│   ├── `Flat vs Layered vs Modular Repository Architectures`
│   ├── `Clean Architecture & Hexagonal Ports-and-Adapters`
│   ├── `Enterprise Monorepos vs Multi-Repo Microservices`
│   └── `Domain-Driven Design (DDD) Bounded Contexts in Go`
├── [[Modules, Workspaces & Versioning (go.mod, go.work)|03. Modules, Workspaces & Versioning (go.mod, go.work)]]
│   ├── `Multi-Module Workspaces with go.work`
│   ├── `Semantic Import Versioning (v2+ Module Paths)`
│   ├── `replace, exclude, and retract Directives in go.mod`
│   ├── `Private Module Configuration (GOPRIVATE, GONOPROXY)`
│   └── `Vendoring Mechanics (go mod vendor & -mod=vendor)`
├── `04. Dependency Injection & Decoupling Patterns`
│   ├── `Manual Constructor Dependency Injection (Idiomatic Go)`
│   ├── `Compile-Time Dependency Injection with Google Wire`
│   ├── `Functional Options Pattern for Flexible Configuration`
│   ├── `Interface-Driven Decoupling & Testability`
│   └── `Uber Fx Framework in Enterprise Go Services`
├── `05. Enterprise API Design & Evolution`
│   ├── `RESTful API Architecture with net-http and chi`
│   ├── `gRPC and Protocol Buffers Service Architecture`
│   ├── `API Versioning Strategies (URL, Header, Subdomain)`
│   ├── `OpenAPI & Swagger Documentation Generation`
│   └── `Backward Compatibility & Breaking Change Prevention`
├── `06. Code Architecture Anti-Patterns & Code Smells`
└── `07. Idiomatic Clean Code & Refactoring Standards`
│   ├── `The Global State & Singleton Anti-Pattern`
│   ├── `The God Package & Package Clutter Anti-Pattern`
│   ├── `The Package Stuttering Anti-Pattern`
│   ├── `Cyclic Dependency Workaround Hacks`
│   └── `Staff-Level Code Architecture Checklist`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Package Design Principles & Encapsulation|01. Package Design Principles & Encapsulation]]
- `Package-Oriented Design (POD) Architecture` — Guidelines for designing packages around domain purpose rather than technical layers (avoiding utils, common).
- `internal- Visibility Enforcement Mechanics` — How the Go compiler strictly restricts access to packages located under /internal/ trees.
- `Circular Dependency Elimination Strategies` — Resolving import cycle not allowed compiler errors via interface extraction and mediator packages.
- `Exported vs Unexported Identifiers & API Surface Minimalism` — Designing minimal, intention-revealing public package interfaces and hiding internals.
- `Package Naming Conventions & Anti-Patterns` — Eliminating stuttering (http.HTTPServer), multi-word packages, and generic naming smells.
### 2. 📂 `02. Project Layouts & Repository Architectures`
- `Standard Go Project Layout (cmd, internal, pkg, api, web)` — Enterprise standard directory taxonomy and separation of operational entry points.
- `Flat vs Layered vs Modular Repository Architectures` — Staff-level evaluation: when to use flat single-package designs vs multi-module layouts.
- `Clean Architecture & Hexagonal Ports-and-Adapters` — Domain entities at the core, use case interactor layers, and pluggable infrastructure adapters.
- `Enterprise Monorepos vs Multi-Repo Microservices` — Tooling, workspace orchestration (go.work), shared libraries, and release boundaries.
- `Domain-Driven Design (DDD) Bounded Contexts in Go` — Structuring aggregates, value objects, domain events, and repositories within isolated packages.
### 3. 📂 [[Modules, Workspaces & Versioning (go.mod, go.work)|03. Modules, Workspaces & Versioning (go.mod, go.work)]]
- `Multi-Module Workspaces with go.work` — Local development across multiple sibling modules without requiring replace directives.
- `Semantic Import Versioning (v2+ Module Paths)` — Major version upgrades (/v2), branch strategies, and breaking API boundary management.
- `replace, exclude, and retract Directives in go.mod` — Local development overrides, CVE-vulnerable release retraction, and dependency exclusions.
- `Private Module Configuration (GOPRIVATE, GONOPROXY)` — Authenticating with enterprise GitLab/GitHub private registries and SSH deploy keys.
- `Vendoring Mechanics (go mod vendor & -mod=vendor)` — Hermetic builds, offline CI pipelines, and vendored dependency auditing.
### 4. 📂 `04. Dependency Injection & Decoupling Patterns`
- `Manual Constructor Dependency Injection (Idiomatic Go)` — Creating explicit NewService(repo, client, logger) constructors without magic frameworks.
- `Compile-Time Dependency Injection with Google Wire` — Generating static dependency injection graphs at build time without reflection overhead.
- `Functional Options Pattern for Flexible Configuration` — Clean, extensible constructor configurations with functional options (WithTimeout, WithRetries).
- `Interface-Driven Decoupling & Testability` — Declaring consumer-side interfaces to enable seamless mock and fake substitutions in tests.
- `Uber Fx Framework in Enterprise Go Services` — Lifecycle management (OnStart, OnStop), dependency injection, and modular container bootstrapping.
### 5. 📂 `05. Enterprise API Design & Evolution`
- `RESTful API Architecture with net-http and chi` — Building fast, idiomatic REST services with sub-routers and context middlewares.
- `gRPC and Protocol Buffers Service Architecture` — Defining .proto service contracts, generating Go stubs, interceptors, and streaming RPCs.
- `API Versioning Strategies (URL, Header, Subdomain)` — Backward-compatible schema evolution, deprecation timelines, and field migrations.
- `OpenAPI & Swagger Documentation Generation` — Auto-generating OpenAPI specs from Go doc comments and declarative annotations.
- `Backward Compatibility & Breaking Change Prevention` — Protobuf field number rules, additive JSON extensions, and contract testing.
### 6. 📂 `06. Code Architecture Anti-Patterns & Code Smells`
- `The Global State & Singleton Anti-Pattern` — Concurrency race hazards, testing pollution, and hidden coupling caused by global variables.
- `The God Package & Package Clutter Anti-Pattern` — Giant monolithic packages containing everything and dumping code into helpers/.
- `The Package Stuttering Anti-Pattern` — Redundant identifiers (user.UserService, client.ClientConfig) harming idiomatic Go readability.
- `Cyclic Dependency Workaround Hacks` — Dangerous anti-patterns: using init() hooks or type casting to bypass circular imports.
- `Staff-Level Code Architecture Checklist` — Pre-production architectural review checklist for high-scale enterprise Go services.

### 7. 📂 `07. Idiomatic Clean Code & Refactoring Standards`
- `Line of Sight & Early Return Idiom` — Aligning the happy path to the left margin, reducing indentation, and eliminating else branches via early returns.
- `Function Signature Design & Minimal Parameter Lists` — Designing expressive signatures, limiting argument counts, and struct parameter objects vs functional options.
- `Naming Idioms (Short-Lived vs Long-Lived Identifiers)` — Single-letter variables in small scopes vs descriptive package-level names, avoiding redundant prefixes.
- `Comment & Documentation Best Practices (Godoc Conventions)` — Writing actionable doc comments explaining WHY rather than WHAT, package comments, and deprecation notices.
- `Refactoring Complex Conditionals & Deep Nesting` — Transforming deeply nested if-else ladders into guard clauses, table-driven lookups, and sub-routines.
- `Avoiding Boolean Parameter Traps in Function Design` — Eliminating boolean flag parameters (e.g. Process(true)) in favor of expressive enum types or functional options.
- `Zero-Value Useful Structs Design Principle` — Designing structs that require no constructors and work out-of-the-box in their zero state (e.g. sync.Mutex, bytes.Buffer).
- `Explicit Over Implicit Engineering Principle in Go` — Why Go deliberately rejects implicit type conversions, operator overloading, and magic annotations.
- `SOLID Principles Applied in Idiomatic Go` — Mapping Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, and Dependency Inversion to Go.
- `Principal Code Review Checklist & Style Guide Synthesis` — Synthesis of Google Go Style Guide, Effective Go, and Uber Go Style Guide for enterprise pull request reviews.

- `Accept Interfaces, Return Structs Principle` — Consumer-driven interface definition: defining narrow single-method interfaces on consumer side rather than producer side.
- `Interface Pollution & Premature Abstraction Defense` — Recognizing when concrete structs are simpler and faster, avoiding 1-to-1 interface-struct boilerplate.
- `Copy Mutex by Value Hazard & Struct Semantics` — Preventing silent synchronization bugs by passing structs containing sync.Mutex by pointer and copylocks analysis.
- `Package Stuttering & Ergonomic Naming Rules` — Eliminating package prefix redundancy (e.g. http.Server not http.HTTPServer, logger.Logger not logger.LogLogger).
- `Errors Are Values Idiom & Repetition Reduction` — Treating errors as first-class domain values (e.g. errWriter pattern) to eliminate repetitive if err != nil boilerplate.
- `Variable Lifetime Minimization & Smallest Viable Scope` — Declaring variables immediately before use with short-statement if blocks to prevent accidental reuse and scope pollution.
- `Side-Effect Free Pure Functions & Immutability` — Writing deterministic transform functions without mutating input pointer state or relying on global variables.
- `Naked Returns Anti-Pattern in Complex Functions` — Why naked returns without explicit variable names harm readability and maintainability in non-trivial functions.
- `Data-Driven Clean Test Refactoring & Subtest Idioms` — Keeping test code as clean and readable as production code using anonymous table structs and expressive sub-test naming.
- `API Deprecation Lifecycles & Godoc Warnings` — Using // Deprecated: notices in Godoc comments, compiler warnings, and phased deprecation lifecycles.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Topics

- [[Idiomatic Clean Code]]
- [[Modules, Workspaces & Versioning (go.mod, go.work)]]
- [[Package Design Principles & Encapsulation]]
- [[Refactoring Standards]]
