---
title: "Buffered Channel Ring Buffer Pointer Math"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# Buffered Channel Ring Buffer Pointer Math in Go

A **buffered channel** in Go uses a circular (ring) buffer internally to store elements temporarily. The key to understanding it is the relationship between:

- `qcount` — number of elements currently buffered
    
- `dataqsiz` — buffer capacity
    
- `sendx` — next position where a sender writes
    
- `recvx` — next position where a receiver reads
    

The important insight is:

> **`sendx` and `recvx` are positions, not counts. They wrap around the buffer.**

---

## 1. Mental Model

Imagine:

```text
capacity = 4

          0       1       2       3
       +-------+-------+-------+-------+
buf →  |   A   |   B   |   C   |       |
       +-------+-------+-------+-------+
           ↑                       ↑
         recvx                   sendx
```

Suppose:

```text
qcount  = 3
dataqsiz = 4
recvx   = 0
sendx   = 3
```

The buffer contains:

```text
A B C _
```

A receiver reads from `recvx = 0`.

A sender writes to `sendx = 3`.

After sending `D`:

```text
A B C D
```

and:

```text
qcount = 4
sendx  = 0
```

`sendx` wraps back to `0`.

---

# 2. Why Do We Need Two Pointers?

You could imagine storing only:

```text
head
tail
```

But Go's channel implementation needs to efficiently track:

```text
recvx → where the next receive happens
sendx → where the next send happens
qcount → how many elements exist
```

This gives O(1) enqueue/dequeue.

There is no shifting:

```text
❌ Remove A:

B C D
```

Instead:

```text
✅ Just advance recvx:

A B C D
↑
recvx

       ↓

A B C D
  ↑
recvx
```

The old slot becomes available for future sends.

---

# 3. The Core Pointer Math

For a buffer of capacity `N`:

```text
sendx = (sendx + 1) % N
recvx = (recvx + 1) % N
```

For example, capacity `4`:

```text
sendx:

0 → 1 → 2 → 3 → 0 → 1 → 2 ...
```

Similarly:

```text
recvx:

0 → 1 → 2 → 3 → 0 → 1 → 2 ...
```

This is the fundamental ring-buffer operation.

---

# 4. Example: Sending

Start:

```text
capacity = 4

buf:
[ _ ][ _ ][ _ ][ _ ]

sendx  = 0
recvx  = 0
qcount = 0
```

Send `A`:

```text
buf:
[ A ][ _ ][ _ ][ _ ]

sendx  = 1
recvx  = 0
qcount = 1
```

Send `B`:

```text
buf:
[ A ][ B ][ _ ][ _ ]

sendx  = 2
qcount = 2
```

Send `C`:

```text
buf:
[ A ][ B ][ C ][ _ ]

sendx  = 3
qcount = 3
```

Send `D`:

```text
buf:
[ A ][ B ][ C ][ D ]

sendx  = 0
qcount = 4
```

Notice:

```text
sendx == recvx
```

does **not** necessarily mean the buffer is empty.

It can mean:

```text
qcount == 0   → empty
qcount == N   → full
```

This is one reason `qcount` is important.

---

# 5. Example: Receiving

Now:

```text
buf:
[ A ][ B ][ C ][ D ]

recvx  = 0
sendx  = 0
qcount = 4
```

Receive:

```text
recvx = 0
```

returns:

```text
A
```

Then:

```text
recvx  = 1
qcount = 3
```

Logical buffer:

```text
    consumed
       ↓
[ _ ][ B ][ C ][ D ]
       ↑
     recvx
```

Receive again:

```text
B
```

Now:

```text
recvx  = 2
qcount = 2
```

---

# 6. The Ring Starts Wrapping

Suppose we continue.

After receiving:

```text
A
B
```

state:

```text
buf:
[ _ ][ _ ][ C ][ D ]

recvx  = 2
sendx  = 0
qcount = 2
```

Now send `E`.

The sender writes:

```text
buf[sendx]
```

which is:

```text
buf[0] = E
```

Then:

```text
sendx = (0 + 1) % 4
      = 1
```

Result:

```text
[ E ][ _ ][ C ][ D ]
  ↑       ↑
 sendx   recvx
```

