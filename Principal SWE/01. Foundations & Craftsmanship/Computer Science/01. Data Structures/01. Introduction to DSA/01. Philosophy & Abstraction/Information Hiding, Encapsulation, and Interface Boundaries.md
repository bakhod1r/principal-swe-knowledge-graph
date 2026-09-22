---
title: "Information Hiding, Encapsulation, and Interface Boundaries"
tags:

  - computer-science
  - data-structures
  - principal-swe
parent: "[[Introduction to DSA]]"
---
# Information Hiding, Encapsulation, and Interface Boundaries

These three concepts are closely related, but they solve **different problems**. A strong engineer should distinguish them because many poor designs come from treating them as synonyms.

---

## 1. The Core Mental Model

Think of a component as a **machine**:

```text
                 Public Boundary
                       │
        ┌──────────────▼──────────────┐
        │          Component          │
        │                             │
Input ──►  Internal State + Algorithms ├──► Output
        │                             │
        │   implementation details    │
        └─────────────────────────────┘
                       ▲
                       │
                Hidden from users
```

The fundamental idea is:

> **Expose what clients need to know; hide what clients do not need to know.**

This creates a boundary between:

- **what the component promises**

- **how the component fulfills that promise**


That boundary is one of the most important ideas in software architecture.

---

# 2. Information Hiding

## Definition

**Information hiding** means deliberately hiding design decisions that are likely to change, so that changes do not propagate to consumers.

The important part is:

> Information hiding is about **hiding decisions**, not merely hiding data.

For example:

```go
type UserRepository struct {
    db *sql.DB
}
```

A caller doesn't need to know:

- which SQL database is used

- which SQL query is executed

- which indexes exist

- whether caching exists

- how connections are managed


The caller only needs:

```go
user, err := repo.Get(ctx, userID)
```

The database implementation is an **information-hiding boundary**.

### Key insight

A common beginner interpretation is:

> "Private fields = information hiding."

Not necessarily.

You can have private fields and still expose too much implementation knowledge through your API.

---

# 3. Encapsulation

**Encapsulation** is the mechanism of grouping state and behavior together while controlling access to that state.

Example:

```go
type Account struct {
    balance int64
}

func (a *Account) Deposit(amount int64) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }

    a.balance += amount
    return nil
}

func (a Account) Balance() int64 {
    return a.balance
}
```

The important property is:

```text
balance
   │
   ├── Deposit()
   └── Balance()
```

The state and operations governing that state live together.

The caller cannot arbitrarily do:

```go
account.balance = -1000
```

because `balance` is private.

Instead:

```go
account.Deposit(100)
```

The object controls how its state changes.

---

# 4. Information Hiding vs Encapsulation

They overlap, but they are not identical.

|Concept|Primary concern|
|---|---|
|Encapsulation|Control access to state/behavior|
|Information hiding|Hide implementation/design decisions|
|Abstraction|Present essential concepts while ignoring irrelevant details|
|Interface boundary|Define the contract between components|

A useful relationship is:

```text
Encapsulation
      │
      ▼
Controls access
      │
      ▼
Helps achieve
      │
      ▼
Information hiding
      │
      ▼
Creates stable boundaries
      │
      ▼
Enables abstraction
```

But none of these automatically guarantees the others.

---

# 5. Interface Boundaries

An **interface boundary** defines what one component expects from another.

In Go:

```go
type UserStore interface {
    Get(ctx context.Context, id int64) (User, error)
    Save(ctx context.Context, user User) error
}
```

The consumer depends on:

```text
UserStore
   │
   ├── Get()
   └── Save()
```

rather than:

```text
PostgreSQL
SQL queries
connection pool
transactions
indexes
database schema
```

A concrete implementation could be:

```go
type PostgresUserStore struct {
    db *sql.DB
}

func (s *PostgresUserStore) Get(
    ctx context.Context,
    id int64,
) (User, error) {
    // SQL implementation
}
```

The consumer doesn't need to know this.

---

# 6. The Most Important Principle

The quality of an interface is determined less by **how small it is** and more by **what knowledge it forces the caller to have**.

Bad boundary:

```go
type Database interface {
    Exec(query string, args ...any) error
    Query(query string, args ...any) (*sql.Rows, error)
    Begin() (*sql.Tx, error)
}
```

Now your application knows about SQL.

That means the abstraction failed to hide the important implementation detail.

Better:

```go
type UserStore interface {
    Get(ctx context.Context, id int64) (User, error)
    Save(ctx context.Context, user User) error
}
```

Now the application expresses its domain requirement.

This is a major Staff/Principal-level distinction:

> **An abstraction should hide volatile implementation decisions, not merely rename them.**

---

# 7. Leaky Abstractions

