---
title: "How to Approach Scalability Patterns and Gotchas"
tags:
  - review
  - system-design
  - architecture
  - distributed-systems
  - introduction
  - principal-swe
parent: "[[How to Approach]]"
---
# How to Approach Scalability Patterns and Common Pitfalls

Scalability is not “add more servers.” It is the discipline of designing a system so that **increasing load does not cause unacceptable degradation in latency, throughput, availability, or cost**.

The key Staff+/Principal-level shift is:

> **Do not start with scalability patterns. Start by identifying the bottleneck, constraint, and scaling dimension.**

---

## 1. Start With the Real Problem

Before choosing a pattern, answer:

```text
What is growing?
    ↓
Requests/sec?
Concurrent users?
Data volume?
Payload size?
Background jobs?
Connections?
Geographic distribution?
    ↓
What resource saturates first?
    ↓
CPU / Memory / Disk / Network / DB / Lock / External dependency
    ↓
What constraint must remain bounded?
    ↓
Latency / throughput / availability / cost
```

For example:

```text
10k req/s
     ↓
API servers: 30% CPU
     ↓
PostgreSQL: 95% CPU
     ↓
Query: SELECT ... WHERE user_id = ?
     ↓
Missing/inefficient index
```

The scalability problem is **not** “we need microservices.”

It is:

> The database query path has become the bottleneck.

That distinction is fundamental.

---

# 2. Mental Model: Scaling Dimensions

Think of scalability across several independent dimensions.

|Dimension|Example|
|---|---|
|Compute|CPU-intensive API|
|Memory|Large working set|
|Storage|Billions of rows|
|Network|Large payloads|
|Concurrency|Millions of simultaneous connections|
|Database|High query/write volume|
|Background work|Millions of jobs|
|Geographic|Users across continents|
|Dependency|Third-party API limits|

A system can scale well in one dimension and badly in another.

For example:

```text
API:
Horizontal scaling     ✅

Database:
Horizontal scaling     ❌

External API:
Rate limited           ❌
```

Adding 100 API instances may therefore make the system **worse** because they generate even more database traffic.

---

# 3. First Principle: Find the Bottleneck

A useful production loop is:

```text
Measure
  ↓
Identify bottleneck
  ↓
Understand why
  ↓
Change architecture
  ↓
Measure again
```

Do not:

```text
"Traffic increased"
      ↓
"Let's add Redis"
      ↓
"Let's add Kafka"
      ↓
"Let's use Kubernetes"
      ↓
"Let's split into microservices"
```

Those are technologies, not diagnoses.

---

# 4. Scalability Patterns

There is no universal hierarchy, but these patterns commonly appear.

## Pattern 1 — Vertical Scaling

Increase resources on one machine.

```text
Before:

     ┌────────────┐
     │  Server    │
     │  4 CPU     │
     │  8 GB RAM  │
     └────────────┘

After:

     ┌────────────┐
     │  Server    │
     │ 32 CPU     │
     │128 GB RAM  │
     └────────────┘
```

### Advantages

- Simple
    
- Low operational complexity
    
- No distributed coordination
    
- Often surprisingly effective
    

### Limitations

Eventually:

```text
hardware ceiling
       ↓
cannot scale further
```

Also creates a large failure domain.

### Principal-level insight

**Vertical scaling is often the best first move.**

Distributed systems introduce enormous complexity. Don't pay that cost before you need it.

---

# 5. Pattern 2 — Horizontal Scaling

Run multiple instances.

```text
                 ┌─────────┐
                 │   LB    │
                 └────┬────┘
          ┌───────────┼───────────┐
          ↓           ↓           ↓
       Server A    Server B    Server C
```

Requests are distributed across instances.

### Works well when

The application is approximately:

```text
request + input → response
```

and instances don't require local shared state.

This leads to:

> **Stateless application design**

---

# 6. Statelessness

Suppose authentication state lives in memory:

```text
Request
   ↓
Server A
   ↓
session[user123] = ...
```

Next request:

```text
Request
   ↓
Server B
   ↓
session[user123] missing
```

