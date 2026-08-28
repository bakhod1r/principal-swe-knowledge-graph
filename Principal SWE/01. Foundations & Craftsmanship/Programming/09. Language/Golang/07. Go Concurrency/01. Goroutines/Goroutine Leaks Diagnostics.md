---
title: "Goroutine Leaks Diagnostics"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Goroutines]]"
---
# Goroutine Leak Diagnostics

A **goroutine leak** occurs when a goroutine remains alive indefinitely even though the work it was created to perform is no longer useful or reachable by the application’s intended control flow.

The important mental model is:

> **A goroutine leak is usually a lifecycle/ownership bug, not a goroutine-count problem.**

A growing goroutine count is a symptom. The root cause is usually **missing cancellation, blocked communication, forgotten cleanup, or incorrect ownership**.

---

## 1. What Problem Does It Solve?

Go makes goroutines cheap:

```go
go worker()
```

But "cheap" does not mean "free" or "automatically managed."

A goroutine that never terminates can retain:

- its stack
    
- references to heap objects
    
- channels
    
- timers/tickers
    
- mutex-related state
    
- network/file resources indirectly
    
- application state captured by closures
    

Over time, enough leaked goroutines can cause:

```text
memory growth
     ↓
scheduler overhead
     ↓
more blocked work
     ↓
resource exhaustion
     ↓
latency / instability
     ↓
process failure
```

---

# 2. Mental Model

Think of every goroutine as a **resource with a lifecycle**:

```text
Created
   ↓
Running / Waiting
   ↓
Completed
   ↓
Released
```

A healthy goroutine has a clear answer to:

> **Who is responsible for making this goroutine stop?**

A leaked goroutine often looks like:

```text
Created
   ↓
Blocked forever
   ↓
Nobody can cancel it
   ↓
Leak
```

This leads to a useful engineering rule:

> **Every goroutine should have an explicit termination condition and an owner.**

---

# 3. The Most Common Leak Patterns

## 3.1 Sending to a Channel Nobody Receives From

Classic example:

```go
func worker() {
    ch := make(chan int)

    go func() {
        ch <- 42
    }()

    return
}
```

The goroutine executes:

```go
ch <- 42
```

but there is no receiver.

Therefore:

```text
goroutine
    ↓
ch <- 42
    ↓
BLOCKED FOREVER
```

### Fix

If the result is optional, don't create an unbounded lifecycle dependency:

```go
func worker() {
    ch := make(chan int, 1)

    go func() {
        ch <- 42
    }()

    return
}
```

But be careful:

**buffering is not a universal leak fix.**

It only changes the blocking behavior.

A better design may be to return the result directly rather than spawning a goroutine.

---

# 4. Receiver Waiting Forever

Another classic:

```go
func worker(ch <-chan int) {
    for {
        value := <-ch
        process(value)
    }
}
```

If the sender disappears without closing the channel:

```text
worker
  ↓
receive
  ↓
wait forever
```

### Better

Use channel closure:

```go
func worker(ch <-chan int) {
    for value := range ch {
        process(value)
    }
}
```

Then the lifecycle becomes:

```text
sender
  ↓
send values
  ↓
close(ch)
  ↓
range terminates
  ↓
worker exits
```

The key principle:

> **The goroutine that owns a channel's sending side should generally own closing it.**

---

# 5. Context Cancellation

One of the most important production patterns:

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

Now the goroutine has two termination conditions:

```text
                    ┌── jobs closed ──→ EXIT
                    │
worker ── select ───┤
                    │
                    └── context cancelled ──→ EXIT
```

This is dramatically safer than:

```go
for job := range jobs {
    process(job)
}
```

when the worker may need to stop independently of the producer.

---

# 6. Context Is Not Magic

This does **not** guarantee cancellation:

```go
func worker(ctx context.Context) {
    <-ctx.Done()

    expensiveOperation()
}
```

The goroutine can still become stuck inside:

```go
expensiveOperation()
```

Cancellation only works if the operation cooperates.

For example:

```go
func worker(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()

    case <-someEvent():
        return nil
    }
}
```

For network operations, use APIs that accept `context.Context`.

---

# 7. Timer and Ticker Leaks

A common pattern:

```go
ticker := time.NewTicker(time.Second)

go func() {
    for {
        <-ticker.C
        doWork()
    }
}()
```

Where is the shutdown?

There isn't one.

The goroutine waits forever.

Better:

