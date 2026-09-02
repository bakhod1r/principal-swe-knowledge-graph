---
title: "How to Approach Core Principles and Architecture"
tags:
  - review
  - system-design
  - architecture
  - distributed-systems
  - introduction
  - principal-swe
parent: "[[How to Approach]]"
---
# How to Approach Core Principles and Architecture

A strong engineer does **not start architecture by choosing technologies**.

The correct starting point is:

> **Problem → Constraints → Principles → Boundaries → Data Flow → Failure Modes → Architecture → Trade-offs**

Architecture is the **result of reasoning**, not the starting point.

---

## 1. Start With the Real Problem

Before drawing boxes, answer:

- What problem are we solving?
    
- Who has the problem?
    
- What does the system need to accomplish?
    
- What is explicitly **out of scope**?
    

For example:

> "We need an API for users to upload and retrieve profile images."

That is much better than:

> "Let's build an image service with Go, PostgreSQL, Redis, Kafka, and S3."

The second statement has already jumped into implementation.

### Principal-level question

Ask:

> **What business capability does this system provide?**

Not:

> What technologies should we use?

---

# 2. Identify Constraints

Architecture is largely determined by constraints.

Separate them into categories.

### Functional

What must the system do?

```text
Upload image
Download image
Delete image
Generate thumbnail
```

### Non-functional

How well must it work?

```text
Availability: 99.9%
p95 latency: < 200 ms
10M users
100K requests/sec
```

### Operational

What can the organization operate?

```text
Small engineering team
Limited DevOps expertise
One deployment region
24/7 service required
```

### Security

```text
Authentication
Authorization
Encryption
Auditability
Data retention
```

### Cost

```text
Infrastructure budget
Storage cost
Network egress
Operational complexity
```

This step prevents a common mistake:

> **Designing for imaginary scale.**

---

# 3. Establish Architecture Principles

Principles are the rules that guide decisions.

Typical principles:

### 1. Correctness before performance

Do not sacrifice correctness because something _might_ become slow.

### 2. Simplicity before abstraction

Every abstraction creates:

- cognitive cost
    
- maintenance cost
    
- debugging cost
    
- operational cost
    

### 3. Explicit ownership

Every important responsibility should have a clear owner.

For example:

```text
Authentication → Auth service
User data      → User domain
Image storage  → Object storage
```

### 4. Minimize distributed state

Distributed systems introduce:

```text
network failures
timeouts
retries
ordering problems
partial failure
consistency problems
```

Therefore:

> Don't distribute something unless there is a reason.

---

# 4. Define Boundaries

This is one of the most important architecture skills.

Ask:

> **Where does one responsibility end and another begin?**

For example:

```text
                    API
                     │
                     ▼
              User Management
                │          │
                ▼          ▼
           PostgreSQL    Redis
```

The important question isn't whether Redis is fast.

The important question is:

> **Who owns the data and responsibility?**

---

# 5. Define Data Ownership

For every piece of important state, identify its owner.

Example:

```text
User
 ├── id
 ├── email
 ├── password_hash
 └── status
```

Who owns this?

```text
User domain
```

Then:

```text
Order
 ├── id
 ├── user_id
 ├── amount
 └── status
```

Who owns it?

```text
Order domain
```

This leads to an important architectural rule:

> **Ownership should be explicit before communication is designed.**

---

# 6. Define API Boundaries

Once responsibilities are clear, define how components communicate.

For example:

```text
Client
  │
  │ HTTP
  ▼
API
  │
  ├── User domain
  │
  ├── Order domain
  │
  └── Payment domain
```

At this point you can ask:

- synchronous or asynchronous?
    
- HTTP or gRPC?
    
- request/response or events?
    
- transactional boundary?
    
- timeout?
    
- retry?
    
- idempotency?
    

---

# 7. Define Transaction Boundaries

This is frequently missed.

Suppose:

```text
CreateOrder()
    ↓
ChargePayment()
    ↓
SendEmail()
```

Should these be one transaction?

Usually **no**.

Database transactions generally cannot atomically include:

```text
PostgreSQL
+
Payment provider
+
Email provider
```

So you need to reason about partial failure.

Example:

```text
Create order
     ↓
Payment succeeds
     ↓
Email fails
```

The system must still remain correct.

This leads to patterns such as:

```text
Outbox
Retry
Idempotency
Saga
Compensation
```

Architecture starts becoming interesting exactly here.

---

# 8. Model Failure Before Success

Do not design only:

```text
Request
 ↓
Service
 ↓
Database
 ↓
Success
```

Design:

```text
Request
 ↓
Service
 ↓
Database
 │
 ├── timeout
 ├── connection failure
 ├── overload
 ├── deadlock
 └── unavailable
```

And for network calls:

```text
Service A
   │
   │ request
   ▼
Service B
   │
   ├── timeout
   ├── duplicate request
   ├── response lost
   ├── partial processing
   └── unavailable
```

This produces much better architecture.

---

# 9. Think in Failure Boundaries

A useful mental model:

```text
                 SYSTEM
──────────────────────────────────
        │               │
        ▼               ▼
     Service A       Service B
        │               │
        ▼               ▼
       DB              DB
```

If Service B fails:

> Does A continue working?

If the answer is no, you have created a dependency/failure boundary.

This is why:

> **A distributed architecture is also a distributed failure model.**

---

# 10. Choose Architecture Style Last

Only now choose:

```text
Monolith
Modular Monolith
Microservices
Event-driven
Serverless
```

Do **not** start with:

> "We need microservices."

Instead ask:

> What constraint requires independent deployment, scaling, ownership, or failure isolation?

### Example

If you have:

```text
5 engineers
50K users
moderate traffic
one database
simple domain
```

A:

```text
Modular Monolith
```

may be the best architecture.

If you have:

```text
500 engineers
multiple teams
independent deployment requirements
very different scaling characteristics
strong domain boundaries
```

Microservices become more defensible.

---

# 11. Use the Simplest Architecture That Satisfies Constraints

A useful progression is:

```text
Single process
     ↓
Modular monolith
     ↓
Service extraction
     ↓
Distributed services
     ↓
Event-driven components
```

Do not climb this ladder prematurely.

Every step adds complexity.

For example:

```text
Monolith
```

has relatively simple:

```text
deployment
transactions
debugging
local development
testing
network communication
observability
```

Microservices introduce:

```text
service discovery
network calls
timeouts
retries
distributed tracing
deployment coordination
schema evolution
distributed transactions
```

Therefore:

> **Microservices solve organizational and scaling problems, but create distributed-systems problems.**

---

# 12. Analyze Data Flow

Architecture diagrams often show components but fail to explain **flow**.

You should be able to trace:

```text
User
 ↓
Load Balancer
 ↓
API
 ↓
Application
 ↓
Database
 ↓
Response
```

Then examine:

### Read path

```text
Client
 ↓
API
 ↓
Cache
 ↓ miss
Database
```

### Write path

```text
Client
 ↓
API
 ↓
Database
 ↓
Outbox
 ↓
Worker
 ↓
Message broker
 ↓
Consumers
```

Now ask:

- Where can requests duplicate?
    
- Where can data become stale?
    
- Where can messages be lost?
    
- Where can ordering break?
    
- Where can backpressure occur?
    

---

# 13. Consider Consistency

Ask:

> Does every reader need the latest value immediately?

If yes:

```text
Strong consistency
```

may be required.

If slight staleness is acceptable:

```text
Cache
Replica
Async processing
Eventual consistency
```

may be acceptable.

Example:

```text
Bank balance
```

usually requires stronger consistency.

While:

```text
View count
```

may tolerate eventual consistency.

The architectural decision comes from the **business invariant**, not from technology preference.

---

# 14. Consider Scaling

Don't simply say:

> "We'll add Kubernetes."

First identify the bottleneck.

```text
Traffic
  ↓
CPU?
Memory?
Database?
Network?
Locks?
External API?
Disk?
```

For example:

```text
API:       10K req/s
Database:  2K writes/s
```

Adding API servers won't solve the database bottleneck.

You need to reason about:

```text
capacity
load
concurrency
latency
saturation
queue depth
```

---

# 15. Add Reliability Mechanisms

Once failure modes are understood:

