---
title: "Context Design Rules"
tags:

  - golang
  - concurrency
  - principal-swe
parent: "[[Context Trees & Request Cancellation]]"
---
# Context Design Rules in Go

`context.Context` is easy to misuse because it looks like a generic parameter object. It is not.

The right mental model is:

> **`context.Context` is a request-lifetime control plane for cancellation, deadlines, and scoped metadata.**

It should answer:

- When should this operation stop?
    
- What deadline applies?
    
- Has the parent operation been cancelled?
    
- What small, request-scoped metadata must travel with the operation?
    

It should **not** become a general-purpose dependency container, configuration object, or business-state bag.

---

## 1. The Core Design Rule

A good Go API generally follows:

```go
func DoSomething(ctx context.Context, input Input) error
```

rather than:

```go
func DoSomething(input Input) error
```

when the operation can reasonably be cancelled, bounded by a deadline, or needs request-scoped metadata.

The important invariant is:

```text
caller owns the lifetime
        │
        ▼
      ctx
        │
        ├── cancellation
        ├── deadline
        └── scoped values
             │
             ▼
          operation
```

The callee **observes** the context; it normally does not own the lifetime of the caller's context.

---

# 2. Rule: Pass `context.Context` Explicitly

Prefer:

```go
func (s *Service) CreateUser(
    ctx context.Context,
    req CreateUserRequest,
) error {
    // ...
}
```

Avoid:

```go
type Service struct {
    ctx context.Context
}
```

and:

```go
type Request struct {
    Context context.Context
    Name    string
}
```

### Why?

A context represents the lifetime of a particular operation.

A service object often lives much longer:

```text
Service lifetime
──────────────────────────────────────────>

Request 1:       ├─────────┤
Request 2:              ├────────────┤
Request 3:                         ├──────┤
```

Storing a request context inside the service creates a lifetime mismatch.

The service may accidentally retain:

- cancellation state
    
- deadlines
    
- request metadata
    
- tracing information
    
- references reachable through context values
    

for much longer than intended.

---

# 3. Rule: `ctx` Should Usually Be the First Parameter

Idiomatic Go:

```go
func QueryUser(ctx context.Context, id int64) (*User, error)
```

Not:

```go
func QueryUser(id int64, ctx context.Context) (*User, error)
```

This creates a predictable API shape:

```go
func Method(
    ctx context.Context,
    explicitArguments ...,
)
```

It also makes context propagation visually obvious during code review.

---

# 4. Rule: Never Pass `nil` Context

Do not do this:

```go
QueryUser(nil, id)
```

A context should generally be:

```go
context.Background()
```

or:

```go
context.TODO()
```

if a real parent context is not yet available.

Why?

Because most context-aware APIs assume:

```go
ctx != nil
```

A nil context can cause:

```go
ctx.Done()
ctx.Err()
ctx.Deadline()
ctx.Value(...)
```

to panic.

---

# 5. Rule: Propagate the Caller Context

Suppose we have:

```go
func (s *Service) CreateOrder(
    ctx context.Context,
    req CreateOrderRequest,
) error {
    return s.repo.Insert(ctx, req)
}
```

This is good.

The cancellation tree remains intact:

```text
HTTP request
     │
     ▼
   handler
     │
     ▼
  service
     │
     ▼
 repository
     │
     ▼
 database
```

If the client disconnects:

```text
client disconnect
       │
       ▼
HTTP context cancelled
       │
       ▼
service observes cancellation
       │
       ▼
DB query cancelled
```

This is one of the primary reasons `context.Context` exists.

---

# 6. Rule: Do Not Replace the Parent Context Arbitrarily

Bad:

```go
func (s *Service) CreateOrder(
    ctx context.Context,
    req Request,
) error {
    ctx = context.Background()

    return s.repo.Insert(ctx, req)
}
```

You just destroyed the cancellation relationship.

If the request was cancelled:

```text
client
  X
  │
  ▼
handler ctx cancelled

but:

service → Background()
             │
             ▼
          DB query
          continues
```

This can cause:

