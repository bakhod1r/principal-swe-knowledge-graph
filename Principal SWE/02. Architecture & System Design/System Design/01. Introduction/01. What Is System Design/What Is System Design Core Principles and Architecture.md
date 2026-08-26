---
title: "What Is System Design Core Principles and Architecture"
tags:
  - review
  - system-design
  - architecture
  - distributed-systems
  - introduction
  - principal-swe
parent: "[[What Is System Design]]"
---

# System Design: Core Principles and Architecture

**System Design** is the discipline of designing software systems that satisfy functional requirements while remaining reliable, scalable, maintainable, secure, and operationally manageable under real-world constraints.

At a high level:

> **System Design = Requirements + Constraints + Components + Data + Communication + Failure Handling + Operational Guarantees**

A Principal Engineer does not start with _“Should we use Kafka or Redis?”_  
They start with:

> **“What problem are we solving, what guarantees do we need, and what constraints determine the design?”**

---

## 1. The Core Mental Model

Think of a system as:

```text
                 ┌──────────────────────┐
                 │       Clients        │
                 └──────────┬───────────┘
                            │
                         Network
                            │
                 ┌──────────▼───────────┐
                 │     API / Gateway    │
                 └──────────┬───────────┘
                            │
              ┌─────────────┼─────────────┐
              │             │             │
        ┌─────▼─────┐ ┌────▼─────┐ ┌─────▼─────┐
        │  Service  │ │  Service │ │  Service  │
        │     A     │ │     B    │ │     C     │
        └─────┬─────┘ └────┬─────┘ └─────┬─────┘
              │             │             │
              └─────────────┼─────────────┘
                            │
                 ┌──────────▼───────────┐
                 │       Data Layer     │
                 │ DB / Cache / Queue   │
                 └──────────────────────┘
```

But this diagram alone is not system design.

Real system design is about answering:

- Who owns the data?
    
- What happens when the database is unavailable?
    
- What happens when a request is duplicated?
    
- What happens when a service is slow?
    
- What happens when traffic increases 100×?
    
- What happens when one component fails?
    
- What consistency do we need?
    
- What latency do users expect?
    
- How do we detect failure?
    
- How do we recover?
    

---

# 2. First Principle: Requirements Before Architecture

There are two categories of requirements.

### Functional Requirements

**What must the system do?**

Example: URL shortener.

```text
POST /shorten
    long URL
       ↓
    short code

GET /abc123
       ↓
    redirect
```

Functional requirements:

- Create short URL
    
- Redirect short URL
    
- Optionally track clicks
    

### Non-Functional Requirements

**How well must it work?**

Examples:

- 100K requests/sec
    
- P99 latency < 100 ms
    
- 99.99% availability
    
- Data must not be lost
    
- Global users
    
- 10 TB data
    
- Eventual consistency acceptable
    

This distinction is fundamental.

A system can be functionally correct but operationally unacceptable.

```text
Functional correctness:
"Request eventually succeeds."

Operational correctness:
"Request succeeds within 100ms,
99.99% of the time,
without losing data."
```

---

# 3. Constraints Drive Architecture

There is no universally "best architecture."

Architecture is a function of constraints:

```text
Architecture
    =
Requirements
    +
Scale
    +
Consistency
    +
Latency
    +
Availability
    +
Cost
    +
Team
    +
Operational maturity
```

For example:

### Small internal application

```text
Go
 ↓
PostgreSQL
```

May be excellent.

Introducing:

```text
Go
 ↓
Kafka
 ↓
5 microservices
 ↓
Redis
 ↓
Elasticsearch
 ↓
Kubernetes
```

may make the system **worse**, not better.

More components mean:

- more failure modes
    
- more deployments
    
- more monitoring
    
- more network calls
    
- more operational complexity
    
- more debugging complexity
    

### Principal-level principle

> **Do not introduce distributed-systems complexity unless the requirements justify it.**

---

# 4. Architecture Is About Boundaries

One of the most important concepts in system design is **boundary**.

You need to determine:

```text
Who owns what?
```

For example:

```text
User Service
    owns
    └── User data

Order Service
    owns
    └── Order data

Payment Service
    owns
    └── Payment data
```

Bad architecture:

```text
Service A ──┐
Service B ──┼──> Same database tables
Service C ──┘
```

Now ownership becomes ambiguous.

Any service can modify anything.

This creates tight coupling.

Better:

```text
Service A → owns A's data
Service B → owns B's data
Service C → owns C's data
```

Communication happens through explicit APIs/events.

---

# 5. Modularity Before Microservices

A common mistake is thinking:

> "Good architecture = microservices."

Not true.

There is an architectural spectrum:

```text
Monolith
   ↓
Modular Monolith
   ↓
Distributed Services
   ↓
Microservices
   ↓
Event-driven distributed system
```

