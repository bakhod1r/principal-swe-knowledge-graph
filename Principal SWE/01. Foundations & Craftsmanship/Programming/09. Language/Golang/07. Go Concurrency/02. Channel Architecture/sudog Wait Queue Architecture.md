---
title: "sudog Wait Queue Architecture"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# `sudog` Wait Queue Architecture in Go

`sudog` is one of the most important internal runtime structures for understanding **how Go blocks and wakes goroutines on channels and synchronization primitives**.

The key idea:

> **A `sudog` is a runtime-side descriptor representing a goroutine that is currently waiting on a synchronization operation.**

It is not the goroutine itself. It is a small object that connects a waiting **G** to a runtime wait queue such as a channel's send/receive queue.

---

## 1. The Problem `sudog` Solves

Consider:

```go
ch := make(chan int)

go func() {
    x := <-ch
    println(x)
}()

ch <- 42
```

The receiving goroutine cannot continue when it executes:

```go
x := <-ch
```

because there is no value available yet.

The runtime needs to:

1. identify the waiting goroutine;
    
2. put it somewhere;
    
3. remove it from runnable execution;
    
4. remember what operation it was waiting for;
    
5. later find it;
    
6. wake it when a sender provides a value.
    

A goroutine alone is not enough to efficiently represent all this queue state.

That's where `sudog` enters.

---

# 2. Mental Model

Think of:

```text
G = Goroutine
sudog = "waiting ticket" for G
wait queue = collection of tickets
```

Conceptually:

```text
        Channel
       ┌─────────────┐
sendq  │ sudog ──► sudog ──► sudog
       └─────────────┘
               │
               ▼
               G
```

For a receiver:

```text
recvq
  │
  ▼
sudog
  │
  ▼
waiting G
```

The `sudog` allows the runtime to connect:

```text
synchronization object
        ↓
    wait queue
        ↓
      sudog
        ↓
    goroutine G
```

---

# 3. Where Does `sudog` Live?

The structure is defined inside the Go runtime, conceptually like:

```go
type sudog struct {
    g *g

    next *sudog
    prev *sudog

    elem unsafe.Pointer

    acquiretime int64
    releasetime int64

    waitlink *sudog
    waittail *sudog

    c *hchan

    ...
}
```

The exact fields can change between Go releases, so when studying a particular Go version, always inspect that version's runtime source.

The important fields conceptually are:

|Field|Purpose|
|---|---|
|`g`|Goroutine associated with this wait|
|`next`|Next element in a wait queue|
|`prev`|Previous element|
|`elem`|Address of the value involved in the operation|
|`c`|Channel associated with the wait|
|wait-related fields|Used by synchronization/wait machinery|

---

# 4. `hchan` + `sudog`

For channels, the relationship is especially important.

A channel's runtime representation (`hchan`) contains queues for blocked senders and receivers.

Conceptually:

```text
                 hchan
          ┌──────────────────┐
          │ qcount            │
          │ dataqsiz          │
          │ buf               │
          │ sendx             │
          │ recvx             │
          │                  │
          │ sendq            │
          │ recvq            │
          └───────┬──────────┘
                  │
       ┌──────────┴──────────┐
       ▼                     ▼

    sendq                   recvq
  ┌───────┐               ┌───────┐
  │sudog  │──► ...        │sudog  │──► ...
  └───┬───┘               └───┬───┘
      │                       │
      ▼                       ▼
      G                       G
```

So the channel does **not** directly maintain a queue of goroutines.

It maintains a queue of `sudog`s.

---

# 5. Blocking a Receiver

Suppose:

```go
ch := make(chan int)

x := <-ch
```

and no sender is currently available.

Conceptually the runtime performs something like:

```text
G1 executes receive
        │
        ▼
Is data available?
        │
       NO
        │
        ▼
Create/acquire sudog
        │
        ▼
sudog.g = G1
sudog.c = ch
sudog.elem = address where value should go
        │
        ▼
enqueue sudog into ch.recvq
        │
        ▼
park G1
```

Now:

```text
ch.recvq
   │
   ▼
sudog
   │
   ▼
  G1
```