```text
Timeouts
Retries
Idempotency
Rate limiting
Backpressure
Circuit breakers
Bulkheads
Graceful shutdown
Health checks
```

But don't add them mechanically.

For every mechanism ask:

> **What failure does this prevent or contain?**

For example:

### Timeout

Prevents:

```text
goroutine/request
        ↓
waiting forever
```

### Retry

Handles transient failure.

But retries can cause:

```text
failure
 ↓
retry
 ↓
failure
 ↓
retry
 ↓
retry storm
```

So retries often require:

```text
bounded attempts
+
exponential backoff
+
jitter
+
idempotency
```

---

# 16. Design Observability

Architecture isn't production-ready until you can answer:

> **What is the system doing right now?**

Think in:

### Metrics

```text
request rate
error rate
p50
p95
p99
CPU
memory
queue depth
DB connections
```

### Logs

Useful for detailed events.

### Traces

Useful for distributed request flow.

```text
Client
  │
  ▼
API ────── 20ms
  │
  ▼
Order ──── 40ms
  │
  ▼
DB ─────── 150ms  ← bottleneck
```

---

# 17. Security Is an Architectural Property

Don't bolt security on afterward.

For every boundary ask:

```text
Who is calling?
What are they allowed to do?
What data can they access?
How is the request validated?
Can it be replayed?
Can it be abused?
```

Think:

```text
Authentication
Authorization
Validation
Least privilege
Encryption
Secrets management
Rate limiting
Audit logs
```

---

# 18. Test the Architecture

Architecture should imply a testing strategy.

Example:

```text
Unit tests
    ↓
Integration tests
    ↓
Contract tests
    ↓
E2E tests
    ↓
Load tests
    ↓
Failure tests
```

For distributed systems especially test:

```text
timeout
duplicate request
duplicate message
out-of-order message
database failure
network partition
consumer crash
retry storm
```

---

# 19. Evaluate Trade-offs Explicitly

Never say:

> "Kafka is better."

Ask:

|Decision|Benefit|Cost|
|---|---|---|
|Async messaging|Decoupling|Operational complexity|
|Redis|Lower read latency|Stale data/cache invalidation|
|Microservices|Independent deployment|Distributed-system complexity|
|PostgreSQL|Strong consistency|Write scalability limits|
|Eventual consistency|Scalability|More complex correctness model|

There is rarely a universally "best" architecture.

There is usually:

> **the best architecture under a specific constraint set.**

---

# 20. The Architecture Decision Loop

Use this repeatedly:

```text
             ┌───────────────┐
             │ Real Problem  │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │  Constraints  │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │ Requirements  │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │   Boundaries  │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │ Data Ownership│
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │   Data Flow   │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │ Failure Modes │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │ Architecture  │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │  Trade-offs   │
             └───────┬───────┘
                     ↓
             ┌───────────────┐
             │ Measurements  │
             └───────┬───────┘
                     │
                     └──────→ Iterate
```

---

# The Most Important Mental Model

When looking at any architecture, ask these **10 questions**:

1. **What is the real problem?**
    
2. **What are the hard constraints?**
    
3. **Who owns each responsibility and piece of data?**
    
4. **Where are the boundaries?**
    
5. **What are the transaction/consistency requirements?**
    
6. **What happens when each dependency fails?**
    
7. **Where is the bottleneck at 10× load?**
    
8. **How do we observe and debug it?**
    
9. **How do we migrate and roll it back?**
    
10. **What complexity are we introducing, and is it justified?**
    

That is the transition from **"I know architecture patterns"** to **"I can reason about architecture."**

### Staff/Principal insight

The highest-level architectural skill is not knowing more patterns.

It is knowing **when not to use them**.

A Principal Engineer should be able to explain:

> "We could use Kafka here, but the business requirement does not need asynchronous processing or independent consumers. PostgreSQL + an outbox would introduce less operational complexity while satisfying the current constraints."

That is architecture judgment.

**Core principle:**

> **Architecture is constrained optimization under uncertainty and failure.**

Not box-drawing.
---

## 🔗 References
- ⬆️ Parent: [[How to Approach]]
- 📚 Module: `Introduction`

