---
title: "CSP Concurrency Model in Go"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Goroutines]]"
---
# CSP Concurrency Model in Go

Go's concurrency model is strongly influenced by **CSP — Communicating Sequential Processes**. The central idea is:

> **Don't communicate by sharing memory; share memory by communicating.**

In Go, this philosophy is primarily expressed through **goroutines** and **channels**.

---

## 1. The Problem CSP Solves

Concurrent programs need multiple execution units to coordinate.

The traditional approach is **shared memory**:

```text
Goroutine A ──┐
              ├──> Shared Memory
Goroutine B ──┘
```

Now you need synchronization:

- mutexes
    
- atomics
    
- condition variables
    
- locks
    
- memory ordering
    
- race prevention
    

For example:

```go
var counter int
var mu sync.Mutex

go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()

go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()
```

This can be perfectly correct, but coordination becomes centered around **who may access the shared state and when**.

CSP proposes a different mental model.

---

# 2. CSP Mental Model

Instead of directly sharing state:

```text
Process A ──> Shared State <── Process B
```

processes communicate through channels:

```text
Process A ──> Channel ──> Process B
```

In Go:

```go
ch := make(chan int)

go func() {
    ch <- 42
}()

value := <-ch
fmt.Println(value)
```

The goroutines don't need to directly manipulate the same variable.

The channel becomes the **communication boundary**.

---

# 3. CSP in Go

The three important concepts are:

```text
CSP
 │
 ├── Processes
 │      ↓
 │   goroutines
 │
 ├── Communication
 │      ↓
 │   channels
 │
 └── Synchronization
        ↓
     channel operations
```

### Process → Goroutine

A CSP process corresponds conceptually to an independently executing activity.

Go provides:

```go
go worker()
```

Each goroutine is a lightweight concurrent execution context.

---

### Communication → Channel

Channels provide typed communication:

```go
ch := make(chan string)

ch <- "hello"

msg := <-ch
```

The channel has a type:

```go
chan string
chan int
chan User
chan error
```

This gives compile-time guarantees about what can travel through the communication boundary.

---

# 4. The Most Important Concept: Send and Receive

A channel has two fundamental operations:

```go
ch <- value    // send
value := <-ch  // receive
```

For an **unbuffered channel**:

```go
ch := make(chan int)
```

a send and receive synchronize with each other.

Conceptually:

```text
Goroutine A                    Goroutine B

    ch <- 42  ───────────────>  <-ch
       │                          │
       └──── synchronization ─────┘
```

A send generally cannot complete until a receiver is ready.

This gives channels an important property:

> **Communication can also provide synchronization.**

That's one of the deepest ideas behind Go's concurrency model.

---

# 5. Unbuffered Channels

Consider:

```go
ch := make(chan int)

go func() {
    ch <- 10
    fmt.Println("sent")
}()

fmt.Println(<-ch)
```

The sender reaches:

```go
ch <- 10
```

and may block until the receiver is ready.

Conceptually:

```text
Sender                         Receiver

ch <- 10
   │
   │ waiting
   │
   └──────────────────────────> <-ch
                                   │
                                   ↓
                                  10
```

This makes an unbuffered channel similar to a **rendezvous point**.

Both parties meet at the communication operation.

---

# 6. Buffered Channels

Go also provides buffered channels:

```go
ch := make(chan int, 3)
```

Now the channel can hold up to three values without a receiver.

```text
             capacity = 3

       ┌─────────────────┐
send → │ 10 │ 20 │ 30    │ → receive
       └─────────────────┘
```

Therefore:

```go
ch <- 10 // doesn't necessarily block
ch <- 20 // doesn't necessarily block
ch <- 30 // doesn't necessarily block
ch <- 40 // blocks until space exists
```

This changes the synchronization semantics.

An important engineering distinction:

> **A buffered channel is not simply a "faster channel." It changes the coordination protocol.**

---

# 7. Channels Are More Than Queues

A common beginner mental model is:

> "A channel is a queue."

That's incomplete.

A channel can provide:

### Communication

```go
ch <- value
```

### Synchronization