Logical queue order is:

```text
C → D → E
```

Even though physically the memory looks like:

```text
E _ C D
```

This is the essence of a **circular buffer**.

---

# 7. Physical Order ≠ Logical Order

This is a very important systems concept.

The underlying array may be:

```text
index:

0   1   2   3
+---+---+---+---+
| E | _ | C | D |
+---+---+---+---+
```

But the logical queue is:

```text
recvx = 2

2 → 3 → 0
C → D → E
```

Therefore:

```text
physical memory order:
E _ C D

logical FIFO order:
C D E
```

Never reason about a ring buffer by looking only at contiguous memory.

Reason using:

```text
recvx
sendx
qcount
capacity
```

---

# 8. How Go Uses It Internally

Conceptually, Go's `hchan` contains fields equivalent to:

```go
type hchan struct {
    qcount   uint
    dataqsiz uint

    buf      unsafe.Pointer

    sendx    uint
    recvx    uint

    // ...
}
```

The important relationships are:

```text
0 <= sendx < dataqsiz
0 <= recvx < dataqsiz
0 <= qcount <= dataqsiz
```

For a buffered channel:

```go
ch := make(chan int, 4)
```

you can mentally model:

```text
dataqsiz = 4
```

and the channel owns storage for four elements.

---

# 9. Send Operation

Conceptually, when the channel has buffer space:

```text
buf[sendx] = value

sendx++

if sendx == dataqsiz {
    sendx = 0
}

qcount++
```

The actual runtime implementation is optimized and uses runtime-specific mechanisms, but this is the correct mental model.

Equivalent mathematical form:

```text
sendx = (sendx + 1) mod dataqsiz
```

---

# 10. Receive Operation

Conceptually:

```text
value = buf[recvx]

clear(buf[recvx])

recvx++

if recvx == dataqsiz {
    recvx = 0
}

qcount--
```

The `clear` step is important when the element contains pointers.

For example:

```go
chan *User
```

The runtime should not unnecessarily keep references to already-consumed objects alive.

So conceptually:

```text
buf[recvx] = nil
```

for pointer-containing element types.

---

# 11. Why Not `%`?

A naive implementation might write:

```go
sendx = (sendx + 1) % dataqsiz
```

But runtime code can use a branch:

```go
sendx++
if sendx == dataqsiz {
    sendx = 0
}
```

Why?

Because `%` can involve a division/modulo operation, while:

```text
increment + compare + conditional reset
```

is simple and predictable.

For a runtime hot path, eliminating unnecessary expensive operations matters.

This is a good example of:

> **First design for correctness; then optimize a measured hot path.**

---

# 12. A Complete Example

Capacity:

```text
N = 4
```

Initial:

```text
recvx  = 0
sendx  = 0
qcount = 0

[ _ ][ _ ][ _ ][ _ ]
  R
  S
```

Send `A`:

```text
[ A ][ _ ][ _ ][ _ ]
  R    S

recvx  = 0
sendx  = 1
qcount = 1
```

Send `B`:

```text
[ A ][ B ][ _ ][ _ ]
  R        S

sendx  = 2
qcount = 2
```

Receive:

```text
A

[ _ ][ B ][ _ ][ _ ]
       R       S

recvx  = 1
sendx  = 2
qcount = 1
```

Send `C`:

```text
[ _ ][ B ][ C ][ _ ]
       R       S

sendx  = 3
qcount = 2
```

Send `D`:

```text
[ _ ][ B ][ C ][ D ]
       R
           S → 0

sendx  = 0
qcount = 3
```

Send `E`:

```text
[ E ][ B ][ C ][ D ]
       R
  S

sendx  = 1
qcount = 4
```

Logical queue:

```text
B → C → D → E
```

Physical array:

```text
E B C D
```

This distinction is critical.

---

# 13. The Invariant You Should Remember

For a buffered channel:

```text
qcount = number of logically buffered elements
```

and:

```text
sendx = next insertion position
recvx = next removal position
```

Therefore:

```text
sendx = recvx
```

is ambiguous.

You need `qcount`:

```text
qcount == 0
    ↓
empty

qcount == dataqsiz
    ↓
full
```

