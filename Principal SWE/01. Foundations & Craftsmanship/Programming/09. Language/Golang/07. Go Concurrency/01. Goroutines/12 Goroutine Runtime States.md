---
title: "12 Goroutine Runtime States"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Goroutines]]"
---
# 12 Goroutine Runtime States

To understand goroutines at a **runtime level**, first separate two concepts:

- **Goroutine (`G`)** — Go runtime's logical execution unit.
    
- **OS thread (`M`)** — actual operating-system thread that executes Go code.
    
- **Processor (`P`)** — runtime resource required for a goroutine to execute Go code.
    

A goroutine can move between several runtime states as the scheduler manages it.

> **Important:** Go does not expose a simple public enum called “12 goroutine states” as an API. The runtime internally has a state machine, and the exact set/naming can change between Go versions. The commonly discussed runtime states are the following.

---

## 1. `_Gidle`

**Meaning:** Goroutine structure has been allocated but is not yet initialized for execution.

Mental model:

```text
G allocated
   ↓
_Gidle
   ↓
initialized
```

This is primarily an **internal runtime state**.

You normally won't observe a goroutine remaining in `_Gidle`.

---

## 2. `_Grunnable`

The goroutine is **ready to execute**, but currently isn't executing on an `M`.

```text
        P
        │
   run queue
        │
        ▼
       G
   _Grunnable
```

Example:

```go
go worker()
```

The runtime creates/schedules the goroutine. It becomes runnable and eventually gets picked by a scheduler.

Important distinction:

```text
Runnable != Running
```

A runnable goroutine is waiting for CPU/runtime scheduling.

---

## 3. `_Grunning`

The goroutine is **currently executing Go code** on an OS thread.

Conceptually:

```text
G (_Grunning)
      │
      P
      │
      M
      │
      CPU
```

This is the state people usually mean by:

> "The goroutine is running."

However, only a limited number of goroutines can actually run simultaneously.

If there are:

```text
GOMAXPROCS = 4
```

then at most approximately **4 goroutines execute Go code simultaneously**.

There can still be millions of goroutines in `_Grunnable` or waiting states.

---

# 4. `_Gsyscall`

The goroutine is executing a **system call**.

For example:

```go
syscall(...)
```

The important distinction is:

```text
_Gsyscall
```

means the goroutine has entered a syscall and the runtime treats its execution differently from ordinary Go code.

Conceptually:

```text
G
│
├── syscall
│
└── _Gsyscall
```

The runtime can potentially detach the `P` from the thread performing the syscall so other goroutines can continue executing.

This is one of the reasons Go can tolerate blocking syscalls much better than a naive one-thread-per-task model.

---

# 5. `_Gwaiting`

The goroutine is **blocked/waiting** for some event.

Examples:

```go
<-ch
```

or:

```go
mutex.Lock()
```

or:

```go
select {
case <-ctx.Done():
}
```

Conceptually:

```text
Running
   │
   ▼
Waiting
   │
   │ event happens
   ▼
Runnable
```

A waiting goroutine consumes essentially **no CPU** while it is blocked.

This is one of the core reasons goroutines are cheap compared with OS threads.

---

# 6. `_Gdead`

The goroutine has **finished execution** and is dead.

Example:

```go
go func() {
    fmt.Println("done")
}()
```

After the function returns:

```text
Running
   │
   ▼
_Gdead
```

A dead goroutine cannot execute again.

The runtime may later **reuse the underlying `G` structure**, which is an implementation optimization.

---

# 7. `_Genqueue`

This is an **internal transition state** associated with enqueueing a goroutine.

Think of it as:

```text
G
│
├── preparing to enter run queue
│
└── _Genqueue
       │
       ▼
   _Grunnable
```

You generally don't reason about this state when writing application code.

It exists to maintain runtime invariants while manipulating scheduler queues.

---

# 8. `_Gcopystack`

This state is associated with **goroutine stack copying**.

Go goroutine stacks are dynamically managed.

For example, a goroutine might begin with a relatively small stack:

```text
G stack

[ small ]
```

and grow when necessary:

```text
[ small ]
    ↓
stack growth
    ↓
[ larger ]
```

During stack copying, the runtime needs to ensure that the goroutine isn't concurrently executing against the stack being relocated.

Hence the internal state:

```text
_Gcopystack
```

This is one of the implementation details behind Go's dynamically growing goroutine stacks.

---

# 9. `_Gpreempted`

The goroutine has been **preempted**.

This is particularly important for understanding modern Go scheduling.

Imagine:

```go
for {
    compute()
}
```

Historically, Go relied heavily on cooperative safe points. Modern Go has **asynchronous preemption**, allowing the runtime to interrupt long-running goroutines at safe points.

Conceptually:

```text
_Grunning
     │
     │ preemption
     ▼
_Gpreempted
     │
     ▼
_Grunnable
```

The goroutine isn't dead.