```go
<-done
```

### Ownership transfer

```go
ch <- object
```

### Backpressure

```go
ch <- job
```

If consumers cannot keep up, producers eventually block.

### Lifecycle signaling

```go
close(done)
```

So channels are better understood as:

> **typed synchronization and communication primitives.**

---

# 8. "Don't Communicate by Sharing Memory"

The famous Go philosophy is:

> **Do not communicate by sharing memory; instead, share memory by communicating.**

Consider:

```go
type Job struct {
    ID   int
    Data []byte
}
```

Instead of:

```text
worker A
   │
   ├── shared Job
   │
worker B
   │
   └── mutex
```

you can transfer ownership conceptually:

```text
Producer
   │
   │ Job
   ▼
Channel
   │
   ▼
Consumer
```

Example:

```go
jobs := make(chan Job)

go func() {
    jobs <- Job{
        ID:   1,
        Data: []byte("hello"),
    }
}()

job := <-jobs
```

The important idea is **ownership and coordination**, not merely avoiding mutexes.

---

# 9. But Channels Do NOT Eliminate Shared Memory

This is a very important correction to a common misconception.

Go still has shared memory.

You can write:

```go
var cache map[string]string
var mu sync.RWMutex
```

and use a mutex correctly.

Channels are not a replacement for every mutex.

A better rule is:

> **Use channels when communication/ownership transfer is the natural abstraction. Use mutexes when protecting shared state is the natural abstraction.**

For example, a concurrent in-memory cache may naturally be:

```go
type Cache struct {
    mu sync.RWMutex
    m  map[string]string
}
```

Trying to turn every cache operation into a channel message can make the design worse.

---

# 10. CSP vs Shared-Memory Concurrency

### Shared-memory model

```text
G1 ──┐
     ├──> State <── lock
G2 ──┘
```

You reason about:

- mutual exclusion
    
- lock ownership
    
- races
    
- deadlocks
    
- memory visibility
    

### CSP-style model

```text
G1 ──> Channel ──> G2
```

You reason about:

- message flow
    
- ownership
    
- blocking
    
- ordering
    
- backpressure
    
- lifecycle
    
- cancellation
    

Neither model is universally superior.

The key is choosing the correct abstraction.

---

# 11. CSP and Pipeline Architecture

One of Go's strongest patterns is the **pipeline**.

Imagine:

```text
Input
  │
  ▼
Stage 1
  │
  ▼
Stage 2
  │
  ▼
Stage 3
  │
  ▼
Output
```

In Go:

```go
func stage1(in <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        defer close(out)

        for n := range in {
            out <- n * 2
        }
    }()

    return out
}
```

Then:

```go
numbers := make(chan int)

result := stage1(numbers)
```

You can compose stages:

```text
numbers
   │
   ▼
stage1
   │
   ▼
stage2
   │
   ▼
stage3
   │
   ▼
result
```

This is CSP thinking applied to software architecture.

---

# 12. Fan-Out

Multiple workers consume from the same channel:

```text
                 ┌── Worker 1
                 │
Jobs ── Channel ─┼── Worker 2
                 │
                 └── Worker 3
```

Example:

```go
jobs := make(chan Job)

for i := 0; i < 3; i++ {
    go worker(jobs)
}
```

Each job is received by one worker.

This is useful for:

- CPU-bound work
    
- I/O-bound work
    
- background processing
    
- bounded concurrency
    

---

# 13. Fan-In

Multiple producers feed one channel:

```text
Producer 1 ──┐
Producer 2 ──┼──> Channel ──> Consumer
Producer 3 ──┘
```

This allows independent concurrent activities to converge into one stream.

---

# 14. `select`: CSP Multiplexing

`select` is one of Go's most important concurrency constructs.

```go
select {
case value := <-ch1:
    fmt.Println(value)

case value := <-ch2:
    fmt.Println(value)
}
```

Conceptually:

```text
             ┌── ch1 ready ──> receive
select ──────┤
             └── ch2 ready ──> receive
```

It allows a goroutine to wait on multiple communication operations.

---

