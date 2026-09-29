---
title: "context.Background() vs context.TODO()"
tags:

  - golang
  - concurrency
  - principal-swe
parent: "[[Context Trees & Request Cancellation]]"
---
# `context.Background()` vs `context.TODO()` in Go

Both `context.Background()` and `context.TODO()` return **empty, non-cancelable root contexts**. Their runtime behavior is essentially identical, but their **semantic intent is different**.

The important distinction is not performance or cancellation—it is **communication of developer intent**.

---

## 1. The Core Idea

Go's `context.Context` represents request-scoped execution state:

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

A context can carry:

- cancellation
    
- deadlines
    
- request-scoped values
    

`Background()` and `TODO()` are both **root contexts**.

```text
                    Context tree
                         │
              ┌──────────┴──────────┐
              │                     │
       Background()              TODO()
              │                     │
        WithCancel()          WithTimeout()
              │                     │
          request ctx          operation ctx
```

---

# 2. `context.Background()`

Use `Background()` when you **intentionally want a root context**.

```go
ctx := context.Background()
```

The semantic meaning is:

> "There is no parent context, and that is intentional."

Typical places:

### Application entry point

```go
func main() {
    ctx := context.Background()

    app, err := NewApp(ctx)
    if err != nil {
        log.Fatal(err)
    }

    app.Run(ctx)
}
```

### Tests

```go
func TestUserService(t *testing.T) {
    ctx := context.Background()

    user, err := service.GetUser(ctx, 42)
    // ...
}
```

### Creating a root application context

```go
func main() {
    ctx := context.Background()

    ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
    defer stop()

    run(ctx)
}
```

Here `Background()` establishes the root of the application's context tree.

---

# 3. `context.TODO()`

`TODO()` is primarily a **semantic marker**.

```go
ctx := context.TODO()
```

It means approximately:

> "I need a context here, but I haven't decided what the correct context should be yet."

This is useful during:

- incremental migration
    
- refactoring
    
- temporarily incomplete APIs
    
- code where the correct parent context is not yet known
    

Example:

```go
func processUser(user User) error {
    ctx := context.TODO()

    return saveUser(ctx, user)
}
```

This compiles and works, but it communicates:

> "The context design here is unfinished."

---

# 4. Their Runtime Behavior

This is the critical technical point:

```go
context.Background()
context.TODO()
```

are both implementations of an empty context.

Conceptually:

```text
Background()
    │
    ├── Deadline() → no deadline
    ├── Done()     → nil
    ├── Err()      → nil
    └── Value()    → nil

TODO()
    │
    ├── Deadline() → no deadline
    ├── Done()     → nil
    ├── Err()      → nil
    └── Value()    → nil
```

So this is **not** a performance distinction:

```go
ctx1 := context.Background()
ctx2 := context.TODO()
```

You should not choose between them based on:

- performance
    
- memory
    
- cancellation behavior
    
- goroutine behavior
    
- scheduling
    
- allocation behavior
    

The distinction is primarily **intent**.

---

# 5. Why Does `TODO()` Exist?

Consider a large codebase being migrated to use contexts.

Before:

```go
func ProcessOrder(order Order) error {
    return saveOrder(order)
}
```

You discover:

```go
func saveOrder(order Order) error {
    // ...
}
```

should become:

```go
func saveOrder(ctx context.Context, order Order) error {
    // ...
}
```

But many call sites have not yet been migrated.

You might temporarily write:

```go
func ProcessOrder(order Order) error {
    return saveOrder(context.TODO(), order)
}
```

This tells future maintainers:

```text
"This context is temporary.
We should eventually propagate the real context."
```

Later:

```go
func ProcessOrder(ctx context.Context, order Order) error {
    return saveOrder(ctx, order)
}
```

Now the `TODO()` disappears.

---

# 6. A Good Mental Model

Think of them as:

```text
Background()
    = "I intentionally have no parent."

TODO()
    = "I don't know what the parent should be yet."
```

This is much more useful than thinking:

```text
Background = production
TODO       = testing
```

That is **incorrect**.

`TODO()` is not a testing context.

---

# 7. `TODO()` Is Not a Fallback for Missing Contexts

This is one of the most important engineering rules.

Suppose you have:

```go
func GetUser(ctx context.Context, id int64) (*User, error)
```

and another function:

```go
func HandleRequest(ctx context.Context, req Request) error {
    return service.GetUser(ctx, req.UserID)
}
```

Do **not** do this:

```go
func HandleRequest(ctx context.Context, req Request) error {
    return service.GetUser(context.TODO(), req.UserID)
}
```

You just destroyed the context propagation chain.

The correct approach is:

```go
func HandleRequest(ctx context.Context, req Request) error {
    return service.GetUser(ctx, req.UserID)
}
```

Because the incoming context may contain:

```text
HTTP request
    │
    ▼
request context
    │
    ├── cancellation
    ├── deadline
    ├── tracing
    └── request-scoped metadata
          │
          ▼
      service
          │
          ▼
       database
```

Replacing it with:

```go
context.TODO()
```

creates:

```text
HTTP request
    │
    ▼
service
    │
    X   context propagation broken
    │
    ▼
TODO()
```

This can cause real production problems.

---

# 8. Example: Lost Cancellation

Suppose an HTTP request is canceled.

Correct:

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    err := service.Process(ctx)
    // ...
}
```

Then:

```text
client disconnects
       │
       ▼
request context canceled
       │
       ▼
service canceled
       │
       ▼
DB query canceled
```

But this is wrong:

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    err := service.Process(context.TODO())
}
```

Now:

```text
client disconnects
       │
       ▼
request context canceled

             X

service continues running
       │
       ▼
database query continues
```

That can waste:

- database connections
    
- CPU
    
