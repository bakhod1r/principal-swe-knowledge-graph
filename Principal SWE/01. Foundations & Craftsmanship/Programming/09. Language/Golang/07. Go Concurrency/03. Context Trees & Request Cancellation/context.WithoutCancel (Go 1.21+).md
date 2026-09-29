---
title: "context.WithoutCancel (Go 1.21+)"
tags:

  - golang
  - concurrency
  - principal-swe
parent: "[[Context Trees & Request Cancellation]]"
---

# `context.WithoutCancel` (Go 1.21+)

`context.WithoutCancel` creates a child `Context` that **inherits values from its parent but is completely detached from the parent's cancellation and deadline propagation**.

This is useful when a piece of work must **outlive the request that initiated it**, while still needing contextual values such as request metadata, tenant ID, trace information, etc.

---

## 1. Problem

Normally, Go contexts form a cancellation tree:

```text
requestCtx
    │
    ├── serviceCtx
    │      │
    │      └── dbCtx
    │
    └── ...
```

If the parent is cancelled:

```text
requestCtx.Cancel()
       │
       ▼
serviceCtx cancelled
       │
       ▼
dbCtx cancelled
```

This is usually exactly what we want.

For example:

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    doWork(ctx)
}
```

If the HTTP client disconnects, the request context can be cancelled.

But sometimes we intentionally need:

```text
HTTP request
     │
     └── response work
     
     └── background audit
             │
             └── must continue after request ends
```

Using the request context directly would be incorrect:

```go
go sendAudit(r.Context()) // ❌ may be cancelled when request ends
```

This is where `WithoutCancel` exists.

---

# 2. Mental Model

Think of a normal derived context as:

```text
Parent
 ├── Values
 ├── Deadline
 └── Cancellation
        │
        ▼
      Child
```

`WithoutCancel` changes the relationship:

```text
Parent
 ├── Values ───────────────► Child
 │
 ├── Deadline ───────────── X
 │
 └── Cancellation ───────── X
```

So:

> **Value inheritance survives; cancellation inheritance does not.**

---

# 3. Basic API

```go
ctx := context.WithoutCancel(parent)
```

Its semantics are approximately:

```text
Value(key)       → parent.Value(key)
Deadline()       → no deadline
Done()            → nil
Err()             → nil
```

That last part is extremely important.

---

# 4. Example

```go
func handler(ctx context.Context) {
    backgroundCtx := context.WithoutCancel(ctx)

    go audit(backgroundCtx)
}
```

Suppose:

```text
ctx
 │
 ├── request_id = "abc-123"
 ├── deadline = 2 seconds
 └── cancellation
```

Then:

```text
backgroundCtx
 │
 ├── request_id = "abc-123"
 ├── deadline = none
 └── cancellation = none
```

Therefore:

```go
backgroundCtx.Value(requestIDKey)
```

still works.

But:

```go
<-backgroundCtx.Done()
```

will block forever because:

```go
backgroundCtx.Done() == nil
```

---

# 5. Why `Done() == nil` Matters

This is one of the most important implementation consequences.

Consider:

```go
select {
case <-ctx.Done():
    return ctx.Err()

case result := <-results:
    return result
}
```

With a normal context:

```go
ctx.Done()
```

eventually becomes readable.

With `WithoutCancel`:

```go
ctx.Done() == nil
```

A receive from a nil channel is permanently disabled.

So:

```go
select {
case <-ctx.Done():
    ...
case result := <-results:
    ...
}
```

effectively becomes:

```text
┌───────────────────────┐
│ cancellation case     │ disabled
│ result case            │ active
└───────────────────────┘
```

This can be useful, but it can also create **unbounded work** if the operation has no independent termination mechanism.

---

# 6. Deadline Is Also Removed

Suppose:

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    5*time.Second,
)
defer cancel()
```

Then:

```go
detached := context.WithoutCancel(ctx)
```

The parent has:

```text
Deadline = now + 5s
```

but:

```go
deadline, ok := detached.Deadline()
```

returns:

```text
ok == false
```

So `WithoutCancel` removes both:

```text
cancellation
deadline
```

not just cancellation.

---

# 7. Values Are Still Inherited

This is the primary reason the API exists.

For example:

```go
type contextKey string

const requestIDKey contextKey = "request-id"

ctx := context.WithValue(
    context.Background(),
    requestIDKey,
    "req-123",
)

detached := context.WithoutCancel(ctx)

fmt.Println(detached.Value(requestIDKey))
```

Output:

```text
req-123
```

So you can preserve contextual metadata while breaking lifecycle propagation.

---

# 8. Internal Implementation

Conceptually, Go implements this using a wrapper around the parent.

The important idea is:

```go
type withoutCancelCtx struct {
    c Context
}
```

Its behavior is essentially:

