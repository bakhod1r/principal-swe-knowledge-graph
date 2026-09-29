---
title: "context.AfterFunc (Go 1.21+)"
tags:

  - golang
  - concurrency
  - principal-swe
parent: "[[Context Trees & Request Cancellation]]"
---
# `context.AfterFunc` (Go 1.21+)

`context.AfterFunc` is a Go 1.21+ API that registers a function to run **asynchronously when a `Context` is canceled**.

It is best understood as:

> **“When this context becomes done, schedule this callback.”**

It is particularly useful when cancellation needs to trigger **cleanup, rollback, notification, or interruption of an external operation**.

---

## 1. The problem it solves

Before `context.AfterFunc`, you commonly wrote:

```go
go func() {
    <-ctx.Done()
    cleanup()
}()
```

This works, but creates a goroutine whose only job is to wait for cancellation.

`AfterFunc` gives the runtime a direct cancellation callback:

```go
stop := context.AfterFunc(ctx, cleanup)
```

Conceptually:

```text
Context
   │
   │ Cancel()
   ▼
Done channel becomes closed
   │
   ▼
AfterFunc callback scheduled
   │
   ▼
cleanup()
```

The important distinction:

**`AfterFunc` does not run `cleanup` synchronously inside `cancel()`.**

The callback runs in its **own goroutine**.

---

# 2. API

The signature is:

```go
func AfterFunc(ctx Context, f func()) (stop func() bool)
```

Example:

```go
stop := context.AfterFunc(ctx, func() {
    fmt.Println("context canceled")
})
```

You receive a `stop` function.

```go
stopped := stop()
```

The return value tells you whether the callback was successfully prevented from running.

---

# 3. Mental model

Think of:

```go
context.AfterFunc(ctx, f)
```

as registering a **one-shot cancellation hook**:

```text
             Context
                │
          cancellation
                │
                ▼
        ┌───────────────┐
        │ callback f()  │
        └───────────────┘
                │
                ▼
        runs exactly once
```

There are three important states:

```text
                AfterFunc
                   │
        ┌──────────┴──────────┐
        │                     │
   ctx canceled           stop() called
        │                     │
        ▼                     ▼
     f runs             f prevented
```

But there is an important race:

```text
                 stop()
                   │
                   ▼
             ┌───────────┐
             │ callback? │
             └─────┬─────┘
                   │
             race with cancel
```

Therefore `stop()` returning `false` means:

> You did **not** successfully prevent the callback from running.

It does **not** mean:

> “The callback has definitely finished.”

That distinction is extremely important.

---

# 4. Basic example

```go
ctx, cancel := context.WithCancel(context.Background())

stop := context.AfterFunc(ctx, func() {
    fmt.Println("cleanup")
})

cancel()

_ = stop
```

Output:

```text
cleanup
```

But because the callback executes asynchronously, this is **not** guaranteed:

```go
cancel()

fmt.Println("done")
```

to produce:

```text
cleanup
done
```

It could be:

```text
done
cleanup
```

because:

```text
cancel()
   │
   ├── marks ctx canceled
   │
   └── schedules callback goroutine
                         │
                         ▼
                      f()
```

---

# 5. `stop()` is not `wait()`

This is one of the most important details.

Consider:

```go
stop := context.AfterFunc(ctx, func() {
    cleanup()
})

if stop() {
    fmt.Println("callback prevented")
}
```

If `stop()` returns `true`:

```text
callback has not started
        │
        ▼
stop() prevents it
```

If it returns `false`:

```text
callback may already be running
        OR
callback has already completed
```

Therefore:

```go
stop()
```

does **not** synchronize with `f()`.

If you need synchronization, you must explicitly provide it.

For example:

```go
var wg sync.WaitGroup

wg.Add(1)

stop := context.AfterFunc(ctx, func() {
    defer wg.Done()
    cleanup()
})

if !stop() {
    wg.Wait()
}
```

Now you have:

```text
stop()
  │
  ├── true  → callback never started
  │
  └── false → wait until callback completes
```

This pattern is useful when callback completion matters.

---

# 6. Why is `AfterFunc` asynchronous?

Imagine:

```go
ctx, cancel := context.WithCancel(parent)

context.AfterFunc(ctx, func() {
    expensiveCleanup()
})

cancel()
```

If `cancel()` waited for:

```go
expensiveCleanup()
```

then cancellation could become unexpectedly expensive.

Instead:

