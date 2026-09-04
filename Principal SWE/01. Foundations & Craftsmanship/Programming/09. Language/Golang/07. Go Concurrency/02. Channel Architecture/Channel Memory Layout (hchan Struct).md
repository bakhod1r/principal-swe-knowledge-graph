---
title: "Channel Memory Layout (hchan Struct)"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# Go Channel Memory Layout (`hchan`)

If you want to understand Go channels at **runtime level**, `hchan` is the key structure. A channel value in Go is small at the language level, but the runtime object behind it contains the synchronization state, buffer metadata, queue pointers, and waiting goroutines.

> **Important:** `hchan` is an **internal runtime implementation detail**, not a public Go API. Its exact fields can change between Go versions.

---

## 1. Mental Model

When you write:

```go
ch := make(chan int, 4)
```

you can mentally model it as:

```text
ch
 │
 │ pointer
 ▼
┌──────────────────────────────┐
│           hchan              │
│                              │
│ qcount = 0                   │
│ dataqsiz = 4                 │
│ buf ──────────────────────┐  │
│ sendx = 0                 │  │
│ recvx = 0                 │  │
│                            │  │
│ sendq ──► waiting senders  │  │
│ recvq ──► waiting receivers│  │
│                            │  │
│ lock                       │  │
└────────────────────────────┼──┘
                             │
                             ▼
                    ┌─────────────────┐
                    │ channel buffer  │
                    ├─────────────────┤
                    │ int             │
                    │ int             │
                    │ int             │
                    │ int             │
                    └─────────────────┘
```

The important idea:

> **The channel variable is not the buffer. It refers to runtime state that manages the buffer and synchronization.**

---

# 2. Conceptual `hchan`

The runtime's structure is roughly conceptually like:

```go
type hchan struct {
    qcount   uint
    dataqsiz uint

    buf      unsafe.Pointer
    elemsize uint16
    closed   uint32

    timer    *timer
    elemtype *_type

    sendx    uint
    recvx    uint

    recvq    waitq
    sendq    waitq

    lock     mutex
}
```

This is **conceptual**, not something application code should depend on. Field layout and additional implementation details can differ across Go releases.

---

# 3. What Each Field Means

## `qcount`

Number of elements currently stored in the channel buffer.

For:

```go
ch := make(chan int, 4)
```

initially:

```text
qcount = 0
```

After:

```go
ch <- 10
ch <- 20
```

you have:

```text
qcount = 2
```

---

## `dataqsiz`

The channel's buffer capacity.

```go
ch := make(chan int, 4)
```

means:

```text
dataqsiz = 4
```

For an unbuffered channel:

```go
ch := make(chan int)
```

conceptually:

```text
dataqsiz = 0
```

This distinction is fundamental:

```text
buffered channel:
sender → buffer → receiver

unbuffered channel:
sender ─────────→ receiver
```

---

# 4. `buf`

`buf` points to the channel's circular buffer.

For:

```go
ch := make(chan int, 4)
```

conceptually:

```text
hchan
  │
  └── buf
       │
       ▼
     ┌────┬────┬────┬────┐
     │ 10 │ 20 │    │    │
     └────┴────┴────┴────┘
```

The buffer contains **elements**, not goroutines.

For example:

```go
ch := make(chan int, 4)

ch <- 10
ch <- 20
```

conceptually:

```text
qcount = 2

buffer:
┌────┬────┬────┬────┐
│ 10 │ 20 │    │    │
└────┴────┴────┴────┘
  ↑
 recvx

      ↑
     sendx
```

The actual runtime uses memory allocation and element metadata to locate elements.

---

# 5. `sendx`

`sendx` is the next position where a send should place an element in the circular buffer.

Suppose:

```go
ch := make(chan int, 4)

ch <- 10
ch <- 20
ch <- 30
```

Conceptually:

```text
buffer:

index       0     1     2     3
           ┌─────┬─────┬─────┬─────┐
           │ 10  │ 20  │ 30  │     │
           └─────┴─────┴─────┴─────┘
                         ↑
                       sendx
```