```go
func (c withoutCancelCtx) Deadline() (time.Time, bool) {
    return time.Time{}, false
}

func (c withoutCancelCtx) Done() <-chan struct{} {
    return nil
}

func (c withoutCancelCtx) Err() error {
    return nil
}

func (c withoutCancelCtx) Value(key any) any {
    return value(c.c, key)
}
```

The critical transformation is:

```text
Value → delegate to parent

Deadline → erase
Done → erase
Err → erase
```

This is a very small abstraction with a very specific semantic purpose.

---

# 9. `WithoutCancel` vs `Background`

A common mistake is thinking:

```go
context.WithoutCancel(ctx)
```

is equivalent to:

```go
context.Background()
```

It is not.

### `Background`

```go
ctx := context.Background()
```

has:

```text
Values      none
Cancellation none
Deadline     none
```

### `WithoutCancel`

```go
ctx := context.WithoutCancel(parent)
```

has:

```text
Values      inherited
Cancellation none
Deadline     none
```

So:

```text
Background
    │
    └── completely independent context

WithoutCancel(parent)
    │
    ├── inherits values
    └── detached lifecycle
```

---

# 10. `WithoutCancel` vs `WithCancel`

These solve opposite problems.

### `WithCancel`

```go
child, cancel := context.WithCancel(parent)
```

creates:

```text
parent
   │
   ▼
child
```

Cancellation flows:

```text
parent ──cancel──► child
```

The child can additionally be cancelled independently:

```text
parent
   │
   ▼
child ──cancel()──► cancelled
```

### `WithoutCancel`

```go
child := context.WithoutCancel(parent)
```

creates:

```text
parent

child
```

with value lookup still connected, but lifecycle propagation removed:

```text
parent ──values──► child

parent ──cancel──X child
```

---

# 11. Very Important: It Is Not a "Forever Context"

A dangerous interpretation is:

> "WithoutCancel means this operation should run forever."

No.

It means:

> **This context does not inherit cancellation or deadlines from its parent.**

You can—and often should—add your own lifecycle.

For example:

```go
detached := context.WithoutCancel(requestCtx)

ctx, cancel := context.WithTimeout(
    detached,
    30*time.Second,
)
defer cancel()

go process(ctx)
```

Now:

```text
requestCtx
    │
    │ values
    ▼
WithoutCancel
    │
    ▼
WithTimeout(30s)
    │
    ▼
background operation
```

The request can terminate:

```text
request cancelled
      X
      │
      └── background operation continues
```

but the background operation still has:

```text
30-second deadline
```

This is usually a much safer design.

---

# 12. Production Pattern

Imagine an HTTP request:

```go
func CreateOrder(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    orderID := createOrder(ctx)

    go publishAudit(ctx, orderID)

    writeResponse(w, orderID)
}
```

Potential problem:

```text
request finishes
       │
       ▼
request context cancelled
       │
       ▼
publishAudit()
       │
       └── cancelled
```

Instead:

```go
func CreateOrder(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    orderID := createOrder(ctx)

    background := context.WithoutCancel(ctx)

    go func() {
        auditCtx, cancel := context.WithTimeout(
            background,
            30*time.Second,
        )
        defer cancel()

        publishAudit(auditCtx, orderID)
    }()

    writeResponse(w, orderID)
}
```

Now the lifecycle is explicit:

```text
HTTP request
     │
     ├── request-scoped work
     │
     └── detached audit
              │
              └── max 30s
```

---

# 13. Why Not Just Use `Background()`?

You could write:

```go
auditCtx, cancel := context.WithTimeout(
    context.Background(),
    30*time.Second,
)
```

But then you lose contextual values:

```go
ctx.Value(requestIDKey)
ctx.Value(traceIDKey)
ctx.Value(tenantIDKey)
```

This can make observability and correlation harder.

With:

```go
context.WithoutCancel(ctx)
```

you preserve those values.

Therefore:

```text
Background()
    = fresh context

WithoutCancel(ctx)
    = detached context with inherited values
```

---

# 14. But There Is an Important Design Warning

`context.Value` should not be treated as a general dependency-injection mechanism.

For example, avoid:

```go
ctx = context.WithValue(ctx, "db", db)
ctx = context.WithValue(ctx, "redis", redis)
ctx = context.WithValue(ctx, "config", config)
```

Then:

```go
context.WithoutCancel(ctx)
```

would accidentally propagate all those values.

Prefer explicit dependencies:

```go
type AuditService struct {
    publisher Publisher
}
```

and use context for:

```text
request-scoped metadata
cancellation
deadlines
tracing
authentication/request information
```

rather than arbitrary application dependencies.

---

# 15. A More Subtle Problem: Stale Values

Because values survive detachment, you can accidentally preserve information whose lifecycle was supposed to end with the request.

For example:

```go
detached := context.WithoutCancel(requestCtx)
```

might preserve:

```text
user identity
authorization metadata
request metadata
trace metadata
```

