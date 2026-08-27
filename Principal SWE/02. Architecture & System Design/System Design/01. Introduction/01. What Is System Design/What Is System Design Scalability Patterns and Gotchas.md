---
title: "What Is System Design Scalability Patterns and Gotchas"
tags:
  - review
  - system-design
  - architecture
  - distributed-systems
  - introduction
  - principal-swe
parent: "[[What Is System Design]]"
---
# System Design: Scalability Patterns

Scalability patterns are **architectural techniques for increasing a system’s capacity while keeping performance, reliability, and operational complexity within acceptable bounds**.

The key idea is:

> **Scalability is not simply “making the server bigger.” It is designing the system so that additional workload can be handled predictably.**

---

## 1. What Problem Does Scalability Solve?

Suppose your system initially handles:

```text
100 requests/sec
```

Then traffic grows:

```text
1K → 10K → 100K requests/sec
```

At some point, a single process or database becomes the bottleneck.

Typical bottlenecks:

- CPU
    
- memory
    
- database connections
    
- database I/O
    
- network bandwidth
    
- disk I/O
    
- locks
    
- connection limits
    
- external APIs
    
- single-instance failure
    
- hot data / hot partitions
    

A scalable architecture provides a way to increase capacity without completely redesigning the system.

---

# 2. Mental Model

Think of scalability as a pipeline:

```text
                 Traffic
                    │
                    ▼
             ┌─────────────┐
             │ Load Balancer│
             └──────┬──────┘
                    │
          ┌─────────┼─────────┐
          ▼         ▼         ▼
       Server 1  Server 2  Server 3
          │         │         │
          └─────────┼─────────┘
                    ▼
                  Cache
                    │
                    ▼
                Database
```

When one component becomes a bottleneck, you introduce a scaling pattern appropriate for that component.

---

# 3. Vertical Scaling

## Definition

Increase the resources of one machine.

```text
4 CPU / 16 GB RAM
        ↓
32 CPU / 128 GB RAM
```

Also called:

**Scale Up**

### Advantages

- Very simple
    
- Minimal architectural changes
    
- Easy deployment
    
- Low operational complexity
    

### Disadvantages

- Hardware limits
    
- Usually expensive at the high end
    
- Still a single failure domain
    
- Eventually hits a ceiling
    

### When to use

Vertical scaling is often the **best first solution**.

Don't introduce distributed systems complexity before you actually need it.

---

# 4. Horizontal Scaling

Increase the number of instances.

```text
          Load Balancer
          /     |     \
         /      |      \
       App1    App2    App3
```

Also called:

**Scale Out**

Instead of:

```text
1 × 32 CPU
```

you might use:

```text
8 × 4 CPU
```

### Requirements

Horizontal scaling becomes much easier when application instances are **stateless**.

Bad:

```text
User → Server A
         │
         └── session stored in RAM
```

Next request:

```text
User → Server B
         │
         └── no session
```

Better:

```text
App instances
      │
      ▼
Shared Session Store / Database
```

---

# 5. Stateless Application Pattern

A horizontally scalable application should ideally not depend on local process state.

Instead of:

```go
var sessions map[string]Session
```

use:

```text
Client
  │
  ▼
Load Balancer
  │
  ├── App 1 ──┐
  ├── App 2 ──┼── Redis / DB
  └── App 3 ──┘
```

The application becomes replaceable:

```text
Instance dies
     ↓
Traffic goes elsewhere
     ↓
No user state is lost
```

This is one of the most important foundations of horizontal scaling.

---

# 6. Load Balancing

A load balancer distributes traffic across instances.

```text
                 ┌── App 1
Client → LB ─────┼── App 2
                 ├── App 3
                 └── App 4
```

Common strategies:

### Round Robin

```text
Request 1 → App1
Request 2 → App2
Request 3 → App3
Request 4 → App1
```

### Least Connections

Send traffic to the instance with the fewest active connections.

### Weighted Routing

```text
App1: weight 70
App2: weight 30
```

Useful when machines have different capacities.

### Consistent Hashing

Useful when requests should consistently map to particular nodes, especially for distributed caches.

---

# 7. Caching

Caching reduces expensive work.

```text
Client
  │
  ▼
Application
  │
  ▼
Cache ── HIT ──→ Response
  │
 MISS
  ▼
Database
```

