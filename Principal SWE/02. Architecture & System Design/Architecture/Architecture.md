---
title: Architecture
tags:
  - architecture
  - systems-architecture
  - clean-architecture
  - ddd
  - microservices
  - eip
  - saas
  - principal-swe
parent: "[[Architecture & System Design]]"
---

# 🏛️ Software Architecture & Enterprise Systems Engineering

Comprehensive, production-grade master architecture covering the complete spectrum of classical and modern styles, Clean/Hexagonal boundaries, Gang of Four (GoF) design patterns, high-performance concurrency patterns, Domain-Driven Design (DDD), distributed data patterns (CQRS, Event Sourcing, Sagas), microservice boundaries, API gateway architectures, fault tolerance & chaos engineering, Enterprise Integration Patterns (EIP), Multi-Tenant SaaS, anti-patterns refactoring, architectural governance, and Web3 blockchain across 14 master pillars and 146 specialized subdomains:

- **Architectural Styles:** Monoliths, Modular Monoliths, SOA vs Microservices, Serverless FaaS, Event-Driven, N-Tier, Microkernel plugins, Space-Based grids, and P2P.
- **Clean & Hexagonal Architecture:** The Dependency Rule, Entities, Use Cases, Ports & Adapters, Onion Architecture, Component Cohesion (REP, CCP, CRP), Coupling (ADP, SDP, SAP), and DTO boundary crossing.
- **GoF & Enterprise Design Patterns:** Creational (Factory, Builder, Singleton, Pool), Structural (Adapter, Bridge, Composite, Decorator, Facade, Flyweight, Proxy), Behavioral (Strategy, Observer, Command, State, Visitor, Iterator).
- **Concurrency & High-Performance Patterns:** Reactor & Proactor event demultiplexing, Active Object, Monitor Object, Half-Sync/Half-Async, Leader-Followers, Worker Pools, Double-Checked Locking, and Disruptor Ring Buffer.
- **Domain-Driven Design (DDD):** Ubiquitous Language, Bounded Contexts, Context Mapping (Shared Kernel, ACL), Entities, Value Objects, Aggregates & Invariants, Domain Events, Repositories, Domain Services, and Event Storming.
- **Distributed Data Patterns:** CQRS read/write segregation, Event Sourcing & event stores, Saga orchestration vs choreography, Transactional Outbox pattern, Change Data Capture (CDC Debezium), and Idempotent Consumers.
- **Microservices & Service Boundaries:** Decomposition by business capability, Database-per-service, Service Chassis, Service Discovery, OpenTelemetry distributed tracing, Envoy/Istio Service Mesh, BFF, and Strangler Fig migrations.
- **API Design & Gateway Architecture:** RESTful resource modeling & HATEOAS, GraphQL Schema Federation, gRPC Protobuf RPC, WebSockets vs SSE, API Gateways, Rate limiting, Webhooks, OAuth 2.1, and OpenAPI 3.1 contracts.
- **Resilience & Chaos Engineering:** Circuit Breaker state transitions, Bulkhead isolation, Exponential backoff with full jitter, Graceful degradation, Timeout deadline budgets, Dead Letter Queues, Load shedding, and Chaos Mesh.
- **Enterprise Integration Patterns (EIP):** Message Channels, Pipes & Filters, Content-Based Routers, Message Translators, Scatter-Gather aggregation, Routing Slips, Resequencers, and Claim Check patterns.
- **Multi-Tenant SaaS Architecture:** Multi-tenancy models (Silo vs Pool vs Bridge), Database-per-tenant vs Schema-per-tenant, Row-Level Security (RLS) isolation, Tenant context propagation, SCIM SAML SSO, Tenant metering, and Noisy Neighbor mitigation.
- **Architectural Anti-Patterns & Refactoring:** Distributed Monolith, God Service, Chatty I/O sprawl, Shared Database anti-pattern, Anemic Domain Model, Golden Hammer, Accidental complexity, Spaghetti code, and Technical Debt quantification.
- **Software Architect Leadership & Governance:** Architecture Decision Records (ADRs), Architecture Review Boards (ARB), RFC engineering processes, Evaluating trade-offs, Technology Radars, Architectural Fitness Functions (ArchUnit), Conway's Law, and Executive Communication.
- **Decentralized & Web3 Blockchain:** Consensus mechanisms (PoW, PoS), EVM internals, Solidity smart contract patterns, IPFS decentralized storage, Token standards (ERC-20, ERC-721), Zero-Knowledge Proofs (zk-SNARKs), and Layer 2 Rollups.