You now need either:

```text
sticky sessions
```

or shared state:

```text
              Redis
             ↗     ↖
        Server A   Server B
```

Prefer externalizing shared state when horizontal scalability matters.

But don't blindly put everything into Redis.

The real question is:

> **What state must be shared, and what consistency guarantees does it require?**

---

# 7. Pattern 3 — Load Balancing

Load balancing distributes traffic.

```text
                 Internet
                    │
                    ↓
              ┌──────────┐
              │ Load Bal.│
              └────┬─────┘
          ┌─────────┼─────────┐
          ↓         ↓         ↓
         A          B         C
```

Common strategies:

- Round robin
    
- Least connections
    
- Weighted routing
    
- Consistent hashing
    
- Latency-aware routing
    

But load balancing doesn't create capacity.

If the database is saturated:

```text
        Load Balancer
       /      |      \
      A       B       C
       \      |      /
        \     |     /
          Database
             🔥
```

You simply distribute requests faster toward the same bottleneck.

---

# 8. Pattern 4 — Caching

Caching reduces expensive work.

```text
Request
   ↓
Cache
 ┌──┴──┐
Hit   Miss
 ↓      ↓
Data   DB
```

The important equation is:

```text
DB load ≈ Request load × (1 - cache hit ratio)
```

For example:

```text
100k req/s
95% hit ratio

DB:
100k × 5%
= 5k req/s
```

That is a huge reduction.

### But caching introduces new problems

- Stale data
    
- Cache invalidation
    
- Memory pressure
    
- Hot keys
    
- Cache stampede
    
- Cold-start load
    
- Inconsistent views
    

The famous difficulty isn't merely:

> “How do we cache?”

It is:

> **What is the correctness model when cached data is stale?**

---

# 9. Cache Stampede

Suppose:

```text
TTL = 60 sec
```

At expiration:

```text
10,000 requests
       ↓
cache miss
       ↓
10,000 DB queries
       ↓
🔥 DB
```

This is a cache stampede.

Typical mitigations:

- Request coalescing
    
- Single-flight
    
- Jittered TTL
    
- Background refresh
    
- Stale-while-revalidate
    
- Distributed locking where justified
    

In Go, `singleflight` is particularly useful for collapsing concurrent duplicate work within a process.

---

# 10. Pattern 5 — Read Replicas

If reads dominate:

```text
                ┌──────────┐
                │ Primary  │
                └────┬─────┘
                     │ replication
            ┌────────┼────────┐
            ↓        ↓        ↓
          Replica  Replica  Replica
```

Writes:

```text
Primary
```

Reads:

```text
Replicas
```

This can increase read capacity.

But replication creates a critical trade-off:

```text
write → primary
           ↓
       replication
           ↓
       replica
```

There can be replication lag.

Therefore:

> **Read replicas usually trade stronger read-after-write semantics for scalability.**

You need to decide whether stale reads are acceptable.

---

# 11. Pattern 6 — Database Sharding

When one database cannot handle the dataset/workload:

```text
                 Application
                      │
               shard routing
          ┌───────────┼───────────┐
          ↓           ↓           ↓
       Shard 1     Shard 2     Shard 3
       users 1M    users 1M    users 1M
```

For example:

```text
hash(user_id) % N
```

### Benefits

- More storage capacity
    
- More write capacity
    
- Smaller indexes
    
- Parallelism
    

### Costs

Now you have distributed-data problems:

- Cross-shard queries
    
- Cross-shard transactions
    
- Rebalancing
    
- Hot shards
    
- Operational complexity
    
- Backup/recovery complexity
    

Sharding should therefore usually be considered **after simpler database scaling techniques**.

---

# 12. Pattern 7 — Partitioning

Don't confuse partitioning with sharding.

Partitioning can happen within one database:

```text
Orders

2026-01
2026-02
2026-03
2026-04
...
```

For time-series data, this can improve:

- Query pruning
    
- Maintenance
    
- Retention
    
- Index size
    
- Data lifecycle management
    