Without cache:

```text
100K requests
      ↓
100K DB queries
```

With cache:

```text
100K requests
      ↓
95K cache hits
      ↓
5K DB queries
```

The exact improvement depends on workload and cacheability.

### Common cache locations

- Browser
    
- CDN
    
- Reverse proxy
    
- Application memory
    
- Redis
    
- Database buffer/cache
    

### Important trade-off

Caching introduces:

> **staleness + invalidation complexity**

The famous problem is:

> **Cache invalidation is hard.**

You need to reason about:

- TTL
    
- invalidation
    
- cache stampede
    
- stale reads
    
- eviction
    
- cache consistency
    
- hot keys
    

---

# 8. CDN

A CDN moves static or cacheable content closer to users.

```text
                Origin
                  │
                  ▼
             CDN Network
          /       |       \
         ▼        ▼        ▼
      User A   User B   User C
```

Good candidates:

- images
    
- JavaScript
    
- CSS
    
- videos
    
- downloads
    
- static HTML
    
- cacheable API responses
    

Instead of:

```text
User in Tokyo
      ↓
US server
```

you can have:

```text
User in Tokyo
      ↓
Tokyo CDN edge
```

This reduces:

- latency
    
- origin traffic
    
- bandwidth
    
- server load
    

---

# 9. Database Read Replicas

A database often becomes the primary bottleneck.

For read-heavy workloads:

```text
             Primary DB
             /       \
          writes     replication
                       │
                ┌──────┴──────┐
                ▼             ▼
             Replica1       Replica2
```

Writes:

```text
Application → Primary
```

Reads:

```text
Application → Replica
```

### Major limitation

Replication introduces consistency considerations.

You may get:

```text
Write → Primary
Immediately read → Replica
```

and receive old data.

This is:

**Replication lag**

Therefore:

> Read replicas are not merely a performance feature; they change your consistency model.

---

# 10. Database Sharding

When one database cannot handle the workload, partition data across multiple databases.

Example:

```text
User ID

1–1M       → Shard 1
1M–2M      → Shard 2
2M–3M      → Shard 3
```

Or hash-based:

```text
shard = hash(user_id) % N
```

Architecture:

```text
                Application
                     │
               Shard Router
              /      |      \
             ▼       ▼       ▼
          DB-1     DB-2     DB-3
```

### Benefits

- More storage capacity
    
- More write capacity
    
- More read capacity
    
- Reduced contention
    

### Costs

Sharding is expensive operationally.

You introduce:

- shard routing
    
- rebalancing
    
- cross-shard queries
    
- cross-shard transactions
    
- hotspot management
    
- migration complexity
    

**Do not shard merely because it sounds scalable.**

---

# 11. Partitioning

Partitioning divides data logically.

Example:

```text
Orders

2026-01 → Partition 1
2026-02 → Partition 2
2026-03 → Partition 3
...
```

This is particularly useful for time-series or large datasets.

Advantages:

- smaller indexes
    
- partition pruning
    
- easier archival
    
- easier data lifecycle management
    

Important distinction:

```text
Partitioning
    ↓
Usually within one database system

Sharding
    ↓
Distribution across independent database nodes
```

The terms are sometimes used loosely, but the operational implications differ.

---

# 12. Asynchronous Processing

Not every operation needs to happen during the HTTP request.

Instead of:

```text
HTTP Request
    ↓
Validate
    ↓
Generate report
    ↓
Send email
    ↓
Process payment
    ↓
Response
```

use:

```text
HTTP Request
    ↓
Validate
    ↓
Write job
    ↓
Response 202
    │
    ▼
   Queue
    │
    ▼
 Worker
    │
    ├── Generate report
    ├── Send email
    └── Process task
```

This is one of the most powerful scalability patterns.

### Why?

The API handles short-lived work while workers handle expensive work.

You can independently scale:

```text
API:     20 instances
Workers: 100 instances
```

---

# 13. Queue-Based Load Leveling

A queue acts as a buffer between producers and consumers.

```text
Producers
    │
    ▼
┌─────────┐
│  Queue  │
└────┬────┘
     │
 ┌───┼────┐
 ▼   ▼    ▼
W1  W2    W3
```