```text
Architecture
│
├── `01. Classical & Modern Architectural Styles`
│   ├── `01. Monolithic Architecture and Boundaries`
│   ├── `02. Modular Monolith Architecture`
│   ├── `03. Service Oriented Architecture Soa vs Microservices`
│   ├── `04. Microservices Architecture Core Invariants`
│   ├── `05. Serverless and Function As a Service Faas Topologies`
│   ├── `06. Event Driven Architecture Eda Pub Sub and Event Streams`
│   ├── `07. Layered N Tier Architecture and Vertical Slicing`
│   ├── `08. Microkernel and Plugin Architecture`
│   ├── `09. Space Based Architecture and in Memory Data Grids`
│   ├── `10. Peer to Peer P2p Distributed Architectures`
│   └── `11. Architectural Style Trade Off Matrix and Evaluation`
├── [[Clean, Hexagonal & Onion Architecture|02. Clean, Hexagonal & Onion Architecture]]
│   ├── `01. The Dependency Rule and Inversion of Control`
│   ├── `02. Entities, Business Rules, and Enterprise Invariants`
│   ├── `03. Use Cases and Interactors Application Business Rules`
│   ├── `04. Hexagonal Architecture Ports and Adapters`
│   ├── `05. Onion Architecture and Domain Core Purity`
│   ├── `06. Interface Adapters, Presenters, and Gateways`
│   ├── `07. Component Cohesion Principles Rep, Ccp, Crp`
│   ├── `08. Component Coupling Principles Adp, Sdp, Sap`
│   ├── `09. Boundaries Crossing and Data Transfer Objects Dtos`
│   ├── `10. Testing Strategies in Clean and Hexagonal Architectures`
│   ├── `11. Clean Architecture Framework Independence and Pragmatism`
│   └── `12. Refactoring Monoliths to Hexagonal Clean Architectures`
├── [[Gang of Four (gof) & Enterprise Design Patterns|03. Gang of Four (GoF) & Enterprise Design Patterns]]
│   ├── `01. Creational Patterns Factory Method and Abstract Factory`
│   ├── `02. Creational Patterns Builder and Prototype`
│   ├── `03. Creational Patterns Singleton and Object Pool`
│   ├── `04. Structural Patterns Adapter and Bridge`
│   ├── `05. Structural Patterns Composite and Decorator`
│   ├── `06. Structural Patterns Facade and Flyweight`
│   ├── `07. Structural Patterns Proxy Virtual, Remote, and Protection`
│   ├── `08. Behavioral Patterns Strategy and Template Method`
│   ├── `09. Behavioral Patterns Observer and Publish Subscribe`
│   ├── `10. Behavioral Patterns Command and Memento`
│   ├── `11. Behavioral Patterns State and Finite State Machines`
│   └── `12. Behavioral Patterns Visitor, Iterator, and Mediator`
├── `04. Concurrency & High-Performance Design Patterns`
│   ├── `01. Reactor Pattern and Event Demultiplexing`
│   ├── `02. Proactor Pattern and Asynchronous I O Completion`
│   ├── `03. Active Object Pattern and Asynchronous Method Invocation`
│   ├── `04. Monitor Object and Thread Synchronization`
│   ├── `05. Half Sync Half Async Architecture`
│   ├── `06. Leader Followers Thread Pool Pattern`
│   ├── `07. Worker Pool and Task Queue Paradigms`
│   ├── `08. Double Checked Locking and Thread Safe Lazy Initialization`
│   ├── `09. Guarded Suspension and Balking Patterns`
│   └── `10. Lmax Disruptor Ring Buffer and Mechanical Sympathy`
├── `05. Domain-Driven Design (DDD) & Strategic Modeling`
│   ├── `01. Ubiquitous Language and Core Domain Identification`
│   ├── `02. Bounded Contexts and Domain Boundaries`
│   ├── `03. Context Mapping Strategies and Integration Relationships`
│   ├── `04. Entities, Identity, and Mutability Lifecycles`
│   ├── `05. Value Objects and Immutability Standards`
│   ├── `06. Aggregates, Aggregate Roots, and Invariant Boundaries`
│   ├── `07. Domain Events and State Mutation Propagation`
│   ├── `08. Repositories, Factories, and Aggregate Persistence`
│   ├── `09. Domain Services vs Application Services`
│   └── `10. Event Storming and Collaborative Domain Modeling`
├── [[Distributed Data Patterns (cqrs, Event Sourcing, Sagas)|06. Distributed Data Patterns (CQRS, Event Sourcing, Sagas)]]
│   ├── `01. Command Query Responsibility Segregation CQRS Architecture`
│   ├── `02. Event Sourcing Mechanics and Event Store Design`
│   ├── `03. Saga Pattern Orchestration vs Choreography`
│   ├── `04. Transactional Outbox Pattern and Reliable Publishing`
│   ├── `05. Change Data Capture CDC with Debezium and Kafka`
│   ├── `06. Idempotent Consumer and Message Deduplication`
│   ├── `07. Materialized View Maintenance and Projection Engines`
│   ├── `08. Event Schema Evolution, Versioning, and Upcasting`
│   ├── `09. Eliminating the Dual Write Problem in Distributed Systems`
│   └── `10. Event Driven Consistency Models and Replayability`
├── `07. Microservice Architecture & Service Boundaries`
│   ├── `01. Service Decomposition Strategies and Boundary Definition`
│   ├── `02. Database Per Service vs Shared Database Anti Pattern`
│   ├── `03. Microservice Chassis and Standardized Service Templates`
│   ├── `04. Dynamic Service Discovery and Registration Topologies`
│   ├── `05. Distributed Tracing, Span Contexts, and Opentelemetry`
│   ├── `06. Service Mesh Architecture Data Plane vs Control Plane`
│   ├── `07. Backend for Frontend BFF Architecture Pattern`
│   ├── `08. Distributed Configuration Management and Dynamic Reloading`
│   ├── `09. Progressive Delivery, Canary Deployments, and Traffic Splitting`
│   └── `10. The Strangler Fig Pattern for Monolith Decomposition`
├── `08. API Design & Gateway Architecture`
│   ├── `01. RESTful API Resource Modeling and Hateoas`
│   ├── `GraphQL Architecture & Apollo Federation Ecosystem`
│   ├── `03. GRPC High Performance RPC and Protocol Buffers`
│   ├── `04. Websockets and Server Sent Events SSE Real Time Protocols`
│   ├── `05. API Gateway Core Patterns and Reverse Proxies`
│   ├── `06. Token Bucket and Leaky Bucket Rate Limiting`
│   ├── `07. Webhook Architecture, Delivery Engines, and Signatures`
│   ├── `08. API Versioning Strategies and Deprecation Lifecycles`
│   ├── `09. API Security Architecture Oauth 2.1, Oidc, and Mtls`
│   ├── `10. Consumer Driven Contract Testing with Pact`
│   ├── `11. Openapi 3.1 Standards and Interactive Documentation`
│   ├── `12. API Mocking, Service Virtualization, and Sandbox APIs`
│   └── `13. High Performance Reverse Proxies Nginx and Envoy`
├── [[Resilience, Fault Tolerance & Chaos Engineering|09. Resilience, Fault Tolerance & Chaos Engineering]]
│   ├── `01. Circuit Breaker Pattern and State Transitions`
│   ├── `02. Bulkhead Pattern and Resource Isolation`
│   ├── `03. Exponential Backoff with Full Jitter`
│   ├── `04. Graceful Degradation and Fallback Strategies`
│   ├── `05. Distributed Timeouts and Deadline Propagation`
│   ├── `06. Dead Letter Queues DLQ and Poison Pill Handling`
│   ├── `07. Rate Limiting, Load Shedding, and Concurrency Limits`
│   ├── `08. Chaos Engineering Principles and Fault Injection`
│   ├── `09. High Availability Topologies and Cross Region Failover`
│   └── `10. Health Check Probes and Synthetic Monitoring`
├── [[Enterprise Integration Patterns (eip)|10. Enterprise Integration Patterns (EIP)]]
│   ├── `01. Message Channels and Point to Point Topologies`
│   ├── `02. Pipes and Filters Architecture Pattern`
│   ├── `03. Message Routing Content Based and Filter Routers`
│   ├── `04. Message Transformation, Translator, and Normalizer`
│   ├── `05. Scatter Gather Pattern and Parallel Aggregation`
│   ├── `06. Recipient List and Dynamic Routing Slips`
│   ├── `07. Resequencer, Aggregator, and Message Correlator`
│   ├── `08. Message Broker Topologies and Enterprise Bridges`
│   ├── `09. Claim Check Pattern for Large Payload Messaging`
│   └── `10. Idempotent Message Receiver and Process Manager`
├── `11. Multi-Tenant SaaS & Data Isolation Architecture`
│   ├── `01. Multi Tenancy Architectural Models Silo, Pool, Bridge`
│   ├── `02. Multi Tenant Storage Partitioning and Isolation Models`
│   ├── `03. Row Level Security RLS for Multi Tenant Data Isolation`
│   ├── `04. Dynamic Tenant Context Propagation in Middleware`
│   ├── `05. Multi Tenant Identity, Scim, and Saml Sso`
│   ├── `06. Tenant Aware Caching, Sharding, and Key Namespacing`
│   ├── `07. Multi Tenant Usage Metering, Quotas, and Billing Engines`
│   ├── `08. Noisy Neighbor Problem and Fair Resource Queuing`
│   ├── `09. Tenant Data Export, Backup, and Gdpr Compliance`
│   └── `10. Custom Domain Routing, Ssl Automation, and Whitelabeling`
├── `12. Architectural Anti-Patterns & Technical Debt Refactoring`
│   ├── `01. The Distributed Monolith Anti Pattern`
│   ├── `02. God Service and Megaservice Architecture Pitfall`
│   ├── `03. Chatty I O and Microservice Sprawl Pitfall`
│   ├── `04. Shared Database Anti Pattern and Database Coupling`
│   ├── `05. Anemic Domain Model Anti Pattern and Procedural Leaks`
│   ├── `06. The Golden Hammer and Trend Driven Architecture`
│   ├── `07. Accidental vs Essential Complexity and Over Engineering`
│   ├── `08. Big Ball of Mud and Spaghetti Code Rot`
│   ├── `09. Vendor Lock In, Cloud Sprawl, and Leaky Abstractions`
│   └── `10. Systemic Technical Debt Quantification and Refactoring Roadmaps`
├── `13. Software Architect Leadership & Governance`
│   ├── `01. Architecture Decision Records ADRs and Decision Governance`
│   ├── `02. Architecture Review Boards ARB and Lightweight Governance`
│   ├── `03. Enterprise RFC Request for Comments Workflows`
│   ├── `04. Evaluating Architectural Trade Offs Cost, Perf, Complexity`
│   ├── `05. Technology Radars, Standardization, and Paved Roads`
│   ├── `06. Evolutionary Architecture and Automated Fitness Functions`
│   ├── `07. Aligning Software Architecture with Business Strategy`
│   ├── `08. Conway's Law, Team Topologies, and Reverse Conway Maneuver`
│   ├── `09. Mentoring, Coaching, and Growing Senior Engineers`
│   └── `10. Executive Communication and Board Level Presentations`
└── [[Decentralized, Web3 & Blockchain Architectures|14. Decentralized, Web3 & Blockchain Architectures]]
│   ├── `01. Blockchain Consensus Mechanisms Pow, Pos, and Tendermint`
│   ├── `02. Ethereum Virtual Machine EVM and Smart Contract Architecture`
│   ├── `03. Solidity Smart Contract Patterns and Upgradability`
│   ├── `04. Decentralized Storage Networks Ipfs, Arweave, and Filecoin`
│   ├── `05. Token Standards Erc 20, Erc 721, and Erc 1155`
│   ├── `06. Zero Knowledge Proofs Zk Snarks and Privacy Scaling`
│   ├── `07. Layer 2 Scaling Rollups Optimistic vs Zk Rollups`
│   └── `08. Web3 Smart Contract Security Auditing and Reentrancy`
```

---

## 🏛️ Core Knowledge Pillars

### 📂 [[Modern Architectural Styles|21. Modern Architectural Styles]]
- 📂 `01. Microservices Architecture Core Invariants`
- 📂 `02. Serverless and Function As a Service Faas Topologies`
- 📂 `03. Event Driven Architecture Eda Pub Sub and Event Streams`
- 📂 `04. Space Based Architecture and in Memory Data Grids`

### 📂 [[Clean, Hexagonal & Onion Architecture|02. Clean, Hexagonal & Onion Architecture]]
- 📂 `01. The Dependency Rule and Inversion of Control` — Source code dependencies must point only inward toward higher-level policies, decoupling business rules from UI, databases, and frameworks.
- 📂 `02. Entities, Business Rules, and Enterprise Invariants` — Encapsulating pure enterprise business logic, rich domain models vs anemic structures, and ensuring business invariants hold true independent of technology.
- 📂 `03. Use Cases and Interactors Application Business Rules` — Orchestrating domain entities, defining application-specific workflows, handling input/output boundary DTOs, and keeping use cases framework-free.
- 📂 `04. Hexagonal Architecture Ports and Adapters` — Driving (Primary) Ports and Adapters (HTTP, CLI) vs Driven (Secondary) Ports and Adapters (PostgreSQL, Kafka), mockability, and swappable infrastructure.
- 📂 `05. Onion Architecture and Domain Core Purity` — Domain Model at the center, Domain Services, Application Services, and Outer UI/Infrastructure ring, ensuring zero database dependencies in the domain core.
- 📂 `06. Interface Adapters, Presenters, and Gateways` — Converting data from use case format to web/UI format, implementing repository interfaces in data access gateways, and separating presentation logic.
- 📂 `07. Component Cohesion Principles Rep, Ccp, Crp` — Release-Reuse Equivalence Principle, Common Closure Principle (classes that change together belong together), and Common Reuse Principle.
- 📂 `08. Component Coupling Principles Adp, Sdp, Sap` — Acyclic dependency graphs, the Stable Dependencies Principle (depend in the direction of stability), and the Main Sequence (Abstractness vs Instability).
- 📂 `09. Boundaries Crossing and Data Transfer Objects Dtos` — Preventing domain entity leaks into web responses, strict boundary DTO mapping, and eliminating accidental coupling between database schema and API payloads.
- 📂 `10. Testing Strategies in Clean and Hexagonal Architectures` — Testing use cases in complete isolation using in-memory mock/fake adapters, testing driving adapters with contract tests, and achieving 100% test velocity.
- 📂 `11. Clean Architecture Framework Independence and Pragmatism` — Treating frameworks as implementation details, protecting core business code from ORM/Web framework breaking upgrades, and pragmatic exceptions.
- 📂 `12. Refactoring Monoliths to Hexagonal Clean Architectures` — Identifying domain boundaries in legacy code, introducing ports around database calls, extracting use case interactors, and verifying with golden master tests.
### 📂 [[Gang of Four (gof) & Enterprise Design Patterns|03. Gang of Four (GoF) & Enterprise Design Patterns]]
- 📂 `01. Creational Patterns Factory Method and Abstract Factory` — Decoupling object creation from consumption, abstract creation interfaces, parameterizing factory families, and dependency injection integration.
- 📂 `02. Creational Patterns Builder and Prototype` — Constructing complex composite objects step-by-step, immutable object construction, and deep cloning with the Prototype pattern.
- 📂 `03. Creational Patterns Singleton and Object Pool` — Lazy initialization, double-checked locking, thread safety, connection pooling (DB, HTTP), and why singletons are often considered anti-patterns.
- 📂 `04. Structural Patterns Adapter and Bridge` — Converting incompatible interfaces, wrapping legacy APIs, and separating an abstraction from its implementation with the Bridge pattern.
- 📂 `05. Structural Patterns Composite and Decorator` — Treating individual objects and compositions uniformly, dynamic runtime behavior wrapping with Decorator, and combining with stream I/O.
- 📂 `06. Structural Patterns Facade and Flyweight` — Providing simplified entry points into complex subsystem graphs (Facade), and sharing fine-grained immutable state to minimize memory footprint (Flyweight).
- 📂 `07. Structural Patterns Proxy Virtual, Remote, and Protection` — Lazy loading large resources with Virtual Proxy, remote RPC stubs, access control verification with Protection Proxy, and dynamic proxies in frameworks.
- 📂 `08. Behavioral Patterns Strategy and Template Method` — Encapsulating interchangeable algorithms behind interfaces (Strategy), and defining algorithmic skeletons with polymorphic hooks (Template Method).
- 📂 `09. Behavioral Patterns Observer and Publish Subscribe` — One-to-many dependency notifications, thread-safe subject state changes, memory leak prevention (Lapsed Listener Problem), and reactive event streams.
- 📂 `10. Behavioral Patterns Command and Memento` — Encapsulating operations as first-class objects (Command), queuing and scheduling commands, and restoring internal object state with Memento snapshots.
- 📂 `11. Behavioral Patterns State and Finite State Machines` — Encapsulating state-specific behavior in dedicated classes, eliminating massive switch statements, and transition logic in workflow engines.
- 📂 `12. Behavioral Patterns Visitor, Iterator, and Mediator` — Adding new operations to object structures without modifying them (Visitor), uniform traversal (Iterator), and decoupling complex multi-object communication (Mediator).
### 📂 [[High-Performance Design Patterns|17. High-Performance Design Patterns]]
- 📂 `01. Proactor Pattern and Asynchronous I O Completion`
- 📂 `02. Leader Followers Thread Pool Pattern`
- 📂 `03. Worker Pool and Task Queue Paradigms`
- 📂 `04. Lmax Disruptor Ring Buffer and Mechanical Sympathy`

### 📂 [[Domain-Driven Design|05. Domain-Driven Design]]
- 📂 `01. Entities, Identity, and Mutability Lifecycles`
- 📂 `02. Value Objects and Immutability Standards`
- 📂 `03. Aggregates, Aggregate Roots, and Invariant Boundaries`
- 📂 `04. Domain Events and State Mutation Propagation`
- 📂 `05. Repositories, Factories, and Aggregate Persistence`
- 📂 `06. Domain Services vs Application Services`

### 📂 [[Strategic Modeling|15. Strategic Modeling]]
- 📂 `01. Ubiquitous Language and Core Domain Identification`
- 📂 `02. Bounded Contexts and Domain Boundaries`
- 📂 `03. Context Mapping Strategies and Integration Relationships`
- 📂 `04. Event Storming and Collaborative Domain Modeling`

### 📂 [[Distributed Data Patterns (cqrs, Event Sourcing, Sagas)|06. Distributed Data Patterns (CQRS, Event Sourcing, Sagas)]]
- 📂 `01. Command Query Responsibility Segregation CQRS Architecture` — Separating write commands (optimizing for consistency/business rules) from read queries (optimizing for high-performance denormalized reads), and read-model sync.
- 📂 `02. Event Sourcing Mechanics and Event Store Design` — Storing system state as a sequence of immutable events rather than current state, replaying events to rebuild state, temporal queries, and snapshotting strategies.
- 📂 `03. Saga Pattern Orchestration vs Choreography` — Managing multi-service distributed transactions without 2PC: Centralized Orchestration (State Machine) vs Decentralized Choreography (Event Routing), and compensating actions.
- 📂 `04. Transactional Outbox Pattern and Reliable Publishing` — Writing business data and outbound event messages atomically within the same database transaction, polling outbox tables, and eliminating dual-write failure windows.
- 📂 `05. Change Data Capture CDC with Debezium and Kafka` — Streaming row-level database WAL mutations directly into Kafka topics without application-level polling, zero-latency event streaming, and schema evolution.
- 📂 `06. Idempotent Consumer and Message Deduplication` — Handling at-least-once message delivery guarantees: Idempotency keys, unique database constraints, Redis deduplication windows, and stateful tracking.
- 📂 `07. Materialized View Maintenance and Projection Engines` — Asynchronous read model projection builders, rebuilding read stores from event streams, handling projection lag, and serving sub-millisecond queries.
- 📂 `08. Event Schema Evolution, Versioning, and Upcasting` — Evolving immutable event structures over time: Adding optional fields, event upcasters converting legacy events at runtime, and avoiding destructive migrations.
- 📂 `09. Eliminating the Dual Write Problem in Distributed Systems` — Why writing to database and message broker concurrently in application code causes silent data loss, and solving via Outbox/CDC/Listen-to-Yourself patterns.
- 📂 `10. Event Driven Consistency Models and Replayability` — Read-your-own-writes consistency in CQRS, tracking event sequence offsets, replaying event history for disaster recovery, and auditing forensic state.
### 📂 [[Microservice Architecture|07. Microservice Architecture]]
- 📂 `01. Microservice Chassis and Standardized Service Templates`
- 📂 `02. Dynamic Service Discovery and Registration Topologies`
- 📂 `03. Distributed Tracing, Span Contexts, and Opentelemetry`
- 📂 `04. Service Mesh Architecture Data Plane vs Control Plane`
- 📂 `05. Distributed Configuration Management and Dynamic Reloading`
- 📂 `06. Progressive Delivery, Canary Deployments, and Traffic Splitting`

### 📂 [[Service Boundaries|20. Service Boundaries]]
- 📂 `01. Service Decomposition Strategies and Boundary Definition`
- 📂 `02. Database Per Service vs Shared Database Anti Pattern`
- 📂 `03. Backend for Frontend BFF Architecture Pattern`
- 📂 `04. The Strangler Fig Pattern for Monolith Decomposition`

### 📂 [[API Design|08. API Design]]
- 📂 `01. RESTful API Resource Modeling and Hateoas`
- 📂 [[GraphQL Architecture|02. GraphQL Architecture]]
- 📂 `03. GRPC High Performance RPC and Protocol Buffers`
- 📂 `04. Websockets and Server Sent Events SSE Real Time Protocols`
- 📂 `05. Webhook Architecture, Delivery Engines, and Signatures`
- 📂 `06. API Versioning Strategies and Deprecation Lifecycles`
- 📂 `07. API Security Architecture Oauth 2.1, Oidc, and Mtls`
- 📂 `08. Consumer Driven Contract Testing with Pact`
- 📂 `09. Openapi 3.1 Standards and Interactive Documentation`
- 📂 `10. API Mocking, Service Virtualization, and Sandbox APIs`
- 📂 [[Apollo Federation Ecosystem|11. Apollo Federation Ecosystem]]

### 📂 [[Gateway Architecture|22. Gateway Architecture]]
- 📂 `01. API Gateway Core Patterns and Reverse Proxies`
- 📂 `02. Token Bucket and Leaky Bucket Rate Limiting`
- 📂 `03. High Performance Reverse Proxies Nginx and Envoy`

### 📂 [[Resilience, Fault Tolerance & Chaos Engineering|09. Resilience, Fault Tolerance & Chaos Engineering]]
- 📂 `01. Circuit Breaker Pattern and State Transitions` — Preventing cascading failures during downstream outages, sliding window failure rate thresholds, slow call thresholds, and automatic recovery probing.
- 📂 `02. Bulkhead Pattern and Resource Isolation` — Isolating thread pools and connection pools per downstream dependency so a slow external API does not exhaust resources for the entire application.
- 📂 `03. Exponential Backoff with Full Jitter` — Why naive retries cause self-inflicted DDoS attacks, calculating exponential backoff with full jitter (AWS algorithm), and decorrelated jitter.
- 📂 `04. Graceful Degradation and Fallback Strategies` — Returning cached stale data, degraded static responses, disabling non-essential UI features (recommendations/reviews), and prioritizing core user flows.
- 📂 `05. Distributed Timeouts and Deadline Propagation` — Setting strict client-side timeouts, propagating remaining deadline budgets across downstream microservice chains, and cancelling dead work early.
- 📂 `06. Dead Letter Queues DLQ and Poison Pill Handling` — Isolating unparseable or crash-inducing messages without blocking consumer pipelines, maximum retry thresholds, and manual replay tooling.
- 📂 `07. Rate Limiting, Load Shedding, and Concurrency Limits` — Little's Law, monitoring queue latency and CPU saturation, dynamically shedding incoming low-priority requests to protect system availability.
- 📂 `08. Chaos Engineering Principles and Fault Injection` — Formulating steady-state hypotheses, introducing real-world turbulence (network latency, killed nodes, packet drop), and validating resilience.
- 📂 `09. High Availability Topologies and Cross Region Failover` — Active-Passive vs Active-Active multi-region deployments, DNS failover with Route 53, data replication lag, and Recovery Time/Point Objectives (RTO/RPO).
- 📂 `10. Health Check Probes and Synthetic Monitoring` — Designing non-cascading health check endpoints, verifying critical downstream dependencies safely, and synthetic browser-driven user journey probes.
### 📂 [[Enterprise Integration Patterns (eip)|10. Enterprise Integration Patterns (EIP)]]
- 📂 `01. Message Channels and Point to Point Topologies` — Establishing decoupled communication channels between heterogeneous enterprise applications, channel adapters, and message endpoints.
- 📂 `02. Pipes and Filters Architecture Pattern` — Building modular data processing pipelines where independent filter components process and pass messages through unidirectional pipes.
- 📂 `03. Message Routing Content Based and Filter Routers` — Inspecting message headers and payload contents to dynamically route messages to specific destinations without sender awareness.
- 📂 `04. Message Transformation, Translator, and Normalizer` — Translating disparate enterprise data formats (XML, JSON, Protobuf, CSV), wrapping payloads in canonical message envelopes, and normalizers.
- 📂 `05. Scatter Gather Pattern and Parallel Aggregation` — Broadcasting a request message to multiple vendor/service endpoints in parallel, aggregating responses, and selecting the best offer with timeouts.
- 📂 `06. Recipient List and Dynamic Routing Slips` — Calculating a dynamic list of recipients based on runtime business rules, and attaching a sequential itinerary (Routing Slip) to the message.
- 📂 `07. Resequencer, Aggregator, and Message Correlator` — Reordering out-of-order event streams based on sequence numbers, and correlating related messages with unique `Correlation-ID` headers.
- 📂 `08. Message Broker Topologies and Enterprise Bridges` — Bridging incompatible message brokers (Kafka to RabbitMQ, SQS to GCP PubSub), message routing hubs, and store-and-forward persistence.
- 📂 `09. Claim Check Pattern for Large Payload Messaging` — Splitting massive message payloads into external cloud blob storage (S3) and passing only a lightweight reference token (claim check) over the message bus.
- 📂 `10. Idempotent Message Receiver and Process Manager` — Ensuring exactly-once business processing semantics on at-least-once message brokers, and orchestrating complex stateful business processes.
### 📂 [[Multi-Tenant SaaS|11. Multi-Tenant SaaS]]
- 📂 `01. Multi Tenancy Architectural Models Silo, Pool, Bridge`
- 📂 `02. Dynamic Tenant Context Propagation in Middleware`
- 📂 `03. Multi Tenant Identity, Scim, and Saml Sso`
- 📂 `04. Multi Tenant Usage Metering, Quotas, and Billing Engines`
- 📂 `05. Noisy Neighbor Problem and Fair Resource Queuing`
- 📂 `06. Custom Domain Routing, Ssl Automation, and Whitelabeling`

### 📂 [[Data Isolation Architecture|16. Data Isolation Architecture]]
- 📂 `01. Multi Tenant Storage Partitioning and Isolation Models`
- 📂 `02. Row Level Security RLS for Multi Tenant Data Isolation`
- 📂 `03. Tenant Aware Caching, Sharding, and Key Namespacing`
- 📂 `04. Tenant Data Export, Backup, and Gdpr Compliance`

### 📂 [[Architectural Anti-Patterns|12. Architectural Anti-Patterns]]
- 📂 `01. The Distributed Monolith Anti Pattern`
- 📂 `02. God Service and Megaservice Architecture Pitfall`
- 📂 `03. Chatty I O and Microservice Sprawl Pitfall`
- 📂 `04. Shared Database Anti Pattern and Database Coupling`
- 📂 `05. Anemic Domain Model Anti Pattern and Procedural Leaks`
- 📂 `06. The Golden Hammer and Trend Driven Architecture`

### 📂 [[Technical Debt Refactoring|18. Technical Debt Refactoring]]
- 📂 `01. Accidental vs Essential Complexity and Over Engineering`
- 📂 `02. Big Ball of Mud and Spaghetti Code Rot`
- 📂 `03. Vendor Lock In, Cloud Sprawl, and Leaky Abstractions`
- 📂 `04. Systemic Technical Debt Quantification and Refactoring Roadmaps`

### 📂 [[Software Architect Leadership|13. Software Architect Leadership]]
- 📂 `01. Evaluating Architectural Trade Offs Cost, Perf, Complexity`
- 📂 `02. Aligning Software Architecture with Business Strategy`
- 📂 `03. Conway's Law, Team Topologies, and Reverse Conway Maneuver`
- 📂 `04. Mentoring, Coaching, and Growing Senior Engineers`
- 📂 `05. Executive Communication and Board Level Presentations`

