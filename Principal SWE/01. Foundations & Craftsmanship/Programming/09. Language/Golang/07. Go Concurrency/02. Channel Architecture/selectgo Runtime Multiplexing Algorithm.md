---
title: "selectgo Runtime Multiplexing Algorithm"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# `selectgo`: Go Runtime Multiplexing Algorithm

`selectgo` is the runtime implementation behind Go's `select` statement. It is where the compiler-generated `select` operation becomes a runtime algorithm that must safely coordinate **multiple channels, blocking goroutines, randomization, locking, wake-up, and race detection**.

The important mental model is:

> **`select` is not "check channels in source order." It constructs a temporary wait-set, probes channels in a randomized order, atomically commits to one ready operation, or parks the goroutine until one becomes ready.**

---

## 1. The Problem

Consider:

```go
select {
case v := <-ch1:
    use(v)

case ch2 <- x:
    sent()

case <-timer.C:
    timeout()

default:
    idle()
}
```

At least three things can happen:

1. `ch1` is immediately receive-ready.
    
2. `ch2` is immediately send-ready.
    
3. Neither is ready.
    

The runtime must guarantee that:

- only one case wins;
    
- channel operations remain synchronized;
    
- multiple goroutines cannot corrupt channel state;
    
- blocked goroutines eventually wake;
    
- `default` does not block;
    
- selection does not systematically favor earlier cases;
    
- closed channels behave correctly;
    
- send/receive semantics remain atomic.
    

This is substantially more complicated than a simple loop:

```go
for {
    if ch1Ready() {
        ...
    }

    if ch2Ready() {
        ...
    }
}
```

That approach introduces races between **checking** readiness and **performing** the operation.

---

# 2. Mental Model

Think of `selectgo` as a small transaction coordinator.

```text
             select
               │
               ▼
        ┌───────────────┐
        │ Build cases   │
        └───────┬───────┘
                │
                ▼
       ┌──────────────────┐
       │ Randomize probe  │
       │ order             │
       └────────┬─────────┘
                │
                ▼
       ┌──────────────────┐
       │ Lock channels    │
       └────────┬─────────┘
                │
                ▼
       ┌──────────────────┐
       │ Find ready case  │
       └───────┬──────────┘
               │
       ┌───────┴────────┐
       │                │
    ready             none
       │                │
       ▼                ▼
  commit operation   enqueue sudogs
       │                │
       ▼                ▼
   unlock channels   park goroutine
       │                │
       ▼                ▼
    return          wake + retry
```

There are actually **two different orderings** involved:

```text
pollOrder
    ↓
Which case do we inspect first?

lockOrder
    ↓
In which order do we acquire channel locks?
```

These are deliberately different concerns.

---

# 3. The Runtime Function

The core runtime function is conceptually:

```go
func selectgo(
    cas0 *scase,
    order0 *uint16,
    pc0 *uintptr,
    nsends, nrecvs int,
    block bool,
) (int, bool)
```

The exact implementation and surrounding details are version-dependent, but the important inputs are:

- `cas0` — select cases
    
- `order0` — generated ordering storage
    
- `pc0` — program-counter information used by instrumentation
    
- `nsends` — number of send cases
    
- `nrecvs` — number of receive cases
    
- `block` — whether blocking is permitted
    

The runtime receives a representation of the select cases rather than repeatedly interpreting Go source syntax.

---

# 4. Step 1 — Build the Case Representation

Each channel operation becomes an internal case.

Conceptually:

```text
scase
 ├── channel
 ├── element
 ├── operation
 └── metadata
```

For example:

```go
select {
case ch1 <- x:
case y := <-ch2:
case <-ch3:
}
```

becomes approximately:

```text
case 0: send ch1
case 1: recv ch2
case 2: recv ch3
```

Nil channels are special:

```go
var ch chan int

select {
case x := <-ch:
    ...
}
```

A nil channel can never communicate.

Therefore its case effectively becomes:

```text
case → disabled
```

This is important because nil channels are frequently used deliberately to dynamically enable/disable select cases.

---

# 5. Step 2 — Generate a Randomized Poll Order