This is a classic ring-buffer design pattern.

---

# 14. Why `sendx` and `recvx` Never Grow Forever

A common misconception is:

```text
sendx = total number of sends
recvx = total number of receives
```

No.

They are **indexes into the physical buffer**.

Therefore:

```text
0 <= sendx < capacity
0 <= recvx < capacity
```

They continuously wrap:

```text
0
↓
1
↓
2
↓
3
↓
0
↓
1
...
```

If you wanted total historical sends, that would be a different counter.

---

# 15. Important Runtime Insight

There are actually **two different buffering paths** to keep in your mental model.

### Buffered channel with waiting receiver

A sender may be able to hand the value directly to a waiting receiver rather than first placing it into the ring buffer.

Conceptually:

```text
sender
   │
   │ direct handoff
   ▼
receiver
```

instead of:

```text
sender
   │
   ▼
ring buffer
   │
   ▼
receiver
```

This is why understanding only the ring buffer is insufficient for understanding Go channels.

You also need:

```text
hchan
 ├── buffer
 ├── sendx
 ├── recvx
 ├── qcount
 ├── send queue
 └── receive queue
```

The send/receive wait queues contain `sudog` structures representing blocked goroutines.

---

# 16. Production-Level Mental Model

Think about a buffered channel as:

```text
                    hchan
                      │
       ┌──────────────┼──────────────┐
       │              │              │
    state           buffer         waiters
       │              │              │
   qcount        circular array   sudog queues
   dataqsiz
   sendx
   recvx
```

The ring buffer answers:

> **Where is the next buffered value written/read?**

The wait queues answer:

> **Which goroutines are blocked waiting to send/receive?**

The scheduler answers:

> **When should those goroutines run again?**

That separation is extremely useful when debugging channel behavior.

---

# 17. Common Mistakes

### Mistake 1 — Treating `sendx` as total sends

Wrong:

```text
sendx = 1,000,000
```

Correct:

```text
sendx ∈ [0, capacity)
```

---

### Mistake 2 — Assuming physical array order is FIFO order

Wrong:

```text
[E][B][C][D]
 ↑
 oldest
```

Correct:

```text
recvx = 1

B → C → D → E
```

---

### Mistake 3 — Assuming `sendx == recvx` means empty

Wrong.

It can mean:

```text
empty OR full
```

Use:

```text
qcount
```

to distinguish them.

---

### Mistake 4 — Assuming every send writes to the buffer

Not necessarily.

With a waiting receiver, the runtime can perform a direct sender→receiver handoff.

---

# 18. Complexity

Ring-buffer operations are:

```text
Send into buffer:    O(1)
Receive from buffer: O(1)
Index advancement:   O(1)
```

No element shifting occurs.

Compare with a naive slice-based queue:

```go
queue = queue[1:]
```

which may create different memory-retention and allocation behaviors depending on implementation.

A proper ring buffer gives predictable constant-time enqueue/dequeue behavior.

---

# 19. Principal Engineer Insight

The deeper lesson isn't merely:

```text
sendx++
recvx++
```

It's **separating logical state from physical storage**.

The same idea appears in:

- network packet rings
    
- lock-free queues
    
- NIC descriptor rings
    
- kernel buffers
    
- storage queues
    
- Kafka-like log structures
    
- bounded worker queues
    
- audio/video streaming buffers
    

You should always ask:

```text
What is the logical ordering?
Where is the physical storage?
What identifies the next read?
What identifies the next write?
How is empty distinguished from full?
What happens when the index wraps?
Who owns each slot?
What happens under concurrency?
```

That mental model transfers far beyond Go channels.

### Key formula

For capacity `N`:

```text
next_send = (sendx + 1) % N
next_recv = (recvx + 1) % N
```

with:

```text
0 <= qcount <= N
0 <= sendx < N
0 <= recvx < N
```

And the most important invariant:

```text
qcount == 0       → buffer empty
qcount == N       → buffer full
```

**Ring buffer pointer math is simple; the difficult part is understanding how that buffer interacts with blocked goroutines, `sudog` wait queues, locking, scheduler wakeups, and the memory model.**
---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
