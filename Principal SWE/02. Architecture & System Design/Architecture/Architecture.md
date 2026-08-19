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

### 1. 📂 `01. Classical & Modern Architectural Styles`
- 📂 `01. Monolithic Architecture and Boundaries` — Shared memory execution, in-process function calls, deployment coupling, database connection pool limits, and vertical scaling boundaries.
- 📂 `02. Modular Monolith Architecture` — Enforcing strict compile-time domain boundaries, internal event buses, avoiding circular dependencies, and facilitating future service extraction.
- 📂 `03. Service Oriented Architecture Soa vs Microservices` — Heavyweight centralized ESB routing vs lightweight dumb pipes and smart endpoints in microservices, service contracts (WSDL/JSON), and enterprise governance.
- 📂 `04. Microservices Architecture Core Invariants` — Single Responsibility at service scale, autonomous deployment pipelines, decentralized data management, polyglot persistence, and organizational Conway's Law alignment.
- 📂 `05. Serverless and Function As a Service Faas Topologies` — Ephemeral execution environments, cold start latencies, stateless computation, event-source mappings, fine-grained pay-per-execution economics, and orchestration.
- 📂 `06. Event Driven Architecture Eda Pub Sub and Event Streams` — Asynchronous messaging, decoupling producers from consumers, temporal decoupling, event broker topologies (Kafka, RabbitMQ), and eventual consistency models.
- 📂 `07. Layered N Tier Architecture and Vertical Slicing` — Horizontal layer separation, dependency flow rules, open vs closed layers, performance hop penalties, and vertical slice architecture by feature.
- 📂 `08. Microkernel and Plugin Architecture` — Core system minimal functionality, plugin contracts and interfaces, dynamic plugin loading, extension points, and isolating third-party runtime extensions.
- 📂 `09. Space Based Architecture and in Memory Data Grids` — Eliminating database bottlenecks via in-memory replicated data grids, processing units, virtualized middleware, and extreme linear write scalability.
- 📂 `10. Peer to Peer P2p Distributed Architectures` — Symmetric node roles, BitTorrent protocol, Kademlia distributed hash tables, gossip protocol node discovery, and Byzantine fault tolerance.
- 📂 `11. Architectural Style Trade Off Matrix and Evaluation` — Evaluating styles against architectural characteristics: Agility, Deployability, Performance, Scalability, Testability, Fault Tolerance, and Operational Cost.
### 2. 📂 [[Clean, Hexagonal & Onion Architecture|02. Clean, Hexagonal & Onion Architecture]]
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
### 3. 📂 [[Gang of Four (gof) & Enterprise Design Patterns|03. Gang of Four (GoF) & Enterprise Design Patterns]]
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
### 4. 📂 `04. Concurrency & High-Performance Design Patterns`
- 📂 `01. Reactor Pattern and Event Demultiplexing` — Event loop architecture (Node.js, Netty, Nginx), synchronous demultiplexing with `epoll`/`kqueue`, non-blocking I/O event dispatching, and handler execution.
- 📂 `02. Proactor Pattern and Asynchronous I O Completion` — Initiating asynchronous I/O operations (Windows IOCP, Linux io_uring), operating system kernel completion notification, and completion handler callbacks.
- 📂 `03. Active Object Pattern and Asynchronous Method Invocation` — Client thread proxy, activation list queue, scheduler thread, servant object execution, and returning asynchronous Future promises.
- 📂 `04. Monitor Object and Thread Synchronization` — Encapsulating shared state with critical sections, monitor locks, condition variables, wait-notify synchronization, and preventing race conditions.
- 📂 `05. Half Sync Half Async Architecture` — Decoupling high-concurrency async network I/O layers from blocking synchronous database/business processing layers via thread-safe queues.
- 📂 `06. Leader Followers Thread Pool Pattern` — One thread acts as the leader waiting for incoming I/O events while followers wait on a synchronization condition, zero-overhead handoffs, and throughput.
- 📂 `07. Worker Pool and Task Queue Paradigms` — Fixed vs elastic worker pools, backpressure handling on queue saturation, work rejection policies, and thread pool starvation prevention.
- 📂 `08. Double Checked Locking and Thread Safe Lazy Initialization` — Volatile read-write semantics, preventing uninitialized memory reads across CPU cores, and modern language-specific idioms (Go `sync.Once`, Java enum).
- 📂 `09. Guarded Suspension and Balking Patterns` — Suspending thread execution until a precondition is met (Guarded Suspension), and immediately returning when object state is inappropriate (Balking).
- 📂 `10. Lmax Disruptor Ring Buffer and Mechanical Sympathy` — Lock-free bounded ring buffers, avoiding cache line false sharing via padding, memory sequence barriers, and single-writer principle for ultra-low latency.
### 5. 📂 `05. Domain-Driven Design (DDD) & Strategic Modeling`
- 📂 `01. Ubiquitous Language and Core Domain Identification` — Creating a shared model-driven vocabulary between domain experts and software engineers, eliminating translation ambiguities, and identifying the Core Domain.
- 📂 `02. Bounded Contexts and Domain Boundaries` — Defining explicit boundaries where a specific domain model applies, separating divergent concepts (e.g. User in Auth vs Customer in Billing), and context isolation.
- 📂 `03. Context Mapping Strategies and Integration Relationships` — Modeling relationships between Bounded Contexts: Partnership, Shared Kernel, Customer-Supplier, Conformist, Open Host Service (OHS), Published Language, and Separate Ways.
- 📂 `04. Entities, Identity, and Mutability Lifecycles` — Distinguishing entities by unique persistent ID rather than attributes, managing lifecycle transitions, and enforcing domain rules during mutations.
- 📂 `05. Value Objects and Immutability Standards` — Self-validating immutable domain building blocks (Money, Address, Email), attribute-based equality, whole value encapsulation, and eliminating primitive obsession.
- 📂 `06. Aggregates, Aggregate Roots, and Invariant Boundaries` — Enforcing transactional consistency within a single aggregate, selecting the Aggregate Root, referencing other aggregates only by ID, and small aggregate sizing.
- 📂 `07. Domain Events and State Mutation Propagation` — Modeling significant business occurrences as immutable past events (`OrderPlaced`, `PaymentReceived`), in-process event dispatching, and asynchronous propagation.
- 📂 `08. Repositories, Factories, and Aggregate Persistence` — Collection-oriented vs persistence-oriented repository interfaces, reconstructing aggregates via Factories, and avoiding leaky database queries in repositories.
- 📂 `09. Domain Services vs Application Services` — Pure stateless business logic involving multiple aggregates (Domain Services) vs use-case transaction orchestration, security, and DTO conversion (Application Services).
- 📂 `10. Event Storming and Collaborative Domain Modeling` — Facilitating collaborative rapid modeling with domain experts: mapping Domain Events (orange), Commands (blue), Aggregates (yellow), Read Models (green), and Policies (lilac).
### 6. 📂 [[Distributed Data Patterns (cqrs, Event Sourcing, Sagas)|06. Distributed Data Patterns (CQRS, Event Sourcing, Sagas)]]
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
### 7. 📂 `07. Microservice Architecture & Service Boundaries`
- 📂 `01. Service Decomposition Strategies and Boundary Definition` — Decomposing monolithic systems by business capability or DDD bounded contexts, evaluating service cohesion, and preventing microservice fragmentation.
- 📂 `02. Database Per Service vs Shared Database Anti Pattern` — Enforcing private data stores per microservice, eliminating cross-service database joins, exposing data via APIs/Events, and managing schema migrations.
- 📂 `03. Microservice Chassis and Standardized Service Templates` — Shared baseline cross-cutting concerns: Health checks, structured JSON logging, metrics, distributed tracing headers, security middleware, and configuration.
- 📂 `04. Dynamic Service Discovery and Registration Topologies` — Client-side vs server-side service discovery, service registry heartbeat checks, dynamic DNS resolution, and load balancing across healthy instances.
- 📂 `05. Distributed Tracing, Span Contexts, and Opentelemetry` — Propagating `traceparent` and `tracestate` HTTP/gRPC headers across microservices, spans, traces, Jaeger/Tempo backends, and latency bottleneck isolation.
- 📂 `06. Service Mesh Architecture Data Plane vs Control Plane` — Sidecar proxy architecture, automatic mTLS mutual authentication, intelligent traffic shifting, fault injection, and operational visibility without code changes.
- 📂 `07. Backend for Frontend BFF Architecture Pattern` — Tailoring backend APIs to specific client form factors (iOS, Android, Desktop Web), optimizing over-fetching/under-fetching, and decoupling client releases.
- 📂 `08. Distributed Configuration Management and Dynamic Reloading` — Centralized configuration stores (Consul, Spring Cloud Config, etcd), dynamic hot-reloading of configs without redeploying, and feature flag management.
- 📂 `09. Progressive Delivery, Canary Deployments, and Traffic Splitting` — Automated canary analysis (Kayenta, Argo Rollouts), routing 5% of production traffic to new versions, monitoring error rates, and automated rollback.
- 📂 `10. The Strangler Fig Pattern for Monolith Decomposition` — Intercepting inbound traffic at the edge proxy/API gateway, routing legacy endpoints to the monolith while routing new features to microservices until sunset.
### 8. 📂 `08. API Design & Gateway Architecture`
- 📂 `01. RESTful API Resource Modeling and Hateoas` — Resource-oriented URIs, proper HTTP status codes (200, 201, 204, 400, 401, 403, 404, 409, 429), safe vs idempotent methods, and hypermedia controls.
- 📂 `GraphQL Architecture & Apollo Federation Ecosystem` — Queries, Mutations, Subscriptions, solving the N+1 problem with DataLoader, Apollo Federation subgraphs, and supergraph gateway composition.
- 📂 `03. GRPC High Performance RPC and Protocol Buffers` — Binary Protobuf serialization, HTTP/2 multiplexing, unary vs client/server/bidirectional streaming RPCs, metadata auth interceptors, and deadlocks.
- 📂 `04. Websockets and Server Sent Events SSE Real Time Protocols` — Full-duplex bidirectional TCP sockets (WebSockets) vs unidirectional HTTP-based event streaming (SSE), connection management, heartbeats, and reconnects.
- 📂 `05. API Gateway Core Patterns and Reverse Proxies` — Edge routing, authentication offloading, SSL termination, request header enrichment, CORS policy management, and backend protocol translation.
- 📂 `06. Token Bucket and Leaky Bucket Rate Limiting` — Implementing distributed rate limiters with Redis sliding window logs, handling bursting with Token Bucket, HTTP 429 headers, and client backoff.
- 📂 `07. Webhook Architecture, Delivery Engines, and Signatures` — Asynchronous webhook worker queues, exponential retry policies, idempotency delivery guarantees, and HMAC-SHA256 signature verification headers.
- 📂 `08. API Versioning Strategies and Deprecation Lifecycles` — Comparing URI versioning (`/v1/`) vs custom header versioning (`Accept: application/vnd.app.v1+json`), deprecation headers (`Sunset`), and sunset timelines.
- 📂 `09. API Security Architecture Oauth 2.1, Oidc, and Mtls` — Authorization Code Flow with PKCE, JWT token verification via JWKS endpoints, scope-based authorization, and mTLS zero-trust client verification.
- 📂 `10. Consumer Driven Contract Testing with Pact` — Eliminating integration test fragility, defining consumer expectations, validating provider compliance in CI, and preventing breaking API releases.
- 📂 `11. Openapi 3.1 Standards and Interactive Documentation` — Writing precise machine-readable API contracts, generating client SDKs, mock servers (Prism), and interactive developer portals (Scalar, Stoplight).
- 📂 `12. API Mocking, Service Virtualization, and Sandbox APIs` — Mocking third-party APIs (Stripe, Twilio) for local development, creating stateful sandbox environments, and simulating network latency and failures.
- 📂 `13. High Performance Reverse Proxies Nginx and Envoy` — Worker process models, asynchronous event loops, upstream connection pooling, keepalive tuning, dynamic DNS resolution, and Lua/Wasm extensions.
### 9. 📂 [[Resilience, Fault Tolerance & Chaos Engineering|09. Resilience, Fault Tolerance & Chaos Engineering]]
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
### 10. 📂 [[Enterprise Integration Patterns (eip)|10. Enterprise Integration Patterns (EIP)]]
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
### 11. 📂 `11. Multi-Tenant SaaS & Data Isolation Architecture`
- 📂 `01. Multi Tenancy Architectural Models Silo, Pool, Bridge` — Comparing compute and storage isolation models: Dedicated tenant silos for compliance vs shared pooled resources for cost efficiency, and hybrid bridge tiers.
- 📂 `02. Multi Tenant Storage Partitioning and Isolation Models` — Evaluating data isolation trade-offs: Blast radius, connection pool overhead, database migration complexity, and disk utilization across models.
- 📂 `03. Row Level Security RLS for Multi Tenant Data Isolation` — Enforcing database-level tenant isolation using session variables (`SET LOCAL app.current_tenant_id`), preventing cross-tenant data leaks in shared tables.
- 📂 `04. Dynamic Tenant Context Propagation in Middleware` — Extracting tenant IDs from JWT claims/subdomains, propagating tenant context across async thread locals and HTTP/gRPC headers, and logging.
- 📂 `05. Multi Tenant Identity, Scim, and Saml Sso` — Managing tenant-specific identity providers (Okta, Azure AD), automated user provisioning with SCIM 2.0, and tenant-scoped role-based access control.
- 📂 `06. Tenant Aware Caching, Sharding, and Key Namespacing` — Prefixing cache keys with tenant IDs, tenant-aware invalidation, and distributing enterprise tier tenants across dedicated database shard clusters.
- 📂 `07. Multi Tenant Usage Metering, Quotas, and Billing Engines` — Tracking API consumption, storage usage, and compute hours per tenant; enforcing hard/soft quotas, and streaming usage events to Stripe/Metronome.
- 📂 `08. Noisy Neighbor Problem and Fair Resource Queuing` — Preventing high-volume tenants from starving shared cluster resources: Tenant-level rate limiting, separate high-priority queues, and noisy tenant isolation.
- 📂 `09. Tenant Data Export, Backup, and Gdpr Compliance` — Extracting full tenant database dumps on demand, executing cryptographic deletion of tenant records across tables, and audit trail verification.
- 📂 `10. Custom Domain Routing, Ssl Automation, and Whitelabeling` — Routing enterprise custom domains (`app.customer.com`) via Cloudflare for SaaS / Envoy, automated Let's Encrypt SSL issuance, and UI whitelabeling.
### 12. 📂 `12. Architectural Anti-Patterns & Technical Debt Refactoring`
- 📂 `01. The Distributed Monolith Anti Pattern` — Identifying distributed monolith symptoms (simultaneous service deployments, shared codebases, cascading synchronous HTTP calls), and refactoring.
- 📂 `02. God Service and Megaservice Architecture Pitfall` — Decomposing oversized services that hold 80% of business logic into bounded subdomains, resolving high contention, and restoring team velocity.
- 📂 `03. Chatty I O and Microservice Sprawl Pitfall` — Diagnosing high-latency request waterfalls caused by excessive inter-service network hops, aggregating requests with BFF, and batch APIs.
- 📂 `04. Shared Database Anti Pattern and Database Coupling` — Why sharing database tables across multiple microservices causes schema lock-in and silent data corruption, and transitioning to API-driven access.
- 📂 `05. Anemic Domain Model Anti Pattern and Procedural Leaks` — Recognizing anemic domain classes (getters/setters only) with business logic scattered in procedural services, and refactoring to Rich Domain Models.
- 📂 `06. The Golden Hammer and Trend Driven Architecture` — Forcing a single technology (Kafka, Kubernetes, GraphQL) onto every problem regardless of fit, evaluating total cost of ownership, and right-tool-for-the-job.
- 📂 `07. Accidental vs Essential Complexity and Over Engineering` — Distinguishing between core business complexity and self-inflicted architectural bloat, applying YAGNI and KISS, and pruning unused abstractions.
- 📂 `08. Big Ball of Mud and Spaghetti Code Rot` — Why systems degenerate into unstructured spaghetti code, defining modular boundaries, creating dependency matrices, and incremental modularization.
- 📂 `09. Vendor Lock In, Cloud Sprawl, and Leaky Abstractions` — Evaluating proprietary cloud service dependencies against open standards, portability trade-offs, and wrapping vendor SDKs in clean domain ports.
- 📂 `10. Systemic Technical Debt Quantification and Refactoring Roadmaps` — Measuring technical debt interest rate, calculating engineering drag, presenting technical debt refactoring ROI to executives, and phasing migrations.
### 13. 📂 `13. Software Architect Leadership & Governance`
- 📂 `01. Architecture Decision Records ADRs and Decision Governance` — Structuring lightweight immutable decision records (Title, Status, Context, Decision, Consequences), version control integration, and maintaining architectural history.
- 📂 `02. Architecture Review Boards ARB and Lightweight Governance` — Establishing peer-driven architectural review sessions, empowering cross-team architects, preventing design bottlenecks, and coaching senior engineers.
- 📂 `03. Enterprise RFC Request for Comments Workflows` — Institutionalizing open technical RFCs, soliciting cross-functional feedback, driving consensus on controversial decisions, and archiving resolved proposals.
- 📂 `04. Evaluating Architectural Trade Offs Cost, Perf, Complexity` — Why there are no right answers, only trade-offs (Mark Richards); balancing latency, throughput, consistency, operational complexity, and infrastructure cost.
- 📂 `05. Technology Radars, Standardization, and Paved Roads` — Building company-specific tech radars, standardizing language/framework selections, establishing 'paved roads' for developer golden paths, and deprecation.
- 📂 `06. Evolutionary Architecture and Automated Fitness Functions` — Writing automated fitness function tests (ArchUnit) to assert layer boundaries, cyclic dependency prevention, performance budgets, and security gates.
- 📂 `07. Aligning Software Architecture with Business Strategy` — Translating executive business objectives into architectural quality attributes, calculating Total Cost of Ownership (TCO), and architectural ROI.
- 📂 `08. Conway's Law, Team Topologies, and Reverse Conway Maneuver` — Aligning software architecture with team communication structures, optimizing team cognitive load, and reorganizing teams to achieve target architectures.
- 📂 `09. Mentoring, Coaching, and Growing Senior Engineers` — Developing architectural thinking in senior engineers, delegating high-impact design challenges, and fostering a culture of technical excellence.
- 📂 `10. Executive Communication and Board Level Presentations` — Presenting multi-year technical strategies to non-technical executives, structuring one-pagers, framing technical risk in financial terms, and building trust.
### 14. 📂 [[Decentralized, Web3 & Blockchain Architectures|14. Decentralized, Web3 & Blockchain Architectures]]
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

## 🗂️ Contents

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
