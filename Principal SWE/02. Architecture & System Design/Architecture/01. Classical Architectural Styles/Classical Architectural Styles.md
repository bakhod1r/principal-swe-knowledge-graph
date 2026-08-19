---
title: Classical Architectural Styles
tags:
  - architecture
  - systems-architecture
  - classical-and-modern-architectural-styles
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Classical & Modern Architectural Styles

Fundamental architectural topologies: Monoliths, Modular Monoliths, Service-Oriented Architecture (SOA), Microservices, Serverless, Event-Driven, Space-Based, Microkernel (Plugin), Peer-to-Peer, and Layered N-Tier architectures.

```text
Classical & Modern Architectural Styles
│
├── [[Monolithic Architecture, Single Deployable Units, and Scaling Trade Offs|01. Monolithic Architecture and Boundaries]]
├── [[Modular Monolith Architecture, Domain Enforced Boundaries, and in Memory Contracts|02. Modular Monolith Architecture]]
├── [[Service Oriented Architecture (soa), Enterprise Service Bus (esb), vs Microservices|03. Service Oriented Architecture Soa vs Microservices]]
├── `04. Microservices Architecture Core Invariants`
├── `05. Serverless and Function As a Service Faas Topologies`
├── `06. Event Driven Architecture Eda Pub Sub and Event Streams`
├── [[Layered N Tier Architecture (presentation, Business, Data) vs Vertical Slicing|07. Layered N Tier Architecture and Vertical Slicing]]
├── [[Microkernel (plugin) Architecture and Extensibility Frameworks|08. Microkernel and Plugin Architecture]]
├── `09. Space Based Architecture and in Memory Data Grids`
├── [[Peer to Peer (p2p) Topologies, Distributed Hash Tables (dht), and Gossip|10. Peer to Peer P2p Distributed Architectures]]
└── [[Architectural Style Selection Framework and Trade Off Matrix|11. Architectural Style Trade Off Matrix and Evaluation]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Monolithic Architecture, Single Deployable Units, and Scaling Trade Offs|01. Monolithic Architecture and Boundaries]] — Shared memory execution, in-process function calls, deployment coupling, database connection pool limits, and vertical scaling boundaries.
- 📂 [[Modular Monolith Architecture, Domain Enforced Boundaries, and in Memory Contracts|02. Modular Monolith Architecture]] — Enforcing strict compile-time domain boundaries, internal event buses, avoiding circular dependencies, and facilitating future service extraction.
- 📂 [[Service Oriented Architecture (soa), Enterprise Service Bus (esb), vs Microservices|03. Service Oriented Architecture Soa vs Microservices]] — Heavyweight centralized ESB routing vs lightweight dumb pipes and smart endpoints in microservices, service contracts (WSDL/JSON), and enterprise governance.
- 📂 `04. Microservices Architecture Core Invariants` — Single Responsibility at service scale, autonomous deployment pipelines, decentralized data management, polyglot persistence, and organizational Conway's Law alignment.
- 📂 `05. Serverless and Function As a Service Faas Topologies` — Ephemeral execution environments, cold start latencies, stateless computation, event-source mappings, fine-grained pay-per-execution economics, and orchestration.
- 📂 `06. Event Driven Architecture Eda Pub Sub and Event Streams` — Asynchronous messaging, decoupling producers from consumers, temporal decoupling, event broker topologies (Kafka, RabbitMQ), and eventual consistency models.
- 📂 [[Layered N Tier Architecture (presentation, Business, Data) vs Vertical Slicing|07. Layered N Tier Architecture and Vertical Slicing]] — Horizontal layer separation, dependency flow rules, open vs closed layers, performance hop penalties, and vertical slice architecture by feature.
- 📂 [[Microkernel (plugin) Architecture and Extensibility Frameworks|08. Microkernel and Plugin Architecture]] — Core system minimal functionality, plugin contracts and interfaces, dynamic plugin loading, extension points, and isolating third-party runtime extensions.
- 📂 `09. Space Based Architecture and in Memory Data Grids` — Eliminating database bottlenecks via in-memory replicated data grids, processing units, virtualized middleware, and extreme linear write scalability.
- 📂 [[Peer to Peer (p2p) Topologies, Distributed Hash Tables (dht), and Gossip|10. Peer to Peer P2p Distributed Architectures]] — Symmetric node roles, BitTorrent protocol, Kademlia distributed hash tables, gossip protocol node discovery, and Byzantine fault tolerance.
- 📂 [[Architectural Style Selection Framework and Trade Off Matrix|11. Architectural Style Trade Off Matrix and Evaluation]] — Evaluating styles against architectural characteristics: Agility, Deployability, Performance, Scalability, Testability, Fault Tolerance, and Operational Cost.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

