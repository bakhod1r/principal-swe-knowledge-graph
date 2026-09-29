---
title: "Context Memory Leaks & Resource Hygiene"
tags:

  - golang
  - concurrency
  - principal-swe
parent: "[[Context Trees & Request Cancellation]]"
---
# Context Memory Leaks & Resource Hygiene in Go

`context.Context` itself rarely causes a memory leak simply because a context exists. The real danger is **retaining a context tree, timers, goroutines, or values longer than their intended lifetime**.

The principal-engineering mental model is:

> **A Context is a lifetime/deadline/cancellation propagation mechanism — not a resource owner by itself.**

Resource hygiene means that every resource whose lifetime you start must have a clearly defined termination path.

---

## 1. What problem are we solving?

Consider:

```go
func handleRequest(parent context.Context) {
    ctx, cancel := context.WithCancel(parent)

    go worker(ctx)

    // forgot cancel()
}
```

If `worker` eventually exits on its own, this may not produce a permanent leak.

But if `worker` waits for:

```go
<-ctx.Done()
```

then nobody may ever cancel `ctx`.

You now have:

```text
request
  │
  ▼
context
  │
  └── goroutine ── waiting forever
```

The memory leak is therefore **not "the Context leaked"**.

It is:

```text
unreleased lifetime
        ↓
goroutine remains reachable
        ↓
goroutine stack + references remain reachable
        ↓
GC cannot reclaim them
```

---

# 2. Mental Model: Reachability

Go's GC asks essentially:

> "Can this object still be reached from a GC root?"

Suppose:

```go
type job struct {
    ctx context.Context
    buf []byte
}
```

and a goroutine retains `job`:

```go
go func() {
    <-job.ctx.Done()
    process(job.buf)
}()
```

If the goroutine never terminates:

```text
GC roots
   │
   ▼
goroutine
   │
   ▼
job
 ┌─┴────────┐
ctx         buf
 │           │
tree       100 MB
```

Even if your application no longer logically needs the request, the objects remain reachable.

That is the important connection:

> **Goroutine leaks frequently become memory leaks.**

---

# 3. Context is a tree

A simplified context hierarchy:

```text
parent
  │
  ├── child A
  │     ├── grandchild A1
  │     └── grandchild A2
  │
  └── child B
```

Cancellation normally propagates downward:

```text
parent cancel
      │
      ├── child A canceled
      │      ├── A1 canceled
      │      └── A2 canceled
      │
      └── child B canceled
```

This means a child can retain references to its descendants.

Therefore, if you create a context subtree and keep it alive unnecessarily, you may retain associated state.

---

# 4. The classic mistake: forgetting `cancel`

```go
func handler(parent context.Context) {
    ctx, cancel := context.WithTimeout(parent, 30*time.Second)

    doSomething(ctx)

    // forgot cancel
}
```

Correct:

```go
func handler(parent context.Context) {
    ctx, cancel := context.WithTimeout(parent, 30*time.Second)
    defer cancel()

    doSomething(ctx)
}
```

Why?

Because `WithTimeout` creates deadline-related machinery.

Conceptually:

```text
WithTimeout
     │
     ├── child context
     └── timer
```

If the operation finishes after 5 ms but the timeout is 30 seconds:

```text
operation finished
       │
       ▼
5 ms

timer still potentially exists
       │
       ▼
30 sec
```

Calling:

```go
cancel()
```

says:

> "This operation is finished; release cancellation-related resources now."

This is **resource hygiene**, even when there is no permanent leak.

---

# 5. `defer cancel()` is usually the right default

```go
func fetch(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    return request(ctx)
}
```

Think of it as:

```text
Acquire lifetime
      │
      ▼
WithTimeout(...)
      │
      ▼
work
      │
      ▼
defer cancel()
      │
      ▼
release lifetime
```

This follows the same general pattern as:

```go
f, err := os.Open(...)
if err != nil {
    return err
}
defer f.Close()
```

or:

```go
mu.Lock()
defer mu.Unlock()
```

or:

```go
tx, err := db.Begin(...)
defer tx.Rollback()
```

The common principle is:

> **Pair acquisition with release as close to the acquisition site as possible.**

