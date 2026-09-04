---
title: "Deadlock"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Goroutines]]"
---
# Goroutine Deadlock in Go

A **goroutine deadlock** occurs when goroutines are unable to make progress because each goroutine is waiting for an event that can never happen.

The key idea is:

> **Deadlock is not “a goroutine is blocked.” Deadlock is a cycle of dependencies where progress becomes impossible.**

---

## 1. The Mental Model

Think of goroutines as workers and synchronization primitives as resources.

```text
G1 ──waiting for──> resource/event owned by G2
 ↑                                │
 └────────waiting for─────────────┘
```

If the dependency graph contains a cycle, nobody can progress.

For example:

```text
Goroutine 1:
    lock A
    wait for lock B

Goroutine 2:
    lock B
    wait for lock A
```

Result:

```text
G1 owns A ──> waits for B
                  ↑
                  │
G2 owns B ────────┘
```

Neither goroutine can continue.

---

# 2. Simplest Example

```go
package main

func main() {
    ch := make(chan int)

    ch <- 42

    println("done")
}
```

This deadlocks.

Why?

An **unbuffered channel** requires a sender and receiver to rendezvous.

```text
main goroutine
      │
      │ send 42
      ▼
   channel
      │
      X
   no receiver
```

The send blocks forever.

Eventually Go detects that all goroutines are blocked:

```text
fatal error: all goroutines are asleep - deadlock!
```

---

# 3. The Receiver Can Deadlock Too

```go
func main() {
    ch := make(chan int)

    value := <-ch

    println(value)
}
```

There is no sender.

```text
Receiver
   │
   ▼
 channel
   │
   X
no sender
```

The receive blocks forever.

---

# 4. Goroutine Deadlock vs Goroutine Leak

These are related but **not the same**.

### Deadlock

The program cannot make progress.

```go
func main() {
    ch := make(chan int)

    go func() {
        ch <- 10
    }()

    <-ch
}
```

This one actually works because sender and receiver synchronize.

A deadlock would be:

```go
func main() {
    ch := make(chan int)

    go func() {
        <-ch
    }()

    <-ch
}
```

Both goroutines are waiting for a sender.

```text
G1 ──receive──> ch <──receive── G2
```

Nobody sends.

---

### Goroutine leak

A goroutine remains blocked longer than intended, even though the rest of the application can continue.

```go
func worker() {
    ch := make(chan int)

    go func() {
        <-ch
    }()
}
```

The caller returns, but the goroutine remains blocked.

Over time:

```text
request 1 → leaked goroutine
request 2 → leaked goroutine
request 3 → leaked goroutine
...
request N → thousands of goroutines
```

A leak may eventually contribute to resource exhaustion, but it does not necessarily immediately deadlock the whole application.

---

# 5. Mutex Deadlock

A classic example:

```go
var mu sync.Mutex

func main() {
    mu.Lock()
    mu.Lock()
}
```

The same goroutine attempts to acquire the same non-reentrant mutex twice.

```text
G1
 │
 ├── Lock()
 │
 │   owns mutex
 │
 └── Lock()
      │
      └── waits for itself forever
```

Go's `sync.Mutex` is **not reentrant**.

---

# 6. Two-Goroutine Mutex Deadlock

More interesting:

```go
var (
    mu1 sync.Mutex
    mu2 sync.Mutex
)

func main() {
    go func() {
        mu1.Lock()
        defer mu1.Unlock()

        mu2.Lock()
        defer mu2.Unlock()
    }()

    mu2.Lock()
    defer mu2.Unlock()

    mu1.Lock()
    defer mu1.Unlock()
}
```

Potential execution:

```text
G1:
    owns mu1
    waits for mu2

G2:
    owns mu2
    waits for mu1
```

Dependency graph:

```text
G1 ──waits for──> G2
 ↑                  │
 │                  │
 └──waits for──────┘
```

This is a **circular wait**.

---