Eventually `sendx` wraps around.

That's why the channel buffer behaves like a **ring buffer**.

---

# 6. `recvx`

`recvx` identifies the next buffer position from which a receiver should consume.

Example:

```text
index       0     1     2     3
           ┌─────┬─────┬─────┬─────┐
           │ 10  │ 20  │ 30  │     │
           └─────┴─────┴─────┴─────┘
             ↑
           recvx
```

After:

```go
x := <-ch
```

`10` is consumed and `recvx` advances.

---

# 7. Circular Buffer

The buffer does **not** shift elements after every receive.

This would be expensive:

```text
receive 10

before:
[10][20][30][40]

after naive shifting:
[20][30][40][ ]
```

Instead Go uses a ring:

```text
[10][20][30][40]
 ↑              ↑
recvx          sendx
```

Pointers/indexes advance:

```text
recvx = (recvx + 1) % dataqsiz
sendx = (sendx + 1) % dataqsiz
```

Conceptually.

Therefore:

```text
enqueue: O(1)
dequeue: O(1)
```

assuming the operation does not block.

---

# 8. `sendq`

`sendq` represents goroutines waiting to send.

Consider:

```go
ch := make(chan int, 1)

ch <- 10
ch <- 20
```

The first send succeeds:

```text
buffer:
┌────┐
│ 10 │
└────┘
```

The second sender cannot put `20` into the full buffer.

Instead, its goroutine can become blocked and enter the sender wait queue:

```text
hchan
 │
 ├── buffer
 │    └── [10]
 │
 └── sendq
      │
      ▼
     G2
```

So:

```text
sendq = blocked senders
```

---

# 9. `recvq`

Similarly, `recvq` contains goroutines waiting to receive.

For:

```go
ch := make(chan int)

x := <-ch
```

there is no buffer.

If no sender is ready:

```text
hchan
 │
 ├── buffer: none
 │
 └── recvq
      │
      ▼
     G1
```

`G1` becomes blocked waiting for a sender.

---

# 10. Wait Queues and `sudog`

The runtime does not literally store raw goroutine pointers in a simplistic list.

It uses runtime structures such as `sudog` to represent a goroutine waiting on a synchronization object.

Conceptually:

```text
hchan
 │
 ├── sendq
 │     │
 │     ▼
 │   sudog → G2
 │
 └── recvq
       │
       ▼
     sudog → G1
```

Think of `sudog` as:

> **runtime bookkeeping for a goroutine that is blocked on a synchronization primitive.**

This mechanism is used by channels and other runtime synchronization facilities.

---

# 11. `closed`

The channel maintains closed state.

Initially:

```go
ch := make(chan int, 4)
```

conceptually:

```text
closed = 0
```

After:

```go
close(ch)
```

the runtime marks it closed.

```text
closed = 1
```

But **closed does not mean the buffer is immediately empty**.

For example:

```go
ch := make(chan int, 2)

ch <- 10
ch <- 20

close(ch)
```

Conceptually:

```text
closed = true

buffer:
┌────┬────┐
│ 10 │ 20 │
└────┴────┘
```

Receivers can still consume:

```go
<-ch // 10
<-ch // 20
```

Then subsequent receives produce the zero value and indicate `ok == false`:

```go
v, ok := <-ch

// v  = 0
// ok = false
```

This is an important mental model:

> **`close` prevents future sends; it does not discard already-buffered values.**

---

# 12. `elemtype`

The runtime needs to know what type of element the channel stores.

For:

```go
chan int
```

the runtime knows the element type is `int`.

For:

```go
chan MyStruct
```

it knows the element type is `MyStruct`.

This information is important for operations involving:

- element size
    
- copying
    
- garbage collection
    
- pointer scanning
    
- zeroing
    
- memory management
    

---

# 13. `elemsize`

This represents the size of each channel element.

For example:

```go
chan int
```

has an element size corresponding to `int` on that architecture.

For:

```go
type User struct {
    ID   int64
    Name string
}

chan User
```

each buffer slot must accommodate a `User`.

Conceptually:

