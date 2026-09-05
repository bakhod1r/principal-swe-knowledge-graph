---
title: "Unbuffered Channel Synchronous Rendezvous"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# Unbuffered Channel — Synchronous Rendezvous

In Go, an **unbuffered channel** (`make(chan T)`) provides a synchronization point where a send and receive must **rendezvous**. There is no channel buffer that can absorb the value.

The core mental model is:

> **A send cannot complete until a receiver is ready, and a receive cannot complete until a sender is ready.**

---

## 1. What problem does it solve?

An unbuffered channel is useful when you want **both communication and synchronization**.

```go
ch := make(chan int)

go func() {
    ch <- 42
}()

v := <-ch
fmt.Println(v)
```

The important part is not merely that `42` moves from one goroutine to another.

The send and receive **synchronize their execution**.

Conceptually:

```text
Sender goroutine                 Receiver goroutine

     ch <- 42
        │
        │   rendezvous
        ▼
   ┌───────────┐
   │   value   │
   └───────────┘
        │
        ▼
                                v := <-ch
```

There is no intermediate queue:

```text
Sender ───────► Receiver
```

rather than:

```text
Sender ─► [buffer] ─► Receiver
```

---

# 2. Buffered vs unbuffered

Consider:

```go
buffered := make(chan int, 3)
unbuffered := make(chan int)
```

### Buffered

```text
Sender
  │
  ▼
┌───────────────┐
│  10  20  30   │  capacity = 3
└───────────────┘
  │
  ▼
Receiver
```

The sender can proceed while there is buffer capacity.

### Unbuffered

```text
Sender
   │
   │
   ▼
  ╔══════════╗
  ║ rendezvous║
  ╚══════════╝
   │
   ▼
Receiver
```

The sender needs a receiver.

---

# 3. What happens internally?

This connects directly to Go's runtime `hchan` and `sudog` wait queues.

An unbuffered channel essentially has:

```text
hchan
 ├── qcount = 0
 ├── dataqsiz = 0
 ├── sendq
 └── recvq
```

There is **no circular buffer** because:

```text
dataqsiz == 0
```

So the important structures become the waiting queues:

```text
sendq                         recvq

sender A ── sudog             receiver B ── sudog
sender C ── sudog             receiver D ── sudog
```

When a sender arrives and a receiver is already waiting:

```text
sender
  │
  ▼
sendq/recvq matching
  │
  ▼
direct handoff
  │
  ▼
receiver
```

The runtime can transfer the value directly between the goroutines rather than placing it into a channel buffer.

---

# 4. The critical distinction: direct handoff

For an unbuffered channel:

```go
ch <- x
```

does **not** mean:

```text
copy x → channel storage
```

because there is no buffer storage.

Instead, when a receiver is waiting, the runtime can perform a **sender → receiver handoff**.

Conceptually:

```text
Before:

Sender                    Receiver

  x                          waiting
  │                            ▲
  │                            │
  └────────── channel ─────────┘


After rendezvous:

Sender                    Receiver

  │                          x
  │                          │
  └────── value transfer ────┘
```

The exact runtime implementation has additional synchronization and scheduling machinery, but this is the correct mental model.

---

# 5. Why does the sender block?

Suppose:

```go
ch := make(chan int)

ch <- 10
```

and there is no receiver.

The send cannot finish.

```text
Goroutine
    │
    ▼
ch <- 10
    │
    ▼
No receiver
    │
    ▼
WAIT
```

The goroutine becomes blocked until another goroutine performs:

```go
<-ch
```

This is why this program deadlocks:

```go
func main() {
    ch := make(chan int)

    ch <- 10

    fmt.Println("done")
}
```

There is no receiver.

Eventually Go reports:

```text
fatal error: all goroutines are asleep - deadlock!
```

---

# 6. The reverse is also true

This:

```go
ch := make(chan int)

v := <-ch
```

also blocks if nobody sends.

```text
Receiver
    │
    ▼
<-ch
    │
    ▼
No sender
    │
    ▼
WAIT
```

So:

```text
                 Unbuffered channel

       sender                         receiver
          │                              │
          │                              │
          ├────── must meet ────────────┤
          │                              │
          ▼                              ▼
       proceed                         proceed
```

This is why the term **synchronous rendezvous** is appropriate.

---

# 7. The rendezvous point

Consider:

```go
func main() {
    ch := make(chan string)

    go func() {
        fmt.Println("before send")
        ch <- "hello"
        fmt.Println("after send")
    }()

    time.Sleep(time.Second)

    fmt.Println("before receive")
    msg := <-ch
    fmt.Println("received:", msg)
}
```

Execution conceptually:

```text
Sender                         Receiver

before send

ch <- "hello"
     │
     │ blocked
     │
     │                       before receive
     │                              │
     └──────── rendezvous ──────────┘
                                    │
                                    ▼
                               receive hello

after send
```

Notice:

```text
send completes
      ↓
receiver receives
```

The send and receive form a synchronization event.

---

# 8. Happens-before relationship

This is one of the most important reasons channels exist.

For a successful channel communication, Go's memory model provides synchronization between the send and corresponding receive.

Conceptually:

```go
x = 42
ch <- struct{}{}
```

and:

```go
<-ch
fmt.Println(x)
```

The communication establishes the necessary ordering so that the receiver can observe the preceding effects appropriately under Go's memory model.

Mental model:

```text
Sender goroutine

x = 42
   │
   ▼
send
   │
   │ synchronization
   ▼
receive
   │
   ▼
read x
```

This makes channels useful not just for **data transfer**, but also for **coordination**.

---

# 9. Unbuffered channel as a synchronization primitive

You can send an empty value:

```go
done := make(chan struct{})

go func() {
    doWork()
    done <- struct{}{}
}()

<-done
fmt.Println("work completed")
```

