---
title: Resilience, Fault Tolerance & Chaos Engineering
tags:
  - architecture
  - systems-architecture
  - resilience,-fault-tolerance-and-chaos-engineering
  - principal-swe
parent: "[[Architecture]]"
---

# 🏛️ Resilience, Fault Tolerance & Chaos Engineering

High-availability and fault-tolerance patterns: Circuit Breakers (Resilience4j), Bulkheads, Exponential Backoff with Jitter, Graceful Degradation, Timeout & Deadline propagation, Dead Letter Queues (DLQ), and Chaos Engineering (Chaos Monkey).

```text
Resilience, Fault Tolerance & Chaos Engineering
│
├── [[Circuit Breaker Pattern: Closed, Open, Half Open State Transitions|01. Circuit Breaker Pattern and State Transitions]]
├── [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation|02. Bulkhead Pattern and Resource Isolation]]
├── [[Exponential Backoff and Jitter: Mitigating the Thundering Herd|03. Exponential Backoff with Full Jitter]]
├── [[Graceful Degradation, Fallback Handlers, and Shedding Load|04. Graceful Degradation and Fallback Strategies]]
├── [[Distributed Timeout Budgets and gRPC Deadline Propagation|05. Distributed Timeouts and Deadline Propagation]]
├── [[Dead Letter Queues (dlq), Poison Pill Handling, and Quarantine Queues|06. Dead Letter Queues DLQ and Poison Pill Handling]]
├── [[Adaptive Load Shedding and Concurrency Limiters (netflix Concurrency Limits)|07. Rate Limiting, Load Shedding, and Concurrency Limits]]
├── [[Chaos Engineering Principles, Fault Injection, and Chaos Mesh|08. Chaos Engineering Principles and Fault Injection]]
├── [[High Availability (ha) Multi Region Topologies and Automated Failover|09. High Availability Topologies and Cross Region Failover]]
└── [[Deep Health Checking Probes (liveness, Readiness, Startup) and Synthetic Canaries|10. Health Check Probes and Synthetic Monitoring]]
```

---

## 🗂️ Core Knowledge Domains

- 📂 [[Circuit Breaker Pattern: Closed, Open, Half Open State Transitions|01. Circuit Breaker Pattern and State Transitions]] — Preventing cascading failures during downstream outages, sliding window failure rate thresholds, slow call thresholds, and automatic recovery probing.
- 📂 [[Bulkhead Pattern: Thread Pool and Semaphore Resource Isolation|02. Bulkhead Pattern and Resource Isolation]] — Isolating thread pools and connection pools per downstream dependency so a slow external API does not exhaust resources for the entire application.
- 📂 [[Exponential Backoff and Jitter: Mitigating the Thundering Herd|03. Exponential Backoff with Full Jitter]] — Why naive retries cause self-inflicted DDoS attacks, calculating exponential backoff with full jitter (AWS algorithm), and decorrelated jitter.
- 📂 [[Graceful Degradation, Fallback Handlers, and Shedding Load|04. Graceful Degradation and Fallback Strategies]] — Returning cached stale data, degraded static responses, disabling non-essential UI features (recommendations/reviews), and prioritizing core user flows.
- 📂 [[Distributed Timeout Budgets and gRPC Deadline Propagation|05. Distributed Timeouts and Deadline Propagation]] — Setting strict client-side timeouts, propagating remaining deadline budgets across downstream microservice chains, and cancelling dead work early.
- 📂 [[Dead Letter Queues (dlq), Poison Pill Handling, and Quarantine Queues|06. Dead Letter Queues DLQ and Poison Pill Handling]] — Isolating unparseable or crash-inducing messages without blocking consumer pipelines, maximum retry thresholds, and manual replay tooling.
- 📂 [[Adaptive Load Shedding and Concurrency Limiters (netflix Concurrency Limits)|07. Rate Limiting, Load Shedding, and Concurrency Limits]] — Little's Law, monitoring queue latency and CPU saturation, dynamically shedding incoming low-priority requests to protect system availability.
- 📂 [[Chaos Engineering Principles, Fault Injection, and Chaos Mesh|08. Chaos Engineering Principles and Fault Injection]] — Formulating steady-state hypotheses, introducing real-world turbulence (network latency, killed nodes, packet drop), and validating resilience.
- 📂 [[High Availability (ha) Multi Region Topologies and Automated Failover|09. High Availability Topologies and Cross Region Failover]] — Active-Passive vs Active-Active multi-region deployments, DNS failover with Route 53, data replication lag, and Recovery Time/Point Objectives (RTO/RPO).
- 📂 [[Deep Health Checking Probes (liveness, Readiness, Startup) and Synthetic Canaries|10. Health Check Probes and Synthetic Monitoring]] — Designing non-cascading health check endpoints, verifying critical downstream dependencies safely, and synthetic browser-driven user journey probes.

---

## 🔗 References
- ⬆️ Parent: [[Architecture]]