### 📂 [[Decentralized, Web3 & Blockchain Architectures|14. Decentralized, Web3 & Blockchain Architectures]]
- 📂 `01. Blockchain Consensus Mechanisms Pow, Pos, and Tendermint` — Nakamoto consensus, Byzantine Fault Tolerance (BFT), validator slashing, finality guarantees, and high-speed consensus (Tendermint/Raft).
- 📂 `02. Ethereum Virtual Machine EVM and Smart Contract Architecture` — Stack-based EVM execution, storage vs memory vs calldata, gas metering, state tries (Merkle Patricia Trie), and smart contract compilation.
- 📂 `03. Solidity Smart Contract Patterns and Upgradability` — Writing gas-efficient Solidity code, Diamond multi-facet proxy pattern (ERC-2535), Universal Upgradeable Proxy Standard (UUPS), and storage layout collisions.
- 📂 `04. Decentralized Storage Networks Ipfs, Arweave, and Filecoin` — Content Identifiers (CID) based on cryptographic hashes, peer-to-peer pinning networks (IPFS), and permanent immutable storage architectures (Arweave).
- 📂 `05. Token Standards Erc 20, Erc 721, and Erc 1155` — Fungible token standard (ERC-20), Non-Fungible token standard (ERC-721), Multi-token batch operations (ERC-1155), and metadata schemas.
- 📂 `06. Zero Knowledge Proofs Zk Snarks and Privacy Scaling` — Prover-Verifier mechanics, arithmetic circuits, non-interactive zero-knowledge proofs, privacy-preserving transactions, and ZK-rollups.
- 📂 `07. Layer 2 Scaling Rollups Optimistic vs Zk Rollups` — Off-chain transaction execution, fraud proofs (Arbitrum, Optimism) vs validity proofs (zkSync, Starknet), sequencer centralization, and bridging.
- 📂 `08. Web3 Smart Contract Security Auditing and Reentrancy` — Reentrancy attack mechanics and CEI (Checks-Effects-Interactions) pattern, flash loan vulnerability exploitation, oracle manipulation, and auditing tools (Slither, Foundry).

