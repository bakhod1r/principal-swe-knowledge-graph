---
title: Concurrency & Synchronization
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Golang]]"
---

# ⚡ Concurrency & Synchronization

Goroutines, channels (hchan), sync primitives, context cancellation trees, the GMP scheduler, lock-free patterns, and distributed resilience.

```text
Concurrency & Synchronization
│
├── [[Goroutines & Fundamentals|01. Goroutines & Fundamentals]]
│   ├── [[CSP Concurrency Model]]
│   ├── [[Goroutine Mechanics]]
│   ├── [[Goroutines vs OS Threads]]
│   ├── [[Stack Growth & Segmented Stacks]]
│   ├── [[morestack Stack Split & Pointer Reallocation]]
│   ├── [[Goroutine Lifecycle & States]]
│   └── [[Goroutine Leaks & Diagnostics]]
├── [[Channel Architecture|02. Channel Architecture]]
│   ├── [[Unbuffered Channels]]
│   ├── [[Buffered Channels]]
│   ├── [[Channel States & Behaviors]]
│   ├── [[Channel Internals (hchan struct)]]
│   ├── [[Channel Send and Receive Flow]]
│   ├── [[select Multiplexing]]
│   └── [[Closing Channels Rules]]
├── [[Sync & Context Primitives|03. Sync & Context Primitives]]
│   ├── [[sync.Mutex (Normal vs Starvation)]]
│   ├── [[sync.RWMutex]]
│   ├── [[sync.WaitGroup]]
│   ├── [[sync.Once & sync.OnceFunc]]
│   ├── [[sync.Pool]]
│   ├── [[sync.Map]]
│   ├── [[sync.Cond]]
│   ├── [[context.Context Tree]]
│   ├── [[x-sync-errgroup]]
│   ├── [[x-sync-singleflight]]
│   ├── [[x-sync-semaphore]]
│   ├── [[sync-atomic Primitives (Load, Store, CAS, Swap, Add, Value)]]
│   └── [[Mutex vs Channel Selection]]
├── [[Concurrency Patterns|04. Concurrency Patterns]]
│   ├── [[Worker Pools]]
│   ├── [[Pipelines & Stream Processing]]
│   ├── [[Fan-In and Fan-Out]]
│   ├── [[Cancellation & Graceful Shutdown]]
│   ├── [[Rate Limiting & Token Bucket]]
│   ├── [[Debounce and Throttle]]
│   ├── [[Heartbeats & Supervisors]]
│   ├── [[Backpressure & Load Shedding]]
│   ├── [[Deadlock, Livelock & Starvation]]
│   └── [[Concurrency Anti-Patterns]]
├── [[Runtime Scheduler & Internals|05. Runtime Scheduler & Internals]]
│   ├── [[GMP Model (G, M, P)]]
│   ├── [[Work Stealing Algorithm]]
│   ├── [[Sysmon Background Daemon]]
│   ├── [[Signal-Based Async Preemption (SIGURG)]]
│   ├── [[Netpoller]]
│   ├── [[Syscall Handling & M Handoff]]
│   ├── [[Lock-Free Programming & CAS]]
│   └── [[GOMAXPROCS Tuning]]
└── [[Distributed Concurrency & Resilience|06. Distributed Concurrency & Resilience]]
│   ├── [[Circuit Breaker Pattern in Go]]
│   ├── [[Distributed Rate Limiting (Redis & Token Bucket)]]
│   ├── [[Idempotency Keys & Deduplication]]
│   ├── [[Distributed Lock (Redlock & Etcd)]]
│   ├── [[Bulkhead Pattern]]
│   └── [[Outbox Pattern in Go]]
```

---

## 🗂️ Core Categories & Topics

