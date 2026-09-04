---
title: "Goroutines vs OS Threads"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Goroutines]]"
---
# Goroutines vs OS Threads

## 1. Definition

### Goroutine

A **goroutine** is a lightweight unit of concurrent execution managed by the **Go runtime**.

```go
go doWork()
```

A goroutine is **not an OS thread**. The Go runtime schedules goroutines onto a much smaller pool of OS threads.

Key characteristics:

- Managed by Go runtime.
    
- Starts with a small stack that can grow/shrink.
    
- Very cheap to create compared with OS threads.
    
- Communication commonly uses channels.
    
- Scheduling is performed by Go's scheduler.
    
- You can have **thousands or millions** of goroutines, depending on workload and resource usage.
    

### OS Thread

An **OS thread** is a kernel-managed execution context provided by the operating system.

Examples:

- Linux `pthread`
    
- Windows threads
    
- macOS threads
    

The OS scheduler decides which runnable thread executes on which CPU core.

OS threads generally have:

- Larger per-thread stack reservation.
    
- Higher creation/destruction overhead.
    
- Kernel scheduling involvement.
    
- More expensive context switching.
    
- A relatively limited practical number compared with goroutines.
    

---

# 2. Mental Model

The most important mental model is:

> **Goroutines are tasks; OS threads are execution resources.**

Go uses an **M:N scheduling model**:

```text
          Goroutines
       G1  G2  G3  G4  G5
        \   |   |   |   /
         \  |   |   |  /
          Go Scheduler
               |
       +-------+-------+
       |       |       |
      M1      M2      M3       OS Threads
       |       |       |
      P1      P2      P3       Ps / processors
       |       |       |
      CPU     CPU     CPU
```

More precisely, Go's scheduler uses the **G-M-P model**:

- **G** = Goroutine
    
- **M** = Machine / OS thread
    
- **P** = Processor, a runtime scheduling resource
    

A simplified execution relationship is:

```text
G ──runs on──> M ──uses──> P ──executes on──> CPU
```

The important distinction:

### Goroutine

Represents **what should execute**.

### OS thread

Represents **where execution actually happens**.

### P

Represents the Go runtime's ability to execute Go code and provides scheduling state/resources.

---

## Why this design exists

Suppose your server receives:

```text
1,000,000 requests
```

A naive model would be:

```text
1 request → 1 OS thread
```

That would be extremely expensive.

Go instead uses:

```text
1 request → 1 goroutine
```

and approximately:

```text
many goroutines → relatively few OS threads
```

For example:

```text
100,000 goroutines
        ↓
      Go scheduler
        ↓
     8–32 OS threads
        ↓
      CPU cores
```

The exact number of threads is dynamic and depends on blocking, syscalls, cgo, runtime behavior, `GOMAXPROCS`, and other factors.

---

# 3. Usage

## Goroutines: default choice for concurrency

Use goroutines when you need many independent or concurrent activities:

```go
go handleConnection(conn)
```

Typical examples:

```text
HTTP requests
TCP connections
background workers
pipeline stages
message consumers
parallel I/O
timers
event processing
```

Example:

```go
func handle(conn net.Conn) {
    defer conn.Close()

    // process connection
}

func main() {
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        go handle(conn)
    }
}
```

You can support a large number of concurrent connections without creating one OS thread per connection.

---

## OS threads: usually an implementation detail

Most Go applications should **not manually manage OS threads**.

The runtime handles them for you.

However, there are situations where OS-thread affinity matters.

For example:

```go
runtime.LockOSThread()
```

This can be necessary when interacting with APIs requiring execution on a particular OS thread, such as certain GUI systems or thread-local OS APIs.

Example:

```go
runtime.LockOSThread()
defer runtime.UnlockOSThread()
```

This is specialized functionality, not normal application concurrency.

---

## CPU-bound work

Suppose:

```go
for i := 0; i < 1_000_000; i++ {
    go expensiveCalculation(i)
}
```

Creating huge numbers of goroutines does **not** mean they all execute simultaneously.

If:

```text
GOMAXPROCS = 8
```

roughly speaking, at most about **8 Ps can execute Go code simultaneously**.

So:

```text
1,000,000 goroutines
        ↓
scheduler
        ↓
~8 executing at once
        ↓
remaining goroutines waiting
```

This is why concurrency and parallelism must be distinguished:

> **Concurrency** = many tasks can make progress independently.

> **Parallelism** = multiple tasks execute simultaneously on multiple CPUs.

Goroutines provide a convenient concurrency abstraction. Actual parallel execution depends on available processors and runtime scheduling.

---

# 4. Gotchas

## 1. Goroutines are cheap, not free

This is a common misconception.

```go
for i := 0; i < 10_000_000; i++ {
    go work()
}
```

is not automatically a good design.

Every goroutine consumes resources:

```text
goroutine
   ↓
stack
scheduler metadata
references
channels / synchronization
heap allocations
```

Millions of goroutines can consume significant memory and scheduling resources.

Use **bounded concurrency** when the workload is large.

For example:

```text
100,000 jobs
      ↓
worker pool
      ↓
100 workers
```

rather than:

```text
100,000 jobs
      ↓
100,000 goroutines doing expensive work
```

---

## 2. Goroutine leaks

A goroutine can remain alive indefinitely if it is blocked.

Classic example:

```go
func worker(ch <-chan int) {
    for {
        value := <-ch
        process(value)
    }
}
```

If nobody closes or sends to `ch`, the goroutine may live forever.

This is a **goroutine leak**.

Production code should have clear lifecycle ownership:

```text
Who starts this goroutine?
Who stops it?
What happens on cancellation?
What happens on shutdown?
```

`context.Context` is often the right mechanism for cancellation.

---

## 3. Blocking does not necessarily mean an OS thread is permanently blocked

This is one of the most important distinctions.

For network I/O:

```go
data, err := conn.Read(buf)
```

the Go runtime can integrate network polling with goroutine scheduling.

Conceptually:

```text
G1
 |
 | waits for network
 ↓
network poller
 |
 ↓
G1 becomes runnable
```

The OS thread can potentially execute another goroutine instead of sitting idle.

This is one reason Go can handle large numbers of concurrent network connections efficiently.

---

## 4. Blocking syscalls and cgo are different

Not every blocking operation behaves like normal Go network I/O.

For example:

```text
G1
 |
blocking syscall
 |
M1
 |
kernel
```

The runtime may need another OS thread to continue executing other goroutines.

Similarly, **cgo calls can cause additional OS threads to be involved**.

Therefore:

> "I have only 8 CPUs, so my process can only have 8 OS threads."

is incorrect.

You can have substantially more OS threads than `GOMAXPROCS`.

`GOMAXPROCS` primarily controls how many processors can execute Go code simultaneously; it is **not a hard limit on the number of OS threads**.

---

## 5. Goroutine scheduling is preemptive

Older explanations sometimes say:

> "A goroutine only yields when it performs I/O or reaches a channel operation."

That mental model is outdated.

Modern Go has **asynchronous/preemptive scheduling mechanisms**, allowing the runtime to interrupt goroutines that run for too long.

Still, badly behaved CPU-bound code can affect scheduler responsiveness, and operations involving cgo or certain runtime boundaries have their own behavior.

---

## 6. `runtime.Gosched()` is not a synchronization primitive

You may see:

```go
runtime.Gosched()
```

It yields the processor so another runnable goroutine can execute.

But this:

```go
runtime.Gosched()
```

does **not** establish the synchronization guarantees provided by:

```go
sync.Mutex
sync.WaitGroup
sync/atomic
channels
```

Do not use scheduling tricks to fix races.

---

## 7. Data races are independent of goroutines vs threads

This is wrong:

> "Go has goroutines, therefore concurrent memory access is safe."

It is not.

This can race:

```go
var counter int

go func() {
    counter++
}()

go func() {
    counter++
}()
```

Use appropriate synchronization:

```go
var mu sync.Mutex
var counter int

mu.Lock()
counter++
mu.Unlock()
```

or another design such as atomics or ownership through channels.

Run:

```bash
go test -race ./...
```

The race detector is one of the most valuable tools for concurrent Go code.

---

## 8. More goroutines ≠ more performance

Consider a CPU-bound workload on an 8-core machine.

These are very different:

```text
8 useful workers
```

versus:

```text
1,000,000 CPU-bound goroutines
```

The latter introduces additional:

- scheduling
    
- memory
    
- synchronization
    
- cache pressure
    
- queueing
    

without creating additional CPU capacity.

A Principal-level question is therefore not:

> "How many goroutines can I create?"

but:

> **"What is the concurrency level that maximizes useful throughput under my bottleneck?"**

---

# The key mental model

Remember this:

```text
                Application
                    |
              Goroutines (G)
          G1 G2 G3 G4 ... G100K
                    |
               Go Scheduler
                    |
          +---------+---------+
          |         |         |
         P1        P2        P3
          |         |         |
         M1        M2        M3
          |         |         |
        CPU 1     CPU 2     CPU 3
```

And distinguish the three concepts:

|Concept|Meaning|
|---|---|
|**Goroutine**|Unit of concurrent work|
|**OS thread (M)**|Kernel execution resource|
|**P**|Go runtime resource that schedules Go code|

### Senior-level takeaway

**Goroutines solve the programming model of concurrency. OS threads provide the underlying execution resources. Go's runtime bridges the two through scheduling.**

Therefore, when debugging production concurrency, think in this order:

```text
Workload
   ↓
Concurrency
   ↓
Blocking / I/O
   ↓
Goroutine lifecycle
   ↓
Scheduler behavior
   ↓
P / GOMAXPROCS
   ↓
OS threads
   ↓
CPU / memory / I/O bottleneck
```

Do not start with **"How many threads do I need?"** In Go, the better question is usually **"How should I model the work and bound its concurrency?"**
---

## 🔗 References
- ⬆️ Parent: [[Goroutines]]
- 📚 Module: `Concurrency & Synchronization`
