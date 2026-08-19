---
title: Microservice Architecture
tags:
  - architecture
  - systems-architecture
  - microservice-architecture-and-service-boundaries
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Microservice Architecture & Service Boundaries

Microservices decomposition and service chassis: Decomposition by business capability, database-per-service, service discovery, distributed tracing (OpenTelemetry), Service Mesh (Envoy/Istio), BFF, canary rollouts, and Strangler Fig monolith migration.

```text
Microservice Architecture & Service Boundaries
│
├── `01. Service Decomposition Strategies and Boundary Definition`
├── `02. Database Per Service vs Shared Database Anti Pattern`
├── [[Microservice Chassis Pattern: Standardized Base Templates and Scaffolding|03. Microservice Chassis and Standardized Service Templates]]
├── [[Dynamic Service Discovery, Registration, and Health Checking (consul, Eureka)|04. Dynamic Service Discovery and Registration Topologies]]
├── [[Distributed Tracing Architecture, W3c Trace Context, and Opentelemetry|05. Distributed Tracing, Span Contexts, and Opentelemetry]]
├── [[Service Mesh Architecture: Envoy Data Plane and Istio Control Plane|06. Service Mesh Architecture Data Plane vs Control Plane]]
├── `07. Backend for Frontend BFF Architecture Pattern`
├── [[Distributed Configuration Management, Centralized Config Servers, and Feature Flags|08. Distributed Configuration Management and Dynamic Reloading]]
├── [[Progressive Delivery: Canary Deployments, Blue Green, and Traffic Splitting|09. Progressive Delivery, Canary Deployments, and Traffic Splitting]]
└── `10. The Strangler Fig Pattern for Monolith Decomposition`
```

---

## 🗂️ Core Knowledge Domains

- 📂 `01. Service Decomposition Strategies and Boundary Definition` — Decomposing monolithic systems by business capability or DDD bounded contexts, evaluating service cohesion, and preventing microservice fragmentation.
- 📂 `02. Database Per Service vs Shared Database Anti Pattern` — Enforcing private data stores per microservice, eliminating cross-service database joins, exposing data via APIs/Events, and managing schema migrations.
- 📂 [[Microservice Chassis Pattern: Standardized Base Templates and Scaffolding|03. Microservice Chassis and Standardized Service Templates]] — Shared baseline cross-cutting concerns: Health checks, structured JSON logging, metrics, distributed tracing headers, security middleware, and configuration.
- 📂 [[Dynamic Service Discovery, Registration, and Health Checking (consul, Eureka)|04. Dynamic Service Discovery and Registration Topologies]] — Client-side vs server-side service discovery, service registry heartbeat checks, dynamic DNS resolution, and load balancing across healthy instances.
- 📂 [[Distributed Tracing Architecture, W3c Trace Context, and Opentelemetry|05. Distributed Tracing, Span Contexts, and Opentelemetry]] — Propagating `traceparent` and `tracestate` HTTP/gRPC headers across microservices, spans, traces, Jaeger/Tempo backends, and latency bottleneck isolation.
- 📂 [[Service Mesh Architecture: Envoy Data Plane and Istio Control Plane|06. Service Mesh Architecture Data Plane vs Control Plane]] — Sidecar proxy architecture, automatic mTLS mutual authentication, intelligent traffic shifting, fault injection, and operational visibility without code changes.
- 📂 `07. Backend for Frontend BFF Architecture Pattern` — Tailoring backend APIs to specific client form factors (iOS, Android, Desktop Web), optimizing over-fetching/under-fetching, and decoupling client releases.
- 📂 [[Distributed Configuration Management, Centralized Config Servers, and Feature Flags|08. Distributed Configuration Management and Dynamic Reloading]] — Centralized configuration stores (Consul, Spring Cloud Config, etcd), dynamic hot-reloading of configs without redeploying, and feature flag management.
- 📂 [[Progressive Delivery: Canary Deployments, Blue Green, and Traffic Splitting|09. Progressive Delivery, Canary Deployments, and Traffic Splitting]] — Automated canary analysis (Kayenta, Argo Rollouts), routing 5% of production traffic to new versions, monitoring error rates, and automated rollback.
- 📂 `10. The Strangler Fig Pattern for Monolith Decomposition` — Intercepting inbound traffic at the edge proxy/API gateway, routing legacy endpoints to the monolith while routing new features to microservices until sunset.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