---

# 6. Important distinction: `cancel` is not always about memory

Consider:

```go
ctx, cancel := context.WithTimeout(parent, time.Second)
defer cancel()

result, err := operation(ctx)
```

Suppose `operation` completes immediately.

The timeout didn't necessarily create a catastrophic leak.

Instead, failing to cancel means:

```text
work complete
       │
       ├── context no longer useful
       │
       └── timer may remain scheduled
```

So:

```text
forgotten cancel
       ≠
always memory leak

forgotten cancel
       →
unnecessary lifetime
       →
resource retention
       →
potential leak under larger workloads
```

This distinction is important when debugging production systems.

---

# 7. The much more dangerous problem: goroutine leaks

Consider:

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        job := <-jobs
        process(job)
    }
}
```

If `jobs` is never closed and there is no cancellation path:

```text
goroutine
    │
    ▼
receive from jobs
    │
    ▼
waiting forever
```

Better:

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return

        case job, ok := <-jobs:
            if !ok {
                return
            }

            process(job)
        }
    }
}
```

Now there are two termination paths:

```text
             ┌── ctx canceled ──→ return
goroutine ───┤
             └── jobs closed ──→ return
```

This is a fundamental production pattern.

---

# 8. Context does not magically stop goroutines

This is a very common misconception.

Wrong mental model:

```text
cancel(ctx)
   ↓
Go runtime kills goroutine
```

That is **not** how Go works.

Correct:

```text
cancel(ctx)
   ↓
ctx.Done() becomes readable
   ↓
goroutine observes cancellation
   ↓
goroutine returns
```

The goroutine must cooperate.

For example:

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return

        default:
            doWork()
        }
    }
}
```

If `doWork()` blocks forever and doesn't observe context, cancellation may not help.

---

# 9. Blocking operations must participate in cancellation

Bad:

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        job := <-jobs
        process(job)
    }
}
```

Better:

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return

        case job := <-jobs:
            process(job)
        }
    }
}
```

Even better when processing itself supports cancellation:

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return

        case job, ok := <-jobs:
            if !ok {
                return
            }

            if err := process(ctx, job); err != nil {
                return
            }
        }
    }
}
```

Now cancellation propagates through the whole execution chain:

```text
HTTP request
     │
     ▼
handler ctx
     │
     ▼
worker ctx
     │
     ▼
DB operation
     │
     ▼
network operation
```

---

# 10. Context values can accidentally retain large objects

This is another subtle issue.

Avoid:

```go
ctx = context.WithValue(ctx, "request", hugeRequest)
```

where:

```go
hugeRequest
 ├── 50 MB body
 ├── buffers
 ├── metadata
 └── other references
```

Now anything retaining the context may indirectly retain that entire object graph.

The context isn't necessarily the root problem.

The problem is:

```text
context
   │
   ▼
large value
   │
   ├── buffer
   ├── objects
   └── references
```

### Better

Use context values for **small request-scoped metadata**, such as:

```text
request ID
trace ID
auth metadata
locale
```

Not:

```text
large request payload
database connection
response object
business entity graph
cache
application state
```

---

# 11. Context should not become a dependency container

This is an architectural smell:

```go
ctx = context.WithValue(ctx, dbKey, db)
ctx = context.WithValue(ctx, cacheKey, cache)
ctx = context.WithValue(ctx, loggerKey, logger)
ctx = context.WithValue(ctx, configKey, config)
```

Now:

```text
Context
 ├── DB
 ├── Cache
 ├── Logger
 ├── Config
 └── ...
```

This makes ownership unclear.

Prefer explicit dependencies:

```go
type Service struct {
    db    *sql.DB
    cache *Cache
    log   *slog.Logger
}
```

And:

```go
func (s *Service) Get(ctx context.Context, id string) error {
    ...
}
```

Context carries **request-scoped execution metadata**, not your application's object graph.

---

# 12. `context.Background()` can break cancellation trees

Suppose:

```go
func handler(ctx context.Context) {
    go backgroundTask(context.Background())
}
```

You've now disconnected:

```text
request ctx
     │
     X
     │
backgroundTask
```

If the request is canceled:

```text
client disconnects
       ↓
request canceled
       ↓
handler stops
```

but:

```text
backgroundTask(context.Background())
       ↓
continues running
```

That may be intentional.

But if it isn't intentional, you've created a lifetime bug.

Prefer:

```go
go backgroundTask(ctx)
```

when the task should die with the request.

Or deliberately detach when you actually want independent lifetime:

```go
bgCtx := context.WithoutCancel(ctx)
```

But remember:

> `WithoutCancel` removes cancellation/deadline propagation; it does not magically establish ownership or cleanup.

---

# 13. A subtle trap: spawning background work from request handlers

Consider:

```go
func handler(ctx context.Context) {
    go expensiveJob(ctx)

    return
}
```

Questions a Principal Engineer should immediately ask:

1. Who owns `expensiveJob`?
    
2. Should it survive request cancellation?
    
3. What is its maximum lifetime?
    
4. What happens during server shutdown?
    
5. How many jobs can exist concurrently?
    
6. What happens under traffic spikes?
    
7. Where is backpressure?
    
8. What happens if the process crashes?
    
9. Is the work durable?
    
10. How is failure observed?
    

The `go` keyword creates a **new concurrency lifetime**.

That means you have created an ownership problem.

---

# 14. Unbounded goroutine creation is a resource leak under load

This looks innocent:

```go
for _, job := range jobs {
    go process(job)
}
```

At:

```text
100 jobs
```

fine.

At:

```text
1,000,000 jobs
```

you have potentially created:

```text
1,000,000 goroutines
```

Even though goroutines are lightweight, they aren't free.

You can exhaust:

```text
memory
scheduler capacity
FDs
DB connections
network connections
downstream capacity
```

Use bounded concurrency:

```go
sem := make(chan struct{}, 100)

for _, job := range jobs {
    sem <- struct{}{}

    go func(job Job) {
        defer func() {
            <-sem
        }()

        process(job)
    }(job)
}
```

But even this requires cancellation and lifecycle handling.

Production systems often benefit from an explicit worker pool when work is continuous.

---

# 15. Resource hygiene applies beyond memory

Think in terms of:

```text
Resource
   │
   ├── memory
   ├── goroutine
   ├── timer
   ├── file descriptor
   ├── socket
   ├── DB connection
   ├── lock
   ├── ticker
   ├── subscription
   └── external lease
```

Every one needs an owner and termination mechanism.

For example:

### Timer

```go
timer := time.NewTimer(...)
defer timer.Stop()
```

### Ticker

```go
ticker := time.NewTicker(...)
defer ticker.Stop()
```

### File

```go
f, err := os.Open(...)
if err != nil {
    return err
}
defer f.Close()
```

### DB rows

```go
rows, err := db.QueryContext(ctx, query)
if err != nil {
    return err
}
defer rows.Close()
```

### HTTP response body

```go
resp, err := client.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()
```

### Context

```go
ctx, cancel := context.WithTimeout(...)
defer cancel()
```

Same engineering principle:

> **Acquire → use → release.**

---

# 16. Context cancellation is a lifetime boundary

A useful mental model:

```text
Context lifetime
───────────────────────────────────>

request starts
     │
     ▼
  ctx created
     │
     ├── DB
     ├── HTTP
     ├── goroutine
     └── timer
     │
     ▼
request finished
     │
     ▼
cancel()
     │
     ├── stop timer
     ├── signal goroutines
     ├── abort context-aware I/O
     └── release references when no longer reachable
```

This makes context cancellation similar to a **structured lifetime boundary**.

---

# 17. Structured concurrency

Go does not enforce structured concurrency everywhere, but you can design toward it.

Bad:

```go
func process(ctx context.Context) {
    go task1(ctx)
    go task2(ctx)
    go task3(ctx)

    return
}
```

The parent returns while children remain alive.

Better:

```go
func process(ctx context.Context) error {
    g, ctx := errgroup.WithContext(ctx)

    g.Go(func() error {
        return task1(ctx)
    })

    g.Go(func() error {
        return task2(ctx)
    })

    g.Go(func() error {
        return task3(ctx)
    })

    return g.Wait()
}
```

Conceptually:

```text
parent operation
       │
 ┌─────┼─────┐
 ▼     ▼     ▼
T1    T2    T3
 │     │     │
 └─────┼─────┘
       ▼
     Wait
       │
       ▼
 parent returns
```

The lifetime of child work is bounded by the parent operation.

This dramatically reduces lifecycle bugs.

---

# 18. Detecting Context-related leaks

Don't start with:

> "Maybe Context is leaking."

Start with:

```text
Symptom
   ↓
Observe
   ↓
Find retained resource
   ↓
Find owner
   ↓
Find missing termination path
```

Useful signals:

### Goroutine count

```go
runtime.NumGoroutine()
```

But this is only a symptom.

Use pprof for real diagnosis.

Look at:

```text
goroutine profile
heap profile
alloc_objects
alloc_space
inuse_objects
inuse_space
```

Typical leak pattern:

```text
goroutine count:
100
120
140
180
250
400
...
```

especially if traffic returns to baseline but goroutines don't.

---

# 19. Heap profiling vs goroutine profiling

This distinction matters.

### Goroutine leak

Look at:

```text
goroutine profile
```

You might find:

```text
goroutine 12345 [chan receive]
```

or:

```text
goroutine 12346 [select]
```

repeated thousands of times.

### Memory retention

Look at:

```text
heap profile
```

You may discover:

```text
large []byte
request objects
context values
buffers
```

being retained.

The root cause might be:

```text
goroutine leak
     ↓
goroutine retains request
     ↓
request retains 10 MB buffer
     ↓
heap grows
```

So:

> **Heap growth can be a consequence of a goroutine leak rather than an allocator problem.**

---

# 20. A production checklist

When creating a context-derived operation, ask:

### Lifetime

```text
Who owns this context?
Who cancels it?
When does it become irrelevant?
```

### Goroutines

```text
Who starts this goroutine?
Who stops it?
What happens on cancellation?
What happens on shutdown?
```

### Timers

```text
Who owns the timer?
Is it stopped when unnecessary?
```

### I/O

```text
Does the operation accept context?
What happens when ctx.Done() fires?
```

### Memory

```text
Does the context contain large values?
Can a long-lived object retain the context?
```

### Concurrency

```text
Can this operation spawn unbounded goroutines?
Is concurrency bounded?
Where is backpressure?
```

### Shutdown

```text
Can the service terminate cleanly?
Are workers joined?
Are resources closed?
```

---

# 21. The strongest mental model

Don't think:

> "I need to prevent Context memory leaks."

Think:

> **"I need explicit ownership and bounded lifetime for every resource and concurrent activity."**

Then Context becomes one mechanism used to express that lifetime.

A production-quality system should have:

```text
                 Ownership
                    │
                    ▼
               Lifetime
                    │
          ┌─────────┼─────────┐
          ▼         ▼         ▼
       Context   Goroutine   Resource
          │         │         │
          ▼         ▼         ▼
      Cancel()    return()   Close()
```

And under failure:

```text
Failure
   │
   ▼
Cancellation
   │
   ▼
Workers stop
   │
   ▼
I/O aborts
   │
   ▼
Resources release
   │
   ▼
Memory becomes unreachable
   │
   ▼
GC reclaims it
```

That is the **Context + resource hygiene** mental model worth carrying into production Go systems.

### Key takeaways

1. **A Context itself is usually not the leak.**
    
2. **Forgotten cancellation can retain timers/context trees unnecessarily.**
    
3. **Goroutine leaks are a major source of memory retention.**
    
4. **Context cancellation is cooperative; goroutines must observe `Done()`.**
    
5. **Avoid storing large objects or application dependencies in context values.**
    
6. **Every goroutine needs an explicit owner and termination path.**
    
7. **Every resource needs Acquire → Use → Release semantics.**
    
8. **Prefer structured lifetimes where child work cannot outlive its parent accidentally.**
    
9. **Use pprof and runtime metrics to prove leaks rather than guessing.**
    
10. **At Principal level, reason about ownership and lifetime—not merely API calls.**

---

## 🔗 References
- ⬆️ Parent: [[Context Trees & Request Cancellation]]
- 📚 Module: `Concurrency & Synchronization`