# 15. `select` + Cancellation

This becomes especially powerful with `context.Context`:

```go
select {
case result := <-results:
    return result

case <-ctx.Done():
    return ctx.Err()
}
```

Now the operation has two possible events:

```text
              ┌── result ──> continue
select ───────┤
              └── cancel ──> stop
```

This is fundamental for production Go systems.

---

# 16. Closing Channels

A channel can be closed:

```go
close(ch)
```

Closing means:

> **No more values will be sent.**

It does **not** mean:

> "Destroy the channel."

Receivers can still receive already-buffered values.

Eventually:

```go
value, ok := <-ch
```

returns:

```go
ok == false
```

after the channel is closed and drained.

---

# 17. `range` Over a Channel

A common CSP-style pattern:

```go
for value := range ch {
    process(value)
}
```

This continues until the channel is closed.

Therefore the producer commonly owns closing:

```go
func producer(out chan<- int) {
    defer close(out)

    for i := 0; i < 10; i++ {
        out <- i
    }
}
```

A useful ownership rule:

> **The sender usually owns closing the channel.**

Receivers generally should not close channels they don't own.

---

# 18. Directional Channels

Go lets you encode communication direction in APIs.

```go
func producer(out chan<- int)
```

means:

```text
send-only
```

And:

```go
func consumer(in <-chan int)
```

means:

```text
receive-only
```

This is excellent API design because the compiler enforces the protocol.

For example:

```go
func producer(out chan<- int) {
    out <- 42
}
```

The producer cannot accidentally receive.

---

# 19. CSP and Ownership

A more advanced mental model is:

> **Channels can represent ownership transfer.**

Suppose a producer constructs a resource:

```go
buf := make([]byte, 4096)
```

and sends it:

```go
ch <- buf
```

The architecture can conceptually become:

```text
Producer owns buf
       │
       │ send
       ▼
    Channel
       │
       ▼
Consumer owns buf
```

This reduces the need for shared mutable state.

But remember: Go does not enforce ownership like Rust.

It is a **design discipline**, not a language-enforced ownership system.

---

# 20. CSP Does Not Mean "Use Goroutines Everywhere"

This is another important engineering lesson.

Bad:

```go
for _, item := range items {
    go process(item)
}
```

If `items` contains 10 million elements, you may create an enormous amount of concurrency.

Production code should usually introduce **bounded concurrency**.

For example:

```text
                ┌── worker
                ├── worker
Jobs ── queue ──┼── worker
                ├── worker
                └── worker
```

The number of workers becomes a deliberate resource limit.

---

# 21. CSP and Backpressure

Suppose:

```text
Producer: 100k jobs/sec
Consumer: 10k jobs/sec
```

Without a bound:

```text
Producer
   │
   ▼
unbounded queue
   │
   ▼
memory exhaustion
```

With a bounded channel:

```go
jobs := make(chan Job, 1000)
```

the producer eventually blocks:

```text
Producer
   │
   ▼
[1000 jobs]
   │
   ▼
Workers
```

That is **backpressure**.

This is one of the most valuable production properties of channels.

---

# 22. A Production Worker Pool

A simplified worker pool:

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

Then:

```go
jobs := make(chan Job, 100)

for i := 0; i < 8; i++ {
    go worker(ctx, jobs)
}
```

Now concurrency is explicitly bounded:

```text
                  ┌── Worker 1
                  ├── Worker 2
                  ├── Worker 3
Jobs ──[100]──────┼── ...
                  ├── Worker 7
                  └── Worker 8
```

This is much closer to production-grade CSP usage.

---

# 23. The Hidden Cost: Goroutine Leaks

CSP-style code can still fail badly.

Example:

```go
func worker(ch <-chan int) {
    go func() {
        value := <-ch
        process(value)
    }()
}
```

If nobody ever sends:

```go
ch <- value
```

the goroutine can remain blocked forever.

That's a **goroutine leak**.

At production scale:

```text
100 leaked goroutines
       ↓
10,000 leaked goroutines
       ↓
memory / scheduler / resources
       ↓
system degradation
```

