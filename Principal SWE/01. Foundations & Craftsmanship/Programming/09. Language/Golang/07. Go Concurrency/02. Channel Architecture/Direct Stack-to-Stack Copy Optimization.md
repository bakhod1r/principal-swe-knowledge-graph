---
title: "Direct Stack-to-Stack Copy Optimization"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# Direct Stack-to-Stack Copy Optimization

**Direct Stack-to-Stack Copy** is a Go runtime optimization used in **synchronous channel communication**, especially when one goroutine sends directly to another waiting goroutine.

The key idea is:

> Instead of copying the value through the channel's buffer or allocating an intermediate object, the runtime can copy the value **directly from the sender's stack to the receiver's stack**.

This is a runtime-level optimization—not a normal Go language feature.

---

## 1. The Problem

Consider:

```go
ch := make(chan int)

go func() {
    x := 42
    ch <- x
}()

y := <-ch
```

Conceptually, we need:

```text
Sender goroutine                 Receiver goroutine

stack                            stack
┌─────────────┐                  ┌─────────────┐
│ x = 42      │                  │ y           │
└──────┬──────┘                  └──────▲──────┘
       │                                │
       └────────── copy ────────────────┘
```

For an **unbuffered channel**, there is no ring buffer:

```text
sender stack ───────► receiver stack
```

The runtime can therefore avoid an unnecessary intermediate storage location.

---

# 2. Mental Model

There are three fundamentally different channel paths.

### Buffered channel

```text
sender
  │
  ▼
┌──────────────┐
│ channel buf  │
└──────┬───────┘
       │
       ▼
   receiver
```

The value must pass through the channel's buffer.

### Unbuffered channel — sender arrives first

```text
sender
  │
  ▼
waiting receiver metadata
  │
  ▼
receiver
```

The sender may park until a receiver arrives.

### Unbuffered channel — receiver already waiting

```text
sender stack
     │
     │ direct copy
     ▼
receiver stack
```

This is the interesting optimization.

---

# 3. Where `sudog` Fits

Go's runtime represents a waiting goroutine using a `sudog`.

Conceptually:

```text
hchan
 │
 ├── sendq ── sudog ── sender G
 │
 └── recvq ── sudog ── receiver G
```

A waiting receiver has information describing where its received value should ultimately go.

Conceptually:

```text
receiver sudog
      │
      ├── waiting G
      │
      └── elem ─────────► receiver's stack slot
```

When a sender arrives, the runtime can discover:

```text
sender value
    │
    │
    ▼
receiver destination
```

and perform the copy directly.

---

# 4. Runtime-Level Flow

For:

```go
ch <- x
```

the runtime eventually reaches channel send machinery roughly along the path:

```text
chansend
   │
   ▼
receiver waiting?
   │
   ├── no ──► buffer / sender queue
   │
   └── yes
        │
        ▼
     direct handoff
        │
        ▼
     typedmemmove
        │
        ▼
 receiver's destination
```

The exact implementation details are version-dependent, but the important architectural distinction is stable:

> If a matching receiver already exists, the runtime can bypass the channel buffer entirely.

---

# 5. Why "Lockless" Needs Care

The phrase **"lockless direct memmove"** can be misleading.

The **copy itself** does not require a mutex around the memory operation.

But channel matching still involves synchronization.

For example:

```text
hchan lock
   │
   ├── inspect recvq
   ├── remove waiting receiver
   ├── establish handoff
   └── unlock
           │
           ▼
       copy value
```

So don't build the mental model:

> "Channel direct handoff = no synchronization."

The correct model is:

> **Channel coordination is synchronized, while the actual value copy can directly move memory between the goroutines' storage locations without an intermediate channel buffer.**

That's an important distinction.

---

# 6. `memmove` Is the Fundamental Operation

For ordinary values, the runtime eventually needs something conceptually equivalent to:

```text
memmove(destination, source, size)
```

For example:

```go
type Message struct {
    ID      int64
    Status  uint32
    Payload [128]byte
}
```

If:

```go
ch := make(chan Message)
```

and a receiver is already waiting:

```text
sender stack                         receiver stack

Message                             Message
┌───────────────┐                   ┌───────────────┐
│ ID            │                   │ ID            │
│ Status        │                   │ Status        │
│ Payload       │                   │ Payload       │
└───────┬───────┘                   └───────▲───────┘
        │                                   │
        └──────────── memmove ─────────────┘
```

There is no:

```text
sender
  ↓
heap temporary
  ↓
channel buffer
  ↓
receiver
```

path.

---

# 7. Why This Matters

Consider a large value:

```go
type Message struct {
    Data [4096]byte
}
```

With an unbuffered channel and a waiting receiver, conceptually:

```text
sender stack
    │
    │ 4096-byte copy
    ▼
receiver stack
```

rather than:

```text
sender stack
    │
    ▼
channel buffer / intermediate storage
    │
    ▼
receiver stack
```

The latter can introduce additional memory traffic.

This matters because copying memory costs:

```text
CPU
memory bandwidth
cache bandwidth
latency
```

---

# 8. Buffered vs Unbuffered

This distinction is critical.

## Unbuffered

```go
ch := make(chan Message)
```

Possible direct handoff:

```text
Sender ───────────────► Receiver
          memmove
```

## Buffered

```go
ch := make(chan Message, 16)
```

If the buffer is used:

```text
Sender
  │
  ▼
┌──────────────────┐
│ channel ring buf │
└────────┬─────────┘
         │
         ▼
      Receiver
```

The channel's buffer is now part of the data path.

However, the runtime has optimizations for buffered channels too; don't reduce the entire channel implementation to "buffer always means two copies." The actual path depends on whether a receiver/sender is already waiting and whether the buffer has capacity.

---

# 9. The More Interesting Case: Stack Relocation

There is a deeper runtime issue.

Go goroutine stacks are **growable**.

A goroutine's stack can move when the runtime grows it.

Suppose:

```text
Receiver stack

old stack
0x1000 ──────────► receiver variable
```

Later:

```text
stack growth

old stack              new stack
0x1000                  0x8000
  │                        │
  └── old data ───────────►┘
```

Therefore, a runtime operation cannot casually retain an arbitrary pointer into another goroutine's stack without considering stack movement.

This is one reason the runtime has specialized machinery around channel communication and stack safety.

---

# 10. Why This Is Not a Normal Go Optimization

You cannot write:

```go
unsafeDirectStackCopy(sender, receiver)
```

and reproduce the runtime behavior safely.

The runtime knows things ordinary Go code does not:

```text
G
│
├── stack bounds
├── stack growth state
├── sudog
├── channel state
├── scheduler state
└── GC metadata
```

The runtime can therefore coordinate:

```text
channel synchronization
        +
stack ownership
        +
GC visibility
        +
goroutine scheduling
        +
typed memory copying
```

Normal application code should not attempt to duplicate this mechanism.

---

# 11. Pointer-Containing Values

Another important detail is GC correctness.

Suppose:

```go
type Message struct {
    Data *Object
}
```

The runtime can't treat every value as merely anonymous bytes.

It needs to preserve the Go garbage collector's understanding of pointers.

Conceptually:

```text
Message
┌───────────────┐
│ Data ─────────┼────► Object
└───────────────┘
```

The runtime's typed memory operations account for pointer-containing types.

This is one reason runtime channel operations are more sophisticated than:

```c
memcpy(dst, src, sizeof(T));
```

---

# 12. The Optimization Is About the Data Path

A useful mental model is:

### Naive conceptual path

```text
sender
  │
  ▼
temporary
  │
  ▼
channel
  │
  ▼
receiver
```

### Direct handoff

```text
sender
  │
  └────────────────► receiver
```

The optimization removes unnecessary intermediate storage.

---

# 13. Performance Perspective

Suppose a message contains `N` bytes.

An intermediate path might require additional memory traffic proportional to:

```text
O(N)
```

while a direct handoff requires approximately:

```text
O(N)
```

for the unavoidable copy itself.

The key improvement isn't asymptotic complexity.

It's the **constant factor**:

```text
less memory traffic
less cache pollution
less intermediate storage
potentially fewer allocations
lower latency
```

This is a classic systems optimization:

> Same algorithmic complexity, better data movement.

---

# 14. But Don't Overestimate It

Direct stack-to-stack copying does **not** mean:

```text
zero-copy
```

The receiver still needs its value.

For:

```go
ch <- x
```

the semantics require the receiver to obtain its own value.

So:

```text
direct copy
```

is not:

```text
zero-copy
```

A better phrase is:

> **single direct value transfer without an intermediate channel buffer.**

That's much more precise.

---

# 15. Important Distinction: Synchronization vs Data Transfer

A channel performs two logically different jobs:

```text
        Channel
           │
    ┌──────┴──────┐
    │             │
Synchronization  Data transfer
    │             │
 rendezvous      memmove
```

For an unbuffered channel:

```go
sender ──────── rendezvous ──────── receiver
                   │
                   ▼
               data copy
```

The rendezvous determines **who communicates with whom**.

The memory copy determines **what value is transferred**.

Keeping those concepts separate makes the runtime much easier to understand.

---

# 16. Production-Level Mental Model

When analyzing a Go channel, think in terms of these paths:

```text
                    channel operation
                           │
                 ┌─────────┴─────────┐
                 │                   │
          matching waiter?       no waiter
                 │                   │
          ┌──────┴──────┐        ┌───┴────┐
          │             │        │        │
       receiver       sender    buffer   queue/park
          │             │
          ▼             ▼
     direct handoff   direct handoff
```

The exact branches depend on whether this is send or receive and whether the channel is buffered, but this model is far more useful than memorizing individual runtime functions.

---

# 17. Common Misconceptions

### ❌ "Unbuffered channels don't copy."

They do.

```text
sender value ──copy──► receiver value
```

They simply don't use a persistent channel buffer for the handoff.

---

### ❌ "Direct handoff means zero-copy."

No.

It means:

```text
sender memory ──memmove──► receiver memory
```

There is still a copy.

---

### ❌ "Lockless means channel operations have no locks."

Incorrect.

The channel's queues/state still require synchronization.

The important optimization is avoiding an **additional synchronization/storage mechanism for the data itself**, not eliminating all synchronization.

---

### ❌ "The receiver reads the sender's stack directly."

Not normally.

The semantic result is that the receiver obtains its own value.

Conceptually:

```text
sender stack ──copy──► receiver destination
```

not:

```text
receiver ──alias──► sender stack
```

---

# 18. Staff+/Principal Insight

The deeper lesson isn't actually about channels.

It is about **data movement architecture**.

Whenever you design a concurrent system, ask:

> **How many times does this data physically move?**

For example:

```text
Producer
   ↓
temporary object
   ↓
queue
   ↓
worker
   ↓
temporary object
   ↓
network
```

could potentially become:

```text
Producer
   ↓
bounded queue
   ↓
worker
   ↓
network buffer
```

Or perhaps, under stronger constraints:

```text
Producer ─────────► Consumer
```

The important optimization dimension is:

```text
data ownership
        +
data lifetime
        +
number of copies
        +
synchronization boundaries
```

This is why runtime engineers care so much about **copy elision, ownership transfer, buffer reuse, batching, cache locality, and allocation behavior**.

---

## Key Takeaways

```text
Unbuffered channel
        │
        ▼
matching receiver already waiting?
        │
       yes
        │
        ▼
direct handoff
        │
        ▼
typed memory copy
        │
        ▼
receiver destination
```

Remember these five points:

1. **Direct stack-to-stack transfer is a runtime optimization.**
    
2. It is most relevant to **synchronous/unbuffered channel rendezvous**.
    
3. It can bypass an **intermediate channel buffer**.
    
4. It is **not zero-copy**; the value still has to be copied.
    
5. "Lockless" should not be interpreted as **"channel synchronization has no locks."**
    

The most important mental model is:

> **A Go channel is both a synchronization mechanism and a data-transfer mechanism. Direct handoff optimizes the data-transfer path while preserving the synchronization semantics.**
---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