Suppose traffic suddenly increases:

```text
Normal:
100 jobs/sec

Spike:
10,000 jobs/sec
```

Instead of overwhelming workers, the queue absorbs the burst.

This is:

**Load leveling**

But there's an important trade-off:

> You exchanged immediate processing for latency.

The queue may grow:

```text
Queue depth:
100
500
5K
50K
```

Therefore queue depth and processing latency become important operational metrics.

---

# 14. Backpressure

Backpressure prevents producers from overwhelming consumers.

```text
Producer
   │
   ▼
Queue
   │
   ▼
Consumer
```

If consumers cannot keep up:

```text
producer rate > consumer rate
```

the system must eventually do something:

- block producers
    
- reject requests
    
- drop work
    
- slow down producers
    
- increase consumers
    
- apply rate limits
    

Without backpressure:

```text
Load
 ↓
Queue
 ↓
Memory
 ↓
OOM
 ↓
Crash
```

A scalable system must have a strategy for overload.

---

# 15. Rate Limiting

Rate limiting controls how much traffic a client or system can generate.

Example:

```text
100 requests / minute / user
```

Algorithms include:

- Fixed Window
    
- Sliding Window
    
- Token Bucket
    
- Leaky Bucket
    

Token Bucket is particularly useful because it allows controlled bursts.

```text
Tokens
  ↓
Request consumes token
  ↓
No token → reject / wait
```

Rate limiting protects:

- CPU
    
- databases
    
- downstream services
    
- APIs
    
- expensive operations
    

---

# 16. Connection Pooling

Creating connections repeatedly is expensive.

Instead:

```text
Application
     │
     ▼
Connection Pool
 ┌───┼────┐
 ▼   ▼    ▼
 C1  C2   C3
     │
     ▼
 Database
```

Reuse existing connections.

In Go:

```go
db.SetMaxOpenConns(...)
db.SetMaxIdleConns(...)
db.SetConnMaxLifetime(...)
```

But don't assume:

> "More connections = more performance."

Too many DB connections can create:

```text
More concurrency
    ↓
More contention
    ↓
More context switching
    ↓
DB saturation
    ↓
Higher latency
```

Capacity must be measured.

---

# 17. Read/Write Separation

For workloads where reads dominate:

```text
                Application
                /         \
             Writes       Reads
                │           │
                ▼           ▼
             Primary     Replicas
```

Example workload:

```text
Reads: 99%
Writes: 1%
```

This can significantly increase read capacity.

But again, consistency must be explicitly designed.

---

# 18. CQRS

**Command Query Responsibility Segregation**

Separate models for:

```text
Commands → write model
Queries  → read model
```

Example:

```text
             Events
               │
        ┌──────┴──────┐
        ▼             ▼
   Write Model    Read Model
                     │
                     ▼
                  Queries
```

Useful when read and write requirements differ dramatically.

But CQRS introduces substantial complexity:

- multiple models
    
- synchronization
    
- eventual consistency
    
- event processing
    
- rebuilding projections
    

It should not be used simply because the system is large.

---

# 19. Event-Driven Architecture

Instead of direct synchronous calls:

```text
Order Service
      │
      ├──→ Payment
      ├──→ Email
      └──→ Inventory
```

publish an event:

```text
Order Service
      │
      ▼
 OrderCreated
      │
      ▼
    Broker
   /   |   \
  ▼    ▼    ▼
Payment Email Inventory
```

This reduces temporal coupling.

Services don't necessarily need to be available at exactly the same moment.

But you now need to reason about:

- duplicate events
    
- ordering
    
- retries
    
- idempotency
    
- dead-letter queues
    
- eventual consistency
    
- schema evolution
    

---

# 20. Bulkheads

Isolate resources so failure in one workload doesn't destroy the entire system.

```text
                Application
              /      |      \
             ▼       ▼       ▼
          Pool A   Pool B   Pool C
          Search   Payments Email
```

If Search becomes overloaded:

```text
Search → exhausted
```

Payments should still work.

This is analogous to compartments in a ship:

> One compartment flooding should not sink the entire ship.

---

# 21. Circuit Breaker

Protect a service from repeatedly calling an unhealthy dependency.