```text
cancel()
 │
 ├── mark context canceled
 │
 └── callback runs asynchronously
```

This keeps cancellation propagation from being blocked by arbitrary callback work.

### Principal-level insight

A cancellation operation should generally not depend on the latency of user-provided cleanup code.

`AfterFunc` therefore creates a **failure/latency boundary** between:

```text
cancellation propagation
```

and:

```text
callback execution
```

---

# 7. Callback executes only once

The callback is one-shot.

```go
stop := context.AfterFunc(ctx, func() {
    fmt.Println("called")
})
```

Even if cancellation is triggered multiple times:

```go
cancel()
cancel()
cancel()
```

the callback is not repeatedly invoked.

Why?

Context cancellation itself is one-way:

```text
ACTIVE
  │
  │ cancel
  ▼
CANCELED
```

There is no:

```text
CANCELED → ACTIVE
```

transition.

Therefore the callback corresponds to the transition into the canceled state.

---

# 8. Already-canceled context

This is an interesting case:

```go
ctx, cancel := context.WithCancel(context.Background())
cancel()

context.AfterFunc(ctx, func() {
    fmt.Println("cleanup")
})
```

The callback is still scheduled asynchronously.

Conceptually:

```text
ctx already canceled
       │
       ▼
AfterFunc registers callback
       │
       ▼
callback scheduled asynchronously
```

It does **not** execute inline during the `AfterFunc` call.

---

# 9. `AfterFunc` vs goroutine waiting on `Done()`

Traditional approach:

```go
go func() {
    <-ctx.Done()
    cleanup()
}()
```

New approach:

```go
context.AfterFunc(ctx, cleanup)
```

The latter expresses the intent more directly:

```text
Context cancellation
        ↓
registered callback
```

rather than:

```text
spawn goroutine
        ↓
wait forever
        ↓
context cancellation
        ↓
cleanup
```

This is particularly valuable when the callback represents a lifecycle action.

---

# 10. The `stop()` pattern

A common pattern is:

```go
stop := context.AfterFunc(ctx, cleanup)

defer stop()
```

For example:

```go
func operation(ctx context.Context) error {
    rollback := func() {
        // rollback operation
    }

    stop := context.AfterFunc(ctx, rollback)
    defer stop()

    if err := doWork(ctx); err != nil {
        return err
    }

    commit()

    return nil
}
```

The idea:

```text
operation starts
      │
      ├── context canceled
      │       │
      │       ▼
      │    rollback
      │
      └── operation succeeds
              │
              ▼
          stop rollback
```

This can be an elegant way to associate cancellation with cleanup.

But there is a subtle race.

---

# 11. The cancellation/commit race

Consider:

```go
stop := context.AfterFunc(ctx, rollback)

doWork()

commit()

stop()
```

Potential execution:

```text
goroutine A                  callback goroutine

doWork()
   │
   ▼
commit()
   │
   │ cancel happens
   ├────────────────────────────► rollback()
   │
stop()
```

Now you potentially have:

```text
commit()
rollback()
```

executing concurrently.

This can be disastrous if the operations are not designed for concurrency.

### Therefore:

`AfterFunc` does **not** automatically make cleanup transactional.

You must reason about:

- synchronization
    
- state transitions
    
- idempotency
    
- commit semantics
    
- cancellation races
    

---

# 12. A safer state-machine model

Instead of thinking:

```text
cancel → rollback
```

think:

```text
             ┌───────────┐
             │ RUNNING   │
             └─────┬─────┘
                   │
          ┌────────┴────────┐
          │                 │
       success           cancel
          │                 │
          ▼                 ▼
      COMMITTED         CANCELED
          │                 │
          └───────┬─────────┘
                  │
             terminal state
```

Your implementation should make the terminal state explicit.

For complicated workflows, use synchronization:

```go
mu.Lock()

switch state {
case running:
    state = committed
}

mu.Unlock()
```

and make rollback check the state before acting.

---

# 13. Example: rollback registration

Suppose an operation allocates a resource:

```go
func createResource(ctx context.Context) error {
    resource := allocate()

    stop := context.AfterFunc(ctx, func() {
        resource.Close()
    })

    if err := initialize(resource); err != nil {
        stop()
        resource.Close()
        return err
    }

    return nil
}
```

There is a problem here:

```text
initialize()
      │
      │ context canceled
      ▼
Close()
      │
      │
initialize() still running?
```