A **modular monolith** can provide excellent boundaries without immediately paying distributed-system costs.

For example:

```text
                    Application
                        │
       ┌────────────────┼────────────────┐
       │                │                │
   Users Module     Orders Module    Payments Module
       │                │                │
       └────────────────┼────────────────┘
                        │
                    PostgreSQL
```

You get:

- clear ownership
    
- modularity
    
- simpler deployment
    
- local function calls
    
- simpler transactions
    
- easier debugging
    

Then, if one module genuinely needs independent scaling/deployment/ownership, it can potentially become a service later.

---

# 6. Communication: In-Process vs Network

This distinction is extremely important.

### In-process

```go
orderService.CreateOrder(ctx, request)
```

Failure modes are relatively limited.

### Network

```text
Service A
   │
   │ HTTP/gRPC
   ▼
Service B
```

Now you introduce:

- latency
    
- packet loss
    
- timeout
    
- connection failure
    
- DNS failure
    
- TLS failure
    
- retry
    
- duplication
    
- ordering problems
    
- partial failure
    

Therefore:

> **Every network boundary is a failure boundary.**

This is one of the most important distributed-systems mental models.

---

# 7. Data Is Usually the Hardest Part

Applications are often easy.

**Data correctness is hard.**

You need to reason about:

### Ownership

Who can modify the data?

### Consistency

When one component updates data, when do others see it?

### Durability

Will the data survive machine failure?

### Concurrency

What happens when two requests modify the same entity simultaneously?

### Transactions

Which operations must succeed or fail together?

Example:

```text
Transfer $100

Account A: -$100
Account B: +$100
```

You don't want:

```text
A = -100
B = failure
```

A transaction may be required.

---

# 8. Consistency Is a Design Decision

Do not automatically choose strong consistency.

Ask what the business requires.

### Strong consistency

```text
Write
 ↓
Read
 ↓
Immediately sees new value
```

Useful for:

- account balances
    
- inventory reservations
    
- financial transactions
    

### Eventual consistency

```text
Write
 ↓
Event
 ↓
Replication
 ↓
Other systems eventually see update
```

Useful for:

- analytics
    
- counters
    
- search indexes
    
- recommendation systems
    
- notifications
    

The key question is:

> **What inconsistency is acceptable to the business?**

Not:

> "Which database is fastest?"

---

# 9. Scalability

Scalability means the system can handle increased workload without unacceptable degradation.

There are two primary dimensions.

### Vertical Scaling

```text
4 CPU / 16 GB
       ↓
32 CPU / 128 GB
```

Simple, but has physical and economic limits.

### Horizontal Scaling

```text
              ┌── Server 1
Load Balancer ├── Server 2
              ├── Server 3
              └── Server N
```

Now traffic can be distributed.

But horizontal scaling requires thinking about:

- statelessness
    
- shared state
    
- session management
    
- database bottlenecks
    
- coordination
    
- distributed locking
    
- consistency
    

Simply adding servers does not automatically scale the system.

---

# 10. The Real Bottleneck Is Usually Somewhere Else

Suppose:

```text
100 application servers
        │
        ▼
   PostgreSQL
```

Adding more application servers doesn't help if PostgreSQL is already saturated.

Think in terms of the bottleneck:

```text
Client
  ↓
Load Balancer
  ↓
Application
  ↓
Cache
  ↓
Database
  ↓
Disk
```

You need to identify the constrained resource.

Possible bottlenecks:

- CPU
    
- memory
    
- disk I/O
    
- network
    
- database connections
    
- locks
    
- cache
    
- queue
    
- external API
    
- GC
    
- goroutines
    
- file descriptors
    

### Principal-level rule

> **Scale the bottleneck, not the component that is easiest to scale.**

---

# 11. Caching

Cache is useful when:

```text
Read frequency >> Write frequency
```

Typical architecture:

```text
Client
  ↓
API
  ↓
Redis
  │
  ├── HIT → return
  │
  └── MISS
       ↓
    Database
       ↓
    Redis
```

But cache introduces new problems:

- stale data
    
- invalidation
    
- cache stampede
    
- cache penetration
    
- memory limits
    
- eviction
    
- consistency
    

The classic problem:

> **Cache invalidation is a correctness problem, not merely a performance problem.**

---

# 12. Asynchronous Processing

Sometimes the user doesn't need the entire operation to happen synchronously.

Instead of:

```text
HTTP Request
     │
     ▼
Generate report
     │
     ▼
Send email
     │
     ▼
Update analytics
     │
     ▼
Response
```

we can use:

```text
HTTP Request
     │
     ▼
Create Job
     │
     ▼
Queue ───────────────┐
                    │
                    ▼
                 Worker
                    │
             ┌──────┼──────┐
             ▼      ▼      ▼
           Report  Email  Analytics
```

