---
title: Code Organization & Architecture
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
│   ├── [[Package-Oriented Design (POD) Architecture]]
│   ├── [[internal- Visibility Enforcement Mechanics]]
│   ├── [[Circular Dependency Elimination Strategies]]
│   ├── [[Exported vs Unexported Identifiers & API Surface Minimalism]]
│   └── [[Package Naming Conventions & Anti-Patterns]]
├── [[Project Layouts & Repository Architectures|02. Project Layouts & Repository Architectures]]
│   ├── [[Standard Go Project Layout (cmd, internal, pkg, api, web)]]
│   ├── [[Flat vs Layered vs Modular Repository Architectures]]
│   ├── [[Clean Architecture & Hexagonal Ports-and-Adapters]]
│   ├── [[Enterprise Monorepos vs Multi-Repo Microservices]]
│   └── [[Domain-Driven Design (DDD) Bounded Contexts in Go]]
├── [[Modules, Workspaces & Versioning (go.mod, go.work)|03. Modules, Workspaces & Versioning (go.mod, go.work)]]
│   ├── [[Multi-Module Workspaces with go.work]]
│   ├── [[Semantic Import Versioning (v2+ Module Paths)]]
│   ├── [[replace, exclude, and retract Directives in go.mod]]
│   ├── [[Private Module Configuration (GOPRIVATE, GONOPROXY)]]
│   └── [[Vendoring Mechanics (go mod vendor & -mod=vendor)]]
├── [[Dependency Injection & Decoupling Patterns|04. Dependency Injection & Decoupling Patterns]]
│   ├── [[Manual Constructor Dependency Injection (Idiomatic Go)]]
│   ├── [[Compile-Time Dependency Injection with Google Wire]]
│   ├── [[Functional Options Pattern for Flexible Configuration]]
│   ├── [[Interface-Driven Decoupling & Testability]]
│   └── [[Uber Fx Framework in Enterprise Go Services]]
├── [[Enterprise API Design & Evolution|05. Enterprise API Design & Evolution]]
│   ├── [[RESTful API Architecture with net-http and chi]]
│   ├── [[gRPC and Protocol Buffers Service Architecture]]
│   ├── [[API Versioning Strategies (URL, Header, Subdomain)]]
│   ├── [[OpenAPI & Swagger Documentation Generation]]
│   └── [[Backward Compatibility & Breaking Change Prevention]]
└── [[Code Architecture Anti-Patterns & Code Smells|06. Code Architecture Anti-Patterns & Code Smells]]
│   ├── [[The Global State & Singleton Anti-Pattern]]
│   ├── [[The God Package & Package Clutter Anti-Pattern]]
│   ├── [[The Package Stuttering Anti-Pattern]]
│   ├── [[Cyclic Dependency Workaround Hacks]]
│   └── [[Staff-Level Code Architecture Checklist]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Package Design Principles & Encapsulation|01. Package Design Principles & Encapsulation]]
- [[Package-Oriented Design (POD) Architecture]] — Guidelines for designing packages around domain purpose rather than technical layers (avoiding utils, common).
- [[internal- Visibility Enforcement Mechanics]] — How the Go compiler strictly restricts access to packages located under /internal/ trees.
- [[Circular Dependency Elimination Strategies]] — Resolving import cycle not allowed compiler errors via interface extraction and mediator packages.
- [[Exported vs Unexported Identifiers & API Surface Minimalism]] — Designing minimal, intention-revealing public package interfaces and hiding internals.
- [[Package Naming Conventions & Anti-Patterns]] — Eliminating stuttering (http.HTTPServer), multi-word packages, and generic naming smells.
### 2. 📂 [[Project Layouts & Repository Architectures|02. Project Layouts & Repository Architectures]]
- [[Standard Go Project Layout (cmd, internal, pkg, api, web)]] — Enterprise standard directory taxonomy and separation of operational entry points.
- [[Flat vs Layered vs Modular Repository Architectures]] — Staff-level evaluation: when to use flat single-package designs vs multi-module layouts.
- [[Clean Architecture & Hexagonal Ports-and-Adapters]] — Domain entities at the core, use case interactor layers, and pluggable infrastructure adapters.
- [[Enterprise Monorepos vs Multi-Repo Microservices]] — Tooling, workspace orchestration (go.work), shared libraries, and release boundaries.
- [[Domain-Driven Design (DDD) Bounded Contexts in Go]] — Structuring aggregates, value objects, domain events, and repositories within isolated packages.
### 3. 📂 [[Modules, Workspaces & Versioning (go.mod, go.work)|03. Modules, Workspaces & Versioning (go.mod, go.work)]]
- [[Multi-Module Workspaces with go.work]] — Local development across multiple sibling modules without requiring replace directives.
- [[Semantic Import Versioning (v2+ Module Paths)]] — Major version upgrades (/v2), branch strategies, and breaking API boundary management.
- [[replace, exclude, and retract Directives in go.mod]] — Local development overrides, CVE-vulnerable release retraction, and dependency exclusions.
- [[Private Module Configuration (GOPRIVATE, GONOPROXY)]] — Authenticating with enterprise GitLab/GitHub private registries and SSH deploy keys.
- [[Vendoring Mechanics (go mod vendor & -mod=vendor)]] — Hermetic builds, offline CI pipelines, and vendored dependency auditing.
### 4. 📂 [[Dependency Injection & Decoupling Patterns|04. Dependency Injection & Decoupling Patterns]]
- [[Manual Constructor Dependency Injection (Idiomatic Go)]] — Creating explicit NewService(repo, client, logger) constructors without magic frameworks.
- [[Compile-Time Dependency Injection with Google Wire]] — Generating static dependency injection graphs at build time without reflection overhead.
- [[Functional Options Pattern for Flexible Configuration]] — Clean, extensible constructor configurations with functional options (WithTimeout, WithRetries).
- [[Interface-Driven Decoupling & Testability]] — Declaring consumer-side interfaces to enable seamless mock and fake substitutions in tests.
- [[Uber Fx Framework in Enterprise Go Services]] — Lifecycle management (OnStart, OnStop), dependency injection, and modular container bootstrapping.
### 5. 📂 [[Enterprise API Design & Evolution|05. Enterprise API Design & Evolution]]
- [[RESTful API Architecture with net-http and chi]] — Building fast, idiomatic REST services with sub-routers and context middlewares.
- [[gRPC and Protocol Buffers Service Architecture]] — Defining .proto service contracts, generating Go stubs, interceptors, and streaming RPCs.
- [[API Versioning Strategies (URL, Header, Subdomain)]] — Backward-compatible schema evolution, deprecation timelines, and field migrations.
- [[OpenAPI & Swagger Documentation Generation]] — Auto-generating OpenAPI specs from Go doc comments and declarative annotations.
- [[Backward Compatibility & Breaking Change Prevention]] — Protobuf field number rules, additive JSON extensions, and contract testing.
### 6. 📂 [[Code Architecture Anti-Patterns & Code Smells|06. Code Architecture Anti-Patterns & Code Smells]]
- [[The Global State & Singleton Anti-Pattern]] — Concurrency race hazards, testing pollution, and hidden coupling caused by global variables.
- [[The God Package & Package Clutter Anti-Pattern]] — Giant monolithic packages containing everything and dumping code into helpers/.
- [[The Package Stuttering Anti-Pattern]] — Redundant identifiers (user.UserService, client.ClientConfig) harming idiomatic Go readability.
- [[Cyclic Dependency Workaround Hacks]] — Dangerous anti-patterns: using init() hooks or type casting to bypass circular imports.
- [[Staff-Level Code Architecture Checklist]] — Pre-production architectural review checklist for high-scale enterprise Go services.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

