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
│   ├── [[atomic.Pointer and atomic.Value Type Safety]]
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

### 📂 [[Goroutines|01. Goroutines]]

### 📂 [[Goroutine Memory Lifecycle|09. Goroutine Memory Lifecycle]]

### 📂 [[Channel Architecture|02. Channel Architecture]]

### 3. 📂 [[Synchronization Primitives (sync)|03. Synchronization Primitives (sync)]]
- [[sync.Mutex (Normal vs Starvation Mode)]] — Bimodal mutex algorithm: high-throughput CPU spin vs fair FIFO handoff at 1ms latency threshold.
- [[sync.RWMutex (Reader Starvation Prevention)]] — Reader-writer lock with atomic reader counters and writer priority signaling.
- [[sync.WaitGroup State Bitpacking]] — 64-bit/128-bit atomic state bitpacking combining task counter and waiter counter.
- [[sync.Once & Fast-Path Double-Checked Locking]] — Atomic fast-path initialization and OnceFunc, OnceValue, OnceValues wrappers (Go 1.21+).
- [[sync.Pool Architecture (poolLocal & Victim Cache)]] — Per-P private slots, shared lock-free deques, poolVictim caches, and GC cleansing cycles.
- [[sync.Map Architecture (readOnly vs dirty Map)]] — Lockless atomic reads from readOnly map, dirty map fallbacks, and amortized promotions.
- [[sync.Cond Condition Variables]] — Coordinating goroutines with Wait(), Signal(), Broadcast(), and lost signal hazards.
- [[Mutex vs Channel Selection Matrix]] — Staff-level architectural decision tree for shared memory vs message passing.
### 📂 [[Atomic Operations|21. Go Synchronization / 03. Atomic Operations]]
- [[sync-atomic Primitives (Load, Store, CAS, Swap, Add)]] — Hardware atomic primitives providing sequential consistency without mutex locks.
- [[Atomic CAS Loop Pattern]] — Optimistic concurrency control with Compare-And-Swap spin loops.
- [[atomic.Pointer and atomic.Value Type Safety]] — Type-safe atomic pointers and atomic value containers in Go 1.19+.

### 📂 [[Lock-Free Concurrency|21. Go Synchronization / 02. Lock-Free Concurrency]]
- [[Lock-Free Stack (Treiber Stack)]] — Implementing a concurrent lock-free LIFO stack using atomic pointer CAS.
- [[Lock-Free Queue (Michael-Scott Queue)]] — Implementing a concurrent lock-free FIFO queue with head and tail pointers.
- [[Lock-Free Ring Buffer]] — High-throughput single-producer single-consumer (SPSC) and MPMC lock-free buffers.
- [[Lock-Free vs Mutex Performance Benchmarks]] — Measuring throughput, CPU cache pressure, and scalability tradeoffs.
### 📂 [[Context Trees & Request Cancellation|03. Context Trees & Request Cancellation]]
- [[context.Background() vs context.TODO()]] — Root context initialization and placeholder context semantics.
- [[cancelCtx & Cancellation Tree Propagation]] — Parent-to-child cancellation propagation and child detach mechanics.
- [[timerCtx & Deadline Scheduling (time.AfterFunc)]] — Scheduling automatic cancellations via system timers and deadline expiration.
- [[valueCtx & Key-Value Immutability]] — Thread-safe immutable request-scoped value lookup and O(N) search depth caveats.
- [[context.WithoutCancel (Go 1.21+)]] — Detaching cancellation from parent context while preserving request values.
- [[context.AfterFunc (Go 1.21+)]] — Registering asynchronous callbacks executed when context is canceled.
- [[Context Memory Leaks & Resource Hygiene]] — Preventing goroutine and memory leaks by always deferring cancel() calls.
- [[Context Design Rules]] — Passing context as first parameter, avoiding context in structs, keeping values clean.
### 📂 [[Advanced Concurrency Patterns|04. Advanced Concurrency Patterns]]
- [[Worker Pools (Static & Elastic Scaling)]] — Bounded concurrency worker pools with dynamic auto-scaling worker counts.
- [[Pipelines & Stream Processing Stages]] — Connecting multi-stage data processing pipelines through channel buffers.
- [[Fan-Out & Fan-In Multiplexing]] — Splitting heavy tasks across worker pools and aggregating results into a single stream.
- [[Or-Done Channel Pattern]] — Cleanly terminating channel reads when context is canceled without deadlocks.
- [[Tee Channel & Bridge Channel Patterns]] — Duplicating streams into multiple channels and unnesting channels of channels.
- [[Singleflight Request Coalescing (x-sync-singleflight)]] — Deduplicating concurrent identical queries into a single in-flight execution.
- [[ErrGroup Concurrency & Context Binding (x-sync-errgroup)]] — Executing concurrent subtasks with error capture and group cancellation.
- [[Weighted Semaphore Resource Limiting (x-sync-semaphore)]] — Restricting access to bounded physical resources using weighted semaphores.
- [[Token Bucket Rate Limiting (x-time-rate)]] — Burst-capable rate limiting using token bucket algorithms.
- [[Leaky Bucket Rate Limiting]] — Smooth constant-rate request processing with leaky bucket queues.
- [[Debounce and Throttle Patterns]] — Suppressing rapid event bursts and enforcing minimum execution intervals.
- [[Graceful Shutdown Coordinator]] — Coordinating multi-service teardown with POSIX signals and context timeouts.
### 📂 [[GMP Scheduler & Runtime Internals|05. GMP Scheduler & Runtime Internals]]
- [[GMP Model (G, M, P, Sched Structs)]] — Goroutines (G), OS Threads (M), Logical Processors (P), and global scheduler state.
- [[Runqueue Architecture (Local vs Global)]] — 256-element lock-free local runqueue per P and mutex-guarded global runqueue.
- [[Work Stealing Algorithm]] — Checking local queue, 1/61 global check, netpoller check, and stealing half from random P.
- [[Netpoller (epoll, kqueue, IOCP) Integration]] — Asynchronous non-blocking I/O event loop integrated directly with scheduler parking.
- [[Syscall Handling & M Handoff]] — entersyscall, exitsyscall, detaching P from blocked M, and waking parked threads.
- [[Sysmon Daemon Thread]] — Background monitoring thread: forcing periodic GC, retaking stuck Ps, and preemption.
- [[Signal-Based Async Preemption (SIGURG)]] — Non-cooperative async preemption of tight compute loops via OS signals.
- [[Thread Parking (notesleep & notewakeup)]] — Low-level OS thread sleeping and futex waking mechanisms in the runtime.
- [[GOMAXPROCS & Container CFS Quota Throttling]] — Kubernetes CPU limits, CFS period/quota calculation, and automaxprocs.
### 📂 [[Go Memory Model|06. Go Memory Model]]

### 📂 [[Hardware Concurrency|11. Hardware Concurrency]]

### 📂 [[Distributed Concurrency|07. Distributed Concurrency]]

### 📂 [[Resilience|10. Resilience]]

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