### 1. 📂 [[Goroutines & Fundamentals|01. Goroutines & Fundamentals]]
- [[CSP Concurrency Model]] — Communicating Sequential Processes: Do not communicate by sharing memory; share memory by communicating.
- [[Goroutine Mechanics]] — Spawning concurrent execution threads with go keyword, 2KB initial stack allocation.
- [[Goroutines vs OS Threads]] — Memory footprint, creation cost, context switch overhead (user vs kernel).
- [[Stack Growth & Segmented Stacks]] — Contiguous stack growth (2KB to 1GB) and stack copying mechanics.
- [[morestack Stack Split & Pointer Reallocation]] — Stack frame boundary checks and pointer adjustment during stack copy.
- [[Goroutine Lifecycle & States]] — _Gidle, _Grunnable, _Grunning, _Gwaiting, _Gsyscall, _Gdead states.
- [[Goroutine Leaks & Diagnostics]] — Identifying blocked goroutines, leaked channels, and diagnostic tools (pprof, runtime.NumGoroutine).
### 2. 📂 [[Channel Architecture|02. Channel Architecture]]
- [[Unbuffered Channels]] — Synchronous rendezvous signaling with direct stack-to-stack copy optimization.
- [[Buffered Channels]] — Asynchronous FIFO ring buffer queues with bounded capacity.
- [[Channel States & Behaviors]] — Send, receive, and close semantics on nil, open, and closed channels.
- [[Channel Internals (hchan struct)]] — hchan fields: ring buffer, lock, sendq/recvq sudog wait queues.
- [[Channel Send and Receive Flow]] — Dissecting chansend and chanrecv runtime step-by-step mechanics.
- [[select Multiplexing]] — Non-blocking select with default, pseudo-random case evaluation, and selectgo() implementation.
- [[Closing Channels Rules]] — Only sender should close, closing closed panic, signaling with close(ch).
### 3. 📂 [[Sync & Context Primitives|03. Sync & Context Primitives]]
- [[sync.Mutex (Normal vs Starvation)]] — Bimodal mutex algorithm: high throughput spin vs fair FIFO handoff.
- [[sync.RWMutex]] — Reader-writer lock with writer starvation prevention.
- [[sync.WaitGroup]] — Atomic counter synchronization for coordinating goroutine completion.
- [[sync.Once & sync.OnceFunc]] — Atomic fast-path initialization and Go 1.21+ OnceFunc/OnceValue wrappers.
- [[sync.Pool]] — Lock-free per-P cache for allocating and reusing short-lived temporary objects.
- [[sync.Map]] — Concurrent map optimized for append-only keys and disjoint key reads.
- [[sync.Cond]] — Condition variables for broadcasting signals to waiting goroutines.
- [[context.Context Tree]] — Cancellation propagation, deadlines, timeouts, and request-scoped values.
- [[x-sync-errgroup]] — Managing concurrent subtasks with error propagation and context cancellation.
- [[x-sync-singleflight]] — Suppressing duplicate in-flight function calls (request coalescing).
- [[x-sync-semaphore]] — Weighted semaphore for limiting concurrent resource access.
- [[sync-atomic Primitives (Load, Store, CAS, Swap, Add, Value)]] — Lockless atomic operations and sequential consistency.
- [[Mutex vs Channel Selection]] — When to use shared memory synchronization vs message passing.
### 4. 📂 [[Concurrency Patterns|04. Concurrency Patterns]]
- [[Worker Pools]] — Bounded concurrency worker pool patterns for throughput and resource control.
- [[Pipelines & Stream Processing]] — Connecting multi-stage concurrent processing steps through channels.
- [[Fan-In and Fan-Out]] — Distributing tasks across multiple workers and multiplexing results into a single channel.
- [[Cancellation & Graceful Shutdown]] — Coordinating graceful process shutdown across long-running background workers.
- [[Rate Limiting & Token Bucket]] — Time-based rate limiting using time.Ticker and x/time/rate token buckets.
- [[Debounce and Throttle]] — Limiting event processing frequency in high-throughput event streams.
- [[Heartbeats & Supervisors]] — Liveness monitoring, health check heartbeats, and worker restart loops.
- [[Backpressure & Load Shedding]] — Handling overload scenarios with bounded buffers and dropped requests.
- [[Deadlock, Livelock & Starvation]] — Detecting and preventing synchronization hazards in concurrent Go programs.
- [[Concurrency Anti-Patterns]] — Unbounded goroutines, variable capture bugs, blocking sends on unbuffered channels.
### 5. 📂 [[Runtime Scheduler & Internals|05. Runtime Scheduler & Internals]]
- [[GMP Model (G, M, P)]] — Goroutines (G), OS Threads (M), Logical Processors (P), and scheduling queues.
- [[Work Stealing Algorithm]] — O(1) local run queue checking, global run queue starvation check, and random P stealing.
- [[Sysmon Background Daemon]] — Monitoring blocked syscalls, forcing periodic GC, and triggering preemption.
- [[Signal-Based Async Preemption (SIGURG)]] — Non-cooperative async signal preemption of tight loops without function calls.
- [[Netpoller]] — Non-blocking I/O event multiplexing (epoll, kqueue, IOCP) integrated with GMP scheduler.
- [[Syscall Handling & M Handoff]] — Entering and exiting syscalls, detaching P from M, and thread parking.
- [[Lock-Free Programming & CAS]] — sync/atomic primitives, Compare-And-Swap algorithms, and memory fences.
- [[GOMAXPROCS Tuning]] — Tuning CPU core allocation and container CPU quota limits (automaxprocs).
### 6. 📂 [[Distributed Concurrency & Resilience|06. Distributed Concurrency & Resilience]]
- [[Circuit Breaker Pattern in Go]] — State machine (Closed, Open, Half-Open) protecting downstream services from cascading failure.
- [[Distributed Rate Limiting (Redis & Token Bucket)]] — Multi-instance rate limiting with Redis sliding window and token buckets.
- [[Idempotency Keys & Deduplication]] — Preventing duplicate execution of concurrent mutation requests.
- [[Distributed Lock (Redlock & Etcd)]] — Safe distributed mutual exclusion with TTL lease renewal and fencing tokens.
- [[Bulkhead Pattern]] — Isolating critical resource pools to prevent total system saturation.
- [[Outbox Pattern in Go]] — Transactional database outbox pattern ensuring atomic database write and message publish.

---

## 🔗 Navigation
- ⬆️ Parent: [[Golang]]
- 💻 Base: `Programming`
- 🎓 Root: [[Principal SWE]]