- unnecessary work
    
- resource consumption
    
- goroutine retention
    
- DB connections remaining busy
    
- wasted CPU
    
- poor tail latency
    

---

# 7. Rule: Derived Contexts Are for Narrowing Lifetime

Creating a child context is appropriate when you want a **stricter constraint**.

For example:

```go
func (s *Service) CreateOrder(
    ctx context.Context,
    req Request,
) error {
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    return s.repo.Insert(ctx, req)
}
```

Mental model:

```text
Parent deadline
       │
       ▼
  10 seconds
       │
       ▼
Child deadline
       │
       ▼
   2 seconds
```

The child cannot outlive its parent.

Formally:

```text
child lifetime ≤ parent lifetime
```

This is a critical design property.

---

# 8. Rule: Always Release Context Resources

Many context constructors allocate resources.

For example:

```go
ctx, cancel := context.WithTimeout(parent, time.Second)
defer cancel()
```

Even when the timeout will eventually fire, explicitly calling:

```go
cancel()
```

is good resource hygiene.

Think:

```text
WithTimeout
    │
    ├── child context
    └── timer
         │
         ▼
      cancel()
         │
         ▼
    release resources
```

So this:

```go
ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()
```

should be your default pattern.

---

# 9. Rule: Don't Use Context for Normal Function Arguments

Bad:

```go
ctx = context.WithValue(ctx, "userID", userID)
```

followed by:

```go
func Process(ctx context.Context) error {
    userID := ctx.Value("userID")
    // ...
}
```

if `userID` is actually required business input.

Prefer:

```go
func Process(
    ctx context.Context,
    userID int64,
) error
```

Why?

Because the function signature communicates its contract.

Compare:

```go
func Process(ctx context.Context) error
```

versus:

```go
func Process(ctx context.Context, userID int64) error
```

The second is much easier to understand, test, refactor, and statically analyze.

---

# 10. Rule: Context Values Are for Request-Scoped Metadata

Good candidates can include things such as:

```text
trace ID
request ID
authenticated principal
locale
request metadata
```

For example:

```go
type contextKey struct {
    name string
}

var traceIDKey contextKey
```

Then:

```go
ctx = context.WithValue(ctx, traceIDKey, traceID)
```

and:

```go
traceID, ok := ctx.Value(traceIDKey).(string)
```

But even here, use restraint.

A useful test is:

> **Would this value naturally disappear when the request/operation ends?**

If yes, context may be appropriate.

If it's application state, configuration, or a dependency, probably not.

---

# 11. Rule: Never Use Basic Types as Context Keys

Avoid:

```go
context.WithValue(ctx, "userID", id)
```

because unrelated packages can accidentally collide:

```go
"userID"
"userID"
```

Prefer a private key type:

```go
type contextKey struct{}

var userIDKey contextKey
```

or:

```go
type userIDKey struct{}
```

The package-private type provides stronger namespace isolation.

---

# 12. Rule: Context Is Not a Dependency Injection Container

This is a common architectural mistake.

Bad:

```go
ctx = context.WithValue(ctx, dbKey, db)
ctx = context.WithValue(ctx, loggerKey, logger)
ctx = context.WithValue(ctx, configKey, config)
ctx = context.WithValue(ctx, cacheKey, cache)
```

Eventually:

```go
func Process(ctx context.Context) error
```

secretly depends on:

```text
DB
Logger
Config
Cache
Metrics
FeatureFlags
User
Tenant
...
```

The signature lies.

Instead:

```go
type Service struct {
    db     *sql.DB
    logger *slog.Logger
    cache  Cache
}
```

Dependencies should normally be explicit and owned by the component that uses them.

---

# 13. Rule: Don't Store Context in Structs

Avoid:

```go
type Worker struct {
    ctx context.Context
}
```

especially when the context belongs to a request.

Prefer:

```go
type Worker struct {
    // long-lived dependencies
}

func (w *Worker) Run(ctx context.Context) error {
    // request/operation lifetime
}
```

There is one important nuance:

For a **long-lived component whose lifetime is itself controlled by a context**, storing a context as part of that component's lifecycle can be valid internally. But even then, the design should make the ownership explicit.

For example:

```go
type Server struct {
    ctx context.Context
}
```

may be reasonable if the server itself owns that lifetime.

The problem isn't literally "context must never be stored."

The real rule is:

> **Do not store a context merely to avoid passing it explicitly.**

---

# 14. Rule: Don't Create `Background()` Deep Inside Business Logic

Bad:

```go
func Save(ctx context.Context, data Data) error {
    ctx = context.Background()

    return db.Save(ctx, data)
}
```

This breaks propagation.

Another subtle version:

```go
func Save(ctx context.Context, data Data) error {
    go func() {
        process(context.Background(), data)
    }()

    return nil
}
```

Now the background task is detached from the request.

Sometimes detachment is intentional, but it should be an explicit architectural decision.

---

# 15. Rule: Understand `WithoutCancel`

Go provides:

```go
ctx2 := context.WithoutCancel(ctx)
```

Conceptually:

```text
parent
  │
  ├── normal child
  │       └── cancellation propagates
  │
  └── WithoutCancel child
          └── cancellation does NOT propagate
```

This is useful when some work should survive cancellation of the parent.

But this should immediately trigger a design question:

> **Why should this work outlive the operation that created it?**

If the answer is unclear, `WithoutCancel` is probably hiding a lifecycle problem.

---

# 16. Rule: Detached Work Needs Its Own Lifetime

Consider:

```go
func Handler(ctx context.Context) {
    go func() {
        ctx := context.WithoutCancel(ctx)
        sendAuditEvent(ctx)
    }()
}
```

You have intentionally detached cancellation.

But now:

```text
HTTP request
     │
     X cancelled
     │
     ▼
audit goroutine
     │
     └── still running
```

What stops it?

Potentially nothing.

A production design may instead give the work its own bounded lifetime:

```go
func Handler(ctx context.Context) {
    go func() {
        ctx := context.WithoutCancel(ctx)

        ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
        defer cancel()

        sendAuditEvent(ctx)
    }()
}
```

Now the lifecycle is:

```text
request cancellation
       │
       X
       │
       ▼
detached operation
       │
       │ max 5s
       ▼
    cancellation
```

But for important background work, a worker queue/job system is often cleaner than spawning detached goroutines from request handlers.

---

# 17. Rule: Cancellation Must Be Observed

Passing a context does not magically cancel your code.

This:

```go
func Work(ctx context.Context) error {
    expensiveOperation()
    return nil
}
```

does not automatically stop `expensiveOperation()`.

You must cooperate:

```go
func Work(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()

        default:
            // perform bounded work
        }
    }
}
```

Or pass the context to APIs that understand it:

```go
db.QueryContext(ctx, query)
http.NewRequestWithContext(ctx, ...)
```

The key distinction:

```text
Context cancellation
        ≠
automatic interruption of arbitrary Go code
```

It is a **cooperative cancellation mechanism**.

---

# 18. Rule: Prefer `ctx.Err()` for Cancellation Cause at Boundaries

A common pattern:

```go
select {
case <-ctx.Done():
    return ctx.Err()

case result := <-results:
    return result
}
```

Typical values:

```go
context.Canceled
context.DeadlineExceeded
```

This allows upper layers to distinguish:

```text
operation succeeded
operation failed
operation was cancelled
operation exceeded deadline
```

That distinction becomes valuable for:

- HTTP status mapping
    
- retry decisions
    
- metrics
    
- logging
    
- tracing
    
- incident debugging
    

---

# 19. Rule: Don't Blindly Retry Context Errors

Bad:

```go
for i := 0; i < 5; i++ {
    err := operation(ctx)
    if err == nil {
        return nil
    }

    time.Sleep(time.Second)
}
```

If:

```text
ctx cancelled
     │
     ▼
operation returns context.Canceled
     │
     ▼
retry
     │
     ▼
retry
     │
     ▼
retry
```

you are doing work after the caller has explicitly asked you to stop.