# 7. The Four Conditions for Deadlock

A useful classical model is **Coffman conditions**.

Deadlock requires all four:

### 1. Mutual exclusion

A resource can be owned by only one goroutine.

```text
G1 owns mutex
G2 cannot acquire it
```

### 2. Hold and wait

A goroutine holds one resource while waiting for another.

```text
G1:
    owns A
    waits for B
```

### 3. No preemption

The resource cannot simply be forcibly taken away.

### 4. Circular wait

There is a cycle:

```text
G1 → G2 → G3 → G1
```

This is an extremely useful mental model when diagnosing synchronization problems.

---

# 8. WaitGroup Deadlock

A common mistake:

```go
func main() {
    var wg sync.WaitGroup

    wg.Add(1)

    wg.Wait()
}
```

You incremented the counter but never called:

```go
wg.Done()
```

Therefore:

```text
counter = 1

Wait()
  ↓
counter != 0
  ↓
wait forever
```

Another dangerous pattern:

```go
wg.Add(1)

go func() {
    // something goes wrong
    return
}()

wg.Wait()
```

If the goroutine exits without `Done()`, the counter never reaches zero.

Prefer:

```go
wg.Add(1)

go func() {
    defer wg.Done()

    // work
}()

wg.Wait()
```

This makes the lifecycle explicit and protects against early returns.

---

# 9. Channel + WaitGroup Deadlock

This is particularly important in production Go code.

```go
func main() {
    jobs := make(chan int)
    var wg sync.WaitGroup

    wg.Add(1)

    go func() {
        defer wg.Done()

        for job := range jobs {
            println(job)
        }
    }()

    wg.Wait()

    jobs <- 42
}
```

Deadlock:

```text
main
 │
 ├── wg.Wait()
 │
 │   waits for worker
 │
worker
 │
 └── waits for jobs
```

But the main goroutine sends the job **only after** waiting.

So:

```text
main ──waits for──> worker
                       │
                       └──waits for──> job
                                         ↑
                                         │
                                  main hasn't sent
```

This is a dependency cycle.

---

# 10. How to Fix It

Send work before waiting:

```go
func main() {
    jobs := make(chan int)
    var wg sync.WaitGroup

    wg.Add(1)

    go func() {
        defer wg.Done()

        for job := range jobs {
            println(job)
        }
    }()

    jobs <- 42
    close(jobs)

    wg.Wait()
}
```

Now the lifecycle is:

```text
main
 │
 ├── start worker
 │
 ├── send job
 │       │
 │       ▼
 │    worker
 │
 ├── close channel
 │
 └── Wait()
```

No cycle exists.

---

# 11. Buffered Channels Change the Dynamics

Compare:

```go
ch := make(chan int)
ch <- 1
```

with:

```go
ch := make(chan int, 1)
ch <- 1
```

The second can proceed because the buffer has capacity.

```text
Unbuffered:

sender ──wait──> receiver


Buffered:

sender ──> [ 1 ]
             buffer
```

But buffering **does not eliminate deadlocks**.

For example:

```go
ch := make(chan int, 1)

ch <- 1
ch <- 2
```

The first send succeeds.

The second blocks because:

```text
capacity = 1
current  = 1
```

There is still no receiver.

---

# 12. Nil Channels

A particularly nasty Go-specific case:

```go
var ch chan int

ch <- 10
```

`ch == nil`.

A send to a nil channel blocks forever.

Likewise:

```go
<-ch
```

blocks forever.

And:

```go
close(ch)
```

panics.

Mental model:

```text
nil channel

send    → blocks forever
receive → blocks forever
close   → panic
```

This is useful intentionally with `select`, but dangerous accidentally.

---

# 13. `select` Can Also Deadlock

```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    select {
    case <-ch1:
    case <-ch2:
    }
}
```

Both channels have no sender.

Therefore all cases block.

Without a `default`:

```text
select
 ├── ch1 → blocked
 └── ch2 → blocked
```