Sharding usually means distributing data across independent database nodes.

---

# 13. Pattern 8 — Asynchronous Processing

Move expensive work out of the synchronous request path.

Naive:

```text
HTTP request
    ↓
Generate PDF
    ↓
Send email
    ↓
Resize images
    ↓
Update analytics
    ↓
Response
```

Better:

```text
HTTP request
    ↓
Persist command
    ↓
Queue
    ↓
202 Accepted
```

Workers:

```text
Queue
 ↓
Worker A
Worker B
Worker C
```

This separates:

```text
request latency
```

from:

```text
background processing
```

But now you must reason about:

- Delivery semantics
    
- Retries
    
- Idempotency
    
- Ordering
    
- Dead-letter queues
    
- Backpressure
    
- Poison messages
    
- Consumer lag
    

---

# 14. Pattern 9 — Backpressure

This is one of the most important scalability concepts.

Suppose:

```text
Producer = 100k jobs/s
Consumer = 20k jobs/s
```

Then:

```text
queue growth = 80k jobs/s
```

Eventually:

```text
memory → full
disk → full
latency → huge
system → failure
```

Backpressure means:

> **When downstream capacity is exhausted, upstream work must slow down, reject work, or shed load.**

Examples:

```text
rate limiting
bounded queues
semaphores
connection limits
worker pools
load shedding
```

A system without backpressure often converts overload into a total outage.

---

# 15. Pattern 10 — Rate Limiting

Protect scarce resources.

```text
Client
  ↓
Rate Limiter
  ↓
Application
```

Common algorithms:

- Token bucket
    
- Leaky bucket
    
- Fixed window
    
- Sliding window
    

Example:

```text
100 requests/sec/user
```

This prevents one client from consuming the entire system capacity.

---

# 16. Pattern 11 — CDN / Edge Caching

For static or cacheable geographically distributed content:

```text
                Users
              /   |   \
             ↓    ↓    ↓
           Edge Edge Edge
             \    |    /
              Origin
```

Benefits:

- Lower latency
    
- Reduced origin traffic
    
- Reduced bandwidth cost
    
- Better global scalability
    

Especially useful for:

- Images
    
- JS/CSS
    
- Videos
    
- Downloads
    
- Cacheable API responses
    

---

# 17. Pattern 12 — Connection Pooling

Creating connections repeatedly is expensive.

Instead:

```text
Application
     │
     ↓
Connection Pool
 ┌───┼────┬────┐
 C1  C2   C3   C4
     │
     ↓
 Database
```

But connection pools have a critical failure mode.

Suppose:

```text
100 application instances
×
100 DB connections
=
10,000 connections
```

Your database may support only:

```text
2,000
```

Horizontal scaling can therefore **destroy database scalability**.

This is a classic distributed-system trap.

---

# 18. Common Pitfall: Scaling the Wrong Layer

One of the most common mistakes:

```text
API CPU = 20%
DB CPU = 100%
```

Engineer:

> “Let's add API instances.”

Result:

```text
More API instances
      ↓
More DB requests
      ↓
DB overload
      ↓
More latency
      ↓
More retries
      ↓
Even more DB requests
      ↓
🔥
```

This is a positive feedback loop.

---

# 19. Common Pitfall: Retry Storm

Imagine a dependency becomes slow.

Clients timeout:

```text
Request
   ↓
Dependency
   ↓
timeout
```

Then retry:

```text
Request
   ↓
retry
   ↓
retry
   ↓
retry
```

Instead of:

```text
1000 requests/s
```

the dependency might receive:

```text
3000 requests/s
```

which makes it slower, causing more retries.

This is a **retry storm**.

Production systems need:

```text
timeouts
+
bounded retries
+
exponential backoff
+
jitter
+
idempotency
+
load shedding
```

---

# 20. Common Pitfall: Unbounded Concurrency

Bad:

```go
for _, job := range jobs {
    go process(job)
}
```

If:

```text
jobs = 10,000,000
```

you've created a resource-management problem.

Better:

```text
          jobs
           ↓
      bounded queue
           ↓
    ┌──────┼──────┐
    ↓      ↓      ↓
 Worker Worker Worker
```

Concurrency should usually be **bounded by the resource being protected**.

---

# 21. Common Pitfall: Hot Keys

Suppose:

```text
celebrity_id = 123
```

receives 50% of all traffic.

Consistent hashing may place it on one node:

```text
Shard A ← 50% traffic 🔥
Shard B ← 10%
Shard C ← 10%
Shard D ← 10%
```

The cluster has spare capacity, but one partition is overloaded.

Solutions may include:

- Key salting
    
- Replication
    
- Local caching
    
- Request coalescing
    
- Better partitioning strategy
    

---

# 22. Common Pitfall: Distributed Monolith

You split:

```text
User Service
Order Service
Payment Service
Notification Service
```

but every request requires:

```text
User
 ↓
Order
 ↓
Payment
 ↓
Notification
 ↓
Inventory
```

Now you have:

```text
distributed deployment
+
distributed failures
+
network latency
+
distributed tracing
+
deployment complexity
```

without gaining meaningful independence.

That's a **distributed monolith**.

Microservices should create useful boundaries, not merely network boundaries.

---

# 23. Common Pitfall: Premature Microservices

Start:

```text
Monolith
```

because it is simple.

Extract services when there is a real boundary such as:

```text
different scaling characteristics
different ownership
different availability requirements
different deployment cadence
different data ownership
strong domain boundary
```

Not because:

> “Netflix uses microservices.”

---

# 24. Common Pitfall: Caching Without a Correctness Model

Bad reasoning:

> “Database is slow, let's cache it.”

You need to answer:

```text
How stale can data be?

0 seconds?
1 second?
1 minute?
1 hour?

What invalidates the cache?

What happens on cache failure?

Can stale data cause financial/security problems?

What happens during cache warm-up?
```

For some data:

```text
stale = acceptable
```

For other data:

```text
stale = incorrect
```

Architecture depends on that distinction.

---

# 25. Common Pitfall: Synchronous Everything

Imagine:

```text
API
 ↓
Service A
 ↓
Service B
 ↓
Service C
 ↓
Service D
 ↓
Database
```

Latency roughly becomes:

```text
L ≈ L_A + L_B + L_C + L_D + L_DB
```

And availability becomes constrained by dependencies.

If each dependency has:

```text
99.9% availability
```

a chain of four independent mandatory dependencies has approximately:

```text
0.999⁴ ≈ 99.60%
```

before considering the application's own failure modes.

This is why dependency topology matters.

---

# 26. Common Pitfall: No Capacity Model

You should estimate capacity before architecture changes.

Suppose:

```text
Peak traffic = 20k req/s
Average = 5k req/s

One instance:
500 req/s

Required:
20,000 / 500 = 40 instances
```

Then add headroom:

```text
40 × 1.5 = 60
```

Why?

Because operating at:

```text
100% capacity
```

leaves no room for:

- traffic spikes
    
- instance failures
    
- deployments
    
- GC pauses
    
- noisy neighbors
    
- dependency degradation
    

Capacity planning is fundamentally about **operating points**, not maximum theoretical throughput.

---

# 27. A Better Scalability Workflow

Use this sequence.

## Step 1 — Define SLOs

Example:

```text
P99 latency < 300 ms
Availability > 99.95%
Error rate < 0.1%
```

Without targets, “scalable” is meaningless.

---

## Step 2 — Quantify Load

Measure:

```text
RPS
concurrency
payload size
read/write ratio
database QPS
connections
storage growth
```

---

## Step 3 — Identify the Bottleneck

Use:

```text
metrics
profiling
tracing
database EXPLAIN
load testing
resource utilization
```

---

## Step 4 — Remove Waste

Before adding infrastructure:

```text
bad query
↓
fix query
```

```text
N+1 requests
↓
batch
```

```text
large payload
↓
compress/paginate
```

```text
duplicate work
↓
cache/coalesce
```

This is usually cheaper than architectural expansion.

---

## Step 5 — Scale the Bottleneck