```text
Client
  │
  ▼
Circuit Breaker
  │
  ├── CLOSED → request
  │
  ├── OPEN → reject immediately
  │
  └── HALF-OPEN → test recovery
```

States:

```text
CLOSED
   ↓ failures
OPEN
   ↓ timeout
HALF-OPEN
   ↓ success
CLOSED
```

This prevents cascading failures.

---

# 22. Replication

Replication creates multiple copies of data or services.

```text
Primary
 /    \
▼      ▼
R1     R2
```

Goals:

- availability
    
- read scaling
    
- disaster recovery
    
- geographic redundancy
    

But replication creates difficult questions:

```text
Who is authoritative?
How quickly does data replicate?
What happens during network partition?
Can replicas diverge?
How is failover handled?
```

Replication is not free redundancy.

It introduces a **consistency problem**.

---

# 23. Geographic Scaling

For global systems:

```text
                 Global DNS / Anycast
                  /       |       \
                 ▼        ▼        ▼
              US-East   EU-West   Asia
                 │        │        │
               Apps      Apps      Apps
                 │        │        │
                 └────────┼────────┘
                          ▼
                       Data
```

Benefits:

- lower latency
    
- regional fault isolation
    
- disaster resilience
    

Costs:

- distributed data
    
- cross-region replication
    
- consistency problems
    
- higher operational complexity
    
- data residency requirements
    

---

# 24. Autoscaling

Automatically change capacity according to workload.

```text
CPU 30%
   ↓
3 instances

CPU 80%
   ↓
10 instances

CPU 20%
   ↓
4 instances
```

Autoscaling signals can include:

- CPU
    
- memory
    
- request rate
    
- queue depth
    
- latency
    
- custom business metrics
    

A strong production design prefers signals that represent **actual saturation**.

For workers, for example:

```text
Queue depth / processing latency
```

may be much better than CPU utilization.

---

# 25. Caching + Async + Horizontal Scaling

Real systems usually combine patterns.

For example:

```text
                    CDN
                     │
                     ▼
                Load Balancer
                     │
          ┌──────────┼──────────┐
          ▼          ▼          ▼
        App1       App2       App3
          │          │          │
          └──────────┼──────────┘
                     ▼
                   Cache
                     │
                ┌────┴─────┐
                ▼          ▼
             Primary    Queue
                │          │
                ▼          ▼
             Replica     Workers
```

Each pattern solves a different bottleneck.

---

# 26. The Scalability Hierarchy

A useful progression is:

```text
                    ┌─────────────────────┐
                    │ Geographic Scaling  │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │    Sharding         │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │ Async / Messaging   │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │ Caching / Replicas  │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │ Horizontal Scaling  │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │ Vertical Scaling    │
                    └─────────────────────┘
```

This is **not a mandatory order**.

The correct pattern depends on the bottleneck.

---

# 27. A Principal Engineer's Scalability Model

Don't ask:

> "How do I make this system scalable?"

Ask:

> **"What resource is currently limiting throughput?"**

Then identify:

```text
Workload
   ↓
Resource
   ↓
Bottleneck
   ↓
Scaling mechanism
```

For example:

### CPU-bound

```text
CPU saturated
    ↓
Horizontal scaling
```

### DB-read-bound

```text
DB reads saturated
    ↓
Caching / read replicas
```

### DB-write-bound

```text
DB writes saturated
    ↓
Schema/query optimization
    ↓
Batching
    ↓
Partitioning
    ↓
Sharding if necessary
```

### Expensive asynchronous work

```text
HTTP latency
    ↓
Queue
    ↓
Workers
```

### Traffic spikes

```text
Burst
 ↓
Queue / autoscaling / rate limiting
```

### Global latency

```text
Users geographically distributed
        ↓
CDN / regional deployment
```

---

# 28. The Most Important Trade-off

Every scalability technique adds complexity.

For example:

```text
Single server
    ↓
Simple
    ↓
Limited capacity
```

versus:

```text
Distributed system
    ↓
Higher capacity
    ↓
More failure modes
    ↓
More operational complexity
    ↓
More consistency problems
```

So:

> **Scalability is not maximizing theoretical throughput. It is achieving required capacity with acceptable complexity, cost, latency, and reliability.**

---

# 29. Common Anti-Patterns

### Premature microservices