One of the most interesting parts of `selectgo` is that it doesn't simply scan:

```text
case 0
case 1
case 2
case 3
```

It constructs a randomized **poll order**.

For:

```go
select {
case <-a:
case <-b:
case <-c:
case <-d:
}
```

the runtime may produce:

```text
pollOrder:

2 → 0 → 3 → 1
```

Another invocation might produce:

```text
1 → 3 → 0 → 2
```

The purpose is fairness.

If several cases are continuously ready, always preferring the first source-level case could cause starvation of later cases.

So conceptually:

```text
source order
     │
     ▼
[ A B C D ]
     │
     ▼
random permutation
     │
     ▼
[ C A D B ]
```

This is **not** random scheduling of goroutines.

It is randomization of the order in which select cases are considered.

---

# 6. Why Randomization Matters

Suppose:

```go
for {
    select {
    case <-fast:
        processFast()

    case <-slow:
        processSlow()
    }
}
```

If both are constantly ready and selection always scanned from the top:

```text
fast → slow
```

then `fast` could repeatedly win.

Randomized polling changes the probability distribution.

For approximately symmetric continuously-ready cases:

```text
P(fast wins) ≈ 1/2
P(slow wins) ≈ 1/2
```

For `N` continuously ready cases:

```text
P(case_i wins) ≈ 1/N
```

This should be understood as **pseudo-random fairness**, not a strict fairness guarantee.

---

# 7. Step 3 — Build the Lock Order

The runtime also creates a **lock order**.

This is separate from the poll order.

Suppose:

```text
pollOrder:

C → A → D → B
```

The channels might be locked in:

```text
A → B → C → D
```

Why?

Because if two goroutines select over overlapping channels and acquire locks in inconsistent orders, deadlock can occur.

Imagine:

```text
G1:
lock(A)
lock(B)

G2:
lock(B)
lock(A)
```

Now:

```text
G1 owns A ──────waiting for B
G2 owns B ──────waiting for A
```

Deadlock.

Therefore channel locks are acquired according to a consistent ordering.

Conceptually:

```text
poll order
    ↓
randomized selection/fairness

lock order
    ↓
deterministic synchronization/deadlock avoidance
```

This distinction is extremely important.

---

# 8. Step 4 — Lock the Channels

The runtime now acquires the relevant channel locks.

Conceptually:

```text
        select cases
             │
             ▼
      unique channels
             │
             ▼
       sorted order
             │
             ▼
       lock channels
```

Why lock multiple channels?

Because readiness and commitment must be evaluated against a stable channel state.

Without synchronization, this would be unsafe:

```text
G1                         G2

check ch1 → ready
                           receive from ch1

perform receive
```

The state changed between the check and the operation.

The runtime needs the check + commitment to behave as one synchronized operation.

---

# 9. Step 5 — First Pass: Find a Ready Case

Once channels are locked, `selectgo` examines the cases in `pollOrder`.

For example:

```text
pollOrder:

C → A → D → B
```

It checks:

```text
C ready?
  ↓ no

A ready?
  ↓ yes
```

Now `A` becomes the winner.

Importantly, the runtime can distinguish several states.

### Receive

A receive can be ready because:

```text
1. waiting sender exists
2. buffered data exists
3. channel is closed
```

### Send

A send can be ready because:

```text
1. waiting receiver exists
2. buffer has free capacity
```

A send to a closed channel is different:

```text
panic
```

---

# 10. Channel Readiness Is More Than `len()`

A common incorrect mental model is:

```text
receive ready ⇔ len(buffer) > 0
```

That's incomplete.

For a channel:

```text
sendq
recvq
buffer
closed
```

readiness may depend on all of them.

For example:

```text
unbuffered channel

sender waiting
     │
     ▼
receiver can immediately receive
```

There may be no buffer at all.

Likewise:

```text
buffered channel

qcount > 0
```

means a receive can proceed even if no sender is waiting.

And:

```text
closed == true
```

makes receive immediately complete.

---

# 11. Fast Path: Immediate Winner

If a ready case is found:

```text
find ready case
       │
       ▼
commit operation
       │
       ▼
unlock channels
       │
       ▼
return selected index
```

The runtime doesn't park the goroutine.

For example:

```go
select {
case v := <-ch:
    fmt.Println(v)

case <-done:
    return
}
```

If `ch` already has data:

```text
selectgo
   │
   ├── ch ready
   │
   ├── receive
   │
   └── return case index
```

This is the common fast path.

---

# 12. What If Nothing Is Ready?

Now consider:

```go
select {
case <-a:
case <-b:
case c <- x:
}
```

and:

```text
a = not ready
b = not ready
c = not ready
```

If blocking is allowed:

```text
block = true
```

the runtime cannot simply spin.

That would be terrible:

```text
check
check
check
check
check
...
```

CPU usage would explode.

Instead, it prepares the goroutine to sleep.

---

# 13. `sudog`: The Waiter Object

This connects directly to Go's channel wait queues.

The runtime uses `sudog` structures to represent goroutines waiting on synchronization operations.

For a select, conceptually:

```text
goroutine G

       │
       ├──── sudog → channel A
       ├──── sudog → channel B
       └──── sudog → channel C
```

The goroutine effectively says:

> "Wake me if any one of these channel operations becomes possible."

This is why select is more sophisticated than:

```go
for {
    if ready(a) ...
    if ready(b) ...
    runtime.Gosched()
}
```

It integrates directly with channel synchronization.

---

# 14. Enqueue on Multiple Channels

Suppose:

```go
select {
case <-a:
case <-b:
case <-c:
}
```

None is ready.

The runtime prepares waiters:

```text
G
│
├── sudog(A)
├── sudog(B)
└── sudog(C)
```

These are inserted into the appropriate channel wait queues.

Conceptually:

```text
channel A recvq
    │
    └── G

channel B recvq
    │
    └── G

channel C recvq
    │
    └── G
```

Now another goroutine sends:

```go
a <- value
```

The runtime can discover:

```text
receiver waiting on A
        │
        ▼
      G wakes
```

---

# 15. The Critical Problem: One Winner

But now there is a subtle problem.

The goroutine is waiting on:

```text
A
B
C
```

Suppose:

```text
A becomes ready
```

and then almost simultaneously:

```text
B becomes ready
```

The goroutine must choose **exactly one** operation.

It cannot successfully receive from both.

Therefore the runtime has a selection/commit protocol around the waiting sudogs.

Conceptually:

```text
A wakes G
B wakes G

        │
        ▼
   choose winner
        │
        ▼
      A wins
        │
        ├── complete A
        │
        └── cancel/remove B,C waiters
```

The other registrations must be cleaned up.

This is one of the central complexities of `selectgo`.

---

# 16. Why `selectgo` Has Two Passes

A useful conceptual model of the algorithm is:

```text
Pass 1:
    Find immediately ready operation.

Pass 2:
    If none is ready:
        enqueue waiters,
        park,
        wake,
        determine winner,
        clean up losing waiters.
```

You can think of it as:

```text
             selectgo
                │
        ┌───────┴────────┐
        ▼                ▼
   ready now?          nothing?
        │                │
       yes              block
        │                │
        ▼                ▼
     commit          enqueue waiters
                         │
                         ▼
                       park
                         │
                         ▼
                       wake
                         │
                         ▼
                     determine
                       winner
```

---

# 17. `default` Changes Everything

Consider:

```go
select {
case <-ch:
    receive()

default:
    idle()
}
```

If `ch` isn't ready, the runtime must **not park**.

Instead:

```text
check ready
    │
    ├── yes → select case
    │
    └── no  → default
```

So:

```text
default
```

effectively means:

```text
block = false
```

after the ready-case scan fails.

This is why:

```go
select {
case <-ch:
default:
}
```

is a non-blocking channel operation.

---

# 18. Empty `select`

This is an interesting edge case:

```go
select {}
```

There are no communication cases.

The only possible semantic result is:

```text
block forever
```

So the goroutine parks permanently.

Conceptually:

```text
select {}
   ↓
no cases
   ↓
nothing can wake it
   ↓
park forever
```