```go
ticker := time.NewTicker(time.Second)
defer ticker.Stop()

go func() {
    for {
        select {
        case <-ctx.Done():
            return

        case <-ticker.C:
            doWork()
        }
    }
}()
```

Now:

```text
ctx cancellation
       ↓
select wakes
       ↓
goroutine exits
       ↓
ticker.Stop()
```

---

# 8. HTTP Request Goroutine Leaks

This is a dangerous pattern:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    go func() {
        doSomething()
    }()

    w.WriteHeader(http.StatusOK)
}
```

The HTTP request finishes, but the goroutine may continue indefinitely.

Sometimes that's intentional.

Often it isn't.

The request's context provides a natural lifecycle boundary:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    go func(ctx context.Context) {
        doSomething(ctx)
    }(r.Context())

    w.WriteHeader(http.StatusOK)
}
```

Then:

```go
func doSomething(ctx context.Context) {
    select {
    case <-ctx.Done():
        return

    case <-workFinished():
        return
    }
}
```

But there is an architectural question here:

> **Should this work actually belong to the HTTP request lifecycle?**

If the answer is no, don't simply detach a goroutine.

Use a durable job system / queue instead.

---

# 9. The Diagnostic Tool: `runtime.NumGoroutine`

Start with:

```go
fmt.Println(runtime.NumGoroutine())
```

Useful for basic diagnostics.

Example:

```go
before := runtime.NumGoroutine()

runTest()

time.Sleep(100 * time.Millisecond)

after := runtime.NumGoroutine()

fmt.Println("before:", before)
fmt.Println("after:", after)
```

But don't treat:

```text
goroutines = 100
```

as automatically bad.

A server may legitimately have hundreds or thousands.

The important signal is:

> **Is the goroutine population returning to its expected steady state?**

---

# 10. Goroutine Profiles with `pprof`

For production diagnostics, `pprof` is much more powerful.

Import:

```go
import _ "net/http/pprof"
```

Expose the debugging endpoint appropriately:

```go
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

Then inspect:

```text
http://localhost:6060/debug/pprof/goroutine
```

Or use:

```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

You can also retrieve the goroutine dump:

```bash
curl http://localhost:6060/debug/pprof/goroutine?debug=2
```

The `debug=2` output is particularly useful because it shows stack traces.

---

# 11. What a Goroutine Leak Looks Like

Suppose you repeatedly see:

```text
goroutine 12345 [chan receive]:
main.worker(...)
    worker.go:42
```

and:

```text
goroutine 12346 [chan receive]:
main.worker(...)
    worker.go:42
```

and:

```text
goroutine 12347 [chan receive]:
main.worker(...)
    worker.go:42
```

Hundreds of identical stacks are a strong signal.

You might discover:

```go
func worker(ch <-chan Job) {
    for {
        job := <-ch
        process(job)
    }
}
```

The question becomes:

> Why is this goroutine still supposed to exist?

That is much more useful than simply asking:

> Why is goroutine count high?

---

# 12. Goroutine States Are Diagnostic Clues

A goroutine profile can show states such as:

```text
[chan receive]
[chan send]
[sleep]
[select]
[sync.Mutex.Lock]
[semacquire]
[IO wait]
```

These states provide clues.

### `[chan receive]`

Investigate:

```go
<-ch
```

Possible causes:

- sender disappeared
    
- channel never closed
    
- cancellation missing
    

---

### `[chan send]`

Investigate:

```go
ch <- value
```

Possible causes:

- receiver disappeared
    
- channel buffer full
    
- consumer stopped
    

---

### `[sync.Mutex.Lock]`

Investigate:

```go
mu.Lock()
```

Possible causes:

- deadlock
    
- lock held indefinitely
    
- slow critical section
    

This isn't necessarily a goroutine leak, but it can produce similar symptoms.

---

### `[IO wait]`

Investigate:

- network connection
    
- filesystem
    
- syscall
    
- missing timeout
    
- stuck external dependency
    

---

# 13. Stack Traces Are Usually More Valuable Than Counts

Imagine:

```text
Goroutines: 25,000
```

That's interesting.

But this is much more useful:

```text
24,800 goroutines blocked here:

worker.go:87
    <-jobs

created by startWorkers
    worker.go:42
```

Now you have:

```text
symptom
   ↓
stack aggregation
   ↓
common blocking point
   ↓
lifecycle owner
   ↓
root cause
```

This is the correct debugging workflow.

---

# 14. Goroutine Leak Testing

You can detect leaks in tests.

A simple approach:

```go
func TestWorkerDoesNotLeak(t *testing.T) {
    before := runtime.NumGoroutine()

    ctx, cancel := context.WithCancel(context.Background())

    jobs := make(chan Job)

    go worker(ctx, jobs)

    cancel()

    time.Sleep(100 * time.Millisecond)

    after := runtime.NumGoroutine()

    if after > before+1 {
        t.Fatalf("possible goroutine leak: before=%d after=%d",
            before, after)
    }
}
```

But this is fragile.

Goroutine scheduling is nondeterministic.

A better test gives the worker an explicit completion signal.

```go
func TestWorkerStops(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    jobs := make(chan Job)
    done := make(chan struct{})

    go func() {
        defer close(done)
        worker(ctx, jobs)
    }()

    cancel()

    select {
    case <-done:
        // success

    case <-time.After(time.Second):
        t.Fatal("worker did not stop")
    }
}
```

This tests the actual contract:

> **Cancellation must cause the worker to terminate.**

That's much stronger.

---

# 15. `goleak`

For Go tests, [`go.uber.org/goleak`](https://github.com/uber-go/goleak) is a useful specialized tool for detecting unexpected goroutines.

Conceptually:

```go
func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}
```

It can identify goroutines that remain after tests complete.

But don't blindly use it without understanding your goroutine ownership.

A legitimate long-lived goroutine may need to be ignored/configured.

---

# 16. Production Diagnostic Workflow

When you suspect a leak, don't immediately modify code.

Use:

```text
1. Reproduce
      ↓
2. Measure goroutine count
      ↓
3. Capture goroutine profile
      ↓
4. Group identical stack traces
      ↓
5. Identify blocking operation
      ↓
6. Find goroutine owner
      ↓
7. Determine expected lifecycle
      ↓
8. Identify missing termination condition
      ↓
9. Fix lifecycle
      ↓
10. Add regression test
      ↓
11. Verify under load
```

This is much better than:

```text
"goroutines are increasing"
        ↓
"add context"
```

---

# 17. Production Observability

A useful metric is:

```text
runtime.NumGoroutine()
```

Expose it through metrics.

But don't alert on an absolute value alone.

For example:

```text
goroutines = 500
```

could be completely normal.

Instead look for:

```text
goroutine count
    ↑
    ↑
    ↑
    ↑
never returning
```

Especially correlate it with:

```text
goroutines
memory
CPU
request rate
latency
GC
open connections
queue depth
```

A strong signal might be:

```text
request rate → stable
goroutine count → continuously increasing
heap → continuously increasing
P99 latency → increasing
```

That strongly suggests a lifecycle/resource leak.

---

# 18. Goroutine Leak vs Deadlock

They are related but different.

### Goroutine leak

A goroutine is alive when it shouldn't be.

```text
goroutine
    ↓
blocked forever
    ↓
leak
```

### Deadlock

Multiple goroutines are waiting on each other and none can progress.

```text
G1 → waits for G2
↑          ↓
└──────────┘
```

A deadlock can cause goroutines to remain forever, but not every goroutine leak is a deadlock.

---

# 19. The Strongest Design Pattern

A production worker often looks like:

```go
func worker(
    ctx context.Context,
    jobs <-chan Job,
) {
    for {
        select {
        case <-ctx.Done():
            return

        case job, ok := <-jobs:
            if !ok {
                return
            }

            if err := process(ctx, job); err != nil {
                // handle error
            }
        }
    }
}
```

Lifecycle:

```text
                 ┌── context cancelled ──→ EXIT
                 │
worker ── select ┤
                 │
                 ├── jobs closed ───────→ EXIT
                 │
                 └── job ───────────────→ process
```

This gives the goroutine **two independent shutdown paths**.

That is an important production property.

---

# 20. Goroutine Ownership

One of the most valuable design questions is:

> **Who owns this goroutine?**

Suppose:

```go
func StartWorker() {
    go worker()
}
```

Who stops it?

If the answer is:

```text
"it will stop somehow"
```

you have a lifecycle design problem.

Prefer:

```go
type Worker struct {
    cancel context.CancelFunc
    done   chan struct{}
}
```

or an `errgroup`-based lifecycle where appropriate.

For example:

```go
g, ctx := errgroup.WithContext(ctx)

g.Go(func() error {
    return worker(ctx)
})

if err := g.Wait(); err != nil {
    return err
}
```

Now the goroutine belongs to a structured concurrency scope.

---

# 21. Structured Concurrency

This is the deeper solution.

Instead of:

```go
go worker()
go worker()
go worker()
```

with no ownership relationship:

```text
Application
 ├── anonymous goroutine
 ├── anonymous goroutine
 └── anonymous goroutine
```

prefer:

```text
Application
    │
    └── lifecycle scope
         ├── worker
         ├── worker
         └── worker
```

The scope controls:

```text
start
cancel
wait
error propagation
shutdown
```

Go does not enforce structured concurrency universally, but `context`, `errgroup`, `WaitGroup`, and explicit lifecycle management allow you to build it.

---

# 22. Common Anti-Patterns

### ❌ Fire-and-forget everywhere

```go
go doWork()
```

Ask:

```text
Who owns it?
Who cancels it?
Who waits for it?
What happens on shutdown?
What happens on error?
```

---

### ❌ Infinite loop without cancellation

```go
for {
    doWork()
}
```

Prefer:

```go
for {
    select {
    case <-ctx.Done():
        return
    default:
        doWork()
    }
}
```

Although even this may be wrong if `doWork()` itself blocks.

---

### ❌ Using buffering as the only fix

```go
make(chan Result, 1)
```

This can hide a lifecycle problem rather than solve it.

---

### ❌ Ignoring context

```go
func fetch() {
    http.Get(...)
}
```

Prefer APIs with cancellation and timeouts.

---

### ❌ Using sleeps in leak tests

```go
time.Sleep(100 * time.Millisecond)
```

Sleep-based tests are timing-dependent.

Prefer explicit synchronization:

```go
done := make(chan struct{})

go func() {
    worker()
    close(done)
}()

select {
case <-done:
case <-time.After(time.Second):
    t.Fatal("worker did not exit")
}
```

---

# 23. A Principal-Level Diagnostic Question

Don't ask only:

> "Where is the goroutine stuck?"

Ask:

> **"What lifecycle invariant was violated?"**

For example:

```text
Invariant:
Every worker must terminate when its parent context is cancelled.
```

Then investigate:

```text
Does worker receive context?
        ↓
Does it select on ctx.Done()?
        ↓
Can process() block forever?
        ↓
Does process() propagate context?
        ↓
Are external I/O operations bounded by deadlines?
```

This moves debugging from symptoms to system design.

---

# 24. Production Checklist

When investigating goroutine leaks:

```text
□ Is goroutine count actually growing?
□ Is the growth sustained?
□ Capture /debug/pprof/goroutine
□ Inspect debug=2 stack traces
□ Group identical stacks
□ Identify blocking operation
□ Identify goroutine owner
□ Define expected lifecycle
□ Check context cancellation
□ Check channel ownership/closure
□ Check WaitGroup usage
□ Check mutex/cond synchronization
□ Check timers/tickers
□ Check network I/O timeouts
□ Check external dependencies
□ Check shutdown path
□ Add explicit completion signal
□ Add regression test
□ Verify under load
□ Add production monitoring
```

---

# 25. Key Takeaways

The most important concepts are:

1. **A goroutine is a resource with a lifecycle.**
    
2. **Every goroutine needs an owner.**
    
3. **Every goroutine should have a termination condition.**
    
4. **`context.Context` is a primary cancellation mechanism.**
    
5. **Channel ownership must be explicit.**
    
6. **`pprof` stack traces are more useful than raw goroutine counts.**
    
7. **Test termination behavior, not merely goroutine numbers.**
    
8. **Don't confuse high goroutine count with a leak.**
    
9. **Buffering can mask blocking but doesn't necessarily fix lifecycle problems.**
    
10. **Structured concurrency is the deeper architectural solution.**
    

The key mental model to retain is:

```text
                ┌───────────────┐
                │ Create        │
                └───────┬───────┘
                        ↓
                ┌───────────────┐
                │ Run / Wait    │
                └───────┬───────┘
                        ↓
              ┌───────────────────┐
              │ Termination event │
              │                   │
              │ ctx cancellation  │
              │ channel close     │
              │ work completed    │
              │ parent shutdown   │
              └─────────┬─────────┘
                        ↓
                ┌───────────────┐
                │ EXIT / CLEAN  │
                └───────────────┘

If termination is undefined:
                ↓
             LEAK
```

**Principal-level rule:** before writing `go foo()`, be able to answer **"who stops this goroutine, and how do I know it actually stopped?"**

---

## 🔗 References
- ⬆️ Parent: `Goroutines & Memory Lifecycle`
- 📚 Module: `Concurrency & Synchronization`