The actual payload is irrelevant.

The channel is being used as a **completion signal**.

A common variation is:

```go
close(done)
```

when you need to notify multiple receivers rather than synchronize with exactly one receiver.

That distinction matters:

```text
send on channel
    ↓
one communication event

close channel
    ↓
broadcast-style notification to receivers
```

---

# 10. Why not always use buffered channels?

Because buffering changes the synchronization semantics.

Compare:

```go
ch := make(chan int)
```

with:

```go
ch := make(chan int, 1)
```

With the unbuffered version:

```go
ch <- 10
```

requires a receiver.

With capacity `1`:

```go
ch <- 10
```

can complete immediately if the buffer is empty.

Therefore:

```text
Unbuffered:

send ──────► receiver
   must synchronize


Buffered:

send ──────► buffer
   can continue
                 │
                 ▼
              receiver
```

So buffer capacity is not merely a performance parameter.

It changes the **coordination contract**.

---

# 11. Producer-consumer example

### Unbuffered

```go
jobs := make(chan Job)

go worker(jobs)

jobs <- job
```

The producer knows that the worker has reached the receive operation when the send successfully completes.

That provides a useful synchronization property:

```text
Producer
   │
   │ job
   ▼
Worker
   │
   ▼
send completes
```

The producer cannot outrun the worker indefinitely.

This gives a natural form of **backpressure**.

---

# 12. But unbuffered does NOT mean "zero latency"

A common misconception is:

> "There is no buffer, so communication must be faster."

Not necessarily.

A rendezvous may involve:

- channel locking
    
- queue manipulation
    
- goroutine state transitions
    
- scheduler interaction
    
- wake-up of another goroutine
    
- synchronization costs
    

Therefore:

```text
unbuffered ≠ automatically faster
```

The correct engineering question is:

> What synchronization semantics does my workload require?

Benchmark if performance matters.

---

# 13. Select interaction

Unbuffered channels become particularly powerful with `select`:

```go
select {
case ch <- value:
    // receiver was available
case <-ctx.Done():
    // cancellation
}
```

This is a production-grade pattern because the send doesn't have to wait forever.

Without cancellation:

```go
ch <- value
```

could block indefinitely.

With:

```go
select {
case ch <- value:
case <-ctx.Done():
}
```

you have an explicit escape path.

---

# 14. Timeout / bounded waiting

Another common pattern:

```go
select {
case ch <- value:
    // sent
case <-time.After(time.Second):
    // timed out
}
```

Although for production code with repeated operations, prefer an appropriately managed `time.Timer` rather than repeatedly allocating timers with `time.After` in hot paths.

The architectural principle is more important:

```text
Communication
      +
Bounded waiting
      =
More resilient system
```

---

# 15. Nil channel behavior

An especially important edge case:

```go
var ch chan int
```

Here:

```go
ch == nil
```

Operations block forever:

```go
ch <- 1  // blocks forever
<-ch     // blocks forever
```

But:

```go
close(ch)
```

panics.

This is useful in `select` because a nil channel can effectively disable a case:

```go
var ch chan int

select {
case ch <- 10:
    // disabled because ch == nil
case <-ctx.Done():
    // still active
}
```

---

# 16. Closed unbuffered channel

Closing is different from sending.

```go
close(ch)
```

means:

> No more values will ever be sent.

After closure:

```go
v, ok := <-ch
```

returns:

```text
v  = zero value
ok = false
```

immediately once all buffered values have been drained.

For an **unbuffered** channel, there are no buffered values to drain.

So after closure:

```go
v, ok := <-ch
```

immediately returns:

```text
zero value, false
```

---

# 17. A useful mental model

Think of an unbuffered channel as a **two-party handshake**:

```text
                UNBUFFERED CHANNEL

       Sender                         Receiver
         │                               │
         │ "I have a value"              │
         │                               │
         ├──────── rendezvous ───────────┤
         │                               │
         │ "I am ready to receive"       │
         │                               │
         ▼                               ▼
      continue                         continue
```

The key invariant is:

> **A successful send and receive happen as one synchronization event.**

---

# 18. Common mistake

Don't think:

```text
channel = queue
```

for every Go channel.

A better model is:

```text
Buffered channel
    =
queue + synchronization


Unbuffered channel
    =
rendezvous + synchronization
```

The runtime `hchan` abstraction supports both, but the operational behavior differs significantly.

---

# 19. Production perspective

When choosing an unbuffered channel, ask:

### Use it when:

- you need direct handoff
    
- sender and receiver should synchronize
    
- natural backpressure is desirable
    
- you want a strong coordination boundary
    
- the producer should not get ahead of the consumer
    

### Be careful when:

- the receiver may disappear
    
- the receiver can become permanently blocked
    
- the producer can wait indefinitely
    
- there is no cancellation path
    
- throughput requires decoupling producers and consumers
    

In production systems, this:

```go
ch <- value
```

is often safer as:

```go
select {
case ch <- value:
case <-ctx.Done():
    return ctx.Err()
}
```

when cancellation is part of the operation's contract.

---

# 20. Principal-level takeaway

The deepest insight is:

```text
Buffered channel
    ↓
decouples producer and consumer
    ↓
absorbs temporary imbalance


Unbuffered channel
    ↓
couples producer and consumer
    ↓
forces rendezvous
    ↓
creates synchronization + backpressure
```

So choosing:

```go
make(chan T)
```

versus:

```go
make(chan T, N)
```

is an **architectural decision**, not merely a runtime tuning parameter.

You are deciding how much temporal independence your components are allowed to have.

**One-line mental model:**

> **An unbuffered channel is a synchronous handoff: neither side can complete the communication alone.**

---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