---

# 19. Closed Channels

Closed channels are especially important.

Consider:

```go
close(ch)

select {
case v, ok := <-ch:
    ...
}
```

The receive is immediately ready.

The runtime can return:

```go
v  = zero value
ok = false
```

For example:

```go
v, ok := <-ch
```

after close:

```text
v  → zero(T)
ok → false
```

Therefore a closed receive participates in `select` as a ready case.

This leads to a common production issue:

```go
for {
    select {
    case <-closedCh:
        // immediately ready forever
    case <-workCh:
        process()
    }
}
```

The closed case remains permanently selectable.

If you don't return or disable it, it can dominate behavior.

---

# 20. Nil Channels

Nil is almost the opposite.

```go
var ch chan int

select {
case <-ch:
    ...
}
```

A nil channel never becomes ready.

Therefore:

```text
nil channel
     ↓
disabled select case
```

This gives Go a powerful dynamic-select pattern:

```go
var in <-chan Item

if enabled {
    in = source
}

select {
case item := <-in:
    process(item)

case <-ctx.Done():
    return
}
```

When:

```go
in = nil
```

the `in` case is effectively disabled.

This is much cleaner than maintaining separate select implementations.

---

# 21. Send to Closed Channel

Receiving from a closed channel is ready.

Sending to a closed channel is different:

```go
ch <- value
```

when:

```text
closed == true
```

causes a panic.

So:

```go
select {
case ch <- value:
    ...
}
```

does **not** mean that a closed channel simply behaves like an unavailable send.

The runtime detects the closed state and the operation panics.

This distinction matters when designing shutdown logic.

---

# 22. The Algorithm at a Higher Level

A simplified pseudocode model is:

```text
selectgo(cases):

    remove/ignore nil cases

    generate randomized poll order

    generate deterministic lock order

    lock all relevant channels

    for case in poll order:
        if case is immediately ready:
            perform case
            unlock
            return winner

    if default exists:
        unlock
        return default

    if non-blocking:
        unlock
        return no winner

    create sudogs for all blocking cases

    enqueue sudogs onto channel queues

    unlock channels

    park current goroutine

    when awakened:
        determine winning case

        reacquire required locks

        remove/cancel losing sudogs

        finish selected operation

        return winner
```

This is a **conceptual model**, not a copy of the runtime source.

The actual runtime implementation contains substantially more machinery around:

- race detection
    
- memory ordering
    
- tracing
    
- timers
    
- select case representation
    
- goroutine state
    
- waitq manipulation
    
- stack management
    
- instrumentation
    
- special-case optimizations.
    

---

# 23. Why Lock Ordering and Poll Ordering Must Be Separate

This is probably the most important algorithmic insight.

Suppose:

```text
Cases:

A
B
C
```

We want:

```text
poll order:
B → C → A
```

because randomized selection improves fairness.

But locking should be:

```text
lock order:
A → B → C
```

because deterministic ordering prevents lock-order deadlocks.

If we used randomized ordering for both:

```text
G1:
B → A

G2:
A → B
```

we could create lock inversion.

Therefore:

```text
randomization
     ↓
fairness

sorting
     ↓
deadlock avoidance
```

Different requirements → different orderings.

This is a classic systems-design pattern.

---

# 24. Complexity

For:

```text
N select cases
```

the runtime generally needs work proportional to the number of cases.

A useful high-level approximation is:

```text
O(N)
```

for scanning/constructing the selection state.

For many channels:

```text
select {
case <-ch1:
case <-ch2:
...
case <-chN:
}
```

the cost of maintaining the select itself grows with `N`.

This is one reason an enormous `select` statement is usually a design smell.

Not necessarily incorrect—but worth questioning.

---

# 25. `select` Is Not an OS `poll()`

It's tempting to think:

```text
Go select
    ≈
Linux epoll/poll
```

That's misleading.

OS-level I/O multiplexers operate over OS resources:

```text
file descriptors
sockets
devices
```

Go `select` multiplexes **Go synchronization operations**, primarily channel communication.

Conceptually:

```text
epoll
    ↓
kernel I/O readiness

select
    ↓
Go channel synchronization
```