It simply lost the CPU so another goroutine can make progress.

---

# 10. `_Gscan`

This is related to the runtime's **garbage collector stack scanning**.

The GC needs to inspect goroutine stacks to find pointers.

Conceptually:

```text
G stack
   │
   ▼
GC scans stack
   │
   ▼
pointer discovery
```

During certain runtime operations, the goroutine's normal state is combined with a scan-related state.

You'll often encounter runtime states represented internally using scan bits such as:

```text
_Gscanrunnable
_Gscanrunning
_Gscanwaiting
_Gscansyscall
```

So `_Gscan*` should not be thought of as an ordinary application-level lifecycle state.

It is more accurate to think:

> **"The runtime is performing GC-related scanning/coordination for this G."**

---

# 11. `_Gscanrunnable`

This is a combination of:

```text
Runnable
+
GC scanning
```

Conceptually:

```text
       G
       │
       ├── runnable
       │
       └── being coordinated/scanned by GC
```

It exists because the runtime needs to represent both:

1. the goroutine's scheduler state, and
    
2. GC coordination state.
    

This is an important distinction when reading Go runtime source code.

---

# 12. `_Gscanwaiting`

Similarly:

```text
_Gscanwaiting
```

represents a goroutine that is fundamentally in a **waiting** state while also participating in GC stack-scanning coordination.

Conceptually:

```text
_Gwaiting
     +
GC scan
     ↓
_Gscanwaiting
```

Again, this is primarily relevant when studying runtime internals rather than application-level concurrency.

---

# The More Useful Mental Model

Instead of memorizing 12 names, I recommend grouping them.

### Execution states

```text
_Gidle
_Grunnable
_Grunning
_Gsyscall
_Gwaiting
_Gdead
```

### Runtime transition/management states

```text
_Genqueue
_Gcopystack
_Gpreempted
```

### GC scanning states

```text
_Gscan*
```

That gives you a much better mental model:

```text
                  ┌───────────────┐
                  │   _Gwaiting   │
                  └───────┬───────┘
                          │ wake
                          ▼
┌──────────────┐     ┌──────────────┐
│ _Grunnable   │────►│  _Grunning   │
└──────────────┘     └──────┬───────┘
       ▲                    │
       │                    ├──── syscall ───► _Gsyscall
       │                    │
       │                    ├──── block ─────► _Gwaiting
       │                    │
       │                    ├──── preempt ───► _Gpreempted
       │                    │
       │                    └──── return ─────► _Gdead
       │
       └──────── _Gpreempted
```

---

# The Critical Distinction

The most important thing to understand as a Go engineer is:

```text
Goroutine state
        ≠
OS thread state
```

For example:

```text
G1 ── runnable
G2 ── waiting
G3 ── runnable
G4 ── syscall
G5 ── waiting
G6 ── running
```

The runtime may have:

```text
6 Goroutines
3 OS Threads
2 Ps
```

or:

```text
1,000,000 Goroutines
8 Ps
10–20 OS Threads
```

depending on workload and runtime behavior.

This is the foundation for understanding the **G-M-P scheduler**.

---

# Production-Level Insight

When debugging goroutine problems, don't ask only:

> "How many goroutines do I have?"

Ask:

```text
How many are runnable?
How many are blocked?
What are they blocked on?
Are they making progress?
Are goroutines accumulating?
Are they stuck waiting for locks?
Are they waiting for channels?
Are they blocked on network I/O?
Are they stuck in syscalls?
Are they being preempted?
```

For example, this is potentially dangerous:

```text
10,000 goroutines
        │
        ▼
   9,900 waiting
        │
        ▼
waiting for one channel
        │
        ▼
producer stopped
```

The problem isn't necessarily:

```text
"10,000 goroutines"
```

The real problem may be:

```text
missing cancellation
      +
unbounded goroutine creation
      +
blocked consumers
```

That is a **lifecycle/backpressure problem**, not simply a goroutine-count problem.

---

## Principal Engineer Mental Model

Think of a goroutine as a **state machine managed by the runtime**:

```text
             ┌─────────────┐
             │   Runnable  │
             └──────┬──────┘
                    │ scheduled
                    ▼
             ┌─────────────┐
             │   Running   │
             └──┬─────┬────┘
                │     │
       blocking │     │ preemption
                │     │
                ▼     ▼
          ┌────────┐  ┌──────────┐
          │Waiting │  │Preempted│
          └───┬────┘  └────┬─────┘
              │            │
              └─────┬──────┘
                    ▼
                Runnable

Running ── syscall ──► Gsyscall
Running ── return ───► Dead
```

Once this mental model is clear, the next layer is **how `G`, `M`, and `P` interact with these states through the scheduler**. That is where goroutine scheduling becomes much more concrete.
---

## 🔗 References
- ⬆️ Parent: `Goroutines & Memory Lifecycle`
- 📚 Module: `Concurrency & Synchronization`