```text
CPU bottleneck
→ horizontal scaling

DB read bottleneck
→ indexes/cache/read replicas

DB write bottleneck
→ batching/sharding/data-model changes

Long-running work
→ asynchronous workers

Network bandwidth
→ compression/CDN

Hot resource
→ partitioning/caching/sharding
```

---

# 28. The Scalability Decision Tree

A useful mental model:

```text
             Load increasing
                   │
                   ↓
             What saturates?
                   │
       ┌───────────┼───────────┐
       ↓           ↓           ↓
      CPU          DB        Network
       │           │           │
       ↓           ↓           ↓
   Scale out    Optimize     CDN/
   instances    queries      compression
                   │
              ┌────┴─────┐
              ↓          ↓
            Reads      Writes
              │          │
              ↓          ↓
           Cache/      partition/
           replicas    shard
```

But always validate with measurements.

---

# 29. Scaling Is Usually About Removing Coupling

A powerful way to think about scalability:

> **A system scales when independent work can be performed independently.**

Bad:

```text
Everything
   ↓
One lock
   ↓
One DB
   ↓
One queue
   ↓
One bottleneck
```

Better:

```text
Partitioned work
      ↓
Independent execution
      ↓
Independent resources
```

This is why these concepts repeatedly appear:

- Statelessness
    
- Partitioning
    
- Sharding
    
- Asynchronous processing
    
- Queueing
    
- Caching
    
- Replication
    
- Bounded concurrency
    

They reduce contention or isolate failure.

---

# 30. The Principal Engineer Perspective

At junior level, the question is:

> “Which scalability pattern should I use?”

At senior level:

> “What is the bottleneck?”

At Staff level:

> “What constraint is preventing the system from scaling?”

At Principal level:

> **“What architecture allows the workload to grow while preserving correctness, reliability, operability, and acceptable cost—and what complexity are we willing to introduce to achieve that?”**

That last part matters.

Every scalability technique has a price:

|Pattern|Benefit|Cost|
|---|---|---|
|Vertical scaling|Simplicity|Hardware ceiling|
|Horizontal scaling|Capacity|Distributed state|
|Cache|Lower load/latency|Staleness/invalidation|
|Read replicas|Read capacity|Replication lag|
|Async queue|Decoupling|Eventual consistency|
|Sharding|Huge scale|Distributed data|
|CDN|Global scale|Cache invalidation|
|Microservices|Independent scaling/ownership|Operational complexity|

The goal isn't:

> **Maximum scalability.**

The goal is:

> **Sufficient scalability with the minimum necessary complexity.**

---

# 31. Production Checklist

Before declaring a system “scalable,” ask:

```text
□ What is the expected peak load?
□ What is the 10x load?
□ What is the current bottleneck?
□ What happens when that bottleneck saturates?
□ Can we scale horizontally?
□ Is application state externalized?
□ Is concurrency bounded?
□ Is backpressure implemented?
□ Are queues bounded?
□ Are retries bounded?
□ Are retries idempotent?
□ Can dependencies fail independently?
□ Can the database handle connection growth?
□ What happens during cache failure?
□ Are there hot keys/partitions?
□ What happens during traffic spikes?
□ What happens when one instance dies?
□ What happens when an entire AZ/region fails?
□ What metrics tell us we're approaching saturation?
□ What is our P95/P99 behavior under load?
□ What is the cost at 1x, 5x, 10x?
□ Can we roll back the scaling change?
```

## Core takeaway

The most important scalability pattern is actually a **reasoning pattern**:

```text
Requirements
    ↓
Workload model
    ↓
Capacity model
    ↓
Bottleneck identification
    ↓
Smallest effective change
    ↓
Measure
    ↓
Failure analysis
    ↓
Scale further only when necessary
```

If you remember only one sentence:

> **Don't scale the system; scale the bottleneck, and make sure the scaling mechanism doesn't simply move the bottleneck somewhere else.**
---

## 🔗 References
- ⬆️ Parent: [[How to Approach]]
- 📚 Module: `Introduction`

