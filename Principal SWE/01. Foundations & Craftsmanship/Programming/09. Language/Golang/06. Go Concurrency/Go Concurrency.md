---
title: Go Concurrency
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Golang]]"
---

# ⚡ Concurrency & Synchronization

Goroutines, channel internals (hchan), sync primitives, lock-free algorithms, context cancellation trees, the GMP scheduler, hardware memory models, and distributed resilience.

```text
Concurrency & Synchronization
│
├── `01. Goroutines & Memory Lifecycle`
│   ├── `CSP Concurrency Model in Go`
│   ├── `Goroutine Spawning Mechanics (go keyword)`
│   ├── `Goroutines vs OS Threads`
│   ├── `Initial 2KB Stack Allocation (stack.go)`
│   ├── `Contiguous Stack Growth (copystack)`
│   ├── `Stack Shrinking during GC`
│   ├── `Stack Splitting Preambles (morestack & nosplit)`
│   ├── `12 Goroutine Runtime States`
│   ├── `Goroutine Allocation Pool (gfget & gfput)`
│   ├── `Goroutine Leaks Diagnostics`
│   └── `Thread Pinning (LockOSThread & UnlockOSThread)`
├── `02. Channel Architecture & Operations`
│   ├── `Channel Send Operation (ch <- v)`
│   ├── `Channel Receive Operation (v <- ch & v, ok <- ch)`
│   ├── `Channel Close Operation (close(ch))`
│   ├── `Channel Length & Capacity (len, cap)`
│   ├── `Unbuffered Channel Synchronous Rendezvous`
│   ├── `Buffered Channel Ring Buffer Pointer Math`
│   ├── `Nil Channel Blocking & Deadlock Patterns`
│   ├── `Closed Channel Read-Write Behavior Matrix`
│   ├── `Channel Memory Layout (hchan Struct)`
│   ├── `sudog Wait Queue Architecture`
│   ├── `Direct Stack-to-Stack Copy Optimization`
│   └── `selectgo Runtime Multiplexing Algorithm`
├── `03. Synchronization Primitives (sync)`
│   ├── `sync.Mutex (Normal vs Starvation Mode)`
│   ├── `sync.RWMutex (Reader Starvation Prevention)`
│   ├── `sync.WaitGroup State Bitpacking`
│   ├── `sync.Once & Fast-Path Double-Checked Locking`
│   ├── `sync.Pool Architecture (poolLocal & Victim Cache)`
│   ├── `sync.Map Architecture (readOnly vs dirty Map)`
│   ├── `sync.Cond Condition Variables`
│   └── `Mutex vs Channel Selection Matrix`
├── `04. Lock-Free & Atomic Concurrency`
│   ├── `sync-atomic Primitives (Load, Store, CAS, Swap, Add)`
│   ├── `Atomic CAS Loop Pattern`
│   ├── [[atomic.Pointer[T] & atomic.Value Type Safety]]
│   ├── `Lock-Free Stack (Treiber Stack)`
│   ├── `Lock-Free Queue (Michael-Scott Queue)`
│   ├── `Lock-Free Ring Buffer`
│   └── `Lock-Free vs Mutex Performance Benchmarks`
├── [[Context Trees & Request Cancellation|05. Context Trees & Request Cancellation]]
│   ├── `context.Background() vs context.TODO()`
│   ├── `cancelCtx & Cancellation Tree Propagation`
│   ├── `timerCtx & Deadline Scheduling (time.AfterFunc)`
│   ├── `valueCtx & Key-Value Immutability`
│   ├── `context.WithoutCancel (Go 1.21+)`
│   ├── `context.AfterFunc (Go 1.21+)`
│   ├── `Context Memory Leaks & Resource Hygiene`
│   └── `Context Design Rules`
├── [[Advanced Concurrency Patterns|06. Advanced Concurrency Patterns]]
│   ├── `Worker Pools (Static & Elastic Scaling)`
│   ├── `Pipelines & Stream Processing Stages`
│   ├── `Fan-Out & Fan-In Multiplexing`
│   ├── `Or-Done Channel Pattern`
│   ├── `Tee Channel & Bridge Channel Patterns`
│   ├── `Singleflight Request Coalescing (x-sync-singleflight)`
│   ├── `ErrGroup Concurrency & Context Binding (x-sync-errgroup)`
│   ├── `Weighted Semaphore Resource Limiting (x-sync-semaphore)`
│   ├── `Token Bucket Rate Limiting (x-time-rate)`
│   ├── `Leaky Bucket Rate Limiting`
│   ├── `Debounce and Throttle Patterns`
│   └── `Graceful Shutdown Coordinator`
├── [[GMP Scheduler & Runtime Internals|07. GMP Scheduler & Runtime Internals]]
│   ├── `GMP Model (G, M, P, Sched Structs)`
│   ├── `Runqueue Architecture (Local vs Global)`
│   ├── `Work Stealing Algorithm`
│   ├── `Netpoller (epoll, kqueue, IOCP) Integration`
│   ├── `Syscall Handling & M Handoff`
│   ├── `Sysmon Daemon Thread`
│   ├── `Signal-Based Async Preemption (SIGURG)`
│   ├── `Thread Parking (notesleep & notewakeup)`
│   └── `GOMAXPROCS & Container CFS Quota Throttling`
├── `08. Go Memory Model & Hardware Concurrency`
│   ├── `Go Memory Model Specification`
│   ├── `Happens-Before Relationship Rules`
│   ├── `Instruction Reordering (Compiler & CPU Out-of-Order)`
│   ├── `Memory Barriers & CPU Store Buffers`
│   ├── `CPU Cache Hierarchy & Cache Lines (64-byte)`
│   ├── `Cache Coherency Protocols (MESI & MOESI)`
│   ├── `False Sharing & Cache Line Invalidation`
│   └── `Data Race vs Race Condition Deep Dive`
└── `09. Distributed Concurrency & Resilience`
│   ├── `Circuit Breaker Pattern in Go`
│   ├── `Distributed Rate Limiting (Redis Sliding Window & Token Bucket)`
│   ├── `Distributed Lock (Redlock Algorithm & Etcd Lease Fencing)`
│   ├── `Idempotency Keys & Deduplication`
│   ├── `Transactional Outbox Pattern with PostgreSQL & Kafka`
│   ├── `Transactional Inbox Pattern for Consumer Idempotency`
│   ├── `Saga Pattern in Distributed Concurrency`
│   ├── `Bulkhead Isolation Pattern`
│   └── `Distributed Consensus & Raft in Go (etcd-raft & hashicorp-raft)`
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 `01. Goroutines & Memory Lifecycle`
- `CSP Concurrency Model in Go` — Communicating Sequential Processes: Do not communicate by sharing memory; share memory by communicating.
- `Goroutine Spawning Mechanics (go keyword)` — Spawning lightweight user-space execution threads with go keyword and 2KB stack.
- `Goroutines vs OS Threads` — Comparing memory footprints (2KB vs 2MB-8MB), creation costs, and context switch overhead.
- `Initial 2KB Stack Allocation (stack.go)` — How the Go runtime allocates minimal contiguous stacks for goroutines.
- `Contiguous Stack Growth (copystack)` — Stack expansion up to 1GB and pointer fixup translation during memory reallocation.
- `Stack Shrinking during GC` — Shrinking underutilized stacks during garbage collection cycles to reclaim memory.
- `Stack Splitting Preambles (morestack & nosplit)` — Compiler-inserted stack boundary checks and //go:nosplit leaf function directives.
- `12 Goroutine Runtime States` — Dissecting _Gidle, _Grunnable, _Grunning, _Gsyscall, _Gwaiting, _Gdead, _Gcopystack, _Gpreempted.
- `Goroutine Allocation Pool (gfget & gfput)` — Reusing dead goroutine memory structures from scheduler free lists.
- `Goroutine Leaks Diagnostics` — Identifying blocked goroutines using pprof, runtime.NumGoroutine, and trace dumps.
- `Thread Pinning (LockOSThread & UnlockOSThread)` — Binding goroutines to dedicated OS threads for Cgo and graphics libraries.
### 2. 📂 `02. Channel Architecture & Operations`
- `Channel Send Operation (ch <- v)` — Synchronous vs asynchronous send execution, locking, and waking receivers.
- `Channel Receive Operation (v <- ch & v, ok <- ch)` — Reading from channels, comma-ok idiom, and unblocking senders.
- `Channel Close Operation (close(ch))` — Closing rules, panic conditions (closing closed or nil channels), and broadcast signaling.
- `Channel Length & Capacity (len, cap)` — Inspecting buffer length and total capacity with len() and cap().
- `Unbuffered Channel Synchronous Rendezvous` — Direct synchronous rendezvous signaling without intermediate buffer storage.
- `Buffered Channel Ring Buffer Pointer Math` — Circular ring buffer pointer math (sendx, recvx, qcount) in hchan.
- `Nil Channel Blocking & Deadlock Patterns` — Why reads and writes to nil channels block forever and how to use them in select.
- `Closed Channel Read-Write Behavior Matrix` — Matrix of reading zero values from closed channels vs panicking on writes.
- `Channel Memory Layout (hchan Struct)` — Dissecting hchan fields: ring buffer, lock mutex, sendq and recvq wait queues.
- `sudog Wait Queue Architecture` — How sudog structs wrap waiting goroutines and integrate with runtime pools.
- `Direct Stack-to-Stack Copy Optimization` — Lockless direct memmove between goroutine stacks bypassing intermediate channel buffer.
- `selectgo Runtime Multiplexing Algorithm` — Pseudo-random case shuffling, lock acquisition ordering, and non-blocking select.
### 3. 📂 `03. Synchronization Primitives (sync)`
- `sync.Mutex (Normal vs Starvation Mode)` — Bimodal mutex algorithm: high-throughput CPU spin vs fair FIFO handoff at 1ms latency threshold.
- `sync.RWMutex (Reader Starvation Prevention)` — Reader-writer lock with atomic reader counters and writer priority signaling.
- `sync.WaitGroup State Bitpacking` — 64-bit/128-bit atomic state bitpacking combining task counter and waiter counter.
- `sync.Once & Fast-Path Double-Checked Locking` — Atomic fast-path initialization and OnceFunc, OnceValue, OnceValues wrappers (Go 1.21+).
- `sync.Pool Architecture (poolLocal & Victim Cache)` — Per-P private slots, shared lock-free deques, poolVictim caches, and GC cleansing cycles.
- `sync.Map Architecture (readOnly vs dirty Map)` — Lockless atomic reads from readOnly map, dirty map fallbacks, and amortized promotions.
- `sync.Cond Condition Variables` — Coordinating goroutines with Wait(), Signal(), Broadcast(), and lost signal hazards.
- `Mutex vs Channel Selection Matrix` — Staff-level architectural decision tree for shared memory vs message passing.
### 4. 📂 `04. Lock-Free & Atomic Concurrency`
- `sync-atomic Primitives (Load, Store, CAS, Swap, Add)` — Hardware atomic primitives providing sequential consistency without mutex locks.
- `Atomic CAS Loop Pattern` — Optimistic concurrency control with Compare-And-Swap spin loops.
- [[atomic.Pointer[T] & atomic.Value Type Safety]] — Type-safe atomic pointers and atomic value containers in Go 1.19+.
- `Lock-Free Stack (Treiber Stack)` — Implementing a concurrent lock-free LIFO stack using atomic pointer CAS.
- `Lock-Free Queue (Michael-Scott Queue)` — Implementing a concurrent lock-free FIFO queue with head and tail pointers.
- `Lock-Free Ring Buffer` — High-throughput single-producer single-consumer (SPSC) and MPMC lock-free buffers.
- `Lock-Free vs Mutex Performance Benchmarks` — Measuring throughput, CPU cache pressure, and scalability tradeoffs.
### 5. 📂 [[Context Trees & Request Cancellation|05. Context Trees & Request Cancellation]]
- `context.Background() vs context.TODO()` — Root context initialization and placeholder context semantics.
- `cancelCtx & Cancellation Tree Propagation` — Parent-to-child cancellation propagation and child detach mechanics.
- `timerCtx & Deadline Scheduling (time.AfterFunc)` — Scheduling automatic cancellations via system timers and deadline expiration.
- `valueCtx & Key-Value Immutability` — Thread-safe immutable request-scoped value lookup and O(N) search depth caveats.
- `context.WithoutCancel (Go 1.21+)` — Detaching cancellation from parent context while preserving request values.
- `context.AfterFunc (Go 1.21+)` — Registering asynchronous callbacks executed when context is canceled.
- `Context Memory Leaks & Resource Hygiene` — Preventing goroutine and memory leaks by always deferring cancel() calls.
- `Context Design Rules` — Passing context as first parameter, avoiding context in structs, keeping values clean.
### 6. 📂 [[Advanced Concurrency Patterns|06. Advanced Concurrency Patterns]]
- `Worker Pools (Static & Elastic Scaling)` — Bounded concurrency worker pools with dynamic auto-scaling worker counts.
- `Pipelines & Stream Processing Stages` — Connecting multi-stage data processing pipelines through channel buffers.
- `Fan-Out & Fan-In Multiplexing` — Splitting heavy tasks across worker pools and aggregating results into a single stream.
- `Or-Done Channel Pattern` — Cleanly terminating channel reads when context is canceled without deadlocks.
- `Tee Channel & Bridge Channel Patterns` — Duplicating streams into multiple channels and unnesting channels of channels.
- `Singleflight Request Coalescing (x-sync-singleflight)` — Deduplicating concurrent identical queries into a single in-flight execution.
- `ErrGroup Concurrency & Context Binding (x-sync-errgroup)` — Executing concurrent subtasks with error capture and group cancellation.
- `Weighted Semaphore Resource Limiting (x-sync-semaphore)` — Restricting access to bounded physical resources using weighted semaphores.
- `Token Bucket Rate Limiting (x-time-rate)` — Burst-capable rate limiting using token bucket algorithms.
- `Leaky Bucket Rate Limiting` — Smooth constant-rate request processing with leaky bucket queues.
- `Debounce and Throttle Patterns` — Suppressing rapid event bursts and enforcing minimum execution intervals.
- `Graceful Shutdown Coordinator` — Coordinating multi-service teardown with POSIX signals and context timeouts.
### 7. 📂 [[GMP Scheduler & Runtime Internals|07. GMP Scheduler & Runtime Internals]]
- `GMP Model (G, M, P, Sched Structs)` — Goroutines (G), OS Threads (M), Logical Processors (P), and global scheduler state.
- `Runqueue Architecture (Local vs Global)` — 256-element lock-free local runqueue per P and mutex-guarded global runqueue.
- `Work Stealing Algorithm` — Checking local queue, 1/61 global check, netpoller check, and stealing half from random P.
- `Netpoller (epoll, kqueue, IOCP) Integration` — Asynchronous non-blocking I/O event loop integrated directly with scheduler parking.
- `Syscall Handling & M Handoff` — entersyscall, exitsyscall, detaching P from blocked M, and waking parked threads.
- `Sysmon Daemon Thread` — Background monitoring thread: forcing periodic GC, retaking stuck Ps, and preemption.
- `Signal-Based Async Preemption (SIGURG)` — Non-cooperative async preemption of tight compute loops via OS signals.
- `Thread Parking (notesleep & notewakeup)` — Low-level OS thread sleeping and futex waking mechanisms in the runtime.
- `GOMAXPROCS & Container CFS Quota Throttling` — Kubernetes CPU limits, CFS period/quota calculation, and automaxprocs.
### 8. 📂 `08. Go Memory Model & Hardware Concurrency`
- `Go Memory Model Specification` — Formal rules defining when a write to a variable by one goroutine is visible to another.
- `Happens-Before Relationship Rules` — Establishing strict happens-before edges via channels, mutexes, sync.Once, and goroutines.
- `Instruction Reordering (Compiler & CPU Out-of-Order)` — How compilers and CPU out-of-order execution reorder memory instructions.
- `Memory Barriers & CPU Store Buffers` — Hardware memory fences (MFENCE, SFENCE, LFENCE) and store buffer flushing.
- `CPU Cache Hierarchy & Cache Lines (64-byte)` — L1/L2/L3 CPU caches, 64-byte cache line granularity, and latency hierarchy.
- `Cache Coherency Protocols (MESI & MOESI)` — Modified, Exclusive, Shared, Invalid cache line state transitions across CPU cores.
- `False Sharing & Cache Line Invalidation` — Contention between independent variables on the same 64-byte cache line and padding fixes.
- `Data Race vs Race Condition Deep Dive` — Distinguishing undefined behavior data races from high-level logical race conditions.
### 9. 📂 `09. Distributed Concurrency & Resilience`
- `Circuit Breaker Pattern in Go` — State machine (Closed, Open, Half-Open) protecting downstream dependencies from cascading failures.
- `Distributed Rate Limiting (Redis Sliding Window & Token Bucket)` — Multi-instance cluster rate limiting using Redis and Lua script token buckets.
- `Distributed Lock (Redlock Algorithm & Etcd Lease Fencing)` — Safe distributed mutual exclusion with TTL lease renewal and fencing tokens.
- `Idempotency Keys & Deduplication` — Guaranteeing exactly-once processing of mutating requests using distributed caches.
- `Transactional Outbox Pattern with PostgreSQL & Kafka` — Atomic database write and message event publishing via outbox transaction tables.
- `Transactional Inbox Pattern for Consumer Idempotency` — Deduplicating incoming message delivery using persistent message ID tracking.
- `Saga Pattern in Distributed Concurrency` — Managing distributed transactions across microservices with compensating rollbacks.
- `Bulkhead Isolation Pattern` — Isolating critical resource pools to prevent complete application exhaustion.
- `Distributed Consensus & Raft in Go (etcd-raft & hashicorp-raft)` — Leader election, log replication, and distributed state machines in Go.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`

---

## 🗂️ Topics

- [[Advanced Concurrency Patterns]]
- [[Channel Architecture]]
- [[Channel Operations]]
- [[Context Trees & Request Cancellation]]
- [[Distributed Concurrency]]
- [[GMP Scheduler & Runtime Internals]]
- [[Go Memory Model]]
- [[Goroutine Memory Lifecycle]]
- [[Goroutines]]
- [[Hardware Concurrency]]
- [[Resilience]]