An abstraction becomes **leaky** when implementation details escape through its boundary.

Suppose:

```go
type UserStore interface {
    Get(ctx context.Context, id int64) (*sql.Row, error)
}
```

This looks abstract, but:

```text
UserStore
    │
    └── *sql.Row
           │
           ▼
       SQL leaked
```

The consumer now depends on `database/sql`.

That's a leaky abstraction.

A better API:

```go
type UserStore interface {
    Get(ctx context.Context, id int64) (User, error)
}
```

Now the storage mechanism is hidden.

---

# 8. Why Information Hiding Matters

Imagine 200 files directly use:

```go
redis.Client
```

Then you decide:

```text
Redis
  ↓
PostgreSQL
```

You potentially have hundreds of consumers to modify.

But if your application depends on:

```go
type SessionStore interface {
    Get(ctx context.Context, id string) (Session, error)
    Put(ctx context.Context, session Session) error
}
```

then:

```text
Application
     │
     ▼
SessionStore
     │
 ┌───┴────┐
 ▼        ▼
Redis    PostgreSQL
```

Changing the implementation becomes localized.

This is the real economic value of information hiding:

> **It limits the blast radius of change.**

---

# 9. Change Amplification

Consider:

```text
Implementation detail
        │
        ├── Service A
        ├── Service B
        ├── Service C
        ├── Service D
        └── Service E
```

Changing the detail causes:

```text
1 change → 5 consumers affected
```

With proper information hiding:

```text
             Interface
                │
       ┌────────┴────────┐
       ▼                 ▼
 Implementation A   Implementation B
```

Changing implementation A may affect:

```text
1 implementation
```

instead of:

```text
N consumers
```

This is why good boundaries reduce **change amplification**.

---

# 10. Abstraction Is Not "Making Everything an Interface"

This is a very common Go mistake.

Bad:

```go
type UserService interface {
    Create(...)
    Update(...)
    Delete(...)
}

type UserServiceImpl struct {
    repo UserRepository
}
```

If there is only one consumer and no meaningful boundary, this may add complexity without hiding anything useful.

You have:

```text
Interface
   ↓
Implementation
   ↓
Implementation
```

with no meaningful architectural benefit.

Go's philosophy strongly favors concrete types by default.

A useful rule:

> **Create an interface when a consumer needs a behavioral boundary, not simply because an implementation exists.**

---

# 11. Consumer-Owned Interfaces

In Go, a particularly useful pattern is:

```go
// package service

type UserReader interface {
    Get(ctx context.Context, id int64) (User, error)
}
```

The consumer defines the smallest capability it needs.

Suppose the implementation provides:

```go
type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) Get(...) ...
func (r *UserRepository) Save(...) ...
func (r *UserRepository) Delete(...) ...
func (r *UserRepository) Search(...) ...
```

The service only needs:

```go
type UserReader interface {
    Get(...)
}
```

The repository automatically satisfies it.

This gives:

```text
Consumer
   │
   │ requires
   ▼
small interface
   ▲
   │ satisfies
   │
Concrete implementation
```

No explicit declaration is necessary.

---

# 12. Dependency Direction

A strong architecture usually makes dependency direction explicit.

Bad:

```text
Domain
  ↓
PostgreSQL
  ↓
SQL
```

The business logic becomes coupled to infrastructure.

Better:

```text
              Interface
             ▲         ▲
             │         │
         Domain      Adapter
             │         │
             │       PostgreSQL
             │
          Business
           Logic
```

The important question isn't:

> "Where should interfaces live?"

The deeper question is:

> **Who owns the dependency?**

The consumer should define the capability it requires when that creates a meaningful boundary.

---

# 13. Information Hiding and API Design

Suppose you expose:

```go
type Config struct {
    Host string
    Port int
    TLS  bool
}
```

You're saying:

> "These fields are part of my public contract."

Changing:

```go
Port int
```

to:

```go
Address string
```

becomes a breaking API change.

Sometimes that's correct.

But if these are implementation details:

```go
type Server struct {
    listener net.Listener
    workers  int
    cache    map[string]Entry
}
```

make them private:

```go
type Server struct {
    listener net.Listener
    workers  int
    cache    map[string]Entry
}
```

Now you can change the internals without forcing callers to change.

---

# 14. Encapsulation Protects Invariants

This is one of the most important reasons for encapsulation.

Suppose:

```go
type Queue struct {
    items []Item
}
```

If `items` is publicly mutable:

```go
q.items = nil
```

the caller can violate internal assumptions.

Instead:

```go
type Queue struct {
    items []Item
}

func (q *Queue) Push(item Item) {
    q.items = append(q.items, item)
}

func (q *Queue) Pop() (Item, bool) {
    if len(q.items) == 0 {
        return Item{}, false
    }

    item := q.items[0]
    q.items = q.items[1:]

    return item, true
}
```