The goroutine waits forever.

---

# 14. Detecting Deadlocks

Start with the runtime error:

```text
fatal error: all goroutines are asleep - deadlock!
```

But don't stop there.

The stack trace is usually the most important evidence.

You might see:

```text
goroutine 1 [chan receive]:
main.main()
```

or:

```text
goroutine 7 [sync.Mutex.Lock]:
...
```

or:

```text
goroutine 12 [semacquire]:
sync.(*WaitGroup).Wait(...)
```

These states tell you **what synchronization primitive the goroutine is waiting on**.

---

# 15. Production Debugging Method

Do not randomly add `Sleep()` or change channel buffering.

Use:

```text
Reproduce
   ↓
Capture goroutine dump
   ↓
Identify blocked goroutines
   ↓
Identify what each goroutine waits for
   ↓
Build dependency graph
   ↓
Find cycle / missing event
   ↓
Fix ownership/lifecycle
   ↓
Add regression test
```

The key question is:

> **“What event is this goroutine waiting for, and which goroutine is responsible for producing that event?”**

Then ask:

> **“Can that producer itself be waiting for me?”**

That question catches a large class of deadlocks.

---

# 16. Preventing Deadlocks

### Establish lock ordering

Bad:

```text
G1: Lock(A) → Lock(B)
G2: Lock(B) → Lock(A)
```

Good:

```text
Everyone:
    Lock(A)
    Lock(B)
```

Define a global ordering:

```text
A < B < C
```

and always acquire locks in that order.

---

### Keep critical sections small

Avoid:

```go
mu.Lock()
defer mu.Unlock()

doNetworkRequest()
doDatabaseQuery()
callExternalService()
```

Holding a mutex across slow I/O increases contention and makes dependency cycles much easier to create.

Prefer:

```text
lock
  ↓
read/update shared state
  ↓
unlock
  ↓
perform slow I/O
```

---

### Make ownership explicit

For channels, determine:

> **Who creates it? Who sends? Who receives? Who closes it?**

A strong convention is:

**The sender owns channel closure.**

For example:

```go
func producer() <-chan int {
    ch := make(chan int)

    go func() {
        defer close(ch)

        // produce values
    }()

    return ch
}
```

The caller cannot accidentally close the channel because it only receives from it.

---

# 17. Deadlock Is Fundamentally a Dependency Problem

The most important mental model is not:

> “Mutexes cause deadlocks.”

It is:

> **Deadlocks happen when progress dependencies form an impossible cycle or a required event has no producer.**

Examples:

```text
Mutex:
A → B → A

Channel:
receive → no sender

WaitGroup:
Wait → missing Done

Pipeline:
stage 1 waits for stage 2
stage 2 waits for stage 1

Resource pool:
worker waits for connection
connection waits for worker
```

So at Staff+/Principal level, don't merely inspect individual locks.

Inspect the **system's dependency graph and ownership model**.

---

## Key Takeaways

1. **Blocked ≠ deadlocked.**
    
2. Deadlock means **no possible progress** under the current state.
    
3. Look for **circular dependencies**.
    
4. Also look for **missing producers/events**.
    
5. `Mutex`, `Channel`, `WaitGroup`, `select`, and resource pools can all participate.
    
6. Buffered channels reduce some blocking but **do not guarantee deadlock freedom**.
    
7. Establish **lock ordering**.
    
8. Make goroutine/channel ownership explicit.
    
9. Use goroutine dumps and synchronization states as evidence.
    
10. Fix the **dependency/lifecycle design**, not merely the symptom.
    

A useful Principal Engineer question is:

> **“For every blocking operation, who guarantees that the condition allowing it to proceed will eventually become true?”**

If you cannot answer that, you probably have a concurrency design problem.

---

## 🔗 References
- ⬆️ Parent: [[Goroutines]]
- 📚 Module: `Concurrency & Synchronization`
