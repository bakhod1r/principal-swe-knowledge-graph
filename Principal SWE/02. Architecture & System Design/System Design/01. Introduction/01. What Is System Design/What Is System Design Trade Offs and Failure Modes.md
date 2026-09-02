---
title: "What Is System Design Trade Offs and Failure Modes"
tags:
  - review
  - system-design
  - architecture
  - distributed-systems
  - introduction
  - principal-swe
parent: "[[What Is System Design]]"
---
# System Design: Trade-offs and Failure Modes

A strong system design is **not** the one that maximizes every desirable property.

It is the one that makes **explicit trade-offs** under real constraints and remains predictable when components fail.

The core mental model is:

> **Every architectural decision buys something by spending something else. Every dependency introduces ways the system can fail.**

---

## 1. What Is a Trade-off?

A **trade-off** is a deliberate choice where improving one system property negatively affects another.

For example:

```text
More consistency
      ↓
More coordination
      ↓
Higher latency / lower availability
```

Or:

```text
More caching
      ↓
Lower database load
      ↓
More stale-data / invalidation complexity
```

There is usually no universally "best" architecture.

The question is:

> **Which property matters most for this particular system?**

---

# 2. The Major System Design Trade-offs

## 2.1 Consistency vs Availability

Consider a replicated database:

```text
          ┌─────────────┐
          │   Primary   │
          └──────┬──────┘
                 │
          replication
        ┌────────┴────────┐
        ↓                 ↓
     Replica A         Replica B
```

Suppose the primary cannot communicate with replicas.

You have a choice.

### Option A — Stop accepting writes

```text
Network partition
       ↓
Cannot guarantee consistency
       ↓
Reject writes
       ↓
Availability ↓
Consistency ↑
```

### Option B — Continue accepting writes

```text
Network partition
       ↓
Continue serving requests
       ↓
Different nodes may diverge
       ↓
Availability ↑
Consistency ↓
```

Neither is automatically correct.

For:

- banking → consistency is usually extremely important
    
- social-media likes → temporary inconsistency may be acceptable
    
- product recommendations → stale data is usually fine
    

---

# 3. Latency vs Consistency

Suppose an API needs data from three services:

```text
Client
  │
  ↓
API
 ├──→ User Service
 ├──→ Order Service
 └──→ Payment Service
```

If the API waits for all three:

```text
Latency ≈ max(
    user_latency,
    order_latency,
    payment_latency
)
```

But you obtain fresher, coordinated information.

Alternatively:

```text
API
 ├──→ Cache
 ├──→ Local data
 └──→ asynchronous updates
```

Latency decreases, but data can become stale.

### Principle

> **Freshness has a latency cost.**

---

# 4. Availability vs Durability

Consider a write request.

### Synchronous replication

```text
Client
  ↓
Primary
  ↓
Replica
  ↓
ACK
```

The system waits for replication before acknowledging.

Advantages:

- stronger durability
    
- lower probability of acknowledged data loss
    

Costs:

- higher latency
    
- replica failure can affect availability
    

### Asynchronous replication

```text
Client
  ↓
Primary
  ↓
ACK
  │
  └────→ Replica
```

Faster, but:

```text
Primary crashes
       ↓
Replication lag
       ↓
Recently acknowledged data may be lost
```

So:

> **Durability is not free.**

---

# 5. Performance vs Simplicity

Suppose a database query takes:

```text
500 ms
```

You could introduce:

```text
Redis
Read replicas
CQRS
Materialized views
Denormalization
Sharding
```

and potentially reduce latency to:

```text
20 ms
```

But now you have:

```text
Application
 ├── PostgreSQL
 ├── Redis
 ├── Replica
 ├── Queue
 └── Materialized View
```

Your system became much harder to reason about.

Therefore:

> **Do not introduce distributed complexity until measurement demonstrates that it is necessary.**

A 500 ms query may simply need:

```sql
CREATE INDEX ...
```

---

# 6. Normalization vs Denormalization

Normalized model:

```text
Users
Orders
OrderItems
Products
```

Advantages:

- less duplication
    
- stronger consistency
    
- easier updates
    

Costs:

- more joins
    
- potentially slower reads
    

Denormalized model:

```text
Order {
    order_id
    customer_name
    product_name
    product_price
}
```

Advantages:

- fast reads
    
- fewer joins
    

Costs:

- duplicated data
    
- synchronization complexity
    
- stale values
    

Mental model:

```text
Normalization
    ↓
Write simplicity / consistency

Denormalization
    ↓
Read performance / operational complexity
```

---

# 7. Strong Guarantees vs Operational Complexity

Compare:

```text
Single PostgreSQL
```

with:

```text
PostgreSQL
    +
Redis
    +
Kafka
    +
Read replicas
    +
CDC
    +
Distributed workers
```