Benefits:

- lower request latency
    
- independent scaling
    
- better resilience
    
- workload smoothing
    

But now you must handle:

- duplicate messages
    
- retries
    
- ordering
    
- dead-letter queues
    
- idempotency
    
- backpressure
    
- poison messages
    

Distributed systems exchange one type of complexity for another.

---

# 13. Failure Is a First-Class Design Concern

A naive architecture assumes:

```text
Request → Success
```

Production architecture assumes:

```text
Request
  │
  ├── success
  ├── timeout
  ├── connection refused
  ├── partial response
  ├── duplicate request
  ├── dependency failure
  ├── overload
  └── process crash
```

For every dependency, ask:

> **What happens if this dependency becomes slow?**

Not just:

> **What happens if it crashes?**

A slow dependency can be more dangerous than a crashed dependency because it can consume resources while requests wait.

---

# 14. Timeouts Are Mandatory at Network Boundaries

Bad:

```go
client.Do(req)
```

without a meaningful timeout strategy.

Production thinking:

```text
Request timeout
       ↓
Dependency timeout
       ↓
Database timeout
```

Timeouts should form a coherent budget.

For example:

```text
Incoming request: 500ms

Service A:
  dependency B: 200ms
  database:      150ms
  remaining:     150ms
```

You don't want each layer independently waiting 5 seconds.

Otherwise:

```text
1 request
   ↓
10 dependencies
   ↓
10 × 5 seconds
   ↓
resource exhaustion
```

---

# 15. Retries Are Dangerous

Retry sounds simple:

```text
Request fails
     ↓
Retry
```

But imagine 10,000 requests hitting an overloaded service.

They all retry.

```text
10,000 requests
      ↓
failure
      ↓
30,000 retries
      ↓
more overload
      ↓
more failures
```

This is a **retry storm**.

Production retry design usually considers:

- bounded retries
    
- exponential backoff
    
- jitter
    
- timeout budget
    
- idempotency
    
- retryable vs non-retryable errors
    

---

# 16. Idempotency

Suppose:

```http
POST /payments
```

The client sends the request.

Server processes payment.

Network fails before response reaches client.

Client retries.

Now:

```text
Payment #1 → $100
Payment #2 → $100
```

Customer was charged twice.

Idempotency solves this class of problem.

For example:

```http
Idempotency-Key: 7f9...
```

Server records:

```text
key → result
```

If the same operation arrives again:

```text
same key
   ↓
return previous result
```

This is critical for:

- payments
    
- order creation
    
- resource provisioning
    
- job submission
    

---

# 17. Backpressure

Suppose producers generate:

```text
100K jobs/sec
```

but consumers process:

```text
10K jobs/sec
```

Then:

```text
Queue:
10K
20K
50K
100K
1M
10M
...
```

Eventually memory/storage fails.

Backpressure means the system has a mechanism to prevent unlimited work accumulation.

Possible strategies:

- bounded queues
    
- rate limiting
    
- load shedding
    
- admission control
    
- producer throttling
    
- consumer scaling
    

### Important mental model

> **A system without backpressure eventually turns overload into failure.**

---

# 18. Availability vs Consistency

Distributed systems often involve trade-offs between:

```text
Consistency
Availability
Partition tolerance
```

Network partitions are unavoidable in distributed systems.

Therefore, when communication breaks, you may have to choose between:

```text
Reject request
```

or:

```text
Accept request with potentially stale/incomplete state
```

There is no magical architecture that eliminates this trade-off.

The correct decision depends on business requirements.

---

# 19. Observability Is Part of Architecture

A production system that cannot explain its failures is incomplete.

Three pillars:

```text
Logs
Metrics
Traces
```

### Metrics

Measure:

```text
Requests/sec
Error rate
P50
P95
P99
CPU
Memory
Queue depth
DB connections
Cache hit ratio
```

### Logs

Useful for events and context:

```text
request_id
user_id
order_id
error
dependency
latency
```

### Traces

Show:

```text
HTTP request
    │
    ├── Service A
    │     ├── Redis
    │     └── PostgreSQL
    │
    └── Service B
          └── Payment API
```

This is extremely valuable for distributed systems.

---

# 20. Security Is a System Property

Security cannot be added only at the API layer.

Think about:

```text
Authentication
Authorization
Input validation
Secrets
Encryption
Network boundaries
Database permissions
Rate limiting
Audit logs
Supply chain
```

For example:

```text
Authenticated ≠ Authorized
```

A user may be authenticated but still must not access another user's resources.

Always ask:

> **Who is allowed to perform this operation on this resource?**

---

# 21. Deployment Is Part of System Design