Channels introduce lifecycle responsibilities.

---

# 24. Every Goroutine Needs a Lifetime

A Principal-level question is:

> **Who stops this goroutine?**

Before creating:

```go
go worker()
```

ask:

1. What starts it?
    
2. What does it wait for?
    
3. What causes it to stop?
    
4. Who closes its input?
    
5. What happens on cancellation?
    
6. What happens if downstream disappears?
    
7. Can it block forever?
    
8. Who waits for it?
    

This mindset is much more important than simply knowing channel syntax.

---

# 25. `WaitGroup` Is Often Part of the Solution

CSP handles communication, but you often need explicit lifecycle synchronization.

For example:

```go
var wg sync.WaitGroup

for i := 0; i < 8; i++ {
    wg.Add(1)

    go func() {
        defer wg.Done()
        worker()
    }()
}

wg.Wait()
```

This illustrates an important point:

> Go's concurrency model is not "channels instead of synchronization."

Rather:

```text
Channels       → communication
Mutexes         → shared-state synchronization
WaitGroup       → lifecycle synchronization
Context         → cancellation/deadlines
Atomics         → low-level shared-state coordination
```

A good Go engineer chooses among them based on the problem.

---

# 26. CSP vs Actor Model

CSP and the Actor Model are related but not identical.

### CSP

```text
Process ── Channel ──> Process
```

Communication occurs through explicit channels.

### Actor Model

```text
Actor A ──message──> Actor B
```

Actors typically encapsulate their own state and communicate through mailboxes.

Go is more naturally CSP-oriented:

```go
ch <- message
```

rather than:

```text
send_message(actor, message)
```

But you can build actor-like systems using goroutines + channels.

---

# 27. CSP Is Not Magic

CSP does not automatically prevent:

### Deadlocks

```go
ch := make(chan int)

ch <- 10
```

If nobody receives, the goroutine blocks forever.

### Goroutine leaks

A goroutine waits indefinitely.

### Livelocks

Goroutines continue executing but make no useful progress.

### Backpressure collapse

Consumers become slower than producers.

### Resource exhaustion

Too many:

- goroutines
    
- channels
    
- queued messages
    
- buffers
    
- external requests
    

### Poor architecture

You can create an enormous web of channels that is harder to understand than a simple mutex-protected data structure.

---

# 28. Common Anti-Pattern: Channel for Everything

Bad reasoning:

> "Go likes channels, therefore every shared variable should become a channel."

For example, don't necessarily create:

```go
requests := make(chan Request)
responses := make(chan Response)
commands := make(chan Command)
state := make(chan State)
```

just because channels are idiomatic.

Sometimes this:

```go
type Counter struct {
    mu sync.Mutex
    n  int64
}
```

is significantly simpler and more efficient.

The correct principle is:

> **Use communication when communication is the problem. Use synchronization primitives when shared state is the problem.**

---

# 29. Performance Model

Channels have costs.

A communication operation may involve:

- synchronization
    
- scheduler interaction
    
- goroutine parking/unparking
    
- memory operations
    
- contention
    
- copying values
    

Therefore:

```go
ch <- smallValue
```

is not equivalent in cost to:

```go
x++
```

You should not replace ordinary sequential code with channel-based coordination without a reason.

For performance-sensitive systems:

```text
Measure
  ↓
Benchmark
  ↓
Profile
  ↓
Identify contention
  ↓
Optimize
```

Use:

- `go test -bench`
    
- `pprof`
    
- race detector
    
- mutex/block profiles where appropriate
    
- production metrics
    

rather than assuming channels are fast or slow.

---

# 30. CSP and Memory Model

Channels also participate in Go's synchronization guarantees.

The important mental model is:

```text
send
  │
  │ happens-before relationship
  ▼
receive
```

This means communication can establish visibility between goroutines.

For example:

```go
var x int

ch := make(chan struct{})

go func() {
    x = 42
    ch <- struct{}{}
}()

<-ch

fmt.Println(x)
```

The synchronization through the channel means the receiver can safely observe the preceding write under Go's memory model.