`G1` is no longer runnable.

---

# 6. What Does "Park" Mean?

Parking is fundamentally:

> **Remove the goroutine from runnable execution until some event makes it runnable again.**

Conceptually:

```text
Before:

Run Queue
──────────────
G1 G2 G3 G4


After G1 blocks:

Run Queue
──────────────
G2 G3 G4

Channel recvq
──────────────
sudog → G1
```

This is why blocking a goroutine does **not** necessarily block an OS thread.

The runtime scheduler can run another goroutine.

---

# 7. Sender Arrives

Now another goroutine executes:

```go
ch <- 42
```

The runtime sees:

```text
ch.recvq != empty
```

Therefore it doesn't necessarily need to put `42` into the channel buffer.

Instead it can directly transfer the value to the waiting receiver.

Conceptually:

```text
Sender G2
   │
   │ send 42
   ▼
hchan
   │
   │ recvq
   ▼
sudog
   │
   ▼
G1
```

The runtime performs the handoff:

```text
42
 │
 ▼
receiver's waiting memory
```

Then:

```text
sudog removed from recvq
        │
        ▼
G1 made runnable
```

---

# 8. Direct Handoff

This is one of the most important ideas.

For an unbuffered channel:

```go
ch := make(chan int)

sender:   ch <- 42
receiver: <-ch
```

there is no channel buffer holding the value.

Instead, sender and receiver can synchronize directly.

Conceptually:

```text
             hchan
              │
       ┌──────┴──────┐
       ▼             ▼
    sender         receiver
      G2              G1
       │               ▲
       └──── 42 ───────┘
```

The `sudog` participates in representing the blocked party.

This is a major reason unbuffered channel communication can be efficient despite requiring synchronization.

---

# 9. `sendq` and `recvq`

For channels:

```text
sendq                  recvq
  │                       │
  ▼                       ▼
sudog → sudog → ...      sudog → sudog → ...
```

### `sendq`

Contains goroutines waiting to send.

Example:

```go
ch <- value
```

when the receiver isn't ready.

### `recvq`

Contains goroutines waiting to receive.

Example:

```go
value := <-ch
```

when no value is available.

---

# 10. Example With Multiple Goroutines

Suppose:

```go
ch := make(chan int)

go func() {
    <-ch
}()

go func() {
    <-ch
}()

go func() {
    <-ch
}()
```

Eventually the runtime can have:

```text
hchan
 │
 └── recvq
       │
       ▼
    sudog(G1)
       │
       ▼
    sudog(G2)
       │
       ▼
    sudog(G3)
```

Three goroutines are blocked.

When senders arrive:

```text
send 10
```

one waiting receiver is selected.

Then:

```text
recvq

sudog(G1) → sudog(G2) → sudog(G3)

     ↓

sudog(G2) → sudog(G3)
```

and `G1` becomes runnable.

---

# 11. Why Not Just Queue `*g`?

A natural question:

> Why not simply have `recvq []*g`?

Because synchronization requires more information than merely:

```text
"G1 is waiting"
```

The runtime may need to associate the waiting goroutine with:

- the synchronization object;
    
- the memory location involved;
    
- queue linkage;
    
- wait timing;
    
- wake-up state;
    
- select-related state;
    
- other synchronization bookkeeping.
    

So:

```text
G
```

represents **execution state**.

Whereas:

```text
sudog
```

represents **participation of that G in a synchronization wait**.

This distinction is extremely important.

---

# 12. One Goroutine Can Have Multiple Wait Relationships

A goroutine may participate in synchronization mechanisms where the runtime needs temporary wait descriptors.

For example:

```go
select {
case x := <-ch1:
case x := <-ch2:
case ch3 <- value:
}
```

The runtime may need to register the goroutine against multiple possible communication cases.

Conceptually:

```text
                 G1
              /   |   \
             /    |    \
            ▼     ▼     ▼
         sudog1 sudog2 sudog3
           │      │      │
           ▼      ▼      ▼
          ch1    ch2    ch3
```

This is one reason `sudog` is more than simply:

```text
queue node = goroutine pointer
```

It is a synchronization-specific runtime object.