The mechanisms are fundamentally different.

---

# 26. Select + Scheduler

When `selectgo` blocks:

```text
goroutine
    ↓
park
```

the Go scheduler can run another goroutine.

So:

```text
G1
 │
 │ select blocks
 ▼
parked

P
 │
 └── run G2
```

This is why a blocked channel operation doesn't consume an OS thread continuously.

Eventually:

```text
G2 sends
   ↓
channel wakes G1
   ↓
G1 becomes runnable
   ↓
scheduler eventually runs G1
```

The channel runtime and scheduler therefore interact closely.

---

# 27. Important Distinction: Ready vs Runnable

Suppose:

```go
select {
case x := <-ch:
}
```

Another goroutine sends:

```go
ch <- x
```

The waiting goroutine becomes eligible to run.

But:

```text
channel operation succeeds
```

doesn't necessarily mean:

```text
goroutine executes immediately
```

There are two different concepts:

```text
channel synchronization
        ↓
operation completed / goroutine made runnable

scheduler
        ↓
when CPU actually executes goroutine
```

This distinction is crucial when reasoning about latency.

---

# 28. Fairness Is Not a Scheduling Guarantee

Do not interpret randomized select order as:

> "Every case gets equal CPU time."

It doesn't guarantee that.

There are multiple layers:

```text
select fairness
      ≠
goroutine scheduler fairness
      ≠
system CPU fairness
```

Randomized selection only addresses selection among simultaneously eligible cases.

---

# 29. Common Incorrect Mental Models

### ❌ "Select checks cases in source order"

Not generally.

Think:

```text
randomized poll order
```

---

### ❌ "Select repeatedly polls channels"

No.

If nothing is ready, it parks the goroutine rather than burning CPU.

---

### ❌ "Select locks one channel"

A select involving multiple channels must coordinate access to multiple channel states.

---

### ❌ "The first ready channel always wins"

The runtime uses randomized polling to avoid deterministic source-order preference.

---

### ❌ "Closed channels are ignored"

No.

Receive from a closed channel is immediately ready.

---

### ❌ "Nil channels block the whole select"

No.

A nil channel case is disabled; other cases can still operate.

---

### ❌ "Wake-up means the goroutine immediately runs"

No.

Wake-up makes the goroutine runnable; scheduling determines execution.

---

# 30. Production Pitfall: Closed Channel Spin

This is one of the most common select bugs.

Bad:

```go
for {
    select {
    case <-done:
        cleanup()

    case item := <-jobs:
        process(item)
    }
}
```

If `done` is closed and the loop doesn't terminate:

```text
done
 ↓
always ready
 ↓
select repeatedly chooses it
 ↓
cleanup repeatedly executes
```

Potentially:

```text
100% CPU
```

Correct shutdown logic is usually:

```go
for {
    select {
    case <-done:
        return

    case item := <-jobs:
        process(item)
    }
}
```

---

# 31. Production Pitfall: Huge Selects

If you have:

```go
select {
case <-ch1:
case <-ch2:
case <-ch3:
...
case <-ch1000:
}
```

ask:

> Why does one goroutine need to multiplex 1000 independent synchronization sources?

Possible alternatives include:

```text
fan-in
worker pool
dispatcher
event loop
single aggregation channel
```

But don't automatically replace it with another abstraction.

The correct question is:

> What responsibility does this goroutine actually own?

---

# 32. Production Pattern: Cancellation

A standard pattern:

```go
select {
case item := <-jobs:
    process(item)

case <-ctx.Done():
    return ctx.Err()
}
```

This works because `ctx.Done()` is a channel.

The select creates a synchronization boundary:

```text
work available
       OR
cancellation
```

The goroutine sleeps until one happens.

---

# 33. Production Pattern: Timeout

For example:

```go
select {
case result := <-resultCh:
    return result

case <-time.After(time.Second):
    return ErrTimeout
}
```

The important runtime mental model is:

```text
resultCh
   │
   ├── ready → result
   │
   └── timeout channel ready → timeout
```

Timers integrate with channel-like synchronization.