Architecture does not end with code.

Consider:

```text
Build
 ↓
Test
 ↓
Deploy
 ↓
Observe
 ↓
Rollback
```

Production architecture should support:

- rolling deployments
    
- health checks
    
- graceful shutdown
    
- readiness/liveness
    
- rollback
    
- configuration management
    
- secret management
    
- capacity management
    

A theoretically perfect system that cannot be safely deployed is not a good production system.

---

# 22. The Core Architecture Principles

These are the principles I would internalize rather than memorize technologies.

### 1. Separation of Concerns

Each component should have a clear responsibility.

### 2. High Cohesion

Related behavior belongs together.

### 3. Low Coupling

Changes should not unnecessarily propagate across the system.

### 4. Explicit Ownership

Every important resource/data domain should have a clear owner.

### 5. Dependency Direction

Dependencies should be intentional.

### 6. Minimize Distributed State

Distributed coordination is expensive.

### 7. Design for Failure

Assume dependencies fail.

### 8. Bound Everything

Bound:

- queues
    
- retries
    
- concurrency
    
- memory
    
- connections
    
- request duration
    

### 9. Prefer Explicit Contracts

APIs and events should define clear contracts.

### 10. Measure Before Optimizing

Use:

```text
metrics
profiling
tracing
benchmarks
load tests
EXPLAIN ANALYZE
```

rather than intuition.

### 11. Make Operations Boring

Good production architecture should be:

```text
predictable
observable
recoverable
deployable
```

### 12. Complexity Has a Cost

Every new component adds:

```text
failure modes
operational burden
latency
maintenance
cognitive load
```

---

# 23. A Practical System Design Process

When you receive a system-design problem, use this sequence:

```text
1. Requirements
       ↓
2. Constraints
       ↓
3. Scale estimation
       ↓
4. API design
       ↓
5. Data model
       ↓
6. Component boundaries
       ↓
7. Data flow
       ↓
8. Consistency model
       ↓
9. Caching
       ↓
10. Async processing
       ↓
11. Failure handling
       ↓
12. Security
       ↓
13. Observability
       ↓
14. Scaling
       ↓
15. Bottlenecks
       ↓
16. Disaster recovery
       ↓
17. Trade-offs
```

Do **not** jump directly to:

```text
Kafka + Redis + Kubernetes + microservices
```

before understanding the problem.

---

# 24. The Principal Engineer Question Set

For every major component, ask:

### Problem

> What problem does this component solve?

### Ownership

> Who owns its data and behavior?

### Scale

> What happens at 10× current traffic?

### Failure

> What happens when it is unavailable?

### Latency

> What is the latency budget?

### Consistency

> What consistency guarantees are required?

### Concurrency

> What happens when 1,000 requests modify the same resource?

### Recovery

> How do we recover after failure?

### Observability

> How do we know it is failing?

### Security

> Who can access it and what can they do?

### Migration

> How do we change this architecture without downtime?

### Rollback

> What happens if deployment N+1 is broken?

### Cost

> What operational and infrastructure cost are we accepting?

---

# 25. The Most Important Mental Shift

Junior thinking:

> "Which technology should I use?"

Senior thinking:

> "What architecture satisfies the requirements?"

Staff thinking:

> "What constraints and trade-offs determine the architecture?"

Principal thinking:

> **"What is the simplest system that satisfies the requirements, remains correct under failure, can evolve safely, and can be operated by the organization?"**

That is the core of system design.

---

## A Compact Mental Model

Keep this in your head:

```text
                 SYSTEM DESIGN
                       │
       ┌───────────────┼────────────────┐
       │               │                │
   Requirements      Scale           Constraints
       │               │                │
       └───────────────┼────────────────┘
                       ↓
                  Architecture
                       │
        ┌──────────────┼───────────────┐
        ↓              ↓               ↓
      Data        Components       Communication
        │              │               │
        └──────────────┼───────────────┘
                       ↓
                 Failure Handling
                       │
              ┌────────┼────────┐
              ↓        ↓        ↓
           Timeout   Retry   Backpressure
              │        │        │
              └────────┼────────┘
                       ↓
               Observability
                       ↓
                 Operations
                       ↓
          Reliability + Scalability
```

**The central idea:** System Design is not primarily about drawing boxes or choosing technologies. It is about **making explicit engineering decisions under constraints and understanding what happens when those decisions meet scale, concurrency, and failure.**

For a Go backend engineer, this becomes especially powerful when you connect these principles to **goroutines, bounded concurrency, HTTP/gRPC, PostgreSQL, Redis, queues, Kubernetes, and distributed-system failure modes** rather than studying each technology in isolation.

---

## 🔗 References
- ⬆️ Parent: [[What Is System Design]]
- 📚 Module: `Introduction`