```text
buf
 │
 ▼
┌──────────────┐
│ User         │
├──────────────┤
│ User         │
├──────────────┤
│ User         │
└──────────────┘
```

---

# 14. `lock`

Channels are synchronization primitives, so concurrent access requires synchronization inside the runtime.

The channel has an internal runtime lock.

Conceptually:

```text
             ┌──────────────┐
G1 ──────────►              │
             │    hchan     │
G2 ──────────►    lock      │
             │              │
G3 ──────────►              │
             └──────────────┘
```

The lock protects channel state such as:

```text
qcount
sendx
recvx
sendq
recvq
closed
```

and other runtime state.

### Important distinction

This does **not** mean:

> "A channel is just a mutex."

A channel provides a much higher-level synchronization mechanism:

```text
data transfer
+
blocking
+
wakeup
+
buffering
+
ordering
+
close semantics
```

The internal lock is merely part of how the runtime implements those semantics.

---

# 15. Buffered Channel Operation

Consider:

```go
ch := make(chan int, 2)

ch <- 10
ch <- 20
```

Initial state:

```text
qcount  = 0
dataqsiz = 2
sendx   = 0
recvx   = 0

buffer:
[ ][ ]
```

### First send

```go
ch <- 10
```

Conceptually:

```text
buffer:
[10][ ]

qcount = 1
sendx  = 1
recvx  = 0
```

### Second send

```go
ch <- 20
```

```text
buffer:
[10][20]

qcount = 2
sendx  = 0
recvx  = 0
```

The `sendx` wraps around.

Now:

```text
buffer FULL
```

A third send cannot immediately insert into the buffer.

---

# 16. Receive

Now:

```go
x := <-ch
```

The runtime retrieves the element at `recvx`.

```text
recvx = 0

buffer:
[10][20]
 ↑
 receive
```

After receiving:

```text
x = 10

qcount = 1
recvx  = 1
```

Conceptually:

```text
buffer:
[ ][20]
    ↑
   recvx
```

Then another receive:

```text
x = 20

qcount = 0
recvx = 0
```

---

# 17. Unbuffered Channels Are Different

This is one of the most important distinctions.

```go
ch := make(chan int)
```

has:

```text
dataqsiz = 0
```

There is no element buffer.

Therefore:

```go
ch <- 42
```

cannot simply put `42` somewhere and continue.

A sender needs a receiver.

Conceptually:

```text
G1                          G2

sender                      receiver
  │                            │
  │       synchronization      │
  └─────────── 42 ─────────────┘
```

This is often described as a **rendezvous**.

---

# 18. Direct Handoff

The runtime can sometimes transfer a value directly between a sender and receiver rather than storing it in a buffer.

Conceptually:

```text
sender
  │
  │ value = 42
  ▼
receiver
```

instead of:

```text
sender
  │
  ▼
buffer
  │
  ▼
receiver
```

This is a critical performance/implementation concept.

For an unbuffered channel:

```text
send ↔ receive
```

is fundamentally a synchronization operation.

---

# 19. The Complete Mental Model

A useful Principal-level model is:

```text
                     CHANNEL
                        │
                        ▼
                ┌───────────────┐
                │     hchan     │
                │               │
                │ qcount        │
                │ dataqsiz      │
                │ buf ──────────┼─────► Ring Buffer
                │ sendx         │
                │ recvx         │
                │ closed        │
                │ elemtype      │
                │ elemsize      │
                │               │
                │ sendq ────────┼─────► blocked senders
                │ recvq ────────┼─────► blocked receivers
                │               │
                │ lock          │
                └───────────────┘
```

Think in four categories:

|Category|Runtime state|
|---|---|
|**Data**|`buf`, `qcount`, `dataqsiz`|
|**Position**|`sendx`, `recvx`|
|**Synchronization**|`sendq`, `recvq`, `lock`|
|**Type/lifecycle**|`elemtype`, `elemsize`, `closed`|

---

# 20. What Happens During `ch <- value`?

At a high level, the runtime considers cases roughly like:

```text
              send
               │
               ▼
       Is channel closed?
          /          \
        yes           no
        │              │
      panic            │
                       ▼
             Is receiver waiting?
                /          \
              yes           no
              │              │
        direct handoff      │
                             ▼
                     Is buffer available?
                       /          \
                     yes           no
                     │              │
                  buffer           block
                  enqueue          sender
                                   ↓
                                  sendq
```

This is not literal source-level pseudocode, but it is an excellent mental model.

---

# 21. What Happens During `<-ch`?

Similarly:

```text
              receive
                 │
                 ▼
        Is buffered value available?
             /          \
           yes           no
           │              │
       dequeue            │
                          ▼
                  Is sender waiting?
                    /          \
                  yes           no
                  │              │
             handoff/block     closed?
                                /   \
                              yes     no
                              │        │
                          zero value   block
                                       recvq
```

Again, the exact runtime path depends on channel state and channel type.

---

# 22. Why `hchan` Matters for Performance

Understanding `hchan` explains why these operations have different costs:

```go
ch <- x
x := <-ch
```

Possible costs include:

```text
                    Cost
                     │
        ┌────────────┼─────────────┐
        ▼            ▼             ▼
      memory       locking       scheduler
      copy         contention    parking/wakeup
```

A channel operation may involve:

- element copying
    
- synchronization
    
- cache-line contention
    
- goroutine parking
    
- goroutine wakeup
    
- scheduler activity
    
- garbage collector interaction
    

Therefore:

> **"Channels are lightweight" does not mean "channel operations are free."**

---

# 23. Common Misconceptions

### ❌ "The channel contains goroutines."

Not exactly.

The `hchan` maintains queues of runtime wait structures associated with blocked goroutines.

---

### ❌ "Buffered channels eliminate synchronization."

No.

Buffered channels reduce the need for sender/receiver rendezvous, but the runtime still synchronizes access to channel state.

---

### ❌ "Closing a channel deletes its buffer."

No.

Buffered values remain available after close.

---

### ❌ "The channel variable itself contains the whole channel."

Think instead:

```text
channel value
     │
     ▼
 runtime channel object
     │
     ├── metadata
     ├── buffer
     ├── wait queues
     └── synchronization state
```

---

### ❌ "A channel is basically a queue."

Incomplete.

A buffered channel contains a queue-like ring buffer, but the **channel abstraction** additionally provides:

```text
queueing
+
synchronization
+
blocking
+
wakeups
+
ordering
+
close semantics
```

---

# 24. Production-Level Mental Model

When debugging channel-heavy Go systems, don't just think:

```text
"goroutine sends to channel"
```

Think:

```text
Goroutine
   │
   ▼
channel operation
   │
   ├── receiver immediately available?
   │       └── direct handoff
   │
   ├── buffer has capacity?
   │       └── enqueue
   │
   └── otherwise
           │
           ▼
        wait queue
           │
           ▼
      goroutine parked
           │
           ▼
      later awakened
```

That model explains a large number of real-world problems:

- goroutine leaks
    
- deadlocks
    
- blocked producers
    
- blocked consumers
    
- backpressure
    
- scheduler activity
    
- unexpected latency
    
- channel contention
    
- memory retention
    

---

## 25. The Most Important Insight

Don't memorize every `hchan` field.

Instead remember:

> **`hchan` is the runtime coordination object behind a Go channel.**

It combines three fundamental things:

```text
                 hchan
                   │
       ┌───────────┼───────────┐
       ▼           ▼           ▼
     DATA       WAITING      STATE
     BUFFER     GOROUTINES   LIFECYCLE
       │           │           │
    sendx       sendq       closed
    recvx       recvq       elemtype
    qcount                  ...
```

From this mental model, you can derive why Go channels behave the way they do rather than memorizing surface-level rules.

### Next concepts worth studying

The natural next step is **`runtime.chansend` and `runtime.chanrecv` internals**—especially how the runtime moves a goroutine between **running → waiting → runnable**, how `sudog` works, and how buffered vs. unbuffered channels take different fast/slow paths.

---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
