---
title: Concurrency Patterns
tags:
  - golang
  - concurrency
  - principal-swe
parent: "[[Concurrency & Synchronization]]"
---

# Concurrency Patterns

Worker pools, fan-in/fan-out, pipeline processing, rate limiting, and deadlock prevention.

```text
Concurrency Patterns
│
├── [[Worker Pools]]
├── [[Pipelines & Stream Processing]]
├── [[Fan-In and Fan-Out]]
├── [[Cancellation & Graceful Shutdown]]
├── [[Rate Limiting & Token Bucket]]
├── [[Debounce and Throttle]]
├── [[Heartbeats & Supervisors]]
├── [[Backpressure & Load Shedding]]
├── [[Deadlock, Livelock & Starvation]]
└── [[Concurrency Anti-Patterns]]
```

---

## 🗂️ Topics

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

---

## 🔗 References
- ⬆️ Parent: [[Concurrency & Synchronization]]
- 🎓 Root: [[Principal SWE]]