---

# 13. `sudog` Lifecycle

A useful mental model is:

```text
      acquire
         │
         ▼
      sudog
         │
         ▼
  initialize wait state
         │
         ▼
    enqueue in queue
         │
         ▼
       park G
         │
         ▼
    synchronization
       occurs
         │
         ▼
    dequeue sudog
         │
         ▼
      wake G
         │
         ▼
      reuse sudog
```

The runtime doesn't necessarily allocate a completely new `sudog` from the general heap every time.

The runtime maintains mechanisms for acquiring and reusing these objects.

That's important for performance because synchronization can happen extremely frequently.

---

# 14. Why Reuse Matters

Imagine a server with:

```text
100,000 goroutines
```

and heavy channel synchronization.

If every block/unblock cycle caused ordinary heap allocation:

```text
goroutine blocks
    ↓
heap allocation
    ↓
goroutine wakes
    ↓
object becomes garbage
```

you would introduce unnecessary:

- allocations;
    
- GC pressure;
    
- memory traffic;
    
- synchronization overhead.
    

Runtime-managed reuse helps avoid this cost.

---

# 15. `sudog` Is Not a Goroutine

This distinction should be firmly memorized.

```text
G
├── stack
├── program counter
├── scheduling state
├── registers/context
└── execution state


sudog
├── pointer to G
├── wait-queue links
├── synchronization metadata
└── operation-specific state
```

Therefore:

```text
sudog ≠ G
```

Instead:

```text
sudog ─────► G
```

---

# 16. Relationship With the Scheduler

The complete blocking path is approximately:

```text
G
│
│ channel receive
▼
runtime channel operation
│
│ no value available
▼
sudog
│
│ enqueue
▼
hchan.recvq
│
│
▼
G becomes waiting
│
▼
scheduler chooses another runnable G
```

Later:

```text
sender
   │
   ▼
hchan.recvq
   │
   ▼
sudog
   │
   ▼
G
   │
   ▼
make runnable
   │
   ▼
scheduler
```

So there are two separate concepts:

### Synchronization

```text
hchan + sudog
```

### Execution scheduling

```text
G + P + M + run queues
```

The `sudog` is the bridge between them.

---

# 17. Critical Distinction: Runnable vs Waiting

A goroutine blocked on a channel isn't simply:

```text
G.state = "sleeping"
```

in isolation.

The runtime has to maintain a relationship:

```text
waiting G
   │
   ▼
sudog
   │
   ▼
synchronization object
```

This enables:

```text
event occurs
   ↓
find waiting participant
   ↓
wake corresponding G
```

That is the architectural reason wait queues exist.

---

# 18. Channel vs `sync.Mutex`

`sudog` is not exclusively a channel concept.

The runtime uses `sudog` machinery in synchronization paths beyond channels.

For example, conceptually:

```text
Mutex
  │
  ▼
waiting goroutines
  │
  ▼
sudog structures
  │
  ▼
G
```

Similarly, runtime synchronization primitives can use the same fundamental waiting/wakeup infrastructure.

So a better mental model is:

> **`sudog` is runtime synchronization infrastructure, while `hchan` is one consumer of that infrastructure.**

---

# 19. Queue Discipline

A wait queue can be viewed as:

```text
head
 │
 ▼
┌────────┐
│ sudog1 │
└───┬────┘
    │ next
    ▼
┌────────┐
│ sudog2 │
└───┬────┘
    │
    ▼
┌────────┐
│ sudog3 │
└────────┘
    ▲
    │
   tail
```

Runtime queue operations must maintain invariants such as:

```text
head == nil  ⇔  tail == nil
```

and linked-list consistency:

```text
head.prev == nil
tail.next == nil
```

The exact runtime implementation should be studied against the Go version you're targeting because internals evolve.

---

# 20. Locking and Concurrency

The queue itself cannot safely be modified by arbitrary goroutines simultaneously.

Channel operations involve synchronization around the channel's internal state.

Conceptually:

```text
goroutine A
     │
     ▼
 acquire channel synchronization
     │
     ▼
 modify sendq/recvq
     │
     ▼
 release
```

This protects invariants such as:

```text
enqueue/dequeue consistency
buffer state
sender/receiver matching
closed state
```

The important architectural point is:

> **The wait queue is shared mutable runtime state, so queue manipulation must participate in the synchronization protocol of the owning primitive.**

---

# 21. Why `sudog` Is a Great Runtime Design

Without a dedicated wait descriptor:

```text
channel
   ↓
goroutine
```

you would have to overload `G` with synchronization-specific state.

That creates poor separation of responsibilities.

Instead:

```text
G
│
│ execution
▼
Scheduler


sudog
│
│ synchronization participation
▼
Wait queue


hchan
│
│ communication
▼
Channel
```

This separation is elegant.

---

# 22. Performance Model

A channel operation potentially involves:

```text
user code
   ↓
runtime entry
   ↓
channel state inspection
   ↓
queue manipulation
   ↓
value transfer
   ↓
goroutine state transition
   ↓
scheduler
```

The expensive case is usually not merely:

```text
pointer manipulation
```

but potentially:

```text
contention
+ synchronization
+ goroutine parking
+ goroutine wakeup
+ scheduler work
+ cache-line traffic
```

Therefore:

> Don't optimize channels by looking only at the `sudog` structure size.

You need to consider the **entire synchronization and scheduling path**.

---

# 23. Common Misconceptions

### ❌ `sudog` is the blocked goroutine

No.

```text
sudog → G
```

---

### ❌ Every channel operation allocates a `sudog`

No.

A `sudog` is relevant primarily when a goroutine must participate in a wait/synchronization path, and runtime mechanisms reuse these objects.

---

### ❌ Buffered channels always use `sudog`

Not necessarily.

If a buffered channel has immediately available capacity/data:

```go
ch <- x
```

or:

```go
x := <-ch
```

can proceed without parking the goroutine.

`sudog` becomes relevant when waiting is required.

---

### ❌ `sudog` is part of the channel buffer

No.

```text
hchan
├── buffer → actual channel elements
├── sendq  → waiting senders (sudog)
└── recvq  → waiting receivers (sudog)
```

Buffer data and wait metadata are separate concepts.

---

# 24. The Most Important Mental Model

Memorize this:

```text
                 Synchronization Object
                         │
                    wait queue
                         │
                       sudog
                         │
                         ▼
                         G
                         │
                         ▼
                     scheduler
```

Or even more compactly:

```text
sudog = G's synchronization wait ticket
```

It tells the runtime:

> "This particular goroutine is waiting here, for this synchronization operation, with this associated state."

---

# 25. Principal-Level Insight

The deeper architectural lesson is not really about `sudog`.

It is about **separating execution state from synchronization state**.

A scalable runtime needs to answer two different questions:

### Execution question

> Which goroutine can execute right now?

Handled through:

```text
G + P + M + run queues
```

### Synchronization question

> Which goroutines are waiting for this event?

Handled through:

```text
synchronization primitive
        ↓
wait queue
        ↓
sudog
        ↓
G
```

This separation allows Go to cheaply support enormous numbers of blocked goroutines without requiring one OS thread per blocked goroutine.

---

## Final Architecture

```text
                    Go Program
                        │
                        ▼
                  Channel Operation
                        │
              ┌─────────┴─────────┐
              │                   │
          can proceed?          cannot
              │                   │
             YES                  ▼
              │              acquire sudog
              │                   │
              │                   ▼
              │              enqueue into
              │              sendq/recvq
              │                   │
              │                   ▼
              │                 park G
              │                   │
              │                   ▼
              │               Scheduler
              │
              │
              │        Later: matching operation
              │                   │
              │                   ▼
              │              dequeue sudog
              │                   │
              │                   ▼
              │                wake G
              │                   │
              └───────────────────┘
```

### Key takeaway

```text
G       = execution entity
sudog   = synchronization wait descriptor
hchan   = channel state + wait queues
scheduler = decides who executes
```

Understanding this relationship is a prerequisite for going deeper into **channel send/receive internals, `select`, `gopark/goready`, `waitq`, `notifyList`, `sync.Mutex`, and Go's scheduler**.
---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