```text
Small application
    ↓
20 microservices
```

You created distributed-system problems before having a scaling problem.

### Cache everything

Caching can create:

- stale data
    
- invalidation bugs
    
- cache stampedes
    
- memory pressure
    

### Add replicas without measuring

Read replicas don't solve:

```text
CPU-bound application
```

### Add more DB connections

More connections can make an overloaded database worse.

### Shard too early

Sharding can turn ordinary queries into distributed queries.

### Async everything

Asynchronous architecture increases:

- latency uncertainty
    
- debugging complexity
    
- operational burden
    
- consistency challenges
    

### Ignore backpressure

Unlimited concurrency eventually becomes system-wide resource exhaustion.

---

# 30. How to Design for Scalability

Use this sequence:

```text
1. Define workload
       ↓
2. Define SLOs
       ↓
3. Estimate capacity
       ↓
4. Identify bottleneck
       ↓
5. Measure current behavior
       ↓
6. Apply the simplest scaling pattern
       ↓
7. Load test
       ↓
8. Observe production behavior
       ↓
9. Repeat
```

For example:

```text
Requirement:
50K requests/sec
P99 < 200 ms

        ↓

Estimate:
Average request cost
DB queries
CPU
memory
network

        ↓

Find bottleneck:
Database reads

        ↓

Solution:
Cache + read replicas

        ↓

Load test

        ↓

Measure:
P50
P95
P99
DB CPU
cache hit rate
error rate

        ↓

Iterate
```

---

# 31. Scalability vs Reliability

These concepts are related but different.

**Scalability:**

> Can the system handle more workload?

**Reliability:**

> Does the system continue behaving correctly despite failures?

A system can be highly scalable but unreliable:

```text
1000 servers
+
massive throughput
+
cascading failure
```

A good production architecture optimizes both.

---

# 32. Key Mental Models

Remember these:

### 1. Bottleneck model

> **Scale the bottleneck, not the entire system blindly.**

### 2. Statelessness

> **Stateless instances are easier to replicate, replace, and autoscale.**

### 3. Queue

> **A queue converts instantaneous load into accumulated work.**

### 4. Cache

> **A cache trades consistency and complexity for lower latency and lower backend load.**

### 5. Replication

> **Replication increases availability/capacity but introduces consistency concerns.**

### 6. Sharding

> **Sharding increases capacity by dividing ownership of data.**

### 7. Backpressure

> **Every system needs a strategy for what happens when demand exceeds capacity.**

### 8. Distributed systems

> **Every new network boundary introduces latency, failure, retries, ordering, and consistency problems.**

---

# 33. Staff/Principal-Level Question

When someone proposes:

> "Let's add Redis."

Don't immediately agree.

Ask:

```text
What is the bottleneck?

Why is Redis the solution?

What data is cacheable?

What is the required consistency?

What is the cache hit rate?

What happens on cache miss?

What happens when Redis is unavailable?

What happens during cache stampede?

How much memory do we need?

How do we invalidate entries?

How do we observe it?

What happens at 10x traffic?
```

That's the difference between:

**"knowing scalability patterns"**

and

**"being able to architect scalable systems."**

---

## Practical Pattern Selection

|Problem|Likely Pattern|
|---|---|
|Single machine CPU limit|Vertical scaling|
|More application traffic|Horizontal scaling|
|Static content load|CDN|
|Expensive repeated reads|Cache|
|Read-heavy DB|Read replicas|
|Huge dataset|Partitioning|
|DB capacity ceiling|Sharding|
|Long-running work|Queue + workers|
|Traffic bursts|Queue / autoscaling|
|Dependency overload|Circuit breaker|
|Resource exhaustion isolation|Bulkhead|
|Excessive client traffic|Rate limiting|
|Global latency|CDN / multi-region|
|Independent processing|Event-driven architecture|
|Different read/write models|CQRS|

### Final principle

> **Scalability is fundamentally a resource-allocation problem under increasing workload.**

A Principal Engineer doesn't start with _"Which technology should we add?"_

They start with:

**Workload → Constraints → Bottleneck → Capacity model → Simplest viable scaling mechanism → Failure modes → Measurement.**


---

## 🔗 References
- ⬆️ Parent: [[What Is System Design]]
- 📚 Module: `Introduction`