Now the component controls its invariant:

```text
Queue invariant:
len(items) >= 0
```

More complex components might have:

```text
head <= tail
size <= capacity
state transitions are valid
ownership is preserved
```

Encapsulation gives the component authority over these invariants.

---

# 15. Encapsulation and Concurrency

This becomes even more important in concurrent systems.

Bad:

```go
type Counter struct {
    Value int64
}
```

Multiple goroutines can mutate it:

```go
counter.Value++
```

The state is exposed without synchronization.

Better:

```go
type Counter struct {
    mu    sync.Mutex
    value int64
}

func (c *Counter) Inc() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}

func (c *Counter) Value() int64 {
    c.mu.Lock()
    defer c.mu.Unlock()

    return c.value
}
```

Now synchronization is encapsulated with the state.

The caller doesn't need to understand:

```text
mutex
critical section
memory ordering
```

to safely use the abstraction.

This is a powerful principle:

> **Encapsulation can hide not only data structures, but also synchronization mechanisms and concurrency invariants.**

---

# 16. Information Hiding vs Performance

Information hiding does **not** mean hiding everything.

Sometimes consumers genuinely need performance-related information.

For example:

```go
cache.Get()
```

may hide the implementation, but the caller may still need to understand:

```text
O(1) average
may block
may perform network I/O
may return stale data
```

An abstraction can hide **implementation**, while still documenting **behavioral guarantees**.

This distinction is critical.

You should hide:

```text
HOW
```

while exposing necessary:

```text
WHAT
GUARANTEES
COST
LIMITATIONS
```

For example:

```go
Get(ctx, key)
```

could document:

```text
- Returns the latest committed value.
- May perform network I/O.
- Respects context cancellation.
- Returns ErrNotFound when absent.
```

The caller knows the operational contract without knowing the implementation.

---

# 17. Semantic vs Syntactic Boundaries

A weak boundary is syntactic:

```go
type UserService interface {
    DoSomething(...)
}
```

A strong boundary is semantic:

```go
type UserRepository interface {
    FindByEmail(ctx context.Context, email string) (User, error)
}
```

The second communicates a domain capability.

Good interfaces answer:

> **What can I ask this component to do?**

Bad interfaces answer:

> **What methods happen to exist on this object?**

---

# 18. The Granularity Problem

Interfaces can be too large:

```go
type Repository interface {
    Create(...)
    Get(...)
    Update(...)
    Delete(...)
    Search(...)
    Count(...)
    Begin(...)
    Commit(...)
    Rollback(...)
}
```

This increases coupling.

Or too small:

```go
type Getter interface {
    Get(...)
}

type Creator interface {
    Create(...)
}

type Updater interface {
    Update(...)
}
```

Now you may have dozens of meaningless abstractions.

The goal is **cohesive boundaries**.

A good interface groups operations that represent a meaningful capability.

---

# 19. The "Stable Core, Volatile Edge" Mental Model

A useful architecture model is:

```text
             Volatile
                │
     ┌──────────┴──────────┐
     │                     │
 HTTP/API              Database
     │                     │
     └──────────┬──────────┘
                │
             Boundary
                │
                ▼
          Stable Domain
             Logic
```

The more volatile something is, the more carefully you should prevent it from leaking into stable parts.

Typical volatile details:

- HTTP framework

- SQL driver

- Redis client

- Kafka client

- cloud SDK

- filesystem

- external API

- serialization format


Stable concepts:

- business rules

- domain invariants

- core workflows

- business decisions


This is one reason hexagonal/ports-and-adapters architecture can work well.

But don't introduce it mechanically. The boundary should solve an actual coupling problem.

---

# 20. Anti-Patterns

### 1. Getter/Setter Everything

```go
type User struct {
    name string
}

func (u *User) SetName(name string) {
    u.name = name
}

func (u *User) GetName() string {
    return u.name
}
```

This may merely recreate public fields with extra ceremony.

Ask:

> What invariant or behavior does the method protect?

---

### 2. Interface for Every Struct

```go
type UserService interface {...}
type UserServiceImpl struct {...}
```

Often unnecessary.

Prefer:

```go
type UserService struct {
    repo UserRepository
}
```

unless a real interface boundary exists.

---

### 3. Generic "Manager" Interfaces

```go
type Manager interface {
    Execute(...)
    Handle(...)
    Process(...)
}
```

These often hide responsibilities instead of clarifying them.

---

### 4. Leaking Infrastructure Types

```go
func GetUser(...) *sql.Row
```

The database has escaped the repository boundary.

---