---

## 🔗 Navigation
- ⬆️ Parent: [[Architecture & System Design]]
- 🏛️ High-Scale Systems: `System Design`
- 🎯 Production Standards: `Best Practices`
- 💻 Computer Science Foundations: `Computer Science`
- 🚀 Infrastructure & DevOps: `DevOps`

---

## 🗂️ Topics

- [[API Design]]
- [[Architectural Anti-Patterns]]
- [[Architecture Governance]]
- [[Classical Architectural Styles]]
- [[Clean, Hexagonal & Onion Architecture]]
- [[Concurrency Architecture Patterns]]
- [[Data Isolation Architecture]]
- [[Decentralized, Web3 & Blockchain Architectures]]
- [[Distributed Data Patterns (cqrs, Event Sourcing, Sagas)]]
- [[Domain-Driven Design]]
- [[Enterprise Integration Patterns (eip)]]
- [[Gang of Four (gof) & Enterprise Design Patterns]]
- [[Gateway Architecture]]
- [[High-Performance Design Patterns]]
- [[Microservice Architecture]]
- [[Modern Architectural Styles]]
- [[Multi-Tenant SaaS]]
- [[Resilience, Fault Tolerance & Chaos Engineering]]
- [[Service Boundaries]]
- [[Software Architect Leadership]]
- [[Strategic Modeling]]
- [[Technical Debt Refactoring]]
