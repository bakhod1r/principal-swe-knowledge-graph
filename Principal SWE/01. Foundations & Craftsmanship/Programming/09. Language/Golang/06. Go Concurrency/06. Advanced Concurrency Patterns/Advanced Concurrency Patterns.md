---
title: Advanced Concurrency Patterns
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Go Concurrency]]"
---

# Advanced Concurrency Patterns

Worker pools, stream pipelines, fan-in/fan-out, singleflight, errgroup, rate limiters, and graceful shutdown.

```text
Advanced Concurrency Patterns
│
├── [[Worker Pools (Static & Elastic Scaling)]]
├── [[Pipelines & Stream Processing Stages]]
├── [[Fan-Out & Fan-In Multiplexing]]
├── [[Or-Done Channel Pattern]]
├── [[Tee Channel & Bridge Channel Patterns]]
├── [[Singleflight Request Coalescing (x-sync-singleflight)]]
├── [[ErrGroup Concurrency & Context Binding (x-sync-errgroup)]]
├── [[Weighted Semaphore Resource Limiting (x-sync-semaphore)]]
├── [[Token Bucket Rate Limiting (x-time-rate)]]
├── [[Leaky Bucket Rate Limiting]]
├── [[Debounce and Throttle Patterns]]
└── [[Graceful Shutdown Coordinator]]
```

---

## 🗂️ Topics

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

---

## 🔗 References
- ⬆️ Parent: `Concurrency & Synchronization`