A retry loop should generally respect cancellation:

```go
for i := 0; i < 5; i++ {
    err := operation(ctx)
    if err == nil {
        return nil
    }

    if ctx.Err() != nil {
        return ctx.Err()
    }

    // bounded backoff...
}
```

---

# 20. Rule: Don't Assume `context` Solves Resource Leaks

Context cancellation is only one part of resource hygiene.

You can still leak:

- goroutines
    
- timers
    
- channels
    
- connections
    
- file descriptors
    
- locks
    
- memory
    
- workers
    

For example:

```go
go func() {
    <-ctx.Done()
}()
```

This goroutine terminates correctly.

But:

```go
go func() {
    result := expensiveOperation()
    results <- result
}()
```

can leak if nobody receives:

```text
worker
  │
  ▼
results <- result
  │
  X
no receiver
```

Context design must therefore be combined with:

```text
Cancellation
+
Ownership
+
Bounded blocking
+
Resource cleanup
+
Backpressure
```

---

# 21. Context API Design Checklist

When designing a function:

```go
func Foo(ctx context.Context, ...)
```

ask:

### Lifetime

- Who owns the operation?
    
- When should it stop?
    
- What happens if the caller disappears?
    

### Cancellation

- Does the operation observe `ctx.Done()`?
    
- Do downstream operations receive the same context?
    

### Deadline

- Does this operation need its own timeout?
    
- Is the child deadline shorter than the parent?
    

### Values

- Is context value actually appropriate?
    
- Is the data request-scoped?
    
- Could it simply be an explicit parameter?
    

### Ownership

- Who creates the context?
    
- Who cancels derived contexts?
    
- Who owns goroutines started by this operation?
    

### Failure

- What happens after cancellation?
    
- Can blocked I/O unblock?
    
- Can workers become stranded?
    

### Resource lifecycle

- Are timers cancelled?
    
- Are goroutines guaranteed to exit?
    
- Are connections released?
    
- Can queues fill indefinitely?
    

---

# 22. The Production Mental Model

Think of context as a **control-flow tree**:

```text
                    Application
                        │
                 root lifecycle
                        │
                 ┌──────┴──────┐
                 │             │
              Request A     Request B
                 │
          ┌──────┴──────┐
          │             │
       Service       External API
          │
      ┌───┴────┐
      │        │
     DB      Cache
```

Cancellation flows downward:

```text
parent cancelled
       │
       ├── child cancelled
       │      ├── DB cancelled
       │      └── cache cancelled
       │
       └── sibling unaffected
```

Deadlines propagate downward:

```text
parent deadline = 10s

service deadline = min(parent, service limit)

DB deadline = min(service, DB limit)
```

So:

```text
effective deadline
    =
min(all applicable deadlines)
```

This is why context is fundamentally about **lifetime propagation**, not generic data passing.

---

# 23. The Rules Worth Memorizing

```text
1. Pass context explicitly.
2. Put context first in the parameter list.
3. Never pass nil context.
4. Propagate the caller's context downward.
5. Derive contexts only when narrowing/changing lifetime intentionally.
6. defer cancel() for derived cancellable contexts.
7. Don't use context as a parameter bag.
8. Don't use context as a DI container.
9. Don't store request contexts in long-lived structs.
10. Use context values only for narrowly scoped metadata.
11. Use private key types for context values.
12. Cancellation is cooperative.
13. Detached work needs explicit ownership and lifetime.
14. Don't retry after cancellation/deadline unless deliberately designed.
15. Treat goroutine/resource ownership as part of context design.
```

### Principal-level mental model

The most important question is not:

> **"Where should I pass `ctx`?"**

It is:

> **"Who owns this operation's lifetime, and how does that lifetime propagate across every boundary?"**

Once you think in terms of **lifetime ownership → cancellation propagation → deadline budgeting → resource cleanup**, most good `context.Context` designs become much easier to derive rather than memorize.

---

## 🔗 References
- ⬆️ Parent: [[Context Trees & Request Cancellation]]
- 📚 Module: `Concurrency & Synchronization`