You need to understand whether the underlying resource permits concurrent:

```text
initialize()
Close()
```

If not, synchronization is required.

This is exactly why `AfterFunc` should be treated as a **concurrency primitive**, not merely a convenient callback API.

---

# 14. `AfterFunc` with `sync.Cond`

One particularly interesting use case is cancellation with `sync.Cond`.

Suppose a goroutine is waiting:

```go
cond.Wait()
```

A context cancellation cannot directly wake a `sync.Cond`.

`AfterFunc` can bridge the two:

```go
stop := context.AfterFunc(ctx, func() {
    cond.Broadcast()
})
defer stop()
```

Now:

```text
context cancellation
       │
       ▼
AfterFunc
       │
       ▼
cond.Broadcast()
       │
       ▼
waiting goroutines wake
```

The waiting code can then inspect:

```go
ctx.Err()
```

to determine whether cancellation caused the wake-up.

This is a powerful pattern because it connects two different synchronization mechanisms:

```text
Context
   │
   ▼
AfterFunc
   │
   ▼
sync.Cond
```

---

# 15. `AfterFunc` with timers

You might initially think:

```go
context.AfterFunc(ctx, f)
```

is similar to:

```go
time.AfterFunc(duration, f)
```

They are conceptually related but triggered by different events.

### `time.AfterFunc`

```text
time elapsed
     │
     ▼
callback
```

### `context.AfterFunc`

```text
context canceled
     │
     ▼
callback
```

So:

```go
time.AfterFunc(5*time.Second, f)
```

means:

> Execute after 5 seconds.

Whereas:

```go
context.AfterFunc(ctx, f)
```

means:

> Execute when `ctx` is canceled.

---

# 16. Relationship with `context.WithTimeout`

This becomes especially useful with:

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    5*time.Second,
)

defer cancel()

context.AfterFunc(ctx, func() {
    cleanup()
})
```

Now:

```text
5-second timer
      │
      ▼
context cancellation
      │
      ▼
AfterFunc callback
      │
      ▼
cleanup()
```

So `AfterFunc` doesn't implement the timer itself.

The context does.

`AfterFunc` simply attaches behavior to the cancellation event.

---

# 17. Context tree interaction

Consider:

```text
parent
  │
  ├── child A
  │
  │     └── AfterFunc
  │
  └── child B
```

When:

```go
cancel(parent)
```

cancellation propagates:

```text
parent
  │
  ├── child A ──► AfterFunc callback
  │
  └── child B
```

The callback therefore follows the cancellation semantics of the particular context it is registered against.

This makes it useful for lifecycle-scoped resources.

---

# 18. Don't use it as a general event system

Bad design:

```go
context.AfterFunc(ctx, func() {
    sendMetrics()
})

context.AfterFunc(ctx, func() {
    notifyUsers()
})

context.AfterFunc(ctx, func() {
    flushCache()
})

context.AfterFunc(ctx, func() {
    writeAuditLog()
})
```

This turns cancellation into an implicit event bus.

That makes the lifecycle difficult to reason about.

A context should primarily represent:

```text
request lifetime
cancellation
deadline
request-scoped values
```

not arbitrary application events.

---

# 19. Don't put critical business logic blindly inside it

Avoid:

```go
context.AfterFunc(ctx, func() {
    chargeCreditCard()
})
```

Cancellation is generally an unsuitable trigger for an irreversible business operation.

Cancellation means:

> The caller no longer wants to continue this operation.

It does not mean:

> Execute arbitrary business logic.

`AfterFunc` is much better suited to:

- cleanup
    
- releasing resources
    
- interrupting waits
    
- rollback mechanisms
    
- cancellation-specific signaling
    

---

# 20. Callback must be cancellation-safe

Because the callback runs asynchronously, it needs the same engineering discipline as any goroutine.

Think about:

### Data races

```go
var state State

context.AfterFunc(ctx, func() {
    state = failed
})
```

while another goroutine reads `state`.

Potential race.

Use:

```go
sync.Mutex
```

or an atomic primitive where appropriate.

---

### Panic

A panic in the callback occurs in its goroutine.

Therefore don't assume the caller's goroutine can recover it:

```go
defer recover()
```

in the caller does not automatically protect another goroutine.

---

### Blocking

Avoid callbacks that can block indefinitely:

```go
context.AfterFunc(ctx, func() {
    <-someChannel
})
```

You've essentially created a goroutine leak after cancellation.

---

# 21. `AfterFunc` and goroutine lifecycle

A useful mental model:

```text
AfterFunc registration
       │
       ▼