The second architecture may support substantially higher scale.

But every additional component creates:

- another failure boundary
    
- another deployment
    
- another monitoring surface
    
- another operational dependency
    
- another recovery procedure
    

A useful approximation is:

```text
System Complexity
    ≈
    Number of components
    ×
    Number of interactions
    ×
    Number of failure states
```

This isn't a mathematical law, but it is an excellent engineering intuition.

---

# 8. Synchronous vs Asynchronous Processing

### Synchronous

```text
Client
  ↓
API
  ↓
Payment
  ↓
Email
  ↓
Response
```

Simple mental model.

But one slow dependency can block the entire request.

### Asynchronous

```text
Client
  ↓
API
  ↓
Queue
  ↓
Worker
 ├── Payment
 └── Email
```

Advantages:

- decoupling
    
- buffering
    
- retry capability
    
- better resilience to traffic spikes
    

But now you need to reason about:

- duplicate messages
    
- ordering
    
- retries
    
- dead-letter queues
    
- consumer lag
    
- idempotency
    
- eventual consistency
    

So:

> **Async architecture trades temporal coupling for coordination complexity.**

---

# 9. What Is a Failure Mode?

A **failure mode** describes how a system can behave incorrectly or become unavailable when something goes wrong.

Don't ask only:

> "What happens when the server crashes?"

Ask:

> **"What happens when each dependency partially fails?"**

---

# 10. Common Failure Modes

## 10.1 Timeout

```text
Service A
   │
   └── request → Service B
                    │
                    X
                 hangs
```

Without timeout:

```text
goroutines/connections
        ↓
wait
        ↓
wait
        ↓
resource exhaustion
```

Therefore every network dependency should have an explicit timeout appropriate to its operation.

---

# 11. Retry Storm

A common mistake:

```text
Service A
   ↓
Service B
   X
```

A retries:

```text
A → B
A → B
A → B
A → B
```

But imagine 10,000 clients doing this simultaneously.

```text
Failure
  ↓
Retries
  ↓
More traffic
  ↓
B becomes even less healthy
  ↓
More failures
  ↓
More retries
```

This is a **retry storm**.

Production retry design generally requires:

- bounded retries
    
- exponential backoff
    
- jitter
    
- deadlines
    
- idempotency
    
- retry classification
    

And importantly:

> **Not every error should be retried.**

---

# 12. Cascading Failure

Example:

```text
                ┌→ Service B
Client → API ───┼→ Service C
                └→ Service D
```

Suppose B becomes slow.

```text
B slow
 ↓
API requests remain active longer
 ↓
Connection pools fill
 ↓
API latency increases
 ↓
Clients retry
 ↓
Traffic increases
 ↓
API collapses
```

One component caused failure elsewhere.

This is a **cascading failure**.

Useful containment mechanisms include:

```text
Timeouts
Circuit breakers
Bulkheads
Rate limits
Bounded concurrency
Backpressure
Load shedding
```

---

# 13. Resource Exhaustion

Every system has finite resources:

```text
CPU
Memory
File descriptors
Connections
Threads
Goroutines
DB connections
Queue capacity
Disk
Network bandwidth
```

A classic failure:

```text
Traffic ↑
   ↓
Concurrency ↑
   ↓
Memory ↑
   ↓
GC pressure ↑
   ↓
Latency ↑
   ↓
Timeouts ↑
   ↓
Retries ↑
   ↓
Traffic ↑
```

This creates a positive feedback loop.

Therefore:

> **Bound concurrency instead of allowing unlimited work.**

---

# 14. Queue Failure Modes

Queues solve some problems but introduce others.

```text
Producer
   ↓
Queue
   ↓
Consumer
```

What if consumers slow down?

```text
Consumer throughput ↓
        ↓
Queue depth ↑
        ↓
Memory/disk ↑
        ↓
Latency ↑
```

Eventually:

```text
Queue full
   ↓
Producer blocked/rejected
```

This is where **backpressure** becomes important.

Monitor:

```text
queue depth
consumer lag
processing latency
oldest message age
dead-letter count
```

---

# 15. Duplicate Processing

Distributed systems commonly provide behavior closer to:

```text
at-least-once delivery
```

rather than true exactly-once execution.

Example:

```text
Worker
  ↓
Process payment
  ↓
Payment succeeds
  ↓
Worker crashes before ACK
  ↓
Message delivered again
  ↓
Payment processed again
```

Therefore:

> **Assume messages and requests can be duplicated.**

Design operations to be **idempotent** where possible.

For example:

```text
POST /payments
Idempotency-Key: abc123
```

The server can ensure:

```text
abc123 → same logical operation
```

instead of charging twice.

---

# 16. Partial Failure

One of the most important distributed-systems concepts.

A machine can be:

```text
alive
```