- goroutines
    
- network resources
    
- downstream capacity
    

At scale, this becomes a reliability problem.

---

# 9. `Background()` Is Also Not a Universal Default

This is another common mistake:

```go
func Process(ctx context.Context) error {
    if ctx == nil {
        ctx = context.Background()
    }

    // ...
}
```

Usually, don't design APIs around silently replacing a missing context.

Better:

```go
func Process(ctx context.Context) error {
    // ...
}
```

and require callers to provide the correct context.

If you're at an application boundary where you genuinely need a root:

```go
func main() {
    ctx := context.Background()

    run(ctx)
}
```

That's appropriate.

---

# 10. Never Pass `nil` as Context

Avoid:

```go
service.Process(nil)
```

Go's context conventions strongly favor a non-nil context.

Instead:

```go
service.Process(context.Background())
```

if you truly have no parent context.

During incomplete migration:

```go
service.Process(context.TODO())
```

may be appropriate.

But `TODO()` should eventually disappear when the correct context becomes available.

---

# 11. Library Code vs Application Code

A useful rule:

### Application root

```go
func main() {
    ctx := context.Background()
    run(ctx)
}
```

Good.

### HTTP handler

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    service.Process(ctx)
}
```

Good.

### gRPC handler

```go
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) {
    s.service.GetUser(ctx, req.Id)
}
```

Good.

### Internal service

```go
func (s *Service) GetUser(ctx context.Context, id int64) (*User, error) {
    return s.repo.GetUser(ctx, id)
}
```

Good.

### Unknown/incomplete context

```go
func legacyFunction() error {
    return newFunction(context.TODO())
}
```

Potentially acceptable **temporarily**.

---

# 12. Don't Store Context in Structs

This is related to the same design philosophy.

Avoid:

```go
type Service struct {
    ctx context.Context
}
```

and:

```go
service := &Service{
    ctx: context.Background(),
}
```

Instead:

```go
type Service struct {
    repo Repository
}
```

and:

```go
func (s *Service) Process(ctx context.Context) error {
    return s.repo.Process(ctx)
}
```

Why?

Because context represents the lifetime of an **operation**, not the lifetime of an object.

Think:

```text
Service lifetime
──────────────────────────────────────>

Request A       Request B       Request C
─────────       ─────────       ─────────
ctx A           ctx B           ctx C
```

A single service may handle thousands of requests.

Therefore the context belongs to the operation.

---

# 13. `context.TODO()` as a Code-Review Signal

In a production codebase, I would treat this:

```go
context.TODO()
```

as a question:

> "Why don't we have the real context here?"

Not necessarily as a bug.

For example:

```go
func migrateLegacyData() error {
    return repository.Run(context.TODO())
}
```

might be perfectly reasonable if this is a standalone process.

But:

```go
func HandleHTTP(ctx context.Context) error {
    return repository.Run(context.TODO())
}
```

is suspicious.

The same line has different engineering meaning depending on its position in the context tree.

---

# 14. Decision Table

|Situation|Use|
|---|---|
|Application root|`context.Background()`|
|Intentionally creating root context|`context.Background()`|
|Test with no parent context|`context.Background()`|
|Temporary missing context during migration|`context.TODO()`|
|Unsure what parent context should be|`context.TODO()`|
|HTTP request|`r.Context()`|
|gRPC request|provided `ctx`|
|Service method|receive `ctx`|
|Repository method|receive `ctx`|
|Background worker with lifecycle|explicit application/worker context|
|Context propagation|**never replace with `TODO()`**|

---

# 15. Production Pattern

A healthy Go application usually looks like:

```text
main
 │
 │ context.Background()
 ▼
application context
 │
 ├── HTTP request
 │      │
 │      ▼
 │   handler(ctx)
 │      │
 │      ▼
 │   service(ctx)
 │      │
 │      ▼
 │   repository(ctx)
 │
 └── background worker
        │
        ▼
      worker(ctx)
```

The context should generally flow **top-down**.

```go
func main() {
    root := context.Background()

    ctx, cancel := signal.NotifyContext(
        root,
        os.Interrupt,
    )
    defer cancel()

    app.Run(ctx)
}
```

Then:

```go
func (a *App) Run(ctx context.Context) error {
    return a.server.Start(ctx)
}
```

Then:

```go
func (s *Service) Process(ctx context.Context) error {
    return s.repo.Save(ctx)
}
```

This gives you a coherent cancellation and deadline propagation model.

---

# 16. Principal-Level Insight

The important distinction is not:

> "`Background()` vs `TODO()` — which one is faster?"

It is:

> **"Where is the ownership boundary of this operation's lifetime?"**

A context should normally originate at a meaningful **operation boundary**:

```text
Process startup
       │
       ▼
    root ctx
       │
       ▼
 incoming request
       │
       ▼
   application
       │
       ▼
     service
       │
       ▼
   repository
```

If you find `context.TODO()` deep inside this chain, ask:

1. Why isn't the caller's context being propagated?
    
2. Is cancellation being lost?
    
3. Is the deadline being lost?
    
4. Is tracing propagation being lost?
    
5. Is this intentionally independent work?
    
6. If independent, should it have its **own lifecycle context** instead?
    

That is the real engineering question.

## Rule of thumb

```go
context.Background()
```

means:

> **"I am intentionally starting a new context tree."**

```go
context.TODO()
```

means:

> **"The correct context design has not been decided/implemented yet."**

So in production application code, prefer **`Background()` at true roots and propagated contexts everywhere else**. Treat persistent `TODO()` usage as technical debt unless there is a clear reason for it.

---

## 🔗 References
- ⬆️ Parent: [[Context Trees & Request Cancellation]]
- 📚 Module: `Concurrency & Synchronization`