callback waiting for cancellation
       │
       │
       ├── stop() succeeds
       │      └── no callback execution
       │
       └── cancellation
              │
              ▼
         callback goroutine
              │
              ▼
           callback
```

The important lifecycle question is:

> What happens if cancellation never occurs?

The callback simply doesn't execute.

Therefore, if `f` captures resources, make sure registration itself doesn't accidentally extend their lifetime through large object graphs.

---

# 22. Memory/lifetime consideration

Suppose:

```go
large := make([]byte, 100<<20)

context.AfterFunc(ctx, func() {
    use(large)
})
```

The callback closure references `large`.

Therefore the registered callback can keep referenced objects reachable for as long as the context/callback registration remains relevant.

This matters for long-lived contexts.

### Rule

Don't register cancellation callbacks on contexts whose lifetime is much longer than the resource they reference unless that lifetime relationship is intentional.

---

# 23. `AfterFunc` vs `defer`

These solve fundamentally different problems.

### `defer`

```go
defer cleanup()
```

means:

> Cleanup when this function returns.

### `AfterFunc`

```go
context.AfterFunc(ctx, cleanup)
```

means:

> Cleanup when this context is canceled.

So:

```text
function lifecycle
       │
       ▼
     defer
```

versus:

```text
context lifecycle
       │
       ▼
   AfterFunc
```

Sometimes you need both.

---

# 24. `AfterFunc` vs `select`

Traditional cancellation:

```go
select {
case <-ctx.Done():
    cleanup()
case result := <-results:
    return result
}
```

`AfterFunc` is useful when cancellation needs to interact with something that isn't naturally represented as a channel operation.

For example:

```go
stop := context.AfterFunc(ctx, func() {
    cond.Broadcast()
})
defer stop()
```

The distinction is:

```text
select
  → your goroutine actively observes cancellation

AfterFunc
  → cancellation invokes an independent callback
```

---

# 25. Production checklist

When using `context.AfterFunc`, ask:

### Correctness

- Is the callback allowed to run concurrently with the main operation?
    
- Is the callback idempotent?
    
- Can it race with successful completion?
    
- Is there a terminal state machine?
    

### Lifecycle

- Who owns the callback?
    
- Who calls `stop()`?
    
- What happens if the context is never canceled?
    

### Synchronization

- Can `stop()` race with callback execution?
    
- Do I need a `WaitGroup`, mutex, channel, or other synchronization?
    

### Reliability

- Can the callback block?
    
- Can it panic?
    
- Can it trigger another failure?
    

### Performance

- Is the callback lightweight?
    
- Am I creating many cancellation callbacks?
    
- Is this preferable to a dedicated goroutine?
    

### Architecture

- Is cancellation being used appropriately?
    
- Am I accidentally using Context as an event bus?
    

---

# 26. The key distinction to remember

The most important mental model is:

```text
context.AfterFunc(ctx, f)
```

does **not** mean:

```text
cancel(ctx) → execute f synchronously
```

It means:

```text
cancel(ctx)
    │
    ▼
context becomes canceled
    │
    ▼
f is scheduled asynchronously
    │
    ▼
f()
```

And:

```go
stop()
```

means:

> **Try to prevent the callback from starting.**

Not:

> Wait for the callback to finish.

---

# 27. Principal-level takeaway

`context.AfterFunc` is fundamentally a **lifecycle-to-action bridge**:

```text
             Context lifecycle
                    │
                    │ cancellation
                    ▼
             ┌───────────────┐
             │   AfterFunc   │
             └───────┬───────┘
                     │
                     ▼
              asynchronous f()
                     │
          ┌──────────┼──────────┐
          ▼          ▼          ▼
       cleanup    rollback    wake-up
```

The API itself is simple.

The difficult engineering question is:

> **What concurrency guarantees must exist between the callback and the operation that registered it?**

That is where production correctness lives.

In particular, remember these three facts:

```text
1. Callback runs asynchronously.
2. stop() returning false does NOT mean callback completed.
3. Cancellation can race with normal completion.
```

Those three rules explain most of the subtle bugs involving `context.AfterFunc`.
---

## 🔗 References
- ⬆️ Parent: [[Context Trees & Request Cancellation]]
- 📚 Module: `Concurrency & Synchronization`