but the network path to it can be broken.

A service can be:

```text
healthy
```

but its database can be unavailable.

A request can:

```text
reach server
→ execute successfully
→ response gets lost
```

The client then sees:

```text
timeout
```

while the server sees:

```text
success
```

This produces an important rule:

> **Timeout does not mean the operation did not happen.**

This is why retries + idempotency are tightly coupled.

---

# 17. Split Brain

Suppose two nodes both believe they are primary:

```text
       Network Partition

   Node A       Node B
  "I'm primary" "I'm primary"
```

Both accept writes.

Now:

```text
A: balance = 100
B: balance = 50
```

When connectivity returns, reconciliation becomes difficult.

Preventing split-brain may require:

- quorum
    
- leader election
    
- fencing
    
- consensus protocols
    
- leases with careful semantics
    

---

# 18. Single Point of Failure

Architecture:

```text
          ┌→ Server A
Client → LB
          └→ Server B

          ↓
       Database
```

If there is only one database:

```text
Database crashes
      ↓
Entire application unavailable
```

That database is an **SPOF**.

But eliminating every SPOF is not always economically justified.

Ask:

```text
What is the required availability?
What is the failure probability?
What downtime is acceptable?
What is the recovery time?
What does redundancy cost?
```

---

# 19. The Important Reliability Model

A useful production framework is:

```text
Prevention
    ↓
Detection
    ↓
Containment
    ↓
Recovery
    ↓
Graceful Degradation
```

For example:

### Prevention

```text
validation
timeouts
capacity limits
```

### Detection

```text
metrics
logs
traces
health checks
alerts
```

### Containment

```text
rate limiting
circuit breakers
bulkheads
load shedding
```

### Recovery

```text
retry
failover
rollback
restore
replay
```

### Graceful degradation

```text
Recommendations unavailable
        ↓
Still show product page
```

The goal isn't:

> "Nothing ever fails."

The goal is:

> **"Failure remains bounded, observable, and recoverable."**

---

# 20. Trade-off Matrix

When evaluating architecture, explicitly score the important dimensions:

|Decision|Benefit|Cost|
|---|---|---|
|Cache|Lower latency/load|Staleness/invalidation|
|Replication|Availability/read scale|Consistency/ops|
|Sharding|Horizontal scale|Complexity|
|Async queue|Decoupling/buffering|Eventual consistency|
|Denormalization|Fast reads|Duplicate data|
|Strong consistency|Correctness|Latency/availability|
|More retries|Transient-failure recovery|Retry storms|
|Microservices|Team/service isolation|Network/ops complexity|
|Multi-region|Disaster resilience|Cost/consistency complexity|
|Larger cache|More hit rate|Memory/cost/staleness|

This table is more valuable than saying:

> "Microservices are scalable."

---

# 21. A Principal-Level Decision Process

When making an architectural decision, use this sequence:

```text
1. Requirements
       ↓
2. Constraints
       ↓
3. Failure assumptions
       ↓
4. Consistency requirements
       ↓
5. Availability requirements
       ↓
6. Latency requirements
       ↓
7. Expected scale
       ↓
8. Operational capabilities
       ↓
9. Cost
       ↓
10. Choose simplest viable architecture
```

For every major decision ask:

### Correctness

> What invariant must never be violated?

### Performance

> What is the actual bottleneck?

### Reliability

> What happens when this dependency fails?

### Concurrency

> What happens when 10,000 requests execute simultaneously?

### Distributed systems

> What happens when the network partitions?

### Operations

> How do we detect the problem?

### Recovery

> How do we restore service?

### Migration

> How do we move from the old design to this one?

### Rollback

> What happens if deployment fails halfway?

---

# 22. The Most Important Mental Model

Don't design only this:

```text
Happy Path

Client
  ↓
API
  ↓
DB
  ↓
Response
```

Design this:

```text
                    ┌─ timeout
                    ├─ retry
                    ├─ duplicate
                    ├─ stale data
                    ├─ overload
                    ├─ network partition
                    ├─ dependency failure
                    ├─ process crash
                    ├─ disk full
                    ├─ DB failure
                    └─ deployment failure

Client → API → DB
```

A system design becomes **production-grade** when failure is part of the architecture rather than an afterthought.

## Principal Engineer takeaway

The key progression is:

```text
Junior:
"How does the system work?"

Senior:
"What can go wrong?"

Staff:
"How does failure propagate?"

Principal:
"How do we bound the blast radius,
detect failure quickly,
preserve critical invariants,
recover safely,
and justify the trade-offs economically?"
```

The deepest principle is:

> **Architecture is the management of trade-offs under failure and constraints—not the selection of technologies.**

## 🔗 References
- ⬆️ Parent: [[What Is System Design]]
- 📚 Module: `Introduction`