This is fundamentally different from unsynchronized shared access.

---

# 31. CSP at System Architecture Level

The most interesting use of CSP isn't syntax.

It's architectural decomposition.

Imagine an HTTP service:

```text
HTTP Handler
     │
     ▼
Request Queue
     │
     ▼
Workers
     │
     ├── Database
     │
     ├── External API
     │
     └── Event Publisher
```

Channels can represent internal boundaries.

But you should ask:

> Does this boundary actually need asynchronous communication?

If the operation must synchronously return a result:

```text
HTTP → Service → DB → HTTP response
```

adding channels may introduce unnecessary complexity.

If the workload is asynchronous:

```text
HTTP
 │
 ▼
enqueue
 │
 ▼
return 202
 │
 ▼
Workers
```

then a queue/channel abstraction may make sense.

---

# 32. Go Channels vs External Queues

A crucial production distinction:

```text
Go channel
```

is **process-local**.

If the process crashes:

```text
Process dies
    ↓
channel disappears
    ↓
queued messages disappear
```

An external queue such as Kafka, NATS, RabbitMQ, SQS, etc. provides very different durability and distribution semantics.

Therefore:

```text
Go channel
    =
in-process coordination
```

not:

```text
Go channel
    =
distributed durable messaging system
```

This distinction matters enormously in distributed systems.

---

# 33. The Principal Engineer Mental Model

When you see:

```go
ch := make(chan Job)
go worker(ch)
```

don't stop at:

> "This uses channels."

Ask:

### Ownership

Who owns the jobs?

### Lifecycle

Who closes the channel?

### Cancellation

What happens when the request is cancelled?

### Backpressure

What happens when workers are slower than producers?

### Capacity

Why is the buffer size `100`?

### Ordering

Does job ordering matter?

### Failure

What happens when `process(job)` fails?

### Retry

Who retries?

### Durability

Can jobs be lost?

### Scaling

What happens at 10× throughput?

### Observability

Can we measure:

- queue depth?
    
- processing latency?
    
- worker utilization?
    
- dropped jobs?
    
- failures?
    

This is the difference between **knowing CSP** and **engineering a reliable concurrent system**.

---

# 34. Practical Decision Framework

When deciding whether to use a channel, ask:

```text
Is there communication between concurrent activities?
                │
               Yes
                │
                ▼
Does ownership/message passing naturally express it?
                │
          ┌─────┴─────┐
         Yes           No
          │             │
          ▼             ▼
       Channel       Shared state
                       │
                       ▼
                 Mutex / Atomic
```

Then ask:

```text
Does the work need asynchronous buffering?
        │
        ├── No → unbuffered/direct coordination
        │
        └── Yes → bounded buffered channel
```

Then:

```text
Does the data need to survive process failure?
        │
        ├── No → channel may be sufficient
        │
        └── Yes → external durable mechanism
```

---

# 35. Key Takeaways

The most important things to internalize are:

1. **CSP models concurrency as independent processes communicating through channels.**
    
2. **Goroutines provide the concurrent execution units.**
    
3. **Channels provide typed communication and synchronization.**
    
4. **Unbuffered channels provide rendezvous-style synchronization.**
    
5. **Buffered channels introduce bounded asynchronous buffering and backpressure.**
    
6. **Channels can represent ownership transfer, not just queues.**
    
7. **`select` multiplexes concurrent communication and cancellation.**
    
8. **Channel-based design does not eliminate shared memory or mutexes.**
    
9. **Every goroutine needs a well-defined lifetime.**
    
10. **Every buffer needs a reason and capacity rationale.**
    
11. **Channels are process-local; they are not durable distributed queues.**
    
12. **CSP is a design model, not a rule that "everything must use channels."**
    

The deepest Go concurrency principle is therefore not:

> **"Use channels."**

It is:

> **Model concurrency around clear ownership, communication, synchronization, lifecycle, and failure boundaries—and use channels when they make those boundaries simpler and safer.**

---

## 🔗 References
- ⬆️ Parent: [[Goroutines]]
- 📚 Module: `Concurrency & Synchronization`