This means `WithoutCancel` is a **lifecycle boundary**, not a complete context sanitization boundary.

Before using it, ask:

> Which values should legitimately survive the request?

This is especially important for:

- security-sensitive metadata
    
- credentials
    
- authorization information
    
- request-specific mutable objects
    
- large objects retained through context values
    

---

# 16. `WithoutCancel` and Goroutine Leaks

This pattern is dangerous:

```go
ctx := context.WithoutCancel(requestCtx)

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            doWork()
        }
    }
}()
```

Because:

```go
ctx.Done() == nil
```

The cancellation branch can never execute.

You have effectively created a goroutine without a termination signal.

Better:

```go
ctx := context.WithoutCancel(requestCtx)

ctx, cancel := context.WithTimeout(ctx, time.Minute)
defer cancel()

go worker(ctx)
```

Or use an explicit application lifecycle:

```text
service shutdown
      │
      ▼
worker context
      │
      ▼
worker exits
```

---

# 17. `WithoutCancel` + `AfterFunc`

Be careful when combining it with APIs whose behavior depends on cancellation.

For example, cancellation-driven mechanisms such as:

```go
context.AfterFunc(...)
```

depend on the context becoming done.

With:

```go
ctx := context.WithoutCancel(parent)
```

there is no cancellation event:

```text
Done() == nil
```

So you should not expect parent cancellation-triggered behavior to propagate through the detached context.

---

# 18. When Should You Use It?

Good candidates include:

### Audit logging

```text
request
   │
   └── audit event
```

when the audit operation must survive client disconnect.

### Asynchronous notification

```text
request
   │
   └── notification
```

if it has its own bounded lifecycle.

### Background persistence

When a small amount of post-request work must continue independently.

### Tracing / metadata propagation

When you need request metadata but not request cancellation.

---

# 19. When Should You NOT Use It?

Do not automatically use it for:

```go
go expensiveOperation(context.WithoutCancel(ctx))
```

just because cancellation is inconvenient.

Ask first:

> Should this operation really outlive the request?

If the answer is no, preserve cancellation.

For normal request work:

```go
ctx
```

is correct.

For detached work:

```go
context.WithoutCancel(ctx)
```

may be correct.

---

# 20. Better Architecture for Significant Background Work

If the operation is important enough that losing it is unacceptable, `WithoutCancel` may not be the real solution.

For example:

```text
HTTP request
     │
     ▼
"send payment confirmation"
```

Launching:

```go
go sendEmail(context.WithoutCancel(ctx))
```

creates a reliability problem.

If the process crashes:

```text
goroutine
   │
   X process dies
```

the work disappears.

A durable architecture might be:

```text
HTTP request
     │
     ▼
DB transaction
     │
     ▼
Outbox
     │
     ▼
Message broker
     │
     ▼
Worker
     │
     ▼
Email provider
```

This gives you much stronger durability semantics.

So:

> `WithoutCancel` solves **context lifecycle**, not **durable background job execution**.

That distinction is a very important production-level mental model.

---

# 21. Decision Table

|Requirement|Approach|
|---|---|
|Request-scoped operation|`ctx`|
|Child can be independently cancelled|`context.WithCancel(ctx)`|
|Child needs shorter deadline|`context.WithTimeout(ctx, ...)`|
|Child must survive parent cancellation|`context.WithoutCancel(ctx)`|
|Survive cancellation + have bounded lifetime|`WithoutCancel` → `WithTimeout`|
|Durable background job|Queue / Outbox / Worker|
|Completely independent context|`context.Background()`|
|Need inherited values|`WithoutCancel(ctx)`|

---

# 22. The Core Mental Model

The most useful way to remember `WithoutCancel` is:

```text
                 context.Context
                       │
          ┌────────────┴────────────┐
          │                         │
       Values                   Lifecycle
          │                         │
          ▼                         ▼
      inherited              deadline/cancel
                                  │
                                  X
                           WithoutCancel
```

It creates a **lifecycle boundary** while retaining **value lookup**.

The production pattern is therefore often:

```go
detached := context.WithoutCancel(parent)

bounded, cancel := context.WithTimeout(
    detached,
    30*time.Second,
)
defer cancel()
```

rather than:

```go
detached := context.WithoutCancel(parent)

// ❌ potentially unbounded
go doWork(detached)
```

### Principal-level takeaway

`context.WithoutCancel` should make you ask a deeper architectural question:

> **Is this work merely request-triggered, or is it actually independent work with its own lifecycle?**

If it is truly independent, give it an **explicit owner, deadline, shutdown mechanism, and—when durability matters—a durable queue/workflow**. `WithoutCancel` only solves the first step: **breaking cancellation inheritance**.

---

## 🔗 References
- ⬆️ Parent: [[Context Trees & Request Cancellation]]
- 📚 Module: `Concurrency & Synchronization`