For repeated/high-throughput code, however, prefer understanding timer lifecycle and allocation behavior rather than blindly creating `time.After` in hot loops.

---

# 34. Production Pattern: Dynamic Case Enablement

This is an advanced but elegant pattern:

```go
var input <-chan Item

if enabled {
    input = source
}

for {
    select {
    case item := <-input:
        process(item)

    case <-ctx.Done():
        return
    }
}
```

Changing:

```go
input = nil
```

changes the select topology without rebuilding the whole control structure.

Mental model:

```text
non-nil channel → enabled
nil channel     → disabled
```

---

# 35. Performance Perspective

For a select with two cases:

```go
select {
case <-a:
case <-b:
}
```

the runtime work is usually tiny compared with network/database I/O.

But if you're doing millions of extremely small operations:

```text
10M select operations/sec
```

then runtime overhead can become measurable.

At that point:

```text
benchmark
↓
pprof
↓
CPU profile
↓
allocation profile
↓
runtime source inspection
```

—not intuition.

---

# 36. A Useful Runtime-Level Diagram

```text
                    select
                      │
                      ▼
               compiler/runtime
                      │
                      ▼
                 selectgo()
                      │
          ┌───────────┴───────────┐
          │                       │
          ▼                       ▼
     pollOrder                lockOrder
     randomized               deterministic
          │                       │
          └───────────┬───────────┘
                      ▼
                lock channels
                      │
                      ▼
              scan ready cases
                      │
             ┌────────┴────────┐
             │                 │
           ready             none
             │                 │
             ▼                 ▼
         commit            default?
             │              /     \
             │            yes      no
             │             │        │
             │             ▼        ▼
             │          return   enqueue sudogs
             │                       │
             │                       ▼
             │                     park
             │                       │
             │                       ▼
             │                     wake
             │                       │
             │                       ▼
             │                 select winner
             │                       │
             └───────────┬───────────┘
                         ▼
                   cleanup losers
                         │
                         ▼
                    unlock/return
```

---

# 37. The Principal-Engineer Mental Model

Don't memorize `selectgo` as a giant runtime function.

Instead remember five responsibilities:

### 1. **Fairness**

```text
randomized poll order
```

### 2. **Synchronization**

```text
channel locks + atomic state transitions
```

### 3. **Blocking**

```text
sudog → channel wait queue → park
```

### 4. **Commit**

```text
exactly one select case wins
```

### 5. **Cleanup**

```text
losing waiters must be removed/cancelled
```

That's the real architecture.

---

# 38. The Deep Connection to Previous Channel Topics

Your previous exploration of Go channels now fits together:

```text
hchan
 │
 ├── qcount
 ├── dataqsiz
 ├── buf
 ├── sendx
 ├── recvx
 ├── sendq
 ├── recvq
 └── lock
```

Then:

```text
ordinary send/receive
        │
        ▼
channel fast path / wait queue
```

Whereas:

```text
select
   │
   ▼
selectgo
   │
   ├── inspect multiple hchan
   ├── lock multiple channels
   ├── randomized poll
   ├── enqueue multiple sudogs
   ├── park G
   ├── wake on one channel
   └── clean up remaining waiters
```

So `selectgo` is essentially the runtime's **multi-channel synchronization coordinator**.

---

## Key Takeaways

```text
select
  ↓
selectgo
  ↓
randomized poll order
  ↓
deterministic lock order
  ↓
lock channels
  ↓
find ready operation
  ↓
 ┌───────────────┐
 │ ready         │ none
 ↓               ↓
commit           enqueue sudogs
                 ↓
                 park G
                 ↓
                 wake
                 ↓
                 choose winner
                 ↓
                 cleanup losers
```

The most important insight is:

> **`select` is not fundamentally a polling construct. It is a synchronization primitive that can atomically coordinate a goroutine's interest in multiple channels and park that goroutine until one synchronization opportunity wins.**

And the three runtime concepts to connect next are:

**`selectgo` → `sudog` wait queues → `gopark/goready` scheduler transitions.**

That chain explains almost the entire blocking lifecycle of a Go `select`.

---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