### 5. Exposing Mutable Internal State

```go
func (c *Cache) Items() map[string]Item
```

The caller can mutate your internal state.

Potentially better:

```go
func (c *Cache) Get(key string) (Item, bool)
```

or return a defensive copy where appropriate.

---

# 21. A Production Design Example

Imagine an order service.

Naive:

```go
func CreateOrder(
    db *sql.DB,
    redis *redis.Client,
    kafka *kafka.Writer,
    req *http.Request,
) error
```

Everything leaks into the business operation.

The business logic now knows:

```text
HTTP
SQL
Redis
Kafka
```

Better:

```go
type OrderStore interface {
    Create(ctx context.Context, order Order) error
}

type EventPublisher interface {
    Publish(ctx context.Context, event Event) error
}

type OrderService struct {
    store     OrderStore
    publisher EventPublisher
}
```

Now:

```text
                 OrderService
                 /           \
                /             \
        OrderStore         EventPublisher
            │                    │
        PostgreSQL              Kafka
```

The service knows:

```text
"persist an order"
"publish an event"
```

It does not need to know:

```text
SQL
Kafka protocol
connection pooling
serialization
```

That's information hiding.

---

# 22. But Don't Over-Abstract

Suppose your application is:

```text
500 LOC
one PostgreSQL database
one deployment
no alternative storage
no testing difficulty
no architectural boundary
```

You might simply write:

```go
type OrderService struct {
    db *sql.DB
}
```

and keep the code straightforward.

The interface can be introduced later when the boundary becomes valuable.

This is an important engineering judgment:

> **Abstraction has a cost. Information hiding is valuable only when the hidden decision is worth isolating.**

---

# 23. Testing Benefit

Good boundaries can make testing easier.

Production:

```text
OrderService
     │
     ▼
PostgresOrderStore
```

Test:

```text
OrderService
     │
     ▼
FakeOrderStore
```

But don't create interfaces solely because "unit tests require mocks."

Often a better test is an integration test against a real PostgreSQL instance.

The interface should represent an architectural dependency, not merely a testing trick.

---

# 24. Security Benefit

Information hiding also reduces accidental access.

For example:

```go
type User struct {
    id           int64
    passwordHash string
}
```

Expose:

```go
func (u User) ID() int64
```

but don't expose the password hash unnecessarily.

Even better, separate concepts:

```go
type User struct {
    id    int64
    email string
}
```

and keep authentication credentials in a dedicated security boundary.

The principle is:

> **Don't expose information merely because the implementation possesses it.**

---

# 25. How to Evaluate a Boundary

When reviewing an API or package, ask:

### Question 1

**What implementation decisions are hidden?**

If the answer is "none", the abstraction may be pointless.

### Question 2

**What changes can happen without affecting consumers?**

This measures the value of information hiding.

### Question 3

**What knowledge must the consumer have?**

If it needs SQL, Redis, Kafka, HTTP, filesystem details, etc., the boundary may be leaking.

### Question 4

**Who owns the invariant?**

The component that owns the invariant should usually control the relevant state.

### Question 5

**Can the implementation change independently?**

If yes, the boundary is providing real value.

### Question 6

**What guarantees remain visible?**

Hiding implementation must not hide important behavioral contracts.

---

# 26. The Principal-Level Mental Model

Think about boundaries in terms of **knowledge and change propagation**.

Suppose:

```text
A → B → C → D
```

If A must understand C and D to use B:

```text
A knows B
A knows C
A knows D
```

then B is a weak boundary.

A stronger design:

```text
A
│
▼
B
│
├── C
└── D
```

A only knows B's contract.

Therefore:

```text
Implementation changes
        ↓
B absorbs change
        ↓
A remains unchanged
```

That is the real purpose.

---

# 27. One-Line Definitions

Keep these mental models:

**Encapsulation**

> Control access to state and behavior so the component can protect its invariants.

**Information hiding**

> Hide implementation decisions that clients should not depend upon.

**Abstraction**

> Expose the essential model while omitting irrelevant details.

**Interface boundary**

> Define the contract through which two components interact.

And the deeper connection:

```text
Encapsulation
      ↓
Protect ownership + invariants

Information Hiding
      ↓
Reduce knowledge + change propagation

Abstraction
      ↓
Expose essential behavior

Interface
      ↓
Make the behavioral boundary explicit
```

## The key engineering lesson

Don't ask:

> **"Where can I add an interface?"**

Ask:

> **"What knowledge should this component NOT need, and which implementation decisions should be able to change without forcing consumers to change?"**

That question leads to much better architecture than mechanically applying OOP patterns.

---

## 🔗 References
- ⬆️ Parent: [[Introduction to DSA]]
- 📚 Module: `Data Structures`
