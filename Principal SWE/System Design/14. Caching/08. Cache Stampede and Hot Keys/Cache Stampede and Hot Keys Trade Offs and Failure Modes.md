---
title: "Cache Stampede and Hot Keys Trade Offs and Failure Modes"
tags:
  - system-design
  - architecture
  - distributed-systems
  - caching
  - principal-swe
parent: "[[Cache Stampede and Hot Keys]]"
---

# Cache Stampede and Hot Keys Trade Offs and Failure Modes

## 1. Definition
**Cache Stampede and Hot Keys Trade Offs and Failure Modes Trade-Offs & Failure Modes** evaluates the architectural compromises, CAP/PACELC theorem implications, split-brain scenarios, and resilience mechanisms for **Cache Stampede and Hot Keys Trade Offs and Failure Modes**.
Designing resilient distributed systems requires choosing explicit trade-offs between consistency, availability, latency, and hardware cost under Byzantine and crash-recovery fault models.

---

## 2. Mental Model
```text
Resilience & Chaos Engineering Invariants:
[ Healthy System ] ─── (Network Partition / Node Crash) ───> [ Degraded Mode ]
                                                                     │
         ┌───────────────────────────────────────────────────────────┘
         ▼
 [ Active Mitigations ]
   ├── Circuit Breakers Open (Fast Failover)
   ├── Rate Limiting / Shed Load (Prioritize Core RPCs)
   └── Fallback to Stale Cache / Graceful Degradation
```

---

## 3. Usage
```go
// Exponential Backoff with Full Jitter Pattern
package main

import (
    "math/rand"
    "time"
)

func BackoffWithJitter(attempt int, base, max time.Duration) time.Duration {
    if attempt < 0 {
        return base
    }
    // Exponential sleep: base * 2^attempt
    temp := base << attempt
    if temp > max || temp <= 0 {
        temp = max
    }
    // Full jitter: Uniform_Random(0, temp) avoids thundering herd synchronized retries
    return time.Duration(rand.Int63n(int64(temp)))
}
```

---

## 4. Gotchas
- **Retry Storms & Cascading Failures:** Retrying failed requests with static intervals synchronizes client traffic, driving already struggling downstream dependencies into total collapse. Always employ **Full Jitter** and dead-letter queues (DLQs).
- **Split-Brain Inconsistency:** When network partitions split a cluster, allowing both partitions to accept writes creates irreconcilable split-brain data divergence. Always require a strict majority quorum ($Q > N/2$) for leader elections and writes.

---

## 🔗 References
- ⬆️ Parent: [[Cache Stampede and Hot Keys]]
- 📚 Module: [[Caching]]

